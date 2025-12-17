<template>
  <v-container fluid class="pa-6">
    <!-- Page Header -->
    <PageHeader
      title="إدارة المديونون"
      subtitle="إدارة حسابات المديونون والمستحقات المالية"
      icon="💳"
    />

    <!-- Statistics Cards -->
    <DebtorStats
      :total-debtors="debtors.length"
      :total-debt="totalDebtAmount"
      :overdue-count="overdueCount"
      :paid-count="paidCount"
    />

    <!-- Search & Filters -->
    <DebtorFilters
      v-model:search-query="searchQuery"
      v-model:status-filter="statusFilter"
      v-model:amount-filter="amountFilter"
      @reset="resetFilters"
    />

    <!-- Debtors Table -->
    <DebtorTable
      :debtors="filteredDebtors"
      :loading="loading"
      :can-write="canWriteDebtors"
      :can-delete="canDeleteDebtors"
      @add="openAddDialog"
      @view="viewDebtor"
      @edit="editDebtor"
      @delete="confirmDeleteDebtor"
      @mark-paid="markAsPaid"
      @view-payments="viewDebtsAndPayments"
      @refresh="refreshData"
      @export="exportData"
    />

    <!-- Add/Edit Debtor Dialog -->
    <DebtorForm
      v-model="dialog"
      :debtor="selectedDebtor"
      :is-edit="isEdit"
      @save="saveDebtor"
      @close="closeDialog"
    />

    <!-- View Debtor Details Dialog -->
    <DebtorDetails
      v-model="viewDialog"
      :debtor="selectedDebtor"
      @close="viewDialog = false"
    />

    <!-- Debts & Payments Dialog -->
    <DebtorPayments
      v-model="debtsPaymentsDialog"
      :debtor="selectedDebtor"
      :can-write="canWriteDebtors"
      :can-delete="canDeleteDebtors"
      @close="debtsPaymentsDialog = false"
      @add-debt="addDebt"
      @edit-debt="editDebt"
      @delete-debt="deleteDebt"
      @add-payment="addPayment"
      @edit-payment="editPayment"
      @delete-payment="deletePayment"
    />

    <!-- Delete Confirmation Dialog -->
    <DeleteConfirmDialog
      v-model="deleteDialog"
      :title="'تأكيد الحذف'"
      :message="`هل أنت متأكد من حذف المدين &quot;${selectedDebtor?.name}&quot;؟`"
      @confirm="confirmDelete"
      @cancel="deleteDialog = false"
    />
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { toast } from 'vue3-toastify'
import { useDebtorsStore } from '@/stores/debtors'
import { usePermissions } from '@/composables/usePermissions'
import { PageHeader } from '@/components/shared'
import {
  DebtorStats,
  DebtorFilters,
  DebtorTable,
  DebtorForm,
  DebtorDetails,
  DebtorPayments
} from '@/components/debtors'
import { DeleteConfirmDialog } from '@/components/projects'

const debtorsStore = useDebtorsStore()
const { canWrite, canDelete } = usePermissions()

// Get reactive state from store
const { debtors, loading } = storeToRefs(debtorsStore)

// Permission checks
const canWriteDebtors = canWrite('/debtors')
const canDeleteDebtors = canDelete('/debtors')

// Local state
const dialog = ref(false)
const viewDialog = ref(false)
const debtsPaymentsDialog = ref(false)
const deleteDialog = ref(false)
const isEdit = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const amountFilter = ref('')
const selectedDebtor = ref(null)

// Computed
const filteredDebtors = computed(() => {
  let filtered = debtors.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(debtor =>
      debtor.name?.toLowerCase().includes(query) ||
      debtor.email?.toLowerCase().includes(query) ||
      debtor.phone?.includes(searchQuery.value)
    )
  }

  if (statusFilter.value) {
    filtered = filtered.filter(debtor => debtor.status === statusFilter.value)
  }

  if (amountFilter.value) {
    filtered = filtered.filter(debtor => {
      const amount = debtor.amount || 0
      switch (amountFilter.value) {
        case 'low': return amount < 10000
        case 'medium': return amount >= 10000 && amount <= 20000
        case 'high': return amount > 20000
        default: return true
      }
    })
  }

  return filtered
})

const totalDebtAmount = computed(() => {
  return debtors.value
    .filter(debtor => debtor.status !== 'paid')
    .reduce((sum, debtor) => sum + (debtor.amount || 0), 0)
})

const overdueCount = computed(() => {
  return debtors.value.filter(debtor => debtor.status === 'overdue').length
})

const paidCount = computed(() => {
  return debtors.value.filter(debtor => debtor.status === 'paid').length
})

// Methods
const openAddDialog = () => {
  isEdit.value = false
  selectedDebtor.value = null
  dialog.value = true
}

const editDebtor = (debtor) => {
  isEdit.value = true
  selectedDebtor.value = { ...debtor }
  dialog.value = true
}

const viewDebtor = (debtor) => {
  selectedDebtor.value = debtor
  viewDialog.value = true
}

const viewDebtsAndPayments = (debtor) => {
  selectedDebtor.value = debtor
  debtsPaymentsDialog.value = true
}

const confirmDeleteDebtor = (debtor) => {
  selectedDebtor.value = debtor
  deleteDialog.value = true
}

const confirmDelete = async () => {
  if (selectedDebtor.value) {
    await debtorsStore.deleteDebtor(selectedDebtor.value.id)
  }
  deleteDialog.value = false
  selectedDebtor.value = null
}

const saveDebtor = async (formData) => {
  if (isEdit.value && selectedDebtor.value) {
    await debtorsStore.updateDebtor(selectedDebtor.value.id, formData)
  } else {
    await debtorsStore.createDebtor({
      ...formData,
      status: 'pending',
      debts: [],
      payments: []
    })
  }
  closeDialog()
}

const closeDialog = () => {
  dialog.value = false
  isEdit.value = false
  selectedDebtor.value = null
}

const markAsPaid = async (debtor) => {
  await debtorsStore.updateDebtor(debtor.id, { ...debtor, status: 'paid' })
  toast.success('تم تحديث حالة المدين إلى مدفوع')
}

const refreshData = () => {
  debtorsStore.fetchDebtors()
}

const exportData = () => {
  toast.info('جاري تصدير البيانات...')
}

const resetFilters = () => {
  searchQuery.value = ''
  statusFilter.value = ''
  amountFilter.value = ''
}

// Debt & Payment methods (placeholders for future API integration)
const addDebt = (debtData) => {
  if (selectedDebtor.value) {
    if (!selectedDebtor.value.debts) selectedDebtor.value.debts = []
    selectedDebtor.value.debts.push({ id: Date.now(), ...debtData })
    toast.success('تم إضافة الدين بنجاح')
  }
}

const editDebt = () => {
  toast.info('تعديل الدين')
}

const deleteDebt = (debt) => {
  if (selectedDebtor.value?.debts) {
    const index = selectedDebtor.value.debts.findIndex(d => d.id === debt.id)
    if (index > -1) {
      selectedDebtor.value.debts.splice(index, 1)
      toast.success('تم حذف الدين بنجاح')
    }
  }
}

const addPayment = (paymentData) => {
  if (selectedDebtor.value) {
    if (!selectedDebtor.value.payments) selectedDebtor.value.payments = []
    selectedDebtor.value.payments.push({ id: Date.now(), ...paymentData })
    toast.success('تم إضافة الدفعة بنجاح')
  }
}

const editPayment = () => {
  toast.info('تعديل الدفعة')
}

const deletePayment = (payment) => {
  if (selectedDebtor.value?.payments) {
    const index = selectedDebtor.value.payments.findIndex(p => p.id === payment.id)
    if (index > -1) {
      selectedDebtor.value.payments.splice(index, 1)
      toast.success('تم حذف الدفعة بنجاح')
    }
  }
}

onMounted(() => {
  debtorsStore.fetchDebtors()
})
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/debtors.css';
</style>
