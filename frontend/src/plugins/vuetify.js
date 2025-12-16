/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'

// Composables
import { createVuetify } from 'vuetify'

// Arabic locale
import { ar } from 'vuetify/locale'

// Centralized colors - SINGLE SOURCE OF TRUTH
import colors from '@/config/colors'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  // RTL & Locale Configuration
  locale: {
    locale: 'ar',
    fallback: 'en',
    messages: { ar },
    rtl: {
      ar: true,
    },
  },

  // Theme Configuration - uses centralized colors
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors,
      },
    },
  },

  // Global Component Defaults
  defaults: {
    global: {
      ripple: true,
    },
    VBtn: {
      elevation: 0,
      rounded: 'lg',
      variant: 'flat',
    },
    VCard: {
      elevation: 2,
      rounded: 'xl',
    },
    VTextField: {
      variant: 'outlined',
      density: 'comfortable',
      hideDetails: 'auto',
      rounded: 'lg',
    },
    VSelect: {
      variant: 'outlined',
      density: 'comfortable',
      hideDetails: 'auto',
      rounded: 'lg',
    },
    VDataTable: {
      hover: true,
      density: 'comfortable',
    },
    VDialog: {
      maxWidth: 600,
    },
    VChip: {
      rounded: 'lg',
    },
    VAvatar: {
      rounded: 'lg',
    },
  },
})
