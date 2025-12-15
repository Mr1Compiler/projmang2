<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">📋</span>
        <h1 class="page-title text-glow fade-in">أنواع المصروفات</h1>
        <p class="page-subtitle fade-in">إدارة وتصنيف جميع أنواع المصروفات في النظام</p>
      </div>

      <!-- Statistics Cards -->
      <v-row class="mb-6">
        <v-col cols="12" md="3">
          <v-card class="stat-card total-types" elevation="3">
            <v-card-text class="text-center">
              <v-icon size="80" color="primary">mdi-format-list-bulleted-type</v-icon>
              <h3 class="stat-number">{{ totalTypes }}</h3>
              <p class="stat-label">إجمالي الأنواع</p>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="stat-card active-types" elevation="3">
            <v-card-text class="text-center">
              <v-icon size="80" color="success">mdi-check-circle</v-icon>
              <h3 class="stat-number">{{ activeTypes }}</h3>
              <p class="stat-label">الأنواع النشطة</p>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="stat-card inactive-types" elevation="3">
            <v-card-text class="text-center">
              <v-icon size="80" color="warning">mdi-pause-circle</v-icon>
              <h3 class="stat-number">{{ inactiveTypes }}</h3>
              <p class="stat-label">الأنواع المعطلة</p>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="stat-card total-expenses" elevation="3">
            <v-card-text class="text-center">
              <v-icon size="80" color="info">mdi-currency-usd</v-icon>
              <h3 class="stat-number">{{ totalExpensesAmount }}</h3>
              <p class="stat-label">إجمالي المصروفات</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Expense Types Table -->
      <v-row>
        <v-col cols="12">
          <v-card class="types-table-card" elevation="3">
            <v-card-title class="table-title">
              <div class="d-flex align-center">
                <v-icon left color="white" class="me-2">mdi-table</v-icon>
                <span class="title-text">أنواع المصروفات</span>
              </div>
              <v-spacer></v-spacer>
              <v-text-field
                v-model="searchQuery"
                prepend-inner-icon="mdi-magnify"
                label="البحث في الأنواع"
                variant="outlined"
                density="compact"
                hide-details
                class="search-field me-3"
              ></v-text-field>
              <v-btn
                color="white"
                variant="elevated"
                @click="addExpenseType"
                size="default"
                class="add-type-btn-header"
              >
                <v-icon left class="me-2">mdi-plus</v-icon>
                إضافة نوع مصروف جديد
              </v-btn>
            </v-card-title>
            <v-card-text>
              <v-data-table
                :headers="typeHeaders"
                :items="expenseTypes"
                :search="searchQuery"
                class="types-table"
                :items-per-page="10"
                :loading="false"
                hover
                no-data-text="لا توجد أنواع مصروفات"
              >
                <!-- Serial Number Column -->
                <template #item.serial="{ index }">
                  <span class="serial-number">{{ index + 1 }}</span>
                </template>

                <!-- Type Name Column -->
                <template #item.name="{ item }">
                  <div class="type-name">
                    <v-icon :color="item.color" class="me-2">{{ item.icon }}</v-icon>
                    <span class="name-text">{{ item.name }}</span>
                  </div>
                </template>

                <!-- Description Column -->
                <template #item.description="{ item }">
                  <span class="description-text">{{ item.description }}</span>
                </template>

                <!-- Status Column -->
                <template #item.status="{ item }">
                  <v-chip
                    :color="item.status === 'active' ? 'success' : 'warning'"
                    size="small"
                    class="status-chip"
                  >
                    {{ item.status === 'active' ? 'نشط' : 'معطل' }}
                  </v-chip>
                </template>

                <!-- Total Expenses Column -->
                <template #item.totalExpenses="{ item }">
                  <span class="amount-text">{{ formatAmount(item.totalExpenses) }}</span>
                </template>

                <!-- Actions Column -->
                <template #item.actions="{ item }">
                  <div class="action-buttons">
                    <v-btn
                      size="small"
                      color="primary"
                      variant="text"
                      @click="viewTypeDetails(item)"
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
                      @click="editType(item)"
                      icon
                      class="action-btn edit-btn"
                      title="تعديل"
                    >
                      <v-icon size="16">mdi-pencil</v-icon>
                    </v-btn>
                    <v-btn
                      size="small"
                      :color="item.status === 'active' ? 'warning' : 'success'"
                      variant="text"
                      @click="toggleTypeStatus(item)"
                      icon
                      class="action-btn toggle-btn"
                      :title="item.status === 'active' ? 'تعطيل' : 'تفعيل'"
                    >
                      <v-icon size="16">{{ item.status === 'active' ? 'mdi-pause' : 'mdi-play' }}</v-icon>
                    </v-btn>
                    <v-btn
                      size="small"
                      color="error"
                      variant="text"
                      @click="deleteType(item)"
                      icon
                      class="action-btn delete-btn"
                      title="حذف"
                    >
                      <v-icon size="16">mdi-delete</v-icon>
                    </v-btn>
                  </div>
                </template>
              </v-data-table>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Add/Edit Expense Type Dialog -->
      <v-dialog v-model="typeDialog" max-width="600" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">{{ isEditing ? 'mdi-pencil' : 'mdi-plus' }}</v-icon>
                <span class="header-title">{{ isEditing ? 'تعديل نوع المصروف' : 'إضافة نوع مصروف جديد' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeTypeDialog"
                class="close-btn"
              />
            </div>
          </div>

          <div class="dialog-body">
            <v-form ref="typeFormRef" v-model="typeFormValid">
              <div class="form-fields">
                <v-row>
                  <!-- Type Name -->
                  <v-col cols="12">
                    <v-text-field
                      v-model="typeForm.name"
                      label="اسم نوع المصروف"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم نوع المصروف مطلوب']"
                      required
                      class="form-field"
                    ></v-text-field>
                  </v-col>

                  <!-- Description -->
                  <v-col cols="12">
                    <v-textarea
                      v-model="typeForm.description"
                      label="الوصف"
                      variant="outlined"
                      rows="3"
                      :rules="[v => !!v || 'الوصف مطلوب']"
                      required
                      class="form-field"
                    ></v-textarea>
                  </v-col>

                  <!-- Icon -->
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="typeForm.icon"
                      :items="iconOptions"
                      item-title="label"
                      item-value="value"
                      label="الأيقونة"
                      variant="outlined"
                      :rules="[v => !!v || 'الأيقونة مطلوبة']"
                      required
                      class="form-field"
                    ></v-select>
                  </v-col>

                  <!-- Color -->
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="typeForm.color"
                      :items="colorOptions"
                      item-title="label"
                      item-value="value"
                      label="اللون"
                      variant="outlined"
                      :rules="[v => !!v || 'اللون مطلوب']"
                      required
                      class="form-field"
                    ></v-select>
                  </v-col>

                  <!-- Status -->
                  <v-col cols="12">
                    <v-select
                      v-model="typeForm.status"
                      :items="statusOptions"
                      item-title="label"
                      item-value="value"
                      label="الحالة"
                      variant="outlined"
                      :rules="[v => !!v || 'الحالة مطلوبة']"
                      required
                      class="form-field"
                    ></v-select>
                  </v-col>
                </v-row>
              </div>
            </v-form>
          </div>

          <div class="dialog-actions">
            <v-spacer></v-spacer>
            <v-btn
              color="grey"
              variant="text"
              @click="closeTypeDialog"
              class="cancel-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveType"
              :disabled="!typeFormValid"
              class="save-btn"
            >
              <v-icon left>mdi-content-save</v-icon>
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
import { toast } from 'vue3-toastify'

// متغيرات الحالة الأساسية
const loading = ref(false)
const typeDialog = ref(false)
const typeFormValid = ref(false)
const typeFormRef = ref(null)
const isEditing = ref(false)
const searchQuery = ref('')

// بيانات أنواع المصروفات
const expenseTypes = ref([
  {
    id: 1,
    name: 'معدات',
    description: 'معدات وأجهزة حاسوبية ومكتبية',
    icon: 'mdi-desktop-classic',
    color: 'blue',
    status: 'active',
    totalExpenses: 150000,
    createdAt: '2024-01-15'
  },
  {
    id: 2,
    name: 'مواد',
    description: 'مواد بناء وأدوات ومستلزمات',
    icon: 'mdi-hammer-screwdriver',
    color: 'orange',
    status: 'active',
    totalExpenses: 250000,
    createdAt: '2024-01-20'
  },
  {
    id: 3,
    name: 'أجور',
    description: 'أجور العمال والموظفين',
    icon: 'mdi-account-cash',
    color: 'green',
    status: 'active',
    totalExpenses: 500000,
    createdAt: '2024-01-25'
  },
  {
    id: 4,
    name: 'نقل',
    description: 'تكاليف النقل والمواصلات',
    icon: 'mdi-truck',
    color: 'purple',
    status: 'active',
    totalExpenses: 75000,
    createdAt: '2024-02-01'
  },
  {
    id: 5,
    name: 'مرافق',
    description: 'فواتير الكهرباء والماء والغاز',
    icon: 'mdi-lightning-bolt',
    color: 'red',
    status: 'active',
    totalExpenses: 120000,
    createdAt: '2024-02-05'
  },
  {
    id: 6,
    name: 'أخرى',
    description: 'مصروفات أخرى متنوعة',
    icon: 'mdi-dots-horizontal',
    color: 'grey',
    status: 'inactive',
    totalExpenses: 30000,
    createdAt: '2024-02-10'
  }
])

// عناوين الجدول
const typeHeaders = ref([
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'اسم النوع', key: 'name', sortable: true, align: 'right' },
  { title: 'الوصف', key: 'description', sortable: true, align: 'right' },
  { title: 'الحالة', key: 'status', sortable: true, align: 'center' },
  { title: 'إجمالي المصروفات', key: 'totalExpenses', sortable: true, align: 'center' },
  { title: 'الإجراءات', key: 'actions', sortable: false, align: 'center' }
])

// خيارات النموذج
const iconOptions = ref([
  { label: 'معدات', value: 'mdi-desktop-classic' },
  { label: 'مواد', value: 'mdi-hammer-screwdriver' },
  { label: 'أجور', value: 'mdi-account-cash' },
  { label: 'نقل', value: 'mdi-truck' },
  { label: 'مرافق', value: 'mdi-lightning-bolt' },
  { label: 'أخرى', value: 'mdi-dots-horizontal' }
])

const colorOptions = ref([
  { label: 'أزرق', value: 'blue' },
  { label: 'برتقالي', value: 'orange' },
  { label: 'أخضر', value: 'green' },
  { label: 'بنفسجي', value: 'purple' },
  { label: 'أحمر', value: 'red' },
  { label: 'رمادي', value: 'grey' }
])

const statusOptions = ref([
  { label: 'نشط', value: 'active' },
  { label: 'معطل', value: 'inactive' }
])

// نموذج البيانات
const typeForm = ref({
  name: '',
  description: '',
  icon: '',
  color: '',
  status: 'active'
})

// الحسابات
const totalTypes = computed(() => expenseTypes.value.length)
const activeTypes = computed(() => expenseTypes.value.filter(type => type.status === 'active').length)
const inactiveTypes = computed(() => expenseTypes.value.filter(type => type.status === 'inactive').length)
const totalExpensesAmount = computed(() => {
  return expenseTypes.value
    .reduce((sum, type) => sum + type.totalExpenses, 0)
    .toLocaleString()
})

// الدوال المساعدة
const formatAmount = (amount) => {
  return `${amount.toLocaleString()} د.ع`
}

// دوال الإجراءات
const addExpenseType = () => {
  isEditing.value = false
  resetTypeForm()
  typeDialog.value = true
}

const editType = (type) => {
  isEditing.value = true
  typeForm.value = { ...type }
  typeDialog.value = true
}

const viewTypeDetails = (type) => {
  toast.info(`${type.name}: ${type.description} - ${type.status === 'active' ? 'نشط' : 'معطل'} - ${formatAmount(type.totalExpenses)}`)
}

const toggleTypeStatus = (type) => {
  type.status = type.status === 'active' ? 'inactive' : 'active'
}

const deleteType = (type) => {
  if (confirm(`هل أنت متأكد من حذف نوع المصروف "${type.name}"؟`)) {
    const index = expenseTypes.value.findIndex(t => t.id === type.id)
    if (index > -1) {
      expenseTypes.value.splice(index, 1)
    }
  }
}

const closeTypeDialog = () => {
  typeDialog.value = false
  resetTypeForm()
}

const resetTypeForm = () => {
  typeForm.value = {
    name: '',
    description: '',
    icon: '',
    color: '',
    status: 'active'
  }
  typeFormValid.value = false
  if (typeFormRef.value) {
    typeFormRef.value.reset()
  }
}

const saveType = () => {
  if (typeFormValid.value) {
    if (isEditing.value) {
      // تحديث نوع موجود
      const index = expenseTypes.value.findIndex(t => t.id === typeForm.value.id)
      if (index > -1) {
        expenseTypes.value[index] = { ...typeForm.value }
      }
    } else {
      // إضافة نوع جديد
      const newType = {
        id: Date.now(),
        ...typeForm.value,
        totalExpenses: 0,
        createdAt: new Date().toISOString().split('T')[0]
      }
      expenseTypes.value.unshift(newType)
    }

    closeTypeDialog()
  }
}

// تهيئة المكون
onMounted(() => {
})
</script>


<style>
/* Import page styles */
@import './styles/expense-types.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
