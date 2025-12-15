<template>
  <v-dialog v-model="dialogModel" max-width="450" persistent>
    <v-card class="delete-dialog">
      <!-- Header -->
      <div class="dialog-header">
        <v-icon size="48" color="white">mdi-alert-circle</v-icon>
      </div>

      <!-- Content -->
      <v-card-text class="text-center pa-6">
        <h3 class="text-h5 font-weight-bold mb-4">{{ title }}</h3>
        <p class="text-body-1 text-grey-darken-1 mb-2">{{ message }}</p>
        <p v-if="itemName" class="text-body-1 font-weight-bold text-primary">
          "{{ itemName }}"
        </p>
        <p class="text-caption text-error mt-4">
          <v-icon size="small" class="me-1">mdi-information</v-icon>
          هذا الإجراء لا يمكن التراجع عنه
        </p>
      </v-card-text>

      <!-- Actions -->
      <v-card-actions class="justify-center pa-4 pt-0">
        <v-btn
          variant="outlined"
          color="grey"
          @click="cancel"
          class="cancel-btn me-3"
          min-width="120"
        >
          إلغاء
        </v-btn>
        <v-btn
          variant="elevated"
          color="error"
          @click="confirm"
          :loading="loading"
          class="delete-btn"
          min-width="120"
        >
          <v-icon class="me-2">mdi-delete</v-icon>
          حذف
        </v-btn>
      </v-card-actions>
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

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for DeleteConfirmDialog */
/* All shared styles (.delete-dialog, .dialog-header, .cancel-btn, .delete-btn)
   are now in projects.css */
</style>
