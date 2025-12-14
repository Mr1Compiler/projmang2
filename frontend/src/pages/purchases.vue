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
            <h3 class="text-h3 font-weight-bold text-error mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; font-size: 1.1rem !important;">{{ formatCurrency(totalExpenses) || '0 د.ع' }}</h3>
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
  console.log('البحث عن:', searchQuery.value)
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
  console.log('عرض تفاصيل الشراء:', item)
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
  console.log('صفحة المشتريات جاهزة')
})
</script>

<style scoped>
/* ========================================
   تنسيق الحقول - نفس تنسيق صفحة المصروفات
   ======================================== */

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
  background: white !important;
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

.dialog-actions {
  padding: 20px 30px;
  background: #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* ========================================
   تنسيقات الصفحة الأساسية - نفس الصفحات الأخرى
   ======================================== */

/* Page Styles */
.data-page {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  padding: 20px;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
}

.purchases-page {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  padding: 20px;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
}

.centered-content {
  max-width: 1400px !important;
  margin: 0 auto !important;
  width: 100% !important;
  box-sizing: border-box !important;
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

.stats-row .v-col {
  padding: 12px;
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
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  border-radius: 16px !important;
  box-shadow: 
    0 8px 24px rgba(59, 130, 246, 0.3),
    0 4px 12px rgba(37, 99, 235, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  width: 100%;
  overflow-x: hidden;
  box-sizing: border-box;
  margin-bottom: 24px;
  position: relative;
  overflow: hidden;
}

.search-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  animation: shimmer 3s infinite;
  pointer-events: none;
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

/* ========================================
   تأثيرات بصرية - Visual Effects
   ======================================== */

.glass-effect {
  background: rgba(255, 255, 255, 0.1) !important;
  backdrop-filter: blur(20px) !important;
  -webkit-backdrop-filter: blur(20px) !important;
}

.gradient-animation {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  animation: gradient-shift 3s ease infinite;
  background-size: 200% 200%;
}

@keyframes gradient-shift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.text-glow {
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3), 0 0 30px rgba(255, 255, 255, 0.2);
}

.fade-in {
  animation: fadeIn 0.6s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.star-twinkle {
  animation: star-twinkle 2s ease-in-out infinite;
}

@keyframes star-twinkle {
  0%, 100% { 
    transform: scale(1) rotate(0deg); 
    opacity: 1;
  }
  50% { 
    transform: scale(1.1) rotate(180deg); 
    opacity: 0.8;
  }
}

.hover-lift {
  transition: all 0.3s ease;
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.card-glow {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.card-glow:hover {
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
}

.smooth-transition {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.icon-glow {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
  transition: all 0.3s ease;
}

.icon-glow:hover {
  filter: drop-shadow(0 6px 12px rgba(0, 0, 0, 0.3));
  transform: scale(1.1);
}

.full-width {
  width: 100% !important;
}

/* ========================================
   تنسيقات الجدول
   ======================================== */

.data-table-card {
  background: rgba(255, 249, 249, 0.95);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin-top: 24px;
}

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

.table-title {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  padding: 20px 24px !important;
  border-radius: 12px 12px 0 0 !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.3) !important;
  margin: 0 !important;
  font-weight: 700 !important;
  font-size: 1.5rem !important;
  text-align: right !important;
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
}

.indigo-title {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
}

.title-text {
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 1.5rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
}

.project-table {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  width: 100% !important;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__th {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.85rem !important;
  text-align: center !important;
  padding: 12px 10px !important;
  border-bottom: 3px solid #6d28d9 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 0.5px !important;
  border-radius: 0 !important;
  position: relative !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
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
  font-weight: 700 !important;
  font-size: 0.85rem !important;
  text-align: center !important;
  padding: 12px 10px !important;
  border-bottom: 3px solid #6d28d9 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 0.5px !important;
  position: relative !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table-header {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  display: table-header-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table-header th {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.85rem !important;
  text-align: center !important;
  padding: 12px 10px !important;
  border-bottom: 3px solid #6d28d9 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 0.5px !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__td {
  padding: 10px 8px !important;
  border-bottom: 1px solid #e5e7eb !important;
  text-align: center !important;
  background: #ffffff !important;
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  vertical-align: middle !important;
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__td *,
.project-table .v-data-table__td span,
.project-table .v-data-table__td div,
.project-table .v-data-table__td p {
  color: #000000 !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__wrapper {
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
}

.project-table .v-data-table__wrapper table {
  width: 100% !important;
  display: table !important;
  visibility: visible !important;
  opacity: 1 !important;
  border-collapse: collapse !important;
}

.project-table .v-data-table__wrapper tbody {
  display: table-row-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__wrapper tbody tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__wrapper tbody td {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
  padding: 10px 8px !important;
  border-bottom: 1px solid #e5e7eb !important;
  text-align: center !important;
  background: #ffffff !important;
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  vertical-align: middle !important;
}

/* إصلاح شامل لعرض الجدول */
.project-table .v-data-table__wrapper tbody td *,
.project-table .v-data-table__wrapper tbody td span,
.project-table .v-data-table__wrapper tbody td div {
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
  color: #000000 !important;
}

/* إصلاح عرض الأعمدة */
.project-table colgroup,
.project-table col {
  display: table-column !important;
  visibility: visible !important;
  opacity: 1 !important;
}

/* إصلاح عرض الجدول بالكامل */
.project-table table {
  table-layout: auto !important;
  width: 100% !important;
  border-collapse: collapse !important;
  border-spacing: 0 !important;
}

.project-table thead {
  display: table-header-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table tbody {
  display: table-row-group !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table tr {
  display: table-row !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table th,
.project-table td {
  display: table-cell !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(even) {
  background: linear-gradient(135deg, #faf5ff 0%, #f3e8ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(odd) {
  background: linear-gradient(135deg, #ffffff 0%, #faf5ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(even) td {
  background: linear-gradient(135deg, #faf5ff 0%, #f3e8ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(odd) td {
  background: linear-gradient(135deg, #ffffff 0%, #faf5ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:hover {
  background: linear-gradient(135deg, #f3e8ff 0%, #e9d5ff 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.15) !important;
  transition: all 0.3s ease !important;
}

.project-table .v-data-table__wrapper tbody tr:hover td {
  background: linear-gradient(135deg, #f3e8ff 0%, #e9d5ff 100%) !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

/* تنسيق النصوص داخل الخلايا - جميعها أسود */
.project-table .serial-number {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 0.8rem !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 4px 10px !important;
  border-radius: 8px !important;
  display: inline-block !important;
  min-width: 35px !important;
  text-align: center !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.project-table .project-name {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  text-align: right !important;
  padding: 4px 8px !important;
  border-radius: 6px !important;
  display: inline-block !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.project-table .date-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.project-table .cost-text {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 0.8rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 6px 12px !important;
  border-radius: 8px !important;
  display: inline-block !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #d1d5db !important;
  white-space: nowrap !important;
  min-width: 90px !important;
  text-align: center !important;
}

.project-table .quantity-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  background: linear-gradient(135deg, #fef3c7 0%, #fde68a 100%) !important;
  padding: 8px 14px !important;
  border-radius: 8px !important;
  display: inline-block !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #fbbf24 !important;
  white-space: nowrap !important;
  min-width: 80px !important;
  text-align: center !important;
}

.project-table .total-amount {
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%) !important;
  border: 1px solid #3b82f6 !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  min-width: 120px !important;
  text-align: center !important;
}

/* تنسيق الشريحات - Chips - نص أسود */
.project-table .v-chip {
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  padding: 4px 10px !important;
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

/* ========================================
   تنسيق القياسات والأرقام
   ======================================== */

/* تنسيق الأرقام في البطاقات الإحصائية */
.stat-number {
  font-size: 1.1rem !important;
  line-height: 1.2 !important;
  letter-spacing: -0.5px !important;
  word-break: break-word !important;
  overflow-wrap: break-word !important;
  font-weight: 800 !important;
  text-align: center !important;
  width: 100% !important;
}

.stat-currency {
  font-size: 1.1rem !important;
  line-height: 1.3 !important;
  word-break: break-word !important;
  overflow-wrap: break-word !important;
  font-weight: 800 !important;
  white-space: normal !important;
  text-align: center !important;
}

.stat-card .text-h4 {
  font-size: 1.75rem !important;
  line-height: 1.3 !important;
  word-break: break-word !important;
  overflow-wrap: break-word !important;
  font-weight: 800 !important;
}

.stat-card .text-h5 {
  font-size: 1.5rem !important;
  line-height: 1.3 !important;
  word-break: break-word !important;
  overflow-wrap: break-word !important;
}

/* تحسين عرض الأرقام الكبيرة */
.stat-card h3 {
  min-height: 70px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  flex-wrap: wrap !important;
  text-align: center !important;
  padding: 8px 4px !important;
  width: 100% !important;
  margin: 0 auto !important;
}

/* تحسين عرض البطاقات الإحصائية */
.stat-card {
  min-height: 180px !important;
  display: flex !important;
  flex-direction: column !important;
  justify-content: center !important;
  align-items: center !important;
  text-align: center !important;
  width: 100% !important;
  margin: 0 auto !important;
}

.stat-card .stat-icon {
  flex-shrink: 0 !important;
  margin: 0 auto !important;
}

.stat-card p {
  text-align: center !important;
  width: 100% !important;
  margin: 0 auto !important;
}

/* تحسين عرض الجدول */
.data-table-card {
  max-width: 100% !important;
  margin: 0 auto !important;
  width: 100% !important;
}

.project-table {
  width: 100% !important;
  max-width: 100% !important;
  margin: 0 auto !important;
  overflow-x: auto !important;
}

/* تحسين عرض شريط البحث */
.search-card {
  max-width: 100% !important;
  margin: 0 auto 24px auto !important;
  width: 100% !important;
}

.search-card .v-card-text {
  padding: 20px !important;
}

.search-card .v-row {
  margin: 0 !important;
  align-items: center !important;
}

.search-card .v-col {
  padding: 8px !important;
}

/* تنسيق حقول البحث والفلترة */
.search-card .search-field :deep(.v-field) {
  background: rgba(255, 255, 255, 1) !important;
  border-radius: 10px !important;
}

.search-card .search-field :deep(.v-field__input) {
  color: #000000 !important;
  background: rgba(255, 255, 255, 1) !important;
  font-weight: 600 !important;
  padding: 12px 16px !important;
}

.search-card .search-field :deep(.v-field__input input) {
  color: #000000 !important;
  font-weight: 600 !important;
}

.search-card .search-field :deep(.v-label) {
  color: #000000 !important;
  font-weight: 600 !important;
  background: rgba(255, 255, 255, 1) !important;
  padding: 0 8px !important;
  margin: 0 8px !important;
}

.search-card .search-field :deep(.v-field--focused .v-label) {
  color: #000000 !important;
  background: rgba(255, 255, 255, 1) !important;
}

.search-card .search-field :deep(.v-field__outline) {
  border-color: rgba(255, 255, 255, 0.8) !important;
  border-width: 2px !important;
}

.search-card .search-field :deep(.v-field--focused .v-field__outline) {
  border-color: rgba(255, 255, 255, 1) !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.2) !important;
}

.search-card .search-field :deep(.v-icon) {
  color: #2563eb !important;
}

.search-card .v-select :deep(.v-field) {
  background: rgba(255, 255, 255, 1) !important;
  border-radius: 10px !important;
}

.search-card .v-select :deep(.v-field__input) {
  color: #000000 !important;
  background: rgba(255, 255, 255, 1) !important;
  font-weight: 600 !important;
  padding: 12px 16px !important;
}

.search-card .v-select :deep(.v-field__input input) {
  color: #000000 !important;
  font-weight: 600 !important;
}

.search-card .v-select :deep(.v-label) {
  color: #000000 !important;
  font-weight: 600 !important;
  background: rgba(255, 255, 255, 1) !important;
  padding: 0 8px !important;
  margin: 0 8px !important;
}

.search-card .v-select :deep(.v-field--focused .v-label) {
  color: #000000 !important;
  background: rgba(255, 255, 255, 1) !important;
}

.search-card .v-select :deep(.v-field__outline) {
  border-color: rgba(255, 255, 255, 0.8) !important;
  border-width: 2px !important;
}

.search-card .v-select :deep(.v-field--focused .v-field__outline) {
  border-color: rgba(255, 255, 255, 1) !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.2) !important;
}

.search-card .v-select :deep(.v-icon) {
  color: #2563eb !important;
}

/* تنسيق زر البحث */
.search-card .search-btn,
.search-card .search-btn.v-btn {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 50%, #3b82f6 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 10px 20px !important;
  font-size: 0.75rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  box-shadow: 
    0 4px 16px rgba(30, 64, 175, 0.4),
    0 2px 8px rgba(37, 99, 235, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
  position: relative;
  overflow: hidden;
}

.search-card .search-btn :deep(.v-btn__content) {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  letter-spacing: 0.3px !important;
}

.search-card .search-btn::before {
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

.search-card .search-btn:hover::before {
  left: 100%;
}

.search-card .search-btn:hover {
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #2563eb 100%) !important;
  box-shadow: 
    0 8px 24px rgba(30, 64, 175, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  transform: translateY(-2px) scale(1.02) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

/* تنسيق زر إضافة عملية شراء */
.search-card .add-expense-btn,
.search-card .add-expense-btn.v-btn {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 10px 20px !important;
  font-size: 0.75rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  box-shadow: 
    0 4px 16px rgba(16, 185, 129, 0.4),
    0 2px 8px rgba(5, 150, 105, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
  position: relative;
  overflow: hidden;
}

.search-card .add-expense-btn :deep(.v-btn__content) {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  letter-spacing: 0.3px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 8px !important;
}

.search-card .add-expense-btn :deep(.v-icon) {
  color: #ffffff !important;
  font-size: 18px !important;
  transition: transform 0.3s ease !important;
}

.search-card .add-expense-btn::before {
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

.search-card .add-expense-btn:hover::before {
  left: 100%;
}

.search-card .add-expense-btn:hover {
  background: linear-gradient(135deg, #047857 0%, #059669 50%, #10b981 100%) !important;
  box-shadow: 
    0 8px 24px rgba(16, 185, 129, 0.5),
    0 4px 12px rgba(5, 150, 105, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  transform: translateY(-2px) scale(1.02) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.search-card .add-expense-btn:hover :deep(.v-icon) {
  transform: rotate(90deg) scale(1.1) !important;
}

/* Responsive Design */
@media (max-width: 1200px) {
  .centered-content,
  .fullscreen-content {
    max-width: 100% !important;
    padding: 0 15px !important;
  }

  .stats-row .v-col {
    padding: 10px !important;
  }
}

@media (max-width: 768px) {
  .purchases-page {
    padding: 10px !important;
  }

  .centered-content,
  .fullscreen-content {
    max-width: 100% !important;
    padding: 0 10px !important;
  }

  .stats-row {
    gap: 12px !important;
  }

  .stats-row .v-col {
    padding: 8px !important;
  }

  .project-table .v-data-table__wrapper {
    overflow-x: auto !important;
  }
  
  .project-table .v-data-table__wrapper table {
    min-width: 900px !important;
  }

  .stat-number {
    font-size: 1rem !important;
  }

  .stat-currency {
    font-size: 1rem !important;
  }

  .stat-card .text-h4 {
    font-size: 1.4rem !important;
  }

  .stat-card .text-h5 {
    font-size: 1.2rem !important;
  }

  .stat-card h3 {
    min-height: 60px !important;
    padding: 4px 2px !important;
  }

  .stat-card {
    min-height: 160px !important;
    padding: 16px !important;
  }

  .project-table .cost-text,
  .project-table .quantity-text,
  .project-table .total-amount {
    font-size: 0.85rem !important;
    padding: 6px 10px !important;
    min-width: auto !important;
  }

  .search-card .v-card-text {
    padding: 12px !important;
  }

  .data-table-card {
    margin-top: 16px !important;
  }
}

@media (max-width: 480px) {
  .centered-content,
  .fullscreen-content {
    padding: 0 8px !important;
  }

  .stat-number {
    font-size: 1rem !important;
  }

  .stat-currency {
    font-size: 1rem !important;
  }

  .stat-card {
    min-height: 140px !important;
    padding: 12px !important;
  }

  .stat-card .stat-icon {
    margin-bottom: 8px !important;
  }

  .stat-card .stat-icon .v-icon {
    font-size: 48px !important;
  }
}
</style>

