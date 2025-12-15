<template>
  <v-dialog v-model="dialogModel" max-width="600" scrollable persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-folder-plus</v-icon>
            <span class="header-title">{{ isEditing ? 'تعديل المشروع' : 'إضافة مشروع' }}</span>
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
            <!-- Project Name -->
            <div class="field-group">
              <label class="field-label">اسم المشروع</label>
              <v-text-field
                v-model="form.name"
                variant="outlined"
                density="comfortable"
                placeholder="أدخل اسم المشروع"
                :rules="[v => !!v || 'اسم المشروع مطلوب']"
                required
                class="custom-input"
                hide-details="auto"
              />
            </div>

            <!-- Project Type -->
            <div class="field-group">
              <label class="field-label">نوع المشروع</label>
              <v-text-field
                v-model="form.type"
                variant="outlined"
                density="comfortable"
                placeholder="أدخل نوع المشروع"
                class="custom-input"
                hide-details="auto"
              />
            </div>

            <!-- Initial Cost -->
            <div class="field-group">
              <label class="field-label">التكلفة المبدئية (د.ع)</label>
              <v-text-field
                v-model.number="form.initialCost"
                type="number"
                variant="outlined"
                density="comfortable"
                placeholder="0"
                :rules="[v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر']"
                required
                class="custom-input"
                hide-details="auto"
              />
            </div>

            <!-- Duration -->
            <div class="field-group">
              <label class="field-label">مدة المشروع (أيام)</label>
              <v-text-field
                v-model.number="form.duration"
                type="number"
                variant="outlined"
                density="comfortable"
                placeholder="0"
                class="custom-input"
                hide-details="auto"
              />
            </div>

            <!-- Total Cost (calculated) -->
            <div class="field-group">
              <label class="field-label">التكلفة الإجمالية (د.ع)</label>
              <v-text-field
                :value="form.initialCost * (form.duration || 1)"
                variant="outlined"
                density="comfortable"
                readonly
                class="custom-input readonly-field"
                hide-details="auto"
              />
            </div>

            <!-- Progress -->
            <div class="field-group">
              <label class="field-label">نسبة الإنجاز (%)</label>
              <v-text-field
                v-model.number="form.progress"
                type="number"
                variant="outlined"
                density="comfortable"
                min="0"
                max="100"
                class="custom-input"
                hide-details="auto"
                :rules="[
                  v => v >= 0 || 'النسبة يجب أن تكون أكبر من أو تساوي 0',
                  v => v <= 100 || 'النسبة يجب أن تكون أقل من أو تساوي 100'
                ]"
              />
            </div>

            <!-- Phone Number -->
            <div class="field-group">
              <label class="field-label">رقم الهاتف</label>
              <v-text-field
                v-model="form.phone"
                variant="outlined"
                density="comfortable"
                placeholder="07XX XXX XXXX"
                class="custom-input"
                hide-details="auto"
              />
            </div>

            <!-- Location -->
            <div class="field-group">
              <label class="field-label">مكان المشروع</label>
              <v-text-field
                v-model="form.location"
                variant="outlined"
                density="comfortable"
                placeholder="أدخل مكان المشروع"
                :rules="[v => !!v || 'مكان المشروع مطلوب']"
                required
                class="custom-input"
                hide-details="auto"
              />
            </div>

            <!-- Notes -->
            <div class="field-group">
              <label class="field-label">ملاحظات</label>
              <v-textarea
                v-model="form.description"
                variant="outlined"
                density="comfortable"
                placeholder="أدخل ملاحظات إضافية"
                rows="3"
                class="custom-input"
                hide-details="auto"
              />
            </div>
          </div>
        </v-form>
      </div>

      <!-- Footer Actions -->
      <div class="dialog-footer">
        <v-btn
          variant="outlined"
          @click="closeDialog"
          class="cancel-btn"
          size="large"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          @click="saveProject"
          :disabled="!formValid"
          class="save-btn"
          size="large"
        >
          {{ isEditing ? 'تحديث المشروع' : 'إضافة المشروع' }}
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
  project: {
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
const form = ref({
  name: '',
  type: '',
  location: '',
  initialCost: 0,
  duration: 0,
  startDate: '',
  criticalCost: 0,
  phone: '',
  user: '',
  status: 'pending',
  priority: 'medium',
  progress: 0,
  description: '',
  category: ''
})

// Watch for project changes (when editing)
watch(() => props.project, (newProject) => {
  if (newProject) {
    form.value = { ...newProject }
  } else {
    resetForm()
  }
}, { immediate: true })

// Reset form to default values
const resetForm = () => {
  form.value = {
    name: '',
    type: '',
    location: '',
    initialCost: 0,
    duration: 0,
    startDate: '',
    criticalCost: 0,
    phone: '',
    user: '',
    status: 'pending',
    priority: 'medium',
    progress: 0,
    description: '',
    category: ''
  }
}

// Close dialog
const closeDialog = () => {
  dialogModel.value = false
  resetForm()
  emit('close')
}

// Save project
const saveProject = () => {
  if (formValid.value) {
    emit('save', { ...form.value })
    closeDialog()
  }
}
</script>

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for ProjectForm */

/* Dialog container */
.image-style-dialog {
  border-radius: 20px !important;
  overflow: hidden !important;
}

/* Header content layout */
.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.header-icon {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  padding: 0.5rem;
}

.close-btn {
  opacity: 0.9;
}

.close-btn:hover {
  opacity: 1;
  background: rgba(255, 255, 255, 0.1);
}

/* Form layout */
.form-fields {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.readonly-field {
  background: #e5e7eb !important;
}
</style>
