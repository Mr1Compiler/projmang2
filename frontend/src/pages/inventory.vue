<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <PageHeader
        title="إدارة المخزون"
        subtitle="إدارة وتتبع جميع عناصر المخزون والمواد"
        icon="📦"
      />

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="primary">mdi-package-variant</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-primary mb-2 stat-number-ltr">{{ totalItems || 0 }}</h3>
            <p class="text-subtitle-1 text-primary mb-0">إجمالي العناصر</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="success">mdi-check-circle</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number-ltr">{{ inStockItems || 0 }}</h3>
            <p class="text-subtitle-1 text-success mb-0">متوفر في المخزون</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="warning">mdi-alert</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number-ltr">{{ lowStockItems || 0 }}</h3>
            <p class="text-subtitle-1 text-warning mb-0">منخفض المخزون</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="info">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h6 font-weight-bold text-info mb-2 stat-number-ltr-nowrap">{{ formatCurrency(totalValue) || '0 د.ع' }}</h3>
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
                class="search-field bg-filter-input"
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
import { formatCurrency } from '@/utils/formatters'
import { PageHeader } from '@/components/shared'

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
/* Import page styles - scoped to this component only */
@import './styles/inventory.css';
</style>
