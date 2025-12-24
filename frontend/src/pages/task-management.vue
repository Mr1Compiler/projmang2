<template>
  <v-container fluid class="task-management-page">
    <!-- Header Section -->
    <div class="engineers-header-card">
      <div class="header-gradient-line"></div>
      <div class="header-content">
        <div class="header-right">
          <div class="engineer-emoji">
            <v-icon size="40" color="white">mdi-clipboard-list</v-icon>
          </div>
          <div class="header-text">
            <h1 class="main-title">إدارة المهام</h1>
            <p class="subtitle">إدارة شاملة للمهام والمشاريع</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Statistics Cards -->
    <div class="stats-section">
      <v-row>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-clipboard-list</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ totalTasks }}</h3>
                <p class="stat-label">إجمالي المهام</p>
              </div>
            </div>
          </v-card>
        </v-col>
        
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ completedTasks }}</h3>
                <p class="stat-label">مهام مكتملة</p>
              </div>
            </div>
          </v-card>
        </v-col>
        
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-clock-outline</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ inProgressTasks }}</h3>
                <p class="stat-label">قيد التنفيذ</p>
              </div>
            </div>
          </v-card>
        </v-col>
        
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-info" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-alert-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ pendingTasks }}</h3>
                <p class="stat-label">مهام معلقة</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <!-- Search and Filter Section -->
    <v-card class="search-filter-card">
      <v-card-text>
        <div class="search-filter-content">
          <v-row>
            <v-col cols="12" md="4">
              <v-text-field
                v-model="searchQuery"
                label="البحث في المهام..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                clearable
                class="search-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="selectedStatus"
                :items="statusOptions"
                label="الحالة"
                variant="outlined"
                clearable
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="selectedPriority"
                :items="priorityOptions"
                label="الأولوية"
                variant="outlined"
                clearable
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-btn
                color="primary"
                variant="elevated"
                @click="showAddTaskDialog = true"
                class="add-btn"
              >
                <v-icon start>mdi-plus</v-icon>
                إضافة مهمة
              </v-btn>
            </v-col>
          </v-row>
        </div>
      </v-card-text>
    </v-card>

    <!-- Tasks Table -->
    <v-card class="tasks-table-card">
      <v-card-title class="table-title">
        <v-icon size="28" color="primary" class="mr-2">mdi-clipboard-list</v-icon>
        <span class="text-h5 font-weight-bold">قائمة المهام</span>
        <v-spacer />
        <v-chip color="primary" variant="elevated" size="large">
          {{ filteredTasks.length }} مهمة
        </v-chip>
      </v-card-title>
      
      <v-data-table
        :headers="headers"
        :items="filteredTasks"
        :items-per-page="10"
        class="tasks-data-table"
        :loading="loading"
      >
        <!-- Status Column -->
        <template #item.status="{ item }">
          <v-chip
            :color="getStatusColor(item.status)"
            variant="elevated"
            size="small"
          >
            <v-icon start>{{ getStatusIcon(item.status) }}</v-icon>
            {{ item.status }}
          </v-chip>
        </template>

        <!-- Priority Column -->
        <template #item.priority="{ item }">
          <v-chip
            :color="getPriorityColor(item.priority)"
            variant="elevated"
            size="small"
          >
            {{ item.priority }}
          </v-chip>
        </template>

        <!-- Actions Column -->
        <template #item.actions="{ item }">
          <div class="action-buttons">
            <v-btn
              icon="mdi-eye"
              size="small"
              color="info"
              variant="elevated"
              @click="viewTask(item)"
              class="action-btn view-btn"
            />
            <v-btn
              v-if="canUpdate"
              icon="mdi-pencil"
              size="small"
              color="warning"
              variant="elevated"
              @click="editTask(item)"
              class="action-btn edit-btn"
            />
            <v-btn
              v-if="canDelete"
              icon="mdi-delete"
              size="small"
              color="error"
              variant="elevated"
              @click="deleteTask(item)"
              class="action-btn delete-btn"
            />
          </div>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Task Dialog -->
    <v-dialog v-model="showAddTaskDialog" max-width="600px" persistent>
      <v-card class="add-task-dialog image-style-dialog">
        <v-card-title class="dialog-header">
          <div class="header-content">
            <div class="header-left">
              <v-icon size="32" color="white" class="mr-2">mdi-plus-circle</v-icon>
              <span class="dialog-title">إضافة مهمة جديدة</span>
            </div>
          </div>
        </v-card-title>
        
        <v-card-text class="dialog-body">
          <v-form ref="addTaskForm" v-model="addFormValid">
            <v-row class="form-fields">
              <v-col cols="12">
                <div class="form-field">
                  <v-text-field
                    v-model="newTask.title"
                    label="عنوان المهمة"
                    variant="outlined"
                    :rules="[v => !!v || 'عنوان المهمة مطلوب']"
                    required
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-select
                    v-model="newTask.status"
                    :items="statusOptions"
                    label="الحالة"
                    variant="outlined"
                    :rules="[v => !!v || 'الحالة مطلوبة']"
                    required
                    class="black-dropdown-select"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-select
                    v-model="newTask.priority"
                    :items="priorityOptions"
                    label="الأولوية"
                    variant="outlined"
                    :rules="[v => !!v || 'الأولوية مطلوبة']"
                    required
                    class="black-dropdown-select"
                  />
                </div>
              </v-col>
              <v-col cols="12">
                <div class="form-field">
                  <v-textarea
                    v-model="newTask.description"
                    label="وصف المهمة"
                    variant="outlined"
                    rows="3"
                  />
                </div>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        
        <v-card-actions class="dialog-actions">
          <v-spacer />
          <v-btn
            color="grey"
            variant="text"
            @click="closeAddTaskDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="primary"
            variant="elevated"
            :disabled="!addFormValid"
            :loading="saving"
            @click="saveNewTask"
          >
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { usePermissions } from '@/composables/usePermissions'

// Permissions - matches database page route
const { canCreate, canUpdate, canDelete } = usePermissions('/projects')

// Reactive data
const loading = ref(false)
const saving = ref(false)
const searchQuery = ref('')
const selectedStatus = ref(null)
const selectedPriority = ref(null)
const showAddTaskDialog = ref(false)
const addFormValid = ref(false)

// Form data
const newTask = ref({
  title: '',
  description: '',
  status: 'معلقة',
  priority: 'متوسطة'
})

// Sample data
const tasks = ref([
  {
    id: 1,
    title: 'تصميم واجهة المستخدم',
    description: 'تصميم واجهة المستخدم للمشروع الجديد',
    status: 'قيد التنفيذ',
    priority: 'عالية',
    assignee: 'أحمد محمد',
    dueDate: '2024-12-31'
  },
  {
    id: 2,
    title: 'مراجعة الكود',
    description: 'مراجعة الكود قبل النشر',
    status: 'مكتملة',
    priority: 'عالية',
    assignee: 'فاطمة حسن',
    dueDate: '2024-12-25'
  },
  {
    id: 3,
    title: 'كتابة الوثائق',
    description: 'كتابة وثائق المشروع',
    status: 'معلقة',
    priority: 'متوسطة',
    assignee: 'محمد عبدالله',
    dueDate: '2025-01-15'
  }
])

// Options
const statusOptions = [
  'معلقة',
  'قيد التنفيذ',
  'مكتملة',
  'ملغاة'
]

const priorityOptions = [
  'منخفضة',
  'متوسطة',
  'عالية',
  'عاجلة'
]

// Table headers
const headers = [
  { 
    title: 'العنوان', 
    key: 'title', 
    sortable: true
  },
  { 
    title: 'الحالة', 
    key: 'status', 
    sortable: true
  },
  { 
    title: 'الأولوية', 
    key: 'priority', 
    sortable: true
  },
  { 
    title: 'المسؤول', 
    key: 'assignee', 
    sortable: true
  },
  { 
    title: 'تاريخ الاستحقاق', 
    key: 'dueDate', 
    sortable: true
  },
  { 
    title: 'الإجراءات', 
    key: 'actions', 
    sortable: false, 
    width: '200px'
  }
]

// Computed properties
const totalTasks = computed(() => tasks.value.length)
const completedTasks = computed(() => tasks.value.filter(task => task.status === 'مكتملة').length)
const inProgressTasks = computed(() => tasks.value.filter(task => task.status === 'قيد التنفيذ').length)
const pendingTasks = computed(() => tasks.value.filter(task => task.status === 'معلقة').length)

const filteredTasks = computed(() => {
  let filtered = tasks.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(task =>
      task.title.toLowerCase().includes(query) ||
      task.description.toLowerCase().includes(query)
    )
  }

  if (selectedStatus.value) {
    filtered = filtered.filter(task => task.status === selectedStatus.value)
  }

  if (selectedPriority.value) {
    filtered = filtered.filter(task => task.priority === selectedPriority.value)
  }

  return filtered
})

// Methods
const getStatusColor = (status) => {
  const colors = {
    'مكتملة': 'success',
    'قيد التنفيذ': 'warning',
    'معلقة': 'info',
    'ملغاة': 'error'
  }
  return colors[status] || 'grey'
}

const getStatusIcon = (status) => {
  const icons = {
    'مكتملة': 'mdi-check-circle',
    'قيد التنفيذ': 'mdi-clock-outline',
    'معلقة': 'mdi-alert-circle',
    'ملغاة': 'mdi-cancel'
  }
  return icons[status] || 'mdi-help-circle'
}

const getPriorityColor = (priority) => {
  const colors = {
    'عاجلة': 'error',
    'عالية': 'warning',
    'متوسطة': 'info',
    'منخفضة': 'success'
  }
  return colors[priority] || 'grey'
}

const viewTask = (task) => {
  console.log('View task:', task)
}

const editTask = (task) => {
  console.log('Edit task:', task)
}

const deleteTask = (task) => {
  console.log('Delete task:', task)
}

const closeAddTaskDialog = () => {
  showAddTaskDialog.value = false
  newTask.value = {
    title: '',
    description: '',
    status: 'معلقة',
    priority: 'متوسطة'
  }
}

const saveNewTask = async () => {
  if (!addFormValid.value) return

  saving.value = true
  
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 1000))
  
  const newId = Math.max(...tasks.value.map(t => t.id)) + 1
  const task = {
    id: newId,
    ...newTask.value,
    assignee: 'غير محدد',
    dueDate: new Date().toISOString().split('T')[0]
  }
  
  tasks.value.push(task)
  closeAddTaskDialog()
  saving.value = false
}

onMounted(() => {
  // Initialize data if needed
})
</script>

<style scoped>
/* Main Page Styles */
.task-management-page {
  background: #ffffff !important;
  min-height: 100vh;
  padding: 0;
  overflow-x: hidden;
}

/* Header Styles */
.engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  border-radius: 0;
  width: 100vw;
  max-width: 100vw;
  box-shadow: 0 8px 32px rgba(25, 118, 210, 0.3);
  position: relative;
  overflow: hidden;
  margin-left: calc(-50vw + 50%);
  margin-right: calc(-50vw + 50%);
  margin-bottom: 1.5rem;
  border: none;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: slideInFromTop 1s ease-out, shimmer 3s ease-in-out infinite;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.3);
  animation: gradientFlow 3s ease-in-out infinite;
  z-index: 2;
}

.header-content {
  padding: 12px 16px !important;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  min-height: auto !important;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.1) 0%, rgba(21, 101, 192, 0.05) 100%);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 3;
  animation: fadeInUp 1.2s ease-out 0.3s both;
  max-width: calc(100vw - 320px);
  margin: 0 auto;
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
  animation: slideInFromRight 1s ease-out 0.9s both, float 3s ease-in-out infinite 2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.main-title {
  color: white !important;
  font-size: 1.2rem !important;
  font-weight: bold !important;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 0.75rem !important;
  margin: 0;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.stats-section {
  margin-bottom: 30px;
  padding: 0 20px;
}

/* Modern Statistics Cards */
.modern-stat-card {
  position: relative !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  cursor: pointer !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  height: 100% !important;
  min-height: 140px !important;
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
  background: #ffffff !important;
  border: 2px solid #3b82f6 !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15) !important;
}

.stat-card-success {
  background: #ffffff !important;
  border: 2px solid #10b981 !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.15) !important;
}

.stat-card-warning {
  background: #ffffff !important;
  border: 2px solid #f59e0b !important;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.15) !important;
}

.stat-card-info {
  background: #ffffff !important;
  border: 2px solid #06b6d4 !important;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.15) !important;
}

.stat-card-content {
  position: relative;
  z-index: 2;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  align-items: center;
  text-align: center;
}

.stat-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 64px;
  height: 64px;
  min-width: 64px;
  min-height: 64px;
  border-radius: 50%;
  margin-bottom: 0.25rem;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  position: relative;
}

.stat-card-primary .stat-icon-wrapper {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.stat-card-success .stat-icon-wrapper {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.stat-card-warning .stat-icon-wrapper {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.stat-card-info .stat-icon-wrapper {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.stat-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  position: relative;
  z-index: 2;
  color: white !important;
}

.stat-card-primary .stat-icon {
  color: #3b82f6 !important;
}

.stat-card-success .stat-icon {
  color: #10b981 !important;
}

.stat-card-warning .stat-icon {
  color: #f59e0b !important;
}

.stat-card-info .stat-icon {
  color: #06b6d4 !important;
}

.stat-value {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 0.5rem;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  text-align: center;
  color: #000000 !important;
}

.stat-label {
  font-size: 1rem;
  font-weight: 500;
  text-align: center;
  color: #64748b;
}

/* Search and Filter */
.search-filter-card {
  border-radius: 16px;
  margin-bottom: 2rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.search-filter-content {
  padding: 1rem 0;
}

.add-btn {
  height: 56px;
  border-radius: 12px;
  font-weight: 600;
  text-transform: none;
}

/* Tasks Table */
.tasks-table-card {
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.table-title {
  background: linear-gradient(135deg, rgba(5, 150, 105, 0.1) 0%, rgba(16, 185, 129, 0.1) 100%);
  border-bottom: 2px solid rgba(5, 150, 105, 0.2);
  padding: 1.5rem;
}

.tasks-data-table {
  background: transparent;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
  flex-wrap: wrap;
}

.action-btn {
  border-radius: 8px;
  font-weight: 600;
  text-transform: none;
  min-width: 36px;
  height: 36px;
}

.view-btn {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
}

.edit-btn {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
}

.delete-btn {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
}

/* Dialog Styles */
.add-task-dialog.image-style-dialog {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%);
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(5, 150, 105, 0.3);
}

.add-task-dialog .dialog-header {
  background: linear-gradient(135deg, #059669, #10b981);
  padding: 20px;
  color: white;
}

.add-task-dialog .dialog-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: white;
}

.add-task-dialog .dialog-body {
  padding: 30px;
  background: white !important;
}

.add-task-dialog .dialog-actions {
  padding: 20px 30px;
  background: #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Animations */
@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-12px);
  }
}

@keyframes shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
}

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

@keyframes gradientFlow {
  0%, 100% {
    background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  }
  50% {
    background: linear-gradient(90deg, #bbdefb 0%, #ffffff 50%, #e3f2fd 100%);
  }
}
</style>
