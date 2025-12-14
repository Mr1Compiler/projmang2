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
  console.log('تم تطبيق الفلاتر')
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
    console.log('تم إضافة المستخدم بنجاح:', user)

  } catch (error) {
    console.error('خطأ في إضافة المستخدم:', error)
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
    console.log('تم تحديث المستخدم بنجاح')

  } catch (error) {
    console.error('خطأ في تحديث المستخدم:', error)
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
    console.log('تم إعادة تعيين كلمة المرور بنجاح')

  } catch (error) {
    console.error('خطأ في إعادة تعيين كلمة المرور:', error)
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
    console.log('تم حذف المستخدم بنجاح')

  } catch (error) {
    console.error('خطأ في حذف المستخدم:', error)
  } finally {
    deleteLoading.value = false
  }
}
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #2e3cbc;
  min-height: 100vh;
  padding: 20px;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
}

/* العنوان المحسن - نفس تصميم الصفحات الأخرى */
.page-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  backdrop-filter: blur(20px);
  border-radius: 16px;
  padding: 20px 30px;
  margin-bottom: 20px;
  text-align: center !important;
  border: 2px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.3);
  animation: gradient-animation 8s ease infinite;
  background-size: 200% 200%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: white;
}

.header-top-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 10px;
  width: 100%;
  text-align: center;
}

.glass-effect {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.3) 0%, rgba(118, 75, 162, 0.3) 100%) !important;
  backdrop-filter: blur(20px) !important;
  border: 1px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.2) !important;
}

.gradient-animation {
  position: relative;
  overflow: hidden;
}

.gradient-animation::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  animation: shimmer 3s infinite;
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

.page-icon {
  font-size: 2rem;
  display: inline-block;
  animation: iconFloat 3s ease-in-out infinite;
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes gradient-animation {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.gradient-animation {
  background-size: 200% 200%;
  animation: gradient-animation 3s ease infinite;
}

.star-twinkle {
  animation: star-twinkle 2s ease-in-out infinite;
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

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
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

.stats-container {
  padding: 0 2rem;
  margin-top: 1.5rem;
}

/* البطاقات المحسنة */
.v-card {
  background: rgba(83, 69, 205, 0.95) !important;
  color: #e2c0cf !important;
  border: 1px solid rgba(224, 12, 118, 0.1) !important;
  box-shadow: 0 8px 32px rgba(79, 70, 229, 0.15) !important;
  backdrop-filter: blur(10px);
  border-radius: 16px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

.v-card:hover {
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
  background: linear-gradient(135deg, rgba(71, 85, 105, 0.9) 0%, rgba(51, 65, 85, 0.8) 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.95rem !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.3) !important;
  padding: 1.5rem 1.2rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
  position: relative;
  overflow: hidden;
  backdrop-filter: blur(10px);
  letter-spacing: 0.025em;
  text-transform: uppercase;
  font-size: 0.875rem;
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
  padding: 1.5rem 1.2rem !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: rgba(248, 250, 252, 0.9);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 2;
  font-weight: 500;
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
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.1) 0%, rgba(147, 51, 234, 0.05) 100%) !important;
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
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

.view-user-dialog .dialog-header,
.edit-user-dialog .dialog-header,
.reset-password-dialog .dialog-header,
.delete-confirm-dialog .dialog-header {
  background: linear-gradient(135deg, rgba(71, 85, 105, 0.9) 0%, rgba(51, 65, 85, 0.8) 100%) !important;
  color: #ffffff !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 1.5rem 2rem !important;
}

.view-user-dialog .v-list-item-title,
.view-user-dialog .v-list-item-subtitle {
  color: #1a1a1a !important;
}

/* تنسيق شامل لنافذة تعديل المستخدم */
.edit-user-dialog .v-field,
.edit-user-dialog .v-field__input,
.edit-user-dialog .v-label,
.edit-user-dialog .v-field__outline,
.edit-user-dialog .v-field__append-inner,
.edit-user-dialog .v-field__prepend-inner {
  color: #1a1a1a !important;
  background: rgba(255, 255, 255, 0.95) !important;
  border: 1px solid rgba(148, 163, 184, 0.3) !important;
  border-radius: 12px !important;
}

.edit-user-dialog .v-field:hover,
.edit-user-dialog .v-field:focus-within {
  border-color: rgba(59, 130, 246, 0.5) !important;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.1) !important;
}

.edit-user-dialog .v-field__input {
  padding: 12px 16px !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

.edit-user-dialog .v-label {
  font-weight: 600 !important;
  color: #374151 !important;
  margin-bottom: 4px !important;
}

.edit-user-dialog .v-icon {
  color: #6b7280 !important;
}

.edit-user-dialog .v-select__selection {
  color: #1a1a1a !important;
  font-weight: 500 !important;
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
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 1.5rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
}

.edit-user-dialog .dialog-title .v-icon {
  color: #10b981 !important;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2)) !important;
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
  background: linear-gradient(135deg, rgba(44, 33, 100, 0.98) 0%, rgba(45, 54, 102, 0.95) 100%) !important;
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
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(67, 56, 202, 0.2);
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 255, 255, 0.2);
  padding: 8px;
}

.filters-card .v-card-text {
  background: transparent !important;
  padding: 16px !important;
}

.filters-card .v-select,
.filters-card .v-text-field {
  background: rgba(255, 255, 255, 0.15) !important;
  border-radius: 12px !important;
}

.filters-card .v-select .v-field__outline,
.filters-card .v-text-field .v-field__outline {
  border-color: rgba(255, 255, 255, 0.4) !important;
  border-width: 2px !important;
}

.filters-card .v-select .v-label,
.filters-card .v-text-field .v-label {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 700 !important;
}

.filters-card .v-select .v-field__input,
.filters-card .v-text-field .v-field__input {
  color: #ffffff !important;
  font-weight: 600 !important;
}

.filters-card .v-select .v-field__append-inner,
.filters-card .v-select .v-field__prepend-inner {
  color: #ffffff !important;
}

.filters-card .v-btn {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0.3) 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  border: 2px solid rgba(255, 255, 255, 0.4) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.filters-card .v-btn:hover {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.3) 0%, rgba(255, 255, 255, 0.4) 100%) !important;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.3) !important;
}

.users-table {
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.1);
  background: #ffffff !important;
  position: relative;
  transition: all 0.4s ease;
}

.users-table .v-card-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
  padding: 1.5rem !important;
  position: relative !important;
  color: #ffffff !important;
}

.users-table .v-card-title * {
  color: #ffffff !important;
}

.users-table .v-card-title .v-icon {
  color: #ffffff !important;
}

.users-table .v-card-title span {
  color: #ffffff !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

.users-table .v-card-title .v-btn {
  color: #ffffff !important;
}

.users-table .v-card-title .v-btn .v-btn__content {
  color: #ffffff !important;
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

  @keyframes fadeInUp {
    from {
      opacity: 0;
      transform: translateY(30px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
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
    background: linear-gradient(135deg, rgba(209, 155, 180, 0.95) 0%, rgba(67, 74, 178, 0.9) 100%);
    backdrop-filter: blur(15px);
    border: 2px solid rgba(210, 0, 171, 0.2);
    border-radius: 20px;
    overflow: hidden;
  }

  .dialog-header {
    background: linear-gradient(135deg, rgba(41, 5, 23, 0.8) 0%, rgba(219, 15, 172, 0.6) 100%);
    color: rgb(189, 111, 111);
    padding: 20px 24px;
    border-bottom: 2px solid rgba(210, 0, 171, 0.3);
  }

  .dialog-title {
    display: flex;
    align-items: center;
    font-weight: 700;
    font-size: 1.5rem;
  }

  .dialog-title h2 {
    color: white;
    margin: 0;
    font-weight: 700;
  }

  .close-btn {
    color: white !important;
  }

  .dialog-content {
    padding: 30px 24px;
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(255, 255, 255, 0.9) 100%);
  }

  .dialog-actions {
    padding: 20px 24px;
    background: linear-gradient(135deg, rgba(71, 72, 135, 0.8) 0%, rgba(76, 86, 194, 0.7) 100%);
    border-top: 2px solid rgba(210, 0, 171, 0.3);
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
    background: rgba(124, 93, 196, 0.9);
    border-radius: 12px;
    border: 1px solid rgba(219, 15, 172, 0.2);
    transition: all 0.3s ease;
  }

  .add-user-dialog .v-text-field .v-field:hover,
  .add-user-dialog .v-select .v-field:hover,
  .add-user-dialog .v-textarea .v-field:hover {
    border-color: rgba(219, 15, 172, 0.4);
    box-shadow: 0 4px 15px rgba(219, 15, 172, 0.1);
  }

  .add-user-dialog .v-text-field.v-field--focused .v-field,
  .add-user-dialog .v-select.v-field--focused .v-field,
  .add-user-dialog .v-textarea.v-field--focused .v-field {
    border-color: rgba(219, 15, 172, 0.6);
    box-shadow: 0 6px 20px rgba(219, 15, 172, 0.2);
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
@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translateY(30px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-10px);
  }
  60% {
    transform: translateY(-5px);
  }
}

@keyframes shimmer {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

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
/* تنسيقات إضافية قوية للتسميات - ضمان الوضوح */
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
.add-user-dialog .v-field--focused .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  background: white !important;
}

.add-user-dialog .v-text-field .v-label,
.add-user-dialog .v-select .v-label,
.add-user-dialog .v-textarea .v-label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}
</style>
