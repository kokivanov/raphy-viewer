<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { _, locale } from 'svelte-i18n';
	import NavBar from '$lib/components/navbar/AppHeader.svelte';
	import { LocalSettings } from '$lib/stores/LocalSettings';
	import { AppConnection, AppStateEnum } from '$lib/stores/AppEvents';
	import { EventsOn } from '$lib/wailsjs/runtime/runtime';
	import { Banner, Spinner } from 'flowbite-svelte';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { defer, delay, some } from 'lodash';
	import { slide } from 'svelte/transition';
	import { GetConnectionState } from '$lib/wailsjs/go/main/App';

	let { children } = $props();

	let showBanner = $state(get(AppConnection) === AppStateEnum.OFFLINE);

	let isConnecting = $state(false);

	onMount(() => {
		if (browser) {
			LocalSettings.subscribe((settings) => {
				console.log(settings.locale);

				if (settings.locale) {
					locale.set(settings.locale);
				}
			});

			AppConnection.subscribe(console.log);

			EventsOn('app.connection', (state) => {
				AppConnection.set(Number(state));
				showBanner = state === '0';
				console.log(showBanner);
			});
		}
	});

	function retry() {
		isConnecting = true;
		GetConnectionState();
		delay(() => {
			isConnecting = false;
		}, 1000);
	}
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<NavBar></NavBar>
<main class="bg-background overflow-y-scroll">
	{@render children?.()}
</main>
<Banner
	bind:open={showBanner}
	class="h-2 bg-red-500 text-sm text-white"
	transition={slide}
	dismissable={false}
	type="bottom"
	>{$_('general.offline')}
	{#if isConnecting}
		<Spinner size="4"></Spinner> {$_('general.reconnecing')}
	{:else}
		<button onclick={retry} class="cursor-pointer underline underline-offset-1"
			>{$_('general.retry')}</button
		>
	{/if}
</Banner>

<style>
	main {
		height: calc(100vh - var(--spacing) * 16);
		width: 100vw;
		overflow-y: auto;
		overflow-x: hidden;
	}
</style>
