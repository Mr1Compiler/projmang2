<template>
  <v-dialog v-model="dialogModel" max-width="700" class="edit-engineer-dialog">
    <v-card class="edit-dialog-card">
      <v-card-title class="edit-dialog-title">
        <div class="d-flex align-center">
          <v-avatar size="32" color="primary" variant="tonal" class="me-3">
            <v-icon>{{ isEditing ? 'mdi-pencil' : 'mdi-plus' }}</v-icon>
          </v-avatar>
          <h2 class="text-h5 font-weight-bold mb-0">
            {{ isEditing ? 'تعديل المهندس' : 'إضافة حساب جديد' }}
          </h2>
        </div>
      </v-card-title>

      <v-card-text>
        <v-form ref="formRef" v-model="formValid">
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="form.name"
                label="اسم المهندس"
                :rules="[v => !!v || 'اسم المهندس مطلوب']"
                required
                variant="outlined"
                density="comfortable"
                prepend-inner-icon="mdi-account"
                class="form-field-enhanced"
                style="font-size: var(--font-size-base-plus) !important; color: #000000 !important;"
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
                density="comfortable"
                prepend-inner-icon="mdi-lock"
                :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                @click:append-inner="showPassword = !showPassword"
                class="form-field-enhanced"
                style="font-size: var(--font-size-base-plus) !important; color: #000000 !important;"
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
                density="comfortable"
                prepend-inner-icon="mdi-email"
                class="form-field-enhanced"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.phone"
                label="رقم الهاتف"
                variant="outlined"
                density="comfortable"
                prepend-inner-icon="mdi-phone"
                class="form-field-enhanced"
              />
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="form.specialization"
                label="التخصص"
                variant="outlined"
                density="comfortable"
                prepend-inner-icon="mdi-briefcase"
                class="form-field-enhanced"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider />
      <v-card-actions class="dialog-actions pa-4">
        <v-spacer />
        <v-btn
          color="grey"
          variant="text"
          @click="closeDialog"
          size="large"
          class="cancel-btn"
        >
          <v-icon class="me-2">mdi-close</v-icon>
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          @click="saveEngineer"
          :disabled="!formValid"
          size="large"
          class="save-btn"
        >
          <v-icon class="me-2">{{ isEditing ? 'mdi-content-save' : 'mdi-plus' }}</v-icon>
          {{ isEditing ? 'تحديث' : 'إضافة' }}
        </v-btn>
      </v-card-actions>
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
.edit-dialog-card {
  border-radius: 20px !important;
  overflow: hidden !important;
}

.edit-dialog-title {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  color: white !important;
  padding: 20px 24px !important;
}

.edit-dialog-title h2 {
  color: white !important;
}

.form-field-enhanced {
  margin-bottom: 8px;
}

.dialog-actions {
  background: #f8fafc;
}

.cancel-btn {
  border-radius: 12px !important;
}

.save-btn {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
}
</style>
