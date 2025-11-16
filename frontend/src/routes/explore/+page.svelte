<script lang="ts">
	import { browser } from '$app/environment';
	import { SortTypesEnum } from '$lib';
	import Tile from '$lib/components/content/Tile.svelte';
	import { QueryAllContent } from '$lib/wailsjs/go/api/ApiService';
	import { GetFavoritesJoin } from '$lib/wailsjs/go/data/DataService';
	import type { core } from '$lib/wailsjs/go/models';
	import { Heading, Spinner } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';

	let popular: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let trending: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let bestScore: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let mostLoved: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let latest: Promise<(core.ShortMedia & { isFav: boolean })[]>;

	onMount(async () => {
		if (browser) {
			popular = getMedia(SortTypesEnum.SORT_POPULARITY_DES);
			trending = getMedia(SortTypesEnum.SORT_TRENDING_DESC);
			mostLoved = getMedia(SortTypesEnum.SORT_FAVOURITES_DES);
			latest = getMedia(SortTypesEnum.SORT_START_DATE_DES);
			bestScore = getMedia(SortTypesEnum.SORT_SCORE_DESC);
		}
	});

	async function getMedia(sort: SortTypesEnum) {
		const rawRes = await QueryAllContent({
			perPage: 20,
			page: 1,
			sort: sort,
			statusIn: ['FINISHED', 'RELEASING']
		});

		const favs = await GetFavoritesJoin(rawRes.map((v) => v.idMal));

		console.log(favs);

		const favList = (favs ?? []).map((v) => v.ID);

		return rawRes.map((v) => {
			return {
				...v,
				isFav: favList.includes(v.idMal),
				convertValues: (a: any, classs: any, asMap: boolean = false) => {}
			};
		});
	}
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
					{#if r}
						{#each r as tile}
							<Tile content={tile}></Tile>
						{/each}
					{/if}
				</div>
			{/await}
		</article>
	</section>
{/snippet}

<div>
	{@render cat($_('explore.trending'), trending)}

	{@render cat($_('explore.popular'), popular)}

	{@render cat($_('explore.latest'), latest)}

	{@render cat($_('explore.best_score'), bestScore)}

	{@render cat($_('explore.most_loved'), mostLoved)}
</div>

<style>
</style>
