<template>
  <!-- Reset Password Dialog -->
  <v-dialog :model-value="resetPasswordDialog" @update:model-value="$emit('update:resetPasswordDialog', $event)" max-width="500px" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-warning) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-key-change</v-icon>
            <span class="header-title">إعادة تعيين كلمة المرور</span>
          </div>
          <v-btn
            icon="mdi-close"
            variant="text"
            size="small"
            color="white"
            @click="$emit('update:resetPasswordDialog', false)"
            class="close-btn"
          />
        </div>
      </div>

      <!-- Dialog Body -->
      <div class="dialog-body" v-if="user">
        <div class="text-center mb-4">
          <v-avatar size="60">
            <v-img :src="user.avatar || defaultAvatar" />
          </v-avatar>
          <h4 class="mt-2">{{ user.name }}</h4>
          <p class="text-caption">{{ user.email }}</p>
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
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn
          color="grey"
          variant="text"
          @click="$emit('update:resetPasswordDialog', false)"
          class="cancel-btn"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="warning"
          variant="elevated"
          @click="$emit('confirm-reset')"
          :loading="resetLoading"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-key-change</v-icon>
          إعادة تعيين
        </v-btn>
      </div>
    </v-card>
  </v-dialog>

  <!-- Delete Confirmation Dialog -->
  <v-dialog :model-value="deleteDialog" @update:model-value="$emit('update:deleteDialog', $event)" max-width="500px" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-error) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-delete-alert</v-icon>
            <span class="header-title">تأكيد الحذف</span>
          </div>
          <v-btn
            icon="mdi-close"
            variant="text"
            size="small"
            color="white"
            @click="$emit('update:deleteDialog', false)"
            class="close-btn"
          />
        </div>
      </div>

      <!-- Dialog Body -->
      <div class="dialog-body" v-if="user">
        <div class="text-center mb-4">
          <v-avatar size="60">
            <v-img :src="user.avatar || defaultAvatar" />
          </v-avatar>
          <h4 class="mt-2">{{ user.name }}</h4>
          <p class="text-caption">{{ user.email }}</p>
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
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn
          color="grey"
          variant="text"
          @click="$emit('update:deleteDialog', false)"
          class="cancel-btn"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="error"
          variant="elevated"
          @click="$emit('confirm-delete')"
          :loading="deleteLoading"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-delete</v-icon>
          حذف نهائي
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
defineProps({
  resetPasswordDialog: {
    type: Boolean,
    default: false
  },
  deleteDialog: {
    type: Boolean,
    default: false
  },
  user: {
    type: Object,
    default: null
  },
  resetLoading: {
    type: Boolean,
    default: false
  },
  deleteLoading: {
    type: Boolean,
    default: false
  }
})

defineEmits([
  'update:resetPasswordDialog',
  'update:deleteDialog',
  'confirm-reset',
  'confirm-delete'
])

const defaultAvatar = 'https://ui-avatars.com/api/?name=User&background=667eea&color=fff&size=128'
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
