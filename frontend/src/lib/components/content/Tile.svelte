<script lang="ts">
	import type { core } from '$lib/wailsjs/go/models';
	import { ImagePlaceholder, Img, P, span } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	import TablerClockFilled from '~icons/tabler/clock-filled';
	import TablerClock from '~icons/tabler/clock';
	import TablerHeartFilled from '~icons/tabler/heart-filled';
	import TablerHeart from '~icons/tabler/heart';
	import { AddFavorite, RemoveFavorite } from '$lib/wailsjs/go/data/DataService';
	import { fade } from 'svelte/transition';
	import { gsap } from 'gsap';

	let isInWatchlist = $state(false);

	let isReady = $state(false);
	let { content }: { content: core.ShortMedia & { isFav: boolean } } = $props();
	let isFav = $state(content.isFav);

	const handleImageLoad = () => {
		isReady = true;
	};

	let status =
		content.status === 'RELEASING'
			? 'general.tile.ongoing'
			: content.status === 'FINISHED'
				? 'general.tile.finished'
				: content.status;
	let statusColor =
		content.status === 'RELEASING' ? 'purple' : content.status === 'FINISHED' ? 'green' : 'red';

	const handleFavClick = async (e: MouseEvent) => {
		e.preventDefault();
		e.stopPropagation();

		gsap.fromTo(
			'#heart',
			{
				scale: 0.7,
				repeat: 0
			},
			{
				scale: 1,
				repeat: 0,
				ease: 'bounce.out'
			}
		);

		if (!isFav) {
			isFav = true;
			isFav = await AddFavorite(content.idMal);
		} else {
			isFav = false;
			isFav = await RemoveFavorite(content.idMal);
		}
	};
</script>

<a
	class="relative h-80 w-52 shrink-0 border-4 border-[var(--color-primary-500)]"
	href={`/media/${content.idMal}`}
	role="button"
	aria-label={$_('general.tile.aria-label', {
		values: {
			to: content.title.userPreferred ?? content.title.english ?? content.title.native
		}
	})}
>
	<div
		class="tile absolute h-full w-full items-center p-3 opacity-0 transition-opacity hover:opacity-100"
	>
		<p
			class="tile-title m-auto h-full w-full content-center overflow-ellipsis text-center font-semibold text-white"
		>
			{content.title.userPreferred ?? content.title.english ?? content.title.native}
		</p>
		<!-- svelte-ignore a11y_consider_explicit_label -->

		<button id="heart" class="absolute right-2 top-1 cursor-pointer" onclick={handleFavClick}>
			{#if isFav}
				<div>
					<TablerHeartFilled class="text-red-400 hover:text-red-400"></TablerHeartFilled>
				</div>
			{:else}
				<div>
					<TablerHeart class="text-white hover:text-red-400"></TablerHeart>
				</div>
			{/if}
		</button>
	</div>

	<div class="absolute bg-{statusColor}-400 rounded-br-md">
		<p class=" px-2 font-semibold text-white">{$_(status)}</p>
	</div>
	{#if !isReady}
		<ImagePlaceholder class="aspect-[4/6] w-1/5 xl:w-[15%]" imgOnly></ImagePlaceholder>
	{/if}
	<Img
		class="aspect-auto h-full w-full object-cover"
		src={`/local/img?id=${content.idMal}&u=${content.coverImage.large ?? content.coverImage.extraLarge ?? content.coverImage.medium}`}
		onload={handleImageLoad}
	></Img>
</a>

<style>
	.tile {
		background: rgba(0, 0, 0, 0.65);
	}

	.tile-title {
		text-overflow: ellipsis;
		overflow: hidden;
	}
</style>
