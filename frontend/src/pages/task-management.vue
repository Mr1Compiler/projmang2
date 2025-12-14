<template>
  <div class="task-management-page">
    <!-- Header Section -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-left">
          <v-icon size="40" color="primary" class="mr-3">mdi-clipboard-list</v-icon>
          <div>
            <h1 class="page-title">إدارة المهام</h1>
            <p class="page-subtitle">نظام شامل لإدارة المهام والمشاريع</p>
          </div>
        </div>
        <div class="header-actions">
          <v-btn
            color="primary"
            variant="elevated"
            size="large"
            prepend-icon="mdi-plus"
            @click="showAddTaskDialog = true"
          >
            إضافة مهمة جديدة
          </v-btn>
        </div>
      </div>
    </div>

    <!-- Statistics Cards -->
    <div class="stats-section">
      <v-row>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="40" color="primary" class="mb-2">mdi-clipboard-list</v-icon>
              <h3 class="text-h4 font-weight-bold">{{ totalTasks }}</h3>
              <p class="text-caption">إجمالي المهام</p>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="40" color="success" class="mb-2">mdi-check-circle</v-icon>
              <h3 class="text-h4 font-weight-bold stat-number">{{ completedTasks }}</h3>
              <p class="text-caption">مهام مكتملة</p>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="40" color="warning" class="mb-2">mdi-clock</v-icon>
              <h3 class="text-h4 font-weight-bold stat-number">{{ pendingTasks }}</h3>
              <p class="text-caption">مهام معلقة</p>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="40" color="error" class="mb-2">mdi-alert-circle</v-icon>
              <h3 class="text-h4 font-weight-bold stat-number">{{ overdueTasks }}</h3>
              <p class="text-caption">مهام متأخرة</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <!-- Filters Section -->
    <div class="filters-section">
      <v-card class="filters-card">
        <v-card-title class="filters-header">
          <v-icon size="32" color="white" class="mr-2">mdi-filter</v-icon>
          <span class="text-h4 font-weight-black" style="color: #ffffff; font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif; text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5), 0 1px 3px rgba(0, 0, 0, 0.3); letter-spacing: 1px; font-size: 1.6rem;">فلترة المهام</span>
        </v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="12" md="3">
              <v-text-field
                v-model="searchQuery"
                label="البحث في المهام"
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                clearable
                density="comfortable"
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="statusFilter"
                :items="statusOptions"
                label="حالة المهمة"
                variant="outlined"
                density="comfortable"
                clearable
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="priorityFilter"
                :items="priorityOptions"
                label="أولوية المهمة"
                variant="outlined"
                density="comfortable"
                clearable
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="assigneeFilter"
                :items="assigneeOptions"
                label="المسؤول"
                variant="outlined"
                density="comfortable"
                clearable
                class="filter-field"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </div>

    <!-- Tasks Table -->
    <div class="tasks-section">
      <v-card class="tasks-card">
        <v-card-title class="tasks-header">
          <v-icon size="18" color="white" class="mr-2">mdi-clipboard-list</v-icon>
          <span class="text-h4 font-weight-black" style="color: #ffffff; font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif; text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3); letter-spacing: 0.2px; font-size: 0.9rem;">قائمة المهام</span>
          <v-spacer />
          <v-btn
            color="primary"
            variant="outlined"
            size="small"
            prepend-icon="mdi-export"
            @click="exportTasks"
          >
            تصدير
          </v-btn>
        </v-card-title>
        
        <v-data-table
          :headers="taskHeaders"
          :items="filteredTasks"
          :items-per-page="10"
          class="tasks-table"
          :loading="tasksLoading"
        >
          <!-- Task ID Column -->
          <template #item.id="{ item }">
            <v-chip
              color="primary"
              variant="elevated"
              size="small"
            >
              #{{ item.id }}
            </v-chip>
          </template>

          <!-- Task Title Column -->
          <template #item.title="{ item }">
            <div class="task-title">
              <h4 class="task-name">{{ item.title }}</h4>
              <p class="task-description">{{ item.description }}</p>
            </div>
          </template>

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

          <!-- Assignee Column -->
          <template #item.assignee="{ item }">
            <div class="assignee-info">
              <v-avatar size="18" class="mr-1">
                <v-img v-if="item.assigneeAvatar" :src="item.assigneeAvatar" />
                <v-icon v-else size="12" color="primary">mdi-account</v-icon>
              </v-avatar>
              <span>{{ item.assignee }}</span>
            </div>
          </template>

          <!-- Due Date Column -->
          <template #item.dueDate="{ item }">
            <div class="due-date">
              <v-icon size="12" color="primary" class="mr-1">mdi-calendar</v-icon>
              <span>{{ formatDate(item.dueDate) }}</span>
            </div>
          </template>

          <!-- Progress Column -->
          <template #item.progress="{ item }">
            <div class="progress-container">
              <v-progress-linear
                :model-value="item.progress"
                :color="getProgressColor(item.progress)"
                height="8"
                rounded
                class="mb-1"
              />
              <span class="progress-text">{{ item.progress }}%</span>
            </div>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="action-buttons">
              <v-btn
                icon="mdi-eye"
                size="x-small"
                color="info"
                variant="elevated"
                @click="viewTask(item)"
              />
              <v-btn
                icon="mdi-pencil"
                size="x-small"
                color="warning"
                variant="elevated"
                @click="editTask(item)"
              />
              <v-btn
                icon="mdi-delete"
                size="x-small"
                color="error"
                variant="elevated"
                @click="deleteTask(item)"
              />
            </div>
          </template>
        </v-data-table>
      </v-card>
    </div>

    <!-- Add/Edit Task Dialog -->
    <v-dialog v-model="showAddTaskDialog" max-width="800">
      <v-card class="task-dialog-card">
        <v-card-title class="task-dialog-header">
          <v-icon size="28" color="white" class="mr-2">mdi-plus</v-icon>
          <span class="text-h5">{{ isEditing ? 'تعديل المهمة' : 'إضافة مهمة جديدة' }}</span>
        </v-card-title>
        
        <v-card-text>
          <v-form ref="taskForm" v-model="taskFormValid">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="taskForm.title"
                  label="عنوان المهمة"
                  variant="outlined"
                  :rules="[v => !!v || 'عنوان المهمة مطلوب']"
                  required
                  class="task-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="taskForm.status"
                  :items="statusOptions"
                  label="حالة المهمة"
                  variant="outlined"
                  required
                  class="task-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  v-model="taskForm.description"
                  label="وصف المهمة"
                  variant="outlined"
                  rows="3"
                  class="task-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="taskForm.priority"
                  :items="priorityOptions"
                  label="أولوية المهمة"
                  variant="outlined"
                  required
                  class="task-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="taskForm.assignee"
                  :items="assigneeOptions"
                  label="المسؤول"
                  variant="outlined"
                  required
                  class="task-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="taskForm.dueDate"
                  label="تاريخ الاستحقاق"
                  type="date"
                  variant="outlined"
                  required
                  class="task-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-slider
                  v-model="taskForm.progress"
                  label="نسبة الإنجاز"
                  min="0"
                  max="100"
                  step="5"
                  thumb-label
                  color="primary"
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        
        <v-card-actions class="dialog-actions">
          <v-spacer />
          <v-btn
            color="grey"
            variant="outlined"
            @click="closeTaskDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="primary"
            variant="elevated"
            :disabled="!taskFormValid"
            @click="saveTask"
          >
            {{ isEditing ? 'تحديث' : 'إضافة' }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Task Details Dialog -->
    <v-dialog v-model="showTaskDetails" max-width="600">
      <v-card>
        <v-card-title class="dialog-header">
          <v-icon size="24" color="primary" class="mr-2">mdi-eye</v-icon>
          <span class="text-h6 font-weight-bold">تفاصيل المهمة</span>
        </v-card-title>
        
        <v-card-text v-if="selectedTask">
          <div class="task-details">
            <div class="detail-row">
              <strong>عنوان المهمة:</strong>
              <span>{{ selectedTask.title }}</span>
            </div>
            <div class="detail-row">
              <strong>الوصف:</strong>
              <span>{{ selectedTask.description }}</span>
            </div>
            <div class="detail-row">
              <strong>الحالة:</strong>
              <v-chip
                :color="getStatusColor(selectedTask.status)"
                variant="elevated"
                size="small"
              >
                {{ selectedTask.status }}
              </v-chip>
            </div>
            <div class="detail-row">
              <strong>الأولوية:</strong>
              <v-chip
                :color="getPriorityColor(selectedTask.priority)"
                variant="elevated"
                size="small"
              >
                {{ selectedTask.priority }}
              </v-chip>
            </div>
            <div class="detail-row">
              <strong>المسؤول:</strong>
              <span>{{ selectedTask.assignee }}</span>
            </div>
            <div class="detail-row">
              <strong>تاريخ الاستحقاق:</strong>
              <span>{{ formatDate(selectedTask.dueDate) }}</span>
            </div>
            <div class="detail-row">
              <strong>نسبة الإنجاز:</strong>
              <div class="progress-container">
                <v-progress-linear
                  :model-value="selectedTask.progress"
                  :color="getProgressColor(selectedTask.progress)"
                  height="8"
                  rounded
                  class="mb-1"
                />
                <span class="progress-text">{{ selectedTask.progress }}%</span>
              </div>
            </div>
          </div>
        </v-card-text>
        
        <v-card-actions class="dialog-actions">
          <v-spacer />
          <v-btn
            color="primary"
            variant="elevated"
            @click="showTaskDetails = false"
          >
            إغلاق
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

// Reactive data
const tasksLoading = ref(false)
const showAddTaskDialog = ref(false)
const showTaskDetails = ref(false)
const isEditing = ref(false)
const taskFormValid = ref(false)
const selectedTask = ref(null)

// Search and filters
const searchQuery = ref('')
const statusFilter = ref('')
const priorityFilter = ref('')
const assigneeFilter = ref('')

// Task form
const taskForm = ref({
  title: '',
  description: '',
  status: 'معلقة',
  priority: 'متوسطة',
  assignee: '',
  dueDate: '',
  progress: 0
})

// Task data
const tasks = ref([
  {
    id: 1,
    title: 'تطوير واجهة المستخدم',
    description: 'تصميم وتطوير واجهة المستخدم للتطبيق الجديد',
    status: 'قيد التنفيذ',
    priority: 'عالية',
    assignee: 'أحمد محمد',
    assigneeAvatar: null,
    dueDate: '2024-01-15',
    progress: 75
  },
  {
    id: 2,
    title: 'اختبار النظام',
    description: 'إجراء اختبارات شاملة للنظام',
    status: 'معلقة',
    priority: 'متوسطة',
    assignee: 'فاطمة علي',
    assigneeAvatar: null,
    dueDate: '2024-01-20',
    progress: 30
  },
  {
    id: 3,
    title: 'توثيق المشروع',
    description: 'كتابة التوثيق التقني للمشروع',
    status: 'مكتملة',
    priority: 'منخفضة',
    assignee: 'محمد حسن',
    assigneeAvatar: null,
    dueDate: '2024-01-10',
    progress: 100
  },
  {
    id: 4,
    title: 'إصلاح الأخطاء',
    description: 'إصلاح الأخطاء المكتشفة في النظام',
    status: 'متأخرة',
    priority: 'عالية',
    assignee: 'سارة أحمد',
    assigneeAvatar: null,
    dueDate: '2024-01-05',
    progress: 60
  }
])

// Table headers
const taskHeaders = ref([
  { title: 'رقم المهمة', key: 'id', sortable: true },
  { title: 'عنوان المهمة', key: 'title', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الأولوية', key: 'priority', sortable: true },
  { title: 'المسؤول', key: 'assignee', sortable: true },
  { title: 'تاريخ الاستحقاق', key: 'dueDate', sortable: true },
  { title: 'نسبة الإنجاز', key: 'progress', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
])

// Filter options
const statusOptions = ref([
  'معلقة',
  'قيد التنفيذ',
  'مكتملة',
  'متأخرة'
])

const priorityOptions = ref([
  'منخفضة',
  'متوسطة',
  'عالية',
  'حرجة'
])

const assigneeOptions = ref([
  'أحمد محمد',
  'فاطمة علي',
  'محمد حسن',
  'سارة أحمد'
])

// Computed properties
const totalTasks = computed(() => tasks.value.length)
const completedTasks = computed(() => tasks.value.filter(task => task.status === 'مكتملة').length)
const pendingTasks = computed(() => tasks.value.filter(task => task.status === 'معلقة').length)
const overdueTasks = computed(() => tasks.value.filter(task => task.status === 'متأخرة').length)

const filteredTasks = computed(() => {
  let filtered = tasks.value

  if (searchQuery.value) {
    filtered = filtered.filter(task =>
      task.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      task.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (statusFilter.value) {
    filtered = filtered.filter(task => task.status === statusFilter.value)
  }

  if (priorityFilter.value) {
    filtered = filtered.filter(task => task.priority === priorityFilter.value)
  }

  if (assigneeFilter.value) {
    filtered = filtered.filter(task => task.assignee === assigneeFilter.value)
  }

  return filtered
})

// Methods
const getStatusColor = (status) => {
  const colors = {
    'معلقة': 'warning',
    'قيد التنفيذ': 'info',
    'مكتملة': 'success',
    'متأخرة': 'error'
  }
  return colors[status] || 'grey'
}

const getStatusIcon = (status) => {
  const icons = {
    'معلقة': 'mdi-clock',
    'قيد التنفيذ': 'mdi-play',
    'مكتملة': 'mdi-check',
    'متأخرة': 'mdi-alert'
  }
  return icons[status] || 'mdi-help'
}

const getPriorityColor = (priority) => {
  const colors = {
    'منخفضة': 'success',
    'متوسطة': 'warning',
    'عالية': 'error',
    'حرجة': 'error'
  }
  return colors[priority] || 'grey'
}

const getProgressColor = (progress) => {
  if (progress < 30) return 'error'
  if (progress < 70) return 'warning'
  return 'success'
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('ar-SA')
}

const viewTask = (task) => {
  selectedTask.value = task
  showTaskDetails.value = true
}

const editTask = (task) => {
  selectedTask.value = task
  taskForm.value = { ...task }
  isEditing.value = true
  showAddTaskDialog.value = true
}

const deleteTask = (task) => {
  if (confirm('هل أنت متأكد من حذف هذه المهمة؟')) {
    const index = tasks.value.findIndex(t => t.id === task.id)
    if (index > -1) {
      tasks.value.splice(index, 1)
    }
  }
}

const saveTask = () => {
  if (!taskFormValid.value) {
    return
  }
  
  if (isEditing.value) {
    const index = tasks.value.findIndex(t => t.id === selectedTask.value.id)
    if (index > -1) {
      tasks.value[index] = { 
        ...taskForm.value,
        id: selectedTask.value.id,
        assigneeAvatar: selectedTask.value.assigneeAvatar || null
      }
    }
  } else {
    // إضافة مهمة جديدة في بداية القائمة
    const newTask = {
      ...taskForm.value,
      id: tasks.value.length > 0 ? Math.max(...tasks.value.map(t => t.id)) + 1 : 1,
      assigneeAvatar: null
    }
    tasks.value.unshift(newTask) // إضافة في البداية بدلاً من النهاية
  }
  
  // إعادة تعيين الفلاتر لإظهار المهمة الجديدة
  searchQuery.value = ''
  statusFilter.value = ''
  priorityFilter.value = ''
  assigneeFilter.value = ''
  
  closeTaskDialog()
}

const closeTaskDialog = () => {
  showAddTaskDialog.value = false
  isEditing.value = false
  selectedTask.value = null
  taskForm.value = {
    title: '',
    description: '',
    status: 'معلقة',
    priority: 'متوسطة',
    assignee: '',
    dueDate: '',
    progress: 0
  }
}

const exportTasks = () => {
  // Export functionality
  console.log('Exporting tasks...')
}

onMounted(() => {
  // Initialize data
  tasksLoading.value = false
})
</script>

<style scoped>
.task-management-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.page-header {
  margin-bottom: 30px;
}

@keyframes headerGradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

@keyframes headerShimmer {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

@keyframes headerFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-5px);
  }
}

@keyframes headerPulse {
  0%, 100% {
    box-shadow: 0 8px 32px rgba(102, 126, 234, 0.2);
  }
  50% {
    box-shadow: 0 12px 40px rgba(102, 126, 234, 0.4);
  }
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.3) 0%, rgba(118, 75, 162, 0.3) 50%, rgba(102, 126, 234, 0.3) 100%);
  background-size: 200% 200%;
  backdrop-filter: blur(10px);
  border-radius: 20px;
  padding: 30px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  position: relative;
  overflow: hidden;
  animation: headerGradient 8s ease infinite, headerFloat 6s ease-in-out infinite, headerPulse 4s ease-in-out infinite;
}

.header-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: headerShimmer 3s infinite;
  z-index: 1;
}

.header-content > * {
  position: relative;
  z-index: 2;
}

.header-left {
  display: flex;
  align-items: center;
}

@keyframes buttonPulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

.header-actions .v-btn {
  animation: buttonPulse 3s ease-in-out infinite;
  transition: all 0.3s ease;
}

.header-actions .v-btn:hover {
  animation: none;
  transform: scale(1.1) translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.4);
}

@keyframes titleGlow {
  0%, 100% {
    text-shadow: 0 0 10px rgba(255, 255, 255, 0.5), 0 0 20px rgba(102, 126, 234, 0.3);
  }
  50% {
    text-shadow: 0 0 20px rgba(255, 255, 255, 0.8), 0 0 30px rgba(102, 126, 234, 0.5);
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0) rotate(0deg);
  }
  50% {
    transform: translateY(-5px) rotate(5deg);
  }
}

.page-title {
  color: white;
  font-size: 2.5rem;
  font-weight: bold;
  margin: 0;
  animation: titleGlow 3s ease-in-out infinite;
  transition: transform 0.3s ease;
}

.page-title:hover {
  transform: scale(1.05);
}

.header-left .v-icon {
  animation: iconBounce 4s ease-in-out infinite;
}

.page-subtitle {
  color: rgba(255, 255, 255, 0.8);
  font-size: 1.1rem;
  margin: 5px 0 0 0;
}

.stats-section {
  margin-bottom: 30px;
}

.stat-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  backdrop-filter: blur(10px);
  border-radius: 16px !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.1) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  overflow: hidden !important;
  position: relative !important;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-6px) scale(1.02);
  box-shadow: 0 12px 32px rgba(59, 130, 246, 0.25) !important;
  border-color: rgba(59, 130, 246, 0.4) !important;
}

.stat-card:hover::before {
  opacity: 1;
}

.stat-card :deep(.v-card-text) {
  padding: 24px 20px !important;
}

.stat-card :deep(.v-icon) {
  margin-bottom: 12px !important;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.stat-card :deep(h3) {
  color: #1e293b !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  margin: 8px 0 !important;
  line-height: 1.2 !important;
}

.stat-number {
  font-size: 1.1rem !important;
  line-height: 1.2 !important;
  font-weight: 800 !important;
  text-align: center !important;
}

.stat-card :deep(.text-caption) {
  color: #64748b !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  margin-top: 4px !important;
}

.filters-section {
  margin-bottom: 30px;
}

.filters-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 50%, #f1f5f9 100%) !important;
  backdrop-filter: blur(15px) !important;
  border-radius: 20px !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.12) !important;
  overflow: hidden !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.filters-card:hover {
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.18) !important;
  border-color: rgba(59, 130, 246, 0.3) !important;
}

.filters-card :deep(.v-card-text) {
  padding: 24px !important;
}

.filters-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 24px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.3) !important;
  display: flex !important;
  align-items: center !important;
  gap: 12px !important;
}

.tasks-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  border-radius: 10px 10px 0 0 !important;
  padding: 6px 10px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.3) !important;
  display: flex !important;
  align-items: center !important;
  gap: 6px !important;
  min-height: auto !important;
}

.tasks-header span {
  font-size: 0.9rem !important;
  font-weight: 600 !important;
  letter-spacing: 0.2px !important;
}

.tasks-section {
  margin-bottom: 30px;
}

.tasks-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  backdrop-filter: blur(10px) !important;
  border-radius: 20px !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.12) !important;
  overflow: hidden !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.tasks-card:hover {
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.18) !important;
  border-color: rgba(59, 130, 246, 0.3) !important;
}

/* تحسين حقول الفلترة */
.filters-card .v-text-field,
.filters-card .v-select,
.filters-card .filter-field {
  background: rgba(255, 255, 255, 0.95) !important;
  border-radius: 12px !important;
  margin-bottom: 0 !important;
}

.filters-card .v-text-field:hover,
.filters-card .v-select:hover,
.filters-card .filter-field:hover {
  transform: translateY(-1px);
}

.filters-card .v-text-field:focus-within,
.filters-card .v-select:focus-within,
.filters-card .filter-field:focus-within {
  transform: translateY(-2px);
}

/* تحسين حقول الفلترة - تنسيق شامل */
.filter-field :deep(.v-field) {
  background-color: #ffffff !important;
  border-radius: 10px !important;
  min-height: 56px !important;
}

.filter-field :deep(.v-field__outline) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.filter-field :deep(.v-field--focused .v-field__outline) {
  border-color: #2563eb !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2) !important;
}

.filter-field :deep(.v-field__input) {
  padding: 14px 16px !important;
  min-height: 56px !important;
}

.filter-field :deep(.v-field__input input),
.filter-field :deep(.v-field__input select) {
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 0 !important;
  line-height: 1.5 !important;
}

.filter-field :deep(.v-label),
.filter-field :deep(.v-field-label) {
  color: #1e40af !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
  padding: 0 8px !important;
}

.filter-field :deep(.v-label--active),
.filter-field :deep(.v-field-label--active),
.filter-field :deep(.v-label--floating),
.filter-field :deep(.v-field-label--floating) {
  color: #1e40af !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.filter-field :deep(.v-field__prepend-inner) {
  color: #3b82f6 !important;
  padding-right: 12px !important;
}

.filter-field :deep(.v-field__prepend-inner .v-icon) {
  color: #3b82f6 !important;
  font-size: 22px !important;
}

/* تحسين أسماء الحقول بشكل أكثر تحديداً */
.filters-card .v-field-label {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
}

.filters-card .v-field-label--active {
  color: #dc2626 !important;
  font-weight: 800 !important;
}

.filters-card .v-field-label--floating {
  color: #dc2626 !important;
  font-weight: 800 !important;
}

/* تحسين شامل لأسماء الحقول */
.filters-card .v-text-field .v-label,
.filters-card .v-select .v-label,
.filters-card .v-text-field .v-field-label,
.filters-card .v-select .v-field-label,
.filters-card .v-text-field .v-field-label--active,
.filters-card .v-select .v-field-label--active,
.filters-card .v-text-field .v-field-label--floating,
.filters-card .v-select .v-field-label--floating {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* استخدام ::deep للوصول للعناصر المتداخلة */
::deep(.filters-card .v-label),
::deep(.filters-card .v-field-label),
::deep(.filters-card .v-field-label--active),
::deep(.filters-card .v-field-label--floating) {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS أقوى وأكثر تحديداً */
.filters-card .v-text-field .v-field__field .v-field-label,
.filters-card .v-select .v-field__field .v-field-label,
.filters-card .v-text-field .v-field__field .v-field-label--active,
.filters-card .v-select .v-field__field .v-field-label--active,
.filters-card .v-text-field .v-field__field .v-field-label--floating,
.filters-card .v-select .v-field__field .v-field-label--floating {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS للعناصر المباشرة */
.filters-card .v-field-label {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

.filters-card .v-field-label--active {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

.filters-card .v-field-label--floating {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS شامل لجميع العناصر */
.filters-card * .v-field-label,
.filters-card * .v-field-label--active,
.filters-card * .v-field-label--floating,
.filters-card * .v-label {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS للعناصر المباشرة */
.filters-card .v-field__field * {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS نهائي شامل - جميع العناصر */
.filters-card .v-field-label,
.filters-card .v-field-label--active,
.filters-card .v-field-label--floating,
.filters-card .v-label,
.filters-card .v-field__field .v-field-label,
.filters-card .v-field__field .v-field-label--active,
.filters-card .v-field__field .v-field-label--floating,
.filters-card .v-field__field .v-label,
.filters-card .v-text-field .v-field-label,
.filters-card .v-text-field .v-field-label--active,
.filters-card .v-text-field .v-field-label--floating,
.filters-card .v-text-field .v-label,
.filters-card .v-select .v-field-label,
.filters-card .v-select .v-field-label--active,
.filters-card .v-select .v-field-label--floating,
.filters-card .v-select .v-label {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS شامل مع ::deep */
::deep(.filters-card .v-field-label),
::deep(.filters-card .v-field-label--active),
::deep(.filters-card .v-field-label--floating),
::deep(.filters-card .v-label) {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* تحسين النصوص داخل حقول الإدخال */
.filters-card .v-field__input,
.filters-card .v-field__input input,
.filters-card .v-field__input textarea,
.filters-card .v-field__input select,
.filters-card .v-field__input option {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* تحسين النصوص مع ::deep */
::deep(.filters-card .v-field__input),
::deep(.filters-card .v-field__input input),
::deep(.filters-card .v-field__input textarea),
::deep(.filters-card .v-field__input select),
::deep(.filters-card .v-field__input option) {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* تحسين النصوص في جميع الحقول */
.filters-card .v-text-field .v-field__input,
.filters-card .v-select .v-field__input,
.filters-card .v-text-field .v-field__input input,
.filters-card .v-select .v-field__input input,
.filters-card .v-text-field .v-field__input textarea,
.filters-card .v-select .v-field__input textarea {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* تنظيم أماكن الخطوط والنصوص - في المنتصف */
.filters-card * {
  color: #dc2626 !important;
  text-align: center !important;
  direction: rtl !important;
}

.filters-card .v-field * {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
}

/* CSS شامل مع ::deep لجميع النصوص - في المنتصف */
::deep(.filters-card *) {
  color: #dc2626 !important;
  text-align: center !important;
  direction: rtl !important;
}

::deep(.filters-card .v-field *) {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
}

/* تحسين محاذاة النصوص في الحقول - في المنتصف */
.filters-card .v-field__input {
  text-align: center !important;
  direction: rtl !important;
  padding-right: 12px !important;
  padding-left: 12px !important;
}

.filters-card .v-field__input input {
  text-align: center !important;
  direction: rtl !important;
  text-indent: 0 !important;
}

/* تحسين محاذاة الأيقونات */
.filters-card .v-field__prepend-inner {
  left: 8px !important;
  right: auto !important;
}

.filters-card .v-field__append-inner {
  right: 8px !important;
  left: auto !important;
}

/* تحسين محاذاة التسميات - في المنتصف */
.filters-card .v-field-label {
  text-align: center !important;
  direction: rtl !important;
  right: 0 !important;
  left: 0 !important;
}

/* تحسين تنظيم الخطوط بشكل شامل - في المنتصف */
.filters-card .v-field {
  text-align: center !important;
  direction: rtl !important;
}

.filters-card .v-field__field {
  text-align: center !important;
  direction: rtl !important;
}

.filters-card .v-field__field .v-field__input {
  text-align: center !important;
  direction: rtl !important;
  unicode-bidi: bidi-override !important;
}

/* تحسين محاذاة النصوص مع ::deep - في المنتصف */
::deep(.filters-card .v-field__input) {
  text-align: center !important;
  direction: rtl !important;
  unicode-bidi: bidi-override !important;
}

::deep(.filters-card .v-field__input input) {
  text-align: center !important;
  direction: rtl !important;
  unicode-bidi: bidi-override !important;
}

/* تحسين تنظيم الخطوط في جميع الحقول - في المنتصف */
.filters-card .v-text-field .v-field__input,
.filters-card .v-select .v-field__input {
  text-align: center !important;
  direction: rtl !important;
  unicode-bidi: bidi-override !important;
  padding-right: 12px !important;
  padding-left: 12px !important;
}

.tasks-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  border-radius: 10px 10px 0 0 !important;
  padding: 6px 10px !important;
}

.tasks-header span {
  font-size: 0.9rem !important;
  font-weight: 600 !important;
  letter-spacing: 0.2px !important;
}

/* جدول المهام - منسق */
.tasks-table {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 250, 252, 0.95) 100%) !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  box-shadow: 0 12px 40px rgba(67, 56, 202, 0.15) !important;
  backdrop-filter: blur(15px) !important;
  border: 2px solid rgba(199, 210, 254, 0.3) !important;
  margin-top: 20px !important;
}

/* رؤوس جدول المهام */
.tasks-table .v-data-table__th {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 50%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 500 !important;
  font-size: 0.45rem !important;
  text-align: center !important;
  padding: 2px 2px !important;
  min-height: 26px !important;
  vertical-align: middle !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0px !important;
  text-transform: none !important;
  position: relative !important;
  z-index: 2 !important;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border-right: 1px solid rgba(255, 255, 255, 0.1) !important;
  line-height: 1.0 !important;
  white-space: nowrap !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
}

/* Force white color for all text elements in table headers */
.tasks-table .v-data-table__th *,
.tasks-table .v-data-table__th span,
.tasks-table .v-data-table__th div,
.tasks-table .v-data-table__th .v-data-table-header__content,
.tasks-table .v-data-table__th .v-data-table-header__content * {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-size: 0.45rem !important;
}

.tasks-table :deep(.v-data-table-header th),
.tasks-table :deep(.v-data-table-header th *),
.tasks-table :deep(.v-data-table-header th span),
.tasks-table :deep(.v-data-table-header th div),
.tasks-table :deep(.v-data-table-header th .v-data-table-header__content),
.tasks-table :deep(.v-data-table-header th .v-data-table-header__content *) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-size: 0.45rem !important;
}

.tasks-table :deep(.v-data-table__wrapper table thead tr th),
.tasks-table :deep(.v-data-table__wrapper table thead tr th *),
.tasks-table :deep(.v-data-table__wrapper table thead tr th span),
.tasks-table :deep(.v-data-table__wrapper table thead tr th div),
.tasks-table :deep(.v-data-table__wrapper table thead tr th .v-data-table-header__content),
.tasks-table :deep(.v-data-table__wrapper table thead tr th .v-data-table-header__content *) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-size: 0.45rem !important;
}

/* Additional rules to ensure white color for all header elements */
.tasks-table :deep(.v-data-table-header),
.tasks-table :deep(.v-data-table-header th),
.tasks-table :deep(.v-data-table-header th span),
.tasks-table :deep(.v-data-table-header th div),
.tasks-table :deep(.v-data-table-header th label),
.tasks-table :deep(.v-data-table-header th .v-data-table-header__content),
.tasks-table :deep(.v-data-table-header th .v-data-table-header__content span),
.tasks-table :deep(.v-data-table-header th .v-data-table-header__content div),
.tasks-table :deep(.v-data-table-header th .v-data-table-header__content label) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* Force white for all text nodes */
.tasks-table :deep(.v-data-table__th),
.tasks-table :deep(.v-data-table__th span),
.tasks-table :deep(.v-data-table__th div),
.tasks-table :deep(.v-data-table__th label),
.tasks-table :deep(.v-data-table__th .v-data-table-header__content),
.tasks-table :deep(.v-data-table__th .v-data-table-header__content span),
.tasks-table :deep(.v-data-table__th .v-data-table-header__content div),
.tasks-table :deep(.v-data-table__th .v-data-table-header__content label) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* Override any inline styles */
.tasks-table .v-data-table__th[style*="color"],
.tasks-table :deep(.v-data-table-header th[style*="color"]) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.tasks-table .v-data-table__th:last-child {
  border-right: none !important;
}

/* صفوف جدول المهام */
.tasks-table .v-data-table__td {
  padding: 10px 8px !important;
  min-height: 60px !important;
  border-bottom: 1px solid rgba(199, 210, 254, 0.3) !important;
  text-align: center !important;
  background: rgba(255, 255, 255, 0.9) !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 400 !important;
  font-size: 0.85rem !important;
  color: #1e293b !important;
  vertical-align: middle !important;
  line-height: 1.6 !important;
  word-wrap: break-word !important;
  overflow-wrap: break-word !important;
}

/* تأثيرات صفوف جدول المهام */
.tasks-table .v-data-table__tr:hover {
  background: linear-gradient(135deg, rgba(67, 56, 202, 0.08) 0%, rgba(99, 102, 241, 0.05) 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(67, 56, 202, 0.2) !important;
  border-radius: 12px !important;
}

.tasks-table .v-data-table__tr:nth-child(even) {
  background: linear-gradient(135deg, rgba(238, 242, 255, 0.8) 0%, rgba(224, 231, 255, 0.6) 100%) !important;
}

.tasks-table .v-data-table__tr:nth-child(odd) {
  background: rgba(255, 255, 255, 0.95) !important;
}

/* تذييل جدول المهام */
.tasks-table .v-data-table-footer {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 50%, #8b5cf6 100%) !important;
  color: white !important;
  border-radius: 0 0 20px 20px !important;
  padding: 20px !important;
  box-shadow: 0 -4px 20px rgba(67, 56, 202, 0.3) !important;
}

/* جميع النصوص في الجدول - نيلي */
.tasks-table .v-data-table__td * {
  color: #4338ca !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.tasks-table .v-data-table__td .v-chip {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  opacity: 1 !important;
  visibility: visible !important;
  z-index: 10 !important;
  position: relative !important;
  font-size: 0.65rem !important;
}

/* رقم المهمة - أصغر */
.tasks-table .v-data-table__td .v-chip[color="primary"] {
  font-size: 0.5rem !important;
  padding: 2px 5px !important;
  min-height: 18px !important;
  height: auto !important;
}

.tasks-table .v-data-table__td .v-chip .v-chip__content {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  opacity: 1 !important;
  visibility: visible !important;
  color: inherit !important;
  font-size: 0.5rem !important;
}

/* جميع الـ chips (الحالة والأولوية) - أصغر */
.tasks-table .v-data-table__td .v-chip {
  font-size: 0.5rem !important;
  padding: 2px 5px !important;
  min-height: 18px !important;
  height: auto !important;
}

.tasks-table .v-data-table__td .v-chip .v-chip__content {
  font-size: 0.5rem !important;
}

.tasks-table .v-data-table__td .v-chip .v-chip__content span {
  opacity: 1 !important;
  visibility: visible !important;
  color: inherit !important;
}

.tasks-table .v-data-table__td .v-chip .v-icon {
  opacity: 1 !important;
  visibility: visible !important;
}

.tasks-table .v-data-table__td .v-btn {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-icon {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content span {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content div {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content p {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content strong {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content b {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content em {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content i {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content u {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content small {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content large {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content h1 {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content h2 {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content h3 {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content h4 {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content h5 {
  color: #4338ca !important;
}

.tasks-table .v-data-table__td .v-progress-linear .v-progress-linear__content .v-progress-linear__label .v-progress-linear__label__content h6 {
  color: #4338ca !important;
}

.task-title {
  max-width: 300px;
  text-align: right !important;
  padding: 8px 12px !important;
  overflow: hidden !important;
}

.task-title .task-name {
  font-size: 0.75rem !important;
  font-weight: 400 !important;
  color: #1e293b !important;
  margin-bottom: 3px !important;
  line-height: 1.3 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  white-space: nowrap !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  max-width: 100% !important;
  display: block !important;
}

.task-title .task-description {
  font-size: 0.65rem !important;
  font-weight: 400 !important;
  color: #64748b !important;
  margin: 0 !important;
  line-height: 1.3 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.assignee-info {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  padding: 2px 6px !important;
  font-size: 0.7rem !important;
  font-weight: 400 !important;
  color: #1e293b !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  white-space: nowrap !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
}

.due-date {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  padding: 2px 6px !important;
  font-size: 0.65rem !important;
  font-weight: 400 !important;
  color: #1e293b !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  white-space: nowrap !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
}

.due-date span {
  font-size: 0.65rem !important;
}

.progress-container {
  padding: 4px 8px !important;
  min-width: 120px;
}

.progress-text {
  font-size: 0.75rem !important;
  font-weight: 400 !important;
  color: #1e293b !important;
  margin-top: 4px !important;
  display: block !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* جميع النصوص في الجدول - محسنة */
.tasks-table .v-data-table__td,
.tasks-table .v-data-table__td div,
.tasks-table .v-data-table__td span,
.tasks-table .v-data-table__td p {
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  line-height: 1.6 !important;
  padding: 4px 8px !important;
  word-wrap: break-word !important;
  overflow-wrap: break-word !important;
}

.tasks-table .v-data-table__td strong,
.tasks-table .v-data-table__td b {
  color: #1e293b !important;
  font-weight: 700 !important;
  font-size: 0.95rem !important;
}

.tasks-table .v-data-table__td em,
.tasks-table .v-data-table__td i {
  color: #1e293b !important;
  font-style: normal !important;
  font-weight: 500 !important;
}

.tasks-table .v-data-table__td h1,
.tasks-table .v-data-table__td h2,
.tasks-table .v-data-table__td h3,
.tasks-table .v-data-table__td h4,
.tasks-table .v-data-table__td h5,
.tasks-table .v-data-table__td h6 {
  color: #1e293b !important;
  font-weight: 700 !important;
  margin: 0 !important;
  padding: 0 !important;
  line-height: 1.5 !important;
}

/* النصوص في الشريط السفلي */
.tasks-table .v-data-table-footer,
.tasks-table .v-data-table-footer div,
.tasks-table .v-data-table-footer span,
.tasks-table .v-data-table-footer p,
.tasks-table .v-data-table-footer strong,
.tasks-table .v-data-table-footer b,
.tasks-table .v-data-table-footer em,
.tasks-table .v-data-table-footer i,
.tasks-table .v-data-table-footer u,
.tasks-table .v-data-table-footer small,
.tasks-table .v-data-table-footer large,
.tasks-table .v-data-table-footer h1,
.tasks-table .v-data-table-footer h2,
.tasks-table .v-data-table-footer h3,
.tasks-table .v-data-table-footer h4,
.tasks-table .v-data-table-footer h5,
.tasks-table .v-data-table-footer h6 {
  color: #4338ca !important;
}

/* جميع العناصر في الجدول - باستثناء العناوين */
.tasks-table *:not(.v-data-table__th):not(.v-data-table-header):not(.v-data-table-header th) {
  color: #4338ca !important;
}

/* استثناء العناوين من القاعدة العامة - إجبار اللون الأبيض */
.tasks-table .v-data-table__th,
.tasks-table .v-data-table__th *,
.tasks-table .v-data-table-header,
.tasks-table .v-data-table-header th,
.tasks-table .v-data-table-header th *,
.tasks-table thead,
.tasks-table thead th,
.tasks-table thead th * {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.task-name {
  font-weight: bold;
  color: #333;
  margin: 0 0 5px 0;
}

.task-description {
  color: #666;
  font-size: 0.9rem;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.assignee-info {
  display: flex;
  align-items: center;
}

.due-date {
  display: flex;
  align-items: center;
  color: #666;
}

.progress-container {
  min-width: 120px;
}

.progress-text {
  font-size: 0.8rem;
  color: #666;
  text-align: center;
  display: block;
}

.action-buttons {
  display: flex;
  gap: 8px;
  justify-content: center;
  align-items: center;
  padding: 8px;
}

.action-buttons .v-btn {
  min-width: 36px !important;
  width: 36px !important;
  height: 36px !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.action-buttons .v-btn:hover {
  transform: translateY(-2px) scale(1.1) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.25) !important;
}

.action-buttons .v-btn:active {
  transform: translateY(0) scale(0.95) !important;
}

.dialog-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 24px !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.3) !important;
}

.dialog-actions {
  padding: 16px 24px 24px 24px !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-radius: 0 0 20px 20px !important;
  border-top: 2px solid rgba(59, 130, 246, 0.1) !important;
  display: flex !important;
  justify-content: flex-end !important;
  align-items: center !important;
  gap: 16px !important;
  margin-top: 0 !important;
}

.dialog-actions .v-btn {
  min-width: 140px !important;
  padding: 14px 28px !important;
  border-radius: 14px !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-transform: none !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
  text-align: center !important;
  display: inline-flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.dialog-actions .v-btn :deep(.v-btn__content) {
  color: inherit !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
  text-align: center !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  width: 100% !important;
  justify-content: center !important;
  align-items: center !important;
}

.dialog-actions .v-btn :deep(.v-btn__content span) {
  text-align: center !important;
  display: inline-block !important;
  width: 100% !important;
}

.dialog-actions .v-btn[color="primary"] {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  box-shadow: 
    0 6px 20px rgba(37, 99, 235, 0.4),
    0 3px 10px rgba(37, 99, 235, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
}

.dialog-actions .v-btn[color="primary"]:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  transform: translateY(-3px) scale(1.03) !important;
  box-shadow: 
    0 10px 30px rgba(37, 99, 235, 0.5),
    0 5px 15px rgba(37, 99, 235, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.5) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.dialog-actions .v-btn[color="primary"]:disabled {
  background: #cbd5e1 !important;
  color: #94a3b8 !important;
  box-shadow: none !important;
  transform: none !important;
}

.dialog-actions .v-btn[color="grey"] {
  background: rgba(255, 255, 255, 0.95) !important;
  backdrop-filter: blur(10px) !important;
  color: #475569 !important;
  border: 2px solid rgba(148, 163, 184, 0.4) !important;
  box-shadow: 
    0 4px 12px rgba(0, 0, 0, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
}

.dialog-actions .v-btn[color="grey"]:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(148, 163, 184, 0.6) !important;
  color: #334155 !important;
  transform: translateY(-3px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(0, 0, 0, 0.18),
    0 4px 12px rgba(148, 163, 184, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
}

.task-details {
  padding: 20px 0;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid #eee;
}

.detail-row:last-child {
  border-bottom: none;
}

.detail-row strong {
  color: #333;
  min-width: 120px;
}

.detail-row span {
  color: #666;
  text-align: right;
}

/* Responsive design */
@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    gap: 20px;
    text-align: center;
  }
  
  .page-title {
    font-size: 2rem;
  }
  
  .task-title {
    max-width: 200px;
  }
}

/* تنسيق نموذج إضافة المهمة */
.task-dialog-card {
  border-radius: 20px !important;
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.25) !important;
  overflow: hidden !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
}

.task-dialog-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  padding: 24px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.3) !important;
}

.task-dialog-header .text-h5 {
  color: white !important;
  font-weight: 400 !important;
  font-size: 1.5rem !important;
  letter-spacing: 0.5px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.task-dialog-header .v-icon {
  color: white !important;
}

/* تنسيق حقول النموذج */
.task-form-field {
  margin-bottom: 16px !important;
}

.task-form-field .v-field {
  border-radius: 12px !important;
  border: 4px solid #000000 !important;
  background: #ffffff !important;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.25) !important;
  transition: all 0.3s ease !important;
}

.task-form-field .v-field:hover {
  border-color: #4338ca !important;
  box-shadow: 0 8px 24px rgba(67, 56, 202, 0.2) !important;
  transform: translateY(-2px) !important;
}

.task-form-field .v-field:focus-within {
  border-color: #dc2626 !important;
  box-shadow: 0 10px 30px rgba(220, 38, 38, 0.25) !important;
  transform: translateY(-3px) !important;
}

/* تنسيق أسماء الحقول في النموذج */
.task-form-field .v-label,
.task-form-field .v-field-label,
.task-form-field .v-field-label--active,
.task-form-field .v-field-label--floating {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.6) !important;
  letter-spacing: 1px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
  background: #ffffff !important;
  padding: 12px 16px !important;
  border-radius: 10px !important;
  border: 3px solid #000000 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
}

/* تنسيق النصوص داخل الحقول */
.task-form-field .v-field__input {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  text-align: center !important;
  direction: rtl !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.8px !important;
  background: #ffffff !important;
}

/* CSS شامل مع ::deep للنموذج */
::deep(.task-form-field .v-label),
::deep(.task-form-field .v-field-label),
::deep(.task-form-field .v-field-label--active),
::deep(.task-form-field .v-field-label--floating) {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.6) !important;
  letter-spacing: 1px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
  background: #ffffff !important;
  padding: 12px 16px !important;
  border-radius: 10px !important;
  border: 3px solid #000000 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
}

::deep(.task-form-field .v-field__input) {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  text-align: center !important;
  direction: rtl !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.8px !important;
  background: #ffffff !important;
}

/* CSS إضافي لضمان الوضوح الكامل */
.task-form-field * {
  color: #000000 !important;
  font-weight: 900 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5) !important;
}

::deep(.task-form-field *) {
  color: #000000 !important;
  font-weight: 900 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5) !important;
}

/* ضمان وضوح جميع النصوص في النموذج */
.task-form-field .v-field__input input,
.task-form-field .v-field__input textarea,
.task-form-field .v-field__input select {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.8px !important;
  background: #ffffff !important;
}

::deep(.task-form-field .v-field__input input),
::deep(.task-form-field .v-field__input textarea),
::deep(.task-form-field .v-field__input select) {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.8px !important;
  background: #ffffff !important;
}
</style>
