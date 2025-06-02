import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Buefy from 'buefy';
// import 'buefy/dist/buefy.css';
import i18n from "./i18n"

import App from './App.vue'
import router from './router'

import './client/client.js'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Buefy);
app.use(i18n);

app.mount('#app')
