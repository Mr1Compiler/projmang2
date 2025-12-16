<template>
  <v-dialog v-model="dialogModel" max-width="1200" scrollable persistent>
    <v-card class="details-dialog">
      <div class="dialog-header details-header">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-account-details</v-icon>
            <span class="header-title">تفاصيل الموظف: {{ employee?.name }}</span>
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

      <div class="dialog-body details-body" v-if="employee">
        <v-tabs v-model="detailsTab" bg-color="transparent" class="details-tabs">
          <v-tab value="info">
            <v-icon class="me-2">mdi-information</v-icon>
            المعلومات الأساسية
          </v-tab>
          <v-tab value="leaves">
            <v-icon class="me-2">mdi-calendar</v-icon>
            الإجازات
          </v-tab>
          <v-tab value="attendance">
            <v-icon class="me-2">mdi-clock-in</v-icon>
            الحضور والانصراف
          </v-tab>
          <v-tab value="evaluations">
            <v-icon class="me-2">mdi-star</v-icon>
            التقييمات
          </v-tab>
          <v-tab value="skills">
            <v-icon class="me-2">mdi-certificate</v-icon>
            المهارات والشهادات
          </v-tab>
          <v-tab value="fingerprint">
            <v-icon class="me-2">mdi-fingerprint</v-icon>
            البصمة
          </v-tab>
          <v-tab value="salary">
            <v-icon class="me-2">mdi-cash</v-icon>
            سجل الرواتب
          </v-tab>
        </v-tabs>

        <v-window v-model="detailsTab">
          <!-- Basic Information Tab -->
          <v-window-item value="info">
            <div class="info-content mt-4">
              <v-row>
                <v-col cols="12" md="6">
                  <v-card class="info-card" elevation="2">
                    <v-card-title class="info-card-title">المعلومات الشخصية</v-card-title>
                    <v-card-text>
                      <div class="info-item">
                        <span class="info-label">الاسم الكامل:</span>
                        <span class="info-value">{{ employee.name }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">رقم الهوية:</span>
                        <span class="info-value">{{ employee.idNumber || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">تاريخ الميلاد:</span>
                        <span class="info-value">{{ employee.birthDate || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">الجنس:</span>
                        <span class="info-value">{{ employee.gender || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">الحالة الاجتماعية:</span>
                        <span class="info-value">{{ employee.maritalStatus || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">الجنسية:</span>
                        <span class="info-value">{{ employee.nationality || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">العنوان:</span>
                        <span class="info-value">{{ employee.address || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">حالة البصمة:</span>
                        <v-chip
                          :color="employee.fingerprint ? 'primary' : 'info'"
                          size="small"
                        >
                          <v-icon start size="small">{{ employee.fingerprint ? 'mdi-check-circle' : 'mdi-alert-circle' }}</v-icon>
                          {{ employee.fingerprint ? 'مسجلة' : 'غير مسجلة' }}
                        </v-chip>
                      </div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" md="6">
                  <v-card class="info-card" elevation="2">
                    <v-card-title class="info-card-title">المعلومات الوظيفية</v-card-title>
                    <v-card-text>
                      <div class="info-item">
                        <span class="info-label">القسم:</span>
                        <span class="info-value">{{ employee.department }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">المنصب:</span>
                        <span class="info-value">{{ employee.position }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">تاريخ التوظيف:</span>
                        <span class="info-value">{{ employee.hireDate }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">نوع العقد:</span>
                        <span class="info-value">{{ employee.contractType || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">الراتب:</span>
                        <span class="info-value">{{ formatCurrency(employee.salary) }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">الحالة:</span>
                        <v-chip class="status-chip" size="small">{{ employee.status }}</v-chip>
                      </div>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
              <v-row class="mt-2">
                <v-col cols="12">
                  <v-card class="info-card" elevation="2">
                    <v-card-title class="info-card-title">معلومات الاتصال</v-card-title>
                    <v-card-text>
                      <v-row>
                        <v-col cols="12" md="6">
                          <div class="info-item">
                            <span class="info-label">رقم الهاتف:</span>
                            <span class="info-value">{{ employee.phone }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">البريد الإلكتروني:</span>
                            <span class="info-value">{{ employee.email }}</span>
                          </div>
                        </v-col>
                        <v-col cols="12" md="6">
                          <div class="info-item">
                            <span class="info-label">جهة الاتصال في الطوارئ:</span>
                            <span class="info-value">{{ employee.emergencyContact || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">رقم هاتف الطوارئ:</span>
                            <span class="info-value">{{ employee.emergencyPhone || '-' }}</span>
                          </div>
                        </v-col>
                      </v-row>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
            </div>
          </v-window-item>

          <!-- Leaves Tab -->
          <v-window-item value="leaves">
            <div class="leaves-content mt-4">
              <div class="d-flex justify-space-between align-center mb-4">
                <h3 class="section-title">سجل الإجازات</h3>
                <v-btn v-if="canWrite" color="primary" @click="$emit('add-leave')">
                  <v-icon class="me-2">mdi-plus</v-icon>
                  إضافة إجازة
                </v-btn>
              </div>
              <v-data-table
                :headers="leaveHeaders"
                :items="employee.leaves || []"
                class="details-table"
                no-data-text="لا توجد إجازات مسجلة"
              >
                <template #item.startDate="{ item }">
                  <span class="date-text">{{ item.startDate }}</span>
                </template>
                <template #item.endDate="{ item }">
                  <span class="date-text">{{ item.endDate }}</span>
                </template>
                <template #item.days="{ item }">
                  <span class="cost-text">{{ item.days }} يوم</span>
                </template>
                <template #item.type="{ item }">
                  <v-chip size="small">{{ item.type }}</v-chip>
                </template>
                <template #item.status="{ item }">
                  <v-chip size="small" :color="item.status === 'موافق' ? 'success' : item.status === 'مرفوض' ? 'error' : 'warning'">
                    {{ item.status }}
                  </v-chip>
                </template>
              </v-data-table>
            </div>
          </v-window-item>

          <!-- Attendance Tab -->
          <v-window-item value="attendance">
            <div class="attendance-content mt-4">
              <div class="d-flex justify-space-between align-center mb-4">
                <h3 class="section-title">سجل الحضور والانصراف</h3>
                <v-btn v-if="canWrite" color="primary" @click="$emit('add-attendance')">
                  <v-icon class="me-2">mdi-plus</v-icon>
                  تسجيل حضور
                </v-btn>
              </div>
              <v-data-table
                :headers="attendanceHeaders"
                :items="employee.attendance || []"
                class="details-table"
                no-data-text="لا توجد سجلات حضور"
              >
                <template #item.date="{ item }">
                  <span class="date-text">{{ item.date }}</span>
                </template>
                <template #item.checkIn="{ item }">
                  <span class="date-text">{{ item.checkIn }}</span>
                </template>
                <template #item.checkOut="{ item }">
                  <span class="date-text">{{ item.checkOut || '-' }}</span>
                </template>
                <template #item.hours="{ item }">
                  <span class="cost-text">{{ item.hours }} ساعة</span>
                </template>
                <template #item.status="{ item }">
                  <v-chip size="small" :color="item.status === 'حاضر' ? 'success' : 'error'">
                    {{ item.status }}
                  </v-chip>
                </template>
              </v-data-table>
            </div>
          </v-window-item>

          <!-- Evaluations Tab -->
          <v-window-item value="evaluations">
            <div class="evaluations-content mt-4">
              <div class="d-flex justify-space-between align-center mb-4">
                <h3 class="section-title">سجل التقييمات</h3>
                <v-btn v-if="canWrite" color="primary" @click="$emit('add-evaluation')">
                  <v-icon class="me-2">mdi-plus</v-icon>
                  إضافة تقييم
                </v-btn>
              </div>
              <v-data-table
                :headers="evaluationHeaders"
                :items="employee.evaluations || []"
                class="details-table"
                no-data-text="لا توجد تقييمات مسجلة"
              >
                <template #item.date="{ item }">
                  <span class="date-text">{{ item.date }}</span>
                </template>
                <template #item.rating="{ item }">
                  <div class="rating-display">
                    <v-icon v-for="i in 5" :key="i" :color="i <= item.rating ? 'warning' : 'grey'" size="small">
                      mdi-star{{ i <= item.rating ? '' : '-outline' }}
                    </v-icon>
                    <span class="ms-2">{{ item.rating }}/5</span>
                  </div>
                </template>
                <template #item.evaluator="{ item }">
                  <span class="project-name">{{ item.evaluator }}</span>
                </template>
              </v-data-table>
            </div>
          </v-window-item>

          <!-- Skills Tab -->
          <v-window-item value="skills">
            <div class="skills-content mt-4">
              <v-row>
                <v-col cols="12" md="6">
                  <div class="d-flex justify-space-between align-center mb-4">
                    <h3 class="section-title">المهارات</h3>
                    <v-btn v-if="canWrite" color="primary" size="small" @click="$emit('add-skill')">
                      <v-icon class="me-2">mdi-plus</v-icon>
                      إضافة مهارة
                    </v-btn>
                  </div>
                  <v-card class="info-card" elevation="2">
                    <v-card-text>
                      <div v-if="employee.skills && employee.skills.length > 0">
                        <v-chip
                          v-for="(skill, index) in employee.skills"
                          :key="index"
                          class="ma-1"
                          color="primary"
                          closable
                          @click:close="$emit('delete-skill', skill)"
                        >
                          {{ skill }}
                        </v-chip>
                      </div>
                      <div v-else class="text-center text-grey">لا توجد مهارات مسجلة</div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" md="6">
                  <div class="d-flex justify-space-between align-center mb-4">
                    <h3 class="section-title">الشهادات</h3>
                    <v-btn v-if="canWrite" color="primary" size="small" @click="$emit('add-certificate')">
                      <v-icon class="me-2">mdi-plus</v-icon>
                      إضافة شهادة
                    </v-btn>
                  </div>
                  <v-card class="info-card" elevation="2">
                    <v-card-text>
                      <div v-if="employee.certificates && employee.certificates.length > 0">
                        <div
                          v-for="(cert, index) in employee.certificates"
                          :key="index"
                          class="certificate-item mb-2"
                        >
                          <div class="d-flex justify-space-between align-center">
                            <div>
                              <div class="font-weight-bold">{{ cert.name }}</div>
                              <div class="text-caption text-grey">{{ cert.issuer }} - {{ cert.date }}</div>
                            </div>
                            <div class="d-flex align-center gap-2">
                              <v-icon color="success">mdi-certificate</v-icon>
                              <v-btn
                                v-if="canWrite"
                                icon
                                size="small"
                                color="error"
                                variant="text"
                                @click="$emit('delete-certificate', cert.id)"
                              >
                                <v-icon size="small">mdi-delete</v-icon>
                              </v-btn>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div v-else class="text-center text-grey">لا توجد شهادات مسجلة</div>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
            </div>
          </v-window-item>

          <!-- Fingerprint Tab -->
          <v-window-item value="fingerprint">
            <div class="fingerprint-content mt-4">
              <v-card class="fingerprint-management-card" elevation="2">
                <v-card-title class="fingerprint-card-title">
                  <v-icon class="me-2">mdi-fingerprint</v-icon>
                  إدارة بصمة الموظف
                </v-card-title>
                <v-card-text>
                  <div class="fingerprint-status-section mb-6">
                    <h3 class="section-title mb-3">حالة البصمة</h3>
                    <v-chip
                      :color="employee.fingerprint ? 'primary' : 'info'"
                      size="large"
                      class="mb-4 fingerprint-chip"
                    >
                      <v-icon start>{{ employee.fingerprint ? 'mdi-check-circle' : 'mdi-alert-circle' }}</v-icon>
                      {{ employee.fingerprint ? 'تم تسجيل البصمة بنجاح' : 'لم يتم تسجيل البصمة بعد' }}
                    </v-chip>
                    <div v-if="employee.fingerprint" class="fingerprint-info mt-4">
                      <div class="info-item">
                        <span class="info-label">تاريخ التسجيل:</span>
                        <span class="info-value">{{ employee.fingerprintDate || '-' }}</span>
                      </div>
                      <div class="info-item">
                        <span class="info-label">حالة البصمة:</span>
                        <v-chip color="primary" size="small">نشطة</v-chip>
                      </div>
                    </div>
                  </div>

                  <v-divider class="my-6"></v-divider>

                  <div class="fingerprint-actions">
                    <h3 class="section-title mb-3">الإجراءات</h3>
                    <v-row>
                      <v-col cols="12" md="6">
                        <v-btn
                          v-if="canWrite"
                          color="primary"
                          variant="elevated"
                          @click="$emit('register-fingerprint')"
                          block
                          size="large"
                          :disabled="loadingFingerprint"
                          class="mb-2 fingerprint-btn"
                        >
                          <v-icon class="me-2">{{ loadingFingerprint ? 'mdi-loading mdi-spin' : 'mdi-fingerprint' }}</v-icon>
                          {{ employee.fingerprint ? 'تحديث البصمة' : 'تسجيل بصمة جديدة' }}
                        </v-btn>
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-btn
                          v-if="canWrite"
                          color="error"
                          variant="elevated"
                          @click="$emit('delete-fingerprint')"
                          block
                          size="large"
                          :disabled="!employee.fingerprint || loadingFingerprint"
                          class="fingerprint-delete-btn"
                        >
                          <v-icon class="me-2">mdi-delete</v-icon>
                          حذف البصمة
                        </v-btn>
                      </v-col>
                    </v-row>
                    <p class="fingerprint-hint mt-4 text-caption text-grey">
                      <v-icon size="small" class="me-1">mdi-information</v-icon>
                      عند الضغط على تسجيل البصمة، سيتم توجيهك لتسجيل بصمة الموظف باستخدام قارئ البصمة المتصل
                    </p>
                  </div>
                </v-card-text>
              </v-card>
            </div>
          </v-window-item>

          <!-- Salary History Tab -->
          <v-window-item value="salary">
            <div class="salary-content mt-4">
              <div class="d-flex justify-space-between align-center mb-4">
                <h3 class="section-title">سجل الرواتب</h3>
                <v-btn v-if="canWrite" color="primary" @click="$emit('add-salary')">
                  <v-icon class="me-2">mdi-plus</v-icon>
                  إضافة راتب
                </v-btn>
              </div>
              <v-data-table
                :headers="salaryHeaders"
                :items="employee.salaryHistory || []"
                class="details-table"
                no-data-text="لا توجد سجلات رواتب"
              >
                <template #item.month="{ item }">
                  <span class="date-text">{{ item.month }}</span>
                </template>
                <template #item.baseSalary="{ item }">
                  <span class="cost-text">{{ formatCurrency(item.baseSalary) }}</span>
                </template>
                <template #item.bonuses="{ item }">
                  <span class="cost-text">{{ formatCurrency(item.bonuses || 0) }}</span>
                </template>
                <template #item.deductions="{ item }">
                  <span class="cost-text text-error">{{ formatCurrency(item.deductions || 0) }}</span>
                </template>
                <template #item.netSalary="{ item }">
                  <span class="cost-text text-success font-weight-bold">{{ formatCurrency(item.netSalary) }}</span>
                </template>
                <template #item.status="{ item }">
                  <v-chip size="small" :color="item.status === 'مدفوع' ? 'success' : 'warning'">
                    {{ item.status }}
                  </v-chip>
                </template>
              </v-data-table>
            </div>
          </v-window-item>
        </v-window>
      </div>

      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeDialog">
          إغلاق
        </v-btn>
        <v-btn v-if="canWrite" color="primary" variant="elevated" @click="$emit('edit', employee)">
          <v-icon class="me-2">mdi-pencil</v-icon>
          تعديل
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
  canWrite: {
    type: Boolean,
    default: true
  },
  loadingFingerprint: {
    type: Boolean,
    default: false
  },
  initialTab: {
    type: String,
    default: 'info'
  }
})

const emit = defineEmits([
  'update:modelValue',
  'close',
  'edit',
  'add-leave',
  'add-attendance',
  'add-evaluation',
  'add-skill',
  'delete-skill',
  'add-certificate',
  'delete-certificate',
  'add-salary',
  'register-fingerprint',
  'delete-fingerprint'
])

const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const detailsTab = ref('info')

watch(() => props.initialTab, (newTab) => {
  if (newTab) {
    detailsTab.value = newTab
  }
}, { immediate: true })

watch(() => props.modelValue, (isOpen) => {
  if (isOpen && props.initialTab) {
    detailsTab.value = props.initialTab
  }
})

const leaveHeaders = [
  { title: 'تاريخ البداية', key: 'startDate', align: 'center' },
  { title: 'تاريخ النهاية', key: 'endDate', align: 'center' },
  { title: 'عدد الأيام', key: 'days', align: 'center' },
  { title: 'نوع الإجازة', key: 'type', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' }
]

const attendanceHeaders = [
  { title: 'التاريخ', key: 'date', align: 'center' },
  { title: 'وقت الدخول', key: 'checkIn', align: 'center' },
  { title: 'وقت الخروج', key: 'checkOut', align: 'center' },
  { title: 'عدد الساعات', key: 'hours', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' }
]

const evaluationHeaders = [
  { title: 'التاريخ', key: 'date', align: 'center' },
  { title: 'التقييم', key: 'rating', align: 'center' },
  { title: 'المقيّم', key: 'evaluator', align: 'right' },
  { title: 'ملاحظات', key: 'comments', align: 'right' }
]

const salaryHeaders = [
  { title: 'الشهر', key: 'month', align: 'center' },
  { title: 'الراتب الأساسي', key: 'baseSalary', align: 'center' },
  { title: 'المكافآت', key: 'bonuses', align: 'center' },
  { title: 'الخصومات', key: 'deductions', align: 'center' },
  { title: 'صافي الراتب', key: 'netSalary', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' }
]

const closeDialog = () => {
  dialogModel.value = false
  detailsTab.value = 'info'
  emit('close')
}

const formatCurrency = (amount) => {
  if (!amount) return '0 د.ع'
  const formatted = new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
  return formatted + ' د.ع'
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

.info-card {
  border-radius: 12px;
}

.info-card-title {
  background: #f8fafc;
  font-size: var(--font-size-base);
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
  border-bottom: 1px solid #f1f5f9;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  color: #64748b;
  font-weight: 500;
}

.info-value {
  color: #1e293b;
  font-weight: 600;
}

.section-title {
  font-size: var(--font-size-base-plus);
  font-weight: 600;
  color: #1e293b;
}

.fingerprint-management-card {
  border-radius: 16px;
}

.fingerprint-card-title {
  background: #f8fafc;
}

.fingerprint-chip {
  width: 100%;
  justify-content: center;
}

.certificate-item {
  padding: 12px;
  border-radius: 8px;
  background: #f8fafc;
}

.rating-display {
  display: flex;
  align-items: center;
}
</style>
