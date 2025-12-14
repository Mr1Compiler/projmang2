<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">💰</span>
        <h1 class="page-title text-glow fade-in">إدارة المبيعات</h1>
        <p class="page-subtitle fade-in">إدارة وتتبع جميع عمليات البيع والمبيعات</p>
      </div>

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="success">mdi-cash-multiple</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalSales || 0 }}</h3>
            <p class="text-subtitle-1 text-success mb-0">إجمالي المبيعات</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="primary">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h6 font-weight-bold text-primary mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; white-space: nowrap;">{{ formatCurrency(totalRevenue) || '0 د.ع' }}</h3>
            <p class="text-subtitle-1 text-primary mb-0">إجمالي الإيرادات</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="info">mdi-account-group</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalCustomers || 0 }}</h3>
            <p class="text-subtitle-1 text-info mb-0">إجمالي العملاء</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ pendingSales || 0 }}</h3>
            <p class="text-subtitle-1 text-warning mb-0">مبيعات معلقة</p>
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
                label="البحث في المبيعات..."
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
                v-model="selectedPaymentMethod"
                :items="paymentMethodOptions"
                label="طريقة الدفع"
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
                @click="searchSales"
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
                @click="openAddSaleDialog"
              >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة عملية بيع
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Sales Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title">
          <span class="title-text">عمليات المبيعات</span>
        </v-card-title>

        <v-data-table
          :headers="headers"
          :items="salesItems"
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

          <!-- Sale Date Column -->
          <template #item.saleDate="{ item }">
            <span class="date-text">{{ item.saleDate }}</span>
          </template>

          <!-- Customer Name Column -->
          <template #item.customerName="{ item }">
            <span class="project-name">{{ item.customerName }}</span>
          </template>

          <!-- Item Name Column -->
          <template #item.itemName="{ item }">
            <span class="project-name">{{ item.itemName }}</span>
          </template>

          <!-- Quantity Column -->
          <template #item.quantity="{ item }">
            <span class="cost-text">{{ item.quantity }} {{ item.unit }}</span>
          </template>

          <!-- Unit Price Column -->
          <template #item.unitPrice="{ item }">
            <span class="cost-text">{{ formatCurrency(item.unitPrice) }}</span>
          </template>

          <!-- Total Amount Column -->
          <template #item.totalAmount="{ item }">
            <span class="cost-text">{{ formatCurrency(item.totalAmount) }}</span>
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
                @click="viewSaleDetails(item)"
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
                @click="editSale(item)"
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
                @click="deleteSale(item)"
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

      <!-- Add/Edit Sale Dialog -->
      <v-dialog v-model="saleDialog" max-width="800" scrollable persistent>
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-cash-multiple</v-icon>
                <span class="header-title">{{ isEditing ? 'تعديل عملية البيع' : 'إضافة عملية بيع جديدة' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeSaleDialog"
                class="close-btn"
              />
            </div>
          </div>

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="saleForm" v-model="saleFormValid">
              <div class="form-fields">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="saleForm.customerName"
                      label="اسم العميل"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم العميل مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="saleForm.saleDate"
                      label="تاريخ البيع"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ البيع مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="saleForm.itemName"
                      label="اسم العنصر"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم العنصر مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="saleForm.unit"
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
                      v-model="saleForm.quantity"
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
                      v-model="saleForm.unitPrice"
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
                      v-model="saleForm.totalAmount"
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
                      v-model="saleForm.paymentMethod"
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
                      v-model="saleForm.status"
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
                      v-model="saleForm.notes"
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
              @click="closeSaleDialog"
              class="action-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveSale"
              :disabled="!saleFormValid"
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

// متغيرات الحالة الأساسية
const loading = ref(false)
const saleDialog = ref(false)
const saleFormValid = ref(false)
const isEditing = ref(false)
const searchQuery = ref('')
const selectedStatus = ref('')
const selectedPaymentMethod = ref('')
const selectedSale = ref(null)

// عناوين الجدول
const headers = ref([
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'تاريخ البيع', key: 'saleDate', sortable: true, align: 'center' },
  { title: 'اسم العميل', key: 'customerName', sortable: true, align: 'right' },
  { title: 'اسم العنصر', key: 'itemName', sortable: true, align: 'right' },
  { title: 'الكمية', key: 'quantity', sortable: true, align: 'center' },
  { title: 'سعر الوحدة', key: 'unitPrice', sortable: true, align: 'center' },
  { title: 'المبلغ الإجمالي', key: 'totalAmount', sortable: true, align: 'center' },
  { title: 'طريقة الدفع', key: 'paymentMethod', sortable: true, align: 'center' },
  { title: 'الحالة', key: 'status', sortable: true, align: 'center' },
  { title: 'الإجراءات', key: 'actions', sortable: false, align: 'center' }
])

// نموذج عملية البيع
const saleForm = ref({
  customerName: '',
  saleDate: '',
  itemName: '',
  unit: '',
  quantity: 0,
  unitPrice: 0,
  totalAmount: 0,
  paymentMethod: '',
  status: 'مكتمل',
  notes: ''
})

// بيانات المبيعات التجريبية
const salesItems = ref([
  {
    id: 1,
    customerName: 'أحمد محمد',
    saleDate: '2024-01-15',
    itemName: 'أسمنت بورتلاند',
    quantity: 50,
    unit: 'كيس',
    unitPrice: 5000,
    totalAmount: 250000,
    paymentMethod: 'نقدي',
    status: 'مكتمل',
    notes: ''
  },
  {
    id: 2,
    customerName: 'شركة البناء الحديث',
    saleDate: '2024-01-16',
    itemName: 'حديد تسليح',
    quantity: 5,
    unit: 'طن',
    unitPrice: 800000,
    totalAmount: 4000000,
    paymentMethod: 'تحويل بنكي',
    status: 'مكتمل',
    notes: 'دفعة أولى'
  },
  {
    id: 3,
    customerName: 'مؤسسة التطوير',
    saleDate: '2024-01-17',
    itemName: 'طلاء أبيض',
    quantity: 20,
    unit: 'علبة',
    unitPrice: 15000,
    totalAmount: 300000,
    paymentMethod: 'شيك',
    status: 'معلق',
    notes: 'في انتظار التحقق'
  }
])

// طرق الدفع
const paymentMethods = [
  'نقدي',
  'تحويل بنكي',
  'شيك',
  'بطاقة ائتمان',
  'آجل'
]

const paymentMethodOptions = ref([
  'جميع الطرق',
  ...paymentMethods
])

const statusOptions = ref([
  'مكتمل',
  'معلق',
  'ملغي',
  'مسترد'
])

// إحصائيات المبيعات
const totalSales = computed(() => salesItems.value.length)
const totalRevenue = computed(() => {
  return salesItems.value.reduce((sum, sale) => sum + sale.totalAmount, 0)
})
const totalCustomers = computed(() => {
  const uniqueCustomers = new Set(salesItems.value.map(sale => sale.customerName))
  return uniqueCustomers.size
})
const pendingSales = computed(() => {
  return salesItems.value.filter(sale => sale.status === 'معلق').length
})

// حساب المبلغ الإجمالي تلقائياً
watch([() => saleForm.value.quantity, () => saleForm.value.unitPrice], () => {
  if (saleForm.value.quantity && saleForm.value.unitPrice) {
    saleForm.value.totalAmount = parseFloat(saleForm.value.quantity) * parseFloat(saleForm.value.unitPrice)
  } else {
    saleForm.value.totalAmount = 0
  }
})

// دوال المبيعات
const searchSales = () => {
  console.log('البحث عن:', searchQuery.value)
}

const openAddSaleDialog = () => {
  saleDialog.value = true
  isEditing.value = false
  selectedSale.value = null
  const today = new Date().toISOString().split('T')[0]
  saleForm.value = {
    customerName: '',
    saleDate: today,
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

const closeSaleDialog = () => {
  saleDialog.value = false
  isEditing.value = false
  selectedSale.value = null
  saleForm.value = {
    customerName: '',
    saleDate: '',
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

const editSale = (sale) => {
  selectedSale.value = sale
  isEditing.value = true
  saleForm.value = { ...sale }
  saleDialog.value = true
}

const viewSaleDetails = (sale) => {
  console.log('عرض تفاصيل عملية البيع:', sale)
  // يمكن إضافة نافذة عرض التفاصيل هنا
}

const deleteSale = (sale) => {
  if (confirm(`هل أنت متأكد من حذف عملية البيع للعميل "${sale.customerName}"؟`)) {
    const index = salesItems.value.findIndex(s => s.id === sale.id)
    if (index > -1) {
      salesItems.value.splice(index, 1)
    }
  }
}

const saveSale = () => {
  if (isEditing.value) {
    // تحديث عملية البيع
    const index = salesItems.value.findIndex(s => s.id === selectedSale.value.id)
    if (index > -1) {
      salesItems.value[index] = {
        ...saleForm.value,
        id: selectedSale.value.id,
        quantity: parseFloat(saleForm.value.quantity),
        unitPrice: parseFloat(saleForm.value.unitPrice),
        totalAmount: parseFloat(saleForm.value.totalAmount)
      }
    }
  } else {
    // إضافة عملية بيع جديدة
    const newSale = {
      ...saleForm.value,
      id: Date.now(),
      quantity: parseFloat(saleForm.value.quantity),
      unitPrice: parseFloat(saleForm.value.unitPrice),
      totalAmount: parseFloat(saleForm.value.totalAmount)
    }
    salesItems.value.push(newSale)
  }
  closeSaleDialog()
}

// دوال مساعدة
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0
  }).format(amount).replace('IQD', 'د.ع')
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

.search-card .v-card-text {
  padding: 16px !important;
  width: 100%;
  box-sizing: border-box;
}

/* زر البحث */
.search-card .search-btn,
.search-card .search-btn.v-btn,
.search-card .v-btn.search-btn {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 50%, #3b82f6 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  padding: 12px 20px !important;
  border-radius: 12px !important;
  box-shadow: 
    0 4px 16px rgba(30, 64, 175, 0.3),
    0 2px 8px rgba(37, 99, 235, 0.2),
    0 0 20px rgba(59, 130, 246, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  text-transform: none !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
}

.search-card .search-btn :deep(.v-btn__content) {
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  text-align: center !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
}

.search-card .search-btn:hover {
  transform: translateY(-2px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(30, 64, 175, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    0 0 40px rgba(59, 130, 246, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #2563eb 100%) !important;
}

/* زر إضافة عملية بيع */
.search-card .add-expense-btn,
.search-card .add-expense-btn.v-btn,
.search-card .v-btn.add-expense-btn {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  padding: 12px 20px !important;
  border-radius: 12px !important;
  box-shadow: 
    0 4px 16px rgba(16, 185, 129, 0.3),
    0 2px 8px rgba(5, 150, 105, 0.2),
    0 0 20px rgba(16, 185, 129, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  text-transform: none !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
}

.search-card .add-expense-btn:hover {
  transform: translateY(-2px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(16, 185, 129, 0.5),
    0 4px 12px rgba(5, 150, 105, 0.4),
    0 0 40px rgba(16, 185, 129, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  background: linear-gradient(135deg, #059669 0%, #047857 50%, #065f46 100%) !important;
}

.search-card .add-expense-btn :deep(.v-btn__content) {
  color: white !important;
  gap: 6px !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
  text-align: center !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
}

.search-card .add-expense-btn :deep(.v-icon) {
  color: white !important;
  font-size: 18px !important;
  margin: 0 !important;
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
   تنسيق جدول المبيعات - Sales Table
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

.project-table .date-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
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


/* قواعد قوية للتسميات - Vuetify 3 - أسود تماماً */
.v-field__label,
.v-label,
.v-field__label *,
.v-label * {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
  text-shadow: none !important;
  -webkit-text-stroke: 0 !important;
  -webkit-font-smoothing: antialiased !important;
}

.v-field--focused .v-field__label,
.v-field--focused .v-label,
.v-field--active .v-field__label,
.v-field--active .v-label,
.v-field--dirty .v-field__label,
.v-field--dirty .v-label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
  text-shadow: none !important;
}

/* قواعد شاملة لجميع التسميات - أسود تماماً */
* .v-label,
* .v-field__label,
* label,
.v-input label,
.v-text-field label,
.v-textarea label,
.v-select label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}

.image-style-dialog * .v-label,
.image-style-dialog * .v-field__label,
.image-style-dialog * label,
.image-style-dialog .v-input label,
.image-style-dialog .v-text-field label,
.image-style-dialog .v-textarea label,
.image-style-dialog .v-select label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
  text-shadow: none !important;
}

/* قواعد قوية جداً للتسميات */
.v-input__details,
.v-messages,
.v-input__details * {
  color: #000000 !important;
}

.image-style-dialog .v-input__details,
.image-style-dialog .v-messages {
  color: #000000 !important;
}

/* إجبار اللون الأسود على جميع التسميات */
.image-style-dialog .v-field__label,
.image-style-dialog .v-label,
.image-style-dialog .v-field__label *,
.image-style-dialog .v-label * {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  text-shadow: none !important;
  -webkit-text-stroke: 0 !important;
}

@keyframes star-twinkle {
  0%, 100% { transform: scale(1) rotate(0deg); }
  50% { transform: scale(1.1) rotate(180deg); }
}

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
}
</style>

