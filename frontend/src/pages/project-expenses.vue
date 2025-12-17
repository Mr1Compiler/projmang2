<template>
  <v-container fluid class="project-expenses-page">
    <!-- Header Section -->
    <PageHeader
      :title="projectName"
      subtitle="تفاصيل مصاريف المشروع"
      mdi-icon="mdi-cash-multiple"
    />

    <v-row class="mb-6">
      <v-col cols="12">
        <div class="d-flex justify-end gap-2">
          <v-btn
            color="success"
            variant="elevated"
            @click="addExpense"
            class="add-btn-header"
            size="large"
          >
            <v-icon left>mdi-plus</v-icon>
            إضافة مصروف
          </v-btn>
          <v-btn
            color="primary"
            variant="outlined"
            @click="goBack"
            class="back-btn"
          >
            <v-icon left>mdi-arrow-right</v-icon>
            العودة
          </v-btn>
        </div>
      </v-col>
    </v-row>

    <!-- Project Info Card -->
    <v-row class="mb-6">
      <v-col cols="12">
        <v-card class="info-card" elevation="2">
          <v-card-title class="info-title">
            <v-icon left color="primary">mdi-information</v-icon>
            معلومات المشروع
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12" md="3">
                <div class="info-item">
                  <span class="info-label">تاريخ البداية:</span>
                  <span class="info-value">{{ projectInfo.startDate }}</span>
                </div>
              </v-col>
              <v-col cols="12" md="3">
                <div class="info-item">
                  <span class="info-label">تاريخ النهاية:</span>
                  <span class="info-value">{{ projectInfo.endDate }}</span>
                </div>
              </v-col>
              <v-col cols="12" md="3">
                <div class="info-item">
                  <span class="info-label">التكلفة الإجمالية:</span>
                  <span class="info-value cost">{{ projectInfo.cost }} د.ع</span>
                </div>
              </v-col>
              <v-col cols="12" md="3">
                <div class="info-item">
                  <span class="info-label">مكان العمل:</span>
                  <span class="info-value">{{ projectInfo.workLocation }}</span>
                </div>
              </v-col>
            </v-row>
            <v-row v-if="projectInfo.notes">
              <v-col cols="12">
                <div class="info-item">
                  <span class="info-label">الملاحظات:</span>
                  <span class="info-value">{{ projectInfo.notes }}</span>
                </div>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Statistics Cards -->
    <v-row class="mb-6">
      <v-col cols="12" md="3">
        <v-card class="stat-card total-expenses" elevation="2">
          <v-card-text class="text-center">
            <v-icon size="40" color="primary">mdi-currency-usd</v-icon>
            <h3 class="stat-number">{{ totalExpenses }}</h3>
            <p class="stat-label">إجمالي المصاريف</p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="stat-card remaining-budget" elevation="2">
          <v-card-text class="text-center">
            <v-icon size="40" color="success">mdi-wallet</v-icon>
            <h3 class="stat-number">{{ remainingBudget }}</h3>
            <p class="stat-label">المتبقي من الميزانية</p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="stat-card expense-count" elevation="2">
          <v-card-text class="text-center">
            <v-icon size="40" color="info">mdi-receipt</v-icon>
            <h3 class="stat-number">{{ projectExpenses.length }}</h3>
            <p class="stat-label">عدد المصاريف</p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="stat-card budget-usage" elevation="2">
          <v-card-text class="text-center">
            <v-icon size="40" color="warning">mdi-chart-pie</v-icon>
            <h3 class="stat-number">{{ budgetUsagePercentage }}%</h3>
            <p class="stat-label">نسبة الاستخدام</p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Expenses Table -->
    <v-row>
      <v-col cols="12">
        <v-card class="expenses-table-card" elevation="2">
          <v-card-title class="table-title">
            <v-icon left color="primary">mdi-table</v-icon>
            مصاريف المشروع
            <v-spacer></v-spacer>
            <v-text-field
              v-model="searchQuery"
              prepend-inner-icon="mdi-magnify"
              label="البحث في المصاريف"
              variant="outlined"
              density="compact"
              hide-details
              class="search-field"
            ></v-text-field>
          </v-card-title>
          <v-card-text>
            <!-- Add Expense Button (Alternative) -->
            <div class="add-expense-section mb-4">
              <v-btn
                color="primary"
                variant="elevated"
                @click="addExpense"
                class="add-expense-btn"
                size="large"
              >
                <v-icon left>mdi-plus</v-icon>
                إضافة صنف جديد
              </v-btn>
            </div>

            <v-data-table
              :headers="expenseHeaders"
              :items="projectExpenses"
              :search="searchQuery"
              class="expenses-table"
              :items-per-page="10"
              :loading="false"
              hover
              no-data-text="لا توجد مصاريف لهذا المشروع"
              :header-props="{
                style: 'background: linear-gradient(135deg, #047857 0%, #059669 100%); color: white; font-weight: 700;'
              }"
            >
              <!-- Serial Number Column -->
              <template #item.serial="{ index }">
                <span class="serial-number">{{ index + 1 }}</span>
              </template>

              <!-- Expense Type Column -->
              <template #item.type="{ item }">
                <v-chip
                  :color="getExpenseTypeColor(item.type)"
                  size="small"
                  class="expense-type-chip"
                >
                  {{ getExpenseTypeLabel(item.type) }}
                </v-chip>
              </template>

              <!-- Amount Column -->
              <template #item.amount="{ item }">
                <span class="amount-text">{{ formatAmount(item.amount) }}</span>
              </template>

              <!-- Date Column -->
              <template #item.date="{ item }">
                <span class="date-text">{{ item.date }}</span>
              </template>

              <!-- Status Column -->
              <template #item.status="{ item }">
                <v-chip
                  :color="getStatusColor(item.status)"
                  size="small"
                  class="status-chip"
                >
                  {{ getStatusLabel(item.status) }}
                </v-chip>
              </template>

              <!-- Actions Column -->
              <template #item.actions="{ item }">
                <div class="action-buttons">
                  <v-btn
                    size="small"
                    color="primary"
                    variant="text"
                    @click="viewExpenseDetails(item)"
                    icon
                    class="action-btn details-btn"
                    title="عرض التفاصيل"
                  >
                    <v-icon size="16">mdi-eye</v-icon>
                  </v-btn>
                  <v-btn
                    size="small"
                    color="success"
                    variant="text"
                    @click="editExpense(item)"
                    icon
                    class="action-btn edit-btn"
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
                    class="action-btn delete-btn"
                    title="حذف"
                  >
                    <v-icon size="16">mdi-delete</v-icon>
                  </v-btn>
                </div>
              </template>
            </v-data-table>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Add Expense Button -->
    <v-fab
      color="primary"
      icon="mdi-plus"
      size="large"
      class="add-expense-fab"
      @click="addExpense"
    ></v-fab>

    <!-- Add New Expense Dialog -->
    <v-dialog v-model="addExpenseDialog" max-width="600" persistent>
      <v-card class="add-expense-dialog">
        <v-card-title class="dialog-header">
          <v-icon left color="primary">mdi-plus-circle</v-icon>
          إضافة صنف جديد
          <v-spacer></v-spacer>
          <v-btn
            icon
            variant="text"
            @click="closeAddExpenseDialog"
            class="close-btn"
          >
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-title>

        <v-card-text class="dialog-content">
          <v-form ref="addExpenseForm" v-model="addExpenseFormValid">
            <v-row>
              <!-- نوع الصرف -->
              <v-col cols="12">
                <v-select
                  v-model="newExpense.type"
                  :items="expenseTypes"
                  item-title="label"
                  item-value="value"
                  label="نوع الصرف"
                  variant="outlined"
                  :rules="[v => !!v || 'نوع الصرف مطلوب']"
                  required
                  class="form-field"
                ></v-select>
              </v-col>

              <!-- التكلفة -->
              <v-col cols="12">
                <v-text-field
                  v-model="newExpense.amount"
                  label="التكلفة (د.ع)"
                  variant="outlined"
                  type="number"
                  :rules="[
                    v => !!v || 'التكلفة مطلوبة',
                    v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر'
                  ]"
                  required
                  class="form-field"
                ></v-text-field>
              </v-col>

              <!-- الملاحظات -->
              <v-col cols="12">
                <v-textarea
                  v-model="newExpense.notes"
                  label="الملاحظات"
                  variant="outlined"
                  rows="3"
                  :rules="[v => !!v || 'الملاحظات مطلوبة']"
                  required
                  class="form-field"
                ></v-textarea>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <v-btn
            color="grey"
            variant="text"
            @click="closeAddExpenseDialog"
            class="cancel-btn"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="primary"
            variant="elevated"
            @click="saveNewExpense"
            :disabled="!addExpenseFormValid"
            class="save-btn"
          >
            <v-icon left>mdi-content-save</v-icon>
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { toast } from 'vue3-toastify'
import { formatAmount } from '@/utils/formatters'
import { PageHeader } from '@/components/shared'

const route = useRoute()
const router = useRouter()

// Project data from route params
const projectName = ref('')
const projectInfo = ref({})
const searchQuery = ref('')

// Add expense dialog
const addExpenseDialog = ref(false)
const addExpenseFormValid = ref(false)
const addExpenseForm = ref(null)

// New expense form data
const newExpense = ref({
  type: '',
  amount: '',
  notes: ''
})

// Expense types for dropdown
const expenseTypes = ref([
  { label: 'معدات', value: 'equipment' },
  { label: 'مواد', value: 'materials' },
  { label: 'أجور', value: 'labor' },
  { label: 'نقل', value: 'transportation' },
  { label: 'مرافق', value: 'utilities' },
  { label: 'أخرى', value: 'other' }
])

// Project expenses data
const projectExpenses = ref([
  {
    id: 1,
    type: 'equipment',
    description: 'شراء معدات حاسوبية',
    amount: 15000,
    date: '15/03/2022',
    status: 'approved',
    notes: 'معدات مطلوبة للمشروع'
  },
  {
    id: 2,
    type: 'materials',
    description: 'مواد بناء وأدوات',
    amount: 25000,
    date: '20/03/2022',
    status: 'pending',
    notes: 'مواد أساسية للمشروع'
  },
  {
    id: 3,
    type: 'labor',
    description: 'أجور العمال',
    amount: 30000,
    date: '25/03/2022',
    status: 'approved',
    notes: 'أجور شهر مارس'
  },
  {
    id: 4,
    type: 'transportation',
    description: 'تكاليف النقل',
    amount: 5000,
    date: '30/03/2022',
    status: 'rejected',
    notes: 'تكاليف نقل المواد'
  },
  {
    id: 5,
    type: 'utilities',
    description: 'فواتير الكهرباء والماء',
    amount: 8000,
    date: '05/04/2022',
    status: 'approved',
    notes: 'فواتير شهر مارس'
  }
])

// Table headers
const expenseHeaders = ref([
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'نوع المصروف', key: 'type', sortable: true, align: 'center' },
  { title: 'الوصف', key: 'description', sortable: true, align: 'right' },
  { title: 'المبلغ', key: 'amount', sortable: true, align: 'center' },
  { title: 'التاريخ', key: 'date', sortable: true, align: 'center' },
  { title: 'الحالة', key: 'status', sortable: true, align: 'center' },
  { title: 'الإجراءات', key: 'actions', sortable: false, align: 'center' }
])

// Computed properties
const totalExpenses = computed(() => {
  return projectExpenses.value
    .filter(expense => expense.status === 'approved')
    .reduce((sum, expense) => sum + expense.amount, 0)
    .toLocaleString()
})

const remainingBudget = computed(() => {
  const total = parseInt(projectInfo.value.cost?.replace(/,/g, '') || 0)
  const used = projectExpenses.value
    .filter(expense => expense.status === 'approved')
    .reduce((sum, expense) => sum + expense.amount, 0)
  return (total - used).toLocaleString()
})

const budgetUsagePercentage = computed(() => {
  const total = parseInt(projectInfo.value.cost?.replace(/,/g, '') || 0)
  if (total === 0) return 0
  const used = projectExpenses.value
    .filter(expense => expense.status === 'approved')
    .reduce((sum, expense) => sum + expense.amount, 0)
  return Math.round((used / total) * 100)
})

// Helper functions
const getExpenseTypeColor = (type) => {
  const colors = {
    equipment: 'blue',
    materials: 'orange',
    labor: 'green',
    transportation: 'purple',
    utilities: 'red',
    other: 'grey'
  }
  return colors[type] || 'grey'
}

const getExpenseTypeLabel = (type) => {
  const labels = {
    equipment: 'معدات',
    materials: 'مواد',
    labor: 'أجور',
    transportation: 'نقل',
    utilities: 'مرافق',
    other: 'أخرى'
  }
  return labels[type] || 'أخرى'
}

const getStatusColor = (status) => {
  const colors = {
    approved: 'success',
    pending: 'warning',
    rejected: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusLabel = (status) => {
  const labels = {
    approved: 'موافق عليه',
    pending: 'في الانتظار',
    rejected: 'مرفوض'
  }
  return labels[status] || 'غير محدد'
}

// formatAmount imported from @/utils/formatters

// Action functions
const goBack = () => {
  router.push('/expenses')
}

const viewExpenseDetails = (expense) => {
  toast.info(`${getExpenseTypeLabel(expense.type)}: ${expense.description} - ${formatAmount(expense.amount)}`)
}

const editExpense = (expense) => {
  // يمكن إضافة منطق التعديل هنا
}

const deleteExpense = (expense) => {
  if (confirm('هل أنت متأكد من حذف هذا المصروف؟')) {
    const index = projectExpenses.value.findIndex(e => e.id === expense.id)
    if (index > -1) {
      projectExpenses.value.splice(index, 1)
    }
  }
}

const addExpense = () => {
  addExpenseDialog.value = true
}

const closeAddExpenseDialog = () => {
  addExpenseDialog.value = false
  resetNewExpenseForm()
}

const resetNewExpenseForm = () => {
  newExpense.value = {
    type: '',
    amount: '',
    notes: ''
  }
  addExpenseFormValid.value = false
  if (addExpenseForm.value) {
    addExpenseForm.value.reset()
  }
}

const saveNewExpense = () => {
  if (addExpenseFormValid.value) {
    // إنشاء مصروف جديد
    const newExpenseItem = {
      id: Date.now(), // استخدام timestamp كـ ID فريد
      type: newExpense.value.type,
      description: getExpenseTypeLabel(newExpense.value.type),
      amount: parseInt(newExpense.value.amount),
      date: new Date().toLocaleDateString('ar-SA'),
      status: 'pending', // حالة افتراضية: في الانتظار
      notes: newExpense.value.notes
    }
    
    // إضافة المصروف إلى القائمة
    projectExpenses.value.unshift(newExpenseItem)

    // إغلاق النموذج وإعادة تعيينه
    closeAddExpenseDialog()
    
    toast.success('تم إضافة المصروف بنجاح!')
  }
}

// Initialize component
onMounted(() => {
  // Get project data from route query parameters
  if (route.query.projectName) {
    projectName.value = route.query.projectName
  }

  // Set project info from query parameters
  projectInfo.value = {
    startDate: route.query.startDate || '10/03/2022',
    endDate: route.query.endDate || '07/03/2024',
    cost: route.query.cost || '500000',
    workLocation: route.query.workLocation || 'لعقوبة',
    notes: route.query.notes || 'لايوجد'
  }
})
</script>


<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/project-expenses.css';
</style>
