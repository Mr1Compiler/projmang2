<template>
  <v-dialog v-model="dialogModel" max-width="800" scrollable persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-currency-usd</v-icon>
            <span class="header-title">{{ isEditing ? 'تعديل المصروف الإداري' : 'إضافة مصروف إداري' }}</span>
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
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.projectName"
                  label="اسم المشروع"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم المشروع مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.cost"
                  label="التكلفة (د.ع)"
                  variant="outlined"
                  type="number"
                  :rules="[v => !!v || 'التكلفة مطلوبة', v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر']"
                  required
                  class="form-field"
                />
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.startDate"
                  label="تاريخ البداية"
                  variant="outlined"
                  type="date"
                  :rules="[v => !!v || 'تاريخ البداية مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.endDate"
                  label="تاريخ الانتهاء"
                  variant="outlined"
                  type="date"
                  :rules="[v => !!v || 'تاريخ الانتهاء مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.workLocation"
                  label="مكان العمل"
                  variant="outlined"
                  :rules="[v => !!v || 'مكان العمل مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="form.expenseType"
                  :items="expenseTypes"
                  label="نوع المصروف"
                  variant="outlined"
                  :rules="[v => !!v || 'نوع المصروف مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12" md="6">
                <v-select
                  v-model="form.status"
                  :items="statusOptions"
                  label="الحالة"
                  variant="outlined"
                  :rules="[v => !!v || 'الحالة مطلوبة']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <!-- Empty column for spacing -->
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12">
                <v-textarea
                  v-model="form.notes"
                  label="الملاحظات"
                  variant="outlined"
                  rows="3"
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
          class="action-btn"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="error"
          variant="elevated"
          @click="saveExpense"
          :disabled="!formValid"
          class="action-btn primary-btn"
        >
          <v-icon class="me-2">mdi-content-save</v-icon>
          {{ isEditing ? 'تحديث' : 'حفظ' }}
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
  expense: {
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
  projectName: '',
  startDate: '',
  endDate: '',
  cost: '',
  workLocation: '',
  expenseType: '',
  status: 'معلق',
  notes: ''
})

const expenseTypes = [
  'تطوير',
  'تحديث',
  'بناء',
  'أمن',
  'صيانة',
  'تدريب',
  'أخرى'
]

const statusOptions = [
  'معتمد',
  'معلق',
  'مرفوض',
  'مسودة'
]

// Reset form to default values
const resetForm = () => {
  form.value = {
    projectName: '',
    startDate: '',
    endDate: '',
    cost: '',
    workLocation: '',
    expenseType: '',
    status: 'معلق',
    notes: ''
  }
}

// Watch for expense changes (when editing)
watch(() => props.expense, (newExpense) => {
  if (newExpense) {
    form.value = { ...newExpense }
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

// Save expense
const saveExpense = () => {
  if (formValid.value) {
    emit('save', { ...form.value, cost: parseFloat(form.value.cost) })
    closeDialog()
  }
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */

/* Use red/error gradient for expenses header */
.dialog-header {
  background: var(--gradient-error-deep) !important;
}

/* Primary button uses error color for expenses */
.primary-btn {
  background: var(--gradient-error) !important;
  color: var(--text-white) !important;
}
</style>
