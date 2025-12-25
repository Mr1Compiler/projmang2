<template>
  <div class="data-page">
    <v-container fluid class="fullscreen-content" style="padding: 0 20px;">
      <!-- Page Header -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="header-text">
              <h1 class="main-title">إدارة التصنيفات</h1>
              <p class="subtitle">تنظيم وإدارة تصنيفات المشاريع والمهام</p>
          </div>
            <div class="engineer-emoji">
              <v-icon size="32" color="white">mdi-tag-multiple</v-icon>
            </div>
          </div>
        </div>
      </div>

      <!-- Statistics Cards -->
      <v-row class="mb-6 stats-row" no-gutters>
        <v-col cols="12" sm="12" md="12" lg="12" xl="12" class="pa-2">
          <div class="stats-cards-container">
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-tag-multiple</v-icon>
            </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ totalCategories }}</h3>
                  <p class="stat-label">إجمالي التصنيفات</p>
                </div>
              </div>
          </v-card>
            <v-card class="modern-stat-card stat-card-success" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
            </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ activeCategories }}</h3>
                  <p class="stat-label">تصنيفات نشطة</p>
                </div>
              </div>
          </v-card>
            <v-card class="modern-stat-card stat-card-warning" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-pause-circle</v-icon>
            </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ inactiveCategories }}</h3>
                  <p class="stat-label">تصنيفات معطلة</p>
                </div>
              </div>
          </v-card>
            <v-card class="modern-stat-card stat-card-info" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-folder-multiple</v-icon>
            </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ totalProjects }}</h3>
                  <p class="stat-label">مشاريع مرتبطة</p>
                </div>
              </div>
          </v-card>
          </div>
        </v-col>
      </v-row>


      <!-- Categories Table -->
      <v-card class="data-table-card" elevation="3">
        <v-card-title class="table-header-enhanced">
          <div class="d-flex align-center">
            <v-icon class="me-2 table-header-icon" color="white" size="18">mdi-tag-multiple</v-icon>
            <div class="table-header-content">
              <h2 class="text-h5 font-weight-bold table-title">قائمة التصنيفات</h2>
            </div>
            <v-spacer />
            <v-btn
              v-if="canCreate"
              class="add-button btn-glow light-sweep smooth-transition"
              @click="openAddCategoryDialog"
              elevation="2"
              color="primary"
              size="small"
              style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
            >
              <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
              إضافة تصنيف جديد
            </v-btn>
            <v-chip color="primary" variant="tonal" class="count-chip ms-3" size="x-small" style="font-size: 0.75rem !important; height: 20px !important;">
              {{ filteredCategories.length }} من {{ totalCategories }}
            </v-chip>
          </div>
        </v-card-title>

        <v-data-table
          :headers="headers"
          :items="filteredCategories"
          :search="searchQuery"
          class="categories-table"
          no-data-text="لا توجد تصنيفات"
          loading-text="جاري التحميل..."
          :items-per-page="-1"
          hide-default-footer
        >
          <!-- Sequence Column -->
          <template v-slot:item.sequence="{ index }">
            <div class="sequence-number">
              {{ (currentPage - 1) * itemsPerPage + index + 1 }}
            </div>
          </template>

          <!-- Category Name Column -->
          <template v-slot:item.name="{ item }">
            <div class="category-name-text">{{ item.name }}</div>
          </template>

          <!-- Status Column -->
          <template v-slot:item.status="{ item }">
            <v-chip
              :color="item.status === 'active' ? 'success' : 'warning'"
              size="x-small"
              variant="tonal"
              class="status-chip"
              :class="item.status === 'active' ? 'active-chip' : 'inactive-chip'"
            >
              <v-icon :icon="item.status === 'active' ? 'mdi-check-circle' : 'mdi-pause-circle'" size="x-small" class="me-1" />
              {{ item.status === 'active' ? 'نشط' : 'معطل' }}
            </v-chip>
          </template>

          <!-- Description Column -->
          <template v-slot:item.description="{ item }">
            <div class="description-display">
              <v-icon class="me-1" color="grey" size="x-small">mdi-text</v-icon>
              <span class="description-text">{{ item.description || 'لا يوجد وصف' }}</span>
            </div>
          </template>

          <!-- Projects Count Column -->
          <template v-slot:item.projectsCount="{ item }">
            <v-chip
              color="info"
              size="x-small"
              variant="tonal"
              class="projects-chip"
            >
              <v-icon icon="mdi-folder-multiple" size="x-small" class="me-1" />
              {{ item.projectsCount }} مشروع
            </v-chip>
          </template>

          <!-- Creation Date Column -->
          <template v-slot:item.createdAt="{ item }">
            <div class="date-display">
              <v-icon class="me-1" color="success" size="x-small">mdi-calendar</v-icon>
              <span class="date-text">{{ item.createdAt || 'غير محدد' }}</span>
            </div>
          </template>

          <!-- Actions Column -->
          <template v-slot:item.actions="{ item }">
            <div class="d-flex align-center gap-2 justify-center">
              <v-btn
                icon="mdi-view-list"
                variant="text"
                size="x-small"
                color="primary"
                @click="showSubCategories(item)"
                class="action-btn"
                title="عرض الأصناف الثانوية"
              />
              <v-btn
                v-if="canUpdate"
                icon="mdi-pencil"
                variant="text"
                size="x-small"
                color="warning"
                @click="editCategory(item)"
                class="action-btn"
                title="تعديل الصنف"
              />
            </div>
          </template>
        </v-data-table>

        <!-- Pagination -->
        <div class="d-flex justify-center pa-4" v-if="totalPages > 0">
          <v-pagination
            v-model="currentPage"
            :length="totalPages"
            :total-visible="7"
            @update:model-value="onPageChange"
            rounded="circle"
            density="comfortable"
            active-color="primary"
          />
        </div>
      </v-card>

      <!-- Add/Edit Category Dialog -->
      <v-dialog v-model="dialog" :max-width="isEditing ? 700 : 500" class="simple-category-dialog">
        <v-card class="simple-dialog-card">
          <!-- Header -->
          <v-card-title class="simple-dialog-header">
            <div class="d-flex align-center justify-space-between w-100">
              <h3 class="text-h6 font-weight-bold text-white mb-0">
                {{ isEditing ? 'تعديل التصنيف' : 'إضافة صنف جديد' }}
              </h3>
              <v-btn
                icon
                color="white"
                variant="text"
                size="small"
                @click="closeDialog"
                class="close-btn"
              >
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </div>
          </v-card-title>

          <!-- Body -->
          <v-card-text class="simple-dialog-body pa-6">
            <v-form ref="form" v-model="formValid">
              <div class="form-field-simple">
                <label class="field-label">اسم القسم :</label>
                <div class="custom-input-container">
                  <input
                    v-model="categoryForm.name"
                    type="text"
                    class="custom-input-field"
                    placeholder="أدخل اسم القسم"
                  />
                </div>
              </div>
            </v-form>

            <!-- Subcategories Section (only when editing) -->
            <div v-if="isEditing && categoryForm.id" class="subcategories-section mt-6">
              <div class="d-flex align-center justify-space-between mb-3">
                <h4 class="text-subtitle-1 font-weight-bold" style="color: #333;">
                  <v-icon class="me-2" size="small" color="primary">mdi-folder-multiple</v-icon>
                  الأصناف الفرعية
                </h4>
                <v-btn
                  size="small"
                  color="primary"
                  variant="tonal"
                  @click="openAddSubCategoryFromEdit"
                >
                  <v-icon class="me-1" size="small">mdi-plus</v-icon>
                  إضافة فرعي
                </v-btn>
              </div>

              <!-- Subcategories List -->
              <div class="subcategories-list" style="max-height: 250px; overflow-y: auto;">
                <div
                  v-for="(subCat, index) in getSubCategories(categoryForm.id)"
                  :key="subCat.id"
                  class="subcategory-item d-flex align-center justify-space-between pa-3 mb-2"
                  style="background: #f8f9fa; border-radius: 8px; border: 1px solid #e9ecef;"
                >
                  <div class="d-flex align-center">
                    <span class="subcategory-number me-3" style="background: #e3f2fd; color: #1976d2; padding: 4px 10px; border-radius: 50%; font-weight: bold; font-size: 0.8rem;">
                      {{ index + 1 }}
                    </span>
                    <div>
                      <div class="font-weight-medium" style="color: #333;">{{ subCat.name }}</div>
                      <div class="text-caption" style="color: #666;">نسبة الإنجاز: {{ subCat.percentage || 0 }}%</div>
                    </div>
                  </div>
                  <div class="d-flex">
                    <v-btn
                      v-if="canUpdate"
                      icon="mdi-pencil"
                      size="x-small"
                      color="primary"
                      variant="text"
                      @click="editSubCategoryFromEdit(subCat)"
                    />
                    <v-btn
                      v-if="canDelete"
                      icon="mdi-delete"
                      size="x-small"
                      color="error"
                      variant="text"
                      @click="deleteSubCategoryHandler(subCat)"
                    />
                  </div>
                </div>

                <!-- Empty state -->
                <div
                  v-if="getSubCategories(categoryForm.id).length === 0"
                  class="text-center pa-4"
                  style="color: #999;"
                >
                  <v-icon size="40" color="grey-lighten-1">mdi-folder-open-outline</v-icon>
                  <p class="mt-2 mb-0">لا توجد أصناف فرعية</p>
                </div>
              </div>
            </div>
          </v-card-text>

          <!-- Actions -->
          <v-card-actions class="simple-dialog-actions pa-6 pt-0">
            <v-btn
              color="grey-darken-1"
              variant="text"
              @click="closeDialog"
              class="save-btn-simple"
            >
              إلغاء
            </v-btn>
            <v-spacer />
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveCategory"
              :disabled="!categoryForm.name"
              :loading="saving"
              class="save-btn-simple"
            >
              {{ isEditing ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Sub Categories Dialog -->
      <v-dialog v-model="subCategoriesDialog" max-width="600" class="sub-categories-dialog">
        <v-card class="sub-categories-card">
          <!-- Header -->
          <v-card-title class="sub-categories-header">
            <div class="d-flex align-center justify-space-between w-100">
              <h3 class="text-h6 font-weight-bold text-white mb-0">
                الأصناف الثانوية - {{ selectedCategory?.name }}
              </h3>
              <v-btn
                icon
                color="white"
                variant="text"
                size="small"
                @click="subCategoriesDialog = false"
                class="close-btn"
              >
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </div>
          </v-card-title>

          <!-- Body -->
          <v-card-text class="sub-categories-body pa-6">
            <div class="sub-categories-list">
              <div 
                v-for="(subCategory, index) in getSubCategories(selectedCategory?.id)" 
                :key="subCategory.id"
                class="sub-category-item"
              >
                <div class="sub-category-number">{{ index + 1 }}</div>
                <div class="sub-category-info">
                  <div class="sub-category-name">{{ subCategory.name }}</div>
                  <div class="sub-category-progress">
                    <span class="progress-label">نسبة الإنجاز:</span>
                    <span class="progress-value">{{ subCategory.percentage || 0 }}%</span>
                  </div>
                </div>
                <div class="sub-category-actions">
                  <v-btn
                    v-if="canUpdate"
                    icon="mdi-pencil"
                    variant="text"
                    size="small"
                    color="primary"
                    @click="editSubCategory(subCategory)"
                    class="sub-action-btn"
                  />
                  <v-btn
                    v-if="canDelete"
                    icon="mdi-delete"
                    variant="text"
                    size="small"
                    color="error"
                    @click="deleteSubCategoryHandler(subCategory)"
                    class="sub-action-btn"
                  />
                </div>
              </div>
            </div>

            <!-- Add New Sub Category -->
            <div class="add-sub-category-section mt-6">
              <v-btn
                color="primary"
                variant="elevated"
                @click="openAddSubCategoryDialog"
                class="add-sub-btn"
                block
              >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة صنف ثانوي جديد
              </v-btn>
            </div>
          </v-card-text>
        </v-card>
      </v-dialog>

      <!-- Add Sub Category Dialog -->
      <v-dialog v-model="addSubCategoryDialog" max-width="500" persistent>
        <v-card class="simple-dialog-card">
          <!-- Header -->
          <v-card-title class="simple-dialog-header">
            <div class="d-flex align-center justify-space-between w-100">
              <h3 class="text-h6 font-weight-bold text-white mb-0">
                {{ isEditingSubCategory ? 'تعديل صنف ثانوي' : 'إضافة صنف ثانوي جديد' }}
              </h3>
              <v-btn
                icon
                color="white"
                variant="text"
                size="small"
                @click="closeAddSubCategoryDialog"
                class="close-btn"
              >
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </div>
          </v-card-title>

          <!-- Body -->
          <v-card-text class="simple-dialog-body pa-6">
            <v-form ref="subCatFormRef">
              <div class="form-field-simple mb-4">
                <label class="field-label">اسم الصنف :</label>
                <div class="custom-input-container">
                  <input
                    v-model="subCategoryForm.name"
                    type="text"
                    class="custom-input-field"
                    placeholder="أدخل اسم الصنف الثانوي"
                  />
                </div>
              </div>
              <div class="form-field-simple">
                <label class="field-label">نسبة الإنجاز (%) :</label>
                <div class="custom-input-container">
                  <input
                    v-model="subCategoryForm.percentage"
                    type="number"
                    class="custom-input-field"
                    placeholder="0"
                    min="0"
                    max="100"
                  />
                </div>
                <!-- Progress Bar Display -->
                <div class="progress-display-container mt-3">
                  <div class="progress-bar-wrapper">
                    <div
                      class="progress-bar-fill"
                      :style="{ width: (subCategoryForm.percentage || 0) + '%' }"
                    ></div>
                  </div>
                  <div class="progress-text-display">
                    <span class="progress-percentage">{{ subCategoryForm.percentage || 0 }}%</span>
                  </div>
                </div>
              </div>
            </v-form>
          </v-card-text>

          <!-- Actions -->
          <v-card-actions class="simple-dialog-actions pa-6 pt-0">
            <v-spacer />
            <v-btn
              color="grey-darken-1"
              variant="text"
              @click="closeAddSubCategoryDialog"
              class="save-btn-simple me-2"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveSubCategory"
              :disabled="!subCategoryForm.name"
              :loading="saving"
              class="save-btn-simple"
            >
              {{ isEditingSubCategory ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Category Details Dialog -->
      <v-dialog v-model="detailsDialog" max-width="700" class="category-details-dialog">
        <v-card class="details-dialog-card">
          <v-card-title class="details-header-enhanced">
            <div class="d-flex align-center">
              <v-avatar size="48" class="me-3 category-avatar">
                <div 
                  class="color-circle-large" 
                  :style="{ backgroundColor: selectedCategory?.color }"
                ></div>
              </v-avatar>
              <div class="category-title-content">
                <h2 class="text-h4 font-weight-bold category-title">{{ selectedCategory?.name }}</h2>
                <p class="text-subtitle-1 category-subtitle mb-0">{{ selectedCategory?.colorName }}</p>
              </div>
            </div>
            <v-btn 
              icon="mdi-close" 
              variant="text" 
              @click="detailsDialog = false" 
              class="close-btn"
            />
          </v-card-title>
          
          <v-divider />
          <v-card-text class="details-content-enhanced">
            <v-row>
              <!-- الاسم -->
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="detail-card info-card">
                  <v-card-title class="detail-card-title">
                    <v-icon class="me-2" color="primary" size="small">mdi-tag</v-icon>
                    اسم التصنيف
                  </v-card-title>
                  <v-card-text class="detail-card-content">
                    <span class="detail-value">{{ selectedCategory?.name }}</span>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- اللون -->
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="detail-card color-card">
                  <v-card-title class="detail-card-title">
                    <v-icon class="me-2" color="info" size="small">mdi-palette</v-icon>
                    اللون
                  </v-card-title>
                  <v-card-text class="detail-card-content">
                    <div class="d-flex align-center">
                      <div 
                        class="color-circle-small me-2" 
                        :style="{ backgroundColor: selectedCategory?.color }"
                      ></div>
                      <span class="detail-value">{{ selectedCategory?.colorName }}</span>
                    </div>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- الحالة -->
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="detail-card status-card">
                  <v-card-title class="detail-card-title">
                    <v-icon 
                      class="me-2" 
                      :color="selectedCategory?.isActive ? 'success' : 'warning'" 
                      size="small"
                    >
                      {{ selectedCategory?.isActive ? 'mdi-check-circle' : 'mdi-pause-circle' }}
                    </v-icon>
                    الحالة
                  </v-card-title>
                  <v-card-text class="detail-card-content">
                    <v-chip 
                      :color="selectedCategory?.isActive ? 'success' : 'warning'"
                      size="small"
                      variant="tonal"
                      class="status-chip-detail"
                    >
                      {{ selectedCategory?.isActive ? 'نشط' : 'معطل' }}
                    </v-chip>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- عدد المشاريع -->
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="detail-card projects-card">
                  <v-card-title class="detail-card-title">
                    <v-icon class="me-2" color="info" size="small">mdi-folder-multiple</v-icon>
                    المشاريع المرتبطة
                  </v-card-title>
                  <v-card-text class="detail-card-content">
                    <v-chip color="info" size="small" variant="tonal" class="projects-chip-detail">
                      {{ selectedCategory?.projectsCount }} مشروع
                    </v-chip>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- تاريخ الإنشاء -->
              <v-col cols="12">
                <v-card variant="outlined" class="detail-card date-card">
                  <v-card-title class="detail-card-title">
                    <v-icon class="me-2" color="success" size="small">mdi-calendar</v-icon>
                    تاريخ الإنشاء
                  </v-card-title>
                  <v-card-text class="detail-card-content">
                    <span class="detail-value date-value">{{ selectedCategory?.createdAt }}</span>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- الوصف -->
              <v-col cols="12" v-if="selectedCategory?.description">
                <v-card variant="outlined" class="detail-card description-card">
                  <v-card-title class="detail-card-title">
                    <v-icon class="me-2" color="primary" size="small">mdi-text</v-icon>
                    الوصف
                  </v-card-title>
                  <v-card-text class="detail-card-content">
                    <p class="description-text-enhanced">{{ selectedCategory?.description }}</p>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
          </v-card-text>
          
          <v-divider />
          <v-card-actions class="details-actions-enhanced">
            <v-spacer />
            <v-btn 
              color="primary" 
              variant="elevated" 
              size="large"
              @click="detailsDialog = false"
              class="confirm-btn"
            >
              <v-icon class="me-2">mdi-check-circle</v-icon>
              موافق
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Delete Confirmation Dialog -->
      <v-dialog v-model="deleteDialog" max-width="400">
        <v-card>
          <v-card-title class="text-h6">تأكيد الحذف</v-card-title>
          <v-card-text>
            هل أنت متأكد من حذف التصنيف "{{ selectedCategory?.name }}"؟
            <br>
            <small class="text-warning">سيتم حذف جميع المشاريع المرتبطة بهذا التصنيف!</small>
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn color="grey" @click="deleteDialog = false">إلغاء</v-btn>
            <v-btn color="error" @click="confirmDelete">حذف</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-container>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { DEFAULT_LIMIT } from '@/constants/pagination'
import {
  listCategories,
  createCategory,
  updateCategory,
  deleteCategory as apiDeleteCategory,
  getCategoryStats,
  listSubCategories,
  createSubCategory,
  updateSubCategory,
  deleteSubCategory as apiDeleteSubCategory
} from '@/api/categories'
import { usePermissions } from '@/composables/usePermissions'

// Permissions - use /categories (stored in DB), router maps to /workCategories for API
const { canCreate, canUpdate, canDelete } = usePermissions('/categories')

// State variables
const dialog = ref(false)
const deleteDialog = ref(false)
const detailsDialog = ref(false)
const subCategoriesDialog = ref(false)
const addSubCategoryDialog = ref(false)
const formValid = ref(false)
const subCategoryFormValid = ref(false)
const isEditing = ref(false)
const isEditingSubCategory = ref(false)
const editingSubCategoryId = ref(null)
const searchQuery = ref('')
const selectedStatus = ref('')
const selectedCategory = ref(null)

// Loading and error states
const loading = ref(false)
const saving = ref(false)
const error = ref('')

// Stats from backend
const stats = ref({
  total: 0,
  active: 0,
  inactive: 0,
  totalSubcategory: 0
})

// Category form - aligned with backend CreateWorkCategory DTO
const categoryForm = ref({
  id: null,
  name: '',
  description: '',
  status: 'active'
})

// Sub Category form - aligned with backend CreateWorkSubCategory DTO
const subCategoryForm = ref({
  name: '',
  description: '',
  percentage: 0
})

// Headers for data table
const headers = [
  { title: 'التسلسل', key: 'sequence', sortable: false, width: '100px', align: 'center' },
  { title: 'اسم الصنف', key: 'name', sortable: true, align: 'center' },
  { title: 'الاجراءات', key: 'actions', sortable: false, width: '150px', align: 'center' }
]

// Status options
const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'معطل', value: 'inactive' }
]

// Categories data from backend
const categories = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(DEFAULT_LIMIT)
const totalItems = ref(0)
const totalPages = ref(0)

// Sub Categories data from backend
const subCategories = ref([])

// Computed properties - use stats from backend
const totalCategories = computed(() => stats.value.total)
const activeCategories = computed(() => stats.value.active)
const inactiveCategories = computed(() => stats.value.inactive)
const totalProjects = computed(() => stats.value.totalSubcategory)

const filteredCategories = computed(() => {
  let filtered = categories.value

  if (searchQuery.value) {
    filtered = filtered.filter(category =>
      category.name?.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      (category.description && category.description.toLowerCase().includes(searchQuery.value.toLowerCase()))
    )
  }

  if (selectedStatus.value) {
    filtered = filtered.filter(category => category.status === selectedStatus.value)
  }

  return filtered
})

// ============ Load Data from Backend ============
const loadCategories = async (page = currentPage.value) => {
  loading.value = true
  error.value = ''
  try {
    const response = await listCategories({ page, limit: itemsPerPage.value })

    // Update pagination state
    categories.value = response.data || []
    totalItems.value = response.total || 0
    totalPages.value = response.totalPages || 0
    currentPage.value = response.page || page
  } catch (err) {
    console.error('Failed to load categories:', err)
    error.value = err?.message || 'فشل تحميل التصنيفات'
  } finally {
    loading.value = false
  }
}

// Handle page change
const onPageChange = (page) => {
  currentPage.value = page
  loadCategories(page)
}

const loadStats = async () => {
  try {
    const result = await getCategoryStats()
    if (result) {
      stats.value = {
        total: result.total || 0,
        active: result.active || 0,
        inactive: result.inactive || 0,
        totalSubcategory: result.totalSubcategory || 0
      }
    }
  } catch (err) {
    console.error('Failed to load stats:', err)
  }
}

const loadSubCategories = async () => {
  try {
    const result = await listSubCategories()
    subCategories.value = result.data || []
  } catch (err) {
    console.error('Failed to load subcategories:', err)
  }
}

// ============ Category Methods ============
const openAddCategoryDialog = () => {
  isEditing.value = false
  categoryForm.value = {
    id: null,
    name: '',
    description: '',
    status: 'active'
  }
  dialog.value = true
}

const editCategory = (category) => {
  isEditing.value = true
  categoryForm.value = {
    id: category.id,
    name: category.name,
    description: category.description || '',
    status: category.status || 'active'
  }
  dialog.value = true
}

const viewCategoryDetails = (category) => {
  selectedCategory.value = category
  detailsDialog.value = true
}

const closeDialog = () => {
  dialog.value = false
  categoryForm.value = {
    id: null,
    name: '',
    description: '',
    status: 'active'
  }
}

const saveCategory = async () => {
  saving.value = true
  try {
    const payload = {
      name: categoryForm.value.name,
      description: categoryForm.value.description || null,
      status: categoryForm.value.status
    }

    if (isEditing.value && categoryForm.value.id) {
      await updateCategory(categoryForm.value.id, payload)
    } else {
      await createCategory(payload)
    }

    await loadCategories()
    await loadStats()
    closeDialog()
  } catch (err) {
    console.error('Failed to save category:', err)
    error.value = err?.message || 'فشل حفظ التصنيف'
  } finally {
    saving.value = false
  }
}

const deleteCategory = (category) => {
  selectedCategory.value = category
  deleteDialog.value = true
}

const confirmDelete = async () => {
  if (!selectedCategory.value) return

  saving.value = true
  try {
    await apiDeleteCategory(selectedCategory.value.id)
    await loadCategories()
    await loadStats()
    deleteDialog.value = false
    selectedCategory.value = null
  } catch (err) {
    console.error('Failed to delete category:', err)
    error.value = err?.message || 'فشل حذف التصنيف'
  } finally {
    saving.value = false
  }
}

// ============ Sub Categories Functions ============
const showSubCategories = (category) => {
  selectedCategory.value = category
  subCategoriesDialog.value = true
}

const getSubCategories = (categoryId) => {
  return subCategories.value.filter(sub => sub.categoryId === categoryId)
}

const openAddSubCategoryDialog = () => {
  isEditingSubCategory.value = false
  editingSubCategoryId.value = null
  subCategoryForm.value = {
    name: '',
    description: '',
    percentage: 0
  }
  addSubCategoryDialog.value = true
}

// Open add subcategory from edit category dialog
const openAddSubCategoryFromEdit = () => {
  // Set selectedCategory to current category being edited
  selectedCategory.value = { id: categoryForm.value.id, name: categoryForm.value.name }
  openAddSubCategoryDialog()
}

// Edit subcategory from edit category dialog
const editSubCategoryFromEdit = (subCat) => {
  selectedCategory.value = { id: categoryForm.value.id, name: categoryForm.value.name }
  editSubCategory(subCat)
}

const closeAddSubCategoryDialog = () => {
  addSubCategoryDialog.value = false
  isEditingSubCategory.value = false
  editingSubCategoryId.value = null
  subCategoryForm.value = {
    name: '',
    description: '',
    percentage: 0
  }
}

const saveSubCategory = async () => {
  if (!subCategoryForm.value.name) return

  saving.value = true
  try {
    if (isEditingSubCategory.value && editingSubCategoryId.value) {
      // Update existing subcategory
      const payload = {
        name: subCategoryForm.value.name,
        description: subCategoryForm.value.description || null,
        percentage: parseFloat(subCategoryForm.value.percentage) || 0
      }
      await updateSubCategory(editingSubCategoryId.value, payload)
    } else {
      // Create new subcategory
      const payload = {
        categoryId: selectedCategory.value.id,
        name: subCategoryForm.value.name,
        description: subCategoryForm.value.description || null,
        percentage: parseFloat(subCategoryForm.value.percentage) || 0,
        status: 'active'
      }
      await createSubCategory(payload)
    }

    await loadSubCategories()
    await loadStats()
    closeAddSubCategoryDialog()
  } catch (err) {
    console.error('Failed to save subcategory:', err)
    error.value = err?.message || 'فشل حفظ التصنيف الفرعي'
  } finally {
    saving.value = false
  }
}

const editSubCategory = (subCategory) => {
  isEditingSubCategory.value = true
  editingSubCategoryId.value = subCategory.id
  subCategoryForm.value = {
    name: subCategory.name || '',
    description: subCategory.description || '',
    percentage: subCategory.percentage || 0
  }
  addSubCategoryDialog.value = true
}

const deleteSubCategoryHandler = async (subCategory) => {
  if (confirm(`هل أنت متأكد من حذف "${subCategory.name}"؟`)) {
    saving.value = true
    try {
      await apiDeleteSubCategory(subCategory.id)
      await loadSubCategories()
      await loadStats()
    } catch (err) {
      console.error('Failed to delete subcategory:', err)
      error.value = err?.message || 'فشل حذف التصنيف الفرعي'
    } finally {
      saving.value = false
    }
  }
}

// ============ Load data on mount ============
onMounted(async () => {
  await Promise.all([
    loadCategories(),
    loadStats(),
    loadSubCategories()
  ])
})
</script>

<style scoped>
/* صفحة البيانات العامة */
.data-page {
  background: #ffffff !important;
  color: var(--text-dark);
  min-height: 100vh;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

/* المحتوى بحجم الشاشة الكامل */
.fullscreen-content {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  text-align: center;
}

/* صف الإحصائيات - نفس تصميم صفحة المهندسين */
.stats-row {
  width: 100%;
  max-width: 100%;
  justify-content: center;
  margin: 0 auto;
  gap: 0;
  display: flex;
  flex-wrap: wrap;
  align-items: stretch;
  padding: 0;
}

.stats-row .v-col {
  margin: 0 !important;
  padding: 0.5rem !important;
  box-sizing: border-box !important;
}

.stats-row .stat-card {
  width: 100% !important;
  margin: 0 !important;
  box-sizing: border-box !important;
}

/* رأس الصفحة المحسن - نفس تصميم صفحة المهندسين */
.engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  background-color: #1976d2 !important;
  background-image: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  border-radius: 0;
  width: 100vw;
  max-width: 100vw;
  box-shadow: 0 8px 32px rgba(25, 118, 210, 0.3);
  position: relative;
  overflow: hidden;
  margin-left: calc(-50vw + 50%);
  margin-right: calc(-50vw + 50%);
  margin-bottom: 1.5rem;
  border: none;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: slideInFromTop 1s ease-out, shimmer 3s ease-in-out infinite;
}

/* التأكد من تطبيق اللون على جميع العناصر */
.data-page .engineers-header-card,
.fullscreen-content .engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  background-color: #1976d2 !important;
}

.engineers-header-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: sweep 2s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: diagonalShimmer 4s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 60px rgba(25, 118, 210, 0.5);
  animation: hoverPulse 0.6s ease-in-out;
}

.engineers-header-card:hover::before {
  animation: sweep 1s ease-in-out infinite;
}

.engineers-header-card:hover::after {
  animation: diagonalShimmer 2s ease-in-out infinite;
}

/* الخط المتدرج في الأعلى */
.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.3);
  animation: gradientFlow 3s ease-in-out infinite;
  z-index: 2;
}

/* محتوى الرأس */
.header-content {
  padding: 12px 16px !important;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  min-height: auto !important;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.1) 0%, rgba(21, 101, 192, 0.05) 100%);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 3;
  animation: fadeInUp 1.2s ease-out 0.3s both;
  max-width: calc(100vw - 320px);
  margin: 0 auto;
}

/* الجانب الأيمن */
.header-right {
  display: flex;
  align-items: center;
  gap: 1.8rem;
  text-align: right;
  padding: 0.8rem 1.5rem;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.15);
}

/* أيقونات المهندس الحديثة */
.engineer-emoji {
  position: relative;
  animation: slideInFromRight 1s ease-out 0.9s both, float 3s ease-in-out infinite 2s, pulse 2s ease-in-out infinite 2s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 0.5rem;
}

.engineer-emoji .v-icon {
  filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3));
  transition: all 0.3s ease;
  background: linear-gradient(135deg, #ffffff, #e3f2fd);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  position: relative;
  animation: iconGlow 2s ease-in-out infinite 3s, iconBounce 3s ease-in-out infinite 3s;
}

.engineer-emoji .v-icon:first-child {
  animation: iconGlow 2s ease-in-out infinite 3s, iconBounce 3s ease-in-out infinite 3s;
}

.engineer-emoji .v-icon:nth-child(2) {
  animation: iconGlow 2s ease-in-out infinite 3.5s, iconWave 4s ease-in-out infinite 3.5s;
}

.engineer-emoji .v-icon:nth-child(3) {
  animation: iconGlow 2s ease-in-out infinite 4s, iconSpin 6s linear infinite 4s, iconShake 2s ease-in-out infinite 4s;
}

.engineer-emoji .v-icon::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, #ffffff, #e3f2fd);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  filter: blur(2px);
  opacity: 0.3;
  z-index: -1;
}

.engineer-emoji:hover .v-icon {
  transform: scale(1.1) rotate(5deg);
  filter: drop-shadow(0 12px 24px rgba(255, 255, 255, 0.4));
}

/* النصوص */
.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.8rem;
  position: relative;
  animation: none !important;
}

.main-title {
  font-size: 1.2rem !important;
  font-weight: 900;
  color: #ffffff;
  margin: 0;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.4);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  letter-spacing: -0.02em;
  position: relative;
  line-height: 1.1;
}

.main-title::after {
  content: '';
  position: absolute;
  bottom: -4px;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 100%);
  border-radius: 2px;
  opacity: 0.8;
}

.subtitle {
  font-size: 0.75rem !important;
  color: #ffffff;
  margin: 0;
  font-weight: 600;
  line-height: 1.4;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  position: relative;
  opacity: 0.95;
}

/* تأثيرات الحركة */
@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-12px) rotate(2deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

@keyframes shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
}

@keyframes slideInFromTop {
  0% {
    transform: translateY(-100px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes slideInFromLeft {
  0% {
    transform: translateX(-50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes slideInFromRight {
  0% {
    transform: translateX(50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes sweep {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes diagonalShimmer {
  0%, 100% {
    transform: translateX(-100%) translateY(-100%) rotate(45deg);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translateX(100%) translateY(100%) rotate(45deg);
    opacity: 0;
  }
}

@keyframes gradientFlow {
  0%, 100% {
    background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  }
  50% {
    background: linear-gradient(90deg, #bbdefb 0%, #ffffff 50%, #e3f2fd 100%);
  }
}

@keyframes hoverPulse {
  0% {
    transform: translateY(-8px) scale(1.02);
  }
  50% {
    transform: translateY(-12px) scale(1.05);
  }
  100% {
    transform: translateY(-8px) scale(1.02);
  }
}

@keyframes fadeInUp {
  0% {
    transform: translateY(30px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes iconGlow {
  0%, 100% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 20px rgba(255, 255, 255, 0.3));
  }
  50% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 30px rgba(255, 255, 255, 0.6));
  }
}

@keyframes iconRotate {
  0% {
    transform: rotate(0deg);
  }
  25% {
    transform: rotate(5deg);
  }
  75% {
    transform: rotate(-5deg);
  }
  100% {
    transform: rotate(0deg);
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-8px);
  }
}

@keyframes iconWave {
  0%, 100% {
    transform: rotate(0deg) scale(1);
  }
  25% {
    transform: rotate(2deg) scale(1.05);
  }
  75% {
    transform: rotate(-2deg) scale(1.05);
  }
}

@keyframes iconSpin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

@keyframes iconShake {
  0%, 100% {
    transform: translateX(0);
  }
  25% {
    transform: translateX(-2px);
  }
  75% {
    transform: translateX(2px);
  }
}

/* Modern Statistics Cards - نفس تصميم الصفحة الرئيسية */
.modern-stat-card {
  position: relative !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  cursor: pointer !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  height: 100% !important;
  min-height: 140px !important;
  background: #ffffff !important;
}

.modern-stat-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.stat-card-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  opacity: 0.1;
  transition: opacity 0.3s ease;
}

.modern-stat-card:hover .stat-card-background {
  opacity: 0.2;
}

.stat-card-primary .stat-card-background {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
}

.stat-card-success .stat-card-background {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.stat-card-warning .stat-card-background {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.stat-card-info .stat-card-background {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
}

.stat-card-primary {
  background: #ffffff !important;
  border: 2px solid #3b82f6 !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15) !important;
}

.stat-card-success {
  background: #ffffff !important;
  border: 2px solid #10b981 !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.15) !important;
}

.stat-card-warning {
  background: #ffffff !important;
  border: 2px solid #f59e0b !important;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.15) !important;
}

.stat-card-info {
  background: #ffffff !important;
  border: 2px solid #06b6d4 !important;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.15) !important;
}

.stat-card-content {
  position: relative;
  z-index: 2;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  align-items: center;
  text-align: center;
}

.stat-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 64px;
  height: 64px;
  min-width: 64px;
  min-height: 64px;
  border-radius: 50%;
  margin-bottom: 0.25rem;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  position: relative;
  overflow: visible;
}

.stat-icon-wrapper::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #ffffff;
  z-index: 1;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.stat-card-primary .stat-icon-wrapper {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.stat-card-primary .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.2) 0%, rgba(37, 99, 235, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-success .stat-icon-wrapper {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.stat-card-success .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.2) 0%, rgba(5, 150, 105, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-warning .stat-icon-wrapper {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.stat-card-warning .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 158, 11, 0.2) 0%, rgba(217, 119, 6, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-card-info .stat-icon-wrapper {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3);
}

.stat-card-info .stat-icon-wrapper::before {
  content: '';
  position: absolute;
  inset: -3px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.2) 0%, rgba(8, 145, 178, 0.2) 100%);
  z-index: -1;
  filter: blur(8px);
}

.stat-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
  min-height: 32px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0;
  padding: 0;
}

.stat-icon-wrapper :deep(.v-icon) {
  font-size: 32px !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
  min-height: 32px !important;
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 !important;
  padding: 0 !important;
}

.stat-icon-wrapper :deep(svg) {
  width: 32px !important;
  height: 32px !important;
  font-size: 32px !important;
}

.stat-card-primary .stat-icon {
  color: #3b82f6 !important;
}

.stat-card-primary .stat-icon-wrapper :deep(.v-icon) {
  color: #3b82f6 !important;
}

.stat-card-success .stat-icon {
  color: #10b981 !important;
}

.stat-card-success .stat-icon-wrapper :deep(.v-icon) {
  color: #10b981 !important;
}

.stat-card-warning .stat-icon {
  color: #f59e0b !important;
}

.stat-card-warning .stat-icon-wrapper :deep(.v-icon) {
  color: #f59e0b !important;
}

.stat-card-info .stat-icon {
  color: #06b6d4 !important;
}

.stat-card-info .stat-icon-wrapper :deep(.v-icon) {
  color: #06b6d4 !important;
}

.check-icon,
.stat-icon-wrapper .check-icon,
.stat-icon-wrapper :deep(.check-icon) {
  transform: scaleX(-1) !important;
}

.stat-info {
  flex: 1;
  text-align: center;
  width: 100%;
}

.stat-value {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 0.5rem;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  text-align: center;
  font-variant-numeric: tabular-nums;
  unicode-bidi: embed;
  color: #000000 !important;
}

.stat-label {
  font-size: 1rem;
  font-weight: 500;
  text-align: center;
  color: #64748b;
}

/* بطاقات الإحصائيات - نفس تصميم صفحة المهندسين */
.stats-cards-container {
  display: flex;
  gap: 0.75rem;
  flex-wrap: nowrap;
  overflow: hidden;
  justify-content: space-between;
  padding: 0.5rem 0;
  width: 100%;
}

.stats-cards-container::-webkit-scrollbar {
  display: none;
}

.stat-card,
.stat-card-enhanced {
  background: linear-gradient(145deg, #9e7272 0%, #f8fafc 100%) !important;
  border-radius: 12px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  height: 100% !important;
  min-height: 100px !important;
  flex: 1 1 0;
  max-width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06), 0 1px 4px rgba(0, 0, 0, 0.03) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  margin: 0 !important;
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-sizing: border-box !important;
}

.stat-card::before,
.stat-card-enhanced::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  border-radius: 24px 24px 0 0;
  z-index: 1;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.stat-card::after,
.stat-card-enhanced::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 50% 0%, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
  border-radius: 24px;
  z-index: 1;
  pointer-events: none;
}

.stat-card:hover,
.stat-card-enhanced:hover {
  transform: translateY(-12px) scale(1.05);
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2), 0 8px 16px rgba(0, 0, 0, 0.1) !important;
  border-color: rgba(219, 206, 206, 0.9);
  background: linear-gradient(145deg, #fffcfc 0%, #f1f5f9 100%) !important;
}

.stat-card:hover::before,
.stat-card-enhanced:hover::before {
  height: 6px;
  opacity: 1;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.4);
}

.stat-card:hover::after,
.stat-card-enhanced:hover::after {
  opacity: 1;
}

.stat-card .v-icon,
.stat-card-enhanced .v-icon {
  filter: drop-shadow(0 6px 12px rgba(0, 0, 0, 0.15));
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 60px !important;
  width: 60px !important;
  height: 60px !important;
}

.stat-card:hover .v-icon,
.stat-card-enhanced:hover .v-icon {
  transform: scale(1.15) rotate(8deg);
  filter: drop-shadow(0 12px 24px rgba(0, 0, 0, 0.25));
}

.stat-card:hover .stat-icon,
.stat-card-enhanced:hover .stat-icon-enhanced {
  transform: scale(1.1);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  background: linear-gradient(135deg, rgba(242, 239, 239, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%);
}

.stat-card:hover p,
.stat-card-enhanced:hover p,
.stat-card:hover .stat-label,
.stat-card-enhanced:hover .stat-label {
  color: #475569 !important;
}

.stat-card h3,
.stat-card .stat-number,
.stat-card-enhanced h3,
.stat-card-enhanced .stat-number {
  font-size: 0.875rem !important;
  font-weight: 700 !important;
  margin-bottom: 4px !important;
  text-align: center !important;
  width: 100%;
  line-height: 1.2;
  position: relative;
  z-index: 2;
  color: #1976d2 !important;
}

.stat-card p,
.stat-card .stat-label,
.stat-card-enhanced p,
.stat-card-enhanced .stat-label {
  font-size: 0.7rem !important;
  font-weight: 500 !important;
  opacity: 1 !important;
  text-align: center !important;
  width: 100%;
  margin: 0;
  color: #64748b !important;
  position: relative;
  z-index: 3;
  visibility: visible !important;
  display: block !important;
}

/* تأثيرات خاصة لكل لون - نفس صفحة المهندسين */
.stat-card:nth-child(1)::before,
.stat-card-enhanced:nth-child(1)::before {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 50%, #7c3aed 100%);
}

.stat-card:nth-child(1):hover,
.stat-card-enhanced:nth-child(1):hover {
  box-shadow: 0 25px 50px rgba(59, 130, 246, 0.25), 0 8px 16px rgba(59, 130, 246, 0.15) !important;
}

.stat-card:nth-child(1) h3,
.stat-card-enhanced:nth-child(1) h3,
.stat-card:nth-child(1) .stat-number,
.stat-card-enhanced:nth-child(1) .stat-number {
  color: #3b82f6 !important;
}

.stat-card:nth-child(2)::before,
.stat-card-enhanced:nth-child(2)::before {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%);
}

.stat-card:nth-child(2):hover,
.stat-card-enhanced:nth-child(2):hover {
  box-shadow: 0 25px 50px rgba(16, 185, 129, 0.25), 0 8px 16px rgba(16, 185, 129, 0.15) !important;
}

.stat-card:nth-child(2) h3,
.stat-card-enhanced:nth-child(2) h3,
.stat-card:nth-child(2) .stat-number,
.stat-card-enhanced:nth-child(2) .stat-number {
  color: #10b981 !important;
}

.stat-card:nth-child(3)::before,
.stat-card-enhanced:nth-child(3)::before {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #b45309 100%);
}

.stat-card:nth-child(3):hover,
.stat-card-enhanced:nth-child(3):hover {
  box-shadow: 0 25px 50px rgba(245, 158, 11, 0.25), 0 8px 16px rgba(245, 158, 11, 0.15) !important;
}

.stat-card:nth-child(3) h3,
.stat-card-enhanced:nth-child(3) h3,
.stat-card:nth-child(3) .stat-number,
.stat-card-enhanced:nth-child(3) .stat-number {
  color: #f59e0b !important;
}

.stat-card:nth-child(4)::before,
.stat-card-enhanced:nth-child(4)::before {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 50%, #0e7490 100%);
}

.stat-card:nth-child(4):hover,
.stat-card-enhanced:nth-child(4):hover {
  box-shadow: 0 25px 50px rgba(6, 182, 212, 0.25), 0 8px 16px rgba(6, 182, 212, 0.15) !important;
}

.stat-card:nth-child(4) h3,
.stat-card-enhanced:nth-child(4) h3,
.stat-card:nth-child(4) .stat-number,
.stat-card-enhanced:nth-child(4) .stat-number {
  color: #06b6d4 !important;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--primary-color), var(--secondary-color));
  opacity: 0;
  transition: opacity 0.3s ease;
}

.hover-lift {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.card-glow {
  position: relative;
  overflow: hidden;
}

.card-glow::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}

.card-glow:hover::before {
  opacity: 1;
}

.smooth-transition {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.icon-glow {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.1));
  transition: filter 0.3s ease;
}

.stat-card:hover .icon-glow {
  filter: drop-shadow(0 6px 12px rgba(0, 0, 0, 0.15));
}

.stat-card:hover {
  transform: translateY(-8px) !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
}

.stat-card:hover::before {
  opacity: 1;
}

.stat-icon {
  transition: all 0.3s ease;
}

.stat-card:hover .stat-icon {
  transform: scale(1.1) rotate(5deg);
}

/* قسم الإجراءات */
.action-section {
  border-radius: 20px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 1px solid #e2e8f0 !important;
  margin-bottom: 1.5rem;
}

.action-section .v-card-text {
  padding: 1.5rem !important;
}

.add-btn {
  border-radius: 16px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  padding: 1rem 2rem !important;
  font-size: 1.1rem !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 8px 20px rgba(var(--primary-color-rgb), 0.3) !important;
}

.add-btn:hover {
  transform: translateY(-4px) !important;
  box-shadow: 0 12px 28px rgba(var(--primary-color-rgb), 0.4) !important;
}

/* بطاقة الجدول */
.data-table-card {
  border-radius: 20px !important;
  overflow: hidden;
  border: 1px solid rgba(0, 0, 0, 0.05);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1) !important;
}

/* رأس الجدول المحسن */
.data-table-card .v-card-title.table-header-enhanced,
.data-table-card .table-header-enhanced,
.table-header-enhanced,
.data-table-card :deep(.v-card-title),
.data-table-card :deep(.v-card-title.table-header-enhanced) {
  background: linear-gradient(135deg, rgb(180, 155, 163) 0%, rgb(150, 122, 129) 50%, rgb(120, 95, 103) 100%) !important;
  background-color: rgb(128, 95, 103) !important;
  background-image: linear-gradient(135deg, rgb(150, 122, 129) 0%, rgb(128, 95, 103) 50%, rgb(100, 70, 78) 100%) !important;
  padding: 6px 10px !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
  color: white !important;
  min-height: 32px !important;
  margin-bottom: 1rem !important;
}

.table-header-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.3)) !important;
  font-size: 16px !important;
  color: white !important;
}

.table-header-content {
  color: white !important;
}

.table-title {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25) !important;
  margin-bottom: 0 !important;
  letter-spacing: 0.2px !important;
  line-height: 1.2 !important;
}

.table-subtitle {
  display: none !important;
}

.add-button,
.add-button.v-btn,
.v-btn.add-button {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 50%, #3b82f6 100%) !important;
  backdrop-filter: blur(10px) !important;
  color: white !important;
  border-radius: 12px !important;
  padding: 12px 24px !important;
  font-weight: 700 !important;
  text-transform: none !important;
  box-shadow: 
    0 4px 16px rgba(30, 64, 175, 0.3),
    0 2px 8px rgba(37, 99, 235, 0.2),
    0 0 20px rgba(59, 130, 246, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  position: relative;
  overflow: hidden;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-align: center !important;
  min-width: auto !important;
  animation: buttonGlow 2s ease-in-out infinite, buttonPulse 3s ease-in-out infinite !important;
}

/* تأثير التوهج للزر */
@keyframes buttonGlow {
  0%, 100% {
    box-shadow: 
      0 4px 16px rgba(30, 64, 175, 0.3),
      0 2px 8px rgba(37, 99, 235, 0.2),
      0 0 20px rgba(59, 130, 246, 0.4),
      inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  }
  50% {
    box-shadow: 
      0 4px 20px rgba(30, 64, 175, 0.5),
      0 2px 12px rgba(37, 99, 235, 0.4),
      0 0 30px rgba(59, 130, 246, 0.6),
      inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  }
}

/* تأثير النبض للزر */
@keyframes buttonPulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.02);
  }
}

.add-button :deep(.v-btn__content) {
  color: white !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  width: 100% !important;
  text-align: center !important;
  gap: 8px !important;
}

.add-button :deep(.v-btn__prepend),
.add-button :deep(.v-btn__append) {
  color: white !important;
  margin: 0 !important;
}

.add-button :deep(.v-icon) {
  color: white !important;
  margin: 0 !important;
}

.add-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.6s ease;
  z-index: 1;
}

.add-button:hover::before {
  left: 100%;
  animation: shimmer 0.6s ease-in-out;
}

.add-button::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.3);
  transform: translate(-50%, -50%);
  transition: width 0.6s ease, height 0.6s ease, opacity 0.6s ease;
  z-index: 0;
  opacity: 0;
}

.add-button:active::after {
  width: 300px;
  height: 300px;
  opacity: 0;
  transition: width 0.3s ease, height 0.3s ease, opacity 0.3s ease;
}

.add-button:hover,
.add-button.v-btn:hover,
.v-btn.add-button:hover {
  transform: translateY(-4px) scale(1.05) !important;
  box-shadow: 
    0 12px 32px rgba(30, 64, 175, 0.5),
    0 6px 16px rgba(37, 99, 235, 0.5),
    0 0 40px rgba(59, 130, 246, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #2563eb 100%) !important;
  border-color: rgba(255, 255, 255, 0.6) !important;
  animation: buttonGlowHover 1.5s ease-in-out infinite !important;
}

@keyframes buttonGlowHover {
  0%, 100% {
    box-shadow: 
      0 12px 32px rgba(30, 64, 175, 0.5),
      0 6px 16px rgba(37, 99, 235, 0.5),
      0 0 40px rgba(59, 130, 246, 0.7),
      inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  }
  50% {
    box-shadow: 
      0 12px 36px rgba(30, 64, 175, 0.6),
      0 6px 20px rgba(37, 99, 235, 0.6),
      0 0 50px rgba(59, 130, 246, 0.8),
      inset 0 1px 0 rgba(255, 255, 255, 0.5) !important;
  }
}

.add-button:active {
  transform: translateY(-1px) scale(1.02) !important;
  box-shadow: 
    0 4px 12px rgba(30, 64, 175, 0.4),
    0 2px 6px rgba(37, 99, 235, 0.3),
    0 0 25px rgba(59, 130, 246, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  animation: buttonClick 0.3s ease !important;
}

@keyframes buttonClick {
  0% {
    transform: translateY(-1px) scale(1.02);
  }
  50% {
    transform: translateY(-1px) scale(0.98);
  }
  100% {
    transform: translateY(-1px) scale(1.02);
  }
}

.add-button :deep(.v-icon) {
  transition: transform 0.3s ease !important;
}

.add-button:hover :deep(.v-icon) {
  transform: rotate(90deg) scale(1.1) !important;
}

/* Simple Dialog Styles */
.simple-category-dialog .v-dialog {
  border-radius: 8px !important;
  box-shadow: none !important;
}

.simple-dialog-card {
  border-radius: 8px !important;
  overflow: hidden !important;
  box-shadow: none !important;
}

.simple-dialog-header {
  background: #1e40af !important;
  color: white !important;
  padding: 1rem 1.5rem !important;
  border-bottom: none !important;
}

.close-btn {
  opacity: 0.8 !important;
  transition: opacity 0.2s ease !important;
}

.close-btn:hover {
  opacity: 1 !important;
}

.simple-dialog-body {
  background: white !important;
  padding: 1.5rem !important;
}

.form-field-simple {
  margin-bottom: 1rem;
}

.field-label {
  display: block;
  color: #1e40af;
  font-size: 0.9rem;
  font-weight: 500;
  margin-bottom: 0.5rem;
  text-align: right;
}

.simple-text-field {
  margin-bottom: 0 !important;
}

.simple-text-field .v-field {
  border: 5px solid #1e40af !important;
  border-radius: 10px !important;
  background: #e2e8f0 !important;
  box-shadow: 0 0 0 2px #1e40af, 0 4px 8px rgba(30, 64, 175, 0.3) !important;
  min-height: 60px !important;
  outline: none !important;
  position: relative !important;
}

.simple-text-field .v-field--focused {
  border-color: #1e40af !important;
  box-shadow: 0 0 0 4px rgba(30, 64, 175, 0.4), 0 4px 8px rgba(30, 64, 175, 0.2) !important;
  background: white !important;
  transform: translateY(-1px) !important;
  outline: 2px solid #1e40af !important;
  outline-offset: 2px !important;
}

.simple-text-field .v-field__input {
  color: #1e40af !important;
  font-weight: 700 !important;
  font-size: 1.2rem !important;
  padding: 1.2rem 1.5rem !important;
  background: #ffffff !important;
  border: 3px solid #1e40af !important;
  border-radius: 8px !important;
  box-shadow: inset 0 3px 6px rgba(30, 64, 175, 0.2) !important;
  min-height: 50px !important;
  margin: 5px !important;
  width: calc(100% - 10px) !important;
}

.simple-text-field .v-field__outline {
  display: none !important;
}

.simple-text-field .v-field__field {
  background: transparent !important;
}

.simple-text-field .v-field::before {
  content: '';
  position: absolute;
  top: -3px;
  left: -3px;
  right: -3px;
  bottom: -3px;
  background: linear-gradient(45deg, #1e40af, #3b82f6);
  border-radius: 12px;
  z-index: -1;
  opacity: 0.5;
}

.simple-text-field .v-field::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: #f1f5f9;
  border-radius: 10px;
  z-index: -1;
}

.simple-text-field:hover .v-field {
  border-color: #1e40af !important;
  box-shadow: 0 0 0 2px #1e40af, 0 4px 8px rgba(30, 64, 175, 0.2) !important;
  transform: translateY(-0.5px) !important;
}

.simple-text-field .v-field__input:focus {
  background: white !important;
  border: 3px solid #1e40af !important;
  box-shadow: inset 0 2px 6px rgba(30, 64, 175, 0.15), 0 0 0 2px rgba(30, 64, 175, 0.2) !important;
  outline: none !important;
  transform: scale(1.02) !important;
}

.simple-text-field .v-field__input::placeholder {
  color: #94a3b8 !important;
  font-weight: 500 !important;
  font-size: 1rem !important;
}

/* Custom Input Field Styles */
.custom-input-container {
  position: relative;
  margin-top: 0.5rem;
}

.custom-input-field {
  width: 100% !important;
  padding: 1rem 1.2rem !important;
  font-size: 1.1rem !important;
  font-weight: 600 !important;
  color: #1e40af !important;
  background: #ffffff !important;
  border: 4px solid #1e40af !important;
  border-radius: 10px !important;
  box-shadow: none !important;
  outline: none !important;
  transition: all 0.3s ease !important;
  min-height: 50px !important;
  box-sizing: border-box !important;
}

.custom-input-field:focus {
  border: 5px solid #1e40af !important;
  box-shadow: none !important;
  transform: none !important;
  background: #f8fafc !important;
}

.custom-input-field::placeholder {
  color: #64748b !important;
  font-weight: 500 !important;
  opacity: 0.8 !important;
}

.custom-input-field:hover {
  border: 4px solid #1e40af !important;
  box-shadow: none !important;
  transform: none !important;
}

/* Table Content Styles */
.sequence-number {
  text-align: center !important;
  font-size: 0.6rem !important;
  font-weight: 500 !important;
  color: #374151 !important;
  padding: 3px 6px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  min-width: 30px !important;
}

.category-name-text {
  font-size: 0.6rem !important;
  font-weight: 500 !important;
  color: #1f2937 !important;
  padding: 2px 4px !important;
  text-align: center !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.action-btn {
  color: #1e40af !important;
  font-weight: 500 !important;
  min-width: 28px !important;
  width: 28px !important;
  height: 28px !important;
}

.action-btn :deep(.v-icon) {
  font-size: 16px !important;
}

.action-btn:hover {
  background: rgba(30, 64, 175, 0.1) !important;
  color: #1e40af !important;
}

/* Table Header Enhancement */
.table-header-enhanced {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  color: white !important;
}

.table-header-enhanced .v-card-title {
  color: white !important;
}

.table-header-enhanced h2,
.table-header-enhanced p {
  color: white !important;
}

/* Data Table Enhancement */
.categories-table {
  margin-top: 1rem !important;
}

.categories-table .v-data-table__wrapper {
  border: 1px solid #e5e7eb !important;
  border-radius: 8px !important;
}

.categories-table .v-data-table-header,
.data-table-card .v-data-table-header {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  font-weight: 600 !important;
  min-height: 28px !important;
  height: auto !important;
  padding: 0 !important;
}

/* تصغير الخط في جميع خلايا الجدول */
.categories-table .v-data-table__td,
.categories-table .v-data-table__tbody td,
.categories-table .v-data-table__wrapper tbody td,
.categories-table .v-data-table__wrapper tbody td *,
.categories-table :deep(.v-data-table__td),
.categories-table :deep(.v-data-table__td *),
.data-table-card .v-data-table__td,
.data-table-card .v-data-table__tbody td,
.data-table-card :deep(.v-data-table__td),
.data-table-card :deep(.v-data-table__td *) {
  font-size: 0.75rem !important;
}

/* تنسيق رأس الجدول - تصغير الخط والحجم */
.categories-table .v-data-table-header th,
.categories-table .v-data-table__th,
.categories-table thead th,
.categories-table thead th *,
.data-table-card .v-data-table-header th,
.data-table-card .v-data-table__th,
.data-table-card thead th,
.data-table-card thead th *,
.data-table-card :deep(.v-data-table-header th),
.data-table-card :deep(.v-data-table-header th *),
.data-table-card :deep(.v-data-table-header__content),
.data-table-card :deep(.v-data-table-header__content *),
.data-table-card :deep(.v-data-table__th),
.data-table-card :deep(.v-data-table__th *),
.data-table-card :deep(.v-data-table__th .v-data-table-header__content),
.data-table-card :deep(.v-data-table__th .v-data-table-header__content *) {
  font-size: 0.6rem !important;
  font-weight: 600 !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  text-align: center !important;
  padding: 2px 4px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.1px !important;
  line-height: 1.2 !important;
  height: auto !important;
  min-height: 20px !important;
  vertical-align: middle !important;
  white-space: nowrap !important;
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
  -webkit-filter: none !important;
  -webkit-text-stroke: none !important;
}

/* إزالة الظلال من جميع العناصر داخل عناوين الجدول */
.categories-table .v-data-table-header th *,
.categories-table .v-data-table__th *,
.data-table-card .v-data-table-header th *,
.data-table-card .v-data-table__th *,
.categories-table :deep(.v-data-table-header th *),
.categories-table :deep(.v-data-table__th *),
.data-table-card :deep(.v-data-table-header th *),
.data-table-card :deep(.v-data-table__th *) {
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
  -webkit-filter: none !important;
}

/* إزالة جميع الظلال من عناوين الجدول */
.categories-table .v-data-table-header th,
.categories-table .v-data-table__th,
.categories-table thead th,
.categories-table thead th *,
.data-table-card .v-data-table-header th,
.data-table-card .v-data-table__th,
.data-table-card thead th,
.data-table-card thead th *,
.data-table-card :deep(.v-data-table-header th),
.data-table-card :deep(.v-data-table-header th *),
.data-table-card :deep(.v-data-table__th),
.data-table-card :deep(.v-data-table__th *),
.data-table-card :deep(.v-data-table-header__content),
.data-table-card :deep(.v-data-table-header__content *),
.data-table-card :deep(.v-data-table-header__content span),
.data-table-card :deep(.v-data-table-header__content div),
.categories-table :deep(.v-data-table-header th),
.categories-table :deep(.v-data-table-header th *),
.categories-table :deep(.v-data-table__th),
.categories-table :deep(.v-data-table__th *),
.categories-table :deep(.v-data-table-header__content),
.categories-table :deep(.v-data-table-header__content *) {
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
  -webkit-text-stroke: none !important;
  -webkit-filter: none !important;
}

/* إزالة الظلال من عناوين الجدول - قواعد إضافية */
.categories-table .v-data-table-header,
.data-table-card .v-data-table-header {
  box-shadow: none !important;
  filter: none !important;
}

.categories-table .v-data-table__wrapper tbody tr {
  border-bottom: 1px solid #f3f4f6 !important;
}

.categories-table .v-data-table__wrapper tbody tr:hover {
  background: #f8fafc !important;
}

.categories-table .v-data-table__wrapper tbody td {
  padding: 3px 4px !important;
  border-bottom: 1px solid #f3f4f6 !important;
  text-align: center !important;
  vertical-align: middle !important;
  font-size: 0.6rem !important;
  min-height: 32px !important;
}

.categories-table .v-data-table__wrapper tbody td > div {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  width: 100% !important;
  font-size: 0.6rem !important;
}

.simple-dialog-actions {
  background: white !important;
  padding: 0 1.5rem 1.5rem !important;
}

/* Sub Categories Dialog Styles */
.sub-categories-dialog .v-dialog {
  border-radius: 8px !important;
  box-shadow: none !important;
}

.sub-categories-card {
  border-radius: 8px !important;
  overflow: hidden !important;
  box-shadow: none !important;
}

.sub-categories-header {
  background: #1e40af !important;
  color: white !important;
  padding: 1rem 1.5rem !important;
  border-bottom: none !important;
}

.sub-categories-body {
  background: white !important;
  padding: 1.5rem !important;
}

.sub-categories-list {
  max-height: 400px;
  overflow-y: auto;
}

.sub-category-item {
  display: flex;
  align-items: center;
  padding: 1rem;
  margin-bottom: 0.5rem;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.sub-category-item:hover {
  background: #f1f5f9;
  border-color: #1e40af;
  transform: translateY(-1px);
}

.sub-category-number {
  width: 40px;
  height: 40px;
  background: #1e40af;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.9rem;
  margin-left: 1rem;
}

.sub-category-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
  text-align: right;
}

.sub-category-name {
  font-size: 1.1rem;
  font-weight: 600;
  color: #1f2937;
}

.sub-category-progress {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.875rem;
  justify-content: flex-end;
}

.progress-label {
  color: #666;
  font-weight: 500;
}

.progress-value {
  color: #1976d2;
  font-weight: 700;
  background: linear-gradient(135deg, #e3f2fd 0%, #bbdefb 100%);
  padding: 2px 10px;
  border-radius: 12px;
  border: 1px solid #90caf9;
}

/* Progress Display in Dialog */
.progress-display-container {
  margin-top: 1rem;
  padding: 1rem;
  background: #f8fafc;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
}

.progress-bar-wrapper {
  width: 100%;
  height: 24px;
  background: #e2e8f0;
  border-radius: 12px;
  overflow: hidden;
  position: relative;
  margin-bottom: 0.5rem;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.1);
}

.progress-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%);
  border-radius: 12px;
  transition: width 0.3s ease;
  position: relative;
  box-shadow: 0 2px 4px rgba(59, 130, 246, 0.3);
}

.progress-bar-fill::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent 0%, rgba(255, 255, 255, 0.3) 50%, transparent 100%);
  animation: shimmer 2s infinite;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.progress-text-display {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 0.5rem;
}

.progress-percentage {
  font-size: 1.2rem;
  font-weight: 700;
  color: #1e40af;
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  padding: 4px 16px;
  border-radius: 20px;
  border: 2px solid #3b82f6;
  min-width: 60px;
  text-align: center;
}

.sub-category-actions {
  display: flex;
  gap: 0.5rem;
}

.sub-action-btn {
  color: #1e40af !important;
}

.sub-action-btn:hover {
  background: rgba(30, 64, 175, 0.1) !important;
}

.add-sub-btn {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  text-transform: none !important;
  padding: 1rem !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(30, 64, 175, 0.3) !important;
}

.add-sub-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 16px rgba(30, 64, 175, 0.4) !important;
}

/* Modern Enhanced Statistics Cards */
.stat-card-enhanced {
  background: linear-gradient(135deg, #ffffff 0%, #f1f5f9 50%, #e2e8f0 100%) !important;
  border: none !important;
  border-radius: 24px !important;
  transition: all 0.4s cubic-bezier(0.25, 0.8, 0.25, 1) !important;
  min-height: 160px !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08), 0 1px 3px rgba(0, 0, 0, 0.1) !important;
}

.stat-card-enhanced::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 6px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  border-radius: 24px 24px 0 0;
}

.stat-card-enhanced::after {
  content: '';
  position: absolute;
  top: -50%;
  right: -50%;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.3) 0%, transparent 70%);
  border-radius: 50%;
  opacity: 0;
  transition: all 0.6s ease;
}

.stat-card-enhanced:hover {
  transform: translateY(-8px) scale(1.03) !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15), 0 8px 16px rgba(0, 0, 0, 0.1) !important;
}

.stat-card-enhanced:hover::after {
  opacity: 1;
  transform: scale(1.5);
}

.stat-icon-enhanced,
.stat-icon {
  position: relative;
  z-index: 2;
  margin-bottom: 0.5rem;
  padding: 8px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 241, 241, 0.8) 0%, rgba(248, 250, 252, 0.8) 100%);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  display: flex;
  justify-content: center;
  align-items: center;
  width: auto;
  height: auto;
  margin: 0 auto 0.5rem auto;
}

.stat-card:hover .stat-icon,
.stat-card-enhanced:hover .stat-icon-enhanced {
  transform: scale(1.1);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  background: linear-gradient(135deg, rgba(242, 239, 239, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%);
}

.stat-icon-enhanced::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-radius: 24px;
  z-index: -1;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.stat-card-enhanced:hover .stat-icon-enhanced,
.stat-card:hover .stat-icon {
  transform: scale(1.1);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  background: linear-gradient(135deg, rgba(242, 239, 239, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%);
}

.stat-card-enhanced:hover .stat-icon-enhanced::before {
  opacity: 1;
}

.stat-number {
  font-family: 'Inter', 'Arial', 'Helvetica', sans-serif !important;
  font-size: 2.5rem !important;
  font-weight: 900 !important;
  line-height: 1.1 !important;
  margin-bottom: 0.75rem !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text !important;
  -webkit-text-fill-color: transparent !important;
  background-clip: text !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.1) !important;
  position: relative;
  z-index: 2;
}

.stat-label {
  font-weight: 700 !important;
  font-size: 1rem !important;
  letter-spacing: 1px !important;
  text-transform: uppercase !important;
  color: #64748b !important;
  opacity: 0.9 !important;
  position: relative;
  z-index: 2;
  font-family: 'Inter', 'Arial', sans-serif !important;
}

.stats-row {
  margin-bottom: 2rem !important;
}

/* Special effects for each card type */
.stat-card-enhanced:nth-child(1) .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.15) 0%, rgba(147, 197, 253, 0.15) 100%);
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.2);
}

.stat-card-enhanced:nth-child(1):hover .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.25) 0%, rgba(147, 197, 253, 0.25) 100%);
  box-shadow: 0 8px 25px rgba(59, 130, 246, 0.3);
}

.stat-card-enhanced:nth-child(2) .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(34, 197, 94, 0.15) 0%, rgba(134, 239, 172, 0.15) 100%);
  box-shadow: 0 4px 15px rgba(34, 197, 94, 0.2);
}

.stat-card-enhanced:nth-child(2):hover .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(34, 197, 94, 0.25) 0%, rgba(134, 239, 172, 0.25) 100%);
  box-shadow: 0 8px 25px rgba(34, 197, 94, 0.3);
}

.stat-card-enhanced:nth-child(3) .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(251, 146, 60, 0.15) 0%, rgba(253, 186, 116, 0.15) 100%);
  box-shadow: 0 4px 15px rgba(251, 146, 60, 0.2);
}

.stat-card-enhanced:nth-child(3):hover .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(251, 146, 60, 0.25) 0%, rgba(253, 186, 116, 0.25) 100%);
  box-shadow: 0 8px 25px rgba(251, 146, 60, 0.3);
}

.stat-card-enhanced:nth-child(4) .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(139, 92, 246, 0.15) 0%, rgba(196, 181, 253, 0.15) 100%);
  box-shadow: 0 4px 15px rgba(139, 92, 246, 0.2);
}

.stat-card-enhanced:nth-child(4):hover .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(139, 92, 246, 0.25) 0%, rgba(196, 181, 253, 0.25) 100%);
  box-shadow: 0 8px 25px rgba(139, 92, 246, 0.3);
}

.save-btn-simple {
  color: #424242 !important;
  font-weight: 500 !important;
  text-transform: none !important;
  padding: 0.5rem 1.5rem !important;
  border-radius: 4px !important;
  border: 1px solid #e0e0e0 !important;
  background: white !important;
}

.save-btn-simple:hover {
  background: #f5f5f5 !important;
  border-color: #bdbdbd !important;
}

.save-btn-simple:disabled {
  color: #bdbdbd !important;
  background: #fafafa !important;
  border-color: #f0f0f0 !important;
}

.count-chip {
  background: rgba(255, 255, 255, 0.3) !important;
  color: white !important;
  border: 1px solid rgba(255, 255, 255, 0.5) !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25) !important;
  backdrop-filter: blur(10px) !important;
  padding: 2px 8px !important;
  height: 20px !important;
}

/* حقول البحث والفلترة */
.search-field,
.filter-field {
  width: 100%;
}

/* جدول التصنيفات */
.categories-table {
  border-radius: 12px !important;
}

.categories-table .v-data-table__wrapper {
  border-radius: 12px;
}

.categories-table .v-data-table-header {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
}

.categories-table .v-data-table-header th {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.7rem !important;
  border-bottom: 2px solid #047857 !important;
  padding: 4px 5px !important;
  text-align: center !important;
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
  min-height: 28px !important;
}

/* تنسيقات إضافية للتأكد من التطبيق */
.categories-table .v-data-table thead th,
.categories-table thead th,
.v-data-table thead th,
.v-data-table-header th {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.7rem !important;
  border-bottom: 2px solid #047857 !important;
  padding: 4px 5px !important;
  min-height: 28px !important;
  text-align: center !important;
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
}



.categories-table .v-data-table__tbody tr {
  transition: all 0.3s ease;
}

.categories-table .v-data-table__tbody tr:hover {
  background: #f8fafc !important;
  transform: scale(1.01);
}

.categories-table .v-data-table__tbody td {
  padding: 1rem 0.75rem !important;
  border-bottom: 1px solid rgba(0, 0, 0, 0.05) !important;
  font-size: 0.75rem !important;
}

/* عرض اللون المحسن */
.color-display {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.5rem;
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%);
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.color-display:hover {
  background: linear-gradient(135deg, #e2e8f0 0%, #f1f5f9 100%);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.color-circle {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  border: 3px solid #ffffff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.color-circle::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: conic-gradient(from 0deg, transparent, rgba(255,255,255,0.3), transparent);
  border-radius: 50%;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.color-display:hover .color-circle::before {
  opacity: 1;
}

.color-display:hover .color-circle {
  transform: scale(1.1);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2);
}

.color-name {
  font-weight: 600;
  font-size: 0.95rem;
  color: #1f2937;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

/* عرض اسم التصنيف */
.category-name-display {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem;
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%);
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.category-name-display::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(59, 130, 246, 0.1), transparent);
  transition: left 0.5s ease;
}

.category-name-display:hover::before {
  left: 100%;
}

.category-name-display:hover {
  background: linear-gradient(135deg, #e2e8f0 0%, #f1f5f9 100%);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #3b82f6;
}

.category-name {
  font-weight: 700;
  font-size: 1rem;
  color: #1f2937;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  transition: color 0.3s ease;
}

.category-name-display:hover .category-name {
  color: #1e40af;
}

.category-name-display .v-icon {
  transition: all 0.3s ease;
}

.category-name-display:hover .v-icon {
  transform: scale(1.1) rotate(5deg);
  filter: drop-shadow(0 2px 4px rgba(59, 130, 246, 0.3));
}


/* عرض الوصف */
.description-display {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  background: linear-gradient(135deg, #f9fafb 0%, #ffffff 100%);
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  transition: all 0.3s ease;
}

.description-display:hover {
  background: linear-gradient(135deg, #f3f4f6 0%, #f9fafb 100%);
  border-color: #9ca3af;
}

.description-text {
  font-weight: 400;
  font-size: 0.6rem;
  color: #6b7280;
  font-style: italic;
}

.description-display:hover .description-text {
  color: #374151;
}

/* عرض التاريخ */
.date-display {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  background: linear-gradient(135deg, #ecfdf5 0%, #ffffff 100%);
  border-radius: 8px;
  border: 1px solid #d1fae5;
  transition: all 0.3s ease;
}

.date-display:hover {
  background: linear-gradient(135deg, #d1fae5 0%, #ecfdf5 100%);
  border-color: #10b981;
  transform: translateY(-1px);
}

.date-text {
  font-weight: 500;
  font-size: 0.6rem;
  color: #059669;
  font-family: 'Arial', 'Helvetica', sans-serif;
}

.date-display:hover .date-text {
  color: #047857;
}

/* تنسيقات النافذة المنبثقة لتفاصيل التصنيف */
.category-details-dialog .v-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border-radius: 20px !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid rgba(59, 130, 246, 0.1) !important;
}

.details-header-enhanced {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 2rem !important;
  margin: 0 !important;
  position: relative !important;
}

.category-avatar {
  border: 3px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.color-circle-large {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  border: 3px solid #ffffff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.category-title-content {
  color: white !important;
}

.category-title {
  color: white !important;
  font-weight: 700 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  margin-bottom: 0.25rem !important;
}

.category-subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500 !important;
  opacity: 0.9 !important;
}

.close-btn {
  color: white !important;
  background: rgba(255, 255, 255, 0.1) !important;
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.2) !important;
  transform: scale(1.1) !important;
}

/* المحتوى المحسن */
.details-content-enhanced {
  background: #ffffff !important;
  padding: 2rem !important;
}

.detail-card {
  border-radius: 16px !important;
  transition: all 0.3s ease !important;
  border: 2px solid #e2e8f0 !important;
  background: linear-gradient(135deg, #ffffff 0%, #fefefe 100%) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
  height: 100%;
}

.detail-card:hover {
  transform: translateY(-4px) !important;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.12) !important;
}

.detail-card-title {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%) !important;
  color: #1e293b !important;
  border-radius: 14px 14px 0 0 !important;
  padding: 12px 16px !important;
  margin: -2px -2px 16px -2px !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
}

.detail-card-content {
  padding: 0 16px 16px 16px !important;
}

.detail-value {
  font-weight: 600 !important;
  font-size: 1rem !important;
  color: #1f2937 !important;
}

.color-circle-small {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 2px solid #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.status-chip-detail, .projects-chip-detail {
  border-radius: 12px !important;
  font-weight: 600 !important;
  padding: 0.5rem 0.75rem !important;
}

.date-value {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  color: #059669 !important;
  font-weight: 600 !important;
}

.description-text-enhanced {
  margin-top: 0.5rem !important;
  padding: 1rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%) !important;
  border-radius: 12px !important;
  border: 1px solid #e2e8f0 !important;
  color: #1f2937 !important;
  line-height: 1.6 !important;
  font-style: italic !important;
}

/* أزرار النافذة المنبثقة */
.details-actions-enhanced {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-radius: 0 0 20px 20px !important;
  padding: 1.5rem 2rem !important;
}

.confirm-btn {
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  border-radius: 12px !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  transition: all 0.3s ease !important;
}

.confirm-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(59, 130, 246, 0.4) !important;
}

/* تنسيقات خاصة لكل نوع بطاقة */
.info-card {
  border-left: 4px solid #3b82f6 !important;
}

.color-card {
  border-left: 4px solid #06b6d4 !important;
}

.status-card {
  border-left: 4px solid #10b981 !important;
}

.projects-card {
  border-left: 4px solid #06b6d4 !important;
}

.date-card {
  border-left: 4px solid #10b981 !important;
}

.description-card {
  border-left: 4px solid #8b5cf6 !important;
}

/* تنسيقات نافذة التعديل */
.edit-category-dialog .v-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border-radius: 20px !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid rgba(59, 130, 246, 0.1) !important;
}

.edit-dialog-title {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 20px 24px !important;
  margin: 0 !important;
}

.edit-dialog-title h2 {
  color: white !important;
  font-weight: 700 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
}

.edit-dialog-title .v-avatar {
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.edit-dialog-content {
  background: #ffffff !important;
  padding: 24px !important;
}

/* بطاقات الحقول */
.form-field-card {
  border-radius: 16px !important;
  transition: all 0.3s ease !important;
  border: 2px solid #e2e8f0 !important;
  background: linear-gradient(135deg, #ffffff 0%, #fefefe 100%) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
  height: 100%;
}

.form-field-card:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.12) !important;
}

.field-card-title {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%) !important;
  color: #1e293b !important;
  border-radius: 14px 14px 0 0 !important;
  padding: 12px 16px !important;
  margin: -2px -2px 16px -2px !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
}

.field-card-content {
  padding: 0 16px 16px 16px !important;
}

/* الحقول المحسنة */
.form-field-enhanced .v-field {
  border-radius: 12px !important;
  transition: all 0.3s ease !important;
}

.form-field-enhanced .v-field:focus-within {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1) !important;
}

.form-field-enhanced .v-label {
  font-weight: 500 !important;
  color: #374151 !important;
}

.form-field-enhanced .v-field__input {
  font-weight: 500 !important;
  color: #1f2937 !important;
}

/* حقل اختيار اللون */
.color-picker-container {
  display: flex;
  align-items: center;
  gap: 12px;
}

.color-input {
  flex: 1;
}

.color-preview {
  width: 40px;
  height: 40px;
  border-radius: 12px;
  border: 3px solid #ffffff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;
}

.color-preview:hover {
  transform: scale(1.1);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2);
}

/* المفتاح */
.status-switch {
  margin-top: 8px;
}

/* الأزرار */
.edit-dialog-actions {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-radius: 0 0 20px 20px !important;
  padding: 20px 24px !important;
}

.save-btn {
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  border-radius: 12px !important;
  font-weight: 600 !important;
  padding: 12px 24px !important;
  transition: all 0.3s ease !important;
}

.save-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(59, 130, 246, 0.4) !important;
}

.cancel-btn {
  border-radius: 12px !important;
  font-weight: 500 !important;
  padding: 12px 24px !important;
  transition: all 0.3s ease !important;
}

/* تنسيقات خاصة لكل نوع بطاقة */
.color-picker-card {
  border-left: 4px solid #f59e0b !important;
}

.color-picker-card .field-card-title {
  background: linear-gradient(135deg, #fef3c7 0%, #fde68a 100%) !important;
}

.status-card {
  border-left: 4px solid #10b981 !important;
}

.status-card .field-card-title {
  background: linear-gradient(135deg, #d1fae5 0%, #a7f3d0 100%) !important;
}

.description-card {
  border-left: 4px solid #8b5cf6 !important;
}

.description-card .field-card-title {
  background: linear-gradient(135deg, #ede9fe 0%, #ddd6fe 100%) !important;
}

/* عناوين الجدول النيلي */
.categories-table .v-data-table-header th {
  font-weight: 700 !important;
  font-size: 0.7rem !important;
  color: #ffffff !important;
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  border-bottom: 2px solid #047857 !important;
  padding: 4px 5px !important;
  min-height: 28px !important;
  text-align: center !important;
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
}


/* أزرار الإجراءات المحسنة */
.actions-cell-enhanced {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
  align-items: center;
  padding: 0.5rem;
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%);
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}

.action-btn-enhanced {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border-radius: 8px !important;
  padding: 8px !important;
  min-width: 36px !important;
  min-height: 36px !important;
  position: relative !important;
  overflow: hidden !important;
}

.action-btn-enhanced::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
  border-radius: 8px;
}

.action-btn-enhanced:hover::before {
  opacity: 1;
}

.action-btn-enhanced:hover {
  transform: translateY(-2px) scale(1.1) !important;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15) !important;
}

/* أزرار محددة */
.view-btn {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3) !important;
}

.view-btn:hover {
  box-shadow: 0 8px 20px rgba(6, 182, 212, 0.4) !important;
}

.edit-btn {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
}

.edit-btn:hover {
  box-shadow: 0 8px 20px rgba(59, 130, 246, 0.4) !important;
}

.delete-btn {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3) !important;
}

.delete-btn:hover {
  box-shadow: 0 8px 20px rgba(239, 68, 68, 0.4) !important;
}

/* الشارات المحسنة */
.status-chip, .projects-chip {
  border-radius: 8px !important;
  font-weight: 500 !important;
  padding: 2px 6px !important;
  font-size: 0.6rem !important;
  height: 20px !important;
  transition: all 0.3s ease !important;
  position: relative !important;
  overflow: hidden !important;
}

.status-chip :deep(.v-icon),
.projects-chip :deep(.v-icon) {
  font-size: 12px !important;
  margin: 0 2px !important;
}

.status-chip::before, .projects-chip::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s ease;
}

.status-chip:hover::before, .projects-chip:hover::before {
  left: 100%;
}

.status-chip:hover, .projects-chip:hover {
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.15) !important;
}

.active-chip {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3) !important;
}

.active-chip:hover {
  box-shadow: 0 8px 20px rgba(16, 185, 129, 0.4) !important;
}

.inactive-chip {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3) !important;
}

.inactive-chip:hover {
  box-shadow: 0 8px 20px rgba(245, 158, 11, 0.4) !important;
}

.projects-chip {
  background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.3) !important;
}

.projects-chip:hover {
  box-shadow: 0 8px 20px rgba(6, 182, 212, 0.4) !important;
}

/* أزرار الحوار */
.dialog-actions {
  padding: 1rem 1.5rem !important;
  background: #f8fafc !important;
  border-top: 1px solid #e2e8f0;
}

/* استجابة للشاشات المختلفة */
@media (max-width: 1200px) {
  .stats-row {
    gap: 1.5rem;
    padding: 0 1rem;
  }
}

@media (max-width: 768px) {
  .page-title {
    font-size: 2rem;
  }
  
  .action-section .v-card-text {
    padding: 1rem !important;
  }
}

/* نافذة التفاصيل */
.details-header {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-bottom: 2px solid rgba(0, 0, 0, 0.05);
  padding: 1.5rem !important;
}

.details-header .color-circle {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 3px solid #ffffff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.details-content {
  padding: 2rem !important;
}

.detail-item {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
  padding: 0.75rem;
  background: #f8fafc;
  border-radius: 12px;
  border-left: 4px solid var(--primary-color);
  transition: all 0.3s ease;
}

.detail-item:hover {
  background: #e2e8f0;
  transform: translateX(4px);
}

.detail-item strong {
  color: var(--text-dark);
  margin-left: 0.5rem;
}

.description-text {
  margin-top: 0.5rem;
  padding: 1rem;
  background: #ffffff;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
  color: var(--text-dark);
  line-height: 1.6;
}

.details-actions {
  padding: 1rem 1.5rem !important;
  background: #f8fafc;
  border-top: 1px solid #e2e8f0;
}

@media (max-width: 600px) {
  .stats-row {
    flex-wrap: wrap;
    gap: 1rem;
  }
  
  .page-header {
    padding: 1rem;
    max-height: 150px;
  }
  
  .page-title {
    font-size: 1.8rem;
  }
  
  .page-subtitle {
    font-size: 1rem;
  }
  
  .details-content {
    padding: 1rem !important;
  }
  
  .detail-item {
    flex-direction: column;
    align-items: flex-start;
    text-align: right;
  }
  
  .detail-item strong {
    margin-left: 0;
    margin-bottom: 0.25rem;
  }
}

/* تنسيقات global للتأكد من التطبيق */
:deep(.categories-table .v-data-table-header th),
:deep(.categories-table .v-data-table thead th),
:deep(.categories-table thead th),
:deep(.v-data-table thead th),
:deep(.v-data-table-header th) {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.7rem !important;
  border-bottom: 2px solid #047857 !important;
  padding: 4px 5px !important;
  min-height: 28px !important;
  text-align: center !important;
  text-shadow: none !important;
  box-shadow: none !important;
  filter: none !important;
}

:deep(.categories-table .v-data-table-header) {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%) !important;
}

/* جعل الأسهم بيضاء واضحة */
:deep(.categories-table .v-data-table-header th .v-data-table-header__content),
:deep(.categories-table .v-data-table-header th .v-icon),
:deep(.categories-table .v-data-table-header th .mdi-arrow-up),
:deep(.categories-table .v-data-table-header th .mdi-arrow-down),
:deep(.categories-table .v-data-table-header th .mdi-menu-up),
:deep(.categories-table .v-data-table-header th .mdi-menu-down),
:deep(.categories-table .v-data-table-header th .v-data-table-header__content .v-icon),
:deep(.categories-table .v-data-table-header th .v-data-table-header__content svg),
:deep(.categories-table .v-data-table-header th svg),
:deep(.categories-table .v-data-table-header th .v-icon svg),
:deep(.categories-table .v-data-table-header th path),
:deep(.categories-table .v-data-table-header th .v-icon path) {
  color: #ffffff !important;
  fill: #ffffff !important;
  stroke: #ffffff !important;
  opacity: 1 !important;
}

/* تنسيقات إضافية للأسهم */
:deep(.categories-table .v-data-table-header th .v-data-table-header__content .v-icon) {
  color: #ffffff !important;
  fill: #ffffff !important;
  opacity: 1 !important;
  filter: brightness(0) invert(1) !important;
  font-size: 18px !important;
  font-weight: bold !important;
}

:deep(.categories-table .v-data-table-header th:hover .v-icon) {
  color: #ffffff !important;
  fill: #ffffff !important;
  opacity: 1 !important;
}

/* جعل الأسهم أكثر وضوحاً */
:deep(.categories-table .v-data-table-header th .v-data-table-header__content) {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 8px !important;
}

:deep(.categories-table .v-data-table-header th .v-data-table-header__content .v-icon) {
  color: #ffffff !important;
  fill: #ffffff !important;
  opacity: 1 !important;
  filter: drop-shadow(0 0 2px rgba(255, 255, 255, 0.8)) !important;
  font-size: 20px !important;
  font-weight: 900 !important;
  text-shadow: 0 0 4px rgba(255, 255, 255, 0.6) !important;
}

/* قواعد إضافية للتأكد من تطبيق اللون الأزرق الداكن */
.data-page .engineers-header-card,
.fullscreen-content .engineers-header-card,
.engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  background-color: #1976d2 !important;
  background-image: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
}

.data-table-card .v-card-title,
.data-table-card .v-card-title.table-header-enhanced,
.data-table-card :deep(.v-card-title),
.data-table-card :deep(.v-card-title.table-header-enhanced),
.table-header-enhanced {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
  background-color: #1976d2 !important;
  background-image: linear-gradient(135deg, #1976d2 0%, #1565c0 100%) !important;
}

/* ========================================
   Animations - نفس صفحة المهندسين
   ======================================== */

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-12px) rotate(2deg);
  }
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

@keyframes shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
}

@keyframes slideInFromTop {
  0% {
    transform: translateY(-100px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes slideInFromLeft {
  0% {
    transform: translateX(-50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes slideInFromRight {
  0% {
    transform: translateX(50px);
    opacity: 0;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes fadeInUp {
  0% {
    transform: translateY(30px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes sweep {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes diagonalShimmer {
  0%, 100% {
    transform: translateX(-100%) translateY(-100%) rotate(45deg);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translateX(100%) translateY(100%) rotate(45deg);
    opacity: 0;
  }
}

@keyframes gradientFlow {
  0%, 100% {
    background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  }
  50% {
    background: linear-gradient(90deg, #bbdefb 0%, #ffffff 50%, #e3f2fd 100%);
  }
}

@keyframes hoverPulse {
  0% {
    transform: translateY(-8px) scale(1.02);
  }
  50% {
    transform: translateY(-12px) scale(1.05);
  }
  100% {
    transform: translateY(-8px) scale(1.02);
  }
}

@keyframes iconGlow {
  0%, 100% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 20px rgba(255, 255, 255, 0.3));
  }
  50% {
    filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 30px rgba(255, 255, 255, 0.6));
  }
}

@keyframes iconRotate {
  0% {
    transform: rotate(0deg);
  }
  25% {
    transform: rotate(5deg);
  }
  75% {
    transform: rotate(-5deg);
  }
  100% {
    transform: rotate(0deg);
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-8px);
  }
}

@keyframes iconWave {
  0%, 100% {
    transform: rotate(0deg) scale(1);
  }
  25% {
    transform: rotate(2deg) scale(1.05);
  }
  75% {
    transform: rotate(-2deg) scale(1.05);
  }
}

@keyframes iconSpin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

@keyframes iconShake {
  0%, 100% {
    transform: translateX(0);
  }
  25% {
    transform: translateX(-2px);
  }
  75% {
    transform: translateX(2px);
  }
}

</style>
