/**
 * plugins/index.js
 *
 * Automatically included in `./src/main.js`
 */

// Plugins
import { createPinia } from 'pinia'
import vuetify from './vuetify'
import router from '@/router'

const pinia = createPinia()

export function registerPlugins (app) {
  app
    .use(pinia)
    .use(vuetify)
    .use(router)
}
