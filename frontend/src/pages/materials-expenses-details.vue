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
}

const searchExpenses = () => {
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
})
</script>

<style>
/* Import shared page styles */
@import './styles/materials-expenses.css';
</style>

<style scoped>
/* Component-specific styles (if any) */
</style>
