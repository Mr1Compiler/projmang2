<template>
  <v-container class="fill-height data-page human-resources-page" fluid>
    <div class="fullscreen-content centered-content">
      <!-- Header Section -->
      <PageHeader
        title="الموارد البشرية"
        subtitle="إدارة وتتبع جميع الموظفين والموارد البشرية"
        icon="👥"
      />

      <!-- Statistics Cards -->
      <HRStats
        :total-employees="totalEmployees"
        :active-employees="activeEmployees"
        :total-departments="totalDepartments"
        :on-leave="onLeave"
        :total-salaries="totalSalaries"
        :average-salary="averageSalary"
      />

      <!-- Search & Filters -->
      <HRFilters
        v-model:search-query="searchQuery"
        v-model:selected-department="selectedDepartment"
        v-model:selected-status="selectedStatus"
        v-model:selected-position="selectedPosition"
        :department-options="departmentOptions"
        :status-options="statusOptions"
        :position-options="positionOptions"
        :can-add="canWriteHR"
        @search="searchEmployees"
        @add="openAddEmployeeDialog"
      />

      <!-- Employees Table -->
      <HRTable
        :employees="filteredEmployees"
        :search-query="searchQuery"
        :loading="loading"
        :can-edit="canWriteHR"
        :can-delete="canDeleteHR"
        @view="viewEmployeeDetails"
        @edit="editEmployee"
        @delete="handleDeleteEmployee"
      />

      <!-- Add/Edit Employee Dialog -->
      <EmployeeForm
        v-model="employeeDialog"
        :employee="selectedEmployee"
        :is-editing="isEditing"
        :loading-fingerprint="loadingFingerprint"
        @save="saveEmployee"
        @close="closeEmployeeDialog"
        @register-fingerprint="handleRegisterFingerprint"
      />

      <!-- Employee Details Dialog -->
      <EmployeeDetails
        v-model="detailsDialog"
        :employee="selectedEmployeeDetails"
        :can-write="canWriteHR"
        :loading-fingerprint="loadingFingerprint"
        :initial-tab="detailsTab"
        @close="closeDetailsDialog"
        @edit="editEmployee"
        @add-leave="openAddLeaveDialog"
        @add-attendance="openAddAttendanceDialog"
        @add-evaluation="openAddEvaluationDialog"
        @add-skill="openAddSkillDialog"
        @delete-skill="handleDeleteSkill"
        @add-certificate="openAddCertificateDialog"
        @delete-certificate="handleDeleteCertificate"
        @add-salary="openAddSalaryDialog"
        @register-fingerprint="handleRegisterFingerprintFromDetails"
        @delete-fingerprint="handleDeleteFingerprint"
      />

      <!-- Sub Dialogs -->
      <EmployeeSubDialogs
        ref="subDialogsRef"
        v-model:leave-dialog="leaveDialog"
        v-model:attendance-dialog="attendanceDialog"
        v-model:evaluation-dialog="evaluationDialog"
        v-model:skill-dialog="skillDialog"
        v-model:certificate-dialog="certificateDialog"
        v-model:salary-dialog="salaryDialog"
        :base-salary="selectedEmployeeDetails?.salary || 0"
        @save-leave="saveLeave"
        @save-attendance="saveAttendance"
        @save-evaluation="saveEvaluation"
        @save-skill="saveSkill"
        @save-certificate="saveCertificate"
        @save-salary="saveSalary"
      />
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { toast } from 'vue3-toastify'
import { useHumanResources } from '@/composables/useHumanResources'
import { usePermissions } from '@/composables/usePermissions'
import { PageHeader } from '@/components/shared'
import {
  HRStats,
  HRFilters,
  HRTable,
  EmployeeForm,
  EmployeeDetails,
  EmployeeSubDialogs
} from '@/components/human-resources'

const route = useRoute()
const { canWrite, canDelete } = usePermissions()

// Composable
const {
  employees,
  loading,
  totalEmployees,
  activeEmployees,
  totalDepartments,
  onLeave,
  totalSalaries,
  averageSalary,
  departmentOptions,
  positionOptions,
  canWriteHR,
  canDeleteHR,
  fetchEmployees,
  createEmployee,
  updateEmployee,
  deleteEmployee,
  filterEmployees,
  addLeave,
  addAttendance,
  addEvaluation,
  addSkill,
  deleteSkill,
  addCertificate,
  deleteCertificate,
  addSalaryRecord,
  registerFingerprint,
  deleteFingerprint
} = useHumanResources()

// Local state
const loadingFingerprint = ref(false)
const employeeDialog = ref(false)
const detailsDialog = ref(false)
const isEditing = ref(false)
const searchQuery = ref('')
const selectedDepartment = ref('')
const selectedStatus = ref('')
const selectedPosition = ref('')
const selectedEmployee = ref(null)
const selectedEmployeeDetails = ref(null)
const detailsTab = ref('info')
const subDialogsRef = ref(null)

// Sub-dialog states
const leaveDialog = ref(false)
const attendanceDialog = ref(false)
const evaluationDialog = ref(false)
const skillDialog = ref(false)
const certificateDialog = ref(false)
const salaryDialog = ref(false)

// Options
const statusOptions = [
  { title: 'نشط', value: 'نشط' },
  { title: 'في إجازة', value: 'في إجازة' },
  { title: 'معطل', value: 'معطل' }
]

// Computed
const filteredEmployees = computed(() => {
  return filterEmployees({
    department: selectedDepartment.value,
    status: selectedStatus.value,
    position: selectedPosition.value
  })
})

// Methods
const searchEmployees = () => {
  // Search handled by filters
}

const openAddEmployeeDialog = () => {
  employeeDialog.value = true
  isEditing.value = false
  selectedEmployee.value = null
}

const closeEmployeeDialog = () => {
  employeeDialog.value = false
  isEditing.value = false
  selectedEmployee.value = null
}

const editEmployee = (item) => {
  selectedEmployee.value = item
  isEditing.value = true
  employeeDialog.value = true
  detailsDialog.value = false
}

const viewEmployeeDetails = (item) => {
  selectedEmployeeDetails.value = item
  detailsTab.value = 'info'
  detailsDialog.value = true
}

const closeDetailsDialog = () => {
  detailsDialog.value = false
  selectedEmployeeDetails.value = null
  detailsTab.value = 'info'
}

const handleDeleteEmployee = async (item) => {
  if (confirm(`هل أنت متأكد من حذف الموظف "${item.name}"؟`)) {
    const success = await deleteEmployee(item.id)
    if (success) {
      toast.success('تم حذف الموظف بنجاح')
    }
  }
}

const saveEmployee = async (formData) => {
  if (isEditing.value && selectedEmployee.value) {
    const updated = await updateEmployee(selectedEmployee.value.id, formData)
    if (updated) toast.success('تم تحديث بيانات الموظف بنجاح')
  } else {
    const created = await createEmployee(formData)
    if (created) toast.success('تم إضافة الموظف بنجاح')
  }
  closeEmployeeDialog()
}

// Fingerprint handlers
const handleRegisterFingerprint = async () => {
  loadingFingerprint.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 2000))
    toast.success('تم تسجيل البصمة بنجاح!')
  } finally {
    loadingFingerprint.value = false
  }
}

const handleRegisterFingerprintFromDetails = async () => {
  if (!selectedEmployeeDetails.value) return
  loadingFingerprint.value = true
  try {
    const result = await registerFingerprint(selectedEmployeeDetails.value.id)
    if (result) {
      selectedEmployeeDetails.value.fingerprint = result.fingerprintId
      selectedEmployeeDetails.value.fingerprintDate = result.fingerprintDate
      toast.success('تم تسجيل البصمة بنجاح!')
    }
  } finally {
    loadingFingerprint.value = false
  }
}

const handleDeleteFingerprint = () => {
  if (!selectedEmployeeDetails.value || !confirm('هل أنت متأكد من حذف بصمة الموظف؟')) return
  const success = deleteFingerprint(selectedEmployeeDetails.value.id)
  if (success) {
    selectedEmployeeDetails.value.fingerprint = null
    selectedEmployeeDetails.value.fingerprintDate = null
    toast.success('تم حذف البصمة بنجاح')
  }
}

// Sub-dialog openers
const openAddLeaveDialog = () => { leaveDialog.value = true }
const openAddAttendanceDialog = () => { attendanceDialog.value = true }
const openAddEvaluationDialog = () => { evaluationDialog.value = true }
const openAddSkillDialog = () => { skillDialog.value = true }
const openAddCertificateDialog = () => { certificateDialog.value = true }
const openAddSalaryDialog = () => {
  if (subDialogsRef.value) {
    subDialogsRef.value.initSalaryForm(selectedEmployeeDetails.value?.salary || 0)
  }
  salaryDialog.value = true
}

// Sub-dialog save handlers
const saveLeave = (data) => {
  if (!selectedEmployeeDetails.value) return
  addLeave(selectedEmployeeDetails.value.id, data)
  toast.success('تم إضافة الإجازة بنجاح')
}

const saveAttendance = (data) => {
  if (!selectedEmployeeDetails.value) return
  addAttendance(selectedEmployeeDetails.value.id, data)
  toast.success('تم تسجيل الحضور بنجاح')
}

const saveEvaluation = (data) => {
  if (!selectedEmployeeDetails.value) return
  addEvaluation(selectedEmployeeDetails.value.id, data)
  toast.success('تم إضافة التقييم بنجاح')
}

const saveSkill = (data) => {
  if (!selectedEmployeeDetails.value || !data.name) return
  addSkill(selectedEmployeeDetails.value.id, data.name)
  toast.success('تم إضافة المهارة بنجاح')
}

const handleDeleteSkill = (skillName) => {
  if (!selectedEmployeeDetails.value) return
  deleteSkill(selectedEmployeeDetails.value.id, skillName)
}

const saveCertificate = (data) => {
  if (!selectedEmployeeDetails.value) return
  addCertificate(selectedEmployeeDetails.value.id, data)
  toast.success('تم إضافة الشهادة بنجاح')
}

const handleDeleteCertificate = (certId) => {
  if (!selectedEmployeeDetails.value) return
  deleteCertificate(selectedEmployeeDetails.value.id, certId)
}

const saveSalary = (data) => {
  if (!selectedEmployeeDetails.value) return
  addSalaryRecord(selectedEmployeeDetails.value.id, data)
  toast.success('تم إضافة الراتب بنجاح')
}

// Watch for fingerprint hash navigation
watch(() => route.hash, (newHash) => {
  if (newHash === '#fingerprint' || newHash === '#بصمة') {
    nextTick(() => {
      if (selectedEmployeeDetails.value) {
        detailsTab.value = 'fingerprint'
      } else if (employees.value.length > 0) {
        viewEmployeeDetails(employees.value[0])
        nextTick(() => { detailsTab.value = 'fingerprint' })
      }
    })
  }
}, { immediate: true })

onMounted(() => {
  fetchEmployees()
  const hash = route.hash || window.location.hash
  if (hash === '#fingerprint' || hash === '#بصمة') {
    if (employees.value.length > 0) {
      setTimeout(() => {
        viewEmployeeDetails(employees.value[0])
        nextTick(() => { detailsTab.value = 'fingerprint' })
      }, 500)
    }
  }
})
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/human-resources.css';
</style>
