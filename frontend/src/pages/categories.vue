<template>
  <div class="data-page">
    <v-container fluid class="fullscreen-content">
      <!-- Page Header -->
      <div class="page-header">
        <div class="header-content">
          <div class="title-row">
            <v-icon size="2rem" class="title-icon">mdi-tag-multiple</v-icon>
            <div class="page-title">إدارة التصنيفات</div>
          </div>
          <div class="page-subtitle">تنظيم وإدارة تصنيفات المشاريع والمهام</div>
        </div>
      </div>

      <!-- Statistics Cards -->
      <v-row class="stats-row">
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card-enhanced pa-4 text-center" elevation="3">
            <div class="stat-icon-enhanced mb-3">
              <v-icon size="64" color="primary">mdi-tag-multiple</v-icon>
            </div>
            <h3 class="stat-number text-h4 font-weight-bold text-primary mb-2">{{ totalCategories }}</h3>
            <p class="stat-label text-body-2 text-medium-emphasis mb-0">إجمالي التصنيفات</p>
          </v-card>
        </v-col>

        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card-enhanced pa-4 text-center" elevation="3">
            <div class="stat-icon-enhanced mb-3">
              <v-icon size="64" color="success">mdi-check-circle</v-icon>
            </div>
            <h3 class="stat-number text-h4 font-weight-bold text-success mb-2">{{ activeCategories }}</h3>
            <p class="stat-label text-body-2 text-medium-emphasis mb-0">تصنيفات نشطة</p>
          </v-card>
        </v-col>

        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card-enhanced pa-4 text-center" elevation="3">
            <div class="stat-icon-enhanced mb-3">
              <v-icon size="64" color="warning">mdi-pause-circle</v-icon>
            </div>
            <h3 class="stat-number text-h4 font-weight-bold text-warning mb-2">{{ inactiveCategories }}</h3>
            <p class="stat-label text-body-2 text-medium-emphasis mb-0">تصنيفات معطلة</p>
          </v-card>
        </v-col>

        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card-enhanced pa-4 text-center" elevation="3">
            <div class="stat-icon-enhanced mb-3">
              <v-icon size="64" color="info">mdi-folder-multiple</v-icon>
            </div>
            <h3 class="stat-number text-h4 font-weight-bold text-info mb-2">{{ totalProjects }}</h3>
            <p class="stat-label text-body-2 text-medium-emphasis mb-0">مشاريع مرتبطة</p>
          </v-card>
        </v-col>
      </v-row>


      <!-- Categories Table -->
      <v-card class="data-table-card" elevation="3">
        <v-card-title class="table-header-enhanced">
          <div class="d-flex align-center">
            <v-icon class="me-3 table-header-icon" color="primary">mdi-tag-multiple</v-icon>
            <div class="table-header-content">
              <h2 class="text-h5 font-weight-bold table-title">قائمة التصنيفات</h2>
              <p class="text-subtitle-2 table-subtitle mb-0">إدارة وتنظيم جميع تصنيفات المشاريع</p>
            </div>
            <v-spacer />
            <v-btn
              color="white"
              size="default"
              variant="elevated"
              @click="openAddCategoryDialog"
              class="add-btn-header"
              rounded="xl"
              elevation="4"
            >
              <v-icon class="me-2 add-btn-icon">mdi-plus</v-icon>
              إضافة تصنيف جديد
            </v-btn>
            <v-chip color="primary" variant="tonal" class="count-chip ms-3">
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
        >
          <!-- Sequence Column -->
          <template v-slot:item.sequence="{ index }">
            <div class="sequence-number">
              {{ index + 1 }}
            </div>
          </template>

          <!-- Category Name Column -->
          <template v-slot:item.name="{ item }">
            <div class="category-name-text">{{ item.name }}</div>
          </template>

          <!-- Status Column -->
          <template v-slot:item.isActive="{ item }">
            <v-chip
              :color="item.isActive ? 'success' : 'warning'"
              size="small"
              variant="tonal"
              class="status-chip"
              :class="item.isActive ? 'active-chip' : 'inactive-chip'"
            >
              <v-icon :icon="item.isActive ? 'mdi-check-circle' : 'mdi-pause-circle'" size="small" class="me-1" />
              {{ item.isActive ? 'نشط' : 'معطل' }}
            </v-chip>
          </template>

          <!-- Description Column -->
          <template v-slot:item.description="{ item }">
            <div class="description-display">
              <v-icon class="me-2" color="grey" size="small">mdi-text</v-icon>
              <span class="description-text">{{ item.description || 'لا يوجد وصف' }}</span>
            </div>
          </template>

          <!-- Projects Count Column -->
          <template v-slot:item.projectsCount="{ item }">
            <v-chip
              color="info"
              size="small"
              variant="tonal"
              class="projects-chip"
            >
              <v-icon icon="mdi-folder-multiple" size="small" class="me-1" />
              {{ item.projectsCount }} مشروع
            </v-chip>
          </template>

          <!-- Creation Date Column -->
          <template v-slot:item.createdAt="{ item }">
            <div class="date-display">
              <v-icon class="me-2" color="success" size="small">mdi-calendar</v-icon>
              <span class="date-text">{{ item.createdAt || 'غير محدد' }}</span>
            </div>
          </template>

          <!-- Actions Column -->
          <template v-slot:item.actions="{ item }">
            <div class="d-flex align-center gap-2 justify-center">
              <v-btn
                icon="mdi-view-list"
                variant="text"
                size="small"
                color="primary"
                @click="showSubCategories(item)"
                class="action-btn"
                title="عرض الأصناف الثانوية"
              />
              <v-btn
                icon="mdi-pencil"
                variant="text"
                size="small"
                color="warning"
                @click="editCategory(item)"
                class="action-btn"
                title="تعديل الصنف"
              />
            </div>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Category Dialog -->
      <v-dialog v-model="dialog" max-width="500" class="simple-category-dialog">
        <v-card class="simple-dialog-card">
          <!-- Header -->
          <v-card-title class="simple-dialog-header">
            <div class="d-flex align-center justify-space-between w-100">
              <h3 class="text-h6 font-weight-bold text-white mb-0">
                إضافة صنف جديد
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
          </v-card-text>

          <!-- Actions -->
          <v-card-actions class="simple-dialog-actions pa-6 pt-0">
            <v-spacer />
            <v-btn
              color="grey-darken-1"
              variant="text"
              @click="saveCategory"
              :disabled="!categoryForm.name"
              class="save-btn-simple"
            >
              حفظ
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
                    <span class="progress-value">{{ subCategory.progress || 0 }}%</span>
                  </div>
                </div>
                <div class="sub-category-actions">
                  <v-btn
                    icon="mdi-pencil"
                    variant="text"
                    size="small"
                    color="primary"
                    @click="editSubCategory(subCategory)"
                    class="sub-action-btn"
                  />
                  <v-btn
                    icon="mdi-delete"
                    variant="text"
                    size="small"
                    color="error"
                    @click="deleteSubCategory(subCategory)"
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
            <v-form ref="subCategoryForm" v-model="subCategoryFormValid">
              <div class="form-field-simple mb-4">
                <label class="field-label">اسم الصنف :</label>
                <div class="custom-input-container">
                  <input
                    v-model="subCategoryForm.name"
                    type="text"
                    class="custom-input-field"
                    placeholder="أدخل اسم الصنف الثانوي"
                    required
                  />
                </div>
              </div>
              <div class="form-field-simple">
                <label class="field-label">نسبة الإنجاز (%) :</label>
                <div class="custom-input-container">
                  <input
                    v-model="subCategoryForm.progress"
                    type="number"
                    class="custom-input-field"
                    placeholder="0"
                    min="0"
                    max="100"
                    required
                  />
                </div>
                <!-- Progress Bar Display -->
                <div class="progress-display-container mt-3">
                  <div class="progress-bar-wrapper">
                    <div 
                      class="progress-bar-fill" 
                      :style="{ width: (subCategoryForm.progress || 0) + '%' }"
                    ></div>
                  </div>
                  <div class="progress-text-display">
                    <span class="progress-percentage">{{ subCategoryForm.progress || 0 }}%</span>
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
              :disabled="!subCategoryFormValid || !subCategoryForm.name"
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

// Category form
const categoryForm = ref({
  name: '',
  colorName: '',
  color: '#3b82f6',
  isActive: true,
  description: ''
})

// Sub Category form
const subCategoryForm = ref({
  name: '',
  progress: 0
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

// Categories data
const categories = ref([
  {
    id: 1,
    name: 'الأعمال الترابية',
    colorName: 'أزرق',
    color: '#3b82f6',
    description: 'جميع الأعمال المتعلقة بالتربة والحفر',
    isActive: true,
    projectsCount: 12,
    createdAt: '2024-01-15'
  },
  {
    id: 2,
    name: 'الحصى الخابط',
    colorName: 'أخضر',
    color: '#10b981',
    description: 'أعمال الحصى والمواد الخابطية',
    isActive: true,
    projectsCount: 8,
    createdAt: '2024-01-20'
  },
  {
    id: 3,
    name: 'اعمال الباور كيرير',
    colorName: 'برتقالي',
    color: '#f59e0b',
    description: 'أعمال الطاقة والكهرباء',
    isActive: true,
    projectsCount: 15,
    createdAt: '2024-01-25'
  },
  {
    id: 4,
    name: 'اعمال التبليط',
    colorName: 'بنفسجي',
    color: '#8b5cf6',
    description: 'أعمال التبليط والتسوية',
    isActive: true,
    projectsCount: 6,
    createdAt: '2024-02-01'
  },
  {
    id: 5,
    name: 'أعمال المقرنص',
    colorName: 'أحمر',
    color: '#ef4444',
    description: 'أعمال المقرنص والزخارف',
    isActive: false,
    projectsCount: 4,
    createdAt: '2024-02-05'
  },
  {
    id: 6,
    name: 'اعمال المحاري',
    colorName: 'أزرق',
    color: '#06b6d4',
    description: 'أعمال المحاري والمواد المحارية',
    isActive: true,
    projectsCount: 9,
    createdAt: '2024-02-10'
  },
  {
    id: 7,
    name: 'اعمال شبكات الماء',
    colorName: 'أخضر',
    color: '#059669',
    description: 'أعمال شبكات المياه والصرف',
    isActive: true,
    projectsCount: 11,
    createdAt: '2024-02-15'
  }
])

// Sub Categories data
const subCategories = ref([
  {
    id: 1,
    parentId: 1,
    name: 'حفر الأساسات',
    description: 'حفر أساسات المباني',
    progress: 75
  },
  {
    id: 2,
    parentId: 1,
    name: 'ردم التربة',
    progress: 50,
    description: 'ردم وتنعيم التربة'
  },
  {
    id: 3,
    parentId: 2,
    name: 'حصى خابط ناعم',
    description: 'حصى خابط ناعم الحبيبات'
  },
  {
    id: 4,
    parentId: 2,
    name: 'حصى خابط خشن',
    description: 'حصى خابط خشن الحبيبات'
  },
  {
    id: 5,
    parentId: 3,
    name: 'أسلاك كهربائية',
    description: 'توصيل الأسلاك الكهربائية'
  },
  {
    id: 6,
    parentId: 3,
    name: 'عدادات الكهرباء',
    description: 'تركيب عداد الكهرباء'
  }
])

// Computed properties
const totalCategories = computed(() => categories.value.length)
const activeCategories = computed(() => categories.value.filter(c => c.isActive).length)
const inactiveCategories = computed(() => categories.value.filter(c => !c.isActive).length)
const totalProjects = computed(() => categories.value.reduce((sum, c) => sum + c.projectsCount, 0))

const filteredCategories = computed(() => {
  let filtered = categories.value

  if (searchQuery.value) {
    filtered = filtered.filter(category =>
      category.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      category.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      category.colorName.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (selectedStatus.value) {
    if (selectedStatus.value === 'active') {
      filtered = filtered.filter(category => category.isActive)
    } else if (selectedStatus.value === 'inactive') {
      filtered = filtered.filter(category => !category.isActive)
    }
  }

  return filtered
})

// Methods
const openAddCategoryDialog = () => {
  isEditing.value = false
  categoryForm.value = {
    name: '',
    colorName: '',
    color: '#3b82f6',
    isActive: true,
    description: ''
  }
  dialog.value = true
}

const editCategory = (category) => {
  isEditing.value = true
  categoryForm.value = { ...category }
  dialog.value = true
}

const viewCategoryDetails = (category) => {
  selectedCategory.value = category
  detailsDialog.value = true
}

const closeDialog = () => {
  dialog.value = false
  categoryForm.value = {
    name: '',
    colorName: '',
    color: '#3b82f6',
    isActive: true,
    description: ''
  }
}

const saveCategory = () => {
  if (isEditing.value) {
    const index = categories.value.findIndex(c => c.id === categoryForm.value.id)
    if (index !== -1) {
      categories.value[index] = {
        ...categories.value[index],
        ...categoryForm.value,
        updatedAt: new Date().toISOString().split('T')[0]
      }
    }
  } else {
    const newCategory = {
      id: Date.now(),
      ...categoryForm.value,
      projectsCount: 0,
      createdAt: new Date().toISOString().split('T')[0]
    }
    categories.value.push(newCategory)
  }
  closeDialog()
}

const deleteCategory = (category) => {
  selectedCategory.value = category
  deleteDialog.value = true
}

// Sub Categories functions
const showSubCategories = (category) => {
  selectedCategory.value = category
  subCategoriesDialog.value = true
}

const getSubCategories = (parentId) => {
  return subCategories.value.filter(sub => sub.parentId === parentId)
}

const openAddSubCategoryDialog = () => {
  isEditingSubCategory.value = false
  editingSubCategoryId.value = null
  subCategoryForm.value = {
    name: '',
    progress: 0
  }
  addSubCategoryDialog.value = true
}

const closeAddSubCategoryDialog = () => {
  addSubCategoryDialog.value = false
  isEditingSubCategory.value = false
  editingSubCategoryId.value = null
  subCategoryForm.value = {
    name: '',
    progress: 0
  }
}

const saveSubCategory = () => {
  if (!subCategoryForm.value.name) return
  
  if (isEditingSubCategory.value && editingSubCategoryId.value) {
    // تعديل صنف ثانوي موجود
    const index = subCategories.value.findIndex(sub => sub.id === editingSubCategoryId.value)
    if (index !== -1) {
      subCategories.value[index].name = subCategoryForm.value.name
      subCategories.value[index].progress = parseFloat(subCategoryForm.value.progress) || 0
    }
  } else {
    // إضافة صنف ثانوي جديد
    const newSubCategory = {
      id: Date.now(),
      parentId: selectedCategory.value.id,
      name: subCategoryForm.value.name,
      progress: parseFloat(subCategoryForm.value.progress) || 0,
      description: ''
    }
    subCategories.value.push(newSubCategory)
  }
  closeAddSubCategoryDialog()
}

const editSubCategory = (subCategory) => {
  console.log('Editing sub category:', subCategory)
  isEditingSubCategory.value = true
  editingSubCategoryId.value = subCategory.id
  subCategoryForm.value = {
    name: subCategory.name || '',
    progress: subCategory.progress || 0
  }
  addSubCategoryDialog.value = true
  console.log('Dialog opened, form:', subCategoryForm.value)
}

const deleteSubCategory = (subCategory) => {
  if (confirm(`هل أنت متأكد من حذف "${subCategory.name}"؟`)) {
    const index = subCategories.value.findIndex(sub => sub.id === subCategory.id)
    if (index !== -1) {
      subCategories.value.splice(index, 1)
    }
  }
}

const confirmDelete = () => {
  const index = categories.value.findIndex(c => c.id === selectedCategory.value.id)
  if (index !== -1) {
    categories.value.splice(index, 1)
  }
  deleteDialog.value = false
  selectedCategory.value = null
}

// Load data from localStorage on mount
onMounted(() => {
  const savedCategories = localStorage.getItem('categories')
  if (savedCategories) {
    categories.value = JSON.parse(savedCategories)
  }
})

// Save to localStorage when categories change
const saveToLocalStorage = () => {
  localStorage.setItem('categories', JSON.stringify(categories.value))
}

// Watch for changes and save
import { watch } from 'vue'
watch(categories, saveToLocalStorage, { deep: true })
</script>

<style scoped>
/* صفحة البيانات العامة */
.data-page {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: var(--text-dark);
  min-height: 100vh;
  padding: 1rem;
  width: 100%;
}

/* المحتوى بحجم الشاشة الكامل */
.fullscreen-content {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  text-align: center;
}

/* صف الإحصائيات */
.stats-row {
  width: 100%;
  justify-content: center;
  margin: 0 auto;
  gap: 2rem;
  display: flex;
  flex-wrap: nowrap;
  align-items: stretch;
  padding: 0 2rem;
}

/* رأس الصفحة المحسن */
.page-header {
  text-align: center;
  padding: 1.5rem 1rem;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-radius: 16px;
  width: 100%;
  box-shadow: 
    0 8px 32px rgba(59, 130, 246, 0.25),
    0 4px 16px rgba(37, 99, 235, 0.2) !important;
  position: relative;
  overflow: hidden;
  margin: 0;
  max-height: 140px;
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.page-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.5) 0%, rgba(255, 255, 255, 0.8) 50%, rgba(255, 255, 255, 0.5) 100%);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.3);
  z-index: 1;
}

.page-header::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.15) 50%, transparent 70%);
  animation: shimmer 3s infinite;
  z-index: 0;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.header-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.title-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 0.5rem;
  padding: 0.25rem 0;
}

.title-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
  transition: all 0.3s ease;
  color: white !important;
  font-size: 2.2rem !important;
  position: relative;
  z-index: 2;
}

.title-icon:hover {
  transform: scale(1.1);
  filter: drop-shadow(0 4px 8px rgba(255, 255, 255, 0.4));
}

.page-title {
  font-size: 1.8rem;
  font-weight: 800;
  color: white !important;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  letter-spacing: 1px;
  line-height: 1.2;
  font-family: 'Arial', 'Helvetica', sans-serif;
  position: relative;
  z-index: 2;
}

.page-subtitle {
  font-size: 0.95rem;
  color: rgba(255, 255, 255, 0.95) !important;
  position: relative;
  z-index: 2;
  color: #475569;
  font-weight: 600;
  text-shadow: 0 2px 4px rgba(71, 85, 105, 0.15);
  letter-spacing: 0.5px;
  line-height: 1.4;
  font-family: 'Arial', 'Helvetica', sans-serif;
}

/* بطاقات الإحصائيات */
.stat-card {
  border-radius: 12px !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 1px solid #e2e8f0 !important;
  overflow: hidden;
  position: relative;
  min-height: 100px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.stat-card h3 {
  font-size: 1.2rem !important;
  font-weight: 700 !important;
  margin-bottom: 4px !important;
  text-align: center !important;
  width: 100%;
  line-height: 1.2;
}

.stat-card p {
  font-size: 0.75rem !important;
  font-weight: 500 !important;
  opacity: 1 !important;
  text-align: center !important;
  width: 100%;
  margin: 0;
  color: #424242 !important;
  visibility: visible !important;
  display: block !important;
}

.stat-card:nth-child(1) h3 {
  color: #3b82f6 !important;
}

.stat-card:nth-child(2) h3 {
  color: #10b981 !important;
}

.stat-card:nth-child(3) h3 {
  color: #f59e0b !important;
}

.stat-card:nth-child(4) h3 {
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
.table-header-enhanced {
  background: linear-gradient(135deg, #1e40af 0%, #1d4ed8 100%) !important;
  padding: 2rem !important;
  border-bottom: 3px solid rgba(255, 255, 255, 0.3) !important;
  color: white !important;
}

.table-header-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.3)) !important;
  color: white !important;
}

.table-header-content {
  color: white !important;
}

.table-title {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.75rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  margin-bottom: 0.5rem !important;
  letter-spacing: 0.5px !important;
}

.table-subtitle {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 600 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3) !important;
  opacity: 1 !important;
}

.add-btn-header {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  color: #1e40af !important;
  font-weight: 700 !important;
  text-transform: none !important;
  padding: 0.75rem 1.5rem !important;
  font-size: 0.95rem !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  box-shadow: 0 4px 16px rgba(30, 64, 175, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.8) !important;
  position: relative !important;
  overflow: hidden !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-align: center !important;
  line-height: 1.4 !important;
  letter-spacing: 0.5px !important;
}

.add-btn-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(30, 64, 175, 0.1), transparent);
  transition: left 0.6s ease;
}

.add-btn-header:hover::before {
  left: 100%;
}

.add-btn-header:hover {
  background: linear-gradient(135deg, #ffffff 0%, #f1f5f9 100%) !important;
  transform: translateY(-3px) scale(1.02) !important;
  box-shadow: 0 8px 24px rgba(30, 64, 175, 0.3) !important;
  border-color: rgba(30, 64, 175, 0.2) !important;
}

.add-btn-header:active {
  transform: translateY(-1px) scale(0.98) !important;
  transition: all 0.1s ease !important;
}

.add-btn-icon {
  transition: all 0.3s ease !important;
  font-size: 1.2rem !important;
  margin-right: 0.5rem !important;
  display: flex !important;
  align-items: center !important;
}

.add-btn-header:hover .add-btn-icon {
  transform: rotate(90deg) scale(1.1) !important;
  filter: drop-shadow(0 2px 4px rgba(30, 64, 175, 0.3)) !important;
}

.add-btn-header .v-btn__content {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  gap: 0.5rem !important;
  width: 100% !important;
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
  font-size: 1rem !important;
  font-weight: 600 !important;
  color: #374151 !important;
  padding: 0.5rem !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.category-name-text {
  font-size: 1.1rem !important;
  font-weight: 600 !important;
  color: #1f2937 !important;
  padding: 0.5rem 0 !important;
  text-align: center !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.action-btn {
  color: #1e40af !important;
  font-weight: 600 !important;
}

.action-btn:hover {
  background: rgba(30, 64, 175, 0.1) !important;
  color: #1e40af !important;
}

/* Table Header Enhancement */
.table-header-enhanced {
  background: #1e40af !important;
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
.categories-table .v-data-table__wrapper {
  border: 1px solid #e5e7eb !important;
  border-radius: 8px !important;
}

.categories-table .v-data-table-header {
  background: #1e40af !important;
  color: white !important;
}

.categories-table .v-data-table-header th {
  background: #1e40af !important;
  color: white !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  text-align: center !important;
}

.categories-table .v-data-table__wrapper tbody tr {
  border-bottom: 1px solid #f3f4f6 !important;
}

.categories-table .v-data-table__wrapper tbody tr:hover {
  background: #f8fafc !important;
}

.categories-table .v-data-table__wrapper tbody td {
  padding: 1rem 0.75rem !important;
  border-bottom: 1px solid #f3f4f6 !important;
  text-align: center !important;
  vertical-align: middle !important;
}

.categories-table .v-data-table__wrapper tbody td > div {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  width: 100% !important;
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

.stat-icon-enhanced {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 88px;
  height: 88px;
  margin: 0 auto;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15) 0%, rgba(118, 75, 162, 0.15) 100%);
  border-radius: 22px;
  transition: all 0.4s cubic-bezier(0.25, 0.8, 0.25, 1);
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.2);
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

.stat-card-enhanced:hover .stat-icon-enhanced {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.25) 0%, rgba(118, 75, 162, 0.25) 100%);
  transform: scale(1.15) rotate(5deg);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.3);
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
  background: rgba(255, 255, 255, 0.25) !important;
  color: white !important;
  border: 2px solid rgba(255, 255, 255, 0.4) !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  backdrop-filter: blur(10px) !important;
  padding: 0.75rem 1.5rem !important;
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
  background: #1e40af !important;
}

.categories-table .v-data-table-header th {
  background: #1e40af !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  border-bottom: 2px solid #1d4ed8 !important;
  padding: 14px 16px !important;
  text-align: center !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

/* تنسيقات إضافية للتأكد من التطبيق */
.categories-table .v-data-table thead th,
.categories-table thead th,
.v-data-table thead th,
.v-data-table-header th {
  background: #1e40af !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  border-bottom: 2px solid #1d4ed8 !important;
  padding: 14px 16px !important;
  text-align: center !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
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
  font-weight: 500;
  font-size: 0.9rem;
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
  font-weight: 600;
  font-size: 0.9rem;
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
  font-size: 1rem !important;
  color: #ffffff !important;
  background: #1e40af !important;
  border-bottom: 2px solid #1d4ed8 !important;
  padding: 14px 16px !important;
  text-align: center !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
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
  border-radius: 12px !important;
  font-weight: 600 !important;
  padding: 0.5rem 0.75rem !important;
  transition: all 0.3s ease !important;
  position: relative !important;
  overflow: hidden !important;
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
  background: #1e40af !important;
  color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  border-bottom: 2px solid #1d4ed8 !important;
  padding: 14px 16px !important;
  text-align: center !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

:deep(.categories-table .v-data-table-header) {
  background: #1e40af !important;
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

</style>
