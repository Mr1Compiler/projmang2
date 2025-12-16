<template>
  <v-card class="data-table-card" elevation="2">
    <v-card-title class="table-title indigo-title">
      <span class="title-text">المشاريع</span>
      <v-chip class="ms-2" size="small" color="white" variant="flat">
        {{ items.length }}
      </v-chip>
    </v-card-title>

    <v-data-table
      :headers="tableHeaders"
      :items="items"
      :search="searchQuery"
      class="project-table"
      :items-per-page="10"
      :loading="loading"
      hover
      no-data-text="لا توجد بيانات"
    >
      <!-- Serial Number Column -->
      <template #item.serial="{ index }">
        <span class="serial-number">{{ index + 1 }}</span>
      </template>

      <!-- Project Name Column -->
      <template #item.projectName="{ item }">
        <span class="project-name">{{ item.projectName }}</span>
      </template>

      <!-- Start Date Column -->
      <template #item.startDate="{ item }">
        <span class="date-text">{{ item.startDate }}</span>
      </template>

      <!-- End Date Column -->
      <template #item.endDate="{ item }">
        <span class="date-text">{{ item.endDate }}</span>
      </template>

      <!-- Cost Column -->
      <template #item.cost="{ item }">
        <span class="cost-text">{{ item.cost }}</span>
      </template>

      <!-- Work Location Column -->
      <template #item.workLocation="{ item }">
        <span class="location-text">{{ item.workLocation }}</span>
      </template>

      <!-- Notes Column -->
      <template #item.notes="{ item }">
        <span class="notes-text">{{ item.notes || 'لايوجد' }}</span>
      </template>

      <!-- Actions Column -->
      <template #item.actions="{ item }">
        <div class="action-buttons">
          <v-btn
            size="small"
            color="primary"
            variant="text"
            @click="$emit('view', item)"
            icon
            class="action-btn details-btn"
            title="عرض التفاصيل"
          >
            <v-icon size="16">mdi-eye</v-icon>
          </v-btn>
          <v-btn
            v-if="canEdit"
            size="small"
            color="warning"
            variant="text"
            @click="$emit('edit', item)"
            icon
            class="action-btn"
            title="تعديل"
          >
            <v-icon size="16">mdi-pencil</v-icon>
          </v-btn>
          <v-btn
            v-if="canDelete"
            size="small"
            color="error"
            variant="text"
            @click="$emit('delete', item)"
            icon
            class="action-btn"
            title="حذف"
          >
            <v-icon size="16">mdi-delete</v-icon>
          </v-btn>
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
defineProps({
  items: {
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
  canEdit: {
    type: Boolean,
    default: true
  },
  canDelete: {
    type: Boolean,
    default: true
  }
})

defineEmits(['view', 'edit', 'delete'])

const tableHeaders = [
  { title: 'التسلسل', key: 'serial', sortable: false, align: 'center' },
  { title: 'اسم المشروع', key: 'projectName', sortable: true, align: 'right' },
  { title: 'تاريخ بدء المشروع', key: 'startDate', sortable: true, align: 'center' },
  { title: 'تاريخ انتهاء المشروع', key: 'endDate', sortable: true, align: 'center' },
  { title: 'التكلفة', key: 'cost', sortable: true, align: 'center' },
  { title: 'مكان العمل', key: 'workLocation', sortable: true, align: 'center' },
  { title: 'الملاحظات', key: 'notes', sortable: false, align: 'right' },
  { title: 'الاجراءات', key: 'actions', sortable: false, align: 'center' }
]
</script>

<style scoped>
.data-table-card {
  border-radius: 16px !important;
  overflow: hidden;
}

.table-title {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: white !important;
  padding: 16px 24px !important;
  display: flex;
  align-items: center;
}

.title-text {
  font-weight: 700;
  font-size: var(--font-size-base-plus);
}

.serial-number {
  font-weight: 600;
  color: #374151;
}

.project-name {
  font-weight: 600;
  color: #1e40af;
}

.date-text {
  color: #4b5563;
}

.cost-text {
  font-weight: 600;
  color: #059669;
}

.location-text {
  color: #4b5563;
}

.notes-text {
  color: #6b7280;
  font-size: var(--font-size-sm);
}

.action-buttons {
  display: flex;
  gap: 4px;
  justify-content: center;
}

.action-btn {
  min-width: 32px !important;
  height: 32px !important;
}
</style>
