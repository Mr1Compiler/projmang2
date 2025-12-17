<template>
  <v-dialog v-model="dialogModel" max-width="700" persistent>
    <v-card v-if="engineer" class="image-style-dialog">
      <!-- Header Section -->
      <div class="dialog-header" style="background: var(--gradient-info) !important;">
        <div class="header-content">
          <div class="header-left">
            <v-avatar size="40" color="white" variant="tonal" class="me-3">
              <v-img v-if="engineer.avatar" :src="engineer.avatar" />
              <v-icon v-else color="info">mdi-account</v-icon>
            </v-avatar>
            <div>
              <span class="header-title">{{ engineer.name }}</span>
              <p class="text-subtitle-2 mb-0" style="color: rgba(255,255,255,0.8);">{{ engineer.email }}</p>
            </div>
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
        <v-row>
          <!-- معلومات أساسية -->
          <v-col cols="12" md="6">
            <v-card variant="outlined" class="pa-4 info-card">
              <v-card-title class="text-h6 pa-0 mb-3">
                <v-icon class="me-2 icon-drop-shadow-info" color="info">mdi-information</v-icon>
                <span class="section-header-info">المعلومات الأساسية</span>
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
                <v-icon class="me-2 icon-drop-shadow-warning" color="warning">mdi-star</v-icon>
                <span class="section-header-warning">التقييم والحالة</span>
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
                  <span class="text-h6 ms-2 rating-value">{{ engineer.rating || 0 }}</span>
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
                <v-icon class="me-2 icon-drop-shadow-success" color="success">mdi-cog</v-icon>
                <span class="section-header-success">المهارات</span>
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
      </div>

      <!-- Dialog Actions -->
      <div class="dialog-actions">
        <v-btn color="grey" variant="text" @click="closeDialog" class="cancel-btn">
          إغلاق
        </v-btn>
        <v-btn v-if="canEdit" color="primary" variant="elevated" @click="$emit('edit', engineer)" class="save-btn">
          <v-icon class="me-2">mdi-pencil</v-icon>
          تعديل
        </v-btn>
      </div>
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
/* Component-specific overrides - base styles in form-dialog.css */
.info-card {
  border-radius: var(--radius-2xl) !important;
  height: 100%;
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: var(--space-3);
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-label {
  font-weight: 700;
  color: var(--color-slate-800);
  font-size: var(--font-size-base-minus);
}

.info-value {
  color: var(--color-slate-600);
  font-size: var(--font-size-base-minus);
}

.info-value.text-primary {
  color: var(--color-blue-600) !important;
  font-weight: 600;
}
</style>
