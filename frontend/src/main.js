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
import '@mdi/font/css/materialdesignicons.css'
import '@/styles/responsive-audit.css'

const app = createApp(App)

registerPlugins(app)

app.mount('#app')
