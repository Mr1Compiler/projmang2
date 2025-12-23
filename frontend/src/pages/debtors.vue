<template>
  <div class="fill-height data-page">
      <v-container fluid class="pa-6" style="padding: 0 20px !important;">
        <!-- رأس الصفحة المحسن -->
        <div class="engineers-header-card">
          <div class="header-gradient-line"></div>
          <div class="header-content">
            <div class="header-right">
              <div class="engineer-emoji">
                <v-icon size="40" color="white">mdi-account-cash</v-icon>
              </div>
              <div class="header-text">
                <h1 class="main-title">إدارة المديونون</h1>
                <p class="subtitle">إدارة حسابات المديونون والمستحقات المالية</p>
              </div>
            </div>
          </div>
        </div>


        <!-- إحصائيات سريعة محسنة -->
        <v-row class="mb-8">
          <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-info pa-6 text-center" elevation="2">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-account-group</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ debtors.length }}</h3>
                  <p class="stat-label">إجمالي المديونون</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-primary pa-6 text-center" elevation="2">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-currency-usd</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ formatCurrency(totalDebt) }}</h3>
                  <p class="stat-label">إجمالي المديونية</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-warning pa-6 text-center" elevation="2">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-clock-alert</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ overdueCount }}</h3>
                  <p class="stat-label">متأخرين</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-success pa-6 text-center" elevation="2">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ paidCount }}</h3>
                  <p class="stat-label">مدفوع</p>
                </div>
              </div>
            </v-card>
          </v-col>
        </v-row>

        <!-- شريط البحث والفلترة المحسن -->
        <v-card class="search-filter-card mb-8" elevation="2">
          <v-card-title class="filters-header-new">
            <div class="d-flex align-center">
              <v-icon size="18" color="white" class="me-2" style="color: #ffffff !important;">mdi-filter</v-icon>
              <span class="filters-header-title">البحث والفلترة</span>
            </div>
          </v-card-title>
          <v-card-text class="filters-content">
            <v-row no-gutters>
              <v-col cols="12" md="5" class="px-2">
                <v-text-field
                  v-model="searchQuery"
                  label="البحث عن مديون"
                  prepend-inner-icon="mdi-magnify"
                  variant="outlined"
                  clearable
                  hide-details
                  density="compact"
                  class="search-field-new"
                />
              </v-col>
              <v-col cols="12" md="2" class="px-2">
                <v-select
                  v-model="statusFilter"
                  label="حالة الدفع"
                  :items="statusOptions"
                  variant="outlined"
                  hide-details
                  clearable
                  density="compact"
                  class="filter-field-new"
                />
              </v-col>
              <v-col cols="12" md="2" class="px-2">
                <v-select
                  v-model="amountFilter"
                  label="نطاق المبلغ"
                  :items="amountOptions"
                  variant="outlined"
                  hide-details
                  clearable
                  density="compact"
                  class="filter-field-new"
                />
              </v-col>
              <v-col cols="12" md="3" class="d-flex justify-end px-2">
                <v-btn
                  color="primary"
                  variant="outlined"
                  size="small"
                  class="reset-button-new"
                  @click="resetFilters"
                  style="height: 36px !important; font-size: 0.875rem !important;"
                >
                  <v-icon class="me-1" size="18">mdi-refresh</v-icon>
                  إعادة تعيين
                </v-btn>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>

        <!-- جدول المديونون المحسن -->
        <v-card class="data-table-card" elevation="2">
          <v-card-title class="table-title-header d-flex align-center justify-space-between">
            <div class="d-flex align-center">
              <v-icon class="me-2" color="white" size="18" style="color: #ffffff !important;">mdi-table</v-icon>
              <span class="title-text">قائمة المديونون</span>
              <v-chip class="ms-3 count-chip" size="x-small" variant="elevated" style="background: rgba(255, 255, 255, 0.2) !important; color: #ffffff !important; font-size: 0.75rem !important; height: 20px !important;">
                {{ filteredDebtors.length }} عنصر
              </v-chip>
            </div>
            <div class="d-flex align-center gap-2">
              <v-btn
                class="add-button btn-glow light-sweep smooth-transition"
                @click="openAddDialog"
                elevation="2"
                color="primary"
                size="small"
                style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
              >
                <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
                إضافة مديون جديد
              </v-btn>
              <v-btn
                icon="mdi-refresh"
                variant="text"
                color="white"
                size="small"
                @click="refreshData"
                class="action-button"
                style="color: #ffffff !important; min-width: 32px !important; width: 32px !important; height: 32px !important;"
              />
              <v-btn
                icon="mdi-download"
                variant="text"
                color="white"
                size="small"
                @click="exportData"
                class="action-button"
                style="color: #ffffff !important; min-width: 32px !important; width: 32px !important; height: 32px !important;"
              />
            </div>
          </v-card-title>
          <v-data-table
            :headers="headers"
            :items="filteredDebtors"
            :loading="loading"
            :items-per-page="-1"
            hide-default-footer
            class="elevation-0 debtors-data-table"
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
            <template v-slot:item.totalDebt="{ item }">
              <div class="text-right">
                <div class="font-weight-bold text-h6">{{ formatCurrency(item.totalDebt) }}</div>
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

          <!-- Pagination -->
          <div class="d-flex justify-center pa-4" v-if="totalPages > 0">
            <v-pagination
              v-model="currentPage"
              :length="totalPages"
              :total-visible="7"
              @update:model-value="onPageChange"
              rounded="circle"
              density="comfortable"
              active-color="primary"
            />
          </div>
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
                  label="الاسم الكامل *"
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
                  :rules="[v => !v || /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صالح']"
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
                  v-model.number="debtorForm.totalDebt"
                  label="المبلغ المطلوب *"
                  type="number"
                  variant="outlined"
                  :rules="[v => v > 0 || 'المبلغ مطلوب ويجب أن يكون أكبر من صفر']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="debtorForm.currency"
                  label="العملة *"
                  :items="currencyOptions"
                  variant="outlined"
                  :rules="[v => !!v || 'العملة مطلوبة']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-menu
                  v-model="dueDateMenu"
                  :close-on-content-click="false"
                  location="bottom"
                >
                  <template v-slot:activator="{ props }">
                    <v-text-field
                      v-model="formattedDueDate"
                      label="تاريخ الاستحقاق"
                      readonly
                      v-bind="props"
                      variant="outlined"
                      prepend-inner-icon="mdi-calendar"
                    />
                  </template>
                  <v-date-picker
                    v-model="selectedDueDate"
                    @update:model-value="onDueDateSelected"
                    color="primary"
                  />
                </v-menu>
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
                {{ formatCurrency(selectedDebtor.totalDebt) }}
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
                  <h3 class="text-h5 font-weight-bold">{{ formatCurrency(selectedDebtor.totalDebt) }}</h3>
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
                            <span class="font-weight-bold text-error">{{ formatCurrency(selectedDebtor.totalDebt) }}</span>
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
import { listDebtors, createDebtor, updateDebtor, deleteDebtor as deleteDebtorApi, getDebtorStats } from '@/api/debtors'
import { DEFAULT_LIMIT } from '@/constants/pagination'

// البيانات التفاعلية
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
const editingDebtorId = ref(null)

// نموذج المدين - aligned with backend DTO
const debtorForm = ref({
  name: '',
  email: '',
  phone: '',
  totalDebt: 0,
  currency: 'IQD',
  dueDate: '',
  status: 'pending',
  notes: ''
})

// قائمة المدينين - from backend
const debtors = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(DEFAULT_LIMIT)
const totalItems = ref(0)
const totalPages = ref(0)

const debtorStats = ref({
  total: 0,
  active: 0,
  paid: 0,
  totalDebt: 0,
  activeDebt: 0,
  averageDebt: 0
})

// عناوين الجدول - aligned with backend DTO
const headers = [
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'المبلغ المطلوب', key: 'totalDebt', sortable: true },
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

// خيارات الفلترة - متطابقة مع قيم الـ backend
const statusOptions = [
  { title: 'جميع الحالات', value: '' },
  { title: 'نشط', value: 'active' },
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

// Date picker state
const dueDateMenu = ref(false)
const selectedDueDate = ref(null)

// Computed property for formatted due date display
const formattedDueDate = computed(() => {
  if (!debtorForm.value.dueDate) return ''
  const date = new Date(debtorForm.value.dueDate)
  return date.toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })
})

// Handle due date selection from picker
const onDueDateSelected = (date) => {
  if (date) {
    const d = new Date(date)
    debtorForm.value.dueDate = d.toISOString().split('T')[0]
  }
  dueDateMenu.value = false
}

// الحسابات
const filteredDebtors = computed(() => {
  let filtered = debtors.value

  // فلترة بالبحث
  if (searchQuery.value) {
    filtered = filtered.filter(debtor =>
      debtor.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      (debtor.email && debtor.email.toLowerCase().includes(searchQuery.value.toLowerCase()))
    )
  }

  // فلترة بالحالة
  if (statusFilter.value) {
    filtered = filtered.filter(debtor => debtor.status === statusFilter.value)
  }

  // فلترة بالمبلغ
  if (amountFilter.value) {
    filtered = filtered.filter(debtor => {
      const amount = debtor.totalDebt
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
  return debtorStats.value.totalDebt || 0
})

const overdueCount = computed(() => {
  // Calculate overdue based on due date (past dates for unpaid debtors)
  const today = new Date()
  return debtors.value.filter(debtor => {
    if (debtor.status === 'paid') return false
    if (!debtor.dueDate) return false
    return new Date(debtor.dueDate) < today
  }).length
})

const paidCount = computed(() => {
  return debtorStats.value.paid || 0
})

// حسابات الديون والتسديدات
const totalPaid = computed(() => {
  if (!selectedDebtor.value || !selectedDebtor.value.payments) return 0
  return selectedDebtor.value.payments.reduce((sum, payment) => sum + payment.amount, 0)
})

const remainingAmount = computed(() => {
  if (!selectedDebtor.value) return 0
  return (selectedDebtor.value.totalDebt || 0) - totalPaid.value
})

const paymentPercentage = computed(() => {
  if (!selectedDebtor.value || !selectedDebtor.value.totalDebt) return 0
  return Math.round((totalPaid.value / selectedDebtor.value.totalDebt) * 100)
})

// الدوال
const formatCurrency = (amount) => {
  if (amount == null) return '0 IQD'
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount) + ' IQD'
}

const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

const getStatusColor = (status) => {
  const colors = {
    'active': 'info',
    'pending': 'warning',
    'paid': 'success'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    'active': 'نشط',
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

// Load debtors from backend
const loadDebtors = async (page = currentPage.value) => {
  loading.value = true
  try {
    // Load debtors list and stats in parallel
    const [response, stats] = await Promise.all([
      listDebtors({ page, limit: itemsPerPage.value }),
      getDebtorStats()
    ])
    console.log('Debtors data received:', response)
    console.log('Debtors stats received:', stats)

    // Update pagination state
    debtors.value = response.data
    totalItems.value = response.total
    totalPages.value = response.totalPages
    currentPage.value = response.page

    if (stats) {
      debtorStats.value = stats
    }
  } catch (err) {
    console.error('Error loading debtors:', err)
  } finally {
    loading.value = false
  }
}

// Handle page change
const onPageChange = (page) => {
  currentPage.value = page
  loadDebtors(page)
}

const openAddDialog = () => {
  isEdit.value = false
  editingDebtorId.value = null
  debtorForm.value = {
    name: '',
    email: '',
    phone: '',
    totalDebt: 0,
    currency: 'IQD',
    dueDate: '',
    status: 'pending',
    notes: ''
  }
  dialog.value = true
}

const editDebtor = (debtor) => {
  isEdit.value = true
  editingDebtorId.value = debtor.id
  // Format dueDate to YYYY-MM-DD for input field
  const formattedDueDate = debtor.dueDate ? new Date(debtor.dueDate).toISOString().split('T')[0] : ''
  debtorForm.value = {
    name: debtor.name,
    email: debtor.email || '',
    phone: debtor.phone || '',
    totalDebt: debtor.totalDebt,
    currency: debtor.currency,
    dueDate: formattedDueDate,
    status: debtor.status || 'pending',
    notes: debtor.notes || ''
  }
  dialog.value = true
}

const viewDebtor = (debtor) => {
  selectedDebtor.value = debtor
  viewDialog.value = true
}

const saveDebtor = async () => {
  try {
    const payload = {
      name: debtorForm.value.name,
      totalDebt: Number(debtorForm.value.totalDebt),
      currency: debtorForm.value.currency
    }
    // Optional fields
    if (debtorForm.value.email && debtorForm.value.email.trim()) {
      payload.email = debtorForm.value.email.trim()
    }
    if (debtorForm.value.phone && debtorForm.value.phone.trim()) {
      payload.phone = debtorForm.value.phone.trim()
    }
    if (debtorForm.value.dueDate) {
      payload.dueDate = new Date(debtorForm.value.dueDate).toISOString()
    }
    if (debtorForm.value.status) {
      payload.status = debtorForm.value.status
    }
    if (debtorForm.value.notes && debtorForm.value.notes.trim()) {
      payload.notes = debtorForm.value.notes.trim()
    }

    if (isEdit.value && editingDebtorId.value) {
      await updateDebtor(editingDebtorId.value, payload)
    } else {
      await createDebtor(payload)
    }
    await loadDebtors()
    closeDialog()
  } catch (err) {
    console.error('Error saving debtor:', err)
  }
}

const closeDialog = () => {
  dialog.value = false
  valid.value = false
  editingDebtorId.value = null
}

const markAsPaid = async (debtor) => {
  try {
    await updateDebtor(debtor.id, { status: 'paid' })
    await loadDebtors()
  } catch (err) {
    console.error('Error marking as paid:', err)
  }
}

const deleteDebtor = async (debtor) => {
  if (!confirm('هل أنت متأكد من حذف هذا المدين؟')) return
  try {
    await deleteDebtorApi(debtor.id)
    await loadDebtors()
  } catch (err) {
    console.error('Error deleting debtor:', err)
  }
}

const refreshData = async () => {
  await loadDebtors()
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

// Load data on mount
onMounted(() => {
  loadDebtors()
})
</script>

<style scoped>
/* ========================================
   تحسين ألوان صفحة المديونون
   ======================================== */

/* خلفية الصفحة */
.data-page {
  background: #ffffff !important;
  color: var(--text-dark);
  min-height: 100vh;
  padding: 0 !important;
  max-width: 100% !important;
  margin: 0 !important;
  width: 100% !important;
  overflow-x: hidden;
}

/* ========================================
   رأس الصفحة - نفس تنسيق صفحة المهندسين
   ======================================== */

.engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  border-radius: 0;
  width: 100vw;
  max-width: 100vw;
  box-shadow: 0 8px 32px rgba(25, 118, 210, 0.3);
  position: relative;
  overflow: hidden;
  margin-left: calc(-50vw + 50%);
  margin-right: calc(-50vw + 50%);
  margin-bottom: 1.5rem;
  border: none;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: slideInFromTop 1s ease-out, shimmer 3s ease-in-out infinite;
}

.engineers-header-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: sweep 2s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: diagonalShimmer 4s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 60px rgba(25, 118, 210, 0.5);
  animation: hoverPulse 0.6s ease-in-out;
}

.engineers-header-card:hover::before {
  animation: sweep 1s ease-in-out infinite;
}

.engineers-header-card:hover::after {
  animation: diagonalShimmer 2s ease-in-out infinite;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.3);
  animation: gradientFlow 3s ease-in-out infinite;
  z-index: 2;
}

.header-content {
  padding: 12px 16px !important;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  min-height: auto !important;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.1) 0%, rgba(21, 101, 192, 0.05) 100%);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 3;
  animation: fadeInUp 1.2s ease-out 0.3s both;
  max-width: calc(100vw - 320px);
  margin: 0 auto;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1.8rem;
  text-align: right;
  padding: 0.8rem 1.5rem;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.15);
}

.engineer-emoji {
  position: relative;
  animation: slideInFromRight 1s ease-out 0.9s both, float 3s ease-in-out infinite 2s, pulse 2s ease-in-out infinite 2s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 0.5rem;
}

.engineer-emoji .v-icon {
  filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3));
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #ffffff, #e3f2fd);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  position: relative;
  animation: iconGlow 2s ease-in-out infinite 3s, iconBounce 3s ease-in-out infinite 3s;
}

.engineer-emoji .v-icon:first-child {
  animation: iconGlow 2s ease-in-out infinite 3s, iconBounce 3s ease-in-out infinite 3s;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.main-title {
  color: white !important;
  font-size: 1.2rem !important;
  font-weight: bold !important;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 0.75rem !important;
  margin: 0;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
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

/* Modern Statistics Cards */
.modern-stat-card {
  position: relative !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  cursor: pointer !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  height: 100% !important;
  min-height: 140px !important;
  background: #ffffff !important;
}

.modern-stat-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.stat-card-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  opacity: 0.1;
  transition: opacity 0.3s ease;
}

.modern-stat-card:hover .stat-card-background {
  opacity: 0.2;
}

.stat-card-primary .stat-card-background {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.stat-card-success .stat-card-background {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.stat-card-warning .stat-card-background {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.stat-card-info .stat-card-background {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.stat-card-primary {
  background: #ffffff !important;
  border: 2px solid #3b82f6 !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15) !important;
}

.stat-card-success {
  background: #ffffff !important;
  border: 2px solid #10b981 !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.15) !important;
}

.stat-card-warning {
  background: #ffffff !important;
  border: 2px solid #f59e0b !important;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.15) !important;
}

.stat-card-info {
  background: #ffffff !important;
  border: 2px solid #06b6d4 !important;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.15) !important;
}

.stat-card-content {
  position: relative;
  z-index: 2;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  align-items: center;
  text-align: center;
}

.stat-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 64px;
  height: 64px;
  min-width: 64px;
  min-height: 64px;
  border-radius: 50%;
  margin-bottom: 0.25rem;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  position: relative;
  overflow: visible;
}

.stat-icon-wrapper::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #ffffff;
  z-index: 1;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.stat-card-primary .stat-icon-wrapper {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.stat-card-primary .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.2) 0%, rgba(37, 99, 235, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-success .stat-icon-wrapper {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.stat-card-success .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.2) 0%, rgba(5, 150, 105, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-warning .stat-icon-wrapper {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.stat-card-warning .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 158, 11, 0.2) 0%, rgba(217, 119, 6, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-info .stat-icon-wrapper {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.stat-card-info .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.2) 0%, rgba(8, 145, 178, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
  min-height: 32px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.stat-icon-wrapper :deep(.v-icon) {
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
  min-height: 32px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 !important;
  padding: 0 !important;
}

.stat-icon-wrapper :deep(svg) {
  width: 32px !important;
  height: 32px !important;
  font-size: 32px !important;
}

.stat-card-primary .stat-icon {
  color: #3b82f6 !important;
}

.stat-card-primary .stat-icon-wrapper :deep(.v-icon) {
  color: #3b82f6 !important;
}

.stat-card-success .stat-icon {
  color: #10b981 !important;
}

.stat-card-success .stat-icon-wrapper :deep(.v-icon) {
  color: #10b981 !important;
}

.stat-card-warning .stat-icon {
  color: #f59e0b !important;
}

.stat-card-warning .stat-icon-wrapper :deep(.v-icon) {
  color: #f59e0b !important;
}

.stat-card-info .stat-icon {
  color: #06b6d4 !important;
}

.stat-card-info .stat-icon-wrapper :deep(.v-icon) {
  color: #06b6d4 !important;
}

.check-icon,
.stat-icon-wrapper .check-icon,
.stat-icon-wrapper :deep(.check-icon) {
  transform: scaleX(-1) !important;
}

.stat-info {
  flex: 1;
  text-align: center;
  width: 100%;
}

.stat-value {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 0.5rem;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  text-align: center;
  font-variant-numeric: tabular-nums;
  unicode-bidi: embed;
  color: #000000 !important;
}

.stat-label {
  font-size: 1rem;
  font-weight: 500;
  text-align: center;
  color: #64748b;
}

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
  background: #ffffff !important;
  border: 1px solid #e2e8f0 !important;
  border-radius: 12px !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
  overflow: hidden !important;
}

/* شريط عنوان الفلترة */
.filters-header-new {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  border-radius: 12px 12px 0 0 !important;
  padding: 0.5rem 0.75rem !important;
  border-bottom: 1px solid #1e40af !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.3) !important;
}

.filters-header-new,
.filters-header-new *,
.filters-header-new div,
.filters-header-new span {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.filters-header-new::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, #93c5fd, #60a5fa, #3b82f6);
  border-radius: 12px 12px 0 0;
}

.filters-header-title {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  letter-spacing: 0.2px !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

.filters-header-new :deep(.v-icon) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
}

/* محتوى الفلترة */
.filters-content {
  padding: 0.5rem 0.75rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 50%, #e2e8f0 100%) !important;
  border-radius: 0 0 12px 12px !important;
}

/* إزالة الأنماط القديمة - تم استبدالها بالأنماط الجديدة أعلاه */

/* حقول البحث والفلترة الجديدة */
.search-field-new :deep(.v-field),
.filter-field-new :deep(.v-field) {
  background: #ffffff !important;
  border-radius: 8px !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease !important;
}

.search-field-new :deep(.v-field:hover),
.filter-field-new :deep(.v-field:hover) {
  border-color: #2563eb !important;
  box-shadow: 0 2px 6px rgba(59, 130, 246, 0.2) !important;
}

.search-field-new :deep(.v-field--focused),
.filter-field-new :deep(.v-field--focused) {
  border-color: #2563eb !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.search-field-new :deep(.v-field__input),
.filter-field-new :deep(.v-field__input) {
  color: #1e293b !important;
  font-size: 0.875rem !important;
  padding: 8px 12px !important;
}

.search-field-new :deep(.v-label),
.filter-field-new :deep(.v-label) {
  color: #3b82f6 !important;
  font-weight: 500 !important;
  font-size: 0.875rem !important;
}

.search-field-new :deep(.v-field--focused .v-label),
.filter-field-new :deep(.v-field--focused .v-label) {
  color: #2563eb !important;
}

.search-field-new :deep(.v-field__prepend-inner .v-icon),
.filter-field-new :deep(.v-field__prepend-inner .v-icon) {
  color: #3b82f6 !important;
}

/* زر إعادة التعيين الجديد */
.reset-button-new {
  border-radius: 8px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  color: #3b82f6 !important;
  border-color: #3b82f6 !important;
  background: #ffffff !important;
  transition: all 0.3s ease !important;
}

.reset-button-new:hover {
  background: #3b82f6 !important;
  color: #ffffff !important;
  border-color: #2563eb !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.3) !important;
}

.reset-button-new :deep(.v-icon) {
  color: inherit !important;
}

/* القوائم المنسدلة */
.search-filter-card :deep(.v-overlay__content) {
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  border: 2px solid #3b82f6 !important;
  overflow: hidden !important;
  background: #ffffff !important;
}

/* ========================================
   جدول البيانات
   ======================================== */

/* تم نقل أنماط .data-table-card إلى common-components.css */

/* شريط عنوان الجدول */
.table-title-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  padding: 8px 12px !important;
  border-radius: 8px 8px 0 0 !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  min-height: 36px !important;
  margin-bottom: 1rem !important;
}

.table-title-header,
.table-title-header *,
.table-title-header div,
.table-title-header span {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.table-title-header .title-text {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.table-title-header :deep(.v-icon) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
}

.table-title-header .count-chip {
  background: rgba(255, 255, 255, 0.2) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.table-title-header .action-button {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.table-title-header .action-button :deep(.v-icon) {
  color: #ffffff !important;
  fill: #ffffff !important;
}

.table-title-header .action-button:hover {
  background: rgba(255, 255, 255, 0.1) !important;
  transform: scale(1.05) !important;
}

.action-button {
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

.action-button:hover {
  background: rgba(255, 255, 255, 0.1) !important;
  transform: scale(1.05) !important;
}

/* مسافة بين شريط العنوان والجدول */
.debtors-data-table {
  margin-top: 1rem !important;
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
/* ========================================
   Animations - نفس صفحة المهندسين
   ======================================== */

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-12px) rotate(2deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

@keyframes shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
}

@keyframes slideInFromTop {
  0% {
    transform: translateY(-100px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes slideInFromLeft {
  0% {
    transform: translateX(-50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes slideInFromRight {
  0% {
    transform: translateX(50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes fadeInUp {
  0% {
    transform: translateY(30px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes sweep {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes diagonalShimmer {
  0%, 100% {
    transform: translateX(-100%) translateY(-100%) rotate(45deg);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translateX(100%) translateY(100%) rotate(45deg);
    opacity: 0;
  }
}

@keyframes gradientFlow {
  0%, 100% {
    background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  }
  50% {
    background: linear-gradient(90deg, #bbdefb 0%, #ffffff 50%, #e3f2fd 100%);
  }
}

@keyframes hoverPulse {
  0% {
    transform: translateY(-8px) scale(1.02);
  }
  50% {
    transform: translateY(-12px) scale(1.05);
  }
  100% {
    transform: translateY(-8px) scale(1.02);
  }
}

@keyframes iconGlow {
  0%, 100% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 20px rgba(255, 255, 255, 0.3));
  }
  50% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 30px rgba(255, 255, 255, 0.6));
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-8px);
  }
}

</style>