<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-currency-usd</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">المصاريف الإدارية</h1>
              <p class="subtitle">إدارة وتتبع جميع المصاريف الإدارية والعامة</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-currency-usd</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ totalExpenses }}</h3>
                <p class="stat-label">إجمالي المصاريف</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ activeExpenses }}</h3>
                <p class="stat-label">مصاريف نشطة</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-clock-alert</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ pendingExpenses }}</h3>
                <p class="stat-label">في الانتظار</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="modern-stat-card stat-card-info" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-chart-line</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ formatCurrency(totalCost) }}</h3>
                <p class="stat-label">إجمالي التكلفة</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Search Bar -->
      <v-card class="search-card filters-card mb-4" elevation="2">
        <v-card-title class="filters-header-new">
          <div class="d-flex align-center">
            <v-icon size="18" color="white" class="me-2" style="color: #ffffff !important;">mdi-filter</v-icon>
            <span class="filters-header-title">فلترة المصاريف</span>
          </div>
        </v-card-title>
        <v-card-text class="filters-content">
          <v-row class="align-center" no-gutters>
            <v-col cols="12" md="5">
              <v-text-field
                v-model="expenseSearchQuery"
                label="البحث في المصاريف الإدارية..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                density="compact"
                clearable
                hide-details
                class="search-field-new"
              />
            </v-col>
            <v-col cols="12" md="2" class="px-2">
              <v-select
                v-model="selectedExpenseType"
                :items="expenseTypeOptions"
                label="نوع المصروف"
                variant="outlined"
                density="compact"
                clearable
                hide-details
                class="filter-field-new"
              />
            </v-col>
            <v-col cols="12" md="2" class="px-2">
              <v-select
                v-model="selectedStatus"
                :items="statusOptions"
                label="الحالة"
                variant="outlined"
                density="compact"
                clearable
                hide-details
                class="filter-field-new"
              />
            </v-col>
            <v-col cols="12" md="3" class="d-flex justify-end">
              <v-btn
                color="primary"
                variant="elevated"
                size="small"
                class="search-btn-new"
                @click="searchExpenses"
              >
                <v-icon class="me-2" size="18">mdi-magnify</v-icon>
                بحث
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Expenses Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title">
          <div class="d-flex align-center justify-space-between" style="width: 100%;">
            <span class="title-text">المشاريع</span>
            <v-btn
              class="add-button-new btn-glow light-sweep smooth-transition"
              @click="openAddExpenseDialog"
              elevation="2"
              color="primary"
              size="small"
              style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important;"
            >
              <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
              إضافة صنف جديد
            </v-btn>
          </div>
        </v-card-title>

        <!-- جدول المصاريف -->
        <v-data-table
          :headers="expenseHeaders"
          :items="filteredExpenses"
          :search="expenseSearchQuery"
          class="project-table"
          :items-per-page="-1"
          :loading="loading"
          hover
          no-data-text="لا توجد بيانات"
          loading-text="جاري التحميل..."
          hide-default-footer
          :header-props="{
            style: 'background: linear-gradient(135deg, #047857 0%, #059669 100%); color: white; font-weight: 700;'
          }"
        >
          <!-- Serial Number Column -->
          <template #item.serial="{ index }">
            <span class="serial-number">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</span>
          </template>

          <!-- Name Column -->
          <template #item.name="{ item }">
            <span class="project-name">{{ item.name }}</span>
          </template>

          <!-- Amount Column -->
          <template #item.amount="{ item }">
            <span class="cost-text">{{ formatCurrency(item.amount) }}</span>
          </template>

          <!-- Type Column -->
          <template #item.type="{ item }">
            <v-chip
              :color="getExpenseTypeColor(item.type)"
              size="small"
              variant="tonal"
            >
              {{ item.type || 'غير محدد' }}
            </v-chip>
          </template>

          <!-- Expense Date Column -->
          <template #item.expenseDate="{ item }">
            <span class="date-text">{{ formatDate(item.expenseDate) }}</span>
          </template>

          <!-- Status Column -->
          <template #item.status="{ item }">
            <v-chip
              :color="getStatusColor(item.status)"
              size="small"
              variant="tonal"
            >
              {{ getStatusText(item.status) }}
            </v-chip>
          </template>

          <!-- Notes Column -->
          <template #item.notes="{ item }">
            <span class="notes-text">{{ item.notes || 'لايوجد' }}</span>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="action-buttons">
              <v-btn
                size="small"
                color="primary"
                variant="text"
                @click="editExpense(item)"
                icon
                class="action-btn"
                title="تعديل"
              >
                <v-icon size="16">mdi-pencil</v-icon>
              </v-btn>
              <v-btn
                size="small"
                color="error"
                variant="text"
                @click="deleteExpense(item)"
                icon
                class="action-btn"
                title="حذف"
              >
                <v-icon size="16">mdi-delete</v-icon>
              </v-btn>
            </div>
          </template>
        </v-data-table>

        <!-- Pagination -->
        <div class="d-flex justify-center pa-4" v-if="totalPages > 0">
          <v-pagination
            v-model="currentPage"
            :length="totalPages"
            :total-visible="7"
            @update:model-value="onPageChange"
            rounded="circle"
            density="comfortable"
            active-color="primary"
          />
        </div>
      </v-card>

      <!-- Add/Edit Administrative Expense Dialog - Clean Form Style -->
      <v-dialog v-model="expenseDialog" max-width="900" scrollable persistent>
        <v-card class="clean-dialog-card clean-form-card">
          <!-- Header Section -->
          <v-card-title class="clean-dialog-header clean-form-header">
            <h2 class="clean-form-title">
              {{ isEditingExpense ? 'تعديل المصروف الإداري' : 'معلومات المصروف الإداري' }}
            </h2>
          </v-card-title>

          <!-- Form Content -->
          <v-card-text class="clean-form-content">
            <p class="clean-form-instruction">
              لإتمام {{ isEditingExpense ? 'تعديل' : 'إضافة' }} المصروف الإداري، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
            </p>

            <v-form ref="expenseFormRef" v-model="expenseFormValid">
              <!-- الصف الأول: الاسم، المبلغ -->
              <v-row class="clean-form-row">
                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      اسم المصروف <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.name"
                      variant="outlined"
                      density="comfortable"
                      placeholder="أدخل اسم المصروف"
                      :rules="[v => !!v || 'اسم المصروف مطلوب']"
                      required
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      المبلغ (د.ع) <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model.number="expenseForm.amount"
                      variant="outlined"
                      density="comfortable"
                      type="number"
                      placeholder="0"
                      :rules="[v => v > 0 || 'المبلغ مطلوب ويجب أن يكون أكبر من صفر']"
                      required
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الثاني: تاريخ المصروف، النوع -->
              <v-row class="clean-form-row">
                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      تاريخ المصروف <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.expenseDate"
                      variant="outlined"
                      density="comfortable"
                      type="date"
                      :rules="[v => !!v || 'تاريخ المصروف مطلوب']"
                      required
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      نوع المصروف
                    </label>
                    <v-text-field
                      v-model="expenseForm.type"
                      variant="outlined"
                      density="comfortable"
                      placeholder="أدخل نوع المصروف"
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الثالث: الحالة -->
              <v-row class="clean-form-row">
                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      الحالة
                    </label>
                    <v-select
                      v-model="expenseForm.status"
                      :items="statusOptions.filter(s => s.value)"
                      item-title="title"
                      item-value="value"
                      variant="outlined"
                      density="comfortable"
                      placeholder="اختر الحالة"
                      required
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الخامس: الملاحظات -->
              <v-row class="clean-form-row">
                <v-col cols="12" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">الملاحظات</label>
                    <v-textarea
                      v-model="expenseForm.notes"
                      variant="outlined"
                      rows="4"
                      density="comfortable"
                      placeholder="أدخل ملاحظات إضافية"
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>

          <!-- Footer Actions -->
          <v-card-actions class="clean-form-actions">
            <v-spacer />
            <v-btn
              class="clean-form-cancel-btn"
              variant="outlined"
              @click="closeExpenseDialog"
            >
              إلغاء
            </v-btn>
            <v-btn
              class="clean-form-continue-btn"
              variant="elevated"
              :disabled="!expenseFormValid"
              @click="saveExpense"
            >
              {{ isEditingExpense ? 'تحديث المصروف' : 'حفظ المصروف' }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { listExpenses, createExpense, updateExpense, deleteExpense as deleteExpenseApi, getExpenseStats } from '@/api/expenses'
import { DEFAULT_LIMIT } from '@/constants/pagination'

const router = useRouter()

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(DEFAULT_LIMIT)
const totalItems = ref(0)
const totalPages = ref(0)

// متغيرات الحالة الأساسية
const loading = ref(false)
const expenseDialog = ref(false)
const expenseFormValid = ref(false)
const isEditingExpense = ref(false)
const expenseSearchQuery = ref('')
const selectedProjectFilter = ref('')
const selectedCostRange = ref('')
const selectedExpense = ref(null)
const selectedExpenseType = ref('')
const selectedStatus = ref('')

// جدول المصاريف الإدارية - aligned with backend DTO
const expenseHeaders = [
  { title: 'التسلسل', key: 'serial', sortable: false, width: '80px' },
  { title: 'الاسم', key: 'name', sortable: true, width: '180px' },
  { title: 'المبلغ', key: 'amount', sortable: true, width: '120px' },
  { title: 'النوع', key: 'type', sortable: true, width: '120px' },
  { title: 'التاريخ', key: 'expenseDate', sortable: true, width: '130px' },
  { title: 'الحالة', key: 'status', sortable: true, width: '100px' },
  { title: 'الملاحظات', key: 'notes', sortable: false, width: '150px' },
  { title: 'الإجراءات', key: 'actions', sortable: false, width: '100px' }
]

// نموذج المصاريف الإدارية - aligned with backend CreateExpense DTO
const expenseForm = ref({
  name: '',
  amount: 0,
  type: '',
  expenseDate: new Date().toISOString().split('T')[0],
  projectId: null,
  status: 'pending',
  notes: ''
})

// Stats from backend API
const expenseStats = ref({
  total: 0,
  totalAmount: 0,
  pending: 0,
  approved: 0,
  rejected: 0,
  averageAmount: 0
})

// بيانات المصاريف الإدارية - from backend API
const expenses = ref([])

// استعلام البحث للمشاريع
const projectSearchQuery = ref('')

// خيارات فلترة المصاريف
const projectFilterOptions = computed(() => {
  const projects = [...new Set(expenses.value.filter(e => e.projectId).map(expense => expense.projectId))]
  return projects.map(project => ({ title: `مشروع ${project}`, value: project }))
})

const costRangeOptions = [
  { title: 'أقل من 50,000 د.ع', value: 'low' },
  { title: '50,000 - 100,000 د.ع', value: 'medium' },
  { title: 'أكثر من 100,000 د.ع', value: 'high' }
]

const expenseTypes = [
  'تطوير',
  'تحديث',
  'بناء',
  'أمن',
  'صيانة',
  'تدريب',
  'أخرى'
]

const expenseTypeOptions = ref([
  'جميع الأنواع',
  'مصاريف إدارية',
  'مصاريف مشاريع',
  'مرتبات وأجور',
  'إيجار ومرافق',
  'معدات وتجهيزات',
  'نقل ومواصلات',
  'تدريب وتطوير',
  'صيانة وإصلاح',
  'مواد مكتبية',
  'تسويق وإعلان',
  'أخرى'
])

const statusOptions = ref([
  { title: 'جميع الحالات', value: '' },
  { title: 'معتمد', value: 'approved' },
  { title: 'معلق', value: 'pending' },
  { title: 'مرفوض', value: 'rejected' }
])

// إحصائيات المصاريف - from backend API
const totalExpenses = computed(() => expenseStats.value.total || expenses.value.length)
const activeExpenses = computed(() => expenseStats.value.approved || 0)
const pendingExpenses = computed(() => expenseStats.value.pending || 0)
const totalCost = computed(() => expenseStats.value.totalAmount || 0)

// دوال المصاريف الإدارية
const searchExpenses = () => {
  // دالة البحث - يمكن إضافة منطق البحث المتقدم هنا
  console.log('البحث عن:', expenseSearchQuery.value)
}

const openAddExpenseDialog = () => {
  expenseDialog.value = true
  isEditingExpense.value = false
  selectedExpense.value = null
  expenseForm.value = {
    name: '',
    amount: 0,
    type: '',
    expenseDate: new Date().toISOString().split('T')[0],
    projectId: null,
    status: 'pending',
    notes: ''
  }
}

const closeExpenseDialog = () => {
  expenseDialog.value = false
  isEditingExpense.value = false
  selectedExpense.value = null
  expenseForm.value = {
    name: '',
    amount: 0,
    type: '',
    expenseDate: new Date().toISOString().split('T')[0],
    projectId: null,
    status: 'pending',
    notes: ''
  }
}

const editExpense = (expense) => {
  selectedExpense.value = expense
  isEditingExpense.value = true
  expenseForm.value = {
    name: expense.name || '',
    amount: expense.amount || 0,
    type: expense.type || '',
    expenseDate: expense.expenseDate ? formatDateForInput(expense.expenseDate) : new Date().toISOString().split('T')[0],
    projectId: expense.projectId || null,
    status: expense.status || 'pending',
    notes: expense.notes || ''
  }
  expenseDialog.value = true
}

// Helper function to format date for input field
const formatDateForInput = (dateString) => {
  if (!dateString) return new Date().toISOString().split('T')[0]
  const date = new Date(dateString)
  return date.toISOString().split('T')[0]
}

// دالة تعديل المشروع
const editProject = (project) => {
  console.log('تعديل المشروع:', project)
  // يمكن إضافة منطق التعديل هنا
}

// دالة عرض تفاصيل المشروع
const viewProjectDetails = (project) => {
  console.log('عرض تفاصيل المشروع:', project)
  console.log('الانتقال إلى صفحة مصاريف المشروع...')
  
  // توجيه إلى صفحة مصاريف المشروع
  router.push({
    path: '/project-expenses',
    query: {
      projectName: project.projectName,
      projectId: project.id,
      startDate: project.startDate,
      endDate: project.endDate,
      cost: project.cost,
      workLocation: project.workLocation,
      notes: project.notes
    }
  }).then(() => {
    console.log('تم الانتقال بنجاح إلى صفحة مصاريف المشروع')
  }).catch((error) => {
    console.error('خطأ في الانتقال:', error)
  })
}

const deleteExpense = async (expense) => {
  if (confirm(`هل أنت متأكد من حذف المصروف "${expense.name}"؟`)) {
    try {
      await deleteExpenseApi(expense.id)
      await loadExpenses()
    } catch (err) {
      console.error('Error deleting expense:', err)
    }
  }
}

const saveExpense = async () => {
  if (expenseFormValid.value) {
    try {
      // Convert date to ISO format for backend
      const dateValue = expenseForm.value.expenseDate
      const isoDate = dateValue ? new Date(dateValue).toISOString() : new Date().toISOString()

      const payload = {
        name: expenseForm.value.name,
        amount: Number(expenseForm.value.amount),
        expenseDate: isoDate
      }

      // Add optional fields if provided
      if (expenseForm.value.type) payload.type = expenseForm.value.type
      if (expenseForm.value.projectId) payload.projectId = expenseForm.value.projectId
      if (expenseForm.value.status) payload.status = expenseForm.value.status
      if (expenseForm.value.notes) payload.notes = expenseForm.value.notes

      if (isEditingExpense.value && selectedExpense.value) {
        // Update existing expense
        await updateExpense(selectedExpense.value.id, payload)
      } else {
        // Create new expense
        await createExpense(payload)
      }

      await loadExpenses()
      closeExpenseDialog()
    } catch (err) {
      console.error('Error saving expense:', err)
    }
  }
}

// فلترة المصاريف
const filteredExpenses = computed(() => {
  let filtered = expenses.value

  // Filter by search query
  if (expenseSearchQuery.value) {
    filtered = filtered.filter(expense =>
      expense.name?.toLowerCase().includes(expenseSearchQuery.value.toLowerCase()) ||
      expense.notes?.toLowerCase().includes(expenseSearchQuery.value.toLowerCase())
    )
  }

  // Filter by type
  if (selectedExpenseType.value && selectedExpenseType.value !== 'جميع الأنواع') {
    filtered = filtered.filter(expense => expense.type === selectedExpenseType.value)
  }

  // Filter by status
  if (selectedStatus.value) {
    filtered = filtered.filter(expense => expense.status === selectedStatus.value)
  }

  // Filter by project
  if (selectedProjectFilter.value) {
    filtered = filtered.filter(expense => expense.projectId === selectedProjectFilter.value)
  }

  // Filter by cost range
  if (selectedCostRange.value) {
    filtered = filtered.filter(expense => {
      const cost = expense.amount
      switch (selectedCostRange.value) {
        case 'low':
          return cost < 50000
        case 'medium':
          return cost >= 50000 && cost <= 100000
        case 'high':
          return cost > 100000
        default:
          return true
      }
    })
  }

  return filtered
})

// دوال مساعدة
const formatCurrency = (amount) => {
  if (amount == null) return '0 IQD'
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount) + ' IQD'
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

// دالة الحصول على لون نوع المصروف
const getExpenseTypeColor = (type) => {
  const colors = {
    'تطوير': 'primary',
    'تحديث': 'info',
    'بناء': 'warning',
    'أمن': 'error',
    'صيانة': 'success',
    'تدريب': 'purple',
    'أخرى': 'grey'
  }
  return colors[type] || 'grey'
}

// دالة الحصول على لون الحالة
const getStatusColor = (status) => {
  const colors = {
    'approved': 'success',
    'pending': 'warning',
    'rejected': 'error',
    'معتمد': 'success',
    'معلق': 'warning',
    'مرفوض': 'error',
    'مسودة': 'grey'
  }
  return colors[status] || 'grey'
}

// دالة الحصول على نص الحالة
const getStatusText = (status) => {
  const texts = {
    'approved': 'معتمد',
    'pending': 'معلق',
    'rejected': 'مرفوض'
  }
  return texts[status] || status || 'غير محدد'
}

// ============ Load Data from Backend ============
const loadExpenses = async (page = currentPage.value) => {
  loading.value = true
  try {
    // Load expenses list and stats in parallel
    const [response, stats] = await Promise.all([
      listExpenses({ page, limit: itemsPerPage.value }),
      getExpenseStats()
    ])
    console.log('Expenses response received:', response)
    console.log('Expenses stats received:', stats)

    // Update pagination state
    expenses.value = response.data
    totalItems.value = response.total
    totalPages.value = response.totalPages
    currentPage.value = response.page

    if (stats) {
      expenseStats.value = stats
    }
  } catch (err) {
    console.error('Error loading expenses:', err)
  } finally {
    loading.value = false
  }
}

// Handle page change
const onPageChange = (page) => {
  currentPage.value = page
  loadExpenses(page)
}

onMounted(() => {
  loadExpenses()
})
</script>

<style scoped>
/* Page Styles */
.data-page {
  background: #ffffff !important;
  color: var(--text-dark);
  min-height: 100vh;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

.fullscreen-content {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  text-align: center;
  padding: 0 20px;
}

/* Header Styles - نفس تنسيق صفحة المهندسين */
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

.engineers-header-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: sweep 2s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: diagonalShimmer 4s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 60px rgba(25, 118, 210, 0.5);
  animation: hoverPulse 0.6s ease-in-out;
}

.engineers-header-card:hover::before {
  animation: sweep 1s ease-in-out infinite;
}

.engineers-header-card:hover::after {
  animation: diagonalShimmer 2s ease-in-out infinite;
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
  animation: slideInFromRight 1s ease-out 0.9s both, float 3s ease-in-out infinite 2s, pulse 2s ease-in-out infinite 2s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 0.5rem;
}

.engineer-emoji .v-icon {
  filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3));
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #ffffff, #e3f2fd);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  position: relative;
  animation: iconGlow 2s ease-in-out infinite 3s, iconBounce 3s ease-in-out infinite 3s;
}

.engineer-emoji .v-icon:first-child {
  animation: iconGlow 2s ease-in-out infinite 3s, iconBounce 3s ease-in-out infinite 3s;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
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

/* Stats Row */
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
  min-height: 140px !important;
  margin: 0 !important;
  box-sizing: border-box !important;
}

.stat-card {
  background: linear-gradient(145deg, #9e7272 0%, #f8fafc 100%) !important;
  border-radius: 12px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  height: 100% !important;
  min-height: 100px !important;
  flex: 1 1 0;
  max-width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06), 0 1px 4px rgba(0, 0, 0, 0.03) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  margin: 0 !important;
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-sizing: border-box !important;
}

.stat-card:hover {
  transform: translateY(-12px) scale(1.05);
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2), 0 8px 16px rgba(0, 0, 0, 0.1) !important;
  border-color: rgba(219, 206, 206, 0.9);
  background: linear-gradient(145deg, #fffcfc 0%, #f1f5f9 100%) !important;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  border-radius: 24px 24px 0 0;
  z-index: 1;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

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

.stat-icon {
  position: relative;
  z-index: 2;
  margin-bottom: 0.5rem;
  padding: 8px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 241, 241, 0.8) 0%, rgba(248, 250, 252, 0.8) 100%);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
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

.stat-card:hover .stat-icon {
  transform: scale(1.1);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  background: linear-gradient(135deg, rgba(242, 239, 239, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%);
}

.stat-card h3 {
  font-size: 1.5rem !important;
  font-weight: 800 !important;
  margin-bottom: 10px !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.15);
  text-align: center !important;
  width: 100%;
  line-height: 1.1;
  position: relative;
  z-index: 2;
  color: #1976d2 !important;
}

.stat-card p {
  font-size: 1rem !important;
  font-weight: 600 !important;
  opacity: 1 !important;
  text-align: center !important;
  width: 100%;
  margin: 0;
  color: #424242 !important;
  position: relative;
  z-index: 3;
  letter-spacing: 0.5px;
  visibility: visible !important;
  display: block !important;
}

/* تأثيرات خاصة لكل لون */
.stat-card:nth-child(1)::before {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 50%, #7c3aed 100%);
}

.stat-card:nth-child(1):hover {
  box-shadow: 0 25px 50px rgba(59, 130, 246, 0.25), 0 8px 16px rgba(59, 130, 246, 0.15) !important;
}

.stat-card:nth-child(1) h3 {
  color: #3b82f6 !important;
}

.stat-card:nth-child(2)::before {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%);
}

.stat-card:nth-child(2):hover {
  box-shadow: 0 25px 50px rgba(16, 185, 129, 0.25), 0 8px 16px rgba(16, 185, 129, 0.15) !important;
}

.stat-card:nth-child(2) h3 {
  color: #10b981 !important;
}

.stat-card:nth-child(3)::before {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #b45309 100%);
}

.stat-card:nth-child(3):hover {
  box-shadow: 0 25px 50px rgba(245, 158, 11, 0.25), 0 8px 16px rgba(245, 158, 11, 0.15) !important;
}

.stat-card:nth-child(3) h3 {
  color: #f59e0b !important;
}

.stat-card:nth-child(4)::before {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 50%, #0e7490 100%);
}

.stat-card:nth-child(4):hover {
  box-shadow: 0 25px 50px rgba(6, 182, 212, 0.25), 0 8px 16px rgba(6, 182, 212, 0.15) !important;
}

.stat-card:nth-child(4) h3 {
  color: #06b6d4 !important;
}

/* Modern Statistics Cards - نفس تصميم الصفحة الرئيسية */
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
  overflow: visible;
}

.stat-icon-wrapper::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 48px;
  height: 48px;
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
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
  min-height: 32px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.stat-icon-wrapper :deep(.v-icon) {
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
  min-height: 32px !important;
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

.check-icon,
.stat-icon-wrapper .check-icon,
.stat-icon-wrapper :deep(.check-icon) {
  transform: scaleX(-1) !important;
}

.stat-info {
  flex: 1;
  text-align: center;
  width: 100%;
}

.stat-value {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 0.5rem;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  text-align: center;
  font-variant-numeric: tabular-nums;
  unicode-bidi: embed;
  color: #000000 !important;
}

.stat-label {
  font-size: 1rem;
  font-weight: 500;
  text-align: center;
  color: #64748b;
}

/* Data Table Card */
.data-table-card {
  background: #ffffff !important;
  border: 1px solid #e2e8f0 !important;
  border-radius: 16px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
  overflow: hidden;
}

/* إصلاح عرض الجدول */
.data-table-card .v-card-title {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.data-table-card .v-data-table {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  margin-top: 1rem !important;
}

/* Search Card Styles */
.search-card,
.filters-card {
  background: #ffffff !important;
  backdrop-filter: blur(15px) !important;
  border-radius: 20px !important;
  border: 1px solid #e2e8f0 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05) !important;
  overflow: hidden !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  width: 100%;
  overflow-x: hidden;
  box-sizing: border-box;
}

.search-card:hover,
.filters-card:hover {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1) !important;
  border-color: #cbd5e1 !important;
}

.filters-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 12px 16px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.3) !important;
  display: flex !important;
  align-items: center !important;
  gap: 8px !important;
  min-height: auto !important;
}

/* شكل جديد للفلترة */
.filters-header-new {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  border-radius: 8px 8px 0 0 !important;
  padding: 0.5rem 0.75rem !important;
  border-bottom: 1px solid #1e40af !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 1px 4px rgba(59, 130, 246, 0.2) !important;
  min-height: 36px !important;
}

.filters-header-new::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, #93c5fd, #60a5fa, #3b82f6);
  border-radius: 8px 8px 0 0;
}

.filters-header-title {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  letter-spacing: 0.2px !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

.filters-content {
  padding: 0.5rem 0.75rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 50%, #e2e8f0 100%) !important;
  border-radius: 0 0 8px 8px !important;
}

.search-field-new :deep(.v-field),
.filter-field-new :deep(.v-field) {
  background: #ffffff !important;
  border-radius: 8px !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease !important;
}

.search-field-new :deep(.v-field:hover),
.filter-field-new :deep(.v-field:hover) {
  border-color: #2563eb !important;
  box-shadow: 0 2px 6px rgba(59, 130, 246, 0.2) !important;
}

.search-field-new :deep(.v-field--focused),
.filter-field-new :deep(.v-field--focused) {
  border-color: #2563eb !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.search-field-new :deep(.v-label),
.filter-field-new :deep(.v-label) {
  color: #1e40af !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
}

.search-field-new :deep(.v-field__prepend-inner .v-icon),
.filter-field-new :deep(.v-field__append-inner .v-icon) {
  color: #3b82f6 !important;
}

.search-btn-new {
  height: 36px !important;
  min-width: 100px !important;
  border-radius: 8px !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-transform: none !important;
  box-shadow: 0 2px 4px rgba(59, 130, 246, 0.2) !important;
  transition: all 0.3s ease !important;
}

.search-btn-new:hover {
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 8px rgba(59, 130, 246, 0.3) !important;
}

.add-button-new {
  height: 36px !important;
  min-width: auto !important;
  border-radius: 8px !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-transform: none !important;
  padding: 0 16px !important;
}

.search-card .v-card-text {
  padding: 16px !important;
  width: 100%;
  box-sizing: border-box;
}

.search-card .v-row {
  margin: 0;
  width: 100%;
}

.search-card .v-col {
  padding: 8px;
  box-sizing: border-box;
  width: 100%;
}

.search-card .v-field__input,
.search-card .v-field__input input,
.filters-card .v-field__input,
.filters-card .v-field__input input,
.search-field .v-field__input,
.search-field .v-field__input input {
  background: #ffffff !important;
  border-radius: 8px !important;
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-shadow: none !important;
  letter-spacing: 0.2px !important;
}

.search-card .v-field__input::placeholder,
.filters-card .v-field__input::placeholder,
.search-field .v-field__input::placeholder,
.search-card .v-field__input input::placeholder,
.filters-card .v-field__input input::placeholder,
.search-field .v-field__input input::placeholder {
  color: #64748b !important;
  font-weight: 500 !important;
  opacity: 0.8 !important;
}

.search-card .v-field__outline,
.filters-card .v-field__outline,
.search-field .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 1px !important;
}

.search-field .v-field:hover .v-field__outline,
.search-card .v-field:hover .v-field__outline,
.filters-card .v-field:hover .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 1.5px !important;
}

.search-card .v-field--focused .v-field__outline,
.filters-card .v-field--focused .v-field__outline,
.search-field .v-field--focused .v-field__outline,
.search-card .v-field:focus-within .v-field__outline,
.filters-card .v-field:focus-within .v-field__outline,
.search-field .v-field:focus-within .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.search-card .v-label,
.search-card .v-field-label,
.filters-card .v-label,
.filters-card .v-field-label {
  color: #475569 !important;
  font-weight: 600 !important;
  font-size: 0.8rem !important;
  text-shadow: none !important;
  letter-spacing: 0.2px !important;
}

.search-card .v-field-label--active,
.search-card .v-field-label--floating,
.filters-card .v-field-label--active,
.filters-card .v-field-label--floating {
  color: #3b82f6 !important;
  font-weight: 700 !important;
  font-size: 0.8rem !important;
}

.search-field .v-label,
.search-field .v-field-label {
  color: #1e40af !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
}

.search-field .v-field-label--active,
.search-field .v-field-label--floating {
  color: #1e40af !important;
  font-weight: 600 !important;
}

.search-card .v-field__prepend-inner .v-icon,
.filters-card .v-field__prepend-inner .v-icon,
.search-field .v-field__prepend-inner .v-icon {
  color: #3b82f6 !important;
  opacity: 0.8 !important;
}

.search-card .v-field--focused .v-field__prepend-inner .v-icon,
.filters-card .v-field--focused .v-field__prepend-inner .v-icon,
.search-field .v-field--focused .v-field__prepend-inner .v-icon {
  color: #2563eb !important;
  opacity: 1 !important;
}

/* حقول الفلترة */
.filter-field .v-field__input,
.filter-field .v-field__input input,
.filter-field .v-field__input select {
  background: #ffffff !important;
  border-radius: 8px !important;
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-shadow: none !important;
  letter-spacing: 0.2px !important;
}

.filter-field .v-field__input::placeholder,
.filter-field .v-field__input input::placeholder,
.filter-field .v-field__input select::placeholder {
  color: #64748b !important;
  font-weight: 500 !important;
  opacity: 0.8 !important;
}

.filter-field .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 1px !important;
}

.filter-field .v-field:hover .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 1.5px !important;
}

.filter-field .v-field--focused .v-field__outline,
.filter-field .v-field:focus-within .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.filter-field .v-label,
.filter-field .v-field-label {
  color: #475569 !important;
  font-weight: 600 !important;
  font-size: 0.8rem !important;
  text-shadow: none !important;
  letter-spacing: 0.2px !important;
}

.filter-field .v-field-label--active,
.filter-field .v-field-label--floating {
  color: #3b82f6 !important;
  font-weight: 700 !important;
  font-size: 0.8rem !important;
}

.filter-field .v-field__append-inner .v-icon {
  color: #3b82f6 !important;
  opacity: 0.8 !important;
}

.filter-field .v-field--focused .v-field__append-inner .v-icon {
  color: #2563eb !important;
  opacity: 1 !important;
}

.filter-field .v-field__append-inner .v-icon {
  color: #64748b !important;
}

.search-btn {
  height: 48px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  background: #1e40af !important;
  color: rgb(248, 248, 248) !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 4px rgba(30, 64, 175, 0.2) !important;
}

.search-btn:hover {
  background: #1e3a8a !important;
  box-shadow: 0 4px 8px rgba(30, 64, 175, 0.3) !important;
  transform: translateY(-1px) !important;
}

/* زر إضافة صنف جديد */
.add-expense-btn,
.add-expense-btn.v-btn,
.v-btn.add-expense-btn {
  height: 48px !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  text-transform: none !important;
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: white !important;
  border-radius: 12px !important;
  box-shadow: 
    0 4px 16px rgba(16, 185, 129, 0.3),
    0 2px 8px rgba(5, 150, 105, 0.2),
    0 0 20px rgba(16, 185, 129, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  position: relative;
  overflow: hidden;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  padding: 10px 20px !important;
  animation: addButtonGlow 2s ease-in-out infinite !important;
}

@keyframes addButtonGlow {
  0%, 100% {
    box-shadow: 
      0 4px 16px rgba(16, 185, 129, 0.3),
      0 2px 8px rgba(5, 150, 105, 0.2),
      0 0 20px rgba(16, 185, 129, 0.4),
      inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  }
  50% {
    box-shadow: 
      0 4px 20px rgba(16, 185, 129, 0.5),
      0 2px 12px rgba(5, 150, 105, 0.4),
      0 0 30px rgba(16, 185, 129, 0.6),
      inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  }
}

.add-expense-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.6s ease;
  z-index: 1;
}

.add-expense-btn:hover::before {
  left: 100%;
}

.add-expense-btn::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.3);
  transform: translate(-50%, -50%);
  transition: width 0.6s ease, height 0.6s ease, opacity 0.6s ease;
  z-index: 0;
  opacity: 0;
}

.add-expense-btn:active::after {
  width: 300px;
  height: 300px;
  opacity: 0;
  transition: width 0.3s ease, height 0.3s ease, opacity 0.3s ease;
}

.add-expense-btn:hover,
.add-expense-btn.v-btn:hover,
.v-btn.add-expense-btn:hover {
  transform: translateY(-3px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(16, 185, 129, 0.5),
    0 4px 12px rgba(5, 150, 105, 0.4),
    0 0 40px rgba(16, 185, 129, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  background: linear-gradient(135deg, #059669 0%, #047857 50%, #065f46 100%) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
  animation: addButtonGlowHover 1.5s ease-in-out infinite !important;
}

@keyframes addButtonGlowHover {
  0%, 100% {
    box-shadow: 
      0 8px 24px rgba(16, 185, 129, 0.5),
      0 4px 12px rgba(5, 150, 105, 0.4),
      0 0 40px rgba(16, 185, 129, 0.7),
      inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  }
  50% {
    box-shadow: 
      0 8px 28px rgba(16, 185, 129, 0.6),
      0 4px 16px rgba(5, 150, 105, 0.5),
      0 0 50px rgba(16, 185, 129, 0.8),
      inset 0 1px 0 rgba(255, 255, 255, 0.5) !important;
  }
}

.add-expense-btn:active {
  transform: translateY(-1px) scale(1.01) !important;
  box-shadow: 
    0 4px 12px rgba(16, 185, 129, 0.4),
    0 2px 6px rgba(5, 150, 105, 0.3),
    0 0 25px rgba(16, 185, 129, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  animation: addButtonClick 0.3s ease !important;
}

@keyframes addButtonClick {
  0% {
    transform: translateY(-1px) scale(1.01);
  }
  50% {
    transform: translateY(-1px) scale(0.98);
  }
  100% {
    transform: translateY(-1px) scale(1.01);
  }
}

.add-expense-btn :deep(.v-btn__content) {
  color: white !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  width: 100% !important;
  text-align: center !important;
  gap: 6px !important;
  position: relative;
  z-index: 2;
  font-size: 0.9rem !important;
  line-height: 1.4 !important;
}

.add-expense-btn :deep(.v-btn__prepend),
.add-expense-btn :deep(.v-btn__append) {
  color: white !important;
  margin: 0 !important;
  position: relative;
  z-index: 2;
}

.add-expense-btn :deep(.v-icon) {
  color: white !important;
  margin: 0 !important;
  transition: transform 0.3s ease !important;
  position: relative;
  z-index: 2;
  font-size: 18px !important;
  width: 18px !important;
  height: 18px !important;
}

.add-expense-btn:hover :deep(.v-icon) {
  transform: rotate(90deg) scale(1.1) !important;
}

/* Table Title Styles */
.table-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  padding: 8px 12px !important;
  border-radius: 8px 8px 0 0 !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  min-height: 36px !important;
  margin-bottom: 1rem !important;
}

/* Indigo Title */
.indigo-title {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

/* رسالة عدم وجود جدول */
.no-table-message {
  text-align: center;
  padding: 60px 20px;
  background: #f8f9fa;
  border-radius: 12px;
  margin: 20px;
}

.no-table-message h3 {
  color: #6c757d;
  margin: 20px 0 10px 0;
  font-size: 1.5rem;
  font-weight: 600;
}

.no-table-message p {
  color: #fefeff;
  font-size: 1rem;
  margin: 0;
}

/* تنسيقات جدول المشاريع */
.project-table {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.project-table .v-data-table__th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: right !important;
  padding: 10px 8px !important;
  border-bottom: 2px solid #1e3a8a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.3px !important;
  border-radius: 0 !important;
  position: relative !important;
}

.project-table .v-data-table__th::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #ffffff 0%, #e5e7eb 50%, #ffffff 100%);
  opacity: 0.3;
}

/* تنسيقات إضافية لضمان وضوح عناوين الجدول */
.project-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: right !important;
  padding: 10px 8px !important;
  border-bottom: 2px solid #1e3a8a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.3px !important;
  position: relative !important;
}

.project-table .v-data-table__wrapper table thead tr th .v-data-table__th__content {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.3px !important;
}

/* تصغير عناوين "التسلسل" و"اسم المشروع" فقط (من اليمين) */
.project-table .v-data-table__wrapper table thead tr th:nth-last-child(8),
.project-table .v-data-table__wrapper table thead tr th:nth-last-child(7),
.project-table .v-data-table__th:nth-last-child(8),
.project-table .v-data-table__th:nth-last-child(7) {
  font-size: 0.85rem !important;
  padding: 12px 10px !important;
  letter-spacing: 0.5px !important;
}

.project-table .v-data-table__wrapper table thead tr th:nth-last-child(8) .v-data-table__th__content,
.project-table .v-data-table__wrapper table thead tr th:nth-last-child(7) .v-data-table__th__content,
.project-table .v-data-table__th:nth-last-child(8) .v-data-table__th__content,
.project-table .v-data-table__th:nth-last-child(7) .v-data-table__th__content {
  font-size: 0.85rem !important;
  letter-spacing: 0.5px !important;
}

/* قواعد إضافية باستخدام :deep() */
.project-table :deep(.v-data-table__wrapper table thead tr th:nth-last-child(8)),
.project-table :deep(.v-data-table__wrapper table thead tr th:nth-last-child(7)) {
  font-size: 0.85rem !important;
  padding: 12px 10px !important;
  letter-spacing: 0.5px !important;
}

.project-table :deep(.v-data-table__wrapper table thead tr th:nth-last-child(8) *),
.project-table :deep(.v-data-table__wrapper table thead tr th:nth-last-child(7) *) {
  font-size: 0.85rem !important;
  letter-spacing: 0.5px !important;
}

/* قواعد قوية لعناوين الجدول */
.v-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: right !important;
  padding: 10px 8px !important;
  border-bottom: 2px solid #1e3a8a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.3px !important;
}

.v-data-table .v-data-table__wrapper table thead tr th * {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.3px !important;
}

.project-table .v-data-table__td {
  padding: 12px 8px !important;
  border-bottom: 1px solid #e0e0e0 !important;
  text-align: center !important;
  background: linear-gradient(135deg, #eff6ff 0%, #dbeafe 100%) !important;
  color: #000000 !important;
  font-weight: 500 !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(even) {
  background: linear-gradient(135deg, #eff6ff 0%, #dbeafe 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(odd) {
  background: linear-gradient(135deg, #ffffff 0%, #eff6ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:hover {
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%) !important;
}

.project-table .serial-number {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
}

.project-table .project-name {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  text-align: right !important;
}

.project-table .date-text {
  color: #000000 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
}

.project-table .cost-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
}

.project-table .location-text {
  color: #000000 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
}

.project-table .notes-text {
  color: #000000 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
  text-align: right !important;
}

.project-table .action-btn {
  color: #000000 !important;
  background: transparent !important;
  border: none !important;
  box-shadow: none !important;
}

.project-table .action-btn:hover {
  background: #f0f0f0 !important;
}

/* تنسيق أزرار الإجراءات */
.action-buttons {
  display: flex;
  gap: 4px;
  align-items: center;
  justify-content: center;
}

.action-btn {
  min-width: 32px !important;
  height: 32px !important;
  border-radius: 6px !important;
  transition: all 0.3s ease !important;
}

.details-btn {
  color: #1e40af !important;
  background: rgba(30, 64, 175, 0.1) !important;
}

.details-btn:hover {
  background: rgba(30, 64, 175, 0.2) !important;
  transform: scale(1.05) !important;
}

.action-btn:not(.details-btn) {
  color: #6b7280 !important;
  background: rgba(107, 114, 128, 0.1) !important;
}

.action-btn:not(.details-btn):hover {
  background: rgba(107, 114, 128, 0.2) !important;
  transform: scale(1.05) !important;
}

.title-text {
  font-size: 0.95rem !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* Administrative Expenses Table Styles */
.expense-table {
  margin-top: 0 !important;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100%;
  overflow-x: auto;
  box-sizing: border-box;
}

/* إصلاح عرض الجدول */
.expense-table .v-data-table {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100%;
  box-sizing: border-box;
}

.expense-table .v-data-table__wrapper {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  overflow-x: auto !important;
  overflow-y: visible !important;
  width: 100%;
  box-sizing: border-box;
}

/* إصلاح عرض الجدول */
.expense-table .v-data-table__wrapper table {
  display: table !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
  min-width: 1000px !important;
  table-layout: auto !important;
  box-sizing: border-box;
}

.expense-table .v-data-table__wrapper table thead {
  display: table-header-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table tbody {
  display: table-row-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table th,
.expense-table .v-data-table__wrapper table td {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper {
  border-radius: 0 0 8px 8px !important;
  overflow: hidden;
  border: 1px solid #e0e0e0 !important;
  border-top: none !important;
}

.expense-table .v-data-table__th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: right !important;
  padding: 10px 8px !important;
  border-bottom: 2px solid #1e3a8a !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* إصلاح عناوين الجدول */
.expense-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: right !important;
  padding: 10px 8px !important;
  border-bottom: 2px solid #1e3a8a !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.expense-table .v-data-table__wrapper table thead tr th .v-data-table__th__content {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  text-align: right !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* قواعد قوية لعناوين الجدول */
.v-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  opacity: 1 !important;
  padding: 10px 8px !important;
  letter-spacing: 0.3px !important;
  visibility: visible !important;
  display: table-cell !important;
  text-align: right !important;
  padding: 16px 12px !important;
  border-bottom: 2px solid #1e3a8a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table .v-data-table__wrapper table thead tr th * {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  opacity: 1 !important;
  visibility: visible !important;
  text-align: right !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* قواعد خاصة بـ Vuetify */
.v-data-table__th {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  opacity: 1 !important;
  visibility: visible !important;
  display: table-cell !important;
  text-align: right !important;
  padding: 10px 8px !important;
  border-bottom: 2px solid #1e3a8a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.v-data-table__th * {
  color: #e3dfdf !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  opacity: 1 !important;
  visibility: visible !important;
  text-align: right !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* إصلاح الجدول نفسه - تبسيط */
.expense-table .v-data-table__wrapper {
  display: block !important;
  overflow: visible !important;
}

.expense-table .v-data-table__wrapper table {
  display: table !important;
  width: 100% !important;
  table-layout: auto !important;
}

.expense-table .v-data-table__wrapper table thead {
  display: table-header-group !important;
}

.expense-table .v-data-table__wrapper table tbody {
  display: table-row-group !important;
}

.expense-table .v-data-table__wrapper table tr {
  display: table-row !important;
}

.expense-table .v-data-table__wrapper table td {
  display: table-cell !important;
}

/* إصلاح عرض الجدول - تبسيط */
.expense-table .v-data-table {
  display: block !important;
}

.expense-table .v-data-table__wrapper {
  display: block !important;
  overflow: visible !important;
}

.expense-table .v-data-table__wrapper table {
  display: table !important;
  width: 100% !important;
}

.expense-table .v-data-table__wrapper table thead {
  display: table-header-group !important;
}

.expense-table .v-data-table__wrapper table tbody {
  display: table-row-group !important;
}

.expense-table .v-data-table__wrapper table tr {
  display: table-row !important;
}

.expense-table .v-data-table__wrapper table th,
.expense-table .v-data-table__wrapper table td {
  display: table-cell !important;
}

/* تحسين ألوان الجدول للوضوح */
.expense-table .v-data-table__td {
  padding: 14px 12px !important;
  border-bottom: 1px solid #ffffff !important;
  text-align: right !important;
  background: rgb(247, 247, 247) !important;
  color: #374151 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
  transition: all 0.3s ease !important;
}

/* إصلاح عرض الجدول */
.expense-table .v-data-table__wrapper {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  overflow: visible !important;
}

.expense-table .v-data-table__wrapper table {
  display: table !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
  table-layout: auto !important;
}

.expense-table .v-data-table__wrapper table thead {
  display: table-header-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table tbody {
  display: table-row-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table th,
.expense-table .v-data-table__wrapper table td {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

/* الصفوف المتناوبة */
.expense-table .v-data-table__wrapper tbody tr:nth-child(even) {
  background: #f7f8f9 !important;
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
  transition: all 0.3s ease !important;
}

.expense-table .v-data-table__wrapper tbody tr:nth-child(odd) {
  background: #f4f1f1 !important;
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
  transition: all 0.3s ease !important;
}

/* تأثير التمرير على الصفوف */
.expense-table .v-data-table__wrapper tbody tr:hover {
  background: #f1f5f9 !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 12px rgba(30, 64, 175, 0.1) !important;
}

/* إصلاح عرض الجدول */
.expense-table .v-data-table__wrapper tbody tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper tbody tr td {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper thead tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper thead tr th {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

/* ألوان النصوص */
.expense-table .serial-number {
  color: #6b7280 !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  background: #f2f2f2 !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
  min-width: 24px !important;
  text-align: center !important;
}

.expense-table .project-name {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.expense-table .date-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.expense-table .cost-text {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  background: #e9ebea !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
}

.expense-table .location-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .notes-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
  line-height: 1.4 !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

/* تحسين ألوان الشريحات */
.expense-table .v-chip {
  font-weight: 600 !important;
  font-size: 0.8rem !important;
  padding: 6px 12px !important;
  border-radius: 20px !important;
  text-transform: none !important;
  letter-spacing: 0.2px !important;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  color: #ffffff !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  transition: all 0.3s ease !important;
}

.expense-table .v-chip:hover {
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15) !important;
}

/* إصلاح عرض الجدول */
.expense-table .v-data-table__wrapper {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  overflow: visible !important;
}

.expense-table .v-data-table__wrapper table {
  display: table !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
  table-layout: auto !important;
}

.expense-table .v-data-table__wrapper table thead {
  display: table-header-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table tbody {
  display: table-row-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.expense-table .v-data-table__wrapper table th,
.expense-table .v-data-table__wrapper table td {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

/* ألوان الشريحات الأصلية */
.expense-table .v-chip--success {
  background: #10b981 !important;
  color: #ffffff !important;
  font-weight: 600 !important;
}

.expense-table .v-chip--info {
  background: #3b82f6 !important;
  color: #ffffff !important;
  font-weight: 600 !important;
}

.expense-table .v-chip--warning {
  background: #f59e0b !important;
  color: #ffffff !important;
  font-weight: 600 !important;
}

.expense-table .v-chip--error {
  background: #ef4444 !important;
  color: #ffffff !important;
  font-weight: 600 !important;
}

/* أزرار الإجراءات */
.expense-table .action-btn {
  color: #b2b2b2 !important;
  transition: all 0.3s ease !important;
  font-weight: 600 !important;
  background: #f5f5f5 !important;
  border: 1px solid #ffffff !important;
  border-radius: 8px !important;
  min-width: 36px !important;
  height: 36px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
}

.expense-table .action-btn:hover {
  color: #059669 !important;
  background: #fafcfd !important;
  border-color: #059669 !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 8px rgba(30, 64, 175, 0.15) !important;
}

/* حدود الجدول */
.expense-table .v-data-table__wrapper {
  border: 1px solid #e2e8f0 !important;
  border-radius: 12px !important;
  overflow: hidden !important;
  background: white !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1) !important;
}

/* حدود الخلايا */
.expense-table .v-data-table__wrapper table td {
  border-right: 1px solid #f1f5f9 !important;
  border-bottom: 1px solid #f1f5f9 !important;
  padding: 14px 12px !important;
  text-align: right !important;
  vertical-align: middle !important;
}

.expense-table .v-data-table__wrapper table td:last-child {
  border-right: none !important;
}

/* حدود الصفوف */
.expense-table .v-data-table__wrapper table tr {
  border-bottom: 1px solid #f1f5f9 !important;
}

.expense-table .v-data-table__wrapper table tr:last-child {
  border-bottom: none !important;
}

.expense-table .v-data-table__wrapper table th:last-child {
  border-right: none !important;
}

/* حدود العناوين */
.expense-table .v-data-table__wrapper table thead tr th {
  border-right: 1px solid rgba(255, 255, 255, 0.2) !important;
}

.expense-table .v-data-table__wrapper table thead tr th:last-child {
  border-right: none !important;
}

/* التمرير */
.expense-table .v-data-table__wrapper::-webkit-scrollbar {
  height: 8px !important;
  width: 8px !important;
}

.expense-table .v-data-table__wrapper::-webkit-scrollbar-track {
  background: #f8fafc !important;
  border-radius: 6px !important;
}

.expense-table .v-data-table__wrapper::-webkit-scrollbar-thumb {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  border-radius: 6px !important;
  transition: all 0.3s ease !important;
}

.expense-table .v-data-table__wrapper::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  box-shadow: 0 2px 4px rgba(5, 150, 105, 0.3) !important;
}

/* تأثيرات التمرير البسيطة */
.expense-table .v-data-table__wrapper tbody tr:hover {
  background: #f8f9fa !important;
}

/* ألوان الروابط البسيطة */
.expense-table a {
  color: #059669 !important;
  text-decoration: none !important;
  font-weight: 500 !important;
}

.expense-table a:hover {
  color: #1d4ed8 !important;
  text-decoration: underline !important;
}

/* ألوان الأيقونات البسيطة */
.expense-table .v-icon {
  color: #6b7280 !important;
}

.expense-table .v-icon:hover {
  color: #059669 !important;
}

/* تحسين ألوان النصوص المهمة */
.expense-table .important-text {
  color: #dc2626 !important;
  font-weight: 700 !important;
}

.expense-table .success-text {
  color: #059669 !important;
  font-weight: 600 !important;
}

.expense-table .warning-text {
  color: #d97706 !important;
  font-weight: 600 !important;
}

/* تحسين ألوان الخلفية */
.expense-table .v-data-table__wrapper {
  background: white !important;
}

/* تحسين ألوان الحدود */
.expense-table .v-data-table__wrapper table {
  border-collapse: separate !important;
  border-spacing: 0 !important;
}

.expense-table .v-data-table__wrapper table td {
  border-right: 1px solid #e5e7eb !important;
}

.expense-table .v-data-table__wrapper table td:last-child {
  border-right: none !important;
}

/* Table Content Styles */
.serial-number {
  font-weight: 600 !important;
  color: #666 !important;
}

.project-name {
  font-weight: 500 !important;
  color: #333 !important;
}

.date-text {
  color: #666 !important;
  font-size: 0.85rem !important;
}

.cost-text {
  font-weight: 600 !important;
  color: #d32f2f !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.location-text {
  color: #666 !important;
  font-size: 0.85rem !important;
}

.notes-text {
  color: #666 !important;
  font-size: 0.85rem !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
  display: inline-block !important;
}

/* ضمان ظهور أزرار الإجراءات */
.expense-table .v-data-table__wrapper table tbody tr td:last-child {
  text-align: center !important;
  padding: 8px !important;
}

.expense-table .v-data-table__wrapper table tbody tr td:last-child .action-btn {
  margin: 0 auto !important;
  display: inline-flex !important;
  visibility: visible !important;
  opacity: 1 !important;
}

/* إصلاح أيقونة الإجراءات */
.expense-table .action-btn .v-icon {
  color: #059669 !important;
  font-size: 18px !important;
  visibility: visible !important;
  opacity: 1 !important;
}

/* تحسين ألوان البحث والفلترة */
.expense-table .search-field .v-field__input,
.expense-table .filter-field .v-field__input {
  color: #1a1a1a !important;
  font-weight: 500 !important;
    font-size: 1rem !important;
  }

.expense-table .search-field .v-field__outline,
.expense-table .filter-field .v-field__outline {
  border-color: #059669 !important;
  border-width: 2px !important;
}

.expense-table .search-field .v-label,
.expense-table .filter-field .v-label {
  color: #059669 !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
}

.expense-table .search-field .v-field--focused .v-field__outline,
.expense-table .filter-field .v-field--focused .v-field__outline {
  border-color: #059669 !important;
  border-width: 2px !important;
}

.expense-table .search-field .v-field--focused .v-label,
.expense-table .filter-field .v-field--focused .v-label {
  color: #059669 !important;
}

/* تحسين أزرار البحث */
.expense-table .search-btn {
  background: #1e40af !important;
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  padding: 12px 24px !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 4px rgba(30, 64, 175, 0.2) !important;
}

.expense-table .search-btn:hover {
  background: #1e3a8a !important;
  box-shadow: 0 4px 8px rgba(30, 64, 175, 0.3) !important;
  transform: translateY(-1px) !important;
}

/* تحسين أزرار الفلترة */
.expense-table .filter-btn {
  background: #059669 !important;
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  padding: 12px 24px !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 4px rgba(5, 150, 105, 0.2) !important;
}

.expense-table .filter-btn:hover {
  background: #047857 !important;
  box-shadow: 0 4px 8px rgba(5, 150, 105, 0.3) !important;
  transform: translateY(-1px) !important;
}

/* تحسين ألوان الفلترة */
.expense-table .v-select .v-field__input {
  color: #1a1a1a !important;
  font-weight: 500 !important;
}

.expense-table .v-select .v-field__outline {
  border-color: #059669 !important;
  border-width: 2px !important;
}

.expense-table .v-select .v-label {
  color: #059669 !important;
  font-weight: 600 !important;
}

.expense-table .v-btn--variant-tonal {
  border-radius: 8px !important;
}

/* Expense Chips */
.expense-table .v-chip--size-small {
  font-weight: 600 !important;
  font-size: 0.8rem !important;
}

/* Expense Cost Display */
.expense-table .text-error {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  font-weight: 700 !important;
    font-size: 1rem !important;
  }

/* Dialog Styles - إصلاح مشكلة النص الأبيض */
.image-style-dialog {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%);
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(5, 150, 105, 0.3);
}

.dialog-header {
  background: linear-gradient(135deg, #059669, #10b981);
  padding: 20px;
  color: white;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-title {
  font-size: 1.5rem;
  font-weight: 700;
}

.close-btn {
  color: white !important;
}

.dialog-body {
  padding: 30px;
  background: #f8fafc;
}

.form-fields {
  margin-bottom: 20px;
}

.form-field {
  margin-bottom: 16px;
}

/* تحسين ألوان الحقول في القائمة المنبثقة */
.form-field .v-field__input {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  background: white !important;
  border-radius: 8px !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
}

/* إصلاح لون النص في جميع الحقول */
.form-field .v-field__input input,
.form-field .v-field__input textarea,
.form-field .v-field__input .v-field__input {
  color: #1a1a1a !important;
  background: white !important;
}

.form-field .v-field__outline {
  border-color: #059669 !important;
  border-width: 2px !important;
}

.form-field .v-label {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  background: white !important;
  padding: 0 8px !important;
}

/* إصلاح لون النص في التسميات */
.form-field .v-field__field .v-label,
.form-field .v-field--focused .v-label {
  color: #059669 !important;
  background: white !important;
}

.form-field .v-field--focused .v-field__outline {
  border-color: #059669 !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(30, 64, 175, 0.1) !important;
}

.form-field .v-field--focused .v-label {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 0.875rem !important;
  background: white !important;
  padding: 0 8px !important;
  margin: 0 !important;
  border-radius: 4px !important;
  text-shadow: none !important;
  letter-spacing: normal !important;
  box-shadow: none !important;
  border: none !important;
  transform: none !important;
  position: static !important;
  z-index: auto !important;
  font-family: inherit !important;
  text-transform: none !important;
}

/* تحسين ألوان القوائم المنسدلة */
.form-field .v-select .v-field__input {
  color: #1a1a1a !important;
  font-weight: 600 !important;
    font-size: 1rem !important;
  padding: 12px 16px !important;
}

.form-field .v-select .v-field__outline {
  border-color: #059669 !important;
  border-width: 2px !important;
}

.form-field .v-select .v-label {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 0.875rem !important;
  background: white !important;
  padding: 0 8px !important;
  margin: 0 !important;
  border-radius: 4px !important;
  text-shadow: none !important;
  letter-spacing: normal !important;
  box-shadow: none !important;
  border: none !important;
  position: static !important;
  z-index: auto !important;
  font-family: inherit !important;
  text-transform: none !important;
}

/* تحسين ألوان رسائل الخطأ */
.form-field .v-messages__message {
  color: #dc2626 !important;
  font-weight: 500 !important;
  font-size: 0.8rem !important;
}

/* تحسين ألوان النصوص في الحقول */
.form-field .v-field__input::placeholder {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

/* تحسين ألوان النصوص المكتوبة */
.form-field .v-field__input input {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

/* إصلاح شامل للنصوص في الحقول */
.form-field input,
.form-field textarea,
.form-field .v-field__input,
.form-field .v-field__input *,
.form-field .v-input__control,
.form-field .v-input__control * {
  color: #1a1a1a !important;
  background: white !important;
}

/* تحسين ألوان النصوص في textarea */
.form-field .v-textarea .v-field__input {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  line-height: 1.5 !important;
}

/* تحسين ألوان الأيقونات */
.form-field .v-icon {
  color: #059669 !important;
}

/* تحسين ألوان القوائم المنسدلة المفتوحة */
.form-field .v-list {
  background: white !important;
  border: 2px solid #059669 !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(5, 150, 105, 0.2) !important;
}

/* إصلاح النصوص في القوائم المنسدلة */
.form-field .v-list *,
.form-field .v-menu *,
.form-field .v-overlay *,
.form-field .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

.form-field .v-list-item {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
}

.form-field .v-list-item:hover {
  background: #f1f5f9 !important;
  color: #059669 !important;
  font-weight: 700 !important;
}

.form-field .v-list-item--active {
  background: #059669 !important;
  color: white !important;
  font-weight: 700 !important;
}

/* إصلاح شامل لجميع النصوص في النافذة */
.image-style-dialog .form-field *,
.image-style-dialog .dialog-body *,
.image-style-dialog .v-field *,
.image-style-dialog .v-input *,
.image-style-dialog .v-select *,
.image-style-dialog .v-textarea * {
  color: #1a1a1a !important;
}

/* استثناء للتسميات */
.image-style-dialog .form-field .v-label {
  color: #059669 !important;
}

/* إصلاح النصوص في الحقول المحددة */
.image-style-dialog .form-field .v-field--focused input,
.image-style-dialog .form-field .v-field--focused textarea,
.image-style-dialog .form-field .v-field--focused .v-field__input,
.image-style-dialog .form-field .v-field--focused .v-field__input * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في القوائم المنسدلة المحددة */
.image-style-dialog .form-field .v-field--focused .v-select__selection,
.image-style-dialog .form-field .v-field--focused .v-field__input {
  color: #1a1a1a !important;
  background: white !important;
}

/* تحسين ألوان النصوص في القوائم المنسدلة */
.form-field .v-list-item .v-list-item-title {
  color: inherit !important;
  font-weight: inherit !important;
  font-size: inherit !important;
}

/* إصلاح نهائي لجميع النصوص */
.dialog-body {
  background: white !important;
}

.dialog-body .v-field,
.dialog-body .v-input,
.dialog-body .v-select,
.dialog-body .v-textarea {
  background: white !important;
}

.dialog-body .v-field__input,
.dialog-body .v-field__input *,
.dialog-body .v-field__field,
.dialog-body .v-field__field * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في القوائم المنسدلة المفتوحة */
.v-overlay__content .v-list,
.v-overlay__content .v-list *,
.v-menu__content .v-list,
.v-menu__content .v-list * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح خاص لقائمة أنواع المصروفات */
.v-list-item,
.v-list-item *,
.v-list-item .v-list-item-title,
.v-list-item .v-list-item-subtitle,
.v-list-item .v-list-item-content,
.v-list-item .v-list-item-content * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح القوائم المنسدلة العامة */
.v-menu .v-list,
.v-menu .v-list *,
.v-menu .v-list-item,
.v-menu .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح القوائم المنسدلة في جميع الحالات */
.v-list,
.v-list *,
.v-list .v-list-item,
.v-list .v-list-item *,
.v-list .v-list-item-title,
.v-list .v-list-item-subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح شامل لجميع القوائم المنسدلة في الصفحة */
.v-select .v-list,
.v-select .v-list *,
.v-select .v-list-item,
.v-select .v-list-item *,
.v-select .v-list-item-title,
.v-select .v-list-item-subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح القوائم المنسدلة في الصفحة الرئيسية */
.data-page .v-list,
.data-page .v-list *,
.data-page .v-list-item,
.data-page .v-list-item *,
.data-page .v-list-item-title,
.data-page .v-list-item-subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح خاص لقوائم أنواع المصروفات */
.v-select[data-type="expense-type"] .v-list,
.v-select[data-type="expense-type"] .v-list *,
.v-select[data-type="expense-type"] .v-list-item,
.v-select[data-type="expense-type"] .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح شامل لجميع القوائم المنسدلة في Vuetify */
.v-list-item,
.v-list-item-title,
.v-list-item-subtitle,
.v-list-item-content,
.v-list-item-content * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح القوائم المنسدلة في جميع الأماكن */
.v-menu__content,
.v-menu__content *,
.v-overlay__content,
.v-overlay__content * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في القوائم المنسدلة المفتوحة */
.v-menu__content .v-list-item,
.v-menu__content .v-list-item *,
.v-overlay__content .v-list-item,
.v-overlay__content .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح نهائي لجميع القوائم المنسدلة */
.v-select__menu,
.v-select__menu *,
.v-select__menu .v-list-item,
.v-select__menu .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح شامل لجميع عناصر القوائم */
.v-list-item__content,
.v-list-item__content *,
.v-list-item__title,
.v-list-item__subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح خاص لصفحة المصروفات */
.expenses-page .v-list,
.expenses-page .v-list *,
.expenses-page .v-list-item,
.expenses-page .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح القوائم المنسدلة في شريط البحث */
.search-card .v-list,
.search-card .v-list *,
.search-card .v-list-item,
.search-card .v-list-item *,
.search-card .v-menu__content,
.search-card .v-menu__content * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح شامل لجميع القوائم المنسدلة في الصفحة */
.v-container .v-list,
.v-container .v-list *,
.v-container .v-list-item,
.v-container .v-list-item *,
.v-container .v-menu__content,
.v-container .v-menu__content * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح نهائي شامل لجميع القوائم المنسدلة */
* .v-list,
* .v-list *,
* .v-list-item,
* .v-list-item *,
* .v-list-item-title,
* .v-list-item-subtitle,
* .v-list-item-content,
* .v-list-item-content * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح خاص للقوائم المنسدلة في Vuetify */
.v-application .v-list,
.v-application .v-list *,
.v-application .v-list-item,
.v-application .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* ========================================
   إصلاح شامل لمشكلة النص الأبيض في القوائم المنسدلة
   - قائمة أنواع المصروفات
   - قائمة الحالات
   - جميع القوائم المنسدلة في الصفحة
   - نافذة الحوار والنماذج
   ======================================== */

/* تحسين وضوح جميع النصوص في الحقول */
.form-field .v-field__input,
.form-field .v-field__input input,
.form-field .v-field__input textarea {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح شامل لجميع العناصر */
.image-style-dialog input,
.image-style-dialog textarea,
.image-style-dialog select,
.image-style-dialog .v-field,
.image-style-dialog .v-input,
.image-style-dialog .v-select,
.image-style-dialog .v-textarea {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في جميع الحالات */
.image-style-dialog .v-field--focused,
.image-style-dialog .v-field--active,
.image-style-dialog .v-field--dirty {
  color: #1a1a1a !important;
  background: white !important;
}

.image-style-dialog .v-field--focused input,
.image-style-dialog .v-field--active input,
.image-style-dialog .v-field--dirty input,
.image-style-dialog .v-field--focused textarea,
.image-style-dialog .v-field--active textarea,
.image-style-dialog .v-field--dirty textarea {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  line-height: 1.5 !important;
}

/* تحسين وضوح التسميات */
.form-field .v-label,
.form-field .v-field__label {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.6rem !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  padding: 10px 24px !important;
  margin: 0 20px !important;
  border-radius: 10px !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.8) !important;
  letter-spacing: 1.2px !important;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.4) !important;
  border: 4px solid #000000 !important;
  position: relative !important;
  z-index: 10 !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  text-transform: uppercase !important;
}

/* تحسين وضوح النصوص في الحقول المحددة */
.form-field .v-field--focused .v-field__input,
.form-field .v-field--focused .v-field__input input,
.form-field .v-field--focused .v-field__input textarea {
  color: #1a1a1a !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

/* تحسين وضوح النصوص في الحقول المملوءة */
.form-field .v-field--filled .v-field__input,
.form-field .v-field--filled .v-field__input input,
.form-field .v-field--filled .v-field__input textarea {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

/* تحسين وضوح التسميات في جميع الحالات */
.form-field .v-field__label,
.form-field .v-label,
.form-field .v-field__label--active,
.form-field .v-label--active {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.6rem !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  padding: 10px 24px !important;
  margin: 0 20px !important;
  border-radius: 10px !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.8) !important;
  letter-spacing: 1.2px !important;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.4) !important;
  border: 4px solid #000000 !important;
  position: relative !important;
  z-index: 10 !important;
  opacity: 1 !important;
  visibility: visible !important;
  display: block !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  text-transform: uppercase !important;
}

.black-list  .v-list-item-title ,.v-list-item--density-default:not(.v-list-item--nav).v-list-item--one-line{
  color: #000000 !important;
}

/* تحسين وضوح التسميات في الحقول المحددة */
.form-field .v-field--focused .v-field__label,
.form-field .v-field--focused .v-label {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.8rem !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  padding: 12px 26px !important;
  margin: 0 22px !important;
  border-radius: 12px !important;
  text-shadow: 0 5px 10px rgba(0, 0, 0, 0.9) !important;
  letter-spacing: 1.4px !important;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.5) !important;
  border: 5px solid #000000 !important;
  transform: scale(1.2) !important;
  position: relative !important;
  z-index: 15 !important;
  opacity: 1 !important;
  visibility: visible !important;
  display: block !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  text-transform: uppercase !important;
}

/* قواعد خاصة بـ Vuetify للتسميات */
.form-field .v-text-field .v-field__label,
.form-field .v-textarea .v-field__label,
.form-field .v-select .v-field__label {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.6rem !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  padding: 10px 24px !important;
  margin: 0 20px !important;
  border-radius: 10px !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.8) !important;
  letter-spacing: 1.2px !important;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.4) !important;
  border: 4px solid #000000 !important;
  position: relative !important;
  z-index: 10 !important;
  opacity: 1 !important;
  visibility: visible !important;
  display: block !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  text-transform: uppercase !important;
}

/* قواعد خاصة بـ Vuetify للتسميات النشطة */
.form-field .v-text-field .v-field__label--active,
.form-field .v-textarea .v-field__label--active,
.form-field .v-select .v-field__label--active {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.8rem !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  padding: 12px 26px !important;
  margin: 0 22px !important;
  border-radius: 12px !important;
  text-shadow: 0 5px 10px rgba(0, 0, 0, 0.9) !important;
  letter-spacing: 1.4px !important;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.5) !important;
  border: 5px solid #000000 !important;
  transform: scale(1.2) !important;
  position: relative !important;
  z-index: 15 !important;
  opacity: 1 !important;
  visibility: visible !important;
  display: block !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  text-transform: uppercase !important;
}

.dialog-actions {
  padding: 20px 30px;
  background: #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.action-btn {
  min-width: 100px;
  font-weight: 600;
  border-radius: 8px !important;
  color: #6b7280 !important;
  background: white !important;
  border: 2px solid #e5e7eb !important;
}

.action-btn:hover {
  color: #059669 !important;
  background: #f0f9ff !important;
  border-color: #059669 !important;
  transform: translateY(-1px) !important;
}

.primary-btn {
  background: linear-gradient(135deg, #059669, #10b981) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3) !important;
}

.primary-btn:hover {
  background: linear-gradient(135deg, #047857, #059669) !important;
  box-shadow: 0 6px 16px rgba(5, 150, 105, 0.4) !important;
  transform: translateY(-2px) !important;
}

/* Animations */
@keyframes star-twinkle {
  0%, 100% { transform: scale(1) rotate(0deg); }
  50% { transform: scale(1.1) rotate(180deg); }
}

@keyframes gradient-animation {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.gradient-animation {
  background-size: 200% 200%;
  animation: gradient-animation 3s ease infinite;
}

@keyframes text-glow {
  0%, 100% { text-shadow: 0 0 20px rgba(255, 255, 255, 0.5); }
  50% { text-shadow: 0 0 30px rgba(255, 255, 255, 0.8); }
}

.text-glow {
  animation: text-glow 2s ease-in-out infinite;
}

@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.fade-in {
  animation: fade-in 0.8s ease-out;
}

@keyframes hover-lift {
  0% { transform: translateY(0); }
  100% { transform: translateY(-8px); }
}

.hover-lift:hover {
  animation: hover-lift 0.3s ease-out forwards;
}

@keyframes card-glow {
  0%, 100% { box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1); }
  50% { box-shadow: 0 15px 40px rgba(0, 0, 0, 0.2); }
}

.card-glow {
  animation: card-glow 3s ease-in-out infinite;
}

@keyframes smooth-transition {
  0% { opacity: 0; transform: scale(0.95); }
  100% { opacity: 1; transform: scale(1); }
}

.smooth-transition {
  animation: smooth-transition 0.5s ease-out;
}

@keyframes icon-glow {
  0%, 100% { filter: drop-shadow(0 0 10px rgba(0, 0, 0, 0.3)); }
  50% { filter: drop-shadow(0 0 20px rgba(0, 0, 0, 0.5)); }
}

.icon-glow {
  animation: icon-glow 2s ease-in-out infinite;
}

@keyframes light-sweep {
  0% { background-position: -200% 0; }
  100% { background-position: 200% 0; }
}

.light-sweep {
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  background-size: 200% 100%;
  animation: light-sweep 2s ease-in-out infinite;
}

.btn-glow {
  position: relative;
  overflow: hidden;
}

.btn-glow::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  transition: left 0.5s;
}

.btn-glow:hover::before {
  left: 100%;
}

/* Responsive Design - Enhanced */
@media (max-width: 1200px) {
  .fullscreen-content {
    max-width: 100%;
    padding: 0 10px;
  }
  
  .stats-row .v-col {
    margin-bottom: 15px;
  }
  
  .search-card .v-row {
    margin: 0;
  }
  
  .search-card .v-col {
    padding: 8px;
  }
}

@media (max-width: 768px) {
  .page-title {
    font-size: 2rem;
  }
  
  .page-subtitle {
    font-size: 1rem;
  }
  
  .page-header {
    padding: 20px;
  }
  
  .stats-row .v-col {
    margin-bottom: 15px;
  }
  
  .search-card .v-row {
    flex-direction: column;
  }
  
  .search-card .v-col {
    width: 100% !important;
    flex: 0 0 100% !important;
    max-width: 100% !important;
    padding: 8px;
  }
  
  .search-btn,
  .add-expense-btn {
    width: 100% !important;
    margin-bottom: 8px;
  }
  
  .dialog-body {
    padding: 20px;
  }
  
  .dialog-actions {
    padding: 15px 20px;
    flex-direction: column;
  }
  
  .action-btn {
    width: 100%;
  }
  
  /* Table Responsive */
  .expense-table .v-data-table__wrapper {
    overflow-x: auto !important;
  }
  
  .expense-table .v-data-table__wrapper table {
    min-width: 800px !important;
  }
  
  .expense-table .v-data-table__wrapper table th,
  .expense-table .v-data-table__wrapper table td {
    white-space: nowrap !important;
    min-width: 100px !important;
  }
}

/* ========================================
   تخصيص عناوين الجداول - Table Headers
   ======================================== */
.project-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: center !important;
  vertical-align: middle !important;
  padding: 10px 8px !important;
  border: none !important;
  box-shadow: 0 3px 12px rgba(4, 120, 87, 0.4) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.project-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.project-table table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}

/* CSS إضافي لضمان التطبيق */
.project-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
  padding: 10px 8px !important;
}

.project-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* CSS لجميع عناوين الجداول */
.v-data-table.project-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
  padding: 10px 8px !important;
}

.v-data-table.project-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* ========================================
   Dialog Form Styles - تنسيقات نموذج الحوار
   ======================================== */
.clean-dialog-card {
  background: #ffffff !important;
  border-radius: 16px !important;
}

.clean-form-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  color: #ffffff !important;
  padding: 20px 24px !important;
}

.clean-form-title {
  color: #ffffff !important;
  font-size: 1.25rem !important;
  font-weight: 600 !important;
  margin: 0 !important;
}

.clean-form-content {
  padding: 24px !important;
  background: #ffffff !important;
}

.clean-form-instruction {
  color: #666666 !important;
  font-size: 0.9rem !important;
  margin-bottom: 24px !important;
  line-height: 1.6 !important;
}

.clean-form-row {
  margin-bottom: 16px !important;
}

.clean-form-field-wrapper {
  margin-bottom: 8px !important;
}

.clean-form-label {
  display: block !important;
  color: #333333 !important;
  font-size: 0.9rem !important;
  font-weight: 500 !important;
  margin-bottom: 8px !important;
}

.required-star {
  color: #f44336 !important;
}

.clean-form-input {
  background: #ffffff !important;
}

.clean-form-input :deep(.v-field) {
  background: #ffffff !important;
  color: #333333 !important;
}

.clean-form-input :deep(.v-field__input) {
  color: #333333 !important;
}

.clean-form-input :deep(.v-field__outline) {
  color: #e0e0e0 !important;
}

.clean-form-input :deep(.v-field--focused .v-field__outline) {
  color: #1976d2 !important;
}

.clean-form-input :deep(input),
.clean-form-input :deep(textarea),
.clean-form-input :deep(.v-select__selection-text) {
  color: #333333 !important;
}

.clean-form-input :deep(input::placeholder),
.clean-form-input :deep(textarea::placeholder) {
  color: #999999 !important;
}

.clean-form-actions {
  padding: 16px 24px !important;
  background: #f5f5f5 !important;
  border-top: 1px solid #e0e0e0 !important;
}

.clean-form-cancel-btn {
  color: #666666 !important;
  border-color: #cccccc !important;
}

.clean-form-continue-btn {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  color: #ffffff !important;
}
</style>

<style>
/* CSS غير محدود النطاق للتأكد من تطبيق التغييرات */
.project-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
  padding: 10px 8px !important;
}

.project-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.project-table table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}

/* CSS إضافي لجميع عناوين الجداول */
.v-data-table.project-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.project-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

/* CSS لضمان التطبيق على جميع عناصر Vuetify */
.v-data-table__wrapper.project-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table__wrapper.project-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* ========================================
   تحسين وضوح بيانات الجدول - Table Data
   ======================================== */
.project-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  background: #ffffff !important;
  border-bottom: 1px solid #e5e7eb !important;
  padding: 1rem 0.8rem !important;
  vertical-align: middle !important;
}

.project-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 700 !important;
}

.project-table table tbody tr:nth-child(even) td {
  background: #f8fafc !important;
}

.project-table table tbody tr:nth-child(even):hover td {
  background: #f1f5f9 !important;
}

/* تحسين وضوح النصوص داخل الخلايا */
.project-table .serial-number {
  color: #374151 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.project-table .project-name {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .date-text {
  color: #374151 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
}

.project-table .cost-text {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
}

.project-table .location-text {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .notes-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

/* ========================================
   تحسين وضوح بيانات الجدول - Table Data
   ======================================== */
.project-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  background: #ffffff !important;
  border-bottom: 1px solid #e5e7eb !important;
  padding: 1rem 0.8rem !important;
  vertical-align: middle !important;
}

.project-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 700 !important;
}

.project-table table tbody tr:nth-child(even) td {
  background: #f8fafc !important;
}

.project-table table tbody tr:nth-child(even):hover td {
  background: #f1f5f9 !important;
}

/* تحسين وضوح النصوص داخل الخلايا */
.project-table .serial-number {
  color: #374151 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.project-table .project-name {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .date-text {
  color: #374151 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
}

.project-table .cost-text {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
}

.project-table .location-text {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .notes-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

/* ========================================
   تحسين وضوح بيانات الجدول - Table Data
   ======================================== */
.project-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  background: #ffffff !important;
  border-bottom: 1px solid #e5e7eb !important;
  padding: 1rem 0.8rem !important;
  vertical-align: middle !important;
}

.project-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 700 !important;
}

.project-table table tbody tr:nth-child(even) td {
  background: #f8fafc !important;
}

.project-table table tbody tr:nth-child(even):hover td {
  background: #f1f5f9 !important;
}

/* تحسين وضوح النصوص داخل الخلايا */
.project-table .serial-number {
  color: #374151 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.project-table .project-name {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .date-text {
  color: #374151 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
}

.project-table .cost-text {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
}

.project-table .location-text {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .notes-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

/* ========================================
   تحسين وضوح بيانات الجدول - Table Data
   ======================================== */
.project-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  background: #ffffff !important;
  border-bottom: 1px solid #e5e7eb !important;
  padding: 1rem 0.8rem !important;
  vertical-align: middle !important;
}

.project-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 700 !important;
}

.project-table table tbody tr:nth-child(even) td {
  background: #f8fafc !important;
}

.project-table table tbody tr:nth-child(even):hover td {
  background: #f1f5f9 !important;
}

/* تحسين وضوح النصوص داخل الخلايا */
.project-table .serial-number {
  color: #374151 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.project-table .project-name {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .date-text {
  color: #374151 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
}

.project-table .cost-text {
  color: #059669 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
}

.project-table .location-text {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-table .notes-text {
  color: #6b7280 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

@media (max-width: 480px) {
  .data-page {
    padding: 10px !important;
  }
  
  .page-header {
    padding: 15px;
    margin-bottom: 20px;
  }
  
  .page-title {
    font-size: 1.5rem;
  }
  
  .stats-row .v-col {
    padding: 8px;
  }
  
  .stat-card {
    padding: 15px !important;
  }
  
  .search-card {
    margin-bottom: 15px;
  }
  
  .search-card .v-card-text {
    padding: 15px !important;
  }
  
  .expense-table .v-data-table__wrapper table th,
  .expense-table .v-data-table__wrapper table td {
    padding: 8px 4px !important;
    font-size: 0.8rem !important;
  }
  
  .expense-table .serial-number,
  .expense-table .project-name,
  .expense-table .date-text,
  .expense-table .cost-text,
  .expense-table .location-text,
  .expense-table .notes-text {
    font-size: 0.8rem !important;
  }
}

/* ========================================
   Animations - نفس صفحة المهندسين
   ======================================== */

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-12px) rotate(2deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
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

@keyframes slideInFromLeft {
  0% {
    transform: translateX(-50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
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

@keyframes gradientFlow {
  0%, 100% {
    background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  }
  50% {
    background: linear-gradient(90deg, #bbdefb 0%, #ffffff 50%, #e3f2fd 100%);
  }
}

@keyframes hoverPulse {
  0% {
    transform: translateY(-8px) scale(1.02);
  }
  50% {
    transform: translateY(-12px) scale(1.05);
  }
  100% {
    transform: translateY(-8px) scale(1.02);
  }
}

@keyframes iconGlow {
  0%, 100% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 20px rgba(255, 255, 255, 0.3));
  }
  50% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 30px rgba(255, 255, 255, 0.6));
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-8px);
  }
}
</style>