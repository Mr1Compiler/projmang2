<template>
  <!-- Add Leave Dialog -->
  <v-dialog v-model="leaveDialogModel" max-width="600" persistent>
    <v-card class="image-style-dialog">
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-calendar-plus</v-icon>
            <span class="header-title">إضافة إجازة</span>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeLeaveDialog" class="close-btn" />
        </div>
      </div>
      <div class="dialog-body">
        <v-form>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="leaveForm.startDate"
                label="تاريخ البداية"
                type="date"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="leaveForm.endDate"
                label="تاريخ النهاية"
                type="date"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-select
                v-model="leaveForm.type"
                :items="leaveTypeOptions"
                label="نوع الإجازة"
                variant="outlined"
                required
                class="form-field black-list"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="leaveForm.status"
                :items="leaveStatusOptions"
                label="الحالة"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-textarea
                v-model="leaveForm.reason"
                label="السبب"
                variant="outlined"
                rows="3"
                class="form-field"
              />
            </v-col>
          </v-row>
        </v-form>
      </div>
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeLeaveDialog">إلغاء</v-btn>
        <v-btn color="primary" variant="elevated" @click="saveLeave">حفظ</v-btn>
      </div>
    </v-card>
  </v-dialog>

  <!-- Add Attendance Dialog -->
  <v-dialog v-model="attendanceDialogModel" max-width="600" persistent>
    <v-card class="image-style-dialog">
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-clock-in</v-icon>
            <span class="header-title">تسجيل حضور</span>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeAttendanceDialog" class="close-btn" />
        </div>
      </div>
      <div class="dialog-body">
        <v-form>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="attendanceForm.date"
                label="التاريخ"
                type="date"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="attendanceForm.status"
                :items="attendanceStatusOptions"
                label="الحالة"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="attendanceForm.checkIn"
                label="وقت الدخول"
                type="time"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="attendanceForm.checkOut"
                label="وقت الخروج"
                type="time"
                variant="outlined"
                class="form-field"
              />
            </v-col>
          </v-row>
        </v-form>
      </div>
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeAttendanceDialog">إلغاء</v-btn>
        <v-btn color="primary" variant="elevated" @click="saveAttendance">حفظ</v-btn>
      </div>
    </v-card>
  </v-dialog>

  <!-- Add Evaluation Dialog -->
  <v-dialog v-model="evaluationDialogModel" max-width="600" persistent>
    <v-card class="image-style-dialog">
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-star-plus</v-icon>
            <span class="header-title">إضافة تقييم</span>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeEvaluationDialog" class="close-btn" />
        </div>
      </div>
      <div class="dialog-body">
        <v-form>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="evaluationForm.date"
                label="التاريخ"
                type="date"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="evaluationForm.evaluator"
                label="المقيّم"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-slider
                v-model="evaluationForm.rating"
                label="التقييم"
                min="1"
                max="5"
                step="1"
                thumb-label="always"
                class="form-field"
              >
                <template #append>
                  <v-text-field
                    v-model="evaluationForm.rating"
                    type="number"
                    style="width: 60px"
                    density="compact"
                    hide-details
                    variant="outlined"
                  />
                </template>
              </v-slider>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-textarea
                v-model="evaluationForm.comments"
                label="ملاحظات"
                variant="outlined"
                rows="3"
                class="form-field"
              />
            </v-col>
          </v-row>
        </v-form>
      </div>
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeEvaluationDialog">إلغاء</v-btn>
        <v-btn color="primary" variant="elevated" @click="saveEvaluation">حفظ</v-btn>
      </div>
    </v-card>
  </v-dialog>

  <!-- Add Skill Dialog -->
  <v-dialog v-model="skillDialogModel" max-width="500" persistent>
    <v-card class="image-style-dialog">
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-certificate</v-icon>
            <span class="header-title">إضافة مهارة</span>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeSkillDialog" class="close-btn" />
        </div>
      </div>
      <div class="dialog-body">
        <v-form>
          <v-text-field
            v-model="skillForm.name"
            label="اسم المهارة"
            variant="outlined"
            required
            class="form-field"
          />
        </v-form>
      </div>
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeSkillDialog">إلغاء</v-btn>
        <v-btn color="primary" variant="elevated" @click="saveSkill">حفظ</v-btn>
      </div>
    </v-card>
  </v-dialog>

  <!-- Add Certificate Dialog -->
  <v-dialog v-model="certificateDialogModel" max-width="600" persistent>
    <v-card class="image-style-dialog">
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-certificate</v-icon>
            <span class="header-title">إضافة شهادة</span>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeCertificateDialog" class="close-btn" />
        </div>
      </div>
      <div class="dialog-body">
        <v-form>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="certificateForm.name"
                label="اسم الشهادة"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="certificateForm.issuer"
                label="المصدر"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="certificateForm.date"
                label="تاريخ الحصول"
                type="date"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="certificateForm.expiryDate"
                label="تاريخ الانتهاء (اختياري)"
                type="date"
                variant="outlined"
                class="form-field"
              />
            </v-col>
          </v-row>
        </v-form>
      </div>
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeCertificateDialog">إلغاء</v-btn>
        <v-btn color="primary" variant="elevated" @click="saveCertificate">حفظ</v-btn>
      </div>
    </v-card>
  </v-dialog>

  <!-- Add Salary Dialog -->
  <v-dialog v-model="salaryDialogModel" max-width="600" persistent>
    <v-card class="image-style-dialog">
      <div class="dialog-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-cash-plus</v-icon>
            <span class="header-title">إضافة راتب</span>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeSalaryDialog" class="close-btn" />
        </div>
      </div>
      <div class="dialog-body">
        <v-form>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="salaryForm.month"
                label="الشهر (YYYY-MM)"
                variant="outlined"
                required
                placeholder="2024-01"
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="salaryForm.status"
                :items="salaryStatusOptions"
                label="الحالة"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="salaryForm.baseSalary"
                label="الراتب الأساسي (د.ع)"
                type="number"
                variant="outlined"
                required
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="salaryForm.bonuses"
                label="المكافآت (د.ع)"
                type="number"
                variant="outlined"
                class="form-field"
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="salaryForm.deductions"
                label="الخصومات (د.ع)"
                type="number"
                variant="outlined"
                class="form-field"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-card class="pa-4" style="background: #f0fdf4;">
                <div class="text-caption text-grey mb-1">صافي الراتب</div>
                <div class="text-h6 text-success font-weight-bold">
                  {{ formatCurrency(netSalary) }}
                </div>
              </v-card>
            </v-col>
          </v-row>
        </v-form>
      </div>
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeSalaryDialog">إلغاء</v-btn>
        <v-btn color="primary" variant="elevated" @click="saveSalary">حفظ</v-btn>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  leaveDialog: { type: Boolean, default: false },
  attendanceDialog: { type: Boolean, default: false },
  evaluationDialog: { type: Boolean, default: false },
  skillDialog: { type: Boolean, default: false },
  certificateDialog: { type: Boolean, default: false },
  salaryDialog: { type: Boolean, default: false },
  baseSalary: { type: Number, default: 0 }
})

const emit = defineEmits([
  'update:leaveDialog',
  'update:attendanceDialog',
  'update:evaluationDialog',
  'update:skillDialog',
  'update:certificateDialog',
  'update:salaryDialog',
  'save-leave',
  'save-attendance',
  'save-evaluation',
  'save-skill',
  'save-certificate',
  'save-salary'
])

// Dialog models
const leaveDialogModel = computed({
  get: () => props.leaveDialog,
  set: (value) => emit('update:leaveDialog', value)
})

const attendanceDialogModel = computed({
  get: () => props.attendanceDialog,
  set: (value) => emit('update:attendanceDialog', value)
})

const evaluationDialogModel = computed({
  get: () => props.evaluationDialog,
  set: (value) => emit('update:evaluationDialog', value)
})

const skillDialogModel = computed({
  get: () => props.skillDialog,
  set: (value) => emit('update:skillDialog', value)
})

const certificateDialogModel = computed({
  get: () => props.certificateDialog,
  set: (value) => emit('update:certificateDialog', value)
})

const salaryDialogModel = computed({
  get: () => props.salaryDialog,
  set: (value) => emit('update:salaryDialog', value)
})

// Forms
const leaveForm = ref({
  startDate: '',
  endDate: '',
  type: '',
  reason: '',
  status: 'قيد المراجعة'
})

const attendanceForm = ref({
  date: new Date().toISOString().split('T')[0],
  checkIn: '08:00',
  checkOut: '17:00',
  status: 'حاضر'
})

const evaluationForm = ref({
  date: new Date().toISOString().split('T')[0],
  rating: 5,
  evaluator: '',
  comments: ''
})

const skillForm = ref({
  name: ''
})

const certificateForm = ref({
  name: '',
  issuer: '',
  date: '',
  expiryDate: ''
})

const salaryForm = ref({
  month: '',
  baseSalary: 0,
  bonuses: 0,
  deductions: 0,
  status: 'قيد الدفع'
})

// Options
const leaveTypeOptions = [
  { title: 'سنوية', value: 'سنوية' },
  { title: 'مرضية', value: 'مرضية' },
  { title: 'طارئة', value: 'طارئة' },
  { title: 'أمومة', value: 'أمومة' },
  { title: 'أخرى', value: 'أخرى' }
]

const leaveStatusOptions = [
  { title: 'قيد المراجعة', value: 'قيد المراجعة' },
  { title: 'موافق', value: 'موافق' },
  { title: 'مرفوض', value: 'مرفوض' }
]

const attendanceStatusOptions = [
  { title: 'حاضر', value: 'حاضر' },
  { title: 'غائب', value: 'غائب' },
  { title: 'متأخر', value: 'متأخر' }
]

const salaryStatusOptions = [
  { title: 'قيد الدفع', value: 'قيد الدفع' },
  { title: 'مدفوع', value: 'مدفوع' },
  { title: 'معلق', value: 'معلق' }
]

// Computed
const netSalary = computed(() => {
  return parseFloat(salaryForm.value.baseSalary || 0) +
         parseFloat(salaryForm.value.bonuses || 0) -
         parseFloat(salaryForm.value.deductions || 0)
})

// Methods
const formatCurrency = (amount) => {
  if (!amount) return '0 د.ع'
  const formatted = new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
  return formatted + ' د.ع'
}

// Leave
const closeLeaveDialog = () => {
  leaveDialogModel.value = false
  leaveForm.value = { startDate: '', endDate: '', type: '', reason: '', status: 'قيد المراجعة' }
}

const saveLeave = () => {
  emit('save-leave', { ...leaveForm.value })
  closeLeaveDialog()
}

// Attendance
const closeAttendanceDialog = () => {
  attendanceDialogModel.value = false
  attendanceForm.value = { date: '', checkIn: '', checkOut: '', status: 'حاضر' }
}

const saveAttendance = () => {
  emit('save-attendance', { ...attendanceForm.value })
  closeAttendanceDialog()
}

// Evaluation
const closeEvaluationDialog = () => {
  evaluationDialogModel.value = false
  evaluationForm.value = { date: '', rating: 5, evaluator: '', comments: '' }
}

const saveEvaluation = () => {
  emit('save-evaluation', { ...evaluationForm.value })
  closeEvaluationDialog()
}

// Skill
const closeSkillDialog = () => {
  skillDialogModel.value = false
  skillForm.value = { name: '' }
}

const saveSkill = () => {
  emit('save-skill', { ...skillForm.value })
  closeSkillDialog()
}

// Certificate
const closeCertificateDialog = () => {
  certificateDialogModel.value = false
  certificateForm.value = { name: '', issuer: '', date: '', expiryDate: '' }
}

const saveCertificate = () => {
  emit('save-certificate', { ...certificateForm.value })
  closeCertificateDialog()
}

// Salary
const closeSalaryDialog = () => {
  salaryDialogModel.value = false
  salaryForm.value = { month: '', baseSalary: 0, bonuses: 0, deductions: 0, status: 'قيد الدفع' }
}

const saveSalary = () => {
  emit('save-salary', { ...salaryForm.value, netSalary: netSalary.value })
  closeSalaryDialog()
}

// Initialize salary form when dialog opens
const initSalaryForm = (baseSalary) => {
  const today = new Date()
  const month = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}`
  salaryForm.value = {
    month: month,
    baseSalary: baseSalary || 0,
    bonuses: 0,
    deductions: 0,
    status: 'قيد الدفع'
  }
}

defineExpose({
  initSalaryForm
})
</script>

<style scoped>
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
}

.dialog-actions {
  padding: 16px 24px;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  border-top: 1px solid #e5e7eb;
}
</style>
