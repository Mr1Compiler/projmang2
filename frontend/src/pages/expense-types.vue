<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-format-list-bulleted-type</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">أنواع المصروفات</h1>
              <p class="subtitle">إدارة وتصنيف جميع أنواع المصروفات في النظام</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Statistics Cards -->
      <v-row class="mb-6">
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-format-list-bulleted-type</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ totalTypes }}</h3>
                <p class="stat-label">إجمالي الأنواع</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ activeTypes }}</h3>
                <p class="stat-label">الأنواع النشطة</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-pause-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ inactiveTypes }}</h3>
                <p class="stat-label">الأنواع المعطلة</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" md="3">
          <v-card class="modern-stat-card stat-card-info" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-currency-usd</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ totalExpensesAmount }}</h3>
                <p class="stat-label">إجمالي المصروفات</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Expense Types Table -->
      <v-row>
        <v-col cols="12">
          <v-card class="types-table-card" elevation="3">
            <v-card-title class="table-title">
              <div class="d-flex align-center">
                <v-icon left color="white" class="me-2" size="18">mdi-table</v-icon>
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
                class="add-button btn-glow light-sweep smooth-transition"
                @click.stop="addExpenseType"
                elevation="2"
                color="primary"
                size="small"
                type="button"
                style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important;"
              >
                <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
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
                      v-if="canUpdate"
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
                      v-if="canUpdate"
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
                      v-if="canDelete"
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
      <v-dialog v-model="typeDialog" max-width="900" persistent>
        <v-card class="task-dialog-card profile-form-card">
          <v-card-title class="task-dialog-header profile-form-header">
            <h2 class="profile-form-title">{{ isEditing ? 'تعديل نوع المصروف' : 'معلومات نوع المصروف' }}</h2>
          </v-card-title>
          
          <v-card-text class="profile-form-content">
            <p class="profile-form-instruction">
              لإتمام {{ isEditing ? 'تعديل' : 'إضافة' }} نوع المصروف، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
            </p>
            
            <v-form ref="typeFormRef" v-model="typeFormValid">
              
              <v-row class="profile-form-row">
                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      اسم نوع المصروف <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="typeForm.name"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم نوع المصروف مطلوب']"
                      required
                      density="comfortable"
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
                
                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      الأيقونة <span class="required-star">*</span>
                    </label>
                    <v-select
                      v-model="typeForm.icon"
                      :items="iconOptions"
                      item-title="label"
                      item-value="value"
                      variant="outlined"
                      :rules="[v => !!v || 'الأيقونة مطلوبة']"
                      required
                      density="comfortable"
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
                
                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      اللون <span class="required-star">*</span>
                    </label>
                    <v-select
                      v-model="typeForm.color"
                      :items="colorOptions"
                      item-title="label"
                      item-value="value"
                      variant="outlined"
                      :rules="[v => !!v || 'اللون مطلوب']"
                      required
                      density="comfortable"
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
              </v-row>
              
              <v-row class="profile-form-row">
                <v-col cols="12" md="6" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      الحالة <span class="required-star">*</span>
                    </label>
                    <v-select
                      v-model="typeForm.status"
                      :items="statusOptions"
                      item-title="label"
                      item-value="value"
                      variant="outlined"
                      :rules="[v => !!v || 'الحالة مطلوبة']"
                      required
                      density="comfortable"
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
              </v-row>
              
              <v-row class="profile-form-row">
                <v-col cols="12" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      الوصف <span class="required-star">*</span>
                    </label>
                    <v-textarea
                      v-model="typeForm.description"
                      variant="outlined"
                      rows="4"
                      :rules="[v => !!v || 'الوصف مطلوب']"
                      required
                      density="comfortable"
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
          
          <v-card-actions class="profile-form-actions">
            <v-spacer />
            <v-btn
              class="profile-form-cancel-btn"
              variant="outlined"
              @click="closeTypeDialog"
            >
              إلغاء
            </v-btn>
            <v-btn
              class="profile-form-continue-btn"
              variant="elevated"
              :disabled="!typeFormValid"
              @click="saveType"
            >
              {{ isEditing ? 'تحديث' : 'حفظ' }}
              <v-icon class="ms-2">mdi-arrow-left</v-icon>
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { usePermissions } from '@/composables/usePermissions'

// Permissions
const { canCreate, canUpdate, canDelete } = usePermissions('/expense-types')

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
  console.log('عرض تفاصيل نوع المصروف:', type)
  alert(`تفاصيل نوع المصروف:\n\nالاسم: ${type.name}\nالوصف: ${type.description}\nالحالة: ${type.status === 'active' ? 'نشط' : 'معطل'}\nإجمالي المصروفات: ${formatAmount(type.totalExpenses)}`)
}

const toggleTypeStatus = (type) => {
  type.status = type.status === 'active' ? 'inactive' : 'active'
  console.log(`تم ${type.status === 'active' ? 'تفعيل' : 'تعطيل'} نوع المصروف:`, type.name)
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
    console.log('تم حفظ نوع المصروف بنجاح')
  }
}

// تهيئة المكون
onMounted(() => {
  console.log('تم تحميل صفحة أنواع المصروفات')
})
</script>

<style scoped>
/* نفس التنسيقات من صفحة expenses.vue */
.data-page {
  background: #ffffff !important;
  min-height: 100vh;
  padding: 0;
  overflow-x: hidden;
}

.fullscreen-content {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 20px;
}

/* Header Styles - نفس تنسيق صفحة المهندسين */
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

/* Modern Statistics Cards - نفس تصميم الصفحة الرئيسية */
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

/* دعم الكروتات القديمة */
.stat-card {
  border-radius: 12px;
  transition: all 0.3s ease;
  border: 2px solid #e2e8f0;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  position: relative;
  overflow: hidden;
}

.stat-card::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: linear-gradient(45deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  transform: rotate(45deg);
  transition: all 0.6s ease;
  opacity: 0;
}

.stat-card:hover::after {
  animation: shimmer 1.5s ease-in-out;
  opacity: 1;
}

@keyframes shimmer {
  0% {
    transform: translateX(-100%) translateY(-100%) rotate(45deg);
  }
  100% {
    transform: translateX(100%) translateY(100%) rotate(45deg);
  }
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: #3b82f6;
  border-radius: 8px 8px 0 0;
}

/* Icon Effects */
.stat-card .v-icon {
  transition: all 0.4s ease;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.stat-card:hover .v-icon {
  transform: scale(1.15) rotate(5deg);
  filter: drop-shadow(0 6px 12px rgba(0, 0, 0, 0.3));
}

.total-types:hover .v-icon {
  animation: bounce 0.6s ease-in-out;
}

.active-types:hover .v-icon {
  animation: pulse 1s ease-in-out infinite;
}

.inactive-types:hover .v-icon {
  animation: shake 0.5s ease-in-out;
}

.total-expenses:hover .v-icon {
  animation: spin 1s linear infinite;
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0) scale(1.15);
  }
  40% {
    transform: translateY(-15px) scale(1.15);
  }
  60% {
    transform: translateY(-8px) scale(1.15);
  }
}

@keyframes pulse {
  0% {
    transform: scale(1.15);
  }
  50% {
    transform: scale(1.3);
  }
  100% {
    transform: scale(1.15);
  }
}

@keyframes shake {
  0%, 100% {
    transform: translateX(0) scale(1.15);
  }
  25% {
    transform: translateX(-8px) scale(1.15);
  }
  75% {
    transform: translateX(8px) scale(1.15);
  }
}

@keyframes spin {
  from {
    transform: rotate(0deg) scale(1.15);
  }
  to {
    transform: rotate(360deg) scale(1.15);
  }
}

.stat-card:hover {
  transform: translateY(-4px) scale(1.02);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

/* Floating particles effect */
.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: #3b82f6;
  border-radius: 8px 8px 0 0;
  transition: all 0.3s ease;
}

.stat-card:hover::before {
  height: 5px;
  background: linear-gradient(90deg, #3b82f6, #1d4ed8, #3b82f6);
  background-size: 200% 100%;
  animation: gradientShift 2s ease-in-out infinite;
}

@keyframes gradientShift {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.total-types {
  border-color: #3b82f6;
  color: #3b82f6;
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
}

.total-types .stat-number {
  color: #1e40af;
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.total-types:hover {
  box-shadow: 0 6px 12px rgba(59, 130, 246, 0.3);
  transform: translateY(-2px);
}

.active-types {
  border-color: #10b981;
  color: #10b981;
  background: linear-gradient(135deg, #d1fae5 0%, #a7f3d0 100%);
}

.active-types .stat-number {
  color: #047857;
  background: linear-gradient(135deg, #047857 0%, #10b981 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.active-types:hover {
  box-shadow: 0 6px 12px rgba(16, 185, 129, 0.3);
  transform: translateY(-2px);
}

.inactive-types {
  border-color: #f59e0b;
  color: #f59e0b;
  background: linear-gradient(135deg, #fef3c7 0%, #fde68a 100%);
}

.inactive-types .stat-number {
  color: #d97706;
  background: linear-gradient(135deg, #d97706 0%, #f59e0b 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.inactive-types:hover {
  box-shadow: 0 6px 12px rgba(245, 158, 11, 0.3);
  transform: translateY(-2px);
}

.total-expenses {
  border-color: #8b5cf6;
  color: #8b5cf6;
  background: linear-gradient(135deg, #e9d5ff 0%, #ddd6fe 100%);
}

.total-expenses .stat-number {
  color: #7c3aed;
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.total-expenses:hover {
  box-shadow: 0 6px 12px rgba(139, 92, 246, 0.3);
  transform: translateY(-2px);
}

.stat-number {
  font-size: 2.5rem;
  font-weight: 800;
  margin: 16px 0 12px 0;
  color: #ffffff;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3);
  background: linear-gradient(135deg, #1e293b 0%, #475569 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  transition: all 0.4s ease;
  position: relative;
}

.stat-card:hover .stat-number {
  transform: scale(1.05);
  text-shadow: 0 5px 10px rgba(0, 0, 0, 0.4);
}

.stat-number::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  border-radius: 8px;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.stat-card:hover .stat-number::before {
  opacity: 1;
  animation: numberGlow 2s ease-in-out infinite;
}

@keyframes numberGlow {
  0%, 100% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
}

.stat-label {
  font-size: 1rem;
  font-weight: 700;
  color: #475569;
  margin: 0;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  transition: all 0.3s ease;
  position: relative;
}

.stat-card:hover .stat-label {
  color: #1e293b;
  transform: translateY(-2px);
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.stat-label::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 50%;
  width: 0;
  height: 2px;
  background: linear-gradient(90deg, transparent, #3b82f6, transparent);
  transition: all 0.3s ease;
  transform: translateX(-50%);
}

.stat-card:hover .stat-label::after {
  width: 100%;
}

/* Add Button */
.add-type-btn {
  background: #3b82f6;
  color: white;
  font-weight: 600;
  font-size: 1rem;
  padding: 10px 20px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(59, 130, 246, 0.2);
  text-transform: none;
  transition: all 0.2s ease;
}

.add-type-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(59, 130, 246, 0.3);
}

/* Table Styles */
.types-table-card {
  border-radius: 8px;
  border: 1px solid #dee0e2;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  background: white;
  overflow: hidden;
}

.types-table-card .v-card-text {
  margin-top: 1rem !important;
}

.table-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  font-weight: 600;
  font-size: 0.95rem !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  padding: 8px 12px !important;
  display: flex !important;
  align-items: center !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
  border-radius: 8px 8px 0 0 !important;
  min-height: 36px !important;
  margin-bottom: 1rem !important;
}

.table-title .title-text {
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.add-type-btn-header,
.add-type-btn-header.v-btn,
.v-btn.add-type-btn-header {
  background: rgba(255, 255, 255, 0.95) !important;
  backdrop-filter: blur(10px) !important;
  color: #2563eb !important;
  border-radius: 12px !important;
  font-weight: 700 !important;
  font-size: 0.9rem !important;
  text-transform: none !important;
  padding: 10px 20px !important;
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.1),
    0 2px 8px rgba(37, 99, 235, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.9) !important;
  border: 2px solid rgba(255, 255, 255, 0.8) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  position: relative;
  overflow: hidden;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.add-type-btn-header :deep(.v-btn__content) {
  color: #2563eb !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 6px !important;
}

.add-type-btn-header :deep(.v-icon) {
  color: #2563eb !important;
  font-size: 18px !important;
}

.add-type-btn-header:hover {
  transform: translateY(-2px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(0, 0, 0, 0.15),
    0 4px 12px rgba(37, 99, 235, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(37, 99, 235, 0.3) !important;
}

.search-field {
  max-width: 300px;
}

.table-title .search-field :deep(.v-field__input),
.table-title .search-field :deep(.v-field__input input) {
  background: rgba(255, 255, 255, 0.95) !important;
  color: #000000 !important;
  font-weight: 600 !important;
}

.table-title .search-field :deep(.v-label),
.table-title .search-field :deep(.v-field-label) {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 700 !important;
}

.table-title .search-field :deep(.v-field-label--active),
.table-title .search-field :deep(.v-field-label--floating) {
  color: rgba(255, 255, 255, 0.95) !important;
}

.table-title .search-field :deep(.v-field__outline) {
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.table-title .search-field :deep(.v-field--focused .v-field__outline) {
  border-color: rgba(255, 255, 255, 0.8) !important;
}

.table-title .search-field :deep(.v-field__prepend-inner .v-icon) {
  color: rgba(255, 255, 255, 0.9) !important;
}

.types-table {
  border-radius: 12px;
  overflow: hidden;
}

.types-table :deep(.v-data-table-header),
.types-table :deep(.v-data-table__thead),
.types-table :deep(.v-data-table-header th),
.types-table :deep(.v-data-table__th) {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.65rem !important;
  text-align: center !important;
  padding: 6px 4px !important;
  min-height: 32px !important;
  line-height: 1.2 !important;
  letter-spacing: 0.2px !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1) !important;
  -webkit-text-fill-color: #ffffff !important;
}

.types-table .v-data-table__th {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.65rem !important;
  text-align: center !important;
  padding: 6px 4px !important;
  min-height: 32px !important;
  line-height: 1.2 !important;
  letter-spacing: 0.2px !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1) !important;
  -webkit-text-fill-color: #ffffff !important;
}

.types-table .v-data-table__td {
  background: white;
  color: #334155;
  font-weight: 500;
  padding: 12px 8px;
  border-bottom: 1px solid #f1f5f9;
  text-align: center;
  transition: all 0.2s ease;
}

.types-table .v-data-table__wrapper tbody tr:nth-child(even) .v-data-table__td {
  background: #f8fafc;
}

.types-table .v-data-table__wrapper tbody tr:hover .v-data-table__td {
  background: #e0f2fe;
  transform: none;
  box-shadow: none;
}

/* Custom Text Styles */
.serial-number {
  font-weight: 600;
  color: #64748b;
  background: #f1f5f9;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.9rem;
  border: 1px solid #e2e8f0;
}

.type-name {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 8px;
}

.name-text {
  font-weight: 600;
  color: #1e293b;
  font-size: 1rem;
}

.description-text {
  font-weight: 500;
  color: #64748b;
  font-size: 0.9rem;
}

.amount-text {
  font-weight: 600;
  color: #059669;
  font-size: 1rem;
  background: #ecfdf5;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #a7f3d0;
}

.status-chip {
  font-weight: 600;
  font-size: 0.85rem;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  gap: 4px;
  align-items: center;
  justify-content: center;
}

.action-btn {
  min-width: 32px;
  height: 32px;
  border-radius: 6px;
  transition: all 0.3s ease;
}

.details-btn {
  color: #3b82f6;
  background: rgba(59, 130, 246, 0.1);
}

.details-btn:hover {
  background: rgba(59, 130, 246, 0.2);
  transform: scale(1.05);
}

.edit-btn {
  color: #10b981;
  background: rgba(16, 185, 129, 0.1);
}

.edit-btn:hover {
  background: rgba(16, 185, 129, 0.2);
  transform: scale(1.05);
}

.toggle-btn {
  color: #f59e0b;
  background: rgba(245, 158, 11, 0.1);
}

.toggle-btn:hover {
  background: rgba(245, 158, 11, 0.2);
  transform: scale(1.05);
}

.delete-btn {
  color: #ef4444;
  background: rgba(239, 68, 68, 0.1);
}

.delete-btn:hover {
  background: rgba(239, 68, 68, 0.2);
  transform: scale(1.05);
}

/* Dialog Styles */
.type-dialog {
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  border: 3px solid rgba(5, 62, 168, 0.3);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(255, 255, 255, 0.9) 100%);
  backdrop-filter: blur(20px);
}

.dialog-header {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%);
  color: rgb(220, 210, 210);
  font-weight: 900;
  font-size: 1.5rem;
  padding: 28px 32px;
  border-bottom: 4px solid #1e3a8a;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4);
  letter-spacing: 0.8px;
  position: relative;
}

.dialog-header::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, transparent 0%, rgba(255, 255, 255, 0.3) 50%, transparent 100%);
}

.close-btn {
  color: white;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 8px;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.05);
}

.dialog-content {
  padding: 32px;
  background: linear-gradient(135deg, #2706fc 0%, #4d93d9 100%);
  border-radius: 0 0 16px 16px;
  position: relative;
}

.dialog-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent 0%, #e2e8f0 50%, transparent 100%);
}

.form-field {
  margin-bottom: 16px;
}

.form-field .v-field {
  border-radius: 12px;
  background: white;
  box-shadow: 0 3px 12px rgba(0, 0, 0, 0.08);
  border: 2px solid #e2e8f0;
  transition: all 0.3s ease;
}

.form-field .v-field:hover {
  border-color: #cbd5e1;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.12);
}

.form-field .v-field__outline {
  border: none;
  border-radius: 12px;
}

.form-field .v-field--focused .v-field__outline {
  border: none;
}

.form-field .v-field--focused {
  border-color: #3b82f6;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.15);
  transform: translateY(-1px);
}

.form-field .v-label {
  color: #374151;
  font-weight: 700;
  font-size: 1rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.form-field .v-field__input {
  color: #1f2937;
  font-weight: 600;
  font-size: 1.1rem;
  padding: 12px 16px;
}

.form-field .v-field__input::placeholder {
  color: #9ca3af;
  font-weight: 500;
}

/* Dropdown Menu Styles */
.form-field .v-list {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 2px solid #e2e8f0 !important;
  border-radius: 12px !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15) !important;
  backdrop-filter: blur(10px) !important;
  max-height: 300px !important;
  overflow-y: auto !important;
}

.form-field .v-list-item {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 16px 20px !important;
  border-bottom: 1px solid #f1f5f9 !important;
  transition: all 0.3s ease !important;
  position: relative !important;
}

.form-field .v-list-item:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  transform: translateX(8px) !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  border-left: 4px solid #1e40af !important;
}

.form-field .v-list-item--active {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 4px 12px rgba(30, 64, 175, 0.4) !important;
  border-left: 4px solid #1e3a8a !important;
}

.form-field .v-list-item:last-child {
  border-bottom: none !important;
}

/* Dropdown Arrow */
.form-field .v-field__append-inner .v-icon {
  color: #6b7280 !important;
  font-size: 1.2rem !important;
  transition: all 0.3s ease !important;
}

.form-field .v-field--focused .v-field__append-inner .v-icon {
  color: #3b82f6 !important;
  transform: rotate(180deg) !important;
}

/* Selected Item Display */
.form-field .v-field__input {
  color: #1f2937 !important;
  font-weight: 600 !important;
  font-size: 1.1rem !important;
  padding: 12px 16px !important;
}

/* Dropdown Scrollbar */
.form-field .v-list::-webkit-scrollbar {
  width: 6px !important;
}

.form-field .v-list::-webkit-scrollbar-track {
  background: #f1f5f9 !important;
  border-radius: 3px !important;
}

.form-field .v-list::-webkit-scrollbar-thumb {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  border-radius: 3px !important;
}

.form-field .v-list::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
}

/* Enhanced Dropdown Visibility */
.form-field .v-overlay__content {
  z-index: 9999 !important;
  border-radius: 12px !important;
  overflow: hidden !important;
}

.form-field .v-menu__content {
  border-radius: 12px !important;
  overflow: hidden !important;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2) !important;
}

/* Dropdown Item Icons */
.form-field .v-list-item__prepend .v-icon {
  color: inherit !important;
  font-size: 1.2rem !important;
  margin-left: 8px !important;
}

/* Dropdown Item Text */
.form-field .v-list-item__content {
  font-weight: 600 !important;
  font-size: 1rem !important;
}

/* Focus States */
.form-field .v-field--focused .v-list {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.1), 0 8px 25px rgba(0, 0, 0, 0.15) !important;
}

/* Animation for Dropdown Opening */
.form-field .v-list {
  animation: dropdownSlide 0.3s ease-out !important;
}

@keyframes dropdownSlide {
  from {
    opacity: 0;
    transform: translateY(-10px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.dialog-actions {
  padding: 24px 32px;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-top: 2px solid #e2e8f0;
  border-radius: 0 0 16px 16px;
  position: relative;
}

.dialog-actions::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent 0%, #cbd5e1 50%, transparent 100%);
}

.cancel-btn,
.cancel-btn.v-btn,
.v-btn.cancel-btn {
  background: rgba(255, 255, 255, 0.95) !important;
  backdrop-filter: blur(10px) !important;
  color: #64748b !important;
  font-weight: 700 !important;
  font-size: 0.95rem !important;
  padding: 12px 24px !important;
  border-radius: 12px !important;
  margin-left: 12px !important;
  border: 2px solid rgba(100, 116, 139, 0.3) !important;
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.08),
    0 2px 8px rgba(100, 116, 139, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.9) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  text-transform: none !important;
  position: relative;
  overflow: hidden;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.cancel-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(100, 116, 139, 0.1), transparent);
  transition: left 0.6s ease;
  z-index: 1;
}

.cancel-btn:hover::before {
  left: 100%;
}

.cancel-btn:hover,
.cancel-btn.v-btn:hover,
.v-btn.cancel-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(100, 116, 139, 0.5) !important;
  transform: translateY(-2px) scale(1.02) !important;
  box-shadow: 
    0 6px 20px rgba(0, 0, 0, 0.12),
    0 3px 10px rgba(100, 116, 139, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
  color: #475569 !important;
}

.cancel-btn:active {
  transform: translateY(0px) scale(0.98) !important;
  box-shadow: 
    0 2px 8px rgba(0, 0, 0, 0.1),
    0 1px 4px rgba(100, 116, 139, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.9) !important;
}

.cancel-btn :deep(.v-btn__content) {
  color: #64748b !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  position: relative;
  z-index: 2;
}

.cancel-btn:hover :deep(.v-btn__content) {
  color: #475569 !important;
}

.save-btn,
.save-btn.v-btn,
.v-btn.save-btn {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 50%, #3b82f6 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  font-size: 0.95rem !important;
  padding: 12px 24px !important;
  border-radius: 12px !important;
  box-shadow: 
    0 4px 16px rgba(30, 64, 175, 0.3),
    0 2px 8px rgba(37, 99, 235, 0.2),
    0 0 20px rgba(59, 130, 246, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  position: relative;
  overflow: hidden;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-transform: none !important;
  animation: saveButtonGlow 2s ease-in-out infinite !important;
}

@keyframes saveButtonGlow {
  0%, 100% {
    box-shadow: 
      0 4px 16px rgba(30, 64, 175, 0.3),
      0 2px 8px rgba(37, 99, 235, 0.2),
      0 0 20px rgba(59, 130, 246, 0.4),
      inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  }
  50% {
    box-shadow: 
      0 4px 20px rgba(30, 64, 175, 0.5),
      0 2px 12px rgba(37, 99, 235, 0.4),
      0 0 30px rgba(59, 130, 246, 0.6),
      inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  }
}

.save-btn::before {
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

.save-btn:hover::before {
  left: 100%;
}

.save-btn:hover,
.save-btn.v-btn:hover,
.v-btn.save-btn:hover {
  transform: translateY(-3px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(30, 64, 175, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    0 0 40px rgba(59, 130, 246, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #2563eb 100%) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
  animation: saveButtonGlowHover 1.5s ease-in-out infinite !important;
}

@keyframes saveButtonGlowHover {
  0%, 100% {
    box-shadow: 
      0 8px 24px rgba(30, 64, 175, 0.5),
      0 4px 12px rgba(37, 99, 235, 0.4),
      0 0 40px rgba(59, 130, 246, 0.7),
      inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  }
  50% {
    box-shadow: 
      0 8px 28px rgba(30, 64, 175, 0.6),
      0 4px 16px rgba(37, 99, 235, 0.5),
      0 0 50px rgba(59, 130, 246, 0.8),
      inset 0 1px 0 rgba(255, 255, 255, 0.5) !important;
  }
}

.save-btn:active {
  transform: translateY(-1px) scale(1.01) !important;
  box-shadow: 
    0 4px 12px rgba(30, 64, 175, 0.4),
    0 2px 6px rgba(37, 99, 235, 0.3),
    0 0 25px rgba(59, 130, 246, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  animation: saveButtonClick 0.3s ease !important;
}

@keyframes saveButtonClick {
  0% {
    transform: translateY(-1px) scale(1.01);
  }
  50% {
    transform: translateY(-1px) scale(0.98);
  }
  100% {
    transform: translateY(-1px) scale(1.01);
  }
}

.save-btn :deep(.v-btn__content) {
  color: white !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 8px !important;
  position: relative;
  z-index: 2;
}

.save-btn :deep(.v-icon) {
  color: white !important;
  transition: transform 0.3s ease !important;
  position: relative;
  z-index: 2;
}

.save-btn:hover :deep(.v-icon) {
  transform: rotate(15deg) scale(1.1) !important;
}

.save-btn:disabled {
  background: #9ca3af;
  border-color: #6b7280;
  box-shadow: none;
  transform: none;
}

.save-btn:disabled {
  background: #e2e8f0;
  color: #94a3b8;
  box-shadow: none;
  transform: none;
}

/* Animations */
@keyframes star-twinkle {
  0%, 100% { transform: scale(1) rotate(0deg); }
  50% { transform: scale(1.1) rotate(5deg); }
}

@keyframes text-glow {
  0% { text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3); }
  100% { text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3), 0 0 20px rgba(255, 255, 255, 0.3); }
}

@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.fade-in {
  animation: fade-in 0.6s ease-out;
}

/* Responsive Design */
@media (max-width: 768px) {
  .data-page {
    padding: 10px;
  }
  
  .page-title {
    font-size: 2rem;
  }
  
  .stat-number {
    font-size: 2rem;
  }
  
  .action-buttons {
    flex-direction: column;
    gap: 2px;
  }
  
  .action-btn {
    min-width: 28px;
    height: 28px;
  }
}

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

/* تنسيقات نموذج نوع المصروف الجديد */
.profile-form-card {
  border-radius: 8px !important;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.12) !important;
  overflow: visible !important;
  background: #ffffff !important;
  border: 2px solid #d1d5db !important;
}

.profile-form-header {
  background: #ffffff !important;
  color: #1e293b !important;
  padding: 20px 24px !important;
  border-bottom: 1px solid #e2e8f0 !important;
  box-shadow: none !important;
}

.profile-form-title {
  font-size: 1.35rem !important;
  font-weight: 700 !important;
  color: #1e293b !important;
  margin: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: none !important;
}

.profile-form-content {
  padding: 24px !important;
  background: #ffffff !important;
  border: 2px solid #d1d5db !important;
  border-radius: 8px !important;
  margin: 20px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}

.profile-form-instruction {
  font-size: 0.85rem !important;
  color: #64748b !important;
  margin-bottom: 24px !important;
  line-height: 1.6 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.profile-form-row {
  margin: 0 !important;
}

.profile-form-column {
  padding: 0 10px !important;
  display: flex !important;
  flex-direction: column !important;
}

.profile-form-column:first-child {
  padding-right: 0 !important;
}

.profile-form-column:last-child {
  padding-left: 0 !important;
}

.profile-form-field-wrapper {
  margin-bottom: 16px !important;
  padding: 12px !important;
  background: #ffffff !important;
  border: 1.5px solid #d1d5db !important;
  border-radius: 6px !important;
  transition: all 0.2s ease !important;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05) !important;
  flex: 1 !important;
}

.profile-form-field-wrapper:hover {
  border-color: #9ca3af !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08) !important;
}

.profile-form-field-wrapper:focus-within {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15), 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.profile-form-label {
  display: block !important;
  font-size: 0.8125rem !important;
  font-weight: 600 !important;
  color: #374151 !important;
  margin-bottom: 6px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-align: right !important;
}

.required-star {
  color: #ef4444 !important;
  font-weight: 700 !important;
  margin-right: 2px !important;
}

.profile-form-input {
  width: 100% !important;
}

.profile-form-input :deep(.v-field) {
  background: #ffffff !important;
  border-radius: 4px !important;
}

.profile-form-input :deep(.v-field__outline) {
  border-color: #d1d5db !important;
  border-width: 1.5px !important;
}

.profile-form-input :deep(.v-field--focused .v-field__outline) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.profile-form-input :deep(.v-field__input) {
  color: #1e293b !important;
  font-size: 0.875rem !important;
  padding: 10px 12px !important;
  text-align: right !important;
  direction: rtl !important;
}

.profile-form-input :deep(.v-label),
.profile-form-input :deep(.v-field-label) {
  color: #6b7280 !important;
  font-size: 0.875rem !important;
  text-align: right !important;
  right: 0 !important;
  left: auto !important;
}

.profile-form-input :deep(.v-select__selection) {
  text-align: right !important;
  direction: rtl !important;
  align-items: center !important;
  min-height: 40px !important;
  font-weight: 500 !important;
}

.profile-form-input :deep(.v-overlay__content) {
  border-radius: 10px !important;
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.18), 0 4px 10px rgba(148, 163, 184, 0.25) !important;
  border: 1px solid #e5e7eb !important;
  overflow: hidden !important;
}

.profile-form-input :deep(.v-overlay__content .v-list-item-title) {
  font-size: 0.85rem !important;
  color: #0f172a !important;
  text-align: right !important;
}

.profile-form-input :deep(.v-overlay__content .v-list-item--active),
.profile-form-input :deep(.v-overlay__content .v-list-item:hover) {
  background: #eff6ff !important;
  color: #1d4ed8 !important;
}

.profile-form-actions {
  padding: 24px 32px !important;
  background: #ffffff !important;
  border-top: 1px solid #e2e8f0 !important;
  display: flex !important;
  justify-content: flex-end !important;
  gap: 12px !important;
  visibility: visible !important;
  opacity: 1 !important;
  z-index: 10 !important;
  position: relative !important;
}

.profile-form-actions .v-btn {
  display: inline-flex !important;
  visibility: visible !important;
  opacity: 1 !important;
  min-height: 44px !important;
  height: auto !important;
}

.profile-form-cancel-btn {
  min-width: 120px !important;
  height: 44px !important;
  border-radius: 6px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  color: #6b7280 !important;
  border-color: #d1d5db !important;
  background: #ffffff !important;
  visibility: visible !important;
  opacity: 1 !important;
  display: inline-flex !important;
  z-index: 10 !important;
  position: relative !important;
  border-width: 1.5px !important;
  border-style: solid !important;
}

.profile-form-cancel-btn .v-btn__content {
  visibility: visible !important;
  opacity: 1 !important;
}

.profile-form-cancel-btn:hover {
  background: #f9fafb !important;
  border-color: #9ca3af !important;
}

.profile-form-continue-btn {
  min-width: 140px !important;
  height: 44px !important;
  border-radius: 6px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  background: #10b981 !important;
  color: #ffffff !important;
  box-shadow: 0 2px 8px rgba(16, 185, 129, 0.3) !important;
  visibility: visible !important;
  opacity: 1 !important;
  display: inline-flex !important;
  z-index: 10 !important;
  position: relative !important;
  border: none !important;
}

.profile-form-continue-btn .v-btn__content {
  visibility: visible !important;
  opacity: 1 !important;
  display: flex !important;
  align-items: center !important;
  gap: 8px !important;
}

.profile-form-continue-btn span {
  visibility: visible !important;
  opacity: 1 !important;
  display: inline-block !important;
}

.profile-form-continue-btn:hover {
  background: #059669 !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4) !important;
  transform: translateY(-1px) !important;
}

.profile-form-continue-btn:disabled {
  background: #d1d5db !important;
  color: #9ca3af !important;
  box-shadow: none !important;
}

.profile-form-continue-btn .v-icon {
  color: #ffffff !important;
}

</style>
