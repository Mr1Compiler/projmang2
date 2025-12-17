<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <PageHeader
        title="المصاريف الإدارية"
        subtitle="إدارة وتتبع جميع المصاريف الإدارية والعامة"
        icon="💰"
      />

      <!-- Summary Cards -->
      <ExpenseStats
        :total-expenses="totalExpensesCount"
        :active-expenses="activeExpensesCount"
        :pending-expenses="pendingExpensesCount"
        :total-cost="totalCostValue"
      />

      <!-- Search & Filters -->
      <ExpenseFilters
        v-model:search-query="searchQuery"
        v-model:expense-type="selectedExpenseType"
        v-model:status="selectedStatus"
        :can-add="canWriteExpenses"
        @search="handleSearch"
        @add="openAddExpenseDialog"
      />

      <!-- Expenses Table -->
      <ExpenseTable
        :items="filteredExpenses"
        :search-query="searchQuery"
        :loading="loading"
        :can-edit="canWriteExpenses"
        :can-delete="canDeleteExpenses"
        @view="viewExpenseDetails"
        @edit="editExpense"
        @delete="confirmDeleteExpense"
      />

      <!-- Add/Edit Administrative Expense Dialog -->
      <ExpenseForm
        v-model="expenseDialog"
        :expense="selectedExpense"
        :is-editing="isEditingExpense"
        @save="saveExpense"
        @close="closeExpenseDialog"
      />

      <!-- Delete Confirmation Dialog -->
      <DeleteConfirmDialog
        v-model="deleteDialog"
        :title="'تأكيد الحذف'"
        :message="`هل أنت متأكد من حذف المصروف؟`"
        @confirm="confirmDelete"
        @cancel="deleteDialog = false"
      />
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useExpensesStore } from '@/stores/expenses'
import { usePermissions } from '@/composables/usePermissions'
import { PageHeader } from '@/components/shared'
import {
  ExpenseStats,
  ExpenseFilters,
  ExpenseTable,
  ExpenseForm
} from '@/components/expenses'
import { DeleteConfirmDialog } from '@/components/projects'

const router = useRouter()
const expensesStore = useExpensesStore()
const { canWrite, canDelete } = usePermissions()

// Get reactive state from store
const { expenses, loading } = storeToRefs(expensesStore)

// Permission checks
const canWriteExpenses = canWrite('/expenses')
const canDeleteExpenses = canDelete('/expenses')

// Local state
const expenseDialog = ref(false)
const deleteDialog = ref(false)
const isEditingExpense = ref(false)
const searchQuery = ref('')
const selectedExpenseType = ref('')
const selectedStatus = ref('')
const selectedExpense = ref(null)

// Computed statistics
const totalExpensesCount = computed(() => expenses.value.length)

const activeExpensesCount = computed(() => {
  const today = new Date()
  return expenses.value.filter(expense => {
    const startDate = expense.startDate ? new Date(expense.startDate) : null
    const endDate = expense.endDate ? new Date(expense.endDate) : null
    if (startDate && endDate) {
      return startDate <= today && endDate >= today
    }
    return expense.status === 'active' || expense.status === 'معتمد'
  }).length
})

const pendingExpensesCount = computed(() => {
  return expenses.value.filter(expense =>
    expense.status === 'pending' || expense.status === 'معلق'
  ).length
})

const totalCostValue = computed(() => {
  return expenses.value.reduce((sum, expense) => sum + (expense.amount || expense.cost || 0), 0)
})

// Filtered expenses
const filteredExpenses = computed(() => {
  let filtered = expenses.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(expense =>
      expense.description?.toLowerCase().includes(query) ||
      expense.projectName?.toLowerCase().includes(query) ||
      expense.notes?.toLowerCase().includes(query)
    )
  }

  if (selectedExpenseType.value) {
    filtered = filtered.filter(expense => expense.expenseType === selectedExpenseType.value)
  }

  if (selectedStatus.value) {
    filtered = filtered.filter(expense => expense.status === selectedStatus.value)
  }

  return filtered
})

// Methods
const handleSearch = () => {
  expensesStore.setFilters({ search: searchQuery.value })
  expensesStore.fetchExpenses()
}

const openAddExpenseDialog = () => {
  expenseDialog.value = true
  isEditingExpense.value = false
  selectedExpense.value = null
}

const closeExpenseDialog = () => {
  expenseDialog.value = false
  isEditingExpense.value = false
  selectedExpense.value = null
}

const editExpense = (expense) => {
  selectedExpense.value = { ...expense }
  isEditingExpense.value = true
  expenseDialog.value = true
}

const viewExpenseDetails = (expense) => {
  router.push({
    path: '/project-expenses',
    query: {
      projectName: expense.projectName || expense.description,
      projectId: expense.projectId || expense.id,
      startDate: expense.startDate,
      endDate: expense.endDate,
      cost: expense.amount || expense.cost,
      workLocation: expense.workLocation,
      notes: expense.notes
    }
  })
}

const confirmDeleteExpense = (expense) => {
  selectedExpense.value = expense
  deleteDialog.value = true
}

const confirmDelete = async () => {
  if (selectedExpense.value) {
    await expensesStore.deleteExpense(selectedExpense.value.id)
  }
  deleteDialog.value = false
  selectedExpense.value = null
}

const saveExpense = async (formData) => {
  if (isEditingExpense.value && selectedExpense.value) {
    await expensesStore.updateExpense(selectedExpense.value.id, formData)
  } else {
    await expensesStore.createExpense(formData)
  }
  closeExpenseDialog()
}

onMounted(() => {
  expensesStore.fetchExpenses()
})
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/expenses.css';
</style>
