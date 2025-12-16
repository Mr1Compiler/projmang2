<template>
  <v-dialog v-model="dialogModel" max-width="1000" scrollable persistent>
    <v-card class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-account-plus</v-icon>
            <span class="header-title">{{ isEditing ? 'تعديل بيانات الموظف' : 'إضافة موظف جديد' }}</span>
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
          <v-tabs v-model="formTab" bg-color="transparent" class="form-tabs">
            <v-tab value="personal">المعلومات الشخصية</v-tab>
            <v-tab value="work">المعلومات الوظيفية</v-tab>
            <v-tab value="contact">معلومات الاتصال</v-tab>
            <v-tab value="additional">معلومات إضافية</v-tab>
          </v-tabs>

          <v-window v-model="formTab">
            <!-- Personal Information Tab -->
            <v-window-item value="personal">
              <div class="form-fields mt-4">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.name"
                      label="اسم الموظف الكامل"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم الموظف مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.idNumber"
                      label="رقم الهوية الوطنية"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="form.birthDate"
                      label="تاريخ الميلاد"
                      variant="outlined"
                      type="date"
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="form.gender"
                      :items="genderOptions"
                      label="الجنس"
                      variant="outlined"
                      class="form-field black-list"
                    />
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="form.maritalStatus"
                      :items="maritalStatusOptions"
                      label="الحالة الاجتماعية"
                      variant="outlined"
                      class="form-field black-list"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.nationality"
                      label="الجنسية"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.address"
                      label="العنوان"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>
              </div>
            </v-window-item>

            <!-- Work Information Tab -->
            <v-window-item value="work">
              <div class="form-fields mt-4">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="form.department"
                      :items="departments"
                      label="القسم"
                      variant="outlined"
                      :rules="[v => !!v || 'القسم مطلوب']"
                      required
                      class="form-field black-list"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.position"
                      label="المنصب"
                      variant="outlined"
                      :rules="[v => !!v || 'المنصب مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.hireDate"
                      label="تاريخ التوظيف"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ التوظيف مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.contractType"
                      label="نوع العقد"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.salary"
                      label="الراتب الأساسي (د.ع)"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'الراتب مطلوب', v => v > 0 || 'الراتب يجب أن يكون أكبر من صفر']"
                      required
                      class="form-field"
                    />
                  </v-col>
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
                </v-row>
              </div>
            </v-window-item>

            <!-- Contact Information Tab -->
            <v-window-item value="contact">
              <div class="form-fields mt-4">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.phone"
                      label="رقم الهاتف"
                      variant="outlined"
                      :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.phone2"
                      label="رقم هاتف إضافي"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.email"
                      label="البريد الإلكتروني"
                      variant="outlined"
                      type="email"
                      :rules="[v => !!v || 'البريد الإلكتروني مطلوب', v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.emergencyContact"
                      label="جهة الاتصال في حالة الطوارئ"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.emergencyPhone"
                      label="رقم هاتف الطوارئ"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                </v-row>
              </div>
            </v-window-item>

            <!-- Additional Information Tab -->
            <v-window-item value="additional">
              <div class="form-fields mt-4">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.education"
                      label="المؤهل العلمي"
                      variant="outlined"
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="form.experience"
                      label="سنوات الخبرة"
                      variant="outlined"
                      type="number"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12">
                    <v-textarea
                      v-model="form.notes"
                      label="ملاحظات"
                      variant="outlined"
                      rows="4"
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12">
                    <v-card class="fingerprint-card" elevation="2">
                      <v-card-title class="fingerprint-title">
                        <v-icon class="me-2">mdi-fingerprint</v-icon>
                        بصمة الموظف
                      </v-card-title>
                      <v-card-text>
                        <div class="fingerprint-status">
                          <v-chip
                            :color="form.fingerprint ? 'primary' : 'info'"
                            size="large"
                            class="mb-3 fingerprint-chip"
                            :class="form.fingerprint ? 'fingerprint-registered' : 'fingerprint-pending'"
                          >
                            <v-icon start>{{ form.fingerprint ? 'mdi-check-circle' : 'mdi-alert-circle' }}</v-icon>
                            {{ form.fingerprint ? 'تم تسجيل البصمة' : 'لم يتم تسجيل البصمة' }}
                          </v-chip>
                        </div>
                        <v-btn
                          color="primary"
                          variant="elevated"
                          @click="$emit('register-fingerprint')"
                          block
                          :disabled="loadingFingerprint"
                          class="fingerprint-btn"
                        >
                          <v-icon class="me-2">{{ loadingFingerprint ? 'mdi-loading mdi-spin' : 'mdi-fingerprint' }}</v-icon>
                          {{ form.fingerprint ? 'تحديث البصمة' : 'تسجيل البصمة' }}
                        </v-btn>
                        <p class="fingerprint-hint mt-3 text-caption text-grey">
                          اضغط على الزر واتبع التعليمات لتسجيل بصمة الموظف
                        </p>
                      </v-card-text>
                    </v-card>
                  </v-col>
                </v-row>
              </div>
            </v-window-item>
          </v-window>
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
          color="primary"
          variant="elevated"
          @click="saveEmployee"
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
  employee: {
    type: Object,
    default: null
  },
  isEditing: {
    type: Boolean,
    default: false
  },
  loadingFingerprint: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'save', 'close', 'register-fingerprint'])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const formRef = ref(null)
const formValid = ref(false)
const formTab = ref('personal')

const form = ref({
  name: '',
  idNumber: '',
  birthDate: '',
  gender: '',
  maritalStatus: '',
  nationality: '',
  address: '',
  phone: '',
  phone2: '',
  email: '',
  emergencyContact: '',
  emergencyPhone: '',
  department: '',
  position: '',
  hireDate: '',
  contractType: '',
  salary: 0,
  status: 'نشط',
  education: '',
  experience: '',
  notes: '',
  fingerprint: null,
  fingerprintDate: null
})

// Options
const statusOptions = [
  { title: 'نشط', value: 'نشط' },
  { title: 'في إجازة', value: 'في إجازة' },
  { title: 'معطل', value: 'معطل' }
]

const genderOptions = [
  { title: 'ذكر', value: 'ذكر' },
  { title: 'أنثى', value: 'أنثى' }
]

const maritalStatusOptions = [
  { title: 'عزباء', value: 'عزباء' },
  { title: 'متزوج', value: 'متزوج' },
  { title: 'مطلق', value: 'مطلق' },
  { title: 'أرمل', value: 'أرمل' }
]

const departments = [
  { title: 'المبيعات', value: 'المبيعات' },
  { title: 'المحاسبة', value: 'المحاسبة' },
  { title: 'الموارد البشرية', value: 'الموارد البشرية' },
  { title: 'التسويق', value: 'التسويق' },
  { title: 'التطوير', value: 'التطوير' },
  { title: 'الدعم الفني', value: 'الدعم الفني' }
]

const resetForm = () => {
  form.value = {
    name: '',
    idNumber: '',
    birthDate: '',
    gender: '',
    maritalStatus: '',
    nationality: '',
    address: '',
    phone: '',
    phone2: '',
    email: '',
    emergencyContact: '',
    emergencyPhone: '',
    department: '',
    position: '',
    hireDate: '',
    contractType: '',
    salary: 0,
    status: 'نشط',
    education: '',
    experience: '',
    notes: '',
    fingerprint: null,
    fingerprintDate: null
  }
  formTab.value = 'personal'
}

watch(() => props.employee, (newEmployee) => {
  if (newEmployee) {
    form.value = { ...newEmployee }
  } else {
    resetForm()
  }
  formTab.value = 'personal'
}, { immediate: true })

const closeDialog = () => {
  dialogModel.value = false
  resetForm()
  emit('close')
}

const saveEmployee = () => {
  if (formValid.value) {
    emit('save', { ...form.value })
    closeDialog()
  }
}
</script>

<style scoped>
@import './styles/human-resources.css';

.dialog-header {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%);
  padding: 16px 24px;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-title {
  color: white;
  font-weight: 600;
  font-size: var(--font-size-base-plus);
}

.dialog-body {
  padding: 24px;
  max-height: 60vh;
  overflow-y: auto;
}

.dialog-actions {
  padding: 16px 24px;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  border-top: 1px solid #e5e7eb;
}

.fingerprint-card {
  border-radius: 12px;
}

.fingerprint-title {
  background: #f8fafc;
  font-size: var(--font-size-base);
}

.fingerprint-chip {
  width: 100%;
  justify-content: center;
}

.fingerprint-btn {
  border-radius: 8px;
}
</style>
