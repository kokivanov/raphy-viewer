<script lang="ts">
	import { browser } from '$app/environment';
	import Tile from '$lib/components/content/Tile.svelte';
	import { searchReq } from '$lib/stores/Search';
	import { QueryAllContent } from '$lib/wailsjs/go/api/ApiService';
	import { GetFavoritesJoin } from '$lib/wailsjs/go/data/DataService';
	import type { core } from '$lib/wailsjs/go/models';
	import { Heading, Spinner } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';

	let res: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let favList: number[];

	async function getSearchResults(q: string) {
		const rawRes = await QueryAllContent({
			perPage: 50,
			page: 1,
			search: q
		});

		const favs = await GetFavoritesJoin(rawRes.map((v) => v.idMal));

		console.log(favs);

		favList = (favs ?? []).map((v) => v.ID);

		return rawRes.map((v) => {
			return {
				...v,
				isFav: favList.includes(v.idMal),
				convertValues: (a: any, classs: any, asMap: boolean = false) => {}
			};
		});
	}

	onMount(async () => {
		if (browser) {
			searchReq.subscribe(async (q) => {
				res = getSearchResults(q);
			});
		}
	});
</script>

<p>SEARCH PAGE</p>

<!-- <div class="grid justify-center gap-5 [grid-template-columns:repeat(auto-fit,208px)]"> -->
{#await res}
	<div class="flex h-full w-full flex-row items-center justify-center">
		<Spinner size="12"></Spinner>
		<Heading tag="h3">{$_('general.loading')}</Heading>
	</div>
{:then medias}
	<div class="flex w-full flex-wrap justify-center gap-5">
		{#if medias}
			{#each medias as media}
				<Tile content={media}></Tile>
			{/each}
		{/if}
	</div>
{/await}
