<template>
  <div class="fill-height data-page">
    <v-container fluid class="pa-6" style="padding: 0 20px !important;">
      <!-- Header -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-shield-account</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">إدارة الأدوار والصلاحيات</h1>
              <p class="subtitle">إنشاء وتعديل الأدوار وتعيين الصفحات والصلاحيات</p>
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
                <v-icon size="48" class="stat-icon">mdi-shield-account</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ roles.length }}</h3>
                <p class="stat-label">إجمالي الأدوار</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
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
      </v-row>

      <!-- Roles Table -->
      <v-card class="users-table" elevation="2">
        <v-card-title class="table-title-header d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" size="18" style="color: #ffffff !important;">mdi-table</v-icon>
            <span class="title-text">قائمة الأدوار</span>
          </div>
          <div class="d-flex table-header-buttons" style="gap: 0.5rem;">
            <v-btn
              v-if="canCreate"
              class="add-button add-user-btn btn-glow light-sweep smooth-transition"
              @click="openAddRoleDialog"
              elevation="2"
              color="primary"
              size="small"
              style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
            >
              <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
              إضافة دور جديد
            </v-btn>
          </div>
        </v-card-title>
        <div class="table-spacer"></div>
        <v-data-table
          :headers="headers"
          :items="roles"
          :loading="loading"
          class="elevation-0 users-data-table"
          :items-per-page="-1"
          hide-default-footer
        >
          <template v-slot:item.name="{ item }">
            <div class="d-flex align-center">
              <v-avatar size="36" color="primary" class="me-3">
                <v-icon color="white" size="20">mdi-shield</v-icon>
              </v-avatar>
              <div>
                <div class="font-weight-medium">{{ item.name }}</div>
                <div class="text-caption text-grey">{{ item.description || 'لا يوجد وصف' }}</div>
              </div>
            </div>
          </template>

          <template v-slot:item.pagesCount="{ item }">
            <v-chip color="info" size="small" variant="flat">
              {{ getRolePagesCount(item.id) }} صفحة
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
              @click="viewRole(item)"
            />
            <v-btn
              v-if="canUpdate"
              icon="mdi-pencil"
              size="small"
              variant="elevated"
              class="edit-btn me-1"
              @click="editRole(item)"
            />
            <v-btn
              v-if="canUpdate"
              icon="mdi-key"
              size="small"
              variant="elevated"
              class="reset-btn me-1"
              @click="openPermissionsDialog(item)"
              title="إدارة الصلاحيات"
            />
            <v-btn
              v-if="canDelete"
              icon="mdi-delete"
              size="small"
              variant="elevated"
              class="delete-btn"
              @click="confirmDeleteRole(item)"
            />
          </template>
        </v-data-table>
      </v-card>
    </v-container>
  </div>

  <!-- Add/Edit Role Dialog -->
  <v-dialog v-model="showRoleDialog" max-width="600px" persistent>
    <v-card class="add-user-dialog" rounded="lg">
      <v-card-title class="dialog-header pa-4">
        <div class="d-flex align-center">
          <v-avatar color="white" size="44" class="me-3">
            <v-icon color="primary" size="26">{{ isEditing ? 'mdi-pencil' : 'mdi-plus' }}</v-icon>
          </v-avatar>
          <div>
            <h3 class="text-h5 font-weight-bold text-white mb-0">
              {{ isEditing ? 'تعديل الدور' : 'إضافة دور جديد' }}
            </h3>
            <span class="text-body-2 text-white-darken-1">
              {{ isEditing ? 'تعديل بيانات الدور' : 'أدخل بيانات الدور الجديد' }}
            </span>
          </div>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          size="default"
          @click="closeRoleDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-card-text class="dialog-content pa-8">
        <v-form ref="roleForm" v-model="formValid" lazy-validation>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="roleData.name"
                label="اسم الدور"
                placeholder="مثال: مدير المشاريع"
                :rules="[v => !!v || 'اسم الدور مطلوب']"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-shield"
                color="primary"
                bg-color="grey-lighten-5"
                class="mb-2"
              />
            </v-col>
            <v-col cols="12">
              <v-textarea
                v-model="roleData.description"
                label="وصف الدور"
                placeholder="وصف مختصر للدور وصلاحياته"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-text"
                color="primary"
                bg-color="grey-lighten-5"
                rows="3"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-card-actions class="dialog-actions pa-5">
        <v-spacer />
        <v-btn variant="outlined" size="large" @click="closeRoleDialog" class="me-3 px-8">
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          size="large"
          @click="saveRole"
          :loading="saving"
          :disabled="!formValid"
          class="px-8"
          prepend-icon="mdi-check"
        >
          <v-icon start size="16">mdi-content-save</v-icon>
          {{ isEditing ? 'حفظ التعديلات' : 'حفظ الدور' }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- View Role Dialog -->
  <v-dialog v-model="showViewDialog" max-width="700px">
    <v-card class="view-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="primary" class="me-3">mdi-shield-account</v-icon>
          <h2>تفاصيل الدور</h2>
        </div>
        <v-btn icon="mdi-close" variant="text" @click="showViewDialog = false" class="close-btn" />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedRole" class="pa-6">
        <v-row>
          <v-col cols="12" class="text-center mb-4">
            <v-avatar size="80" color="primary">
              <v-icon color="white" size="40">mdi-shield</v-icon>
            </v-avatar>
            <h3 class="mt-3">{{ selectedRole.name }}</h3>
            <p class="text-caption">{{ selectedRole.description || 'لا يوجد وصف' }}</p>
          </v-col>

          <v-col cols="12">
            <h4 class="mb-3">الصفحات والصلاحيات:</h4>
            <v-list density="compact">
              <v-list-item
                v-for="rp in getViewRolePages()"
                :key="rp.id"
                class="mb-2"
              >
                <template v-slot:prepend>
                  <v-icon color="primary">{{ rp.page?.icon || 'mdi-file' }}</v-icon>
                </template>
                <v-list-item-title>{{ rp.page?.name || 'صفحة غير معروفة' }}</v-list-item-title>
                <v-list-item-subtitle>
                  <v-chip
                    v-for="perm in rp.permissions"
                    :key="perm"
                    size="x-small"
                    :color="getPermissionColor(perm)"
                    class="me-1"
                  >
                    {{ getPermissionText(perm) }}
                  </v-chip>
                </v-list-item-subtitle>
              </v-list-item>
              <v-list-item v-if="getViewRolePages().length === 0">
                <v-list-item-title class="text-grey">لا توجد صفحات مخصصة لهذا الدور</v-list-item-title>
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

  <!-- Permissions Dialog -->
  <v-dialog v-model="showPermissionsDialog" max-width="900px" persistent>
    <v-card class="add-user-dialog" rounded="lg">
      <v-card-title class="dialog-header pa-4">
        <div class="d-flex align-center">
          <v-avatar color="white" size="44" class="me-3">
            <v-icon color="primary" size="26">mdi-key</v-icon>
          </v-avatar>
          <div>
            <h3 class="text-h5 font-weight-bold text-white mb-0">
              إدارة صلاحيات الدور
            </h3>
            <span class="text-body-2 text-white-darken-1">
              {{ selectedRole?.name }} - تعيين الصفحات والصلاحيات
            </span>
          </div>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          size="default"
          @click="closePermissionsDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-card-text class="dialog-content pa-6" style="max-height: 70vh; overflow-y: auto;">
        <v-alert type="info" variant="tonal" class="mb-4">
          حدد الصفحات التي يمكن لهذا الدور الوصول إليها، واختر الصلاحيات المناسبة لكل صفحة
        </v-alert>

        <v-expansion-panels variant="accordion" class="permissions-panels">
          <v-expansion-panel
            v-for="page in pages"
            :key="page.id"
            class="mb-2"
          >
            <v-expansion-panel-title class="page-panel-title">
              <div class="d-flex align-center justify-space-between w-100">
                <div class="d-flex align-center">
                  <v-checkbox
                    v-model="pageEnabled[page.id]"
                    hide-details
                    density="compact"
                    color="primary"
                    @click.stop
                    @change="togglePagePermissions(page.id)"
                  />
                  <v-icon class="me-2" size="24" color="primary">{{ page.icon || 'mdi-file' }}</v-icon>
                  <div>
                    <span class="font-weight-bold">{{ page.name }}</span>
                    <span class="text-caption text-grey ms-2">({{ page.route }})</span>
                  </div>
                </div>
                <v-chip
                  v-if="getSelectedPermissionsCount(page.id) > 0"
                  color="primary"
                  size="small"
                  class="me-4"
                >
                  {{ getSelectedPermissionsCount(page.id) }} صلاحية
                </v-chip>
              </div>
            </v-expansion-panel-title>
            <v-expansion-panel-text class="page-panel-content">
              <div class="permissions-grid">
                <div
                  v-for="perm in getPermissionsForPage(page)"
                  :key="perm.value"
                  class="permission-checkbox"
                >
                  <v-checkbox-btn
                    v-model="pagePermissions[page.id]"
                    :value="perm.value"
                    :color="perm.color"
                  />
                  <span class="permission-label">{{ perm.label }}</span>
                </div>
              </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-card-text>

      <v-card-actions class="dialog-actions pa-5">
        <v-spacer />
        <v-btn variant="outlined" size="large" @click="closePermissionsDialog" class="me-3 px-8">
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          size="large"
          @click="savePermissions"
          :loading="savingPermissions"
          class="px-8"
          prepend-icon="mdi-check"
        >
          <v-icon start size="16">mdi-content-save</v-icon>
          حفظ الصلاحيات
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

      <v-card-text v-if="selectedRole" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60" color="error">
            <v-icon color="white" size="30">mdi-shield</v-icon>
          </v-avatar>
          <h4 class="mt-2">{{ selectedRole.name }}</h4>
        </div>

        <v-alert type="error" variant="tonal" class="mb-4">
          تحذير: هذا الإجراء لا يمكن التراجع عنه!
        </v-alert>

        <p class="text-body-2 text-center">
          هل أنت متأكد من حذف هذا الدور نهائياً؟
        </p>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn color="grey" variant="outlined" @click="showDeleteDialog = false" class="me-2">
          إلغاء
        </v-btn>
        <v-btn color="error" variant="elevated" @click="deleteRole" :loading="deleting">
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
import { ref, reactive, onMounted } from 'vue'
import { listRoles, createRole, updateRole, deleteRole as deleteRoleApi } from '@/api/roles'
import { listPages } from '@/api/pages'
import { listRolePages, getRolePagesByRoleId, bulkUpdateRolePages } from '@/api/rolePages'
import { usePermissions } from '@/composables/usePermissions'

// Permissions
const { canCreate, canUpdate, canDelete } = usePermissions('/roles')

// Data
const loading = ref(false)
const roles = ref([])
const pages = ref([])
const rolePages = ref([])

// Dialog states
const showRoleDialog = ref(false)
const showViewDialog = ref(false)
const showPermissionsDialog = ref(false)
const showDeleteDialog = ref(false)
const isEditing = ref(false)
const selectedRole = ref(null)

// Form data
const roleForm = ref(null)
const formValid = ref(false)
const saving = ref(false)
const deleting = ref(false)
const savingPermissions = ref(false)

const roleData = reactive({
  name: '',
  description: ''
})

// Permissions matrix - pageId -> array of permissions
const pagePermissions = reactive({})

// Track which pages are enabled (have at least one permission)
const pageEnabled = reactive({})

// Base permissions for all pages
const basePermissions = [
  { value: 'read', label: 'قراءة', color: 'success' },
  { value: 'create', label: 'إنشاء', color: 'primary' },
  { value: 'update', label: 'تعديل', color: 'warning' },
  { value: 'delete', label: 'حذف', color: 'error' }
]

// Extra permissions for specific pages
const extraPermissions = {
  '/users': [
    { value: 'updatePassword', label: 'تغيير كلمة المرور', color: 'purple' }
  ]
}

// Get available permissions for a specific page
const getPermissionsForPage = (page) => {
  const extra = extraPermissions[page.route] || []
  return [...basePermissions, ...extra]
}

// Snackbar
const snackbar = reactive({
  show: false,
  message: '',
  color: 'success'
})

// Table headers
const headers = ref([
  { title: 'الدور', key: 'name', sortable: true },
  { title: 'عدد الصفحات', key: 'pagesCount', sortable: false },
  { title: 'تاريخ الإنشاء', key: 'createdAt', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
])

// Fetch data
async function fetchData() {
  loading.value = true
  try {
    const [rolesData, pagesData, rolePagesData] = await Promise.all([
      listRoles(),
      listPages(),
      listRolePages()
    ])
    roles.value = Array.isArray(rolesData) ? rolesData : rolesData?.data || []
    pages.value = Array.isArray(pagesData) ? pagesData : pagesData?.data || []
    rolePages.value = Array.isArray(rolePagesData) ? rolePagesData : rolePagesData?.data || []
  } catch (error) {
    console.error('Failed to fetch data:', error)
    showSnackbar('فشل في تحميل البيانات', 'error')
  } finally {
    loading.value = false
  }
}

// Get count of pages for a role
function getRolePagesCount(roleId) {
  return rolePages.value.filter(rp => rp.roleId === roleId).length
}

// Get role pages for view dialog
function getViewRolePages() {
  if (!selectedRole.value) return []
  const rps = rolePages.value.filter(rp => rp.roleId === selectedRole.value.id)
  return rps.map(rp => ({
    ...rp,
    page: pages.value.find(p => p.id === rp.pageId)
  }))
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

// Permission helpers
function getPermissionColor(perm) {
  const colors = {
    read: 'success',
    create: 'primary',
    update: 'warning',
    delete: 'error'
  }
  return colors[perm] || 'grey'
}

function getPermissionText(perm) {
  const texts = {
    read: 'قراءة',
    create: 'إنشاء',
    update: 'تعديل',
    delete: 'حذف'
  }
  return texts[perm] || perm
}

// Toggle all permissions for a page
function togglePagePermissions(pageId) {
  if (pageEnabled[pageId]) {
    // Enable with read permission by default
    if (!pagePermissions[pageId] || pagePermissions[pageId].length === 0) {
      pagePermissions[pageId] = ['read']
    }
  } else {
    // Disable - clear all permissions
    pagePermissions[pageId] = []
  }
}

// Get count of selected permissions for a page
function getSelectedPermissionsCount(pageId) {
  return pagePermissions[pageId]?.length || 0
}

// Dialog functions
function openAddRoleDialog() {
  isEditing.value = false
  roleData.name = ''
  roleData.description = ''
  showRoleDialog.value = true
}

function editRole(role) {
  isEditing.value = true
  selectedRole.value = role
  roleData.name = role.name
  roleData.description = role.description || ''
  showRoleDialog.value = true
}

function closeRoleDialog() {
  showRoleDialog.value = false
  roleData.name = ''
  roleData.description = ''
  selectedRole.value = null
}

function viewRole(role) {
  selectedRole.value = role
  showViewDialog.value = true
}

function confirmDeleteRole(role) {
  selectedRole.value = role
  showDeleteDialog.value = true
}

async function openPermissionsDialog(role) {
  selectedRole.value = role

  // Reset permissions and pageEnabled
  Object.keys(pagePermissions).forEach(key => {
    delete pagePermissions[key]
  })
  Object.keys(pageEnabled).forEach(key => {
    delete pageEnabled[key]
  })

  // Initialize all pages with empty arrays
  pages.value.forEach(page => {
    pagePermissions[page.id] = []
    pageEnabled[page.id] = false
  })

  // Load existing permissions for this role
  try {
    const existingRolePages = await getRolePagesByRoleId(role.id)
    existingRolePages.forEach(rp => {
      if (rp.pageId && rp.permissions) {
        pagePermissions[rp.pageId] = [...rp.permissions]
        pageEnabled[rp.pageId] = rp.permissions.length > 0
      }
    })
  } catch (error) {
    console.error('Failed to load role permissions:', error)
  }

  showPermissionsDialog.value = true
}

function closePermissionsDialog() {
  showPermissionsDialog.value = false
  selectedRole.value = null
}

// Save functions
async function saveRole() {
  if (!roleForm.value?.validate()) return

  saving.value = true
  try {
    if (isEditing.value) {
      await updateRole(selectedRole.value.id, roleData)
      showSnackbar('تم تحديث الدور بنجاح', 'success')
    } else {
      await createRole(roleData)
      showSnackbar('تم إنشاء الدور بنجاح', 'success')
    }
    closeRoleDialog()
    await fetchData()
  } catch (error) {
    console.error('Failed to save role:', error)
    showSnackbar('فشل في حفظ الدور', 'error')
  } finally {
    saving.value = false
  }
}

async function deleteRole() {
  if (!selectedRole.value) return

  deleting.value = true
  try {
    await deleteRoleApi(selectedRole.value.id)
    showSnackbar('تم حذف الدور بنجاح', 'success')
    showDeleteDialog.value = false
    selectedRole.value = null
    await fetchData()
  } catch (error) {
    console.error('Failed to delete role:', error)
    showSnackbar('فشل في حذف الدور', 'error')
  } finally {
    deleting.value = false
  }
}

async function savePermissions() {
  if (!selectedRole.value) return

  savingPermissions.value = true
  try {
    // Build page permissions array
    const permissionsToSave = []
    Object.entries(pagePermissions).forEach(([pageId, perms]) => {
      if (perms && perms.length > 0) {
        permissionsToSave.push({
          pageId: Number(pageId),
          permissions: perms
        })
      }
    })

    await bulkUpdateRolePages(selectedRole.value.id, permissionsToSave)
    showSnackbar('تم حفظ الصلاحيات بنجاح', 'success')
    closePermissionsDialog()
    await fetchData()
  } catch (error) {
    console.error('Failed to save permissions:', error)
    showSnackbar('فشل في حفظ الصلاحيات', 'error')
  } finally {
    savingPermissions.value = false
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

.reset-btn {
  background: #fff3e0 !important;
  color: #f57c00 !important;
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

/* تحسين قراءة النصوص في جدول الصلاحيات */
.v-table tbody td {
  color: #1a1a1a !important;
  background: #ffffff !important;
}

.v-table tbody tr:nth-child(even) td {
  background: #fafafa !important;
}

.v-table tbody tr:hover td {
  background: #f0f7ff !important;
}

.v-table thead th {
  color: #1a1a1a !important;
  font-weight: 700 !important;
  background: #e3f2fd !important;
  border-bottom: 2px solid #1976d2 !important;
}

/* تحسين مظهر الـ checkboxes في جدول الصلاحيات */
.v-table :deep(.v-checkbox) {
  display: flex;
  justify-content: center;
}

.v-table :deep(.v-selection-control) {
  min-height: 40px !important;
}

.v-table :deep(.v-selection-control__wrapper) {
  width: 24px !important;
  height: 24px !important;
}

.v-table :deep(.v-selection-control__input) {
  width: 24px !important;
  height: 24px !important;
}

.v-table :deep(.v-selection-control__input > .v-icon) {
  font-size: 24px !important;
}

/* ألوان الـ checkboxes */
.v-table :deep(.v-checkbox .v-selection-control--dirty .v-icon) {
  opacity: 1 !important;
}

.v-table :deep(.v-checkbox:not(.v-selection-control--dirty) .v-icon) {
  color: #9e9e9e !important;
}

/* تحسين ظهور اسم الصفحة والأيقونة */
.v-table tbody td .v-icon {
  color: #1976d2 !important;
}

.v-table tbody td span {
  color: #1a1a1a !important;
  font-weight: 500 !important;
}

.v-table tbody td .text-caption {
  color: #666666 !important;
  font-weight: 400 !important;
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

.add-user-dialog :deep(.v-list-item-title) {
  color: #1a1a1a !important;
}

/* Permissions Panels Styles */
.permissions-panels {
  border-radius: 12px;
}

.permissions-panels .v-expansion-panel {
  background: #ffffff !important;
  border: 1px solid #e0e0e0;
  border-radius: 8px !important;
  margin-bottom: 8px;
}

.permissions-panels .v-expansion-panel--active {
  border-color: #1976d2;
}

.permissions-panels :deep(.v-expansion-panel-title) {
  background: #fafafa !important;
  color: #1a1a1a !important;
}

.permissions-panels :deep(.v-expansion-panel-title__overlay) {
  background: transparent !important;
}

.permissions-panels :deep(.v-expansion-panel-title .font-weight-bold) {
  color: #1a1a1a !important;
}

.permissions-panels :deep(.v-expansion-panel-title .text-caption) {
  color: #666666 !important;
}

.page-panel-title {
  min-height: 56px;
}

.page-panel-content {
  background: #ffffff !important;
  padding: 16px !important;
}

.permissions-panels :deep(.v-expansion-panel-text__wrapper) {
  background: #ffffff !important;
  padding: 16px !important;
}

.permissions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 12px;
}

.permission-checkbox {
  background: #f5f5f5 !important;
  border-radius: 8px;
  padding: 8px 16px;
  border: 1px solid #e0e0e0;
  transition: all 0.2s;
  margin: 0 !important;
  display: flex;
  flex-direction: row-reverse;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  cursor: pointer;
}

.permission-checkbox:hover {
  border-color: #1976d2;
  background: #e3f2fd !important;
}

.permission-checkbox .permission-label {
  color: #1a1a1a !important;
  font-weight: 500;
  font-size: 14px;
  flex: 1;
}

.permission-checkbox :deep(.v-selection-control) {
  min-height: unset !important;
  flex: none;
}
</style>
