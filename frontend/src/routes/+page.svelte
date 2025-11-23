<script lang="ts">
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { SortTypesEnum } from '$lib';
	import Tile from '$lib/components/content/Tile.svelte';
	import Alien from '$lib/svgs/Alien.svelte';
	import { getMedia } from '$lib/utils/media';
	import { core } from '$lib/wailsjs/go/models';
	import {
		Button,
		Carousel,
		CarouselIndicators,
		ControlButton,
		Controls,
		GradientButton,
		Heading,
		ImagePlaceholder,
		P,
		Skeleton,
		Spinner,
		VideoPlaceholder
	} from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';

	import OouiNextLtr from '~icons/ooui/next-ltr';
	import OouiPreviousLtr from '~icons/ooui/previous-ltr';

	let watching = new Promise<(core.ShortMedia & { isFav: boolean })[] | null>((resolve) => {
		setTimeout(() => resolve(null), 1000);
	});
	let popular: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let trending: Promise<(core.ShortMedia & { isFav: boolean })[]>;
	let banners: Promise<{ src: string; idMal: number; title: string }[]>;

	onMount(async () => {
		if (browser) {
			popular = getMedia(SortTypesEnum.SORT_POPULARITY_DES);
			trending = getMedia(SortTypesEnum.SORT_TRENDING_DESC);
			banners = new Promise((resolve, reject) => {
				trending
					.catch((err) => {
						reject(err);
						return [];
					})
					.then((medias) => {
						console.log(medias);
						const finalList = medias
							.filter((media) => media.bannerImage)
							.slice(0, 5)
							.map((media) => {
								return {
									src: media.bannerImage,
									idMal: media.idMal,
									title: media.title.userPreferred ?? media.title.english ?? media.title.native
								};
							});
						resolve(finalList);
					});
			});
		}
	});
</script>

{#snippet cat(
	title: string,
	cn: Promise<
		| (core.ShortMedia & {
				isFav: boolean;
		  })[]
		| null
	>
)}
	{#await cn}
		<article class="mt-4 h-72">
			<Heading tag="h3" class="pl-12">{title}</Heading>
			<div class="flex h-full w-full flex-row items-center justify-center">
				<Spinner size="12"></Spinner>
				<Heading tag="h3">{$_('general.loading')}</Heading>
			</div>
		</article>
	{:then r}
		{#if r}
			<section class="mt-4 h-72">
				<Heading tag="h3" class="pl-12">{title}</Heading>
				<article>
					<div class="custom-scroll flex w-dvw gap-5 overflow-x-scroll py-2 pl-20">
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
				</article>
			</section>
		{/if}
	{/await}
{/snippet}

{#if banners}
	{#await banners}
		<div class="relative mt-4 flex h-56 w-full flex-row items-center justify-center">
			<Spinner size="12" class="z-10"></Spinner>
			<Heading tag="h3" class="z-10">{$_('general.loading')}</Heading>
			<ImagePlaceholder imgOnly class="image-placeholder banner-placeholder absolute w-full px-10"
			></ImagePlaceholder>
		</div>
	{:then ban}
		<Carousel images={ban} duration={10000} class="mx-10 mt-4 overflow-hidden rounded-2xl">
			{#snippet slide({ index, Slide })}
				<div class="relative">
					<Slide image={ban[index]}></Slide>
					<div class="absolute h-full w-full bg-black opacity-35"></div>
					<div class="absolute bottom-0 w-full px-16 py-10">
						<Heading tag="h3" class="text-white">{ban[index].title}</Heading>
						<Button class="mt-2 cursor-pointer" onclick={() => goto(`/media/${ban[index].idMal}`)}
							>{$_('home.banner.watch')}</Button
						>
						<Button color="secondary" class="ml-2 mt-2 cursor-pointer"
							>{$_('home.banner.add')}</Button
						>
					</div>
				</div>
			{/snippet}

			<Controls>
				{#snippet children(changeSlide)}
					<button
						class="hover:text-accent absolute start-4 top-1/2 -translate-y-1/2 cursor-pointer p-2 font-bold text-white"
						onclick={() => changeSlide(false)}
						><OouiPreviousLtr></OouiPreviousLtr>
					</button>
					<button
						class="hover:text-accent absolute end-4 top-1/2 -translate-y-1/2 cursor-pointer p-2 font-bold text-white"
						onclick={() => changeSlide(true)}><OouiNextLtr></OouiNextLtr></button
					>
				{/snippet}
			</Controls>
			<CarouselIndicators />
		</Carousel>
	{/await}
{/if}

<div>
	{@render cat($_('explore.watching'), watching)}

	{@render cat($_('explore.trending'), trending)}

	{@render cat($_('explore.popular'), popular)}
</div>

<style>
	:global(.banner-placeholder > div) {
		width: 100% !important;
		margin: 0px;
	}

	div::-webkit-scrollbar {
		display: none;
	}

	div {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
</style>
