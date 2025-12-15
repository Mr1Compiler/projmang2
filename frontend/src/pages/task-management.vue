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
}

onMounted(() => {
  // Initialize data
  tasksLoading.value = false
})
</script>


<style>
/* Import page styles */
@import './styles/task-management.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
