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
  console.log('تصدير البيانات')
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
  console.log('إضافة دين جديد')
}

const editDebt = (debt) => {
  // منطق تعديل الدين
  console.log('تعديل الدين:', debt)
}

const deleteDebt = (debt) => {
  // منطق حذف الدين
  console.log('حذف الدين:', debt)
}

const addPayment = () => {
  // منطق إضافة تسديد جديد
  console.log('إضافة تسديد جديد')
}

const editPayment = (payment) => {
  // منطق تعديل التسديد
  console.log('تعديل التسديد:', payment)
}

const deletePayment = (payment) => {
  // منطق حذف التسديد
  console.log('حذف التسديد:', payment)
}

onMounted(() => {
  // تهيئة البيانات
})
</script>

<style scoped>
/* ========================================
   تحسين ألوان صفحة المديونون
   ======================================== */

/* خلفية الصفحة */
.data-page {
  background: #ffffff;
  color: #1e293b;
  min-height: 100vh;
}

/* ========================================
   رأس الصفحة
   ======================================== */

.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #667eea 100%) !important;
  backdrop-filter: blur(20px);
  border-radius: 16px;
  padding: 20px 30px;
  margin-bottom: 20px;
  text-align: center !important;
  border: 2px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
  animation: gradient-animation 8s ease infinite;
  background-size: 200% 200%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: white;
}

.header-top-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 10px;
  width: 100%;
  text-align: center;
}

.page-icon {
  font-size: 2rem;
  display: inline-block;
  animation: iconFloat 3s ease-in-out infinite;
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes gradient-animation {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.gradient-animation {
  background-size: 200% 200%;
  animation: gradient-animation 3s ease infinite;
}

.action-buttons-container {
  padding: 0 2rem;
  display: flex;
  justify-content: flex-start;
}


.add-button {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.95rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
  padding: 10px 20px !important;
  border-radius: 10px !important;
  box-shadow: 0 4px 16px rgba(16, 185, 129, 0.4), 0 2px 6px rgba(5, 150, 105, 0.3) !important;
  transition: all 0.3s ease !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  position: relative;
  overflow: hidden;
  text-transform: none !important;
  height: auto !important;
}

.add-button :deep(.v-btn__content) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 700 !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
}

.add-button :deep(.v-icon) {
  color: #ffffff !important;
  fill: #ffffff !important;
  font-size: 1.1rem !important;
  margin-inline-end: 8px !important;
}

.add-button :deep(.v-btn__content) {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 700 !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
}

.add-button :deep(.v-icon) {
  font-size: 1.2rem !important;
  margin-inline-end: 8px !important;
}

.add-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.5s ease;
}

.add-button:hover::before {
  left: 100%;
}

.add-button:hover {
  transform: translateY(-3px) scale(1.02) !important;
  box-shadow: 0 10px 30px rgba(16, 185, 129, 0.5), 0 4px 12px rgba(5, 150, 105, 0.4) !important;
  background: linear-gradient(135deg, #059669 0%, #047857 50%, #065f46 100%) !important;
}

.add-button:active {
  transform: translateY(-1px) scale(0.98) !important;
}

/* ========================================
   بطاقات الإحصائيات
   ======================================== */

.stat-card {
  background: var(--bg-white) !important;
  border-radius: 16px !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  min-height: 200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center !important;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08) !important;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(135deg, #2563eb 0%, #7c3aed 100%);
}

.stat-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
}

.stat-icon {
  position: relative;
  z-index: 2;
}

.stat-card .v-icon {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.1));
  transition: all 0.3s ease;
}

.stat-card:hover .v-icon {
  transform: scale(1.1);
  filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.2));
}

.stat-card h3 {
  font-size: 2.5rem !important;
  font-weight: 700 !important;
  margin-bottom: 8px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center !important;
  width: 100%;
}

.stat-card p {
  font-size: 1rem !important;
  font-weight: 500 !important;
  opacity: 0.8;
  text-align: center !important;
  width: 100%;
}

/* تأثيرات خاصة لكل لون */
.stat-card:nth-child(1)::before {
  background: var(--gradient-primary);
}

.stat-card:nth-child(2)::before {
  background: var(--gradient-error);
}

.stat-card:nth-child(3)::before {
  background: var(--gradient-warning);
}

.stat-card:nth-child(4)::before {
  background: var(--gradient-success);
}

/* ========================================
   شريط البحث والفلترة
   ======================================== */

.search-filter-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #667eea 100%) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  border-radius: 16px !important;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3) !important;
}

.search-filter-card,
.search-filter-card .v-card,
.search-filter-card .v-card-title,
.search-filter-card .v-card-text {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #667eea 100%) !important;
}

/* ========================================
   إصلاح لون النصوص في شريط البحث والفلترة
   - تغيير لون النص من الأبيض إلى الأسود الداكن
   - تحسين وضوح جميع النصوص في شريط البحث
   - إصلاح التسميات والحقول والقوائم المنسدلة
   ======================================== */

/* إصلاح عنوان شريط البحث والفلترة */
.search-filter-card .v-card-title {
  color: #ffffff !important;
  font-weight: 700 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.search-filter-card .v-card-title * {
  color: #ffffff !important;
}

/* إصلاح النصوص في شريط البحث والفلترة */
.search-filter-card .v-card-text,
.search-filter-card .v-card-text *,
.search-filter-card .v-row,
.search-filter-card .v-row *,
.search-filter-card .v-col,
.search-filter-card .v-col * {
  color: #1a1a1a !important;
}

/* إصلاح التسميات في شريط البحث والفلترة */
.search-filter-card :deep(.v-label),
.search-filter-card :deep(.v-label *),
.search-filter-card :deep(.v-field__label),
.search-filter-card :deep(.v-field__label *),
.search-filter-card :deep(.v-field-label) {
  color: #ffffff !important;
}

/* تسميات حالة الدفع ونطاق المبلغ - الحالة الافتراضية */
.search-filter-card :deep(.filter-field .v-field__label),
.search-filter-card :deep(.filter-field .v-label),
.search-filter-card :deep(.filter-field .v-field-label),
.search-filter-card :deep(.filter-field .v-field__label *),
.search-filter-card :deep(.filter-field .v-label *) {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  font-weight: 600 !important;
}

.search-filter-card :deep(.filter-field .v-field--focused .v-field__label),
.search-filter-card :deep(.filter-field .v-field--active .v-field__label),
.search-filter-card :deep(.filter-field .v-field--focused .v-field__label *),
.search-filter-card :deep(.filter-field .v-field--active .v-field__label *) {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* إصلاح النصوص في الحقول */
.search-filter-card .v-text-field .v-field__input,
.search-filter-card .v-select .v-field__input,
.search-filter-card .v-field__input,
.search-filter-card .v-field__input * {
  color: #1a1a1a !important;
}

/* إصلاح خلفية الحقول */
.search-filter-card .v-text-field .v-field,
.search-filter-card .v-select .v-field {
  background: rgba(255, 255, 255, 0.9) !important;
  border: 1px solid rgba(14, 165, 233, 0.3) !important;
}

.search-field,
.filter-field {
  background: rgba(255, 255, 255, 0.9) !important;
  border-radius: 12px !important;
}

/* إصلاح القوائم المنسدلة في شريط البحث والفلترة */
.search-filter-card .v-menu__content,
.search-filter-card .v-menu__content *,
.search-filter-card .v-overlay__content,
.search-filter-card .v-overlay__content * {
  color: #1a1a1a !important;
  background: white !important;
}

.search-filter-card .v-list,
.search-filter-card .v-list *,
.search-filter-card .v-list-item,
.search-filter-card .v-list-item *,
.search-filter-card .v-list-item-title,
.search-filter-card .v-list-item-subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح خاص للقوائم المنسدلة */
.search-filter-card .v-select__menu,
.search-filter-card .v-select__menu *,
.search-filter-card .v-select__menu .v-list-item,
.search-filter-card .v-select__menu .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في جميع الحقول */
.search-filter-card .v-input,
.search-filter-card .v-input *,
.search-filter-card .v-text-field,
.search-filter-card .v-text-field *,
.search-filter-card .v-select,
.search-filter-card .v-select * {
  color: #1a1a1a !important;
}

/* إصلاح شامل لجميع النصوص في شريط البحث */
.search-filter-card *,
.search-filter-card * * {
  color: #1a1a1a !important;
}

/* استثناء التسميات من القاعدة العامة */
.search-filter-card :deep(.v-label),
.search-filter-card :deep(.v-field__label),
.search-filter-card :deep(.v-field-label),
.search-filter-card :deep(.v-label *),
.search-filter-card :deep(.v-field__label *) {
  color: #ffffff !important;
}

/* استثناء تسميات حالة الدفع ونطاق المبلغ - لون أسود */
.search-filter-card :deep(.filter-field .v-label),
.search-filter-card :deep(.filter-field .v-field__label),
.search-filter-card :deep(.filter-field .v-field-label),
.search-filter-card :deep(.filter-field .v-label *),
.search-filter-card :deep(.filter-field .v-field__label *) {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.reset-button {
  border-radius: 12px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  color: #ffffff !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
  background: rgba(255, 255, 255, 0.1) !important;
}

.reset-button,
.reset-button :deep(.v-btn__content),
.reset-button :deep(.v-btn__content *),
.reset-button :deep(.v-icon),
.reset-button :deep(.v-icon *) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.reset-button:hover {
  background: rgba(255, 255, 255, 0.2) !important;
  border-color: rgba(255, 255, 255, 0.7) !important;
}

.reset-button:hover :deep(.v-btn__content),
.reset-button:hover :deep(.v-icon) {
  color: #ffffff !important;
}

/* ========================================
   جدول البيانات
   ======================================== */

.data-table-card {
  background: #ffffff !important;
  border: 1px solid #e2e8f0 !important;
  border-radius: 16px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
  overflow: hidden;
}

.data-table-card .v-card-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3);
  padding: 1.5rem;
  position: relative;
}

.data-table-card .v-card-title .v-icon {
  color: #ffffff !important;
}

.data-table-card .v-card-title span {
  color: #ffffff !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

.data-table-card .v-card-title .v-chip {
  color: #ffffff !important;
  background: rgba(255, 255, 255, 0.2) !important;
  border: 1px solid rgba(255, 255, 255, 0.3) !important;
}

.action-button {
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

.action-button:hover {
  background: #f1f5f9 !important;
  transform: scale(1.1) !important;
}

/* الجداول */
.v-data-table {
  background: #ffffff !important;
  color: #1e293b !important;
  border-radius: 0 0 16px 16px !important;
}

.v-data-table th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2) !important;
  padding: 16px !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
  text-transform: none !important;
}

/* ========================================
   إصلاح لون عناوين الجدول في صفحة المديونين
   - تغيير لون النص من الفايت إلى الأسود الداكن
   - تحسين وضوح جميع عناوين الجدول
   - محاذاة العناوين في المنتصف
   ======================================== */

/* إصلاح شامل لعناوين الجدول */
.v-data-table .v-data-table-header th,
.v-data-table .v-data-table-header th *,
.v-data-table thead th,
.v-data-table thead th * {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
  text-transform: none !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2) !important;
}

/* قواعد إضافية لضمان اللون الأبيض على جميع عناصر عناوين الجدول */
.v-data-table :deep(.v-data-table-header th),
.v-data-table :deep(.v-data-table-header th *),
.v-data-table :deep(.v-data-table-header th span),
.v-data-table :deep(.v-data-table-header th div),
.v-data-table :deep(.v-data-table-header th .v-data-table-header__content),
.v-data-table :deep(.v-data-table-header th .v-data-table-header__content *),
.v-data-table :deep(.v-data-table__wrapper table thead tr th),
.v-data-table :deep(.v-data-table__wrapper table thead tr th *),
.v-data-table :deep(.v-data-table__wrapper table thead tr th span),
.v-data-table :deep(.v-data-table__wrapper table thead tr th div),
.v-data-table :deep(.v-data-table__th),
.v-data-table :deep(.v-data-table__th *),
.v-data-table :deep(.v-data-table__th span),
.v-data-table :deep(.v-data-table__th div) {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
  font-size: 0.75rem !important;
}

/* تطبيق اللون الأبيض على أيقونات في عناوين الجدول */
.v-data-table :deep(.v-data-table-header th .v-icon),
.v-data-table :deep(.v-data-table-header th i),
.v-data-table :deep(.v-data-table-header th svg),
.v-data-table :deep(.v-data-table-header th .v-data-table-header__sort-badge),
.v-data-table :deep(.v-data-table-header th .v-data-table-header__sort-icon) {
  color: #ffffff !important;
  fill: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* إصلاح خلايا الجدول */
.v-data-table td {
  color: #1a1a1a !important;
  border-bottom: 1px solid #f1f5f9 !important;
  padding: 16px !important;
  font-size: 0.9rem !important;
  text-align: center !important;
  vertical-align: middle !important;
}

/* إصلاح النصوص في خلايا الجدول */
.v-data-table tbody td,
.v-data-table tbody td *,
.v-data-table tbody tr td,
.v-data-table tbody tr td * {
  color: #1a1a1a !important;
}

/* إصلاح نهائي شامل لجميع عناصر الجدول */
.v-data-table,
.v-data-table *,
.v-data-table tbody,
.v-data-table tbody *,
.v-data-table tbody tr,
.v-data-table tbody tr *,
.v-data-table tbody tr td,
.v-data-table tbody tr td * {
  color: #1a1a1a !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-style: normal !important;
}

/* إصلاح النصوص في جميع العناصر */
.v-data-table .font-weight-bold,
.v-data-table .text-error,
.v-data-table .text-success,
.v-data-table .text-warning,
.v-data-table span,
.v-data-table div,
.v-data-table p {
  color: #1a1a1a !important;
}

/* استثناء للنصوص الملونة */
.v-data-table .text-error {
  color: #dc2626 !important;
}

.v-data-table .text-success {
  color: #059669 !important;
}

.v-data-table .text-warning {
  color: #d97706 !important;
}

/* ========================================
   إصلاح لون النصوص في عمود الإجراءات
   - تحسين وضوح أزرار الإجراءات في الجدول
   - ضمان وضوح الأيقونات والنصوص
   - تحسين التباين والوضوح
   ======================================== */

/* إصلاح أزرار الإجراءات في الجدول */
.v-data-table .v-btn,
.v-data-table .v-btn *,
.v-data-table .v-btn .v-icon,
.v-data-table .v-btn .v-btn__content {
  color: #1a1a1a !important;
}

/* إصلاح أزرار الإجراءات المحددة */
.v-data-table .v-btn[color="success"],
.v-data-table .v-btn[color="success"] *,
.v-data-table .v-btn[color="success"] .v-icon,
.v-data-table .v-btn[color="success"] .v-btn__content {
  color: #059669 !important;
}

.v-data-table .v-btn[color="error"],
.v-data-table .v-btn[color="error"] *,
.v-data-table .v-btn[color="error"] .v-icon,
.v-data-table .v-btn[color="error"] .v-btn__content {
  color: #dc2626 !important;
}

.v-data-table .v-btn[color="primary"],
.v-data-table .v-btn[color="primary"] *,
.v-data-table .v-btn[color="primary"] .v-icon,
.v-data-table .v-btn[color="primary"] .v-btn__content {
  color: #1976d2 !important;
}

/* إصلاح أزرار الإجراءات عند التمرير */
.v-data-table .v-btn:hover,
.v-data-table .v-btn:hover *,
.v-data-table .v-btn:hover .v-icon,
.v-data-table .v-btn:hover .v-btn__content {
  color: #1a1a1a !important;
}

/* إصلاح أزرار الإجراءات الملونة عند التمرير */
.v-data-table .v-btn[color="success"]:hover,
.v-data-table .v-btn[color="success"]:hover *,
.v-data-table .v-btn[color="success"]:hover .v-icon {
  color: #047857 !important;
}

.v-data-table .v-btn[color="error"]:hover,
.v-data-table .v-btn[color="error"]:hover *,
.v-data-table .v-btn[color="error"]:hover .v-icon {
  color: #b91c1c !important;
}

.v-data-table .v-btn[color="primary"]:hover,
.v-data-table .v-btn[color="primary"]:hover *,
.v-data-table .v-btn[color="primary"]:hover .v-icon {
  color: #1565c0 !important;
}

/* إصلاح شامل لجميع أزرار الإجراءات */
.v-data-table .d-flex,
.v-data-table .d-flex *,
.v-data-table .gap-1,
.v-data-table .gap-1 *,
.v-data-table .align-center,
.v-data-table .align-center * {
  color: #1a1a1a !important;
}

/* إصلاح أزرار الإجراءات في جميع الحالات */
.v-data-table .v-btn.v-btn--size-small,
.v-data-table .v-btn.v-btn--size-small *,
.v-data-table .v-btn.v-btn--variant-text,
.v-data-table .v-btn.v-btn--variant-text * {
  color: #1a1a1a !important;
}

/* إصلاح خاص للأزرار الملونة */
.v-data-table .v-btn[color="success"].v-btn--variant-text,
.v-data-table .v-btn[color="success"].v-btn--variant-text * {
  color: #059669 !important;
}

.v-data-table .v-btn[color="error"].v-btn--variant-text,
.v-data-table .v-btn[color="error"].v-btn--variant-text * {
  color: #dc2626 !important;
}

/* إصلاح أزرار الإجراءات في جميع الجداول الفرعية */
.v-data-table .v-data-table .v-btn,
.v-data-table .v-data-table .v-btn * {
  color: #1a1a1a !important;
}

/* إصلاح أزرار الإجراءات الملونة في الجداول الفرعية */
.v-data-table .v-data-table .v-btn[color="success"],
.v-data-table .v-data-table .v-btn[color="success"] * {
  color: #059669 !important;
}

.v-data-table .v-data-table .v-btn[color="error"],
.v-data-table .v-data-table .v-btn[color="error"] * {
  color: #dc2626 !important;
}

/* إصلاح نهائي شامل لجميع أزرار الإجراءات */
.v-data-table .v-btn .mdi,
.v-data-table .v-btn .mdi *,
.v-data-table .v-btn .mdi-eye,
.v-data-table .v-btn .mdi-pencil,
.v-data-table .v-btn .mdi-credit-card,
.v-data-table .v-btn .mdi-delete {
  color: inherit !important;
}

/* إصلاح أزرار الإجراءات في جميع الجداول */
.v-data-table .v-btn[icon="mdi-eye"],
.v-data-table .v-btn[icon="mdi-pencil"],
.v-data-table .v-btn[icon="mdi-credit-card"],
.v-data-table .v-btn[icon="mdi-delete"] {
  color: #1a1a1a !important;
}

/* إصلاح أزرار الإجراءات الملونة */
.v-data-table .v-btn[icon="mdi-credit-card"][color="success"] {
  color: #059669 !important;
}

.v-data-table .v-btn[icon="mdi-delete"][color="error"] {
  color: #dc2626 !important;
}

/* إصلاح أزرار الإجراءات عند التمرير */
.v-data-table .v-btn:hover .mdi,
.v-data-table .v-btn:hover .mdi * {
  color: inherit !important;
}

/* إصلاح شامل لجميع النصوص في الصفحة */
.debtors-page *,
.debtors-page * * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع الكروت */
.v-card,
.v-card *,
.v-card .v-card-title,
.v-card .v-card-title *,
.v-card .v-card-text,
.v-card .v-card-text * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع الأزرار */
.v-btn,
.v-btn *,
.v-btn .v-btn__content,
.v-btn .v-btn__content * {
  color: white !important;
}

/* استثناء للأزرار المحددة */
.v-btn.v-btn--variant-outlined,
.v-btn.v-btn--variant-outlined * {
  color: #1976d2 !important;
}

.v-data-table tbody tr:hover {
  background: #f8fafc !important;
}

/* ========================================
   الأزرار والتفاعل
   ======================================== */

.v-btn {
  font-weight: 500 !important;
  text-transform: none !important;
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

.v-btn--variant-elevated {
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%) !important;
  color: #ffffff !important;
  box-shadow: 0 4px 12px rgba(37, 99, 235, 0.3) !important;
}

.v-btn--variant-elevated:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(37, 99, 235, 0.4) !important;
}

.v-btn--variant-outlined {
  border: 2px solid #2563eb !important;
  color: #2563eb !important;
  background: transparent !important;
}

.v-btn--variant-outlined:hover {
  background: #2563eb !important;
  color: #ffffff !important;
  transform: translateY(-1px) !important;
}

/* ========================================
   الشارات والرموز
   ======================================== */

.v-chip {
  color: #1e293b !important;
  background: #f1f5f9 !important;
  border: 1px solid #e2e8f0 !important;
  font-weight: 500 !important;
  border-radius: 8px !important;
}

/* ========================================
   النماذج
   ======================================== */

.v-text-field input,
.v-select input,
.v-textarea textarea {
  color: #1e293b !important;
  background: #ffffff !important;
  font-size: 0.9rem !important;
}

.v-text-field label,
.v-select label,
.v-textarea label {
  color: #64748b !important;
  font-weight: 500 !important;
}

/* ========================================
   البطاقات العامة
   ======================================== */

.v-card {
  background: #ffffff !important;
  color: #1e293b !important;
  border: 1px solid #e2e8f0 !important;
  border-radius: 16px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}

.v-card-title {
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 1.1rem !important;
}

.v-card-text {
  color: #64748b !important;
  font-size: 0.9rem !important;
}
.sidebar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  direction: rtl;
  text-align: right;
}

.rtl-sidebar {
  direction: rtl;
  text-align: right;
}

.rtl-sidebar .v-list-item {
  text-align: right;
  direction: rtl;
}

.rtl-sidebar .v-list-item-title {
  text-align: right;
  direction: rtl;
}

.menu-item {
  transition: all 0.3s ease;
}

.menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.logo-section {
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.v-card {
  border-radius: 12px;
}

.v-data-table {
  border-radius: 12px;
}

.gap-1 {
  gap: 4px;
}

.gap-2 {
  gap: 8px;
}

/* دعم RTL شامل */
.v-main {
  direction: rtl;
  text-align: right;
}

.v-container {
  direction: rtl;
  text-align: right;
}

.v-card {
  direction: rtl;
  text-align: right;
}

.v-card-title {
  text-align: right;
  direction: rtl;
}

.v-card-text {
  text-align: right;
  direction: rtl;
}

.v-data-table {
  direction: rtl;
  text-align: right;
}

.v-data-table th,
.v-data-table td {
  text-align: right;
  direction: rtl;
}

.v-text-field input {
  text-align: right;
  direction: rtl;
}

.v-select .v-field__input {
  text-align: right;
  direction: rtl;
}

.v-btn {
  direction: rtl;
}

.v-icon {
  transform: scaleX(-1);
}

h1, h2, h3, h4, h5, h6, p, span, div {
  text-align: right;
  direction: rtl;
}

.v-list {
  direction: rtl;
  text-align: right;
}

.v-dialog .v-card {
  direction: rtl;
  text-align: right;
  border-radius: 16px !important;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.2) !important;
  background: white !important;
}

/* ========================================
   إصلاح لون النصوص في نافذة إضافة مديون جديد
   - تغيير لون النص من الأبيض إلى الأسود الداكن
   - تحسين وضوح جميع النصوص في النافذة المنبثقة
   - إصلاح التسميات والحقول والقوائم المنسدلة
   ======================================== */

/* إصلاح عنوان النافذة المنبثقة */
.v-dialog .v-card-title {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 16px 16px 0 0;
  border-bottom: 1px solid #e2e8f0;
  font-weight: 600 !important;
  color: #1a1a1a !important;
  text-align: center !important;
}

/* إصلاح النصوص في النافذة المنبثقة */
.v-dialog .v-card-text,
.v-dialog .v-card-text *,
.v-dialog .v-form,
.v-dialog .v-form * {
  color: #1a1a1a !important;
}

/* إصلاح التسميات في النافذة المنبثقة */
.v-dialog .v-label,
.v-dialog .v-label *,
.v-dialog .v-field__label,
.v-dialog .v-field__label * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في الحقول */
.v-dialog .v-text-field .v-field__input,
.v-dialog .v-textarea .v-field__input,
.v-dialog .v-select .v-field__input,
.v-dialog .v-field__input,
.v-dialog .v-field__input * {
  color: #1a1a1a !important;
}

/* إصلاح خلفية الحقول */
.v-dialog .v-text-field .v-field,
.v-dialog .v-textarea .v-field,
.v-dialog .v-select .v-field {
  background: white !important;
  border: 1px solid #e2e8f0 !important;
}

/* إصلاح القوائم المنسدلة في النافذة المنبثقة */
.v-dialog .v-menu__content,
.v-dialog .v-menu__content *,
.v-dialog .v-overlay__content,
.v-dialog .v-overlay__content * {
  color: #1a1a1a !important;
  background: white !important;
}

.v-dialog .v-list,
.v-dialog .v-list *,
.v-dialog .v-list-item,
.v-dialog .v-list-item *,
.v-dialog .v-list-item-title,
.v-dialog .v-list-item-subtitle {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في جميع عناصر النافذة */
.v-dialog .v-card-text span,
.v-dialog .v-card-text div,
.v-dialog .v-card-text p,
.v-dialog .v-card-text label {
  color: #1a1a1a !important;
}

/* إصلاح رسائل التحقق */
.v-dialog .v-messages,
.v-dialog .v-messages *,
.v-dialog .v-message {
  color: #1a1a1a !important;
}

/* إصلاح شامل لجميع النصوص في النافذة المنبثقة */
.v-dialog *,
.v-dialog * *,
.v-dialog .v-card *,
.v-dialog .v-card * * {
  color: #1a1a1a !important;
}

/* إصلاح خاص للقوائم المنسدلة */
.v-dialog .v-select__menu,
.v-dialog .v-select__menu *,
.v-dialog .v-select__menu .v-list-item,
.v-dialog .v-select__menu .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

/* إصلاح النصوص في جميع الحقول */
.v-dialog .v-input,
.v-dialog .v-input *,
.v-dialog .v-text-field,
.v-dialog .v-text-field *,
.v-dialog .v-textarea,
.v-dialog .v-textarea *,
.v-dialog .v-select,
.v-dialog .v-select * {
  color: #1a1a1a !important;
}

/* إصلاح خلفية النافذة المنبثقة */
.v-dialog .v-card-text {
  background: white !important;
}

/* إصلاح نهائي شامل لجميع النصوص في النافذة المنبثقة */
.v-dialog .v-application,
.v-dialog .v-application *,
.v-dialog .v-application .v-card,
.v-dialog .v-application .v-card * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع العناصر */
.v-dialog .v-card-title span,
.v-dialog .v-card-title *,
.v-dialog .v-card-actions,
.v-dialog .v-card-actions * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في الأزرار (استثناء) */
.v-dialog .v-btn {
  color: white !important;
}

.v-dialog .v-btn .v-btn__content {
  color: white !important;
}

/* إصلاح أزرار النافذة */
.v-dialog .v-card-actions .v-btn {
  border-radius: 8px !important;
  text-transform: none !important;
  font-weight: 500 !important;
}

.v-dialog .v-card-actions .v-btn--variant-elevated {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15) !important;
}

/* إصلاح شامل لجميع النصوص في النافذة المنبثقة */
.v-dialog .v-card-text .v-row,
.v-dialog .v-card-text .v-row *,
.v-dialog .v-card-text .v-col,
.v-dialog .v-card-text .v-col * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع الحقول المحددة */
.v-dialog .v-text-field[data-field="name"] *,
.v-dialog .v-text-field[data-field="email"] *,
.v-dialog .v-text-field[data-field="phone"] *,
.v-dialog .v-text-field[data-field="amount"] *,
.v-dialog .v-select[data-field="currency"] *,
.v-dialog .v-text-field[data-field="dueDate"] *,
.v-dialog .v-textarea[data-field="notes"] * {
  color: #1a1a1a !important;
}

/* إصلاح النصوص في جميع العناصر */
.v-dialog .v-card-text .v-input,
.v-dialog .v-card-text .v-input *,
.v-dialog .v-card-text .v-field,
.v-dialog .v-card-text .v-field * {
  color: #1a1a1a !important;
}

.v-tabs {
  direction: rtl;
}

.v-tab {
  direction: rtl;
  text-align: right;
}
</style>