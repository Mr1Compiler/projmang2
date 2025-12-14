<template>
  <span 
    :class="['icon-fallback', iconClass]" 
    :style="{ fontSize: size + 'px', color: color }"
  >
    {{ getIcon() }}
  </span>
</template>

<script setup>
const props = defineProps({
  icon: {
    type: String,
    required: true
  },
  size: {
    type: [String, Number],
    default: 24
  },
  color: {
    type: String,
    default: 'currentColor'
  },
  iconClass: {
    type: String,
    default: ''
  }
})

const iconMap = {
  'mdi-home': '🏠',
  'mdi-trending-up': '📈',
  'mdi-trending-down': '📉',
  'mdi-currency-usd': '💵',
  'mdi-chart-line': '📊',
  'mdi-plus': '➕',
  'mdi-minus': '➖',
  'mdi-check': '✅',
  'mdi-close': '❌',
  'mdi-delete': '🗑️',
  'mdi-pencil': '✏️',
  'mdi-edit': '✏️',
  'mdi-arrow-left': '←',
  'mdi-arrow-right': '→',
  'mdi-arrow-up': '↑',
  'mdi-arrow-down': '↓',
  'mdi-heart': '❤️',
  'mdi-star': '⭐',
  'mdi-account': '👤',
  'mdi-cash': '💰',
  'mdi-wallet': '👛',
  'mdi-credit-card': '💳',
  'mdi-bank': '🏦',
  'mdi-receipt': '🧾',
  'mdi-calculator': '🧮',
  'mdi-calendar': '📅',
  'mdi-clock': '🕐',
  'mdi-settings': '⚙️',
  'mdi-information': 'ℹ️',
  'mdi-alert': '⚠️',
  'mdi-success': '✅',
  'mdi-error': '❌'
}

function getIcon() {
  return iconMap[props.icon] || '❓'
}
</script>

<style scoped>
.icon-fallback {
  display: inline-block;
  line-height: 1;
  vertical-align: middle;
}
</style>
