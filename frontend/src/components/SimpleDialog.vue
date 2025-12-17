<template>
  <v-dialog v-model="isOpen" max-width="600" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">{{ icon }}</v-icon>
            <span class="header-title">{{ title }}</span>
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

      <!-- Content -->
      <div class="dialog-body">
        <slot />
      </div>

      <!-- Footer -->
      <div class="dialog-actions">
        <v-btn
          color="grey"
          variant="text"
          @click="closeDialog"
          class="cancel-btn"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          @click="saveAction"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-content-save</v-icon>
          {{ saveText }}
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  open: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: 'إضافة عنصر'
  },
  saveText: {
    type: String,
    default: 'حفظ'
  },
  icon: {
    type: String,
    default: 'mdi-plus'
  }
})

const emit = defineEmits(['close', 'save'])

const isOpen = computed({
  get: () => props.open,
  set: (value) => {
    if (!value) {
      emit('close')
    }
  }
})

const closeDialog = () => {
  emit('close')
}

const saveAction = () => {
  emit('save')
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
