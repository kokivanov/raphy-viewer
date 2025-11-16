<script lang="ts">
	import { _ } from 'svelte-i18n';
	import { tick } from 'svelte';
	import { fade, scale } from 'svelte/transition';
	import { sineIn } from 'svelte/easing';
	import LetsIconsSearchAlt from '~icons/lets-icons/search-alt';
	import LetsIconsCloseRound from '~icons/lets-icons/close-round';
	import { scaleFromSide } from '$lib/transitions/ScaleFromSide';
	import { debounce } from 'lodash';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { searchReq } from '../../stores/Search';
	import { Tooltip } from 'flowbite-svelte';

	let searchrequest = $state('');
	let showSearch = $state(false);

	// svelte-ignore non_reactive_update
	let searchInput: HTMLInputElement;

	export const handleSearchClick = async () => {
		if (!showSearch) {
			showSearch = true;
			await tick();
			searchInput.focus();
		}

		if (!searchrequest.length && page.route.id !== '/explore') {
			goto('/explore', { keepFocus: true });
		} else if (searchrequest.length && page.route.id !== '/search') {
			goto(`/search`, { keepFocus: true });
		}
	};

	export const handleBlur = () => {
		showSearch = !!searchrequest.length || page.route.id === '/search';
	};

	export const handleKeyPress = (e: KeyboardEvent) => {
		if (e.key === 'Enter') {
			e.preventDefault();
			e.stopPropagation();
			searchReq.set(searchrequest);
			if (searchrequest.length > 0) {
				goto(`/search`, { keepFocus: true });
			} else {
				goto(`/explore`, { keepFocus: true });
			}
		}
	};

	export const handleSearch = debounce(() => {
		searchReq.set(searchrequest);
		if (searchrequest.length > 0) {
			goto(`/search`, { keepFocus: true });
		} else {
			goto(`/explore`, { keepFocus: true });
		}
	}, 1000);

	export const handleClearClick = async () => {
		showSearch = true;
		await tick();
		searchInput.focus();
		if (page.route.id === '/search') {
			goto('/explore', { keepFocus: true });
		}
		searchrequest = '';
	};
</script>

<div class="relative">
	<button onclick={handleClearClick} class="absolute right-0">
		<LetsIconsCloseRound
			class={`size-10 cursor-pointer p-2 transition-[opacity] delay-150 duration-200 ${showSearch ? 'opacity-100' : 'opacity-0'}`}
		></LetsIconsCloseRound>
		<Tooltip>$_('navbar.search.clear_tooltip')</Tooltip>
	</button>
	<button onclick={handleSearchClick} class={`absolute left-0 cursor-pointer rounded-3xl `}>
		<LetsIconsSearchAlt
			color={`${showSearch ? 'black' : 'white'}`}
			class={`size-10 p-2 transition `}
			role="link"
		></LetsIconsSearchAlt>
		<Tooltip
			>{showSearch
				? $_('navbar.search.extended_tooltip')
				: $_('navbar.search.collapsed_tooltip')}</Tooltip
		>
	</button>
	<input
		class={`h-10 rounded-3xl transition-[width,padding,background] duration-300 ease-in-out ${showSearch ? 'w-72 bg-amber-50 px-10' : 'w-10 bg-none px-5'}`}
		bind:this={searchInput}
		bind:value={searchrequest}
		placeholder={$_('navbar.search.placeholder')}
		aria-label="search"
		onblur={handleBlur}
		oninput={handleSearch}
		onkeypress={handleKeyPress}
		transition:scaleFromSide={{ duration: 200, axis: 'x' }}
	/>
</div>
