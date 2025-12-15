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
  isEditingSubCategory.value = true
  editingSubCategoryId.value = subCategory.id
  subCategoryForm.value = {
    name: subCategory.name || '',
    progress: subCategory.progress || 0
  }
  addSubCategoryDialog.value = true
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


<style>
/* Import page styles */
@import './styles/categories.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
