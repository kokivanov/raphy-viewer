<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { locale } from 'svelte-i18n';
	import NavBar from '$lib/components/navbar/AppHeader.svelte';
	import { LocalSettings } from '$lib/stores/LocalSettings';

	let { children } = $props();

	LocalSettings.subscribe((settings) => {
		console.log(settings.locale);

		if (settings.locale) {
			locale.set(settings.locale);
		}
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<NavBar></NavBar>
<main class="bg-background overflow-y-scroll">
	{@render children?.()}
</main>

<style>
	main {
		height: calc(100vh - var(--spacing) * 20.5);
		width: 100vw;
		overflow-y: auto;
		overflow-x: hidden;
	}
</style>
