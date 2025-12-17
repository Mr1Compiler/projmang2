<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <PageHeader
        title="إدارة المبيعات"
        subtitle="إدارة وتتبع جميع عمليات البيع والمبيعات"
        icon="💰"
      />

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="success">mdi-cash-multiple</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number-ltr">{{ totalSales || 0 }}</h3>
            <p class="text-subtitle-1 text-success mb-0">إجمالي المبيعات</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="primary">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h6 font-weight-bold text-primary mb-2 stat-number-ltr-nowrap">{{ formatCurrency(totalRevenue) || '0 د.ع' }}</h3>
            <p class="text-subtitle-1 text-primary mb-0">إجمالي الإيرادات</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="info">mdi-account-group</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-info mb-2 stat-number-ltr">{{ totalCustomers || 0 }}</h3>
            <p class="text-subtitle-1 text-info mb-0">إجمالي العملاء</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3">
          <v-card class="stat-card pa-6 pb-8 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number-ltr">{{ pendingSales || 0 }}</h3>
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
                class="search-field bg-filter-input"
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
import { formatCurrency } from '@/utils/formatters'
import { PageHeader } from '@/components/shared'

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
const getStatusText = (status) => {
  return status || 'غير محدد'
}

onMounted(() => {
  // تهيئة الصفحة
})
</script>


<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/sales.css';
</style>
