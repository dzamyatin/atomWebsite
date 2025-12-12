import { createI18n } from "vue-i18n";
import ru from "./locales/ru.js";
import en from "./locales/en.js";

const messages = {
    ru,
    en,
};

export default createI18n({
    locale: import.meta.env.VITE_DEFAULT_LOCALE,
    fallbackLocale: import.meta.env.VITE_FALLBACK_LOCALE,
    legacy: false,
    messages,
});
