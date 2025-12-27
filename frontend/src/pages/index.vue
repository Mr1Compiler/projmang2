<template>
  <div class="modern-dashboard">
    <v-container fluid style="padding: 0 20px;">
      <!-- Hero Section - Modern Design -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-view-dashboard-variant</v-icon>
          </div>
            <div class="header-text">
              <h1 class="main-title">لوحة التحكم الرئيسية</h1>
              <p class="subtitle">نظرة شاملة على مشاريعك وإحصائياتك المهمة</p>
            </div>
          </div>
        </div>
        </div>

      <!-- Main Statistics Cards - Modern Grid -->
      <v-row class="stats-grid" no-gutters>
        <!-- Total Projects Card -->
        <v-col cols="12" sm="6" md="3" class="pa-2">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="40" class="stat-icon">mdi-folder-multiple</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.totalProjects }}</h3>
                <p class="stat-label">إجمالي المشاريع</p>
              </div>
              <div class="stat-trend">
                <v-icon size="18" color="success">mdi-trending-up</v-icon>
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
                <v-icon size="40" class="stat-icon">mdi-check-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.activeProjects }}</h3>
                <p class="stat-label">مشاريع نشطة</p>
              </div>
              <div class="stat-trend">
                <v-icon size="18" color="success">mdi-trending-up</v-icon>
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
                <v-icon size="40" class="stat-icon">mdi-check-all</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.completedProjects }}</h3>
                <p class="stat-label">مشاريع مكتملة</p>
              </div>
              <div class="stat-trend">
                <v-icon size="18" color="success">mdi-trending-up</v-icon>
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
                <v-icon size="40" class="stat-icon">mdi-account-hard-hat</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ projectStats.totalEngineers }}</h3>
                <p class="stat-label">المهندسين</p>
              </div>
              <div class="stat-trend">
                <v-icon size="18" color="success">mdi-trending-up</v-icon>
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
            <div class="financial-card-background"></div>
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
            <div class="financial-card-background"></div>
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
            <div class="financial-card-background"></div>
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
                      stroke="#e2e8f0"
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
                    <span class="legend-color" style="background: #8b5cf6;"></span>
                    <span class="legend-text">مكتمل {{ completedPercentage }}%</span>
                  </div>
                  <div class="legend-item">
                    <span class="legend-color" style="background: #3b82f6;"></span>
                    <span class="legend-text">قيد التنفيذ {{ activePercentage }}%</span>
                  </div>
                  <div class="legend-item">
                    <span class="legend-color" style="background: #94a3b8;"></span>
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
                    <span class="legend-dot" style="background: #8b5cf6;"></span>
                    <span>مهام مكتملة</span>
                  </div>
                  <div class="legend-item">
                    <span class="legend-dot" style="background: #94a3b8;"></span>
                    <span>مهام معلقة</span>
                  </div>
                </div>
                <div class="line-chart">
                  <svg viewBox="0 0 400 200" class="line-svg">
                    <!-- Grid lines -->
                    <line x1="0" y1="40" x2="400" y2="40" stroke="#e2e8f0" stroke-width="1" />
                    <line x1="0" y1="80" x2="400" y2="80" stroke="#e2e8f0" stroke-width="1" />
                    <line x1="0" y1="120" x2="400" y2="120" stroke="#e2e8f0" stroke-width="1" />
                    <line x1="0" y1="160" x2="400" y2="160" stroke="#e2e8f0" stroke-width="1" />
                    
                    <!-- Completed tasks line -->
                    <polyline
                      :points="completedTasksLine"
                      fill="none"
                      stroke="#8b5cf6"
                      stroke-width="3"
                    />
                    <!-- Pending tasks line -->
                    <polyline
                      :points="pendingTasksLine"
                      fill="none"
                      stroke="#94a3b8"
                      stroke-width="3"
                    />
                    
                    <!-- Data points -->
                    <circle
                      v-for="(point, index) in completedTasksPoints"
                      :key="'completed-' + index"
                      :cx="point.x"
                      :cy="point.y"
                      r="4"
                      fill="#8b5cf6"
                    />
                    <circle
                      v-for="(point, index) in pendingTasksPoints"
                      :key="'pending-' + index"
                      :cx="point.x"
                      :cy="point.y"
                      r="4"
                      fill="#94a3b8"
                    />
                  </svg>
                  <div class="line-chart-labels">
                    <span v-for="(label, index) in chartLabels" :key="index">{{ label }}</span>
                  </div>
                </div>
                <div class="chart-stats">
                  <div class="stat-item">
                    <span class="stat-label">مكتملة</span>
                    <span class="stat-number">{{ totalCompletedTasks.toLocaleString() }}</span>
                  </div>
                  <div class="stat-item">
                    <span class="stat-label">معلقة</span>
                    <span class="stat-number">{{ totalPendingTasks.toLocaleString() }}</span>
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
import { ref, computed, onMounted, watch } from 'vue'
import {
  getDashboardStats,
  getFinancialStats,
  getProjectProgress,
  getTaskSummary,
  getActivities
} from '@/api/dashboard'

// Loading states
const loading = ref({
  stats: true,
  financial: true,
  progress: true,
  tasks: true,
  activities: true
})

// Error states
const errors = ref({
  stats: null,
  financial: null,
  progress: null,
  tasks: null,
  activities: null
})

// Financial data from API
const incomeData = ref(0)
const expenseData = ref(0)

// Project stats from API
const projectStats = ref({
  totalProjects: 0,
  completedProjects: 0,
  activeProjects: 0,
  totalEngineers: 0
})

// Task summary data from API
const taskSummaryData = ref({
  labels: [],
  completed: [],
  pending: []
})

// Recent Activities from API
const recentActivities = ref([])

// Action translation map for Arabic
const actionMapAr = {
  create: 'أضاف',
  update: 'حدّث',
  delete: 'حذف',
  login: 'سجّل دخول',
  logout: 'سجّل خروج',
  status_change: 'غيّر الحالة',
  refresh: 'جدّد الجلسة'
}

// Target type translation map for Arabic
const targetTypeMapAr = {
  project: 'مشروع',
  expense: 'مصروف',
  income: 'دخل',
  debtor: 'مدين',
  user: 'مستخدم',
  work_category: 'فئة عمل',
  work_subcategory: 'فئة فرعية',
  team_member: 'عضو فريق',
  auth: 'الحساب',
  role: 'دور',
  role_page: 'صلاحية الدور',
  page: 'صفحة',
  user_role: 'دور المستخدم',
  workday: 'يوم عمل',
  workday_material: 'مواد يوم العمل',
  workday_labor: 'عمالة يوم العمل',
  workday_equipment: 'معدات يوم العمل'
}

// Action icon map
const actionIconMap = {
  create: 'mdi-plus-circle',
  update: 'mdi-pencil',
  delete: 'mdi-delete',
  login: 'mdi-login',
  logout: 'mdi-logout',
  status_change: 'mdi-swap-horizontal',
  refresh: 'mdi-refresh'
}

// Action color map
const actionColorMap = {
  create: 'success',
  update: 'primary',
  delete: 'error',
  login: 'info',
  logout: 'warning',
  status_change: 'primary',
  refresh: 'info'
}

// Target type icon map (for better visual representation)
const targetIconMap = {
  project: 'mdi-folder-multiple',
  expense: 'mdi-cash-minus',
  income: 'mdi-cash-plus',
  debtor: 'mdi-account-cash',
  user: 'mdi-account',
  work_category: 'mdi-shape',
  work_subcategory: 'mdi-shape-outline',
  team_member: 'mdi-account-group',
  auth: 'mdi-shield-account',
  role: 'mdi-badge-account',
  role_page: 'mdi-key',
  page: 'mdi-file-document',
  user_role: 'mdi-account-key',
  workday: 'mdi-calendar-today',
  workday_material: 'mdi-package-variant',
  workday_labor: 'mdi-account-hard-hat',
  workday_equipment: 'mdi-tools'
}

// Format activity for display
function formatActivity(activity) {
  const action = actionMapAr[activity.action] || activity.action
  const targetType = targetTypeMapAr[activity.targetType] || activity.targetType
  return `${activity.actorName} ${action} ${targetType} "${activity.targetName}"`
}

// Format relative time in Arabic
function formatRelativeTime(dateString) {
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now - date
  const diffSecs = Math.floor(diffMs / 1000)
  const diffMins = Math.floor(diffSecs / 60)
  const diffHours = Math.floor(diffMins / 60)
  const diffDays = Math.floor(diffHours / 24)

  if (diffSecs < 60) return 'الآن'
  if (diffMins < 60) return `منذ ${diffMins} دقيقة`
  if (diffHours < 24) return `منذ ${diffHours} ساعة`
  if (diffDays === 1) return 'منذ يوم'
  if (diffDays < 7) return `منذ ${diffDays} أيام`
  if (diffDays < 30) return `منذ ${Math.floor(diffDays / 7)} أسبوع`
  return `منذ ${Math.floor(diffDays / 30)} شهر`
}

// Fetch dashboard stats
async function fetchDashboardStats() {
  loading.value.stats = true
  errors.value.stats = null
  try {
    const response = await getDashboardStats()
    if (response.success && response.data) {
      projectStats.value = {
        totalProjects: response.data.totalProjects || 0,
        completedProjects: response.data.completedProjects || 0,
        activeProjects: response.data.activeProjects || 0,
        totalEngineers: response.data.totalEngineers || 0
      }
    }
  } catch (error) {
    console.error('Failed to fetch dashboard stats:', error)
    errors.value.stats = error.message
  } finally {
    loading.value.stats = false
  }
}

// Fetch financial stats
async function fetchFinancialStats() {
  loading.value.financial = true
  errors.value.financial = null
  try {
    const response = await getFinancialStats()
    if (response.success && response.data) {
      incomeData.value = response.data.totalIncome || 0
      expenseData.value = response.data.totalExpenses || 0
    }
  } catch (error) {
    console.error('Failed to fetch financial stats:', error)
    errors.value.financial = error.message
  } finally {
    loading.value.financial = false
  }
}

// Fetch project progress
async function fetchProjectProgress() {
  loading.value.progress = true
  errors.value.progress = null
  try {
    const status = progressFilter.value === 'قيد التنفيذ' ? 'in_progress'
      : progressFilter.value === 'مكتمل' ? 'completed'
      : progressFilter.value === 'في الانتظار' ? 'pending'
      : 'all'
    const response = await getProjectProgress({ status, limit: 4 })
    if (response.success && response.data?.projects) {
      projectProgressData.value = response.data.projects
    }
  } catch (error) {
    console.error('Failed to fetch project progress:', error)
    errors.value.progress = error.message
  } finally {
    loading.value.progress = false
  }
}

// Project progress data from API
const projectProgressData = ref([])

// Fetch task summary
async function fetchTaskSummary() {
  loading.value.tasks = true
  errors.value.tasks = null
  try {
    // Convert Arabic month filter to YYYY-MM format
    const monthMap = {
      'ديسمبر 2025': '2025-12',
      'نوفمبر 2025': '2025-11',
      'أكتوبر 2025': '2025-10',
      'سبتمبر 2025': '2025-09',
      'أغسطس 2025': '2025-08',
      'يوليو 2025': '2025-07',
      'يونيو 2025': '2025-06',
      'مايو 2025': '2025-05',
      'أبريل 2025': '2025-04',
      'مارس 2025': '2025-03',
      'فبراير 2025': '2025-02',
      'يناير 2025': '2025-01'
    }
    const month = monthMap[taskSummaryFilter.value] || getCurrentMonth()
    const response = await getTaskSummary(month)
    if (response.success && response.data) {
      taskSummaryData.value = {
        labels: response.data.labels || [],
        completed: response.data.completed || [],
        pending: response.data.pending || []
      }
    }
  } catch (error) {
    console.error('Failed to fetch task summary:', error)
    errors.value.tasks = error.message
  } finally {
    loading.value.tasks = false
  }
}

// Helper to get current month in YYYY-MM format
function getCurrentMonth() {
  const now = new Date()
  return `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}`
}

// Fetch activities
async function fetchActivities() {
  loading.value.activities = true
  errors.value.activities = null
  try {
    const response = await getActivities({ limit: 5 })
    if (response.success && response.data?.data) {
      recentActivities.value = response.data.data.map(activity => ({
        icon: targetIconMap[activity.targetType] || actionIconMap[activity.action] || 'mdi-information',
        color: actionColorMap[activity.action] || 'grey',
        text: formatActivity(activity),
        time: formatRelativeTime(activity.createdAt)
      }))
    }
  } catch (error) {
    console.error('Failed to fetch activities:', error)
    errors.value.activities = error.message
  } finally {
    loading.value.activities = false
  }
}

// Fetch all data on mount
onMounted(() => {
  fetchDashboardStats()
  fetchFinancialStats()
  fetchProjectProgress()
  fetchTaskSummary()
  fetchActivities()
})

// Reminders
const reminders = ref([
  {
    icon: 'mdi-clipboard-text',
    color: '#3b82f6',
    title: 'مراجعة المهام والتحديثات',
    description: 'مراجعة المهام المكتملة وتحديث لوحة المشروع'
  },
  {
    icon: 'mdi-star',
    color: '#f59e0b',
    title: 'متابعة ملاحظات العميل',
    description: 'مراجعة ملاحظات العميل للمهام الجارية'
  },
  {
    icon: 'mdi-lightbulb',
    color: '#10b981',
    title: 'جلسة تخطيط المشروع',
    description: 'جمع ملاحظات العميل للمهام الجارية'
  }
])

// Progress Chart Data
const progressFilter = ref('قيد التنفيذ')
const progressFilterItems = ['قيد التنفيذ', 'مكتمل', 'في الانتظار', 'الكل']

// Watch progress filter and refetch data
watch(progressFilter, () => {
  fetchProjectProgress()
})

const progressBars = computed(() => {
  const colors = ['#3b82f6', '#8b5cf6', '#10b981', '#f59e0b']
  if (projectProgressData.value.length === 0) {
    return []
  }
  return projectProgressData.value.map((project, index) => ({
    height: project.progress || 0,
    value: `${Math.round(project.progress || 0)}%`,
    label: project.name || `مشروع ${index + 1}`,
    color: colors[index % colors.length]
  }))
})

// Donut Chart Data
const completedColor = ref('#8b5cf6')
const inProgressColor = ref('#3b82f6')
const pendingColor = ref('#94a3b8')

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
const taskSummaryFilter = ref('ديسمبر 2025')
const taskSummaryFilterItems = ['ديسمبر 2025', 'نوفمبر 2025', 'أكتوبر 2025', 'سبتمبر 2025']

// Watch task summary filter and refetch data
watch(taskSummaryFilter, () => {
  fetchTaskSummary()
})

const chartLabels = computed(() => {
  if (taskSummaryData.value.labels.length === 0) {
    return ['1', '2', '3', '4']
  }
  return taskSummaryData.value.labels
})

// Convert task data to chart points
const completedTasksPoints = computed(() => {
  const data = taskSummaryData.value.completed
  if (data.length === 0) return []

  const maxValue = Math.max(...data, ...taskSummaryData.value.pending, 1)
  const chartHeight = 160
  const startX = 50
  const stepX = 300 / Math.max(data.length - 1, 1)

  return data.map((value, index) => ({
    x: startX + index * stepX,
    y: 180 - (value / maxValue) * chartHeight
  }))
})

const pendingTasksPoints = computed(() => {
  const data = taskSummaryData.value.pending
  if (data.length === 0) return []

  const maxValue = Math.max(...taskSummaryData.value.completed, ...data, 1)
  const chartHeight = 160
  const startX = 50
  const stepX = 300 / Math.max(data.length - 1, 1)

  return data.map((value, index) => ({
    x: startX + index * stepX,
    y: 180 - (value / maxValue) * chartHeight
  }))
})

const completedTasksLine = computed(() => {
  return completedTasksPoints.value.map(p => `${p.x},${p.y}`).join(' ')
})

const pendingTasksLine = computed(() => {
  return pendingTasksPoints.value.map(p => `${p.x},${p.y}`).join(' ')
})

// Total tasks for display
const totalCompletedTasks = computed(() => {
  return taskSummaryData.value.completed.reduce((sum, val) => sum + val, 0)
})

const totalPendingTasks = computed(() => {
  return taskSummaryData.value.pending.reduce((sum, val) => sum + val, 0)
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

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount)
}
</script>

<style scoped>
/* ========================================
   لوحة التحكم الحديثة - Modern Dashboard
   ======================================== */

.modern-dashboard {
  min-height: 100vh;
  background: #ffffff !important;
  padding: 0;
  margin: 0;
}

@keyframes gradientShift {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}

/* ========================================
   Hero Section
   ======================================== */

/* Header Styles - نفس تنسيق صفحة المهندسين */
.engineers-header-card {
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.7) 0%, rgba(21, 101, 192, 0.7) 100%);
  border-radius: 0 0 24px 24px;
  width: 100%;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(25, 118, 210, 0.3);
  position: relative;
  overflow: hidden;
  margin-bottom: 0;
  border: none;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: slideInFromTop 0.8s ease-out;
}

.engineers-header-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.15), transparent);
  animation: sweep 3s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.08) 50%, transparent 70%);
  animation: diagonalShimmer 5s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(25, 118, 210, 0.4);
}

.engineers-header-card:hover::before {
  animation: sweep 2s ease-in-out infinite;
}

.engineers-header-card:hover::after {
  animation: diagonalShimmer 3s ease-in-out infinite;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.3);
  animation: gradientFlow 4s ease-in-out infinite;
  z-index: 2;
}

.header-content {
  padding: 1.5rem 2rem !important;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.1) 0%, rgba(21, 101, 192, 0.05) 100%);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 3;
  animation: fadeInUp 0.8s ease-out 0.2s both;
  width: 100%;
  margin: 0 auto;
  border-radius: 0;
  border: none;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1.8rem;
  text-align: right;
  padding: 0.8rem 1.5rem;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.15);
}

.engineer-emoji {
  position: relative;
  animation: slideInFromRight 0.8s ease-out 0.4s both;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 0.5rem;
  transition: transform 0.3s ease;
}

.engineer-emoji:hover {
  transform: scale(1.1);
}

.engineer-emoji .v-icon {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
  transition: all 0.3s ease;
  color: white !important;
  position: relative;
}

.engineer-emoji .v-icon:first-child {
  animation: iconBounce 3s ease-in-out infinite 1.5s;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.main-title {
  color: white !important;
  font-size: 1rem !important;
  font-weight: 600 !important;
  margin: 0;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.3px;
}

.subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 0.7rem !important;
  margin: 0;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.2px;
}

/* ========================================
   Modern Statistics Cards
   ======================================== */

.stats-grid {
  margin-top: 0;
  margin-bottom: 2rem;
}

.modern-stat-card {
  position: relative !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  cursor: pointer !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  height: 100% !important;
  min-height: 120px !important;
  background: #ffffff !important;
}

.modern-stat-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.stat-card-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  opacity: 0.1;
  transition: opacity 0.3s ease;
}

.modern-stat-card:hover .stat-card-background {
  opacity: 0.2;
}

.stat-card-primary .stat-card-background {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.stat-card-success .stat-card-background {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.stat-card-warning .stat-card-background {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.stat-card-info .stat-card-background {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.stat-card-primary {
  background: #ffffff;
  border: 2px solid #3b82f6;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15);
}

.stat-card-success {
  background: #ffffff;
  border: 2px solid #10b981;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.15);
}

.stat-card-warning {
  background: #ffffff;
  border: 2px solid #f59e0b;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.15);
}

.stat-card-info {
  background: #ffffff;
  border: 2px solid #06b6d4;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.15);
}

.stat-card-content {
  position: relative;
  z-index: 2;
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
  align-items: center;
  text-align: center;
}

.stat-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  min-width: 56px;
  min-height: 56px;
  border-radius: 50%;
  margin-bottom: 0.2rem;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  position: relative;
  overflow: visible;
}

.stat-icon-wrapper::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 42px;
  height: 42px;
  border-radius: 50%;
  background: #ffffff;
  z-index: 1;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.stat-card-primary .stat-icon-wrapper {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.stat-card-primary .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.2) 0%, rgba(37, 99, 235, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-success .stat-icon-wrapper {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.stat-card-success .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.2) 0%, rgba(5, 150, 105, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-warning .stat-icon-wrapper {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.stat-card-warning .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 158, 11, 0.2) 0%, rgba(217, 119, 6, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-info .stat-icon-wrapper {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.stat-card-info .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.2) 0%, rgba(8, 145, 178, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
  font-size: 28px !important;
  width: 28px !important;
  height: 28px !important;
  min-width: 28px !important;
  min-height: 28px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.stat-icon-wrapper :deep(.v-icon) {
  font-size: 28px !important;
  width: 28px !important;
  height: 28px !important;
  min-width: 28px !important;
  min-height: 28px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 !important;
  padding: 0 !important;
}

.stat-icon-wrapper :deep(svg) {
  width: 32px !important;
  height: 32px !important;
  font-size: 32px !important;
}

.stat-card-primary .stat-icon {
  color: #3b82f6 !important;
}

.stat-card-primary .stat-icon-wrapper :deep(.v-icon) {
  color: #3b82f6 !important;
}

.stat-card-success .stat-icon {
  color: #10b981 !important;
}

.stat-card-success .stat-icon-wrapper :deep(.v-icon) {
  color: #10b981 !important;
}

.stat-card-warning .stat-icon {
  color: #f59e0b !important;
}

.stat-card-warning .stat-icon-wrapper :deep(.v-icon) {
  color: #f59e0b !important;
}

.stat-card-info .stat-icon {
  color: #06b6d4 !important;
}

.stat-card-info .stat-icon-wrapper :deep(.v-icon) {
  color: #06b6d4 !important;
}

.stat-info {
  flex: 1;
  text-align: center;
  width: 100%;
}

.stat-value {
    font-size: 2rem;
  font-weight: 800;
  margin-bottom: 0.4rem;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  text-align: center;
  font-variant-numeric: tabular-nums;
  unicode-bidi: embed;
  color: #000000 !important;
}

.stat-label {
  font-size: 0.85rem;
  font-weight: 500;
  text-align: center;
  color: #64748b;
  letter-spacing: 0.2px;
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  font-weight: 600;
  padding: 0.5rem 1rem;
  border-radius: 12px;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  font-variant-numeric: tabular-nums;
  background: rgba(0, 0, 0, 0.05);
  color: #64748b;
}

.stat-trend span {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  font-variant-numeric: tabular-nums;
}

/* ========================================
   Financial Cards
   ======================================== */

.financial-stats {
  margin-bottom: 2rem;
}

.financial-card {
  position: relative;
  border-radius: 20px;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.financial-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.financial-card-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  opacity: 0.1;
}

.income-card {
  background: #ffffff;
  border: 2px solid #10b981;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.15);
}

.income-card .financial-card-background {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.expense-card {
  background: #ffffff;
  border: 2px solid #ef4444;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.15);
}

.expense-card .financial-card-background {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
}

.profit-card {
  background: #ffffff;
  border: 2px solid #3b82f6;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15);
}

.profit-card .financial-card-background {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.financial-card-content {
  position: relative;
  z-index: 2;
  padding: 2rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.financial-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 80px;
  height: 80px;
  border-radius: 16px;
  background: rgba(0, 0, 0, 0.05);
}

.income-card .financial-icon {
  background: rgba(16, 185, 129, 0.1);
}

.expense-card .financial-icon {
  background: rgba(239, 68, 68, 0.1);
}

.profit-card .financial-icon {
  background: rgba(59, 130, 246, 0.1);
}

.financial-info {
  flex: 1;
}

.financial-label {
  font-size: 0.8rem;
  color: #64748b;
  margin-bottom: 0.5rem;
  font-weight: 500;
  letter-spacing: 0.2px;
}

.financial-value {
  font-size: 0.75rem;
  font-weight: 800;
  margin-bottom: 0.5rem;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  font-variant-numeric: tabular-nums;
  unicode-bidi: embed;
  color: #000000 !important;
}

.financial-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  border-radius: 12px;
  font-size: 0.9rem;
  font-weight: 600;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  font-variant-numeric: tabular-nums;
  background: rgba(0, 0, 0, 0.05);
  color: #64748b;
}

.financial-badge span {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  font-variant-numeric: tabular-nums;
}

.financial-badge.success {
  background: rgba(16, 185, 129, 0.1);
  color: #10b981;
}

.financial-badge.error {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

.financial-badge.primary {
  background: rgba(59, 130, 246, 0.1);
  color: #3b82f6;
}

/* ========================================
   Charts Section
   ======================================== */

.charts-section {
  margin-bottom: 2rem;
}

.chart-card {
  background: #ffffff;
  border-radius: 20px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  transition: all 0.3s ease;
  height: 100%;
}

.chart-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.chart-card .card-header {
  padding: 0.75rem 1rem !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 48px !important;
  display: flex;
  align-items: center;
}

.chart-card .header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chart-card .header-actions {
  display: flex;
  gap: 0.5rem;
}

.action-icon,
.chart-card .action-icon {
  color: white;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 18px !important;
  width: 18px !important;
  height: 18px !important;
}

.action-icon:hover,
.chart-card .action-icon:hover {
  transform: scale(1.1);
  opacity: 0.8;
}

.chart-filter {
  max-width: 150px;
}

.chart-filter :deep(.v-field__input) {
  color: white !important;
  font-size: 0.75rem;
}

.chart-filter :deep(.v-field__prepend-inner) {
  color: white !important;
}

.chart-filter :deep(.v-field__outline) {
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.chart-card .card-subtitle {
  font-size: 0.65rem !important;
  color: rgba(255, 255, 255, 0.9) !important;
  margin-top: 0.2rem !important;
  margin-bottom: 0 !important;
  line-height: 1.2 !important;
  letter-spacing: 0.2px;
}

.progress-chart-card .card-header {
  padding: 0.15rem 0.5rem !important;
}

.progress-chart-card .header-content {
  align-items: center !important;
  gap: 0.5rem;
}

/* Title styles are unified in .card-title */

.progress-chart-card .chart-filter {
  max-width: 140px !important;
}

.progress-chart-card .chart-filter :deep(.v-field) {
  min-height: 28px !important;
}

.progress-chart-card .chart-filter :deep(.v-field__input) {
  font-size: 0.7rem !important;
  padding: 4px 8px !important;
}

.progress-chart-card .chart-filter :deep(.v-field__control) {
  min-height: 28px !important;
}

/* Completed Chart - Same styling as progress chart */
.completed-chart-card .card-header {
  padding: 0.15rem 0.5rem !important;
}

/* Title styles are unified in .card-title */
/* Subtitle styles are unified in .chart-card .card-subtitle */

/* Task Summary Chart - Same styling as progress chart */
.task-summary-card .card-header {
  padding: 0.15rem 0.5rem !important;
}

.task-summary-card .header-content {
  align-items: center !important;
  gap: 0.5rem;
}

/* Title styles are unified in .card-title */

.task-summary-card .chart-filter {
  max-width: 140px !important;
}

.task-summary-card .chart-filter :deep(.v-field) {
  min-height: 28px !important;
}

.task-summary-card .chart-filter :deep(.v-field__input) {
  font-size: 0.7rem !important;
  padding: 4px 8px !important;
}

.task-summary-card .chart-filter :deep(.v-field__control) {
  min-height: 28px !important;
}

/* Progress Chart */
.progress-chart-card .card-body {
  padding: 1rem;
}

.chart-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.bar-chart {
  display: flex;
  justify-content: space-around;
  align-items: flex-end;
  height: 200px;
  gap: 1rem;
}

.bar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.01rem;
  flex: 1;
}

.bar-wrapper {
  width: 100%;
  height: 200px;
  display: flex;
  align-items: flex-end;
  position: relative;
}

.bar {
  width: 100%;
  border-radius: 8px 8px 0 0;
  position: relative;
  transition: all 0.3s ease;
  min-height: 20px;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding-top: 0.5rem;
}

.bar:hover {
  opacity: 0.8;
  transform: scaleY(1.05);
}

.bar-value {
  color: white;
  font-size: 0.7rem;
  font-weight: 700;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.bar-label {
  font-size: 0.7rem;
  color: #64748b;
  text-align: center;
  font-weight: 500;
}

.chart-summary {
  text-align: center;
  padding: 1rem;
  background: #f8fafc;
  border-radius: 12px;
}

.summary-number {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
}

.summary-label {
  font-size: 0.7rem;
  color: #64748b;
  font-weight: 600;
  letter-spacing: 0.2px;
}

.summary-value {
  font-size: 2.5rem;
  font-weight: 800;
  color: #1e293b;
  font-family: 'Arial', 'Helvetica', sans-serif;
  direction: ltr;
}

.summary-text {
  font-size: 0.75rem;
  color: #64748b;
  margin: 0;
  letter-spacing: 0.2px;
}

/* Donut Chart */
.completed-chart-card .card-body {
  padding: 1rem !important;
}

.donut-chart-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}

.donut-chart {
  position: relative;
  width: 200px;
  height: 200px;
}

.donut-svg {
  width: 100%;
  height: 100%;
}

.donut-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
}

.donut-value {
  font-size: 2rem;
  font-weight: 800;
  color: #1e293b;
  margin: 0;
  font-family: 'Arial', 'Helvetica', sans-serif;
  direction: ltr;
}

.donut-label {
  font-size: 0.75rem;
  color: #64748b;
  margin: 0;
  letter-spacing: 0.2px;
}

.donut-legend {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  width: 100%;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.legend-color {
  width: 16px;
  height: 16px;
  border-radius: 4px;
  flex-shrink: 0;
}

.legend-text {
  font-size: 0.75rem;
  color: #475569;
  font-weight: 500;
  letter-spacing: 0.2px;
}

/* Line Chart */
.task-summary-card .card-body {
  padding: 1rem !important;
}

.line-chart-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.line-chart-legend {
  display: flex;
  gap: 1.5rem;
  justify-content: center;
}

.legend-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  display: inline-block;
  margin-left: 0.5rem;
}

.line-chart {
  position: relative;
  width: 100%;
  height: 200px;
}

.line-svg {
  width: 100%;
  height: 100%;
}

.line-chart-labels {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem 0;
  font-size: 0.7rem;
  color: #64748b;
}

.chart-stats {
  display: flex;
  justify-content: space-around;
  padding: 1rem;
  background: #f8fafc;
  border-radius: 12px;
  gap: 1rem;
}

.chart-stats .stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
}

.chart-stats .stat-label {
  font-size: 0.75rem;
  color: #64748b;
}

.chart-stats .stat-number {
  font-size: 1.5rem;
  font-weight: 800;
  color: #1e293b;
  font-family: 'Arial', 'Helvetica', sans-serif;
  direction: ltr;
}

/* Reminders Card */
.reminders-section {
  margin-bottom: 2rem;
}

.reminders-card {
  background: #ffffff;
  border-radius: 20px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  transition: all 0.3s ease;
  height: 100%;
}

.reminders-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
}

.reminders-card .card-header {
  padding: 0.15rem 0.5rem !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.reminder-item {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  padding: 1rem;
  margin-bottom: 1rem;
  background: #f8fafc;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.reminder-item:hover {
  background: #f1f5f9;
  transform: translateX(-4px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.reminder-item:last-child {
  margin-bottom: 0;
}

.reminder-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  min-width: 48px;
  min-height: 48px;
  border-radius: 12px;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;
}

.reminder-icon-wrapper:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.reminder-icon-wrapper .v-icon {
  color: #ffffff !important;
}

.reminder-content {
  flex: 1;
}

.reminder-title {
  font-size: 0.8rem;
  font-weight: 600;
  color: #1e293b;
  margin: 0 0 0.25rem 0;
  letter-spacing: 0.2px;
}

.reminder-description {
  font-size: 0.75rem;
  color: #64748b;
  margin: 0;
  line-height: 1.5;
}

/* ========================================
   Progress & Activity Cards
   ======================================== */

.progress-section {
  margin-bottom: 2rem;
}

.progress-card,
.recent-activity-card {
  background: #ffffff;
  border-radius: 20px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  transition: all 0.3s ease;
}

.progress-card:hover,
.recent-activity-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15);
}

.card-header {
  padding: 0.75rem;
  border-bottom: 1px solid #e2e8f0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.card-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.95rem !important;
  font-weight: 600 !important;
  color: white !important;
  margin: 0 !important;
  line-height: 1.3 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  letter-spacing: 0.3px !important;
}

/* توحيد حجم العناوين في جميع الكروت */
.chart-card .card-title,
.progress-chart-card .card-title,
.completed-chart-card .card-title,
.task-summary-card .card-title,
.reminders-card .card-title,
.activity-card .card-title {
  font-size: 0.95rem !important;
  font-weight: 600 !important;
  line-height: 1.3 !important;
}

.title-icon {
  color: white !important;
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.2)) !important;
  font-size: 18px !important;
  width: 18px !important;
  height: 18px !important;
  flex-shrink: 0 !important;
}

/* توحيد حجم الأيقونات في جميع الكروت */
.chart-card .title-icon,
.progress-chart-card .title-icon,
.completed-chart-card .title-icon,
.task-summary-card .title-icon,
.reminders-card .title-icon,
.activity-card .title-icon {
  font-size: 18px !important;
  width: 18px !important;
  height: 18px !important;
}

.card-body {
  padding: 1rem !important;
}

.progress-item {
  margin-bottom: 1.5rem;
}

.progress-item:last-child {
  margin-bottom: 0;
}

.progress-label {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
  color: #475569;
  font-size: 0.9rem;
  font-weight: 500;
}

.progress-percentage {
  font-weight: 700;
  color: #1e293b;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  font-variant-numeric: tabular-nums;
}

.progress-bar {
  border-radius: 8px;
  overflow: hidden;
}

.activity-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  margin-bottom: 1rem;
  background: #f8fafc;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.activity-item:hover {
  background: #f1f5f9;
  transform: translateX(-4px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.activity-item:last-child {
  margin-bottom: 0;
}

.activity-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  border-radius: 12px;
  background: rgba(0, 0, 0, 0.05);
}

.activity-content {
  flex: 1;
}

.activity-text {
  font-size: 0.75rem;
  color: #1e293b;
  margin-bottom: 0.25rem;
  font-weight: 500;
  letter-spacing: 0.2px;
}

.activity-time {
  font-size: 0.75rem;
  color: #64748b;
}

/* ========================================
   Responsive Design
   ======================================== */

@media (max-width: 960px) {
  .hero-title {
    font-size: 1.1rem;
  }

  .hero-icon {
    font-size: 40px !important;
    width: 40px !important;
    height: 40px !important;
  }

  .stat-value {
    font-size: 2rem;
  }

  .financial-value {
    font-size: 0.7rem;
  }
}

@media (max-width: 600px) {
  .modern-dashboard {
    padding: 1rem 0;
  }

  .dashboard-hero {
    padding: 1.5rem 1rem;
  }

  .hero-content {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }

  .hero-text-wrapper {
    text-align: center;
  }

  .hero-title {
    font-size: 1rem;
    text-align: center;
  }

  .hero-subtitle {
    font-size: 0.7rem;
    text-align: center;
  }

  .hero-icon {
    font-size: 36px !important;
    width: 36px !important;
    height: 36px !important;
  }

  .stat-card-content {
    padding: 1.5rem;
  }

  .stat-value {
    font-size: 1.75rem;
  }

  .financial-value {
    font-size: 0.65rem;
  }

  .financial-card-content {
    flex-direction: column;
    text-align: center;
  }
}

/* ========================================
   Animations - نفس صفحة إدارة المهام
   ======================================== */

@keyframes slideInFromTop {
  0% {
    transform: translateY(-100px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes fadeInUp {
  0% {
    transform: translateY(30px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes sweep {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes diagonalShimmer {
  0%, 100% {
    transform: translateX(-100%) translateY(-100%) rotate(45deg);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translateX(100%) translateY(100%) rotate(45deg);
    opacity: 0;
  }
}

@keyframes slideInFromRight {
  0% {
    transform: translateX(50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes gradientFlow {
  0%, 100% {
    background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  }
  50% {
    background: linear-gradient(90deg, #bbdefb 0%, #ffffff 50%, #e3f2fd 100%);
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px) scale(1);
  }
  50% {
    transform: translateY(-6px) scale(1.05);
  }
}

</style>
