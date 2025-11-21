<script>
	import {
		Avatar,
		Heading,
		Img,
		Navbar,
		NavBrand,
		NavHamburger,
		NavLi,
		NavUl,
		Tooltip
	} from 'flowbite-svelte';
	import SearchBar from './SearchBar.svelte';
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
	import { page } from '$app/state';
	import Logo from '../../svgs/Logo.svelte';

	let isFullScreen = $state(false);

	const activeUrl = $derived(page.url.pathname);

	onMount(async () => {
		isFullScreen = await WindowIsMaximised();
	});

	const handleDoubleClick = async () => {
		await WindowToggleMaximise();
		isFullScreen = !isFullScreen;
	};

	const handleMaximiseClick = async () => {
		if (isFullScreen) {
			await WindowUnmaximise();
			isFullScreen = false;
		} else {
			await WindowMaximise();
			isFullScreen = true;
		}
	};
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	class="draggable unselectable sticky start-0 top-0 z-20 flex grow-0 flex-row items-center justify-between bg-white shadow-sm"
	ondblclick={handleDoubleClick}
>
	<svelte-css-wrapper class="grow justify-between">
		<Navbar>
			<NavBrand href="/about" class="">
				<Logo></Logo>
				<Heading tag="h5" class="text-accent ml-1">Raphy Viewer</Heading>
			</NavBrand>

			<NavUl {activeUrl} activeClass="nav-active" nonActiveClass="nav-inactive" class="order-1">
				<NavLi href="/">{$_('navbar.home')}</NavLi>
				<NavLi href="/explore">{$_('navbar.explore')}</NavLi>
				<NavLi href="/my">{$_('navbar.my')}</NavLi>
			</NavUl>

			<div class="flex md:order-2">
				<NavHamburger></NavHamburger>
				<Avatar id="avatar-menu"></Avatar>
			</div>
		</Navbar>
	</svelte-css-wrapper>
	<div
		id="btns-window-actions"
		class="not-draggable mx-2 flex h-full flex-row content-center items-center justify-center"
	>
		<button class="action-button" onclick={WindowMinimise}>
			<QlementineIconsWindowsMinimize16></QlementineIconsWindowsMinimize16>
		</button>
		<div class="relative">
			{#if isFullScreen}
				<button class="action-button" onclick={handleMaximiseClick}>
					<QlementineIconsWindowsUnmaximize16></QlementineIconsWindowsUnmaximize16>
				</button>
			{:else}
				<button class="action-button" onclick={handleMaximiseClick}>
					<QlementineIconsWindowsMaximize16></QlementineIconsWindowsMaximize16>
				</button>
			{/if}
		</div>
		<button class="action-button" onclick={Quit}>
			<QlementineIconsWindowsClose16></QlementineIconsWindowsClose16>
		</button>
	</div>
</div>

<style>
	.unselectable {
		-webkit-user-select: none;
		-moz-user-select: none;
		-ms-user-select: none;
		user-select: none;
	}

	.draggable {
		--wails-draggable: drag;
	}

	.not-draggable {
		--wails-draggable: no-drag;
	}

	:global(.nav-active) {
		border-bottom: 2px var(--color-accent) solid;
		border-radius: 0px;
		box-sizing: border-box;
		transition: all;
		transition-duration: 100ms;
		padding-bottom: 0px;
	}

	:global(.nav-inactive) {
		border-bottom: 2px rgba(0, 0, 0, 0) solid;
		border-radius: 0px;
		box-sizing: border-box;
		transition: all;
		transition-duration: 100ms;
		padding-bottom: 0px;
	}
</style>
