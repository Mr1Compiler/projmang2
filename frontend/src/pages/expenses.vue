<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <div class="header-top-content">
          <h1 class="page-title">المصاريف الإدارية</h1>
          <span class="page-icon">💰</span>
        </div>
        <p class="page-subtitle">إدارة وتتبع جميع المصاريف الإدارية والعامة</p>
      </div>

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="error">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-error mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalExpenses || 5 }}</h3>
            <p class="text-subtitle-1 text-error mb-0">إجمالي المصاريف</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="success">mdi-check-circle</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ activeExpenses || 3 }}</h3>
            <p class="text-subtitle-1 text-success mb-0">مصاريف نشطة</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ pendingExpenses || 2 }}</h3>
            <p class="text-subtitle-1 text-warning mb-0">في الانتظار</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="info">mdi-chart-line</v-icon>
            </div>
            <h3 class="text-h6 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; white-space: nowrap;">{{ formatCurrency(totalCost) || '430,000 د.ع' }}</h3>
            <p class="text-subtitle-1 text-info mb-0">إجمالي التكلفة</p>
          </v-card>
        </v-col>
      </v-row>

      <!-- Search Bar -->
      <v-card class="search-card mb-4" elevation="2">
        <v-card-text class="pa-4">
          <v-row class="align-center">
            <v-col cols="12" md="4">
              <v-text-field
                v-model="expenseSearchQuery"
                label="البحث في المصاريف الإدارية..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="search-field"
                style="background: #f5f5f5;"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-select
                v-model="selectedExpenseType"
                :items="expenseTypeOptions"
                label="نوع المصروف"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="black-list filter-field"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-select
                v-model="selectedStatus"
                :items="statusOptions"
                label="الحالة"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="2">
        <v-btn
                color="error"
                variant="elevated"
          size="large"
                class="search-btn"
                @click="searchExpenses"
        >
                بحث
        </v-btn>
            </v-col>
            <v-col cols="12" md="2">
            <v-btn
                color="success"
                variant="elevated"
          size="large"
                class="add-expense-btn"
                @click="openAddExpenseDialog"
        >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة صنف جديد
            </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Expenses Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title">
          <span class="title-text">المشاريع</span>
        </v-card-title>

        <!-- جدول المشاريع -->
        <v-data-table
          :headers="projectHeaders"
          :items="projectData"
          :search="projectSearchQuery"
          class="project-table"
          :items-per-page="10"
          :loading="false"
          hover
          no-data-text="لا توجد بيانات"
          :header-props="{
            style: 'background: linear-gradient(135deg, #047857 0%, #059669 100%); color: white; font-weight: 700;'
          }"
        >
          <!-- Serial Number Column -->
          <template #item.serial="{ index }">
            <span class="serial-number">{{ index + 1 }}</span>
          </template>

          <!-- Project Name Column -->
          <template #item.projectName="{ item }">
            <span class="project-name">{{ item.projectName }}</span>
          </template>

          <!-- Start Date Column -->
          <template #item.startDate="{ item }">
            <span class="date-text">{{ item.startDate }}</span>
          </template>

          <!-- End Date Column -->
          <template #item.endDate="{ item }">
            <span class="date-text">{{ item.endDate }}</span>
          </template>

          <!-- Cost Column -->
          <template #item.cost="{ item }">
            <span class="cost-text">{{ item.cost }}</span>
          </template>

          <!-- Work Location Column -->
          <template #item.workLocation="{ item }">
            <span class="location-text">{{ item.workLocation }}</span>
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
                @click="viewProjectDetails(item)"
                icon
                class="action-btn details-btn"
                title="عرض التفاصيل"
              >
                <v-icon size="16">mdi-eye</v-icon>
            </v-btn>
            <v-btn
              size="small"
                color="black"
                variant="text"
                @click="editProject(item)"
                icon
                class="action-btn"
                title="تعديل"
              >
                <v-icon size="16">mdi-dots-horizontal</v-icon>
            </v-btn>
            </div>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Administrative Expense Dialog -->
      <v-dialog v-model="expenseDialog" max-width="800" scrollable persistent>
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-currency-usd</v-icon>
                <span class="header-title">{{ isEditingExpense ? 'تعديل المصروف الإداري' : 'إضافة مصروف إداري' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeExpenseDialog"
                class="close-btn"
              />
            </div>
          </div>

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="expenseForm" v-model="expenseFormValid">
              <div class="form-fields">
                <v-row>
                  <v-col cols="12" md="6">
              <v-text-field
                      v-model="expenseForm.projectName"
                      label="اسم المشروع"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم المشروع مطلوب']"
                required
                      class="form-field"
              />
                  </v-col>
                  <v-col cols="12" md="6">
              <v-text-field
                      v-model="expenseForm.cost"
                      label="التكلفة (د.ع)"
                      variant="outlined"
                type="number"
                      :rules="[v => !!v || 'التكلفة مطلوبة', v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر']"
                required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.startDate"
                      label="تاريخ البداية"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ البداية مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.endDate"
                      label="تاريخ الانتهاء"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ الانتهاء مطلوب']"
                required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.workLocation"
                      label="مكان العمل"
                      variant="outlined"
                      :rules="[v => !!v || 'مكان العمل مطلوب']"
                required
                      class="form-field"
              />
                  </v-col>
                  <v-col cols="12" md="6">
              <v-select
                      v-model="expenseForm.expenseType"
                      :items="expenseTypes"
                      label="نوع المصروف"
                      variant="outlined"
                      :rules="[v => !!v || 'نوع المصروف مطلوب']"
                required
                      class="form-field black-list"
              />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
              <v-select
                v-model="expenseForm.status"
                :items="statusOptions"
                label="الحالة"
                      variant="outlined"
                :rules="[v => !!v || 'الحالة مطلوبة']"
                required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <!-- Empty column for spacing -->
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12">
              <v-textarea
                v-model="expenseForm.notes"
                      label="الملاحظات"
                      variant="outlined"
                rows="3"
                      class="form-field"
              />
                  </v-col>
                </v-row>
              </div>
            </v-form>
          </div>

          <!-- Dialog Actions -->
          <div class="dialog-actions">
            <v-btn
              color="grey"
              variant="text"
              @click="closeExpenseDialog"
              class="action-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="error"
              variant="elevated"
              @click="saveExpense"
              :disabled="!expenseFormValid"
              class="action-btn primary-btn"
            >
              <v-icon class="me-2">mdi-content-save</v-icon>
              {{ isEditingExpense ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </div>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

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

// جدول المصاريف الإدارية
const expenseHeaders = [
  { title: 'التسلسل', key: 'serial', sortable: false, width: '80px' },
  { title: 'اسم المشروع', key: 'projectName', sortable: true, width: '180px' },
  { title: 'نوع المصروف', key: 'expenseType', sortable: true, width: '120px' },
  { title: 'تاريخ بدء المشروع', key: 'startDate', sortable: true, width: '130px' },
  { title: 'تاريخ انتهاء المشروع', key: 'endDate', sortable: true, width: '130px' },
  { title: 'التكلفة', key: 'cost', sortable: true, width: '120px' },
  { title: 'مكان العمل', key: 'workLocation', sortable: true, width: '100px' },
  { title: 'الحالة', key: 'status', sortable: true, width: '100px' },
  { title: 'الملاحظات', key: 'notes', sortable: false, width: '150px' },
  { title: 'الإجراءات', key: 'actions', sortable: false, width: '100px' }
]

// نموذج المصاريف الإدارية
const expenseForm = ref({
  projectName: '',
  startDate: '',
  endDate: '',
  cost: '',
  workLocation: '',
  expenseType: '',
  status: 'معلق',
  notes: ''
})

// بيانات المصاريف الإدارية التجريبية
const administrativeExpenses = ref([
  {
    id: 1,
    projectName: 'المشروع الأول',
    startDate: '2022-03-10',
    endDate: '2024-03-07',
    cost: 500000,
    workLocation: 'لعقوبة',
    expenseType: 'تطوير',
    status: 'معتمد',
    notes: 'لايوجد'
  },
  {
    id: 2,
    projectName: 'dfghjkl',
    startDate: '2022-04-03',
    endDate: '2022-04-03',
    cost: 12.313,
    workLocation: 'akjsda',
    expenseType: 'تحديث',
    status: 'معلق',
    notes: 'لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد'
  },
  {
    id: 3,
    projectName: 'gfdhcgh',
    startDate: '2025-08-25',
    endDate: '2025-08-25',
    cost: 2000000000,
    workLocation: 'aqz',
    expenseType: 'بناء',
    status: 'مرفوض',
    notes: 'لايوجد'
  },
  {
    id: 4,
    projectName: 'مشروع تحديث المختبرات',
    startDate: '2024-04-01',
    endDate: '2024-07-01',
    cost: 125000,
    workLocation: 'الرياض',
    expenseType: 'تحديث',
    status: 'معتمد',
    notes: 'تحديث وتطوير المختبرات العلمية'
  },
  {
    id: 5,
    projectName: 'مشروع الأمن السيبراني',
    startDate: '2024-05-01',
    endDate: '2024-08-01',
    cost: 80000,
    workLocation: 'جدة',
    expenseType: 'أمن',
    status: 'معلق',
    notes: 'تطوير أنظمة الأمن السيبراني'
  }
])

// بيانات المشاريع
const projectData = ref([
  {
    id: 1,
    projectName: 'المشروع الأول',
    startDate: '10/03/2022',
    endDate: '07/03/2024',
    cost: '500000',
    workLocation: 'لعقوبة',
    notes: 'لايوجد'
  },
  {
    id: 2,
    projectName: 'dfghjkl',
    startDate: '03/04/2022',
    endDate: '03/04/2022',
    cost: '12.313',
    workLocation: 'akjsda',
    notes: 'لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد لايوجد'
  },
  {
    id: 3,
    projectName: 'gfdhcgh',
    startDate: '25/08/2025',
    endDate: '25/08/2025',
    cost: '2000000000',
    workLocation: 'aqz',
    notes: 'لايوجد'
  }
])

// عناوين جدول المشاريع
const projectHeaders = ref([
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'اسم المشروع', key: 'projectName', sortable: true, align: 'right' },
  { title: 'تاريخ بدء المشروع', key: 'startDate', sortable: true, align: 'center' },
  { title: 'تاريخ انتهاء المشروع', key: 'endDate', sortable: true, align: 'center' },
  { title: 'التكلفة', key: 'cost', sortable: true, align: 'center' },
  { title: 'مكان العمل', key: 'workLocation', sortable: true, align: 'center' },
  { title: 'الملاحظات', key: 'notes', sortable: false, align: 'right' },
  { title: 'الاجراءات', key: 'actions', sortable: false, align: 'center' }
])

// استعلام البحث للمشاريع
const projectSearchQuery = ref('')

// خيارات فلترة المصاريف
const projectFilterOptions = computed(() => {
  const projects = [...new Set(administrativeExpenses.value.map(expense => expense.projectName))]
  return projects.map(project => ({ title: project, value: project }))
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
  'جميع الحالات',
  'معتمد',
  'معلق',
  'مرفوض',
  'مسودة'
])

// إحصائيات المصاريف
const totalExpenses = computed(() => administrativeExpenses.value.length)
const activeExpenses = computed(() => {
  const today = new Date()
  return administrativeExpenses.value.filter(expense => {
    const startDate = new Date(expense.startDate)
    const endDate = new Date(expense.endDate)
    return startDate <= today && endDate >= today
  }).length
})
const pendingExpenses = computed(() => {
  const today = new Date()
  return administrativeExpenses.value.filter(expense => {
    const startDate = new Date(expense.startDate)
    return startDate > today
  }).length
})
const totalCost = computed(() => {
  return administrativeExpenses.value.reduce((sum, expense) => sum + expense.cost, 0)
})

// دوال المصاريف الإدارية
const searchExpenses = () => {
  // دالة البحث - يمكن إضافة منطق البحث المتقدم هنا
}

const openAddExpenseDialog = () => {
  expenseDialog.value = true
  isEditingExpense.value = false
  selectedExpense.value = null
  expenseForm.value = {
    projectName: '',
    startDate: '',
    endDate: '',
    cost: '',
    workLocation: '',
    expenseType: '',
    status: 'معلق',
    notes: ''
  }
}

const closeExpenseDialog = () => {
  expenseDialog.value = false
  isEditingExpense.value = false
  selectedExpense.value = null
  expenseForm.value = {
    projectName: '',
    startDate: '',
    endDate: '',
    cost: '',
    workLocation: '',
    expenseType: '',
    status: 'معلق',
    notes: ''
  }
}

const editExpense = (expense) => {
  selectedExpense.value = expense
  isEditingExpense.value = true
  expenseForm.value = { ...expense }
  expenseDialog.value = true
}

// دالة تعديل المشروع
const editProject = (project) => {
  // يمكن إضافة منطق التعديل هنا
}

// دالة عرض تفاصيل المشروع
const viewProjectDetails = (project) => {
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
  })
}

const deleteExpense = (expense) => {
  const index = administrativeExpenses.value.findIndex(e => e.id === expense.id)
    if (index > -1) {
    administrativeExpenses.value.splice(index, 1)
  }
}

const saveExpense = () => {
  if (isEditingExpense.value) {
    // تحديث المصروف
    const index = administrativeExpenses.value.findIndex(e => e.id === selectedExpense.value.id)
      if (index > -1) {
      administrativeExpenses.value[index] = {
          ...expenseForm.value,
        id: selectedExpense.value.id
        }
      }
    } else {
    // إضافة مصروف جديد
      const newExpense = {
        ...expenseForm.value,
        id: Date.now(),
      cost: parseFloat(expenseForm.value.cost)
    }
    administrativeExpenses.value.push(newExpense)
  }
  closeExpenseDialog()
}

// فلترة المصاريف
const filteredExpenses = computed(() => {
  let filtered = administrativeExpenses.value

  if (selectedProjectFilter.value) {
    filtered = filtered.filter(expense => expense.projectName === selectedProjectFilter.value)
  }

  if (selectedCostRange.value) {
    filtered = filtered.filter(expense => {
      const cost = expense.cost
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
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount).replace('IQD', 'د.ع')
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ar-SA')
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
    'معتمد': 'success',
    'معلق': 'warning',
    'مرفوض': 'error',
    'مسودة': 'grey'
  }
  return colors[status] || 'grey'
}

// دالة الحصول على نص الحالة
const getStatusText = (status) => {
  return status || 'غير محدد'
}

onMounted(() => {
  // تهيئة الصفحة
})
</script>


<style>
/* Import page styles */
@import './styles/expenses.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
