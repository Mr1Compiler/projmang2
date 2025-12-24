<template>
  <div class="equipment-details-page">
    <!-- Header Section -->
    <div class="page-header glass-effect gradient-animation">
      <div class="header-content">
        <div class="header-text">
          <h1 class="page-title text-glow fade-in">تفاصيل الآليات</h1>
          <p class="page-subtitle fade-in">إدارة المعدات والآلات المستخدمة في المشروع</p>
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

    <!-- Equipment Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="white">mdi-truck</v-icon>
        الآليات والمعدات
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="equipmentSearch"
              label="البحث في الآليات..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchEquipment"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchEquipment"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="indigo" 
              variant="elevated" 
              class="add-btn"
              @click="openAddDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة آلة
            </v-btn>
          </div>
        </div>
      </v-card-text>

    </v-card>

    <!-- Equipment Table Card -->
    <v-card class="data-table-card" elevation="2">
      <v-card-title class="table-title indigo-title">
        <span class="title-text">قائمة الآليات</span>
      </v-card-title>

      <v-data-table
        v-model:page="currentPage"
        :headers="equipmentHeaders"
        :items="equipmentData"
        :search="equipmentSearch"
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

        <!-- Equipment Name Column -->
        <template #item.equipmentName="{ item }">
          <span class="project-name">{{ item.equipmentName }}</span>
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

        <!-- Notes Column -->
        <template #item.notes="{ item }">
          <span class="project-name">{{ item.notes || '-' }}</span>
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
              title="تعديل الآلة"
              class="action-btn"
            />
            <v-btn
              icon="mdi-delete"
              size="small"
              color="error"
              variant="text"
              @click="deleteEquipment(item)"
              title="حذف الآلة"
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

    <!-- Add New Equipment Dialog - Clean Form Style -->
    <v-dialog v-model="showAddDialog" max-width="900" scrollable persistent>
      <v-card class="clean-dialog-card clean-form-card">
        <!-- Header Section -->
        <v-card-title class="clean-dialog-header clean-form-header">
          <h2 class="clean-form-title">معلومات الآلة</h2>
        </v-card-title>

        <!-- Form Content -->
        <v-card-text class="clean-form-content">
          <p class="clean-form-instruction">
            لإتمام إضافة الآلة، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
          </p>

          <v-form ref="form" v-model="formValid">
            <!-- الصف الأول: اسم الآلة -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    اسم الآلة <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model="newEquipment.equipmentName"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل اسم الآلة"
                    :rules="[v => !!v || 'اسم الآلة مطلوب']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثاني: العدد، التكلفة -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    العدد <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="newEquipment.quantity"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'العدد يجب أن يكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    التكلفة (د.ع) <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="newEquipment.cost"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'التكلفة يجب أن تكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثالث: المجموع -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    المجموع (د.ع)
                  </label>
                  <v-text-field
                    :value="(newEquipment.quantity * newEquipment.cost) || 0"
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
            </v-row>

            <!-- الصف الرابع: الملاحظات -->
            <v-row class="clean-form-row">
              <v-col cols="12" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">ملاحظات</label>
                  <v-textarea
                    v-model="newEquipment.notes"
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
            @click="saveEquipment"
          >
            حفظ الآلة
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Edit Equipment Dialog -->
    <v-dialog v-model="showEditDialog" max-width="900" scrollable persistent>
      <v-card class="clean-dialog-card clean-form-card">
        <!-- Header Section -->
        <v-card-title class="clean-dialog-header clean-form-header">
          <h2 class="clean-form-title">تعديل معلومات الآلة</h2>
        </v-card-title>

        <!-- Form Content -->
        <v-card-text class="clean-form-content">
          <p class="clean-form-instruction">
            قم بتعديل المعلومات المطلوبة. جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
          </p>

          <v-form ref="editForm" v-model="editFormValid">
            <!-- الصف الأول: اسم الآلة -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    اسم الآلة <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model="editEquipment.equipmentName"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل اسم الآلة"
                    :rules="[v => !!v || 'اسم الآلة مطلوب']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثاني: العدد، التكلفة -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    العدد <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="editEquipment.quantity"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'العدد يجب أن يكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    التكلفة (د.ع) <span class="required-star">*</span>
                  </label>
                  <v-text-field
                    v-model.number="editEquipment.cost"
                    variant="outlined"
                    density="comfortable"
                    type="number"
                    placeholder="0"
                    :rules="[v => (v > 0) || 'التكلفة يجب أن تكون أكبر من صفر']"
                    required
                    hide-details="auto"
                    class="clean-form-input"
                  />
                </div>
              </v-col>
            </v-row>

            <!-- الصف الثالث: المجموع -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">
                    المجموع (د.ع)
                  </label>
                  <v-text-field
                    :value="(editEquipment.quantity * editEquipment.cost) || 0"
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
            </v-row>

            <!-- الصف الرابع: الملاحظات -->
            <v-row class="clean-form-row">
              <v-col cols="12" class="clean-form-column">
                <div class="clean-form-field-wrapper">
                  <label class="clean-form-label">ملاحظات</label>
                  <v-textarea
                    v-model="editEquipment.notes"
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
            @click="saveEditEquipment"
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
import { listEquipmentByWorkDay, createEquipment, updateEquipment, deleteEquipment as deleteEquipmentApi } from '@/api/materials'
import { DEFAULT_LIMIT } from '@/constants/pagination'

const router = useRouter()
const route = useRoute()

// Get workDayId and projectId from route query
const workDayId = computed(() => route.query.workDayId || null)
const projectId = computed(() => route.query.projectId || null)

// Reactive data
const equipmentSearch = ref('')
const showAddDialog = ref(false)
const formValid = ref(false)
const loading = ref(false)
const form = ref(null)

// Edit dialog state
const showEditDialog = ref(false)
const editFormValid = ref(false)
const editForm = ref(null)
const editingEquipmentId = ref(null)
const editEquipment = ref({
  equipmentName: '',
  quantity: 0,
  cost: 0,
  notes: ''
})

// Pagination state
const currentPage = ref(1)

// Form data - aligned with backend DTO
const newEquipment = ref({
  equipmentName: '',
  quantity: 0,
  cost: 0,
  notes: ''
})

// Table headers - aligned with backend fields
const equipmentHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'اسم الآلة', key: 'equipmentName', sortable: true },
  { title: 'العدد', key: 'quantity', sortable: true },
  { title: 'التكلفة', key: 'cost', sortable: true },
  { title: 'المجموع', key: 'total', sortable: true },
  { title: 'الملاحظات', key: 'notes', sortable: false },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Data from backend
const equipmentData = ref([])

// Computed total pages for pagination
const totalPages = computed(() => Math.ceil(equipmentData.value.length / DEFAULT_LIMIT))

// Methods
const goBack = () => {
  const query = {}
  if (workDayId.value) query.id = workDayId.value
  if (projectId.value) query.projectId = projectId.value
  router.push({ path: '/work-day-details', query })
}

// Load data from backend
const loadEquipment = async () => {
  if (!workDayId.value) return
  loading.value = true
  try {
    const data = await listEquipmentByWorkDay(workDayId.value)
    equipmentData.value = data.map(item => ({
      ...item,
      total: item.quantity * item.cost
    }))
  } catch (err) {
    console.error('Error loading equipment:', err)
  } finally {
    loading.value = false
  }
}

const searchEquipment = () => {
  console.log('البحث في الآليات:', equipmentSearch.value)
}

const openAddDialog = () => {
  newEquipment.value = {
    equipmentName: '',
    quantity: 0,
    cost: 0,
    notes: ''
  }
  showAddDialog.value = true
}

const closeAddDialog = () => {
  showAddDialog.value = false
  if (form.value) {
    form.value.reset()
  }
  newEquipment.value = {
    equipmentName: '',
    quantity: 0,
    cost: 0,
    notes: ''
  }
}

const saveEquipment = async () => {
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
        equipmentName: newEquipment.value.equipmentName,
        quantity: Number(newEquipment.value.quantity),
        cost: Number(newEquipment.value.cost)
      }
      // Optional field
      if (newEquipment.value.notes && newEquipment.value.notes.trim()) {
        payload.notes = newEquipment.value.notes.trim()
      }

      await createEquipment(payload)
      await loadEquipment()

      // Reset form
      form.value.reset()
      newEquipment.value = {
        equipmentName: '',
        quantity: 0,
        cost: 0,
        notes: ''
      }

      showAddDialog.value = false
    } catch (err) {
      console.error('Error creating equipment:', err)
    }
  }
}

const deleteEquipment = async (item) => {
  if (confirm('هل أنت متأكد من حذف هذه الآلة؟')) {
    try {
      await deleteEquipmentApi(item.id)
      await loadEquipment()
    } catch (err) {
      console.error('Error deleting equipment:', err)
    }
  }
}

// Edit functions
const openEditDialog = (item) => {
  editingEquipmentId.value = item.id
  editEquipment.value = {
    equipmentName: item.equipmentName,
    quantity: item.quantity,
    cost: item.cost,
    notes: item.notes || ''
  }
  showEditDialog.value = true
}

const closeEditDialog = () => {
  showEditDialog.value = false
  if (editForm.value) {
    editForm.value.reset()
  }
  editingEquipmentId.value = null
  editEquipment.value = {
    equipmentName: '',
    quantity: 0,
    cost: 0,
    notes: ''
  }
}

const saveEditEquipment = async () => {
  const { valid } = await editForm.value.validate()
  if (valid) {
    try {
      const payload = {
        equipmentName: editEquipment.value.equipmentName,
        quantity: Number(editEquipment.value.quantity),
        cost: Number(editEquipment.value.cost)
      }
      if (editEquipment.value.notes && editEquipment.value.notes.trim()) {
        payload.notes = editEquipment.value.notes.trim()
      }

      await updateEquipment(editingEquipmentId.value, payload)
      await loadEquipment()
      closeEditDialog()
    } catch (err) {
      console.error('Error updating equipment:', err)
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
    loadEquipment()
  }
}, { immediate: true })
</script>

<style scoped>
.equipment-details-page {
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

.content-container {
  padding: 0 2rem 2rem 2rem !important;
  max-width: 100%;
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

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  position: relative;
  z-index: 1;
}

.header-text {
  flex: 1;
  position: relative;
  z-index: 1;
}

.back-btn {
  background: rgba(255, 255, 255, 0.15) !important;
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  position: relative;
  z-index: 1;
  min-width: 36px !important;
  height: 36px !important;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.25) !important;
  transform: translateX(2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.page-title {
  font-size: 1.4rem !important;
  font-weight: 700 !important;
  margin: 0;
  padding: 0.4rem 0.8rem !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.3px !important;
  color: #ffffff !important;
  border-radius: 6px !important;
}

.page-subtitle {
  font-size: 0.9rem !important;
  opacity: 0.95;
  margin: 0.3rem 0 0 0 !important;
  font-weight: 400;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
  letter-spacing: 0.3px !important;
  color: rgba(255, 255, 255, 0.95);
}

.section-card {
  border-radius: 16px;
  overflow: hidden;
}

.section-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 100%);
  color: white;
  font-weight: 600 !important;
  font-size: 1rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  padding: 0.6rem 1rem !important;
}

.section-title :deep(.v-icon) {
  font-size: 18px !important;
  color: #ffffff !important;
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
  background: linear-gradient(135deg, #3f51b5 0%, #303f9f 100%) !important;
  box-shadow: 0 2px 8px rgba(63, 81, 181, 0.25);
}

.add-btn :deep(.v-icon) {
  font-size: 18px !important;
}

.add-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(63, 81, 181, 0.4);
  background: linear-gradient(135deg, #303f9f 0%, #283593 100%) !important;
}

.data-table {
  margin-top: 1rem;
}

/* تحسين ألوان النصوص في الجداول */
.data-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  padding: 3px 4px !important;
  min-height: 32px !important;
}

.data-table :deep(.v-data-table__th) {
  color: #ffffff !important;
  font-weight: 500 !important;
  font-size: 0.55rem !important;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  padding: 2px 4px !important;
  min-height: 20px !important;
}

.data-table :deep(.v-data-table__tr:hover) {
  background-color: #f8f9fa !important;
}

.data-table :deep(.v-data-table__tr:hover .v-data-table__td) {
  color: #000 !important;
}

/* تحسين ألوان النصوص في الجداول الفرعية */
.data-table :deep(.v-data-table__wrapper) {
  border: 2px solid #e8eaf6;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(63, 81, 181, 0.1);
  background: white;
}

.data-table :deep(.v-data-table__wrapper table) {
  border-collapse: separate;
  border-spacing: 0;
}

.data-table :deep(.v-data-table__wrapper table td) {
  color: #333 !important;
  border-color: #e0e0e0 !important;
}

.data-table :deep(.v-data-table__wrapper table th) {
  color: #ffffff !important;
  background: linear-gradient(135deg, #3f51b5 0%, #303f9f 100%) !important;
  border-color: #283593 !important;
}

/* أنماط مخصصة لجداول الآليات */
.equipment-table :deep(.v-data-table__td) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.6rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  padding: 3px 4px !important;
  min-height: 32px !important;
}

.equipment-table :deep(.v-data-table__th) {
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

/* تحسين ألوان الأزرار في الجداول */
.data-table :deep(.v-btn) {
  color: #dc3545 !important;
  transition: all 0.3s ease;
  min-width: 28px !important;
  height: 28px !important;
}

.data-table :deep(.v-btn .v-icon) {
  font-size: 16px !important;
}

.data-table :deep(.v-btn:hover) {
  background-color: #f8d7da !important;
  color: #721c24 !important;
  transform: scale(1.05);
  box-shadow: 0 2px 6px rgba(220, 53, 69, 0.25);
}

.dialog-card {
  border-radius: 16px;
  overflow: hidden;
}

.dialog-title {
  background: linear-gradient(135deg, #3f51b5 0%, #303f9f 100%);
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
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.15) 50%, transparent 70%);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
  border-radius: 20px;
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
  background: radial-gradient(circle, rgba(255, 255, 255, 0.15) 0%, rgba(66, 165, 245, 0.1) 50%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
  border-radius: 50%;
}

@keyframes pulse {
  0%, 100% { opacity: 0.3; transform: scale(0.8); }
  50% { opacity: 0.6; transform: scale(1.2); }
}

/* ========================================
   Table Card Styles - Same as Materials/Labor
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

/* ========================================
   Clean Form Styles - Same as Materials/Labor
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

/* Responsive */
@media (max-width: 768px) {
  .equipment-details-page {
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
</style>
