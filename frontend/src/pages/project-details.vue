<template>
  <div class="project-details-page">
    <v-container fluid class="fill-height">
      <!-- Header Section -->
      <div class="page-header">
        <v-btn 
          icon="mdi-arrow-right" 
          @click="goBack" 
          class="back-btn"
          size="large"
          variant="text"
          color="primary"
        >
          <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
        <div class="header-content">
          <h1 class="page-title">تفاصيل المشروع</h1>
          <p class="page-subtitle">{{ project?.name || 'تفاصيل المشروع' }}</p>
        </div>
      </div>

      <!-- Project Details Cards -->
      <v-row v-if="project" class="details-cards">
        <!-- Basic Information Card -->
        <v-col cols="12" md="6">
          <v-card class="detail-card info-card" elevation="3">
            <v-card-title class="card-header info-header">
              <v-icon class="me-2">mdi-information</v-icon>
              المعلومات الأساسية
            </v-card-title>
            <v-card-text class="card-content">
              <div class="detail-item">
                <strong>اسم المشروع:</strong>
                <span>{{ project.name }}</span>
              </div>
              <div class="detail-item">
                <strong>الموقع:</strong>
                <span>{{ project.location }}</span>
              </div>
              <div class="detail-item">
                <strong>المسؤول:</strong>
                <span>{{ project.manager }}</span>
              </div>
              <div class="detail-item">
                <strong>فئة المشروع:</strong>
                <v-chip color="primary" variant="tonal">{{ project.category }}</v-chip>
              </div>
            </v-card-text>
          </v-card>
        </v-col>

        <!-- Status and Priority Card -->
        <v-col cols="12" md="6">
          <v-card class="detail-card status-card" elevation="3">
            <v-card-title class="card-header status-header">
              <v-icon class="me-2">mdi-account</v-icon>
              الحالة والأولوية
            </v-card-title>
            <v-card-text class="card-content">
              <div class="detail-item">
                <strong>الحالة:</strong>
                <v-chip :color="getStatusColor(project.status)" variant="tonal">
                  {{ getStatusText(project.status) }}
                </v-chip>
              </div>
              <div class="detail-item">
                <strong>الأولوية:</strong>
                <v-chip :color="getPriorityColor(project.priority)" variant="tonal">
                  {{ getPriorityText(project.priority) }}
                </v-chip>
              </div>
              <div class="detail-item">
                <strong>تاريخ البدء:</strong>
                <span>{{ project.startDate }}</span>
              </div>
            </v-card-text>
          </v-card>
        </v-col>

        <!-- Work Days Card -->
        <v-col cols="12" md="6">
          <v-card class="detail-card work-days-card" elevation="3">
            <v-card-title class="card-header work-days-header">
              <v-icon class="me-2">mdi-calendar</v-icon>
              أيام العمل
              <v-spacer />
              <v-btn
                icon="mdi-pencil"
                size="small"
                variant="text"
                color="white"
                @click="editWorkDays"
                class="edit-days-btn"
              >
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
            </v-card-title>
            <v-card-text class="card-content">
              <div v-for="day in workDays" :key="day.name" class="work-day-item">
                <div class="day-info">
                  <v-icon :color="getDayIconColor(day)" class="me-2">
                    {{ getDayIcon(day) }}
                  </v-icon>
                  <span class="day-name">{{ day.name }}</span>
                </div>
                <span class="day-hours">{{ day.hours }} ساعات</span>
              </div>
              <v-btn
                block
                variant="outlined"
                color="success"
                @click="editWorkDays"
                class="mt-3"
              >
                <v-icon class="me-2">mdi-calendar-clock</v-icon>
                إدارة أيام العمل
              </v-btn>
            </v-card-text>
          </v-card>
        </v-col>

        <!-- Financial Information Card -->
        <v-col cols="12" md="6">
          <v-card class="detail-card financial-card" elevation="3">
            <v-card-title class="card-header financial-header">
              <v-icon class="me-2">mdi-chart-line</v-icon>
              المعلومات المالية
            </v-card-title>
            <v-card-text class="card-content">
              <div class="detail-item financial-item">
                <strong>التكلفة المبدئية:</strong>
                <span class="initial-cost">{{ formatCurrency(project.initialCost) }}</span>
              </div>
              <div class="detail-item financial-item">
                <strong>التكلفة الحرجة:</strong>
                <span class="critical-cost">{{ formatCurrency(project.criticalCost) }}</span>
              </div>
              <div class="detail-item financial-item">
                <strong>الفرق:</strong>
                <span class="difference">{{ formatCurrency(project.criticalCost - project.initialCost) }}</span>
              </div>
            </v-card-text>
          </v-card>
        </v-col>

        <!-- Project Description Card -->
        <v-col cols="12" v-if="project.description">
          <v-card class="detail-card description-card" elevation="3">
            <v-card-title class="card-header description-header">
              <v-icon class="me-2">mdi-text</v-icon>
              وصف المشروع
            </v-card-title>
            <v-card-text class="card-content">
              <p class="description-text">{{ project.description }}</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Loading State -->
      <v-row v-else class="justify-center">
        <v-col cols="12" class="text-center">
          <v-progress-circular indeterminate color="primary" size="64"></v-progress-circular>
          <p class="mt-4">جاري تحميل تفاصيل المشروع...</p>
        </v-col>
      </v-row>

      <!-- Actions -->
      <v-row v-if="project" class="mt-6">
        <v-col cols="12" class="text-center">
          <v-btn 
            color="primary" 
            size="large" 
            @click="editProject"
            class="me-4"
          >
            <v-icon class="me-2">mdi-pencil</v-icon>
            تعديل المشروع
          </v-btn>
          <v-btn 
            color="error" 
            size="large" 
            @click="deleteProject"
            variant="outlined"
          >
            <v-icon class="me-2">mdi-delete</v-icon>
            حذف المشروع
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const project = ref(null)
const projectId = computed(() => route.params.id)

// Work days data
const workDays = ref([
  { name: 'السبت', hours: 8, status: 'active' },
  { name: 'الأحد', hours: 8, status: 'active' },
  { name: 'الاثنين', hours: 8, status: 'active' },
  { name: 'الثلاثاء', hours: 8, status: 'active' },
  { name: 'الأربعاء', hours: 8, status: 'active' },
  { name: 'الخميس', hours: 6, status: 'partial' },
  { name: 'الجمعة', hours: 0, status: 'off' }
])

// Sample project data - in real app, this would come from API
const sampleProjects = [
  {
    id: 1,
    name: 'مشروع تطوير الموقع الإلكتروني',
    location: 'الرياض',
    manager: 'أحمد محمد',
    category: 'تطوير',
    status: 'active',
    priority: 'high',
    startDate: '15/01/2024',
    initialCost: 50000,
    criticalCost: 75000,
    description: 'مشروع تطوير موقع إلكتروني متكامل للمؤسسة مع نظام إدارة المحتوى والمدفوعات الإلكترونية.'
  }
]

// Methods
const loadProject = () => {
  // In real app, fetch from API using projectId
  const foundProject = sampleProjects.find(p => p.id === parseInt(projectId.value))
  if (foundProject) {
    project.value = foundProject
  } else {
    // Redirect to projects list if project not found
    router.push('/project-management')
  }
}

const goBack = () => {
  router.push('/project-management')
}

const formatCurrency = (amount) => {
  if (!amount) return '0 د.ع'
  return `${amount.toLocaleString()} د.ع`
}

const getStatusColor = (status) => {
  const colors = {
    active: 'success',
    pending: 'warning',
    completed: 'info',
    cancelled: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    active: 'نشط',
    pending: 'في الانتظار',
    completed: 'مكتمل',
    cancelled: 'ملغي'
  }
  return texts[status] || status
}

const getPriorityColor = (priority) => {
  const colors = {
    high: 'error',
    medium: 'warning',
    low: 'success'
  }
  return colors[priority] || 'grey'
}

const getPriorityText = (priority) => {
  const texts = {
    high: 'عالية',
    medium: 'متوسطة',
    low: 'منخفضة'
  }
  return texts[priority] || priority
}

const getDayIcon = (day) => {
  const icons = {
    active: 'mdi-check-circle',
    partial: 'mdi-clock-outline',
    off: 'mdi-close-circle'
  }
  return icons[day.status] || 'mdi-help-circle'
}

const getDayIconColor = (day) => {
  const colors = {
    active: 'success',
    partial: 'warning',
    off: 'grey'
  }
  return colors[day.status] || 'grey'
}

const editProject = () => {
  // Navigate to edit project page
  router.push(`/project-management?edit=${projectId.value}`)
}

const editWorkDays = () => {
  // Navigate to work days management page
  router.push('/work-days')
}

const deleteProject = () => {
  // Implement delete functionality
  // After deletion, redirect to projects list
  router.push('/project-management')
}

onMounted(() => {
  loadProject()
})
</script>


<style>
/* Import page styles */
@import './styles/project-details.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
