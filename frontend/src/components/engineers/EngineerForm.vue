<template>
  <v-dialog v-model="dialogModel" max-width="700" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">{{ isEditing ? 'mdi-pencil' : 'mdi-account-plus' }}</v-icon>
            <span class="header-title">{{ isEditing ? 'تعديل المهندس' : 'إضافة حساب جديد' }}</span>
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

      <!-- Form Content -->
      <div class="dialog-body">
        <v-form ref="formRef" v-model="formValid">
          <div class="form-fields">
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="form.name"
                  label="اسم المهندس"
                  :rules="[v => !!v || 'اسم المهندس مطلوب']"
                  required
                  variant="outlined"
                  prepend-inner-icon="mdi-account"
                  class="form-field"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                  v-model="form.password"
                  label="كلمة المرور"
                  :type="showPassword ? 'text' : 'password'"
                  :rules="isEditing ? [] : [v => !!v || 'كلمة المرور مطلوبة']"
                  :required="!isEditing"
                  variant="outlined"
                  prepend-inner-icon="mdi-lock"
                  :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                  @click:append-inner="showPassword = !showPassword"
                  class="form-field"
                  :hint="isEditing ? 'اتركه فارغاً إذا لم ترد تغيير كلمة المرور' : ''"
                  persistent-hint
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.email"
                  label="البريد الإلكتروني"
                  type="email"
                  variant="outlined"
                  prepend-inner-icon="mdi-email"
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                  prepend-inner-icon="mdi-phone"
                  class="form-field"
                />
              </v-col>
              <v-col cols="12">
                <v-text-field
                  v-model="form.specialization"
                  label="التخصص"
                  variant="outlined"
                  prepend-inner-icon="mdi-briefcase"
                  class="form-field"
                />
              </v-col>
            </v-row>
          </div>
        </v-form>
      </div>

      <!-- Dialog Actions -->
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
          @click="saveEngineer"
          :disabled="!formValid"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-content-save</v-icon>
          {{ isEditing ? 'تحديث' : 'إضافة' }}
        </v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  engineer: {
    type: Object,
    default: null
  },
  isEditing: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'save', 'close'])

// Dialog model
const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// Form state
const formRef = ref(null)
const formValid = ref(false)
const showPassword = ref(false)
const form = ref({
  name: '',
  password: '',
  email: '',
  phone: '',
  specialization: ''
})

// Reset form to default values
const resetForm = () => {
  form.value = {
    name: '',
    password: '',
    email: '',
    phone: '',
    specialization: ''
  }
  showPassword.value = false
}

// Watch for engineer changes (when editing)
watch(() => props.engineer, (newEngineer) => {
  if (newEngineer) {
    form.value = {
      name: newEngineer.name || '',
      password: '',
      email: newEngineer.email || '',
      phone: newEngineer.phone || '',
      specialization: newEngineer.specialization || ''
    }
  } else {
    resetForm()
  }
}, { immediate: true })

// Close dialog
const closeDialog = () => {
  dialogModel.value = false
  resetForm()
  emit('close')
}

// Save engineer
const saveEngineer = () => {
  if (formValid.value) {
    emit('save', { ...form.value })
    closeDialog()
  }
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
