<template>
  <v-dialog v-model="dialogModel" max-width="600px" persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">{{ isEdit ? 'mdi-pencil' : 'mdi-account-plus' }}</v-icon>
            <span class="header-title">{{ isEdit ? 'تعديل بيانات المديون' : 'إضافة مديون جديد' }}</span>
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
                  v-model="form.name"
                  label="الاسم الكامل"
                  variant="outlined"
                  :rules="[v => !!v || 'الاسم مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.email"
                  label="البريد الإلكتروني"
                  type="email"
                  variant="outlined"
                  :rules="[v => !!v || 'البريد الإلكتروني مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.amount"
                  label="المبلغ المطلوب"
                  type="number"
                  variant="outlined"
                  :rules="[v => !!v || 'المبلغ مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="form.currency"
                  label="العملة"
                  :items="currencyOptions"
                  variant="outlined"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="form.dueDate"
                  label="تاريخ الاستحقاق"
                  type="date"
                  variant="outlined"
                  :rules="[v => !!v || 'تاريخ الاستحقاق مطلوب']"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="form.status"
                  label="الحالة"
                  :items="statusOptions"
                  variant="outlined"
                  required
                  class="form-field"
                />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  v-model="form.notes"
                  label="ملاحظات"
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
          class="cancel-btn"
        >
          إلغاء
        </v-btn>
        <v-btn
          color="primary"
          variant="elevated"
          :disabled="!formValid"
          @click="saveDebtor"
          class="save-btn"
        >
          <v-icon class="me-2">mdi-content-save</v-icon>
          {{ isEdit ? 'تحديث' : 'حفظ' }}
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
  debtor: {
    type: Object,
    default: null
  },
  isEdit: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'save', 'close'])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const formRef = ref(null)
const formValid = ref(false)
const form = ref({
  name: '',
  email: '',
  phone: '',
  amount: '',
  currency: 'IQD',
  dueDate: '',
  status: 'pending',
  notes: ''
})

const currencyOptions = ['IQD', 'USD', 'EUR', 'SAR']
const statusOptions = [
  { title: 'قيد الانتظار', value: 'pending' },
  { title: 'مدفوع', value: 'paid' },
  { title: 'متأخر', value: 'overdue' }
]

const resetForm = () => {
  form.value = {
    name: '',
    email: '',
    phone: '',
    amount: '',
    currency: 'IQD',
    dueDate: '',
    status: 'pending',
    notes: ''
  }
}

watch(() => props.debtor, (newDebtor) => {
  if (newDebtor) {
    form.value = { ...newDebtor }
  } else {
    resetForm()
  }
}, { immediate: true })

const closeDialog = () => {
  dialogModel.value = false
  resetForm()
  emit('close')
}

const saveDebtor = () => {
  if (formValid.value) {
    emit('save', { ...form.value, amount: parseFloat(form.value.amount) })
    closeDialog()
  }
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
</style>
