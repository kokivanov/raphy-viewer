<script>
	import Logo from '$lib/svgs/Logo.svelte';
	import { Card, Footer, FooterCopyright, FooterLink, FooterLinkGroup, P } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
</script>

<section class="flex w-full justify-center p-5">
	<Card class="flex flex-col items-center justify-center p-5">
		<Logo></Logo>
		<P>Raphy Viewer 0.0.1a</P>
	</Card>
</section>

<Footer class="absolute bottom-0 w-full">
	<FooterCopyright href="/" by="Koki" year={2025} />
	<FooterLinkGroup
		class="mt-3 flex flex-wrap items-center text-sm text-gray-500 sm:mt-0 dark:text-gray-400"
	>
		<FooterLink href="/about">{$_('footer.about')}</FooterLink>
		<FooterLink>Github</FooterLink>
		<FooterLink>{$_('footer.contact')}</FooterLink>
	</FooterLinkGroup>
</Footer>
