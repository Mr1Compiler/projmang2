<template>
  <div>
      <!-- الشريط العلوي -->
      <v-app-bar
        flat
        height="70"
        class="top-bar"
      >
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
          color="primary"
          prepend-icon="mdi-plus"
          class="me-3"
          @click="showAddUserDialog = true"
        >
          إضافة مستخدم
        </v-btn>

        <!-- الإشعارات -->
        <v-btn
          icon="mdi-bell"
          variant="text"
          class="me-2"
        >
          <v-badge
            color="pink"
            dot
            floating
          />
        </v-btn>

        <!-- صورة المستخدم -->
        <v-avatar
          size="40"
          class="me-2"
        >
          <v-img src="https://randomuser.me/api/portraits/men/1.jpg" />
        </v-avatar>
      </v-app-bar>

      <!-- المحتوى -->
      <div class="main-content pa-6">
        <!-- شريط العنوان الرئيسي -->
        <div class="page-header glass-effect gradient-animation">
          <div class="header-top-content">
            <h1 class="page-title">إدارة المستخدمين</h1>
            <span class="page-icon">👥</span>
          </div>
          <p class="page-subtitle">نظام شامل لإدارة حسابات المستخدمين والصلاحيات</p>
        </div>

        <!-- الإحصائيات -->
        <div class="stats-container mb-6">
          <v-row>
            <v-col cols="12" md="3">
              <v-card class="pa-4 text-center" color="primary" variant="tonal">
                <v-icon size="48" color="primary" class="mb-2">mdi-account-multiple</v-icon>
                <h3 class="text-h4 font-weight-bold">{{ totalUsers }}</h3>
                <p class="text-subtitle-1">إجمالي المستخدمين</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="pa-4 text-center" color="success" variant="tonal">
                <v-icon size="48" color="success" class="mb-2">mdi-check-circle</v-icon>
                <h3 class="text-h4 font-weight-bold">{{ activeUsers }}</h3>
                <p class="text-subtitle-1">مستخدمين نشطين</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="pa-4 text-center" color="warning" variant="tonal">
                <v-icon size="48" color="warning" class="mb-2">mdi-account-clock</v-icon>
                <h3 class="text-h4 font-weight-bold">{{ pendingUsers }}</h3>
                <p class="text-subtitle-1">مستخدمين معلقين</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="pa-4 text-center" color="info" variant="tonal">
                <v-icon size="48" color="info" class="mb-2">mdi-shield-account</v-icon>
                <h3 class="text-h4 font-weight-bold">{{ adminUsers }}</h3>
                <p class="text-subtitle-1">مدراء النظام</p>
              </v-card>
            </v-col>
          </v-row>
        </div>

        <!-- فلاتر البحث -->
        <v-card class="filters-card mb-6" elevation="2">
          <v-card-text class="pa-4">
            <v-row>
              <v-col cols="12" md="3">
                <v-select
                  v-model="selectedRole"
                  :items="roles"
                  label="الدور"
                  variant="outlined"
                  density="compact"
                  clearable
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-select
                  v-model="selectedStatus"
                  :items="statusOptions"
                  label="الحالة"
                  variant="outlined"
                  density="compact"
                  clearable
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-select
                  v-model="selectedDepartment"
                  :items="departments"
                  label="القسم"
                  variant="outlined"
                  density="compact"
                  clearable
                />
              </v-col>
              <v-col cols="12" md="3">
                <v-btn
                  color="primary"
                  prepend-icon="mdi-filter"
                  @click="applyFilters"
                  class="mt-2"
                >
                  تطبيق الفلاتر
                </v-btn>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>

        <!-- جدول المستخدمين -->
        <v-card class="users-table" elevation="2">
          <v-card-title class="d-flex align-center justify-space-between">
            <span class="text-h5 font-weight-bold">قائمة المستخدمين</span>
            <div class="d-flex gap-2">
              <v-btn
                color="success"
                prepend-icon="mdi-download"
                variant="outlined"
              >
                تصدير البيانات
              </v-btn>
              <v-btn
                color="info"
                prepend-icon="mdi-account-plus"
                variant="outlined"
              >
                إضافة متعدد
              </v-btn>
            </div>
          </v-card-title>

          <v-data-table
            :headers="headers"
            :items="filteredUsers"
            :search="searchQuery"
            class="elevation-0"
            :items-per-page="10"
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
                icon="mdi-pencil"
                size="small"
                variant="elevated"
                class="edit-btn"
                data-action="edit"
                @click="editUser(item)"
              />
              <v-btn
                icon="mdi-key"
                size="small"
                variant="elevated"
                class="reset-btn"
                data-action="reset"
                @click="resetPassword(item)"
              />
              <v-btn
                icon="mdi-delete"
                size="small"
                variant="elevated"
                class="delete-btn"
                data-action="delete"
                @click="deleteUser(item)"
              />
            </template>
          </v-data-table>
        </v-card>

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
  </div>

  <!-- نافذة حوار إضافة مستخدم جديد -->
  <v-dialog v-model="showAddUserDialog" max-width="800px" persistent>
    <v-card class="add-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="primary" class="me-3">mdi-account-plus</v-icon>
          <h2>إضافة مستخدم جديد</h2>
        </div>
        <v-btn 
          icon="mdi-close" 
          variant="text" 
          @click="closeAddUserDialog"
          class="close-btn"
        />
      </v-card-title>
      
      <v-divider />
      
      <v-card-text class="dialog-content">
        <v-form ref="addUserForm" v-model="formValid" lazy-validation>
          <v-row>
            <!-- الصورة الشخصية -->
            <v-col cols="12" class="text-center mb-4">
              <v-avatar size="100" class="user-avatar-upload">
                <v-img 
                  :src="newUser.avatar || 'https://via.placeholder.com/100x100?text=صورة'"
                  alt="صورة المستخدم"
                />
              </v-avatar>
              <div class="mt-2">
                <v-btn 
                  size="small" 
                  color="primary" 
                  variant="outlined"
                  prepend-icon="mdi-camera"
                >
                  تحديد صورة
                </v-btn>
              </div>
            </v-col>

            <!-- الاسم الأول -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.firstName"
                label="الاسم الأول *"
                :rules="nameRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account"
              />
            </v-col>

            <!-- الاسم الأخير -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.lastName"
                label="الاسم الأخير *"
                :rules="nameRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account"
              />
            </v-col>

            <!-- البريد الإلكتروني -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.email"
                label="البريد الإلكتروني *"
                :rules="emailRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-email"
                type="email"
              />
            </v-col>

            <!-- رقم الهاتف -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.phone"
                label="رقم الهاتف"
                variant="outlined"
                prepend-inner-icon="mdi-phone"
                type="tel"
              />
            </v-col>

            <!-- الدور -->
            <v-col cols="12" md="6">
              <v-select
                v-model="newUser.role"
                :items="roleOptions"
                label="الدور *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account-tie"
                class="black-dropdown-select"
              />
            </v-col>

            <!-- القسم -->
            <v-col cols="12" md="6">
              <v-select
                v-model="newUser.department"
                :items="departmentOptions"
                label="القسم *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-office-building"
                class="black-dropdown-select"
              />
            </v-col>

            <!-- الحالة -->
            <v-col cols="12" md="6">
              <v-select
                v-model="newUser.status"
                :items="statusOptions"
                label="الحالة *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account-check"
                class="black-dropdown-select"
              />
            </v-col>

            <!-- كلمة المرور -->
            <v-col cols="12" md="6">
              <v-text-field
                v-model="newUser.password"
                label="كلمة المرور *"
                :rules="passwordRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-lock"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                @click:append-inner="showPassword = !showPassword"
              />
            </v-col>

            <!-- ملاحظات -->
            <v-col cols="12">
              <v-textarea
                v-model="newUser.notes"
                label="ملاحظات"
                variant="outlined"
                prepend-inner-icon="mdi-note-text"
                rows="3"
                auto-grow
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="closeAddUserDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          @click="saveNewUser"
          :loading="saving"
          :disabled="!formValid"
        >
          حفظ المستخدم
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- نافذة عرض تفاصيل المستخدم -->
  <v-dialog v-model="showViewUserDialog" max-width="600px">
    <v-card class="view-user-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="primary" class="me-3">mdi-account-details</v-icon>
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
      
      <v-card-text v-if="selectedUser" class="pa-6">
        <v-row>
          <v-col cols="12" class="text-center mb-4">
            <v-avatar size="100">
              <v-img :src="selectedUser.avatar" />
            </v-avatar>
            <h3 class="mt-3">{{ selectedUser.name }}</h3>
            <v-chip 
              :color="getStatusColor(selectedUser.status)" 
              size="small" 
              class="mt-2"
            >
              {{ getStatusText(selectedUser.status) }}
            </v-chip>
          </v-col>
          
          <v-col cols="12" md="6">
            <v-list density="compact">
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="primary">mdi-email</v-icon>
                </template>
                <v-list-item-title>البريد الإلكتروني</v-list-item-title>
                <v-list-item-subtitle>{{ selectedUser.email }}</v-list-item-subtitle>
              </v-list-item>
              
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="success">mdi-phone</v-icon>
                </template>
                <v-list-item-title>رقم الهاتف</v-list-item-title>
                <v-list-item-subtitle>{{ selectedUser.phone }}</v-list-item-subtitle>
              </v-list-item>
              
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="warning">mdi-account-tie</v-icon>
                </template>
                <v-list-item-title>الدور</v-list-item-title>
                <v-list-item-subtitle>{{ getRoleText(selectedUser.role) }}</v-list-item-subtitle>
              </v-list-item>
            </v-list>
          </v-col>
          
          <v-col cols="12" md="6">
            <v-list density="compact">
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="info">mdi-office-building</v-icon>
                </template>
                <v-list-item-title>القسم</v-list-item-title>
                <v-list-item-subtitle>{{ selectedUser.department }}</v-list-item-subtitle>
              </v-list-item>
              
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="purple">mdi-clock-outline</v-icon>
                </template>
                <v-list-item-title>آخر دخول</v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(selectedUser.lastLogin) }}</v-list-item-subtitle>
              </v-list-item>
              
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="teal">mdi-calendar-plus</v-icon>
                </template>
                <v-list-item-title>تاريخ الإنشاء</v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(selectedUser.createdAt) }}</v-list-item-subtitle>
              </v-list-item>
            </v-list>
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
          <v-icon size="32" color="success" class="me-3">mdi-account-edit</v-icon>
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
            
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.name"
                label="الاسم *"
                :rules="nameRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account"
              />
            </v-col>
            
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.email"
                label="البريد الإلكتروني *"
                :rules="emailRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-email"
              />
            </v-col>
            
            <v-col cols="12" md="6">
              <v-text-field
                v-model="selectedUser.phone"
                label="رقم الهاتف"
                variant="outlined"
                prepend-inner-icon="mdi-phone"
              />
            </v-col>
            
            <v-col cols="12" md="6">
              <v-select
                v-model="selectedUser.role"
                :items="roleOptions"
                label="الدور *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account-tie"
              />
            </v-col>
            
            <v-col cols="12" md="6">
              <v-select
                v-model="selectedUser.department"
                :items="departmentOptions"
                label="القسم *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-office-building"
              />
            </v-col>
            
            <v-col cols="12" md="6">
              <v-select
                v-model="selectedUser.status"
                :items="statusOptions"
                label="الحالة *"
                :rules="requiredRules"
                required
                variant="outlined"
                prepend-inner-icon="mdi-account-check"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider />

      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="closeEditUserDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="success"
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
          <v-icon size="32" color="warning" class="me-3">mdi-key-change</v-icon>
          <h2>إعادة تعيين كلمة المرور</h2>
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
            <v-img :src="selectedUser.avatar" />
          </v-avatar>
          <h4 class="mt-2">{{ selectedUser.name }}</h4>
          <p class="text-caption">{{ selectedUser.email }}</p>
        </div>
        
        <v-alert 
          type="warning" 
          variant="tonal" 
          class="mb-4"
        >
          سيتم إرسال كلمة مرور جديدة إلى البريد الإلكتروني للمستخدم
        </v-alert>
        
        <p class="text-body-2 text-center">
          هل أنت متأكد من إعادة تعيين كلمة المرور لهذا المستخدم؟
        </p>
      </v-card-text>
      
      <v-divider />
      
      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="closeResetPasswordDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="warning"
          variant="elevated"
          @click="confirmResetPassword"
          :loading="resetLoading"
        >
          إعادة تعيين
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <!-- نافذة تأكيد الحذف -->
  <v-dialog v-model="showDeleteConfirmDialog" max-width="500px">
    <v-card class="delete-confirm-dialog">
      <v-card-title class="dialog-header">
        <div class="dialog-title">
          <v-icon size="32" color="error" class="me-3">mdi-delete-alert</v-icon>
          <h2>تأكيد الحذف</h2>
        </div>
        <v-btn 
          icon="mdi-close" 
          variant="text" 
          @click="closeDeleteConfirmDialog"
          class="close-btn"
        />
      </v-card-title>
      
      <v-divider />
      
      <v-card-text v-if="selectedUser" class="pa-6">
        <div class="text-center mb-4">
          <v-avatar size="60">
            <v-img :src="selectedUser.avatar" />
          </v-avatar>
          <h4 class="mt-2">{{ selectedUser.name }}</h4>
          <p class="text-caption">{{ selectedUser.email }}</p>
        </div>
        
        <v-alert 
          type="error" 
          variant="tonal" 
          class="mb-4"
        >
          تحذير: هذا الإجراء لا يمكن التراجع عنه!
        </v-alert>
        
        <p class="text-body-2 text-center">
          هل أنت متأكد من حذف هذا المستخدم نهائياً؟
        </p>
      </v-card-text>
      
      <v-divider />
      
      <v-card-actions class="dialog-actions">
        <v-spacer />
        <v-btn
          color="grey"
          variant="outlined"
          @click="closeDeleteConfirmDialog"
          class="me-2"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="error"
          variant="elevated"
          @click="confirmDeleteUser"
          :loading="deleteLoading"
        >
          حذف نهائي
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'

// البيانات التفاعلية
const drawer = ref(true)
const searchQuery = ref('')
const selectedRole = ref(null)
const selectedStatus = ref(null)
const selectedDepartment = ref(null)

// متغيرات نافذة إضافة المستخدم
const showAddUserDialog = ref(false)
const formValid = ref(false)
const saving = ref(false)
const showPassword = ref(false)
const addUserForm = ref(null)

// متغيرات نوافذ الإجراءات
const showViewUserDialog = ref(false)
const showEditUserDialog = ref(false)
const showResetPasswordDialog = ref(false)
const showDeleteConfirmDialog = ref(false)
const selectedUser = ref(null)
const editUserForm = ref(null)
const editFormValid = ref(false)
const editSaving = ref(false)
const resetLoading = ref(false)
const deleteLoading = ref(false)

// قائمة القوائم الرئيسية
const mainMenuItems = ref([
  { title: 'الرئيسية', icon: 'mdi-view-dashboard', to: '/', active: false },
  { title: 'المدينون', icon: 'mdi-account-group', to: '/debtors', active: false },
  { title: 'المشاريع', icon: 'mdi-folder-multiple', to: '/projects', active: false },
  { title: 'المهندسين', icon: 'mdi-account-hard-hat', to: '/engineers', active: false },
  { title: 'المصاريف الإدارية', icon: 'mdi-cash-multiple', to: '/administrative-expenses', active: false },
  { title: 'المستخدمين', icon: 'mdi-account-multiple', to: '/users', active: true }
])

// المشاريع
const projects = ref([
  { name: 'تخطيط الحدث', color: 'purple' },
  { name: 'خطة الإفطار', color: 'green' }
])

// خيارات الفلاتر
const roles = ref([
  'مدير النظام',
  'مدير المشاريع',
  'مهندس',
  'محاسب',
  'موظف إداري',
  'مراجع',
  'مستخدم عادي'
])

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

// بيانات المستخدمين
const users = ref([
  {
    id: 1,
    name: 'أحمد محمد العلي',
    email: 'ahmed@example.com',
    phone: '+966501234567',
    role: 'admin',
    department: 'تقنية المعلومات',
    status: 'active',
    lastLogin: '2024-01-25T10:30:00',
    avatar: 'https://randomuser.me/api/portraits/men/1.jpg'
  },
  {
    id: 2,
    name: 'فاطمة السعد',
    email: 'fatima@example.com',
    phone: '+966507654321',
    role: 'project_manager',
    department: 'المشاريع',
    status: 'active',
    lastLogin: '2024-01-25T09:15:00',
    avatar: 'https://randomuser.me/api/portraits/women/1.jpg'
  },
  {
    id: 3,
    name: 'محمد عبدالله',
    email: 'mohammed@example.com',
    phone: '+966509876543',
    role: 'engineer',
    department: 'تقنية المعلومات',
    status: 'inactive',
    lastLogin: '2024-01-20T14:45:00',
    avatar: 'https://randomuser.me/api/portraits/men/2.jpg'
  },
  {
    id: 4,
    name: 'نورا أحمد',
    email: 'nora@example.com',
    phone: '+966501112233',
    role: 'accountant',
    department: 'المحاسبة',
    status: 'active',
    lastLogin: '2024-01-25T11:20:00',
    avatar: 'https://randomuser.me/api/portraits/women/2.jpg'
  },
  {
    id: 5,
    name: 'خالد السالم',
    email: 'khalid@example.com',
    phone: '+966504445566',
    role: 'admin',
    department: 'الإدارة',
    status: 'active',
    lastLogin: '2024-01-25T08:30:00',
    avatar: 'https://randomuser.me/api/portraits/men/3.jpg'
  },
  {
    id: 6,
    name: 'سارة محمد',
    email: 'sara@example.com',
    phone: '+966505556677',
    role: 'employee',
    department: 'الموارد البشرية',
    status: 'pending',
    lastLogin: null,
    avatar: 'https://randomuser.me/api/portraits/women/3.jpg'
  },
  {
    id: 7,
    name: 'عبدالرحمن علي',
    email: 'abdulrahman@example.com',
    phone: '+966506667788',
    role: 'engineer',
    department: 'الصيانة',
    status: 'active',
    lastLogin: '2024-01-24T16:10:00',
    avatar: 'https://randomuser.me/api/portraits/men/4.jpg'
  },
  {
    id: 8,
    name: 'مريم حسن',
    email: 'mariam@example.com',
    phone: '+966507778899',
    role: 'reviewer',
    department: 'التسويق',
    status: 'active',
    lastLogin: '2024-01-25T12:45:00',
    avatar: 'https://randomuser.me/api/portraits/women/4.jpg'
  }
])

// بيانات المستخدم الجديد
const newUser = ref({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  role: '',
  department: '',
  status: 'active',
  password: '',
  notes: '',
  avatar: ''
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

const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' },
  { title: 'معلق', value: 'pending' },
  { title: 'محظور', value: 'banned' }
]

// قواعد التحقق من صحة البيانات
const nameRules = [
  v => !!v || 'الاسم مطلوب',
  v => (v && v.length >= 2) || 'الاسم يجب أن يكون على الأقل حرفين',
  v => (v && v.length <= 50) || 'الاسم يجب أن يكون أقل من 50 حرف'
]

const emailRules = [
  v => !!v || 'البريد الإلكتروني مطلوب',
  v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح',
  v => !users.value.some(user => user.email === v) || 'البريد الإلكتروني مستخدم بالفعل'
]

const passwordRules = [
  v => !!v || 'كلمة المرور مطلوبة',
  v => (v && v.length >= 6) || 'كلمة المرور يجب أن تكون على الأقل 6 أحرف',
  v => (v && v.length <= 20) || 'كلمة المرور يجب أن تكون أقل من 20 حرف'
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

// الدوال المساعدة
const formatDate = (date) => {
  if (!date) return 'لم يسجل دخول'
  return new Date(date).toLocaleDateString('ar-SA', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
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

const deleteUser = (user) => {
  selectedUser.value = { ...user }
  showDeleteConfirmDialog.value = true
}

// دوال إدارة نافذة إضافة المستخدم
const closeAddUserDialog = () => {
  showAddUserDialog.value = false
  resetForm()
}

const resetForm = () => {
  newUser.value = {
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    role: '',
    department: '',
    status: 'active',
    password: '',
    notes: '',
    avatar: ''
  }
  if (addUserForm.value) {
    addUserForm.value.resetValidation()
  }
  formValid.value = false
  showPassword.value = false
}

const saveNewUser = async () => {
  if (!addUserForm.value.validate()) {
    return
  }

  saving.value = true

  try {
    // محاكاة عملية الحفظ
    await new Promise(resolve => setTimeout(resolve, 1500))

    // إنشاء المستخدم الجديد
    const user = {
      id: users.value.length + 1,
      name: `${newUser.value.firstName} ${newUser.value.lastName}`,
      email: newUser.value.email,
      phone: newUser.value.phone || 'غير محدد',
      role: newUser.value.role,
      department: newUser.value.department,
      status: newUser.value.status,
      lastLogin: null,
      avatar: newUser.value.avatar || `https://ui-avatars.com/api/?name=${encodeURIComponent(newUser.value.firstName + '+' + newUser.value.lastName)}&background=667eea&color=fff&size=128`,
      notes: newUser.value.notes,
      createdAt: new Date().toISOString()
    }

    // إضافة المستخدم إلى القائمة
    users.value.unshift(user)

    // إغلاق النافذة وإعادة تعيين النموذج
    closeAddUserDialog()

    // عرض رسالة نجاح (يمكن استخدام مكتبة toast)

  } catch (error) {
  } finally {
    saving.value = false
  }
}

// دوال إدارة نوافذ الإجراءات
const closeViewUserDialog = () => {
  showViewUserDialog.value = false
  selectedUser.value = null
}

const closeEditUserDialog = () => {
  showEditUserDialog.value = false
  selectedUser.value = null
  editFormValid.value = false
}

const saveEditUser = async () => {
  if (!editUserForm.value.validate()) {
    return
  }

  editSaving.value = true

  try {
    // محاكاة عملية الحفظ
    await new Promise(resolve => setTimeout(resolve, 1000))

    // العثور على المستخدم وتحديثه
    const index = users.value.findIndex(u => u.id === selectedUser.value.id)
    if (index !== -1) {
      users.value[index] = { ...selectedUser.value }
    }

    closeEditUserDialog()

  } catch (error) {
  } finally {
    editSaving.value = false
  }
}

const closeResetPasswordDialog = () => {
  showResetPasswordDialog.value = false
  selectedUser.value = null
}

const confirmResetPassword = async () => {
  resetLoading.value = true

  try {
    // محاكاة عملية إعادة تعيين كلمة المرور
    await new Promise(resolve => setTimeout(resolve, 1500))

    closeResetPasswordDialog()

  } catch (error) {
  } finally {
    resetLoading.value = false
  }
}

const closeDeleteConfirmDialog = () => {
  showDeleteConfirmDialog.value = false
  selectedUser.value = null
}

const confirmDeleteUser = async () => {
  deleteLoading.value = true

  try {
    // محاكاة عملية الحذف
    await new Promise(resolve => setTimeout(resolve, 1000))

    // حذف المستخدم من القائمة
    const index = users.value.findIndex(u => u.id === selectedUser.value.id)
    if (index !== -1) {
      users.value.splice(index, 1)
    }

    closeDeleteConfirmDialog()

  } catch (error) {
  } finally {
    deleteLoading.value = false
  }
}
</script>


<style>
/* Import page styles */
@import './styles/users.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
