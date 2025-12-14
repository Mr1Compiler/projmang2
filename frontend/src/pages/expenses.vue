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
  console.log('البحث عن:', expenseSearchQuery.value)
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

<style scoped>
/* Page Styles */
.data-page {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  padding: 20px;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
}

.fullscreen-content {
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
  box-sizing: border-box;
  overflow-x: hidden;
}

/* Header Styles */
.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #667eea 100%) !important;
  backdrop-filter: blur(20px);
  border-radius: 16px;
  padding: 20px 30px;
  margin-bottom: 20px;
  text-align: center !important;
  border: 2px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
  animation: gradient-animation 8s ease infinite;
  background-size: 200% 200%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.header-top-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 10px;
  width: 100%;
  text-align: center;
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-5px);
  }
}

.page-icon {
  font-size: 2.5rem;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  margin: 0;
  animation: iconFloat 3s ease-in-out infinite;
  vertical-align: middle;
  line-height: 1;
  transition: transform 0.3s ease;
}

.page-title {
  font-size: 2rem;
  font-weight: 800;
  color: rgb(246, 246, 246);
  margin: 0;
  padding: 0;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
  text-align: center !important;
  width: auto;
  line-height: 1.2;
  letter-spacing: 0.5px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.page-subtitle {
  font-size: 1rem;
  color: rgba(243, 240, 240, 0.9);
  margin: 0;
  padding: 0;
  text-align: center !important;
  width: 100%;
  line-height: 1.4;
  letter-spacing: 0.3px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  font-weight: 500;
}

/* Stats Row */
.stats-row {
  margin-bottom: 30px;
  width: 100%;
  overflow-x: hidden;
}

.stats-row .v-row {
  margin: 0;
  width: 100%;
}

.stats-row .v-col {
  padding: 8px;
  box-sizing: border-box;
}

.stat-card {
  background: rgba(229, 226, 226, 0.95);
  border-radius: 16px;
  transition: all 0.3s ease;
  border: 1px solid rgba(252, 252, 252, 0.3);
}

.stat-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.stat-icon {
  transition: all 0.3s ease;
}

.stat-icon:hover {
  transform: scale(1.1);
}

/* Data Table Card */
.data-table-card {
  background: rgba(255, 249, 249, 0.95);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
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
}

/* Search Card Styles */
.search-card {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border: 2px solid rgba(255, 255, 255, 0.2) !important;
  border-radius: 16px !important;
  box-shadow: 
    0 8px 32px rgba(59, 130, 246, 0.25),
    0 4px 16px rgba(37, 99, 235, 0.2) !important;
  width: 100%;
  overflow-x: hidden;
  box-sizing: border-box;
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

.search-field .v-field__input,
.search-field .v-field__input input {
  background: rgba(255, 255, 255, 0.95) !important;
  border-radius: 8px !important;
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.search-field .v-field__outline {
  border-color: rgba(255, 255, 255, 0.5) !important;
  border-width: 2px !important;
}

.search-field .v-field--focused .v-field__outline {
  border-color: rgba(255, 255, 255, 0.8) !important;
  border-width: 3px !important;
}

.search-field .v-label,
.search-field .v-field-label {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.search-field .v-field-label--active,
.search-field .v-field-label--floating {
  color: #000000 !important;
  font-weight: 700 !important;
}

.search-field .v-field__prepend-inner .v-icon {
  color: rgba(255, 255, 255, 0.9) !important;
}

/* حقول الفلترة */
.filter-field .v-field__input,
.filter-field .v-field__input input,
.filter-field .v-field__input select {
  background: rgba(255, 255, 255, 0.95) !important;
  border-radius: 8px !important;
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.filter-field .v-field__outline {
  border-color: rgba(255, 255, 255, 0.5) !important;
  border-width: 2px !important;
}

.filter-field .v-field--focused .v-field__outline {
  border-color: rgba(255, 255, 255, 0.8) !important;
  border-width: 3px !important;
}

.filter-field .v-label,
.filter-field .v-field-label {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.filter-field .v-field-label--active,
.filter-field .v-field-label--floating {
  color: #000000 !important;
  font-weight: 700 !important;
}

.filter-field .v-field__append-inner .v-icon {
  color: rgba(255, 255, 255, 0.9) !important;
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
  padding: 20px 24px !important;
  border-radius: 12px 12px 0 0 !important;
  box-shadow: 
    0 4px 16px rgba(59, 130, 246, 0.3),
    0 2px 8px rgba(37, 99, 235, 0.2) !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2);
}

/* Indigo Title */
.indigo-title {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 1.3rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
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
  font-size: 1.4rem !important;
  font-weight: 800 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
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
</style>