import { init, register } from "svelte-i18n";
const defaultLocale = 'en'

register('en', () => import("./locales/en.json"))
register('ua', () => import("./locales/ua.json"))

console.log('initializing i18n')

init({
    fallbackLocale: defaultLocale,
    initialLocale: defaultLocale,
    handleMissingMessage: (options) => {
        return `Missing ${options.locale} translation for {${options.id}}`
    }
})