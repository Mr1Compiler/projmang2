<template>
  <v-card
    class="project-card card-glow hover-lift smooth-transition"
    elevation="4"
    hover
    min-height="400"
  >
    <!-- Project Header -->
    <v-card-title class="project-card-title">
      <div class="d-flex align-center">
        <v-avatar size="40" color="primary" variant="tonal">
          <v-icon>mdi-folder-multiple</v-icon>
        </v-avatar>
        <div class="ms-3 flex-grow-1">
          <h3 class="project-name">{{ project.name }}</h3>
          <v-chip
            size="small"
            :color="statusColor"
            variant="tonal"
            class="status-chip-main"
          >
            <v-icon :icon="statusIcon" size="small" class="me-1" />
            {{ statusText }}
          </v-chip>
        </div>
      </div>
    </v-card-title>

    <v-card-text class="project-card-content">
      <!-- Project Description -->
      <div class="project-description-wrapper">
        <p class="project-description">{{ project.description || 'لا يوجد وصف للمشروع' }}</p>
      </div>

      <!-- Project Details -->
      <div class="project-details">
        <div class="detail-item">
          <v-icon size="small" color="primary">mdi-map-marker</v-icon>
          <span class="detail-text">{{ project.location }}</span>
        </div>

        <div class="detail-item">
          <v-icon size="small" color="success">mdi-currency-usd</v-icon>
          <span class="detail-text">{{ formatCurrency(project.initialCost) }}</span>
        </div>

        <div class="detail-item">
          <v-icon size="small" color="warning">mdi-alert</v-icon>
          <span class="detail-text">{{ formatCurrency(project.criticalCost || project.initialCost * 1.5) }}</span>
        </div>

        <div class="detail-item">
          <v-icon size="small" color="info">mdi-calendar</v-icon>
          <span class="detail-text">{{ formatDate(project.startDate) }}</span>
        </div>

        <div class="detail-item">
          <v-icon size="small" color="secondary">mdi-account</v-icon>
          <span class="detail-text">{{ project.user || project.manager || 'غير محدد' }}</span>
        </div>
      </div>

      <!-- Progress Bar -->
      <div class="progress-section" style="margin-top: 1rem;">
        <div class="progress-header" style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 0.5rem;">
          <span class="progress-label" style="font-family: 'Arial', 'Helvetica', sans-serif; color: #424242; font-size: var(--font-size-sm-plus); font-weight: 600;">نسبة الإنجاز</span>
          <span class="progress-percentage" style="font-family: 'Arial', 'Helvetica', sans-serif; color: #1976d2; font-size: var(--font-size-sm-plus); font-weight: 700;">{{ project.progress || 0 }}%</span>
        </div>
        <v-progress-linear
          :model-value="project.progress || 0"
          color="primary"
          height="8"
          rounded
          class="progress-bar"
          style="border-radius: 4px;"
        />
      </div>
    </v-card-text>

    <v-card-actions class="project-card-actions">
      <v-spacer />
      <v-btn
        color="primary"
        variant="elevated"
        size="default"
        @click="$emit('view-details', project)"
        class="details-btn"
        block
      >
        <v-icon class="me-2">mdi-eye</v-icon>
        عرض التفاصيل
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  project: {
    type: Object,
    required: true
  }
})

defineEmits(['view-details'])

// Helper functions
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const formatDate = (dateString) => {
  if (!dateString) return 'غير محدد'
  const date = new Date(dateString)
  return date.toLocaleDateString('en-GB', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

// Status helpers
const statusColor = computed(() => {
  const colors = {
    pending: 'warning',
    active: 'success',
    completed: 'primary',
    cancelled: 'error'
  }
  return colors[props.project.status] || 'grey'
})

const statusText = computed(() => {
  const texts = {
    pending: 'في الانتظار',
    active: 'نشط',
    completed: 'مكتمل',
    cancelled: 'ملغي'
  }
  return texts[props.project.status] || props.project.status
})

const statusIcon = computed(() => {
  const icons = {
    pending: 'mdi-clock-outline',
    active: 'mdi-play-circle',
    completed: 'mdi-check-circle',
    cancelled: 'mdi-cancel'
  }
  return icons[props.project.status] || 'mdi-help-circle'
})
</script>

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for ProjectCard */

/* Project details layout */
.project-details {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.project-description-wrapper {
  margin-bottom: 1.5rem;
}

/* Progress bar styling */
.progress-bar {
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

/* Status chip styling */
.status-chip-main {
  font-weight: 600 !important;
  font-size: var(--font-size-sm-minus) !important;
}

/* Advanced hover effects for cards */
.project-card .v-avatar {
  transition: all 0.3s ease !important;
}

.project-card:hover .v-avatar {
  transform: scale(1.1) rotate(5deg) !important;
}

.project-card .v-chip {
  transition: all 0.3s ease !important;
}

.project-card:hover .v-chip {
  transform: scale(1.05) !important;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .project-card-actions {
    padding: 0.75rem 1rem 1rem 1rem !important;
  }

  .project-name {
    font-size: var(--font-size-base) !important;
  }

  .detail-item {
    padding: 0.5rem;
  }

  .detail-text {
    font-size: var(--font-size-xs-plus) !important;
  }
}
</style>
