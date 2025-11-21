<script lang="ts">
	import { browser } from '$app/environment';
	import { SortTypesEnum } from '$lib';
	import Tile from '$lib/components/content/Tile.svelte';
	import Alien from '$lib/svgs/Alien.svelte';
	import { getMedia } from '$lib/utils/media';
	import { core } from '$lib/wailsjs/go/models';
	import { Button, Heading, P, Spinner } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';

	let watching = new Promise<(core.ShortMedia & { isFav: boolean })[]>((resolve) => {
		resolve([]);
	});
	let popular: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let trending: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	onMount(async () => {
		if (browser) {
			popular = getMedia(SortTypesEnum.SORT_POPULARITY_DES);
			trending = getMedia(SortTypesEnum.SORT_TRENDING_DESC);
		}
	});
</script>

{#snippet cat(
	title: string,
	cn: Promise<
		(core.ShortMedia & {
			isFav: boolean;
		})[]
	>
)}
	<section>
		<Heading tag="h3">{title}</Heading>
		<article class="h-96">
			{#await cn}
				<div class="flex h-full w-full flex-row items-center justify-center">
					<Spinner size="12"></Spinner>
					<Heading tag="h3">{$_('general.loading')}</Heading>
				</div>
			{:then r}
				<div class="custom-scroll flex w-dvw gap-5 overflow-x-scroll p-7">
					{#if r?.length}
						{#each r as tile}
							<Tile content={tile}></Tile>
						{/each}
					{:else}
						<div class="flex h-60 w-full flex-col items-center justify-center">
							<Alien></Alien>
							<P class="font-semibold text-zinc-500">{$_('home.no-media')}</P>
							<Button class="bg-secondary-400 mt-2">{$_('home.explore_more')} +</Button>
						</div>
					{/if}
				</div>
			{/await}
		</article>
	</section>
{/snippet}

<div>
	{@render cat($_('explore.watching'), watching)}

	{@render cat($_('explore.trending'), trending)}

	{@render cat($_('explore.popular'), popular)}
</div>
