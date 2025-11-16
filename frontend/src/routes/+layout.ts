import { browser } from "$app/environment";
import { locale, waitLocale } from "svelte-i18n";
import type { LayoutLoad } from "./$types";
import '$lib/localization/i18n';
import { GetUserLocalSettings } from "$lib/wailsjs/go/main/App";
import { LocalSettings } from "$lib/stores/LocalSettings";


export const prerender = 'auto'; 
export const ssr = false;

export const load: LayoutLoad = async () => {
    const settings = await GetUserLocalSettings()
    LocalSettings.set(settings)

    if (browser) {
        locale.set(settings.locale ?? "en");
    }
    await waitLocale();

};
