<template>
  <v-container class="fill-height engineers-page" max-width="1200">
    <div class="centered-content">
      <!-- Header Section - Card Style -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="header-text">
              <h1 class="main-title">المهندسين</h1>
              <p class="subtitle">إدارة وتتبع جميع المهندسين والمشاريع</p>
            </div>
            <div class="engineer-emoji">
              <v-icon size="60" color="white">mdi-account-hard-hat</v-icon>
              <v-icon size="45" color="white" class="secondary-icon">mdi-account-tie</v-icon>
              <v-icon size="35" color="white" class="tertiary-icon">mdi-account-cog</v-icon>
            </div>
          </div>
        </div>
      </div>

        <!-- Summary Cards -->
        <v-row class="mb-6 stats-row">
            <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
              <div class="stat-icon mb-3 icon-glow">
                <v-icon size="64" color="primary">mdi-account-group</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-primary mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ engineers.length }}</h3>
              <p class="text-subtitle-1 text-primary mb-0">إجمالي المهندسين</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
              <div class="stat-icon mb-3 icon-glow">
                <v-icon size="64" color="success">mdi-check-circle</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ activeEngineers }}</h3>
              <p class="text-subtitle-1 text-success mb-0">مهندسين نشطين</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
              <div class="stat-icon mb-3 icon-glow">
                <v-icon size="64" color="info">mdi-briefcase</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalProjects }}</h3>
              <p class="text-subtitle-1 text-info mb-0">إجمالي المشاريع</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
            <v-card class="stat-card pa-6 text-center hover-lift card-glow smooth-transition" elevation="2">
              <div class="stat-icon mb-3 icon-glow">
                <v-icon size="64" color="warning">mdi-star</v-icon>
              </div>
              <h3 class="text-h3 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ averageRating }}</h3>
              <p class="text-subtitle-1 text-warning mb-0">متوسط التقييم</p>
              </v-card>
            </v-col>
          </v-row>

      <!-- Engineers Table -->
      <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
        <v-card-title class="d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" style="color: white !important;">mdi-account-group</v-icon>
            <span class="text-h6 font-weight-bold" style="color: white !important; font-family: 'Arial', 'Helvetica', sans-serif;">قائمة المهندسين</span>
            <v-chip class="ms-3" style="background-color: rgba(255, 255, 255, 0.3) !important; color: white !important; border: 1px solid rgba(255, 255, 255, 0.5);" size="small">{{ engineers.length }}</v-chip>
        </div>
          <v-btn
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddEngineerDialog"
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
            v-model="searchQuery"
            prepend-inner-icon="mdi-magnify"
            label="البحث في المهندسين..."
                  variant="outlined"
            density="comfortable"
                  clearable
            hide-details
            class="mb-4 search-field-enhanced"
            style="font-size: 1.1rem !important; font-weight: 500 !important;"
            placeholder="ابحث بالاسم أو البريد الإلكتروني أو التخصص..."
            color="primary"
            bg-color="white"
            persistent-hint
            hint="اكتب للبحث في أسماء المهندسين أو تخصصاتهم"
            hint-class="search-hint-text"
          />
          </v-card-text>

        <v-data-table
          :headers="tableHeaders"
          :items="engineers"
          :search="searchQuery"
          class="elevation-0"
          no-data-text="لا يوجد مهندسين"
          loading-text="جاري التحميل..."
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
              <div class="font-weight-bold text-primary" style="font-size: 1rem !important; color: #1976d2 !important;">{{ item.name }}</div>
              <div class="text-caption text-medium-emphasis" style="font-size: 0.875rem !important; color: #666 !important;">{{ item.email }}</div>
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
                v-for="skill in item.skills.slice(0, 2)"
                      :key="skill"
                size="small"
                      color="primary"
                variant="tonal"
              >
                      {{ skill }}
              </v-chip>
                  <v-chip
                v-if="item.skills.length > 2"
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
                @click="viewEngineer(item)"
              />
        <v-btn
                icon="mdi-pencil"
                    size="small"
          variant="text"
                color="warning"
                @click="editEngineer(item)"
              />
              <v-btn
                icon="mdi-briefcase-plus"
                size="small"
                variant="text"
                color="success"
                @click="openProjectsDialog(item)"
                title="إضافة مشروع"
              />
                  <v-btn
                icon="mdi-delete"
                    size="small"
                variant="text"
                color="error"
                @click="deleteEngineer(item)"
              />
            </div>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Engineer Dialog -->
      <v-dialog v-model="dialog" max-width="700" class="edit-engineer-dialog">
        <v-card class="edit-dialog-card">
          <v-card-title class="edit-dialog-title">
            <div class="d-flex align-center">
              <v-avatar size="32" color="primary" variant="tonal" class="me-3">
                <v-icon>{{ isEditing ? 'mdi-pencil' : 'mdi-plus' }}</v-icon>
        </v-avatar>
              <h2 class="text-h5 font-weight-bold mb-0">
                {{ isEditing ? 'تعديل المهندس' : 'إضافة حساب جديد' }}
              </h2>
            </div>
          </v-card-title>

          <v-card-text>
            <v-form ref="form" v-model="formValid">
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    v-model="engineerForm.name"
                    label="اسم المهندس"
                    :rules="[v => !!v || 'اسم المهندس مطلوب']"
                    required
                    variant="outlined"
                    density="comfortable"
                    prepend-inner-icon="mdi-account"
                    class="form-field-enhanced"
                    style="font-size: 1.1rem !important; color: #000000 !important; --v-field-input-color: #000000 !important; --v-input-color: #000000 !important; --v-field-input-text-color: #000000 !important; --v-field-input-fill-color: #000000 !important;"
                    color="black"
                    label-color="black"
                    input-color="black"
                    text-color="black"
                  />
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    v-model="engineerForm.password"
                    label="كلمة المرور"
                    :type="showPassword ? 'text' : 'password'"
                    :rules="[v => !!v || 'كلمة المرور مطلوبة']"
                    required
                    variant="outlined"
                    density="comfortable"
                    prepend-inner-icon="mdi-lock"
                    :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                    @click:append-inner="showPassword = !showPassword"
                    class="form-field-enhanced"
                    style="font-size: 1.1rem !important; color: #000000 !important; --v-field-input-color: #000000 !important; --v-input-color: #000000 !important; --v-field-input-text-color: #000000 !important; --v-field-input-fill-color: #000000 !important;"
                    color="black"
                    label-color="black"
                    input-color="black"
                    text-color="black"
                  />
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>

          <v-divider />
          <v-card-actions class="dialog-actions pa-4">
            <v-spacer />
            <v-btn
              color="grey"
              variant="text"
              @click="closeDialog"
              size="large"
              class="cancel-btn"
            >
              <v-icon class="me-2">mdi-close</v-icon>
              إلغاء
            </v-btn>
                <v-btn
                  color="primary"
              variant="elevated"
              @click="saveEngineer"
              :disabled="!formValid"
              size="large"
              class="save-btn"
            >
              <v-icon class="me-2">{{ isEditing ? 'mdi-content-save' : 'mdi-plus' }}</v-icon>
              {{ isEditing ? 'تحديث' : 'إضافة' }}
                </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Delete Confirmation Dialog -->
      <v-dialog v-model="deleteDialog" max-width="400">
        <v-card>
          <v-card-title class="text-h6">تأكيد الحذف</v-card-title>
          <v-card-text>
            هل أنت متأكد من حذف المهندس "{{ selectedEngineer?.name }}"؟
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn color="grey" variant="text" @click="deleteDialog = false">
              إلغاء
            </v-btn>
            <v-btn color="error" variant="elevated" @click="confirmDelete">
                    حذف
                  </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Engineer Projects Dialog -->
      <v-dialog v-model="projectsDialog" max-width="800" class="projects-dialog">
        <v-card v-if="selectedEngineerForProjects">
          <v-card-title class="projects-dialog-title">
            <div class="d-flex align-center">
              <v-avatar size="32" color="success" variant="tonal" class="me-3">
                <v-icon>mdi-briefcase</v-icon>
              </v-avatar>
              <h2 class="text-h5 font-weight-bold mb-0">
                مشاريع المهندس - {{ selectedEngineerForProjects.name }}
              </h2>
            </div>
          </v-card-title>

          <v-card-text>
            <!-- Add Project Section -->
            <v-card variant="outlined" class="mb-4 pa-4">
              <v-card-title class="text-subtitle-1 pa-0 mb-3">
                <v-icon class="me-2" color="success" size="small">mdi-plus-circle</v-icon>
                إضافة مشروع جديد
              </v-card-title>
              <v-row>
                <v-col cols="12" md="8">
            <v-autocomplete
              v-model="selectedProject"
              :items="availableProjects"
              item-title="name"
              item-value="id"
              label="اختر المشروع"
              variant="outlined"
              density="comfortable"
              prepend-inner-icon="mdi-briefcase"
              class="form-field-enhanced enhanced-select black-dropdown-select"
              style="color: #000000 !important; font-size: 1.1rem !important; font-weight: 600 !important;"
              color="black"
              label-color="black"
              item-color="black"
              menu-color="white"
              :menu-props="{ 
                contentClass: 'black-dropdown-menu',
                attach: 'body',
                maxHeight: 300,
                zIndex: 50000
              }"
              bg-color="white"
            />
                </v-col>
                <v-col cols="12" md="4">
                  <v-btn
                    color="success"
                    variant="elevated"
                    @click="addProjectToEngineer"
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
            <v-card variant="outlined">
              <v-card-title class="text-subtitle-1 pa-3">
                <v-icon class="me-2" color="info" size="small">mdi-format-list-bulleted</v-icon>
                مشاريع المهندس الحالية
              </v-card-title>
              <v-card-text class="pa-0">
                <div v-if="selectedEngineerForProjects.projects && selectedEngineerForProjects.projects.length > 0">
                  <v-list>
                    <v-list-item
                      v-for="project in selectedEngineerForProjects.projects"
                      :key="project.id"
                      class="project-item"
                    >
                      <template v-slot:prepend>
                        <v-icon color="success">mdi-briefcase-check</v-icon>
                      </template>
                      <v-list-item-title class="font-weight-medium">
                        {{ project.name }}
                      </v-list-item-title>
                      <template v-slot:append>
                        <v-btn
                          icon="mdi-delete"
                          size="small"
                          variant="text"
                          color="error"
                          @click="removeProjectFromEngineer(project.id)"
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
          </v-card-text>

          <v-divider />
          <v-card-actions class="pa-4">
            <v-spacer />
            <v-btn color="grey" variant="text" @click="projectsDialog = false">
              إغلاق
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Engineer Details Dialog -->
      <v-dialog v-model="detailsDialog" max-width="700" class="engineer-details-dialog">
        <v-card v-if="selectedEngineer">
          <v-card-title class="d-flex align-center justify-space-between">
            <div class="d-flex align-center">
              <v-avatar size="40" color="primary" variant="tonal" class="me-3">
                <v-img v-if="selectedEngineer.avatar" :src="selectedEngineer.avatar" />
                <v-icon v-else>mdi-account</v-icon>
                  </v-avatar>
              <div>
                <h2 class="text-h5 font-weight-bold" style="color: white; text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);">{{ selectedEngineer.name }}</h2>
                <p class="text-subtitle-2 mb-0" style="color: rgba(255, 255, 255, 0.9); font-weight: 500;">{{ selectedEngineer.email }}</p>
                </div>
            </div>
            <v-btn icon="mdi-close" variant="text" @click="detailsDialog = false" style="color: white; background: rgba(255, 255, 255, 0.1); border-radius: 8px;" />
          </v-card-title>

          <v-divider />

          <v-card-text class="pa-6">
            <v-row>
              <!-- معلومات أساسية -->
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="pa-4">
                  <v-card-title class="text-h6 pa-0 mb-3">
                    <v-icon class="me-2" color="info" style="filter: drop-shadow(0 2px 4px rgba(59, 130, 246, 0.3));">mdi-information</v-icon>
                    <span style="color: #1e40af !important; font-weight: 700 !important; font-size: 1.1rem !important;">المعلومات الأساسية</span>
                  </v-card-title>
                  <div class="space-y-3">
                    <div class="d-flex justify-space-between">
                      <span class="font-weight-bold" style="color: #1e293b !important; font-size: 0.95rem;">الاسم:</span>
                      <span class="font-weight-semibold" style="color: #2563eb !important; font-size: 0.95rem;">{{ selectedEngineer.name }}</span>
                  </div>
                    <div class="d-flex justify-space-between">
                      <span class="font-weight-bold" style="color: #1e293b !important; font-size: 0.95rem;">البريد الإلكتروني:</span>
                      <span style="color: #334155 !important; font-size: 0.95rem;">{{ selectedEngineer.email }}</span>
                  </div>
                    <div class="d-flex justify-space-between">
                      <span class="font-weight-bold" style="color: #1e293b !important; font-size: 0.95rem;">رقم الهاتف:</span>
                      <span style="color: #334155 !important; font-size: 0.95rem;">{{ selectedEngineer.phone }}</span>
                  </div>
                    <div class="d-flex justify-space-between">
                      <span class="font-weight-bold" style="color: #1e293b !important; font-size: 0.95rem;">التخصص:</span>
                      <span style="color: #334155 !important; font-size: 0.95rem;">{{ selectedEngineer.specialization }}</span>
                  </div>
                    <div class="d-flex justify-space-between">
                      <span class="font-weight-bold" style="color: #1e293b !important; font-size: 0.95rem;">سنوات الخبرة:</span>
                      <span style="color: #334155 !important; font-size: 0.95rem;">{{ selectedEngineer.experience }} سنة</span>
                </div>
                  </div>
                </v-card>
              </v-col>

              <!-- التقييم والحالة -->
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="pa-4">
                  <v-card-title class="text-h6 pa-0 mb-3">
                    <v-icon class="me-2" color="warning" style="filter: drop-shadow(0 2px 4px rgba(245, 158, 11, 0.3));">mdi-star</v-icon>
                    <span style="color: #b45309 !important; font-weight: 700 !important; font-size: 1.1rem !important;">التقييم والحالة</span>
                  </v-card-title>
                  <div class="space-y-4">
                    <div class="d-flex align-center">
                      <span class="font-weight-bold me-3" style="color: #1e293b !important; font-size: 0.95rem;">التقييم:</span>
                      <v-rating
                        :model-value="selectedEngineer.rating"
                        readonly
                        size="small"
                        color="warning"
                        density="compact"
                      />
                      <span class="text-h6 ms-2" style="color: #f59e0b !important; font-weight: 700 !important;">{{ selectedEngineer.rating }}</span>
                    </div>
                    <div class="d-flex justify-space-between">
                      <span class="font-weight-bold" style="color: #1e293b !important; font-size: 0.95rem;">الحالة:</span>
                    <v-chip
                        :color="selectedEngineer.status === 'active' ? 'success' : 'error'"
                      size="small"
                        variant="tonal"
                    >
                        {{ selectedEngineer.status === 'active' ? 'نشط' : 'غير نشط' }}
                    </v-chip>
                  </div>
                </div>
                </v-card>
              </v-col>

              <!-- المهارات -->
              <v-col cols="12">
                <v-card variant="outlined" class="pa-4">
                  <v-card-title class="text-h6 pa-0 mb-3">
                    <v-icon class="me-2" color="success" style="filter: drop-shadow(0 2px 4px rgba(16, 185, 129, 0.3));">mdi-cog</v-icon>
                    <span style="color: #047857 !important; font-weight: 700 !important; font-size: 1.1rem !important;">المهارات</span>
                  </v-card-title>
                  <div class="d-flex flex-wrap gap-2">
                  <v-chip
                      v-for="skill in selectedEngineer.skills"
                      :key="skill"
                      color="primary"
                      variant="tonal"
                    size="small"
                  >
                      {{ skill }}
                  </v-chip>
                </div>
                </v-card>
              </v-col>
            </v-row>
          </v-card-text>

          <v-divider />

          <v-card-actions class="pa-4">
            <v-spacer />
            <v-btn color="primary" variant="elevated" @click="editEngineer(selectedEngineer)">
              <v-icon class="me-2">mdi-pencil</v-icon>
                    تعديل
                  </v-btn>
            <v-btn color="grey" variant="text" @click="detailsDialog = false">
              إغلاق
                  </v-btn>
          </v-card-actions>
            </v-card>
      </v-dialog>
      </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

// Reactive data
const engineers = ref([
  {
    id: 1,
    name: 'أحمد محمد',
    email: 'ahmed@example.com',
    phone: '0501234567',
    specialization: 'مهندس برمجيات',
    rating: 4.5,
    status: 'active',
    skills: ['JavaScript', 'Vue.js', 'Node.js', 'Python'],
    experience: '5 سنوات خبرة في تطوير التطبيقات',
    avatar: null
  },
  {
    id: 2,
    name: 'فاطمة علي',
    email: 'fatima@example.com',
    phone: '0507654321',
    specialization: 'مهندسة معمارية',
    rating: 4.8,
    status: 'active',
    skills: ['AutoCAD', 'Revit', 'SketchUp', 'Photoshop'],
    experience: '7 سنوات خبرة في التصميم المعماري',
    avatar: null
  },
  {
    id: 3,
    name: 'محمد السعيد',
    email: 'mohamed@example.com',
    phone: '0509876543',
    specialization: 'مهندس كهربائي',
    rating: 4.2,
    status: 'inactive',
    skills: ['PLC', 'Automation', 'Electrical Design'],
    experience: '3 سنوات خبرة في الهندسة الكهربائية',
    avatar: null
  }
])

const searchQuery = ref('')
const dialog = ref(false)
const deleteDialog = ref(false)
const detailsDialog = ref(false)
const isEditing = ref(false)
const formValid = ref(false)
const selectedEngineer = ref(null)
const showPassword = ref(false)
const projectsDialog = ref(false)
const selectedEngineerForProjects = ref(null)
const availableProjects = ref([
  { id: 1, name: 'مشروع تطوير النظام', status: 'active' },
  { id: 2, name: 'مشروع إدارة الموارد', status: 'active' },
  { id: 3, name: 'مشروع التصميم المعماري', status: 'active' },
  { id: 4, name: 'مشروع البنية التحتية', status: 'active' },
  { id: 5, name: 'مشروع الأمان السيبراني', status: 'active' }
])
const selectedProject = ref(null)

const engineerForm = ref({
  name: '',
  password: ''
})

const statusOptions = [
  { title: 'نشط', value: 'active' },
  { title: 'غير نشط', value: 'inactive' }
]

const tableHeaders = [
  { title: 'الصورة', key: 'avatar', sortable: false },
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'التقييم', key: 'rating', sortable: true },
  { title: 'المهارات', key: 'skills', sortable: false },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Computed properties
const filteredEngineers = computed(() => {
  if (!searchQuery.value) return engineers.value
  return engineers.value.filter(engineer =>
    engineer.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
    engineer.email.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
    engineer.specialization.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const activeEngineers = computed(() => {
  return engineers.value.filter(e => e.status === 'active').length
})

const totalProjects = computed(() => {
  return engineers.value.reduce((total, engineer) => total + Math.floor(Math.random() * 5) + 1, 0)
})

const averageRating = computed(() => {
  const total = engineers.value.reduce((sum, engineer) => sum + engineer.rating, 0)
  return (total / engineers.value.length).toFixed(1)
})

// Methods
const openAddEngineerDialog = () => {
  isEditing.value = false
  engineerForm.value = {
    name: '',
    password: ''
  }
  dialog.value = true
}

const viewEngineer = (engineer) => {
  selectedEngineer.value = engineer
  detailsDialog.value = true
}

const editEngineer = (engineer) => {
  isEditing.value = true
  engineerForm.value = { ...engineer }
  selectedEngineer.value = engineer
  dialog.value = true
}

const deleteEngineer = (engineer) => {
  selectedEngineer.value = engineer
  deleteDialog.value = true
}

const saveEngineer = () => {
  if (isEditing.value) {
    const index = engineers.value.findIndex(e => e.id === selectedEngineer.value.id)
    if (index !== -1) {
      engineers.value[index] = { 
        ...engineers.value[index],
        name: engineerForm.value.name,
        password: engineerForm.value.password
      }
    }
  } else {
    const newEngineer = {
      id: Date.now(),
      name: engineerForm.value.name,
      email: `${engineerForm.value.name.toLowerCase().replace(/\s+/g, '')}@example.com`,
      phone: '0500000000',
      specialization: 'مهندس',
      rating: 0,
      status: 'active',
      skills: [],
      experience: 'خبرة جديدة',
      avatar: null,
      password: engineerForm.value.password
    }
    engineers.value.push(newEngineer)
  }
  closeDialog()
}

const closeDialog = () => {
  dialog.value = false
  isEditing.value = false
  selectedEngineer.value = null
  engineerForm.value = {
    name: '',
    password: ''
  }
}

const confirmDelete = () => {
  if (selectedEngineer.value) {
    const index = engineers.value.findIndex(e => e.id === selectedEngineer.value.id)
    if (index !== -1) {
      engineers.value.splice(index, 1)
    }
  }
  deleteDialog.value = false
  selectedEngineer.value = null
}

const openProjectsDialog = (engineer) => {
  selectedEngineerForProjects.value = engineer
  selectedProject.value = null
  projectsDialog.value = true
}

const addProjectToEngineer = () => {
  if (selectedProject.value && selectedEngineerForProjects.value) {
    const engineerIndex = engineers.value.findIndex(e => e.id === selectedEngineerForProjects.value.id)
    if (engineerIndex !== -1) {
      if (!engineers.value[engineerIndex].projects) {
        engineers.value[engineerIndex].projects = []
      }
      
      const projectExists = engineers.value[engineerIndex].projects.some(p => p.id === selectedProject.value.id)
      if (!projectExists) {
        engineers.value[engineerIndex].projects.push(selectedProject.value)
      }
    }
    selectedProject.value = null
  }
}

const removeProjectFromEngineer = (projectId) => {
  if (selectedEngineerForProjects.value) {
    const engineerIndex = engineers.value.findIndex(e => e.id === selectedEngineerForProjects.value.id)
    if (engineerIndex !== -1 && engineers.value[engineerIndex].projects) {
      engineers.value[engineerIndex].projects = engineers.value[engineerIndex].projects.filter(p => p.id !== projectId)
    }
  }
}

// Lifecycle
onMounted(() => {
  // Initialize data if needed
})
</script>


<style>
/* Import page styles */
@import './styles/engineers.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
