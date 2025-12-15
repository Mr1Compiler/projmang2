
<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">🏗️</span>
        <h1 class="page-title text-glow fade-in">
          <span class="title-icon">🏗️</span>
          <span class="title-text">إدارة المشاريع الهندسية</span>
          <span class="title-decoration"></span>
        </h1>
        <p class="page-subtitle fade-in">نظام متكامل لإدارة وتتبع جميع المشاريع والمهام الهندسية</p>
        <div class="page-main-title">
          <h2 class="main-title">الصفحة الرئيسية للمشاريع</h2>
          <p class="main-subtitle">مرحباً بك في نظام إدارة المشاريع الهندسية المتقدم</p>
        </div>
        
      <!-- Simple Title Section -->
      <div class="simple-title-section">
        <div class="title-container">
          <v-icon size="48" color="primary" class="title-icon">mdi-folder-multiple-outline</v-icon>
          <h2 class="simple-title">قائمة المشاريع المتاحة</h2>
          <p class="simple-subtitle">استعرض وأدر جميع مشاريعك الهندسية من مكان واحد</p>
          <div class="title-stats-simple">
            <span class="stat-item">{{ totalProjects || 3 }} مشروع</span>
            <span class="stat-separator">•</span>
            <span class="stat-item">{{ activeProjects || 1 }} نشط</span>
            <span class="stat-separator">•</span>
            <span class="stat-item">{{ pendingProjects || 1 }} في الانتظار</span>
            <span class="stat-separator">•</span>
            <span class="stat-item">{{ averageProgress || 0 }}% متوسط التقدم</span>
          </div>
          <div class="title-actions">
            <v-btn
              :color="showTeamManagement ? 'success' : 'primary'"
              :variant="showTeamManagement ? 'elevated' : 'outlined'"
              size="large"
              prepend-icon="mdi-account-group"
              @click="showTeamManagement = !showTeamManagement"
              class="team-toggle-btn"
            >
              {{ showTeamManagement ? 'إخفاء إدارة الفريق' : 'إدارة الفريق' }}
            </v-btn>
          </div>
        </div>
      </div>


      <!-- Team Management Section -->
      <v-expand-transition>
        <div v-if="showTeamManagement" class="team-management-section">
          <v-card class="team-management-card">
            <v-card-title class="team-section-header">
              <v-icon size="32" color="success" class="mr-2">mdi-account-group</v-icon>
              <span class="text-h5 font-weight-bold">إدارة الفريق</span>
              <v-spacer />
              <v-btn
                color="success"
                variant="elevated"
                size="small"
                prepend-icon="mdi-plus"
                @click="showAddMemberDialog = true"
              >
                إضافة عضو
              </v-btn>
            </v-card-title>
            
            <v-card-text>
              <!-- Team Statistics -->
              <div class="team-stats-row mb-4">
                <v-row>
                  <v-col cols="12" sm="6" md="3">
                    <v-card class="team-stat-card">
                      <v-card-text class="text-center">
                        <v-icon size="32" color="primary" class="mb-2">mdi-account-multiple</v-icon>
                        <h3 class="text-h6 font-weight-bold">{{ teamMembers.length }}</h3>
                        <p class="text-caption">إجمالي الأعضاء</p>
                      </v-card-text>
                    </v-card>
                  </v-col>
                  <v-col cols="12" sm="6" md="3">
                    <v-card class="team-stat-card">
                      <v-card-text class="text-center">
                        <v-icon size="32" color="success" class="mb-2">mdi-account-check</v-icon>
                        <h3 class="text-h6 font-weight-bold">{{ activeTeamMembers }}</h3>
                        <p class="text-caption">أعضاء نشطين</p>
                      </v-card-text>
                    </v-card>
                  </v-col>
                  <v-col cols="12" sm="6" md="3">
                    <v-card class="team-stat-card">
                      <v-card-text class="text-center">
                        <v-icon size="32" color="info" class="mb-2">mdi-office-building</v-icon>
                        <h3 class="text-h6 font-weight-bold">{{ teamDepartments }}</h3>
                        <p class="text-caption">الأقسام</p>
                      </v-card-text>
                    </v-card>
                  </v-col>
                  <v-col cols="12" sm="6" md="3">
                    <v-card class="team-stat-card">
                      <v-card-text class="text-center">
                        <v-icon size="32" color="warning" class="mb-2">mdi-clipboard-list</v-icon>
                        <h3 class="text-h6 font-weight-bold">{{ totalTeamTasks }}</h3>
                        <p class="text-caption">المهام</p>
                      </v-card-text>
                    </v-card>
                  </v-col>
                </v-row>
              </div>

              <!-- Team Members Table -->
              <v-data-table
                :headers="teamHeaders"
                :items="teamMembers"
                :items-per-page="5"
                class="team-members-table"
                :loading="teamLoading"
              >
                <!-- Avatar Column -->
                <template #item.avatar="{ item }">
                  <v-avatar size="32" class="member-avatar">
                    <v-img v-if="item.avatar" :src="item.avatar" />
                    <v-icon v-else size="20" color="primary">mdi-account</v-icon>
                  </v-avatar>
                </template>

                <!-- Name Column -->
                <template #item.name="{ item }">
                  <div class="member-info">
                    <h4 class="member-name">{{ item.name }}</h4>
                    <p class="member-email">{{ item.email }}</p>
                  </div>
                </template>

                <!-- Department Column -->
                <template #item.department="{ item }">
                  <v-chip
                    :color="getDepartmentColor(item.department)"
                    variant="elevated"
                    size="small"
                  >
                    {{ item.department }}
                  </v-chip>
                </template>

                <!-- Role Column -->
                <template #item.role="{ item }">
                  <v-chip
                    :color="getRoleColor(item.role)"
                    variant="elevated"
                    size="small"
                  >
                    {{ item.role }}
                  </v-chip>
                </template>

                <!-- Status Column -->
                <template #item.status="{ item }">
                  <v-chip
                    :color="item.status === 'نشط' ? 'success' : 'error'"
                    variant="elevated"
                    size="small"
                  >
                    <v-icon start>{{ item.status === 'نشط' ? 'mdi-check-circle' : 'mdi-close-circle' }}</v-icon>
                    {{ item.status }}
                  </v-chip>
                </template>

                <!-- Tasks Column -->
                <template #item.tasks="{ item }">
                  <div class="tasks-info">
                    <span class="task-count">{{ item.tasksCount }}</span>
                    <span class="task-label">مهمة</span>
                  </div>
                </template>

                <!-- Actions Column -->
                <template #item.actions="{ item }">
                  <div class="action-buttons">
                    <v-btn
                      icon="mdi-eye"
                      size="x-small"
                      color="info"
                      variant="elevated"
                      @click="viewTeamMember(item)"
                    />
                    <v-btn
                      icon="mdi-pencil"
                      size="x-small"
                      color="warning"
                      variant="elevated"
                      @click="editTeamMember(item)"
                    />
                    <v-btn
                      icon="mdi-delete"
                      size="x-small"
                      color="error"
                      variant="elevated"
                      @click="deleteTeamMember(item)"
                    />
                  </div>
                </template>
              </v-data-table>
            </v-card-text>
          </v-card>
        </div>
      </v-expand-transition>

      <!-- Add Team Member Dialog -->
      <v-dialog v-model="showAddMemberDialog" max-width="600">
        <v-card class="team-member-dialog-card">
          <v-card-title class="team-member-dialog-header">
            <v-icon size="28" color="white" class="mr-2">mdi-account-plus</v-icon>
            <span class="text-h5 font-weight-black">إضافة عضو جديد</span>
          </v-card-title>
          
          <v-card-text>
            <v-form ref="teamMemberForm" v-model="teamMemberFormValid">
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="teamMemberForm.name"
                    label="اسم العضو"
                    variant="outlined"
                    :rules="[v => !!v || 'اسم العضو مطلوب']"
                    required
                    class="team-member-form-field"
                    style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="teamMemberForm.email"
                    label="البريد الإلكتروني"
                    variant="outlined"
                    type="email"
                    :rules="[v => !!v || 'البريد الإلكتروني مطلوب', v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح']"
                    required
                    class="team-member-form-field"
                    style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="teamMemberForm.phone"
                    label="رقم الهاتف"
                    variant="outlined"
                    :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                    required
                    class="team-member-form-field"
                    style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="teamMemberForm.department"
                    :items="departmentOptions"
                    label="القسم"
                    variant="outlined"
                    required
                    class="team-member-form-field"
                    style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="teamMemberForm.role"
                    :items="roleOptions"
                    label="المنصب"
                    variant="outlined"
                    required
                    class="team-member-form-field"
                    style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="teamMemberForm.status"
                    :items="statusOptions"
                    label="الحالة"
                    variant="outlined"
                    required
                    class="team-member-form-field"
                    style="--v-theme-primary: #000000; --v-field-label-color: #000000; --v-field-label-active-color: #000000; --v-field-label-floating-color: #000000; text-align: center; direction: rtl;"
                  />
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
          
          <v-card-actions class="pa-4">
            <v-spacer></v-spacer>
            <v-btn
              color="grey"
              variant="outlined"
              @click="closeAddMemberDialog"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="success"
              variant="elevated"
              @click="saveTeamMember"
              :disabled="!teamMemberFormValid"
            >
              إضافة العضو
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
        
        <div class="page-badge">
          <v-chip color="primary" variant="elevated" size="large">
            <v-icon start>mdi-star</v-icon>
            نظام إدارة متقدم
          </v-chip>
        </div>
      </div>

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row" no-gutters>
        <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
          <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="48" color="info">mdi-folder-multiple</v-icon>
        </div>
            <h3 class="text-h4 font-weight-bold text-info mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ totalProjects || 3 }}</h3>
            <p class="text-caption text-info mb-0">إجمالي المشاريع</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
          <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="48" color="success">mdi-check-circle</v-icon>
      </div>
            <h3 class="text-h4 font-weight-bold text-success mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ activeProjects || 1 }}</h3>
            <p class="text-caption text-success mb-0">مشاريع نشطة</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="2" lg="2" xl="2" class="pa-2">
          <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="48" color="warning">mdi-clock-alert</v-icon>
            </div>
            <h3 class="text-h4 font-weight-bold text-warning mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ pendingProjects || 1 }}</h3>
            <p class="text-caption text-warning mb-0">في الانتظار</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
          <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="48" color="error">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h6 font-weight-bold text-error mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; white-space: nowrap;">{{ formatCurrency(totalBudget) || '225,000 د.ع' }}</h3>
            <p class="text-caption text-error mb-0">إجمالي الميزانية</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3" lg="3" xl="3" class="pa-2">
          <v-card class="stat-card pa-4 pb-6 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-3 icon-glow">
              <v-icon size="48" color="primary">mdi-chart-line</v-icon>
            </div>
            <h3 class="text-h4 font-weight-bold text-primary mb-2" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ averageProgress || 0 }}%</h3>
            <p class="text-caption text-primary mb-0">متوسط التقدم</p>
          </v-card>
        </v-col>
      </v-row>

      <!-- Projects Table -->
      <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
        <v-card-title class="d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" style="color: #4338ca;" size="28">mdi-folder-multiple</v-icon>
            <span class="text-h4 font-weight-black" style="color: #ffffff; font-family: 'Arial', 'Helvetica', sans-serif; text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2); letter-spacing: 0.5px;">قائمة المشاريع</span>
            <v-chip class="ms-3" color="primary" size="small" variant="elevated">{{ projects.length || 3 }}</v-chip>
          </div>
          <v-btn
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddProjectDialog"
            elevation="2"
          >
            <span>إضافة مشروع جديد <v-icon class="icon-glow inline-icon">mdi-plus</v-icon></span>
          </v-btn>
        </v-card-title>

        <!-- Search Bar -->
        <v-card-text class="pb-0 modern-filter-section" style="display: none;">
          <div class="filter-header">
            <div class="filter-icon-container">
              <v-icon class="filter-icon" color="primary" size="24">mdi-tune-variant</v-icon>
              <div class="icon-glow"></div>
            </div>
            <h3 class="filter-title">أدوات البحث والفلترة الذكية</h3>
            <p class="filter-subtitle">استخدم الذكاء الاصطناعي للبحث والفلترة المتقدمة</p>
            <div class="filter-decoration">
              <div class="decoration-line"></div>
              <div class="decoration-dot"></div>
            </div>
          </div>
          <v-row class="align-center modern-fields-row">
            <v-col cols="12" md="6">
              <div class="modern-search-container">
        <v-text-field
          v-model="searchQuery"
                  label="البحث الذكي في المشاريع..."
          prepend-inner-icon="mdi-magnify"
          variant="outlined"
                density="comfortable"
                clearable
          hide-details
                  class="modern-search-field"
                  style="--v-theme-primary: #dc2626; --v-field-label-color: #dc2626; --v-field-label-active-color: #dc2626; --v-field-label-floating-color: #dc2626; text-align: center; direction: rtl;"
              />
                <div class="search-glow"></div>
              </div>
            </v-col>
            <v-col cols="12" md="6">
              <div class="modern-filters-container">
                <div class="filter-item">
                <v-select
                  v-model="selectedCategory"
                  :items="filterCategories"
                  item-title="title"
                  item-value="value"
                  label="فلترة حسب الفئة"
                  variant="outlined"
                  density="comfortable"
                  clearable
                  hide-details
                    class="modern-filter-field"
                    style="--v-theme-primary: #dc2626; --v-field-label-color: #dc2626; --v-field-label-active-color: #dc2626; --v-field-label-floating-color: #dc2626; text-align: center; direction: rtl;"
                />
                  <div class="filter-glow"></div>
                </div>
                <div class="filter-item">
                <v-select
                  v-model="selectedStatus"
                  :items="filterStatuses"
                  item-title="title"
                  item-value="value"
                  label="فلترة حسب الحالة"
                  variant="outlined"
                  density="comfortable"
                  clearable
                  hide-details
                    class="modern-filter-field"
                    style="--v-theme-primary: #dc2626; --v-field-label-color: #dc2626; --v-field-label-active-color: #dc2626; --v-field-label-floating-color: #dc2626; text-align: center; direction: rtl;"
                />
                  <div class="filter-glow"></div>
                </div>
              </div>
            </v-col>
          </v-row>
        </v-card-text>

        <!-- Projects Grid -->
        <v-row v-if="(filteredProjects.length > 0) || true" class="projects-grid-row">
          <v-col
            v-for="project in (filteredProjects.length > 0 ? filteredProjects : sampleProjects)"
            :key="project.id"
            cols="12"
            sm="6"
            md="6"
            lg="4"
            xl="4"
          >
            <v-card
              class="project-card"
              elevation="4"
              hover
              min-height="400"
            >
              <!-- Project Header -->
              <v-card-title class="project-card-title">
                <div class="d-flex align-center">
                  <v-avatar size="40" color="primary" variant="tonal">
                    <v-icon>mdi-folder-multiple</v-icon>
                  </v-avatar>
                  <div class="ms-3 flex-grow-1">
                    <h3 class="project-name">{{ project.name }}</h3>
            <v-chip
              size="small"
                      :color="getStatusColor(project.status)"
                      variant="tonal"
                      class="status-chip-main"
            >
                      <v-icon :icon="getStatusIcon(project.status)" size="small" class="me-1" />
                      {{ getStatusText(project.status) }}
            </v-chip>
                  </div>
                </div>
              </v-card-title>

              <v-card-text class="project-card-content">
                <!-- Project Description -->
                <div class="project-description-wrapper">
                  <p class="project-description">{{ project.description || 'لا يوجد وصف للمشروع' }}</p>
                </div>

                <!-- Project Details -->
                <div class="project-details">
                  <div class="detail-item">
                    <v-icon size="small" color="primary">mdi-map-marker</v-icon>
                    <span class="detail-text">{{ project.location }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="success">mdi-currency-usd</v-icon>
                    <span class="detail-text">{{ formatCurrency(project.initialCost) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="warning">mdi-alert</v-icon>
                    <span class="detail-text">{{ formatCurrency(project.criticalCost || project.initialCost * 1.5) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="info">mdi-calendar</v-icon>
                    <span class="detail-text">{{ formatDate(project.startDate) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="secondary">mdi-account</v-icon>
                    <span class="detail-text">{{ project.user || project.manager || 'غير محدد' }}</span>
                  </div>
                </div>

                <!-- Progress Bar -->
                <div class="progress-section" style="margin-top: 1rem;">
                  <div class="progress-header" style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 0.5rem;">
                    <span class="progress-label" style="font-family: 'Arial', 'Helvetica', sans-serif; color: #424242; font-size: 0.9rem; font-weight: 600;">نسبة الإنجاز</span>
                    <span class="progress-percentage" style="font-family: 'Arial', 'Helvetica', sans-serif; color: #1976d2; font-size: 0.9rem; font-weight: 700;">{{ project.progress || 0 }}%</span>
                  </div>
                  <v-progress-linear
                    :model-value="project.progress || 0"
                    color="primary"
                    height="8"
                    rounded
                    class="progress-bar"
                    style="border-radius: 4px;"
                  />
                </div>
              </v-card-text>

              <v-card-actions class="project-card-actions">
                <v-spacer />
        <v-btn
          color="primary"
                  variant="elevated"
                  size="default"
                  @click="viewProjectDetails(project)"
                  class="details-btn"
                  block
                >
                  <v-icon class="me-2">mdi-eye</v-icon>
                  عرض التفاصيل
        </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>

        <!-- No Projects Message -->
        <v-card v-else class="no-projects-card" elevation="2">
          <v-card-text class="text-center py-8">
            <v-icon size="4rem" color="grey-lighten-1">mdi-folder-open-outline</v-icon>
            <h3 class="text-h5 text-grey-lighten-1 mt-4">لا يوجد مشاريع</h3>
            <p class="text-body-1 text-grey-lighten-1">لم يتم العثور على أي مشاريع تطابق معايير البحث</p>
          </v-card-text>
        </v-card>

      </v-card>

      <!-- Administrative Expenses Table - Hidden -->
      <v-card v-show="false" class="data-table-card card-glow smooth-transition centered-table mt-6" elevation="2">
        <v-card-title class="d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" style="color: #4338ca;" size="28">mdi-currency-usd</v-icon>
            <span class="text-h4 font-weight-black" style="color: #000000; font-family: 'Arial', 'Helvetica', sans-serif; text-shadow: 0 3px 6px rgba(0, 0, 0, 0.2), 0 1px 3px rgba(0, 0, 0, 0.1); letter-spacing: 0.5px;">المصاريف الإدارية</span>
            <v-chip class="ms-3" color="error" size="small" variant="elevated">{{ administrativeExpenses.length || 5 }}</v-chip>
          </div>
          <v-btn
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddExpenseDialog"
            elevation="2"
            color="error"
          >
            <v-icon class="me-2 icon-glow">mdi-plus</v-icon>
            إضافة مصروف إداري
          </v-btn>
        </v-card-title>

        <!-- Search Bar for Expenses -->
        <v-card-text class="pb-0 modern-expense-filter-section">
          <div class="filter-header">
            <div class="filter-icon-container">
              <v-icon class="filter-icon" color="error" size="24">mdi-chart-line-variant</v-icon>
              <div class="icon-glow expense-glow"></div>
            </div>
            <h3 class="filter-title expense-title">أدوات تحليل المصاريف المتقدمة</h3>
            <p class="filter-subtitle">تحليل ذكي للمصاريف مع تقارير مفصلة</p>
            <div class="filter-decoration">
              <div class="decoration-line expense-line"></div>
              <div class="decoration-dot expense-dot"></div>
            </div>
          </div>
          <v-row class="align-center modern-fields-row">
            <v-col cols="12" md="6">
              <div class="modern-search-container">
              <v-text-field
                v-model="expenseSearchQuery"
                  label="البحث الذكي في المصاريف..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                  class="modern-search-field"
              />
                <div class="search-glow expense-search-glow"></div>
              </div>
            </v-col>
            <v-col cols="12" md="6">
              <div class="modern-filters-container">
                <div class="filter-item">
                <v-select
                  v-model="selectedProjectFilter"
                  :items="projectFilterOptions"
                  item-title="title"
                  item-value="value"
                  label="فلترة حسب المشروع"
                  variant="outlined"
                  density="comfortable"
                  clearable
                  hide-details
                    class="modern-filter-field"
                />
                  <div class="filter-glow expense-filter-glow"></div>
                </div>
                <div class="filter-item">
                <v-select
                  v-model="selectedCostRange"
                  :items="costRangeOptions"
                  item-title="title"
                  item-value="value"
                  label="فلترة حسب التكلفة"
                  variant="outlined"
                  density="comfortable"
                  clearable
                  hide-details
                    class="modern-filter-field"
                />
                  <div class="filter-glow expense-filter-glow"></div>
                </div>
              </div>
            </v-col>
          </v-row>
        </v-card-text>

        <!-- Expenses Table -->
        <v-data-table
          :headers="expenseHeaders"
          :items="filteredExpenses"
          :search="expenseSearchQuery"
          class="expense-table"
          :items-per-page="10"
          :loading="false"
          hover
        >
          <!-- Project Name Column -->
          <template #item.projectName="{ item }">
            <div class="d-flex align-center">
              <v-avatar size="32" color="primary" variant="tonal" class="me-2">
                <v-icon size="16">mdi-folder-multiple</v-icon>
              </v-avatar>
              <span class="font-weight-medium">{{ item.projectName }}</span>
            </div>
          </template>

          <!-- Start Date Column -->
          <template #item.startDate="{ item }">
            <v-chip size="small" color="primary" variant="elevated" class="date-chip">
              {{ formatDate(item.startDate) }}
            </v-chip>
          </template>

          <!-- End Date Column -->
          <template #item.endDate="{ item }">
            <v-chip size="small" color="secondary" variant="elevated" class="date-chip">
              {{ formatDate(item.endDate) }}
            </v-chip>
          </template>

          <!-- Cost Column -->
          <template #item.cost="{ item }">
            <span class="font-weight-bold cost-display" style="font-family: 'Arial', 'Helvetica', sans-serif;">
              {{ formatCurrency(item.cost) }}
            </span>
          </template>

          <!-- Work Location Column -->
          <template #item.workLocation="{ item }">
            <v-chip size="small" color="success" variant="elevated" class="location-chip">
              {{ item.workLocation }}
            </v-chip>
          </template>

          <!-- Notes Column -->
          <template #item.notes="{ item }">
            <span class="text-caption" style="max-width: 200px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">
              {{ item.notes || 'لا توجد ملاحظات' }}
            </span>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="d-flex gap-2">
              <v-btn
                size="small"
                color="primary"
                variant="elevated"
                @click="editExpense(item)"
                icon
                class="action-btn edit-btn"
              >
                <v-icon size="16">mdi-pencil</v-icon>
              </v-btn>
              <v-btn
                size="small"
                color="error"
                variant="elevated"
                @click="deleteExpense(item)"
                icon
                class="action-btn delete-btn"
              >
                <v-icon size="16">mdi-delete</v-icon>
              </v-btn>
            </div>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Project Dialog - Image Style Design -->
      <v-dialog v-model="dialog" max-width="600" scrollable persistent>
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-folder-plus</v-icon>
                <span class="header-title">{{ isEditing ? 'تعديل المشروع' : 'إضافة مشروع' }}</span>
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

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="form" v-model="formValid">
              <div class="form-fields">
                <!-- Project Name -->
                <div class="field-group">
                  <label class="field-label">اسم المشروع</label>
                  <v-text-field
                    v-model="projectForm.name"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل اسم المشروع"
                    :rules="[v => !!v || 'اسم المشروع مطلوب']"
                    required
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>

                <!-- Project Type -->
                <div class="field-group">
                  <label class="field-label">نوع المشروع</label>
                  <v-text-field
                    v-model="projectForm.type"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل نوع المشروع"
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>

                <!-- Initial Cost -->
                <div class="field-group">
                  <label class="field-label">التكلفة المبدئية (د.ع)</label>
                  <v-text-field
                    v-model.number="projectForm.initialCost"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    placeholder="0"
                    :rules="[v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر']"
                    required
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>

                <!-- Duration -->
                <div class="field-group">
                  <label class="field-label">مدة المشروع (أيام)</label>
                  <v-text-field
                    v-model.number="projectForm.duration"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    placeholder="0"
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>

                <!-- Total Cost (calculated) -->
                <div class="field-group">
                  <label class="field-label">التكلفة الإجمالية (د.ع)</label>
                  <v-text-field
                    :value="projectForm.initialCost * (projectForm.duration || 1)"
                    variant="outlined"
                    density="comfortable"
                    readonly
                    class="custom-input readonly-field"
                    hide-details="auto"
                  />
                </div>

                <!-- Progress -->
                <div class="field-group">
                  <label class="field-label">نسبة الإنجاز (%)</label>
                  <v-text-field
                    v-model.number="projectForm.progress"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    min="0"
                    max="100"
                    class="custom-input"
                    hide-details="auto"
                    :rules="[
                      v => v >= 0 || 'النسبة يجب أن تكون أكبر من أو تساوي 0',
                      v => v <= 100 || 'النسبة يجب أن تكون أقل من أو تساوي 100'
                    ]"
                  />
                </div>

                <!-- Phone Number -->
                <div class="field-group">
                  <label class="field-label">رقم الهاتف</label>
                  <v-text-field
                    v-model="projectForm.phone"
                    variant="outlined"
                    density="comfortable"
                    placeholder="07XX XXX XXXX"
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>

                <!-- Location -->
                <div class="field-group">
                  <label class="field-label">مكان المشروع</label>
                  <v-text-field
                    v-model="projectForm.location"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل مكان المشروع"
                    :rules="[v => !!v || 'مكان المشروع مطلوب']"
                    required
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>

                <!-- Notes -->
                <div class="field-group">
                  <label class="field-label">ملاحظات</label>
                  <v-textarea
                    v-model="projectForm.description"
                    variant="outlined"
                    density="comfortable"
                    placeholder="أدخل ملاحظات إضافية"
                    rows="3"
                    class="custom-input"
                    hide-details="auto"
                  />
                </div>
              </div>
            </v-form>
          </div>

          <!-- Footer Actions -->
          <div class="dialog-footer">
            <v-btn
              variant="outlined"
              @click="closeDialog"
              class="cancel-btn"
              size="large"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveProject"
              :disabled="!formValid"
              class="save-btn"
              size="large"
            >
              {{ isEditing ? 'تحديث المشروع' : 'إضافة المشروع' }}
            </v-btn>
          </div>
        </v-card>
      </v-dialog>
          
      <!-- Project Details Dialog - Removed, now using separate page -->
      <v-dialog v-model="detailsDialog" max-width="900">
        <v-card class="details-card">
          <v-card-title class="d-flex align-center justify-space-between details-header">
          <div class="d-flex align-center">
              <v-icon class="me-2" color="info" size="large">mdi-information</v-icon>
              <span class="text-h5 font-weight-bold">تفاصيل المشروع</span>
            </div>
            <v-btn
              icon="mdi-close"
              variant="text"
              @click="closeDetailsDialog"
              class="close-btn"
            />
          </v-card-title>

          <v-card-text class="details-content" v-if="selectedProjectDetails">
            <v-row>
              <!-- Basic Information -->
              <v-col cols="12" md="6">
                <v-card class="info-card mb-4" elevation="2">
                  <v-card-title class="info-card-title">
                    <v-icon class="me-2" color="primary">mdi-folder-information</v-icon>
                    المعلومات الأساسية
                  </v-card-title>
                  <v-card-text>
                    <div class="info-item">
                      <span class="info-label">اسم المشروع:</span>
                      <span class="info-value">{{ selectedProjectDetails.name }}</span>
          </div>
                    <div class="info-item">
                      <span class="info-label">الموقع:</span>
                      <span class="info-value">
                        <v-icon size="small" class="me-1">mdi-map-marker</v-icon>
                        {{ selectedProjectDetails.location }}
                      </span>
        </div>
                    <div class="info-item">
                      <span class="info-label">المسؤول:</span>
                      <span class="info-value">{{ selectedProjectDetails.user }}</span>
                    </div>
                    <div class="info-item">
                      <span class="info-label">فئة المشروع:</span>
                      <span class="info-value">
                        <v-chip size="small" color="primary">{{ selectedProjectDetails.category }}</v-chip>
                      </span>
                    </div>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- Status and Priority -->
              <v-col cols="12" md="6">
                <v-card class="info-card mb-4" elevation="2">
                  <v-card-title class="info-card-title">
                    <v-icon class="me-2" color="success">mdi-chart-line</v-icon>
                    الحالة والأولوية
                  </v-card-title>
                  <v-card-text>
                    <div class="info-item">
                      <span class="info-label">الحالة:</span>
                      <v-chip 
                        :color="getStatusColor(selectedProjectDetails.status)"
                        size="small"
                      >
                        {{ getStatusText(selectedProjectDetails.status) }}
                      </v-chip>
              </div>
                    <div class="info-item">
                      <span class="info-label">الأولوية:</span>
                      <v-chip 
                        :color="getPriorityColor(selectedProjectDetails.priority)"
                        size="small"
                      >
                        {{ getPriorityText(selectedProjectDetails.priority) }}
                      </v-chip>
            </div>
                    <div class="info-item">
                      <span class="info-label">تاريخ البدء:</span>
                      <span class="info-value">{{ formatDate(selectedProjectDetails.startDate) }}</span>
                    </div>
                    </v-card-text>
                  </v-card>
              </v-col>

              <!-- Financial Information -->
              <v-col cols="12" md="6">
                <v-card class="info-card mb-4" elevation="2">
                  <v-card-title class="info-card-title">
                    <v-icon class="me-2" color="warning">mdi-currency-usd</v-icon>
                    المعلومات المالية
                  </v-card-title>
                  <v-card-text>
                    <div class="info-item">
                      <span class="info-label">التكلفة المبدئية:</span>
                      <span class="info-value text-success font-weight-bold">
                        {{ formatCurrency(selectedProjectDetails.initialCost) }}
                      </span>
                    </div>
                    <div class="info-item">
                      <span class="info-label">التكلفة الحرجة:</span>
                      <span class="info-value text-warning font-weight-bold">
                        {{ formatCurrency(selectedProjectDetails.criticalCost) }}
                      </span>
                    </div>
                    <div class="info-item">
                      <span class="info-label">الفرق:</span>
                      <span 
                        class="info-value font-weight-bold"
                        :class="selectedProjectDetails.criticalCost - selectedProjectDetails.initialCost > 0 ? 'text-error' : 'text-success'"
                      >
                        {{ formatCurrency(selectedProjectDetails.criticalCost - selectedProjectDetails.initialCost) }}
                      </span>
                    </div>
                  </v-card-text>
                </v-card>
              </v-col>

              <!-- Working Days -->
              <v-col cols="12" md="6">
                <v-card class="info-card" elevation="2">
                  <v-card-title class="info-card-title">
                    <v-icon class="me-2" color="success">mdi-calendar-check</v-icon>
                    أيام العمل
                  </v-card-title>
                  <v-card-text>
                    <div class="working-days-list">
                      <div 
                        v-for="day in selectedProjectDetails.workingDays || getDefaultWorkingDays()" 
                        :key="day.name"
                        class="working-day-item"
                        :class="{ 'active': day.isActive }"
                      >
                        <v-icon 
                          :color="day.isActive ? 'success' : 'grey-lighten-1'" 
                          size="small"
                          class="me-2"
                        >
                          {{ day.isActive ? 'mdi-check-circle' : 'mdi-circle-outline' }}
                        </v-icon>
                        <span class="day-name">{{ day.name }}</span>
                        <span class="day-hours">{{ day.hours }}</span>
                      </div>
                      </div>
                    <div class="total-hours">
                      <v-divider class="my-3"></v-divider>
                      <div class="d-flex justify-space-between align-center">
                        <span class="font-weight-bold">إجمالي ساعات العمل:</span>
                        <v-chip color="success" size="small">
                          {{ getTotalWorkingHours() }} ساعة
                        </v-chip>
                      </div>
                      <div class="mt-3">
                        <v-btn
                          color="success"
                          variant="elevated"
                          size="small"
                          @click="exportWorkingDaysToExcel"
                          class="export-excel-btn"
                        >
                          <v-icon left size="small">mdi-file-excel</v-icon>
                          تحميل Excel
                        </v-btn>
                      </div>
                    </div>
                    </v-card-text>
                  </v-card>
              </v-col>

              <!-- Project Actions -->
              <v-col cols="12" md="6">
                <v-card class="info-card" elevation="2">
                  <v-card-title class="info-card-title">
                    <v-icon class="me-2" color="primary">mdi-cog</v-icon>
                    إجراءات المشروع
                  </v-card-title>
                  <v-card-text>
                    <div class="actions-list">
                      <div 
                        v-for="action in selectedProjectDetails.actions || getDefaultActions()" 
                        :key="action.id"
                        class="action-item"
                      >
                        <div class="action-header">
                          <v-icon 
                            :color="action.status === 'completed' ? 'success' : action.status === 'in-progress' ? 'warning' : 'grey'"
                            size="small"
                            class="me-2"
                          >
                            {{ action.status === 'completed' ? 'mdi-check' : action.status === 'in-progress' ? 'mdi-clock' : 'mdi-circle-outline' }}
                          </v-icon>
                          <span class="action-title">{{ action.title }}</span>
                          <v-chip 
                            :color="getActionStatusColor(action.status)"
                            size="x-small"
                            class="ms-auto"
                          >
                            {{ getActionStatusText(action.status) }}
                          </v-chip>
                </div>
                        <div class="action-details">
                          <small class="text-medium-emphasis">
                            {{ action.description }}
                          </small>
                          <div class="action-date">
                            <v-icon size="x-small" class="me-1">mdi-calendar</v-icon>
                            <small>{{ action.dueDate || 'غير محدد' }}</small>
              </div>
            </div>
          </div>
                    </div>
                  </v-card-text>
        </v-card>
              </v-col>

              <!-- Description -->
              <v-col cols="12">
                <v-card class="info-card" elevation="2">
                  <v-card-title class="info-card-title">
                    <v-icon class="me-2" color="info">mdi-text</v-icon>
                    وصف المشروع
                  </v-card-title>
                  <v-card-text>
                    <div class="description-text">
                      {{ selectedProjectDetails.description || 'لا يوجد وصف متاح' }}
      </div>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
          </v-card-text>

          <v-card-actions class="details-actions">
            <v-spacer />
            <v-btn
              color="primary"
              variant="outlined"
              @click="closeDetailsDialog"
              class="close-details-btn"
              size="large"
              rounded="xl"
            >
              <v-icon class="me-2">mdi-close</v-icon>
              إغلاق
            </v-btn>
            <v-btn
              color="warning"
              variant="elevated"
              @click="editProjectFromDetails"
              class="edit-details-btn"
              size="large"
              rounded="xl"
            >
              <v-icon class="me-2">mdi-pencil</v-icon>
              تعديل المشروع
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Delete Confirmation Dialog -->
      <v-dialog v-model="deleteDialog" max-width="400">
        <v-card>
          <v-card-title class="text-h6">تأكيد الحذف</v-card-title>
          <v-card-text>
            هل أنت متأكد من حذف المشروع "{{ selectedProject?.name }}"؟
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
    </div>
      <!-- Add/Edit Administrative Expense Dialog - Hidden -->
      <v-dialog v-model="expenseDialog" max-width="700" scrollable persistent v-show="false">
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-currency-usd</v-icon>
                <span class="header-title">{{ isEditingExpense ? 'تعديل المصروف الإداري' : 'إضافة مصروف إداري' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeExpenseDialog"
                class="close-btn"
              />
            </div>
          </div>

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="expenseForm" v-model="expenseFormValid">
              <div class="form-fields">
                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.projectName"
                      label="اسم المشروع"
                      variant="outlined"
                      :rules="[v => !!v || 'اسم المشروع مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.cost"
                      label="التكلفة (د.ع)"
                      variant="outlined"
                      type="number"
                      :rules="[v => !!v || 'التكلفة مطلوبة', v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.startDate"
                      label="تاريخ البداية"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ البداية مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.endDate"
                      label="تاريخ الانتهاء"
                      variant="outlined"
                      type="date"
                      :rules="[v => !!v || 'تاريخ الانتهاء مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="expenseForm.workLocation"
                      label="مكان العمل"
                      variant="outlined"
                      :rules="[v => !!v || 'مكان العمل مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-select
                      v-model="expenseForm.expenseType"
                      :items="expenseTypes"
                      label="نوع المصروف"
                      variant="outlined"
                      :rules="[v => !!v || 'نوع المصروف مطلوب']"
                      required
                      class="form-field"
                    />
                  </v-col>
                </v-row>

                <v-row>
                  <v-col cols="12">
                    <v-textarea
                      v-model="expenseForm.notes"
                      label="الملاحظات"
                      variant="outlined"
                      rows="3"
                      class="form-field"
                    />
                  </v-col>
                </v-row>
              </div>
            </v-form>
          </div>

          <!-- Dialog Actions -->
          <div class="dialog-actions">
            <v-btn
              color="grey"
              variant="text"
              @click="closeExpenseDialog"
              class="action-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="error"
              variant="elevated"
              @click="saveExpense"
              :disabled="!expenseFormValid"
              class="action-btn primary-btn"
            >
              <v-icon class="me-2">mdi-content-save</v-icon>
              {{ isEditingExpense ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </div>
        </v-card>
      </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

// عنوان الصفحة
document.title = 'إدارة المشاريع الهندسية - نظام إدارة المشاريع'

const router = useRouter()

// متغيرات الحالة الأساسية
const loading = ref(false)
const dialog = ref(false)
const deleteDialog = ref(false)
const detailsDialog = ref(false)
const formValid = ref(false)
const isEditing = ref(false)
const searchQuery = ref('')
const selectedCategory = ref('')
const selectedStatus = ref('')
const selectedProject = ref(null)
const selectedProjectDetails = ref(null)

// Team Management
const showTeamManagement = ref(false)
const teamLoading = ref(false)
const showAddMemberDialog = ref(false)
const teamMemberFormValid = ref(false)
const teamMemberForm = ref({
  name: '',
  email: '',
  phone: '',
  department: '',
  role: '',
  status: 'نشط'
})
const teamMembers = ref([
  {
    id: 1,
    name: 'أحمد محمد علي',
    email: 'ahmed.mohamed@company.com',
    phone: '+964 770 123 4567',
    department: 'الهندسة المدنية',
    role: 'مهندس أول',
    status: 'نشط',
    tasksCount: 5,
    avatar: null
  },
  {
    id: 2,
    name: 'فاطمة حسن',
    email: 'fatima.hassan@company.com',
    phone: '+964 770 234 5678',
    department: 'الهندسة المعمارية',
    role: 'مهندسة معمارية',
    status: 'نشط',
    tasksCount: 3,
    avatar: null
  },
  {
    id: 3,
    name: 'محمد عبدالله',
    email: 'mohamed.abdullah@company.com',
    phone: '+964 770 345 6789',
    department: 'إدارة المشاريع',
    role: 'مدير مشروع',
    status: 'نشط',
    tasksCount: 8,
    avatar: null
  },
  {
    id: 4,
    name: 'سارة أحمد',
    email: 'sara.ahmed@company.com',
    phone: '+964 770 456 7890',
    department: 'الهندسة الكهربائية',
    role: 'مهندسة كهربائية',
    status: 'غير نشط',
    tasksCount: 2,
    avatar: null
  },
  {
    id: 5,
    name: 'علي محمود',
    email: 'ali.mahmoud@company.com',
    phone: '+964 770 567 8901',
    department: 'الهندسة الميكانيكية',
    role: 'مهندس ميكانيكي',
    status: 'نشط',
    tasksCount: 4,
    avatar: null
  }
])

const teamHeaders = [
  { title: 'الصورة', key: 'avatar', sortable: false, width: '60px' },
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'القسم', key: 'department', sortable: true },
  { title: 'الدور', key: 'role', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'المهام', key: 'tasks', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false, width: '120px' }
]

// خيارات القوائم المنسدلة
const departmentOptions = [
  'الهندسة المدنية',
  'الهندسة المعمارية', 
  'الهندسة الكهربائية',
  'الهندسة الميكانيكية',
  'إدارة المشاريع',
  'الموارد البشرية',
  'المحاسبة',
  'التسويق'
]

const roleOptions = [
  'مدير المشروع',
  'مهندس أول',
  'مهندس',
  'مساعد مهندس',
  'فني',
  'مساعد إداري',
  'محاسب',
  'مستشار'
]

const statusOptions = [
  'نشط',
  'غير نشط',
  'في إجازة',
  'مستقيل'
]

// متغيرات المصاريف الإدارية
const expenseDialog = ref(false)
const expenseFormValid = ref(false)
const isEditingExpense = ref(false)
const expenseSearchQuery = ref('')
const selectedProjectFilter = ref('')
const selectedCostRange = ref('')
const selectedExpense = ref(null)

// جدول المصاريف الإدارية
const expenseHeaders = [
  { title: 'اسم المشروع', key: 'projectName', sortable: true },
  { title: 'تاريخ البداية', key: 'startDate', sortable: true, width: '120px' },
  { title: 'تاريخ الانتهاء', key: 'endDate', sortable: true, width: '120px' },
  { title: 'التكلفة', key: 'cost', sortable: true, width: '120px' },
  { title: 'مكان العمل', key: 'workLocation', sortable: true, width: '120px' },
  { title: 'الملاحظات', key: 'notes', sortable: false, width: '200px' },
  { title: 'الإجراءات', key: 'actions', sortable: false, width: '120px' }
]

// نموذج المصاريف الإدارية
const expenseForm = ref({
  projectName: '',
  startDate: '',
  endDate: '',
  cost: '',
  workLocation: '',
  expenseType: '',
  notes: ''
})

// بيانات المصاريف الإدارية التجريبية
const administrativeExpenses = ref([
  {
    id: 1,
    projectName: 'مشروع تطوير الموقع الإلكتروني',
    startDate: '2024-01-15',
    endDate: '2024-03-15',
    cost: 50000,
    workLocation: 'بغداد',
    expenseType: 'تطوير',
    notes: 'مصروفات تطوير وتصميم الموقع'
  },
  {
    id: 2,
    projectName: 'مشروع تحديث النظام الأكاديمي',
    startDate: '2024-02-01',
    endDate: '2024-04-01',
    cost: 75000,
    workLocation: 'البصرة',
    expenseType: 'تحديث',
    notes: 'تحديث النظام الأكاديمي الحالي'
  },
  {
    id: 3,
    projectName: 'مشروع بناء المكتبة الرقمية',
    startDate: '2024-03-01',
    endDate: '2024-06-01',
    cost: 100000,
    workLocation: 'أربيل',
    expenseType: 'بناء',
    notes: 'إنشاء مكتبة رقمية شاملة'
  },
  {
    id: 4,
    projectName: 'مشروع تحديث المختبرات',
    startDate: '2024-04-01',
    endDate: '2024-07-01',
    cost: 125000,
    workLocation: 'الموصل',
    expenseType: 'تحديث',
    notes: 'تحديث وتطوير المختبرات العلمية'
  },
  {
    id: 5,
    projectName: 'مشروع الأمن السيبراني',
    startDate: '2024-05-01',
    endDate: '2024-08-01',
    cost: 80000,
    workLocation: 'النجف',
    expenseType: 'أمن',
    notes: 'تطوير أنظمة الأمن السيبراني'
  }
])

// خيارات فلترة المصاريف
const projectFilterOptions = computed(() => {
  const projects = [...new Set(administrativeExpenses.value.map(expense => expense.projectName))]
  return projects.map(project => ({ title: project, value: project }))
})

const costRangeOptions = [
  { title: 'أقل من 50,000 د.ع', value: 'low' },
  { title: '50,000 - 100,000 د.ع', value: 'medium' },
  { title: 'أكثر من 100,000 د.ع', value: 'high' }
]

const expenseTypes = [
  'تطوير',
  'تحديث',
  'بناء',
  'أمن',
  'صيانة',
  'تدريب',
  'أخرى'
]

// Table headers
const tableHeaders = [
  { title: 'اسم المشروع', key: 'name', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true, width: '120px' },
  { title: 'الموقع', key: 'location', sortable: true, width: '120px' },
  { title: 'الميزانية', key: 'budget', sortable: false, width: '150px' },
  { title: 'تاريخ البدء', key: 'startDate', sortable: true, width: '120px' },
  { title: 'المسؤول', key: 'manager', sortable: true, width: '120px' },
  { title: 'الإجراءات', key: 'actions', sortable: false, width: '120px' }
]


// Sample projects data
const sampleProjects = [
    {
      id: 1,
    name: 'مشروع تطوير الموقع الإلكتروني',
    description: 'تطوير موقع إلكتروني جديد للجامعة',
    location: 'بغداد',
    status: 'active',
    initialCost: 50000,
    criticalCost: 75000,
    startDate: '2024-01-15',
    user: 'أحمد محمد'
    },
    {
      id: 2,
    name: 'مشروع تحديث النظام الأكاديمي',
    description: 'تحديث النظام الأكاديمي الحالي',
    location: 'البصرة',
    status: 'pending',
    initialCost: 75000,
    criticalCost: 100000,
    startDate: '2024-02-01',
    user: 'فاطمة علي'
    },
    {
      id: 3,
    name: 'مشروع بناء المكتبة الرقمية',
    description: 'بناء مكتبة رقمية شاملة',
    location: 'أربيل',
    status: 'completed',
    initialCost: 100000,
    criticalCost: 150000,
    startDate: '2024-01-30',
    user: 'محمد السعيد'
  }
]

// نموذج البيانات
const projectForm = ref({
  name: '',
  type: '',
  location: '',
  initialCost: 0,
  duration: 0,
  startDate: '',
  criticalCost: 0,
  phone: '',
  user: '',
  status: 'pending',
  priority: 'medium',
  progress: 0,
  description: '',
  category: ''
})

// البيانات الأساسية
const projects = ref([
    {
      id: 1,
    name: 'مشروع تطوير الموقع الإلكتروني',
    location: 'بغداد',
    initialCost: 50000,
    startDate: '2024-01-15',
    criticalCost: 75000,
    user: 'أحمد محمد',
    status: 'active',
    priority: 'high',
    progress: 65,
    description: 'تطوير موقع إلكتروني جديد للجامعة',
    category: 'تطوير',
    workingDays: [
      { name: 'السبت', hours: '8 ساعات', isActive: true },
      { name: 'الأحد', hours: '8 ساعات', isActive: true },
      { name: 'الاثنين', hours: '8 ساعات', isActive: true },
      { name: 'الثلاثاء', hours: '8 ساعات', isActive: true },
      { name: 'الأربعاء', hours: '6 ساعات', isActive: true },
      { name: 'الخميس', hours: '0 ساعة', isActive: false },
      { name: 'الجمعة', hours: '0 ساعة', isActive: false }
    ],
    actions: [
      { id: 1, title: 'تحليل المتطلبات', description: 'تحليل متطلبات المشروع وتحديد المواصفات', status: 'completed', dueDate: '2024-01-20' },
      { id: 2, title: 'تصميم الواجهة', description: 'تصميم واجهة المستخدم وتجربة المستخدم', status: 'in-progress', dueDate: '2024-02-15' },
      { id: 3, title: 'تطوير الواجهة الأمامية', description: 'تطوير واجهة المستخدم باستخدام React', status: 'pending', dueDate: '2024-03-01' },
      { id: 4, title: 'تطوير الخادم الخلفي', description: 'تطوير API والخادم الخلفي', status: 'pending', dueDate: '2024-03-15' },
      { id: 5, title: 'الاختبار والمراجعة', description: 'اختبار النظام ومراجعة الكود', status: 'pending', dueDate: '2024-04-01' }
      ]
    },
    {
      id: 2,
    name: 'مشروع تحديث النظام الأكاديمي',
    location: 'البصرة',
    initialCost: 75000,
    startDate: '2024-02-01',
    criticalCost: 100000,
    user: 'فاطمة علي',
    status: 'pending',
    priority: 'medium',
    progress: 30,
    description: 'تحديث النظام الأكاديمي الحالي',
    category: 'تحديث',
    workingDays: [
      { name: 'السبت', hours: '6 ساعات', isActive: true },
      { name: 'الأحد', hours: '8 ساعات', isActive: true },
      { name: 'الاثنين', hours: '8 ساعات', isActive: true },
      { name: 'الثلاثاء', hours: '8 ساعات', isActive: true },
      { name: 'الأربعاء', hours: '8 ساعات', isActive: true },
      { name: 'الخميس', hours: '4 ساعات', isActive: true },
      { name: 'الجمعة', hours: '0 ساعة', isActive: false }
    ],
    actions: [
      { id: 1, title: 'مراجعة النظام الحالي', description: 'مراجعة النظام الأكاديمي الموجود وتحديد نقاط التحسين', status: 'completed', dueDate: '2024-01-25' },
      { id: 2, title: 'تحديث قاعدة البيانات', description: 'تحديث هيكل قاعدة البيانات وتحسين الأداء', status: 'in-progress', dueDate: '2024-02-20' },
      { id: 3, title: 'تطوير واجهات جديدة', description: 'تطوير واجهات مستخدم محسنة', status: 'pending', dueDate: '2024-03-10' },
      { id: 4, title: 'تدريب المستخدمين', description: 'تدريب الموظفين على النظام الجديد', status: 'pending', dueDate: '2024-03-25' }
      ]
    },
    {
      id: 3,
    name: 'مشروع بناء المكتبة الرقمية',
    location: 'أربيل',
    initialCost: 100000,
    startDate: '2024-01-30',
    criticalCost: 150000,
    user: 'محمد السعيد',
    status: 'completed',
    priority: 'high',
    progress: 100,
    description: 'بناء مكتبة رقمية شاملة',
    category: 'بناء',
    workingDays: [
      { name: 'السبت', hours: '8 ساعات', isActive: true },
      { name: 'الأحد', hours: '8 ساعات', isActive: true },
      { name: 'الاثنين', hours: '8 ساعات', isActive: true },
      { name: 'الثلاثاء', hours: '8 ساعات', isActive: true },
      { name: 'الأربعاء', hours: '8 ساعات', isActive: true },
      { name: 'الخميس', hours: '0 ساعة', isActive: false },
      { name: 'الجمعة', hours: '0 ساعة', isActive: false }
    ],
    actions: [
      { id: 1, title: 'تصميم المكتبة الرقمية', description: 'تصميم هيكل المكتبة الرقمية وتحديد الميزات المطلوبة', status: 'completed', dueDate: '2024-01-30' },
      { id: 2, title: 'تطوير نظام البحث', description: 'تطوير نظام بحث متقدم للكتب والمصادر', status: 'completed', dueDate: '2024-02-10' },
      { id: 3, title: 'إضافة الكتب والمصادر', description: 'إضافة مجموعة شاملة من الكتب والمصادر الرقمية', status: 'in-progress', dueDate: '2024-02-28' },
      { id: 4, title: 'تطوير نظام الإعارة', description: 'تطوير نظام إعارة الكتب الإلكترونية', status: 'pending', dueDate: '2024-03-15' },
      { id: 5, title: 'الاختبار النهائي', description: 'اختبار جميع وظائف المكتبة الرقمية', status: 'pending', dueDate: '2024-03-30' }
    ]
  }
])

// قوائم الاختيار - تم حذف statusOptions المكرر

const priorityOptions = [
  { title: 'منخفضة', value: 'low' },
  { title: 'متوسطة', value: 'medium' },
  { title: 'عالية', value: 'high' },
  { title: 'عاجلة', value: 'urgent' }
]

// قوائم الفلترة
const filterCategories = [
  { title: 'تطوير', value: 'تطوير' },
  { title: 'تحديث', value: 'تحديث' },
  { title: 'بناء', value: 'بناء' },
  { title: 'صيانة', value: 'صيانة' },
  { title: 'استشارات', value: 'استشارات' }
]

const filterStatuses = [
  { title: 'في الانتظار', value: 'pending' },
  { title: 'نشط', value: 'active' },
  { title: 'مكتمل', value: 'completed' },
  { title: 'ملغي', value: 'cancelled' }
]


// الخصائص المحسوبة
const totalProjects = computed(() => {
  return projects.value.length
})

const activeProjects = computed(() => {
  return projects.value.filter(p => p.status === 'active').length
})

const pendingProjects = computed(() => {
  return projects.value.filter(p => p.status === 'pending').length
})

const averageProgress = computed(() => {
  if (projects.value.length === 0) return 0
  const totalProgress = projects.value.reduce((sum, project) => sum + (project.progress || 0), 0)
  return Math.round(totalProgress / projects.value.length)
})

// إحصائيات الفريق
const activeTeamMembers = computed(() => teamMembers.value.filter(member => member.status === 'نشط').length)
const teamDepartments = computed(() => new Set(teamMembers.value.map(member => member.department)).size)
const totalTeamTasks = computed(() => teamMembers.value.reduce((sum, member) => sum + member.tasksCount, 0))

const completedProjects = computed(() => {
  return projects.value.filter(p => p.progress === 100).length
})

const totalBudget = computed(() => {
  return projects.value.reduce((sum, project) => sum + project.initialCost, 0)
})

const filteredProjects = computed(() => {
  let filtered = projects.value

  // فلترة حسب البحث النصي
  if (searchQuery.value) {
    filtered = filtered.filter(project =>
      project.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      project.location.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      project.user.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      project.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      project.category.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  // فلترة حسب الفئة
  if (selectedCategory.value) {
    filtered = filtered.filter(project => project.category === selectedCategory.value)
  }

  // فلترة حسب الحالة
  if (selectedStatus.value) {
    filtered = filtered.filter(project => project.status === selectedStatus.value)
  }

  return filtered
})

// الدوال المساعدة
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const getStatusColor = (status) => {
  const colors = {
    pending: 'warning',
    active: 'success',
    completed: 'primary',
    cancelled: 'error'
  }
  return colors[status] || 'grey'
}

const getStatusText = (status) => {
  const texts = {
    pending: 'في الانتظار',
    active: 'نشط',
    completed: 'مكتمل',
    cancelled: 'ملغي'
  }
  return texts[status] || status
}

const getStatusIcon = (status) => {
  const icons = {
    pending: 'mdi-clock-outline',
    active: 'mdi-play-circle',
    completed: 'mdi-check-circle',
    cancelled: 'mdi-cancel'
  }
  return icons[status] || 'mdi-help-circle'
}


const getPriorityText = (priority) => {
  const texts = {
    low: 'منخفضة',
    medium: 'متوسطة',
    high: 'عالية',
    urgent: 'عاجلة'
  }
  return texts[priority] || priority
}

const getPriorityColor = (priority) => {
  const colors = {
    low: 'success',
    medium: 'warning',
    high: 'error',
    urgent: 'purple'
  }
  return colors[priority] || 'grey'
}

const formatDate = (dateString) => {
  if (!dateString) return 'غير محدد'
  const date = new Date(dateString)
  return date.toLocaleDateString('en-GB', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

// دوال أيام العمل والإجراءات
const getDefaultWorkingDays = () => {
  return [
    { name: 'السبت', hours: '8 ساعات', isActive: true },
    { name: 'الأحد', hours: '8 ساعات', isActive: true },
    { name: 'الاثنين', hours: '8 ساعات', isActive: true },
    { name: 'الثلاثاء', hours: '8 ساعات', isActive: true },
    { name: 'الأربعاء', hours: '8 ساعات', isActive: true },
    { name: 'الخميس', hours: '4 ساعات', isActive: true },
    { name: 'الجمعة', hours: '0 ساعة', isActive: false }
  ]
}

const getDefaultActions = () => {
  return [
    { id: 1, title: 'تحديد المتطلبات', description: 'تحديد متطلبات المشروع الأساسية', status: 'pending', dueDate: null },
    { id: 2, title: 'التخطيط', description: 'وضع خطة زمنية للمشروع', status: 'pending', dueDate: null },
    { id: 3, title: 'التنفيذ', description: 'بدء تنفيذ المشروع', status: 'pending', dueDate: null },
    { id: 4, title: 'الاختبار', description: 'اختبار المشروع والتأكد من جودته', status: 'pending', dueDate: null }
  ]
}

const getTotalWorkingHours = () => {
  if (!selectedProjectDetails.value?.workingDays) {
    return getDefaultWorkingDays().reduce((total, day) => {
      return day.isActive ? total + parseInt(day.hours) : total
    }, 0)
  }
  
  return selectedProjectDetails.value.workingDays.reduce((total, day) => {
    return day.isActive ? total + parseInt(day.hours) : total
  }, 0)
}

// دالة تصدير أيام العمل إلى Excel
const exportWorkingDaysToExcel = () => {
  try {
    const workingDays = selectedProjectDetails.value?.workingDays || getDefaultWorkingDays()
    const projectName = selectedProjectDetails.value?.name || 'مشروع غير محدد'
    
    // إعداد البيانات للتصدير
    const excelData = [
      ['اسم المشروع', projectName],
      ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
      [''],
      ['يوم العمل', 'الحالة', 'عدد الساعات', 'ملاحظات']
    ]
    
    // إضافة بيانات أيام العمل
    workingDays.forEach(day => {
      excelData.push([
        day.name,
        day.isActive ? 'نشط' : 'غير نشط',
        day.hours + ' ساعة',
        day.isActive ? 'يوم عمل عادي' : 'يوم إجازة'
      ])
    })
    
    // إضافة إجمالي الساعات
    excelData.push(['', '', '', ''])
    excelData.push(['إجمالي ساعات العمل', '', getTotalWorkingHours() + ' ساعة', ''])
    
    // تحويل البيانات إلى CSV
    const csvContent = excelData.map(row => 
      row.map(cell => `"${cell}"`).join(',')
    ).join('\n')
    
    // إضافة BOM للدعم العربي
    const BOM = '\uFEFF'
    const blob = new Blob([BOM + csvContent], { type: 'text/csv;charset=utf-8;' })
    
    // إنشاء رابط التحميل
    const link = document.createElement('a')
    const url = URL.createObjectURL(blob)
    link.setAttribute('href', url)
    link.setAttribute('download', `أيام_العمل_${projectName.replace(/\s+/g, '_')}_${new Date().toISOString().split('T')[0]}.csv`)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)

    // إظهار رسالة نجاح

  } catch (error) {
  }
}

const getActionStatusText = (status) => {
  const statusTexts = {
    'pending': 'في الانتظار',
    'in-progress': 'قيد التنفيذ',
    'completed': 'مكتمل',
    'cancelled': 'ملغي'
  }
  return statusTexts[status] || status
}

const getActionStatusColor = (status) => {
  const statusColors = {
    'pending': 'grey',
    'in-progress': 'warning',
    'completed': 'success',
    'cancelled': 'error'
  }
  return statusColors[status] || 'grey'
}

// دوال إدارة البيانات
const openAddProjectDialog = () => {
  isEditing.value = false
  projectForm.value = {
    name: '',
    location: '',
    initialCost: 0,
    startDate: '',
    criticalCost: 0,
    user: '',
    status: 'pending',
    priority: 'medium',
    description: '',
    category: ''
  }
  dialog.value = true
}

const viewProject = (project) => {
}

const viewProjectDetails = (project) => {
  // Navigate directly to work days management page
  router.push('/work-days')
}

const closeDetailsDialog = () => {
  detailsDialog.value = false
  selectedProjectDetails.value = null
}

const editProjectFromDetails = () => {
  if (selectedProjectDetails.value) {
    editProject(selectedProjectDetails.value)
    closeDetailsDialog()
  }
}

const editProject = (project) => {
  isEditing.value = true
  projectForm.value = { ...project }
  selectedProject.value = project
  dialog.value = true
}

const deleteProject = (project) => {
  selectedProject.value = project
  deleteDialog.value = true
}

const saveProject = () => {
  if (formValid.value) {
    if (isEditing.value) {
      const index = projects.value.findIndex(p => p.id === selectedProject.value.id)
      if (index > -1) {
        projects.value[index] = { ...projectForm.value, id: selectedProject.value.id }
      }
    } else {
      const newProject = {
        ...projectForm.value,
        id: Date.now()
      }
      projects.value.unshift(newProject)
    }
    closeDialog()
  }
}

const closeDialog = () => {
  dialog.value = false
  isEditing.value = false
  selectedProject.value = null
  projectForm.value = {
    name: '',
    location: '',
    initialCost: 0,
    startDate: '',
    criticalCost: 0,
    user: '',
    status: 'pending',
    priority: 'medium',
    description: '',
    category: ''
  }
}

// دوال المصاريف الإدارية
const openAddExpenseDialog = () => {
  expenseDialog.value = true
  isEditingExpense.value = false
  selectedExpense.value = null
  expenseForm.value = {
    projectName: '',
    startDate: '',
    endDate: '',
    cost: '',
    workLocation: '',
    expenseType: '',
    notes: ''
  }
}

const closeExpenseDialog = () => {
  expenseDialog.value = false
  isEditingExpense.value = false
  selectedExpense.value = null
  expenseForm.value = {
    projectName: '',
    startDate: '',
    endDate: '',
    cost: '',
    workLocation: '',
    expenseType: '',
    notes: ''
  }
}

const editExpense = (expense) => {
  selectedExpense.value = expense
  isEditingExpense.value = true
  expenseForm.value = { ...expense }
  expenseDialog.value = true
}

const deleteExpense = (expense) => {
  const index = administrativeExpenses.value.findIndex(e => e.id === expense.id)
  if (index > -1) {
    administrativeExpenses.value.splice(index, 1)
  }
}

const saveExpense = () => {
  if (isEditingExpense.value) {
    // تحديث المصروف
    const index = administrativeExpenses.value.findIndex(e => e.id === selectedExpense.value.id)
    if (index > -1) {
      administrativeExpenses.value[index] = {
        ...expenseForm.value,
        id: selectedExpense.value.id
      }
    }
  } else {
    // إضافة مصروف جديد
    const newExpense = {
      ...expenseForm.value,
      id: Date.now(),
      cost: parseFloat(expenseForm.value.cost)
    }
    administrativeExpenses.value.push(newExpense)
  }
  closeExpenseDialog()
}

// فلترة المصاريف
const filteredExpenses = computed(() => {
  let filtered = administrativeExpenses.value

  if (selectedProjectFilter.value) {
    filtered = filtered.filter(expense => expense.projectName === selectedProjectFilter.value)
  }

  if (selectedCostRange.value) {
    filtered = filtered.filter(expense => {
      const cost = expense.cost
      switch (selectedCostRange.value) {
        case 'low':
          return cost < 50000
        case 'medium':
          return cost >= 50000 && cost <= 100000
        case 'high':
          return cost > 100000
        default:
          return true
      }
    })
  }

  return filtered
})

const confirmDelete = () => {
  if (selectedProject.value) {
    const index = projects.value.findIndex(p => p.id === selectedProject.value.id)
    if (index > -1) {
      projects.value.splice(index, 1)
    }
  }
  deleteDialog.value = false
  selectedProject.value = null
}



// وظائف إدارة الفريق
const getDepartmentColor = (department) => {
  const colors = {
    'الهندسة المدنية': 'blue',
    'الهندسة المعمارية': 'purple',
    'الهندسة الكهربائية': 'orange',
    'الهندسة الميكانيكية': 'green',
    'إدارة المشاريع': 'red',
    'الموارد البشرية': 'pink',
    'المحاسبة': 'teal',
    'التسويق': 'indigo'
  }
  return colors[department] || 'grey'
}

const getRoleColor = (role) => {
  if (role.includes('مدير')) return 'red'
  if (role.includes('أول')) return 'blue'
  if (role.includes('مهندس')) return 'green'
  return 'grey'
}

const viewTeamMember = (member) => {
  // TODO: إضافة نافذة عرض التفاصيل
}

const editTeamMember = (member) => {
  // TODO: إضافة نافذة التعديل
}

const deleteTeamMember = (member) => {
  // TODO: إضافة تأكيد الحذف
}

const closeAddMemberDialog = () => {
  showAddMemberDialog.value = false
  teamMemberForm.value = {
    name: '',
    email: '',
    phone: '',
    department: '',
    role: '',
    status: 'نشط'
  }
  teamMemberFormValid.value = false
}

const saveTeamMember = () => {
  if (teamMemberFormValid.value) {
    const newMember = {
      ...teamMemberForm.value,
      id: Date.now(),
      tasksCount: 0,
      avatar: null
    }
    teamMembers.value.push(newMember)
    closeAddMemberDialog()
  }
}

onMounted(() => {
  // Initialize data if needed
})
</script>

<style>
/* Import page styles */
@import './styles/project-management.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
