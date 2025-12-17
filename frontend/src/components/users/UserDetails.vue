<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="600px" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-account-details</v-icon>
            <span class="header-title">تفاصيل المستخدم</span>
          </div>
          <v-btn
            icon="mdi-close"
            variant="text"
            size="small"
            color="white"
            @click="$emit('update:modelValue', false)"
            class="close-btn"
          />
        </div>
      </div>

      <!-- Dialog Body -->
      <div v-if="user" class="dialog-body">
        <v-row>
          <v-col cols="12" class="text-center mb-4">
            <v-avatar size="100">
              <v-img :src="user.avatar || defaultAvatar" />
            </v-avatar>
            <h3 class="mt-3">{{ user.name }}</h3>
            <v-chip
              :color="getStatusColor(user.status)"
              size="small"
              class="mt-2"
            >
              {{ getStatusText(user.status) }}
            </v-chip>
          </v-col>

          <v-col cols="12" md="6">
            <v-list density="compact">
              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="primary">mdi-email</v-icon>
                </template>
                <v-list-item-title>البريد الإلكتروني</v-list-item-title>
                <v-list-item-subtitle>{{ user.email }}</v-list-item-subtitle>
              </v-list-item>

              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="success">mdi-phone</v-icon>
                </template>
                <v-list-item-title>رقم الهاتف</v-list-item-title>
                <v-list-item-subtitle>{{ user.phone || 'غير محدد' }}</v-list-item-subtitle>
              </v-list-item>

              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="warning">mdi-account-tie</v-icon>
                </template>
                <v-list-item-title>الدور</v-list-item-title>
                <v-list-item-subtitle>{{ getRoleText(user.role) }}</v-list-item-subtitle>
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
                <v-list-item-subtitle>{{ user.department || 'غير محدد' }}</v-list-item-subtitle>
              </v-list-item>

              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="purple">mdi-clock-outline</v-icon>
                </template>
                <v-list-item-title>آخر دخول</v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(user.lastLogin) }}</v-list-item-subtitle>
              </v-list-item>

              <v-list-item>
                <template v-slot:prepend>
                  <v-icon color="teal">mdi-calendar-plus</v-icon>
                </template>
                <v-list-item-title>تاريخ الإنشاء</v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(user.createdAt) }}</v-list-item-subtitle>
              </v-list-item>
            </v-list>
          </v-col>
        </v-row>
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn
          color="primary"
          variant="elevated"
          @click="$emit('update:modelValue', false)"
          class="save-btn"
        >
          إغلاق
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { formatDateTime } from '@/utils/formatters'

defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  user: {
    type: Object,
    default: null
  }
})

defineEmits(['update:modelValue'])

const defaultAvatar = 'https://ui-avatars.com/api/?name=User&background=667eea&color=fff&size=128'

// Use centralized formatter with custom fallback
const formatDate = (date) => {
  if (!date) return 'لم يسجل دخول'
  return formatDateTime(date)
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
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
