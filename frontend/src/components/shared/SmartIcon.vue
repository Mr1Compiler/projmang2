<template>
  <span class="smart-icon-wrapper" :style="wrapperStyle">
    <!-- Try MDI first -->
    <i
      v-if="useMdi"
      :class="['mdi', mdiClass, customClass]"
      :style="iconStyle"
      ref="mdiIcon"
    ></i>

    <!-- Fallback to emoji -->
    <span
      v-else
      :class="['emoji-fallback', customClass]"
      :style="iconStyle"
    >
      {{ emojiIcon }}
    </span>
  </span>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'

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
  customClass: {
    type: String,
    default: ''
  }
})

const useMdi = ref(true)
const mdiIcon = ref(null)

const mdiClass = computed(() => {
  return props.icon.startsWith('mdi-') ? props.icon : `mdi-${props.icon}`
})

const iconStyle = computed(() => ({
  fontSize: typeof props.size === 'number' ? `${props.size}px` : props.size,
  color: props.color
}))

const wrapperStyle = computed(() => ({
  display: 'inline-block',
  lineHeight: '1'
}))

// Emoji fallbacks
const emojiMap = {
  'home': '🏠',
  'trending-up': '📈',
  'trending-down': '📉',
  'currency-usd': '💵',
  'chart-line': '📊',
  'plus': '➕',
  'minus': '➖',
  'check': '✅',
  'close': '❌',
  'delete': '🗑️',
  'pencil': '✏️',
  'edit': '✏️',
  'arrow-left': '←',
  'arrow-right': '→',
  'arrow-up': '↑',
  'arrow-down': '↓',
  'heart': '❤️',
  'star': '⭐',
  'account': '👤',
  'cash': '💰',
  'wallet': '👛',
  'credit-card': '💳',
  'bank': '🏦',
  'receipt': '🧾',
  'calculator': '🧮',
  'calendar': '📅',
  'clock': '🕐',
  'settings': '⚙️',
  'information': 'ℹ️',
  'alert': '⚠️',
  'success': '✅',
  'error': '❌',
  'folder': '📁',
  'file': '📄',
  'search': '🔍',
  'filter': '🔽',
  'refresh': '🔄',
  'download': '⬇️',
  'upload': '⬆️',
  'print': '🖨️',
  'email': '📧',
  'phone': '📞',
  'location': '📍',
  'map': '🗺️',
  'user': '👤',
  'users': '👥',
  'team': '👥',
  'project': '📊',
  'task': '✓',
  'money': '💰',
  'expense': '💸',
  'income': '💵'
}

const emojiIcon = computed(() => {
  const iconName = props.icon.replace('mdi-', '')
  return emojiMap[iconName] || '❓'
})

// Check if MDI font is loaded
onMounted(async () => {
  await nextTick()

  // Try to detect if MDI font is loaded
  const testElement = document.createElement('span')
  testElement.className = 'mdi mdi-home'
  testElement.style.position = 'absolute'
  testElement.style.visibility = 'hidden'
  testElement.style.fontSize = '24px'
  document.body.appendChild(testElement)

  // Check if the icon renders properly
  setTimeout(() => {
    const computedStyle = window.getComputedStyle(testElement)
    const fontFamily = computedStyle.fontFamily

    // If font family doesn't include "Material Design Icons", use emoji fallback
    if (!fontFamily.includes('Material Design Icons')) {
      useMdi.value = false
    }

    document.body.removeChild(testElement)
  }, 100)
})
</script>

<style scoped>
@import './styles/shared.css';

.smart-icon-wrapper {
  display: inline-block;
  line-height: 1;
  vertical-align: middle;
}

.mdi {
  font-family: "Material Design Icons" !important;
  font-weight: normal !important;
  font-style: normal !important;
  display: inline-block !important;
  line-height: 1 !important;
  text-rendering: auto !important;
  -webkit-font-smoothing: antialiased !important;
  -moz-osx-font-smoothing: grayscale !important;
}

.emoji-fallback {
  display: inline-block;
  line-height: 1;
  font-family: "Apple Color Emoji", "Segoe UI Emoji", "Noto Color Emoji", "Android Emoji", sans-serif;
}
</style>