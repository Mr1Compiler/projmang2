<template>
  <div class="modern-dashboard">
    <v-container fluid class="px-5">
      <!-- Hero Section -->
      <PageHeader
        title="لوحة التحكم الرئيسية"
        subtitle="نظرة شاملة على مشاريعك وإحصائياتك المهمة"
        mdi-icon="mdi-view-dashboard-variant"
      />

      <!-- Main Statistics Cards - Modern Grid -->
      <v-row class="stats-grid" no-gutters>
        <!-- Total Projects Card -->
        <v-col cols="12" sm="6" md="3" class="pa-2">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-folder-multiple</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.totalProjects }}</h3>
                <p class="stat-label">إجمالي المشاريع</p>
              </div>
              <div class="stat-trend">
                <v-icon size="20" color="success">mdi-trending-up</v-icon>
                <span>+12%</span>
              </div>
            </div>
          </v-card>
        </v-col>

        <!-- Active Projects Card -->
        <v-col cols="12" sm="6" md="3" class="pa-2">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-check-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.activeProjects }}</h3>
                <p class="stat-label">مشاريع نشطة</p>
              </div>
              <div class="stat-trend">
                <v-icon size="20" color="success">mdi-trending-up</v-icon>
                <span>+5%</span>
              </div>
            </div>
          </v-card>
        </v-col>

        <!-- Completed Projects Card -->
        <v-col cols="12" sm="6" md="3" class="pa-2">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-check-all</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.completedProjects }}</h3>
                <p class="stat-label">مشاريع مكتملة</p>
              </div>
              <div class="stat-trend">
                <v-icon size="20" color="success">mdi-trending-up</v-icon>
                <span>+8%</span>
              </div>
            </div>
          </v-card>
        </v-col>

        <!-- Total Engineers Card -->
        <v-col cols="12" sm="6" md="3" class="pa-2">
          <v-card class="modern-stat-card stat-card-info" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-account-hard-hat</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.totalEngineers }}</h3>
                <p class="stat-label">المهندسين</p>
              </div>
              <div class="stat-trend">
                <v-icon size="20" color="success">mdi-trending-up</v-icon>
                <span>+3</span>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Financial Statistics Row -->
      <v-row class="financial-stats" no-gutters>
        <v-col cols="12" md="4" class="pa-2">
          <v-card class="financial-card income-card" elevation="0">
            <div class="financial-card-content">
              <div class="financial-icon">
                <v-icon size="56" color="success">mdi-arrow-down-circle</v-icon>
              </div>
              <div class="financial-info">
                <p class="financial-label">إجمالي الإيرادات</p>
                <h2 class="financial-value">{{ formatCurrency(totalIncome) }}</h2>
                <div class="financial-badge success">
                  <v-icon size="16">mdi-trending-up</v-icon>
                  <span>+15%</span>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>

        <v-col cols="12" md="4" class="pa-2">
          <v-card class="financial-card expense-card" elevation="0">
            <div class="financial-card-content">
              <div class="financial-icon">
                <v-icon size="56" color="error">mdi-arrow-up-circle</v-icon>
              </div>
              <div class="financial-info">
                <p class="financial-label">إجمالي المصروفات</p>
                <h2 class="financial-value">{{ formatCurrency(totalExpenses) }}</h2>
                <div class="financial-badge error">
                  <v-icon size="16">mdi-trending-down</v-icon>
                  <span>-8%</span>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>

        <v-col cols="12" md="4" class="pa-2">
          <v-card class="financial-card profit-card" elevation="0">
            <div class="financial-card-content">
              <div class="financial-icon">
                <v-icon size="56" color="primary">mdi-chart-line</v-icon>
              </div>
              <div class="financial-info">
                <p class="financial-label">صافي الربح</p>
                <h2 class="financial-value">{{ formatCurrency(netProfit) }}</h2>
                <div class="financial-badge primary">
                  <v-icon size="16">mdi-trending-up</v-icon>
                  <span>+23%</span>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Charts Section -->
      <v-row class="charts-section" no-gutters>
        <!-- Project Progress Chart -->
        <v-col cols="12" md="4" class="pa-2">
          <v-card class="chart-card progress-chart-card" elevation="0">
            <div class="card-header">
              <div class="header-content">
                <h3 class="card-title">
                  <v-icon class="title-icon">mdi-chart-bar</v-icon>
                  تقدم المشاريع
                </h3>
                <v-select
                  v-model="progressFilter"
                  :items="progressFilterItems"
                  density="compact"
                  variant="outlined"
                  hide-details
                  class="chart-filter"
                ></v-select>
              </div>
            </div>
            <div class="card-body">
              <div class="chart-container">
                <div class="bar-chart">
                  <div class="bar-item" v-for="(bar, index) in progressBars" :key="index">
                    <div class="bar-wrapper">
                      <div 
                        class="bar" 
                        :style="{ height: bar.height + '%', backgroundColor: bar.color }"
                      >
                        <span class="bar-value">{{ bar.value }}</span>
                      </div>
                    </div>
                    <span class="bar-label">{{ bar.label }}</span>
                  </div>
                </div>
                <div class="chart-summary">
                  <div class="summary-number">
                    <span class="summary-label">30 يوم</span>
                    <h2 class="summary-value">{{ projectStats.activeProjects }}</h2>
                    <p class="summary-text">مشاريع قيد التنفيذ هذا الشهر</p>
                  </div>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>

        <!-- Project Completed Chart -->
        <v-col cols="12" md="4" class="pa-2">
          <v-card class="chart-card completed-chart-card" elevation="0">
            <div class="card-header">
              <h3 class="card-title">
                <v-icon class="title-icon">mdi-chart-donut</v-icon>
                المشاريع المكتملة
              </h3>
              <p class="card-subtitle">إجمالي المشاريع {{ projectStats.totalProjects }}</p>
            </div>
            <div class="card-body">
              <div class="donut-chart-container">
                <div class="donut-chart">
                  <svg viewBox="0 0 200 200" class="donut-svg">
                    <circle
                      cx="100"
                      cy="100"
                      r="80"
                      fill="none"
                      stroke="var(--color-slate-200)"
                      stroke-width="20"
                    />
                    <circle
                      cx="100"
                      cy="100"
                      r="80"
                      fill="none"
                      :stroke="completedColor"
                      stroke-width="20"
                      :stroke-dasharray="completedCircumference"
                      :stroke-dashoffset="completedOffset"
                      stroke-linecap="round"
                      transform="rotate(-90 100 100)"
                    />
                    <circle
                      cx="100"
                      cy="100"
                      r="60"
                      fill="none"
                      :stroke="inProgressColor"
                      stroke-width="20"
                      :stroke-dasharray="inProgressCircumference"
                      :stroke-dashoffset="inProgressOffset"
                      stroke-linecap="round"
                      transform="rotate(-90 100 100)"
                    />
                    <circle
                      cx="100"
                      cy="100"
                      r="40"
                      fill="none"
                      :stroke="pendingColor"
                      stroke-width="20"
                      :stroke-dasharray="pendingCircumference"
                      :stroke-dashoffset="pendingOffset"
                      stroke-linecap="round"
                      transform="rotate(-90 100 100)"
                    />
                  </svg>
                  <div class="donut-center">
                    <h3 class="donut-value">{{ completedPercentage }}%</h3>
                    <p class="donut-label">مكتمل</p>
                  </div>
                </div>
                <div class="donut-legend">
                  <div class="legend-item">
                    <span class="legend-color" style="background: var(--color-violet-500);"></span>
                    <span class="legend-text">مكتمل {{ completedPercentage }}%</span>
                  </div>
                  <div class="legend-item">
                    <span class="legend-color" style="background: var(--color-blue-500);"></span>
                    <span class="legend-text">قيد التنفيذ {{ activePercentage }}%</span>
                  </div>
                  <div class="legend-item">
                    <span class="legend-color" style="background: var(--color-slate-400);"></span>
                    <span class="legend-text">في الانتظار {{ pendingPercentage }}%</span>
                  </div>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>

        <!-- Task Summary Chart -->
        <v-col cols="12" md="4" class="pa-2">
          <v-card class="chart-card task-summary-card" elevation="0">
            <div class="card-header">
              <div class="header-content">
                <h3 class="card-title">
                  <v-icon class="title-icon">mdi-chart-line</v-icon>
                  ملخص المهام
                </h3>
                <v-select
                  v-model="taskSummaryFilter"
                  :items="taskSummaryFilterItems"
                  density="compact"
                  variant="outlined"
                  hide-details
                  class="chart-filter"
                ></v-select>
              </div>
            </div>
            <div class="card-body">
              <div class="line-chart-container">
                <div class="line-chart-legend">
                  <div class="legend-item">
                    <span class="legend-dot" style="background: var(--color-violet-500);"></span>
                    <span>مهام مكتملة</span>
                  </div>
                  <div class="legend-item">
                    <span class="legend-dot" style="background: var(--color-slate-400);"></span>
                    <span>مهام معلقة</span>
                  </div>
                </div>
                <div class="line-chart">
                  <svg viewBox="0 0 400 200" class="line-svg">
                    <!-- Grid lines -->
                    <line x1="0" y1="40" x2="400" y2="40" stroke="var(--color-slate-200)" stroke-width="1" />
                    <line x1="0" y1="80" x2="400" y2="80" stroke="var(--color-slate-200)" stroke-width="1" />
                    <line x1="0" y1="120" x2="400" y2="120" stroke="var(--color-slate-200)" stroke-width="1" />
                    <line x1="0" y1="160" x2="400" y2="160" stroke="var(--color-slate-200)" stroke-width="1" />
                    
                    <!-- Completed tasks line -->
                    <polyline
                      :points="completedTasksLine"
                      fill="none"
                      stroke="var(--color-violet-500)"
                      stroke-width="3"
                    />
                    <!-- Pending tasks line -->
                    <polyline
                      :points="pendingTasksLine"
                      fill="none"
                      stroke="var(--color-slate-400)"
                      stroke-width="3"
                    />
                    
                    <!-- Data points -->
                    <circle
                      v-for="(point, index) in completedTasksPoints"
                      :key="'completed-' + index"
                      :cx="point.x"
                      :cy="point.y"
                      r="4"
                      fill="var(--color-violet-500)"
                    />
                    <circle
                      v-for="(point, index) in pendingTasksPoints"
                      :key="'pending-' + index"
                      :cx="point.x"
                      :cy="point.y"
                      r="4"
                      fill="var(--color-slate-400)"
                    />
                  </svg>
                  <div class="line-chart-labels">
                    <span v-for="(label, index) in chartLabels" :key="index">{{ label }}</span>
                  </div>
                </div>
                <div class="chart-stats">
                  <div class="stat-item">
                    <span class="stat-label">مكتملة</span>
                    <span class="stat-number">49,225</span>
                  </div>
                  <div class="stat-item">
                    <span class="stat-label">معلقة</span>
                    <span class="stat-number">82,214</span>
                  </div>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Reminders and Activity Section -->
      <v-row class="reminders-section" no-gutters>
        <v-col cols="12" md="6" class="pa-2">
          <v-card class="reminders-card" elevation="0">
            <div class="card-header">
              <div class="header-content">
                <h3 class="card-title">
                  <v-icon class="title-icon">mdi-bell-outline</v-icon>
                  التذكيرات
                </h3>
                <div class="header-actions">
                  <v-icon size="20" class="action-icon">mdi-information-outline</v-icon>
                  <v-icon size="20" class="action-icon">mdi-pencil</v-icon>
                  <v-icon size="20" class="action-icon">mdi-dots-vertical</v-icon>
                </div>
              </div>
            </div>
            <div class="card-body">
              <div class="reminder-item" v-for="(reminder, index) in reminders" :key="index">
                <div class="reminder-icon-wrapper" :style="{ background: reminder.color }">
                  <v-icon color="white" size="24">{{ reminder.icon }}</v-icon>
                </div>
                <div class="reminder-content">
                  <h4 class="reminder-title">{{ reminder.title }}</h4>
                  <p class="reminder-description">{{ reminder.description }}</p>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>

        <v-col cols="12" md="6" class="pa-2">
          <v-card class="recent-activity-card" elevation="0">
            <div class="card-header">
              <h3 class="card-title">
                <v-icon class="title-icon">mdi-clock-outline</v-icon>
                النشاط الأخير
              </h3>
            </div>
            <div class="card-body">
              <div class="activity-item" v-for="(activity, index) in recentActivities" :key="index">
                <div class="activity-icon">
                  <v-icon :color="activity.color" size="24">{{ activity.icon }}</v-icon>
                </div>
                <div class="activity-content">
                  <p class="activity-text">{{ activity.text }}</p>
                  <span class="activity-time">{{ activity.time }}</span>
                </div>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { formatCurrency } from '@/utils/formatters'
import { PageHeader } from '@/components/shared'

// Sample data - in real app, this would come from API
const incomeData = ref(175000) // Total income
const expenseData = ref(103000) // Total expenses

const projectStats = ref({
  totalProjects: 24,
  completedProjects: 18,
  activeProjects: 6,
  totalEngineers: 12
})

// Recent Activities
const recentActivities = ref([
  {
    icon: 'mdi-check-circle',
    color: 'success',
    text: 'تم إكمال مشروع "بناء مجمع سكني"',
    time: 'منذ ساعتين'
  },
  {
    icon: 'mdi-folder-plus',
    color: 'primary',
    text: 'تم إضافة مشروع جديد "مركز تجاري"',
    time: 'منذ 4 ساعات'
  },
  {
    icon: 'mdi-account-plus',
    color: 'info',
    text: 'تم إضافة مهندس جديد للفريق',
    time: 'منذ يوم'
  },
  {
    icon: 'mdi-currency-usd',
    color: 'success',
    text: 'تم تسجيل إيراد جديد بقيمة 50,000 د.ع',
    time: 'منذ يومين'
  }
])

// Reminders
const reminders = ref([
  {
    icon: 'mdi-clipboard-text',
    color: 'var(--color-blue-500)',
    title: 'مراجعة المهام والتحديثات',
    description: 'مراجعة المهام المكتملة وتحديث لوحة المشروع'
  },
  {
    icon: 'mdi-star',
    color: 'var(--color-amber-500)',
    title: 'متابعة ملاحظات العميل',
    description: 'مراجعة ملاحظات العميل للمهام الجارية'
  },
  {
    icon: 'mdi-lightbulb',
    color: 'var(--color-emerald-500)',
    title: 'جلسة تخطيط المشروع',
    description: 'جمع ملاحظات العميل للمهام الجارية'
  }
])

// Progress Chart Data
const progressFilter = ref('قيد التنفيذ')
const progressFilterItems = ['قيد التنفيذ', 'مكتمل', 'في الانتظار', 'الكل']

const progressBars = computed(() => [
  { height: 75, value: '+12%', label: 'مشروع 1', color: 'var(--color-blue-500)' },
  { height: 45, value: '+4%', label: 'مشروع 2', color: 'var(--color-violet-500)' },
  { height: 85, value: '+17%', label: 'مشروع 3', color: 'var(--color-blue-500)' },
  { height: 60, value: '+8%', label: 'مشروع 4', color: 'var(--color-violet-500)' }
])

// Donut Chart Data
const completedColor = ref('var(--color-violet-500)')
const inProgressColor = ref('var(--color-blue-500)')
const pendingColor = ref('var(--color-slate-400)')

const completedCircumference = computed(() => {
  const radius = 80
  const circumference = 2 * Math.PI * radius
  return `${circumference} ${circumference}`
})

const completedOffset = computed(() => {
  const radius = 80
  const circumference = 2 * Math.PI * radius
  return circumference - (completedPercentage.value / 100) * circumference
})

const inProgressCircumference = computed(() => {
  const radius = 60
  const circumference = 2 * Math.PI * radius
  return `${circumference} ${circumference}`
})

const inProgressOffset = computed(() => {
  const radius = 60
  const circumference = 2 * Math.PI * radius
  return circumference - (activePercentage.value / 100) * circumference
})

const pendingCircumference = computed(() => {
  const radius = 40
  const circumference = 2 * Math.PI * radius
  return `${circumference} ${circumference}`
})

const pendingOffset = computed(() => {
  const radius = 40
  const circumference = 2 * Math.PI * radius
  return circumference - (pendingPercentage.value / 100) * circumference
})

// Task Summary Chart Data
const taskSummaryFilter = ref('أبريل 2025')
const taskSummaryFilterItems = ['أبريل 2025', 'مارس 2025', 'فبراير 2025', 'يناير 2025']

const chartLabels = ref(['18', '19', '20', '21', '22', '23', '24', '25'])

const completedTasksPoints = computed(() => [
  { x: 50, y: 120 },
  { x: 100, y: 100 },
  { x: 150, y: 90 },
  { x: 200, y: 85 },
  { x: 250, y: 80 },
  { x: 300, y: 75 },
  { x: 350, y: 70 }
])

const pendingTasksPoints = computed(() => [
  { x: 50, y: 160 },
  { x: 100, y: 150 },
  { x: 150, y: 140 },
  { x: 200, y: 130 },
  { x: 250, y: 125 },
  { x: 300, y: 120 },
  { x: 350, y: 115 }
])

const completedTasksLine = computed(() => {
  return completedTasksPoints.value.map(p => `${p.x},${p.y}`).join(' ')
})

const pendingTasksLine = computed(() => {
  return pendingTasksPoints.value.map(p => `${p.x},${p.y}`).join(' ')
})

// Computed properties
const totalIncome = computed(() => incomeData.value)
const totalExpenses = computed(() => expenseData.value)
const netProfit = computed(() => totalIncome.value - totalExpenses.value)

const activePercentage = computed(() => {
  if (projectStats.value.totalProjects === 0) return 0
  return Math.round((projectStats.value.activeProjects / projectStats.value.totalProjects) * 100)
})

const completedPercentage = computed(() => {
  if (projectStats.value.totalProjects === 0) return 0
  return Math.round((projectStats.value.completedProjects / projectStats.value.totalProjects) * 100)
})

const pendingPercentage = computed(() => {
  if (projectStats.value.totalProjects === 0) return 0
  const pending = projectStats.value.totalProjects - projectStats.value.activeProjects - projectStats.value.completedProjects
  return Math.round((pending / projectStats.value.totalProjects) * 100)
})
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/index.css';
</style>
