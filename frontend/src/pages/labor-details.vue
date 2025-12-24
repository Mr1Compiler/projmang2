<template>
  <div class="labor-details-page">
    <!-- Header Section -->
    <div class="page-header glass-effect gradient-animation">
      <div class="header-content">
        <div class="header-text">
          <h1 class="page-title text-glow fade-in">تفاصيل الأيدي العاملة</h1>
          <p class="page-subtitle fade-in">إدارة العمال والموظفين المشاركين في المشروع</p>
        </div>
        <v-btn 
          icon="mdi-arrow-right" 
          @click="goBack" 
          class="back-btn"
          size="small"
        >
          <v-icon size="20" color="white">mdi-arrow-right</v-icon>
        </v-btn>
      </div>
    </div>

    <!-- Content Container -->
    <div class="content-container">

    <!-- Labor Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="success">mdi-account-group</v-icon>
        الأيدي العاملة
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="laborSearch"
              label="البحث في العمال..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchLabor"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchLabor"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="success" 
              variant="elevated" 
              class="add-btn"
              @click="openAddDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة عامل
            </v-btn>
          </div>
        </div>
      </v-card-text>
    </v-card>

    <!-- Labor Table Card -->
    <v-card class="data-table-card" elevation="2">
      <v-card-title class="table-title indigo-title">
        <span class="title-text">قائمة العمال</span>
      </v-card-title>

      <v-data-table
        v-model:page="currentPage"
        :headers="laborHeaders"
        :items="laborData"
        :search="laborSearch"
        :items-per-page="DEFAULT_LIMIT"
        class="project-table"
        no-data-text="لا توجد بيانات متاحة"
        loading-text="جاري التحميل..."
        hover
        hide-default-footer
      >
        <!-- Serial Number Column -->
        <template #item.id="{ index }">
          <span class="serial-number">{{ index + 1 }}</span>
        </template>

        <!-- Worker Name Column -->
        <template #item.workerName="{ item }">
          <span class="project-name">{{ item.workerName }}</span>
        </template>

        <!-- Job Title Column -->
        <template #item.jobTitle="{ item }">
          <span class="project-name">{{ item.jobTitle || '-' }}</span>
        </template>

        <!-- Quantity Column -->
        <template #item.quantity="{ item }">
          <span class="date-text">{{ item.quantity }}</span>
        </template>

        <!-- Cost Column -->
        <template #item.cost="{ item }">
          <span class="cost-text">{{ formatCurrency(item.cost) }}</span>
        </template>

        <!-- Total Column -->
        <template #item.total="{ item }">
          <span class="cost-text">{{ formatCurrency(item.total) }}</span>
        </template>

        <!-- Phone Column -->
        <template #item.phone="{ item }">
          <span class="date-text">{{ item.phone || '-' }}</span>
        </template>

        <!-- Address Column -->
        <template #item.address="{ item }">
          <span class="project-name">{{ item.address || '-' }}</span>
        </template>

        <!-- Actions Column -->
        <template #item.actions="{ item }">
          <div class="action-buttons">
            <v-btn
              icon="mdi-pencil"
              size="small"
              color="primary"
              variant="text"
              @click="openEditDialog(item)"
              title="تعديل العامل"
              class="action-btn"
            />
            <v-btn
              icon="mdi-delete"
              size="small"
              color="error"
              variant="text"
              @click="deleteLabor(item)"
              title="حذف العامل"
              class="action-btn"
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
          rounded="circle"
          density="comfortable"
          active-color="primary"
        />
      </div>
    </v-card>

    <!-- Add New Labor Dialog - Clean Form Style -->
    <v-dialog v-model="showAddDialog" max-width="900" scrollable persistent>
      <v-card class="clean-dialog-card clean-form-card">
        <!-- Header Section -->
        <v-card-title class="clean-dialog-header clean-form-header">
          <h2 class="clean-form-title">معلومات العامل</h2>
        </v-card-title>

        <!-- Form Content -->
        <v-card-text class="clean-form-content">
          <p class="clean-form-instruction">
            لإتمام إضافة العامل، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
          </p>

          <v-form ref="form" v-model="formValid">
            <!-- الصف الأول: اسم العامل، المهنة -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    اسم العامل <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model="newLabor.workerName"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل اسم العامل"
                    :rules="[v => !!v || 'اسم العامل مطلوب']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    المهنة
                  </label>
                  <v-text-field
                    v-model="newLabor.jobTitle"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل المهنة"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثاني: العدد، الأجرة اليومية -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    عدد الأيام <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="newLabor.quantity"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'عدد الأيام يجب أن يكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    الأجرة اليومية (د.ع) <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="newLabor.cost"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'الأجرة يجب أن تكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثالث: المجموع، رقم الهاتف -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    المجموع (د.ع)
                  </label>
                  <v-text-field
                    :value="(newLabor.quantity * newLabor.cost) || 0"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    readonly
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    رقم الهاتف
                  </label>
                  <v-text-field
                    v-model="newLabor.phone"
                    variant="outlined"
                    density="comfortable"
                    placeholder="07XX XXX XXXX"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الرابع: العنوان -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    العنوان
                  </label>
                  <v-text-field
                    v-model="newLabor.address"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل العنوان"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الخامس: الملاحظات -->
            <v-row class="clean-form-row">
              <v-col cols="12" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">ملاحظات</label>
                  <v-textarea
                    v-model="newLabor.notes"
                    variant="outlined"
                    rows="4"
                    density="comfortable"
                    placeholder="أدخل ملاحظات إضافية"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="clean-form-actions">
          <v-spacer />
          <v-btn
            class="clean-form-cancel-btn"
            variant="outlined"
            @click="closeAddDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            class="clean-form-continue-btn"
            variant="elevated"
            :disabled="!formValid"
            @click="saveLabor"
          >
            حفظ العامل
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Edit Labor Dialog -->
    <v-dialog v-model="showEditDialog" max-width="900" scrollable persistent>
      <v-card class="clean-dialog-card clean-form-card">
        <!-- Header Section -->
        <v-card-title class="clean-dialog-header clean-form-header">
          <h2 class="clean-form-title">تعديل معلومات العامل</h2>
        </v-card-title>

        <!-- Form Content -->
        <v-card-text class="clean-form-content">
          <p class="clean-form-instruction">
            قم بتعديل المعلومات المطلوبة. جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
          </p>

          <v-form ref="editForm" v-model="editFormValid">
            <!-- الصف الأول: اسم العامل، المهنة -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    اسم العامل <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model="editLabor.workerName"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل اسم العامل"
                    :rules="[v => !!v || 'اسم العامل مطلوب']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    المهنة
                  </label>
                  <v-text-field
                    v-model="editLabor.jobTitle"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل المهنة"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثاني: العدد، الأجرة اليومية -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    عدد الأيام <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="editLabor.quantity"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'عدد الأيام يجب أن يكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    الأجرة اليومية (د.ع) <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="editLabor.cost"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'الأجرة يجب أن تكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثالث: المجموع، رقم الهاتف -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    المجموع (د.ع)
                  </label>
                  <v-text-field
                    :value="(editLabor.quantity * editLabor.cost) || 0"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    readonly
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    رقم الهاتف
                  </label>
                  <v-text-field
                    v-model="editLabor.phone"
                    variant="outlined"
                    density="comfortable"
                    placeholder="07XX XXX XXXX"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الرابع: العنوان -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    العنوان
                  </label>
                  <v-text-field
                    v-model="editLabor.address"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل العنوان"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الخامس: الملاحظات -->
            <v-row class="clean-form-row">
              <v-col cols="12" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">ملاحظات</label>
                  <v-textarea
                    v-model="editLabor.notes"
                    variant="outlined"
                    rows="4"
                    density="comfortable"
                    placeholder="أدخل ملاحظات إضافية"
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="clean-form-actions">
          <v-spacer />
          <v-btn
            class="clean-form-cancel-btn"
            variant="outlined"
            @click="closeEditDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            class="clean-form-continue-btn"
            variant="elevated"
            :disabled="!editFormValid"
            @click="saveEditLabor"
          >
            حفظ التعديلات
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { listLaborByWorkDay, createLabor, updateLabor, deleteLabor as deleteLaborApi } from '@/api/materials'
import { DEFAULT_LIMIT } from '@/constants/pagination'

const router = useRouter()
const route = useRoute()

// Get workDayId and projectId from route query
const workDayId = computed(() => route.query.workDayId || null)
const projectId = computed(() => route.query.projectId || null)

// Reactive data
const laborSearch = ref('')
const showAddDialog = ref(false)
const showEditDialog = ref(false)
const formValid = ref(false)
const editFormValid = ref(false)
const loading = ref(false)
const form = ref(null)
const editForm = ref(null)
const editingLaborId = ref(null)

// Form data - aligned with backend DTO
const newLabor = ref({
  workerName: '',
  jobTitle: '',
  quantity: 0,
  cost: 0,
  phone: '',
  address: '',
  notes: ''
})

// Edit form data
const editLabor = ref({
  workerName: '',
  jobTitle: '',
  quantity: 0,
  cost: 0,
  phone: '',
  address: '',
  notes: ''
})

// Table headers - aligned with backend fields
const laborHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'اسم العامل', key: 'workerName', sortable: true },
  { title: 'المهنة', key: 'jobTitle', sortable: true },
  { title: 'عدد الأيام', key: 'quantity', sortable: true },
  { title: 'الأجرة اليومية', key: 'cost', sortable: true },
  { title: 'المجموع', key: 'total', sortable: true },
  { title: 'رقم الهاتف', key: 'phone', sortable: true },
  { title: 'العنوان', key: 'address', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Data from backend
const laborData = ref([])

// Pagination state
const currentPage = ref(1)

// Computed total pages for pagination
const totalPages = computed(() => Math.ceil(laborData.value.length / DEFAULT_LIMIT))

// Methods
const goBack = () => {
  const query = {}
  if (workDayId.value) query.id = workDayId.value
  if (projectId.value) query.projectId = projectId.value
  router.push({ path: '/work-day-details', query })
}

// Load data from backend
const loadLabor = async () => {
  if (!workDayId.value) return
  loading.value = true
  try {
    const data = await listLaborByWorkDay(workDayId.value)
    laborData.value = data.map(item => ({
      ...item,
      total: item.quantity * item.cost
    }))
  } catch (err) {
    console.error('Error loading labor:', err)
  } finally {
    loading.value = false
  }
}

const searchLabor = () => {
  console.log('البحث في العمال:', laborSearch.value)
}

const openAddDialog = () => {
  newLabor.value = {
    workerName: '',
    jobTitle: '',
    quantity: 0,
    cost: 0,
    phone: '',
    address: '',
    notes: ''
  }
  showAddDialog.value = true
}

const closeAddDialog = () => {
  showAddDialog.value = false
  if (form.value) {
    form.value.reset()
  }
  newLabor.value = {
    workerName: '',
    jobTitle: '',
    quantity: 0,
    cost: 0,
    phone: '',
    address: '',
    notes: ''
  }
}

const saveLabor = async () => {
  const { valid } = await form.value.validate()
  if (valid) {
    try {
      const wdId = parseInt(workDayId.value)
      if (!wdId || isNaN(wdId)) {
        console.error('Invalid workDayId:', workDayId.value)
        return
      }

      const payload = {
        workDayId: wdId,
        workerName: newLabor.value.workerName,
        quantity: Number(newLabor.value.quantity),
        cost: Number(newLabor.value.cost)
      }
      // Optional fields
      if (newLabor.value.jobTitle && newLabor.value.jobTitle.trim()) {
        payload.jobTitle = newLabor.value.jobTitle.trim()
      }
      if (newLabor.value.phone && newLabor.value.phone.trim()) {
        payload.phone = newLabor.value.phone.trim()
      }
      if (newLabor.value.address && newLabor.value.address.trim()) {
        payload.address = newLabor.value.address.trim()
      }
      if (newLabor.value.notes && newLabor.value.notes.trim()) {
        payload.notes = newLabor.value.notes.trim()
      }

      await createLabor(payload)
      await loadLabor()

      // Reset form
      form.value.reset()
      newLabor.value = {
        workerName: '',
        jobTitle: '',
        quantity: 0,
        cost: 0,
        phone: '',
        address: '',
        notes: ''
      }

      showAddDialog.value = false
    } catch (err) {
      console.error('Error creating labor:', err)
    }
  }
}

const deleteLabor = async (item) => {
  if (confirm('هل أنت متأكد من حذف هذا العامل؟')) {
    try {
      await deleteLaborApi(item.id)
      await loadLabor()
    } catch (err) {
      console.error('Error deleting labor:', err)
    }
  }
}

// Edit functions
const openEditDialog = (item) => {
  editingLaborId.value = item.id
  editLabor.value = {
    workerName: item.workerName || '',
    jobTitle: item.jobTitle || '',
    quantity: item.quantity || 0,
    cost: item.cost || 0,
    phone: item.phone || '',
    address: item.address || '',
    notes: item.notes || ''
  }
  showEditDialog.value = true
}

const closeEditDialog = () => {
  showEditDialog.value = false
  editingLaborId.value = null
  if (editForm.value) {
    editForm.value.reset()
  }
  editLabor.value = {
    workerName: '',
    jobTitle: '',
    quantity: 0,
    cost: 0,
    phone: '',
    address: '',
    notes: ''
  }
}

const saveEditLabor = async () => {
  const { valid } = await editForm.value.validate()
  if (valid && editingLaborId.value) {
    try {
      const payload = {
        workerName: editLabor.value.workerName,
        quantity: Number(editLabor.value.quantity),
        cost: Number(editLabor.value.cost)
      }
      // Optional fields
      if (editLabor.value.jobTitle && editLabor.value.jobTitle.trim()) {
        payload.jobTitle = editLabor.value.jobTitle.trim()
      }
      if (editLabor.value.phone && editLabor.value.phone.trim()) {
        payload.phone = editLabor.value.phone.trim()
      }
      if (editLabor.value.address && editLabor.value.address.trim()) {
        payload.address = editLabor.value.address.trim()
      }
      if (editLabor.value.notes && editLabor.value.notes.trim()) {
        payload.notes = editLabor.value.notes.trim()
      }

      await updateLabor(editingLaborId.value, payload)
      await loadLabor()
      closeEditDialog()
    } catch (err) {
      console.error('Error updating labor:', err)
    }
  }
}

// Format currency
const formatCurrency = (value) => {
  if (!value) return '0 IQD'
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(value) + ' IQD'
}

// Lifecycle - watch workDayId and load data when available
watch(workDayId, (newId) => {
  if (newId) {
    loadLabor()
  }
}, { immediate: true })
</script>

<style scoped>
.labor-details-page {
  min-height: 100vh;
  background: #f5f5f5;
  padding: 0 !important;
  padding-top: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  position: relative;
  overflow-x: hidden;
  direction: rtl;
  margin: 0;
  width: 100%;
}

.page-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white;
  padding: 0.4rem 0.75rem !important;
  border-radius: 0 !important;
  margin: 0 !important;
  margin-bottom: 1rem !important;
  margin-top: 0 !important;
  position: relative;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(25, 118, 210, 0.15);
  border: none !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
  backdrop-filter: blur(15px);
  z-index: 1;
  width: 100% !important;
  max-width: 100% !important;
  box-sizing: border-box !important;
}

.content-container {
  padding: 0 2rem 2rem 2rem !important;
  max-width: 100%;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  max-width: 100%;
}

.back-btn {
  background: rgba(255, 255, 255, 0.25) !important;
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 600;
  font-size: 0.875rem;
  letter-spacing: 0.3px;
  line-height: 1.2;
  min-width: 36px !important;
  height: 36px !important;
}

.back-btn :deep(.v-icon) {
  color: #ffffff !important;
}

.back-btn :deep(.v-btn__overlay) {
  color: #ffffff !important;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.35) !important;
  transform: translateX(2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
}

.page-title {
  font-size: 1.1rem !important;
  font-weight: 600 !important;
  margin: 0 !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(0, 0, 0, 0.1);
  color: #ffffff !important;
  letter-spacing: 0.3px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.3 !important;
  padding: 0.2rem 0.6rem !important;
  border-radius: 6px !important;
  display: inline-block;
  backdrop-filter: blur(10px);
  box-shadow: 0 2px 8px rgba(26, 35, 126, 0.3);
}

.page-subtitle {
  font-size: 0.875rem !important;
  opacity: 0.9 !important;
  margin: 0.25rem 0 0 0 !important;
  color: #ffffff !important;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
  font-weight: 500 !important;
  letter-spacing: 0.3px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.3 !important;
}

.section-card {
  border-radius: 16px;
  overflow: hidden;
}

.section-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  background-size: 200% 200%;
  animation: gradientShift 3s ease infinite;
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 4px rgba(0, 0, 0, 0.3), 0 0 12px rgba(255, 255, 255, 0.2);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.3px !important;
  line-height: 1.3 !important;
  margin-bottom: 0.75rem !important;
  padding: 0.5rem 0.875rem !important;
  position: relative;
  overflow: hidden;
  box-shadow: 0 2px 10px rgba(59, 130, 246, 0.25);
}

.section-title :deep(.v-icon) {
  font-size: 18px !important;
}

.section-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
}

.section-title::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.15) 0%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
}

.section-title:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(59, 130, 246, 0.5), 0 0 60px rgba(96, 165, 250, 0.3);
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

@keyframes pulse {
  0%, 100% { 
    opacity: 0.3; 
    transform: scale(0.8); 
  }
  50% { 
    opacity: 0.6; 
    transform: scale(1.2); 
  }
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* تحسين تصميم شريط البحث */
.search-container {
  padding: 0.75rem 1rem !important;
  background: linear-gradient(135deg, #eff6ff 0%, #dbeafe 50%, #bfdbfe 100%);
  border-radius: 8px !important;
  margin-bottom: 0.75rem !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.15);
  border: 1px solid #93c5fd !important;
  position: relative;
}

.search-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px !important;
  background: linear-gradient(90deg, #60a5fa, #3b82f6, #2563eb);
  border-radius: 8px 8px 0 0;
}

.search-box {
  display: flex;
  align-items: center;
  gap: 0.5rem !important;
  max-width: 600px;
  margin: 0 auto;
}

.search-field {
  flex: 1;
  margin-bottom: 0;
}

.search-field :deep(.v-field) {
  background-color: #ffffff !important;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease;
  min-height: 36px !important;
}

.search-field :deep(.v-field__input) {
  background-color: #ffffff !important;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid #3b82f6 !important;
  transition: all 0.3s ease;
}

.search-field :deep(.v-field__outline) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.search-field :deep(.v-field--focused .v-field__outline) {
  border-color: #2563eb !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2) !important;
}

.search-field :deep(.v-field__input):focus {
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.4) !important;
  border-color: #2563eb !important;
  transform: translateY(-1px);
}

.search-field :deep(.v-label) {
  color: #1e40af !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  opacity: 1 !important;
}

.search-field :deep(.v-label--active),
.search-field :deep(.v-field__label--active) {
  color: #1e40af !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.search-field :deep(.v-field__input input) {
  color: #111827 !important;
  font-weight: 500 !important;
  font-size: 0.875rem !important;
}

.search-field :deep(.v-field__prepend-inner) {
  color: #3b82f6 !important;
}

.search-field :deep(.v-field__prepend-inner .v-icon) {
  color: #3b82f6 !important;
  font-size: 18px !important;
}

.search-field :deep(.v-field__prepend-inner i) {
  color: #3b82f6 !important;
}

.search-btn {
  height: 36px !important;
  min-width: 80px !important;
  border-radius: 8px;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-transform: none;
  box-shadow: 0 2px 8px rgba(25, 118, 210, 0.25);
  transition: all 0.3s ease;
}

.search-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(25, 118, 210, 0.4);
}

/* أنماط أزرار الإضافة */
.add-btn {
  height: 36px !important;
  min-width: 100px !important;
  border-radius: 8px;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-transform: none;
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #4caf50 0%, #388e3c 100%) !important;
  box-shadow: 0 2px 8px rgba(76, 175, 80, 0.25);
}

.add-btn :deep(.v-icon) {
  font-size: 18px !important;
}

.add-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(76, 175, 80, 0.4);
  background: linear-gradient(135deg, #388e3c 0%, #2e7d32 100%) !important;
}

.data-table {
  margin-top: 1rem;
}

/* ========================================
   تنسيقات الجدول - نفس تنسيقات الموارد البشرية
   ======================================== */

.data-table-card {
  background: rgba(255, 249, 249, 0.95);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin-top: 24px;
}

.table-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  padding: 8px 12px !important;
  border-radius: 8px 8px 0 0 !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.25) !important;
  margin: 0 !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-align: right !important;
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
  min-height: 36px !important;
}

.indigo-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  font-weight: 600 !important;
}

.title-text {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.25) !important;
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

.project-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: 500 !important;
  font-size: 0.55rem !important;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  padding: 2px 4px !important;
  min-height: 20px !important;
  border-bottom: 1px solid #0d47a1 !important;
  position: relative;
}

.project-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  padding: 3px 4px !important;
  min-height: 32px !important;
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

.project-table .serial-number {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.6rem !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 3px 6px !important;
  border-radius: 6px !important;
  display: inline-block !important;
  min-width: 30px !important;
  text-align: center !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
}

.project-table .project-name {
  color: #000000 !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  text-align: right !important;
  padding: 2px 4px !important;
  border-radius: 4px !important;
  display: inline-block !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.project-table .date-text {
  color: #000000 !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.project-table .cost-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.6rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 3px 6px !important;
  border-radius: 6px !important;
  display: inline-block !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #d1d5db !important;
  white-space: nowrap !important;
  min-width: 80px !important;
  text-align: center !important;
}

.project-table .category-chip {
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  padding: 3px 6px !important;
  border-radius: 6px !important;
  color: #000000 !important;
  background: #f3f4f6 !important;
  border: 1px solid #e5e7eb !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
}

.action-buttons {
  display: flex;
  gap: 4px;
  align-items: center;
  justify-content: center;
}

.action-btn {
  min-width: 28px !important;
  height: 28px !important;
  border-radius: 4px !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.08) !important;
}

.action-btn .v-icon {
  font-size: 16px !important;
}

.action-btn:hover {
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15) !important;
}

.dialog-card {
  border-radius: 16px;
  overflow: hidden;
}

.dialog-title {
  background: linear-gradient(135deg, #4caf50 0%, #388e3c 100%);
  color: white;
  font-weight: bold;
  font-size: 1.3rem;
}

.dialog-content {
  padding: 2rem;
}

.dialog-actions {
  padding: 1rem 2rem;
  background: #f8f9fa;
}

/* Modal Styles - تطبيق تصميم إضافة يوم عمل */
.modal-card {
  border-radius: 12px !important;
  overflow: hidden;
}

.modal-header {
  background: #007bff !important;
  color: white !important;
  padding: 1rem 1.5rem !important;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-title {
  font-size: 1.4rem;
  font-weight: 800;
  color: white;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  letter-spacing: 0.5px;
}

.close-btn {
  color: white !important;
  min-width: 32px !important;
  height: 32px !important;
}

.modal-content {
  padding: 2rem !important;
  background: white;
}

.modal-content :deep(.v-field) {
  border-radius: 8px !important;
  background: #fafafa !important;
  transition: all 0.3s ease !important;
}

.modal-content :deep(.v-field:hover) {
  background: #f5f5f5 !important;
  border-color: #666 !important;
}

.modal-content :deep(.v-field__input) {
  color: #1e3a8a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  text-shadow: none !important;
}

.modal-content :deep(.v-label),
.modal-content :deep(.v-field__label),
.modal-content :deep(.v-label--active),
.modal-content :deep(.v-field__label--active),
.modal-content :deep(.v-label--floating),
.modal-content :deep(.v-field__label--floating) {
  color: #000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  text-shadow: none !important;
  opacity: 1 !important;
  background: white !important;
  padding: 0 8px !important;
}

.modal-content :deep(.v-field--focused .v-label),
.modal-content :deep(.v-field--focused .v-field__label) {
  color: #000 !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-field__outline) {
  border-color: #333 !important;
  border-width: 2px !important;
  border-style: solid !important;
  box-shadow: 0 0 0 1px rgba(51, 51, 51, 0.2) !important;
}

.modal-content :deep(.v-field--focused .v-field__outline) {
  border-color: #007bff !important;
  border-width: 3px !important;
  border-style: solid !important;
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.3) !important;
  outline: none !important;
}

.modal-content :deep(.v-field__input::placeholder) {
  color: #1e3a8a !important;
  opacity: 0.7 !important;
  font-weight: 500 !important;
}

/* Dropdown Menu Styling */
.modal-content :deep(.v-list),
.modal-content :deep(.v-menu__content),
.modal-content :deep(.v-overlay__content) {
  background: white !important;
  border: 2px solid #333 !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  max-height: 300px !important;
  overflow-y: auto !important;
}

.modal-content :deep(.v-list .v-list-item__content),
.modal-content :deep(.v-menu__content .v-list-item__content),
.modal-content :deep(.v-overlay__content .v-list-item__content) {
  color: #1e3a8a !important;
}

.modal-content :deep(.v-list .v-list-item-title),
.modal-content :deep(.v-menu__content .v-list-item-title),
.modal-content :deep(.v-overlay__content .v-list-item-title) {
  color: #1e3a8a !important;
  font-weight: 600 !important;
}

.modal-content :deep(.v-list-item),
.modal-content :deep(.v-menu__content .v-list-item),
.modal-content :deep(.v-overlay__content .v-list-item) {
  color: #1e3a8a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 0.75rem 1rem !important;
  border-bottom: 1px solid #e9ecef !important;
  transition: all 0.2s ease !important;
}

.modal-content :deep(.v-list-item:hover),
.modal-content :deep(.v-menu__content .v-list-item:hover),
.modal-content :deep(.v-overlay__content .v-list-item:hover) {
  background: #f8f9fa !important;
  color: #007bff !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-list-item--active),
.modal-content :deep(.v-menu__content .v-list-item--active),
.modal-content :deep(.v-overlay__content .v-list-item--active) {
  background: #007bff !important;
  color: white !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-list-item--active:hover),
.modal-content :deep(.v-menu__content .v-list-item--active:hover),
.modal-content :deep(.v-overlay__content .v-list-item--active:hover) {
  background: #0056b3 !important;
  color: white !important;
}

.modal-footer {
  padding: 1rem 1.5rem !important;
  background: #f8f9fa !important;
  border-top: 1px solid #dee2e6 !important;
}

.save-btn {
  background: #007bff !important;
  color: white !important;
  font-weight: 600;
  padding: 0.75rem 2rem !important;
}

/* Animations */
.glass-effect {
  backdrop-filter: blur(15px);
}

.page-header.glass-effect {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
}

.gradient-animation {
  background-size: 400% 400%;
  animation: gradientShift 20s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.text-glow {
  text-shadow: 0 0 30px rgba(255, 255, 255, 0.6);
}

.fade-in {
  animation: fadeIn 1.2s ease-out;
}

@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(30px) scale(0.95); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0) scale(1); 
  }
}

/* تأثيرات إضافية للرأس */
.page-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.page-header::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
}

@keyframes pulse {
  0%, 100% { opacity: 0.3; transform: scale(0.8); }
  50% { opacity: 0.6; transform: scale(1.2); }
}

/* Responsive */
@media (max-width: 768px) {
  .labor-details-page {
    padding: 1rem;
  }
  
  .page-title {
    font-size: 2rem;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
  }
  
  .dialog-content {
    padding: 1rem;
  }
  
  .search-box {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .search-btn, .add-btn {
    width: 100%;
    min-width: auto;
  }

  .search-container {
    padding: 0.75rem;
  }

  .search-field {
    margin-bottom: 0.5rem;
  }
}

/* ========================================
   Clean Form Styles - Same as Materials
   ======================================== */
.clean-form-card {
  background: #ffffff !important;
}

.clean-dialog-header,
.clean-form-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  color: #ffffff !important;
  padding: 1rem 1.5rem !important;
}

.clean-form-title {
  font-size: 1.25rem !important;
  font-weight: 600 !important;
  margin: 0 !important;
  color: #ffffff !important;
}

.clean-form-content {
  padding: 1.5rem !important;
  background: #ffffff !important;
}

.clean-form-instruction {
  color: #666666 !important;
  font-size: 0.9rem !important;
  margin-bottom: 1.5rem !important;
  line-height: 1.6 !important;
}

.clean-form-row {
  margin-bottom: 0.5rem !important;
}

.clean-form-column {
  padding: 0.5rem !important;
}

.clean-form-field-wrapper {
  margin-bottom: 0.5rem;
}

.clean-form-label {
  display: block;
  font-size: 0.875rem !important;
  font-weight: 600 !important;
  color: #333333 !important;
  margin-bottom: 0.5rem !important;
}

.required-star {
  color: #d32f2f !important;
}

.clean-form-input :deep(.v-field) {
  background-color: #f5f5f5 !important;
  border-radius: 8px !important;
}

.clean-form-input :deep(.v-field__input) {
  color: #333333 !important;
  font-size: 0.95rem !important;
}

.clean-form-input :deep(.v-field__outline) {
  border-color: #e0e0e0 !important;
}

.clean-form-input :deep(.v-field--focused .v-field__outline) {
  border-color: #1976d2 !important;
  border-width: 2px !important;
}

.clean-form-input :deep(.v-label) {
  color: #666666 !important;
}

.clean-form-input :deep(input::placeholder),
.clean-form-input :deep(textarea::placeholder) {
  color: #999999 !important;
}

.clean-form-actions {
  padding: 1rem 1.5rem !important;
  background: #f5f5f5 !important;
  border-top: 1px solid #e0e0e0 !important;
}

.clean-form-cancel-btn {
  color: #666666 !important;
  border-color: #cccccc !important;
}

.clean-form-cancel-btn:hover {
  background: #eeeeee !important;
}

.clean-form-continue-btn {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  color: #ffffff !important;
}

.clean-form-continue-btn:hover {
  background: linear-gradient(135deg, #1565c0 0%, #0d47a1 100%) !important;
}

.clean-form-continue-btn:disabled {
  background: #cccccc !important;
  color: #999999 !important;
}

</style>
