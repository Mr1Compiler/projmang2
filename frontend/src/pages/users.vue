<template>
  <div class="fill-height data-page">
    <v-container fluid class="pa-6" style="padding: 0 20px !important;">
      <!-- رأس الصفحة المحسن -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-account-group</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">إدارة المستخدمين</h1>
              <p class="subtitle">إدارة وتتبع جميع المستخدمين والصلاحيات</p>
            </div>
          </div>
        </div>
      </div>

      <!-- إحصائيات سريعة محسنة -->
      <v-row class="mb-8">
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-account-multiple</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ totalUsers }}</h3>
                  <p class="stat-label">إجمالي المستخدمين</p>
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
                  <h3 class="stat-value">{{ activeUsers }}</h3>
                  <p class="stat-label">مستخدمين نشطين</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-warning" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-account-clock</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ pendingUsers }}</h3>
                  <p class="stat-label">مستخدمين معلقين</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-info" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-shield-account</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ adminUsers }}</h3>
                  <p class="stat-label">مدراء النظام</p>
                </div>
              </div>
            </v-card>
          </v-col>
        </v-row>

        <!-- فلاتر البحث -->
        <v-card class="filters-card mb-6" elevation="2">
          <v-card-title class="filters-header-new">
            <div class="d-flex align-center justify-start" style="width: 100%;">
              <v-icon size="20" class="me-2" style="color: #ffffff !important;">mdi-menu-down</v-icon>
              <span class="filters-header-title" style="color: #ef4444 !important;">البحث والفلترة</span>
            </div>
          </v-card-title>
          <v-card-text class="filters-content">
            <v-row no-gutters>
              <v-col cols="12" md="3" class="px-2">
                <v-text-field
                  v-model="searchQuery"
                  placeholder="البحث في المستخدمين"
                  prepend-inner-icon="mdi-magnify"
                  variant="outlined"
                  clearable
                  hide-details
                  density="compact"
                  class="search-field-new"
                />
              </v-col>
              <v-col cols="12" md="3" class="px-2">
                <v-select
                  v-model="selectedRole"
                  :items="roles"
                  label="الدور"
                  variant="outlined"
                  density="compact"
                  clearable
                  hide-details
                  class="filter-field-new"
                />
              </v-col>
              <v-col cols="12" md="3" class="px-2">
                <v-select
                  v-model="selectedStatus"
                  :items="statusOptions"
                  label="الحالة"
                  variant="outlined"
                  density="compact"
                  clearable
                  hide-details
                  class="filter-field-new"
                />
              </v-col>
              <v-col cols="12" md="3" class="px-2">
                <v-select
                  v-model="selectedDepartment"
                  :items="departments"
                  label="القسم"
                  variant="outlined"
                  density="compact"
                  clearable
                  hide-details
                  class="filter-field-new"
                />
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>

        <!-- جدول المستخدمين -->
        <v-card class="users-table" elevation="2">
          <v-card-title class="table-title-header d-flex align-center justify-space-between">
            <div class="d-flex align-center">
              <v-icon class="me-2" size="18" style="color: #ffffff !important;">mdi-table</v-icon>
              <span class="title-text">قائمة المستخدمين</span>
            </div>
            <div class="d-flex table-header-buttons" style="gap: 0.5rem;">
              <v-btn
                class="table-header-btn"
                size="small"
                style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important; color: #ffffff !important;"
              >
                <v-icon class="me-1" size="18">mdi-download</v-icon>
                تصدير البيانات
              </v-btn>
              <v-btn
                v-if="canCreate"
                class="add-button add-user-btn btn-glow light-sweep smooth-transition"
                @click="showAddUserDialog = true"
                elevation="2"
                color="primary"
                size="small"
                style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
              >
                <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
                إضافة مستخدم
              </v-btn>
            </div>
          </v-card-title>
          <div class="table-spacer"></div>
          <v-data-table
            :headers="headers"
            :items="filteredUsers"
            :search="searchQuery"
            class="elevation-0 users-data-table"
            :items-per-page="-1"
            hide-default-footer
          >
            <template v-slot:item.user="{ item }">
              <div class="d-flex align-center">
                <v-avatar size="40" class="me-3">
                  <v-img :src="item.avatar" />
                </v-avatar>
                <div>
                  <div class="font-weight-medium">{{ item.name }}</div>
                  <div class="text-caption text-grey">{{ item.email }}</div>
                </div>
              </div>
            </template>

            <template v-slot:item.role="{ item }">
              <v-chip
                :color="getRoleColor(item.role)"
                size="small"
                variant="flat"
              >
                {{ getRoleText(item.role) }}
              </v-chip>
            </template>

            <template v-slot:item.status="{ item }">
              <v-chip
                :color="getStatusColor(item.status)"
                size="small"
                variant="flat"
              >
                {{ getStatusText(item.status) }}
              </v-chip>
            </template>

            <template v-slot:item.lastLogin="{ item }">
              <span class="text-body-2">{{ formatDate(item.lastLogin) }}</span>
            </template>

            <template v-slot:item.actions="{ item }">
              <v-btn
                icon="mdi-eye"
                size="small"
                variant="elevated"
                class="view-btn"
                data-action="view"
                @click="viewUser(item)"
              />
              <v-btn
                v-if="canUpdate && canEditUser(item)"
                icon="mdi-pencil"
                size="small"
                variant="elevated"
                class="edit-btn"
                data-action="edit"
                @click="editUser(item)"
              />
              <v-btn
                v-if="canUpdatePassword && canEditUser(item)"
                icon="mdi-lock-reset"
                size="small"
                variant="elevated"
                class="password-btn"
                data-action="password"
                @click="resetPassword(item)"
              />
              <v-btn
                v-if="canUpdate && canEditUser(item)"
                :icon="item.status === 'active' ? 'mdi-account-off' : 'mdi-account-check'"
                size="small"
                variant="elevated"
                :class="item.status === 'active' ? 'deactivate-btn' : 'activate-btn'"
                :data-action="item.status === 'active' ? 'deactivate' : 'activate'"
                @click="toggleUserStatus(item)"
              />
            </template>
          </v-data-table>

          <!-- Pagination -->
          <div class="d-flex justify-center pa-4" v-if="pagination.totalPages > 0">
            <v-pagination
              v-model="pagination.page"
              :length="pagination.totalPages"
              :total-visible="7"
              @update:model-value="onPageChange"
              rounded="circle"
              density="comfortable"
              active-color="primary"
            />
          </div>
        </v-card>

    </v-container>
  </div>

  <!-- نافذة حوار إضافة مستخدم جديد -->
  <v-dialog v-model="showAddUserDialog" max-width="950px" persistent>
    <v-card class="add-user-dialog" rounded="lg">
      <!-- Header -->
      <v-card-title class="dialog-header pa-4">
        <div class="d-flex align-center">
          <v-avatar color="white" size="44" class="me-3">
            <v-icon color="primary" size="26">mdi-account-plus</v-icon>
          </v-avatar>
          <div>
            <h3 class="text-h5 font-weight-bold text-white mb-0">إضافة مستخدم جديد</h3>
            <span class="text-body-2 text-white-darken-1">أدخل بيانات المستخدم الجديد</span>
          </div>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          size="default"
          @click="closeAddUserDialog"
          class="close-btn"
        />
      </v-card-title>

      <!-- Form Content -->
      <v-card-text class="dialog-content pa-8">
        <v-form ref="addUserForm" v-model="formValid" lazy-validation>
          <v-row>
            <!-- اسم المستخدم -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.username"
                label="اسم المستخدم"
                placeholder="أدخل اسم المستخدم"
                :rules="usernameRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-account"
                color="primary"
                class="mb-2"
              />
            </v-col>

            <!-- الاسم الكامل -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.fullName"
                label="الاسم الكامل"
                placeholder="أدخل الاسم الكامل"
                :rules="fullNameRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-account-outline"
                color="primary"
                class="mb-2"
              />
            </v-col>

            <!-- البريد الإلكتروني -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.email"
                label="البريد الإلكتروني"
                placeholder="example@domain.com"
                :rules="emailRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-email-outline"
                type="email"
                color="primary"
                class="mb-2"
              />
            </v-col>

            <!-- رقم الهاتف -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.phone"
                label="رقم الهاتف"
                placeholder="07xxxxxxxx"
                :rules="phoneRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-phone-outline"
                type="tel"
                color="primary"
                class="mb-2"
              />
            </v-col>

            <!-- المسمى الوظيفي -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.jobTitle"
                label="المسمى الوظيفي"
                placeholder="مثال: مهندس برمجيات"
                :rules="jobTitleRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-briefcase-outline"
                color="primary"
                class="mb-2"
              />
            </v-col>

            <!-- الدور -->
            <v-col cols="12" md="6">
              <v-select
                v-model="newUser.roleId"
                :items="roleSelectOptions"
                item-title="title"
                item-value="value"
                label="الدور"
                placeholder="اختر الدور"
                :rules="[v => !!v || 'الدور مطلوب']"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-shield-account-outline"
                color="primary"
                class="mb-2"
              />
            </v-col>

            <!-- Divider for password section -->
            <v-col cols="12" class="py-3">
              <v-divider class="my-3" />
              <div class="text-body-2 text-grey-darken-1 mb-3">
                <v-icon size="20" class="me-2">mdi-lock-outline</v-icon>
                معلومات تسجيل الدخول
              </div>
            </v-col>

            <!-- كلمة المرور -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.password"
                label="كلمة المرور"
                placeholder="أدخل كلمة المرور"
                :rules="passwordRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-lock-outline"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                @click:append-inner="showPassword = !showPassword"
                color="primary"
                hint="8 أحرف على الأقل"
                persistent-hint
                class="mb-2"
              />
            </v-col>

            <!-- تأكيد كلمة المرور -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.confirmPassword"
                label="تأكيد كلمة المرور"
                placeholder="أعد إدخال كلمة المرور"
                :rules="confirmPasswordRules"
                variant="outlined"
                density="default"
                prepend-inner-icon="mdi-lock-check-outline"
                :type="showConfirmPassword ? 'text' : 'password'"
                :append-inner-icon="showConfirmPassword ? 'mdi-eye-off' : 'mdi-eye'"
                @click:append-inner="showConfirmPassword = !showConfirmPassword"
                color="primary"
                class="mb-2"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <!-- Actions -->
      <v-card-actions class="dialog-actions pa-5">
        <v-spacer />
        <v-btn
          variant="outlined"
          size="large"
          @click="closeAddUserDialog"
          class="me-3 px-8"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          size="large"
          @click="saveNewUser"
          :loading="saving"
          :disabled="!formValid"
          class="px-8"
          prepend-icon="mdi-check"
        >
          <v-icon start size="16">mdi-content-save</v-icon>
          حفظ المستخدم
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- نافذة عرض تفاصيل المستخدم -->
  <v-dialog v-model="showViewUserDialog" max-width="800px">
    <v-card class="view-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" class="me-3">mdi-account-details</v-icon>
          <h2>تفاصيل المستخدم</h2>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          @click="closeViewUserDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedUser" class="dialog-content">
        <v-row>
          <v-col cols="12" class="text-center mb-4">
            <v-avatar size="80">
              <v-img :src="selectedUser.avatar" />
            </v-avatar>
          </v-col>

          <!-- اسم المستخدم -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="selectedUser.username"
              label="اسم المستخدم"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- الاسم الكامل -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="selectedUser.fullName"
              label="الاسم الكامل"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- البريد الإلكتروني -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="selectedUser.email"
              label="البريد الإلكتروني"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- رقم الهاتف -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="selectedUser.phone"
              label="رقم الهاتف"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- المسمى الوظيفي -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="selectedUser.jobTitle"
              label="المسمى الوظيفي"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- الحالة -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="getStatusText(selectedUser.status)"
              label="الحالة"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- تاريخ الإنشاء -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="formatDate(selectedUser.createdAt)"
              label="تاريخ الإنشاء"
              variant="outlined"
              readonly
            />
          </v-col>

          <!-- آخر دخول -->
          <v-col cols="12" md="6">
            <v-text-field
              :model-value="formatDate(selectedUser.lastLogin)"
              label="آخر دخول"
              variant="outlined"
              readonly
            />
          </v-col>
        </v-row>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="primary"
          variant="elevated"
          @click="closeViewUserDialog"
        >
          إغلاق
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- نافذة تعديل المستخدم -->
  <v-dialog v-model="showEditUserDialog" max-width="800px" persistent>
    <v-card class="edit-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" class="me-3">mdi-account-edit</v-icon>
          <h2>تعديل المستخدم</h2>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          @click="closeEditUserDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedUser" class="dialog-content">
        <v-form ref="editUserForm" v-model="editFormValid" lazy-validation>
          <v-row>
            <v-col cols="12" class="text-center mb-4">
              <v-avatar size="80">
                <v-img :src="selectedUser.avatar" />
              </v-avatar>
            </v-col>

            <!-- اسم المستخدم -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.username"
                label="اسم المستخدم"
                variant="outlined"
                hint="يستخدم لتسجيل الدخول"
              />
            </v-col>

            <!-- الاسم الكامل -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.fullName"
                label="الاسم الكامل"
                variant="outlined"
              />
            </v-col>

            <!-- البريد الإلكتروني -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.email"
                label="البريد الإلكتروني"
                :rules="emailRules"
                variant="outlined"
              />
            </v-col>

            <!-- رقم الهاتف -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.phone"
                label="رقم الهاتف"
                variant="outlined"
              />
            </v-col>

            <!-- المسمى الوظيفي -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.jobTitle"
                label="المسمى الوظيفي"
                variant="outlined"
              />
            </v-col>

            <!-- الحالة -->
            <v-col cols="12" md="6">
              <v-select
                v-model="selectedUser.status"
                :items="statusOptions"
                label="الحالة"
                variant="outlined"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          variant="outlined"
          @click="closeEditUserDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          @click="saveEditUser"
          :loading="editSaving"
          :disabled="!editFormValid"
        >
          حفظ التعديلات
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- نافذة إعادة تعيين كلمة المرور -->
  <v-dialog v-model="showResetPasswordDialog" max-width="500px">
    <v-card class="reset-password-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" class="me-3">mdi-lock-reset</v-icon>
          <h2>تغيير كلمة المرور</h2>
        </div>
        <v-btn
          icon="mdi-close"
          variant="text"
          @click="closeResetPasswordDialog"
          class="close-btn"
        />
      </v-card-title>

      <v-divider />

      <v-card-text v-if="selectedUser" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60">
            <v-img v-if="selectedUser.avatar" :src="selectedUser.avatar" />
            <v-icon v-else size="30">mdi-account</v-icon>
          </v-avatar>
          <h4 class="mt-2">{{ selectedUser.name }}</h4>
          <p class="text-caption">{{ selectedUser.email }}</p>
        </div>

        <v-form ref="resetPasswordForm" @submit.prevent="confirmResetPassword">
          <v-text-field
            v-model="resetPasswordData.password"
            label="كلمة المرور الجديدة"
            :rules="resetPasswordRules"
            :type="showResetPassword ? 'text' : 'password'"
            :append-inner-icon="showResetPassword ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append-inner="showResetPassword = !showResetPassword"
            variant="outlined"
            hint="8 أحرف على الأقل"
            persistent-hint
            class="mb-3"
          />
          <v-text-field
            v-model="resetPasswordData.confirmPassword"
            label="تأكيد كلمة المرور"
            :rules="resetConfirmPasswordRules"
            :type="showResetConfirmPassword ? 'text' : 'password'"
            :append-inner-icon="showResetConfirmPassword ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append-inner="showResetConfirmPassword = !showResetConfirmPassword"
            variant="outlined"
          />
        </v-form>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          variant="outlined"
          @click="closeResetPasswordDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          @click="confirmResetPassword"
          :loading="resetLoading"
        >
          تغيير كلمة المرور
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { listUsers, createUser, updateUser, updateUserPassword, updateUserStatus } from '@/api/users'
import { listRoles, assignRoleToUser } from '@/api/roles'
import { DEFAULT_LIMIT } from '@/constants/pagination'
import { usePermissions } from '@/composables/usePermissions'
import { useAuthStore } from '@/stores/auth'

// Auth store
const authStore = useAuthStore()
const currentUserId = computed(() => authStore.user?.id)

// Check if user can be edited (not self, not super admin)
const canEditUser = (user) => {
  // Can't edit yourself
  if (user.id === currentUserId.value) return false
  // Can't edit super admin (assuming role name or id check)
  if (user.role === 'admin' || user.roleName === 'Super Admin' || user.roleName === 'admin') return false
  return true
}

// Permissions
const { canCreate, canUpdate, canUpdatePassword } = usePermissions('/users')

// البيانات التفاعلية
const loading = ref(false)
const searchQuery = ref('')
const selectedRole = ref(null)
const selectedStatus = ref(null)
const selectedDepartment = ref(null)

// متغيرات نافذة إضافة المستخدم
const showAddUserDialog = ref(false)
const formValid = ref(false)
const saving = ref(false)
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const addUserForm = ref(null)

// متغيرات نوافذ الإجراءات
const showViewUserDialog = ref(false)
const showEditUserDialog = ref(false)
const showResetPasswordDialog = ref(false)
const selectedUser = ref(null)
const editUserForm = ref(null)
const editFormValid = ref(false)
const editSaving = ref(false)
const resetLoading = ref(false)

// متغيرات نافذة تغيير كلمة المرور
const resetPasswordForm = ref(null)
const showResetPassword = ref(false)
const showResetConfirmPassword = ref(false)
const resetPasswordData = ref({
  password: '',
  confirmPassword: ''
})

// خيارات الفلاتر - يتم تحميلها من الـ API
const roles = ref([])
const rolesData = ref([]) // البيانات الكاملة للأدوار من الـ API

const departments = ref([
  'تقنية المعلومات',
  'الموارد البشرية',
  'المحاسبة',
  'المشاريع',
  'التسويق',
  'الإدارة',
  'الصيانة'
])

// رؤوس الجدول
const headers = ref([
  { title: 'المستخدم', key: 'user', sortable: true },
  { title: 'الدور', key: 'role', sortable: true },
  { title: 'القسم', key: 'department', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'آخر دخول', key: 'lastLogin', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
])

// بيانات المستخدمين - يتم تحميلها من الـ API
const users = ref([])
const pagination = ref({
  total: 0,
  page: 1,
  limit: DEFAULT_LIMIT,
  totalPages: 0
})

// بيانات المستخدم الجديد (متوافق مع Backend CreateUser DTO)
const newUser = ref({
  username: '',
  fullName: '',
  email: '',
  phone: '',
  jobTitle: '',
  password: '',
  confirmPassword: '',
  roleId: null // الدور المحدد
})

// خيارات النماذج
const roleOptions = [
  { title: 'مدير عام', value: 'admin' },
  { title: 'مدير مشروع', value: 'project_manager' },
  { title: 'مهندس', value: 'engineer' },
  { title: 'محاسب', value: 'accountant' },
  { title: 'موظف', value: 'employee' },
  { title: 'مراجع', value: 'reviewer' },
  { title: 'مستخدم', value: 'user' }
]

const departmentOptions = [
  { title: 'الإدارة', value: 'الإدارة' },
  { title: 'المحاسبة', value: 'المحاسبة' },
  { title: 'الهندسة', value: 'الهندسة' },
  { title: 'الموارد البشرية', value: 'الموارد البشرية' },
  { title: 'الصيانة', value: 'الصيانة' },
  { title: 'التسويق', value: 'التسويق' },
  { title: 'المبيعات', value: 'المبيعات' },
  { title: 'تكنولوجيا المعلومات', value: 'تكنولوجيا المعلومات' }
]

// خيارات الحالة (متوافقة مع Backend - active أو inactive فقط)
const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' }
]

// قواعد التحقق من صحة البيانات (متوافقة مع Backend validations)
const usernameRules = [
  v => !!v || 'اسم المستخدم مطلوب',
  v => (v && v.length >= 3) || 'اسم المستخدم يجب أن يكون على الأقل 3 أحرف',
  v => (v && v.length <= 50) || 'اسم المستخدم يجب أن يكون أقل من 50 حرف',
  v => /^[a-zA-Z0-9_]+$/.test(v) || 'اسم المستخدم يجب أن يحتوي على حروف وأرقام و _ فقط'
]

const fullNameRules = [
  v => !!v || 'الاسم الكامل مطلوب',
  v => (v && v.length >= 2) || 'الاسم يجب أن يكون على الأقل حرفين',
  v => (v && v.length <= 100) || 'الاسم يجب أن يكون أقل من 100 حرف'
]

const emailRules = [
  v => !!v || 'البريد الإلكتروني مطلوب',
  v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح'
]

const phoneRules = [
  v => !!v || 'رقم الهاتف مطلوب',
  v => (v && v.length >= 7) || 'رقم الهاتف غير صحيح'
]

const jobTitleRules = [
  v => !!v || 'المسمى الوظيفي مطلوب'
]

const passwordRules = [
  v => !!v || 'كلمة المرور مطلوبة',
  v => (v && v.length >= 8) || 'كلمة المرور يجب أن تكون على الأقل 8 أحرف'
]

const confirmPasswordRules = [
  v => !!v || 'تأكيد كلمة المرور مطلوب',
  v => v === newUser.value.password || 'كلمة المرور غير متطابقة'
]

const resetPasswordRules = [
  v => !!v || 'كلمة المرور مطلوبة',
  v => (v && v.length >= 8) || 'كلمة المرور يجب أن تكون على الأقل 8 أحرف'
]

const resetConfirmPasswordRules = [
  v => !!v || 'تأكيد كلمة المرور مطلوب',
  v => v === resetPasswordData.value.password || 'كلمة المرور غير متطابقة'
]

const requiredRules = [
  v => !!v || 'هذا الحقل مطلوب'
]

// الإحصائيات المحسوبة
const totalUsers = computed(() => users.value.length)
const activeUsers = computed(() => users.value.filter(user => user.status === 'active').length)
const pendingUsers = computed(() => users.value.filter(user => user.status === 'pending').length)
const adminUsers = computed(() => users.value.filter(user => user.role === 'admin').length)

// المستخدمين المفلترين
const filteredUsers = computed(() => {
  let filtered = users.value

  // فلترة بالبحث
  if (searchQuery.value) {
    filtered = filtered.filter(user =>
      user.name?.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.email?.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.phone?.includes(searchQuery.value)
    )
  }

  if (selectedRole.value) {
    filtered = filtered.filter(user => user.role === selectedRole.value)
  }

  if (selectedStatus.value) {
    filtered = filtered.filter(user => user.status === selectedStatus.value)
  }

  return filtered
})

// الدوال المساعدة
const formatDate = (date) => {
  if (!date) return 'لم يسجل دخول'
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getRoleColor = (role) => {
  const colors = {
    'admin': 'red',
    'project_manager': 'blue',
    'engineer': 'green',
    'accountant': 'purple',
    'employee': 'orange',
    'reviewer': 'teal',
    'user': 'grey'
  }
  return colors[role] || 'grey'
}

const getRoleText = (role) => {
  const texts = {
    'admin': 'مدير النظام',
    'project_manager': 'مدير المشاريع',
    'engineer': 'مهندس',
    'accountant': 'محاسب',
    'employee': 'موظف إداري',
    'reviewer': 'مراجع',
    'user': 'مستخدم عادي'
  }
  return texts[role] || 'غير محدد'
}

const getStatusColor = (status) => {
  const colors = {
    'active': 'success',
    'inactive': 'error',
    'pending': 'warning',
    'banned': 'grey'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    'active': 'نشط',
    'inactive': 'غير نشط',
    'pending': 'معلق',
    'banned': 'محظور'
  }
  return texts[status] || 'غير محدد'
}

const applyFilters = () => {
  console.log('تم تطبيق الفلاتر')
}

const resetFilters = () => {
  searchQuery.value = ''
  selectedRole.value = null
  selectedStatus.value = null
}

const viewUser = (user) => {
  selectedUser.value = { ...user }
  showViewUserDialog.value = true
}

const editUser = (user) => {
  selectedUser.value = { ...user }
  showEditUserDialog.value = true
}

const resetPassword = (user) => {
  selectedUser.value = { ...user }
  showResetPasswordDialog.value = true
}

const toggleUserStatus = async (user) => {
  try {
    const newStatus = user.status === 'active' ? 'inactive' : 'active'
    await updateUserStatus(user.id, newStatus)
    await loadUsers()
    console.log(`تم ${newStatus === 'active' ? 'تفعيل' : 'تعطيل'} المستخدم بنجاح`)
  } catch (error) {
    console.error('خطأ في تغيير حالة المستخدم:', error)
  }
}

// دوال إدارة نافذة إضافة المستخدم
const closeAddUserDialog = () => {
  showAddUserDialog.value = false
  resetForm()
}

const resetForm = () => {
  newUser.value = {
    username: '',
    fullName: '',
    email: '',
    phone: '',
    jobTitle: '',
    password: '',
    confirmPassword: '',
    roleId: null
  }
  if (addUserForm.value) {
    addUserForm.value.resetValidation()
  }
  formValid.value = false
  showPassword.value = false
  showConfirmPassword.value = false
}

const saveNewUser = async () => {
  if (!addUserForm.value.validate()) {
    return
  }

  saving.value = true

  try {
    // إنشاء payload للـ API (متوافق مع Backend CreateUser DTO)
    const userData = {
      username: newUser.value.username,
      email: newUser.value.email,
      password: newUser.value.password,
      fullName: newUser.value.fullName,
      phone: newUser.value.phone,
      jobTitle: newUser.value.jobTitle
    }

    // إرسال البيانات إلى الـ API
    const createdUser = await createUser(userData)

    // تعيين الدور للمستخدم إذا تم اختياره
    if (newUser.value.roleId && createdUser?.id) {
      try {
        await assignRoleToUser(createdUser.id, newUser.value.roleId)
        console.log('تم تعيين الدور للمستخدم بنجاح')
      } catch (roleError) {
        console.error('خطأ في تعيين الدور:', roleError)
      }
    }

    // إعادة تحميل البيانات من الـ API
    await loadUsers()

    // إغلاق النافذة وإعادة تعيين النموذج
    closeAddUserDialog()

    console.log('تم إضافة المستخدم بنجاح')

  } catch (error) {
    console.error('خطأ في إضافة المستخدم:', error)
  } finally {
    saving.value = false
  }
}

// دوال إدارة نوافذ الإجراءات
const closeAllDialogs = () => {
  showViewUserDialog.value = false
  showEditUserDialog.value = false
  showResetPasswordDialog.value = false
  selectedUser.value = null
  editFormValid.value = false
}

const closeViewUserDialog = () => {
  closeAllDialogs()
}

const closeEditUserDialog = () => {
  closeAllDialogs()
}

const saveEditUser = async () => {
  if (!editUserForm.value.validate()) {
    return
  }

  editSaving.value = true

  try {
    // إنشاء payload للتحديث (متوافق مع Backend UpdateUser DTO)
    // جميع الحقول اختيارية في التحديث
    const userData = {
      username: selectedUser.value.username || undefined,
      email: selectedUser.value.email || undefined,
      fullName: selectedUser.value.fullName || undefined,
      phone: selectedUser.value.phone || undefined,
      jobTitle: selectedUser.value.jobTitle || undefined,
      status: selectedUser.value.status || undefined
    }

    // إرسال التحديث إلى الـ API
    await updateUser(selectedUser.value.id, userData)

    // إعادة تحميل البيانات
    await loadUsers()

    closeEditUserDialog()
    console.log('تم تحديث المستخدم بنجاح')

  } catch (error) {
    console.error('خطأ في تحديث المستخدم:', error)
  } finally {
    editSaving.value = false
  }
}

const closeResetPasswordDialog = () => {
  closeAllDialogs()
  resetPasswordData.value = { password: '', confirmPassword: '' }
  showResetPassword.value = false
  showResetConfirmPassword.value = false
}

const confirmResetPassword = async () => {
  // التحقق من صحة النموذج
  const { valid } = await resetPasswordForm.value.validate()
  if (!valid) return

  resetLoading.value = true

  try {
    // إرسال كلمة المرور الجديدة إلى الـ API
    await updateUserPassword(selectedUser.value.id, resetPasswordData.value.password)

    closeResetPasswordDialog()
    console.log('تم تغيير كلمة المرور بنجاح')

  } catch (error) {
    console.error('خطأ في تغيير كلمة المرور:', error)
  } finally {
    resetLoading.value = false
  }
}

// ========================================
// دالة تحميل البيانات من الـ API
// ========================================
const loadUsers = async () => {
  loading.value = true
  try {
    const response = await listUsers({ page: pagination.value.page, limit: pagination.value.limit })
    console.log('Users data received:', response)

    // Map backend data to frontend format
    const mapUser = (user) => ({
      id: user.id,
      // Backend fields (للتعديل)
      username: user.userName || user.username || '',
      fullName: user.fullName || '',
      email: user.email || '',
      phone: user.phone || '',
      jobTitle: user.jobTitle || '',
      status: user.status || 'active',
      lastLogin: user.lastLogin,
      createdAt: user.createdAt,
      // Display fields (للعرض في الجدول)
      name: user.fullName || '',
      role: mapJobTitleToRole(user.jobTitle),
      avatar: user.avatar || `https://ui-avatars.com/api/?name=${encodeURIComponent(user.fullName || 'User')}&background=667eea&color=fff&size=128`
    })

    if (response.data) {
      users.value = response.data.map(mapUser)
      pagination.value = {
        total: response.total || 0,
        page: response.page || 1,
        limit: response.limit || DEFAULT_LIMIT,
        totalPages: response.totalPages || 0
      }
    } else if (Array.isArray(response)) {
      users.value = response.map(mapUser)
    }
  } catch (error) {
    console.error('Error loading users:', error)
  } finally {
    loading.value = false
  }
}

// Helper function to map job title to role
const mapJobTitleToRole = (jobTitle) => {
  if (!jobTitle) return 'user'
  const title = jobTitle.toLowerCase()
  if (title.includes('admin') || title.includes('مدير')) return 'admin'
  if (title.includes('manager') || title.includes('مشروع')) return 'project_manager'
  if (title.includes('engineer') || title.includes('مهندس')) return 'engineer'
  if (title.includes('accountant') || title.includes('محاسب')) return 'accountant'
  if (title.includes('reviewer') || title.includes('مراجع')) return 'reviewer'
  return 'employee'
}

// Handle page change
const onPageChange = (page) => {
  pagination.value.page = page
  loadUsers()
}

// ========================================
// دالة تحميل الأدوار من الـ API
// ========================================
const loadRoles = async () => {
  try {
    const response = await listRoles()
    console.log('Roles data received:', response)

    // تخزين البيانات الكاملة - تأكد من أنها مصفوفة
    rolesData.value = Array.isArray(response) ? response : []

    // إنشاء قائمة للفلتر (أسماء الأدوار فقط)
    roles.value = rolesData.value.map(role => role.name)

  } catch (error) {
    console.error('Error loading roles:', error)
    rolesData.value = []
    roles.value = []
  }
}

// قائمة الأدوار للاختيار في النموذج
const roleSelectOptions = computed(() => {
  if (!Array.isArray(rolesData.value)) {
    return []
  }
  return rolesData.value.map(role => ({
    title: role.name,
    value: role.id
  }))
})

// ========================================
// دورة الحياة
// ========================================
onMounted(() => {
  loadUsers()
  loadRoles()
})
</script>

<style scoped>
/* ========================================
   صفحة إدارة المستخدمين - تصميم عصري محدث
   نسق ألوان متدرج: بنفسجي داكن → وردي → أسود
   جدول محسن + نافذة إضافة مستخدم كاملة
   تأثيرات بصرية متقدمة ونماذج تفاعلية
   برمجة وتصميم قسم تكنولوجيا المعلومات جامعة التراث
   ارتكاز للحلول البرمجية erticaz.com
   ======================================== */

/* تحسين ألوان صفحة المستخدمين - نفس تصميم الصفحات الأخرى */
.data-page {
  background: #dbeafe !important;
  color: #2e3cbc;
  min-height: 100vh;
  padding: 0;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
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

.header-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.add-user-btn {
  background: rgba(255, 255, 255, 0.2) !important;
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  color: white !important;
  font-weight: 600 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
  transition: all 0.3s ease !important;
}

.add-user-btn:hover {
  background: rgba(255, 255, 255, 0.3) !important;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3) !important;
}

.add-user-btn :deep(.v-btn__prepend) {
  margin-inline-end: 8px;
}

@keyframes star-twinkle {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.1); opacity: 0.8; }
}

.page-title {
  font-size: 2rem;
  font-weight: 800;
  color: rgb(246, 246, 246);
  margin: 0;
  padding: 0;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
  text-align: center !important;
  width: auto;
  line-height: 1.2;
  letter-spacing: 0.5px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.text-glow {
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3), 0 0 30px rgba(255, 255, 255, 0.2);
}

.fade-in {
  animation: fadeInUp 0.8s ease-out;
}

.page-subtitle {
  font-size: 1rem;
  color: rgba(243, 240, 240, 0.9);
  margin: 0;
  padding: 0;
  text-align: center !important;
  width: 100%;
  line-height: 1.4;
  letter-spacing: 0.3px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  font-weight: 500;
}

.stats-section {
  margin-bottom: 30px;
  padding: 0 20px;
}

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

/* استثناء الكروت الإحصائية من تنسيقات v-card العامة */
.modern-stat-card,
.modern-stat-card.v-card,
.data-page .modern-stat-card,
.data-page .modern-stat-card.v-card {
  background: #ffffff !important;
  color: #000000 !important;
  border: 2px solid !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  backdrop-filter: none !important;
  border-radius: 20px !important;
}

.modern-stat-card:hover,
.modern-stat-card.v-card:hover,
.data-page .modern-stat-card:hover,
.data-page .modern-stat-card.v-card:hover {
  transform: translateY(-8px) scale(1.02) !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2) !important;
}

/* التأكد من أن الكروت الإحصائية لا تتأثر بتنسيقات v-card العامة */
.modern-stat-card .stat-card-content,
.modern-stat-card .stat-icon-wrapper,
.modern-stat-card .stat-info,
.modern-stat-card .stat-value,
.modern-stat-card .stat-label {
  color: inherit !important;
}

/* تنسيقات v-card العامة (لا تنطبق على الكروت الإحصائية) */
.v-card:not(.modern-stat-card) {
  background: rgba(83, 69, 205, 0.95) !important;
  color: #e2c0cf !important;
  border: 1px solid rgba(224, 12, 118, 0.1) !important;
  box-shadow: 0 8px 32px rgba(79, 70, 229, 0.15) !important;
  backdrop-filter: blur(10px);
  border-radius: 16px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

.v-card:not(.modern-stat-card):hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(79, 70, 229, 0.25) !important;
  border-color: rgba(79, 70, 229, 0.2) !important;
}

.v-card-title {
  color: #1e293b !important;
  font-weight: 700 !important;
  font-size: 1.2rem !important;
  background: linear-gradient(135deg, rgba(27, 18, 186, 0.05) 0%, rgba(124, 58, 237, 0.08) 100%);
  border-bottom: 1px solid rgba(79, 70, 229, 0.1);
  padding: 1.5rem !important;
}

/* استثناء شريط عنوان جدول المستخدمين */
.users-table .v-card-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
}

.v-card-text {
  color: #475569 !important;
  padding: 1.5rem !important;
}

/* الجداول المحسنة */
.v-data-table {
  background: linear-gradient(135deg, rgba(248, 250, 252, 0.95) 0%, rgba(241, 245, 249, 0.9) 100%) !important;
  color: #1a1a1a !important;
  border-radius: 24px !important;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.1), 0 0 0 1px rgba(148, 163, 184, 0.1) !important;
  border: 1px solid rgba(148, 163, 184, 0.2);
  backdrop-filter: blur(20px);
  position: relative;
}

.v-data-table::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.05) 0%, rgba(147, 51, 234, 0.05) 100%);
  pointer-events: none;
  z-index: 1;
}

.v-data-table th {
  background: #ffffff !important;
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 0.7rem !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.2) !important;
  padding: 0.5rem 0.75rem !important;
  text-shadow: none !important;
  position: relative;
  overflow: hidden;
  backdrop-filter: none;
  letter-spacing: 0.025em;
  text-transform: uppercase;
  min-height: 32px !important;
}

.v-data-table th::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 0;
  height: 2px;
  background: linear-gradient(90deg, #3b82f6, #8b5cf6);
  transition: width 0.3s ease;
}

.v-data-table th:hover::after {
  width: 60%;
}

.v-data-table td {
  color: #1a1a1a !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.1) !important;
  padding: 0.5rem 0.75rem !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: rgba(248, 250, 252, 0.9);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 2;
  font-weight: 400;
  font-size: 0.75rem !important;
  min-height: 40px !important;
}

.v-data-table td::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 3px;
  background: linear-gradient(180deg, transparent, rgba(59, 130, 246, 0.5), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.v-data-table tbody tr:hover td::before {
  opacity: 1;
}

.v-data-table tbody tr:hover td {
  background: linear-gradient(135deg, rgba(69, 46, 46, 0.1) 0%, rgba(147, 51, 234, 0.05) 100%) !important;
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(54, 47, 47, 0.15);
  border-radius: 12px;
  margin: 4px 8px;
}

/* الأزرار المحسنة */
.v-btn {
  font-weight: 600 !important;
  border-radius: 12px !important;
  text-transform: none !important;
  box-shadow: 0 4px 12px rgba(202, 176, 191, 0.15) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.v-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(79, 70, 229, 0.25) !important;
}

.v-btn.v-btn--variant-elevated {
  background: linear-gradient(135deg, #4f46e5 0%, #7c3aed 100%) !important;
  color: white !important;
  box-shadow: 0 4px 20px rgba(79, 70, 229, 0.3) !important;
}

.v-btn.v-btn--variant-elevated:hover {
  background: linear-gradient(135deg, #4338ca 0%, #6d28d9 100%) !important;
  box-shadow: 0 8px 30px rgba(79, 70, 229, 0.4) !important;
}

/* أزرار الجدول المحسنة */
.v-data-table .v-btn {
  min-width: 40px !important;
  height: 40px !important;
  border-radius: 12px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(148, 163, 184, 0.2);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* ========================================
   تحديث ألوان أزرار الإجراءات
   - ألوان جديدة ومميزة لكل إجراء
   - تحسين الوضوح والتباين
   - تصميم عصري وجذاب
   ======================================== */

/* ========================================
   أزرار الإجراءات بألوان متنوعة
   - ألوان مميزة لكل إجراء
   - نصوص بيضاء للوضوح
   - ظلال ملونة متناسقة
   ======================================== */

/* قواعد أكثر تحديداً للأزرار */
.v-data-table .v-btn[icon="mdi-eye"],
.v-data-table .v-btn .v-icon[mdi-eye] {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.4) !important;
  border: none !important;
}

.v-data-table .v-btn[icon="mdi-pencil"],
.v-data-table .v-btn .v-icon[mdi-pencil] {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4) !important;
  border: none !important;
}

.v-data-table .v-btn[icon="mdi-key"],
.v-data-table .v-btn .v-icon[mdi-key] {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.4) !important;
  border: none !important;
}

.v-data-table .v-btn[icon="mdi-delete"],
.v-data-table .v-btn .v-icon[mdi-delete] {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4) !important;
  border: none !important;
}

/* تأثيرات التمرير للأزرار الملونة */
.v-data-table .v-btn[icon="mdi-eye"]:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 8px 25px rgba(59, 130, 246, 0.6) !important;
}

.v-data-table .v-btn[icon="mdi-pencil"]:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 8px 25px rgba(16, 185, 129, 0.6) !important;
}

.v-data-table .v-btn[icon="mdi-key"]:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 8px 25px rgba(245, 158, 11, 0.6) !important;
}

.v-data-table .v-btn[icon="mdi-delete"]:hover {
  transform: translateY(-3px) scale(1.05);
  box-shadow: 0 8px 25px rgba(239, 68, 68, 0.6) !important;
}

/* تأثير الضغط */
.v-data-table .v-btn:active {
  transform: translateY(-1px) scale(0.98);
}

/* تحسينات إضافية للأزرار */
.v-data-table .v-btn {
  border: none !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2) !important;
}

/* إصلاح النصوص في الأزرار الملونة */
.v-data-table .v-btn .v-icon {
  color: white !important;
  font-size: 18px !important;
}

/* تأثيرات إضافية للأزرار الملونة */
.v-data-table .v-btn[icon="mdi-eye"] .v-icon {
  color: white !important;
}

.v-data-table .v-btn[icon="mdi-pencil"] .v-icon {
  color: white !important;
}

.v-data-table .v-btn[icon="mdi-key"] .v-icon {
  color: white !important;
}

.v-data-table .v-btn[icon="mdi-delete"] .v-icon {
  color: white !important;
}

/* تأثيرات نهائية للأزرار */
.v-data-table .v-btn {
  position: relative;
  overflow: hidden;
}

.v-data-table .v-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.v-data-table .v-btn:hover::before {
  left: 100%;
}

/* قواعد إضافية لضمان تطبيق الألوان */
.v-data-table .v-btn:has(.v-icon.mdi-eye) {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
}

.v-data-table .v-btn:has(.v-icon.mdi-pencil) {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
}

.v-data-table .v-btn:has(.v-icon.mdi-key) {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
}

.v-data-table .v-btn:has(.v-icon.mdi-delete) {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
}

/* قواعد عامة للأزرار */
.v-data-table .v-btn {
  filter: contrast(1.1) brightness(1.05);
  border: none !important;
  min-width: 40px !important;
  height: 40px !important;
  border-radius: 12px !important;
}

/* تأثيرات إضافية للأزرار الملونة */
.v-data-table .v-btn:focus {
  outline: 2px solid rgba(255, 255, 255, 0.4) !important;
  outline-offset: 2px !important;
}

/* قواعد بديلة باستخدام class names */
.v-data-table .v-btn.view-btn,
.v-data-table .v-btn[data-action="view"] {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
}

.v-data-table .v-btn.edit-btn,
.v-data-table .v-btn[data-action="edit"] {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
}

.v-data-table .v-btn.reset-btn,
.v-data-table .v-btn[data-action="reset"] {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
}

.v-data-table .v-btn.delete-btn,
.v-data-table .v-btn[data-action="delete"] {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
}

.v-btn.v-btn--variant-outlined {
  border: 2px solid currentColor !important;
  background: rgba(65, 7, 7, 0.9) !important;
}

.v-btn.v-btn--variant-outlined:hover {
  background: currentColor !important;
  color: rgb(224, 209, 209) !important;
}

/* قواعد نهائية لضمان تطبيق الألوان */
.v-data-table .view-btn,
.v-data-table .view-btn .v-btn__content,
.v-data-table .view-btn .v-icon {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
}

.v-data-table .edit-btn,
.v-data-table .edit-btn .v-btn__content,
.v-data-table .edit-btn .v-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
}

.v-data-table .reset-btn,
.v-data-table .reset-btn .v-btn__content,
.v-data-table .reset-btn .v-icon {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
}

.v-data-table .delete-btn,
.v-data-table .delete-btn .v-btn__content,
.v-data-table .delete-btn .v-icon {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
}

/* تنسيق النوافذ المنبثقة الجديدة */
.view-user-dialog .v-card,
.edit-user-dialog .v-card,
.reset-password-dialog .v-card,
.delete-confirm-dialog .v-card {
  background: linear-gradient(135deg, rgba(248, 250, 252, 0.98) 0%, rgba(241, 245, 249, 0.95) 100%) !important;
  border-radius: 20px !important;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15) !important;
  backdrop-filter: blur(20px);
  border: 1px solid rgba(148, 163, 184, 0.2);
}

.reset-password-dialog .dialog-header,
.delete-confirm-dialog .dialog-header {
  background: linear-gradient(135deg, rgba(71, 85, 105, 0.9) 0%, rgba(51, 65, 85, 0.8) 100%) !important;
  color: #ffffff !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 1.5rem 2rem !important;
}

.view-user-dialog .dialog-header {
  background: #ffffff !important;
  color: #1a1a1a !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 1.5rem 2rem !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.2);
}

.edit-user-dialog .dialog-header {
  background: #ffffff !important;
  color: #1a1a1a !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 1.5rem 2rem !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.2);
}

/* تنسيق عنوان نافذة تفاصيل المستخدم */
.view-user-dialog .dialog-title h2 {
  color: #1a1a1a !important;
  font-weight: 700 !important;
  font-size: 1.5rem !important;
}

.view-user-dialog .dialog-title .v-icon {
  color: #1a1a1a !important;
}

.view-user-dialog .close-btn {
  color: #64748b !important;
}

.view-user-dialog .close-btn:hover {
  color: #1a1a1a !important;
  background: rgba(0, 0, 0, 0.05) !important;
}

/* تنسيق حقول نافذة تفاصيل المستخدم */
.view-user-dialog .v-field {
  background: #ffffff !important;
}

.view-user-dialog .v-field * {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.view-user-dialog .v-field input,
.view-user-dialog .v-field input[readonly],
.view-user-dialog .v-field textarea,
.view-user-dialog .v-field__input,
.view-user-dialog .v-text-field input,
.view-user-dialog input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  opacity: 1 !important;
}

.view-user-dialog .v-label,
.view-user-dialog label,
.view-user-dialog .v-field-label {
  color: #374151 !important;
  -webkit-text-fill-color: #374151 !important;
  opacity: 1 !important;
}

.view-user-dialog .v-card-text {
  background: #ffffff !important;
}

.view-user-dialog .dialog-content {
  background: #ffffff !important;
}

/* تنسيق شامل لنافذة تعديل المستخدم */
.edit-user-dialog .v-field {
  background: #ffffff !important;
}

.edit-user-dialog .v-field * {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.edit-user-dialog .v-field input,
.edit-user-dialog .v-field textarea,
.edit-user-dialog .v-field__input,
.edit-user-dialog .v-text-field input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  caret-color: #1a1a1a !important;
}

.edit-user-dialog .v-label,
.edit-user-dialog label,
.edit-user-dialog .v-field-label {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  opacity: 1 !important;
}

.edit-user-dialog .v-select__selection,
.edit-user-dialog .v-select__selection-text {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.edit-user-dialog .v-list-item {
  color: #1a1a1a !important;
  background: white !important;
}

.edit-user-dialog .v-list-item:hover {
  background: rgba(59, 130, 246, 0.1) !important;
}

.edit-user-dialog .v-list-item--active {
  background: rgba(59, 130, 246, 0.2) !important;
  color: #1d4ed8 !important;
}

/* تنسيق عنوان نافذة التعديل */
.edit-user-dialog .dialog-title h2 {
  color: #1a1a1a !important;
  font-weight: 700 !important;
  font-size: 1.5rem !important;
}

.edit-user-dialog .dialog-title .v-icon {
  color: #1a1a1a !important;
}

/* تنسيق أزرار نافذة التعديل */
.edit-user-dialog .dialog-actions .v-btn {
  border-radius: 12px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  padding: 12px 24px !important;
  min-width: 120px !important;
}

.edit-user-dialog .dialog-actions .v-btn[color="success"] {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  box-shadow: 0 4px 15px rgba(16, 185, 129, 0.3) !important;
}

.edit-user-dialog .dialog-actions .v-btn[color="success"]:hover {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  box-shadow: 0 6px 20px rgba(16, 185, 129, 0.4) !important;
  transform: translateY(-2px) !important;
}

.edit-user-dialog .dialog-actions .v-btn[color="grey"] {
  background: linear-gradient(135deg, #6b7280 0%, #4b5563 100%) !important;
  color: white !important;
  box-shadow: 0 4px 15px rgba(107, 114, 128, 0.3) !important;
}

.edit-user-dialog .dialog-actions .v-btn[color="grey"]:hover {
  background: linear-gradient(135deg, #4b5563 0%, #374151 100%) !important;
  box-shadow: 0 6px 20px rgba(107, 114, 128, 0.4) !important;
  transform: translateY(-2px) !important;
}

/* تنسيق الصورة الشخصية */
.edit-user-dialog .v-avatar {
  border: 3px solid rgba(59, 130, 246, 0.3) !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15) !important;
  transition: all 0.3s ease !important;
}

.edit-user-dialog .v-avatar:hover {
  border-color: rgba(59, 130, 246, 0.6) !important;
  transform: scale(1.05) !important;
}

/* تحسينات إضافية لنافذة التعديل */
.edit-user-dialog .v-card-text {
  background: #ffffff !important;
  padding: 2rem !important;
}

.edit-user-dialog .v-row {
  margin: 0 !important;
}

.edit-user-dialog .v-col {
  padding: 8px 12px !important;
}

/* تأثيرات الإدخال */
.edit-user-dialog .v-field--focused {
  border-color: rgba(59, 130, 246, 0.8) !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
  transform: translateY(-1px) !important;
}

.edit-user-dialog .v-field--error {
  border-color: rgba(239, 68, 68, 0.8) !important;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.15) !important;
}

/* رسائل التحقق */
.edit-user-dialog .v-messages {
  color: #ef4444 !important;
  font-size: 0.875rem !important;
  font-weight: 500 !important;
}

.edit-user-dialog .v-messages__message {
  color: #ef4444 !important;
}

/* تأثيرات التمرير */
.edit-user-dialog .v-field,
.edit-user-dialog .v-btn {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

/* تحسين الظلال */
.edit-user-dialog {
  box-shadow: 0 25px 80px rgba(0, 0, 0, 0.2) !important;
}

/* النماذج */
.v-text-field input,
.v-select input,
.v-textarea textarea {
  color: #1e293b !important;
  background: #ffffff !important;
}

.v-text-field label,
.v-select label,
.v-textarea label {
  color: #64748b !important;
}

.logo-section {
  border-bottom: 1px solid #e0e0e0;
  margin-bottom: 16px;
}

.menu-item {
  transition: all 0.3s ease;
  border-radius: 12px;
}

.menu-item:hover {
  background-color: rgba(79, 70, 229, 0.08);
}

.menu-item.v-list-item--active {
  background-color: rgba(79, 70, 229, 0.12);
  color: #4f46e5;
}

.project-item {
  transition: all 0.3s ease;
  border-radius: 8px;
}

.project-item:hover {
  background-color: rgba(0, 0, 0, 0.04);
}

.top-bar {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(255, 255, 255, 0.9) 100%);
  border-bottom: 1px solid rgba(79, 70, 229, 0.1);
  box-shadow: 0 4px 20px rgba(79, 70, 229, 0.1);
  backdrop-filter: blur(10px);
}

.search-field {
  border-radius: 8px;
}

.main-content {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: calc(100vh - 70px);
}

.page-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-radius: 20px;
  padding: 24px;
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.3);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.filters-card {
  background: #ffffff !important;
  background-color: #ffffff !important;
  border: 1px solid #e2e8f0 !important;
  border-radius: 12px !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
  overflow: hidden !important;
}

.filters-card .v-card-text {
  background: #ffffff !important;
  background-color: #ffffff !important;
  padding: 16px !important;
}

/* شريط عنوان الفلترة الجديد */
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
  color: #ef4444 !important;
  -webkit-text-fill-color: #ef4444 !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  letter-spacing: 0.2px !important;
}

.filters-header-new :deep(.v-icon) {
  color: #ef4444 !important;
  -webkit-text-fill-color: #ef4444 !important;
  fill: #ef4444 !important;
}

/* محتوى الفلترة */
.filters-content {
  padding: 0.5rem 0.75rem !important;
  background: #ffffff !important;
  background-color: #ffffff !important;
  border-radius: 0 0 12px 12px !important;
}

.filters-content,
.filters-content *,
.filters-content .v-row,
.filters-content .v-col {
  background: #ffffff !important;
  background-color: #ffffff !important;
}

/* حقول البحث والفلترة الجديدة */
.search-field-new :deep(.v-field),
.filter-field-new :deep(.v-field),
.search-field-new :deep(.v-field__input),
.filter-field-new :deep(.v-field__input),
.search-field-new :deep(.v-input),
.filter-field-new :deep(.v-input) {
  background: #ffffff !important;
  background-color: #ffffff !important;
}

.search-field-new :deep(.v-field),
.filter-field-new :deep(.v-field) {
  border-radius: 8px !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #fca5a5 !important;
  transition: all 0.3s ease !important;
}

.search-field-new :deep(.v-field:hover),
.filter-field-new :deep(.v-field:hover) {
  border-color: #f87171 !important;
  box-shadow: 0 2px 6px rgba(239, 68, 68, 0.1) !important;
}

.search-field-new :deep(.v-field--focused),
.filter-field-new :deep(.v-field--focused) {
  border-color: #f87171 !important;
  box-shadow: 0 0 0 2px rgba(239, 68, 68, 0.1) !important;
}

.search-field-new :deep(.v-field__input),
.filter-field-new :deep(.v-field__input) {
  color: #1e293b !important;
  font-size: 0.875rem !important;
  padding: 8px 12px !important;
}

.search-field-new :deep(.v-label),
.filter-field-new :deep(.v-label) {
  color: #1e40af !important;
  font-weight: 500 !important;
  font-size: 0.875rem !important;
}

.search-field-new :deep(.v-field--focused .v-label),
.filter-field-new :deep(.v-field--focused .v-label) {
  color: #1e40af !important;
}

.search-field-new :deep(.v-field__prepend-inner .v-icon) {
  color: #1e40af !important;
}

.search-field-new :deep(.v-field__input::placeholder) {
  color: #94a3b8 !important;
  opacity: 1 !important;
}


/* زر إعادة التعيين */
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
.filters-card :deep(.v-overlay__content) {
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid #fca5a5 !important;
  overflow: hidden !important;
  background: #ffffff !important;
}

.filter-field-new :deep(.v-field__append-inner .v-icon) {
  color: #1e40af !important;
}

.users-table {
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.1);
  background: #ffffff !important;
  position: relative;
  transition: all 0.4s ease;
}

/* شريط عنوان الجدول */
.table-title-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  padding: 8px 12px !important;
  border-radius: 8px 8px 0 0 !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2) !important;
  min-height: 36px !important;
  margin-bottom: 0 !important;
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

.users-table .v-card-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2) !important;
  padding: 8px 12px !important;
  position: relative !important;
  color: #ffffff !important;
  min-height: 36px !important;
  margin-bottom: 0 !important;
}

.users-table .v-card-title * {
  color: #ffffff !important;
}

.users-table .v-card-title .v-icon {
  color: #ffffff !important;
}

.users-table .v-card-title span {
  color: #ffffff !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  font-size: 0.95rem !important;
  font-weight: 600 !important;
}

.users-table .v-card-title .v-btn {
  color: #ffffff !important;
}

.users-table .v-card-title .v-btn .v-btn__content {
  color: #ffffff !important;
}

/* مسافة بين شريط العنوان والجدول */
.table-spacer {
  height: 1rem !important;
  background: #ffffff !important;
  width: 100% !important;
}

.users-data-table {
  margin-top: 0 !important;
  background: #ffffff !important;
}

/* إزالة التمرير من الجدول */
.users-data-table .v-data-table__wrapper {
  overflow-x: visible !important;
  overflow-y: visible !important;
}

.users-data-table .v-table__wrapper {
  overflow-x: visible !important;
  overflow-y: visible !important;
}

.users-data-table .v-table {
  overflow-x: visible !important;
  overflow-y: visible !important;
}

.v-data-table .v-data-table__wrapper {
  overflow-x: visible !important;
  overflow-y: visible !important;
}

.v-data-table .v-table__wrapper {
  overflow-x: visible !important;
  overflow-y: visible !important;
}

/* تنسيق موحد لأزرار شريط الجدول */
.table-header-buttons .table-header-btn {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  border: none !important;
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.3) !important;
  transition: all 0.3s ease !important;
}

.table-header-buttons .table-header-btn:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.4) !important;
}

.table-header-buttons .table-header-btn :deep(.v-icon) {
  color: #ffffff !important;
}

.table-header-buttons .table-header-btn :deep(.v-btn__content) {
  color: #ffffff !important;
}

.table-header-buttons {
  gap: 0.5rem !important;
  display: flex !important;
  align-items: center !important;
}

.chart-card {
    border-radius: 16px;
    box-shadow: 0 8px 32px rgba(79, 70, 229, 0.15);
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(255, 255, 255, 0.9) 100%);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(79, 70, 229, 0.1);
  }

  /* تحسينات إضافية للجدول */
  .v-data-table th::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
    animation: shimmerHeader 3s infinite;
  }

  .v-data-table tbody tr {
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .v-data-table tbody tr:nth-child(even) td {
    background: rgba(76, 86, 194, 0.2);
  }

  .v-data-table tbody tr:nth-child(odd) td {
    background: rgba(71, 72, 135, 0.2);
  }

  /* تأثيرات للأعمدة */
  .v-data-table td:first-child {
    border-right: 3px solid rgba(219, 15, 172, 0.3);
  }

  .v-data-table td:last-child {
    border-left: 3px solid rgba(219, 15, 172, 0.3);
  }

  /* تأثير للصور الشخصية */
  .v-avatar {
    border: 2px solid rgba(59, 130, 246, 0.3);
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    backdrop-filter: blur(10px);
  }

  .v-avatar:hover {
    transform: scale(1.05);
    box-shadow: 0 8px 25px rgba(59, 130, 246, 0.3);
    border-color: rgba(59, 130, 246, 0.6);
  }

  /* تحديث الـ chips */
  .v-chip {
    backdrop-filter: blur(10px);
    border: 1px solid rgba(148, 163, 184, 0.2);
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
  }

  .v-chip:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }

  @keyframes shimmerHeader {
    0% { transform: translateX(-100%); }
    100% { transform: translateX(100%); }
  }

  /* تحسين شريط التنقل (Pagination) */
  .v-pagination {
    background: linear-gradient(135deg, rgba(71, 72, 135, 0.9) 0%, rgba(76, 86, 194, 0.8) 100%);
    border-radius: 15px;
    padding: 15px;
    box-shadow: 0 8px 25px rgba(219, 15, 172, 0.2);
    border: 2px solid rgba(210, 0, 171, 0.15);
  }

  .v-pagination .v-btn {
    background: linear-gradient(135deg, rgba(219, 15, 172, 0.7) 0%, rgba(236, 72, 153, 0.5) 100%) !important;
    color: white !important;
    margin: 0 3px;
  border-radius: 12px;
    min-width: 40px;
    height: 40px;
    box-shadow: 0 4px 12px rgba(219, 15, 172, 0.3);
  }

  .v-pagination .v-btn:hover {
    background: linear-gradient(135deg, rgba(236, 72, 153, 0.8) 0%, rgba(219, 15, 172, 0.6) 100%) !important;
    transform: translateY(-2px);
    box-shadow: 0 6px 18px rgba(219, 15, 172, 0.4);
  }

  .v-pagination .v-btn--active {
    background: linear-gradient(135deg, rgba(41, 5, 23, 0.9) 0%, rgba(219, 15, 172, 0.8) 100%) !important;
    box-shadow: 0 6px 20px rgba(41, 5, 23, 0.5);
    transform: scale(1.1);
  }

  /* تحسين عداد العناصر */
  .v-data-table-footer {
    background: linear-gradient(135deg, rgba(71, 85, 105, 0.9) 0%, rgba(51, 65, 85, 0.8) 100%);
    color: #ffffff !important;
    border-radius: 0 0 24px 24px;
    padding: 20px 24px;
    backdrop-filter: blur(10px);
    border-top: 1px solid rgba(148, 163, 184, 0.3);
  }

  /* إصلاح النصوص في التذييل */
  .v-data-table-footer *,
  .v-data-table-footer * * {
    color: #ffffff !important;
  }

  .v-data-table-footer .v-select {
    background: rgba(59, 130, 246, 0.1);
    border-radius: 12px;
    border: 1px solid rgba(59, 130, 246, 0.3);
    backdrop-filter: blur(10px);
  }

  .v-data-table-footer .v-pagination {
    background: transparent;
  }

  .v-data-table-footer .v-pagination .v-btn {
    background: rgba(59, 130, 246, 0.1);
    border: 1px solid rgba(59, 130, 246, 0.3);
    color: #f8fafc;
    backdrop-filter: blur(10px);
  }

  .v-data-table-footer .v-pagination .v-btn:hover {
    background: rgba(59, 130, 246, 0.2);
    transform: translateY(-2px);
  }

  /* تأثيرات إضافية للجدول الحديث */
  .v-data-table {
    animation: fadeInUp 0.6s ease-out;
  }


  /* تأثير للصفوف */
  .v-data-table tbody tr {
    animation: slideInRow 0.4s ease-out;
    animation-fill-mode: both;
  }

  .v-data-table tbody tr:nth-child(1) { animation-delay: 0.1s; }
  .v-data-table tbody tr:nth-child(2) { animation-delay: 0.2s; }
  .v-data-table tbody tr:nth-child(3) { animation-delay: 0.3s; }
  .v-data-table tbody tr:nth-child(4) { animation-delay: 0.4s; }
  .v-data-table tbody tr:nth-child(5) { animation-delay: 0.5s; }

  @keyframes slideInRow {
    from {
      opacity: 0;
      transform: translateX(-20px);
    }
    to {
      opacity: 1;
      transform: translateX(0);
    }
  }

  /* تأثير للعناصر التفاعلية */
  .v-data-table td {
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .v-data-table td:hover {
    transform: scale(1.02);
  }

  /* ========================================
     إصلاح وضوح النصوص في جميع الحقول
     - ضمان وضوح جميع النصوص في الجدول
     - تحسين التباين والقراءة
     - إصلاح النصوص المختلطة
     ======================================== */

  /* إصلاح شامل لجميع النصوص في الجدول */
  .v-data-table *,
  .v-data-table * *,
  .v-data-table .d-flex,
  .v-data-table .d-flex *,
  .v-data-table .align-center,
  .v-data-table .align-center * {
    color: #1a1a1a !important;
  }

  /* إصلاح النصوص في خلايا المستخدم */
  .v-data-table .font-weight-medium,
  .v-data-table .text-caption,
  .v-data-table .text-body-2 {
    color: #1a1a1a !important;
    font-weight: 500 !important;
  }

  /* إصلاح النصوص في الـ chips */
  .v-data-table .v-chip,
  .v-data-table .v-chip *,
  .v-data-table .v-chip .v-chip__content,
  .v-data-table .v-chip .v-chip__content * {
    color: #ffffff !important;
    font-weight: 600 !important;
  }

  /* إصلاح النصوص في الصور الشخصية */
  .v-data-table .v-avatar,
  .v-data-table .v-avatar *,
  .v-data-table .v-img,
  .v-data-table .v-img * {
    color: #1a1a1a !important;
  }

  /* إصلاح النصوص في أزرار الإجراءات */
  .v-data-table .v-btn,
  .v-data-table .v-btn *,
  .v-data-table .v-btn .v-icon,
  .v-data-table .v-btn .v-btn__content {
    color: #ffffff !important;
  }

  /* إصلاح النصوص في التواريخ */
  .v-data-table .text-body-2 {
    color: #1a1a1a !important;
    font-weight: 500 !important;
  }

  /* إصلاح شامل لجميع العناصر النصية */
  .v-data-table .text-grey,
  .v-data-table .text-caption,
  .v-data-table .text-body-1,
  .v-data-table .text-body-2,
  .v-data-table .text-subtitle-1,
  .v-data-table .text-subtitle-2 {
    color: #1a1a1a !important;
    font-weight: 500 !important;
  }

  /* إصلاح النصوص في جميع العناصر */
  .v-data-table span,
  .v-data-table div,
  .v-data-table p,
  .v-data-table label {
    color: #1a1a1a !important;
  }

  /* إصلاح النصوص في العناصر المحددة */
  .v-data-table .v-list-item,
  .v-data-table .v-list-item *,
  .v-data-table .v-list-item-title,
  .v-data-table .v-list-item-subtitle {
    color: #1a1a1a !important;
  }

  /* إصلاح نهائي شامل لجميع النصوص */
  .v-data-table .v-application,
  .v-data-table .v-application *,
  .v-data-table .v-application .v-data-table,
  .v-data-table .v-application .v-data-table * {
    color: #1a1a1a !important;
  }

  /* إصلاح النصوص في جميع الحالات */
  .v-data-table .v-data-table__wrapper,
  .v-data-table .v-data-table__wrapper *,
  .v-data-table .v-data-table__wrapper table,
  .v-data-table .v-data-table__wrapper table * {
    color: #1a1a1a !important;
  }

  /* إصلاح خاص للنصوص المختلطة */
  .v-data-table .me-3,
  .v-data-table .me-3 *,
  .v-data-table .gap-1,
  .v-data-table .gap-1 * {
    color: #1a1a1a !important;
  }

  /* تنسيق نافذة إضافة المستخدم */
  .add-user-dialog {
    background: #ffffff !important;
    border-radius: 16px !important;
    overflow: visible !important;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15) !important;
    display: flex !important;
    flex-direction: column !important;
    max-height: 90vh !important;
  }

  .add-user-dialog .v-card-text {
    overflow-y: auto !important;
    flex: 1 !important;
    display: block !important;
    visibility: visible !important;
    opacity: 1 !important;
  }

  .add-user-dialog .dialog-actions {
    flex-shrink: 0;
    position: sticky;
    bottom: 0;
    background: #f8fafc;
    z-index: 10;
  }

  .dialog-header {
    background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
    color: #ffffff;
    padding: 20px 24px;
    border-bottom: none;
    display: flex;
    align-items: center;
    justify-content: space-between;
    min-height: 60px;
  }

  .dialog-title {
    display: flex;
    align-items: center;
    font-weight: 600;
    font-size: 1rem;
    gap: 8px;
  }

  .dialog-title h2 {
    color: #ffffff;
    margin: 0;
    font-weight: 600;
    font-size: 1.1rem;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
    line-height: 1.2;
  }

  .dialog-title .v-icon {
    color: #ffffff;
    filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.2));
    font-size: 24px !important;
    width: 24px !important;
    height: 24px !important;
  }

  .close-btn {
    color: #ffffff !important;
    background: rgba(255, 255, 255, 0.2) !important;
    border-radius: 50% !important;
    transition: all 0.3s ease !important;
    width: 32px !important;
    height: 32px !important;
    min-width: 32px !important;
    padding: 0 !important;
  }

  .close-btn:hover {
    background: rgba(255, 255, 255, 0.3) !important;
    transform: rotate(90deg);
  }

  .close-btn .v-icon {
    font-size: 18px !important;
    width: 18px !important;
    height: 18px !important;
  }

  .dialog-content {
    padding: 32px !important;
    background: #ffffff !important;
  }

  /* تنسيق الحقول - عرض كامل */
  .add-user-dialog .v-row {
    margin: -12px !important;
  }

  .add-user-dialog .v-col {
    padding: 12px !important;
  }

  .add-user-dialog .v-text-field,
  .add-user-dialog .v-select,
  .add-user-dialog .v-textarea {
    width: 100% !important;
  }

  .add-user-dialog .v-input {
    width: 100% !important;
  }

  .add-user-dialog .v-input__control {
    width: 100% !important;
  }

  .add-user-dialog .v-field {
    width: 100% !important;
  }

  .add-user-dialog .v-field__field {
    width: 100% !important;
  }

  .add-user-dialog .v-field__input {
    padding: 16px 14px !important;
    min-height: 56px !important;
    font-size: 1rem !important;
  }

  .add-user-dialog .v-field__prepend-inner {
    padding-right: 12px !important;
  }

  .add-user-dialog .v-field__append-inner {
    padding-left: 12px !important;
  }

  .add-user-dialog .v-input__details {
    padding-top: 6px !important;
    padding-inline-start: 16px !important;
  }

  .add-user-dialog .v-messages__message {
    font-size: 0.85rem !important;
  }

  .dialog-actions {
    padding: 12px 16px;
    background: #f8fafc;
    border-top: 1px solid #e2e8f0;
    display: flex !important;
    justify-content: flex-end;
    gap: 8px;
    align-items: center;
    min-height: 50px;
    visibility: visible !important;
    opacity: 1 !important;
  }

  .add-user-dialog .dialog-actions {
    display: flex !important;
    visibility: visible !important;
    opacity: 1 !important;
  }

  .add-user-dialog .dialog-actions .v-btn {
    display: inline-flex !important;
    visibility: visible !important;
    opacity: 1 !important;
  }

  /* تنسيق أزرار نافذة إضافة المستخدم */
  .add-user-dialog .dialog-actions .v-btn {
    border-radius: 8px !important;
    font-weight: 600 !important;
    text-transform: none !important;
    padding: 6px 16px !important;
    min-width: 100px !important;
    font-size: 0.8rem !important;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.12) !important;
    transition: all 0.3s ease !important;
    height: 32px !important;
    min-height: 32px !important;
  }

  .add-user-dialog .dialog-actions .v-btn[color="primary"] {
    background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
    color: white !important;
    box-shadow: 0 2px 8px rgba(25, 118, 210, 0.25) !important;
  }

  .add-user-dialog .dialog-actions .v-btn[color="primary"]:hover {
    background: linear-gradient(135deg, #1565c0 0%, #0d47a1 100%) !important;
    box-shadow: 0 4px 12px rgba(25, 118, 210, 0.35) !important;
    transform: translateY(-1px) !important;
  }

  .add-user-dialog .dialog-actions .v-btn[color="primary"]:disabled,
  .add-user-dialog .save-btn-dialog:disabled {
    background: #9ca3af !important;
    color: white !important;
    opacity: 0.6 !important;
    cursor: not-allowed !important;
  }

  .add-user-dialog .save-btn-dialog {
    background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
    color: white !important;
    box-shadow: 0 2px 8px rgba(25, 118, 210, 0.25) !important;
    min-width: 110px !important;
    z-index: 100 !important;
    position: relative !important;
    visibility: visible !important;
    opacity: 1 !important;
    display: inline-flex !important;
    height: 32px !important;
    min-height: 32px !important;
    padding: 6px 16px !important;
    font-size: 0.8rem !important;
    font-weight: 600 !important;
    border: none !important;
    cursor: pointer !important;
    border-radius: 8px !important;
  }

  .add-user-dialog .save-btn-dialog .v-icon {
    display: inline-block !important;
    visibility: visible !important;
    opacity: 1 !important;
    margin-inline-end: 4px !important;
    font-size: 16px !important;
    width: 16px !important;
    height: 16px !important;
  }

  .add-user-dialog .save-btn-dialog span {
    display: inline-block !important;
    visibility: visible !important;
    opacity: 1 !important;
  }

  .add-user-dialog .save-btn-dialog:hover {
    background: linear-gradient(135deg, #1565c0 0%, #0d47a1 100%) !important;
    box-shadow: 0 4px 12px rgba(25, 118, 210, 0.35) !important;
    transform: translateY(-1px) !important;
  }

  .add-user-dialog .cancel-btn-dialog {
    background: #ffffff !important;
    color: #6b7280 !important;
    border: 1.5px solid #d1d5db !important;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08) !important;
    min-width: 80px !important;
    height: 32px !important;
    min-height: 32px !important;
    padding: 6px 16px !important;
    font-size: 0.8rem !important;
    border-radius: 8px !important;
  }

  .add-user-dialog .cancel-btn-dialog:hover {
    background: #f9fafb !important;
    border-color: #9ca3af !important;
    transform: translateY(-1px) !important;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.12) !important;
  }

  .add-user-dialog .dialog-actions .v-btn[color="grey"] {
    background: #ffffff !important;
    color: #6b7280 !important;
    border: 2px solid #d1d5db !important;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
  }

  .add-user-dialog .dialog-actions .v-btn[color="grey"]:hover {
    background: #f9fafb !important;
    border-color: #9ca3af !important;
    transform: translateY(-2px) !important;
  }

  .user-avatar-upload {
    border: 3px solid rgba(219, 15, 172, 0.5);
    box-shadow: 0 8px 25px rgba(219, 15, 172, 0.3);
    transition: all 0.3s ease;
  }

  .user-avatar-upload:hover {
    transform: scale(1.05);
    box-shadow: 0 12px 35px rgba(219, 15, 172, 0.5);
  }

  /* تنسيق الحقول في النافذة */
  .add-user-dialog .v-text-field,
  .add-user-dialog .v-select,
  .add-user-dialog .v-textarea {
    margin-bottom: 8px;
  }

  .add-user-dialog .v-text-field .v-field,
  .add-user-dialog .v-select .v-field,
  .add-user-dialog .v-textarea .v-field {
    background: #ffffff !important;
    border-radius: 12px;
    border: 1px solid #d1d5db !important;
    transition: all 0.3s ease;
  }

  .add-user-dialog .v-text-field .v-field:hover,
  .add-user-dialog .v-select .v-field:hover,
  .add-user-dialog .v-textarea .v-field:hover {
    border-color: #1976d2 !important;
    box-shadow: 0 2px 8px rgba(25, 118, 210, 0.15);
  }

  .add-user-dialog .v-field--focused .v-field,
  .add-user-dialog .v-text-field .v-field--focused,
  .add-user-dialog .v-select .v-field--focused,
  .add-user-dialog .v-textarea .v-field--focused {
    border-color: #1976d2 !important;
    box-shadow: 0 0 0 2px rgba(25, 118, 210, 0.2) !important;
  }

  /* Fix input text color - make it black/dark */
  .add-user-dialog .v-field__input,
  .add-user-dialog .v-field__input input,
  .add-user-dialog .v-text-field input,
  .add-user-dialog .v-select input,
  .add-user-dialog .v-textarea textarea {
    color: #1a1a1a !important;
    caret-color: #1976d2 !important;
  }

  .add-user-dialog .v-field__input::placeholder {
    color: #9ca3af !important;
  }

  /* Fix select display text */
  .add-user-dialog .v-select__selection,
  .add-user-dialog .v-select__selection-text {
    color: #1a1a1a !important;
  }

  /* Fix icons color */
  .add-user-dialog .v-field__prepend-inner .v-icon,
  .add-user-dialog .v-field__append-inner .v-icon {
    color: #6b7280 !important;
  }

  .add-user-dialog .v-field--focused .v-field__prepend-inner .v-icon,
  .add-user-dialog .v-field--focused .v-field__append-inner .v-icon {
    color: #1976d2 !important;
  }

  /* تأثيرات إضافية للنافذة */
  .add-user-dialog .v-btn {
    border-radius: 12px;
    font-weight: 600;
    text-transform: none;
    transition: all 0.3s ease;
  }

  .add-user-dialog .v-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
  }

  /* تحسين التحقق من صحة البيانات */
  .add-user-dialog .v-messages {
    color: #d32f2f;
    font-size: 0.875rem;
    font-weight: 500;
  }

  /* تأثيرات الأيقونات */
  .add-user-dialog .v-icon {
    transition: all 0.3s ease;
  }

  .add-user-dialog .v-text-field:hover .v-icon,
  .add-user-dialog .v-select:hover .v-icon,
  .add-user-dialog .v-textarea:hover .v-icon {
    color: rgba(219, 15, 172, 0.8) !important;
}

.chart-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
  background: linear-gradient(135deg, rgba(79, 70, 229, 0.05) 0%, rgba(124, 58, 237, 0.08) 100%);
  border-radius: 12px;
  border: 2px dashed rgba(79, 70, 229, 0.2);
}

/* التأثيرات المتحركة */

/* بطاقات الإحصائيات المحسنة */
.v-card[color="primary"] {
  background: linear-gradient(135deg, rgba(79, 70, 229, 0.1) 0%, rgba(79, 70, 229, 0.2) 100%) !important;
  border: 1px solid rgba(79, 70, 229, 0.3) !important;
}

.v-card[color="success"] {
  background: linear-gradient(135deg, rgba(34, 197, 94, 0.1) 0%, rgba(34, 197, 94, 0.2) 100%) !important;
  border: 1px solid rgba(34, 197, 94, 0.3) !important;
}

.v-card[color="warning"] {
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.1) 0%, rgba(249, 115, 22, 0.2) 100%) !important;
  border: 1px solid rgba(249, 115, 22, 0.3) !important;
}

.v-card[color="info"] {
  background: linear-gradient(135deg, rgba(124, 58, 237, 0.1) 0%, rgba(124, 58, 237, 0.2) 100%) !important;
  border: 1px solid rgba(124, 58, 237, 0.3) !important;
}

/* الشرائح المحسنة */
.v-chip {
  border-radius: 15px !important;
  font-weight: 700 !important;
  box-shadow: 0 4px 15px rgba(219, 15, 172, 0.3) !important;
  transition: all 0.3s ease;
  background: linear-gradient(135deg, rgba(219, 15, 172, 0.8) 0%, rgba(236, 72, 153, 0.6) 100%) !important;
  color: #ffffff !important;
  border: 1px solid rgba(255, 255, 255, 0.2);
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
}

.v-chip:hover {
  transform: translateY(-2px) scale(1.05);
  box-shadow: 0 8px 25px rgba(219, 15, 172, 0.5) !important;
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.9) 0%, rgba(219, 15, 172, 0.7) 100%) !important;
}

/* شرائح الحالة */
.v-chip.status-active {
  background: linear-gradient(135deg, rgba(34, 197, 94, 0.8) 0%, rgba(16, 185, 129, 0.6) 100%) !important;
}

.v-chip.status-inactive {
  background: linear-gradient(135deg, rgba(239, 68, 68, 0.8) 0%, rgba(220, 38, 38, 0.6) 100%) !important;
}

.v-chip.status-pending {
  background: linear-gradient(135deg, rgba(249, 115, 22, 0.8) 0%, rgba(234, 88, 12, 0.6) 100%) !important;
}

/* التصميم المتجاوب */
@media (max-width: 960px) {
  .main-content {
    padding: 16px;
  }
  
  .header-content {
    padding: 1.5rem;
  }
  
  .page-title {
    font-size: 2.2rem !important;
  }
  
  .title-icon {
    font-size: 3rem;
  }
}

@media (max-width: 600px) {
  .page-header {
    padding: 16px;
  }
  
  .header-content {
    padding: 1rem;
    flex-direction: column;
    text-align: center;
  }
  
  .page-title {
    font-size: 2rem !important;
  }
  
  .page-subtitle {
    font-size: 1.1rem !important;
  }
  
  .title-icon {
    font-size: 2.5rem;
  }
}

/* ========================================
   تنسيق القوائم المنسدلة في حوار إضافة المستخدم
   ======================================== */

/* تنسيق التسميات - ألوان سوداء قوية وواضحة */
.add-user-dialog .v-label,
.add-user-dialog .v-field__label,
.add-user-dialog .v-label--active,
.add-user-dialog .v-field__label--active,
.add-user-dialog .v-label--floating,
.add-user-dialog .v-field__label--floating {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
  background: white !important;
  padding: 0 8px !important;
  text-shadow: none !important;
  letter-spacing: 0.3px !important;
}

.add-user-dialog .v-field--focused .v-label,
.add-user-dialog .v-field--focused .v-field__label,
.add-user-dialog .v-field--focused .v-label--active,
.add-user-dialog .v-field--focused .v-field__label--active,
.add-user-dialog .v-field--focused .v-label--floating,
.add-user-dialog .v-field--focused .v-field__label--floating {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 0.875rem !important;
  opacity: 1 !important;
  background: white !important;
  padding: 0 8px !important;
}

/* تنسيق خاص لكل نوع من الحقول */
.add-user-dialog .v-text-field .v-label,
.add-user-dialog .v-text-field .v-field__label,
.add-user-dialog .v-select .v-label,
.add-user-dialog .v-select .v-field__label,
.add-user-dialog .v-textarea .v-label,
.add-user-dialog .v-textarea .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  background: white !important;
}

.add-user-dialog .v-text-field .v-field--focused .v-label,
.add-user-dialog .v-text-field .v-field--focused .v-field__label,
.add-user-dialog .v-select .v-field--focused .v-label,
.add-user-dialog .v-select .v-field--focused .v-field__label,
.add-user-dialog .v-textarea .v-field--focused .v-label,
.add-user-dialog .v-textarea .v-field--focused .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  background: white !important;
}

/* تنسيق الحدود - رمادي فاتح */
.add-user-dialog .v-field__outline {
  border-color: #d1d5db !important;
  border-width: 2px !important;
}

.add-user-dialog .v-field--focused .v-field__outline {
  border-color: #6b7280 !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(107, 114, 128, 0.1) !important;
}

.add-user-dialog .v-field:hover .v-field__outline {
  border-color: #9ca3af !important;
}

/* تنسيق القوائم المنسدلة - ألوان سوداء */
.add-user-dialog .black-dropdown-select :deep(.v-menu__content),
.add-user-dialog .black-dropdown-select :deep(.v-overlay__content) {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.add-user-dialog .black-dropdown-select :deep(.v-menu__content .v-list-item),
.add-user-dialog .black-dropdown-select :deep(.v-overlay__content .v-list-item) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  min-height: 48px !important;
  background: white !important;
}

.add-user-dialog .black-dropdown-select :deep(.v-menu__content .v-list-item:hover),
.add-user-dialog .black-dropdown-select :deep(.v-overlay__content .v-list-item:hover) {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.add-user-dialog .black-dropdown-select :deep(.v-menu__content .v-list-item--active),
.add-user-dialog .black-dropdown-select :deep(.v-overlay__content .v-list-item--active),
.add-user-dialog .black-dropdown-select :deep(.v-menu__content .v-list-item[aria-selected="true"]),
.add-user-dialog .black-dropdown-select :deep(.v-overlay__content .v-list-item[aria-selected="true"]) {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

.add-user-dialog .black-dropdown-select :deep(.v-list-item-title),
.add-user-dialog .black-dropdown-select :deep(.v-list-item__title),
.add-user-dialog .black-dropdown-select :deep(.v-list-item-content),
.add-user-dialog .black-dropdown-select :deep(.v-list-item__content) {
  color: inherit !important;
  background: inherit !important;
}

/* تنسيق شامل لجميع القوائم المنسدلة في الحوار */
.add-user-dialog .v-menu__content .v-list-item,
.add-user-dialog .v-overlay__content .v-list-item {
  color: #000000 !important;
  background: white !important;
}

.add-user-dialog .v-menu__content .v-list-item:hover,
.add-user-dialog .v-overlay__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.add-user-dialog .v-menu__content .v-list-item--active,
.add-user-dialog .v-overlay__content .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
}
</style>

<style>
/* تنسيقات شاملة للقوائم المنسدلة في حوار إضافة المستخدم */
.add-user-dialog .v-menu__content,
.add-user-dialog .v-overlay__content {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.add-user-dialog .v-menu__content .v-list-item,
.add-user-dialog .v-overlay__content .v-list-item {
  color: #000000 !important;
  background: white !important;
}

.add-user-dialog .v-menu__content .v-list-item:hover,
.add-user-dialog .v-overlay__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.add-user-dialog .v-menu__content .v-list-item--active,
.add-user-dialog .v-overlay__content .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
}
</style>



<style>
/* ========================================
   إصلاح شامل لألوان النصوص في حوار إضافة المستخدم
   ======================================== */

/* إصلاح لون النص المكتوب في الحقول - أسود داكن */
.add-user-dialog .v-field__input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.add-user-dialog .v-field__input input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  caret-color: #1976d2 !important;
}

.add-user-dialog .v-text-field input,
.add-user-dialog .v-text-field .v-field input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.add-user-dialog .v-select .v-field__input,
.add-user-dialog .v-select input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.add-user-dialog .v-textarea textarea {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

/* إصلاح خلفية الحقول - أبيض */
.add-user-dialog .v-field,
.add-user-dialog .v-field__field,
.add-user-dialog .v-field__overlay {
  background: #ffffff !important;
  background-color: #ffffff !important;
}

/* إصلاح نص القائمة المنسدلة */
.add-user-dialog .v-select__selection,
.add-user-dialog .v-select__selection-text,
.add-user-dialog .v-select .v-select__selection {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

/* ========================================
   تنسيق القائمة المنسدلة للأدوار - خفيف
   ======================================== */
.v-menu > .v-overlay__content > .v-list,
.v-select__content .v-list {
  background: #ffffff !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #e5e7eb !important;
  padding: 8px !important;
}

.v-menu > .v-overlay__content > .v-list .v-list-item,
.v-select__content .v-list-item {
  background: #ffffff !important;
  color: #1a1a1a !important;
  border-radius: 6px !important;
  margin-bottom: 4px !important;
  min-height: 44px !important;
}

.v-menu > .v-overlay__content > .v-list .v-list-item:hover,
.v-select__content .v-list-item:hover {
  background: #f3f4f6 !important;
}

.v-menu > .v-overlay__content > .v-list .v-list-item--active,
.v-select__content .v-list-item--active {
  background: #e0f2fe !important;
  color: #1976d2 !important;
}

.v-menu > .v-overlay__content > .v-list .v-list-item-title,
.v-select__content .v-list-item-title {
  color: #1a1a1a !important;
  font-size: 0.95rem !important;
  font-weight: 500 !important;
}

.v-menu > .v-overlay__content > .v-list .v-list-item--active .v-list-item-title,
.v-select__content .v-list-item--active .v-list-item-title {
  color: #1976d2 !important;
  font-weight: 600 !important;
}

/* تنسيقات التسميات - العناوين */
.add-user-dialog .v-label,
.add-user-dialog .v-field-label {
  color: #374151 !important;
  font-weight: 500 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
}

.add-user-dialog .v-field-label--floating,
.add-user-dialog .v-label--floating {
  color: #1976d2 !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  background: white !important;
  padding: 0 8px !important;
  transform: translateY(-50%) !important;
}

/* تنسيق حقول outlined مع العناوين */
.add-user-dialog .v-field--variant-outlined .v-field__outline__notch {
  border-width: 0 !important;
}

.add-user-dialog .v-field--variant-outlined .v-field__outline__start,
.add-user-dialog .v-field--variant-outlined .v-field__outline__end {
  border-color: #d1d5db !important;
  border-width: 2px !important;
}

.add-user-dialog .v-field--variant-outlined:hover .v-field__outline__start,
.add-user-dialog .v-field--variant-outlined:hover .v-field__outline__end {
  border-color: #9ca3af !important;
}

.add-user-dialog .v-field--variant-outlined.v-field--focused .v-field__outline__start,
.add-user-dialog .v-field--variant-outlined.v-field--focused .v-field__outline__end {
  border-color: #1976d2 !important;
}

.add-user-dialog .v-field--focused .v-label,
.add-user-dialog .v-field--focused .v-field__label {
  color: #1976d2 !important;
  font-weight: 600 !important;
  opacity: 1 !important;
  background: white !important;
}

.add-user-dialog .v-text-field .v-label,
.add-user-dialog .v-select .v-label,
.add-user-dialog .v-textarea .v-label {
  color: #374151 !important;
  font-weight: 600 !important;
  opacity: 1 !important;
}

/* إصلاح ألوان الأيقونات */
.add-user-dialog .v-field__prepend-inner .v-icon,
.add-user-dialog .v-field__append-inner .v-icon {
  color: #6b7280 !important;
}

/* إصلاح الحدود */
.add-user-dialog .v-field--variant-outlined .v-field__outline__start,
.add-user-dialog .v-field--variant-outlined .v-field__outline__end,
.add-user-dialog .v-field--variant-outlined .v-field__outline__notch::before,
.add-user-dialog .v-field--variant-outlined .v-field__outline__notch::after {
  border-color: #d1d5db !important;
}

.add-user-dialog .v-field--focused.v-field--variant-outlined .v-field__outline__start,
.add-user-dialog .v-field--focused.v-field--variant-outlined .v-field__outline__end,
.add-user-dialog .v-field--focused.v-field--variant-outlined .v-field__outline__notch::before,
.add-user-dialog .v-field--focused.v-field--variant-outlined .v-field__outline__notch::after {
  border-color: #1976d2 !important;
}

/* إصلاح رسائل التلميحات */
.add-user-dialog .v-messages__message {
  color: #6b7280 !important;
}

/* ========================================
   إصلاح شامل لحجم وعرض حقول نموذج إضافة المستخدم
   ======================================== */
.add-user-dialog {
  width: 100% !important;
  max-width: 950px !important;
}

.add-user-dialog .v-card {
  width: 100% !important;
}

.add-user-dialog .v-card-text,
.add-user-dialog .dialog-content {
  width: 100% !important;
  padding: 32px 40px !important;
  box-sizing: border-box !important;
}

.add-user-dialog .v-form {
  width: 100% !important;
}

.add-user-dialog .v-row {
  width: 100% !important;
  margin: 0 -16px !important;
}

.add-user-dialog .v-col {
  padding: 16px !important;
  flex: 0 0 50% !important;
  max-width: 50% !important;
}

.add-user-dialog .v-col[class*="cols-12"]:not([class*="md-6"]):not([class*="sm-6"]) {
  flex: 0 0 100% !important;
  max-width: 100% !important;
}

.add-user-dialog .v-text-field,
.add-user-dialog .v-select {
  width: 100% !important;
  min-width: 100% !important;
}

.add-user-dialog .v-input {
  width: 100% !important;
  min-width: 100% !important;
}

.add-user-dialog .v-input__control {
  width: 100% !important;
  min-width: 100% !important;
}

.add-user-dialog .v-field {
  width: 100% !important;
  min-width: 100% !important;
  min-height: 56px !important;
}

.add-user-dialog .v-field__field {
  width: 100% !important;
  min-height: 56px !important;
}

.add-user-dialog .v-field__input {
  width: 100% !important;
  min-height: 56px !important;
  padding: 16px !important;
  font-size: 1rem !important;
}

.add-user-dialog .v-field__input input {
  width: 100% !important;
  font-size: 1rem !important;
}

.add-user-dialog .v-field__outline {
  width: 100% !important;
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

/* ========================================
   إصلاح شامل لألوان النصوص في حوار تعديل المستخدم
   ======================================== */

/* إصلاح لون النص المكتوب في الحقول */
.edit-user-dialog .v-field__input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.edit-user-dialog .v-field__input input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  caret-color: #1976d2 !important;
}

.edit-user-dialog .v-text-field input,
.edit-user-dialog .v-text-field .v-field input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.edit-user-dialog .v-select .v-field__input,
.edit-user-dialog .v-select input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

/* إصلاح خلفية الحقول */
.edit-user-dialog .v-field,
.edit-user-dialog .v-field__field,
.edit-user-dialog .v-field__overlay {
  background: #ffffff !important;
  background-color: #ffffff !important;
}

/* إصلاح نص القائمة المنسدلة */
.edit-user-dialog .v-select__selection,
.edit-user-dialog .v-select__selection-text,
.edit-user-dialog .v-select .v-select__selection {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

/* تنسيقات التسميات - العناوين */
.edit-user-dialog .v-label,
.edit-user-dialog .v-field-label {
  color: #374151 !important;
  -webkit-text-fill-color: #374151 !important;
  font-weight: 500 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
}

.edit-user-dialog .v-field-label--floating,
.edit-user-dialog .v-label--floating {
  color: #1976d2 !important;
  -webkit-text-fill-color: #1976d2 !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
  background: white !important;
  padding: 0 8px !important;
}

.edit-user-dialog .v-text-field .v-label,
.edit-user-dialog .v-select .v-label {
  color: #374151 !important;
  -webkit-text-fill-color: #374151 !important;
  font-weight: 600 !important;
  opacity: 1 !important;
}

/* إصلاح ألوان الأيقونات */
.edit-user-dialog .v-field__prepend-inner .v-icon,
.edit-user-dialog .v-field__append-inner .v-icon {
  color: #6b7280 !important;
}

/* إصلاح رسائل التلميحات */
.edit-user-dialog .v-messages__message {
  color: #6b7280 !important;
}

/* ========================================
   إصلاح شامل لألوان نافذة تغيير كلمة المرور
   ======================================== */

.reset-password-dialog .dialog-header {
  background: #ffffff !important;
  color: #1a1a1a !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.2);
}

.reset-password-dialog .dialog-title h2 {
  color: #1a1a1a !important;
}

.reset-password-dialog .dialog-title .v-icon {
  color: #1a1a1a !important;
}

.reset-password-dialog .v-card-text {
  background: #ffffff !important;
}

.reset-password-dialog .v-field {
  background: #ffffff !important;
}

.reset-password-dialog .v-field * {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.reset-password-dialog .v-field input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  caret-color: #1976d2 !important;
}

.reset-password-dialog .v-label,
.reset-password-dialog .v-field-label {
  color: #374151 !important;
  -webkit-text-fill-color: #374151 !important;
  opacity: 1 !important;
}

.reset-password-dialog .v-field-label--floating,
.reset-password-dialog .v-label--floating {
  color: #1976d2 !important;
  -webkit-text-fill-color: #1976d2 !important;
  background: white !important;
}

.reset-password-dialog .v-messages__message {
  color: #6b7280 !important;
}

.reset-password-dialog h4 {
  color: #1a1a1a !important;
}

.reset-password-dialog .text-caption {
  color: #6b7280 !important;
}

/* ========================================
   إصلاح ألوان نافذة تفاصيل المستخدم
   ======================================== */

.view-user-dialog .v-card {
  background: #ffffff !important;
}

.view-user-dialog .dialog-header {
  background: #ffffff !important;
  color: #1a1a1a !important;
}

.view-user-dialog .dialog-title h2 {
  color: #1a1a1a !important;
}

.view-user-dialog .dialog-title .v-icon {
  color: #1a1a1a !important;
}

.view-user-dialog .dialog-content {
  background: #ffffff !important;
}

.view-user-dialog .v-field {
  background: #ffffff !important;
}

.view-user-dialog .v-field__input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.view-user-dialog .v-field__input input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
}

.view-user-dialog input {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  opacity: 1 !important;
}

.view-user-dialog input[readonly] {
  color: #1a1a1a !important;
  -webkit-text-fill-color: #1a1a1a !important;
  opacity: 1 !important;
}

.view-user-dialog .v-label {
  color: #374151 !important;
  -webkit-text-fill-color: #374151 !important;
  opacity: 1 !important;
}

.view-user-dialog .v-field-label {
  color: #374151 !important;
  -webkit-text-fill-color: #374151 !important;
  opacity: 1 !important;
}

.view-user-dialog .close-btn {
  color: #64748b !important;
}

</style>
