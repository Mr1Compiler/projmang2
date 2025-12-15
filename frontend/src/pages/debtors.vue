<template>
  <div>
      <v-container fluid class="pa-6">
        <!-- رأس الصفحة المحسن -->
        <div class="page-header glass-effect gradient-animation">
          <div class="header-top-content">
            <h1 class="page-title">إدارة المديونون</h1>
            <span class="page-icon">💳</span>
          </div>
          <p class="page-subtitle">إدارة حسابات المديونون والمستحقات المالية</p>
        </div>


        <!-- إحصائيات سريعة محسنة -->
        <v-row class="mb-8">
          <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center" elevation="2">
              <div class="stat-icon mb-3">
                <v-icon size="64" color="info">mdi-account-group</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ debtors.length }}</h3>
              <p class="text-subtitle-1 text-info mb-0">إجمالي المديونون</p>
            </v-card>
          </v-col>
          <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center" elevation="2">
              <div class="stat-icon mb-3">
                <v-icon size="64" color="error">mdi-currency-usd</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-error mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatCurrency(totalDebt) }}</h3>
              <p class="text-subtitle-1 text-error mb-0">إجمالي المديونية</p>
            </v-card>
          </v-col>
          <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center" elevation="2">
              <div class="stat-icon mb-3">
                <v-icon size="64" color="warning">mdi-clock-alert</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ overdueCount }}</h3>
              <p class="text-subtitle-1 text-warning mb-0">متأخرين</p>
            </v-card>
          </v-col>
          <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center" elevation="2">
              <div class="stat-icon mb-3">
                <v-icon size="64" color="success">mdi-check-circle</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ paidCount }}</h3>
              <p class="text-subtitle-1 text-success mb-0">مدفوع</p>
            </v-card>
          </v-col>
        </v-row>

        <!-- شريط البحث والفلترة المحسن -->
        <v-card class="search-filter-card mb-8" elevation="2">
          <v-card-title class="d-flex align-center">
            <v-icon class="me-2" color="primary">mdi-filter</v-icon>
            البحث والفلترة
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12" md="4">
                <v-text-field
                  v-model="searchQuery"
                  label="البحث عن مديون"
                  prepend-inner-icon="mdi-magnify"
                  variant="outlined"
                  clearable
                  hide-details
                  class="search-field"
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-select
                  v-model="statusFilter"
                  label="حالة الدفع"
                  :items="statusOptions"
                  variant="outlined"
                  hide-details
                  clearable
                  class="filter-field"
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-select
                  v-model="amountFilter"
                  label="نطاق المبلغ"
                  :items="amountOptions"
                  variant="outlined"
                  hide-details
                  clearable
                  class="filter-field"
                />
              </v-col>
              <v-col cols="12" md="2">
                <v-btn
                  color="primary"
                  variant="outlined"
                  block
                  class="reset-button"
                  @click="resetFilters"
                >
                  <v-icon class="me-1">mdi-refresh</v-icon>
                  إعادة تعيين
                </v-btn>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>

        <!-- جدول المديونون المحسن -->
        <v-card class="data-table-card" elevation="2">
          <v-card-title class="d-flex align-center justify-space-between pa-6">
            <div class="d-flex align-center">
              <v-icon class="me-2" color="primary">mdi-table</v-icon>
              <span class="text-h6 font-weight-bold">قائمة المديونون</span>
              <v-chip class="ms-3" color="primary" variant="tonal">
                {{ filteredDebtors.length }} عنصر
              </v-chip>
            </div>
            <div class="d-flex align-center gap-2">
              <v-btn
                color="primary"
                size="default"
                prepend-icon="mdi-plus"
                class="add-button"
                @click="openAddDialog"
              >
                إضافة مديون جديد
              </v-btn>
              <v-btn
                icon="mdi-refresh"
                variant="text"
                color="primary"
                @click="refreshData"
                class="action-button"
              />
              <v-btn
                icon="mdi-download"
                variant="text"
                color="success"
                @click="exportData"
                class="action-button"
              />
            </div>
          </v-card-title>
          <v-data-table
            :headers="headers"
            :items="filteredDebtors"
            :loading="loading"
            class="elevation-0"
            no-data-text="لا توجد بيانات"
            loading-text="جاري التحميل..."
          >
            <!-- عمود الاسم -->
            <template v-slot:item.name="{ item }">
              <div class="d-flex align-center" style="cursor: pointer;" @click="viewDebtsAndPayments(item)">
                <v-avatar
                  :color="getStatusColor(item.status)"
                  size="32"
                  class="me-3"
                >
                  <span class="text-white font-weight-bold">
                    {{ item.name.charAt(0) }}
                  </span>
                </v-avatar>
                <div>
                  <div class="font-weight-medium text-primary">{{ item.name }}</div>
                  <div class="text-caption text-grey-darken-1">{{ item.email }}</div>
                  <div class="text-caption text-primary">انقر لعرض الديون والتسديدات</div>
                </div>
              </div>
            </template>

            <!-- عمود المبلغ -->
            <template v-slot:item.amount="{ item }">
              <div class="text-right">
                <div class="font-weight-bold text-h6">{{ formatCurrency(item.amount) }}</div>
                <div class="text-caption text-grey-darken-1">
                  {{ item.currency }}
                </div>
              </div>
            </template>

            <!-- عمود تاريخ الاستحقاق -->
            <template v-slot:item.dueDate="{ item }">
              <div class="text-center">
                <div class="font-weight-medium">{{ formatDate(item.dueDate) }}</div>
                <v-chip
                  :color="getDueDateColor(item.dueDate)"
                  size="small"
                  variant="tonal"
                >
                  {{ getDueDateStatus(item.dueDate) }}
                </v-chip>
              </div>
            </template>

            <!-- عمود الحالة -->
            <template v-slot:item.status="{ item }">
              <v-chip
                :color="getStatusColor(item.status)"
                :variant="item.status === 'paid' ? 'flat' : 'tonal'"
                size="small"
              >
                {{ getStatusText(item.status) }}
              </v-chip>
            </template>

            <!-- عمود الإجراءات -->
            <template v-slot:item.actions="{ item }">
              <div class="d-flex align-center gap-1">
                <v-btn
                  icon="mdi-eye"
                  size="small"
                  variant="text"
                  @click="viewDebtor(item)"
                />
                <v-btn
                  icon="mdi-pencil"
                  size="small"
                  variant="text"
                  @click="editDebtor(item)"
                />
                <v-btn
                  icon="mdi-credit-card"
                  size="small"
                  variant="text"
                  color="success"
                  @click="markAsPaid(item)"
                  v-if="item.status !== 'paid'"
                />
                <v-btn
                  icon="mdi-delete"
                  size="small"
                  variant="text"
                  color="error"
                  @click="deleteDebtor(item)"
                />
              </div>
            </template>
          </v-data-table>
        </v-card>
      </v-container>

    <!-- نافذة إضافة/تعديل مديون -->
    <v-dialog v-model="dialog" max-width="600px">
      <v-card>
        <v-card-title class="text-h5 font-weight-bold">
          {{ isEdit ? 'تعديل بيانات المديون' : 'إضافة مديون جديد' }}
        </v-card-title>
        <v-card-text>
          <v-form ref="form" v-model="valid">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="debtorForm.name"
                  label="الاسم الكامل"
                  variant="outlined"
                  :rules="[v => !!v || 'الاسم مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="debtorForm.email"
                  label="البريد الإلكتروني"
                  type="email"
                  variant="outlined"
                  :rules="[v => !!v || 'البريد الإلكتروني مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="debtorForm.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="debtorForm.amount"
                  label="المبلغ المطلوب"
                  type="number"
                  variant="outlined"
                  :rules="[v => !!v || 'المبلغ مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="debtorForm.currency"
                  label="العملة"
                  :items="currencyOptions"
                  variant="outlined"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="debtorForm.dueDate"
                  label="تاريخ الاستحقاق"
                  type="date"
                  variant="outlined"
                  required
                />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  v-model="debtorForm.notes"
                  label="ملاحظات"
                  variant="outlined"
                  rows="3"
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            color="grey"
            variant="text"
            @click="closeDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="primary"
            @click="saveDebtor"
            :disabled="!valid"
          >
            {{ isEdit ? 'تحديث' : 'إضافة' }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- نافذة عرض تفاصيل المديون -->
    <v-dialog v-model="viewDialog" max-width="500px">
      <v-card>
        <v-card-title class="text-h5 font-weight-bold">
          تفاصيل المديون
        </v-card-title>
        <v-card-text v-if="selectedDebtor">
          <v-row>
            <v-col cols="12">
              <div class="text-center mb-4">
                <v-avatar
                  :color="getStatusColor(selectedDebtor.status)"
                  size="64"
                >
                  <span class="text-white text-h4 font-weight-bold">
                    {{ selectedDebtor.name.charAt(0) }}
                  </span>
                </v-avatar>
                <h3 class="text-h5 font-weight-bold mt-2">{{ selectedDebtor.name }}</h3>
              </div>
            </v-col>
            <v-col cols="6">
              <strong>البريد الإلكتروني:</strong>
              <p>{{ selectedDebtor.email }}</p>
            </v-col>
            <v-col cols="6">
              <strong>رقم الهاتف:</strong>
              <p>{{ selectedDebtor.phone || 'غير محدد' }}</p>
            </v-col>
            <v-col cols="6">
              <strong>المبلغ المطلوب:</strong>
              <p class="text-h6 font-weight-bold text-error">
                {{ formatCurrency(selectedDebtor.amount) }}
              </p>
            </v-col>
            <v-col cols="6">
              <strong>تاريخ الاستحقاق:</strong>
              <p>{{ formatDate(selectedDebtor.dueDate) }}</p>
            </v-col>
            <v-col cols="12">
              <strong>الحالة:</strong>
              <v-chip
                :color="getStatusColor(selectedDebtor.status)"
                :variant="selectedDebtor.status === 'paid' ? 'flat' : 'tonal'"
                class="mt-1"
              >
                {{ getStatusText(selectedDebtor.status) }}
              </v-chip>
            </v-col>
            <v-col cols="12" v-if="selectedDebtor.notes">
              <strong>ملاحظات:</strong>
              <p>{{ selectedDebtor.notes }}</p>
            </v-col>
          </v-row>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            color="primary"
            @click="viewDialog = false"
          >
            إغلاق
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- نافذة عرض الديون والتسديدات -->
    <v-dialog v-model="debtsPaymentsDialog" max-width="800px">
      <v-card>
        <v-card-title class="text-h5 font-weight-bold d-flex align-center">
          <v-icon class="me-2" color="primary">mdi-credit-card</v-icon>
          الديون والتسديدات - {{ selectedDebtor?.name }}
        </v-card-title>
        <v-card-text v-if="selectedDebtor">
          <!-- إحصائيات سريعة -->
          <v-row class="mb-4">
            <v-col cols="12" md="4">
              <v-card color="error" variant="tonal" class="pa-3">
                <div class="text-center">
                  <v-icon size="32" color="error" class="mb-2">mdi-currency-usd</v-icon>
                  <h3 class="text-h5 font-weight-bold">{{ formatCurrency(selectedDebtor.amount) }}</h3>
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

          <!-- تبويبات الديون والتسديدات -->
          <v-tabs v-model="activeTab" class="mb-4">
            <v-tab value="debts">الديون</v-tab>
            <v-tab value="payments">التسديدات</v-tab>
            <v-tab value="summary">الملخص</v-tab>
          </v-tabs>

          <v-window v-model="activeTab">
            <!-- تبويب الديون -->
            <v-window-item value="debts">
              <v-card variant="outlined">
                <v-card-title class="d-flex align-center justify-space-between">
                  <span>قائمة الديون</span>
                  <v-btn
                    color="primary"
                    size="small"
                    prepend-icon="mdi-plus"
                    @click="addDebt"
                  >
                    إضافة دين
                  </v-btn>
                </v-card-title>
                <v-data-table
                  :headers="debtHeaders"
                  :items="selectedDebtor.debts || []"
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
                      icon="mdi-pencil"
                      size="small"
                      variant="text"
                      @click="editDebt(item)"
                    />
                    <v-btn
                      icon="mdi-delete"
                      size="small"
                      variant="text"
                      color="error"
                      @click="deleteDebt(item)"
                    />
                  </template>
                </v-data-table>
              </v-card>
            </v-window-item>

            <!-- تبويب التسديدات -->
            <v-window-item value="payments">
              <v-card variant="outlined">
                <v-card-title class="d-flex align-center justify-space-between">
                  <span>قائمة التسديدات</span>
                  <v-btn
                    color="success"
                    size="small"
                    prepend-icon="mdi-plus"
                    @click="addPayment"
                  >
                    إضافة تسديد
                  </v-btn>
                </v-card-title>
                <v-data-table
                  :headers="paymentHeaders"
                  :items="selectedDebtor.payments || []"
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
                      icon="mdi-pencil"
                      size="small"
                      variant="text"
                      @click="editPayment(item)"
                    />
                    <v-btn
                      icon="mdi-delete"
                      size="small"
                      variant="text"
                      color="error"
                      @click="deletePayment(item)"
                    />
                  </template>
                </v-data-table>
              </v-card>
            </v-window-item>

            <!-- تبويب الملخص -->
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
                            <span class="font-weight-bold text-error">{{ formatCurrency(selectedDebtor.amount) }}</span>
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
                      <v-list v-if="selectedDebtor.payments && selectedDebtor.payments.length > 0">
                        <v-list-item
                          v-for="payment in selectedDebtor.payments.slice(0, 3)"
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
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            color="primary"
            @click="debtsPaymentsDialog = false"
          >
            إغلاق
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

// البيانات التفاعلية
const drawer = ref(true)
const loading = ref(false)
const dialog = ref(false)
const viewDialog = ref(false)
const debtsPaymentsDialog = ref(false)
const valid = ref(false)
const isEdit = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const amountFilter = ref('')
const selectedDebtor = ref(null)
const activeTab = ref('debts')

// نموذج المدين
const debtorForm = ref({
  name: '',
  email: '',
  phone: '',
  amount: '',
  currency: 'IQD',
  dueDate: '',
  notes: ''
})

// قائمة المدينين
const debtors = ref([
  {
    id: 1,
    name: 'أحمد محمد العلي',
    email: 'ahmed.ali@example.com',
    phone: '+966501234567',
    amount: 15000,
    currency: 'IQD',
    dueDate: '2024-01-15',
    status: 'overdue',
    notes: 'مدين من مشروع تطوير الموقع',
    debts: [
      { id: 1, amount: 10000, date: '2024-01-01', description: 'دفعة أولى - تطوير الموقع', status: 'unpaid' },
      { id: 2, amount: 5000, date: '2024-01-10', description: 'دفعة ثانية - إضافات الموقع', status: 'unpaid' }
    ],
    payments: [
      { id: 1, amount: 2000, date: '2024-01-05', method: 'bank_transfer', description: 'دفعة جزئية' },
      { id: 2, amount: 1000, date: '2024-01-12', method: 'cash', description: 'دفعة نقدية' }
    ]
  },
  {
    id: 2,
    name: 'فاطمة عبدالله السعد',
    email: 'fatima.saad@example.com',
    phone: '+966502345678',
    amount: 8500,
    currency: 'IQD',
    dueDate: '2024-02-20',
    status: 'pending',
    notes: 'مدين من خدمات الاستشارة',
    debts: [
      { id: 3, amount: 8500, date: '2024-01-15', description: 'رسوم الاستشارة', status: 'unpaid' }
    ],
    payments: []
  },
  {
    id: 3,
    name: 'محمد سالم القحطاني',
    email: 'mohammed.qhtani@example.com',
    phone: '+966503456789',
    amount: 25000,
    currency: 'IQD',
    dueDate: '2024-01-10',
    status: 'paid',
    notes: 'تم السداد كاملاً',
    debts: [
      { id: 4, amount: 25000, date: '2024-01-01', description: 'مشروع تطوير تطبيق', status: 'paid' }
    ],
    payments: [
      { id: 3, amount: 25000, date: '2024-01-10', method: 'bank_transfer', description: 'سداد كامل' }
    ]
  },
  {
    id: 4,
    name: 'نورا عبدالرحمن الشمري',
    email: 'nora.shamri@example.com',
    phone: '+966504567890',
    amount: 12000,
    currency: 'IQD',
    dueDate: '2024-03-05',
    status: 'pending',
    notes: 'مدين من مشروع التصميم',
    debts: [
      { id: 5, amount: 8000, date: '2024-01-20', description: 'تصميم الهوية البصرية', status: 'unpaid' },
      { id: 6, amount: 4000, date: '2024-02-01', description: 'تصميم الموقع', status: 'unpaid' }
    ],
    payments: [
      { id: 4, amount: 3000, date: '2024-02-15', method: 'credit_card', description: 'دفعة جزئية' }
    ]
  },
  {
    id: 5,
    name: 'خالد أحمد المطيري',
    email: 'khalid.mutairi@example.com',
    phone: '+966505678901',
    amount: 18000,
    currency: 'IQD',
    dueDate: '2024-01-25',
    status: 'overdue',
    notes: 'مدين من خدمات البرمجة',
    debts: [
      { id: 7, amount: 12000, date: '2024-01-01', description: 'برمجة النظام الأساسي', status: 'unpaid' },
      { id: 8, amount: 6000, date: '2024-01-15', description: 'إضافات النظام', status: 'unpaid' }
    ],
    payments: [
      { id: 5, amount: 5000, date: '2024-01-20', method: 'bank_transfer', description: 'دفعة جزئية' }
    ]
  }
])

// عناوين الجدول
const headers = [
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'المبلغ المطلوب', key: 'amount', sortable: true },
  { title: 'تاريخ الاستحقاق', key: 'dueDate', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// عناوين جدول الديون
const debtHeaders = [
  { title: 'المبلغ', key: 'amount', sortable: true },
  { title: 'التاريخ', key: 'date', sortable: true },
  { title: 'الوصف', key: 'description', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// عناوين جدول التسديدات
const paymentHeaders = [
  { title: 'المبلغ', key: 'amount', sortable: true },
  { title: 'التاريخ', key: 'date', sortable: true },
  { title: 'طريقة الدفع', key: 'method', sortable: true },
  { title: 'الوصف', key: 'description', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// خيارات الفلترة
const statusOptions = [
  { title: 'جميع الحالات', value: '' },
  { title: 'متأخر', value: 'overdue' },
  { title: 'معلق', value: 'pending' },
  { title: 'مدفوع', value: 'paid' }
]

const amountOptions = [
  { title: 'جميع المبالغ', value: '' },
  { title: 'أقل من 10,000', value: 'low' },
  { title: '10,000 - 20,000', value: 'medium' },
  { title: 'أكثر من 20,000', value: 'high' }
]

const currencyOptions = [
  { title: 'دينار عراقي', value: 'IQD' },
  { title: 'دولار أمريكي', value: 'USD' },
  { title: 'يورو', value: 'EUR' }
]

// عناصر القائمة الرئيسية
const mainMenuItems = [
  { title: 'الرئيسية', icon: 'mdi-view-dashboard', to: '/', active: false },
  { title: 'المشاريع', icon: 'mdi-folder', to: '/project-management', active: false },
  { title: 'المهندسين', icon: 'mdi-account-hard-hat', to: '/engineers', active: false },
  { title: 'المصاريف الإدارية', icon: 'mdi-currency-usd', to: '/administrative-expenses', active: false },
  { title: 'المصروفات العامة', icon: 'mdi-chart-line', to: '/expenses', active: false },
  { title: 'الإيرادات', icon: 'mdi-trending-up', to: '/income', active: false },
  { title: 'المدينين', icon: 'mdi-credit-card', to: '/debtors', active: true },
  { title: 'المستخدمين', icon: 'mdi-account-group', to: '/users', active: false }
]

// الحسابات
const filteredDebtors = computed(() => {
  let filtered = debtors.value

  // فلترة بالبحث
  if (searchQuery.value) {
    filtered = filtered.filter(debtor =>
      debtor.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      debtor.email.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  // فلترة بالحالة
  if (statusFilter.value) {
    filtered = filtered.filter(debtor => debtor.status === statusFilter.value)
  }

  // فلترة بالمبلغ
  if (amountFilter.value) {
    filtered = filtered.filter(debtor => {
      const amount = debtor.amount
      switch (amountFilter.value) {
        case 'low': return amount < 10000
        case 'medium': return amount >= 10000 && amount <= 20000
        case 'high': return amount > 20000
        default: return true
      }
    })
  }

  return filtered
})

const totalDebt = computed(() => {
  return debtors.value
    .filter(debtor => debtor.status !== 'paid')
    .reduce((sum, debtor) => sum + debtor.amount, 0)
})

const overdueCount = computed(() => {
  return debtors.value.filter(debtor => debtor.status === 'overdue').length
})

const paidCount = computed(() => {
  return debtors.value.filter(debtor => debtor.status === 'paid').length
})

// حسابات الديون والتسديدات
const totalPaid = computed(() => {
  if (!selectedDebtor.value || !selectedDebtor.value.payments) return 0
  return selectedDebtor.value.payments.reduce((sum, payment) => sum + payment.amount, 0)
})

const remainingAmount = computed(() => {
  if (!selectedDebtor.value) return 0
  return selectedDebtor.value.amount - totalPaid.value
})

const paymentPercentage = computed(() => {
  if (!selectedDebtor.value || selectedDebtor.value.amount === 0) return 0
  return Math.round((totalPaid.value / selectedDebtor.value.amount) * 100)
})

// الدوال
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD'
  }).format(amount)
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('ar-SA')
}

const getStatusColor = (status) => {
  const colors = {
    'overdue': 'error',
    'pending': 'warning',
    'paid': 'success'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    'overdue': 'متأخر',
    'pending': 'معلق',
    'paid': 'مدفوع'
  }
  return texts[status] || 'غير محدد'
}

const getDueDateColor = (dueDate) => {
  const today = new Date()
  const due = new Date(dueDate)
  const diffTime = due - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays < 0) return 'error'
  if (diffDays <= 7) return 'warning'
  return 'success'
}

const getDueDateStatus = (dueDate) => {
  const today = new Date()
  const due = new Date(dueDate)
  const diffTime = due - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays < 0) return 'متأخر'
  if (diffDays === 0) return 'اليوم'
  if (diffDays <= 7) return `${diffDays} أيام`
  return 'مستقبلي'
}

const openAddDialog = () => {
  isEdit.value = false
  debtorForm.value = {
    name: '',
    email: '',
    phone: '',
  amount: '',
  currency: 'IQD',
  dueDate: '',
    notes: ''
  }
  dialog.value = true
}

const editDebtor = (debtor) => {
  isEdit.value = true
  debtorForm.value = { ...debtor }
  dialog.value = true
}

const viewDebtor = (debtor) => {
  selectedDebtor.value = debtor
  viewDialog.value = true
}

const saveDebtor = () => {
  if (isEdit.value) {
    const index = debtors.value.findIndex(d => d.id === debtorForm.value.id)
    if (index !== -1) {
      debtors.value[index] = { ...debtorForm.value }
    }
  } else {
    const newDebtor = {
      ...debtorForm.value,
      id: Date.now(),
      status: 'pending'
    }
    debtors.value.push(newDebtor)
  }
  closeDialog()
}

const closeDialog = () => {
  dialog.value = false
  valid.value = false
}

const markAsPaid = (debtor) => {
  debtor.status = 'paid'
}

const deleteDebtor = (debtor) => {
  const index = debtors.value.findIndex(d => d.id === debtor.id)
  if (index !== -1) {
    debtors.value.splice(index, 1)
  }
}

const refreshData = () => {
  loading.value = true
  setTimeout(() => {
    loading.value = false
  }, 1000)
}

const exportData = () => {
  // منطق التصدير
}

const resetFilters = () => {
  searchQuery.value = ''
  statusFilter.value = ''
  amountFilter.value = ''
}

// دوال الديون والتسديدات
const viewDebtsAndPayments = (debtor) => {
  selectedDebtor.value = debtor
  activeTab.value = 'debts'
  debtsPaymentsDialog.value = true
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

const addDebt = () => {
  // منطق إضافة دين جديد
}

const editDebt = (debt) => {
  // منطق تعديل الدين
}

const deleteDebt = (debt) => {
  // منطق حذف الدين
}

const addPayment = () => {
  // منطق إضافة تسديد جديد
}

const editPayment = (payment) => {
  // منطق تعديل التسديد
}

const deletePayment = (payment) => {
  // منطق حذف التسديد
}

onMounted(() => {
  // تهيئة البيانات
})
</script>


<style>
/* Import page styles */
@import './styles/debtors.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
