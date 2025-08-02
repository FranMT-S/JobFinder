import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import vuetify from './plugins/vuetify'

import './assets/tailwinds.css'
// tailwind
import './assets/style.css'


const pinia = createPinia()
const app = createApp(App)

app
.use(pinia)
.use(vuetify)
.mount('#app')

