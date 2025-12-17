<template>
  <v-dialog v-model="dialogModel" max-width="800px" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-info) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-credit-card</v-icon>
            <span class="header-title">الديون والتسديدات - {{ debtor?.name }}</span>
          </div>
          <v-btn
            icon="mdi-close"
            variant="text"
            size="small"
            color="white"
            @click="closeDialog"
            class="close-btn"
          />
        </div>
      </div>

      <!-- Dialog Body -->
      <div v-if="debtor" class="dialog-body">
        <!-- Quick Stats -->
        <v-row class="mb-4">
          <v-col cols="12" md="4">
            <v-card color="error" variant="tonal" class="pa-3">
              <div class="text-center">
                <v-icon size="32" color="error" class="mb-2">mdi-currency-usd</v-icon>
                <h3 class="text-h5 font-weight-bold">{{ formatCurrency(debtor.amount) }}</h3>
                <p class="text-subtitle-2 mb-0">إجمالي المديونية</p>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" md="4">
            <v-card color="success" variant="tonal" class="pa-3">
              <div class="text-center">
                <v-icon size="32" color="success" class="mb-2">mdi-check-circle</v-icon>
                <h3 class="text-h5 font-weight-bold">{{ formatCurrency(totalPaid) }}</h3>
                <p class="text-subtitle-2 mb-0">إجمالي المدفوع</p>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" md="4">
            <v-card color="warning" variant="tonal" class="pa-3">
              <div class="text-center">
                <v-icon size="32" color="warning" class="mb-2">mdi-clock-alert</v-icon>
                <h3 class="text-h5 font-weight-bold">{{ formatCurrency(remainingAmount) }}</h3>
                <p class="text-subtitle-2 mb-0">المبلغ المتبقي</p>
              </div>
            </v-card>
          </v-col>
        </v-row>

        <!-- Tabs -->
        <v-tabs v-model="activeTab" class="mb-4">
          <v-tab value="debts">الديون</v-tab>
          <v-tab value="payments">التسديدات</v-tab>
          <v-tab value="summary">الملخص</v-tab>
        </v-tabs>

        <v-window v-model="activeTab">
          <!-- Debts Tab -->
          <v-window-item value="debts">
            <v-card variant="outlined">
              <v-card-title class="d-flex align-center justify-space-between">
                <span>قائمة الديون</span>
                <v-btn
                  v-if="canWrite"
                  color="primary"
                  size="small"
                  prepend-icon="mdi-plus"
                  @click="$emit('add-debt')"
                >
                  إضافة دين
                </v-btn>
              </v-card-title>
              <v-data-table
                :headers="debtHeaders"
                :items="debtor.debts || []"
                class="elevation-0"
                no-data-text="لا توجد ديون"
              >
                <template v-slot:item.amount="{ item }">
                  <span class="font-weight-bold text-error">{{ formatCurrency(item.amount) }}</span>
                </template>
                <template v-slot:item.date="{ item }">
                  {{ formatDate(item.date) }}
                </template>
                <template v-slot:item.status="{ item }">
                  <v-chip
                    :color="item.status === 'paid' ? 'success' : 'warning'"
                    size="small"
                    variant="tonal"
                  >
                    {{ item.status === 'paid' ? 'مدفوع' : 'غير مدفوع' }}
                  </v-chip>
                </template>
                <template v-slot:item.actions="{ item }">
                  <v-btn
                    v-if="canWrite"
                    icon="mdi-pencil"
                    size="small"
                    variant="text"
                    @click="$emit('edit-debt', item)"
                  />
                  <v-btn
                    v-if="canDelete"
                    icon="mdi-delete"
                    size="small"
                    variant="text"
                    color="error"
                    @click="$emit('delete-debt', item)"
                  />
                </template>
              </v-data-table>
            </v-card>
          </v-window-item>

          <!-- Payments Tab -->
          <v-window-item value="payments">
            <v-card variant="outlined">
              <v-card-title class="d-flex align-center justify-space-between">
                <span>قائمة التسديدات</span>
                <v-btn
                  v-if="canWrite"
                  color="success"
                  size="small"
                  prepend-icon="mdi-plus"
                  @click="$emit('add-payment')"
                >
                  إضافة تسديد
                </v-btn>
              </v-card-title>
              <v-data-table
                :headers="paymentHeaders"
                :items="debtor.payments || []"
                class="elevation-0"
                no-data-text="لا توجد تسديدات"
              >
                <template v-slot:item.amount="{ item }">
                  <span class="font-weight-bold text-success">{{ formatCurrency(item.amount) }}</span>
                </template>
                <template v-slot:item.date="{ item }">
                  {{ formatDate(item.date) }}
                </template>
                <template v-slot:item.method="{ item }">
                  <v-chip
                    :color="getPaymentMethodColor(item.method)"
                    size="small"
                    variant="tonal"
                  >
                    {{ getPaymentMethodText(item.method) }}
                  </v-chip>
                </template>
                <template v-slot:item.actions="{ item }">
                  <v-btn
                    v-if="canWrite"
                    icon="mdi-pencil"
                    size="small"
                    variant="text"
                    @click="$emit('edit-payment', item)"
                  />
                  <v-btn
                    v-if="canDelete"
                    icon="mdi-delete"
                    size="small"
                    variant="text"
                    color="error"
                    @click="$emit('delete-payment', item)"
                  />
                </template>
              </v-data-table>
            </v-card>
          </v-window-item>

          <!-- Summary Tab -->
          <v-window-item value="summary">
            <v-card variant="outlined">
              <v-card-title>ملخص الديون والتسديدات</v-card-title>
              <v-card-text>
                <v-row>
                  <v-col cols="12" md="6">
                    <h4 class="text-h6 font-weight-bold mb-3">تفاصيل المديونية</h4>
                    <v-list>
                      <v-list-item>
                        <v-list-item-title>إجمالي المديونية</v-list-item-title>
                        <template v-slot:append>
                          <span class="font-weight-bold text-error">{{ formatCurrency(debtor.amount) }}</span>
                        </template>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>إجمالي المدفوع</v-list-item-title>
                        <template v-slot:append>
                          <span class="font-weight-bold text-success">{{ formatCurrency(totalPaid) }}</span>
                        </template>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>المبلغ المتبقي</v-list-item-title>
                        <template v-slot:append>
                          <span class="font-weight-bold text-warning">{{ formatCurrency(remainingAmount) }}</span>
                        </template>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>نسبة السداد</v-list-item-title>
                        <template v-slot:append>
                          <span class="font-weight-bold">{{ paymentPercentage }}%</span>
                        </template>
                      </v-list-item>
                    </v-list>
                  </v-col>
                  <v-col cols="12" md="6">
                    <h4 class="text-h6 font-weight-bold mb-3">آخر التسديدات</h4>
                    <v-list v-if="debtor.payments && debtor.payments.length > 0">
                      <v-list-item
                        v-for="payment in debtor.payments.slice(0, 3)"
                        :key="payment.id"
                      >
                        <template v-slot:prepend>
                          <v-icon color="success">mdi-check-circle</v-icon>
                        </template>
                        <v-list-item-title>{{ formatCurrency(payment.amount) }}</v-list-item-title>
                        <v-list-item-subtitle>{{ formatDate(payment.date) }}</v-list-item-subtitle>
                      </v-list-item>
                    </v-list>
                    <p v-else class="text-grey-darken-1">لا توجد تسديدات</p>
                  </v-col>
                </v-row>
              </v-card-text>
            </v-card>
          </v-window-item>
        </v-window>
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn
          color="primary"
          variant="elevated"
          @click="closeDialog"
          class="save-btn"
        >
          إغلاق
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'
import { formatCurrency, formatDate } from '@/utils/formatters'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  debtor: {
    type: Object,
    default: null
  },
  canWrite: {
    type: Boolean,
    default: true
  },
  canDelete: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits([
  'update:modelValue',
  'close',
  'add-debt',
  'edit-debt',
  'delete-debt',
  'add-payment',
  'edit-payment',
  'delete-payment'
])

const activeTab = ref('debts')

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const debtHeaders = [
  { title: 'المبلغ', key: 'amount', sortable: true },
  { title: 'التاريخ', key: 'date', sortable: true },
  { title: 'الوصف', key: 'description', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

const paymentHeaders = [
  { title: 'المبلغ', key: 'amount', sortable: true },
  { title: 'التاريخ', key: 'date', sortable: true },
  { title: 'طريقة الدفع', key: 'method', sortable: true },
  { title: 'الوصف', key: 'description', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

const totalPaid = computed(() => {
  if (!props.debtor?.payments) return 0
  return props.debtor.payments.reduce((sum, payment) => sum + payment.amount, 0)
})

const remainingAmount = computed(() => {
  if (!props.debtor) return 0
  return props.debtor.amount - totalPaid.value
})

const paymentPercentage = computed(() => {
  if (!props.debtor || props.debtor.amount === 0) return 0
  return Math.round((totalPaid.value / props.debtor.amount) * 100)
})

const closeDialog = () => {
  dialogModel.value = false
  emit('close')
}

const getPaymentMethodColor = (method) => {
  const colors = {
    'bank_transfer': 'primary',
    'credit_card': 'success',
    'cash': 'warning',
    'check': 'info'
  }
  return colors[method] || 'grey'
}

const getPaymentMethodText = (method) => {
  const texts = {
    'bank_transfer': 'تحويل بنكي',
    'credit_card': 'بطاقة ائتمان',
    'cash': 'نقدي',
    'check': 'شيك'
  }
  return texts[method] || 'غير محدد'
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
