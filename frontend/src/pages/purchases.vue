<template>
  <v-container class="fill-height data-page purchases-page" fluid>
    <div class="fullscreen-content centered-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">🛒</span>
        <h1 class="page-title text-glow fade-in">إدارة المشتريات</h1>
        <p class="page-subtitle fade-in">إدارة وتتبع جميع عمليات الشراء والمشتريات</p>
      </div>

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="primary">mdi-cart</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-primary mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(totalPurchases) || '0' }}</h3>
            <p class="text-subtitle-1 text-primary mb-0">إجمالي المشتريات</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="error">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-error mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; font-size: var(--font-size-base-plus) !important;">{{ formatCurrency(totalExpenses) || '0 د.ع' }}</h3>
            <p class="text-subtitle-1 text-error mb-0">إجمالي المصروفات</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="info">mdi-truck-delivery</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-info mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(totalSuppliers) || '0' }}</h3>
            <p class="text-subtitle-1 text-info mb-0">إجمالي الموردين</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(pendingPurchases) || '0' }}</h3>
            <p class="text-subtitle-1 text-warning mb-0">مشتريات معلقة</p>
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
                label="البحث في المشتريات..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="search-field"
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
                class="black-list"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-select
                v-model="selectedSupplier"
                :items="supplierOptions"
                label="المورد"
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
                @click="searchPurchases"
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
                @click="openAddPurchaseDialog"
              >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة عملية شراء
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Purchases Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title">
          <span class="title-text">عمليات المشتريات</span>
        </v-card-title>

        <v-data-table
          :headers="headers"
          :items="purchaseItems"
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

          <!-- Purchase Date Column -->
          <template #item.purchaseDate="{ item }">
            <span class="date-text">{{ item.purchaseDate }}</span>
          </template>

          <!-- Supplier Name Column -->
          <template #item.supplierName="{ item }">
            <span class="project-name">{{ item.supplierName }}</span>
          </template>

          <!-- Item Name Column -->
          <template #item.itemName="{ item }">
            <span class="project-name">{{ item.itemName }}</span>
          </template>

          <!-- Quantity Column -->
          <template #item.quantity="{ item }">
            <span class="quantity-text">{{ formatNumber(item.quantity) }} {{ item.unit }}</span>
          </template>

          <!-- Unit Price Column -->
          <template #item.unitPrice="{ item }">
            <span class="cost-text">{{ formatCurrency(item.unitPrice) }}</span>
          </template>

          <!-- Total Amount Column -->
          <template #item.totalAmount="{ item }">
            <span class="cost-text total-amount">{{ formatCurrency(item.totalAmount) }}</span>
          </template>

          <!-- Payment Method Column -->
          <template #item.paymentMethod="{ item }">
            <v-chip class="category-chip" size="small">
              {{ item.paymentMethod }}
            </v-chip>
          </template>

          <!-- Status Column -->
          <template #item.status="{ item }">
            <v-chip class="status-chip" size="small">
              {{ getStatusText(item.status) }}
            </v-chip>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="action-buttons">
              <v-btn
                size="small"
                color="primary"
                variant="text"
                @click="viewPurchaseDetails(item)"
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
                @click="editPurchase(item)"
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
                @click="deletePurchase(item)"
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

      <!-- Add/Edit Purchase Dialog -->
      <v-dialog v-model="purchaseDialog" max-width="800" scrollable persistent>
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-cart</v-icon>
                <span class="header-title">{{ isEditing ? 'تعديل عملية الشراء' : 'إضافة عملية شراء جديدة' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closePurchaseDialog"
                class="close-btn"
              />
            </div>
          </div>

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="purchaseForm" v-model="purchaseFormValid">
              <div class="form-fields">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="purchaseForm.supplierName"
                      label="اسم المورد"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم المورد مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="purchaseForm.purchaseDate"
                      label="تاريخ الشراء"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ الشراء مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="purchaseForm.itemName"
                      label="اسم العنصر"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم العنصر مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="purchaseForm.unit"
                      label="وحدة القياس"
                      variant="outlined"
                      :rules="[v => !!v || 'وحدة القياس مطلوبة']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="purchaseForm.quantity"
                      label="الكمية"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'الكمية مطلوبة', v => v > 0 || 'الكمية يجب أن تكون أكبر من صفر']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="purchaseForm.unitPrice"
                      label="سعر الوحدة (د.ع)"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'سعر الوحدة مطلوب', v => v > 0 || 'سعر الوحدة يجب أن يكون أكبر من صفر']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="purchaseForm.totalAmount"
                      label="المبلغ الإجمالي (د.ع)"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'المبلغ الإجمالي مطلوب', v => v > 0 || 'المبلغ الإجمالي يجب أن يكون أكبر من صفر']"
                      required
                      class="form-field"
                      readonly
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="purchaseForm.paymentMethod"
                      :items="paymentMethods"
                      label="طريقة الدفع"
                      variant="outlined"
                      :rules="[v => !!v || 'طريقة الدفع مطلوبة']"
                      required
                      class="form-field black-list"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="purchaseForm.status"
                      :items="statusOptions"
                      label="الحالة"
                      variant="outlined"
                      :rules="[v => !!v || 'الحالة مطلوبة']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12">
                    <v-textarea
                      v-model="purchaseForm.notes"
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
              @click="closePurchaseDialog"
              class="action-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="savePurchase"
              :disabled="!purchaseFormValid"
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
import { ref, computed, onMounted, watch } from 'vue'

// ========================================
// متغيرات الحالة الأساسية
// ========================================
const loading = ref(false)
const purchaseDialog = ref(false)
const purchaseFormValid = ref(false)
const isEditing = ref(false)
const selectedPurchase = ref(null)
const searchQuery = ref('')
const selectedStatus = ref('')
const selectedSupplier = ref('')

// ========================================
// بيانات النموذج
// ========================================
const purchaseForm = ref({
  supplierName: '',
  purchaseDate: '',
  itemName: '',
  unit: '',
  quantity: 0,
  unitPrice: 0,
  totalAmount: 0,
  paymentMethod: '',
  status: 'مكتمل',
  notes: ''
})

// ========================================
// بيانات المشتريات
// ========================================
const purchaseItems = ref([
  {
    id: 1,
    supplierName: 'مورد المواد الأولية',
    purchaseDate: '2024-01-15',
    itemName: 'أسمنت',
    unit: 'كيس',
    quantity: 100,
    unitPrice: 5000,
    totalAmount: 500000,
    paymentMethod: 'نقدي',
    status: 'مكتمل',
    notes: 'تم استلام الكمية كاملة'
  },
  {
    id: 2,
    supplierName: 'شركة الحديد',
    purchaseDate: '2024-01-20',
    itemName: 'حديد تسليح',
    unit: 'طن',
    quantity: 5,
    unitPrice: 1200000,
    totalAmount: 6000000,
    paymentMethod: 'شيك',
    status: 'معلق',
    notes: 'في انتظار التأكيد'
  },
  {
    id: 3,
    supplierName: 'مورد الخشب',
    purchaseDate: '2024-01-25',
    itemName: 'خشب',
    unit: 'متر مكعب',
    quantity: 10,
    unitPrice: 800000,
    totalAmount: 8000000,
    paymentMethod: 'تحويل بنكي',
    status: 'مكتمل',
    notes: ''
  }
])

// ========================================
// الخيارات
// ========================================
const statusOptions = ref([
  { title: 'مكتمل', value: 'مكتمل' },
  { title: 'معلق', value: 'معلق' },
  { title: 'ملغي', value: 'ملغي' }
])

const paymentMethods = ref([
  { title: 'نقدي', value: 'نقدي' },
  { title: 'شيك', value: 'شيك' },
  { title: 'تحويل بنكي', value: 'تحويل بنكي' },
  { title: 'آجل', value: 'آجل' }
])

const supplierOptions = computed(() => {
  const suppliers = [...new Set(purchaseItems.value.map(item => item.supplierName))]
  return suppliers.map(supplier => ({ title: supplier, value: supplier }))
})

// ========================================
// رؤوس الجدول
// ========================================
const headers = ref([
  { title: '#', key: 'serial', align: 'center', sortable: false },
  { title: 'تاريخ الشراء', key: 'purchaseDate', align: 'center' },
  { title: 'اسم المورد', key: 'supplierName', align: 'right' },
  { title: 'اسم العنصر', key: 'itemName', align: 'right' },
  { title: 'الكمية', key: 'quantity', align: 'center' },
  { title: 'سعر الوحدة', key: 'unitPrice', align: 'center' },
  { title: 'المبلغ الإجمالي', key: 'totalAmount', align: 'center' },
  { title: 'طريقة الدفع', key: 'paymentMethod', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
])

// ========================================
// Computed Properties
// ========================================
const totalPurchases = computed(() => purchaseItems.value.length)
const totalExpenses = computed(() => {
  return purchaseItems.value.reduce((sum, item) => sum + (item.totalAmount || 0), 0)
})
const totalSuppliers = computed(() => {
  return new Set(purchaseItems.value.map(item => item.supplierName)).size
})
const pendingPurchases = computed(() => {
  return purchaseItems.value.filter(item => item.status === 'معلق').length
})

// ========================================
// Methods
// ========================================
const formatCurrency = (amount) => {
  if (!amount) return '0 د.ع'
  const formatted = new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
  return formatted + ' د.ع'
}

const formatNumber = (number) => {
  if (!number) return '0'
  return new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(number)
}

const getStatusText = (status) => {
  return status
}

const searchPurchases = () => {
}

const openAddPurchaseDialog = () => {
  purchaseDialog.value = true
  isEditing.value = false
  selectedPurchase.value = null
  purchaseForm.value = {
    supplierName: '',
    purchaseDate: '',
    itemName: '',
    unit: '',
    quantity: 0,
    unitPrice: 0,
    totalAmount: 0,
    paymentMethod: '',
    status: 'مكتمل',
    notes: ''
  }
}

const closePurchaseDialog = () => {
  purchaseDialog.value = false
  isEditing.value = false
  selectedPurchase.value = null
  purchaseForm.value = {
    supplierName: '',
    purchaseDate: '',
    itemName: '',
    unit: '',
    quantity: 0,
    unitPrice: 0,
    totalAmount: 0,
    paymentMethod: '',
    status: 'مكتمل',
    notes: ''
  }
}

const editPurchase = (item) => {
  selectedPurchase.value = item
  isEditing.value = true
  purchaseForm.value = { ...item }
  purchaseDialog.value = true
}

const viewPurchaseDetails = (item) => {
  // يمكن إضافة نافذة عرض التفاصيل هنا
}

const deletePurchase = (item) => {
  if (confirm(`هل أنت متأكد من حذف عملية الشراء "${item.itemName}"؟`)) {
    const index = purchaseItems.value.findIndex(p => p.id === item.id)
    if (index > -1) {
      purchaseItems.value.splice(index, 1)
    }
  }
}

const savePurchase = () => {
  if (!purchaseFormValid.value) return

  // حساب المبلغ الإجمالي
  const total = parseFloat(purchaseForm.value.quantity) * parseFloat(purchaseForm.value.unitPrice)
  purchaseForm.value.totalAmount = total

  if (isEditing.value && selectedPurchase.value) {
    // تحديث المشتري الموجود
    const index = purchaseItems.value.findIndex(p => p.id === selectedPurchase.value.id)
    if (index > -1) {
      purchaseItems.value[index] = { ...purchaseForm.value, id: selectedPurchase.value.id }
    }
  } else {
    // إضافة مشتري جديد
    const newPurchase = {
      ...purchaseForm.value,
      id: Date.now()
    }
    purchaseItems.value.push(newPurchase)
  }

  closePurchaseDialog()
}

// حساب المبلغ الإجمالي تلقائياً
watch([() => purchaseForm.value.quantity, () => purchaseForm.value.unitPrice], ([quantity, unitPrice]) => {
  if (quantity && unitPrice) {
    purchaseForm.value.totalAmount = parseFloat(quantity) * parseFloat(unitPrice)
  } else {
    purchaseForm.value.totalAmount = 0
  }
})

onMounted(() => {
})
</script>


<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/purchases.css';
</style>
