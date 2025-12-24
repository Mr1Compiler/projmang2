<template>
  <div class="fill-height data-page">
    <v-container fluid class="pa-6" style="padding: 0 20px !important;">
      <!-- Header -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-account-group</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">إدارة فرق العمل</h1>
              <p class="subtitle">إضافة وإدارة أعضاء الفريق في المشاريع</p>
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
                <v-icon size="48" class="stat-icon">mdi-account-multiple</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ stats.total || 0 }}</h3>
                <p class="stat-label">إجمالي الأعضاء</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-success" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-account-check</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ stats.uniqueUsers || 0 }}</h3>
                <p class="stat-label">المستخدمين الفريدين</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-warning" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-folder-multiple</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ stats.uniqueProjects || 0 }}</h3>
                <p class="stat-label">المشاريع النشطة</p>
              </div>
            </div>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="modern-stat-card stat-card-info" elevation="0">
            <div class="stat-card-background"></div>
            <div class="stat-card-content">
              <div class="stat-icon-wrapper">
                <v-icon size="48" class="stat-icon">mdi-chart-bar</v-icon>
              </div>
              <div class="stat-info">
                <h3 class="stat-value">{{ (stats.avgPerProject || 0).toFixed(1) }}</h3>
                <p class="stat-label">متوسط الأعضاء/مشروع</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </v-row>

      <!-- Team Members Table -->
      <v-card class="users-table" elevation="2">
        <v-card-title class="table-title-header d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" size="18" style="color: #ffffff !important;">mdi-table</v-icon>
            <span class="title-text">قائمة أعضاء الفرق</span>
          </div>
          <div class="d-flex table-header-buttons" style="gap: 0.5rem;">
            <!-- Filter by Project -->
            <v-select
              v-model="filterProjectId"
              :items="projectsForFilter"
              item-title="name"
              item-value="id"
              label="تصفية حسب المشروع"
              variant="outlined"
              density="compact"
              clearable
              hide-details
              class="filter-select"
              style="min-width: 200px; background: white; border-radius: 8px;"
            />
            <v-btn
              v-if="canCreate"
              class="add-button add-user-btn btn-glow light-sweep smooth-transition"
              @click="openAddDialog"
              elevation="2"
              color="primary"
              size="small"
              style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
            >
              <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
              إضافة عضو للمشروع
            </v-btn>
          </div>
        </v-card-title>
        <div class="table-spacer"></div>
        <v-data-table
          :headers="headers"
          :items="filteredTeamMembers"
          :loading="loading"
          class="elevation-0 users-data-table"
          :items-per-page="-1"
          hide-default-footer
        >
          <template v-slot:item.project="{ item }">
            <div class="d-flex align-center">
              <v-avatar size="36" color="primary" class="me-3">
                <v-icon color="white" size="20">mdi-folder</v-icon>
              </v-avatar>
              <div>
                <div class="font-weight-medium">{{ getProjectName(item.projectId) }}</div>
              </div>
            </div>
          </template>

          <template v-slot:item.user="{ item }">
            <div class="d-flex align-center">
              <v-avatar size="36" color="success" class="me-3">
                <v-icon color="white" size="20">mdi-account</v-icon>
              </v-avatar>
              <div>
                <div class="font-weight-medium">{{ getUserName(item.userId) }}</div>
                <div class="text-caption text-grey">{{ getUserJobTitle(item.userId) }}</div>
              </div>
            </div>
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
              @click="viewMember(item)"
            />
            <v-btn
              v-if="canDelete"
              icon="mdi-delete"
              size="small"
              variant="elevated"
              class="delete-btn"
              @click="confirmDeleteMember(item)"
            />
          </template>
        </v-data-table>
      </v-card>
    </v-container>
  </div>

  <!-- Add Member Dialog -->
  <v-dialog v-model="showAddDialog" max-width="600px" persistent>
    <v-card class="add-user-dialog" rounded="lg">
      <v-card-title class="dialog-header pa-4">
        <div class="d-flex align-center">
          <v-avatar color="white" size="44" class="me-3">
            <v-icon color="primary" size="26">mdi-account-plus</v-icon>
          </v-avatar>
          <div>
            <h3 class="text-h5 font-weight-bold text-white mb-0">
              إضافة عضو للمشروع
            </h3>
            <span class="text-body-2 text-white-darken-1">
              اختر المشروع والمستخدم لإضافته للفريق
            </span>
          </div>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          size="default"
          @click="closeAddDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-card-text class="dialog-content pa-8">
        <v-form ref="addForm" v-model="formValid" lazy-validation>
          <v-row>
            <v-col cols="12">
              <v-select
                v-model="formData.projectId"
                :items="projects"
                item-title="name"
                item-value="id"
                label="المشروع"
                placeholder="اختر المشروع"
                :rules="[v => !!v || 'المشروع مطلوب']"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-folder"
                color="primary"
                bg-color="grey-lighten-5"
                class="mb-2"
              />
            </v-col>
            <v-col cols="12">
              <v-select
                v-model="formData.userId"
                :items="availableUsers"
                item-title="fullName"
                item-value="id"
                label="المستخدم"
                placeholder="اختر المستخدم"
                :rules="[v => !!v || 'المستخدم مطلوب']"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-account"
                color="primary"
                bg-color="grey-lighten-5"
                :disabled="!formData.projectId"
                :hint="!formData.projectId ? 'اختر المشروع أولاً' : ''"
                persistent-hint
              >
                <template v-slot:item="{ item, props }">
                  <v-list-item v-bind="props">
                    <template v-slot:prepend>
                      <v-avatar size="32" color="primary">
                        <v-icon color="white" size="18">mdi-account</v-icon>
                      </v-avatar>
                    </template>
                    <template v-slot:subtitle>
                      {{ item.raw.jobTitle || 'لا يوجد مسمى وظيفي' }}
                    </template>
                  </v-list-item>
                </template>
              </v-select>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-card-actions class="dialog-actions pa-5">
        <v-spacer />
        <v-btn variant="outlined" size="large" @click="closeAddDialog" class="me-3 px-8">
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          size="large"
          @click="saveMember"
          :loading="saving"
          :disabled="!formValid"
          class="px-8"
          prepend-icon="mdi-check"
        >
          <v-icon start size="16">mdi-content-save</v-icon>
          إضافة العضو
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- View Member Dialog -->
  <v-dialog v-model="showViewDialog" max-width="600px">
    <v-card class="view-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="primary" class="me-3">mdi-account-details</v-icon>
          <h2>تفاصيل العضو</h2>
        </div>
        <v-btn icon="mdi-close" variant="text" @click="showViewDialog = false" class="close-btn" />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedMember" class="pa-6">
        <v-row>
          <v-col cols="12" class="text-center mb-4">
            <v-avatar size="80" color="primary">
              <v-icon color="white" size="40">mdi-account</v-icon>
            </v-avatar>
            <h3 class="mt-3">{{ getUserName(selectedMember.userId) }}</h3>
            <p class="text-caption">{{ getUserJobTitle(selectedMember.userId) }}</p>
          </v-col>

          <v-col cols="12">
            <v-list density="compact">
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="primary">mdi-folder</v-icon>
                </template>
                <v-list-item-title>المشروع</v-list-item-title>
                <v-list-item-subtitle>{{ getProjectName(selectedMember.projectId) }}</v-list-item-subtitle>
              </v-list-item>
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="primary">mdi-calendar</v-icon>
                </template>
                <v-list-item-title>تاريخ الإضافة</v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(selectedMember.createdAt) }}</v-list-item-subtitle>
              </v-list-item>
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="primary">mdi-email</v-icon>
                </template>
                <v-list-item-title>البريد الإلكتروني</v-list-item-title>
                <v-list-item-subtitle>{{ getUserEmail(selectedMember.userId) }}</v-list-item-subtitle>
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

      <v-card-text v-if="selectedMember" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60" color="error">
            <v-icon color="white" size="30">mdi-account-remove</v-icon>
          </v-avatar>
          <h4 class="mt-2">{{ getUserName(selectedMember.userId) }}</h4>
          <p class="text-caption">{{ getProjectName(selectedMember.projectId) }}</p>
        </div>

        <v-alert type="error" variant="tonal" class="mb-4">
          تحذير: سيتم إزالة هذا العضو من المشروع!
        </v-alert>

        <p class="text-body-2 text-center">
          هل أنت متأكد من إزالة هذا العضو من فريق المشروع؟
        </p>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn color="grey" variant="outlined" @click="showDeleteDialog = false" class="me-2">
          إلغاء
        </v-btn>
        <v-btn color="error" variant="elevated" @click="deleteMember" :loading="deleting">
          إزالة العضو
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
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { listTeamMembers, getTeamMemberStats, createTeamMember, deleteTeamMember } from '@/api/teamMembers'
import { listProjects } from '@/api/projects'
import { listUsers } from '@/api/users'
import { usePermissions } from '@/composables/usePermissions'

// Permissions
const { canCreate, canUpdate, canDelete } = usePermissions('/team-members')

// Data
const loading = ref(false)
const teamMembers = ref([])
const projects = ref([])
const users = ref([])
const stats = ref({})
const filterProjectId = ref(null)

// Dialog states
const showAddDialog = ref(false)
const showViewDialog = ref(false)
const showDeleteDialog = ref(false)
const selectedMember = ref(null)

// Form data
const addForm = ref(null)
const formValid = ref(false)
const saving = ref(false)
const deleting = ref(false)

const formData = reactive({
  projectId: null,
  userId: null
})

// Snackbar
const snackbar = reactive({
  show: false,
  message: '',
  color: 'success'
})

// Table headers
const headers = ref([
  { title: 'المشروع', key: 'project', sortable: true },
  { title: 'العضو', key: 'user', sortable: true },
  { title: 'تاريخ الإضافة', key: 'createdAt', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
])

// Computed: Projects for filter (with "all" option)
const projectsForFilter = computed(() => {
  return [{ id: null, name: 'جميع المشاريع' }, ...projects.value]
})

// Computed: Filtered team members
const filteredTeamMembers = computed(() => {
  if (!filterProjectId.value) {
    return teamMembers.value
  }
  return teamMembers.value.filter(tm => tm.projectId === filterProjectId.value)
})

// Computed: Available users (not already in selected project)
const availableUsers = computed(() => {
  if (!formData.projectId) return users.value

  const projectMemberUserIds = teamMembers.value
    .filter(tm => tm.projectId === formData.projectId)
    .map(tm => tm.userId)

  return users.value.filter(user => !projectMemberUserIds.includes(user.id))
})

// Fetch data
async function fetchData() {
  loading.value = true
  try {
    const [teamMembersData, projectsData, usersData, statsData] = await Promise.all([
      listTeamMembers({ limit: 100 }),
      listProjects({ limit: 100 }),
      listUsers({ limit: 100 }),
      getTeamMemberStats()
    ])
    teamMembers.value = teamMembersData?.data || []
    projects.value = projectsData?.data || []
    users.value = usersData?.data || []
    stats.value = statsData || {}
  } catch (error) {
    console.error('Failed to fetch data:', error)
    showSnackbar('فشل في تحميل البيانات', 'error')
  } finally {
    loading.value = false
  }
}

// Helper functions
function getProjectName(projectId) {
  const project = projects.value.find(p => p.id === projectId)
  return project?.name || 'مشروع غير معروف'
}

function getUserName(userId) {
  const user = users.value.find(u => u.id === userId)
  return user?.fullName || 'مستخدم غير معروف'
}

function getUserJobTitle(userId) {
  const user = users.value.find(u => u.id === userId)
  return user?.jobTitle || 'لا يوجد مسمى وظيفي'
}

function getUserEmail(userId) {
  const user = users.value.find(u => u.id === userId)
  return user?.email || '-'
}

function formatDate(date) {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

// Dialog functions
function openAddDialog() {
  formData.projectId = null
  formData.userId = null
  showAddDialog.value = true
}

function closeAddDialog() {
  showAddDialog.value = false
  formData.projectId = null
  formData.userId = null
}

function viewMember(member) {
  selectedMember.value = member
  showViewDialog.value = true
}

function confirmDeleteMember(member) {
  selectedMember.value = member
  showDeleteDialog.value = true
}

// Save functions
async function saveMember() {
  if (!addForm.value?.validate()) return

  saving.value = true
  try {
    await createTeamMember({
      projectId: formData.projectId,
      userId: formData.userId
    })
    showSnackbar('تم إضافة العضو للمشروع بنجاح', 'success')
    closeAddDialog()
    await fetchData()
  } catch (error) {
    console.error('Failed to add team member:', error)
    showSnackbar('فشل في إضافة العضو للمشروع', 'error')
  } finally {
    saving.value = false
  }
}

async function deleteMember() {
  if (!selectedMember.value) return

  deleting.value = true
  try {
    await deleteTeamMember(selectedMember.value.id)
    showSnackbar('تم إزالة العضو من المشروع بنجاح', 'success')
    showDeleteDialog.value = false
    selectedMember.value = null
    await fetchData()
  } catch (error) {
    console.error('Failed to delete team member:', error)
    showSnackbar('فشل في إزالة العضو من المشروع', 'error')
  } finally {
    deleting.value = false
  }
}

function showSnackbar(message, color = 'success') {
  snackbar.message = message
  snackbar.color = color
  snackbar.show = true
}

// Reset userId when project changes
watch(() => formData.projectId, () => {
  formData.userId = null
})

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

.stat-card-info {
  background: linear-gradient(135deg, #00acc1 0%, #0097a7 100%);
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

.filter-select :deep(.v-field__input) {
  color: #1a1a1a !important;
}

.filter-select :deep(.v-label) {
  color: #666666 !important;
}

/* Table text improvements */
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

/* View dialog text improvements */
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

/* Add dialog text improvements */
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

.add-user-dialog :deep(.v-list-item-title) {
  color: #1a1a1a !important;
}

/* Delete dialog text improvements */
.delete-confirm-dialog h4 {
  color: #1a1a1a !important;
}

.delete-confirm-dialog .text-caption {
  color: #555555 !important;
}

.delete-confirm-dialog .text-body-2 {
  color: #333333 !important;
}
</style>
