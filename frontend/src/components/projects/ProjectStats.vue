<template>
  <v-row class="mb-6 stats-row" no-gutters>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="info">mdi-folder-multiple</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalProjects || 0 }}</h3>
        <p class="text-caption text-info mb-0">إجمالي المشاريع</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="success">mdi-check-circle</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ activeProjects || 0 }}</h3>
        <p class="text-caption text-success mb-0">مشاريع نشطة</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="warning">mdi-clock-alert</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ pendingProjects || 0 }}</h3>
        <p class="text-caption text-warning mb-0">في الانتظار</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="error">mdi-currency-usd</v-icon>
        </div>
        <h3 class="text-h6 font-weight-bold text-error mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; white-space: nowrap;">{{ formattedBudget }}</h3>
        <p class="text-caption text-error mb-0">إجمالي الميزانية</p>
      </v-card>
    </v-col>
    <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
      <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
        <div class="stat-icon mb-3 icon-glow">
          <v-icon size="48" color="primary">mdi-chart-line</v-icon>
        </div>
        <h3 class="text-h4 font-weight-bold text-primary mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ averageProgress || 0 }}%</h3>
        <p class="text-caption text-primary mb-0">متوسط التقدم</p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  totalProjects: {
    type: Number,
    default: 0
  },
  activeProjects: {
    type: Number,
    default: 0
  },
  pendingProjects: {
    type: Number,
    default: 0
  },
  totalBudget: {
    type: Number,
    default: 0
  },
  averageProgress: {
    type: Number,
    default: 0
  }
})

// Local format function to match original component
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const formattedBudget = computed(() => {
  return formatCurrency(props.totalBudget) || '0 د.ع'
})
</script>

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for ProjectStats */

/* Stats row layout */
.stats-row {
  width: 100%;
  justify-content: center;
  margin: 0 auto;
  gap: 0;
  display: flex;
  flex-wrap: wrap;
  align-items: stretch;
  padding: 0;
}

.stats-row .v-col {
  margin: 0 !important;
  padding: 8px !important;
  box-sizing: border-box !important;
}

.stats-row .stat-card {
  width: 100% !important;
  margin: 0 !important;
  box-sizing: border-box !important;
}

/* Enhanced stat card hover effects */
.stat-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 50% 0%, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  border-radius: 24px;
  z-index: 1;
  pointer-events: none;
}

.stat-card .v-icon {
  filter: drop-shadow(0 6px 12px rgba(0, 0, 0, 0.15));
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 60px !important;
  width: 60px !important;
  height: 60px !important;
}

.stat-card:hover .v-icon {
  transform: scale(1.15) rotate(8deg);
  filter: drop-shadow(0 12px 24px rgba(0, 0, 0, 0.25));
}

/* Color-specific effects for each stat card */
.stat-card:nth-child(1)::before {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 50%, #7c3aed 100%);
}

.stat-card:nth-child(1):hover {
  box-shadow: 0 25px 50px rgba(59, 130, 246, 0.25), 0 8px 16px rgba(59, 130, 246, 0.15) !important;
}

.stat-card:nth-child(2)::before {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%);
}

.stat-card:nth-child(2):hover {
  box-shadow: 0 25px 50px rgba(16, 185, 129, 0.25), 0 8px 16px rgba(16, 185, 129, 0.15) !important;
}

.stat-card:nth-child(3)::before {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #b45309 100%);
}

.stat-card:nth-child(3):hover {
  box-shadow: 0 25px 50px rgba(245, 158, 11, 0.25), 0 8px 16px rgba(245, 158, 11, 0.15) !important;
}

.stat-card:nth-child(4)::before {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 50%, #b91c1c 100%);
}

.stat-card:nth-child(4):hover {
  box-shadow: 0 25px 50px rgba(239, 68, 68, 0.25), 0 8px 16px rgba(239, 68, 68, 0.15) !important;
}

.stat-card:nth-child(5)::before {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%);
}

.stat-card:nth-child(5):hover {
  box-shadow: 0 25px 50px rgba(59, 130, 246, 0.25), 0 8px 16px rgba(59, 130, 246, 0.15) !important;
}

/* Responsive adjustments specific to stats row */
@media (max-width: 1200px) {
  .stats-row {
    gap: 0;
    padding: 0;
  }

  .stats-row .v-col {
    padding: 6px !important;
  }
}

@media (max-width: 768px) {
  .stats-row .v-col {
    padding: 4px !important;
  }

  .stat-card .v-icon {
    font-size: 40px !important;
    width: 40px !important;
    height: 40px !important;
  }
}

@media (max-width: 480px) {
  .stats-row .v-col {
    padding: 2px !important;
  }

  .stat-card .v-icon {
    font-size: 32px !important;
    width: 32px !important;
    height: 32px !important;
  }
}
</style>
