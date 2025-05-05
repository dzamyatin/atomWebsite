import { createI18n } from "vue-i18n";
import ru from "./locales/ru.js";
import en from "./locales/en.js";

const messages = {
    ru,
    en,
};

console.log("Одын", ru)
console.log("Два", {
    tst: `123 тестdsa das`,
    "nav.home": "Home",
    "nav.about": "About",
    "home.header": "Welcome to the Vue 3 I18n tutorial!",
    "home.created_by": "This tutorial was brought to you by Lokalise.",
    "about.header": "About us"
})

export default createI18n({
    locale: import.meta.env.VITE_DEFAULT_LOCALE,
    fallbackLocale: import.meta.env.VITE_FALLBACK_LOCALE,
    legacy: false,
    messages,
});
