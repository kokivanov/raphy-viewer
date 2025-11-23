<script lang="ts">
	import { browser } from '$app/environment';
	import { page } from '$app/state';
	import { QueryByID } from '$lib/wailsjs/go/api/ApiService';
	import type { core } from '$lib/wailsjs/go/models';
	import { Heading, ImagePlaceholder, Img, P, Spinner } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';
	import { blur } from 'svelte/transition';
	import nobanner from '$lib/assets/images/no-banner.jpg';

	let id = page.params.id;
	let media: Promise<core.FullMedia> | null;

	onMount(() => {
		if (browser) {
			const mediaId = Number(id);

			if (!isNaN(mediaId)) {
				media = QueryByID(mediaId);
			}
		}
	});
</script>

{#if media}
	{#await media}
		<section out:blur={{ duration: 100 }} class="relative">
			<article class="my-12 flex w-full items-center justify-center">
				<Spinner class="mr-12 size-16"></Spinner>
				<P>{$_('general.loading')}</P>
			</article>
			<ImagePlaceholder class="h-full w-full" />
		</section>
	{:then mediaData}
		<section class="relative pr-5" in:blur={{ duration: 300, delay: 100 }}>
			<div class="opacity-gradient fixed max-h-56 w-full overflow-clip">
				{#if mediaData.bannerImage}
					<Img
						class="h-auto w-full"
						src={`/local/img?id=${mediaData.idMal}&u=${mediaData.bannerImage}`}
					/>
				{:else}
					<Img class="max-h-[220px] w-full" src={nobanner} />
				{/if}
			</div>
			<article class="relative top-60 flex">
				<Img
					class="relative -top-32 m-5 h-80 border-4 border-amber-500"
					src={`/local/img?id=${mediaData.idMal}&u=${mediaData.coverImage.extraLarge ?? mediaData.coverImage.large}`}
				/>
				<div class="mt-2">
					<Heading tag="h1">
						{mediaData.title.userPreferred ?? mediaData.title.english ?? mediaData.title.native}
					</Heading>
					{#if mediaData.title.native && mediaData.title.native !== mediaData.title.userPreferred && (mediaData.title.userPreferred || mediaData.title.english)}
						<Heading tag="h2" class="text-2xl text-stone-500">
							{mediaData.title.native}
						</Heading>
					{/if}
					<P class="mt-4">
						{@html mediaData.description}
					</P>
				</div>
			</article>
		</section>
	{/await}
{/if}

<style>
	.opacity-gradient {
		mask-mode: alpha;
		mask-image: linear-gradient(to top, rgba(0, 0, 0, 0), rgba(0, 0, 0, 1));
	}
</style>
