/**
 * Centralized Color Configuration
 *
 * This file is the SINGLE SOURCE OF TRUTH for all colors.
 * Both Vuetify (vuetify.js) and CSS (colors.css) should use these values.
 *
 * To change a color:
 * 1. Update it here
 * 2. Run the app - Vuetify components will update automatically
 * 3. Update colors.css with the same value (or use CSS custom properties)
 */

export const colors = {
  // Primary - Indigo
  primary: '#4338ca',
  'primary-light': '#6366f1',
  'primary-dark': '#3730a3',

  // Secondary - Gray
  secondary: '#424242',
  'secondary-light': '#6d6d6d',
  'secondary-dark': '#212121',

  // Success - Green
  success: '#4caf50',
  'success-light': '#81c784',
  'success-dark': '#388e3c',

  // Warning - Orange
  warning: '#ff9800',
  'warning-light': '#ffb74d',
  'warning-dark': '#f57c00',

  // Error - Red
  error: '#f44336',
  'error-light': '#e57373',
  'error-dark': '#d32f2f',

  // Info - Blue
  info: '#2196f3',
  'info-light': '#64b5f6',
  'info-dark': '#1976d2',

  // Background & Surface
  background: '#ffffff',
  surface: '#ffffff',
}

export default colors
