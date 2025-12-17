<template>
  <v-dialog v-model="dialogModel" max-width="800" persistent>
    <v-card v-if="engineer" class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-success) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-icon size="24" color="white" class="header-icon">mdi-briefcase</v-icon>
            <span class="header-title">مشاريع المهندس - {{ engineer.name }}</span>
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

      <!-- Dialog Body -->
      <div class="dialog-body">
        <!-- Add Project Section -->
        <v-card variant="outlined" class="mb-4 pa-4 add-project-section">
          <v-card-title class="text-subtitle-1 pa-0 mb-3">
            <v-icon class="me-2" color="success" size="small">mdi-plus-circle</v-icon>
            إضافة مشروع جديد
          </v-card-title>
          <v-row>
            <v-col cols="12" md="8">
              <v-autocomplete
                v-model="selectedProject"
                :items="filteredAvailableProjects"
                item-title="name"
                item-value="id"
                label="اختر المشروع"
                variant="outlined"
                density="comfortable"
                prepend-inner-icon="mdi-briefcase"
                class="form-field-enhanced enhanced-select input-text-dark"
                color="black"
                :menu-props="{
                  maxHeight: 300,
                  zIndex: 50000
                }"
                bg-color="white"
                no-data-text="لا توجد مشاريع متاحة"
                return-object
              />
            </v-col>
            <v-col cols="12" md="4">
              <v-btn
                color="success"
                variant="elevated"
                @click="addProject"
                :disabled="!selectedProject"
                block
                class="add-project-btn"
              >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة المشروع
              </v-btn>
            </v-col>
          </v-row>
        </v-card>

        <!-- Engineer's Projects List -->
        <v-card variant="outlined" class="projects-list-card">
          <v-card-title class="text-subtitle-1 pa-3">
            <v-icon class="me-2" color="info" size="small">mdi-format-list-bulleted</v-icon>
            مشاريع المهندس الحالية
            <v-chip class="ms-2" size="x-small" color="info" variant="tonal">
              {{ engineerProjects.length }}
            </v-chip>
          </v-card-title>
          <v-card-text class="pa-0">
            <div v-if="engineerProjects.length > 0">
              <v-list>
                <v-list-item
                  v-for="project in engineerProjects"
                  :key="project.id"
                  class="project-item"
                >
                  <template v-slot:prepend>
                    <v-icon color="success">mdi-briefcase-check</v-icon>
                  </template>
                  <v-list-item-title class="font-weight-medium">
                    {{ project.name }}
                  </v-list-item-title>
                  <v-list-item-subtitle v-if="project.status">
                    <v-chip size="x-small" :color="getStatusColor(project.status)" variant="tonal">
                      {{ getStatusLabel(project.status) }}
                    </v-chip>
                  </v-list-item-subtitle>
                  <template v-slot:append>
                    <v-btn
                      icon="mdi-delete"
                      size="small"
                      variant="text"
                      color="error"
                      @click="removeProject(project.id)"
                    />
                  </template>
                </v-list-item>
              </v-list>
            </div>
            <div v-else class="text-center pa-4 text-medium-emphasis">
              <v-icon size="48" color="grey-lighten-1" class="mb-2">mdi-briefcase-outline</v-icon>
              <p class="text-subtitle-1">لا توجد مشاريع مخصصة لهذا المهندس</p>
            </div>
          </v-card-text>
        </v-card>
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeDialog" class="cancel-btn">
          إغلاق
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
  engineer: {
    type: Object,
    default: null
  },
  availableProjects: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:modelValue', 'add-project', 'remove-project', 'close'])

// Dialog model
const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// Local state
const selectedProject = ref(null)

// Computed: engineer's current projects
const engineerProjects = computed(() => {
  return props.engineer?.projects || []
})

// Computed: filter out already assigned projects
const filteredAvailableProjects = computed(() => {
  const assignedIds = engineerProjects.value.map(p => p.id)
  return props.availableProjects.filter(p => !assignedIds.includes(p.id))
})

// Reset selected project when dialog opens/closes
watch(() => props.modelValue, (newValue) => {
  if (newValue) {
    selectedProject.value = null
  }
})

// Status helpers
const getStatusColor = (status) => {
  const colors = {
    active: 'success',
    completed: 'info',
    pending: 'warning',
    cancelled: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusLabel = (status) => {
  const labels = {
    active: 'نشط',
    completed: 'مكتمل',
    pending: 'في الانتظار',
    cancelled: 'ملغي'
  }
  return labels[status] || status
}

// Add project
const addProject = () => {
  if (selectedProject.value && props.engineer) {
    emit('add-project', {
      engineerId: props.engineer.id,
      project: selectedProject.value
    })
    selectedProject.value = null
  }
}

// Remove project
const removeProject = (projectId) => {
  if (props.engineer) {
    emit('remove-project', {
      engineerId: props.engineer.id,
      projectId
    })
  }
}

// Close dialog
const closeDialog = () => {
  dialogModel.value = false
  selectedProject.value = null
  emit('close')
}
</script>

<style scoped>
/* Component-specific overrides - base styles in form-dialog.css */
.add-project-section {
  border-radius: var(--radius-2xl) !important;
  background: linear-gradient(145deg, var(--color-green-50) 0%, var(--color-green-100) 100%) !important;
  border-color: var(--color-green-300) !important;
}

.add-project-btn {
  height: var(--space-14) !important;
  border-radius: var(--radius-xl) !important;
  font-weight: 600 !important;
}

.projects-list-card {
  border-radius: var(--radius-2xl) !important;
}

.project-item {
  border-bottom: var(--space-px) solid var(--border-light);
}

.project-item:last-child {
  border-bottom: none;
}
</style>
