<script>
	import { Avatar, Img, Tooltip } from 'flowbite-svelte';
	import SearchBar from './SearchBar.svelte';
	import favicon from '$lib/assets/favicon.svg';
	import {
		EventsOn,
		Quit,
		WindowFullscreen,
		WindowIsFullscreen,
		WindowIsMaximised,
		WindowMaximise,
		WindowMinimise,
		WindowToggleMaximise,
		WindowUnfullscreen,
		WindowUnmaximise
	} from '$lib/wailsjs/runtime/runtime';
	import { GetPlatform } from '$lib/wailsjs/go/main/App';

	import QlementineIconsWindowsUnmaximize16 from '~icons/qlementine-icons/windows-unmaximize-16';
	import QlementineIconsWindowsMinimize16 from '~icons/qlementine-icons/windows-minimize-16';
	import QlementineIconsWindowsMaximize16 from '~icons/qlementine-icons/windows-maximize-16';
	import QlementineIconsWindowsClose16 from '~icons/qlementine-icons/windows-close-16';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { _ } from 'svelte-i18n';

	let platform = '';
	let isFullScreen = $state(false);

	onMount(async () => {
		platform = await GetPlatform();
		if (platform === 'darwin') {
			isFullScreen = await WindowIsFullscreen();
		} else {
			isFullScreen = await WindowIsMaximised();
		}
	});

	const handleDoubleClick = async () => {
		await WindowToggleMaximise();
		isFullScreen = !isFullScreen && platform !== 'darwin';
	};

	const handleMaximiseClick = async () => {
		if (platform === 'darwin') {
			if (isFullScreen) {
				await WindowUnfullscreen();
				isFullScreen = false;
			} else {
				await WindowFullscreen();
				isFullScreen = true;
			}
		} else {
			if (isFullScreen) {
				await WindowMaximise();
				isFullScreen = false;
			} else {
				await WindowUnmaximise();
				isFullScreen = true;
			}
		}
	};
</script>

<header class="sticky top-0 z-10 flex flex-row bg-amber-600 px-4 py-2">
	<section
		ondblclick={handleDoubleClick}
		class="flex-12 draggable flex flex-row items-center justify-between"
		role="navigation"
	>
		<!-- svelte-ignore component_name_lowercase -->
		<button
			role="link"
			onclick={() => goto('/')}
			class="not-draggable m-0 w-8 shrink-0 p-0"
			aria-details="Home page"
		>
			<Img src={favicon}></Img>
		</button>
		<Tooltip>{$_('navbar.home')}</Tooltip>
		<div class="w-full"></div>
		<div class="not-draggable">
			<SearchBar></SearchBar>
		</div>
		<div class="w-full"></div>
		<!-- svelte-ignore component_name_lowercase -->
		<a class="not-draggable h-8 w-8" aria-details="User profile" href="/user">
			<Avatar class="h-8 w-8"></Avatar>
			<Tooltip>{$_('navbar.profile_tooltip')}</Tooltip>
		</a>
	</section>
	<section class="flex items-center pl-4">
		<!-- svelte-ignore component_name_lowercase -->
		<button class="size-8 hover:bg-amber-500" onclick={WindowMinimise}>
			<QlementineIconsWindowsMinimize16 class="size-8" color="white"
			></QlementineIconsWindowsMinimize16>
			<Tooltip>{$_('navbar.minimise')}</Tooltip>
		</button>
		<!-- svelte-ignore component_name_lowercase -->
		<button class="size-8 hover:bg-amber-500" onclick={handleMaximiseClick}>
			{#if isFullScreen}
				<QlementineIconsWindowsUnmaximize16 class="size-8" color="white"
				></QlementineIconsWindowsUnmaximize16>
				<Tooltip>{$_('navbar.unmaximise')}</Tooltip>
			{:else}
				<QlementineIconsWindowsMaximize16 class="size-8" color="white"
				></QlementineIconsWindowsMaximize16>
				<Tooltip>{$_('navbar.maximise')}</Tooltip>
			{/if}
		</button>
		<!-- svelte-ignore component_name_lowercase -->
		<button class="size-8 hover:bg-red-500" onclick={Quit}>
			<QlementineIconsWindowsClose16 class="size-8" color="white"></QlementineIconsWindowsClose16>
			<Tooltip>{$_('navbar.close')}</Tooltip>
		</button>
	</section>
</header>

<style>
	.draggable {
		--wails-draggable: drag;
	}

	.not-draggable {
		--wails-draggable: no-drag;
	}
</style>
