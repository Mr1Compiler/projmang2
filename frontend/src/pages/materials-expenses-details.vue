<template>
  <div class="materials-expenses-page">
    <!-- Modern Dialog for Add Material -->
    <ModernDialog
      :open="showAddDialog"
      title="إضافة مادة"
      save-text="حفظ"
      icon="mdi-hammer-screwdriver"
      @close="closeAddDialog"
      @save="saveMaterial"
    >
      <MaterialForm
        v-model="newMaterial"
        ref="materialFormRef"
      />
    </ModernDialog>
    <!-- Header Section -->
    <div class="page-header glass-effect gradient-animation">
      <div class="header-content">
        <v-btn 
          icon="mdi-arrow-left" 
          @click="goBack" 
          class="back-btn"
          size="large"
          color="white"
        >
          <v-icon>mdi-arrow-left</v-icon>
        </v-btn>
        <div class="header-text">
          <h1 class="page-title text-glow fade-in">المواد والمصاريف اليومية</h1>
          <p class="page-subtitle fade-in">إدارة المواد والمصاريف اليومية للمشروع</p>
        </div>
      </div>
    </div>

    <!-- Materials Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="error">mdi-hammer-screwdriver</v-icon>
        المواد
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="materialsSearch"
              label="البحث في المواد..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchMaterials"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchMaterials"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="success" 
              variant="elevated" 
              class="add-btn"
              @click="openAddMaterialDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة مادة
            </v-btn>
          </div>
        </div>
      </v-card-text>

      <!-- Materials Table -->
      <v-data-table
        :headers="materialsHeaders"
        :items="materialsData"
        :search="materialsSearch"
        :items-per-page="10"
        class="data-table materials-table"
        no-data-text="لا توجد بيانات متاحة"
        loading-text="جاري التحميل..."
        density="comfortable"
      >
        <template v-slot:item.actions="{ item }">
          <v-btn
            icon="mdi-delete"
            size="small"
            color="error"
            variant="text"
            @click="deleteMaterial(item)"
            title="حذف المادة"
          />
        </template>
      </v-data-table>
    </v-card>

    <!-- Daily Expenses Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="error">mdi-cash-multiple</v-icon>
        المصاريف اليومية
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="expensesSearch"
              label="البحث في المصاريف..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchExpenses"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchExpenses"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="warning" 
              variant="elevated" 
              class="add-btn"
              @click="openAddExpenseDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة مصروف
            </v-btn>
          </div>
        </div>
      </v-card-text>

      <!-- Expenses Table -->
      <v-data-table
        :headers="expensesHeaders"
        :items="expensesData"
        :search="expensesSearch"
        :items-per-page="10"
        class="data-table expenses-table"
        no-data-text="لا توجد بيانات متاحة"
        loading-text="جاري التحميل..."
        density="comfortable"
      >
        <template v-slot:item.actions="{ item }">
          <v-btn
            icon="mdi-delete"
            size="small"
            color="error"
            variant="text"
            @click="deleteExpense(item)"
            title="حذف المصروف"
          />
        </template>
      </v-data-table>
    </v-card>

    <!-- Simple Dialog for Add Material -->
    <SimpleDialog
      :open="showAddDialog"
      title="إضافة مادة"
      save-text="حفظ"
      icon="mdi-hammer-screwdriver"
      @close="closeAddDialog"
      @save="saveMaterial"
    >
      <MaterialForm
        v-model="newMaterial"
        ref="materialFormRef"
      />
    </SimpleDialog>

    <!-- Simple Dialog for Add Expense -->
    <SimpleDialog
      :open="showAddExpenseDialog"
      title="إضافة مصروف جديد"
      save-text="حفظ"
      icon="mdi-cash-multiple"
      @close="closeAddExpenseDialog"
      @save="saveExpense"
    >
      <ExpenseForm
        v-model="newExpense"
        ref="expenseFormRef"
      />
    </SimpleDialog>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import SimpleDialog from '@/components/SimpleDialog.vue'
import MaterialForm from '@/components/MaterialForm.vue'
import ExpenseForm from '@/components/ExpenseForm.vue'

const router = useRouter()

// Reactive data
const materialsSearch = ref('')
const expensesSearch = ref('')
const showAddDialog = ref(false)
const showAddExpenseDialog = ref(false)
const isMaterial = ref(true)
const formValid = ref(false)

// Form data
const newItem = ref({
  type: '',
  quantity: '',
  price: '',
  totalPrice: '',
  driver: '',
  details: ''
})

// Modern dialog data
const newMaterial = ref({
  type: '',
  quantity: '',
  price: '',
  totalPrice: '',
  driver: '',
  details: ''
})

const newExpense = ref({
  type: '',
  quantity: '',
  price: '',
  totalPrice: '',
  details: '',
  additionalNotes: ''
})

const materialFormRef = ref(null)
const expenseFormRef = ref(null)

// Table headers
const materialsHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'نوع المادة', key: 'type', sortable: true },
  { title: 'الكمية', key: 'quantity', sortable: true },
  { title: 'السعر', key: 'price', sortable: true },
  { title: 'السعر الكلي', key: 'total', sortable: true },
  { title: 'اسم السائق', key: 'driver', sortable: true },
  { title: 'التفاصيل', key: 'details', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

const expensesHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'نوع النثرية', key: 'type', sortable: true },
  { title: 'العدد', key: 'quantity', sortable: true },
  { title: 'الأجرة اليومية', key: 'price', sortable: true },
  { title: 'المجموع', key: 'total', sortable: true },
  { title: 'التفاصيل', key: 'details', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Sample data
const materialsData = ref([
  {
    id: 1,
    type: 'أسمنت',
    quantity: 30,
    price: 120,
    total: 3600,
    driver: 'أحمد علي',
    details: 'أسمنت عالي الجودة'
  },
  {
    id: 2,
    type: 'رمل',
    quantity: 50,
    price: 2500,
    total: 125000,
    driver: 'محمد حسن',
    details: 'رمل ناعم للبناء'
  },
  {
    id: 3,
    type: 'حديد',
    quantity: 2,
    price: 50000,
    total: 100000,
    driver: 'علي أحمد',
    details: 'حديد تسليح 12 مم'
  }
])

const expensesData = ref([
  {
    id: 1,
    type: 'وجبة غداء',
    quantity: 8,
    price: 5000,
    total: 40000,
    details: 'وجبة غداء للعمال'
  },
  {
    id: 2,
    type: 'نقل',
    quantity: 4,
    price: 10000,
    total: 40000,
    details: 'تكلفة النقل اليومية'
  }
])

// Methods
const goBack = () => {
  router.push('/work-day-details')
}

const searchMaterials = () => {
  console.log('البحث في المواد:', materialsSearch.value)
}

const searchExpenses = () => {
  console.log('البحث في المصاريف:', expensesSearch.value)
}

const openAddDialog = (material) => {
  isMaterial.value = material
  newItem.value = {
    type: '',
    quantity: '',
    price: '',
    driver: ''
  }
  showAddDialog.value = true
}

const openAddMaterialDialog = () => {
  showAddDialog.value = true
  // Reset form data
  newMaterial.value = {
    type: '',
    quantity: '',
    price: '',
    totalPrice: '',
    driver: '',
    details: ''
  }
}

const openAddExpenseDialog = () => {
  showAddExpenseDialog.value = true
}

const closeAddDialog = () => {
  showAddDialog.value = false
  newItem.value = {
    type: '',
    quantity: '',
    price: '',
    totalPrice: '',
    driver: '',
    details: ''
  }
}

const closeAddExpenseDialog = () => {
  showAddExpenseDialog.value = false
  if (expenseFormRef.value) {
    expenseFormRef.value.resetForm()
  }
  newExpense.value = {
    type: '',
    quantity: '',
    price: '',
    totalPrice: '',
    details: '',
    additionalNotes: ''
  }
}

const saveItem = () => {
  if (isMaterial.value) {
    const newMaterial = {
      id: materialsData.value.length + 1,
      type: newItem.value.type,
      quantity: parseInt(newItem.value.quantity),
      price: parseInt(newItem.value.price),
      total: parseInt(newItem.value.totalPrice) || (parseInt(newItem.value.quantity) * parseInt(newItem.value.price)),
      driver: newItem.value.driver,
      details: newItem.value.details
    }
    materialsData.value.push(newMaterial)
  } else {
    const newExpense = {
      id: expensesData.value.length + 1,
      type: newItem.value.type,
      quantity: parseInt(newItem.value.quantity),
      price: parseInt(newItem.value.price),
      total: parseInt(newItem.value.totalPrice) || (parseInt(newItem.value.quantity) * parseInt(newItem.value.price)),
      details: newItem.value.details
    }
    expensesData.value.push(newExpense)
  }
  closeAddDialog()
}

// Modern dialog functions
const saveMaterial = () => {
  if (materialFormRef.value && materialFormRef.value.isValid()) {
    const newMaterialItem = {
      id: materialsData.value.length + 1,
      type: newMaterial.value.type,
      quantity: parseInt(newMaterial.value.quantity),
      price: parseInt(newMaterial.value.price),
      total: parseInt(newMaterial.value.totalPrice) || (parseInt(newMaterial.value.quantity) * parseInt(newMaterial.value.price)),
      driver: newMaterial.value.driver,
      details: newMaterial.value.details
    }
    
    materialsData.value.push(newMaterialItem)
    
    // Reset form
    materialFormRef.value.resetForm()
    newMaterial.value = {
      type: '',
      quantity: '',
      price: '',
      totalPrice: '',
      driver: '',
      details: ''
    }
    
    showAddDialog.value = false
  }
}

// Modern dialog functions for expenses
const saveExpense = () => {
  if (expenseFormRef.value && expenseFormRef.value.isValid()) {
    const newExpenseItem = {
      id: expensesData.value.length + 1,
      type: newExpense.value.type,
      quantity: parseInt(newExpense.value.quantity),
      price: parseInt(newExpense.value.price),
      total: parseInt(newExpense.value.totalPrice) || (parseInt(newExpense.value.quantity) * parseInt(newExpense.value.price)),
      details: newExpense.value.details,
      additionalNotes: newExpense.value.additionalNotes
    }
    
    expensesData.value.push(newExpenseItem)
    
    // Reset form
    expenseFormRef.value.resetForm()
    newExpense.value = {
      type: '',
      quantity: '',
      price: '',
      totalPrice: '',
      details: '',
      additionalNotes: ''
    }
    
    showAddExpenseDialog.value = false
  }
}

const deleteMaterial = (item) => {
  if (confirm('هل أنت متأكد من حذف هذه المادة؟')) {
    const index = materialsData.value.findIndex(m => m.id === item.id)
    if (index > -1) {
      materialsData.value.splice(index, 1)
    }
  }
}

const deleteExpense = (item) => {
  if (confirm('هل أنت متأكد من حذف هذا المصروف؟')) {
    const index = expensesData.value.findIndex(e => e.id === item.id)
    if (index > -1) {
      expensesData.value.splice(index, 1)
    }
  }
}

// Lifecycle
onMounted(() => {
  console.log('✅ صفحة المواد والمصاريف تم تحميلها بنجاح!')
})
</script>

<style scoped>
.materials-expenses-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 2rem;
}

.page-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white;
  padding: 2.5rem;
  border-radius: 20px;
  margin-bottom: 2rem;
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.header-content {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.back-btn {
  background: rgba(255, 255, 255, 0.15) !important;
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.25) !important;
  transform: translateX(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
}

.page-title {
  font-size: 2.8rem;
  font-weight: 900;
  margin: 0;
  text-shadow: 0 3px 12px rgba(0, 0, 0, 0.6);
  letter-spacing: -0.5px;
  color: #ffffff !important;
  text-stroke: 1px rgba(0, 0, 0, 0.3);
  -webkit-text-stroke: 1px rgba(0, 0, 0, 0.3);
}

.page-subtitle {
  font-size: 1.3rem;
  opacity: 1;
  margin: 0.8rem 0 0 0;
  font-weight: 600;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
  color: #ffffff !important;
  text-stroke: 0.5px rgba(0, 0, 0, 0.2);
  -webkit-text-stroke: 0.5px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.5px;
}

.section-card {
  border-radius: 16px;
  overflow: hidden;
}

.section-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%);
  background-size: 200% 200%;
  animation: gradientShift 3s ease infinite;
  color: #ffffff !important;
  font-weight: 800;
  font-size: 1.4rem;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.3), 0 0 20px rgba(255, 255, 255, 0.3);
  padding: 1.25rem 1.5rem !important;
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(59, 130, 246, 0.4), 0 0 40px rgba(96, 165, 250, 0.2);
  transition: all 0.3s ease;
}

.section-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
}

.section-title::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.15) 0%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
}

.section-title:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(59, 130, 246, 0.5), 0 0 60px rgba(96, 165, 250, 0.3);
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

@keyframes pulse {
  0%, 100% { 
    opacity: 0.3; 
    transform: scale(0.8); 
  }
  50% { 
    opacity: 0.6; 
    transform: scale(1.2); 
  }
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* تحسين تصميم شريط البحث */
.search-container {
  padding: 1.5rem;
  background: linear-gradient(135deg, #eff6ff 0%, #dbeafe 50%, #bfdbfe 100%);
  border-radius: 12px;
  margin-bottom: 1rem;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.2);
  border: 2px solid #93c5fd;
  position: relative;
}

.search-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #60a5fa, #3b82f6, #2563eb);
  border-radius: 12px 12px 0 0;
}

.search-box {
  display: flex;
  align-items: center;
  gap: 1rem;
  max-width: 600px;
  margin: 0 auto;
}

.search-field {
  flex: 1;
  margin-bottom: 0;
}

.search-field :deep(.v-field) {
  background-color: #ffffff !important;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease;
}

.search-field :deep(.v-field__input) {
  background-color: #ffffff !important;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease;
}

.search-field :deep(.v-field__outline) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.search-field :deep(.v-field--focused .v-field__outline) {
  border-color: #2563eb !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2) !important;
}

.search-field :deep(.v-field__input):focus {
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.4) !important;
  border-color: #2563eb !important;
  transform: translateY(-1px);
}

.search-btn {
  height: 56px;
  min-width: 120px;
  border-radius: 8px;
  font-weight: bold;
  text-transform: none;
  box-shadow: 0 4px 12px rgba(25, 118, 210, 0.3);
  transition: all 0.3s ease;
}

.search-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(25, 118, 210, 0.4);
}

.search-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 8px rgba(25, 118, 210, 0.3);
}

/* أنماط أزرار الإضافة */
.add-btn {
  height: 56px;
  min-width: 140px;
  border-radius: 8px;
  font-weight: bold;
  text-transform: none;
  transition: all 0.3s ease;
}

.add-btn.success {
  background: linear-gradient(135deg, #4caf50 0%, #388e3c 100%) !important;
  box-shadow: 0 4px 12px rgba(76, 175, 80, 0.3);
}

.add-btn.warning {
  background: linear-gradient(135deg, #ff9800 0%, #f57c00 100%) !important;
  box-shadow: 0 4px 12px rgba(255, 152, 0, 0.3);
}

.add-btn:hover {
  transform: translateY(-2px);
}

.add-btn.success:hover {
  box-shadow: 0 6px 16px rgba(76, 175, 80, 0.4);
  background: linear-gradient(135deg, #388e3c 0%, #2e7d32 100%) !important;
}

.add-btn.warning:hover {
  box-shadow: 0 6px 16px rgba(255, 152, 0, 0.4);
  background: linear-gradient(135deg, #f57c00 0%, #ef6c00 100%) !important;
}

.add-btn:active {
  transform: translateY(0);
}

.add-btn :deep(.v-btn__content) {
  font-size: 1rem;
  font-weight: 600;
}

.add-btn :deep(.v-icon) {
  font-size: 1.2rem;
  transition: transform 0.3s ease;
}

.add-btn:hover :deep(.v-icon) {
  transform: scale(1.1);
}

.data-table {
  margin-top: 1rem;
}

/* تحسين ألوان النصوص في الجداول */
.data-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 600;
  font-size: 1rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.data-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: 900;
  font-size: 1.1rem;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  text-stroke: 0.5px rgba(0, 0, 0, 0.2);
  -webkit-text-stroke: 0.5px rgba(0, 0, 0, 0.2);
}

.data-table :deep(.v-data-table__tr:hover) {
  background-color: #f8f9fa !important;
}

.data-table :deep(.v-data-table__tr:hover .v-data-table__td) {
  color: #000 !important;
}

/* تحسين ألوان النصوص في الجداول الفرعية */
.data-table :deep(.v-data-table__wrapper) {
  background-color: white;
}

.data-table :deep(.v-data-table__wrapper table) {
  background-color: white;
}

.data-table :deep(.v-data-table__wrapper table td) {
  color: #333 !important;
  border-color: #e0e0e0 !important;
}

.data-table :deep(.v-data-table__wrapper table th) {
  color: #1976d2 !important;
  background-color: #f5f5f5 !important;
  border-color: #e0e0e0 !important;
}

/* أنماط مخصصة لجداول المواد والمصاريف */
.materials-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 700;
  font-size: 1rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.materials-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: 900;
  font-size: 1.1rem;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4);
  text-stroke: 0.5px rgba(0, 0, 0, 0.2);
  -webkit-text-stroke: 0.5px rgba(0, 0, 0, 0.2);
  border-bottom: 2px solid #0d47a1 !important;
  position: relative;
}

.expenses-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 700;
  font-size: 1rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.expenses-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: bold;
  font-size: 1rem;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
  border-bottom: 2px solid #0d47a1 !important;
  position: relative;
}

/* تحسين ألوان الأزرار في الجداول */
.data-table :deep(.v-btn) {
  color: #dc3545 !important;
  transition: all 0.3s ease;
}

.data-table :deep(.v-btn:hover) {
  background-color: #f8d7da !important;
  color: #721c24 !important;
  transform: scale(1.1);
  box-shadow: 0 2px 8px rgba(220, 53, 69, 0.3);
}

/* تحسينات إضافية لرؤوس الجداول */
.materials-table :deep(.v-data-table__th:hover),
.expenses-table :deep(.v-data-table__th:hover) {
  background: linear-gradient(135deg, #1565c0 0%, #0d47a1 100%) !important;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

/* تحسين ألوان الأيقونات في الرؤوس */
.materials-table :deep(.v-data-table__th .v-icon),
.expenses-table :deep(.v-data-table__th .v-icon) {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.3));
}

/* تحسين ألوان النصوص المحددة للفرز */
.materials-table :deep(.v-data-table__th.v-data-table__th--sorted),
.expenses-table :deep(.v-data-table__th.v-data-table__th--sorted) {
  background: linear-gradient(135deg, #0d47a1 0%, #1976d2 100%) !important;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.2);
}

/* تحسينات إضافية للوضوح */
.data-table :deep(.v-data-table__wrapper) {
  border: 2px solid #e3f2fd;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(25, 118, 210, 0.1);
  background: white;
}

/* تحسين ألوان الحدود */
.data-table :deep(.v-data-table__wrapper table) {
  border-collapse: separate;
  border-spacing: 0;
}

.data-table :deep(.v-data-table__wrapper table td) {
  border-right: 1px solid #f0f0f0;
  border-bottom: 1px solid #f0f0f0;
}

.data-table :deep(.v-data-table__wrapper table th) {
  border-right: 1px solid rgba(255, 255, 255, 0.2);
  border-bottom: 2px solid #0d47a1;
}

.data-table :deep(.v-data-table__wrapper table th:last-child) {
  border-right: none;
}

.data-table :deep(.v-data-table__wrapper table td:last-child) {
  border-right: none;
}

.data-table :deep(.v-data-table__wrapper table) {
  border-collapse: separate;
  border-spacing: 0;
}

.data-table :deep(.v-data-table__wrapper table tr:nth-child(even)) {
  background-color: #fafafa;
}

.data-table :deep(.v-data-table__wrapper table tr:nth-child(odd)) {
  background-color: #ffffff;
}

.data-table :deep(.v-data-table__wrapper table tr:hover) {
  background-color: #e3f2fd !important;
}

.data-table :deep(.v-data-table__wrapper table tr:hover td) {
  color: #000 !important;
  font-weight: 700;
}

/* تحسين ألوان النصوص في حالة عدم وجود بيانات */
.data-table :deep(.v-data-table__empty) {
  color: #666 !important;
  font-weight: 500;
  font-size: 1.1rem;
}

/* تحسين ألوان النصوص في حالة التحميل */
.data-table :deep(.v-data-table__loading) {
  color: #1976d2 !important;
  font-weight: 500;
  font-size: 1.1rem;
}

.dialog-card {
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border: none;
}

.dialog-title {
  background: #1976d2;
  color: white;
  font-weight: bold;
  font-size: 1.2rem;
  padding: 1.2rem 1.5rem;
  display: flex;
  align-items: center;
  position: relative;
  text-align: center;
  justify-content: center;
}

.close-btn {
  color: white !important;
  opacity: 0.9;
  transition: opacity 0.3s ease;
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
}

.close-btn:hover {
  opacity: 1;
  background: rgba(255, 255, 255, 0.1) !important;
}

.dialog-content {
  padding: 1.5rem;
  background: white;
}

.dialog-content .v-text-field,
.dialog-content .v-textarea {
  margin-bottom: 1.5rem;
}

/* تحسين تصميم الحقول لتبدو مثل الصورة */
.dialog-content .v-field {
  background-color: #ffffff !important;
  border-radius: 4px !important;
}

.dialog-content .v-field__input {
  color: #2c3e50 !important;
  font-weight: 500 !important;
  font-size: 1rem !important;
  background-color: #ffffff !important;
  padding: 12px 16px !important;
  min-height: 48px !important;
}

.dialog-content .v-label {
  color: #424242 !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
  letter-spacing: 0.3px;
}

.dialog-content .v-field--focused .v-label {
  color: #1976d2 !important;
  font-weight: 600 !important;
  transform: scale(0.85) !important;
}

.dialog-content .v-field__outline {
  border-color: #e0e0e0 !important;
  border-width: 1px !important;
  border-radius: 4px !important;
}

.dialog-content .v-field--focused .v-field__outline {
  border-color: #1976d2 !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 1px rgba(25, 118, 210, 0.2) !important;
}

.dialog-content .v-field__outline__start,
.dialog-content .v-field__outline__end {
  border-color: #e0e0e0 !important;
  border-width: 1px !important;
}

.dialog-content .v-field--focused .v-field__outline__start,
.dialog-content .v-field--focused .v-field__outline__end {
  border-color: #1976d2 !important;
  border-width: 2px !important;
}

/* تحسين وضوح النصوص في الحقول */
.dialog-content .v-field__input::placeholder {
  color: #95a5a6 !important;
  font-weight: 400 !important;
  opacity: 0.8 !important;
  font-size: 0.95rem !important;
}

.dialog-content .v-field__input:focus {
  color: #2c3e50 !important;
  font-weight: 600 !important;
  background-color: #f8f9fa !important;
  border-color: #1976d2 !important;
}

/* تحسين وضوح رسائل الخطأ */
.dialog-content .v-messages__message {
  color: #d32f2f !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  text-shadow: 0 2px 4px rgba(211, 47, 47, 0.3);
  background-color: #ffebee !important;
  padding: 4px 8px !important;
  border-radius: 4px !important;
  border-left: 4px solid #d32f2f !important;
}

/* تحسين وضوح الأيقونات */
.dialog-content .v-field__append-inner .v-icon {
  color: #0d47a1 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 2px 4px rgba(13, 71, 161, 0.4);
}

.dialog-content .v-field--focused .v-field__append-inner .v-icon {
  color: #1565c0 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(21, 101, 192, 0.5);
}

/* تحسين وضوح الحقول بشكل عام */
.dialog-content .v-text-field,
.dialog-content .v-textarea {
  background-color: #ffffff !important;
  border-radius: 6px !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  margin-bottom: 1.5rem !important;
}

/* تحسين وضوح عناوين المربعات */
.dialog-content .v-field__field {
  position: relative;
}

.dialog-content .v-field__field::before {
  content: '';
  position: absolute;
  top: -8px;
  left: 12px;
  right: 12px;
  height: 1px;
  background: linear-gradient(90deg, #1a237e, #1976d2, #1a237e);
  opacity: 0.3;
  z-index: 1;
}

.dialog-content .v-text-field:hover,
.dialog-content .v-textarea:hover {
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15) !important;
  transition: all 0.3s ease !important;
}

/* تحسين وضوح النصوص في منطقة النص */
.dialog-content .v-textarea .v-field__input {
  min-height: 80px !important;
  padding: 12px !important;
  line-height: 1.5 !important;
  color: #2c3e50 !important;
  font-weight: 500 !important;
  font-size: 1rem !important;
}

/* تحسين وضوح النصوص في رأس الحوار */
.dialog-title {
  color: #ffffff !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.5);
  letter-spacing: 0.5px;
}

/* تحسين وضوح زر الحفظ */
.save-btn {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  letter-spacing: 0.5px;
}

.dialog-actions {
  padding: 1rem 1.5rem;
  background: white;
  border-top: 1px solid #e0e0e0;
  display: flex;
  justify-content: flex-end;
}

.save-btn {
  background: #1976d2 !important;
  color: white !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  border-radius: 4px !important;
  text-transform: none !important;
  box-shadow: 0 2px 4px rgba(25, 118, 210, 0.3) !important;
  transition: all 0.3s ease !important;
}

.save-btn:hover {
  background: #1565c0 !important;
  box-shadow: 0 4px 8px rgba(25, 118, 210, 0.4) !important;
}

.actions-card {
  background: white;
  border-radius: 16px;
}

/* Animations */
.glass-effect {
  backdrop-filter: blur(15px);
}

.page-header.glass-effect {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
}

.gradient-animation {
  background-size: 400% 400%;
  animation: gradientShift 20s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.text-glow {
  text-shadow: 0 0 30px rgba(255, 255, 255, 0.6);
}

.fade-in {
  animation: fadeIn 1.2s ease-out;
}

@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(30px) scale(0.95); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0) scale(1); 
  }
}

/* تأثيرات إضافية للرأس */
.page-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.page-header::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
}

@keyframes pulse {
  0%, 100% { opacity: 0.3; transform: scale(0.8); }
  50% { opacity: 0.6; transform: scale(1.2); }
}

/* تحسينات إضافية للبحث */
.search-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, #1976d2, #42a5f5, #1976d2);
  border-radius: 12px 12px 0 0;
}

.search-container {
  position: relative;
}

.search-field :deep(.v-label) {
  color: #1e40af !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
}

.search-field :deep(.v-label--active),
.search-field :deep(.v-field__label--active) {
  color: #1e40af !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.search-field :deep(.v-field--focused .v-label) {
  color: #1e40af !important;
  font-weight: 700 !important;
}

.search-field :deep(.v-field__input input) {
  color: #111827 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

/* تحسين ألوان الأيقونات */
.search-field :deep(.v-field__prepend-inner) {
  color: #3b82f6 !important;
}

.search-field :deep(.v-field__prepend-inner .v-icon) {
  color: #3b82f6 !important;
  font-size: 24px !important;
  transition: color 0.3s ease;
}

.search-field :deep(.v-field__prepend-inner i) {
  color: #3b82f6 !important;
}

.search-field :deep(.v-field--focused .v-field__prepend-inner .v-icon) {
  color: #2563eb !important;
  transform: scale(1.1);
}

/* تحسين ألوان زر البحث */
.search-btn :deep(.v-btn__content) {
  font-size: 1rem;
  font-weight: 600;
}

.search-btn :deep(.v-icon) {
  font-size: 1.2rem;
  transition: transform 0.3s ease;
}

.search-btn:hover :deep(.v-icon) {
  transform: scale(1.1);
}

/* Responsive */
@media (max-width: 768px) {
  .materials-expenses-page {
    padding: 1rem;
  }
  
  .page-title {
    font-size: 2rem;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
  }
  
  .dialog-content {
    padding: 1rem;
  }
  
  .search-box {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .search-btn, .add-btn {
    width: 100%;
    min-width: auto;
  }
  
  .search-container {
    padding: 0.75rem;
  }
  
  .search-field {
    margin-bottom: 0.5rem;
  }
}
</style>
