<template>
  <div class="fill-height data-page">
    <div>
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <div class="header-top-content">
          <h1 class="page-title">الإيرادات</h1>
          <span class="page-icon">💰</span>
        </div>
        <p class="page-subtitle">إدارة وتتبع جميع مصادر الإيرادات</p>
      </div>

      <!-- Summary Cards -->
      <div class="cards-container">
        <v-row class="mb-6" no-gutters>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card revenue-card" elevation="0">
            <div class="card-header">
              <div class="card-icon success">
                <i class="mdi mdi-currency-usd"></i>
              </div>
              <div class="card-title">إجمالي الإيرادات</div>
            </div>
            <div class="card-content">
              <div class="card-value success">{{ formatCurrency(totalIncome) }}</div>
              <div class="card-subtitle">جميع الإيرادات المؤكدة</div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="stat-card monthly-card" elevation="0">
            <div class="card-header">
              <div class="card-icon info">
                <i class="mdi mdi-calendar-month"></i>
              </div>
              <div class="card-title">إيرادات هذا الشهر</div>
            </div>
            <div class="card-content">
              <div class="card-value info">{{ formatCurrency(monthlyIncome) }}</div>
              <div class="card-subtitle">شهر {{ new Date().toLocaleDateString('ar-SA', { month: 'long' }) }}</div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="stat-card growth-card" elevation="0">
            <div class="card-header">
              <div class="card-icon warning">
                <i class="mdi mdi-trending-up"></i>
              </div>
              <div class="card-title">نمو الإيرادات</div>
            </div>
            <div class="card-content">
              <div class="card-value warning">{{ incomeGrowth }}%</div>
              <div class="card-subtitle">مقارنة بالشهر الماضي</div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="stat-card sources-card" elevation="0">
            <div class="card-header">
              <div class="card-icon primary">
                <i class="mdi mdi-source-branch"></i>
              </div>
              <div class="card-title">مصادر الإيرادات</div>
            </div>
            <div class="card-content">
              <div class="card-value primary">{{ incomeSources.length }}</div>
              <div class="card-subtitle">إجمالي المصادر</div>
            </div>
          </v-card>
        </v-col>
        </v-row>
      </div>

      <!-- Action Buttons -->
      <div class="action-buttons mb-6">
        <v-btn
          class="add-button"
          size="large"
          @click="showAddDialog = true"
        >
          <i class="mdi mdi-plus me-2"></i>
          إضافة إيراد جديد
        </v-btn>
        <v-btn
          class="export-button"
          size="large"
          @click="exportToCSV"
        >
          <i class="mdi mdi-download me-2"></i>
          تصدير Excel
        </v-btn>
      </div>

      <!-- Income Sources Table -->
      <div class="table-container">
        <v-card class="data-table-card mb-6">
        <v-card-title class="d-flex align-center">
          <v-icon class="me-2">mdi-format-list-bulleted</v-icon>
          مصادر الإيرادات
        </v-card-title>
        <v-data-table
          :headers="headers"
          :items="incomeSources"
          :loading="loading"
          class="elevation-1 income-table"
        >
          <template v-slot:item.amount="{ item }">
            <span class="font-weight-bold text-success">{{ formatCurrency(item.amount) }}</span>
          </template>
          <template v-slot:item.date="{ item }">
            {{ formatDate(item.date) }}
          </template>
          <template v-slot:item.actions="{ item }">
            <v-btn
              size="small"
              color="primary"
              @click="editIncome(item)"
            >
              <i class="mdi mdi-pencil"></i>
            </v-btn>
            <v-btn
              size="small"
              color="error"
              @click="deleteIncome(item)"
            >
              <i class="mdi mdi-delete"></i>
            </v-btn>
          </template>
        </v-data-table>
        </v-card>
      </div>

      <!-- Add/Edit Income Dialog -->
      <v-dialog v-model="showAddDialog" max-width="600">
        <v-card>
          <v-card-title>
            <span class="text-h5">{{ editingIncome ? 'تعديل الإيراد' : 'إضافة إيراد جديد' }}</span>
          </v-card-title>
          <v-card-text>
            <v-form ref="form" v-model="valid">
              <v-text-field
                v-model="incomeForm.description"
                label="وصف الإيراد"
                :rules="[v => !!v || 'الوصف مطلوب']"
                required
              />
              <v-text-field
                v-model.number="incomeForm.amount"
                label="المبلغ"
                type="number"
                :rules="[v => v > 0 || 'المبلغ يجب أن يكون أكبر من صفر']"
                required
              />
              <v-select
                v-model="incomeForm.category"
                :items="incomeCategories"
                label="الفئة"
                :rules="[v => !!v || 'الفئة مطلوبة']"
                required
              />
              <v-textarea
                v-model="incomeForm.notes"
                label="ملاحظات"
                rows="3"
              />
            </v-form>
          </v-card-text>
          <v-card-actions class="dialog-actions">
            <v-spacer />
            <v-btn color="grey" @click="closeDialog">إلغاء</v-btn>
            <v-btn color="success" @click="saveIncome" :disabled="!valid">حفظ</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'

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
// نموذج البيانات
// ========================================
const incomeForm = ref({
  description: '',
  amount: 0,
  category: '',
  notes: '',
  date: new Date().toISOString().split('T')[0]
})

// ========================================
// البيانات الأساسية
// ========================================
const incomeSources = ref([
  {
    id: 1,
    description: 'رسوم التسجيل',
    amount: 50000,
    category: 'رسوم طلابية',
    date: '2024-01-15',
    notes: 'رسوم الفصل الدراسي الأول',
    status: 'confirmed',
    source: 'طلاب'
  },
  {
    id: 2,
    description: 'منحة حكومية',
    amount: 100000,
    category: 'منح',
    date: '2024-01-10',
    notes: 'منحة وزارة التعليم العالي',
    status: 'confirmed',
    source: 'حكومي'
  },
  {
    id: 3,
    description: 'استشارات أكاديمية',
    amount: 25000,
    category: 'خدمات',
    date: '2024-01-05',
    notes: 'استشارات للقطاع الخاص',
    status: 'confirmed',
    source: 'خاص'
  },
  {
    id: 4,
    description: 'إيرادات الكافتيريا',
    amount: 15000,
    category: 'خدمات',
    date: '2024-01-20',
    notes: 'إيرادات شهر يناير',
    status: 'pending',
    source: 'داخلي'
  },
  {
    id: 5,
    description: 'تبرع خيري',
    amount: 75000,
    category: 'تبرعات',
    date: '2024-01-25',
    notes: 'تبرع من مؤسسة خيرية',
    status: 'confirmed',
    source: 'خيري'
  }
])

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
  { title: 'مؤكد', value: 'confirmed', color: 'success' },
  { title: 'معلق', value: 'pending', color: 'warning' },
  { title: 'ملغي', value: 'cancelled', color: 'error' }
]

const sourceOptions = [
  { title: 'طلاب', value: 'طلاب' },
  { title: 'حكومي', value: 'حكومي' },
  { title: 'خاص', value: 'خاص' },
  { title: 'داخلي', value: 'داخلي' },
  { title: 'خيري', value: 'خيري' }
]

// ========================================
// عناوين الجدول
// ========================================
const headers = [
  { title: 'الوصف', key: 'description', align: 'start', sortable: true },
  { title: 'المبلغ', key: 'amount', align: 'center', sortable: true },
  { title: 'الفئة', key: 'category', align: 'center', sortable: true },
  { title: 'المصدر', key: 'source', align: 'center', sortable: true },
  { title: 'الحالة', key: 'status', align: 'center', sortable: true },
  { title: 'التاريخ', key: 'date', align: 'center', sortable: true },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
]

// ========================================
// الخصائص المحسوبة
// ========================================
const totalIncome = computed(() => {
  return incomeSources.value
    .filter(item => item.status === 'confirmed')
    .reduce((sum, item) => sum + item.amount, 0)
})

const monthlyIncome = computed(() => {
  const currentMonth = new Date().getMonth()
  const currentYear = new Date().getFullYear()
  
  return incomeSources.value
    .filter(item => {
      const itemDate = new Date(item.date)
      return itemDate.getMonth() === currentMonth && 
             itemDate.getFullYear() === currentYear &&
             item.status === 'confirmed'
    })
    .reduce((sum, item) => sum + item.amount, 0)
})

const pendingIncome = computed(() => {
  return incomeSources.value
    .filter(item => item.status === 'pending')
    .reduce((sum, item) => sum + item.amount, 0)
})

const incomeGrowth = computed(() => {
  const currentMonth = new Date().getMonth()
  const currentYear = new Date().getFullYear()
  const lastMonth = currentMonth === 0 ? 11 : currentMonth - 1
  const lastMonthYear = currentMonth === 0 ? currentYear - 1 : currentYear
  
  const currentMonthIncome = incomeSources.value
    .filter(item => {
      const itemDate = new Date(item.date)
      return itemDate.getMonth() === currentMonth && 
             itemDate.getFullYear() === currentYear &&
             item.status === 'confirmed'
    })
    .reduce((sum, item) => sum + item.amount, 0)
  
  const lastMonthIncome = incomeSources.value
    .filter(item => {
      const itemDate = new Date(item.date)
      return itemDate.getMonth() === lastMonth && 
             itemDate.getFullYear() === lastMonthYear &&
             item.status === 'confirmed'
    })
    .reduce((sum, item) => sum + item.amount, 0)
  
  if (lastMonthIncome === 0) return 0
  return ((currentMonthIncome - lastMonthIncome) / lastMonthIncome * 100).toFixed(1)
})

const filteredIncomeSources = computed(() => {
  let filtered = incomeSources.value
  
  // فلترة حسب البحث
  if (searchQuery.value) {
    filtered = filtered.filter(item =>
      item.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      item.category.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      item.notes.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  
  // فلترة حسب الفئة
  if (selectedCategory.value) {
    filtered = filtered.filter(item => item.category === selectedCategory.value)
  }
  
  // فلترة حسب التاريخ
  if (dateRange.value && dateRange.value.length === 2) {
    const startDate = new Date(dateRange.value[0])
    const endDate = new Date(dateRange.value[1])
    filtered = filtered.filter(item => {
      const itemDate = new Date(item.date)
      return itemDate >= startDate && itemDate <= endDate
    })
  }
  
  return filtered
})

const incomeByCategory = computed(() => {
  const categories = {}
  incomeSources.value
    .filter(item => item.status === 'confirmed')
    .forEach(item => {
      if (!categories[item.category]) {
        categories[item.category] = 0
      }
      categories[item.category] += item.amount
    })
  return categories
})

// ========================================
// الدوال المساعدة
// ========================================
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD'
  }).format(amount)
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('ar-SA')
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
// دوال إدارة البيانات
// ========================================
const openAddDialog = () => {
  editingIncome.value = null
  incomeForm.value = {
    description: '',
    amount: 0,
    category: '',
    notes: '',
    date: new Date().toISOString().split('T')[0]
  }
  showAddDialog.value = true
}

const editIncome = (item) => {
  editingIncome.value = item
  incomeForm.value = { 
    ...item,
    date: formatDateForInput(item.date)
  }
  showAddDialog.value = true
}

const deleteIncome = (item) => {
  if (confirm(`هل أنت متأكد من حذف الإيراد "${item.description}"؟`)) {
    const index = incomeSources.value.findIndex(i => i.id === item.id)
    if (index > -1) {
      incomeSources.value.splice(index, 1)
      // حفظ في localStorage
      saveToLocalStorage()
    }
  }
}

const saveIncome = () => {
  if (valid.value) {
    if (editingIncome.value) {
      // تحديث الإيراد الموجود
      const index = incomeSources.value.findIndex(i => i.id === editingIncome.value.id)
      if (index > -1) {
        incomeSources.value[index] = {
          ...incomeForm.value,
          id: editingIncome.value.id,
          status: editingIncome.value.status,
          source: editingIncome.value.source
        }
      }
    } else {
      // إضافة إيراد جديد
      const newIncome = {
        ...incomeForm.value,
        id: Date.now(),
        status: 'confirmed',
        source: 'داخلي'
      }
      incomeSources.value.unshift(newIncome)
    }
    
    // حفظ في localStorage
    saveToLocalStorage()
    closeDialog()
  }
}

const closeDialog = () => {
  showAddDialog.value = false
  editingIncome.value = null
  incomeForm.value = {
    description: '',
    amount: 0,
    category: '',
    notes: '',
    date: new Date().toISOString().split('T')[0]
  }
}

const updateIncomeStatus = (item, newStatus) => {
  const index = incomeSources.value.findIndex(i => i.id === item.id)
  if (index > -1) {
    incomeSources.value[index].status = newStatus
    saveToLocalStorage()
  }
}

// ========================================
// دوال التخزين المحلي
// ========================================
const saveToLocalStorage = () => {
  try {
    localStorage.setItem('incomeSources', JSON.stringify(incomeSources.value))
  } catch (error) {
    console.error('خطأ في حفظ البيانات:', error)
  }
}

const loadFromLocalStorage = () => {
  try {
    const saved = localStorage.getItem('incomeSources')
    if (saved) {
      incomeSources.value = JSON.parse(saved)
    }
  } catch (error) {
    console.error('خطأ في تحميل البيانات:', error)
  }
}

// ========================================
// دوال التصدير والطباعة
// ========================================
const exportToCSV = () => {
  const headers = ['الوصف', 'المبلغ', 'الفئة', 'المصدر', 'الحالة', 'التاريخ', 'ملاحظات']
  const csvContent = [
    headers.join(','),
    ...filteredIncomeSources.value.map(item => [
      `"${item.description}"`,
      item.amount,
      `"${item.category}"`,
      `"${item.source}"`,
      `"${getStatusText(item.status)}"`,
      `"${formatDate(item.date)}"`,
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
        const itemDate = new Date(item.date)
        return itemDate.getMonth() === date.getMonth() &&
               itemDate.getFullYear() === date.getFullYear() &&
               item.status === 'confirmed'
      })
      .reduce((sum, item) => sum + item.amount, 0)
    
    trends.push({
      month: date.toLocaleDateString('ar-SA', { month: 'short' }),
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
// مراقبة التغييرات
// ========================================
watch(incomeSources, () => {
  // إعادة حساب الإحصائيات عند تغيير البيانات
}, { deep: true })

// ========================================
// دورة الحياة
// ========================================
onMounted(() => {
  loading.value = true
  
  // تحميل البيانات من localStorage
  loadFromLocalStorage()
  
  // محاكاة تحميل البيانات
  setTimeout(() => {
    loading.value = false
  }, 1000)
})
</script>

<style scoped>
/* ========================================
   صفحة الإيرادات - CSS منظم ومتناسق
   ======================================== */

/* صفحة البيانات العامة */
.data-page {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  color: var(--text-dark);
  min-height: 100vh;
  padding: 0 !important;
  max-width: 100% !important;
  margin: 0 !important;
  width: 100% !important;
}

/* حاويات المحتوى */
.cards-container {
  padding: 0 2rem;
  margin-top: 1.5rem;
}

.action-buttons {
  padding: 0 2rem;
}

.table-container {
  padding: 0 2rem;
}

/* رأس الصفحة المحسن */
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
  color: white;
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

.page-icon {
  font-size: 2rem;
  display: inline-block;
  animation: iconFloat 3s ease-in-out infinite;
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
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

.page-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grain" width="100" height="100" patternUnits="userSpaceOnUse"><circle cx="25" cy="25" r="1" fill="white" opacity="0.1"/><circle cx="75" cy="75" r="1" fill="white" opacity="0.1"/><circle cx="50" cy="10" r="0.5" fill="white" opacity="0.1"/><circle cx="10" cy="60" r="0.5" fill="white" opacity="0.1"/><circle cx="90" cy="40" r="0.5" fill="white" opacity="0.1"/></pattern></defs><rect width="100" height="100" fill="url(%23grain)"/></svg>');
  opacity: 0.3;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 2rem;
  position: relative;
  z-index: 2;
  width: 100%;
  justify-content: flex-start;
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

/* بطاقات الإحصائيات المحسنة */
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

.data-table-card .v-data-table {
  background: transparent !important;
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

.add-button .v-btn__content {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-align: center !important;
}

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

.export-button .v-btn__content {
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
</style>
