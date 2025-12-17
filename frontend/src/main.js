/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Styles
import 'unfonts.css'
import '@/styles/main.css'
import '@/styles/colors.css'
import '@/styles/rtl.css'
import '@/styles/visual-effects.css'
import '@/styles/table-headers.css'
import '@/styles/form-dialog.css'
import '@mdi/font/css/materialdesignicons.css'
import 'vue3-toastify/dist/index.css'

const app = createApp(App)

registerPlugins(app)

app.mount('#app')
