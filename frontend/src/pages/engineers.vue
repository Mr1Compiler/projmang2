<template>
  <v-container class="fill-height engineers-page" max-width="1200" style="padding: 0;">
    <div class="centered-content" style="padding: 0 20px;">
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
              <v-icon size="32" color="white">mdi-account-hard-hat</v-icon>
            </div>
          </div>
        </div>
      </div>

        <!-- Summary Cards -->
        <v-row class="mb-2 stats-row">
            <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-account-group</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ engineers.length }}</h3>
                  <p class="stat-label">إجمالي المهندسين</p>
                </div>
              </div>
            </v-card>
            </v-col>
            <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-success" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon check-icon">mdi-check-circle</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ activeEngineers }}</h3>
                  <p class="stat-label">مهندسين نشطين</p>
                </div>
              </div>
            </v-card>
            </v-col>
            <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-info" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-briefcase</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ totalProjects }}</h3>
                  <p class="stat-label">إجمالي المشاريع</p>
                </div>
              </div>
            </v-card>
            </v-col>
            <v-col cols="12" md="3">
            <v-card class="modern-stat-card stat-card-warning" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-star</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ averageRating }}</h3>
                  <p class="stat-label">متوسط التقييم</p>
                </div>
              </div>
            </v-card>
            </v-col>
          </v-row>

      <!-- Engineers Table -->
      <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
        <v-card-title class="d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" style="color: white !important; font-size: 18px !important;">mdi-account-group</v-icon>
            <span class="text-h6 font-weight-bold" style="color: white !important; font-family: 'Arial', 'Helvetica', sans-serif; font-size: 0.95rem !important;">قائمة المهندسين</span>
            <v-chip class="ms-3" style="background-color: rgba(255, 255, 255, 0.3) !important; color: white !important; border: 1px solid rgba(255, 255, 255, 0.5); font-size: 0.75rem !important; height: 20px !important;" size="x-small">{{ engineers.length }}</v-chip>
        </div>
          <v-btn
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddEngineerDialog"
            elevation="2"
            color="primary"
            size="small"
            style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important; font-size: 0.875rem !important;"
          >
            <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
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
          :loading="loading"
          class="elevation-0"
          no-data-text="لا يوجد مهندسين"
          loading-text="جاري التحميل..."
          :items-per-page="-1"
          hide-default-footer
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

        <!-- Pagination -->
        <div class="d-flex justify-center pa-4 pagination-container" v-if="totalPages > 0">
          <v-pagination
            v-model="currentPage"
            :length="totalPages"
            :total-visible="7"
            @update:model-value="onPageChange"
            rounded="circle"
            color="primary"
          />
        </div>
      </v-card>

      <!-- Add/Edit Engineer Dialog - Profile Form Style (مثل نموذج إضافة المهمة) -->
      <v-dialog v-model="dialog" max-width="900">
        <v-card class="task-dialog-card profile-form-card">
          <v-card-title class="task-dialog-header profile-form-header">
            <h2 class="profile-form-title">
              {{ isEditing ? 'تعديل المهندس' : 'معلومات المهندس' }}
            </h2>
          </v-card-title>

          <v-card-text class="profile-form-content">
            <p class="profile-form-instruction">
              لإتمام {{ isEditing ? 'تعديل' : 'إضافة' }} حساب المهندس، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
            </p>

            <v-form ref="form" v-model="formValid">
              <!-- الصف الأول: اسم المستخدم، الاسم الكامل -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="6" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      اسم المستخدم <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="engineerForm.username"
                      variant="solo-filled"
                      flat
                      density="comfortable"
                      :rules="[
                        v => !!v || 'اسم المستخدم مطلوب',
                        v => (v && v.length >= 3) || 'اسم المستخدم يجب أن يكون 3 أحرف على الأقل',
                        v => /^[a-zA-Z0-9_]+$/.test(v) || 'اسم المستخدم يجب أن يحتوي على حروف وأرقام و _ فقط'
                      ]"
                      required
                      prepend-inner-icon="mdi-account"
                      hide-details="auto"
                      class="profile-form-input"
                      placeholder="اسم المستخدم لتسجيل الدخول"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="6" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      الاسم الكامل <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="engineerForm.fullName"
                      variant="solo-filled"
                      flat
                      density="comfortable"
                      :rules="[v => !!v || 'الاسم الكامل مطلوب']"
                      required
                      prepend-inner-icon="mdi-account-outline"
                      hide-details="auto"
                      class="profile-form-input"
                      placeholder="الاسم الكامل للمهندس"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الثاني: البريد الإلكتروني، رقم الهاتف -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="6" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      البريد الإلكتروني <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="engineerForm.email"
                      type="email"
                      variant="solo-filled"
                      flat
                      density="comfortable"
                      :rules="[
                        v => !!v || 'البريد الإلكتروني مطلوب',
                        v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح'
                      ]"
                      required
                      prepend-inner-icon="mdi-email"
                      hide-details="auto"
                      class="profile-form-input"
                      placeholder="example@domain.com"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="6" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      رقم الهاتف <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="engineerForm.phone"
                      variant="solo-filled"
                      flat
                      density="comfortable"
                      :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                      required
                      prepend-inner-icon="mdi-phone"
                      hide-details="auto"
                      class="profile-form-input"
                      placeholder="07XX XXX XXXX"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الثالث: كلمة المرور -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="6" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      كلمة المرور <span class="required-star" v-if="!isEditing">*</span>
                    </label>
                    <v-text-field
                      v-model="engineerForm.password"
                      :type="showPassword ? 'text' : 'password'"
                      variant="solo-filled"
                      flat
                      density="comfortable"
                      :rules="[
                        v => isEditing || !!v || 'كلمة المرور مطلوبة',
                        v => isEditing || (v && v.length >= 8) || 'كلمة المرور يجب أن تكون 8 أحرف على الأقل'
                      ]"
                      prepend-inner-icon="mdi-lock"
                      :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                      @click:append-inner="showPassword = !showPassword"
                      hide-details="auto"
                      class="profile-form-input"
                      placeholder="كلمة المرور (8 أحرف على الأقل)"
                    />
                  </div>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>

          <v-card-actions class="profile-form-actions">
            <v-spacer />
            <v-btn
              class="profile-form-cancel-btn"
              variant="outlined"
              @click="closeDialog"
              :disabled="saving"
            >
              إلغاء
            </v-btn>
            <v-btn
              class="profile-form-continue-btn"
              variant="elevated"
              :disabled="!formValid || saving"
              :loading="saving"
              @click="saveEngineer"
            >
              {{ isEditing ? 'تحديث' : 'حفظ' }}
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
import { listUsers, createUser, updateUser, deleteUser as deleteUserApi } from '@/api/users'
import { DEFAULT_LIMIT } from '@/constants/pagination'

// Reactive data - يتم تحميل البيانات من API المستخدمين
const engineers = ref([])
const loading = ref(false)

// Pagination state
const currentPage = ref(1)
const totalItems = ref(0)
const totalPages = ref(0)

const searchQuery = ref('')
const dialog = ref(false)
const deleteDialog = ref(false)
const detailsDialog = ref(false)
const isEditing = ref(false)
const formValid = ref(false)
const form = ref(null)
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

// نموذج المهندس (متوافق مع Backend CreateUser DTO)
// المسمى الوظيفي يتم تعيينه تلقائياً كـ "مهندس"
const engineerForm = ref({
  username: '',
  fullName: '',
  email: '',
  phone: '',
  password: ''
})
const saving = ref(false)

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
  // عدد المشاريع المرتبطة بالمهندسين
  return engineers.value.reduce((total, engineer) => {
    return total + (engineer.projects?.length || 0)
  }, 0)
})

const averageRating = computed(() => {
  if (engineers.value.length === 0) return '0.0'
  const total = engineers.value.reduce((sum, engineer) => sum + (engineer.rating || 0), 0)
  return (total / engineers.value.length).toFixed(1)
})

// Methods
const openAddEngineerDialog = () => {
  isEditing.value = false
  engineerForm.value = {
    username: '',
    fullName: '',
    email: '',
    phone: '',
    password: ''
  }
  dialog.value = true
}

const viewEngineer = (engineer) => {
  console.log('View engineer:', engineer)
  selectedEngineer.value = engineer
  detailsDialog.value = true
}

const editEngineer = (engineer) => {
  isEditing.value = true
  engineerForm.value = {
    username: engineer.username || '',
    fullName: engineer.fullName || engineer.name || '',
    email: engineer.email || '',
    phone: engineer.phone || '',
    password: ''
  }
  selectedEngineer.value = engineer
  dialog.value = true
}

const deleteEngineer = (engineer) => {
  selectedEngineer.value = engineer
  deleteDialog.value = true
}

const saveEngineer = async () => {
  if (!form.value.validate()) return

  saving.value = true
  try {
    if (isEditing.value) {
      // تحديث المهندس
      const userData = {
        username: engineerForm.value.username || undefined,
        fullName: engineerForm.value.fullName || undefined,
        email: engineerForm.value.email || undefined,
        phone: engineerForm.value.phone || undefined
      }
      await updateUser(selectedEngineer.value.id, userData)
    } else {
      // إضافة مهندس جديد - المسمى الوظيفي "مهندس" تلقائياً
      const userData = {
        username: engineerForm.value.username,
        fullName: engineerForm.value.fullName,
        email: engineerForm.value.email,
        phone: engineerForm.value.phone,
        password: engineerForm.value.password,
        jobTitle: 'مهندس'  // يتم تعيينه تلقائياً
      }
      await createUser(userData)
    }

    // إعادة تحميل البيانات
    await loadEngineers()
    closeDialog()
    console.log('تم حفظ المهندس بنجاح')
  } catch (error) {
    console.error('خطأ في حفظ المهندس:', error)
  } finally {
    saving.value = false
  }
}

const closeDialog = () => {
  dialog.value = false
  isEditing.value = false
  selectedEngineer.value = null
  engineerForm.value = {
    username: '',
    fullName: '',
    email: '',
    phone: '',
    password: ''
  }
}

const confirmDelete = async () => {
  if (selectedEngineer.value) {
    try {
      await deleteUserApi(selectedEngineer.value.id)
      await loadEngineers()
      console.log('تم حذف المهندس بنجاح')
    } catch (error) {
      console.error('خطأ في حذف المهندس:', error)
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

// ========================================
// دالة تحميل المهندسين من API المستخدمين
// ========================================

// دالة للتحقق إذا كان المستخدم مهندس بناءً على الدور أو المسمى الوظيفي
const isEngineer = (user) => {
  // أولاً: التحقق من حقل الدور إذا كان موجوداً
  if (user.role) {
    const role = String(user.role).toLowerCase()
    if (role === 'engineer' || role === 'مهندس') return true
  }

  // ثانياً: التحقق من المسمى الوظيفي
  if (!user.jobTitle) return false
  const title = user.jobTitle.toLowerCase()
  // التحقق من الكلمات المفتاحية للمهندسين بالعربية والإنجليزية
  const engineerKeywords = [
    'engineer', 'مهندس', 'مهندسة', 'هندسة',
    'engineering', 'eng.', 'eng '
  ]
  return engineerKeywords.some(keyword => title.includes(keyword))
}

// جميع المهندسين (بعد الفلترة)
const allEngineers = ref([])

const loadEngineers = async () => {
  loading.value = true
  try {
    // جلب جميع المستخدمين (بحد أقصى 100) لأن الفلترة تتم على الواجهة
    const response = await listUsers({ page: 1, limit: 100 })
    console.log('Users data received for engineers:', response)

    if (response.data) {
      // فلترة المستخدمين الذين هم مهندسين فقط بناءً على الدور أو المسمى الوظيفي
      const engineerUsers = response.data.filter(user => isEngineer(user))

      // تحويل بيانات المستخدمين إلى تنسيق المهندسين
      allEngineers.value = engineerUsers.map(user => ({
        id: user.id,
        name: user.fullName || '',
        email: user.email || '',
        phone: user.phone || '',
        specialization: user.jobTitle || 'مهندس',
        rating: 0,
        status: user.status || 'active',
        skills: [],
        experience: '',
        avatar: user.avatar || null,
        // حفظ البيانات الأصلية للتعديل
        username: user.userName || user.username || '',
        fullName: user.fullName || '',
        jobTitle: user.jobTitle || ''
      }))

      // حساب عدد الصفحات بناءً على عدد المهندسين المفلترين
      totalItems.value = allEngineers.value.length
      totalPages.value = Math.ceil(allEngineers.value.length / DEFAULT_LIMIT)
      currentPage.value = 1

      // عرض الصفحة الأولى
      updateDisplayedEngineers()
    }
  } catch (error) {
    console.error('Error loading engineers:', error)
  } finally {
    loading.value = false
  }
}

// تحديث المهندسين المعروضين بناءً على الصفحة الحالية
const updateDisplayedEngineers = () => {
  const start = (currentPage.value - 1) * DEFAULT_LIMIT
  const end = start + DEFAULT_LIMIT
  engineers.value = allEngineers.value.slice(start, end)
}

// Handle page change
const onPageChange = (page) => {
  currentPage.value = page
  updateDisplayedEngineers()
}

// Lifecycle
onMounted(() => {
  loadEngineers()
})
</script>

<style scoped>
/* ========================================
   صفحة المهندسين - CSS منظم ومتناسق
   ======================================== */

/* تنسيق نموذج إضافة المهندس - نفس نمط نموذج إضافة المهمة */
.profile-form-card {
  border-radius: 8px !important;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.12) !important;
  overflow: visible !important;
  background: #ffffff !important;
  border: 2px solid #d1d5db !important;
}

.profile-form-header {
  background: #ffffff !important;
  color: #1e293b !important;
  padding: 20px 24px !important;
  border-bottom: 1px solid #e2e8f0 !important;
  box-shadow: none !important;
}

.profile-form-title {
  font-size: 1.35rem !important;
  font-weight: 700 !important;
  color: #1e293b !important;
  margin: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: none !important;
}

.profile-form-content {
  padding: 24px !important;
  background: #ffffff !important;
  border: 2px solid #d1d5db !important;
  border-radius: 8px !important;
  margin: 20px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}

.profile-form-instruction {
  font-size: 0.85rem !important;
  color: #64748b !important;
  margin-bottom: 24px !important;
  line-height: 1.6 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.profile-form-row {
  margin: 0 !important;
}

.profile-form-column {
  padding: 0 10px !important;
  display: flex !important;
  flex-direction: column !important;
}

.profile-form-column:first-child {
  padding-right: 0 !important;
}

.profile-form-column:last-child {
  padding-left: 0 !important;
}

.profile-form-field-wrapper {
  margin-bottom: 16px !important;
  padding: 0 !important;
  background: transparent !important;
  border: none !important;
  border-radius: 0 !important;
  transition: none !important;
  box-shadow: none !important;
  flex: 1 !important;
}

.profile-form-field-wrapper:hover {
  border-color: transparent !important;
  box-shadow: none !important;
}

.profile-form-field-wrapper:focus-within {
  border-color: transparent !important;
  box-shadow: none !important;
}

.profile-form-label {
  display: block !important;
  font-size: 0.8125rem !important;
  font-weight: 600 !important;
  color: #374151 !important;
  margin-bottom: 6px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-align: right !important;
}

.required-star {
  color: #ef4444 !important;
  font-weight: 700 !important;
  margin-right: 2px !important;
}

.profile-form-input {
  width: 100% !important;
}

.profile-form-input :deep(.v-field) {
  background: #f9fafb !important;
  border-radius: 6px !important;
  border: 1.5px solid #d1d5db !important;
  transition: all 0.2s ease !important;
}

.profile-form-input :deep(.v-field:hover) {
  border-color: #9ca3af !important;
}

.profile-form-input :deep(.v-field--focused) {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.profile-form-input :deep(.v-field__outline) {
  border-color: #d1d5db !important;
  border-width: 1.5px !important;
}

.profile-form-input :deep(.v-field--focused .v-field__outline) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.profile-form-input :deep(.v-field__outline__start),
.profile-form-input :deep(.v-field__outline__notch),
.profile-form-input :deep(.v-field__outline__end) {
  border-color: #d1d5db !important;
  border-width: 1.5px !important;
}

.profile-form-input :deep(.v-field--focused .v-field__outline__start),
.profile-form-input :deep(.v-field--focused .v-field__outline__notch),
.profile-form-input :deep(.v-field--focused .v-field__outline__end) {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.profile-form-input :deep(.v-field__input) {
  color: #1e293b !important;
  font-size: 0.875rem !important;
  padding: 10px 12px !important;
  text-align: right !important;
  direction: rtl !important;
}

.profile-form-input :deep(.v-label),
.profile-form-input :deep(.v-field-label) {
  color: #6b7280 !important;
  font-size: 0.875rem !important;
  text-align: right !important;
  right: 0 !important;
  left: auto !important;
}

.profile-form-input :deep(.v-field__prepend-inner),
.profile-form-input :deep(.v-field__append-inner) {
  padding-right: 12px !important;
}

.profile-form-input :deep(.v-select__selection) {
  text-align: right !important;
  direction: rtl !important;
  align-items: center !important;
  min-height: 40px !important;
  font-weight: 500 !important;
}

.profile-form-input :deep(.v-field__append-inner .v-icon) {
  color: #1e40af !important;
  opacity: 1 !important;
}

.profile-form-input :deep(.v-overlay__content) {
  border-radius: 10px !important;
  box-shadow:
    0 10px 30px rgba(15, 23, 42, 0.18),
    0 4px 10px rgba(148, 163, 184, 0.25) !important;
  border: 1px solid #e5e7eb !important;
  overflow: hidden !important;
  background: #ffffff !important;
}

.profile-form-input :deep(.v-overlay__content .v-list) {
  padding: 4px 0 !important;
}

.profile-form-input :deep(.v-overlay__content .v-list-item) {
  min-height: 34px !important;
  padding-inline: 12px !important;
}

.profile-form-input :deep(.v-overlay__content .v-list-item-title) {
  font-size: 0.85rem !important;
  color: #0f172a !important;
  text-align: right !important;
}

.profile-form-input :deep(.v-overlay__content .v-list-item--active),
.profile-form-input :deep(.v-overlay__content .v-list-item:hover) {
  background: #eff6ff !important;
  color: #1d4ed8 !important;
}

.profile-form-actions {
  padding: 24px 32px !important;
  background: #ffffff !important;
  border-top: 1px solid #e2e8f0 !important;
  display: flex !important;
  justify-content: flex-end !important;
  gap: 12px !important;
  visibility: visible !important;
  opacity: 1 !important;
  z-index: 10 !important;
  position: relative !important;
}

.profile-form-actions .v-btn {
  display: inline-flex !important;
  visibility: visible !important;
  opacity: 1 !important;
  min-height: 44px !important;
  height: auto !important;
}

.profile-form-cancel-btn {
  min-width: 120px !important;
  height: 44px !important;
  border-radius: 6px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  color: #6b7280 !important;
  border-color: #d1d5db !important;
  background: #ffffff !important;
  visibility: visible !important;
  opacity: 1 !important;
  display: inline-flex !important;
  z-index: 10 !important;
  position: relative !important;
  border-width: 1.5px !important;
  border-style: solid !important;
}

.profile-form-cancel-btn .v-btn__content {
  visibility: visible !important;
  opacity: 1 !important;
}

.profile-form-cancel-btn:hover {
  background: #f9fafb !important;
  border-color: #9ca3af !important;
}

.profile-form-continue-btn {
  min-width: 140px !important;
  height: 44px !important;
  border-radius: 6px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  background: #10b981 !important;
  color: #ffffff !important;
  box-shadow: 0 2px 8px rgba(16, 185, 129, 0.3) !important;
  visibility: visible !important;
  opacity: 1 !important;
  display: inline-flex !important;
  z-index: 10 !important;
  position: relative !important;
  border: none !important;
}

.profile-form-continue-btn .v-btn__content {
  visibility: visible !important;
  opacity: 1 !important;
  display: flex !important;
  align-items: center !important;
  gap: 8px !important;
}

.profile-form-continue-btn span {
  visibility: visible !important;
  opacity: 1 !important;
  display: inline-block !important;
}

.profile-form-continue-btn:hover {
  background: #059669 !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4) !important;
  transform: translateY(-1px) !important;
}

.profile-form-continue-btn:disabled {
  background: #d1d5db !important;
  color: #9ca3af !important;
  box-shadow: none !important;
}

.profile-form-continue-btn .v-icon {
  color: #ffffff !important;
}

/* صفحة المهندسين */
.engineers-page {
  background: #ffffff !important;
  color: #1e293b;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0;
  overflow-x: hidden;
}

/* المحتوى المركزي */
.centered-content {
  width: 100%;
  max-width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2.5rem;
  text-align: center;
}

/* بطاقة رأس المهندسين */
.engineers-header-card {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
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

/* الأيقونة الثانوية */
.secondary-icon {
  animation: slideInFromRight 1s ease-out 1.2s both, float 3s ease-in-out infinite 2.5s, pulse 2s ease-in-out infinite 2.5s;
  opacity: 0.8;
  transform: translateX(10px);
}

.secondary-icon:hover {
  transform: scale(1.15) rotate(-3deg) translateX(10px);
  opacity: 1;
  filter: drop-shadow(0 10px 20px rgba(255, 255, 255, 0.5));
}

/* الأيقونة الثالثة */
.tertiary-icon {
  animation: slideInFromRight 1s ease-out 1.5s both, float 3s ease-in-out infinite 3s, pulse 2s ease-in-out infinite 3s;
  opacity: 0.7;
  transform: translateX(20px);
}

.tertiary-icon:hover {
  transform: scale(1.2) rotate(3deg) translateX(20px);
  opacity: 1;
  filter: drop-shadow(0 8px 16px rgba(255, 255, 255, 0.6));
}

/* النصوص */
.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.8rem;
  position: relative;
  animation: slideInFromLeft 1s ease-out 0.6s both;
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

/* صف الإحصائيات */
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

/* تصميم متجاوب للرأس */
@media (max-width: 768px) {
  .engineers-header-card {
    max-width: 100%;
    margin: 0 auto 1.5rem auto;
  }
  
  .header-content {
    padding: 1rem;
    min-height: 70px;
  }
  
  .header-right {
    gap: 1rem;
    padding: 0.6rem 1rem;
  }
  
  .engineer-emoji {
    font-size: 2rem;
  }
  
  .main-title {
    font-size: 1.1rem !important;
  }
  
  .subtitle {
    font-size: 0.7rem !important;
  }
}

@media (max-width: 480px) {
  .header-content {
    padding: 1rem;
    flex-direction: column;
    text-align: center;
  }
  
  .header-right {
    flex-direction: column;
    gap: 0.8rem;
  }
  
  .engineer-emoji {
    font-size: 2rem;
  }
  
  .main-title {
    font-size: 1rem !important;
  }
  
  .subtitle {
    font-size: 0.65rem !important;
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

/* دعم الكروتات القديمة */
.stat-card {
  background: linear-gradient(145deg, #9e7272 0%, #f8fafc 100%) !important;
  border-radius: 12px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  height: 100% !important;
  min-height: 140px !important;
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

.stat-card::before {
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

.stat-card::after {
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

.stat-card:hover {
  transform: translateY(-12px) scale(1.05);
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2), 0 8px 16px rgba(0, 0, 0, 0.1) !important;
  border-color: rgba(219, 206, 206, 0.9);
  background: linear-gradient(145deg, #fffcfc 0%, #f1f5f9 100%) !important;
}

.stat-card:hover::before {
  height: 6px;
  opacity: 1;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.4);
}

.stat-card:hover::after {
  opacity: 1;
}

.stat-icon {
  position: relative;
  z-index: 2;
  margin-bottom: 0.5rem;
  padding: 8px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 241, 241, 0.8) 0%, rgba(248, 250, 252, 0.8) 100%);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.stat-card .v-icon {
  filter: drop-shadow(0 6px 12px rgba(0, 0, 0, 0.15));
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 36px !important;
  width: 36px !important;
  height: 36px !important;
}

.stat-card:hover .v-icon {
  transform: scale(1.15) rotate(8deg);
  filter: drop-shadow(0 12px 24px rgba(0, 0, 0, 0.25));
}

.stat-card:hover .stat-icon {
  transform: scale(1.1);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  background: linear-gradient(135deg, rgba(242, 239, 239, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%);
}

.stat-card h3 {
  font-size: 1.5rem !important;
  font-weight: 800 !important;
  margin-bottom: 10px !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.15);
  text-align: center !important;
  width: 100%;
  line-height: 1.1;
  position: relative;
  z-index: 2;
  color: #1976d2 !important;
  text-align: center !important;
  width: 100%;
  line-height: 1.1;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  position: relative;
  z-index: 2;
  transition: all 0.5s ease;
  letter-spacing: -1px;
}


.stat-card p,
.stat-card .text-subtitle-1 {
  font-size: 1rem !important;
  font-weight: 600 !important;
  opacity: 1 !important;
  text-align: center !important;
  width: 100%;
  margin: 0;
  color: #424242 !important;
  position: relative;
  z-index: 3;
  letter-spacing: 0.5px;
  visibility: visible !important;
  display: block !important;
}

.stat-card:hover p {
  color: #475569 !important;
  transform: translateY(-2px);
}

/* تأثيرات خاصة لكل لون */
.stat-card:nth-child(1)::before {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 50%, #7c3aed 100%);
}

.stat-card:nth-child(1):hover {
  box-shadow: 0 25px 50px rgba(59, 130, 246, 0.25), 0 8px 16px rgba(59, 130, 246, 0.15) !important;
}

.stat-card:nth-child(1) h3 {
  color: #3b82f6 !important;
}

.stat-card:nth-child(2)::before {
  background: linear-gradient(135deg, #10b981 0%, #059669 50%, #047857 100%);
}

.stat-card:nth-child(2):hover {
  box-shadow: 0 25px 50px rgba(16, 185, 129, 0.25), 0 8px 16px rgba(16, 185, 129, 0.15) !important;
}

.stat-card:nth-child(2) h3 {
  color: #10b981 !important;
}

.stat-card:nth-child(2) .stat-icon {
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.15) 0%, rgba(52, 211, 153, 0.2) 100%);
}

.stat-card:nth-child(2) .v-icon {
  background: linear-gradient(135deg, #10b981 0%, #34d399 50%, #6ee7b7 100%) !important;
  -webkit-background-clip: text !important;
  -webkit-text-fill-color: transparent !important;
  background-clip: text !important;
}

.stat-card:nth-child(2) h3 {
  background: linear-gradient(135deg, #059669 0%, #10b981 50%, #34d399 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.stat-card:nth-child(3)::before {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #b45309 100%);
}

.stat-card:nth-child(3):hover {
  box-shadow: 0 25px 50px rgba(245, 158, 11, 0.25), 0 8px 16px rgba(245, 158, 11, 0.15) !important;
}

.stat-card:nth-child(3) h3 {
  color: #f59e0b !important;
}

.stat-card:nth-child(4)::before {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #b45309 100%);
}

.stat-card:nth-child(4):hover {
  box-shadow: 0 25px 50px rgba(245, 158, 11, 0.25), 0 8px 16px rgba(245, 158, 11, 0.15) !important;
}

.stat-card:nth-child(4) h3 {
  color: #f59e0b !important;
}

/* بطاقة الجدول - نفس تصميم صفحة المشاريع */
.data-table-card {
  border-radius: 20px !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.data-table-card:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15) !important;
}

/* الجدول المركزي */
.centered-table {
  width: 100%;
  max-width: 100%;
  margin: 0 auto;
  border-radius: 20px !important;
  overflow: hidden;
  padding: 0 1rem;
}

.data-table-card:hover {
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12) !important;
}

.data-table-card .v-card-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-bottom: 1px solid rgba(37, 99, 235, 0.3);
  padding: 8px 12px !important;
  position: relative;
  color: white !important;
  min-height: 36px !important;
}

.data-table-card .v-card-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px !important;
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.5) 0%, rgba(255, 255, 255, 0.8) 50%, rgba(255, 255, 255, 0.5) 100%);
  box-shadow: 0 1px 4px rgba(255, 255, 255, 0.25);
}

.data-table-card .v-data-table {
  background: transparent !important;
}

.data-table-card .v-data-table-header {
  /* رأس جدول محايد بدون لون أزرق */
  background: #f8fafc !important;
  border-bottom: 1px solid #ffffff !important;
  font-weight: 600 !important;
  min-height: 32px !important;
  height: auto !important;
  padding: 0 !important;
}

/* تنسيق رأس الجدول - تصغير الخط والحجم */
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
  font-size: 0.65rem !important;
  font-weight: 600 !important;
  color: #1e293b !important;
  -webkit-text-fill-color: #1e293b !important;
  text-align: center !important;
  padding: 6px 4px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.2px !important;
  line-height: 1.2 !important;
  height: auto !important;
  min-height: 32px !important;
  vertical-align: middle !important;
  white-space: nowrap !important;
}

.data-table-card tbody tr:hover {
  background: #f1f5f9 !important;
  transform: scale(1.01);
  transition: all 0.2s ease;
}

.data-table-card tbody tr {
  transition: all 0.2s ease;
}

/* جعل خطوط الجدول بيضاء */
.data-table-card :deep(.v-data-table__wrapper table),
.data-table-card :deep(.v-data-table__wrapper table td),
.data-table-card :deep(.v-data-table__wrapper table th),
.data-table-card :deep(.v-data-table__wrapper table tr),
.data-table-card :deep(.v-data-table__td),
.data-table-card :deep(.v-data-table__th) {
  border-color: #ffffff !important;
}

.data-table-card :deep(.v-data-table__wrapper table td),
.data-table-card :deep(.v-data-table__td) {
  border-right: 1px solid #ffffff !important;
  border-bottom: 1px solid #ffffff !important;
}

.data-table-card :deep(.v-data-table__wrapper table th),
.data-table-card :deep(.v-data-table__th) {
  border-right: 1px solid #ffffff !important;
  border-bottom: 1px solid #ffffff !important;
}

/* زر الإضافة */
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

/* النوافذ المنبثقة */
.v-dialog .v-card {
  border-radius: 16px !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.v-dialog .v-card-title {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-bottom: 1px solid #e2e8f0;
  padding: 1.5rem;
  position: relative;
}

.v-dialog .v-card-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: var(--gradient-primary);
}

.dialog-actions .v-btn {
  border-radius: 8px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  transition: all 0.3s ease !important;
}

.dialog-actions .v-btn:hover {
  transform: translateY(-1px) !important;
}

/* تأثيرات النصوص */
.text-primary {
  color: #3b82f6 !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.text-success {
  color: #10b981 !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.text-info {
  color: #06b6d4 !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.text-warning {
  color: #f59e0b !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

/* تأثيرات إضافية للعناصر */
.v-chip {
  transition: all 0.3s ease !important;
}

.v-chip:hover {
  transform: scale(1.05) !important;
}

.v-avatar {
  transition: all 0.3s ease !important;
}

.v-avatar:hover {
  transform: scale(1.1) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

/* تصميم متجاوب */
@media (max-width: 960px) {
  .engineers-page {
    padding: 1.5rem 1rem;
  }
  
  .centered-content {
    gap: 2rem;
  }
  
  .stats-row {
    max-width: 100%;
    gap: 1rem;
    padding: 0 0.5rem;
  }
  
  .stat-card {
    max-width: 220px;
    min-height: 200px;
  }
  
  .stat-card h3 {
    font-size: 2.2rem !important;
  }
  
  .stat-card .stat-icon .v-icon {
    font-size: 60px !important;
    width: 60px !important;
    height: 60px !important;
  }
  
  .centered-table {
    max-width: 100%;
    padding: 0 0.5rem;
  }
}

@media (max-width: 600px) {
  .engineers-page {
    padding: 1rem 0.5rem;
  }
  
  .centered-content {
    gap: 1.5rem;
  }
  
  .stats-row {
    max-width: 100%;
    gap: 0.8rem;
    flex-wrap: wrap;
    padding: 0 0.25rem;
  }
  
  .stat-card {
    max-width: 180px;
    min-height: 180px;
    flex: 0 0 calc(50% - 0.4rem);
  }
  
  .stat-card h3 {
    font-size: 2rem !important;
  }
  
  .stat-card .stat-icon .v-icon {
    font-size: 52px !important;
    width: 52px !important;
    height: 52px !important;
  }
  
  .add-button {
    padding: 10px 20px !important;
    font-size: 1rem !important;
  }
  
  .centered-table {
    max-width: 100%;
    padding: 0 0.25rem;
  }
}

/* شاشات صغيرة جداً */
@media (max-width: 480px) {
  .stats-row {
    flex-direction: column;
    gap: 1rem;
  }
  
  .stat-card {
    max-width: 100%;
    flex: none;
    min-height: 160px;
  }
}

/* تحسينات عامة */
.v-card {
  border-radius: 16px !important;
  transition: all 0.3s ease !important;
}

.v-btn {
  border-radius: 8px !important;
  text-transform: none !important;
  transition: all 0.3s ease !important;
}

.v-text-field,
.v-select,
.v-textarea {
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

.v-text-field:focus-within,
.v-select:focus-within,
.v-textarea:focus-within {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.v-chip {
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

/* تأثيرات إضافية للجدول */
.v-data-table tbody tr {
  transition: all 0.2s ease;
}

.v-data-table tbody tr:hover {
  background: #f1f5f9 !important;
  transform: scale(1.01);
}

/* تنسيقات حقل البحث المحسن */
.search-field-enhanced {
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%) !important;
  border-radius: 16px !important;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15) !important;
  margin: 16px 0 !important;
}

.search-field-enhanced .v-field {
  border: 3px solid #3b82f6 !important;
  border-radius: 16px !important;
  background: white !important;
  transition: all 0.4s ease !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
}

.search-field-enhanced .v-field:focus-within {
  border-color: #1d4ed8 !important;
  box-shadow: 0 0 0 4px rgba(29, 78, 216, 0.15) !important;
  transform: translateY(-3px) !important;
  background: #fefefe !important;
}

.search-field-enhanced .v-field__input {
  color: #0f172a !important;
  font-weight: 600 !important;
  font-size: 1.2rem !important;
  padding: 16px !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.search-field-enhanced .v-label {
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 1.1rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.search-field-enhanced .v-field__prepend-inner {
  color: #1d4ed8 !important;
  font-size: 1.4rem !important;
}

.search-field-enhanced .v-field__append-inner {
  color: #64748b !important;
}

.search-field-enhanced .v-messages {
  color: #64748b !important;
  font-weight: 500 !important;
  font-size: 0.95rem !important;
}

/* تأثيرات إضافية للوضوح */
.search-field-enhanced:hover .v-field {
  border-color: #1d4ed8 !important;
  box-shadow: 0 4px 16px rgba(29, 78, 216, 0.25) !important;
}

/* تنسيق النص التوضيحي */
.search-hint-text {
  color: #475569 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  margin-top: 8px !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* تحسين وضوح النص المدخل */
.search-field-enhanced .v-field__input::placeholder {
  color: #94a3b8 !important;
  font-weight: 500 !important;
  font-size: 1.1rem !important;
  opacity: 0.8 !important;
}

/* تأثيرات للتصنيف */
.v-rating .v-rating__wrapper {
  transition: all 0.3s ease !important;
}

.v-rating .v-rating__wrapper:hover {
  transform: scale(1.1);
}

/* تنسيقات نافذة التفاصيل */
.space-y-3 > * + * {
  margin-top: 12px !important;
}

.space-y-4 > * + * {
  margin-top: 16px !important;
}

/* تنسيقات النافذة المنبثقة */
.v-dialog .v-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border-radius: 20px !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
  border: 1px solid rgba(59, 130, 246, 0.1) !important;
}

.v-dialog .v-card-title {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 20px 24px !important;
  margin: 0 !important;
}

.v-dialog .v-card-title h2 {
  color: white !important;
  font-weight: 700 !important;
  font-size: 1.5rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
}

.v-dialog .v-card-title p {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500 !important;
}

.v-dialog .v-avatar {
  border: 3px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

/* تنسيقات المحتوى */
.v-dialog .v-card-text {
  background: #ffffff !important;
  padding: 24px !important;
}

.v-dialog .v-card.v-card--variant-outlined {
  border: 2px solid #e2e8f0 !important;
  border-radius: 16px !important;
  background: linear-gradient(135deg, #ffffff 0%, #fefefe 100%) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}

/* عناوين الأقسام */
.v-dialog .v-card-title.text-h6 {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%) !important;
  color: #1e293b !important;
  border-radius: 12px 12px 0 0 !important;
  padding: 12px 16px !important;
  margin: -16px -16px 16px -16px !important;
  font-weight: 600 !important;
  font-size: 1.1rem !important;
}

/* الأيقونات */
.v-dialog .v-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1)) !important;
}

/* النصوص */
.v-dialog .font-weight-medium {
  color: #374151 !important;
  font-weight: 600 !important;
}

.v-dialog .text-primary {
  color: #1d4ed8 !important;
  font-weight: 600 !important;
}

/* التقييم */
.v-dialog .v-rating {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1)) !important;
}

/* الشارات */
.v-dialog .v-chip {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
  font-weight: 500 !important;
}

.v-dialog .v-chip--variant-tonal {
  border: 1px solid rgba(59, 130, 246, 0.2) !important;
}

/* الأزرار */
.v-dialog .v-card-actions {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-radius: 0 0 20px 20px !important;
  padding: 20px 24px !important;
}

.v-dialog .v-btn--variant-elevated {
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  border-radius: 12px !important;
  font-weight: 600 !important;
}

.v-dialog .v-btn--variant-text {
  border-radius: 8px !important;
  font-weight: 500 !important;
}

/* تحسينات إضافية للألوان */
.v-dialog .v-divider {
  border-color: rgba(59, 130, 246, 0.2) !important;
  margin: 0 !important;
}

/* تأثيرات hover */
.v-dialog .v-card:hover {
  transform: translateY(-2px) !important;
  transition: all 0.3s ease !important;
}

/* تحسين ألوان الشارات */
.v-dialog .v-chip--color-success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
}

.v-dialog .v-chip--color-error {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
}

.v-dialog .v-chip--color-primary {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
}

/* تحسين ألوان النجوم */
.v-dialog .v-rating .v-icon {
  color: #f59e0b !important;
  filter: drop-shadow(0 2px 4px rgba(245, 158, 11, 0.3)) !important;
}

/* تنسيقات نافذة التعديل */
.edit-engineer-dialog .v-card {
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

/* تنسيقات الحقول */
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

/* تنسيقات البطاقات الداخلية */
.rating-card, .skills-card, .experience-card {
  background: linear-gradient(135deg, #ffffff 0%, #fefefe 100%) !important;
  border: 2px solid #e2e8f0 !important;
  border-radius: 16px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}

.rating-card .v-card-title, .skills-card .v-card-title, .experience-card .v-card-title {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%) !important;
  color: #1e293b !important;
  border-radius: 12px 12px 0 0 !important;
  padding: 8px 12px !important;
  margin: -12px -12px 12px -12px !important;
  font-weight: 600 !important;
  font-size: 0.9rem !important;
}

/* تنسيقات الأزرار */
.dialog-actions {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-radius: 0 0 20px 20px !important;
}

.save-btn,
.save-btn.v-btn,
.v-btn.save-btn {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 50%, #3b82f6 100%) !important;
  box-shadow: 
    0 4px 16px rgba(30, 64, 175, 0.3),
    0 2px 8px rgba(37, 99, 235, 0.2),
    0 0 20px rgba(59, 130, 246, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border-radius: 12px !important;
  font-weight: 700 !important;
  padding: 14px 28px !important;
  color: white !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  position: relative;
  overflow: hidden;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  text-align: center !important;
  min-width: 140px !important;
  animation: saveButtonGlow 2s ease-in-out infinite !important;
}

@keyframes saveButtonGlow {
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

.save-btn::before {
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

.save-btn:hover::before {
  left: 100%;
}

.save-btn::after {
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

.save-btn:active::after {
  width: 300px;
  height: 300px;
  opacity: 0;
  transition: width 0.3s ease, height 0.3s ease, opacity 0.3s ease;
}

.save-btn:hover,
.save-btn.v-btn:hover,
.v-btn.save-btn:hover {
  transform: translateY(-3px) scale(1.03) !important;
  box-shadow: 
    0 8px 24px rgba(30, 64, 175, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    0 0 40px rgba(59, 130, 246, 0.7),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #2563eb 100%) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
  animation: saveButtonGlowHover 1.5s ease-in-out infinite !important;
}

@keyframes saveButtonGlowHover {
  0%, 100% {
    box-shadow: 
      0 8px 24px rgba(30, 64, 175, 0.5),
      0 4px 12px rgba(37, 99, 235, 0.4),
      0 0 40px rgba(59, 130, 246, 0.7),
      inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  }
  50% {
    box-shadow: 
      0 8px 28px rgba(30, 64, 175, 0.6),
      0 4px 16px rgba(37, 99, 235, 0.5),
      0 0 50px rgba(59, 130, 246, 0.8),
      inset 0 1px 0 rgba(255, 255, 255, 0.5) !important;
  }
}

.save-btn:active {
  transform: translateY(-1px) scale(1.01) !important;
  box-shadow: 
    0 4px 12px rgba(30, 64, 175, 0.4),
    0 2px 6px rgba(37, 99, 235, 0.3),
    0 0 25px rgba(59, 130, 246, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  animation: saveButtonClick 0.3s ease !important;
}

@keyframes saveButtonClick {
  0% {
    transform: translateY(-1px) scale(1.01);
  }
  50% {
    transform: translateY(-1px) scale(0.98);
  }
  100% {
    transform: translateY(-1px) scale(1.01);
  }
}

.save-btn :deep(.v-btn__content) {
  color: white !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  width: 100% !important;
  text-align: center !important;
  gap: 8px !important;
  position: relative;
  z-index: 2;
}

.save-btn :deep(.v-btn__prepend),
.save-btn :deep(.v-btn__append) {
  color: white !important;
  margin: 0 !important;
  position: relative;
  z-index: 2;
}

.save-btn :deep(.v-icon) {
  color: white !important;
  margin: 0 !important;
  transition: transform 0.3s ease !important;
  position: relative;
  z-index: 2;
}

.save-btn:hover :deep(.v-icon) {
  transform: rotate(90deg) scale(1.1) !important;
}

.save-btn:disabled {
  opacity: 0.6 !important;
  cursor: not-allowed !important;
  animation: none !important;
}

.cancel-btn {
  border-radius: 12px !important;
  font-weight: 500 !important;
  padding: 12px 24px !important;
}

/* تحسينات إضافية */
.edit-engineer-dialog .v-card-text {
  background: #ffffff !important;
  padding: 24px !important;
}

.edit-engineer-dialog .v-divider {
  border-color: rgba(59, 130, 246, 0.2) !important;
  margin: 0 !important;
}

/* ألوان سوداء للقائمة المنبثقة */
.edit-engineer-dialog .v-text-field input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-label {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__outline {
  color: #000000 !important;
}

.edit-engineer-dialog .v-icon {
  color: #000000 !important;
}

/* قواعد أقوى لضمان اللون الأسود */
.edit-engineer-dialog input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-input__control input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-text-field__input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-input input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__field input {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input::placeholder {
  color: #666666 !important;
}

.edit-engineer-dialog .v-field__input input:focus {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input:not(:placeholder-shown) {
  color: #000000 !important;
}

/* قواعد إضافية قوية جداً */
.edit-engineer-dialog * {
  color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input,
.edit-engineer-dialog .v-field__input textarea,
.edit-engineer-dialog .v-field__input select {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input:focus,
.edit-engineer-dialog .v-field__input textarea:focus,
.edit-engineer-dialog .v-field__input select:focus {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input:active,
.edit-engineer-dialog .v-field__input textarea:active,
.edit-engineer-dialog .v-field__input select:active {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.edit-engineer-dialog .v-field__input input:hover,
.edit-engineer-dialog .v-field__input textarea:hover,
.edit-engineer-dialog .v-field__input select:hover {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* إجبار اللون الأسود على جميع النصوص */
.edit-engineer-dialog .v-text-field .v-field__input input {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.edit-engineer-dialog .v-text-field .v-field__input input:focus {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.edit-engineer-dialog .v-text-field .v-field__input input:not(:placeholder-shown) {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* CSS لمشاريع المهندس */
.projects-dialog .projects-dialog-title {
  background: linear-gradient(135deg, #4caf50 0%, #388e3c 100%);
  color: #000000 !important;
  padding: 1.5rem !important;
}

.projects-dialog .project-item {
  border-bottom: 1px solid #e0e0e0;
  transition: background-color 0.2s ease;
}

.projects-dialog .project-item:hover {
  background-color: #f5f5f5;
}

.projects-dialog .add-project-btn {
  height: 48px !important;
  font-weight: 600 !important;
}

/* تحسين ألوان القائمة المنسدلة */
.projects-dialog .v-select {
  color: #000000 !important;
}

.projects-dialog .v-select .v-field__input {
  color: #000000 !important;
}

.projects-dialog .v-select .v-label {
  color: #000000 !important;
  font-weight: 600 !important;
}

.projects-dialog .v-select .v-icon {
  color: #000000 !important;
}

.projects-dialog .v-select .v-field__outline {
  border-color: #1976d2 !important;
  border-width: 2px !important;
}

.projects-dialog .v-select .v-field--focused .v-field__outline {
  border-color: #1976d2 !important;
  border-width: 2px !important;
}

/* تحسين ألوان قائمة تفاصيل المهندس */
.engineer-details-dialog .v-card-text {
  color: #1e293b !important;
  background: #ffffff !important;
}

.engineer-details-dialog .v-card-text span {
  color: #1e293b !important;
}

.engineer-details-dialog .v-card-title:first-child {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
}

.engineer-details-dialog .v-card {
  background: #ffffff !important;
}

.engineer-details-dialog .v-card.variant-outlined {
  border: 2px solid #e2e8f0 !important;
  background: #ffffff !important;
}

.engineer-details-dialog .v-card-title.text-h6 {
  color: #1e40af !important;
  font-weight: 700 !important;
}

.engineer-details-dialog .v-card-text .v-card-title {
  color: #1e40af !important;
  font-weight: 700 !important;
}

.engineer-details-dialog .v-card-text .font-weight-medium,
.engineer-details-dialog .v-card-text .font-weight-bold {
  color: #1e293b !important;
}

.engineer-details-dialog .v-card-text .text-primary {
  color: #2563eb !important;
  font-weight: 600 !important;
}

.projects-dialog .v-select .v-field--focused .v-label {
  color: #1976d2 !important;
  font-weight: 700 !important;
}

/* تحسين قائمة الخيارات */
.projects-dialog .v-list {
  background: #ffffff !important;
  border: 1px solid #e0e0e0 !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

.projects-dialog .v-list-item {
  color: #000000 !important;
  font-weight: 500 !important;
  padding: 12px 16px !important;
  border-bottom: 1px solid #f0f0f0 !important;
}

.projects-dialog .v-list-item:hover {
  background-color: #e3f2fd !important;
  color: #1976d2 !important;
}

.projects-dialog .v-list-item--active {
  background-color: #1976d2 !important;
  color: #ffffff !important;
}

.projects-dialog .v-list-item--active:hover {
  background-color: #1565c0 !important;
  color: #ffffff !important;
}

/* تحسين النص المحدد */
.projects-dialog .v-select .v-field__input input {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.projects-dialog .v-select .v-field__input input::placeholder {
  color: #666666 !important;
  font-weight: 500 !important;
}

/* تحسينات إضافية للقائمة المنسدلة */
.projects-dialog .enhanced-select {
  font-size: 1.1rem !important;
  font-weight: 600 !important;
}

.projects-dialog .enhanced-select .v-field__input {
  font-size: 1.1rem !important;
  font-weight: 600 !important;
}

.projects-dialog .enhanced-select .v-label {
  font-size: 1rem !important;
  font-weight: 600 !important;
  color: #1e40af !important;
}

.projects-dialog .enhanced-select .v-field--focused .v-label {
  color: #1e40af !important;
  font-weight: 700 !important;
  transform: translateY(-8px) scale(0.85) !important;
}

.projects-dialog .enhanced-select .v-field__outline {
  border-color: #1e40af !important;
  border-width: 2px !important;
  border-radius: 8px !important;
}

.projects-dialog .enhanced-select .v-field--focused .v-field__outline {
  border-color: #1e40af !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 2px rgba(30, 64, 175, 0.2) !important;
}

.projects-dialog .enhanced-select .v-icon {
  color: #1e40af !important;
  font-size: 1.2rem !important;
}

/* تحسين قائمة الخيارات المنسدلة باللون النيلي */
.projects-dialog .v-menu .v-list {
  background: linear-gradient(135deg, #ffffff 0%, #f1f5f9 50%, #e2e8f0 100%) !important;
  border: 2px solid #1e40af !important;
  border-radius: 12px !important;
  box-shadow: 0 12px 32px rgba(30, 64, 175, 0.25), 0 4px 16px rgba(30, 64, 175, 0.15) !important;
  max-height: 300px !important;
  overflow-y: auto !important;
  backdrop-filter: blur(10px) !important;
}

.projects-dialog .v-menu .v-list-item {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 16px 20px !important;
  border-bottom: 1px solid rgba(0, 0, 0, 0.2) !important;
  transition: all 0.2s ease !important;
}

.projects-dialog .v-menu .v-list-item:hover {
  background-color: rgba(0, 0, 0, 0.1) !important;
  color: #000000 !important;
  font-weight: 700 !important;
  transform: translateX(4px) !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3) !important;
}

.projects-dialog .v-menu .v-list-item--active {
  background-color: rgba(0, 0, 0, 0.2) !important;
  color: #000000 !important;
  font-weight: 700 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4) !important;
}

.projects-dialog .v-menu .v-list-item--active:hover {
  background-color: rgba(0, 0, 0, 0.3) !important;
  color: #000000 !important;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.5) !important;
}

/* قواعد إضافية لضمان اللون الأسود في القائمة المنسدلة */
.projects-dialog .v-menu .v-list-item .v-list-item-title {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list-item .v-list-item-subtitle {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list-item .v-list-item-content {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list-item .v-list-item__content {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list-item .v-list-item__title {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
}

/* قواعد شاملة لجميع عناصر القائمة */
.projects-dialog .v-menu .v-list .v-list-item * {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item span {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item div {
  color: #000000 !important;
}

/* قواعد إضافية قوية جداً للون الأسود */
.projects-dialog .v-menu .v-list .v-list-item {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item::before {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item::after {
  color: #000000 !important;
}

/* قواعد خاصة بـ Vuetify */
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title {
  color: #000000 !important;
  font-weight: 600 !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
}

/* قواعد للعناصر المباشرة */
.projects-dialog .v-menu .v-list .v-list-item > * {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item > div {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item > span {
  color: #000000 !important;
}

/* قواعد للعناصر المتداخلة */
.projects-dialog .v-menu .v-list .v-list-item div * {
  color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item span * {
  color: #000000 !important;
}

/* جميع النصوص داخل القائمة المنبثقة باللون الأسود */
.projects-dialog * {
  color: #000000 !important;
}

.projects-dialog .v-card-title {
  color: #000000 !important;
}

.projects-dialog .v-card-text {
  color: #000000 !important;
}

.projects-dialog .v-card-subtitle {
  color: #000000 !important;
}

.projects-dialog .v-list-item-title {
  color: #000000 !important;
}

.projects-dialog .v-list-item-subtitle {
  color: #000000 !important;
}

.projects-dialog .v-btn {
  color: #000000 !important;
}

.projects-dialog .v-btn .v-btn__content {
  color: #000000 !important;
}

.projects-dialog .v-icon {
  color: #000000 !important;
}

.projects-dialog .v-chip {
  color: #000000 !important;
}

.projects-dialog .v-chip .v-chip__content {
  color: #000000 !important;
}

.projects-dialog .text-h5 {
  color: #000000 !important;
}

.projects-dialog .text-h6 {
  color: #000000 !important;
}

.projects-dialog .text-subtitle-1 {
  color: #000000 !important;
}

.projects-dialog .text-subtitle-2 {
  color: #000000 !important;
}

.projects-dialog .text-body-1 {
  color: #000000 !important;
}

.projects-dialog .text-body-2 {
  color: #000000 !important;
}

.projects-dialog .text-caption {
  color: #000000 !important;
}

.projects-dialog .text-overline {
  color: #000000 !important;
}

/* استثناءات للعناصر التي يجب أن تبقى سوداء */
.projects-dialog .v-menu .v-list-item {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 18px 24px !important;
  border-bottom: 1px solid rgba(0, 0, 0, 0.15) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  position: relative !important;
  overflow: hidden !important;
}

.projects-dialog .v-menu .v-list-item:hover {
  background: linear-gradient(135deg, rgba(0, 0, 0, 0.08) 0%, rgba(0, 0, 0, 0.15) 100%) !important;
  color: #000000 !important;
  font-weight: 800 !important;
  transform: translateX(6px) scale(1.02) !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2), 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.15) !important;
  border-left: 4px solid #000000 !important;
}

.projects-dialog .v-menu .v-list-item--active {
  background: linear-gradient(135deg, rgba(0, 0, 0, 0.12) 0%, rgba(0, 0, 0, 0.2) 100%) !important;
  color: #000000 !important;
  font-weight: 800 !important;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.25), 0 2px 8px rgba(0, 0, 0, 0.15) !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.15) !important;
  border-left: 4px solid #000000 !important;
  transform: translateX(2px) !important;
}

.projects-dialog .v-menu .v-list-item--active:hover {
  background: linear-gradient(135deg, rgba(0, 0, 0, 0.2) 0%, rgba(0, 0, 0, 0.3) 100%) !important;
  color: #000000 !important;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3), 0 4px 16px rgba(0, 0, 0, 0.2) !important;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.2) !important;
  transform: translateX(4px) scale(1.01) !important;
  border-left: 6px solid #000000 !important;
}

/* قواعد إضافية قوية للون الأسود في القائمة المنسدلة */
.projects-dialog .v-menu .v-list .v-list-item .v-list-item-title {
  color: #000000 !important;
  font-weight: 700 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item-subtitle {
  color: #000000 !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item-content {
  color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content {
  color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title {
  color: #000000 !important;
  font-weight: 700 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* قواعد شاملة لجميع عناصر القائمة المنسدلة */
.projects-dialog .v-menu .v-list .v-list-item * {
  color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item span {
  color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item div {
  color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* قواعد قوية جداً للون الأسود */
.projects-dialog .v-menu .v-list .v-list-item {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.projects-dialog .v-menu .v-list .v-list-item * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* قواعد إضافية قوية جداً للون الأسود */
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  font-weight: 700 !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* قواعد خاصة بـ Vuetify v-list-item */
.projects-dialog .v-list-item {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-list-item .v-list-item__content {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-list-item .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  font-weight: 700 !important;
}

/* قواعد شاملة لجميع العناصر */
.projects-dialog .v-menu .v-list .v-list-item,
.projects-dialog .v-menu .v-list .v-list-item *,
.projects-dialog .v-menu .v-list .v-list-item div,
.projects-dialog .v-menu .v-list .v-list-item span,
.projects-dialog .v-menu .v-list .v-list-item p {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* قواعد خاصة للمشاريع المختارة */
.projects-dialog .v-menu .v-list-item--selected {
  background-color: rgba(0, 0, 0, 0.1) !important;
  color: #000000 !important;
  font-weight: 800 !important;
}

.projects-dialog .v-menu .v-list-item--selected * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-menu .v-list-item--selected .v-list-item-title {
  color: #000000 !important;
  font-weight: 800 !important;
}

.projects-dialog .v-menu .v-list-item--selected .v-list-item__title {
  color: #000000 !important;
  font-weight: 800 !important;
}

/* قواعد للمشاريع المحددة في Vuetify */
.projects-dialog .v-menu .v-list-item[aria-selected="true"] {
  background-color: rgba(0, 0, 0, 0.1) !important;
  color: #000000 !important;
  font-weight: 800 !important;
}

.projects-dialog .v-menu .v-list-item[aria-selected="true"] * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* قواعد للمشاريع المحددة بالكليك */
.projects-dialog .v-menu .v-list-item.v-list-item--active {
  background-color: rgba(0, 0, 0, 0.1) !important;
  color: #000000 !important;
  font-weight: 800 !important;
}

.projects-dialog .v-menu .v-list-item.v-list-item--active * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* تأثيرات بصرية إضافية للقائمة */
.projects-dialog .v-menu .v-list-item::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: -100% !important;
  width: 100% !important;
  height: 100% !important;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent) !important;
  transition: left 0.5s ease !important;
}

.projects-dialog .v-menu .v-list-item:hover::before {
  left: 100% !important;
}

/* تحسين شريط التمرير */
.projects-dialog .v-menu .v-list::-webkit-scrollbar {
  width: 8px !important;
}

.projects-dialog .v-menu .v-list::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1) !important;
  border-radius: 4px !important;
}

.projects-dialog .v-menu .v-list::-webkit-scrollbar-thumb {
  background: linear-gradient(135deg, #000000, #333333) !important;
  border-radius: 4px !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
}

.projects-dialog .v-menu .v-list::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(135deg, #111111, #444444) !important;
}

/* قواعد قوية جداً لضمان اللون الأسود */
.projects-dialog .v-menu .v-list .v-list-item {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2) !important;
  font-weight: 700 !important;
}

.projects-dialog .v-menu .v-list .v-list-item * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

/* قواعد خاصة بـ Vuetify v-list-item */
.projects-dialog .v-list-item {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

.projects-dialog .v-list-item * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

/* قواعد شاملة لجميع العناصر */
.projects-dialog .v-menu .v-list .v-list-item,
.projects-dialog .v-menu .v-list .v-list-item *,
.projects-dialog .v-menu .v-list .v-list-item div,
.projects-dialog .v-menu .v-list .v-list-item span,
.projects-dialog .v-menu .v-list .v-list-item p,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2) !important;
  font-weight: 700 !important;
}

/* قواعد خاصة بـ Vuetify v-list-item */
.projects-dialog .v-list-item,
.projects-dialog .v-list-item *,
.projects-dialog .v-list-item div,
.projects-dialog .v-list-item span,
.projects-dialog .v-list-item p,
.projects-dialog .v-list-item .v-list-item__content,
.projects-dialog .v-list-item .v-list-item__title,
.projects-dialog .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2) !important;
  font-weight: 700 !important;
}

/* قواعد قوية جداً لضمان اللون الأسود - إصدار محسن */
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

/* قواعد خاصة بـ Vuetify v-list-item */
.projects-dialog .v-list-item .v-list-item__content .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

.projects-dialog .v-list-item .v-list-item__content {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.projects-dialog .v-list-item .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

/* قواعد شاملة لجميع العناصر - إصدار محسن */
.projects-dialog .v-menu .v-list .v-list-item,
.projects-dialog .v-menu .v-list .v-list-item *,
.projects-dialog .v-menu .v-list .v-list-item div,
.projects-dialog .v-menu .v-list .v-list-item span,
.projects-dialog .v-menu .v-list .v-list-item p,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle,
.projects-dialog .v-list-item,
.projects-dialog .v-list-item *,
.projects-dialog .v-list-item div,
.projects-dialog .v-list-item span,
.projects-dialog .v-list-item p,
.projects-dialog .v-list-item .v-list-item__content,
.projects-dialog .v-list-item .v-list-item__title,
.projects-dialog .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
}

/* قواعد خاصة للقائمة المنسدلة السوداء */
.black-dropdown-menu .v-list-item {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

.black-dropdown-menu .v-list-item * {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
}

.black-dropdown-menu .v-list-item .v-list-item__content {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

.black-dropdown-menu .v-list-item .v-list-item__title {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

.black-dropdown-menu .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* قواعد شاملة للقائمة المنسدلة السوداء */
.black-dropdown-menu .v-list-item,
.black-dropdown-menu .v-list-item *,
.black-dropdown-menu .v-list-item div,
.black-dropdown-menu .v-list-item span,
.black-dropdown-menu .v-list-item p {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  font-weight: 800 !important;
}

/* قواعد قوية جداً لضمان اللون الأسود - نهائي */
.projects-dialog .v-menu .v-list .v-list-item,
.projects-dialog .v-menu .v-list .v-list-item *,
.projects-dialog .v-menu .v-list .v-list-item div,
.projects-dialog .v-menu .v-list .v-list-item span,
.projects-dialog .v-menu .v-list .v-list-item p,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  font-weight: 900 !important;
  font-size: 1.2rem !important;
}

/* قواعد خاصة بـ Vuetify v-list-item - نهائي */
.projects-dialog .v-list-item,
.projects-dialog .v-list-item *,
.projects-dialog .v-list-item div,
.projects-dialog .v-list-item span,
.projects-dialog .v-list-item p,
.projects-dialog .v-list-item .v-list-item__content,
.projects-dialog .v-list-item .v-list-item__title,
.projects-dialog .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  font-weight: 900 !important;
  font-size: 1.2rem !important;
}

/* قواعد شاملة لجميع العناصر - نهائي */
.projects-dialog .v-menu .v-list .v-list-item,
.projects-dialog .v-menu .v-list .v-list-item *,
.projects-dialog .v-menu .v-list .v-list-item div,
.projects-dialog .v-menu .v-list .v-list-item span,
.projects-dialog .v-menu .v-list .v-list-item p,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__content,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__title,
.projects-dialog .v-menu .v-list .v-list-item .v-list-item__subtitle,
.projects-dialog .v-list-item,
.projects-dialog .v-list-item *,
.projects-dialog .v-list-item div,
.projects-dialog .v-list-item span,
.projects-dialog .v-list-item p,
.projects-dialog .v-list-item .v-list-item__content,
.projects-dialog .v-list-item .v-list-item__title,
.projects-dialog .v-list-item .v-list-item__subtitle,
.black-dropdown-menu .v-list-item,
.black-dropdown-menu .v-list-item *,
.black-dropdown-menu .v-list-item div,
.black-dropdown-menu .v-list-item span,
.black-dropdown-menu .v-list-item p,
.black-dropdown-menu .v-list-item .v-list-item__content,
.black-dropdown-menu .v-list-item .v-list-item__title,
.black-dropdown-menu .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  font-weight: 900 !important;
  font-size: 1.2rem !important;
}

/* قواعد إضافية قوية جداً لضمان اللون الأسود */
.v-menu .v-list .v-list-item,
.v-menu .v-list .v-list-item *,
.v-menu .v-list .v-list-item div,
.v-menu .v-list .v-list-item span,
.v-menu .v-list .v-list-item p,
.v-menu .v-list .v-list-item .v-list-item__content,
.v-menu .v-list .v-list-item .v-list-item__title,
.v-menu .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  font-weight: 900 !important;
  font-size: 1.2rem !important;
}

/* قواعد عامة لجميع القوائم المنسدلة */
.v-list .v-list-item,
.v-list .v-list-item *,
.v-list .v-list-item div,
.v-list .v-list-item span,
.v-list .v-list-item p,
.v-list .v-list-item .v-list-item__content,
.v-list .v-list-item .v-list-item__title,
.v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  font-weight: 900 !important;
  font-size: 1.2rem !important;
}

/* قواعد قوية جداً للقائمة المنسدلة السوداء */
.black-dropdown-select .v-field__input,
.black-dropdown-select .v-field__outline,
.black-dropdown-select .v-label,
.black-dropdown-select .v-field__append-inner,
.black-dropdown-select .v-field__prepend-inner {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* قواعد نهائية للقائمة المنسدلة */
.black-dropdown-menu,
.black-dropdown-menu *,
.black-dropdown-menu .v-list,
.black-dropdown-menu .v-list *,
.black-dropdown-menu .v-list-item,
.black-dropdown-menu .v-list-item *,
.black-dropdown-menu .v-list-item div,
.black-dropdown-menu .v-list-item span,
.black-dropdown-menu .v-list-item p,
.black-dropdown-menu .v-list-item .v-list-item__content,
.black-dropdown-menu .v-list-item .v-list-item__title,
.black-dropdown-menu .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
}

/* قواعد خاصة بـ Vuetify v-select */
.v-select .v-field__input,
.v-select .v-field__outline,
.v-select .v-label,
.v-select .v-field__append-inner,
.v-select .v-field__prepend-inner {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* قواعد شاملة لجميع عناصر القائمة المنسدلة */
.v-menu .v-overlay__content .v-list,
.v-menu .v-overlay__content .v-list *,
.v-menu .v-overlay__content .v-list-item,
.v-menu .v-overlay__content .v-list-item *,
.v-menu .v-overlay__content .v-list-item div,
.v-menu .v-overlay__content .v-list-item span,
.v-menu .v-overlay__content .v-list-item p,
.v-menu .v-overlay__content .v-list-item .v-list-item__content,
.v-menu .v-overlay__content .v-list-item .v-list-item__title,
.v-menu .v-overlay__content .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
}

/* قواعد نهائية قوية جداً للقائمة المنسدلة */
.black-dropdown-select .v-list-item,
.black-dropdown-select .v-list-item *,
.black-dropdown-select .v-list-item div,
.black-dropdown-select .v-list-item span,
.black-dropdown-select .v-list-item p,
.black-dropdown-select .v-list-item .v-list-item__content,
.black-dropdown-select .v-list-item .v-list-item__title,
.black-dropdown-select .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.6) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
}

/* قواعد شاملة لجميع عناصر v-list-item */
.v-list-item,
.v-list-item *,
.v-list-item div,
.v-list-item span,
.v-list-item p,
.v-list-item .v-list-item__content,
.v-list-item .v-list-item__title,
.v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.6) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
}

/* قواعد نهائية لجميع عناصر القائمة المنسدلة */
.v-overlay__content .v-list,
.v-overlay__content .v-list *,
.v-overlay__content .v-list-item,
.v-overlay__content .v-list-item *,
.v-overlay__content .v-list-item div,
.v-overlay__content .v-list-item span,
.v-overlay__content .v-list-item p,
.v-overlay__content .v-list-item .v-list-item__content,
.v-overlay__content .v-list-item .v-list-item__title,
.v-overlay__content .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.6) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
}

/* قواعد خاصة بـ template #item */
.black-dropdown-select .v-list-item span,
.black-dropdown-select .v-list-item .v-list-item__title,
.black-dropdown-select .v-list-item .v-list-item__content {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.6) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد شاملة لجميع عناصر القائمة */
.v-list-item span,
.v-list-item .v-list-item__title,
.v-list-item .v-list-item__content,
.v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.6) !important;
  font-weight: 900 !important;
  font-size: 1.3rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد نهائية قوية جداً للقائمة المنسدلة */
.v-menu .v-overlay__content .v-list .v-list-item,
.v-menu .v-overlay__content .v-list .v-list-item *,
.v-menu .v-overlay__content .v-list .v-list-item div,
.v-menu .v-overlay__content .v-list .v-list-item span,
.v-menu .v-overlay__content .v-list .v-list-item p,
.v-menu .v-overlay__content .v-list .v-list-item .v-list-item__content,
.v-menu .v-overlay__content .v-list .v-list-item .v-list-item__title,
.v-menu .v-overlay__content .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد خاصة بـ Vuetify v-select dropdown */
.v-select .v-field__input,
.v-select .v-field__outline,
.v-select .v-label,
.v-select .v-field__append-inner,
.v-select .v-field__prepend-inner {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
}

/* قواعد شاملة لجميع عناصر القائمة المنسدلة */
.v-list .v-list-item,
.v-list .v-list-item *,
.v-list .v-list-item div,
.v-list .v-list-item span,
.v-list .v-list-item p,
.v-list .v-list-item .v-list-item__content,
.v-list .v-list-item .v-list-item__title,
.v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* إصلاح مشكلة z-index للقائمة المنسدلة */
.v-menu .v-overlay__content,
.v-menu .v-overlay__content .v-list,
.v-menu .v-overlay__content .v-list-item {
  z-index: 9999 !important;
  position: relative !important;
}

/* قواعد خاصة للقائمة المنسدلة في dialog */
.projects-dialog .v-menu .v-overlay__content,
.projects-dialog .v-menu .v-overlay__content .v-list,
.projects-dialog .v-menu .v-overlay__content .v-list-item {
  z-index: 10000 !important;
  position: relative !important;
}

/* قواعد شاملة لجميع القوائم المنسدلة */
.v-menu,
.v-menu .v-overlay,
.v-menu .v-overlay__content {
  z-index: 9999 !important;
}

/* قواعد خاصة بـ v-select */
.v-select .v-menu,
.v-select .v-menu .v-overlay,
.v-select .v-menu .v-overlay__content {
  z-index: 10000 !important;
}

/* قواعد قوية جداً لضمان ظهور القائمة في المقدمة */
.projects-dialog .v-select .v-menu,
.projects-dialog .v-select .v-menu .v-overlay,
.projects-dialog .v-select .v-menu .v-overlay__content,
.projects-dialog .v-select .v-menu .v-overlay__content .v-list,
.projects-dialog .v-select .v-menu .v-overlay__content .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد شاملة لجميع القوائم المنسدلة */
.v-menu .v-overlay__content,
.v-menu .v-overlay__content .v-list,
.v-menu .v-overlay__content .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد نهائية للقائمة المنسدلة */
.black-dropdown-menu,
.black-dropdown-menu .v-list,
.black-dropdown-menu .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد شاملة لجميع القوائم المنسدلة في body */
body .v-menu .v-overlay__content,
body .v-menu .v-overlay__content .v-list,
body .v-menu .v-overlay__content .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد خاصة للقائمة المنسدلة المرفقة بـ body */
body .black-dropdown-menu,
body .black-dropdown-menu .v-list,
body .black-dropdown-menu .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد نهائية لجميع القوائم المنسدلة */
.v-menu .v-overlay__content,
.v-menu .v-overlay__content .v-list,
.v-menu .v-overlay__content .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد خاصة بـ v-combobox */
.v-combobox .v-menu,
.v-combobox .v-menu .v-overlay,
.v-combobox .v-menu .v-overlay__content,
.v-combobox .v-menu .v-overlay__content .v-list,
.v-combobox .v-menu .v-overlay__content .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد خاصة للقائمة المنسدلة في v-combobox */
.v-combobox .v-list .v-list-item,
.v-combobox .v-list .v-list-item *,
.v-combobox .v-list .v-list-item div,
.v-combobox .v-list .v-list-item span,
.v-combobox .v-list .v-list-item p,
.v-combobox .v-list .v-list-item .v-list-item__content,
.v-combobox .v-list .v-list-item .v-list-item__title,
.v-combobox .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد قوية جداً لضمان اللون الأسود في v-combobox */
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item *,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item div,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item span,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item p,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item .v-list-item__content,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item .v-list-item__title,
.v-combobox .v-menu .v-overlay__content .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد شاملة لجميع عناصر v-combobox */
.v-combobox .v-list-item,
.v-combobox .v-list-item *,
.v-combobox .v-list-item div,
.v-combobox .v-list-item span,
.v-combobox .v-list-item p,
.v-combobox .v-list-item .v-list-item__content,
.v-combobox .v-list-item .v-list-item__title,
.v-combobox .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد خاصة بـ v-autocomplete */
.v-autocomplete .v-menu,
.v-autocomplete .v-menu .v-overlay,
.v-autocomplete .v-menu .v-overlay__content,
.v-autocomplete .v-menu .v-overlay__content .v-list,
.v-autocomplete .v-menu .v-overlay__content .v-list-item {
  z-index: 50000 !important;
  position: fixed !important;
}

/* قواعد خاصة للقائمة المنسدلة في v-autocomplete */
.v-autocomplete .v-list .v-list-item,
.v-autocomplete .v-list .v-list-item *,
.v-autocomplete .v-list .v-list-item div,
.v-autocomplete .v-list .v-list-item span,
.v-autocomplete .v-list .v-list-item p,
.v-autocomplete .v-list .v-list-item .v-list-item__content,
.v-autocomplete .v-list .v-list-item .v-list-item__title,
.v-autocomplete .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد قوية جداً لضمان اللون الأسود في v-autocomplete */
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item *,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item div,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item span,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item p,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item .v-list-item__content,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item .v-list-item__title,
.v-autocomplete .v-menu .v-overlay__content .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* خلفية سوداء للقائمة المنسدلة */
.v-autocomplete .v-menu .v-overlay__content .v-list {
  background: #000000 !important;
  border: 2px solid #333333 !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.8) !important;
}

/* قواعد نهائية قوية جداً للون الأسود */
.v-autocomplete .v-list-item,
.v-autocomplete .v-list-item *,
.v-autocomplete .v-list-item div,
.v-autocomplete .v-list-item span,
.v-autocomplete .v-list-item p,
.v-autocomplete .v-list-item .v-list-item__content,
.v-autocomplete .v-list-item .v-list-item__title,
.v-autocomplete .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* قواعد شاملة لجميع عناصر القائمة المنسدلة */
.v-menu .v-overlay__content .v-list .v-list-item,
.v-menu .v-overlay__content .v-list .v-list-item *,
.v-menu .v-overlay__content .v-list .v-list-item div,
.v-menu .v-overlay__content .v-list .v-list-item span,
.v-menu .v-overlay__content .v-list .v-list-item p,
.v-menu .v-overlay__content .v-list .v-list-item .v-list-item__content,
.v-menu .v-overlay__content .v-list .v-list-item .v-list-item__title,
.v-menu .v-overlay__content .v-list .v-list-item .v-list-item__subtitle {
  color: #000000 !important;
  -webkit-text-fill-color: #000000 !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* خلفية سوداء لجميع القوائم المنسدلة */
.v-menu .v-overlay__content .v-list {
  background: #000000 !important;
  border: 2px solid #333333 !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.8) !important;
}

/* قواعد نهائية لجميع عناصر القائمة */
.v-list-item,
.v-list-item *,
.v-list-item div,
.v-list-item span,
.v-list-item p,
.v-list-item .v-list-item__content,
.v-list-item .v-list-item__title,
.v-list-item .v-list-item__subtitle {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.8) !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  opacity: 1 !important;
  display: block !important;
  visibility: visible !important;
}

/* خلفية سوداء لجميع عناصر القائمة */
.v-list-item {
  background: #000000 !important;
  border-bottom: 1px solid #333333 !important;
}

/* تأثيرات حركية لعناصر القائمة المنسدلة */
.v-list-item {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  transform: translateX(0) !important;
  position: relative !important;
  overflow: hidden !important;
}

.v-list-item:hover {
  transform: translateX(8px) scale(1.02) !important;
  background: linear-gradient(135deg, rgba(30, 64, 175, 0.1), rgba(99, 102, 241, 0.1)) !important;
  box-shadow: 0 8px 25px rgba(30, 64, 175, 0.3) !important;
  border-left: 4px solid #1e40af !important;
}

.v-list-item:active {
  transform: translateX(4px) scale(0.98) !important;
  background: linear-gradient(135deg, rgba(30, 64, 175, 0.2), rgba(99, 102, 241, 0.2)) !important;
}

/* تأثير shimmer للعناصر */
.v-list-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  transition: left 0.5s;
  z-index: 1;
}

.v-list-item:hover::before {
  left: 100%;
}

/* تأثير pulse للعناصر */
.v-list-item::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  background: radial-gradient(circle, rgba(30, 64, 175, 0.3), transparent);
  border-radius: 50%;
  transform: translate(-50%, -50%);
  transition: all 0.3s ease;
  z-index: 0;
}

.v-list-item:hover::after {
  width: 200px;
  height: 200px;
}

/* تأثير bounce للعناصر عند الظهور */
@keyframes slideInBounce {
  0% {
    opacity: 0;
    transform: translateY(-20px) scale(0.8);
  }
  50% {
    opacity: 0.8;
    transform: translateY(5px) scale(1.05);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.v-list-item {
  animation: slideInBounce 0.4s cubic-bezier(0.68, -0.55, 0.265, 1.55) !important;
}

/* تأثير wave للعناصر */
@keyframes wave {
  0%, 100% {
    transform: translateX(0);
  }
  25% {
    transform: translateX(2px);
  }
  75% {
    transform: translateX(-2px);
  }
}

.v-list-item:hover {
  animation: wave 0.6s ease-in-out infinite !important;
}

/* تأثير glow للنص */
.v-list-item .v-list-item__title,
.v-list-item .v-list-item__subtitle {
  transition: all 0.3s ease !important;
  position: relative !important;
}

.v-list-item:hover .v-list-item__title,
.v-list-item:hover .v-list-item__subtitle {
  text-shadow: 0 0 10px rgba(30, 64, 175, 0.8), 0 4px 8px rgba(0, 0, 0, 0.7) !important;
  transform: scale(1.05) !important;
}

/* تأثير ripple للعناصر */
.v-list-item {
  position: relative !important;
  overflow: hidden !important;
}

.v-list-item .ripple {
  position: absolute;
  border-radius: 50%;
  background: rgba(30, 64, 175, 0.3);
  transform: scale(0);
  animation: ripple 0.6s linear;
  pointer-events: none;
}

@keyframes ripple {
  to {
    transform: scale(4);
    opacity: 0;
  }
}

/* Pagination Styles */
.pagination-container {
  background: #f5f5f5;
  border-radius: 8px;
  margin: 0 16px 16px 16px;
}

.pagination-container :deep(.v-pagination__item) {
  color: #333 !important;
}

.pagination-container :deep(.v-pagination__item--is-active) {
  background: #1976d2 !important;
  color: white !important;
}
</style>