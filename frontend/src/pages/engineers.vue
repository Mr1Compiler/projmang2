<template>
  <v-container class="fill-height engineers-page" fluid>
    <div class="centered-content">
      <!-- Header Section -->
      <PageHeader
        title="المهندسين"
        subtitle="إدارة وتتبع جميع المهندسين والمشاريع"
        mdi-icon="mdi-account-hard-hat"
      />

      <!-- Summary Cards -->
      <EngineerStats
        :total-engineers="engineers.length"
        :active-engineers="activeEngineersCount"
        :total-projects="totalProjectsCount"
        :average-rating="averageRatingValue"
      />

      <!-- Engineers Table -->
      <EngineerTable
        :engineers="engineers"
        v-model:search-query="searchQuery"
        :loading="loading"
        :can-write="canWriteEngineers"
        :can-delete="canDeleteEngineers"
        @add="openAddEngineerDialog"
        @view="viewEngineer"
        @edit="editEngineer"
        @delete="confirmDeleteEngineer"
        @manage-projects="openProjectsDialog"
      />

      <!-- Add/Edit Engineer Dialog -->
      <EngineerForm
        v-model="dialog"
        :engineer="selectedEngineer"
        :is-editing="isEditing"
        @save="saveEngineer"
        @close="closeDialog"
      />

      <!-- Delete Confirmation Dialog -->
      <DeleteConfirmDialog
        v-model="deleteDialog"
        :title="'تأكيد الحذف'"
        :message="`هل أنت متأكد من حذف المهندس &quot;${selectedEngineer?.name}&quot;؟`"
        @confirm="confirmDelete"
        @cancel="deleteDialog = false"
      />

      <!-- Engineer Details Dialog -->
      <EngineerDetails
        v-model="detailsDialog"
        :engineer="selectedEngineer"
        :can-edit="canWriteEngineers"
        @edit="editEngineer"
        @close="detailsDialog = false"
      />

      <!-- Engineer Projects Dialog -->
      <EngineerProjects
        v-model="projectsDialog"
        :engineer="selectedEngineerForProjects"
        :available-projects="availableProjects"
        @add-project="handleAddProject"
        @remove-project="handleRemoveProject"
        @close="projectsDialog = false"
      />
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useEngineersStore } from '@/stores/engineers'
import { usePermissions } from '@/composables/usePermissions'
import { PageHeader } from '@/components/shared'
import {
  EngineerStats,
  EngineerTable,
  EngineerForm,
  EngineerDetails,
  EngineerProjects
} from '@/components/engineers'
import { DeleteConfirmDialog } from '@/components/projects'

// Store & Permissions
const engineersStore = useEngineersStore()
const { canWrite, canDelete } = usePermissions()

// Get reactive state from store
const { engineers, loading, availableProjects } = storeToRefs(engineersStore)

// Permission checks
const canWriteEngineers = canWrite('/engineers')
const canDeleteEngineers = canDelete('/engineers')

// Local state
const searchQuery = ref('')
const dialog = ref(false)
const deleteDialog = ref(false)
const detailsDialog = ref(false)
const projectsDialog = ref(false)
const isEditing = ref(false)
const selectedEngineer = ref(null)
const selectedEngineerForProjects = ref(null)

// Computed properties
const activeEngineersCount = computed(() => {
  return engineers.value.filter(e => e.status === 'active').length
})

const totalProjectsCount = computed(() => {
  return engineers.value.reduce((total, engineer) => total + (engineer.projects?.length || 0), 0)
})

const averageRatingValue = computed(() => {
  if (engineers.value.length === 0) return '0.0'
  const total = engineers.value.reduce((sum, engineer) => sum + (engineer.rating || 0), 0)
  return (total / engineers.value.length).toFixed(1)
})

// Methods
const openAddEngineerDialog = () => {
  isEditing.value = false
  selectedEngineer.value = null
  dialog.value = true
}

const viewEngineer = (engineer) => {
  selectedEngineer.value = engineer
  detailsDialog.value = true
}

const editEngineer = (engineer) => {
  isEditing.value = true
  selectedEngineer.value = engineer
  detailsDialog.value = false
  dialog.value = true
}

const confirmDeleteEngineer = (engineer) => {
  selectedEngineer.value = engineer
  deleteDialog.value = true
}

const saveEngineer = async (formData) => {
  if (isEditing.value && selectedEngineer.value) {
    await engineersStore.updateEngineer(selectedEngineer.value.id, formData)
  } else {
    await engineersStore.createEngineer(formData)
  }
  closeDialog()
}

const closeDialog = () => {
  dialog.value = false
  isEditing.value = false
  selectedEngineer.value = null
}

const confirmDelete = async () => {
  if (selectedEngineer.value) {
    await engineersStore.deleteEngineer(selectedEngineer.value.id)
  }
  deleteDialog.value = false
  selectedEngineer.value = null
}

const openProjectsDialog = (engineer) => {
  selectedEngineerForProjects.value = engineer
  projectsDialog.value = true
}

const handleAddProject = async ({ engineerId, project }) => {
  await engineersStore.addProjectToEngineer(engineerId, project.id)
  // Update local reference
  const engineer = engineers.value.find(e => e.id === engineerId)
  if (engineer) {
    selectedEngineerForProjects.value = engineer
  }
}

const handleRemoveProject = async ({ engineerId, projectId }) => {
  await engineersStore.removeProjectFromEngineer(engineerId, projectId)
  // Update local reference
  const engineer = engineers.value.find(e => e.id === engineerId)
  if (engineer) {
    selectedEngineerForProjects.value = engineer
  }
}

// Lifecycle
onMounted(async () => {
  await engineersStore.fetchEngineers()
  await engineersStore.fetchAvailableProjects()
})
</script>

<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/engineers.css';
</style>
