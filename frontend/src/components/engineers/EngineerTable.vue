<template>
  <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
    <v-card-title class="d-flex align-center justify-space-between table-header">
      <div class="d-flex align-center">
        <v-icon class="me-2" style="color: white !important;">mdi-account-group</v-icon>
        <span class="text-h6 font-weight-bold" style="color: white !important; font-family: 'Arial', 'Helvetica', sans-serif;">قائمة المهندسين</span>
        <v-chip class="ms-3" style="background-color: rgba(255, 255, 255, 0.3) !important; color: white !important; border: 1px solid rgba(255, 255, 255, 0.5);" size="small">{{ engineers.length }}</v-chip>
      </div>
      <v-btn
        v-if="canWrite"
        class="add-button btn-glow light-sweep smooth-transition"
        @click="$emit('add')"
        elevation="2"
        color="primary"
        style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;"
      >
        <v-icon class="me-2 icon-glow">mdi-plus</v-icon>
        إضافة حساب جديد
      </v-btn>
    </v-card-title>

    <!-- Search Field -->
    <v-card-text class="pb-0">
      <v-text-field
        :model-value="searchQuery"
        @update:model-value="$emit('update:searchQuery', $event)"
        prepend-inner-icon="mdi-magnify"
        label="البحث في المهندسين..."
        variant="outlined"
        density="comfortable"
        clearable
        hide-details
        class="mb-4 search-field-enhanced"
        style="font-size: var(--font-size-base-plus) !important; font-weight: 500 !important;"
        placeholder="ابحث بالاسم أو البريد الإلكتروني أو التخصص..."
        color="primary"
        bg-color="white"
      />
    </v-card-text>

    <v-data-table
      :headers="tableHeaders"
      :items="engineers"
      :search="searchQuery"
      class="elevation-0"
      no-data-text="لا يوجد مهندسين"
      loading-text="جاري التحميل..."
      :loading="loading"
      items-per-page="10"
      show-current-page
    >
      <!-- Avatar Column -->
      <template v-slot:item.avatar="{ item }">
        <v-avatar size="40" class="me-3">
          <v-img v-if="item.avatar" :src="item.avatar" />
          <v-icon v-else color="primary">mdi-account</v-icon>
        </v-avatar>
      </template>

      <!-- Name Column -->
      <template v-slot:item.name="{ item }">
        <div>
          <div class="font-weight-bold text-primary" style="font-size: var(--font-size-base) !important; color: #1976d2 !important;">{{ item.name }}</div>
          <div class="text-caption text-medium-emphasis" style="font-size: var(--font-size-sm) !important; color: #666 !important;">{{ item.email }}</div>
        </div>
      </template>

      <!-- Rating Column -->
      <template v-slot:item.rating="{ item }">
        <div class="d-flex align-center">
          <v-rating
            :model-value="item.rating"
            readonly
            size="small"
            color="warning"
            density="compact"
          />
          <span class="text-caption ms-1">{{ item.rating }}</span>
        </div>
      </template>

      <!-- Skills Column -->
      <template v-slot:item.skills="{ item }">
        <div class="d-flex flex-wrap gap-1">
          <v-chip
            v-for="skill in (item.skills || []).slice(0, 2)"
            :key="skill"
            size="small"
            color="primary"
            variant="tonal"
          >
            {{ skill }}
          </v-chip>
          <v-chip
            v-if="(item.skills || []).length > 2"
            size="small"
            color="grey"
            variant="tonal"
          >
            +{{ item.skills.length - 2 }}
          </v-chip>
        </div>
      </template>

      <!-- Status Column -->
      <template v-slot:item.status="{ item }">
        <v-chip
          :color="item.status === 'active' ? 'success' : 'error'"
          size="small"
          variant="tonal"
        >
          {{ item.status === 'active' ? 'نشط' : 'غير نشط' }}
        </v-chip>
      </template>

      <!-- Actions Column -->
      <template v-slot:item.actions="{ item }">
        <div class="d-flex gap-2">
          <v-btn
            icon="mdi-eye"
            size="small"
            variant="text"
            color="primary"
            @click="$emit('view', item)"
          />
          <v-btn
            v-if="canWrite"
            icon="mdi-pencil"
            size="small"
            variant="text"
            color="warning"
            @click="$emit('edit', item)"
          />
          <v-btn
            v-if="canWrite"
            icon="mdi-briefcase-plus"
            size="small"
            variant="text"
            color="success"
            @click="$emit('manage-projects', item)"
            title="إضافة مشروع"
          />
          <v-btn
            v-if="canDelete"
            icon="mdi-delete"
            size="small"
            variant="text"
            color="error"
            @click="$emit('delete', item)"
          />
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
const props = defineProps({
  engineers: {
    type: Array,
    default: () => []
  },
  searchQuery: {
    type: String,
    default: ''
  },
  loading: {
    type: Boolean,
    default: false
  },
  canWrite: {
    type: Boolean,
    default: true
  },
  canDelete: {
    type: Boolean,
    default: true
  }
})

defineEmits(['add', 'view', 'edit', 'delete', 'manage-projects', 'update:searchQuery'])

const tableHeaders = [
  { title: 'الصورة', key: 'avatar', sortable: false },
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'التقييم', key: 'rating', sortable: true },
  { title: 'المهارات', key: 'skills', sortable: false },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]
</script>

<style scoped>
.data-table-card {
  border-radius: 24px !important;
  overflow: hidden;
}

.table-header {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  padding: 16px 24px !important;
}

.search-field-enhanced {
  max-width: 500px;
}

.add-button {
  border-radius: 12px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  letter-spacing: 0 !important;
}

.btn-glow:hover {
  box-shadow: 0 8px 25px rgba(59, 130, 246, 0.4) !important;
}
</style>
