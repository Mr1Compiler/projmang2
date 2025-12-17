<template>
  <v-dialog v-model="dialogModel" max-width="500px" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-info) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-account-details</v-icon>
            <span class="header-title">تفاصيل المديون</span>
          </div>
          <v-btn
            icon="mdi-close"
            variant="text"
            size="small"
            color="white"
            @click="closeDialog"
            class="close-btn"
          />
        </div>
      </div>

      <!-- Dialog Body -->
      <div v-if="debtor" class="dialog-body">
        <v-row>
          <v-col cols="12">
            <div class="text-center mb-4">
              <v-avatar
                :color="getStatusColor(debtor.status)"
                size="64"
              >
                <span class="text-white text-h4 font-weight-bold">
                  {{ debtor.name?.charAt(0) }}
                </span>
              </v-avatar>
              <h3 class="text-h5 font-weight-bold mt-2">{{ debtor.name }}</h3>
            </div>
          </v-col>
          <v-col cols="6">
            <strong>البريد الإلكتروني:</strong>
            <p>{{ debtor.email }}</p>
          </v-col>
          <v-col cols="6">
            <strong>رقم الهاتف:</strong>
            <p>{{ debtor.phone || 'غير محدد' }}</p>
          </v-col>
          <v-col cols="6">
            <strong>المبلغ المطلوب:</strong>
            <p class="text-h6 font-weight-bold text-error">
              {{ formatCurrency(debtor.amount) }}
            </p>
          </v-col>
          <v-col cols="6">
            <strong>تاريخ الاستحقاق:</strong>
            <p>{{ formatDate(debtor.dueDate) }}</p>
          </v-col>
          <v-col cols="12">
            <strong>الحالة:</strong>
            <v-chip
              :color="getStatusColor(debtor.status)"
              :variant="debtor.status === 'paid' ? 'flat' : 'tonal'"
              class="mt-1"
            >
              {{ getStatusText(debtor.status) }}
            </v-chip>
          </v-col>
          <v-col cols="12" v-if="debtor.notes">
            <strong>ملاحظات:</strong>
            <p>{{ debtor.notes }}</p>
          </v-col>
        </v-row>
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn
          color="primary"
          variant="elevated"
          @click="closeDialog"
          class="save-btn"
        >
          إغلاق
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed } from 'vue'
import { formatCurrency, formatDate } from '@/utils/formatters'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  debtor: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['update:modelValue', 'close'])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const closeDialog = () => {
  dialogModel.value = false
  emit('close')
}

const getStatusColor = (status) => {
  const colors = {
    'overdue': 'error',
    'pending': 'warning',
    'paid': 'success'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    'overdue': 'متأخر',
    'pending': 'معلق',
    'paid': 'مدفوع'
  }
  return texts[status] || 'غير محدد'
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
