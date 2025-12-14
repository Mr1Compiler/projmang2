<template>
  <v-container fluid class="project-expenses-page">
    <!-- Header Section -->
    <v-row class="mb-6">
      <v-col cols="12">
        <v-card class="header-card" elevation="3">
          <v-card-text class="pa-6">
            <div class="d-flex align-center justify-space-between">
              <div>
                <h1 class="page-title">{{ projectName }}</h1>
                <p class="page-subtitle">تفاصيل مصاريف المشروع</p>
              </div>
              <div class="header-actions">
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
            </div>
          </v-card-text>
        </v-card>
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

const formatAmount = (amount) => {
  return `${amount.toLocaleString()} د.ع`
}

// Action functions
const goBack = () => {
  router.push('/expenses')
}

const viewExpenseDetails = (expense) => {
  console.log('عرض تفاصيل المصروف:', expense)
  alert(`تفاصيل المصروف:\n\nالنوع: ${getExpenseTypeLabel(expense.type)}\nالوصف: ${expense.description}\nالمبلغ: ${formatAmount(expense.amount)}\nالتاريخ: ${expense.date}\nالحالة: ${getStatusLabel(expense.status)}\nالملاحظات: ${expense.notes}`)
}

const editExpense = (expense) => {
  console.log('تعديل المصروف:', expense)
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
  console.log('فتح نموذج إضافة مصروف جديد')
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
    
    console.log('تم إضافة مصروف جديد:', newExpenseItem)
    
    // إغلاق النموذج وإعادة تعيينه
    closeAddExpenseDialog()
    
    // رسالة نجاح (يمكن استبدالها بـ snackbar)
    alert('تم إضافة المصروف بنجاح!')
  }
}

// Initialize component
onMounted(() => {
  console.log('تم تحميل صفحة مصاريف المشروع')
  console.log('معاملات URL:', route.query)
  
  // Get project data from route query parameters
  if (route.query.projectName) {
    projectName.value = route.query.projectName
    console.log('اسم المشروع:', projectName.value)
  }
  
  // Set project info from query parameters
  projectInfo.value = {
    startDate: route.query.startDate || '10/03/2022',
    endDate: route.query.endDate || '07/03/2024',
    cost: route.query.cost || '500000',
    workLocation: route.query.workLocation || 'لعقوبة',
    notes: route.query.notes || 'لايوجد'
  }
  
  console.log('معلومات المشروع:', projectInfo.value)
})
</script>

<style scoped>
.project-expenses-page {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  min-height: 100vh;
  padding: 20px;
}

/* Header Styles */
.header-card {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: white !important;
  border-radius: 16px !important;
}

.page-title {
  font-size: 2.5rem !important;
  font-weight: 800 !important;
  margin-bottom: 8px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
}

.page-subtitle {
  font-size: 1.2rem !important;
  opacity: 0.9 !important;
  margin: 0 !important;
}

.header-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.add-btn-header {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  padding: 12px 20px !important;
  border-radius: 12px !important;
  box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3) !important;
  text-transform: none !important;
  letter-spacing: 0.5px !important;
}

.add-btn-header:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 16px rgba(5, 150, 105, 0.4) !important;
}

.add-btn-header .v-icon {
  margin-right: 8px !important;
  font-size: 1.1rem !important;
}

.back-btn {
  background: rgba(255, 255, 255, 0.2) !important;
  color: white !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  border-radius: 12px !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.3) !important;
  transform: translateY(-2px) !important;
}

/* Info Card Styles */
.info-card {
  border-radius: 16px !important;
  border: 2px solid #e2e8f0 !important;
}

.info-title {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%) !important;
  color: #1e40af !important;
  font-weight: 700 !important;
  font-size: 1.3rem !important;
  border-bottom: 2px solid #e2e8f0 !important;
}

.info-item {
  margin-bottom: 12px;
}

.info-label {
  font-weight: 600 !important;
  color: #64748b !important;
  display: block !important;
  margin-bottom: 4px !important;
  font-size: 0.9rem !important;
}

.info-value {
  font-weight: 700 !important;
  color: #1e293b !important;
  font-size: 1.1rem !important;
}

.info-value.cost {
  color: #059669 !important;
  font-size: 1.3rem !important;
}

/* Statistics Cards */
.stat-card {
  border-radius: 16px !important;
  transition: all 0.3s ease !important;
  border: 2px solid transparent !important;
}

.stat-card:hover {
  transform: translateY(-4px) !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15) !important;
}

.total-expenses {
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%) !important;
  border-color: #3b82f6 !important;
}

.remaining-budget {
  background: linear-gradient(135deg, #dcfce7 0%, #bbf7d0 100%) !important;
  border-color: #22c55e !important;
}

.expense-count {
  background: linear-gradient(135deg, #e0e7ff 0%, #c7d2fe 100%) !important;
  border-color: #6366f1 !important;
}

.budget-usage {
  background: linear-gradient(135deg, #fef3c7 0%, #fde68a 100%) !important;
  border-color: #f59e0b !important;
}

.stat-number {
  font-size: 2.5rem !important;
  font-weight: 800 !important;
  margin: 12px 0 8px 0 !important;
  color: #1e293b !important;
}

.stat-label {
  font-size: 1rem !important;
  font-weight: 600 !important;
  color: #64748b !important;
  margin: 0 !important;
}

/* Table Styles */
.expenses-table-card {
  border-radius: 16px !important;
  border: 2px solid #e9d5ff !important;
  box-shadow: 0 8px 25px rgba(124, 58, 237, 0.1) !important;
  background: linear-gradient(135deg, #ffffff 0%, #faf5ff 100%) !important;
}

.table-title {
  background: linear-gradient(135deg, #f3e8ff 0%, #e9d5ff 100%) !important;
  color: #7c3aed !important;
  font-weight: 800 !important;
  font-size: 1.5rem !important;
  border-bottom: 3px solid #c4b5fd !important;
  text-shadow: 0 2px 4px rgba(124, 58, 237, 0.2) !important;
  letter-spacing: 0.5px !important;
}

.search-field {
  max-width: 300px !important;
}

.expenses-table {
  border-radius: 12px !important;
  overflow: hidden !important;
}

.expenses-table .v-data-table__th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-align: center !important;
  padding: 1.2rem 0.8rem !important;
  border-bottom: 4px solid #065f46 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  position: relative !important;
}

.expenses-table .v-data-table__th::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, transparent 0%, #ffffff 50%, transparent 100%);
  opacity: 0.3;
}

.expenses-table .v-data-table__td {
  background: #ffffff !important;
  color: #1f2937 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 1rem 0.8rem !important;
  border-bottom: 1px solid #e5e7eb !important;
  text-align: center !important;
  transition: all 0.3s ease !important;
  vertical-align: middle !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table .v-data-table__wrapper tbody tr:nth-child(even) .v-data-table__td {
  background: #f8fafc !important;
}

.expenses-table .v-data-table__wrapper tbody tr:hover .v-data-table__td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 700 !important;
  transform: scale(1.01) !important;
  box-shadow: 0 4px 12px rgba(4, 120, 87, 0.15) !important;
}

/* Custom Text Styles */
.serial-number {
  font-weight: 700 !important;
  color: #374151 !important;
  background: #f3f4f6 !important;
  padding: 6px 10px !important;
  border-radius: 8px !important;
  font-size: 1rem !important;
  border: 1px solid #d1d5db !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.expense-type-chip {
  font-weight: 700 !important;
  font-size: 0.9rem !important;
  border-radius: 12px !important;
  padding: 4px 12px !important;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1) !important;
}

.amount-text {
  font-weight: 700 !important;
  color: #059669 !important;
  font-size: 1rem !important;
  background: #ecfdf5 !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
  border: 1px solid #a7f3d0 !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
}

.date-text {
  font-weight: 600 !important;
  color: #374151 !important;
  font-size: 0.95rem !important;
  background: #f3f4f6 !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
  border: 1px solid #d1d5db !important;
}

.status-chip {
  font-weight: 700 !important;
  font-size: 0.9rem !important;
  border-radius: 12px !important;
  padding: 6px 12px !important;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1) !important;
}

/* Action Buttons */
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

.edit-btn {
  color: #059669 !important;
  background: rgba(5, 150, 105, 0.1) !important;
}

.edit-btn:hover {
  background: rgba(5, 150, 105, 0.2) !important;
  transform: scale(1.05) !important;
}

.delete-btn {
  color: #dc2626 !important;
  background: rgba(220, 38, 38, 0.1) !important;
}

.delete-btn:hover {
  background: rgba(220, 38, 38, 0.2) !important;
  transform: scale(1.05) !important;
}

/* Add Expense Section */
.add-expense-section {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding: 16px 0;
  border-bottom: 2px solid #e2e8f0;
  margin-bottom: 20px;
}

.add-expense-btn {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 12px 24px !important;
  border-radius: 12px !important;
  box-shadow: 0 4px 12px rgba(30, 64, 175, 0.3) !important;
  text-transform: none !important;
  letter-spacing: 0.5px !important;
}

.add-expense-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 16px rgba(30, 64, 175, 0.4) !important;
}

.add-expense-btn .v-icon {
  margin-right: 8px !important;
  font-size: 1.2rem !important;
}

/* FAB Button */
.add-expense-fab {
  position: fixed !important;
  bottom: 24px !important;
  right: 24px !important;
  z-index: 1000 !important;
}

/* Add Expense Dialog Styles */
.add-expense-dialog {
  border-radius: 20px !important;
  overflow: hidden !important;
  box-shadow: 0 20px 40px rgba(210, 40, 176, 0.2) !important;
  border: 3px solid #0c24ad !important;
  background: linear-gradient(135deg, #190244 0%, #7300e7 100%) !important;
}

.dialog-header {
  background: linear-gradient(135deg, #2a38d1 0%, #423cac 100%) !important;
  color: rgb(255, 255, 255) !important;
  font-weight: 800 !important;
    font-size: 1.4rem !important;
  padding: 24px 28px !important;
  border-bottom: 3px solid #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.close-btn {
  color: rgb(14, 7, 7) !important;
  background: rgba(255, 255, 255, 0.2) !important;
  border-radius: 8px !important;
}

.close-btn:hover {
  background: rgba(7, 1, 1, 0.3) !important;
  transform: scale(1.05) !important;
}

.dialog-content {
  padding: 28px !important;
  background: linear-gradient(135deg, #7505e5 0%, #5d35cb 100%) !important;
  border-radius: 0 0 16px 16px !important;
}

.form-field {
  margin-bottom: 16px !important;
}

.form-field .v-field {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #eeecec 0%, #b2095b 100%) !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.1) !important;
  border: 2px solid #7409e7 !important;
}

.form-field .v-field__outline {
  border: 2px solid #f8f6f6 !important;
  border-radius: 12px !important;
}

.form-field .v-field--focused .v-field__outline {
  border-color: #7c3aed !important;
  box-shadow: 0 0 0 3px rgba(124, 58, 237, 0.2) !important;
}

.form-field .v-label {
  color: #7c3aed !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.form-field .v-field__input {
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 1.05rem !important;
}

/* Dropdown Menu Styling */
.form-field .v-list {
  background: linear-gradient(135deg, #000000 0%, #1a0033 100%) !important;
  border-radius: 12px !important;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.8) !important;
  border: 4px solid #ff6b35 !important;
  backdrop-filter: blur(15px) !important;
}

.form-field .v-list-item {
  color: #ffffff !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  padding: 20px 28px !important;
  transition: all 0.3s ease !important;
  border-bottom: 3px solid rgba(255, 255, 255, 0.5) !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.9) !important;
  letter-spacing: 1px !important;
  background: linear-gradient(135deg, #2a0033 0%, #1a0022 100%) !important;
}

.form-field .v-list-item:hover {
  background: linear-gradient(135deg, #ff6b35 0%, #ff8c42 100%) !important;
  color: #000000 !important;
  transform: translateX(10px) !important;
  box-shadow: 0 8px 25px rgba(255, 107, 53, 0.7) !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  border-left: 6px solid #000000 !important;
  font-weight: 900 !important;
}

.form-field .v-list-item--active {
  background: linear-gradient(135deg, #ff6b35 0%, #ff8c42 100%) !important;
  color: #000000 !important;
  font-weight: 900 !important;
  box-shadow: 0 10px 30px rgba(255, 107, 53, 0.8) !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  border-left: 6px solid #000000 !important;
  transform: scale(1.05) !important;
}

.dialog-actions {
  padding: 20px 28px !important;
  background: linear-gradient(135deg, #ffffff 0%, #faf5ff 100%) !important;
  border-top: 2px solid #e9d5ff !important;
  border-radius: 0 0 16px 16px !important;
}

.cancel-btn {
  color: #7c3aed !important;
  font-weight: 700 !important;
  padding: 14px 28px !important;
  border-radius: 12px !important;
  margin-left: 12px !important;
  background: linear-gradient(135deg, #1a05d5 0%, #fdfbff 100%) !important;
  border: 2px solid #fdfdff !important;
  box-shadow: 0 2px 8px rgba(124, 58, 237, 0.2) !important;
}

.cancel-btn:hover {
  background: linear-gradient(135deg, #f0efef 0%, #ddd6fe 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.3) !important;
}

.save-btn {
  background: linear-gradient(135deg, #7c3aed 0%, #a855f7 100%) !important;
  color: white !important;
  font-weight: 800 !important;
  padding: 14px 28px !important;
  border-radius: 12px !important;
  box-shadow: 0 6px 16px rgba(124, 58, 237, 0.4) !important;
  border: 2px solid #6d28d9 !important;
}

.save-btn:hover {
  transform: translateY(-3px) !important;
  box-shadow: 0 8px 20px rgba(124, 58, 237, 0.5) !important;
}

.save-btn:disabled {
  background: linear-gradient(135deg, #e2e8f0 0%, #cbd5e1 100%) !important;
  color: #94a3b8 !important;
  box-shadow: none !important;
  transform: none !important;
  border-color: #cbd5e1 !important;
}

/* Responsive Design */
@media (max-width: 768px) {
  .project-expenses-page {
    padding: 10px;
  }
  
  .page-title {
    font-size: 2rem !important;
  }
  
  .stat-number {
    font-size: 2rem !important;
  }
  
  .action-buttons {
    flex-direction: column;
    gap: 2px;
  }
  
  .action-btn {
    min-width: 28px !important;
    height: 28px !important;
  }
}

/* ========================================
   تحسين إضافي لعناوين الجدول - Additional Table Headers
   ======================================== */
.expenses-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.expenses-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.expenses-table table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}

/* CSS إضافي لضمان التطبيق */
.expenses-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.expenses-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}
</style>

<style>
/* CSS غير محدود النطاق لضمان تطبيق التغييرات */
.expenses-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  background: #ffffff !important;
  border-bottom: 1px solid #e5e7eb !important;
  padding: 1rem 0.8rem !important;
  vertical-align: middle !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 800 !important;
}

.expenses-table table tbody tr:nth-child(even) td {
  background: #f8fafc !important;
}

.expenses-table table tbody tr:nth-child(even):hover td {
  background: #f1f5f9 !important;
}

/* تحسين وضوح النصوص داخل الخلايا */
.expenses-table .serial-number {
  color: #374151 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  background: #f3f4f6 !important;
  padding: 8px 12px !important;
  border-radius: 8px !important;
  border: 2px solid #d1d5db !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table .amount-text {
  color: #059669 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  background: #ecfdf5 !important;
  padding: 6px 10px !important;
  border-radius: 6px !important;
  border: 2px solid #a7f3d0 !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table .date-text {
  color: #374151 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  background: #f3f4f6 !important;
  padding: 6px 10px !important;
  border-radius: 6px !important;
  border: 2px solid #d1d5db !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* CSS إضافي لضمان التطبيق */
.v-data-table.expenses-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  background: #ffffff !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.v-data-table.expenses-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 800 !important;
}

/* ========================================
   تحسين وضوح بيانات الجدول - Table Data Clarity
   ======================================== */
.expenses-table table tbody tr td {
  color: #1f2937 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  background: #ffffff !important;
  border-bottom: 1px solid #e5e7eb !important;
  padding: 1rem 0.8rem !important;
  vertical-align: middle !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table table tbody tr:hover td {
  background: #f9fafb !important;
  color: #111827 !important;
  font-weight: 800 !important;
}

.expenses-table table tbody tr:nth-child(even) td {
  background: #f8fafc !important;
}

.expenses-table table tbody tr:nth-child(even):hover td {
  background: #f1f5f9 !important;
}

/* تحسين وضوح النصوص داخل الخلايا */
.expenses-table .serial-number {
  color: #374151 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  background: #f3f4f6 !important;
  padding: 8px 12px !important;
  border-radius: 8px !important;
  border: 2px solid #d1d5db !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table .amount-text {
  color: #059669 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  background: #ecfdf5 !important;
  padding: 6px 10px !important;
  border-radius: 6px !important;
  border: 2px solid #a7f3d0 !important;
  font-family: 'Arial', sans-serif !important;
  direction: ltr !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expenses-table .date-text {
  color: #374151 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  background: #f3f4f6 !important;
  padding: 6px 10px !important;
  border-radius: 6px !important;
  border: 2px solid #d1d5db !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}
</style>
