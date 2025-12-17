<template>
  <div class="page-header">
    <div class="page-header-content">
      <span v-if="icon" class="page-header-icon">{{ icon }}</span>
      <v-icon v-else-if="mdiIcon" :size="iconSize" color="white" class="page-header-icon">{{ mdiIcon }}</v-icon>
      <div class="page-header-text">
        <h1 class="page-header-title">{{ title }}</h1>
        <p v-if="subtitle" class="page-header-subtitle">{{ subtitle }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * PageHeader Component
 * Centralized page header for consistent styling across all pages
 * Following FRONTEND_RULES.md - Smart Pages, Dumb Components
 */

defineProps({
  title: {
    type: String,
    required: true
  },
  subtitle: {
    type: String,
    default: ''
  },
  icon: {
    type: String,
    default: ''
  },
  mdiIcon: {
    type: String,
    default: ''
  },
  iconSize: {
    type: [Number, String],
    default: 40
  }
})
</script>

<style scoped>
@import './styles/shared.css';
</style>
