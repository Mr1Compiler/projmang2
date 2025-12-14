<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">📦</span>
        <h1 class="page-title text-glow fade-in">إدارة المخزون</h1>
        <p class="page-subtitle fade-in">إدارة وتتبع جميع عناصر المخزون والمواد</p>
      </div>

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="primary">mdi-package-variant</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-primary mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalItems || 0 }}</h3>
            <p class="text-subtitle-1 text-primary mb-0">إجمالي العناصر</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="success">mdi-check-circle</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ inStockItems || 0 }}</h3>
            <p class="text-subtitle-1 text-success mb-0">متوفر في المخزون</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="warning">mdi-alert</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ lowStockItems || 0 }}</h3>
            <p class="text-subtitle-1 text-warning mb-0">منخفض المخزون</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="info">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h6 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; white-space: nowrap;">{{ formatCurrency(totalValue) || '0 د.ع' }}</h3>
            <p class="text-subtitle-1 text-info mb-0">إجمالي القيمة</p>
          </v-card>
        </v-col>
      </v-row>

      <!-- Search Bar -->
      <v-card class="search-card mb-4" elevation="2">
        <v-card-text class="pa-4">
          <v-row class="align-center">
            <v-col cols="12" md="4">
              <v-text-field
                v-model="searchQuery"
                label="البحث في المخزون..."
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
                v-model="selectedCategory"
                :items="categoryOptions"
                label="الفئة"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="black-list"
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
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-btn
                color="primary"
                variant="elevated"
                size="large"
                class="search-btn"
                @click="searchInventory"
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
                @click="openAddItemDialog"
              >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة عنصر جديد
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Inventory Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title">
          <span class="title-text">عناصر المخزون</span>
        </v-card-title>

        <v-data-table
          :headers="headers"
          :items="inventoryItems"
          :search="searchQuery"
          class="project-table"
          :items-per-page="10"
          :loading="loading"
          hover
          no-data-text="لا توجد بيانات"
        >
          <!-- Serial Number Column -->
          <template #item.serial="{ index }">
            <span class="serial-number">{{ index + 1 }}</span>
          </template>

          <!-- Item Name Column -->
          <template #item.name="{ item }">
            <span class="project-name">{{ item.name }}</span>
          </template>

          <!-- Category Column -->
          <template #item.category="{ item }">
            <v-chip class="category-chip" size="small">
              {{ item.category }}
            </v-chip>
          </template>

          <!-- Quantity Column -->
          <template #item.quantity="{ item }">
            <span class="cost-text">{{ item.quantity }} {{ item.unit }}</span>
          </template>

          <!-- Unit Price Column -->
          <template #item.unitPrice="{ item }">
            <span class="cost-text">{{ formatCurrency(item.unitPrice) }}</span>
          </template>

          <!-- Total Value Column -->
          <template #item.totalValue="{ item }">
            <span class="cost-text">{{ formatCurrency(item.totalValue) }}</span>
          </template>

          <!-- Status Column -->
          <template #item.status="{ item }">
            <v-chip class="status-chip" size="small">
              {{ getStatusText(item.status) }}
            </v-chip>
          </template>

          <!-- Location Column -->
          <template #item.location="{ item }">
            <span class="location-text">{{ item.location || 'غير محدد' }}</span>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="action-buttons">
              <v-btn
                size="small"
                color="primary"
                variant="text"
                @click="viewItemDetails(item)"
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
                @click="editItem(item)"
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
                @click="deleteItem(item)"
                icon
                class="action-btn"
                title="حذف"
              >
                <v-icon size="16">mdi-delete</v-icon>
              </v-btn>
            </div>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Item Dialog -->
      <v-dialog v-model="itemDialog" max-width="800" scrollable persistent>
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-package-variant</v-icon>
                <span class="header-title">{{ isEditing ? 'تعديل عنصر المخزون' : 'إضافة عنصر جديد' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeItemDialog"
                class="close-btn"
              />
            </div>
          </div>

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="itemForm" v-model="itemFormValid">
              <div class="form-fields">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="itemForm.name"
                      label="اسم العنصر"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم العنصر مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="itemForm.category"
                      :items="categories"
                      label="الفئة"
                      variant="outlined"
                      :rules="[v => !!v || 'الفئة مطلوبة']"
                      required
                      class="form-field black-list"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="itemForm.quantity"
                      label="الكمية"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'الكمية مطلوبة', v => v >= 0 || 'الكمية يجب أن تكون أكبر من أو تساوي صفر']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="itemForm.unit"
                      label="وحدة القياس"
                      variant="outlined"
                      :rules="[v => !!v || 'وحدة القياس مطلوبة']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="itemForm.unitPrice"
                      label="سعر الوحدة (د.ع)"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'سعر الوحدة مطلوب', v => v > 0 || 'سعر الوحدة يجب أن يكون أكبر من صفر']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="itemForm.minStock"
                      label="الحد الأدنى للمخزون"
                      variant="outlined"
                      type="number"
                      :rules="[v => v >= 0 || 'الحد الأدنى يجب أن يكون أكبر من أو يساوي صفر']"
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="itemForm.maxStock"
                      label="الحد الأقصى للمخزون"
                      variant="outlined"
                      type="number"
                      :rules="[v => v >= 0 || 'الحد الأقصى يجب أن يكون أكبر من أو يساوي صفر']"
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="itemForm.location"
                      label="مكان التخزين"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="itemForm.status"
                      :items="statusOptions"
                      label="الحالة"
                      variant="outlined"
                      :rules="[v => !!v || 'الحالة مطلوبة']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="itemForm.supplier"
                      label="المورد"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12">
                    <v-textarea
                      v-model="itemForm.notes"
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
              @click="closeItemDialog"
              class="action-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveItem"
              :disabled="!itemFormValid"
              class="action-btn primary-btn"
            >
              <v-icon class="me-2">mdi-content-save</v-icon>
              {{ isEditing ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </div>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

// متغيرات الحالة الأساسية
const loading = ref(false)
const itemDialog = ref(false)
const itemFormValid = ref(false)
const isEditing = ref(false)
const searchQuery = ref('')
const selectedCategory = ref('')
const selectedStatus = ref('')
const selectedItem = ref(null)

// عناوين الجدول
const headers = ref([
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'اسم العنصر', key: 'name', sortable: true, align: 'right' },
  { title: 'الفئة', key: 'category', sortable: true, align: 'center' },
  { title: 'الكمية', key: 'quantity', sortable: true, align: 'center' },
  { title: 'سعر الوحدة', key: 'unitPrice', sortable: true, align: 'center' },
  { title: 'القيمة الإجمالية', key: 'totalValue', sortable: true, align: 'center' },
  { title: 'الحالة', key: 'status', sortable: true, align: 'center' },
  { title: 'مكان التخزين', key: 'location', sortable: true, align: 'center' },
  { title: 'الإجراءات', key: 'actions', sortable: false, align: 'center' }
])

// نموذج العنصر
const itemForm = ref({
  name: '',
  category: '',
  quantity: 0,
  unit: '',
  unitPrice: 0,
  minStock: 0,
  maxStock: 0,
  location: '',
  status: 'متوفر',
  supplier: '',
  notes: ''
})

// بيانات المخزون التجريبية
const inventoryItems = ref([
  {
    id: 1,
    name: 'أسمنت بورتلاند',
    category: 'مواد بناء',
    quantity: 500,
    unit: 'كيس',
    unitPrice: 5000,
    totalValue: 2500000,
    minStock: 100,
    maxStock: 1000,
    location: 'المستودع الرئيسي',
    status: 'متوفر',
    supplier: 'شركة مواد البناء',
    notes: 'جودة عالية'
  },
  {
    id: 2,
    name: 'حديد تسليح',
    category: 'مواد بناء',
    quantity: 50,
    unit: 'طن',
    unitPrice: 800000,
    totalValue: 40000000,
    minStock: 20,
    maxStock: 100,
    location: 'المستودع الرئيسي',
    status: 'متوفر',
    supplier: 'شركة الحديد',
    notes: ''
  },
  {
    id: 3,
    name: 'طلاء أبيض',
    category: 'دهانات',
    quantity: 5,
    unit: 'علبة',
    unitPrice: 15000,
    totalValue: 75000,
    minStock: 10,
    maxStock: 50,
    location: 'المستودع الفرعي',
    status: 'منخفض',
    supplier: 'شركة الدهانات',
    notes: 'يحتاج إعادة طلب'
  }
])

// الفئات
const categories = [
  'مواد بناء',
  'دهانات',
  'أدوات',
  'كهرباء',
  'سباكة',
  'أخرى'
]

const categoryOptions = ref([
  'جميع الفئات',
  ...categories
])

const statusOptions = ref([
  'متوفر',
  'منخفض',
  'نفد',
  'محجوز'
])

// إحصائيات المخزون
const totalItems = computed(() => inventoryItems.value.length)
const inStockItems = computed(() => {
  return inventoryItems.value.filter(item => item.status === 'متوفر').length
})
const lowStockItems = computed(() => {
  return inventoryItems.value.filter(item => {
    return item.status === 'منخفض' || (item.minStock > 0 && item.quantity <= item.minStock)
  }).length
})
const totalValue = computed(() => {
  return inventoryItems.value.reduce((sum, item) => sum + item.totalValue, 0)
})

// دوال المخزون
const searchInventory = () => {
  console.log('البحث عن:', searchQuery.value)
}

const openAddItemDialog = () => {
  itemDialog.value = true
  isEditing.value = false
  selectedItem.value = null
  itemForm.value = {
    name: '',
    category: '',
    quantity: 0,
    unit: '',
    unitPrice: 0,
    minStock: 0,
    maxStock: 0,
    location: '',
    status: 'متوفر',
    supplier: '',
    notes: ''
  }
}

const closeItemDialog = () => {
  itemDialog.value = false
  isEditing.value = false
  selectedItem.value = null
  itemForm.value = {
    name: '',
    category: '',
    quantity: 0,
    unit: '',
    unitPrice: 0,
    minStock: 0,
    maxStock: 0,
    location: '',
    status: 'متوفر',
    supplier: '',
    notes: ''
  }
}

const editItem = (item) => {
  selectedItem.value = item
  isEditing.value = true
  itemForm.value = { ...item }
  itemDialog.value = true
}

const viewItemDetails = (item) => {
  console.log('عرض تفاصيل العنصر:', item)
  // يمكن إضافة نافذة عرض التفاصيل هنا
}

const deleteItem = (item) => {
  if (confirm(`هل أنت متأكد من حذف "${item.name}"؟`)) {
    const index = inventoryItems.value.findIndex(i => i.id === item.id)
    if (index > -1) {
      inventoryItems.value.splice(index, 1)
    }
  }
}

const saveItem = () => {
  const totalValue = parseFloat(itemForm.value.quantity) * parseFloat(itemForm.value.unitPrice)
  
  if (isEditing.value) {
    // تحديث العنصر
    const index = inventoryItems.value.findIndex(i => i.id === selectedItem.value.id)
    if (index > -1) {
      inventoryItems.value[index] = {
        ...itemForm.value,
        id: selectedItem.value.id,
        quantity: parseFloat(itemForm.value.quantity),
        unitPrice: parseFloat(itemForm.value.unitPrice),
        totalValue: totalValue
      }
    }
  } else {
    // إضافة عنصر جديد
    const newItem = {
      ...itemForm.value,
      id: Date.now(),
      quantity: parseFloat(itemForm.value.quantity),
      unitPrice: parseFloat(itemForm.value.unitPrice),
      totalValue: totalValue
    }
    inventoryItems.value.push(newItem)
  }
  closeItemDialog()
}

// دوال مساعدة
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount).replace('IQD', 'د.ع')
}

const getCategoryColor = (category) => {
  const colors = {
    'مواد بناء': 'primary',
    'دهانات': 'info',
    'أدوات': 'warning',
    'كهرباء': 'success',
    'سباكة': 'error',
    'أخرى': 'grey'
  }
  return colors[category] || 'grey'
}

const getStatusColor = (status) => {
  const colors = {
    'متوفر': 'success',
    'منخفض': 'warning',
    'نفد': 'error',
    'محجوز': 'info'
  }
  return colors[status] || 'grey'
}

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

.page-header {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(20px);
  border-radius: 12px;
  padding: 16px 24px;
  margin-bottom: 20px;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.2);
  position: relative;
  overflow: hidden;
}

.page-icon {
  font-size: 2rem;
  display: block;
  margin-bottom: 8px;
  animation: star-twinkle 2s ease-in-out infinite;
}

.page-title {
  font-size: 1.5rem;
  font-weight: 800;
  color: rgb(246, 246, 246);
  margin-bottom: 6px;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.page-subtitle {
  font-size: 0.9rem;
  color: rgba(243, 240, 240, 0.9);
  margin: 0;
}

.stats-row {
  margin-bottom: 30px;
  width: 100%;
  overflow-x: hidden;
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

.search-card {
  background: linear-gradient(135deg, #4783cc 0%, #4d5abb 100%) !important;
  border: 2px solid #365e93 !important;
  border-radius: 12px !important;
  box-shadow: 0 4px 12px rgba(200, 11, 11, 0.1) !important;
  width: 100%;
  overflow-x: hidden;
  box-sizing: border-box;
}

.data-table-card {
  background: rgba(255, 249, 249, 0.95);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.table-title {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  padding: 20px 24px !important;
  border-radius: 12px 12px 0 0 !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.3) !important;
}

.title-text {
  font-size: 1.4rem !important;
  font-weight: 800 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* ========================================
   تنسيق جدول المخزون - Inventory Table
   ======================================== */
.project-table {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

/* عناوين الجدول - Header Styles */
.project-table .v-data-table__th {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-align: center !important;
  padding: 18px 14px !important;
  border-bottom: 3px solid #6d28d9 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 1px !important;
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
  background: linear-gradient(90deg, #ffffff 0%, rgba(255, 255, 255, 0.3) 50%, #ffffff 100%);
  opacity: 0.3;
}

.project-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-align: center !important;
  padding: 18px 14px !important;
  border-bottom: 3px solid #6d28d9 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 1px !important;
  position: relative !important;
}

.project-table .v-data-table__wrapper table thead tr th .v-data-table__th__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 1px !important;
}

/* بيانات الجدول - Table Data Styles - نص أسود تماماً */
.project-table .v-data-table__td {
  padding: 16px 12px !important;
  border-bottom: 1px solid #e5e7eb !important;
  text-align: center !important;
  background: #ffffff !important;
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  vertical-align: middle !important;
}

/* جميع النصوص في الجدول - أسود */
.project-table .v-data-table__td *,
.project-table .v-data-table__td span,
.project-table .v-data-table__td div,
.project-table .v-data-table__td p {
  color: #000000 !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(even) {
  background: linear-gradient(135deg, #faf5ff 0%, #f3e8ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(odd) {
  background: linear-gradient(135deg, #ffffff 0%, #faf5ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:hover {
  background: linear-gradient(135deg, #f3e8ff 0%, #e9d5ff 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.15) !important;
  transition: all 0.3s ease !important;
}

.project-table .v-data-table__wrapper tbody tr:hover td {
  color: #000000 !important;
  font-weight: 700 !important;
}

/* تنسيق النصوص داخل الخلايا - جميعها أسود */
.project-table .serial-number {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 6px 12px !important;
  border-radius: 8px !important;
  display: inline-block !important;
  min-width: 40px !important;
  text-align: center !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.project-table .project-name {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  text-align: right !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
  display: inline-block !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.project-table .cost-text {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 6px 12px !important;
  border-radius: 8px !important;
  display: inline-block !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #d1d5db !important;
}

.project-table .location-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
  display: inline-block !important;
  background: #f9fafb !important;
}

/* تنسيق الشريحات - Chips - نص أسود */
.project-table .v-chip {
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  padding: 8px 14px !important;
  border-radius: 8px !important;
  color: #000000 !important;
  background: #f3f4f6 !important;
  border: 1px solid #e5e7eb !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  transition: all 0.3s ease !important;
}

.project-table .v-chip:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15) !important;
  background: #e5e7eb !important;
}

/* شريحات الفئة - نص أسود */
.project-table .category-chip {
  color: #000000 !important;
  background: #f3f4f6 !important;
  border: 1px solid #d1d5db !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
}

.project-table .category-chip:hover {
  background: #e5e7eb !important;
  border-color: #9ca3af !important;
  color: #000000 !important;
}

/* شريحات الحالة - نص أسود */
.project-table .status-chip {
  color: #000000 !important;
  background: #f3f4f6 !important;
  border: 1px solid #d1d5db !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
}

.project-table .status-chip:hover {
  background: #e5e7eb !important;
  border-color: #9ca3af !important;
  color: #000000 !important;
}

/* حدود الجدول */
.project-table .v-data-table__wrapper {
  border: 1px solid #e5e7eb !important;
  border-radius: 12px !important;
  overflow: hidden !important;
  background: white !important;
}

.project-table .v-data-table__wrapper table {
  border-collapse: separate !important;
  border-spacing: 0 !important;
}

.project-table .v-data-table__wrapper table td {
  border-right: 1px solid #f1f5f9 !important;
}

.project-table .v-data-table__wrapper table td:last-child {
  border-right: none !important;
}

.project-table .v-data-table__wrapper table thead tr th {
  border-right: 1px solid rgba(255, 255, 255, 0.2) !important;
}

.project-table .v-data-table__wrapper table thead tr th:last-child {
  border-right: none !important;
}

.action-buttons {
  display: flex;
  gap: 4px;
  align-items: center;
  justify-content: center;
}

.action-btn {
  min-width: 36px !important;
  height: 36px !important;
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.action-btn:hover {
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15) !important;
}

.details-btn {
  color: #7c3aed !important;
  background: rgba(124, 58, 237, 0.1) !important;
}

.details-btn:hover {
  background: rgba(124, 58, 237, 0.2) !important;
  color: #6d28d9 !important;
}

.action-btn:not(.details-btn) {
  margin: 0 2px !important;
}

.action-btn[color="success"] {
  color: #059669 !important;
  background: rgba(5, 150, 105, 0.1) !important;
}

.action-btn[color="success"]:hover {
  background: rgba(5, 150, 105, 0.2) !important;
  color: #047857 !important;
}

.action-btn[color="error"] {
  color: #dc2626 !important;
  background: rgba(220, 38, 38, 0.1) !important;
}

.action-btn[color="error"]:hover {
  background: rgba(220, 38, 38, 0.2) !important;
  color: #b91c1c !important;
}

.image-style-dialog {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%);
  border-radius: 20px;
  overflow: hidden;
}

.dialog-header {
  background: linear-gradient(135deg, #059669, #10b981);
  padding: 20px;
  color: white;
}

.dialog-body {
  padding: 30px;
  background: #f8fafc;
}

.dialog-actions {
  padding: 20px 30px;
  background: #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

@keyframes star-twinkle {
  0%, 100% { transform: scale(1) rotate(0deg); }
  50% { transform: scale(1.1) rotate(180deg); }
}

/* ========================================
   تحسين وضوح النصوص في نموذج التعديل
   ======================================== */

/* ========================================
   تنسيق الحقول - نفس تنسيق صفحة المصروفات
   ======================================== */

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

/* ========================================
   تنسيق القوائم المنسدلة - Dropdown Menus
   ======================================== */

/* القائمة المنسدلة نفسها */
.v-menu__content,
.v-overlay__content {
  background: white !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid #e5e7eb !important;
  overflow: hidden !important;
}

.v-list {
  background: white !important;
  padding: 4px 0 !important;
}



.image-style-dialog .v-field--focused .v-field__label,
.image-style-dialog .v-field--focused .v-label {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.dialog-body .v-label,
.dialog-body .v-field__label,
.dialog-body .v-label--active,
.dialog-body .v-field__label--active,
.dialog-body .v-label--floating,
.dialog-body .v-field__label--floating {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  background: white !important;
  opacity: 1 !important;
  text-shadow: none !important;
}

.dialog-body .v-field--focused .v-label,
.dialog-body .v-field--focused .v-field__label,
.dialog-body .v-field--active .v-label,
.dialog-body .v-field--active .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
}

/* جميع التسميات في النموذج */
.image-style-dialog .v-label,
.image-style-dialog .v-field__label,
.image-style-dialog .v-label--active,
.image-style-dialog .v-field__label--active {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  opacity: 1 !important;
  text-shadow: none !important;
}

.image-style-dialog .v-field--focused .v-label,
.image-style-dialog .v-field--focused .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
}

/* التسميات داخل النموذج - قواعد إضافية */
.image-style-dialog .v-field__label,
.image-style-dialog .v-label,
.image-style-dialog .v-field__label--active,
.image-style-dialog .v-label--active {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  opacity: 1 !important;
  background: white !important;
  text-shadow: none !important;
}

.image-style-dialog .v-field--focused .v-field__label,
.image-style-dialog .v-field--focused .v-label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
}

/* القوائم المنسدلة المفتوحة */
/* جميع القوائم المنسدلة في الصفحة - نص أسود تماماً */
.v-menu__content .v-list,
.v-menu__content .v-list-item,
.v-menu__content .v-list-item-title,
.v-menu__content .v-list-item *,
.v-overlay__content .v-list,
.v-overlay__content .v-list-item,
.v-overlay__content .v-list-item * {
  background: white !important;
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.v-menu__content .v-list-item:hover,
.v-menu__content .v-list-item:hover *,
.v-overlay__content .v-list-item:hover,
.v-overlay__content .v-list-item:hover * {
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%) !important;
  color: #000000 !important;
  font-weight: 700 !important;
  transform: translateX(-4px) !important;
  border-right: 3px solid #059669 !important;
}

.v-menu__content .v-list-item--active,
.v-menu__content .v-list-item--active *,
.v-overlay__content .v-list-item--active,
.v-overlay__content .v-list-item--active * {
  background: #f3f4f6 !important;
  color: #000000 !important;
  font-weight: 700 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
  border: 2px solid #059669 !important;
}

/* القائمة المنسدلة في النموذج - نص أسود تماماً */
.dialog-body .v-list-item,
.dialog-body .v-list-item-title,
.dialog-body .v-list-item *,
.dialog-body .v-list-item-title * {
  background: white !important;
  color: #000000 !important;
  padding: 12px 16px !important;
}

.dialog-body .v-list-item:hover,
.dialog-body .v-list-item:hover * {
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%) !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.dialog-body .v-list-item--active,
.dialog-body .v-list-item--active *,
.dialog-body .v-list-item--active .v-list-item-title {
  background: #f3f4f6 !important;
  color: #000000 !important;
  font-weight: 700 !important;
  border: 2px solid #059669 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

/* القائمة المنسدلة الخاصة بالفئة - نص أسود تماماً */
.black-list .v-list-item,
.black-list .v-list-item-title,
.black-list .v-list-item *,
.black-list .v-list-item-title * {
  background: white !important;
  color: #000000 !important;
  font-weight: 600 !important;
  padding: 12px 16px !important;
}

.black-list .v-list-item:hover,
.black-list .v-list-item:hover * {
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%) !important;
  color: #000000 !important;
  font-weight: 700 !important;
  transform: translateX(-4px) !important;
  border-right: 3px solid #059669 !important;
}

.black-list .v-list-item--active,
.black-list .v-list-item--active *,
.black-list .v-list-item--active .v-list-item-title {
  background: #f3f4f6 !important;
  color: #000000 !important;
  font-weight: 700 !important;
  border: 2px solid #059669 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}


/* تحسين التمرير */
.project-table .v-data-table__wrapper::-webkit-scrollbar {
  height: 8px !important;
  width: 8px !important;
}

.project-table .v-data-table__wrapper::-webkit-scrollbar-track {
  background: #f8fafc !important;
  border-radius: 6px !important;
}

.project-table .v-data-table__wrapper::-webkit-scrollbar-thumb {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  border-radius: 6px !important;
  transition: all 0.3s ease !important;
}

.project-table .v-data-table__wrapper::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(135deg, #6d28d9 0%, #7c3aed 100%) !important;
  box-shadow: 0 2px 4px rgba(124, 58, 237, 0.3) !important;
}

/* التصميم المتجاوب */
@media (max-width: 768px) {
  .page-title {
    font-size: 2rem;
  }
  
  .page-subtitle {
    font-size: 1rem;
  }
  
  .project-table .v-data-table__wrapper {
    overflow-x: auto !important;
  }
  
  .project-table .v-data-table__wrapper table {
    min-width: 900px !important;
  }
  
  .project-table .v-data-table__td {
    padding: 12px 8px !important;
    font-size: 0.9rem !important;
  }
  
  .project-table .v-data-table__th {
    padding: 14px 10px !important;
    font-size: 1rem !important;
  }
}

@media (max-width: 480px) {
  .project-table .v-data-table__wrapper table {
    min-width: 800px !important;
  }
  
  .project-table .serial-number,
  .project-table .project-name,
  .project-table .cost-text,
  .project-table .location-text {
    font-size: 0.85rem !important;
  }
}
</style>

