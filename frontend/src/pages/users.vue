<template>
  <div>
    <!-- الشريط العلوي -->
    <v-app-bar flat height="70" class="top-bar">
      <!-- شريط البحث -->
      <v-text-field
        v-model="searchQuery"
        placeholder="البحث في المستخدمين"
        prepend-inner-icon="mdi-magnify"
        variant="outlined"
        density="compact"
        hide-details
        class="search-field"
        style="max-width: 400px;"
      />

      <v-spacer />

      <!-- زر إضافة مستخدم جديد -->
      <v-btn
        v-if="canWriteUsers"
        color="primary"
        prepend-icon="mdi-plus"
        class="me-3"
        @click="openAddUserDialog"
      >
        إضافة مستخدم
      </v-btn>

      <!-- الإشعارات -->
      <v-btn icon="mdi-bell" variant="text" class="me-2">
        <v-badge color="pink" dot floating />
      </v-btn>

      <!-- صورة المستخدم -->
      <v-avatar size="40" class="me-2">
        <v-img src="https://randomuser.me/api/portraits/men/1.jpg" />
      </v-avatar>
    </v-app-bar>

    <!-- المحتوى -->
    <div class="main-content pa-6">
      <!-- شريط العنوان الرئيسي -->
      <PageHeader
        title="إدارة المستخدمين"
        subtitle="نظام شامل لإدارة حسابات المستخدمين والصلاحيات"
        icon="👥"
      />

      <!-- الإحصائيات -->
      <UserStats
        :total-users="totalUsersCount"
        :active-users="activeUsersCount"
        :pending-users="pendingUsersCount"
        :admin-users="adminUsersCount"
      />

      <!-- فلاتر البحث -->
      <UserFilters
        v-model:selected-role="selectedRole"
        v-model:selected-status="selectedStatus"
        v-model:selected-department="selectedDepartment"
        :role-options="roleFilterOptions"
        :status-options="statusFilterOptions"
        :department-options="departmentFilterOptions"
        @apply-filters="applyFilters"
      />

      <!-- جدول المستخدمين -->
      <UserTable
        :users="filteredUsers"
        :search-query="searchQuery"
        :loading="loading"
        @view="viewUser"
        @edit="editUser"
        @reset-password="openResetPasswordDialog"
        @delete="openDeleteDialog"
        @export="exportData"
        @bulk-add="bulkAdd"
      />

      <!-- إحصائيات إضافية -->
      <v-row class="mt-6">
        <v-col cols="12" md="6">
          <v-card class="chart-card" elevation="2">
            <v-card-title class="text-h6 font-weight-bold">توزيع المستخدمين حسب الدور</v-card-title>
            <v-card-text>
              <div class="chart-placeholder">
                <v-icon size="64" color="primary">mdi-chart-pie</v-icon>
                <p class="text-body-1 mt-2">رسم بياني دائري للأدوار</p>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" md="6">
          <v-card class="chart-card" elevation="2">
            <v-card-title class="text-h6 font-weight-bold">نشاط المستخدمين</v-card-title>
            <v-card-text>
              <div class="chart-placeholder">
                <v-icon size="64" color="success">mdi-chart-line</v-icon>
                <p class="text-body-1 mt-2">رسم بياني خطي للنشاط</p>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <!-- نوافذ الحوار -->
    <UserForm
      v-model="showFormDialog"
      :user="selectedUser"
      :loading="saving"
      @save="saveUser"
    />

    <UserDetails
      v-model="showViewDialog"
      :user="selectedUser"
    />

    <UserDialogs
      v-model:reset-password-dialog="showResetPasswordDialog"
      v-model:delete-dialog="showDeleteDialog"
      :user="selectedUser"
      :reset-loading="resetLoading"
      :delete-loading="deleteLoading"
      @confirm-reset="confirmResetPassword"
      @confirm-delete="confirmDeleteUser"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { toast } from 'vue3-toastify'
import { useUsers } from '@/composables'
import { PageHeader } from '@/components/shared'
import {
  UserStats,
  UserFilters,
  UserTable,
  UserForm,
  UserDetails,
  UserDialogs
} from '@/components/users'

// Composable
const {
  users,
  loading,
  totalUsers,
  activeUsers,
  pendingUsers,
  adminUsers,
  canWriteUsers,
  canDeleteUsers,
  fetchUsers,
  createUser,
  updateUser,
  deleteUser,
  changePassword
} = useUsers()

// Search and filters
const searchQuery = ref('')
const selectedRole = ref(null)
const selectedStatus = ref(null)
const selectedDepartment = ref(null)

// Dialog states
const showFormDialog = ref(false)
const showViewDialog = ref(false)
const showResetPasswordDialog = ref(false)
const showDeleteDialog = ref(false)
const selectedUser = ref(null)

// Loading states
const saving = ref(false)
const resetLoading = ref(false)
const deleteLoading = ref(false)

// Filter options
const roleFilterOptions = [
  'مدير النظام',
  'مدير المشاريع',
  'مهندس',
  'محاسب',
  'موظف إداري',
  'مراجع',
  'مستخدم عادي'
]

const statusFilterOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' },
  { title: 'معلق', value: 'pending' },
  { title: 'محظور', value: 'banned' }
]

const departmentFilterOptions = [
  'تقنية المعلومات',
  'الموارد البشرية',
  'المحاسبة',
  'المشاريع',
  'التسويق',
  'الإدارة',
  'الصيانة'
]

// Computed statistics
const totalUsersCount = computed(() => totalUsers.value || users.value.length)
const activeUsersCount = computed(() => activeUsers.value)
const pendingUsersCount = computed(() => pendingUsers.value)
const adminUsersCount = computed(() => adminUsers.value)

// Filtered users
const filteredUsers = computed(() => {
  let filtered = users.value

  if (selectedRole.value) {
    filtered = filtered.filter(user => user.role === selectedRole.value)
  }

  if (selectedStatus.value) {
    filtered = filtered.filter(user => user.status === selectedStatus.value)
  }

  if (selectedDepartment.value) {
    filtered = filtered.filter(user => user.department === selectedDepartment.value)
  }

  return filtered
})

// Methods
const applyFilters = () => {
  fetchUsers({
    role: selectedRole.value,
    status: selectedStatus.value,
    department: selectedDepartment.value
  })
}

const openAddUserDialog = () => {
  selectedUser.value = null
  showFormDialog.value = true
}

const viewUser = (user) => {
  selectedUser.value = { ...user }
  showViewDialog.value = true
}

const editUser = (user) => {
  selectedUser.value = { ...user }
  showFormDialog.value = true
}

const openResetPasswordDialog = (user) => {
  selectedUser.value = { ...user }
  showResetPasswordDialog.value = true
}

const openDeleteDialog = (user) => {
  selectedUser.value = { ...user }
  showDeleteDialog.value = true
}

const saveUser = async (userData) => {
  saving.value = true

  try {
    if (userData.id) {
      // Update existing user
      await updateUser(userData.id, userData)
    } else {
      // Create new user
      await createUser(userData)
    }
    showFormDialog.value = false
    selectedUser.value = null
  } catch (error) {
    toast.error('حدث خطأ أثناء حفظ المستخدم')
  } finally {
    saving.value = false
  }
}

const confirmResetPassword = async () => {
  if (!selectedUser.value) return

  resetLoading.value = true

  try {
    await changePassword(selectedUser.value.id, null) // null triggers reset
    showResetPasswordDialog.value = false
    selectedUser.value = null
  } catch (error) {
    toast.error('حدث خطأ أثناء إعادة تعيين كلمة المرور')
  } finally {
    resetLoading.value = false
  }
}

const confirmDeleteUser = async () => {
  if (!selectedUser.value || !canDeleteUsers.value) return

  deleteLoading.value = true

  try {
    await deleteUser(selectedUser.value.id)
    showDeleteDialog.value = false
    selectedUser.value = null
  } catch (error) {
    toast.error('حدث خطأ أثناء حذف المستخدم')
  } finally {
    deleteLoading.value = false
  }
}

const exportData = () => {
  toast.info('جاري تصدير البيانات...')
}

const bulkAdd = () => {
  toast.info('إضافة متعددة غير متاحة حالياً')
}

// Load data on mount
onMounted(() => {
  fetchUsers()
})
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/users.css';
</style>
