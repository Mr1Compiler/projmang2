
<template>
  <v-container class="fill-height data-page" fluid>
    <div class="fullscreen-content">
      <!-- Header Section -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="48" color="white">mdi-folder-multiple</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">المشاريع</h1>
              <p class="subtitle">إدارة وتتبع جميع المشاريع</p>
            </div>
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

      <!-- Summary Cards -->
      <v-row class="mb-6 stats-row" no-gutters>
        <v-col cols="12" sm="12" md="12" lg="12" xl="12" class="pa-2">
          <div class="stats-cards-container">
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-folder-multiple</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ displayTotalProjects }}</h3>
                  <p class="stat-label">إجمالي المشاريع</p>
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
                  <h3 class="stat-value">{{ displayActiveProjects }}</h3>
                  <p class="stat-label">مشاريع نشطة</p>
                </div>
              </div>
            </v-card>
            <v-card class="modern-stat-card stat-card-warning" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-clock-alert</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ displayPendingProjects }}</h3>
                  <p class="stat-label">في الانتظار</p>
                </div>
              </div>
            </v-card>
            <v-card class="modern-stat-card stat-card-info" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-currency-usd</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ displayTotalBudget }}</h3>
                  <p class="stat-label">إجمالي الميزانية</p>
                </div>
              </div>
            </v-card>
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-chart-line</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ displayAverageProgress }}</h3>
                  <p class="stat-label">متوسط التقدم</p>
                </div>
              </div>
            </v-card>
          </div>
        </v-col>
      </v-row>

      <!-- Projects Table -->
      <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
        <v-card-title class="d-flex align-center justify-space-between">
          <div class="d-flex align-center">
            <v-icon class="me-2" style="color: #4338ca;" size="28">mdi-folder-multiple</v-icon>
            <span class="text-h4 font-weight-black" style="color: #ffffff; font-family: 'Arial', 'Helvetica', sans-serif; text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2); letter-spacing: 0.5px;">قائمة مشاريع</span>
            <v-chip class="ms-3" color="primary" size="small" variant="elevated">{{ projects.length }}</v-chip>
          </div>
          <v-btn
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddProjectDialog"
            elevation="2"
            color="primary"
            style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;"
          >
            <v-icon class="me-2 icon-glow">mdi-plus</v-icon>
            إضافة مشروع جديد
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
        <v-row v-if="filteredProjects.length > 0" class="projects-grid-row">
          <v-col
            v-for="project in filteredProjects"
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
              :class="{ 'warning-border': project.currentSpending >= project.warningCost && project.warningCost > 0 }"
            >
              <!-- Warning Banner -->
              <v-alert
                v-if="project.currentSpending >= project.warningCost && project.warningCost > 0"
                type="warning"
                variant="tonal"
                density="compact"
                class="ma-0 rounded-0"
              >
                <v-icon class="me-1">mdi-alert</v-icon>
                تحذير: المصروفات تجاوزت حد التحذير!
              </v-alert>

              <!-- Project Header -->
              <v-card-title class="project-card-title">
                <div class="d-flex align-center">
                  <v-avatar size="40" :color="project.currentSpending >= project.warningCost && project.warningCost > 0 ? 'warning' : 'primary'" variant="tonal">
                    <v-icon>{{ project.currentSpending >= project.warningCost && project.warningCost > 0 ? 'mdi-alert' : 'mdi-folder-multiple' }}</v-icon>
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
                    <span class="detail-text">{{ project.location || 'غير محدد' }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="success">mdi-currency-usd</v-icon>
                    <span class="detail-text">الميزانية: {{ formatCurrency(project.totalCost) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="warning">mdi-alert-circle</v-icon>
                    <span class="detail-text">حد التنبيه: {{ formatCurrency(project.warningCost) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" :color="project.currentSpending > project.warningCost ? 'error' : 'info'">mdi-cash</v-icon>
                    <span class="detail-text" :class="project.currentSpending > project.warningCost ? 'text-error' : ''">المصروف: {{ formatCurrency(project.currentSpending) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="info">mdi-calendar</v-icon>
                    <span class="detail-text">{{ formatDate(project.startDate) }}</span>
                  </div>

                  <div class="detail-item">
                    <v-icon size="small" color="secondary">mdi-phone</v-icon>
                    <span class="detail-text">{{ project.clientPhone || 'غير محدد' }}</span>
                  </div>
                </div>

                <!-- Progress Bar -->
                <div class="progress-section">
                  <div class="progress-header">
                    <span class="progress-label">نسبة الإنجاز</span>
                    <span class="progress-percentage">{{ project.progress || 0 }}%</span>
                  </div>
                  <v-progress-linear
                    :model-value="project.progress || 0"
                    color="primary"
                    height="6"
                    rounded
                    class="progress-bar"
                  />
                </div>
              </v-card-text>

              <v-card-actions class="project-card-actions d-flex ga-2">
                <v-btn
                  color="warning"
                  variant="elevated"
                  size="default"
                  @click="editProject(project)"
                  class="edit-btn flex-grow-1 d-flex justify-center align-center"
                >
                  <v-icon class="me-2">mdi-pencil</v-icon>
                  <span>تعديل</span>
                </v-btn>
                <v-btn
                  color="primary"
                  variant="elevated"
                  size="default"
                  @click="viewProjectDetails(project)"
                  class="details-btn flex-grow-1 d-flex justify-center align-center"
                >
                  <v-icon class="me-2">mdi-eye</v-icon>
                  <span>عرض التفاصيل</span>
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>

        <!-- Pagination -->
        <div class="d-flex justify-center pa-4" v-if="totalPages > 0">
          <v-pagination
            v-model="currentPage"
            :length="totalPages"
            :total-visible="7"
            @update:model-value="onPageChange"
            rounded="circle"
          />
        </div>

        <div v-else class="pa-6 text-center text-medium-emphasis">
          <v-icon size="4rem" color="grey-lighten-1">mdi-folder-open-outline</v-icon>
          <h3 class="text-h5 text-grey-lighten-1 mt-4">لا يوجد مشاريع</h3>
          <p class="text-body-1 text-grey-lighten-1">لم يتم العثور على أي مشاريع تطابق معايير البحث</p>
        </div>

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

      <!-- Add/Edit Project Dialog - Profile Form Style (مثل نموذج معلومات المهمة) -->
      <v-dialog v-model="dialog" max-width="900" persistent>
        <v-card class="task-dialog-card profile-form-card">
          <!-- Header Section -->
          <v-card-title class="task-dialog-header profile-form-header">
            <h2 class="profile-form-title">
              {{ isEditing ? 'تعديل المشروع' : 'معلومات المشروع' }}
            </h2>
          </v-card-title>

          <!-- Form Content -->
          <v-card-text class="profile-form-content">
            <p class="profile-form-instruction">
              لإتمام {{ isEditing ? 'تعديل' : 'إضافة' }} المشروع، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
            </p>

            <v-form ref="projectFormRef" v-model="formValid">
              <!-- الصف الأول: اسم، نوع، مكان المشروع -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="projectForm.name"
                    label="اسم المشروع *"
                    variant="outlined"
                    density="comfortable"
                    :rules="[v => !!v || 'اسم المشروع مطلوب']"
                    required
                    hide-details="auto"
                  />
                </v-col>

                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="projectForm.type"
                    label="نوع المشروع"
                    variant="outlined"
                    density="comfortable"
                    hide-details="auto"
                  />
                </v-col>

                <v-col cols="12" md="4">
                  <v-text-field
                    v-model="projectForm.location"
                    label="مكان المشروع *"
                    variant="outlined"
                    density="comfortable"
                    :rules="[v => !!v || 'مكان المشروع مطلوب']"
                    required
                    hide-details="auto"
                  />
                </v-col>
              </v-row>

              <!-- الصف الثاني: التكلفة والمدة والإجمالي -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="4">
                  <v-text-field
                    v-model.number="projectForm.initialCost"
                    label="حد التنبيه (د.ع) *"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    :rules="warningCostRules"
                    required
                    hide-details="auto"
                  />
                </v-col>

                <v-col cols="12" md="4">
                  <v-text-field
                    v-model.number="projectForm.duration"
                    label="مدة المشروع (أيام)"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    hide-details="auto"
                  />
                </v-col>

                <v-col cols="12" md="4">
                  <v-text-field
                    v-model.number="projectForm.totalCost"
                    label="التكلفة الإجمالية (د.ع) *"
                    type="number"
                    variant="outlined"
                    density="comfortable"
                    :rules="totalCostRules"
                    required
                    hide-details="auto"
                  />
                </v-col>
              </v-row>

              <!-- الصف الثالث: تاريخ البدء ورقم الهاتف -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="6">
                  <v-menu
                    v-model="startDateMenu"
                    :close-on-content-click="false"
                    location="bottom"
                  >
                    <template v-slot:activator="{ props }">
                      <v-text-field
                        v-model="formattedStartDate"
                        label="تاريخ البدء"
                        readonly
                        v-bind="props"
                        variant="outlined"
                        density="comfortable"
                        hide-details="auto"
                        prepend-inner-icon="mdi-calendar"
                        dir="ltr"
                        style="text-align: left;"
                      />
                    </template>
                    <v-date-picker
                      v-model="selectedStartDate"
                      @update:model-value="onStartDateSelected"
                      color="primary"
                      show-adjacent-months
                      :first-day-of-week="6"
                      class="rtl-date-picker"
                    />
                  </v-menu>
                </v-col>

                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="projectForm.phone"
                    label="رقم العميل"
                    variant="outlined"
                    density="comfortable"
                    hide-details="auto"
                  />
                </v-col>
              </v-row>

              <!-- الصف الرابع: الوصف -->
              <v-row class="profile-form-row">
                <v-col cols="12">
                  <v-textarea
                    v-model="projectForm.description"
                    label="وصف المشروع"
                    variant="outlined"
                    rows="3"
                    density="comfortable"
                    hide-details="auto"
                  />
                </v-col>
              </v-row>

              <!-- الصف الخامس: الملاحظات -->
              <v-row class="profile-form-row">
                <v-col cols="12">
                  <v-textarea
                    v-model="projectForm.notes"
                    label="ملاحظات"
                    variant="outlined"
                    rows="2"
                    density="comfortable"
                    hide-details="auto"
                  />
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>

          <!-- Footer Actions -->
          <v-card-actions class="profile-form-actions">
            <v-spacer />
            <v-btn
              class="profile-form-cancel-btn"
              variant="outlined"
              @click="closeDialog"
            >
              إلغاء
            </v-btn>
            <v-btn
              class="profile-form-continue-btn"
              variant="elevated"
              :disabled="!formValid"
              @click="saveProject"
            >
              {{ isEditing ? 'تحديث المشروع' : 'حفظ المشروع' }}
              <v-icon class="ms-2">mdi-arrow-left</v-icon>
            </v-btn>
          </v-card-actions>
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
                        {{ selectedProjectDetails.location || 'غير محدد' }}
                      </span>
        </div>
                    <div class="info-item">
                      <span class="info-label">رقم العميل:</span>
                      <span class="info-value">
                        <v-icon size="small" class="me-1">mdi-phone</v-icon>
                        {{ selectedProjectDetails.clientPhone || 'غير محدد' }}
                      </span>
                    </div>
                    <div class="info-item">
                      <span class="info-label">نوع المشروع:</span>
                      <span class="info-value">
                        <v-chip size="small" color="primary">{{ selectedProjectDetails.type || selectedProjectDetails.category || 'غير محدد' }}</v-chip>
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
                      <span class="info-label">الميزانية الإجمالية:</span>
                      <span class="info-value text-success font-weight-bold">
                        {{ formatCurrency(selectedProjectDetails.totalCost) }}
                      </span>
                    </div>
                    <div class="info-item">
                      <span class="info-label">حد التحذير:</span>
                      <span class="info-value text-warning font-weight-bold">
                        {{ formatCurrency(selectedProjectDetails.warningCost) }}
                      </span>
                    </div>
                    <div class="info-item">
                      <span class="info-label">المصروفات الحالية:</span>
                      <span
                        class="info-value font-weight-bold"
                        :class="selectedProjectDetails.currentSpending > selectedProjectDetails.warningCost ? 'text-error' : 'text-success'"
                      >
                        {{ formatCurrency(selectedProjectDetails.currentSpending) }}
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
      <!-- Add/Edit Administrative Expense Dialog - Profile Form Style (مثل نموذج إضافة المهمة) -->
      <v-dialog v-model="expenseDialog" max-width="900" persistent v-show="false">
        <v-card class="task-dialog-card profile-form-card">
          <!-- Header Section -->
          <v-card-title class="task-dialog-header profile-form-header">
            <h2 class="profile-form-title">
              {{ isEditingExpense ? 'تعديل المصروف الإداري' : 'معلومات المصروف الإداري' }}
            </h2>
          </v-card-title>

          <!-- Form Content -->
          <v-card-text class="profile-form-content">
            <p class="profile-form-instruction">
              لإتمام {{ isEditingExpense ? 'تعديل' : 'إضافة' }} المصروف الإداري، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
            </p>

            <v-form ref="expenseForm" v-model="expenseFormValid">
              <!-- الصف الأول: اسم المشروع، التكلفة، مكان العمل -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      اسم المشروع <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.projectName"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'اسم المشروع مطلوب']"
                      required
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      التكلفة (د.ع) <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.cost"
                      type="number"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'التكلفة مطلوبة', v => v > 0 || 'التكلفة يجب أن تكون أكبر من صفر']"
                      required
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      مكان العمل <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.workLocation"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'مكان العمل مطلوب']"
                      required
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الثاني: تاريخ البداية، تاريخ الانتهاء، نوع المصروف -->
              <v-row class="profile-form-row">
                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      تاريخ البداية <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.startDate"
                      type="date"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'تاريخ البداية مطلوب']"
                      required
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      تاريخ الانتهاء <span class="required-star">*</span>
                    </label>
                    <v-text-field
                      v-model="expenseForm.endDate"
                      type="date"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'تاريخ الانتهاء مطلوب']"
                      required
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>

                <v-col cols="12" md="4" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      نوع المصروف <span class="required-star">*</span>
                    </label>
                    <v-select
                      v-model="expenseForm.expenseType"
                      :items="expenseTypes"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'نوع المصروف مطلوب']"
                      required
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
              </v-row>

              <!-- الصف الثالث: الملاحظات -->
              <v-row class="profile-form-row">
                <v-col cols="12" class="profile-form-column">
                  <div class="profile-form-field-wrapper">
                    <label class="profile-form-label">
                      الملاحظات
                    </label>
                    <v-textarea
                      v-model="expenseForm.notes"
                      variant="outlined"
                      rows="3"
                      density="comfortable"
                      hide-details="auto"
                      class="profile-form-input"
                    />
                  </div>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>

          <!-- Dialog Actions -->
          <v-card-actions class="profile-form-actions">
            <v-spacer />
            <v-btn
              class="profile-form-cancel-btn"
              variant="outlined"
              @click="closeExpenseDialog"
            >
              إلغاء
            </v-btn>
            <v-btn
              class="profile-form-continue-btn"
              variant="elevated"
              :disabled="!expenseFormValid"
              @click="saveExpense"
            >
              {{ isEditingExpense ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Snackbar for error/success messages -->
      <v-snackbar
        v-model="showSuccessMessage"
        :timeout="5000"
        color="error"
        location="top"
      >
        {{ successMessage }}
        <template v-slot:actions>
          <v-btn
            variant="text"
            @click="showSuccessMessage = false"
          >
            إغلاق
          </v-btn>
        </template>
      </v-snackbar>
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { listProjects, getProjectStats, getProjectWorkdays, getProject, createProject, updateProject } from '@/api/projects'
import { DEFAULT_LIMIT } from '@/constants/pagination'

// عنوان الصفحة
document.title = 'إدارة المشاريع الهندسية - نظام إدارة المشاريع'

const router = useRouter()

// متغيرات الحالة الأساسية
const loading = ref(false)

// Pagination state
const currentPage = ref(1)
const totalItems = ref(0)
const totalPages = ref(0)

const dialog = ref(false)
const deleteDialog = ref(false)
const detailsDialog = ref(false)
const formValid = ref(false)
const projectFormRef = ref(null)
const isEditing = ref(false)
const searchQuery = ref('')
const selectedCategory = ref('')
const selectedStatus = ref('')
const selectedProject = ref(null)
const selectedProjectDetails = ref(null)
const detailsLoading = ref(false)
const savingProject = ref(false)
const showSuccessMessage = ref(false)
const successMessage = ref('')
const workdaysLoading = ref(false)
const workdaysError = ref('')

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


// نموذج البيانات
const projectForm = ref({
  name: '',
  type: '',
  location: '',
  initialCost: 0,
  totalCost: 0,
  duration: 0,
  startDate: '',
  phone: '',
  user: '',
  status: 'pending',
  priority: 'medium',
  progress: 0,
  description: '',
  notes: '',
  category: ''
})

// Validation rules for cost fields (computed to react to changes)
const warningCostRules = computed(() => [
  v => Number(v) > 0 || 'حد التنبيه يجب أن يكون أكبر من صفر',
  v => Number(v) <= Number(projectForm.value.totalCost) || 'حد التنبيه يجب أن يكون أقل من أو يساوي التكلفة الإجمالية'
])

const totalCostRules = computed(() => [
  v => Number(v) > 0 || 'التكلفة يجب أن تكون أكبر من صفر',
  v => Number(v) >= Number(projectForm.value.initialCost) || 'التكلفة الإجمالية يجب أن تكون أكبر من أو تساوي حد التنبيه'
])

// Watch cost fields to trigger form revalidation for cross-field validation
watch(() => projectForm.value.initialCost, () => {
  if (projectFormRef.value) {
    projectFormRef.value.validate()
  }
})

watch(() => projectForm.value.totalCost, () => {
  if (projectFormRef.value) {
    projectFormRef.value.validate()
  }
})

// Date picker state for start date
const startDateMenu = ref(false)
const selectedStartDate = ref(null)

// Computed property for formatted start date display
const formattedStartDate = computed(() => {
  if (!projectForm.value.startDate) return ''
  const date = new Date(projectForm.value.startDate)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}/${month}/${day}`
})

// Handle start date selection from picker
const onStartDateSelected = (date) => {
  if (date) {
    const d = new Date(date)
    projectForm.value.startDate = d.toISOString().split('T')[0]
  }
  startDateMenu.value = false
}

// البيانات الأساسية (يتم جلبها من الـ API)
const projects = ref([])
const projectStats = ref({})

console.log('Projects data loaded:', projects.value.length, 'projects')

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
const totalProjects = computed(() => projectStats.value.total ?? projects.value.length)
const activeProjects = computed(() => {
  if (projectStats.value.inProgress != null) return projectStats.value.inProgress
  return projects.value.filter(p => p.status === 'active' || p.status === 'in_progress').length
})
const pendingProjects = computed(() => {
  if (projectStats.value.pending != null) return projectStats.value.pending
  return projects.value.filter(p => p.status === 'pending').length
})
const averageProgress = computed(() => {
  if (projectStats.value.averageProgress != null) return projectStats.value.averageProgress
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
  if (projectStats.value.totalBudget != null) return projectStats.value.totalBudget
  return projects.value.reduce((sum, project) => sum + (project.budget ?? project.initialCost ?? 0), 0)
})

const displayTotalProjects = computed(() => (loading.value ? '...' : totalProjects.value ?? 0))
const displayActiveProjects = computed(() => (loading.value ? '...' : activeProjects.value ?? 0))
const displayPendingProjects = computed(() => (loading.value ? '...' : pendingProjects.value ?? 0))
const displayAverageProgress = computed(() => (loading.value ? '...' : (averageProgress.value ?? 0) + '%'))
const displayTotalBudget = computed(() => {
  if (loading.value) return '...'
  if (totalBudget.value != null) return formatCurrency(totalBudget.value)
  return '-'
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
  const workingDays = selectedProjectDetails.value?.workingDays || getDefaultWorkingDays()
  return workingDays.reduce((total, day) => {
    const hours = parseInt(day.hours) || 0
    return day.isActive ? total + hours : total
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
      ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
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
    console.log('تم تصدير بيانات أيام العمل بنجاح')
    
  } catch (error) {
    console.error('خطأ في تصدير البيانات:', error)
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

async function loadWorkdays(projectId) {
  if (!projectId) return
  workdaysLoading.value = true
  workdaysError.value = ''
  try {
    const data = await getProjectWorkdays(projectId)
    if (Array.isArray(data) && data.length) {
      selectedProjectDetails.value = {
        ...selectedProjectDetails.value,
        workingDays: data.map(day => ({
          name: day.dayName || day.name || '',
          hours: day.hours ? `${day.hours}` : '0',
          isActive: day.isActive ?? true,
        })),
      }
    }
  } catch (err) {
    console.error('فشل تحميل أيام العمل', err)
    workdaysError.value = 'فشل تحميل أيام العمل من الخادم'
  } finally {
    workdaysLoading.value = false
  }
}

// دوال إدارة البيانات
const openAddProjectDialog = () => {
  isEditing.value = false
  // Set today's date as default
  const today = new Date().toISOString().split('T')[0]
  projectForm.value = {
    name: '',
    type: '',
    location: '',
    initialCost: 0,
    totalCost: 0,
    duration: 0,
    startDate: today,
    phone: '',
    user: '',
    status: 'pending',
    priority: 'medium',
    progress: 0,
    description: '',
    notes: '',
    category: ''
  }
  // Initialize selected date for picker
  selectedStartDate.value = new Date()
  dialog.value = true
}

const viewProject = (project) => {
  console.log('View project:', project)
}

const viewProjectDetails = (project) => {
  if (!project?.id) return
  router.push({ path: '/work-days', query: { projectId: project.id } })
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
  // Map API fields to form fields
  projectForm.value = {
    name: project.name || '',
    type: project.type || '',
    location: project.location || '',
    initialCost: project.warningCost || 0,
    totalCost: project.totalCost || 0,
    duration: project.duration || 0,
    startDate: project.startDate ? project.startDate.split('T')[0] : '',
    phone: project.clientPhone || '',
    user: '',
    status: project.status || 'pending',
    priority: 'medium',
    progress: project.progressPercentage || 0,
    description: project.description || '',
    notes: project.notes || '',
    category: ''
  }
  // Initialize selected date for picker with the project's start date
  selectedStartDate.value = project.startDate ? new Date(project.startDate) : null
  selectedProject.value = project
  dialog.value = true
}

const deleteProject = (project) => {
  selectedProject.value = project
  deleteDialog.value = true
}

const saveProject = async () => {
  if (!formValid.value) return
  savingProject.value = true
  try {
    // Raw values - no calculations
    const payload = {
      name: projectForm.value.name,
      type: projectForm.value.type || '',
      description: projectForm.value.description || '',
      clientPhone: projectForm.value.phone || '',
      location: projectForm.value.location || '',
      startDate: projectForm.value.startDate ? new Date(projectForm.value.startDate).toISOString() : null,
      duration: projectForm.value.duration ? Number(projectForm.value.duration) : null,
      warningCost: Number(projectForm.value.initialCost || 0),
      totalCost: Number(projectForm.value.totalCost || 0),
      notes: projectForm.value.notes || '',
    }

    if (isEditing.value && selectedProject.value?.id) {
      await updateProject(selectedProject.value.id, payload)
    } else {
      await createProject(payload)
    }

    await loadProjectsFromApi()
    closeDialog()
  } catch (err) {
    console.error('فشل حفظ المشروع', err)

    // Handle validation errors from backend
    let errorMessage = 'فشل حفظ المشروع'
    if (err?.data?.errors && typeof err.data.errors === 'object') {
      const fieldNames = {
        name: 'اسم المشروع',
        warningCost: 'تكلفة التحذير',
        totalCost: 'التكلفة الإجمالية',
        location: 'الموقع',
        startDate: 'تاريخ البدء',
        duration: 'المدة',
        clientPhone: 'رقم الهاتف',
        description: 'الوصف',
      }
      const errorMessages = Object.entries(err.data.errors).map(([field, msg]) => {
        const fieldName = fieldNames[field] || field
        return `${fieldName} مطلوب`
      })
      errorMessage = errorMessages.join('، ')
    } else if (err?.message) {
      errorMessage = err.message
    }

    successMessage.value = errorMessage
    showSuccessMessage.value = true
  } finally {
    savingProject.value = false
  }
}

const closeDialog = () => {
  dialog.value = false
  isEditing.value = false
  selectedProject.value = null
  projectForm.value = {
    name: '',
    type: '',
    location: '',
    initialCost: 0,
    totalCost: 0,
    duration: 0,
    startDate: '',
    phone: '',
    user: '',
    status: 'pending',
    priority: 'medium',
    progress: 0,
    description: '',
    notes: '',
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
  console.log('عرض عضو الفريق:', member)
  // TODO: إضافة نافذة عرض التفاصيل
}

const editTeamMember = (member) => {
  console.log('تعديل عضو الفريق:', member)
  // TODO: إضافة نافذة التعديل
}

const deleteTeamMember = (member) => {
  console.log('حذف عضو الفريق:', member)
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

const statusMap = {
  pending: 'pending',
  in_progress: 'active',
  completed: 'completed',
}

function mapStatus(status) {
  return statusMap[status] || status || 'pending'
}

async function loadProjectsFromApi(page = 1) {
  loading.value = true
  try {
    const [list, stats] = await Promise.all([
      listProjects({ page, limit: DEFAULT_LIMIT }),
      getProjectStats({ period: 'all' })
    ])
    const listData = Array.isArray(list)
      ? list
      : (Array.isArray(list?.data) ? list.data : [])

    const mapped = listData.map(p => ({
      id: p.id,
      name: p.name,
      description: p.description || '',
      location: p.location || '',
      status: mapStatus(p.status),
      totalCost: p.totalCost ?? 0,
      warningCost: p.warningCost ?? 0,
      currentSpending: p.currentSpending ?? 0,
      initialCost: p.totalCost ?? p.initialCost ?? 0,
      budget: p.totalCost ?? p.initialCost ?? 0,
      criticalCost: p.warningCost ?? p.criticalCost ?? 0,
      startDate: p.startDate,
      duration: p.duration ?? null,
      clientPhone: p.clientPhone || '',
      type: p.type || '',
      user: p.createdByName || p.createdBy || p.manager || '',
      manager: p.createdByName || p.createdBy || p.manager || '',
      progress: p.progressPercentage ?? p.progress ?? 0,
      category: p.type || '',
      notes: p.notes || '',
      createdBy: p.createdBy,
      createdAt: p.createdAt,
    }))
    if (mapped.length) {
      projects.value = mapped
    }
    projectStats.value = stats?.data || stats || {}

    // Update pagination
    totalItems.value = list.total || 0
    totalPages.value = list.totalPages || 0
    currentPage.value = list.page || page
  } catch (err) {
    console.error('فشل تحميل المشاريع من الخادم', err)
  } finally {
    loading.value = false
  }
}

// Handle page change
function onPageChange(page) {
  currentPage.value = page
  loadProjectsFromApi(page)
}

onMounted(() => {
  loadProjectsFromApi()
})
</script>

<style scoped>
/* ========================================
   صفحة إدارة المشاريع - CSS منظم
   Project Management Page - Organized CSS
   ======================================== */

/* ========================================
   1. الصفحة العامة
   ======================================== */

.data-page {
  background: #ffffff !important;
  color: var(--text-dark);
  min-height: 100vh;
  padding: 0 !important;
  padding-top: 0 !important;
  margin: 0 !important;
  margin-top: 0 !important;
  width: 100%;
  overflow-x: hidden;
}

.data-page :deep(.v-container) {
  padding: 0 1rem !important;
  padding-top: 0 !important;
  overflow-x: hidden;
  max-width: 100% !important;
}

.fullscreen-content {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  text-align: center;
  padding: 0 20px;
  padding-top: 0 !important;
  margin-top: 0 !important;
}

/* ========================================
   2. رأس الصفحة
   ======================================== */

/* صف الإحصائيات */
.stats-row {
  width: 100%;
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
  padding: 8px !important;
  box-sizing: border-box !important;
}

.stats-row .stat-card {
  width: 100% !important;
  min-height: 140px !important;
  margin: 0 !important;
  box-sizing: border-box !important;
}

.stats-row.full-width {
  max-width: 100%;
}

/* ========================================
   4. رأس الصفحة المحسن
   ======================================== */

.engineers-header-card {
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.7) 0%, rgba(21, 101, 192, 0.7) 100%);
  border-radius: 0;
  width: 100vw;
  max-width: 100vw;
  margin-left: calc(-50vw + 50%);
  margin-right: calc(-50vw + 50%);
  margin-top: 0;
  margin-bottom: 1.5rem;
  box-shadow: 0 8px 32px rgba(25, 118, 210, 0.3);
  position: relative;
  overflow: hidden;
  border: none;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: slideInFromTop 0.8s ease-out;
  z-index: 1000 !important;
  isolation: isolate;
  backdrop-filter: blur(10px);
}

.engineers-header-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.15), transparent);
  animation: sweep 3s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.08) 50%, transparent 70%);
  animation: diagonalShimmer 5s ease-in-out infinite;
  z-index: 1;
}

.engineers-header-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(25, 118, 210, 0.4);
}

.engineers-header-card:hover::before {
  animation: sweep 2s ease-in-out infinite;
}

.engineers-header-card:hover::after {
  animation: diagonalShimmer 3s ease-in-out infinite;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(90deg, #ffffff 0%, #e3f2fd 50%, #bbdefb 100%);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.3);
  animation: gradientFlow 4s ease-in-out infinite;
  z-index: 2;
}

.header-content {
  padding: 12px 16px !important;
  padding-top: 0 !important;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  min-height: auto !important;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.1) 0%, rgba(21, 101, 192, 0.05) 100%);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 1001 !important;
  animation: fadeInUp 0.8s ease-out 0.2s both;
  max-width: calc(100vw - 320px);
  margin: 0 auto;
  margin-top: 0 !important;
}

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

.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.8rem;
  position: relative;
}

.engineer-emoji {
  position: relative;
  animation: slideInFromRight 0.8s ease-out 0.4s both;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 0;
  transition: transform 0.3s ease;
  width: 56px;
  height: 56px;
  min-width: 56px;
  min-height: 56px;
  flex-shrink: 0;
}

.engineer-emoji:hover {
  transform: scale(1.05);
}

.engineer-emoji .v-icon {
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
  transition: all 0.3s ease;
  color: white !important;
  position: relative;
  font-size: 48px !important;
  width: 48px !important;
  height: 48px !important;
  min-width: 48px !important;
  min-height: 48px !important;
}

.engineer-emoji :deep(.v-icon) {
  font-size: 48px !important;
  width: 48px !important;
  height: 48px !important;
  min-width: 48px !important;
  min-height: 48px !important;
}

.engineer-emoji :deep(svg) {
  width: 48px !important;
  height: 48px !important;
  font-size: 48px !important;
}

.engineer-emoji .v-icon:first-child {
  animation: iconBounce 3s ease-in-out infinite 1.5s;
}

/* Old page-header styles removed - using engineers-header-card instead */

.title-icon {
  font-size: 3rem;
  animation: bounce 2s infinite;
  filter: drop-shadow(0 4px 8px rgba(26, 35, 126, 0.3));
  transition: all 0.3s ease;
}

.title-icon:hover {
  transform: scale(1.2) rotate(10deg);
  filter: drop-shadow(0 6px 12px rgba(26, 35, 126, 0.5));
}

.title-text {
  background: linear-gradient(135deg, #1a237e 0%, #3949ab 30%, #5c6bc0 60%, #7986cb 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  font-weight: 800;
  position: relative;
  z-index: 2;
}

.title-decoration {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 100%;
  background: linear-gradient(45deg, transparent 30%, rgba(26, 35, 126, 0.1) 50%, transparent 70%);
  animation: shimmer 4s infinite;
  z-index: 1;
}

.page-subtitle {
  color: #283593 !important;
  font-size: 0.85rem !important;
  margin-bottom: 1rem;
  text-shadow: 0 1px 4px rgba(40, 53, 147, 0.2);
  opacity: 0.9;
  font-weight: 500;
  line-height: 1.4;
}

.page-badge {
  margin-top: 1rem;
  display: flex;
  justify-content: center;
  animation: fadeInUp 0.8s ease-out 0.3s both;
}

.page-badge .v-chip {
  font-weight: 600;
  font-size: 0.9rem;
  padding: 0.5rem 1rem;
  box-shadow: 0 4px 12px rgba(25, 118, 210, 0.3);
  transition: all 0.3s ease;
}

.page-badge .v-chip:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(25, 118, 210, 0.4);
}

.page-main-title {
  margin: 2rem 0;
  text-align: center;
  animation: fadeInUp 0.8s ease-out 0.5s both;
  padding: 1.5rem 2rem;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.03) 0%, rgba(57, 73, 171, 0.05) 50%, rgba(92, 107, 192, 0.03) 100%);
  border-radius: 20px;
  border: 1px solid rgba(26, 35, 126, 0.1);
  backdrop-filter: blur(15px);
  box-shadow: 0 8px 32px rgba(26, 35, 126, 0.1);
  position: relative;
  overflow: hidden;
}

.page-main-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  animation: shimmer 3s infinite;
}

.page-main-title:hover {
  transform: translateY(-2px);
  box-shadow: 0 12px 40px rgba(26, 35, 126, 0.15);
  border-color: rgba(26, 35, 126, 0.15);
}

.main-title {
  color: white !important;
  font-size: 1.2rem !important;
  font-weight: bold !important;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 0.75rem !important;
  margin: 0;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
  -moz-osx-font-smoothing: grayscale;
}

.main-subtitle {
  color: #283593 !important;
  font-size: 1.3rem !important;
  margin-bottom: 0;
  text-shadow: 0 2px 8px rgba(40, 53, 147, 0.3);
  opacity: 0.95;
  font-weight: 600;
  line-height: 1.6;
  font-style: italic;
  background: linear-gradient(135deg, #283593 0%, #3949ab 50%, #5c6bc0 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  padding: 0.3rem 0.8rem;
  border-radius: 10px;
  background-color: rgba(226, 14, 84, 0.03);
  backdrop-filter: blur(5px);
  border: 1px solid rgba(40, 53, 147, 0.08);
  box-shadow: 0 4px 16px rgba(40, 53, 147, 0.1);
  transition: all 0.3s ease;
  display: inline-block;
  margin-top: 0.5rem;
}

.main-subtitle:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(40, 53, 147, 0.15);
  border-color: rgba(40, 53, 147, 0.12);
}

.page-icon {
  font-size: 3.5rem !important;
  margin-bottom: 0.8rem;
  display: block;
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.15));
}

/* ========================================
   5. بطاقات الإحصائيات المحسنة
   ======================================== */

.stats-cards-container {
  display: flex;
  gap: 0.75rem;
  flex-wrap: nowrap;
  overflow: hidden;
  padding: 0.5rem 0;
  width: 100%;
  justify-content: space-between;
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
  background: #ffffff;
  border: 2px solid #3b82f6;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15);
}

.stat-card-success {
  background: #ffffff;
  border: 2px solid #10b981;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.15);
}

.stat-card-warning {
  background: #ffffff;
  border: 2px solid #f59e0b;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.15);
}

.stat-card-info {
  background: #ffffff;
  border: 2px solid #06b6d4;
  box-shadow: 0 4px 12px rgba(6, 182, 212, 0.15);
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
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
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
  margin: 0;
  padding: 0;
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

/* ========================================
   6. شريط الفلترة
   ======================================== */

.search-card,
.filters-card {
  background: #ffffff !important;
  backdrop-filter: blur(15px) !important;
  border-radius: 20px !important;
  border: 1px solid #e2e8f0 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05) !important;
  overflow: hidden !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.search-card:hover,
.filters-card:hover {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1) !important;
  border-color: #cbd5e1 !important;
}

.filters-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: white !important;
  border-radius: 20px 20px 0 0 !important;
  padding: 12px 16px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.3) !important;
  display: flex !important;
  align-items: center !important;
  gap: 8px !important;
  min-height: auto !important;
}

.search-card .v-field__input,
.search-card .v-field__input input,
.filters-card .v-field__input,
.filters-card .v-field__input input,
.search-field .v-field__input,
.search-field .v-field__input input {
  background: #ffffff !important;
  border-radius: 8px !important;
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
  text-shadow: none !important;
  letter-spacing: 0.2px !important;
}

.search-card .v-field__outline,
.filters-card .v-field__outline,
.search-field .v-field__outline {
  border-color: #cbd5e1 !important;
  border-width: 1px !important;
}

.search-card .v-field--focused .v-field__outline,
.filters-card .v-field--focused .v-field__outline,
.search-field .v-field--focused .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.search-card .v-label,
.search-card .v-field-label,
.filters-card .v-label,
.filters-card .v-field-label {
  color: #475569 !important;
  font-weight: 600 !important;
  font-size: 0.8rem !important;
}

.search-card .v-field-label--active,
.search-card .v-field-label--floating,
.filters-card .v-field-label--active,
.filters-card .v-field-label--floating {
  color: #3b82f6 !important;
  font-weight: 700 !important;
  font-size: 0.8rem !important;
}

.filter-field .v-field__input,
.filter-field .v-field__input input,
.filter-field .v-field__input select {
  background: #ffffff !important;
  border-radius: 8px !important;
  color: #1e293b !important;
  font-weight: 600 !important;
  font-size: 0.875rem !important;
}

.filter-field .v-field__outline {
  border-color: #cbd5e1 !important;
  border-width: 1px !important;
}

.filter-field .v-field--focused .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15) !important;
}

.filter-field .v-label,
.filter-field .v-field-label {
  color: #475569 !important;
  font-weight: 600 !important;
  font-size: 0.8rem !important;
}

.filter-field .v-field-label--active,
.filter-field .v-field-label--floating {
  color: #3b82f6 !important;
  font-weight: 700 !important;
  font-size: 0.8rem !important;
}

.search-btn {
  height: 48px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 8px !important;
}

.add-expense-btn {
  height: 48px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 8px !important;
}

/* ========================================
   7. بطاقة الجدول
   ======================================== */
/* تم نقل الأنماط الأساسية إلى common-components.css */

.data-table-card {
  border-radius: 20px !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.data-table-card:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15) !important;
}

.data-table-card .v-card-title {
  padding: 10px 16px !important;
  min-height: auto !important;
}

.data-table-card .v-card-title .v-icon {
  font-size: 20px !important;
  width: 20px !important;
  height: 20px !important;
}

.data-table-card .v-card-title .text-h4 {
  font-size: 1.25rem !important;
  line-height: 1.4 !important;
}

.data-table-card .v-card-title .v-btn {
  padding: 6px 16px !important;
  font-size: 0.875rem !important;
  min-height: 36px !important;
  height: 36px !important;
}

.data-table-card .v-card-title .v-btn .v-icon {
  font-size: 18px !important;
  width: 18px !important;
  height: 18px !important;
}

.data-table-card .v-card-title .v-chip {
  font-size: 0.75rem !important;
  height: 24px !important;
  padding: 0 8px !important;
}

.data-table-card .v-card-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: var(--gradient-primary);
}

.projects-grid-row {
  margin-top: 1.5rem !important;
  padding: 0 1rem !important;
}

.data-table-card .v-data-table {
  background: transparent !important;
}

.data-table-card .v-data-table-header {
  background: #f8fafc !important;
  font-weight: 600 !important;
}

.data-table-card tbody tr:hover {
  background: #f1f5f9 !important;
  transform: scale(1.01);
  transition: all 0.2s ease;
}

.data-table-card tbody tr {
  transition: all 0.2s ease;
}

.centered-table {
  width: 100%;
  margin: 0;
  border-radius: 20px !important;
  overflow: hidden;
}

/* ========================================
   7. شريط البحث والفلترة
   ======================================== */

.search-field,
.filter-field {
  border-radius: 12px !important;
  transition: all 0.3s ease !important;
}

.search-field .v-field,
.filter-field .v-field {
  border-radius: 12px !important;
  transition: all 0.3s ease !important;
}

.search-field .v-field:focus-within,
.filter-field .v-field:focus-within {
  border-color: var(--primary-color) !important;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.1) !important;
  transform: translateY(-1px);
}

.search-field .v-field__input {
  padding: 12px 16px !important;
  font-size: 0.95rem !important;
}

.filter-field .v-field__input {
  padding: 12px 16px !important;
  font-size: 0.95rem !important;
}

.search-field .v-field__prepend-inner {
  margin-inline-start: 16px !important;
  color: var(--primary-color) !important;
}

.filter-field .v-field__prepend-inner {
  margin-inline-start: 16px !important;
  color: var(--primary-color) !important;
}

.search-field .v-field__clearable {
  margin-inline-end: 16px !important;
}

.filter-field .v-field__clearable {
  margin-inline-end: 16px !important;
}

/* ========================================
   8. تفاصيل المشروع Dialog
   ======================================== */
.details-card {
  border-radius: 20px !important;
  overflow: hidden;
}

.details-header {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  color: rgb(229, 222, 222) !important;
  padding: 1.5rem !important;
  border-radius: 20px 20px 0 0 !important;
}

.details-header .v-icon {
  color: white !important;
}

.close-btn {
  color: white !important;
}

.close-btn:hover {
  background-color: rgba(255, 255, 255, 0.1) !important;
}

.details-content {
  padding: 2rem !important;
  background: var(--bg-light);
}

.info-card {
  border-radius: 16px !important;
  transition: all 0.3s ease;
  border: 2px solid #e5e7eb !important;
  background: #ffffff !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08) !important;
}

.info-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15) !important;
  border-color: #cbd5e1 !important;
}

.info-card-title {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  color: #1e293b !important;
  font-size: 1.1rem !important;
  font-weight: 700 !important;
  padding: 1rem 1.5rem !important;
  border-bottom: 2px solid #e5e7eb !important;
}

.info-card-title .v-icon {
  color: var(--primary-color) !important;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 0;
  border-bottom: 1px solid #e5e7eb !important;
  background: #ffffff !important;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  font-weight: 700 !important;
  color: #374151 !important;
  font-size: 1rem !important;
  min-width: 140px;
  text-align: right;
}

.info-value {
  font-weight: 600 !important;
  color: #1e293b !important;
  font-size: 1rem !important;
  text-align: left;
  padding: 4px 12px;
  background: #f9fafb;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
}

.description-text {
  background: #f8fafc;
  padding: 1.5rem;
  border-radius: 12px;
  border-left: 4px solid var(--primary-color);
  font-size: 1rem;
  line-height: 1.6;
  color: var(--text-dark);
  min-height: 100px;
}

/* أيام العمل */
.working-days-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.working-day-item {
  display: flex;
  align-items: center;
  padding: 0.75rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  background: #f8fafc;
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.working-day-item.active {
  background: linear-gradient(135deg, #e8f5e8 0%, #f0f9f0 100%);
  border-color: rgba(34, 197, 94, 0.2);
}

.working-day-item:hover {
  transform: translateX(4px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.day-name {
  font-weight: 600;
  color: var(--text-dark);
  flex: 1;
}

.day-hours {
  font-weight: 500;
  color: var(--text-dark-muted);
  font-size: 0.9rem;
}

.total-hours {
  margin-top: 1rem;
}

/* إجراءات المشروع */
.actions-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.action-item {
  padding: 1rem;
  border-radius: 12px;
  background: #f8fafc;
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.action-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: rgba(59, 130, 246, 0.2);
}

.action-header {
  display: flex;
  align-items: center;
  margin-bottom: 0.5rem;
}

.action-title {
  font-weight: 600;
  color: var(--text-dark);
  flex: 1;
  margin: 0 0.5rem;
}

.action-details {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.action-date {
  display: flex;
  align-items: center;
  color: var(--text-dark-muted);
  font-size: 0.85rem;
}
  
.details-actions {
  padding: 2rem 2.5rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-top: 1px solid rgba(0, 0, 0, 0.08);
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 1rem;
  position: relative;
}

.details-actions::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent 0%, rgba(59, 130, 246, 0.3) 50%, transparent 100%);
}

.close-details-btn,
.edit-details-btn {
  border-radius: 16px !important;
  font-weight: 700 !important;
  text-transform: none !important;
  padding: 0.875rem 2rem !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  min-width: 140px !important;
  height: 52px !important;
  position: relative;
  overflow: hidden;
}

.close-details-btn {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 2px solid var(--primary-color) !important;
  color: var(--primary-color) !important;
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.15) !important;
}

.close-details-btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 8px 25px rgba(59, 130, 246, 0.25) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f1f5f9 100%) !important;
  border-color: var(--secondary-color) !important;
  color: var(--secondary-color) !important;
}

.close-details-btn .v-icon {
  transition: all 0.3s ease !important;
}

.close-details-btn:hover .v-icon {
  transform: scale(1.1) rotate(90deg);
}

.edit-details-btn {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 50%, #b45309 100%) !important;
  color: white !important;
  border: none !important;
  box-shadow: 0 6px 20px rgba(245, 158, 11, 0.3) !important;
  position: relative;
}

.edit-details-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2) 0%, transparent 50%);
  border-radius: 16px;
  z-index: 1;
}

.edit-details-btn .v-icon,
.edit-details-btn span {
  position: relative;
  z-index: 2;
}

.edit-details-btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 12px 35px rgba(245, 158, 11, 0.4) !important;
  background: linear-gradient(135deg, #f59e0b 0%, #ea580c 50%, #dc2626 100%) !important;
}

.edit-details-btn:hover .v-icon {
  transform: scale(1.15) rotate(-5deg);
}

.edit-details-btn:active {
  transform: translateY(-1px) scale(1.01);
  box-shadow: 0 4px 15px rgba(245, 158, 11, 0.3) !important;
}

.close-details-btn:active {
  transform: translateY(-1px) scale(1.01);
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.2) !important;
}

/* تأثيرات إضافية للأزرار */
.close-details-btn::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, transparent 50%);
  border-radius: 16px;
  opacity: 0;
  transition: opacity 0.3s ease;
  z-index: 1;
}

.close-details-btn:hover::after {
  opacity: 1;
}

.edit-details-btn::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.3) 0%, transparent 70%);
  border-radius: 16px;
  opacity: 0;
  transition: opacity 0.3s ease;
  z-index: 3;
}

.edit-details-btn:hover::after {
  opacity: 1;
}

/* Animations for header */
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

/* تأثير النقر */
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

.close-details-btn:active,
.edit-details-btn:active {
  animation: buttonClick 0.2s ease;
}

/* زر العودة للصفحة الرئيسية */
.back-details-btn {
  border-radius: 16px !important;
  font-weight: 700 !important;
  text-transform: none !important;
  padding: 0.875rem 2rem !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  min-width: 180px !important;
  height: 52px !important;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 2px solid var(--secondary-color) !important;
  color: var(--secondary-color) !important;
  box-shadow: 0 4px 15px rgba(107, 114, 128, 0.15) !important;
}

.back-details-btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 8px 25px rgba(107, 114, 128, 0.25) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f1f5f9 100%) !important;
  border-color: var(--primary-color) !important;
  color: var(--primary-color) !important;
}

.back-details-btn:hover .v-icon {
  transform: scale(1.1) translateX(3px);
}

.back-details-btn:active {
  transform: translateY(-1px) scale(1.01);
  box-shadow: 0 4px 15px rgba(107, 114, 128, 0.2) !important;
  animation: buttonClick 0.2s ease;
}

/* أزرار الإجراءات في الجدول */
.action-btn {
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
  min-width: 40px !important;
  height: 40px !important;
}

.action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

.action-btn .v-icon {
  transition: all 0.3s ease !important;
}

.action-btn:hover .v-icon {
  transform: scale(1.1);
}

/* زر الإضافة */
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

@keyframes buttonPulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.02);
    box-shadow: 
      0 10px 28px rgba(79, 70, 229, 0.6),
      0 5px 14px rgba(99, 102, 241, 0.5),
      0 3px 8px rgba(129, 140, 248, 0.4),
      inset 0 2px 0 rgba(255, 255, 255, 0.4),
      inset 0 -2px 0 rgba(0, 0, 0, 0.15);
  }
}

@keyframes buttonFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-3px);
  }
}

@keyframes buttonGradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px) scale(1);
  }
  50% {
    transform: translateY(-6px) scale(1.05);
  }
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
  transition: transform 0.3s ease !important;
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

.add-button:active {
  transform: translateY(-1px) scale(1.02) !important;
  box-shadow: 
    0 4px 12px rgba(30, 64, 175, 0.4),
    0 2px 6px rgba(37, 99, 235, 0.3),
    0 0 25px rgba(59, 130, 246, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  animation: buttonClick 0.3s ease !important;
}

.add-button:hover :deep(.v-icon) {
  transform: rotate(90deg) scale(1.1) !important;
}

/* النوافذ المنبثقة */
.v-dialog .v-card {
  border-radius: 20px !important;
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
  color: var(--primary-color) !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.text-success {
  color: var(--success-color) !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.text-info {
  color: var(--info-color) !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.text-warning {
  color: var(--warning-color) !important;
  font-weight: 600 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}


/* تصميم متجاوب */
@media (max-width: 1200px) {
  .data-page {
    padding: 1rem 0.5rem;
  }
  
  .fullscreen-content {
    gap: 1.5rem;
  }
  
  .stats-row {
    gap: 1.5rem;
    justify-content: center;
    flex-wrap: nowrap;
    padding: 0 1.5rem;
  }
  
  .stat-card {
    height: 180px;
    flex: 0 0 200px;
    width: 200px;
  }
  
  .stat-card h3 {
    font-size: 1.9rem !important;
  }
  
  .stat-card .stat-icon .v-icon {
    font-size: 56px !important;
    width: 56px !important;
    height: 56px !important;
  }
  
  .page-title {
    font-size: 2rem !important;
    padding: 0.8rem 1.5rem;
    gap: 0.8rem;
  }
  
  .title-icon {
    font-size: 2.5rem;
  }
  
  .title-text {
    font-size: 2rem;
  }
  
  .page-subtitle {
    font-size: 1.1rem !important;
  }
  
  .page-header {
    padding: 1.5rem 1rem;
    max-height: 160px;
  }
  
  .page-icon {
    font-size: 3rem !important;
  }
}

@media (max-width: 768px) {
  .data-page {
    padding: 0.5rem;
  }
  
  .fullscreen-content {
    gap: 1rem;
  }
  
  .main-title {
    font-size: 1.8rem !important;
    padding: 0.4rem 0.8rem;
  }
  
  .main-subtitle {
    font-size: 1.1rem !important;
    padding: 0.2rem 0.6rem;
  }
  
  .page-main-title {
    margin: 1.5rem 0;
    padding: 1rem 1.5rem;
  }
  
  .stats-row {
    gap: 1rem;
    flex-wrap: nowrap;
    justify-content: center;
    padding: 0 1rem;
  }
  
  .stat-card {
    height: 160px;
    flex: 0 0 180px;
    width: 180px;
  }
  
  .stat-card h3 {
    font-size: 1.7rem !important;
  }
  
  .stat-card .stat-icon .v-icon {
    font-size: 52px !important;
    width: 52px !important;
    height: 52px !important;
  }
  
  .page-title {
    font-size: 2.2rem !important;
  }
  
  .page-subtitle {
    font-size: 1rem !important;
  }
  
  .page-header {
    padding: 1.5rem 1rem;
  }
  
  .page-icon {
    font-size: 3.5rem !important;
  }
  
  .add-button {
    padding: 10px 20px !important;
    font-size: 1rem !important;
  }
  
}

/* شاشات صغيرة جداً */
@media (max-width: 480px) {
  .stats-row {
    gap: 0.8rem;
    flex-wrap: nowrap;
    justify-content: center;
    padding: 0 0.5rem;
  }
  
  .stat-card {
    height: 140px;
    flex: 0 0 160px;
    width: 160px;
  }
  
  .stat-card h3 {
    font-size: 1.5rem !important;
  }
  
  .stat-card .stat-icon .v-icon {
    font-size: 44px !important;
    width: 44px !important;
    height: 44px !important;
  }
  
  .stat-card p {
    font-size: 0.85rem !important;
  }
  
  .page-title {
    font-size: 1.8rem !important;
  }
  
  .page-subtitle {
    font-size: 0.9rem !important;
  }
  
  .page-header {
    padding: 1rem;
  }
  
  .page-icon {
    font-size: 3rem !important;
  }
}

/* تحسينات للشاشات الكبيرة */
@media (min-width: 1400px) {
  .data-page {
    padding: 2rem;
  }
  
  .fullscreen-content {
    gap: 3rem;
  }
  
  .stats-row {
    gap: 2.5rem;
    padding: 0 3rem;
  }
  
  .stat-card {
    height: 220px;
    flex: 0 0 240px;
    width: 240px;
  }
  
  .stat-card h3 {
    font-size: 2.4rem !important;
  }
  
  .stat-card .stat-icon .v-icon {
    font-size: 68px !important;
    width: 68px !important;
    height: 68px !important;
  }
  
  .page-title {
    font-size: 3.5rem !important;
  }
  
  .page-subtitle {
    font-size: 1.5rem !important;
  }
  
  .page-header {
    padding: 4rem 2rem;
  }
  
  .page-icon {
    font-size: 6rem !important;
  }
}

/* تحسينات عامة للمكونات */
.v-card {
  border-radius: 20px !important;
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




/* ========================================
   9. تنسيق مربعات المشاريع
   ======================================== */

.project-card {
  height: 100%;
  border-radius: 24px !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border: 2px solid #e2e8f0 !important;
  overflow: hidden;
  position: relative;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08) !important;
}

.project-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background: linear-gradient(90deg, #3b82f6 0%, #8b5cf6 50%, #3b82f6 100%);
  background-size: 200% 100%;
  opacity: 0;
  transition: all 0.4s ease;
  animation: gradientShift 3s ease infinite;
  z-index: 1;
}

.project-card::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(59, 130, 246, 0.1) 0%, transparent 70%);
  opacity: 0;
  transition: opacity 0.4s ease;
  pointer-events: none;
}

.project-card:hover {
  transform: translateY(-12px) scale(1.03) !important;
  box-shadow: 0 24px 48px rgba(59, 130, 246, 0.25), 0 8px 16px rgba(0, 0, 0, 0.15) !important;
  border-color: #3b82f6 !important;
}

.project-card:hover::before {
  opacity: 1;
  height: 6px;
}

.project-card:hover::after {
  opacity: 1;
}

/* عنوان البطاقة */
.project-card-title {
  padding: 1.75rem 1.75rem 1.25rem !important;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
}

.project-card-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s ease;
}

.project-card:hover .project-card-title::before {
  left: 100%;
}

.project-card-title .v-icon {
  color: #ffffff !important;
}

.project-card-title .v-avatar {
  background: rgba(255, 255, 255, 0.25) !important;
  color: #ffffff !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
  transition: all 0.3s ease !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
}

.project-card:hover .project-card-title .v-avatar {
  transform: scale(1.15) rotate(5deg) !important;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3) !important;
  background: rgba(255, 255, 255, 0.35) !important;
}

.project-card-title .v-avatar .v-icon {
  color: #ffffff !important;
}

.project-name {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-size: 0.95rem !important;
  font-weight: 800 !important;
  color: #ffffff !important;
  margin: 0 !important;
  line-height: 1.3 !important;
  letter-spacing: 0.3px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  text-align: right !important;
  word-wrap: break-word !important;
  overflow-wrap: break-word !important;
  hyphens: auto !important;
}

/* محتوى البطاقة */
.project-card-content {
  padding: 1.75rem !important;
  flex-grow: 1;
  background: linear-gradient(180deg, #ffffff 0%, #fafbfc 100%) !important;
}

/* تنسيق النصوص في بطاقات المشاريع - نظام نيلي موحد */
.project-description {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  color: #4b5563 !important;
  font-size: 0.8rem !important;
  line-height: 1.4 !important;
  margin: 0 auto 0.75rem auto !important;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  font-weight: 500 !important;
  letter-spacing: 0.1px !important;
  text-align: center !important;
  padding: 0.5rem 0.6rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 8px;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
  width: fit-content;
  max-width: 100%;
  margin-left: auto !important;
  margin-right: auto !important;
  justify-content: center;
  align-items: center;
}

.project-card:hover .project-description {
  background: linear-gradient(135deg, #e0f2fe 0%, #dbeafe 100%);
  border-color: #3b82f6;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.1);
}

.detail-text {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  color: #1e293b !important;
  font-size: 0.8rem !important;
  font-weight: 600 !important;
  line-height: 1.5 !important;
  letter-spacing: 0.1px !important;
  text-align: center !important;
  word-wrap: break-word !important;
  overflow-wrap: break-word !important;
  max-width: 100% !important;
  display: inline-block !important;
  transition: color 0.3s ease;
}

.detail-item:hover .detail-text {
  color: #3b82f6 !important;
  font-weight: 700 !important;
}

/* تفاصيل المشروع */
.project-details {
  display: flex;
    flex-direction: column;
  gap: 0.75rem;
}

.detail-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 12px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid #e2e8f0;
  position: relative;
  overflow: hidden;
}

.detail-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(59, 130, 246, 0.1), transparent);
  transition: left 0.5s ease;
}

.detail-item:hover {
  background: linear-gradient(135deg, #e0f2fe 0%, #dbeafe 100%);
  transform: translateY(-2px);
  border-color: #3b82f6;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.15);
}

.detail-item:hover::before {
  left: 100%;
}

.detail-item .v-icon {
  transition: all 0.3s ease;
}

.detail-item:hover .v-icon {
  transform: scale(1.2);
  color: #3b82f6 !important;
}


/* أزرار البطاقة */
.project-card-actions {
  padding: 1.25rem 1.75rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-top: 2px solid #e2e8f0;
  position: relative;
}

.project-card-actions::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #3b82f6 0%, #8b5cf6 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.project-card:hover .project-card-actions::before {
  opacity: 1;
}

/* Progress Section Styles */
.progress-section {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.08) 0%, rgba(139, 92, 246, 0.05) 100%);
  border-radius: 12px;
  padding: 0.75rem;
  border: 1.5px solid rgba(59, 130, 246, 0.15);
  margin-top: 1rem;
  position: relative;
  overflow: hidden;
}

.progress-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #3b82f6 0%, #8b5cf6 100%);
  border-radius: 12px 12px 0 0;
}

.progress-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 6px;
}

.progress-label {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  color: #1e293b !important;
  font-size: 0.85rem !important;
  font-weight: 700 !important;
  letter-spacing: 0.2px !important;
  text-align: center !important;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.4rem;
}

.progress-percentage {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  color: #3b82f6 !important;
  font-size: 0.85rem !important;
  font-weight: 800 !important;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.15) 0%, rgba(139, 92, 246, 0.1) 100%) !important;
  padding: 4px 10px !important;
  border-radius: 8px !important;
  border: 1.5px solid rgba(59, 130, 246, 0.3) !important;
  text-shadow: 0 1px 2px rgba(59, 130, 246, 0.2) !important;
  letter-spacing: 0.3px !important;
  box-shadow: 0 1px 4px rgba(59, 130, 246, 0.2) !important;
  transition: all 0.3s ease !important;
}

.project-card:hover .progress-percentage {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
}

.progress-bar {
  border-radius: 6px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  height: 6px !important;
}

.progress-bar .v-progress-linear__determinate {
  background: linear-gradient(90deg, #3b82f6 0%, #8b5cf6 50%, #3b82f6 100%);
  background-size: 200% 100%;
  border-radius: 8px;
  animation: gradientShift 3s ease infinite;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.4);
}

.progress-bar .v-progress-linear__background {
  background: rgba(59, 130, 246, 0.1);
  border-radius: 4px;
}

.details-btn {
  border-radius: 14px !important;
  font-weight: 700 !important;
  text-transform: none !important;
  padding: 0.75rem 2rem !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  position: relative;
  overflow: hidden;
}

.details-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.5s ease;
}

.details-btn:hover {
  transform: translateY(-3px) scale(1.02) !important;
  box-shadow: 0 8px 24px rgba(59, 130, 246, 0.4) !important;
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%) !important;
}

.details-btn:hover::before {
  left: 100%;
}

.details-btn:active {
  transform: translateY(-1px) scale(0.98) !important;
}

/* بطاقة عدم وجود مشاريع */
.no-projects-card {
  border-radius: 16px !important;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%) !important;
  border: 2px dashed #cbd5e1 !important;
}

/* تأثيرات متقدمة للبطاقات */
.project-card .v-avatar {
  transition: all 0.3s ease !important;
}

.project-card:hover .v-avatar {
  transform: scale(1.15) rotate(5deg) !important;
}

.project-card .v-chip {
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

.project-card:hover .v-chip {
  transform: scale(1.08) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

/* Animation للـ gradient */
@keyframes gradientShift {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

/* استجابة للشاشات المختلفة */
@media (max-width: 1200px) {
  .stats-row {
    gap: 0;
    padding: 0;
  }
  
  .stats-row .v-col {
    padding: 6px !important;
  }
  
  .stat-card {
    min-height: 120px !important;
    padding: 0.8rem !important;
  }
}

@media (max-width: 768px) {
  .stats-row .v-col {
    padding: 4px !important;
  }
  
  .stat-card {
    min-height: 100px !important;
    padding: 0.6rem !important;
  }
  
  .stat-card h3 {
    font-size: 1.5rem !important;
  }
  
  .stat-card .v-icon {
    font-size: 40px !important;
    width: 40px !important;
    height: 40px !important;
  }
}

@media (max-width: 480px) {
  .stats-row .v-col {
    padding: 2px !important;
  }
  
  .stat-card {
    min-height: 90px !important;
    padding: 0.5rem !important;
  }
  
  .stat-card h3 {
    font-size: 1.2rem !important;
  }
  
  .stat-card .v-icon {
    font-size: 32px !important;
    width: 32px !important;
    height: 32px !important;
  }
}

@media (max-width: 768px) {
  .project-card-title {
    padding: 1rem !important;
  }
  
  .project-card-content {
    padding: 1rem !important;
  }
  
  .project-card-actions {
    padding: 0.75rem 1rem !important;
  }
  
  .project-name {
    font-size: 1rem !important;
  }
}

@media (max-width: 600px) {
  .stats-row {
    flex-wrap: wrap;
    gap: 1rem;
  }
  
  .detail-item {
    padding: 0.4rem;
  }
  
  .detail-text {
    font-size: 0.8rem !important;
  }
  
  .page-title {
    font-size: 1.8rem !important;
    padding: 0.6rem 1rem;
    gap: 0.6rem;
    flex-direction: column;
  }
  
  .title-icon {
    font-size: 2rem;
  }
  
  .title-text {
    font-size: 1.8rem;
  }
}

/* Enhanced Project Cards */
.project-card {
  height: 100% !important;
  min-height: 400px !important;
  border-radius: 16px !important;
  transition: all 0.3s ease !important;
  border: 1px solid rgba(0, 0, 0, 0.08) !important;
}

.project-card.warning-border {
  border: 2px solid #f59e0b !important;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.25) !important;
}

.project-card.warning-border:hover {
  border-color: #d97706 !important;
  box-shadow: 0 12px 24px rgba(245, 158, 11, 0.35) !important;
}

.project-card:hover {
  transform: translateY(-4px) !important;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15) !important;
  border-color: #3b82f6 !important;
}

.project-card-title {
  padding: 1.5rem 1.5rem 1rem 1.5rem !important;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.3) !important;
}

.project-card-content {
  padding: 1.5rem !important;
  flex-grow: 1 !important;
}

.project-card-actions {
  padding: 1rem 1.5rem 1.5rem 1.5rem !important;
  background: #fafbfc !important;
  border-top: 1px solid rgba(0, 0, 0, 0.05) !important;
}

.project-description-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  margin-bottom: 0.75rem;
}


.project-details {
  display: flex;
    flex-direction: column;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.detail-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem;
  background: #f8fafc;
  border-radius: 12px;
  border-left: 3px solid #3b82f6;
  transition: all 0.3s ease;
}

.detail-item:hover {
  background: #e2e8f0;
  transform: translateX(4px);
}

.detail-text {
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  color: #4b5563 !important;
  font-size: 0.75rem !important;
  font-weight: 600 !important;
  line-height: 1.4 !important;
  letter-spacing: 0.1px !important;
  text-align: right !important;
  word-wrap: break-word !important;
  overflow-wrap: break-word !important;
  max-width: 100% !important;
  display: inline-block !important;
}

/* تم نقل أنماط .status-chip-main إلى common-components.css */

.details-btn {
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  letter-spacing: 0.5px !important;
  padding: 0.75rem 1.5rem !important;
  border-radius: 12px !important;
  text-transform: none !important;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .project-card {
    min-height: 350px !important;
  }
  
  .project-card-title {
    padding: 1rem !important;
  }
  
  .project-card-content {
    padding: 1rem !important;
  }
  
  .project-card-actions {
    padding: 0.75rem 1rem 1rem 1rem !important;
  }
  
  .project-name {
    font-size: 1.1rem !important;
  }
}

/* Modern Project Dialog Styles */
.modern-project-dialog {
  border-radius: 20px !important;
  overflow: hidden !important;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15) !important;
  border: 4px solid #1e3a8a !important;
}

/* ========================================
   التصميم الجديد الحديث والواضح
   ======================================== */

/* النموذج الرئيسي */
.modern-project-dialog {
  border-radius: 24px !important;
  overflow: hidden !important;
  box-shadow: 0 30px 100px rgba(25, 118, 210, 0.4), 0 0 0 1px rgba(255, 255, 255, 0.1) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  position: relative !important;
  border: 2px solid rgba(25, 118, 210, 0.2) !important;
  backdrop-filter: blur(20px) !important;
}

/* الرأس الجديد */
.dialog-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 25%, #1e88e5 50%, #42a5f5 75%, #64b5f6 100%) !important;
  background-size: 400% 400% !important;
  animation: gradientShift 8s ease infinite !important;
  padding: 2.5rem !important;
  display: flex;
  align-items: center;
  gap: 2rem;
  position: relative;
  box-shadow: 0 8px 30px rgba(25, 118, 210, 0.4), 0 0 0 1px rgba(255, 255, 255, 0.2) !important;
  border-bottom: 3px solid rgba(255, 255, 255, 0.3) !important;
}

.header-icon {
  width: 70px;
  height: 70px;
  background: rgba(255, 255, 255, 0.25);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(15px);
  border: 3px solid rgba(255, 255, 255, 0.4);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(255, 255, 255, 0.2);
  animation: pulse 3s ease-in-out infinite;
}

.header-content {
  flex: 1;
}

.dialog-title {
  font-size: 2.2rem !important;
  font-weight: 800 !important;
  color: #4b5563 !important;
  margin: 0 0 0.8rem 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 4px 12px rgba(0, 0, 0, 0.5), 0 0 0 2px rgba(255, 255, 255, 0.1) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.2 !important;
  -webkit-text-stroke: 0.5px rgba(255, 255, 255, 0.3) !important;
}

.dialog-subtitle {
  font-size: 1.1rem !important;
  color: rgba(255, 255, 255, 0.95) !important;
  margin: 0 !important;
  font-weight: 500 !important;
  opacity: 1 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
}

.close-btn {
  opacity: 0.8 !important;
  transition: all 0.3s ease !important;
  background: rgba(255, 255, 255, 0.1) !important;
  border-radius: 50% !important;
  width: 40px !important;
  height: 40px !important;
}

.close-btn:hover {
  opacity: 1 !important;
  background: rgba(255, 255, 255, 0.2) !important;
  transform: scale(1.1) !important;
}

/* محتوى النموذج */
.dialog-body {
  padding: 1rem !important;
  background: #f8f9fa !important;
  max-height: 70vh !important;
  overflow-y: auto !important;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* أقسام النموذج */
.form-section {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border-radius: 16px !important;
  padding: 2rem !important;
  border: 2px solid rgba(25, 118, 210, 0.1) !important;
  box-shadow: 0 8px 25px rgba(25, 118, 210, 0.1), 0 0 0 1px rgba(255, 255, 255, 0.5) !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  position: relative !important;
  overflow: hidden !important;
  backdrop-filter: blur(10px) !important;
}

.form-section::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  height: 4px !important;
  background: linear-gradient(90deg, #1976d2, #42a5f5, #64b5f6) !important;
  border-radius: 16px 16px 0 0 !important;
}

.form-section:hover {
  box-shadow: 0 15px 40px rgba(25, 118, 210, 0.2), 0 0 0 1px rgba(255, 255, 255, 0.8) !important;
  transform: translateY(-4px) scale(1.02) !important;
  border-color: rgba(25, 118, 210, 0.3) !important;
}

.section-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
  padding: 1rem 1.5rem;
  background: linear-gradient(135deg, #1976d2 0%, #42a5f5 100%);
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(25, 118, 210, 0.2);
  position: relative;
  overflow: hidden;
}

.section-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmer 3s infinite;
}

.section-title {
  font-size: 1.3rem !important;
  font-weight: 700 !important;
  color: #4b5563 !important;
  margin: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.3 !important;
  -webkit-text-stroke: 0.5px rgba(255, 255, 255, 0.2) !important;
}

/* شبكة الحقول */
.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.75rem;
}

@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }
}

/* حقول النموذج */
.form-field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.field-label {
  font-size: 1.1rem !important;
  font-weight: 700 !important;
  color: #1976d2 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  margin-bottom: 0.8rem !important;
  text-shadow: 0 2px 4px rgba(25, 118, 210, 0.1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.3 !important;
  display: block !important;
  position: relative !important;
}

.field-label::before {
  content: '•' !important;
  color: #1976d2 !important;
  font-weight: 900 !important;
  margin-left: 0.5rem !important;
  font-size: 1.2rem !important;
}

/* حقول الإدخال الحديثة */
.modern-input .v-field {
  border-radius: 12px !important;
  border: 2px solid rgba(26, 35, 126, 0.3) !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.1) 0%, rgba(30, 58, 138, 0.15) 100%) !important;
  min-height: 48px !important;
  box-shadow: 0 4px 15px rgba(26, 35, 126, 0.2) !important;
  position: relative !important;
  backdrop-filter: blur(10px) !important;
}

/* تحسين خاص لحقل الملاحظات */
.modern-input.v-textarea .v-field {
  min-height: 140px !important;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.15) 0%, rgba(30, 58, 138, 0.2) 100%) !important;
  border: 3px solid rgba(26, 35, 126, 0.4) !important;
  box-shadow: 0 6px 20px rgba(26, 35, 126, 0.25) !important;
}

.modern-input.v-textarea .v-field:hover {
  border-color: #1a237e !important;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.25) 0%, rgba(30, 58, 138, 0.3) 100%) !important;
  box-shadow: 0 10px 30px rgba(26, 35, 126, 0.35), 0 0 0 2px rgba(26, 35, 126, 0.5) !important;
  transform: translateY(-3px) !important;
}

.modern-input.v-textarea .v-field--focused {
  border-color: #1a237e !important;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.3) 0%, rgba(30, 58, 138, 0.35) 100%) !important;
  box-shadow: 0 15px 40px rgba(26, 35, 126, 0.4), 0 0 0 3px rgba(26, 35, 126, 0.6) !important;
  transform: translateY(-4px) scale(1.02) !important;
}

.modern-input .v-field:hover {
  border-color: #1a237e !important;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.2) 0%, rgba(30, 58, 138, 0.25) 100%) !important;
  box-shadow: 0 8px 25px rgba(26, 35, 126, 0.3), 0 0 0 1px rgba(26, 35, 126, 0.4) !important;
  transform: translateY(-2px) !important;
}

.modern-input .v-field--focused {
  border-color: #1a237e !important;
  background: linear-gradient(135deg, rgba(26, 35, 126, 0.25) 0%, rgba(30, 58, 138, 0.3) 100%) !important;
  box-shadow: 0 12px 30px rgba(26, 35, 126, 0.4), 0 0 0 2px rgba(26, 35, 126, 0.5) !important;
  transform: translateY(-3px) scale(1.02) !important;
}

.modern-input .v-field__input {
  color: #1a237e !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 0.75rem 1rem !important;
  line-height: 1.4 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: transparent !important;
  border-radius: 8px !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.2) !important;
  letter-spacing: 0.5px !important;
}

/* ضمان أن النص المكتوب في حقل اسم المشروع يكون نيلي */
.modern-input .v-field__input,
.modern-input .v-field__input input,
.modern-input .v-field__input textarea {
  color: #1a237e !important;
  font-weight: 700 !important;
}

/* تحسين خاص لحقل اسم المشروع */
.modern-input .v-text-field .v-field__input {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
}

/* تحسين خاص لحقل الملاحظات */
.modern-input.v-textarea .v-field__input {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  padding: 1rem !important;
  line-height: 1.6 !important;
  text-shadow: 0 3px 6px rgba(26, 35, 126, 0.3) !important;
  letter-spacing: 0.7px !important;
  min-height: 120px !important;
}

.modern-input.v-textarea .v-field__input::placeholder {
  color: rgba(26, 35, 126, 0.6) !important;
  font-weight: 600 !important;
  opacity: 0.8 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(26, 35, 126, 0.2) !important;
}

.modern-input .v-field__input::placeholder {
  color: rgba(26, 35, 126, 0.5) !important;
  font-weight: 600 !important;
  opacity: 0.8 !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(26, 35, 126, 0.1) !important;
}

.modern-input .v-field__append-inner {
  color: #1a237e !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  text-shadow: 0 1px 2px rgba(26, 35, 126, 0.2) !important;
}

.modern-input .v-field__outline {
  display: none !important;
}

/* ضمان أن النص في جميع أنواع الحقول يكون نيلي */
.modern-input input,
.modern-input textarea,
.modern-input .v-field__input,
.modern-input .v-field__input input,
.modern-input .v-field__input textarea,
.modern-input .v-input__control input,
.modern-input .v-input__control textarea {
  color: #1a237e !important;
  font-weight: 700 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* تحسين خاص لحقول النص */
.modern-input .v-text-field input {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
}

/* تحسين خاص لحقول التاريخ */
.modern-input .v-text-field input[type="date"] {
  color: #1a237e !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
}

/* تحسين خاص لحقول الأرقام */
.modern-input .v-text-field input[type="number"] {
  color: #1a237e !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
}

.modern-input .v-field__overlay {
  display: none !important;
}

/* CSS قوي جداً لضمان ظهور النص النيلي */
.modern-project-dialog .modern-input input,
.modern-project-dialog .modern-input textarea,
.modern-project-dialog .modern-input .v-field__input,
.modern-project-dialog .v-text-field input,
.modern-project-dialog .v-textarea textarea,
.modern-project-dialog .v-input input,
.modern-project-dialog .v-input textarea {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS إضافي لجميع عناصر الإدخال */
.modern-project-dialog input[type="text"],
.modern-project-dialog input[type="date"],
.modern-project-dialog input[type="number"],
.modern-project-dialog textarea {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS للعناصر المباشرة */
.modern-project-dialog .v-field__input input,
.modern-project-dialog .v-field__input textarea {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS باستخدام ::deep لضمان التطبيق */
::deep(.modern-input input),
::deep(.modern-input textarea),
::deep(.modern-input .v-field__input),
::deep(.v-text-field input),
::deep(.v-textarea textarea) {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS إضافي لجميع عناصر Vuetify */
::deep(.v-field__input),
::deep(.v-field__input input),
::deep(.v-field__input textarea),
::deep(.v-input__control input),
::deep(.v-input__control textarea) {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* ========================================
   التصميم الجديد الفائق الحداثة
   ======================================== */

/* القائمة المنبثقة الرئيسية */
.ultra-modern-dialog {
  border-radius: 24px !important;
  overflow: hidden !important;
  box-shadow: 0 40px 120px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(255, 255, 255, 0.1) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  position: relative !important;
  border: 2px solid rgba(26, 35, 126, 0.1) !important;
  backdrop-filter: blur(20px) !important;
}

/* الرأس الفائق الحداثة */
.ultra-dialog-header {
  background: linear-gradient(135deg, #1a237e 0%, #283593 25%, #3949ab 50%, #5c6bc0 75%, #7986cb 100%) !important;
  background-size: 400% 400% !important;
  animation: gradientShift 8s ease infinite !important;
  padding: 3rem 2rem !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 10px 40px rgba(26, 35, 126, 0.4) !important;
}

.header-background {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: radial-gradient(circle at 20% 80%, rgba(255, 255, 255, 0.1) 0%, transparent 50%),
              radial-gradient(circle at 80% 20%, rgba(255, 255, 255, 0.1) 0%, transparent 50%) !important;
  pointer-events: none !important;
}

.header-pattern {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.05) 50%, transparent 70%) !important;
  animation: shimmer 3s infinite !important;
}

.header-content {
  display: flex !important;
  align-items: center !important;
  gap: 2rem !important;
  position: relative !important;
  z-index: 1 !important;
}

.header-icon-container {
  position: relative !important;
}

.icon-circle {
  width: 80px !important;
  height: 80px !important;
  background: rgba(255, 255, 255, 0.2) !important;
  border-radius: 50% !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  backdrop-filter: blur(10px) !important;
  border: 3px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2) !important;
  animation: pulse 3s ease-in-out infinite !important;
}

.icon-glow {
  position: absolute !important;
  top: -10px !important;
  left: -10px !important;
  right: -10px !important;
  bottom: -10px !important;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.3) 0%, transparent 70%) !important;
  border-radius: 50% !important;
  animation: glow 2s ease-in-out infinite alternate !important;
}

.header-text {
  flex: 1 !important;
}

.ultra-dialog-title {
  color: #4b5563 !important;
  font-size: 2.5rem !important;
  font-weight: 900 !important;
  margin: 0 0 1rem 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 4px 12px rgba(0, 0, 0, 0.5), 0 0 0 2px rgba(255, 255, 255, 0.1) !important;
  letter-spacing: 1px !important;
  line-height: 1.2 !important;
  -webkit-text-stroke: 1px rgba(255, 255, 255, 0.3) !important;
}

.ultra-dialog-subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 1.2rem !important;
  margin: 0 !important;
  font-weight: 500 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.4 !important;
}

.ultra-close-btn {
  background: rgba(255, 255, 255, 0.1) !important;
  border-radius: 50% !important;
  width: 50px !important;
  height: 50px !important;
  backdrop-filter: blur(10px) !important;
  border: 2px solid rgba(255, 255, 255, 0.2) !important;
  transition: all 0.3s ease !important;
}

.ultra-close-btn:hover {
  background: rgba(255, 255, 255, 0.2) !important;
  transform: scale(1.1) !important;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2) !important;
}

/* الأقسام الفائقة الحداثة */
.ultra-form-container {
  padding: 2rem !important;
  display: flex !important;
  flex-direction: column !important;
  gap: 2rem !important;
}

.ultra-form-section {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border-radius: 20px !important;
  padding: 2rem !important;
  border: 2px solid rgba(26, 35, 126, 0.1) !important;
  box-shadow: 0 10px 30px rgba(26, 35, 126, 0.1), 0 0 0 1px rgba(255, 255, 255, 0.5) !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  position: relative !important;
  overflow: hidden !important;
  backdrop-filter: blur(10px) !important;
}

.ultra-form-section::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  height: 4px !important;
  background: linear-gradient(90deg, #1a237e, #3949ab, #5c6bc0) !important;
  border-radius: 20px 20px 0 0 !important;
}

.ultra-form-section:hover {
  box-shadow: 0 20px 50px rgba(26, 35, 126, 0.2), 0 0 0 1px rgba(255, 255, 255, 0.8) !important;
  transform: translateY(-5px) scale(1.02) !important;
  border-color: rgba(26, 35, 126, 0.3) !important;
}

.ultra-section-header {
  display: flex !important;
  align-items: center !important;
  gap: 1.5rem !important;
  margin-bottom: 2rem !important;
  padding: 1.5rem 2rem !important;
  background: linear-gradient(135deg, #1a237e 0%, #3949ab 100%) !important;
  border-radius: 16px !important;
  box-shadow: 0 6px 20px rgba(26, 35, 126, 0.3) !important;
  position: relative !important;
  overflow: hidden !important;
}

.ultra-section-header::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%) !important;
  animation: shimmer 3s infinite !important;
}

.section-icon {
  width: 50px !important;
  height: 50px !important;
  background: rgba(255, 255, 255, 0.2) !important;
  border-radius: 50% !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  backdrop-filter: blur(10px) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2) !important;
}

.ultra-section-title {
  font-size: 1.5rem !important;
  font-weight: 800 !important;
  color: #4b5563 !important;
  margin: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.3 !important;
  -webkit-text-stroke: 0.5px rgba(255, 255, 255, 0.2) !important;
}

.section-line {
  flex: 1 !important;
  height: 2px !important;
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.3) 0%, transparent 100%) !important;
  border-radius: 1px !important;
}

/* الشبكة الفائقة الحداثة */
.ultra-form-grid {
  display: grid !important;
  grid-template-columns: 1fr 1fr !important;
  gap: 2rem !important;
}

.ultra-form-field {
  display: flex !important;
  flex-direction: column !important;
  gap: 0.8rem !important;
}

.ultra-field-label {
  font-size: 1.1rem !important;
  font-weight: 700 !important;
  color: #1a237e !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.3 !important;
  display: flex !important;
  align-items: center !important;
  gap: 0.5rem !important;
}

/* الحقول الفائقة الحداثة */
.ultra-modern-input .v-field {
  border-radius: 16px !important;
  border: 3px solid rgba(26, 35, 126, 0.2) !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%) !important;
  min-height: 56px !important;
  box-shadow: 0 6px 20px rgba(26, 35, 126, 0.15) !important;
  position: relative !important;
  backdrop-filter: blur(10px) !important;
}

.ultra-modern-input .v-field:hover {
  border-color: #1a237e !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(248, 250, 252, 1) 100%) !important;
  box-shadow: 0 12px 35px rgba(26, 35, 126, 0.25), 0 0 0 2px rgba(26, 35, 126, 0.3) !important;
  transform: translateY(-3px) !important;
}

.ultra-modern-input .v-field--focused {
  border-color: #1a237e !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(248, 250, 252, 1) 100%) !important;
  box-shadow: 0 15px 45px rgba(26, 35, 126, 0.3), 0 0 0 3px rgba(26, 35, 126, 0.4) !important;
  transform: translateY(-4px) scale(1.02) !important;
}

.ultra-modern-input .v-field__input {
  color: #1a237e !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  padding: 1rem 1.5rem !important;
  line-height: 1.4 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: transparent !important;
  border-radius: 12px !important;
  text-shadow: 0 2px 4px rgba(26, 35, 126, 0.2) !important;
  letter-spacing: 0.5px !important;
}

.ultra-modern-input .v-field__input::placeholder {
  color: rgba(26, 35, 126, 0.5) !important;
  font-weight: 600 !important;
  opacity: 0.8 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(26, 35, 126, 0.1) !important;
}

/* الأزرار الفائقة الحداثة */
.ultra-dialog-footer {
  padding: 3rem 2rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%) !important;
  border-top: 3px solid rgba(26, 35, 126, 0.1) !important;
  position: relative !important;
  overflow: hidden !important;
  backdrop-filter: blur(10px) !important;
}

.footer-background {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: radial-gradient(circle at 50% 0%, rgba(26, 35, 126, 0.05) 0%, transparent 50%) !important;
  pointer-events: none !important;
}

.footer-content {
  display: flex !important;
  align-items: center !important;
  justify-content: flex-end !important;
  gap: 1.5rem !important;
  position: relative !important;
  z-index: 1 !important;
}

.ultra-cancel-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 16px !important;
  padding: 1rem 2rem !important;
  border: 3px solid rgba(156, 163, 175, 0.3) !important;
  color: #6b7280 !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: rgba(255, 255, 255, 0.9) !important;
  backdrop-filter: blur(10px) !important;
  box-shadow: 0 6px 20px rgba(156, 163, 175, 0.15) !important;
  letter-spacing: 0.5px !important;
  min-width: 140px !important;
}

.ultra-cancel-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: #9ca3af !important;
  color: #374151 !important;
  transform: translateY(-3px) scale(1.05) !important;
  box-shadow: 0 12px 35px rgba(156, 163, 175, 0.25) !important;
}

.ultra-save-btn {
  font-weight: 900 !important;
  text-transform: none !important;
  border-radius: 16px !important;
  padding: 1rem 2.5rem !important;
  background: linear-gradient(135deg, #1a237e 0%, #3949ab 50%, #5c6bc0 100%) !important;
  background-size: 200% 200% !important;
  animation: gradientShift 3s ease infinite !important;
  box-shadow: 0 10px 30px rgba(26, 35, 126, 0.4) !important;
  font-size: 1.2rem !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  color: #4b5563 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  border: 3px solid rgba(255, 255, 255, 0.2) !important;
  backdrop-filter: blur(10px) !important;
  min-width: 180px !important;
}

.ultra-save-btn:hover {
  background: linear-gradient(135deg, #0d47a1 0%, #1a237e 50%, #283593 100%) !important;
  transform: translateY(-4px) scale(1.08) !important;
  box-shadow: 0 20px 50px rgba(26, 35, 126, 0.5) !important;
  border-color: rgba(255, 255, 255, 0.4) !important;
}

.ultra-save-btn:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #d1d5db 100%) !important;
  color: #6b7280 !important;
  transform: none !important;
  box-shadow: 0 4px 15px rgba(156, 163, 175, 0.2) !important;
  border-color: rgba(156, 163, 175, 0.3) !important;
}

/* الرسوم المتحركة */
@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

@keyframes glow {
  0% { opacity: 0.5; transform: scale(1); }
  100% { opacity: 0.8; transform: scale(1.1); }
}

@keyframes pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translateY(30px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-10px);
  }
  60% {
    transform: translateY(-5px);
  }
}

/* التجاوب */
@media (max-width: 768px) {
  .ultra-dialog-header {
    padding: 2rem 1rem !important;
  }
  
  .header-content {
    flex-direction: column !important;
    text-align: center !important;
    gap: 1rem !important;
  }
  
  .ultra-dialog-title {
    font-size: 2rem !important;
  }
  
  .ultra-form-container {
    padding: 1rem !important;
  }
  
  .ultra-form-section {
    padding: 1.5rem !important;
  }
  
  .ultra-form-grid {
    grid-template-columns: 1fr !important;
    gap: 1.5rem !important;
  }
  
  .footer-content {
    flex-direction: column !important;
    gap: 1rem !important;
  }
  
  .ultra-cancel-btn,
  .ultra-save-btn {
    width: 100% !important;
    min-width: auto !important;
  }
}

/* ========================================
   التصميم الجديد الفائق الوضوح
   ======================================== */

/* القائمة المنبثقة الجديدة */
.new-ultra-dialog {
  border-radius: 28px !important;
  overflow: hidden !important;
  box-shadow: 0 50px 150px rgba(0, 0, 0, 0.4), 0 0 0 1px rgba(255, 255, 255, 0.2) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  position: relative !important;
  border: 3px solid rgba(26, 35, 126, 0.15) !important;
  backdrop-filter: blur(25px) !important;
}

/* الرأس الجديد */
.new-dialog-header {
  background: linear-gradient(135deg, #0d47a1 0%, #1a237e 25%, #283593 50%, #3949ab 75%, #5c6bc0 100%) !important;
  background-size: 500% 500% !important;
  animation: newGradientShift 10s ease infinite !important;
  padding: 3.5rem 2.5rem !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 15px 50px rgba(13, 71, 161, 0.5) !important;
}

.header-gradient-bg {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: radial-gradient(circle at 30% 70%, rgba(255, 255, 255, 0.15) 0%, transparent 60%),
              radial-gradient(circle at 70% 30%, rgba(255, 255, 255, 0.1) 0%, transparent 60%) !important;
  pointer-events: none !important;
}

.header-content-wrapper {
  position: relative !important;
  z-index: 1 !important;
}

.header-main-content {
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
  gap: 2rem !important;
}

.title-section {
  display: flex !important;
  align-items: center !important;
  gap: 2rem !important;
}

.title-icon {
  width: 90px !important;
  height: 90px !important;
  background: rgba(255, 255, 255, 0.25) !important;
  border-radius: 50% !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  backdrop-filter: blur(15px) !important;
  border: 4px solid rgba(255, 255, 255, 0.4) !important;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3) !important;
  animation: newPulse 4s ease-in-out infinite !important;
}

.title-text {
  flex: 1 !important;
}

.new-dialog-title {
  color: #4b5563 !important;
  font-size: 3rem !important;
  font-weight: 900 !important;
  margin: 0 0 1.2rem 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 6px 15px rgba(0, 0, 0, 0.6), 0 0 0 3px rgba(255, 255, 255, 0.2) !important;
  letter-spacing: 1.5px !important;
  line-height: 1.1 !important;
  -webkit-text-stroke: 1.5px rgba(255, 255, 255, 0.4) !important;
}

.new-dialog-subtitle {
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 1.4rem !important;
  margin: 0 !important;
  font-weight: 600 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 3px 8px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.8px !important;
  line-height: 1.4 !important;
}

.new-close-btn {
  background: rgba(255, 255, 255, 0.15) !important;
  border-radius: 50% !important;
  width: 60px !important;
  height: 60px !important;
  backdrop-filter: blur(15px) !important;
  border: 3px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2) !important;
}

.new-close-btn:hover {
  background: rgba(255, 255, 255, 0.25) !important;
  transform: scale(1.15) rotate(90deg) !important;
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.3) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

/* الأقسام الجديدة */
.new-form-container {
  padding: 3rem 2.5rem !important;
  display: flex !important;
  flex-direction: column !important;
  gap: 2.5rem !important;
}

.new-form-section {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  border-radius: 24px !important;
  padding: 2.5rem !important;
  border: 3px solid rgba(26, 35, 126, 0.12) !important;
  box-shadow: 0 15px 40px rgba(26, 35, 126, 0.15), 0 0 0 1px rgba(255, 255, 255, 0.6) !important;
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  position: relative !important;
  overflow: hidden !important;
  backdrop-filter: blur(15px) !important;
}

.new-form-section::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  height: 6px !important;
  background: linear-gradient(90deg, #0d47a1, #1a237e, #283593, #3949ab, #5c6bc0) !important;
  border-radius: 24px 24px 0 0 !important;
}

.new-form-section:hover {
  box-shadow: 0 25px 60px rgba(26, 35, 126, 0.25), 0 0 0 1px rgba(255, 255, 255, 0.9) !important;
  transform: translateY(-8px) scale(1.03) !important;
  border-color: rgba(26, 35, 126, 0.4) !important;
}

.new-section-header {
  display: flex !important;
  align-items: center !important;
  gap: 2rem !important;
  margin-bottom: 2.5rem !important;
  padding: 2rem 2.5rem !important;
  background: linear-gradient(135deg, #0d47a1 0%, #1a237e 25%, #283593 50%, #3949ab 75%, #5c6bc0 100%) !important;
  border-radius: 20px !important;
  box-shadow: 0 8px 25px rgba(13, 71, 161, 0.4) !important;
  position: relative !important;
  overflow: hidden !important;
}

.new-section-header::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.12) 50%, transparent 70%) !important;
  animation: newShimmer 4s infinite !important;
}

.section-icon-wrapper {
  width: 60px !important;
  height: 60px !important;
  background: rgba(255, 255, 255, 0.25) !important;
  border-radius: 50% !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  backdrop-filter: blur(15px) !important;
  border: 3px solid rgba(255, 255, 255, 0.4) !important;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.25) !important;
}

.new-section-title {
  font-size: 1.8rem !important;
  font-weight: 800 !important;
  color: #4b5563 !important;
  margin: 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 3px 10px rgba(0, 0, 0, 0.4) !important;
  letter-spacing: 0.8px !important;
  line-height: 1.3 !important;
  -webkit-text-stroke: 0.8px rgba(255, 255, 255, 0.3) !important;
}

.section-divider {
  flex: 1 !important;
  height: 3px !important;
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.4) 0%, transparent 100%) !important;
  border-radius: 2px !important;
}

/* الشبكة الجديدة */
.new-form-grid {
  display: grid !important;
  grid-template-columns: 1fr 1fr !important;
  gap: 2.5rem !important;
}

.new-form-field {
  display: flex !important;
  flex-direction: column !important;
  gap: 1rem !important;
}

.new-field-label {
  font-size: 1.3rem !important;
  font-weight: 800 !important;
  color: #0d47a1 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.2) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.4 !important;
  display: flex !important;
  align-items: center !important;
  gap: 0.8rem !important;
  margin-bottom: 0.5rem !important;
}

/* الحقول الجديدة مع وضوح فائق */
.new-ultra-input .v-field {
  border-radius: 20px !important;
  border: 4px solid rgba(13, 71, 161, 0.25) !important;
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(248, 250, 252, 0.95) 100%) !important;
  min-height: 64px !important;
  box-shadow: 0 8px 25px rgba(13, 71, 161, 0.2) !important;
  position: relative !important;
  backdrop-filter: blur(15px) !important;
}

.new-ultra-input .v-field:hover {
  border-color: #0d47a1 !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(248, 250, 252, 1) 100%) !important;
  box-shadow: 0 15px 45px rgba(13, 71, 161, 0.3), 0 0 0 3px rgba(13, 71, 161, 0.4) !important;
  transform: translateY(-4px) !important;
}

.new-ultra-input .v-field--focused {
  border-color: #0d47a1 !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(248, 250, 252, 1) 100%) !important;
  box-shadow: 0 20px 60px rgba(13, 71, 161, 0.4), 0 0 0 4px rgba(13, 71, 161, 0.5) !important;
  transform: translateY(-6px) scale(1.03) !important;
}

/* النص فائق الوضوح */
.new-ultra-input .v-field__input {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  padding: 1.2rem 1.8rem !important;
  line-height: 1.5 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: transparent !important;
  border-radius: 16px !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  letter-spacing: 0.8px !important;
}

.new-ultra-input .v-field__input::placeholder {
  color: rgba(13, 71, 161, 0.6) !important;
  font-weight: 700 !important;
  opacity: 0.9 !important;
  font-size: 1.3rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(13, 71, 161, 0.2) !important;
}

/* CSS قوي جداً لضمان ظهور النص المكتوب */
.new-ultra-dialog input,
.new-ultra-dialog textarea,
.new-ultra-dialog .v-field__input,
.new-ultra-dialog .v-field__input input,
.new-ultra-dialog .v-field__input textarea,
.new-ultra-dialog .v-input__control input,
.new-ultra-dialog .v-input__control textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS إضافي لجميع عناصر الإدخال */
.new-ultra-dialog input[type="text"],
.new-ultra-dialog input[type="date"],
.new-ultra-dialog input[type="number"],
.new-ultra-dialog textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS للعناصر المباشرة */
.new-ultra-dialog .v-field__input input,
.new-ultra-dialog .v-field__input textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS باستخدام ::deep لضمان التطبيق */
::deep(.new-ultra-input input),
::deep(.new-ultra-input textarea),
::deep(.new-ultra-input .v-field__input),
::deep(.v-text-field input),
::deep(.v-textarea textarea) {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS إضافي لجميع عناصر Vuetify */
::deep(.v-field__input),
::deep(.v-field__input input),
::deep(.v-field__input textarea),
::deep(.v-input__control input),
::deep(.v-input__control textarea) {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS قوي جداً لضمان ظهور النص في جميع الحالات */
.new-ultra-dialog .v-text-field .v-field__input,
.new-ultra-dialog .v-textarea .v-field__input,
.new-ultra-dialog .v-field__input input,
.new-ultra-dialog .v-field__input textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS للعناصر المباشرة في القائمة المنبثقة */
.new-ultra-dialog input,
.new-ultra-dialog textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS قوي جداً لضمان ظهور النص المكتوب */
.new-ultra-dialog * input,
.new-ultra-dialog * textarea,
.new-ultra-dialog * .v-field__input,
.new-ultra-dialog * .v-field__input input,
.new-ultra-dialog * .v-field__input textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS إضافي لجميع العناصر */
.new-ultra-dialog .v-field__input,
.new-ultra-dialog .v-field__input input,
.new-ultra-dialog .v-field__input textarea,
.new-ultra-dialog .v-input__control input,
.new-ultra-dialog .v-input__control textarea,
.new-ultra-dialog .v-text-field input,
.new-ultra-dialog .v-textarea textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS للعناصر المباشرة */
.new-ultra-dialog input[type="text"],
.new-ultra-dialog input[type="date"],
.new-ultra-dialog input[type="number"],
.new-ultra-dialog textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* CSS قوي جداً لضمان ظهور النص */
.new-ultra-dialog .v-field__input input,
.new-ultra-dialog .v-field__input textarea,
.new-ultra-dialog .v-input__control input,
.new-ultra-dialog .v-input__control textarea {
  color: #0d47a1 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(13, 71, 161, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
}

/* الأزرار الجديدة */
.new-dialog-footer {
  padding: 3.5rem 2.5rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%) !important;
  border-top: 4px solid rgba(13, 71, 161, 0.15) !important;
  position: relative !important;
  overflow: hidden !important;
  backdrop-filter: blur(15px) !important;
}

.footer-gradient-bg {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: radial-gradient(circle at 50% 0%, rgba(13, 71, 161, 0.08) 0%, transparent 60%) !important;
  pointer-events: none !important;
}

.footer-content-wrapper {
  display: flex !important;
  align-items: center !important;
  justify-content: flex-end !important;
  gap: 2rem !important;
  position: relative !important;
  z-index: 1 !important;
}

.new-cancel-btn {
  font-weight: 800 !important;
  text-transform: none !important;
  border-radius: 20px !important;
  padding: 1.2rem 2.5rem !important;
  border: 4px solid rgba(156, 163, 175, 0.4) !important;
  color: #6b7280 !important;
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  font-size: 1.2rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: rgba(255, 255, 255, 0.95) !important;
  backdrop-filter: blur(15px) !important;
  box-shadow: 0 8px 25px rgba(156, 163, 175, 0.2) !important;
  letter-spacing: 0.8px !important;
  min-width: 160px !important;
}

.new-cancel-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: #9ca3af !important;
  color: #374151 !important;
  transform: translateY(-4px) scale(1.08) !important;
  box-shadow: 0 15px 45px rgba(156, 163, 175, 0.3) !important;
}

.new-save-btn {
  font-weight: 900 !important;
  text-transform: none !important;
  border-radius: 20px !important;
  padding: 1.2rem 3rem !important;
  background: linear-gradient(135deg, #0d47a1 0%, #1a237e 25%, #283593 50%, #3949ab 75%, #5c6bc0 100%) !important;
  background-size: 300% 300% !important;
  animation: newGradientShift 4s ease infinite !important;
  box-shadow: 0 15px 40px rgba(13, 71, 161, 0.5) !important;
  font-size: 1.3rem !important;
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  color: #4b5563 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3) !important;
  border: 4px solid rgba(255, 255, 255, 0.3) !important;
  backdrop-filter: blur(15px) !important;
  min-width: 200px !important;
}

.new-save-btn:hover {
  background: linear-gradient(135deg, #0d47a1 0%, #1a237e 50%, #283593 100%) !important;
  transform: translateY(-6px) scale(1.1) !important;
  box-shadow: 0 25px 60px rgba(13, 71, 161, 0.6) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.new-save-btn:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #d1d5db 100%) !important;
  color: #6b7280 !important;
  transform: none !important;
  box-shadow: 0 6px 20px rgba(156, 163, 175, 0.25) !important;
  border-color: rgba(156, 163, 175, 0.4) !important;
}

/* الرسوم المتحركة الجديدة */
@keyframes newGradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

@keyframes newShimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

@keyframes newPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.08); }
}

/* ========================================
   التصميم الجديد المختلف تماماً
   ======================================== */

/* القائمة المنبثقة الجديدة */
.super-modern-dialog {
  border-radius: 20px !important;
  overflow: hidden !important;
  box-shadow: 0 25px 80px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(255, 255, 255, 0.1) !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
  position: relative !important;
  border: 2px solid rgba(25, 118, 210, 0.1) !important;
  backdrop-filter: blur(20px) !important;
}

/* الرأس الجديد */
.super-dialog-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 25%, #0d47a1 50%, #1565c0 75%, #1976d2 100%) !important;
  background-size: 300% 300% !important;
  animation: superGradientShift 6s ease infinite !important;
  padding: 2rem !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 8px 30px rgba(25, 118, 210, 0.4) !important;
}

.header-pattern-bg {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: radial-gradient(circle at 25% 75%, rgba(255, 255, 255, 0.1) 0%, transparent 50%),
              radial-gradient(circle at 75% 25%, rgba(255, 255, 255, 0.08) 0%, transparent 50%) !important;
  pointer-events: none !important;
}

.header-main {
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
  gap: 1.5rem !important;
  position: relative !important;
  z-index: 1 !important;
}

.header-left-section {
  display: flex !important;
  align-items: center !important;
  gap: 1.5rem !important;
}

.icon-container {
  width: 60px !important;
  height: 60px !important;
  background: rgba(255, 255, 255, 0.2) !important;
  border-radius: 50% !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  backdrop-filter: blur(10px) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.2) !important;
  animation: superPulse 3s ease-in-out infinite !important;
}

.title-container {
  flex: 1 !important;
}

.super-dialog-title {
  color: #4b5563 !important;
  font-size: 2.2rem !important;
  font-weight: 800 !important;
  margin: 0 0 0.5rem 0 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 3px 10px rgba(0, 0, 0, 0.4), 0 0 0 2px rgba(255, 255, 255, 0.1) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.2 !important;
}

.super-dialog-subtitle {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 1.1rem !important;
  margin: 0 !important;
  font-weight: 500 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
}

.super-close-btn {
  background: rgba(255, 255, 255, 0.15) !important;
  border-radius: 50% !important;
  width: 50px !important;
  height: 50px !important;
  backdrop-filter: blur(10px) !important;
  border: 2px solid rgba(255, 255, 255, 0.2) !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2) !important;
}

.super-close-btn:hover {
  background: rgba(255, 255, 255, 0.25) !important;
  transform: scale(1.1) rotate(90deg) !important;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.3) !important;
  border-color: rgba(255, 255, 255, 0.4) !important;
}

/* المحتوى الجديد */
.super-dialog-body {
  padding: 2rem !important;
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%) !important;
}

.super-form-container {
  display: flex !important;
  flex-direction: column !important;
  gap: 1.5rem !important;
}

.super-form-section {
  background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%) !important;
  border-radius: 16px !important;
  padding: 1.5rem !important;
  border: 2px solid rgba(25, 118, 210, 0.1) !important;
  box-shadow: 0 4px 20px rgba(25, 118, 210, 0.1) !important;
  transition: all 0.3s ease !important;
  position: relative !important;
  overflow: hidden !important;
}

.super-form-section:hover {
  box-shadow: 0 8px 30px rgba(25, 118, 210, 0.15) !important;
  transform: translateY(-2px) !important;
  border-color: rgba(25, 118, 210, 0.2) !important;
}

.field-group {
  display: flex !important;
  flex-direction: column !important;
  gap: 0.8rem !important;
  margin-bottom: 1.5rem !important;
}

.field-group:last-child {
  margin-bottom: 0 !important;
}

.super-field-label {
  font-size: 1.1rem !important;
  font-weight: 700 !important;
  color: #1976d2 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(25, 118, 210, 0.1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.3 !important;
  display: flex !important;
  align-items: center !important;
  gap: 0.5rem !important;
  margin-bottom: 0.5rem !important;
}

/* الحقول الجديدة مع وضوح فائق */
.super-modern-input .v-field {
  border-radius: 12px !important;
  border: 2px solid rgba(25, 118, 210, 0.2) !important;
  transition: all 0.3s ease !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.9) 0%, rgba(248, 250, 252, 0.9) 100%) !important;
  min-height: 52px !important;
  box-shadow: 0 4px 15px rgba(25, 118, 210, 0.1) !important;
  position: relative !important;
  backdrop-filter: blur(10px) !important;
}

.super-modern-input .v-field:hover {
  border-color: #1976d2 !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(248, 250, 252, 1) 100%) !important;
  box-shadow: 0 8px 25px rgba(25, 118, 210, 0.2), 0 0 0 2px rgba(25, 118, 210, 0.3) !important;
  transform: translateY(-2px) !important;
}

.super-modern-input .v-field--focused {
  border-color: #1976d2 !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 1) 0%, rgba(248, 250, 252, 1) 100%) !important;
  box-shadow: 0 12px 35px rgba(25, 118, 210, 0.3), 0 0 0 3px rgba(25, 118, 210, 0.4) !important;
  transform: translateY(-3px) scale(1.02) !important;
}

/* النص فائق الوضوح */
.super-modern-input .v-field__input {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  padding: 1rem 1.5rem !important;
  line-height: 1.4 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: transparent !important;
  border-radius: 10px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  letter-spacing: 0.5px !important;
}

.super-modern-input .v-field__input::placeholder {
  color: rgba(25, 118, 210, 0.5) !important;
  font-weight: 600 !important;
  opacity: 0.8 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(25, 118, 210, 0.1) !important;
}

/* الأزرار الجديدة */
.super-dialog-footer {
  padding: 2rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%) !important;
  border-top: 2px solid rgba(25, 118, 210, 0.1) !important;
  position: relative !important;
  overflow: hidden !important;
  backdrop-filter: blur(10px) !important;
}

.footer-pattern-bg {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: radial-gradient(circle at 50% 0%, rgba(25, 118, 210, 0.05) 0%, transparent 60%) !important;
  pointer-events: none !important;
}

.footer-main {
  display: flex !important;
  align-items: center !important;
  justify-content: flex-end !important;
  gap: 1rem !important;
  position: relative !important;
  z-index: 1 !important;
}

.super-cancel-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 0.8rem 1.5rem !important;
  border: 2px solid rgba(156, 163, 175, 0.3) !important;
  color: #6b7280 !important;
  transition: all 0.3s ease !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: rgba(255, 255, 255, 0.9) !important;
  backdrop-filter: blur(10px) !important;
  box-shadow: 0 4px 15px rgba(156, 163, 175, 0.1) !important;
  letter-spacing: 0.3px !important;
  min-width: 120px !important;
}

.super-cancel-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: #9ca3af !important;
  color: #374151 !important;
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 8px 25px rgba(156, 163, 175, 0.2) !important;
}

.super-save-btn {
  font-weight: 800 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 0.8rem 2rem !important;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 50%, #0d47a1 100%) !important;
  background-size: 200% 200% !important;
  animation: superGradientShift 3s ease infinite !important;
  box-shadow: 0 8px 25px rgba(25, 118, 210, 0.4) !important;
  font-size: 1.1rem !important;
  transition: all 0.3s ease !important;
  color: #4b5563 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.2) !important;
  backdrop-filter: blur(10px) !important;
  min-width: 160px !important;
}

.super-save-btn:hover {
  background: linear-gradient(135deg, #0d47a1 0%, #1976d2 50%, #1565c0 100%) !important;
  transform: translateY(-3px) scale(1.08) !important;
  box-shadow: 0 15px 40px rgba(25, 118, 210, 0.5) !important;
  border-color: rgba(255, 255, 255, 0.3) !important;
}

.super-save-btn:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #d1d5db 100%) !important;
  color: #6b7280 !important;
  transform: none !important;
  box-shadow: 0 4px 15px rgba(156, 163, 175, 0.2) !important;
  border-color: rgba(156, 163, 175, 0.3) !important;
}

/* الرسوم المتحركة الجديدة */
@keyframes superGradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

@keyframes superPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

/* CSS قوي جداً لضمان ظهور النص المكتوب */
.super-modern-dialog input,
.super-modern-dialog textarea,
.super-modern-dialog .v-field__input,
.super-modern-dialog .v-field__input input,
.super-modern-dialog .v-field__input textarea,
.super-modern-dialog .v-input__control input,
.super-modern-dialog .v-input__control textarea {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
}

/* CSS إضافي لجميع عناصر الإدخال */
.super-modern-dialog input[type="text"],
.super-modern-dialog input[type="date"],
.super-modern-dialog input[type="number"],
.super-modern-dialog textarea {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
}

/* CSS للعناصر المباشرة */
.super-modern-dialog .v-field__input input,
.super-modern-dialog .v-field__input textarea {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
}

/* CSS باستخدام ::deep لضمان التطبيق */
::deep(.super-modern-input input),
::deep(.super-modern-input textarea),
::deep(.super-modern-input .v-field__input),
::deep(.v-text-field input),
::deep(.v-textarea textarea) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
}

/* CSS إضافي لجميع عناصر Vuetify */
::deep(.v-field__input),
::deep(.v-field__input input),
::deep(.v-field__input textarea),
::deep(.v-input__control input),
::deep(.v-input__control textarea) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
}

/* ========================================
   تصميم مشابه للصورة - إضافة مشروع
   ======================================== */

/* القائمة المنبثقة الجديدة */
.image-style-dialog {
  border-radius: 8px !important;
  overflow: hidden !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15) !important;
  background: #ffffff !important;
  position: relative !important;
}

/* الرأس الأزرق */
.dialog-header {
  background: #1976d2 !important;
  padding: 1rem 1.5rem !important;
  position: relative !important;
}

.header-content {
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
}

.header-left {
  display: flex !important;
  align-items: center !important;
  gap: 0.75rem !important;
}

.header-icon {
  color: #1976d2 !important;
}

.header-title {
  color: #1976d2 !important;
  font-size: 1.1rem !important;
  font-weight: 700 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(255, 255, 255, 0.8) !important;
}

.close-btn {
  color: #4b5563 !important;
  min-width: 32px !important;
  width: 32px !important;
  height: 32px !important;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1) !important;
}

/* محتوى النموذج */
.dialog-body {
  padding: 1.5rem !important;
  background: #ffffff !important;
}

.form-fields {
  display: flex !important;
  flex-direction: column !important;
  gap: 1rem !important;
}

.field-group {
  display: flex !important;
  flex-direction: column !important;
  gap: 0.5rem !important;
}

.field-label {
  font-size: 1rem !important;
  font-weight: 700 !important;
  color: #000000 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  margin-bottom: 0.25rem !important;
}

/* الحقول المخصصة */
.custom-input .v-field {
  border-radius: 4px !important;
  border: 2px solid #000000 !important;
  background: #ffffff !important;
  min-height: 40px !important;
  transition: all 0.2s ease !important;
}

.custom-input .v-field:hover {
  border-color: #333333 !important;
  box-shadow: 0 0 0 1px #333333 !important;
}

.custom-input .v-field--focused {
  border-color: #000000 !important;
  box-shadow: 0 0 0 2px #000000 !important;
}

.custom-input .v-field__input {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 0.75rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.custom-input .v-field__input::placeholder {
  color: #666666 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* الحقل للقراءة فقط */
.readonly-field .v-field {
  background: #f0f0f0 !important;
  border-color: #000000 !important;
}

.readonly-field .v-field__input {
  color: #000000 !important;
  font-weight: 700 !important;
  cursor: not-allowed !important;
}

/* منطقة الأزرار */
.dialog-footer {
  padding: 1rem 1.5rem !important;
  background: #f9fafb !important;
  border-top: 1px solid #e5e7eb !important;
  display: flex !important;
  align-items: center !important;
  justify-content: flex-end !important;
  gap: 0.75rem !important;
}

.cancel-btn,
.cancel-btn.v-btn {
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 10px 20px !important;
  border: 2px solid rgba(156, 163, 175, 0.3) !important;
  color: #6b7280 !important;
  background: rgba(255, 255, 255, 0.95) !important;
  backdrop-filter: blur(10px) !important;
  font-size: 0.75rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: 100px !important;
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.08),
    0 2px 8px rgba(156, 163, 175, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.9) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
  position: relative;
  overflow: hidden;
}

.cancel-btn :deep(.v-btn__content) {
  color: #6b7280 !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  letter-spacing: 0.3px !important;
}

.cancel-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.cancel-btn:hover::before {
  left: 100%;
}

.cancel-btn:hover,
.cancel-btn.v-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(156, 163, 175, 0.5) !important;
  color: #374151 !important;
  transform: translateY(-2px) scale(1.02) !important;
  box-shadow: 
    0 6px 20px rgba(0, 0, 0, 0.12),
    0 3px 10px rgba(156, 163, 175, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
}

.save-btn,
.save-btn.v-btn {
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 10px 20px !important;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #4b5563 !important;
  font-size: 0.75rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: 140px !important;
  box-shadow: 
    0 4px 16px rgba(59, 130, 246, 0.4),
    0 2px 8px rgba(37, 99, 235, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
  position: relative;
  overflow: hidden;
}

.save-btn :deep(.v-btn__content) {
  color: #4b5563 !important;
  font-weight: 600 !important;
  font-size: 0.75rem !important;
  text-align: center !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  letter-spacing: 0.3px !important;
  line-height: 1.4 !important;
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

.save-btn:hover,
.save-btn.v-btn:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  box-shadow: 
    0 8px 24px rgba(59, 130, 246, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  transform: translateY(-2px) scale(1.02) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.save-btn:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #6b7280 100%) !important;
  color: #4b5563 !important;
  box-shadow: 0 2px 8px rgba(156, 163, 175, 0.2) !important;
  transform: none !important;
  opacity: 0.6 !important;
  cursor: not-allowed !important;
}

.save-btn:disabled::before {
  display: none !important;
}

/* التأكد من تطبيق اللون الأزرق الفاتح حتى مع color="primary" */
.save-btn.v-btn--variant-elevated {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #4b5563 !important;
}

.save-btn.v-btn--variant-elevated:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
}

/* CSS قوي لضمان ظهور النص المكتوب */
.image-style-dialog input,
.image-style-dialog textarea,
.image-style-dialog .v-field__input,
.image-style-dialog .v-field__input input,
.image-style-dialog .v-field__input textarea {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS إضافي لجميع عناصر الإدخال */
.image-style-dialog input[type="text"],
.image-style-dialog input[type="number"],
.image-style-dialog textarea {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS باستخدام ::deep لضمان التطبيق */
::deep(.custom-input input),
::deep(.custom-input textarea),
::deep(.custom-input .v-field__input) {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS قوي جداً لضمان وضوح النص المكتوب */
.image-style-dialog .v-field__input input,
.image-style-dialog .v-field__input textarea,
.image-style-dialog .v-input__control input,
.image-style-dialog .v-input__control textarea,
.image-style-dialog .v-text-field input,
.image-style-dialog .v-textarea textarea {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS إضافي لجميع عناصر Vuetify */
::deep(.v-field__input),
::deep(.v-field__input input),
::deep(.v-field__input textarea),
::deep(.v-input__control input),
::deep(.v-input__control textarea),
::deep(.v-text-field input),
::deep(.v-textarea textarea) {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS شامل لجميع النصوص في القائمة المنبثقة */
.image-style-dialog * {
  color: #000000 !important;
}

.image-style-dialog .v-field__input,
.image-style-dialog .v-field__input *,
.image-style-dialog input,
.image-style-dialog textarea,
.image-style-dialog label,
.image-style-dialog .field-label,
.image-style-dialog .v-label,
.image-style-dialog .v-field__outline,
.image-style-dialog .v-field__outline * {
  color: #000000 !important;
  font-weight: 700 !important;
}

/* CSS قوي جداً لضمان السواد الكامل */
.image-style-dialog .v-field__input input,
.image-style-dialog .v-field__input textarea,
.image-style-dialog .v-input__control input,
.image-style-dialog .v-input__control textarea,
.image-style-dialog .v-text-field input,
.image-style-dialog .v-textarea textarea,
.image-style-dialog .v-field__input,
.image-style-dialog .v-field__input *,
.image-style-dialog input[type="text"],
.image-style-dialog input[type="number"],
.image-style-dialog textarea {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* منطقة الأزرار */
.dialog-footer {
  padding: 2rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%) !important;
  border-top: 3px solid rgba(25, 118, 210, 0.2) !important;
  display: flex !important;
  align-items: center !important;
  justify-content: flex-end !important;
  gap: 1rem !important;
  border-radius: 0 0 24px 24px !important;
  box-shadow: 0 -4px 20px rgba(25, 118, 210, 0.1) !important;
  backdrop-filter: blur(10px) !important;
}

.cancel-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 0.75rem 1.5rem !important;
  border: 2px solid rgba(156, 163, 175, 0.3) !important;
  color: #6b7280 !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: rgba(255, 255, 255, 0.8) !important;
  backdrop-filter: blur(10px) !important;
  box-shadow: 0 4px 15px rgba(156, 163, 175, 0.1) !important;
  letter-spacing: 0.3px !important;
}

.cancel-btn:hover {
  background: rgba(239, 68, 68, 0.1) !important;
  border-color: #ef4444 !important;
  color: #dc2626 !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 25px rgba(239, 68, 68, 0.2) !important;
}

.cancel-btn:hover {
  background: #f5f5f5 !important;
  border-color: #bdbdbd !important;
  color: #424242 !important;
  transform: translateY(-1px) !important;
}

.save-btn {
  font-weight: 800 !important;
  text-transform: none !important;
  border-radius: 12px !important;
  padding: 0.75rem 2rem !important;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  background-size: 200% 200% !important;
  animation: gradientShift 3s ease infinite !important;
  box-shadow: 0 8px 25px rgba(59, 130, 246, 0.4) !important;
  font-size: 1.1rem !important;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275) !important;
  color: #4b5563 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.2) !important;
  backdrop-filter: blur(10px) !important;
}

.save-btn:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  box-shadow: 0 12px 35px rgba(59, 130, 246, 0.5) !important;
  transform: translateY(-3px) scale(1.05) !important;
  border-color: rgba(255, 255, 255, 0.4) !important;
}

.save-btn:disabled {
  background: linear-gradient(135deg, #9ca3af 0%, #6b7280 100%) !important;
  color: #4b5563 !important;
  box-shadow: 0 4px 15px rgba(156, 163, 175, 0.2) !important;
  transform: none !important;
  animation: none !important;
  border-color: rgba(255, 255, 255, 0.1) !important;
}

.modern-dialog-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  padding: 1.5rem !important;
  position: relative !important;
}

.modern-dialog-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grain" width="100" height="100" patternUnits="userSpaceOnUse"><circle cx="25" cy="25" r="1" fill="white" opacity="0.1"/><circle cx="75" cy="75" r="1" fill="white" opacity="0.1"/><circle cx="50" cy="10" r="0.5" fill="white" opacity="0.05"/><circle cx="10" cy="60" r="0.5" fill="white" opacity="0.05"/></pattern></defs><rect width="100" height="100" fill="url(%23grain)"/></svg>');
  opacity: 0.3;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 1rem;
  position: relative;
  z-index: 2;
}

.header-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.dialog-main-title {
  font-size: 1.25rem !important;
  font-weight: 800 !important;
  color: white !important;
  margin: 0 !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
}

.dialog-subtitle {
  font-size: 0.9rem !important;
  color: rgba(255, 255, 255, 0.9) !important;
  margin: 0.25rem 0 0 0 !important;
  font-weight: 500 !important;
}

.project-dialog-header {
  background: linear-gradient(135deg, #f1f5f9 0%, #e2e8f0 100%) !important;
  padding: 1.5rem 2rem !important;
  border-bottom: 1px solid rgba(0, 0, 0, 0.1) !important;
}

.dialog-title {
  font-size: 1.25rem !important;
  font-weight: 700 !important;
  color: #1e293b !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
}

.project-dialog-body {
  padding: 2rem !important;
  background: #ffffff !important;
}

.compact-dialog-body {
  padding: 2rem !important;
  background: linear-gradient(135deg, #fef3c7 0%, #fde68a 100%) !important;
  border-radius: 0 !important;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.1) !important;
  border: 2px solid #e5e7eb !important;
  border-top: none !important;
  border-bottom: none !important;
}

.form-sections {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-section {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 12px;
  padding: 1.25rem;
  border: 1px solid #e2e8f0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  position: relative;
  overflow: hidden;
}

.form-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #667eea, #764ba2);
  border-radius: 16px 16px 0 0;
}

.section-title {
  font-size: 1rem !important;
  font-weight: 700 !important;
  color: #1e293b !important;
  margin-bottom: 1rem !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.section-title::before {
  content: '';
  width: 4px;
  height: 20px;
  background: linear-gradient(135deg, #3b82f6, #1d4ed8);
  border-radius: 2px;
}

.form-field-container {
  margin-bottom: 1rem;
}

.form-section .form-field-container:last-child {
  margin-bottom: 0;
}

/* Compact Field Styles */
.compact-field-label {
  display: block;
  font-size: 0.8rem !important;
  font-weight: 600 !important;
  color: #374151 !important;
  margin-bottom: 0.5rem !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
}

.compact-form-field .v-field {
  border-radius: 8px !important;
  border: 1px solid #d1d5db !important;
  transition: all 0.2s ease !important;
  background: #ffffff !important;
  min-height: 40px !important;
}

.compact-form-field .v-field--focused {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.1) !important;
}

.compact-form-field .v-field__input {
  color: #111827 !important;
  font-weight: 500 !important;
  font-size: 0.9rem !important;
  padding: 0.5rem 0.75rem !important;
}

.field-label {
  display: block;
  font-size: 0.95rem !important;
  font-weight: 700 !important;
  color: #1f2937 !important;
  margin-bottom: 0.75rem !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.05) !important;
}

.project-form-field .v-field {
  border-radius: 10px !important;
  border: 2px solid #d1d5db !important;
  transition: all 0.3s ease !important;
  background: #ffffff !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05) !important;
  min-height: 56px !important;
}

.project-form-field .v-field:hover {
  border-color: #9ca3af !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1) !important;
}

.project-form-field .v-field--focused {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1) !important;
  background: #ffffff !important;
}

.project-form-field .v-field__input {
  color: #111827 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 0.75rem 1rem !important;
}

.project-form-field .v-field__outline {
  display: none !important;
}

.project-form-field .v-field__overlay {
  background: transparent !important;
}

.project-form-field .v-field__append-inner {
  color: #6b7280 !important;
}

.project-form-field .v-select .v-field__input {
  color: #111827 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.project-form-field .v-field--active .v-field__input {
  color: #111827 !important;
  opacity: 1 !important;
}

.project-form-field .v-field--focused .v-field__input {
  color: #111827 !important;
  opacity: 1 !important;
}

.project-dialog-actions {
  padding: 1rem 2rem !important;
  background: #f8fafc !important;
  border-top: 1px solid rgba(0, 0, 0, 0.1) !important;
}

.cancel-btn {
  font-weight: 600 !important;
  text-transform: none !important;
}

.save-btn {
  font-weight: 600 !important;
  text-transform: none !important;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 100%) !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  color: #4b5563 !important;
  transition: all 0.3s ease !important;
}

.save-btn:hover {
  transform: translateY(-1px) !important;
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4) !important;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
}

/* Modern Dialog Actions */
.modern-dialog-actions {
  padding: 1rem 1.5rem !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-top: 1px solid rgba(0, 0, 0, 0.08) !important;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
}

.actions-info {
  display: flex;
  align-items: center;
}

.actions-buttons {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.modern-cancel-btn {
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 8px !important;
  padding: 0 1rem !important;
  min-width: 100px !important;
  font-size: 0.9rem !important;
}

.modern-save-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 8px !important;
  padding: 0 1.25rem !important;
  min-width: 140px !important;
  font-size: 0.9rem !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.4) !important;
}

.modern-save-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.5) !important;
}

.modern-save-btn:disabled {
  background: #e2e8f0 !important;
  color: #9ca3af !important;
  box-shadow: none !important;
  transform: none !important;
}

/* Ultra Modern Styles */
.ultra-modern-header {
  background: #1e3a8a !important;
  padding: 1.25rem 1.75rem !important;
  position: relative !important;
  border-radius: 0 !important;
  box-shadow: 0 2px 8px rgba(30, 58, 138, 0.3) !important;
}

.header-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.ultra-modern-title {
  font-size: 1.4rem !important;
  font-weight: 800 !important;
  color: white !important;
  margin: 0 !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
}

.close-header-btn {
  opacity: 0.9 !important;
  transition: all 0.2s ease !important;
  background: rgba(255, 255, 255, 0.15) !important;
  border-radius: 4px !important;
  padding: 0.5rem !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
}

.close-header-btn:hover {
  opacity: 1 !important;
  background: rgba(255, 255, 255, 0.25) !important;
  border-color: rgba(255, 255, 255, 0.4) !important;
  transform: scale(1.1) !important;
}

.ultra-modern-form {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  padding: 0.5rem;
}

.ultra-form-row {
  display: flex;
  gap: 1.5rem;
  align-items: flex-start;
}

.ultra-form-field {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  flex: 1;
}

.ultra-field-label {
  font-size: 1.1rem !important;
  font-weight: 700 !important;
  color: #1f2937 !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  margin-bottom: 0.5rem !important;
  padding: 0.25rem 0 !important;
  border-bottom: 2px solid rgba(30, 58, 138, 0.1) !important;
}

.ultra-input-field .v-field {
  border-radius: 8px !important;
  border: 4px solid #1e3a8a !important;
  transition: all 0.3s ease !important;
  background: linear-gradient(135deg, #ffffff 0%, #fef3c7 100%) !important;
  min-height: 56px !important;
  box-shadow: 0 4px 12px rgba(30, 58, 138, 0.4) !important;
  position: relative !important;
  margin-bottom: 0.5rem !important;
  border-style: solid !important;
  border-width: 4px !important;
}

.ultra-input-field .v-field:hover {
  border-color: #1d4ed8 !important;
  background: linear-gradient(135deg, #ffffff 0%, #fde68a 100%) !important;
  box-shadow: 0 6px 16px rgba(30, 64, 175, 0.4) !important;
  transform: translateY(-2px) !important;
  border-width: 4px !important;
  border-style: solid !important;
}

.ultra-input-field .v-field--focused {
  border-color: #1e40af !important;
  background: linear-gradient(135deg, #ffffff 0%, #fef3c7 100%) !important;
  box-shadow: 0 0 0 6px rgba(30, 64, 175, 0.3) !important;
  transform: translateY(-2px) !important;
  border-width: 4px !important;
  border-style: solid !important;
}

.ultra-input-field .v-field__input {
  color: #1e3a8a !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 1rem 1.5rem !important;
  line-height: 1.7 !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  background: #ffffff !important;
  border-radius: 6px !important;
  opacity: 1 !important;
  z-index: 1 !important;
  text-shadow: none !important;
  letter-spacing: 0.3px !important;
  transition: all 0.3s ease !important;
}

.ultra-input-field .v-field__input::placeholder {
  color: #9ca3af !important;
  font-weight: 400 !important;
  opacity: 0.8 !important;
  font-style: italic !important;
  letter-spacing: 0.2px !important;
  text-shadow: none !important;
}

.ultra-input-field .v-field--active .v-field__input {
  color: #1e40af !important;
  opacity: 1 !important;
  background: #ffffff !important;
  text-shadow: none !important;
}

.ultra-input-field .v-field--focused .v-field__input {
  color: #1e40af !important;
  font-weight: 700 !important;
  opacity: 1 !important;
  background: #ffffff !important;
  text-shadow: none !important;
  transform: none !important;
}

.ultra-modern-actions {
  padding: 1.5rem 2rem !important;
  background: linear-gradient(135deg, #fde68a 0%, #f59e0b 100%) !important;
  border-top: 2px solid #e5e7eb !important;
  border-left: 2px solid #e5e7eb !important;
  border-right: 2px solid #e5e7eb !important;
  border-bottom: 2px solid #e5e7eb !important;
  display: flex !important;
  justify-content: flex-end !important;
  border-radius: 0 0 8px 8px !important;
}

.ultra-save-btn {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 6px !important;
  padding: 0.75rem 2rem !important;
  min-width: 100px !important;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  font-size: 1rem !important;
  letter-spacing: 0.5px !important;
}

.ultra-save-btn:hover {
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%) !important;
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4) !important;
  transform: translateY(-1px) !important;
}

/* Special styling for cost field */
.ultra-input-field .v-field__append-inner {
  color: #374151 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

/* Textarea specific styling */
.ultra-input-field textarea {
  color: #1e3a8a !important;
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  line-height: 1.8 !important;
  resize: vertical !important;
  min-height: 100px !important;
  padding: 1rem 1.5rem !important;
  font-family: 'Inter', 'Arial', sans-serif !important;
  border-radius: 6px !important;
  text-shadow: none !important;
  letter-spacing: 0.3px !important;
  transition: all 0.3s ease !important;
  background: #ffffff !important;
}

.ultra-input-field textarea::placeholder {
  color: #9ca3af !important;
  font-weight: 400 !important;
  opacity: 0.8 !important;
  font-style: italic !important;
  letter-spacing: 0.2px !important;
  text-shadow: none !important;
}

/* Calendar icon styling */
.ultra-input-field .v-field__prepend-inner {
  color: #1e3a8a !important;
  font-size: 1.2rem !important;
  text-shadow: 0 1px 2px rgba(30, 58, 138, 0.2) !important;
  transition: all 0.3s ease !important;
}

.ultra-input-field .v-field__prepend-inner:hover {
  color: #1e40af !important;
  transform: scale(1.1) !important;
}

/* Currency suffix styling */
.ultra-input-field .v-field__append-inner {
  color: #1e3a8a !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  padding-left: 0.5rem !important;
  text-shadow: 0 1px 2px rgba(30, 58, 138, 0.1) !important;
  letter-spacing: 0.5px !important;
}

/* Inner field styling */
.ultra-input-field .v-field__field {
  padding: 0 !important;
  border: none !important;
  border-radius: 6px !important;
  margin: 3px !important;
  background: transparent !important;
}

/* Force all input fields to have consistent borders */
.ultra-input-field .v-field:not(.v-field--error) {
  border: 4px solid #1e3a8a !important;
  border-style: solid !important;
  border-width: 4px !important;
  box-shadow: 0 4px 12px rgba(30, 58, 138, 0.4) !important;
}

/* Ensure all fields have the same border visibility */
.ultra-input-field .v-field,
.ultra-input-field .v-text-field .v-field,
.ultra-input-field .v-select .v-field,
.ultra-input-field .v-textarea .v-field {
  border: 4px solid #1e3a8a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(30, 58, 138, 0.4) !important;
  min-height: 56px !important;
}

.ultra-input-field .v-field__outline {
  display: none !important;
}

.ultra-input-field .v-field__overlay {
  display: none !important;
}

/* Ensure text visibility - Beautiful Navy Blue Theme */
.ultra-input-field input,
.ultra-input-field textarea {
  color: #1e3a8a !important;
  font-weight: 700 !important;
  background: #ffffff !important;
  border: none !important;
  outline: none !important;
  opacity: 1 !important;
  z-index: 2 !important;
  text-shadow: none !important;
  letter-spacing: 0.3px !important;
  font-size: 1.1rem !important;
}

.ultra-input-field input:focus,
.ultra-input-field textarea:focus {
  color: #1e40af !important;
  background: #ffffff !important;
  opacity: 1 !important;
  text-shadow: none !important;
  transform: none !important;
}

/* Error state styling - Changed to Navy Blue */
.ultra-input-field .v-field--error {
  border: 3px solid #1e3a8a !important;
  background: #f0f4ff !important;
  border-style: solid !important;
  border-width: 3px !important;
  box-shadow: 0 3px 8px rgba(30, 58, 138, 0.3) !important;
}

.ultra-input-field .v-field--error .v-field__input {
  color: #1e3a8a !important;
}

.ultra-input-field .v-field--error:hover {
  border: 3px solid #1d4ed8 !important;
  background: #e0e7ff !important;
  border-style: solid !important;
  border-width: 3px !important;
}

.ultra-input-field .v-field--error.v-field--focused {
  border: 3px solid #1e3a8a !important;
  background: #ffffff !important;
  box-shadow: 0 0 0 2px rgba(30, 58, 138, 0.2) !important;
  border-style: solid !important;
  border-width: 3px !important;
}

/* Error messages - Changed to Navy Blue */
.ultra-input-field .v-messages {
  margin-top: 0.25rem !important;
  min-height: 1.25rem !important;
}

.ultra-input-field .v-messages__message {
  color: #1e3a8a !important;
  font-size: 0.75rem !important;
  font-weight: 600 !important;
  line-height: 1.25rem !important;
  display: flex !important;
  align-items: center !important;
}

.ultra-input-field .v-messages__message::before {
  content: "ℹ" !important;
  margin-left: 0.25rem !important;
  font-size: 0.7rem !important;
  color: #1e3a8a !important;
}

/* Dialog Container Enhancements */

/* Administrative Expenses Table Styles - منسق */
.expense-table {
  margin-top: 20px !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 250, 252, 0.95) 100%) !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  box-shadow: 0 12px 40px rgba(67, 56, 202, 0.15) !important;
  backdrop-filter: blur(15px) !important;
  border: 2px solid rgba(199, 210, 254, 0.3) !important;
}

.expense-table .v-data-table__wrapper {
  border-radius: 16px;
  overflow: hidden;
  background: transparent;
}

.expense-table .v-data-table__th {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 50%, #8b5cf6 100%) !important;
  color: #4b5563 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-align: center !important;
  padding: 24px 20px !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.4) !important;
  border-bottom: 3px solid rgba(255, 255, 255, 0.3) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.8px !important;
  text-transform: uppercase !important;
  position: relative !important;
  z-index: 2 !important;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border-right: 1px solid rgba(255, 255, 255, 0.1) !important;
}

.expense-table .v-data-table__th:last-child {
  border-right: none !important;
}

.expense-table .v-data-table__th::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0.05) 100%);
  pointer-events: none;
  z-index: 1;
}

.expense-table .v-data-table__th > * {
  position: relative;
  z-index: 3;
}

/* Enhanced Header Visibility */
.expense-table .v-data-table__th {
  background-attachment: fixed !important;
  background-size: 200% 200% !important;
  animation: headerShimmer 3s ease-in-out infinite !important;
}

@keyframes headerShimmer {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

/* Header Text Enhancement */
.expense-table .v-data-table__th .v-data-table-header__content {
  color: #4b5563 !important;
  font-weight: 900 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5) !important;
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.3)) !important;
}

/* Header Hover Effect */
.expense-table .v-data-table__th:hover {
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #3b82f6 100%) !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2) !important;
}

/* Additional Header Clarity */
.expense-table .v-data-table__th {
  -webkit-font-smoothing: antialiased !important;
  -moz-osx-font-smoothing: grayscale !important;
  font-smooth: always !important;
  text-rendering: optimizeLegibility !important;
}

/* Header Text Contrast */
.expense-table .v-data-table__th .v-data-table-header__content {
  color: #4b5563 !important;
  font-weight: 900 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5) !important;
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.3)) !important;
  -webkit-text-stroke: 0.5px rgba(0, 0, 0, 0.3) !important;
}

.expense-table .v-data-table__td {
  padding: 20px 16px !important;
  border-bottom: 1px solid rgba(199, 210, 254, 0.3) !important;
  text-align: center !important;
  background: rgba(255, 255, 255, 0.9) !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 500 !important;
  font-size: 1rem !important;
  color: #374151 !important;
}

.expense-table .v-data-table__tr:hover {
  background: linear-gradient(135deg, rgba(67, 56, 202, 0.08) 0%, rgba(99, 102, 241, 0.05) 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(67, 56, 202, 0.2) !important;
  border-radius: 12px !important;
}

.expense-table .v-data-table__tr:nth-child(even) {
  background: linear-gradient(135deg, rgba(238, 242, 255, 0.8) 0%, rgba(224, 231, 255, 0.6) 100%) !important;
}

.expense-table .v-data-table__tr:nth-child(odd) {
  background: rgba(255, 255, 255, 0.95) !important;
}

/* Expense Table Card */
.data-table-card.mt-6 {
  margin-top: 24px !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
  border-radius: 20px !important;
  box-shadow: 0 12px 40px rgba(59, 130, 246, 0.15) !important;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(248, 250, 252, 0.9) 100%) !important;
  backdrop-filter: blur(15px);
  position: relative;
  overflow: hidden;
}

.data-table-card.mt-6::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #3b82f6 0%, #1d4ed8 50%, #60a5fa 100%);
  border-radius: 20px 20px 0 0;
}

/* Expense Table Footer - منسق */
.expense-table .v-data-table-footer {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 50%, #8b5cf6 100%) !important;
  color: white !important;
  border-radius: 0 0 20px 20px !important;
  padding: 20px !important;
  box-shadow: 0 -4px 20px rgba(67, 56, 202, 0.3) !important;
}

.expense-table .v-data-table-footer .v-pagination {
  color: white !important;
}

.expense-table .v-data-table-footer .v-pagination .v-btn {
  color: white !important;
  background: rgba(255, 255, 255, 0.2) !important;
}

.expense-table .v-data-table-footer .v-pagination .v-btn:hover {
  background: rgba(255, 255, 255, 0.3) !important;
}

.expense-table .v-data-table-footer .v-pagination .v-btn--active {
  background: rgba(255, 255, 255, 0.4) !important;
  color: white !important;
}


/* Modern Filter Sections - نظام نيلي موحد */
.modern-filter-section {
  background: linear-gradient(135deg, rgba(238, 242, 255, 0.8) 0%, rgba(224, 231, 255, 0.9) 50%, rgba(199, 210, 254, 0.8) 100%) !important;
  border-radius: 20px !important;
  margin: 20px !important;
  padding: 32px !important;
  border: 1px solid rgba(199, 210, 254, 0.4) !important;
  box-shadow: 0 8px 32px rgba(67, 56, 202, 0.15), 0 4px 16px rgba(0, 0, 0, 0.08) !important;
  backdrop-filter: blur(20px) !important;
  position: relative !important;
  overflow: hidden !important;
}

.modern-expense-filter-section {
  background: linear-gradient(135deg, rgba(67, 56, 202, 0.12) 0%, rgba(79, 70, 229, 0.18) 50%, rgba(99, 102, 241, 0.12) 100%) !important;
  border: 1px solid rgba(67, 56, 202, 0.25) !important;
  box-shadow: 0 8px 32px rgba(67, 56, 202, 0.15), 0 4px 16px rgba(0, 0, 0, 0.08) !important;
}

.filter-header {
  text-align: center !important;
  margin-bottom: 32px !important;
  position: relative !important;
}

.filter-icon-container {
  position: relative !important;
  display: inline-block !important;
  margin-bottom: 16px !important;
}

.filter-icon {
  position: relative !important;
  z-index: 2 !important;
  animation: modernPulse 2s ease-in-out infinite !important;
}

.icon-glow {
  position: absolute !important;
  top: 50% !important;
  left: 50% !important;
  transform: translate(-50%, -50%) !important;
  width: 60px !important;
  height: 60px !important;
  background: radial-gradient(circle, rgba(67, 56, 202, 0.3) 0%, rgba(99, 102, 241, 0.2) 50%, transparent 70%) !important;
  border-radius: 50% !important;
  animation: modernGlow 2s ease-in-out infinite !important;
}

.expense-glow {
  background: radial-gradient(circle, rgba(67, 56, 202, 0.3) 0%, transparent 70%) !important;
}

.filter-title {
  font-size: 1.6rem !important;
  font-weight: 900 !important;
  color: #000000 !important;
  margin-bottom: 12px !important;
  letter-spacing: 0.8px !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.2), 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  line-height: 1.3 !important;
  position: relative !important;
  z-index: 2 !important;
}

.expense-title {
  font-size: 1.6rem !important;
  font-weight: 900 !important;
  color: #000000 !important;
  margin-bottom: 12px !important;
  letter-spacing: 0.8px !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.2), 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  line-height: 1.3 !important;
  position: relative !important;
  z-index: 2 !important;
}

.filter-subtitle {
  font-size: 1rem !important;
  font-weight: 500 !important;
  color: #000000 !important;
  margin-bottom: 20px !important;
  opacity: 0.9 !important;
}

.filter-decoration {
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
  gap: 12px !important;
  margin-top: 16px !important;
}

.decoration-line {
  width: 40px !important;
  height: 2px !important;
  background: linear-gradient(90deg, transparent 0%, #6366f1 30%, #4338ca 50%, #6366f1 70%, transparent 100%) !important;
  border-radius: 1px !important;
  animation: modernShimmer 2s ease-in-out infinite !important;
  box-shadow: 0 2px 4px rgba(67, 56, 202, 0.2) !important;
}

.expense-line {
  background: linear-gradient(90deg, transparent 0%, #4338ca 50%, transparent 100%) !important;
}

.decoration-dot {
  width: 8px !important;
  height: 8px !important;
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 100%) !important;
  border-radius: 50% !important;
  animation: modernBounce 2s ease-in-out infinite !important;
}

.expense-dot {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 100%) !important;
}

/* Modern Fields */
.modern-fields-row {
  margin-top: 24px !important;
}

.modern-search-container,
.modern-filters-container {
  position: relative !important;
}

.modern-search-container {
  margin-bottom: 16px !important;
}

.modern-filters-container {
  display: flex !important;
  gap: 16px !important;
  flex-wrap: wrap !important;
}

.filter-item {
  flex: 1 !important;
  min-width: 200px !important;
  position: relative !important;
}

.modern-search-field,
.modern-filter-field {
  background: rgba(255, 255, 255, 0.95) !important;
  border-radius: 16px !important;
  backdrop-filter: blur(10px) !important;
  border: 2px solid rgba(199, 210, 254, 0.4) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  box-shadow: 0 4px 16px rgba(67, 56, 202, 0.08) !important;
}

/* تغيير لون النص في حقل البحث الذكي */
.modern-search-field .v-field__input {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

.modern-search-field .v-label {
  color: #dc2626 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 3px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.3px !important;
  opacity: 1 !important;
}
الل
/* CSS أكثر تحديداً لأسماء الحقول */
.modern-search-field .v-field__outline .v-field__outline__start,
.modern-search-field .v-field__outline .v-field__outline__end,
.modern-search-field .v-field__outline .v-field__outline__notch {
  color: #dc2626 !important;
}

/* CSS إضافي لأسماء الحقول */
.modern-search-field .v-field-label--active,
.modern-search-field .v-field-label--floating {
  color: #dc2626 !important;
}

.modern-filter-field .v-field-label--active,
.modern-filter-field .v-field-label--floating {
  color: #dc2626 !important;
}

.modern-search-field .v-label.v-field-label {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

.modern-search-field .v-field-label {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

/* CSS باستخدام ::deep للتأكد من التطبيق - محاذاة في المنتصف */
::deep(.modern-search-field .v-label) {
  color: #dc2626 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 3px rgba(220, 38, 38, 0.3) !important;
  text-align: center !important;
  direction: rtl !important;
  letter-spacing: 0.3px !important;
}

::deep(.modern-search-field .v-field-label) {
  color: #dc2626 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

::deep(.modern-search-field .v-field__input) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

::deep(.modern-filter-field .v-field__input) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* تحسين وضوح حقول الفلترة */
.modern-filter-field .v-label {
  color: #dc2626 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 3px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.3px !important;
  opacity: 1 !important;
}

.modern-filter-field .v-field__input {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS أكثر تحديداً لحقول الفلترة */
.modern-filter-field .v-field__outline .v-field__outline__start,
.modern-filter-field .v-field__outline .v-field__outline__end,
.modern-filter-field .v-field__outline .v-field__outline__notch {
  color: #000000 !important;
}

.modern-filter-field .v-label.v-field-label {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

.modern-filter-field .v-field-label {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
}

/* CSS مع ::deep لحقول الفلترة - محاذاة في المنتصف */
::deep(.modern-filter-field .v-label) {
  color: #dc2626 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 1px 3px rgba(220, 38, 38, 0.3) !important;
  text-align: center !important;
  direction: rtl !important;
  letter-spacing: 0.3px !important;
}

::deep(.modern-filter-field .v-field-label) {
  color: #dc2626 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

/* تحسين النصوص المختارة في حقول الفلترة */
::deep(.modern-filter-field .v-field__input input) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-align: center !important;
  direction: rtl !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

::deep(.modern-search-field .v-field__input input) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

.modern-search-field:hover,
.modern-filter-field:hover {
  border-color: rgba(67, 56, 202, 0.5) !important;
  box-shadow: 0 8px 24px rgba(67, 56, 202, 0.2) !important;
  transform: translateY(-2px) !important;
}

.modern-search-field:focus-within,
.modern-filter-field:focus-within {
  border-color: #4338ca !important;
  box-shadow: 0 0 0 4px rgba(67, 56, 202, 0.2), 0 8px 24px rgba(67, 56, 202, 0.3) !important;
  transform: translateY(-2px) !important;
}

.search-glow,
.filter-glow {
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: linear-gradient(45deg, transparent 30%, rgba(99, 102, 241, 0.12) 50%, rgba(199, 210, 254, 0.08) 70%, transparent 90%) !important;
  border-radius: 16px !important;
  opacity: 0 !important;
  transition: opacity 0.3s ease !important;
  pointer-events: none !important;
}

.expense-search-glow,
.expense-filter-glow {
  background: linear-gradient(45deg, transparent 30%, rgba(99, 102, 241, 0.12) 50%, rgba(199, 210, 254, 0.08) 70%, transparent 90%) !important;
}

.modern-search-field:hover .search-glow,
.modern-filter-field:hover .filter-glow {
  opacity: 1 !important;
}

/* Animations */
@keyframes modernPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

@keyframes modernGlow {
  0%, 100% { opacity: 0.3; transform: translate(-50%, -50%) scale(1); }
  50% { opacity: 0.6; transform: translate(-50%, -50%) scale(1.2); }
}

@keyframes modernShimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

@keyframes modernBounce {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-8px); }
}

/* Search and Filter Fields - General */
.search-field .v-field__input,
.filter-field .v-field__input {
  color: #1e40af !important;
  font-weight: 700 !important;
  background: white !important;
  border-radius: 12px !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  border: 2px solid #e2e8f0 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

.search-field .v-field__input:focus,
.filter-field .v-field__input:focus {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1) !important;
}

.search-field .v-label,
.filter-field .v-label {
  color: #1e40af !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.search-field .v-field__outline,
.filter-field .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.search-field .v-field__outline--focused,
.filter-field .v-field__outline--focused {
  border-color: #1d4ed8 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2) !important;
}

/* Expense Search and Filter Fields */
.expense-table .search-field .v-field__input,
.expense-table .filter-field .v-field__input {
  color: #1e40af !important;
  font-weight: 700 !important;
  background: white !important;
  border-radius: 12px !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  border: 2px solid #e2e8f0 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

.expense-table .search-field .v-field__input:focus,
.expense-table .filter-field .v-field__input:focus {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1) !important;
}

.expense-table .search-field .v-label,
.expense-table .filter-field .v-label {
  color: #1e40af !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.expense-table .search-field .v-field__outline,
.expense-table .filter-field .v-field__outline {
  border-color: #3b82f6 !important;
  border-width: 2px !important;
}

.expense-table .search-field .v-field__outline--focused,
.expense-table .filter-field .v-field__outline--focused {
  border-color: #1d4ed8 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2) !important;
}

/* Expense Action Buttons */
.expense-table .v-btn--size-small {
  min-width: 36px !important;
  height: 36px !important;
  border-radius: 10px !important;
  transition: all 0.3s ease !important;
}

.expense-table .v-btn--variant-tonal {
  border-radius: 10px !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1) !important;
}

.expense-table .v-btn--variant-tonal:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

/* Expense Chips */
.expense-table .v-chip--size-small {
  font-weight: 700 !important;
  font-size: 0.8rem !important;
  border-radius: 12px !important;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1) !important;
  transition: all 0.3s ease !important;
}

.expense-table .v-chip--size-small:hover {
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15) !important;
}

/* Expense Cost Display */
.expense-table .text-error {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  color: #059669 !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

/* Expense Avatar */
.expense-table .v-avatar {
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
}

/* Expense Notes */
.expense-table .text-caption {
  color: #64748b !important;
  font-weight: 500 !important;
  font-style: italic !important;
}

/* Expense Chips Styling */
.expense-table .date-chip {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
}

.expense-table .location-chip {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
}

/* Expense Cost Display */
.expense-table .cost-display {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  color: white !important;
  padding: 8px 16px !important;
  border-radius: 12px !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
  display: inline-block !important;
  min-width: 100px !important;
  text-align: center !important;
}

/* Expense Action Buttons */
.expense-table .action-btn {
  min-width: 40px !important;
  height: 40px !important;
  border-radius: 12px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  transition: all 0.3s ease !important;
}

.expense-table .edit-btn {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
}

.expense-table .edit-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.4) !important;
}

.expense-table .delete-btn {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
}

.expense-table .delete-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(239, 68, 68, 0.4) !important;
}

/* Expense Table Responsive Design */
@media (max-width: 768px) {
  .expense-table .v-data-table__th {
    padding: 16px 12px !important;
    font-size: 1rem !important;
    font-weight: 800 !important;
    letter-spacing: 0.3px !important;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.6) !important;
    background: linear-gradient(135deg, #1e3a8a 0%, #3b82f6 100%) !important;
  }
  
  .expense-table .v-data-table__th .v-data-table-header__content {
    font-size: 0.9rem !important;
    font-weight: 900 !important;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.7) !important;
  }
}

@media (max-width: 480px) {
  .expense-table .v-data-table__th {
    padding: 12px 8px !important;
    font-size: 0.9rem !important;
    font-weight: 900 !important;
    letter-spacing: 0.2px !important;
  }
  
  .expense-table .v-data-table__th .v-data-table-header__content {
    font-size: 0.8rem !important;
    font-weight: 900 !important;
  }
}

/* Expense Table Animation */
.expense-table .v-data-table__tr {
  animation: fadeInUp 0.5s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.v-dialog .v-card {
  box-shadow: 0 20px 40px rgba(30, 58, 138, 0.2) !important;
  border-radius: 12px !important;
  overflow: hidden !important;
  border: 2px solid rgba(30, 58, 138, 0.3) !important;
}

/* Enhanced field visibility */
.ultra-input-field .v-field::before {
  content: '' !important;
  position: absolute !important;
  top: -2px !important;
  left: -2px !important;
  right: -2px !important;
  bottom: -2px !important;
  border-radius: 8px !important;
  background: linear-gradient(135deg, rgba(30, 58, 138, 0.1) 0%, rgba(30, 58, 138, 0.05) 100%) !important;
  z-index: -1 !important;
}

/* Navy Blue Theme for Internal Tables and Data */
.ultra-modern-dialog .v-data-table,
.ultra-modern-dialog table {
  color: #1e3a8a !important;
  border-color: #1e3a8a !important;
}

.ultra-modern-dialog .v-data-table th,
.ultra-modern-dialog th {
  color: #1e3a8a !important;
  background: rgba(30, 58, 138, 0.1) !important;
  border-color: #1e3a8a !important;
}

.ultra-modern-dialog .v-data-table td,
.ultra-modern-dialog td {
  color: #1e3a8a !important;
  border-color: rgba(30, 58, 138, 0.2) !important;
}

.ultra-modern-dialog .v-data-table tr:hover,
.ultra-modern-dialog tr:hover {
  background: rgba(30, 58, 138, 0.05) !important;
}

/* Navy Blue for dropdowns and selects */
.ultra-modern-dialog .v-select .v-field__input {
  color: #1e3a8a !important;
}

.ultra-modern-dialog .v-list-item {
  color: #1e3a8a !important;
}

.ultra-modern-dialog .v-list-item:hover {
  background: rgba(30, 58, 138, 0.1) !important;
}

/* Final border consistency for all form elements */
.ultra-modern-dialog .v-field,
.ultra-modern-dialog .v-text-field .v-field,
.ultra-modern-dialog .v-select .v-field,
.ultra-modern-dialog .v-textarea .v-field,
.ultra-modern-dialog input,
.ultra-modern-dialog textarea,
.ultra-modern-dialog select {
  border: 3px solid #1e3a8a !important;
  border-style: solid !important;
  border-width: 3px !important;
  border-radius: 8px !important;
  box-shadow: 0 3px 8px rgba(30, 58, 138, 0.3) !important;
}

/* Override any conflicting border styles */
.ultra-modern-dialog .v-field * {
  border-color: #1e3a8a !important;
}

/* Ensure text input visibility */
.ultra-modern-dialog input[type="text"],
.ultra-modern-dialog input[type="date"],
.ultra-modern-dialog input[type="number"],
.ultra-modern-dialog textarea {
  color: #1e3a8a !important;
  font-weight: 700 !important;
  background: #ffffff !important;
  border: none !important;
  outline: none !important;
  opacity: 1 !important;
  font-size: 1.1rem !important;
  letter-spacing: 0.3px !important;
  text-shadow: none !important;
}

.ultra-modern-dialog input[type="text"]:focus,
.ultra-modern-dialog input[type="date"]:focus,
.ultra-modern-dialog input[type="number"]:focus,
.ultra-modern-dialog textarea:focus {
  color: #1e40af !important;
  background: #ffffff !important;
  text-shadow: none !important;
  transform: none !important;
}

/* Force text visibility in all input states */
.ultra-modern-dialog .v-field__input,
.ultra-modern-dialog .v-text-field__input,
.ultra-modern-dialog .v-select__selection,
.ultra-modern-dialog .v-textarea__input {
  color: #1e3a8a !important;
  font-weight: 700 !important;
  background: #ffffff !important;
  opacity: 1 !important;
  text-shadow: none !important;
}

.v-overlay__scrim {
  backdrop-filter: blur(4px) !important;
  background: rgba(0, 0, 0, 0.5) !important;
}

/* Mobile Responsive for Form */
@media (max-width: 768px) {
  .form-sections {
    gap: 1.5rem;
  }
  
  .form-section {
    padding: 1rem;
  }
  
  .section-title {
    font-size: 1rem !important;
    margin-bottom: 1rem !important;
  }
  
  .project-dialog-body {
    padding: 1rem !important;
  }
  
  .project-dialog-actions {
    padding: 1rem !important;
  }
  
  .modern-dialog-header {
    padding: 1.5rem !important;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
    gap: 0.75rem;
  }
  
  .header-icon {
    width: 48px;
    height: 48px;
  }
  
  .dialog-main-title {
    font-size: 1.25rem !important;
  }
  
  .modern-dialog-actions {
    flex-direction: column;
    gap: 1rem;
    padding: 1rem !important;
  }
  
  .actions-buttons {
    width: 100%;
    justify-content: center;
  }
  
  .modern-cancel-btn,
  .modern-save-btn {
    flex: 1;
    min-width: auto !important;
  }
  
  .ultra-form-row {
    flex-direction: column;
    gap: 0;
  }
  
  .ultra-modern-header {
    padding: 1rem 1.5rem !important;
  }
  
  .ultra-modern-actions {
    padding: 1rem 1.5rem !important;
  }
  
  .compact-dialog-body {
    padding: 1rem !important;
  }
}

/* Simple Title Section Styles */
.simple-title-section {
  margin: 2rem 0;
  text-align: center;
  padding: 2rem 1rem;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.05) 0%, rgba(139, 92, 246, 0.05) 100%);
  border-radius: 16px;
  border: 2px solid rgba(59, 130, 246, 0.1);
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
  position: relative;
  z-index: 10;
}

.title-container {
  max-width: 800px;
  margin: 0 auto;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
}

.title-icon {
  margin-bottom: 1rem;
  animation: pulse 2s infinite;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  color: #3b82f6 !important;
  margin-left: auto !important;
  margin-right: auto !important;
}

.simple-title {
  font-size: 2.5rem !important;
  font-weight: 900 !important;
  color: #1e293b !important;
  margin-bottom: 1rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.2), 0 2px 4px rgba(0, 0, 0, 0.1) !important;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  text-align: center !important;
  line-height: 1.2 !important;
  letter-spacing: 0.5px !important;
}

.simple-subtitle {
  font-size: 1.2rem !important;
  color: #475569 !important;
  margin-bottom: 1.5rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 600 !important;
  line-height: 1.6 !important;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  text-align: center !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

.title-stats-simple {
  display: flex !important;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  flex-wrap: wrap;
  visibility: visible !important;
  opacity: 1 !important;
  width: 100% !important;
}

.stat-item {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
  color: white;
  padding: 8px 16px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 0.9rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
  transition: all 0.3s ease;
  display: inline-block !important;
  visibility: visible !important;
  opacity: 1 !important;
}

.stat-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.4);
}

.stat-separator {
  color: #94a3b8;
  font-size: 1.2rem;
  font-weight: bold;
}

/* Animations */
@keyframes pulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .simple-title {
    font-size: 1.8rem !important;
  }
  
  .simple-subtitle {
    font-size: 1rem !important;
  }
  
  .title-stats-simple {
    gap: 0.5rem;
  }
  
  .stat-item {
    font-size: 0.8rem;
    padding: 6px 12px;
  }
  
  .simple-title-section {
    padding: 1.5rem 1rem;
  }
}

/* Title Actions */
.title-actions {
  margin-top: 2rem;
  display: flex;
  justify-content: center;
  gap: 1rem;
  flex-wrap: wrap;
}

/* تنسيق نموذج إضافة المشروع - نفس تنسيق نموذج معلومات المهمة */
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
  padding: 16px !important;
  background: #f8fafc !important;
  border: 2px solid #d1d5db !important;
  border-radius: 8px !important;
  margin: 12px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}

.profile-form-content :deep(.v-text-field),
.profile-form-content :deep(.v-textarea) {
  background: #ffffff !important;
  border-radius: 8px !important;
}

.profile-form-content :deep(.v-field) {
  background: #ffffff !important;
}

.profile-form-content :deep(.v-field__outline) {
  color: #cbd5e1 !important;
}

.profile-form-content :deep(.v-field--focused .v-field__outline) {
  color: #3b82f6 !important;
}

.profile-form-content :deep(.v-label) {
  color: #475569 !important;
  font-weight: 500 !important;
}

.profile-form-content :deep(.v-field__input) {
  color: #1e293b !important;
}

.profile-form-instruction {
  font-size: 0.85rem !important;
  color: #64748b !important;
  margin-bottom: 12px !important;
  line-height: 1.4 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.profile-form-row {
  margin: 0 !important;
  row-gap: 4px !important;
}

.profile-form-column {
  padding: 0 6px !important;
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
  margin-bottom: 8px !important;
  padding: 8px !important;
  background: #ffffff !important;
  border: 1.5px solid #d1d5db !important;
  border-radius: 6px !important;
  transition: all 0.2s ease !important;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05) !important;
  flex: 1 !important;
}

.profile-form-field-wrapper:hover {
  border-color: #9ca3af !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08) !important;
}

.profile-form-field-wrapper:focus-within {
  border-color: #3b82f6 !important;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15), 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.profile-form-label {
  display: block !important;
  font-size: 0.8125rem !important;
  font-weight: 600 !important;
  color: #374151 !important;
  margin-bottom: 4px !important;
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
  background: #ffffff !important;
  border-radius: 4px !important;
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

.profile-form-slider {
  margin-top: 6px !important;
  margin-bottom: 4px !important;
}

.progress-value-display {
  text-align: right !important;
  margin-top: 0 !important;
}

.progress-value {
  display: inline-block !important;
  margin-right: auto !important;
  min-width: auto !important;
  padding: 0 !important;
  border-radius: 0 !important;
  border: none !important;
  background: transparent !important;
  font-size: 0.75rem !important;
  font-weight: 600 !important;
  color: #1d4ed8 !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.profile-form-label-row {
  display: flex !important;
  align-items: center !important;
  justify-content: space-between !important;
  gap: 8px !important;
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

.team-toggle-btn {
  border-radius: 25px;
  font-weight: 700;
  text-transform: none;
  padding: 0 2rem;
  height: 56px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.team-toggle-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.team-toggle-btn:hover::before {
  left: 100%;
}

.team-toggle-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.team-toggle-btn.v-btn--variant-elevated {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  box-shadow: 0 6px 24px rgba(16, 185, 129, 0.4);
}

.team-toggle-btn.v-btn--variant-outlined {
  border: 2px solid #3b82f6;
  color: #3b82f6;
  background: rgba(59, 130, 246, 0.05);
}

.team-toggle-btn.v-btn--variant-outlined:hover {
  background: rgba(59, 130, 246, 0.1);
  border-color: #2563eb;
  color: #2563eb;
}

/* Responsive Design for Title Actions */
@media (max-width: 768px) {
  .title-actions {
    margin-top: 1.5rem;
  }
  
  .team-toggle-btn {
    width: 100%;
    max-width: 300px;
  }
}


/* Team Management Section */
.team-management-section {
  margin: 2rem 0;
  animation: slideDown 0.5s ease-out;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.team-management-card {
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(16, 185, 129, 0.2);
  border: 1px solid rgba(16, 185, 129, 0.3);
  backdrop-filter: blur(10px);
}

.team-section-header {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
  border-radius: 20px 20px 0 0;
  padding: 1.5rem;
}

.team-stats-row {
  margin-bottom: 1.5rem;
}

.team-stat-card {
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
}

.team-stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
}

.team-members-table {
  background: transparent;
}

.team-members-table th {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white !important;
  font-weight: 700;
  font-size: 0.9rem;
  text-align: center;
  vertical-align: middle;
  padding: 1rem 0.5rem;
  border: none;
}

.team-members-table td {
  background: rgba(248, 250, 252, 0.9);
  color: #1a1a1a !important;
  font-weight: 500;
  text-align: center;
  vertical-align: middle;
  padding: 1rem 0.5rem;
  border-bottom: 1px solid rgba(226, 232, 240, 0.5);
}

.team-members-table tbody tr:hover td {
  background: rgba(16, 185, 129, 0.05);
  transform: scale(1.01);
}

.member-info {
  text-align: right;
}

.member-name {
  font-size: 1rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 0.25rem 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-email {
  font-size: 0.8rem;
  color: #64748b;
  margin: 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-avatar {
  border: 2px solid rgba(16, 185, 129, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.tasks-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
}

.task-count {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1e293b;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.task-label {
  font-size: 0.7rem;
  color: #64748b;
  font-weight: 500;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.action-buttons {
  display: flex;
  gap: 0.25rem;
  justify-content: center;
  flex-wrap: wrap;
}

.action-buttons .v-btn {
  border-radius: 6px;
  font-weight: 600;
  text-transform: none;
  min-width: 28px;
  height: 28px;
}

/* Responsive Design for Team Management */
@media (max-width: 768px) {
  .team-management-section {
    margin: 1rem 0;
  }
  
  .team-section-header {
    padding: 1rem;
  }
  
  .team-stats-row .v-col {
    margin-bottom: 1rem;
  }
  
  .action-buttons {
    flex-direction: column;
    gap: 0.125rem;
  }
  
  .member-info {
    text-align: center;
  }
}



/* زر تصدير Excel */
.export-excel-btn {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  text-transform: none !important;
  border-radius: 8px !important;
  box-shadow: 0 2px 8px rgba(16, 185, 129, 0.3) !important;
  transition: all 0.3s ease !important;
}

.export-excel-btn:hover {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4) !important;
}

.export-excel-btn .v-icon {
  color: white !important;
  font-size: 1.1rem !important;
}

/* CSS شامل لجميع أسماء الحقول - محاذاة في المنتصف */
.modern-search-field .v-label,
.modern-filter-field .v-label,
.modern-search-field .v-field-label,
.modern-filter-field .v-field-label,
.modern-search-field .v-field-label--active,
.modern-filter-field .v-field-label--active,
.modern-search-field .v-field-label--floating,
.modern-filter-field .v-field-label--floating {
  text-align: center !important;
  direction: rtl !important;
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS شامل مع ::deep لجميع أسماء الحقول */
::deep(.modern-search-field .v-label),
::deep(.modern-filter-field .v-label),
::deep(.modern-search-field .v-field-label),
::deep(.modern-filter-field .v-field-label),
::deep(.modern-search-field .v-field-label--active),
::deep(.modern-filter-field .v-field-label--active),
::deep(.modern-search-field .v-field-label--floating),
::deep(.modern-filter-field .v-field-label--floating) {
  text-align: center !important;
  direction: rtl !important;
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
}

/* CSS نهائي شامل لجميع أسماء الحقول */
.modern-filter-section * .v-label,
.modern-filter-section * .v-field-label,
.modern-filter-section * .v-field-label--active,
.modern-filter-section * .v-field-label--floating,
.modern-search-field * .v-label,
.modern-filter-field * .v-label,
.modern-search-field * .v-field-label,
.modern-filter-field * .v-field-label {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
}

/* CSS نهائي شامل مع ::deep */
::deep(.modern-filter-section * .v-label),
::deep(.modern-filter-section * .v-field-label),
::deep(.modern-filter-section * .v-field-label--active),
::deep(.modern-filter-section * .v-field-label--floating),
::deep(.modern-search-field * .v-label),
::deep(.modern-filter-field * .v-label),
::deep(.modern-search-field * .v-field-label),
::deep(.modern-filter-field * .v-field-label) {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
}

/* CSS إضافي لضمان التطبيق في صفحة المشاريع */
.v-data-table .v-field-label,
.v-data-table .v-field-label--active,
.v-data-table .v-field-label--floating,
.v-data-table .v-label {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
}

/* تنسيق نموذج إضافة عضو الفريق */
.team-member-dialog-card {
  border-radius: 20px !important;
  box-shadow: 0 12px 40px rgba(67, 56, 202, 0.2) !important;
  overflow: hidden !important;
}

/* CSS قوي لضمان التطبيق */
.team-member-dialog-card .v-text-field,
.team-member-dialog-card .v-select {
  margin-bottom: 20px !important;
}

.team-member-dialog-card .v-text-field .v-field,
.team-member-dialog-card .v-select .v-field {
  border: 4px solid #dc2626 !important;
  border-radius: 16px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  box-shadow: 0 8px 24px rgba(220, 38, 38, 0.3) !important;
  min-height: 80px !important;
  padding: 12px !important;
}

.team-member-dialog-card .v-text-field .v-field__input,
.team-member-dialog-card .v-select .v-field__input {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-align: center !important;
  direction: rtl !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  padding: 16px 20px !important;
}

.team-member-dialog-card .v-text-field .v-label,
.team-member-dialog-card .v-select .v-label {
  color: #4b5563 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  padding: 12px 16px !important;
  border-radius: 12px !important;
  border: 3px solid #7f1d1d !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.5) !important;
  text-align: center !important;
  direction: rtl !important;
  margin-bottom: 8px !important;
  display: block !important;
  width: 100% !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* CSS قوي مع ::deep */
::deep(.team-member-dialog-card .v-text-field),
::deep(.team-member-dialog-card .v-select) {
  margin-bottom: 20px !important;
}

::deep(.team-member-dialog-card .v-text-field .v-field),
::deep(.team-member-dialog-card .v-select .v-field) {
  border: 4px solid #dc2626 !important;
  border-radius: 16px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  box-shadow: 0 8px 24px rgba(220, 38, 38, 0.3) !important;
  min-height: 80px !important;
  padding: 12px !important;
}

::deep(.team-member-dialog-card .v-text-field .v-field__input),
::deep(.team-member-dialog-card .v-select .v-field__input) {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-align: center !important;
  direction: rtl !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  padding: 16px 20px !important;
}

::deep(.team-member-dialog-card .v-text-field .v-label),
::deep(.team-member-dialog-card .v-select .v-label) {
  color: #4b5563 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  padding: 12px 16px !important;
  border-radius: 12px !important;
  border: 3px solid #7f1d1d !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.5) !important;
  text-align: center !important;
  direction: rtl !important;
  margin-bottom: 8px !important;
  display: block !important;
  width: 100% !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.team-member-dialog-header {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 50%, #8b5cf6 100%) !important;
  color: white !important;
  padding: 24px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
}

.team-member-dialog-header .text-h5 {
  color: white !important;
  font-weight: 800 !important;
  font-size: 1.5rem !important;
  letter-spacing: 0.5px !important;
}

.team-member-dialog-header .v-icon {
  color: white !important;
}

/* تنسيق حقول نموذج عضو الفريق */
.team-member-form-field {
  margin-bottom: 20px !important;
  position: relative !important;
  display: block !important;
  width: 100% !important;
}

.team-member-form-field .v-field {
  border-radius: 16px !important;
  border: 3px solid #0ea5e9 !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  box-shadow: 0 8px 24px rgba(14, 165, 233, 0.2) !important;
  transition: all 0.3s ease !important;
  margin-bottom: 12px !important;
  padding: 8px !important;
  position: relative !important;
  overflow: visible !important;
  min-height: 80px !important;
  display: flex !important;
  flex-direction: column !important;
  justify-content: center !important;
}

.team-member-form-field .v-field:hover {
  border-color: #4338ca !important;
  box-shadow: 0 12px 32px rgba(67, 56, 202, 0.3) !important;
  transform: translateY(-3px) !important;
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%) !important;
}

.team-member-form-field .v-field:focus-within {
  border-color: #dc2626 !important;
  box-shadow: 0 16px 40px rgba(220, 38, 38, 0.4) !important;
  transform: translateY(-4px) !important;
  background: linear-gradient(135deg, #fef2f2 0%, #fee2e2 100%) !important;
}

/* تنسيق أسماء الحقول في نموذج عضو الفريق */
.team-member-form-field .v-label,
.team-member-form-field .v-field-label,
.team-member-form-field .v-field-label--active,
.team-member-form-field .v-field-label--floating {
  color: #4b5563 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.8) !important;
  letter-spacing: 1px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  padding: 12px 16px !important;
  border-radius: 12px !important;
  border: 3px solid #7f1d1d !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.5) !important;
  margin-bottom: 8px !important;
  display: block !important;
  width: 100% !important;
  position: relative !important;
  z-index: 10 !important;
  line-height: 1.4 !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* تنسيق النصوص داخل حقول نموذج عضو الفريق */
.team-member-form-field .v-field__input {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.3rem !important;
  text-align: center !important;
  direction: rtl !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  padding: 16px 20px !important;
  border-radius: 10px !important;
  border: 3px solid #0ea5e9 !important;
  box-shadow: inset 0 2px 4px rgba(14, 165, 233, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* CSS شامل مع ::deep لنموذج عضو الفريق */
::deep(.team-member-form-field .v-label),
::deep(.team-member-form-field .v-field-label),
::deep(.team-member-form-field .v-field-label--active),
::deep(.team-member-form-field .v-field-label--floating) {
  color: #4b5563 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.8) !important;
  letter-spacing: 1px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  padding: 12px 16px !important;
  border-radius: 12px !important;
  border: 3px solid #7f1d1d !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.5) !important;
  margin-bottom: 8px !important;
  display: block !important;
  width: 100% !important;
  position: relative !important;
  z-index: 10 !important;
  line-height: 1.4 !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

::deep(.team-member-form-field .v-field__input) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1.3rem !important;
  text-align: center !important;
  direction: rtl !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  padding: 16px 20px !important;
  border-radius: 10px !important;
  border: 3px solid #0ea5e9 !important;
  box-shadow: inset 0 2px 4px rgba(14, 165, 233, 0.1) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* CSS إضافي لضمان الوضوح الكامل لنموذج عضو الفريق */
.team-member-form-field * {
  color: #000000 !important;
  font-weight: 800 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

::deep(.team-member-form-field *) {
  color: #000000 !important;
  font-weight: 800 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* ضمان وضوح جميع النصوص في نموذج عضو الفريق */
.team-member-form-field .v-field__input input,
.team-member-form-field .v-field__input textarea,
.team-member-form-field .v-field__input select {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.5rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 0.8px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  padding: 20px 24px !important;
  border-radius: 12px !important;
  border: 4px solid #0ea5e9 !important;
  box-shadow: inset 0 4px 8px rgba(14, 165, 233, 0.2) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

::deep(.team-member-form-field .v-field__input input),
::deep(.team-member-form-field .v-field__input textarea),
::deep(.team-member-form-field .v-field__input select) {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.5rem !important;
  text-shadow: 0 3px 6px rgba(0, 0, 0, 0.5) !important;
  letter-spacing: 0.8px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  padding: 20px 24px !important;
  border-radius: 12px !important;
  border: 4px solid #0ea5e9 !important;
  box-shadow: inset 0 4px 8px rgba(14, 165, 233, 0.2) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* CSS شامل لجميع العناصر في صفحة المشاريع */
.project-management-page * .v-label,
.project-management-page * .v-field-label,
.project-management-page * .v-field-label--active,
.project-management-page * .v-field-label--floating {
  color: #dc2626 !important;
  font-weight: 800 !important;
  font-size: 1.2rem !important;
  text-shadow: 0 2px 4px rgba(220, 38, 38, 0.3) !important;
  letter-spacing: 0.5px !important;
  opacity: 1 !important;
  text-align: center !important;
  direction: rtl !important;
}

/* تحسين تنسيق النموذج العام */
.team-member-dialog-card .v-card-text {
  padding: 32px !important;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  min-height: 400px !important;
}

.team-member-dialog-card .v-row {
  margin: 0 -12px !important;
  align-items: stretch !important;
}

.team-member-dialog-card .v-col {
  padding: 12px !important;
  display: flex !important;
  flex-direction: column !important;
  justify-content: flex-start !important;
}

/* تنسيق محسن للحقول */
.team-member-dialog-card .v-form {
  width: 100% !important;
}

.team-member-dialog-card .v-form .v-row {
  width: 100% !important;
  margin: 0 !important;
}

.team-member-dialog-card .v-form .v-col {
  width: 100% !important;
  padding: 8px !important;
  margin-bottom: 16px !important;
}

/* تحسين أزرار الإجراءات */
.team-member-dialog-card .v-card-actions {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%) !important;
  border-top: 2px solid #e5e7eb !important;
  padding: 20px 24px !important;
}

.team-member-dialog-card .v-btn {
  font-weight: 700 !important;
  font-size: 1.1rem !important;
  padding: 12px 24px !important;
  border-radius: 12px !important;
  text-transform: none !important;
  letter-spacing: 0.5px !important;
}

.team-member-dialog-card .v-btn--variant-elevated {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

.team-member-dialog-card .v-btn--variant-outlined {
  border-width: 2px !important;
}

/* تحسين تنسيق الحقول بشكل عام */
.team-member-dialog-card .v-text-field,
.team-member-dialog-card .v-select {
  margin-bottom: 16px !important;
}

.team-member-dialog-card .v-text-field .v-field,
.team-member-dialog-card .v-select .v-field {
  min-height: 80px !important;
}

/* تحسين تنسيق الصفوف والأعمدة */
.team-member-dialog-card .v-row {
  margin: 0 !important;
  align-items: flex-start !important;
}

.team-member-dialog-card .v-col {
  margin-bottom: 16px !important;
  padding: 0 8px !important;
}

/* تحسين تنسيق النموذج */
.team-member-dialog-card .v-form {
  padding: 0 !important;
}

.team-member-dialog-card .v-form .v-row {
  margin: 0 !important;
}

.team-member-dialog-card .v-form .v-col {
  margin-bottom: 16px !important;
  padding: 0 8px !important;
}

/* تحسين تنسيق الحقول الفردية */
.team-member-dialog-card .v-text-field,
.team-member-dialog-card .v-select {
  width: 100% !important;
  margin: 0 !important;
}

.team-member-dialog-card .v-text-field .v-field,
.team-member-dialog-card .v-select .v-field {
  width: 100% !important;
  margin: 0 !important;
}

/* CSS شامل لضمان التطبيق */
.team-member-dialog-card * .v-field {
  border: 4px solid #dc2626 !important;
  border-radius: 16px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  box-shadow: 0 8px 24px rgba(220, 38, 38, 0.3) !important;
  min-height: 80px !important;
  padding: 12px !important;
}

.team-member-dialog-card * .v-field__input {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-align: center !important;
  direction: rtl !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  padding: 16px 20px !important;
}

.team-member-dialog-card * .v-label {
  color: #4b5563 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  padding: 12px 16px !important;
  border-radius: 12px !important;
  border: 3px solid #7f1d1d !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.5) !important;
  text-align: center !important;
  direction: rtl !important;
  margin-bottom: 8px !important;
  display: block !important;
  width: 100% !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* CSS شامل مع ::deep */
::deep(.team-member-dialog-card * .v-field) {
  border: 4px solid #dc2626 !important;
  border-radius: 16px !important;
  background: linear-gradient(135deg, #ffffff 0%, #f0f9ff 100%) !important;
  box-shadow: 0 8px 24px rgba(220, 38, 38, 0.3) !important;
  min-height: 80px !important;
  padding: 12px !important;
}

::deep(.team-member-dialog-card * .v-field__input) {
  color: #000000 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  text-align: center !important;
  direction: rtl !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  padding: 16px 20px !important;
}

::deep(.team-member-dialog-card * .v-label) {
  color: #4b5563 !important;
  font-weight: 900 !important;
  font-size: 1.4rem !important;
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  padding: 12px 16px !important;
  border-radius: 12px !important;
  border: 3px solid #7f1d1d !important;
  box-shadow: 0 6px 16px rgba(220, 38, 38, 0.5) !important;
  text-align: center !important;
  direction: rtl !important;
  margin-bottom: 8px !important;
  display: block !important;
  width: 100% !important;
  min-height: 50px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
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

/* RTL Date Picker Styles */
.rtl-date-picker {
  direction: rtl !important;
}

.rtl-date-picker :deep(.v-date-picker-header) {
  direction: rtl !important;
}

.rtl-date-picker :deep(.v-date-picker-month__day) {
  direction: rtl !important;
}

</style>