<template>
  <div class="fill-height data-page">
    <v-container fluid class="pa-6" style="padding: 0 20px !important;">
      <!-- Header -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-file-document-multiple</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">إدارة الصفحات</h1>
              <p class="subtitle">إنشاء وتعديل صفحات النظام</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Statistics -->
      <v-row class="mb-8">
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-primary" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-file-document-multiple</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ pages.length }}</h3>
                <p class="stat-label">إجمالي الصفحات</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ activePages }}</h3>
                <p class="stat-label">صفحات نشطة</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-pause-circle</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ inactivePages }}</h3>
                <p class="stat-label">صفحات غير نشطة</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Pages Table -->
      <v-card class="users-table" elevation="2">
        <v-card-title class="table-title-header d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" size="18" style="color: #ffffff !important;">mdi-table</v-icon>
            <span class="title-text">قائمة الصفحات</span>
          </div>
          <div class="d-flex table-header-buttons" style="gap: 0.5rem;">
            <v-btn
              v-if="canCreate"
              class="add-button add-user-btn btn-glow light-sweep smooth-transition"
              @click="openAddPageDialog"
              elevation="2"
              color="primary"
              size="small"
              style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
            >
              <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
              إضافة صفحة جديدة
            </v-btn>
          </div>
        </v-card-title>
        <div class="table-spacer"></div>
        <v-data-table
          :headers="headers"
          :items="pages"
          :loading="loading"
          class="elevation-0 users-data-table"
          :items-per-page="-1"
          hide-default-footer
        >
          <template v-slot:item.name="{ item }">
            <div class="d-flex align-center">
              <v-avatar size="36" color="primary" class="me-3">
                <v-icon color="white" size="20">{{ item.icon || 'mdi-file' }}</v-icon>
              </v-avatar>
              <div>
                <div class="font-weight-medium">{{ item.name }}</div>
                <div class="text-caption text-grey">{{ item.route }}</div>
              </div>
            </div>
          </template>

          <template v-slot:item.icon="{ item }">
            <v-icon>{{ item.icon || 'mdi-file' }}</v-icon>
            <span class="text-caption text-grey ms-2">{{ item.icon }}</span>
          </template>

          <template v-slot:item.status="{ item }">
            <v-chip
              :color="item.status === 'active' ? 'success' : 'warning'"
              size="small"
              variant="flat"
            >
              {{ item.status === 'active' ? 'نشط' : 'غير نشط' }}
            </v-chip>
          </template>

          <template v-slot:item.createdAt="{ item }">
            <span class="text-body-2">{{ formatDate(item.createdAt) }}</span>
          </template>

          <template v-slot:item.actions="{ item }">
            <v-btn
              icon="mdi-eye"
              size="small"
              variant="elevated"
              class="view-btn me-1"
              @click="viewPage(item)"
            />
            <v-btn
              v-if="canUpdate"
              icon="mdi-pencil"
              size="small"
              variant="elevated"
              class="edit-btn me-1"
              @click="editPage(item)"
            />
            <v-btn
              v-if="canDelete"
              icon="mdi-delete"
              size="small"
              variant="elevated"
              class="delete-btn"
              @click="confirmDeletePage(item)"
            />
          </template>
        </v-data-table>
      </v-card>
    </v-container>
  </div>

  <!-- Add/Edit Page Dialog -->
  <v-dialog v-model="showPageDialog" max-width="700px" persistent>
    <v-card class="add-user-dialog" rounded="lg">
      <v-card-title class="dialog-header pa-4">
        <div class="d-flex align-center">
          <v-avatar color="white" size="44" class="me-3">
            <v-icon color="primary" size="26">{{ isEditing ? 'mdi-pencil' : 'mdi-plus' }}</v-icon>
          </v-avatar>
          <div>
            <h3 class="text-h5 font-weight-bold text-white mb-0">
              {{ isEditing ? 'تعديل الصفحة' : 'إضافة صفحة جديدة' }}
            </h3>
            <span class="text-body-2 text-white-darken-1">
              {{ isEditing ? 'تعديل بيانات الصفحة' : 'أدخل بيانات الصفحة الجديدة' }}
            </span>
          </div>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          size="default"
          @click="closePageDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-card-text class="dialog-content pa-8">
        <v-form ref="pageForm" v-model="formValid" lazy-validation>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="pageData.name"
                label="اسم الصفحة"
                placeholder="مثال: المشاريع"
                :rules="[v => !!v || 'اسم الصفحة مطلوب']"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-file-document"
                color="primary"
                bg-color="grey-lighten-5"
                class="mb-2"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="pageData.route"
                label="مسار الصفحة"
                placeholder="مثال: /projects"
                :rules="routeRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-link"
                color="primary"
                bg-color="grey-lighten-5"
                class="mb-2"
                dir="ltr"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-combobox
                v-model="pageData.icon"
                :items="iconOptions"
                label="أيقونة الصفحة"
                placeholder="اختر أو اكتب اسم الأيقونة"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-emoticon"
                color="primary"
                bg-color="grey-lighten-5"
                class="mb-2"
              >
                <template v-slot:prepend-inner>
                  <v-icon>{{ pageData.icon || 'mdi-file' }}</v-icon>
                </template>
                <template v-slot:item="{ item, props }">
                  <v-list-item v-bind="props">
                    <template v-slot:prepend>
                      <v-icon>{{ item.value }}</v-icon>
                    </template>
                  </v-list-item>
                </template>
              </v-combobox>
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="pageData.status"
                :items="statusOptions"
                label="حالة الصفحة"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-toggle-switch"
                color="primary"
                bg-color="grey-lighten-5"
                class="mb-2"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-card-actions class="dialog-actions pa-5">
        <v-spacer />
        <v-btn variant="outlined" size="large" @click="closePageDialog" class="me-3 px-8">
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          size="large"
          @click="savePage"
          :loading="saving"
          :disabled="!formValid"
          class="px-8"
          prepend-icon="mdi-check"
        >
          <v-icon start size="16">mdi-content-save</v-icon>
          {{ isEditing ? 'حفظ التعديلات' : 'حفظ الصفحة' }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- View Page Dialog -->
  <v-dialog v-model="showViewDialog" max-width="600px">
    <v-card class="view-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="primary" class="me-3">mdi-file-document</v-icon>
          <h2>تفاصيل الصفحة</h2>
        </div>
        <v-btn icon="mdi-close" variant="text" @click="showViewDialog = false" class="close-btn" />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedPage" class="pa-6">
        <v-row>
          <v-col cols="12" class="text-center mb-4">
            <v-avatar size="80" color="primary">
              <v-icon color="white" size="40">{{ selectedPage.icon || 'mdi-file' }}</v-icon>
            </v-avatar>
            <h3 class="mt-3">{{ selectedPage.name }}</h3>
            <v-chip
              :color="selectedPage.status === 'active' ? 'success' : 'warning'"
              size="small"
              class="mt-2"
            >
              {{ selectedPage.status === 'active' ? 'نشط' : 'غير نشط' }}
            </v-chip>
          </v-col>

          <v-col cols="12">
            <v-list density="compact">
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="primary">mdi-link</v-icon>
                </template>
                <v-list-item-title>المسار</v-list-item-title>
                <v-list-item-subtitle dir="ltr" class="text-left">{{ selectedPage.route }}</v-list-item-subtitle>
              </v-list-item>

              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="success">mdi-emoticon</v-icon>
                </template>
                <v-list-item-title>الأيقونة</v-list-item-title>
                <v-list-item-subtitle>{{ selectedPage.icon || 'غير محدد' }}</v-list-item-subtitle>
              </v-list-item>

              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="info">mdi-calendar</v-icon>
                </template>
                <v-list-item-title>تاريخ الإنشاء</v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(selectedPage.createdAt) }}</v-list-item-subtitle>
              </v-list-item>
            </v-list>
          </v-col>
        </v-row>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn color="primary" variant="elevated" @click="showViewDialog = false">
          إغلاق
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- Delete Confirmation Dialog -->
  <v-dialog v-model="showDeleteDialog" max-width="500px">
    <v-card class="delete-confirm-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="error" class="me-3">mdi-delete-alert</v-icon>
          <h2>تأكيد الحذف</h2>
        </div>
        <v-btn icon="mdi-close" variant="text" @click="showDeleteDialog = false" class="close-btn" />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedPage" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60" color="error">
            <v-icon color="white" size="30">{{ selectedPage.icon || 'mdi-file' }}</v-icon>
          </v-avatar>
          <h4 class="mt-2">{{ selectedPage.name }}</h4>
          <p class="text-caption">{{ selectedPage.route }}</p>
        </div>

        <v-alert type="error" variant="tonal" class="mb-4">
          تحذير: هذا الإجراء لا يمكن التراجع عنه! سيتم إزالة هذه الصفحة من جميع الأدوار.
        </v-alert>

        <p class="text-body-2 text-center">
          هل أنت متأكد من حذف هذه الصفحة نهائياً؟
        </p>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn color="grey" variant="outlined" @click="showDeleteDialog = false" class="me-2">
          إلغاء
        </v-btn>
        <v-btn color="error" variant="elevated" @click="deletePage" :loading="deleting">
          حذف نهائي
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- Snackbar -->
  <v-snackbar v-model="snackbar.show" :color="snackbar.color" :timeout="3000">
    {{ snackbar.message }}
  </v-snackbar>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { listPages, createPage, updatePage, deletePage as deletePageApi } from '@/api/pages'
import { usePermissions } from '@/composables/usePermissions'

// Permissions
const { canCreate, canUpdate, canDelete } = usePermissions('/pages')

// Data
const loading = ref(false)
const pages = ref([])

// Dialog states
const showPageDialog = ref(false)
const showViewDialog = ref(false)
const showDeleteDialog = ref(false)
const isEditing = ref(false)
const selectedPage = ref(null)

// Form data
const pageForm = ref(null)
const formValid = ref(false)
const saving = ref(false)
const deleting = ref(false)

const pageData = reactive({
  name: '',
  route: '',
  icon: '',
  status: 'active'
})

// Snackbar
const snackbar = reactive({
  show: false,
  message: '',
  color: 'success'
})

// Options
const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' }
]

const iconOptions = [
  'mdi-view-dashboard',
  'mdi-folder-multiple',
  'mdi-clipboard-list',
  'mdi-account-hard-hat',
  'mdi-tag-multiple',
  'mdi-chart-line',
  'mdi-format-list-bulleted-type',
  'mdi-trending-up',
  'mdi-credit-card',
  'mdi-account-multiple',
  'mdi-account-group',
  'mdi-fingerprint',
  'mdi-shield-account',
  'mdi-file-document-multiple',
  'mdi-cog',
  'mdi-bell',
  'mdi-calendar',
  'mdi-chart-bar',
  'mdi-home',
  'mdi-email'
]

// Validation rules
const routeRules = [
  v => !!v || 'مسار الصفحة مطلوب',
  v => /^\//.test(v) || 'المسار يجب أن يبدأ بـ /'
]

// Table headers
const headers = ref([
  { title: 'الصفحة', key: 'name', sortable: true },
  { title: 'الأيقونة', key: 'icon', sortable: false },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'تاريخ الإنشاء', key: 'createdAt', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
])

// Computed
const activePages = computed(() => pages.value.filter(p => p.status === 'active').length)
const inactivePages = computed(() => pages.value.filter(p => p.status !== 'active').length)

// Fetch data
async function fetchData() {
  loading.value = true
  try {
    const result = await listPages()
    pages.value = Array.isArray(result) ? result : result?.data || []
  } catch (error) {
    console.error('Failed to fetch pages:', error)
    showSnackbar('فشل في تحميل البيانات', 'error')
  } finally {
    loading.value = false
  }
}

// Format date
function formatDate(date) {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

// Dialog functions
function openAddPageDialog() {
  isEditing.value = false
  pageData.name = ''
  pageData.route = ''
  pageData.icon = ''
  pageData.status = 'active'
  showPageDialog.value = true
}

function editPage(page) {
  isEditing.value = true
  selectedPage.value = page
  pageData.name = page.name
  pageData.route = page.route
  pageData.icon = page.icon || ''
  pageData.status = page.status || 'active'
  showPageDialog.value = true
}

function closePageDialog() {
  showPageDialog.value = false
  pageData.name = ''
  pageData.route = ''
  pageData.icon = ''
  pageData.status = 'active'
  selectedPage.value = null
}

function viewPage(page) {
  selectedPage.value = page
  showViewDialog.value = true
}

function confirmDeletePage(page) {
  selectedPage.value = page
  showDeleteDialog.value = true
}

// Save functions
async function savePage() {
  if (!pageForm.value?.validate()) return

  saving.value = true
  try {
    if (isEditing.value) {
      await updatePage(selectedPage.value.id, pageData)
      showSnackbar('تم تحديث الصفحة بنجاح', 'success')
    } else {
      await createPage(pageData)
      showSnackbar('تم إنشاء الصفحة بنجاح', 'success')
    }
    closePageDialog()
    await fetchData()
  } catch (error) {
    console.error('Failed to save page:', error)
    showSnackbar('فشل في حفظ الصفحة', 'error')
  } finally {
    saving.value = false
  }
}

async function deletePage() {
  if (!selectedPage.value) return

  deleting.value = true
  try {
    await deletePageApi(selectedPage.value.id)
    showSnackbar('تم حذف الصفحة بنجاح', 'success')
    showDeleteDialog.value = false
    selectedPage.value = null
    await fetchData()
  } catch (error) {
    console.error('Failed to delete page:', error)
    showSnackbar('فشل في حذف الصفحة', 'error')
  } finally {
    deleting.value = false
  }
}

function showSnackbar(message, color = 'success') {
  snackbar.message = message
  snackbar.color = color
  snackbar.show = true
}

// Mount
onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.data-page {
  background: #f5f5f5;
  min-height: 100vh;
}

.engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 24px;
  position: relative;
  overflow: hidden;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #64b5f6, #1976d2, #64b5f6);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.engineer-emoji {
  width: 64px;
  height: 64px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.main-title {
  color: white;
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
}

.subtitle {
  color: rgba(255, 255, 255, 0.8);
  font-size: 0.875rem;
  margin: 0;
}

.modern-stat-card {
  border-radius: 16px;
  position: relative;
  overflow: hidden;
}

.stat-card-primary {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
}

.stat-card-success {
  background: linear-gradient(135deg, #43a047 0%, #2e7d32 100%);
}

.stat-card-warning {
  background: linear-gradient(135deg, #fb8c00 0%, #f57c00 100%);
}

.stat-card-content {
  display: flex;
  align-items: center;
  padding: 20px;
  position: relative;
  z-index: 1;
}

.stat-icon-wrapper {
  margin-left: 16px;
}

.stat-icon {
  color: rgba(255, 255, 255, 0.9);
}

.stat-info {
  color: white;
}

.stat-value {
  font-size: 2rem;
  font-weight: 700;
  margin: 0;
}

.stat-label {
  font-size: 0.875rem;
  opacity: 0.9;
  margin: 0;
}

.table-title-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  color: white;
  padding: 16px 20px;
}

.title-text {
  font-weight: 600;
}

.table-spacer {
  height: 8px;
}

.view-btn {
  background: #e3f2fd !important;
  color: #1976d2 !important;
}

.edit-btn {
  background: #e8f5e9 !important;
  color: #43a047 !important;
}

.delete-btn {
  background: #ffebee !important;
  color: #e53935 !important;
}

.dialog-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.dialog-title {
  display: flex;
  align-items: center;
}

.close-btn {
  color: white !important;
}

.dialog-actions {
  background: #f5f5f5;
}

/* تحسين قراءة النصوص في الجدول */
.users-data-table :deep(.v-data-table__tr td) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
}

.users-data-table :deep(.font-weight-medium) {
  color: #1a1a1a !important;
  font-weight: 600 !important;
}

.users-data-table :deep(.text-caption) {
  color: #555555 !important;
}

.users-data-table :deep(.text-grey) {
  color: #666666 !important;
}

.users-data-table :deep(.text-body-2) {
  color: #333333 !important;
}

/* تحسين قراءة النصوص في الديالوج */
.view-user-dialog h3,
.view-user-dialog h4 {
  color: #1a1a1a !important;
}

.view-user-dialog .text-caption {
  color: #555555 !important;
}

.view-user-dialog .v-list-item-title {
  color: #1a1a1a !important;
  font-weight: 600 !important;
}

.view-user-dialog .v-list-item-subtitle {
  color: #444444 !important;
}

/* تحسين قراءة النصوص في نموذج الإضافة/التعديل */
.add-user-dialog .dialog-content {
  background: #ffffff !important;
}

.add-user-dialog :deep(.v-field__input) {
  color: #1a1a1a !important;
  font-weight: 500 !important;
}

.add-user-dialog :deep(.v-field input) {
  color: #1a1a1a !important;
}

.add-user-dialog :deep(.v-field textarea) {
  color: #1a1a1a !important;
}

.add-user-dialog :deep(.v-label) {
  color: #333333 !important;
  font-weight: 600 !important;
}

.add-user-dialog :deep(.v-field--focused .v-label) {
  color: #1976d2 !important;
}

.add-user-dialog :deep(.v-select__selection-text) {
  color: #1a1a1a !important;
}

.add-user-dialog :deep(.v-field__input input::placeholder) {
  color: #888888 !important;
}

.add-user-dialog :deep(.v-combobox .v-field__input) {
  color: #1a1a1a !important;
}

.add-user-dialog :deep(.v-list-item-title) {
  color: #1a1a1a !important;
}
</style>
