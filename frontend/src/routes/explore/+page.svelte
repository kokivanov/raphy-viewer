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

<style>
</style>
