<script>
	import { _ } from 'svelte-i18n';
	import '$lib/localization/i18n';
</script>

<p>{$_('home.message', { values: { name: 'Maks' } })}</p>
