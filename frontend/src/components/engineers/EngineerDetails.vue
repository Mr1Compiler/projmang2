<template>
  <v-dialog v-model="dialogModel" max-width="700" class="engineer-details-dialog">
    <v-card v-if="engineer">
      <v-card-title class="d-flex align-center justify-space-between details-header">
        <div class="d-flex align-center">
          <v-avatar size="40" color="primary" variant="tonal" class="me-3">
            <v-img v-if="engineer.avatar" :src="engineer.avatar" />
            <v-icon v-else>mdi-account</v-icon>
          </v-avatar>
          <div>
            <h2 class="text-h5 font-weight-bold" style="color: white; text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);">{{ engineer.name }}</h2>
            <p class="text-subtitle-2 mb-0" style="color: rgba(255, 255, 255, 0.9); font-weight: 500;">{{ engineer.email }}</p>
          </div>
        </div>
        <v-btn icon="mdi-close" variant="text" @click="closeDialog" style="color: white; background: rgba(255, 255, 255, 0.1); border-radius: 8px;" />
      </v-card-title>

      <v-divider />

      <v-card-text class="pa-6">
        <v-row>
          <!-- معلومات أساسية -->
          <v-col cols="12" md="6">
            <v-card variant="outlined" class="pa-4 info-card">
              <v-card-title class="text-h6 pa-0 mb-3">
                <v-icon class="me-2" color="info" style="filter: drop-shadow(0 2px 4px rgba(59, 130, 246, 0.3));">mdi-information</v-icon>
                <span style="color: #1e40af !important; font-weight: 700 !important; font-size: var(--font-size-base-plus) !important;">المعلومات الأساسية</span>
              </v-card-title>
              <div class="info-list">
                <div class="info-item">
                  <span class="info-label">الاسم:</span>
                  <span class="info-value text-primary">{{ engineer.name }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">البريد الإلكتروني:</span>
                  <span class="info-value">{{ engineer.email }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">رقم الهاتف:</span>
                  <span class="info-value">{{ engineer.phone || 'غير محدد' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">التخصص:</span>
                  <span class="info-value">{{ engineer.specialization || 'غير محدد' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">سنوات الخبرة:</span>
                  <span class="info-value">{{ engineer.experience || 'غير محدد' }}</span>
                </div>
              </div>
            </v-card>
          </v-col>

          <!-- التقييم والحالة -->
          <v-col cols="12" md="6">
            <v-card variant="outlined" class="pa-4 info-card">
              <v-card-title class="text-h6 pa-0 mb-3">
                <v-icon class="me-2" color="warning" style="filter: drop-shadow(0 2px 4px rgba(245, 158, 11, 0.3));">mdi-star</v-icon>
                <span style="color: #b45309 !important; font-weight: 700 !important; font-size: var(--font-size-base-plus) !important;">التقييم والحالة</span>
              </v-card-title>
              <div class="info-list">
                <div class="info-item d-flex align-center">
                  <span class="info-label me-3">التقييم:</span>
                  <v-rating
                    :model-value="engineer.rating || 0"
                    readonly
                    size="small"
                    color="warning"
                    density="compact"
                  />
                  <span class="text-h6 ms-2" style="color: #f59e0b !important; font-weight: 700 !important;">{{ engineer.rating || 0 }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">الحالة:</span>
                  <v-chip
                    :color="engineer.status === 'active' ? 'success' : 'error'"
                    size="small"
                    variant="tonal"
                  >
                    {{ engineer.status === 'active' ? 'نشط' : 'غير نشط' }}
                  </v-chip>
                </div>
              </div>
            </v-card>
          </v-col>

          <!-- المهارات -->
          <v-col cols="12">
            <v-card variant="outlined" class="pa-4 info-card">
              <v-card-title class="text-h6 pa-0 mb-3">
                <v-icon class="me-2" color="success" style="filter: drop-shadow(0 2px 4px rgba(16, 185, 129, 0.3));">mdi-cog</v-icon>
                <span style="color: #047857 !important; font-weight: 700 !important; font-size: var(--font-size-base-plus) !important;">المهارات</span>
              </v-card-title>
              <div v-if="engineer.skills && engineer.skills.length > 0" class="d-flex flex-wrap gap-2">
                <v-chip
                  v-for="skill in engineer.skills"
                  :key="skill"
                  color="primary"
                  variant="tonal"
                  size="small"
                >
                  {{ skill }}
                </v-chip>
              </div>
              <div v-else class="text-medium-emphasis">
                لا توجد مهارات مسجلة
              </div>
            </v-card>
          </v-col>
        </v-row>
      </v-card-text>

      <v-divider />

      <v-card-actions class="pa-4">
        <v-spacer />
        <v-btn v-if="canEdit" color="primary" variant="elevated" @click="$emit('edit', engineer)">
          <v-icon class="me-2">mdi-pencil</v-icon>
          تعديل
        </v-btn>
        <v-btn color="grey" variant="text" @click="closeDialog">
          إغلاق
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  engineer: {
    type: Object,
    default: null
  },
  canEdit: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue', 'edit', 'close'])

// Dialog model
const dialogModel = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

// Close dialog
const closeDialog = () => {
  dialogModel.value = false
  emit('close')
}
</script>

<style scoped>
.engineer-details-dialog .v-card {
  border-radius: 20px !important;
  overflow: hidden !important;
}

.details-header {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  padding: 20px 24px !important;
}

.info-card {
  border-radius: 16px !important;
  height: 100%;
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-label {
  font-weight: 700;
  color: #1e293b;
  font-size: var(--font-size-base-minus);
}

.info-value {
  color: #334155;
  font-size: var(--font-size-base-minus);
}

.info-value.text-primary {
  color: #2563eb !important;
  font-weight: 600;
}
</style>
