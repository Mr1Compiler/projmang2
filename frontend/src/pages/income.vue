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


<style>
/* Import page styles */
@import './styles/income.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
