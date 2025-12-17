<template>
  <v-dialog v-model="dialogModel" max-width="450" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-error) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-alert-circle</v-icon>
            <span class="header-title">{{ title }}</span>
          </div>
          <v-btn
            icon="mdi-close"
            variant="text"
            size="small"
            color="white"
            @click="cancel"
            class="close-btn"
          />
        </div>
      </div>

      <!-- Dialog Body -->
      <div class="dialog-body text-center">
        <p class="text-body-1 text-grey-darken-1 mb-2">{{ message }}</p>
        <p v-if="itemName" class="text-body-1 font-weight-bold text-primary">
          "{{ itemName }}"
        </p>
        <v-alert
          type="error"
          variant="tonal"
          class="mt-4"
        >
          <v-icon size="small" class="me-1">mdi-information</v-icon>
          هذا الإجراء لا يمكن التراجع عنه
        </v-alert>
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn
          color="grey"
          variant="text"
          @click="cancel"
          class="cancel-btn"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="error"
          variant="elevated"
          @click="confirm"
          :loading="loading"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-delete</v-icon>
          حذف
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: 'تأكيد الحذف'
  },
  message: {
    type: String,
    default: 'هل أنت متأكد من حذف هذا العنصر؟'
  },
  itemName: {
    type: String,
    default: ''
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'confirm', 'cancel'])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const confirm = () => {
  emit('confirm')
}

const cancel = () => {
  dialogModel.value = false
  emit('cancel')
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
