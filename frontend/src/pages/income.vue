<template>
  <div class="fill-height data-page">
    <div>
      <!-- Header Section -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-currency-usd</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">الإيرادات</h1>
              <p class="subtitle">إدارة وتتبع جميع مصادر الإيرادات</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Summary Cards -->
      <div class="cards-container">
        <v-row class="mb-6" no-gutters>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-currency-usd</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ formatCurrency(totalIncome) }}</h3>
                <p class="stat-label">إجمالي الإيرادات</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-calendar-month</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ formatCurrency(monthlyIncome) }}</h3>
                <p class="stat-label">إيرادات هذا الشهر</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-trending-up</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ incomeGrowth }}%</h3>
                <p class="stat-label">نمو الإيرادات</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-info" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-source-branch</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ incomeSources.length }}</h3>
                <p class="stat-label">مصادر الإيرادات</p>
              </div>
            </div>
          </v-card>
        </v-col>
        </v-row>
      </div>

      <!-- Income Sources Table -->
      <div class="table-container">
        <v-card class="data-table-card mb-6">
        <v-card-title class="table-title-header d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" color="white" size="18">mdi-format-list-bulleted</v-icon>
            <span class="title-text">مصادر الإيرادات</span>
          </div>
          <div class="d-flex align-center gap-2">
            <v-btn
              class="export-button btn-glow smooth-transition"
              size="small"
              @click="exportToCSV"
              style="background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
            >
              <v-icon class="me-2" size="18">mdi-download</v-icon>
              تصدير Excel
            </v-btn>
            <v-btn
              v-if="canCreate"
              class="add-button btn-glow light-sweep smooth-transition"
              @click="showAddDialog = true"
              elevation="2"
              color="primary"
              size="small"
              style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
            >
              <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
              إضافة إيراد جديد
            </v-btn>
          </div>
        </v-card-title>
        <v-data-table
          :headers="headers"
          :items="incomeSources"
          :loading="loading"
          :items-per-page="-1"
          hide-default-footer
          class="elevation-1 income-table"
        >
          <template v-slot:item.amount="{ item }">
            <span class="font-weight-bold text-success">{{ formatCurrency(item.amount) }}</span>
          </template>
          <template v-slot:item.incomeDate="{ item }">
            {{ formatDate(item.incomeDate) }}
          </template>
          <template v-slot:item.status="{ item }">
            {{ getStatusText(item.status) }}
          </template>
          <template v-slot:item.actions="{ item }">
            <v-btn
              v-if="canUpdate"
              size="small"
              color="primary"
              @click="editIncome(item)"
            >
              <i class="mdi mdi-pencil"></i>
            </v-btn>
            <v-btn
              v-if="canDelete"
              size="small"
              color="error"
              @click="deleteIncome(item)"
            >
              <i class="mdi mdi-delete"></i>
            </v-btn>
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
      </div>

      <!-- Add/Edit Income Dialog -->
      <v-dialog v-model="showAddDialog" max-width="900" scrollable persistent>
        <v-card class="clean-dialog-card clean-form-card">
          <!-- Header Section -->
          <v-card-title class="clean-dialog-header clean-form-header">
            <h2 class="clean-form-title">
              {{ editingIncome ? 'تعديل الإيراد' : 'معلومات الإيراد' }}
            </h2>
          </v-card-title>

          <!-- Form Content -->
          <v-card-text class="clean-form-content">
            <p class="clean-form-instruction">
              لإتمام {{ editingIncome ? 'تعديل' : 'إضافة' }} الإيراد، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
            </p>

            <v-form ref="form" v-model="valid">
              <v-row class="clean-form-row">
                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      اسم الإيراد <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="incomeForm.name"
                      variant="outlined"
                      density="comfortable"
                      placeholder="أدخل اسم الإيراد"
                      :rules="[v => !!v || 'الاسم مطلوب']"
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
                      v-model.number="incomeForm.amount"
                      type="number"
                      variant="outlined"
                      density="comfortable"
                      placeholder="0"
                      :rules="[v => v > 0 || 'المبلغ يجب أن يكون أكبر من صفر']"
                      required
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <v-row class="clean-form-row">
                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">
                      تاريخ الإيراد <span class="required-star">*</span>
                    </label>
                    <v-menu
                      v-model="incomeDateMenu"
                      :close-on-content-click="false"
                      location="bottom"
                    >
                      <template v-slot:activator="{ props }">
                        <v-text-field
                          v-model="formattedIncomeDate"
                          readonly
                          v-bind="props"
                          variant="outlined"
                          density="comfortable"
                          :rules="[v => !!v || 'التاريخ مطلوب']"
                          required
                          hide-details="auto"
                          class="clean-form-input"
                          prepend-inner-icon="mdi-calendar"
                        />
                      </template>
                      <v-date-picker
                        v-model="selectedIncomeDate"
                        @update:model-value="onIncomeDateSelected"
                        color="primary"
                      />
                    </v-menu>
                  </div>
                </v-col>

                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">النوع</label>
                    <v-text-field
                      v-model="incomeForm.type"
                      variant="outlined"
                      density="comfortable"
                      placeholder="أدخل نوع الإيراد"
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <v-row class="clean-form-row">
                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">الحالة</label>
                    <v-select
                      v-model="incomeForm.status"
                      :items="statusOptions"
                      variant="outlined"
                      density="comfortable"
                      placeholder="اختر الحالة"
                      clearable
                      hide-details="auto"
                      class="clean-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="6" class="clean-form-column">
                  <div class="clean-form-field-wrapper">
                    <label class="clean-form-label">ملاحظات</label>
                    <v-textarea
                      v-model="incomeForm.notes"
                      variant="outlined"
                      rows="3"
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

          <v-card-actions class="clean-form-actions">
            <v-spacer />
            <v-btn
              class="clean-form-cancel-btn"
              variant="outlined"
              @click="closeDialog"
            >
              إلغاء
            </v-btn>
            <v-btn
              class="clean-form-continue-btn"
              variant="elevated"
              :disabled="!valid"
              @click="saveIncome"
            >
              {{ editingIncome ? 'تحديث الإيراد' : 'حفظ الإيراد' }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { listIncome, createIncome, updateIncome, deleteIncome as deleteIncomeApi, getIncomeStats } from '@/api/income'
import { DEFAULT_LIMIT } from '@/constants/pagination'
import { usePermissions } from '@/composables/usePermissions'

// Permissions
const { canCreate, canUpdate, canDelete } = usePermissions('/income')

// ========================================
// متغيرات الحالة الأساسية
// ========================================
const loading = ref(false)
const showAddDialog = ref(false)
const valid = ref(false)
const editingIncome = ref(null)
const searchQuery = ref('')
const selectedCategory = ref('')
const dateRange = ref([])

// ========================================
// نموذج البيانات - aligned with backend DTO
// ========================================
const incomeForm = ref({
  name: '',
  amount: 0,
  type: '',
  incomeDate: new Date().toISOString().split('T')[0],
  status: 'received',
  notes: ''
})

// ========================================
// Date picker state
// ========================================
const incomeDateMenu = ref(false)
const selectedIncomeDate = ref(null)

// Computed property for formatted income date display
const formattedIncomeDate = computed(() => {
  if (!incomeForm.value.incomeDate) return ''
  const date = new Date(incomeForm.value.incomeDate)
  return date.toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })
})

// Handle income date selection from picker
const onIncomeDateSelected = (date) => {
  if (date) {
    const d = new Date(date)
    incomeForm.value.incomeDate = d.toISOString().split('T')[0]
  }
  incomeDateMenu.value = false
}

// ========================================
// البيانات الأساسية
// ========================================
const incomeSources = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(DEFAULT_LIMIT)
const totalItems = ref(0)
const totalPages = ref(0)

const incomeStats = ref({
  total: 0,
  totalAmount: 0,
  pending: 0,
  approved: 0,
  rejected: 0,
  averageAmount: 0
})

// ========================================
// قوائم الاختيار
// ========================================
const incomeCategories = [
  { title: 'رسوم طلابية', value: 'رسوم طلابية', color: 'primary' },
  { title: 'منح', value: 'منح', color: 'success' },
  { title: 'استثمارات', value: 'استثمارات', color: 'info' },
  { title: 'خدمات', value: 'خدمات', color: 'warning' },
  { title: 'تبرعات', value: 'تبرعات', color: 'error' },
  { title: 'أخرى', value: 'أخرى', color: 'grey' }
]

const statusOptions = [
  { title: 'مستلم', value: 'received', color: 'success' },
  { title: 'معلق', value: 'pending', color: 'warning' },
  { title: 'مؤكد', value: 'approved', color: 'info' },
  { title: 'مرفوض', value: 'rejected', color: 'error' }
]

const sourceOptions = [
  { title: 'طلاب', value: 'طلاب' },
  { title: 'حكومي', value: 'حكومي' },
  { title: 'خاص', value: 'خاص' },
  { title: 'داخلي', value: 'داخلي' },
  { title: 'خيري', value: 'خيري' }
]

// ========================================
// عناوين الجدول - aligned with backend DTO
// ========================================
const headers = [
  { title: 'الاسم', key: 'name', align: 'start', sortable: true },
  { title: 'المبلغ', key: 'amount', align: 'center', sortable: true },
  { title: 'النوع', key: 'type', align: 'center', sortable: true },
  { title: 'الحالة', key: 'status', align: 'center', sortable: true },
  { title: 'التاريخ', key: 'incomeDate', align: 'center', sortable: true },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// ========================================
// الخصائص المحسوبة - تستخدم البيانات من الـ API
// ========================================
const totalIncome = computed(() => {
  return incomeStats.value.totalAmount || 0
})

const monthlyIncome = computed(() => {
  // This will show monthly stats when using period=month in the API
  return incomeStats.value.totalAmount || 0
})

const pendingIncome = computed(() => {
  return incomeStats.value.pending || 0
})

const incomeGrowth = computed(() => {
  // Calculate growth based on average - could be enhanced with historical API data
  if (incomeStats.value.total === 0) return 0
  return 0 // Growth would need comparison data from API
})

const filteredIncomeSources = computed(() => {
  let filtered = incomeSources.value

  // فلترة حسب البحث
  if (searchQuery.value) {
    filtered = filtered.filter(item =>
      (item.name || '').toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      (item.type || '').toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      (item.notes || '').toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  // فلترة حسب النوع
  if (selectedCategory.value) {
    filtered = filtered.filter(item => item.type === selectedCategory.value)
  }

  // فلترة حسب التاريخ
  if (dateRange.value && dateRange.value.length === 2) {
    const startDate = new Date(dateRange.value[0])
    const endDate = new Date(dateRange.value[1])
    filtered = filtered.filter(item => {
      const itemDate = new Date(item.incomeDate)
      return itemDate >= startDate && itemDate <= endDate
    })
  }

  return filtered
})

const incomeByCategory = computed(() => {
  const categories = {}
  incomeSources.value
    .filter(item => item.status === 'received')
    .forEach(item => {
      if (!categories[item.type]) {
        categories[item.type] = 0
      }
      categories[item.type] += item.amount
    })
  return categories
})

// ========================================
// الدوال المساعدة
// ========================================
const formatCurrency = (amount) => {
  if (amount == null) return '0 IQD'
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount) + ' IQD'
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

const formatDateForInput = (dateString) => {
  return new Date(dateString).toISOString().split('T')[0]
}

const getCategoryColor = (category) => {
  const cat = incomeCategories.find(c => c.value === category)
  return cat ? cat.color : 'grey'
}

const getStatusColor = (status) => {
  const stat = statusOptions.find(s => s.value === status)
  return stat ? stat.color : 'grey'
}

const getStatusText = (status) => {
  const stat = statusOptions.find(s => s.value === status)
  return stat ? stat.title : status
}

// ========================================
// دوال تحميل البيانات من API
// ========================================
const loadIncome = async (page = currentPage.value) => {
  loading.value = true
  try {
    // Load income list and stats in parallel
    const [response, stats] = await Promise.all([
      listIncome({ page, limit: itemsPerPage.value }),
      getIncomeStats()
    ])
    console.log('Income data received:', response)
    console.log('Income stats received:', stats)

    // Update pagination state
    incomeSources.value = response.data
    totalItems.value = response.total
    totalPages.value = response.totalPages
    currentPage.value = response.page

    if (stats) {
      incomeStats.value = stats
    }
  } catch (err) {
    console.error('Error loading income:', err)
  } finally {
    loading.value = false
  }
}

// Handle page change
const onPageChange = (page) => {
  currentPage.value = page
  loadIncome(page)
}

// ========================================
// دوال إدارة البيانات
// ========================================
const openAddDialog = () => {
  editingIncome.value = null
  incomeForm.value = {
    name: '',
    amount: 0,
    type: '',
    incomeDate: new Date().toISOString().split('T')[0],
    status: 'received',
    notes: ''
  }
  showAddDialog.value = true
}

const editIncome = (item) => {
  editingIncome.value = item
  incomeForm.value = {
    name: item.name || '',
    amount: item.amount || 0,
    type: item.type || '',
    incomeDate: item.incomeDate ? formatDateForInput(item.incomeDate) : new Date().toISOString().split('T')[0],
    status: item.status || 'received',
    notes: item.notes || ''
  }
  showAddDialog.value = true
}

const deleteIncome = async (item) => {
  if (confirm(`هل أنت متأكد من حذف الإيراد "${item.name}"؟`)) {
    try {
      await deleteIncomeApi(item.id)
      await loadIncome()
    } catch (err) {
      console.error('Error deleting income:', err)
    }
  }
}

const saveIncome = async () => {
  if (valid.value) {
    try {
      // Convert date string to ISO datetime format for backend
      const dateValue = incomeForm.value.incomeDate
      const isoDate = dateValue ? new Date(dateValue).toISOString() : new Date().toISOString()

      // Build payload with required fields
      const payload = {
        name: incomeForm.value.name,
        amount: Number(incomeForm.value.amount),
        incomeDate: isoDate
      }

      // Add optional fields if they have values
      if (incomeForm.value.type) {
        payload.type = incomeForm.value.type
      }
      if (incomeForm.value.status) {
        payload.status = incomeForm.value.status
      }
      if (incomeForm.value.notes) {
        payload.notes = incomeForm.value.notes
      }

      if (editingIncome.value) {
        await updateIncome(editingIncome.value.id, payload)
      } else {
        await createIncome(payload)
      }

      await loadIncome()
      closeDialog()
    } catch (err) {
      console.error('Error saving income:', err)
    }
  }
}

const closeDialog = () => {
  showAddDialog.value = false
  editingIncome.value = null
  incomeForm.value = {
    name: '',
    amount: 0,
    type: '',
    incomeDate: new Date().toISOString().split('T')[0],
    status: 'received',
    notes: ''
  }
}

// ========================================
// دوال التصدير والطباعة
// ========================================
const exportToCSV = () => {
  const csvHeaders = ['الاسم', 'المبلغ', 'النوع', 'الحالة', 'التاريخ', 'ملاحظات']
  const csvContent = [
    csvHeaders.join(','),
    ...filteredIncomeSources.value.map(item => [
      `"${item.name || ''}"`,
      item.amount,
      `"${item.type || ''}"`,
      `"${getStatusText(item.status)}"`,
      `"${formatDate(item.incomeDate)}"`,
      `"${item.notes || ''}"`
    ].join(','))
  ].join('\n')

  const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = `income-report-${new Date().toISOString().split('T')[0]}.csv`
  link.click()
}

const printReport = () => {
  window.print()
}

// ========================================
// دوال الإحصائيات المتقدمة
// ========================================
const getIncomeTrend = (months = 6) => {
  const trends = []
  const currentDate = new Date()

  for (let i = months - 1; i >= 0; i--) {
    const date = new Date(currentDate.getFullYear(), currentDate.getMonth() - i, 1)
    const monthIncome = incomeSources.value
      .filter(item => {
        const itemDate = new Date(item.incomeDate)
        return itemDate.getMonth() === date.getMonth() &&
               itemDate.getFullYear() === date.getFullYear() &&
               item.status === 'received'
      })
      .reduce((sum, item) => sum + item.amount, 0)

    trends.push({
      month: date.toLocaleDateString('en-US', { month: 'short' }),
      amount: monthIncome
    })
  }

  return trends
}

const getTopCategories = (limit = 5) => {
  return Object.entries(incomeByCategory.value)
    .sort(([,a], [,b]) => b - a)
    .slice(0, limit)
    .map(([category, amount]) => ({ category, amount }))
}

// ========================================
// دورة الحياة
// ========================================
onMounted(() => {
  loadIncome()
})
</script>

<style scoped>
/* ========================================
   صفحة الإيرادات - CSS منظم ومتناسق
   ======================================== */

/* صفحة البيانات العامة */
.data-page {
  background: #ffffff !important;
  color: var(--text-dark);
  min-height: 100vh;
  padding: 0 !important;
  max-width: 100% !important;
  margin: 0 !important;
  width: 100% !important;
  overflow-x: hidden;
}

/* حاويات المحتوى */
.cards-container {
  padding: 0 2rem 0 2rem;
  margin-top: 1.5rem;
}

.action-buttons {
  padding: 0 2rem;
}

.table-container {
  padding: 0 2rem;
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

.header-icon {
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.header-icon i {
  font-size: 2.5rem;
  color: white;
}

.header-text {
  flex: 1;
}

.page-title {
  color: white !important;
  font-weight: 800 !important;
  font-size: 1.5rem !important;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.5px;
}

.page-subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 1rem !important;
  font-weight: 400;
  margin: 0;
  text-align: center;
}

.header-stats {
  position: relative;
  z-index: 2;
  text-align: left;
  min-width: 300px;
  margin-left: auto;
}

.stat-item {
  background: rgba(255, 255, 255, 0.15);
  padding: 1rem 1.5rem;
  border-radius: 12px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.stat-label {
  display: block;
  font-size: 0.9rem;
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: 0.5rem;
}

.stat-value {
  display: block;
  font-size: 1.2rem;
  font-weight: 600;
  color: white;
  direction: ltr;
  text-align: left;
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

.stat-card-success .stat-icon {
  color: #10b981 !important;
}

.stat-card-warning .stat-icon {
  color: #f59e0b !important;
}

.stat-card-info .stat-icon {
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

/* دعم الكروتات القديمة */
.stat-card {
  background: white !important;
  border-radius: 20px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  min-height: 180px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08) !important;
  border: 2px solid rgba(0, 0, 0, 0.15) !important;
}

.stat-card:hover {
  transform: translateY(-12px) scale(1.03);
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15) !important;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem 1.5rem 1rem 1.5rem;
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
}

.card-icon {
  width: 50px;
  height: 50px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  color: white;
}

.card-icon.success {
  background: linear-gradient(135deg, #10b981, #059669);
}

.card-icon.info {
  background: linear-gradient(135deg, #3b82f6, #1d4ed8);
}

.card-icon.warning {
  background: linear-gradient(135deg, #f59e0b, #d97706);
}

.card-icon.primary {
  background: linear-gradient(135deg, #8b5cf6, #7c3aed);
}

.card-title {
  font-size: 1rem;
  font-weight: 600;
  color: #374151;
  flex: 1;
}

.card-content {
  padding: 1rem 1.5rem 1.5rem 1.5rem;
}

.card-value {
  font-size: 1.6rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
  direction: ltr;
  text-align: left;
}

.card-value.success {
  color: #10b981;
}

.card-value.info {
  color: #3b82f6;
}

.card-value.warning {
  color: #f59e0b;
}

.card-value.primary {
  color: #8b5cf6;
}

.card-subtitle {
  font-size: 0.9rem;
  color: #6b7280;
  font-weight: 500;
}

/* تأثيرات خاصة لكل لون */
.stat-card:nth-child(1)::before {
  background: var(--gradient-success);
}

.stat-card:nth-child(2)::before {
  background: var(--gradient-info);
}

.stat-card:nth-child(3)::before {
  background: var(--gradient-warning);
}

.stat-card:nth-child(4)::before {
  background: var(--gradient-primary);
}

/* بطاقة الجدول المحسنة */
.data-table-card {
  background: var(--bg-white) !important;
  border-radius: 16px !important;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08) !important;
}

.data-table-card .v-card-title {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 16px 16px 0 0;
  border-bottom: 1px solid #e2e8f0;
  font-weight: 600 !important;
  color: var(--text-dark) !important;
}

.table-title-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  padding: 8px 12px !important;
  border-radius: 8px 8px 0 0 !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  min-height: 36px !important;
  margin-bottom: 1rem !important;
}

.table-title-header .title-text {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.data-table-card .v-data-table {
  background: transparent !important;
  margin-top: 1rem !important;
}

.data-table-card .v-data-table-header {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
}

.data-table-card .v-data-table-header th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
  text-transform: none !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2) !important;
}

.data-table-card .v-data-table tbody tr {
  border-bottom: 1px solid #f1f5f9 !important;
}

.data-table-card .v-data-table tbody tr:hover {
  background: #f8fafc !important;
}

/* ========================================
   إصلاح لون النص داخل جدول الإيرادات
   - تغيير لون النص من الأبيض إلى الأسود الداكن
   - تحسين وضوح جميع النصوص في الجدول
   - الحفاظ على النصوص الخضراء للمبالغ
   ======================================== */

/* ========================================
   إصلاح محاذاة رؤوس الجدول في صفحة الإيرادات
   - إصلاح النصوص المائلة في رؤوس الجدول
   - محاذاة النصوص في المنتصف
   - إزالة التحويلات غير المرغوب فيها
   - ضمان التنسيق الصحيح للعناوين
   ======================================== */

/* إصلاح محاذاة رؤوس الجدول */
.data-table-card .v-data-table-header th,
.data-table-card .v-data-table-header th *,
.data-table-card .v-data-table thead th,
.data-table-card .v-data-table thead th * {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
  text-transform: none !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2) !important;
}

/* إصلاح لون النص داخل الجدول */
.data-table-card .v-data-table tbody td {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
  text-align: center !important;
  vertical-align: middle !important;
}

.data-table-card .v-data-table tbody td * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في خلايا الجدول */
.data-table-card .v-data-table .v-data-table__td {
  color: #1a1a1a !important;
}

.data-table-card .v-data-table .v-data-table__td * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في القوالب المخصصة */
.data-table-card .v-data-table .font-weight-bold {
  color: #1a1a1a !important;
}

.data-table-card .v-data-table .text-success {
  color: #059669 !important;
}

/* إصلاح شامل لجميع النصوص في الجدول */
.data-table-card .v-data-table tbody tr td,
.data-table-card .v-data-table tbody tr td *,
.data-table-card .v-data-table tbody tr td span,
.data-table-card .v-data-table tbody tr td div,
.data-table-card .v-data-table tbody tr td p {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في أزرار الإجراءات */
.data-table-card .v-data-table .v-btn {
  color: white !important;
}

.data-table-card .v-data-table .v-btn .v-icon {
  color: white !important;
}

/* إصلاح النصوص في القوالب المخصصة */
.data-table-card .v-data-table template,
.data-table-card .v-data-table template * {
  color: #1a1a1a !important;
}

/* إصلاح شامل لجميع الجداول في الصفحة */
.v-data-table tbody td,
.v-data-table tbody td *,
.v-data-table tbody tr td,
.v-data-table tbody tr td *,
.v-data-table .v-data-table__td,
.v-data-table .v-data-table__td * {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
  text-align: center !important;
  vertical-align: middle !important;
}

/* إصلاح رؤوس جميع الجداول في الصفحة */
.v-data-table thead th,
.v-data-table thead th *,
.v-data-table .v-data-table-header th,
.v-data-table .v-data-table-header th * {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
  text-transform: none !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2) !important;
}

/* إصلاح النصوص في جميع خلايا الجدول */
.v-data-table tbody tr td span,
.v-data-table tbody tr td div,
.v-data-table tbody tr td p,
.v-data-table tbody tr td button {
  color: #1a1a1a !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
}

/* إصلاح نهائي شامل لجميع عناصر الجدول */
.data-page .v-data-table,
.data-page .v-data-table *,
.data-page .v-data-table tbody,
.data-page .v-data-table tbody *,
.data-page .v-data-table tbody tr,
.data-page .v-data-table tbody tr *,
.data-page .v-data-table tbody tr td,
.data-page .v-data-table tbody tr td * {
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
}

/* إصلاح رؤوس الجدول في الصفحة الرئيسية */
.data-page .v-data-table thead,
.data-page .v-data-table thead *,
.data-page .v-data-table thead th,
.data-page .v-data-table thead th * {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
  text-transform: none !important;
}

/* تنسيق شامل لعناوين جدول الإيرادات */
.income-table :deep(.v-data-table-header th),
.income-table :deep(.v-data-table__wrapper table thead tr th),
.income-table :deep(.v-data-table__wrapper table thead tr th *) {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  padding: 12px 10px !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2) !important;
}

.income-table :deep(.v-data-table-header th *),
.income-table :deep(.v-data-table-header th span),
.income-table :deep(.v-data-table-header th div),
.income-table :deep(.v-data-table-header th .v-data-table-header__content),
.income-table :deep(.v-data-table-header th .v-data-table-header__content *),
.income-table :deep(.v-data-table__wrapper table thead tr th *),
.income-table :deep(.v-data-table__wrapper table thead tr th span),
.income-table :deep(.v-data-table__wrapper table thead tr th div),
.income-table :deep(.v-data-table-header th i),
.income-table :deep(.v-data-table-header th .v-icon),
.income-table :deep(.v-data-table-header th .v-data-table-header__sort-badge),
.income-table :deep(.v-data-table-header th .v-data-table-header__sort-icon) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
}

/* ضمان أن جميع النصوص في عناوين الجدول بيضاء */
.data-table-card .income-table :deep(.v-data-table-header),
.data-table-card .income-table :deep(.v-data-table-header th),
.data-table-card .income-table :deep(.v-data-table-header th *),
.data-table-card .income-table :deep(.v-data-table-header th span),
.data-table-card .income-table :deep(.v-data-table-header th div),
.data-table-card .income-table :deep(.v-data-table-header th .v-data-table-header__content),
.data-table-card .income-table :deep(.v-data-table-header th .v-data-table-header__content *),
.data-table-card .income-table :deep(.v-data-table__wrapper table thead),
.data-table-card .income-table :deep(.v-data-table__wrapper table thead tr),
.data-table-card .income-table :deep(.v-data-table__wrapper table thead tr th),
.data-table-card .income-table :deep(.v-data-table__wrapper table thead tr th *) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
}

/* قواعد شاملة إضافية لضمان اللون الأبيض */
.data-table-card :deep(.v-data-table-header),
.data-table-card :deep(.v-data-table-header th),
.data-table-card :deep(.v-data-table-header th *),
.data-table-card :deep(.v-data-table-header th span),
.data-table-card :deep(.v-data-table-header th div),
.data-table-card :deep(.v-data-table-header th p),
.data-table-card :deep(.v-data-table-header th label),
.data-table-card :deep(.v-data-table-header th .v-data-table-header__content),
.data-table-card :deep(.v-data-table-header th .v-data-table-header__content *),
.data-table-card :deep(.v-data-table-header th .v-data-table-header__content span),
.data-table-card :deep(.v-data-table-header th .v-data-table-header__content div),
.data-table-card :deep(.v-data-table__wrapper table thead),
.data-table-card :deep(.v-data-table__wrapper table thead tr),
.data-table-card :deep(.v-data-table__wrapper table thead tr th),
.data-table-card :deep(.v-data-table__wrapper table thead tr th *),
.data-table-card :deep(.v-data-table__wrapper table thead tr th span),
.data-table-card :deep(.v-data-table__wrapper table thead tr th div),
.data-table-card :deep(.v-data-table__wrapper table thead tr th p),
.data-table-card :deep(.v-data-table__wrapper table thead tr th label),
.data-table-card :deep(.v-data-table__th),
.data-table-card :deep(.v-data-table__th *),
.data-table-card :deep(.v-data-table__th span),
.data-table-card :deep(.v-data-table__th div) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
}

/* تطبيق اللون الأبيض على أيقونات وأزرار في عناوين الجدول */
.data-table-card :deep(.v-data-table-header th .v-icon),
.data-table-card :deep(.v-data-table-header th i),
.data-table-card :deep(.v-data-table-header th svg),
.data-table-card :deep(.v-data-table-header th .v-data-table-header__sort-badge),
.data-table-card :deep(.v-data-table-header th .v-data-table-header__sort-icon),
.data-table-card :deep(.v-data-table__wrapper table thead tr th .v-icon),
.data-table-card :deep(.v-data-table__wrapper table thead tr th i),
.data-table-card :deep(.v-data-table__wrapper table thead tr th svg) {
  color: #ffffff !important;
  fill: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* إصلاح خاص لصفحة الإيرادات */
.income-page .v-data-table tbody td,
.income-page .v-data-table tbody td * {
  color: #1a1a1a !important;
}

/* إصلاح نهائي شامل لجميع النصوص في الجدول */
.data-page .v-data-table,
.data-page .v-data-table *,
.data-page .v-data-table tbody,
.data-page .v-data-table tbody *,
.data-page .v-data-table tbody tr,
.data-page .v-data-table tbody tr *,
.data-page .v-data-table tbody tr td,
.data-page .v-data-table tbody tr td * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع العناصر */
.data-page .v-data-table .font-weight-bold,
.data-page .v-data-table .text-success,
.data-page .v-data-table span,
.data-page .v-data-table div,
.data-page .v-data-table p {
  color: #1a1a1a !important;
}

/* استثناء للنصوص الخضراء (المبالغ) */
.data-page .v-data-table .text-success {
  color: #059669 !important;
}

/* أزرار الإجراءات */
.action-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.add-button {
  background: linear-gradient(135deg, #10b981, #059669) !important;
  color: white !important;
  border-radius: 16px !important;
  padding: 16px 32px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  box-shadow: 0 8px 24px rgba(16, 185, 129, 0.3) !important;
  transition: all 0.3s ease !important;
  font-size: 1rem !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-align: center !important;
}

/* تم نقل أنماط .add-button و .export-button الأساسية إلى common-components.css */

.add-button:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 32px rgba(16, 185, 129, 0.4) !important;
}

.export-button {
  background: linear-gradient(135deg, #3b82f6, #1d4ed8) !important;
  color: white !important;
  border-radius: 16px !important;
  padding: 16px 32px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  box-shadow: 0 8px 24px rgba(59, 130, 246, 0.3) !important;
  transition: all 0.3s ease !important;
  font-size: 1rem !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-align: center !important;
}


.export-button:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 32px rgba(59, 130, 246, 0.4) !important;
}

/* نافذة الحوار المحسنة */
.v-dialog .v-card {
  border-radius: 16px !important;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.2) !important;
}

.v-dialog .v-card-title {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 16px 16px 0 0;
  border-bottom: 1px solid #e2e8f0;
  font-weight: 600 !important;
  color: #1a1a1a !important;
}

/* ========================================
   إصلاح لون النصوص في نافذة إضافة الإيراد
   - تغيير لون النص من الأبيض إلى الأسود الداكن
   - تحسين وضوح جميع النصوص في النافذة المنبثقة
   - إصلاح التسميات والحقول والقوائم المنسدلة
   ======================================== */

/* إصلاح النصوص في النافذة المنبثقة */
.v-dialog .v-card-text,
.v-dialog .v-card-text *,
.v-dialog .v-form,
.v-dialog .v-form * {
  color: #1a1a1a !important;
}

/* إصلاح التسميات في النافذة المنبثقة */
.v-dialog .v-label,
.v-dialog .v-label *,
.v-dialog .v-field__label,
.v-dialog .v-field__label * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في الحقول */
.v-dialog .v-text-field .v-field__input,
.v-dialog .v-textarea .v-field__input,
.v-dialog .v-select .v-field__input,
.v-dialog .v-field__input,
.v-dialog .v-field__input * {
  color: #1a1a1a !important;
}

/* إصلاح خلفية الحقول */
.v-dialog .v-text-field .v-field,
.v-dialog .v-textarea .v-field,
.v-dialog .v-select .v-field {
  background: white !important;
  border: 1px solid #e2e8f0 !important;
}

/* إصلاح القوائم المنسدلة في النافذة المنبثقة */
.v-dialog .v-menu__content,
.v-dialog .v-menu__content *,
.v-dialog .v-overlay__content,
.v-dialog .v-overlay__content * {
  color: #1a1a1a !important;
  background: white !important;
}

.v-dialog .v-list,
.v-dialog .v-list *,
.v-dialog .v-list-item,
.v-dialog .v-list-item *,
.v-dialog .v-list-item-title,
.v-dialog .v-list-item-subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في جميع عناصر النافذة */
.v-dialog .v-card-text span,
.v-dialog .v-card-text div,
.v-dialog .v-card-text p,
.v-dialog .v-card-text label {
  color: #1a1a1a !important;
}

/* إصلاح رسائل التحقق */
.v-dialog .v-messages,
.v-dialog .v-messages *,
.v-dialog .v-message {
  color: #1a1a1a !important;
}

/* إصلاح شامل لجميع النصوص في النافذة المنبثقة */
.v-dialog *,
.v-dialog * *,
.v-dialog .v-card *,
.v-dialog .v-card * * {
  color: #1a1a1a !important;
}

/* إصلاح خاص للقوائم المنسدلة */
.v-dialog .v-select__menu,
.v-dialog .v-select__menu *,
.v-dialog .v-select__menu .v-list-item,
.v-dialog .v-select__menu .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في جميع الحقول */
.v-dialog .v-input,
.v-dialog .v-input *,
.v-dialog .v-text-field,
.v-dialog .v-text-field *,
.v-dialog .v-textarea,
.v-dialog .v-textarea *,
.v-dialog .v-select,
.v-dialog .v-select * {
  color: #1a1a1a !important;
}

/* إصلاح خلفية النافذة المنبثقة */
.v-dialog .v-card {
  background: white !important;
}

.v-dialog .v-card-text {
  background: white !important;
}

/* إصلاح نهائي شامل لجميع النصوص في النافذة المنبثقة */
.v-dialog .v-application,
.v-dialog .v-application *,
.v-dialog .v-application .v-card,
.v-dialog .v-application .v-card * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع العناصر */
.v-dialog .v-card-title span,
.v-dialog .v-card-title *,
.v-dialog .v-card-actions,
.v-dialog .v-card-actions * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في الأزرار (استثناء) */
.v-dialog .v-btn {
  color: white !important;
}

.v-dialog .v-btn .v-btn__content {
  color: white !important;
}

/* أزرار النافذة */
.dialog-actions .v-btn {
  border-radius: 8px !important;
  text-transform: none !important;
  font-weight: 500 !important;
}

.dialog-actions .v-btn--variant-elevated {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15) !important;
}

/* التصميم المتجاوب */
@media (max-width: 960px) {
  .stat-card {
    min-height: 180px;
  }
  
  .stat-card h3 {
    font-size: 2rem !important;
  }
  
  .stat-card .stat-icon {
    font-size: 2.5rem !important;
  }
  
  .page-title {
    font-size: 2rem !important;
  }
}

@media (max-width: 600px) {
  .stat-card {
    min-height: 160px;
    padding: 16px !important;
  }
  
  .stat-card h3 {
    font-size: 1.8rem !important;
  }
  
  .stat-card .stat-icon {
    font-size: 2rem !important;
  }
  
  .page-title {
    font-size: 1.8rem !important;
  }
  
  .page-subtitle {
    font-size: 1rem !important;
  }
}

/* تحسينات إضافية */
.v-card {
  border-radius: 16px !important;
}

.v-btn {
  border-radius: 8px !important;
  text-transform: none !important;
  font-weight: 500 !important;
}

.v-text-field .v-field {
  border-radius: 8px !important;
}

.v-select .v-field {
  border-radius: 8px !important;
}

.v-textarea .v-field {
  border-radius: 8px !important;
}

/* تأثيرات النصوص */
.text-success {
  color: var(--success-color) !important;
  font-weight: 600 !important;
}

.text-info {
  color: var(--info-color) !important;
  font-weight: 600 !important;
}

.text-warning {
  color: var(--warning-color) !important;
  font-weight: 600 !important;
}

.text-primary {
  color: var(--primary-color) !important;
  font-weight: 600 !important;
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
