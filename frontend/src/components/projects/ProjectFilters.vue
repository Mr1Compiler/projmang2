<template>
  <v-card class="filters-card mb-6" elevation="2">
    <v-card-text class="pa-4">
      <v-row align="center" dense>
        <!-- Search Input -->
        <v-col cols="12" sm="6" md="4">
          <v-text-field
            v-model="searchQuery"
            prepend-inner-icon="mdi-magnify"
            label="بحث في المشاريع"
            variant="outlined"
            density="comfortable"
            hide-details
            clearable
            class="search-input"
            @update:model-value="emitFilters"
          />
        </v-col>

        <!-- Status Filter -->
        <v-col cols="12" sm="6" md="3">
          <v-select
            v-model="selectedStatus"
            :items="statusOptions"
            label="الحالة"
            variant="outlined"
            density="comfortable"
            hide-details
            clearable
            class="filter-select"
            @update:model-value="emitFilters"
          />
        </v-col>

        <!-- Sort By -->
        <v-col cols="12" sm="6" md="3">
          <v-select
            v-model="sortBy"
            :items="sortOptions"
            label="ترتيب حسب"
            variant="outlined"
            density="comfortable"
            hide-details
            class="filter-select"
            @update:model-value="emitFilters"
          />
        </v-col>

        <!-- Add Button -->
        <v-col cols="12" sm="6" md="2" class="d-flex justify-end">
          <v-btn
            v-if="canAdd"
            color="primary"
            variant="elevated"
            prepend-icon="mdi-plus"
            @click="$emit('add-project')"
            class="add-btn"
          >
            إضافة مشروع
          </v-btn>
        </v-col>
      </v-row>

      <!-- Active Filters Chips -->
      <v-row v-if="hasActiveFilters" class="mt-3">
        <v-col cols="12">
          <div class="d-flex flex-wrap gap-2 align-center">
            <span class="text-caption text-grey">الفلاتر النشطة:</span>

            <v-chip
              v-if="searchQuery"
              size="small"
              closable
              color="primary"
              variant="tonal"
              @click:close="clearSearch"
            >
              بحث: {{ searchQuery }}
            </v-chip>

            <v-chip
              v-if="selectedStatus"
              size="small"
              closable
              color="success"
              variant="tonal"
              @click:close="clearStatus"
            >
              الحالة: {{ getStatusText(selectedStatus) }}
            </v-chip>

            <v-btn
              v-if="hasActiveFilters"
              size="small"
              variant="text"
              color="error"
              @click="clearAllFilters"
            >
              مسح الكل
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  canAdd: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['filter-change', 'add-project'])

// Filter state
const searchQuery = ref('')
const selectedStatus = ref(null)
const sortBy = ref('createdAt')

// Options
const statusOptions = [
  { title: 'الكل', value: null },
  { title: 'نشط', value: 'active' },
  { title: 'في الانتظار', value: 'pending' },
  { title: 'مكتمل', value: 'completed' },
  { title: 'ملغي', value: 'cancelled' }
]

const sortOptions = [
  { title: 'الأحدث', value: 'createdAt' },
  { title: 'الاسم', value: 'name' },
  { title: 'التكلفة', value: 'initialCost' },
  { title: 'التقدم', value: 'progress' }
]

// Computed
const hasActiveFilters = computed(() => {
  return searchQuery.value || selectedStatus.value
})

// Methods
const getStatusText = (status) => {
  const option = statusOptions.find(opt => opt.value === status)
  return option ? option.title : status
}

const emitFilters = () => {
  emit('filter-change', {
    search: searchQuery.value,
    status: selectedStatus.value,
    sortBy: sortBy.value
  })
}

const clearSearch = () => {
  searchQuery.value = ''
  emitFilters()
}

const clearStatus = () => {
  selectedStatus.value = null
  emitFilters()
}

const clearAllFilters = () => {
  searchQuery.value = ''
  selectedStatus.value = null
  sortBy.value = 'createdAt'
  emitFilters()
}
</script>

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for ProjectFilters */
/* All shared styles (.filters-card, .search-input, .filter-select, .add-btn)
   are now in projects.css */

@media (max-width: 600px) {
  .add-btn {
    width: 100%;
  }
}
</style>
