<template>
  <v-container class="fill-height data-page human-resources-page" fluid>
    <div class="fullscreen-content centered-content">
      <!-- Header Section -->
      <div class="engineers-header-card">
        <div class="header-gradient-line"></div>
        <div class="header-content">
          <div class="header-right">
            <div class="engineer-emoji">
              <v-icon size="40" color="white">mdi-account-group</v-icon>
            </div>
            <div class="header-text">
              <h1 class="main-title">الموارد البشرية</h1>
              <p class="subtitle">إدارة وتتبع جميع الموظفين والموارد البشرية</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Statistics Cards -->
      <div class="stats-section">
        <v-row>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-account-group</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ formatNumber(totalEmployees) || '0' }}</h3>
                  <p class="stat-label">إجمالي الموظفين</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-success" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon check-icon">mdi-account-check</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ formatNumber(activeEmployees) || '0' }}</h3>
                  <p class="stat-label">موظفين نشطين</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-info" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-office-building</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ formatNumber(totalDepartments) || '0' }}</h3>
                  <p class="stat-label">إجمالي الأقسام</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-warning" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-calendar-clock</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value">{{ formatNumber(onLeave) || '0' }}</h3>
                  <p class="stat-label">في إجازة</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-primary" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-currency-usd</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value" style="font-size: 1.5rem !important;">{{ formatCurrency(totalSalaries) || '0 د.ع' }}</h3>
                  <p class="stat-label">إجمالي الرواتب</p>
                </div>
              </div>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <v-card class="modern-stat-card stat-card-info" elevation="0">
              <div class="stat-card-background"></div>
              <div class="stat-card-content">
                <div class="stat-icon-wrapper">
                  <v-icon size="48" class="stat-icon">mdi-chart-line</v-icon>
                </div>
                <div class="stat-info">
                  <h3 class="stat-value" style="font-size: 1.5rem !important;">{{ formatCurrency(averageSalary) || '0 د.ع' }}</h3>
                  <p class="stat-label">متوسط الراتب</p>
                </div>
              </div>
            </v-card>
          </v-col>
        </v-row>
      </div>

      <!-- Search Bar -->
      <v-card class="filters-card mb-4" elevation="2">
        <v-card-title class="filters-header-new">
          <div class="d-flex align-center justify-start">
            <v-icon size="20" color="white" class="me-2" style="color: #ffffff !important;">mdi-menu-down</v-icon>
            <span class="filters-header-title" style="color: #ef4444 !important;">البحث والفلترة</span>
          </div>
        </v-card-title>
        <v-card-text class="filters-content">
          <v-row no-gutters>
            <v-col cols="12" md="3" class="px-2">
              <v-text-field
                v-model="searchQuery"
                placeholder="البحث في الموظفين"
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                clearable
                hide-details
                density="compact"
                class="search-field-new"
              />
            </v-col>
            <v-col cols="12" md="3" class="px-2">
              <v-select
                v-model="selectedDepartment"
                :items="departmentOptions"
                label="القسم"
                variant="outlined"
                density="compact"
                clearable
                hide-details
                class="filter-field-new"
              />
            </v-col>
            <v-col cols="12" md="3" class="px-2">
              <v-select
                v-model="selectedStatus"
                :items="statusOptions"
                label="الحالة"
                variant="outlined"
                density="compact"
                clearable
                hide-details
                class="filter-field-new"
              />
            </v-col>
            <v-col cols="12" md="3" class="px-2">
              <v-select
                v-model="selectedPosition"
                :items="positionOptions"
                label="المنصب"
                variant="outlined"
                density="compact"
                clearable
                hide-details
                class="filter-field-new"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Employees Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title d-flex align-center justify-space-between">
          <span class="title-text">قائمة الموظفين</span>
          <v-btn
            v-if="canCreate"
            class="add-button btn-glow light-sweep smooth-transition"
            @click="openAddEmployeeDialog"
            elevation="2"
            color="primary"
            size="small"
            style="background: linear-gradient(135deg, #93c5fd 0%, #60a5fa 50%, #3b82f6 100%) !important; color: #ffffff !important; height: 36px !important; font-size: 0.875rem !important;"
          >
            <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
            إضافة موظف
          </v-btn>
        </v-card-title>

        <v-data-table
          :headers="headers"
          :items="filteredEmployees"
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

          <!-- Employee Name Column -->
          <template #item.name="{ item }">
            <span class="project-name">{{ item.name }}</span>
          </template>

          <!-- ID Number Column -->
          <template #item.idNumber="{ item }">
            <span class="date-text">{{ item.idNumber || '-' }}</span>
          </template>

          <!-- Department Column -->
          <template #item.department="{ item }">
            <v-chip class="category-chip" size="small">
              {{ item.department }}
            </v-chip>
          </template>

          <!-- Position Column -->
          <template #item.position="{ item }">
            <span class="project-name">{{ item.position }}</span>
          </template>

          <!-- Phone Column -->
          <template #item.phone="{ item }">
            <span class="date-text">{{ item.phone }}</span>
          </template>

          <!-- Email Column -->
          <template #item.email="{ item }">
            <span class="project-name">{{ item.email }}</span>
          </template>

          <!-- Salary Column -->
          <template #item.salary="{ item }">
            <span class="cost-text">{{ formatCurrency(item.salary) }}</span>
          </template>

          <!-- Hire Date Column -->
          <template #item.hireDate="{ item }">
            <span class="date-text">{{ item.hireDate }}</span>
          </template>

          <!-- Status Column -->
          <template #item.status="{ item }">
            <v-chip class="status-chip" size="small">
              {{ getStatusText(item.status) }}
            </v-chip>
          </template>

          <!-- Actions Column -->
          <template #item.actions="{ item }">
            <div class="action-buttons">
              <v-btn
                size="small"
                color="primary"
                variant="text"
                @click="viewEmployeeDetails(item)"
                icon
                class="action-btn details-btn"
                title="عرض التفاصيل"
              >
                <v-icon size="14">mdi-eye</v-icon>
              </v-btn>
              <v-btn
                v-if="canUpdate"
                size="small"
                color="success"
                variant="text"
                @click="editEmployee(item)"
                icon
                class="action-btn"
                title="تعديل"
              >
                <v-icon size="14">mdi-pencil</v-icon>
              </v-btn>
              <v-btn
                v-if="canDelete"
                size="small"
                color="error"
                variant="text"
                @click="deleteEmployee(item)"
                icon
                class="action-btn"
                title="حذف"
              >
                <v-icon size="14">mdi-delete</v-icon>
              </v-btn>
            </div>
          </template>
        </v-data-table>
      </v-card>

      <!-- Add/Edit Employee Dialog -->
      <v-dialog v-model="employeeDialog" max-width="1000" scrollable persistent>
        <v-card class="image-style-dialog">
          <!-- Header Section -->
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-account-plus</v-icon>
                <span class="header-title">{{ isEditing ? 'تعديل بيانات الموظف' : 'إضافة موظف جديد' }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeEmployeeDialog"
                class="close-btn"
              />
            </div>
          </div>

          <!-- Form Content -->
          <div class="dialog-body">
            <v-form ref="employeeForm" v-model="employeeFormValid">
              <v-tabs v-model="formTab" bg-color="transparent" class="form-tabs">
                <v-tab value="personal">المعلومات الشخصية</v-tab>
                <v-tab value="work">المعلومات الوظيفية</v-tab>
                <v-tab value="contact">معلومات الاتصال</v-tab>
                <v-tab value="additional">معلومات إضافية</v-tab>
              </v-tabs>

              <v-window v-model="formTab">
                <!-- Personal Information Tab -->
                <v-window-item value="personal">
                  <div class="form-fields mt-4">
                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.name"
                          label="اسم الموظف الكامل"
                          variant="outlined"
                          :rules="[v => !!v || 'اسم الموظف مطلوب']"
                          required
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.idNumber"
                          label="رقم الهوية الوطنية"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12" md="4">
                        <v-text-field
                          v-model="employeeForm.birthDate"
                          label="تاريخ الميلاد"
                          variant="outlined"
                          type="date"
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="4">
                        <v-select
                          v-model="employeeForm.gender"
                          :items="genderOptions"
                          label="الجنس"
                          variant="outlined"
                          class="form-field black-list"
                        />
                      </v-col>
                      <v-col cols="12" md="4">
                        <v-select
                          v-model="employeeForm.maritalStatus"
                          :items="maritalStatusOptions"
                          label="الحالة الاجتماعية"
                          variant="outlined"
                          class="form-field black-list"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.nationality"
                          label="الجنسية"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.address"
                          label="العنوان"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>
                  </div>
                </v-window-item>

                <!-- Work Information Tab -->
                <v-window-item value="work">
                  <div class="form-fields mt-4">
                    <v-row>
                      <v-col cols="12" md="6">
                        <v-select
                          v-model="employeeForm.department"
                          :items="departments"
                          label="القسم"
                          variant="outlined"
                          :rules="[v => !!v || 'القسم مطلوب']"
                          required
                          class="form-field black-list"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.position"
                          label="المنصب"
                          variant="outlined"
                          :rules="[v => !!v || 'المنصب مطلوب']"
                          required
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.hireDate"
                          label="تاريخ التوظيف"
                          variant="outlined"
                          type="date"
                          :rules="[v => !!v || 'تاريخ التوظيف مطلوب']"
                          required
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.contractType"
                          label="نوع العقد"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.salary"
                          label="الراتب الأساسي (د.ع)"
                          variant="outlined"
                          type="number"
                          :rules="[v => !!v || 'الراتب مطلوب', v => v > 0 || 'الراتب يجب أن يكون أكبر من صفر']"
                          required
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-select
                          v-model="employeeForm.status"
                          :items="statusOptions"
                          label="الحالة"
                          variant="outlined"
                          :rules="[v => !!v || 'الحالة مطلوبة']"
                          required
                          class="form-field"
                        />
                      </v-col>
                    </v-row>
                  </div>
                </v-window-item>

                <!-- Contact Information Tab -->
                <v-window-item value="contact">
                  <div class="form-fields mt-4">
                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.phone"
                          label="رقم الهاتف"
                          variant="outlined"
                          :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                          required
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.phone2"
                          label="رقم هاتف إضافي"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.email"
                          label="البريد الإلكتروني"
                          variant="outlined"
                          type="email"
                          :rules="[v => !!v || 'البريد الإلكتروني مطلوب', v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح']"
                          required
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.emergencyContact"
                          label="جهة الاتصال في حالة الطوارئ"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.emergencyPhone"
                          label="رقم هاتف الطوارئ"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>
                  </div>
                </v-window-item>

                <!-- Additional Information Tab -->
                <v-window-item value="additional">
                  <div class="form-fields mt-4">
                    <v-row>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.education"
                          label="المؤهل العلمي"
                          variant="outlined"
                          class="form-field"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          v-model="employeeForm.experience"
                          label="سنوات الخبرة"
                          variant="outlined"
                          type="number"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12">
                        <v-textarea
                          v-model="employeeForm.notes"
                          label="ملاحظات"
                          variant="outlined"
                          rows="4"
                          class="form-field"
                        />
                      </v-col>
                    </v-row>

                    <v-row>
                      <v-col cols="12">
                        <v-card class="fingerprint-card" elevation="2">
                          <v-card-title class="fingerprint-title">
                            <v-icon class="me-2">mdi-fingerprint</v-icon>
                            بصمة الموظف
                          </v-card-title>
                          <v-card-text>
                            <div class="fingerprint-status">
                              <v-chip
                                :color="employeeForm.fingerprint ? 'primary' : 'info'"
                                size="large"
                                class="mb-3 fingerprint-chip"
                                :class="employeeForm.fingerprint ? 'fingerprint-registered' : 'fingerprint-pending'"
                              >
                                <v-icon start>{{ employeeForm.fingerprint ? 'mdi-check-circle' : 'mdi-alert-circle' }}</v-icon>
                                {{ employeeForm.fingerprint ? 'تم تسجيل البصمة' : 'لم يتم تسجيل البصمة' }}
                              </v-chip>
                            </div>
                            <v-btn
                              color="primary"
                              variant="elevated"
                              @click="registerFingerprint"
                              block
                              :disabled="loadingFingerprint"
                              class="fingerprint-btn"
                            >
                              <v-icon class="me-2">{{ loadingFingerprint ? 'mdi-loading mdi-spin' : 'mdi-fingerprint' }}</v-icon>
                              {{ employeeForm.fingerprint ? 'تحديث البصمة' : 'تسجيل البصمة' }}
                            </v-btn>
                            <p class="fingerprint-hint mt-3 text-caption text-grey">
                              اضغط على الزر واتبع التعليمات لتسجيل بصمة الموظف
                            </p>
                          </v-card-text>
                        </v-card>
                      </v-col>
                    </v-row>
                  </div>
                </v-window-item>
              </v-window>
            </v-form>
          </div>

          <!-- Dialog Actions -->
          <div class="dialog-actions">
            <v-btn
              color="grey"
              variant="text"
              @click="closeEmployeeDialog"
              class="action-btn"
            >
              إلغاء
            </v-btn>
            <v-btn
              color="primary"
              variant="elevated"
              @click="saveEmployee"
              :disabled="!employeeFormValid"
              class="action-btn primary-btn"
            >
              <v-icon class="me-2">mdi-content-save</v-icon>
              {{ isEditing ? 'تحديث' : 'حفظ' }}
            </v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Employee Details Dialog -->
      <v-dialog v-model="detailsDialog" max-width="1200" scrollable persistent>
        <v-card class="details-dialog">
          <div class="dialog-header details-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-account-details</v-icon>
                <span class="header-title">تفاصيل الموظف: {{ selectedEmployeeDetails?.name }}</span>
              </div>
              <v-btn
                icon="mdi-close"
                variant="text"
                size="small"
                color="white"
                @click="closeDetailsDialog"
                class="close-btn"
              />
            </div>
          </div>

          <div class="dialog-body details-body" v-if="selectedEmployeeDetails">
            <v-tabs v-model="detailsTab" bg-color="transparent" class="details-tabs">
              <v-tab value="info">
                <v-icon class="me-2">mdi-information</v-icon>
                المعلومات الأساسية
              </v-tab>
              <v-tab value="leaves">
                <v-icon class="me-2">mdi-calendar</v-icon>
                الإجازات
              </v-tab>
              <v-tab value="attendance">
                <v-icon class="me-2">mdi-clock-in</v-icon>
                الحضور والانصراف
              </v-tab>
              <v-tab value="evaluations">
                <v-icon class="me-2">mdi-star</v-icon>
                التقييمات
              </v-tab>
              <v-tab value="skills">
                <v-icon class="me-2">mdi-certificate</v-icon>
                المهارات والشهادات
              </v-tab>
              <v-tab value="fingerprint">
                <v-icon class="me-2">mdi-fingerprint</v-icon>
                البصمة
              </v-tab>
              <v-tab value="salary">
                <v-icon class="me-2">mdi-cash</v-icon>
                سجل الرواتب
              </v-tab>
            </v-tabs>

            <v-window v-model="detailsTab">
              <!-- Basic Information Tab -->
              <v-window-item value="info">
                <div class="info-content mt-4">
                  <v-row>
                    <v-col cols="12" md="6">
                      <v-card class="info-card" elevation="2">
                        <v-card-title class="info-card-title">المعلومات الشخصية</v-card-title>
                        <v-card-text>
                          <div class="info-item">
                            <span class="info-label">الاسم الكامل:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.name }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">رقم الهوية:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.idNumber || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">تاريخ الميلاد:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.birthDate || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">الجنس:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.gender || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">الحالة الاجتماعية:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.maritalStatus || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">الجنسية:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.nationality || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">العنوان:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.address || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">حالة البصمة:</span>
                            <v-chip
                              :color="selectedEmployeeDetails.fingerprint ? 'primary' : 'info'"
                              size="small"
                              :class="selectedEmployeeDetails.fingerprint ? 'fingerprint-registered-small' : 'fingerprint-pending-small'"
                            >
                              <v-icon start size="small">{{ selectedEmployeeDetails.fingerprint ? 'mdi-check-circle' : 'mdi-alert-circle' }}</v-icon>
                              {{ selectedEmployeeDetails.fingerprint ? 'مسجلة' : 'غير مسجلة' }}
                            </v-chip>
                          </div>
                        </v-card-text>
                      </v-card>
                    </v-col>
                    <v-col cols="12" md="6">
                      <v-card class="info-card" elevation="2">
                        <v-card-title class="info-card-title">المعلومات الوظيفية</v-card-title>
                        <v-card-text>
                          <div class="info-item">
                            <span class="info-label">القسم:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.department }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">المنصب:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.position }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">تاريخ التوظيف:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.hireDate }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">نوع العقد:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.contractType || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">الراتب:</span>
                            <span class="info-value">{{ formatCurrency(selectedEmployeeDetails.salary) }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">الحالة:</span>
                            <v-chip class="status-chip" size="small">{{ selectedEmployeeDetails.status }}</v-chip>
                          </div>
                        </v-card-text>
                      </v-card>
                    </v-col>
                  </v-row>
                  <v-row class="mt-2">
                    <v-col cols="12">
                      <v-card class="info-card" elevation="2">
                        <v-card-title class="info-card-title">معلومات الاتصال</v-card-title>
                        <v-card-text>
                          <v-row>
                            <v-col cols="12" md="6">
                              <div class="info-item">
                                <span class="info-label">رقم الهاتف:</span>
                                <span class="info-value">{{ selectedEmployeeDetails.phone }}</span>
                              </div>
                              <div class="info-item">
                                <span class="info-label">البريد الإلكتروني:</span>
                                <span class="info-value">{{ selectedEmployeeDetails.email }}</span>
                              </div>
                            </v-col>
                            <v-col cols="12" md="6">
                              <div class="info-item">
                                <span class="info-label">جهة الاتصال في الطوارئ:</span>
                                <span class="info-value">{{ selectedEmployeeDetails.emergencyContact || '-' }}</span>
                              </div>
                              <div class="info-item">
                                <span class="info-label">رقم هاتف الطوارئ:</span>
                                <span class="info-value">{{ selectedEmployeeDetails.emergencyPhone || '-' }}</span>
                              </div>
                            </v-col>
                          </v-row>
                        </v-card-text>
                      </v-card>
                    </v-col>
                  </v-row>
                </div>
              </v-window-item>

              <!-- Leaves Tab -->
              <v-window-item value="leaves">
                <div class="leaves-content mt-4">
                  <div class="d-flex justify-space-between align-center mb-4">
                    <h3 class="section-title">سجل الإجازات</h3>
                    <v-btn color="primary" @click="openAddLeaveDialog">
                      <v-icon class="me-2">mdi-plus</v-icon>
                      إضافة إجازة
                    </v-btn>
                  </div>
                  <v-data-table
                    :headers="leaveHeaders"
                    :items="selectedEmployeeDetails.leaves || []"
                    class="details-table"
                    no-data-text="لا توجد إجازات مسجلة"
                  >
                    <template #item.startDate="{ item }">
                      <span class="date-text">{{ item.startDate }}</span>
                    </template>
                    <template #item.endDate="{ item }">
                      <span class="date-text">{{ item.endDate }}</span>
                    </template>
                    <template #item.days="{ item }">
                      <span class="cost-text">{{ item.days }} يوم</span>
                    </template>
                    <template #item.type="{ item }">
                      <v-chip size="small">{{ item.type }}</v-chip>
                    </template>
                    <template #item.status="{ item }">
                      <v-chip size="small" :color="item.status === 'موافق' ? 'success' : item.status === 'مرفوض' ? 'error' : 'warning'">
                        {{ item.status }}
                      </v-chip>
                    </template>
                  </v-data-table>
                </div>
              </v-window-item>

              <!-- Attendance Tab -->
              <v-window-item value="attendance">
                <div class="attendance-content mt-4">
                  <div class="d-flex justify-space-between align-center mb-4">
                    <h3 class="section-title">سجل الحضور والانصراف</h3>
                    <v-btn color="primary" @click="openAddAttendanceDialog">
                      <v-icon class="me-2">mdi-plus</v-icon>
                      تسجيل حضور
                    </v-btn>
                  </div>
                  <v-data-table
                    :headers="attendanceHeaders"
                    :items="selectedEmployeeDetails.attendance || []"
                    class="details-table"
                    no-data-text="لا توجد سجلات حضور"
                  >
                    <template #item.date="{ item }">
                      <span class="date-text">{{ item.date }}</span>
                    </template>
                    <template #item.checkIn="{ item }">
                      <span class="date-text">{{ item.checkIn }}</span>
                    </template>
                    <template #item.checkOut="{ item }">
                      <span class="date-text">{{ item.checkOut || '-' }}</span>
                    </template>
                    <template #item.hours="{ item }">
                      <span class="cost-text">{{ item.hours }} ساعة</span>
                    </template>
                    <template #item.status="{ item }">
                      <v-chip size="small" :color="item.status === 'حاضر' ? 'success' : 'error'">
                        {{ item.status }}
                      </v-chip>
                    </template>
                  </v-data-table>
                </div>
              </v-window-item>

              <!-- Evaluations Tab -->
              <v-window-item value="evaluations">
                <div class="evaluations-content mt-4">
                  <div class="d-flex justify-space-between align-center mb-4">
                    <h3 class="section-title">سجل التقييمات</h3>
                    <v-btn color="primary" @click="openAddEvaluationDialog">
                      <v-icon class="me-2">mdi-plus</v-icon>
                      إضافة تقييم
                    </v-btn>
                  </div>
                  <v-data-table
                    :headers="evaluationHeaders"
                    :items="selectedEmployeeDetails.evaluations || []"
                    class="details-table"
                    no-data-text="لا توجد تقييمات مسجلة"
                  >
                    <template #item.date="{ item }">
                      <span class="date-text">{{ item.date }}</span>
                    </template>
                    <template #item.rating="{ item }">
                      <div class="rating-display">
                        <v-icon v-for="i in 5" :key="i" :color="i <= item.rating ? 'warning' : 'grey'" size="small">
                          mdi-star{{ i <= item.rating ? '' : '-outline' }}
                        </v-icon>
                        <span class="ms-2">{{ item.rating }}/5</span>
                      </div>
                    </template>
                    <template #item.evaluator="{ item }">
                      <span class="project-name">{{ item.evaluator }}</span>
                    </template>
                  </v-data-table>
                </div>
              </v-window-item>

              <!-- Skills Tab -->
              <v-window-item value="skills">
                <div class="skills-content mt-4">
                  <v-row>
                    <v-col cols="12" md="6">
                      <div class="d-flex justify-space-between align-center mb-4">
                        <h3 class="section-title">المهارات</h3>
                        <v-btn color="primary" size="small" @click="openAddSkillDialog">
                          <v-icon class="me-2">mdi-plus</v-icon>
                          إضافة مهارة
                        </v-btn>
                      </div>
                      <v-card class="info-card" elevation="2">
                        <v-card-text>
                          <div v-if="selectedEmployeeDetails.skills && selectedEmployeeDetails.skills.length > 0">
                            <v-chip
                              v-for="(skill, index) in selectedEmployeeDetails.skills"
                              :key="index"
                              class="ma-1"
                              color="primary"
                              closable
                              @click:close="deleteSkill(skill)"
                            >
                              {{ skill }}
                            </v-chip>
                          </div>
                          <div v-else class="text-center text-grey">لا توجد مهارات مسجلة</div>
                        </v-card-text>
                      </v-card>
                    </v-col>
                    <v-col cols="12" md="6">
                      <div class="d-flex justify-space-between align-center mb-4">
                        <h3 class="section-title">الشهادات</h3>
                        <v-btn color="primary" size="small" @click="openAddCertificateDialog">
                          <v-icon class="me-2">mdi-plus</v-icon>
                          إضافة شهادة
                        </v-btn>
                      </div>
                      <v-card class="info-card" elevation="2">
                        <v-card-text>
                          <div v-if="selectedEmployeeDetails.certificates && selectedEmployeeDetails.certificates.length > 0">
                            <div
                              v-for="(cert, index) in selectedEmployeeDetails.certificates"
                              :key="index"
                              class="certificate-item mb-2"
                            >
                              <div class="d-flex justify-space-between align-center">
                                <div>
                                  <div class="font-weight-bold">{{ cert.name }}</div>
                                  <div class="text-caption text-grey">{{ cert.issuer }} - {{ cert.date }}</div>
                                </div>
                                <div class="d-flex align-center gap-2">
                                  <v-icon color="success">mdi-certificate</v-icon>
                                  <v-btn
                                    icon
                                    size="small"
                                    color="error"
                                    variant="text"
                                    @click="deleteCertificate(cert.id)"
                                  >
                                    <v-icon size="small">mdi-delete</v-icon>
                                  </v-btn>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div v-else class="text-center text-grey">لا توجد شهادات مسجلة</div>
                        </v-card-text>
                      </v-card>
                    </v-col>
                  </v-row>
                </div>
              </v-window-item>

              <!-- Fingerprint Tab -->
              <v-window-item value="fingerprint">
                <div class="fingerprint-content mt-4">
                  <v-card class="fingerprint-management-card" elevation="2">
                    <v-card-title class="fingerprint-card-title">
                      <v-icon class="me-2">mdi-fingerprint</v-icon>
                      إدارة بصمة الموظف
                    </v-card-title>
                    <v-card-text>
                      <div class="fingerprint-status-section mb-6">
                        <h3 class="section-title mb-3">حالة البصمة</h3>
                        <v-chip
                          :color="selectedEmployeeDetails.fingerprint ? 'primary' : 'info'"
                          size="large"
                          class="mb-4 fingerprint-chip"
                          :class="selectedEmployeeDetails.fingerprint ? 'fingerprint-registered' : 'fingerprint-pending'"
                        >
                          <v-icon start>{{ selectedEmployeeDetails.fingerprint ? 'mdi-check-circle' : 'mdi-alert-circle' }}</v-icon>
                          {{ selectedEmployeeDetails.fingerprint ? 'تم تسجيل البصمة بنجاح' : 'لم يتم تسجيل البصمة بعد' }}
                        </v-chip>
                        <div v-if="selectedEmployeeDetails.fingerprint" class="fingerprint-info mt-4">
                          <div class="info-item">
                            <span class="info-label">تاريخ التسجيل:</span>
                            <span class="info-value">{{ selectedEmployeeDetails.fingerprintDate || '-' }}</span>
                          </div>
                          <div class="info-item">
                            <span class="info-label">حالة البصمة:</span>
                            <v-chip color="primary" size="small" class="fingerprint-status-chip">نشطة</v-chip>
                          </div>
                        </div>
                      </div>

                      <v-divider class="my-6"></v-divider>

                      <div class="fingerprint-actions">
                        <h3 class="section-title mb-3">الإجراءات</h3>
                        <v-row>
                          <v-col cols="12" md="6">
                            <v-btn
                              color="primary"
                              variant="elevated"
                              @click="registerFingerprintFromDetails"
                              block
                              size="large"
                              :disabled="loadingFingerprint"
                              class="mb-2 fingerprint-btn"
                            >
                              <v-icon class="me-2">{{ loadingFingerprint ? 'mdi-loading mdi-spin' : 'mdi-fingerprint' }}</v-icon>
                              {{ selectedEmployeeDetails.fingerprint ? 'تحديث البصمة' : 'تسجيل بصمة جديدة' }}
                            </v-btn>
                          </v-col>
                          <v-col cols="12" md="6">
                            <v-btn
                              color="error"
                              variant="elevated"
                              @click="deleteFingerprint"
                              block
                              size="large"
                              :disabled="!selectedEmployeeDetails.fingerprint || loadingFingerprint"
                              class="fingerprint-delete-btn"
                            >
                              <v-icon class="me-2">mdi-delete</v-icon>
                              حذف البصمة
                            </v-btn>
                          </v-col>
                        </v-row>
                        <p class="fingerprint-hint mt-4 text-caption text-grey">
                          <v-icon size="small" class="me-1">mdi-information</v-icon>
                          عند الضغط على تسجيل البصمة، سيتم توجيهك لتسجيل بصمة الموظف باستخدام قارئ البصمة المتصل
                        </p>
                      </div>
                    </v-card-text>
                  </v-card>
                </div>
              </v-window-item>

              <!-- Salary History Tab -->
              <v-window-item value="salary">
                <div class="salary-content mt-4">
                  <div class="d-flex justify-space-between align-center mb-4">
                    <h3 class="section-title">سجل الرواتب</h3>
                    <v-btn color="primary" @click="openAddSalaryDialog">
                      <v-icon class="me-2">mdi-plus</v-icon>
                      إضافة راتب
                    </v-btn>
                  </div>
                  <v-data-table
                    :headers="salaryHeaders"
                    :items="selectedEmployeeDetails.salaryHistory || []"
                    class="details-table"
                    no-data-text="لا توجد سجلات رواتب"
                  >
                    <template #item.month="{ item }">
                      <span class="date-text">{{ item.month }}</span>
                    </template>
                    <template #item.baseSalary="{ item }">
                      <span class="cost-text">{{ formatCurrency(item.baseSalary) }}</span>
                    </template>
                    <template #item.bonuses="{ item }">
                      <span class="cost-text">{{ formatCurrency(item.bonuses || 0) }}</span>
                    </template>
                    <template #item.deductions="{ item }">
                      <span class="cost-text text-error">{{ formatCurrency(item.deductions || 0) }}</span>
                    </template>
                    <template #item.netSalary="{ item }">
                      <span class="cost-text text-success font-weight-bold">{{ formatCurrency(item.netSalary) }}</span>
                    </template>
                    <template #item.status="{ item }">
                      <v-chip size="small" :color="item.status === 'مدفوع' ? 'success' : 'warning'">
                        {{ item.status }}
                      </v-chip>
                    </template>
                  </v-data-table>
                </div>
              </v-window-item>
            </v-window>
          </div>

          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeDetailsDialog">
              إغلاق
            </v-btn>
            <v-btn color="primary" variant="elevated" @click="editEmployee(selectedEmployeeDetails)">
              <v-icon class="me-2">mdi-pencil</v-icon>
              تعديل
            </v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Add Leave Dialog -->
      <v-dialog v-model="leaveDialog" max-width="600" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-calendar-plus</v-icon>
                <span class="header-title">إضافة إجازة</span>
              </div>
              <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeLeaveDialog" class="close-btn" />
            </div>
          </div>
          <div class="dialog-body">
            <v-form>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="leaveForm.startDate"
                    label="تاريخ البداية"
                    type="date"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="leaveForm.endDate"
                    label="تاريخ النهاية"
                    type="date"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="leaveForm.type"
                    :items="leaveTypeOptions"
                    label="نوع الإجازة"
                    variant="outlined"
                    required
                    class="form-field black-list"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="leaveForm.status"
                    :items="leaveStatusOptions"
                    label="الحالة"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-textarea
                    v-model="leaveForm.reason"
                    label="السبب"
                    variant="outlined"
                    rows="3"
                    class="form-field"
                  />
                </v-col>
              </v-row>
            </v-form>
          </div>
          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeLeaveDialog">إلغاء</v-btn>
            <v-btn color="primary" variant="elevated" @click="saveLeave">حفظ</v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Add Attendance Dialog -->
      <v-dialog v-model="attendanceDialog" max-width="600" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-clock-in</v-icon>
                <span class="header-title">تسجيل حضور</span>
              </div>
              <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeAttendanceDialog" class="close-btn" />
            </div>
          </div>
          <div class="dialog-body">
            <v-form>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="attendanceForm.date"
                    label="التاريخ"
                    type="date"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="attendanceForm.status"
                    :items="attendanceStatusOptions"
                    label="الحالة"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="attendanceForm.checkIn"
                    label="وقت الدخول"
                    type="time"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="attendanceForm.checkOut"
                    label="وقت الخروج"
                    type="time"
                    variant="outlined"
                    class="form-field"
                  />
                </v-col>
              </v-row>
            </v-form>
          </div>
          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeAttendanceDialog">إلغاء</v-btn>
            <v-btn color="primary" variant="elevated" @click="saveAttendance">حفظ</v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Add Evaluation Dialog -->
      <v-dialog v-model="evaluationDialog" max-width="600" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-star-plus</v-icon>
                <span class="header-title">إضافة تقييم</span>
              </div>
              <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeEvaluationDialog" class="close-btn" />
            </div>
          </div>
          <div class="dialog-body">
            <v-form>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="evaluationForm.date"
                    label="التاريخ"
                    type="date"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="evaluationForm.evaluator"
                    label="المقيّم"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-slider
                    v-model="evaluationForm.rating"
                    label="التقييم"
                    min="1"
                    max="5"
                    step="1"
                    thumb-label="always"
                    class="form-field"
                  >
                    <template #append>
                      <v-text-field
                        v-model="evaluationForm.rating"
                        type="number"
                        style="width: 60px"
                        density="compact"
                        hide-details
                        variant="outlined"
                      />
                    </template>
                  </v-slider>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-textarea
                    v-model="evaluationForm.comments"
                    label="ملاحظات"
                    variant="outlined"
                    rows="3"
                    class="form-field"
                  />
                </v-col>
              </v-row>
            </v-form>
          </div>
          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeEvaluationDialog">إلغاء</v-btn>
            <v-btn color="primary" variant="elevated" @click="saveEvaluation">حفظ</v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Add Skill Dialog -->
      <v-dialog v-model="skillDialog" max-width="500" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-certificate</v-icon>
                <span class="header-title">إضافة مهارة</span>
              </div>
              <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeSkillDialog" class="close-btn" />
            </div>
          </div>
          <div class="dialog-body">
            <v-form>
              <v-text-field
                v-model="skillForm.name"
                label="اسم المهارة"
                variant="outlined"
                required
                class="form-field"
              />
            </v-form>
          </div>
          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeSkillDialog">إلغاء</v-btn>
            <v-btn color="primary" variant="elevated" @click="saveSkill">حفظ</v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Add Certificate Dialog -->
      <v-dialog v-model="certificateDialog" max-width="600" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-certificate</v-icon>
                <span class="header-title">إضافة شهادة</span>
              </div>
              <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeCertificateDialog" class="close-btn" />
            </div>
          </div>
          <div class="dialog-body">
            <v-form>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="certificateForm.name"
                    label="اسم الشهادة"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="certificateForm.issuer"
                    label="المصدر"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="certificateForm.date"
                    label="تاريخ الحصول"
                    type="date"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="certificateForm.expiryDate"
                    label="تاريخ الانتهاء (اختياري)"
                    type="date"
                    variant="outlined"
                    class="form-field"
                  />
                </v-col>
              </v-row>
            </v-form>
          </div>
          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeCertificateDialog">إلغاء</v-btn>
            <v-btn color="primary" variant="elevated" @click="saveCertificate">حفظ</v-btn>
          </div>
        </v-card>
      </v-dialog>

      <!-- Add Salary Dialog -->
      <v-dialog v-model="salaryDialog" max-width="600" persistent>
        <v-card class="image-style-dialog">
          <div class="dialog-header">
            <div class="header-content">
              <div class="header-left">
                <v-icon size="24" color="white" class="header-icon">mdi-cash-plus</v-icon>
                <span class="header-title">إضافة راتب</span>
              </div>
              <v-btn icon="mdi-close" variant="text" size="small" color="white" @click="closeSalaryDialog" class="close-btn" />
            </div>
          </div>
          <div class="dialog-body">
            <v-form>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="salaryForm.month"
                    label="الشهر (YYYY-MM)"
                    variant="outlined"
                    required
                    placeholder="2024-01"
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-select
                    v-model="salaryForm.status"
                    :items="salaryStatusOptions"
                    label="الحالة"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="salaryForm.baseSalary"
                    label="الراتب الأساسي (د.ع)"
                    type="number"
                    variant="outlined"
                    required
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="salaryForm.bonuses"
                    label="المكافآت (د.ع)"
                    type="number"
                    variant="outlined"
                    class="form-field"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="salaryForm.deductions"
                    label="الخصومات (د.ع)"
                    type="number"
                    variant="outlined"
                    class="form-field"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <v-card class="pa-4" style="background: #f0fdf4;">
                    <div class="text-caption text-grey mb-1">صافي الراتب</div>
                    <div class="text-h6 text-success font-weight-bold">
                      {{ formatCurrency((parseFloat(salaryForm.baseSalary || 0) + parseFloat(salaryForm.bonuses || 0) - parseFloat(salaryForm.deductions || 0))) }}
                    </div>
                  </v-card>
                </v-col>
              </v-row>
            </v-form>
          </div>
          <div class="dialog-actions">
            <v-btn color="grey" variant="text" @click="closeSalaryDialog">إلغاء</v-btn>
            <v-btn color="primary" variant="elevated" @click="saveSalary">حفظ</v-btn>
          </div>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { usePermissions } from '@/composables/usePermissions'

// Permissions
const { canCreate, canUpdate, canDelete } = usePermissions('/human-resources')

// ========================================
// متغيرات الحالة الأساسية
// ========================================
const loading = ref(false)
const employeeDialog = ref(false)
const employeeFormValid = ref(false)
const isEditing = ref(false)
const selectedEmployee = ref(null)
const searchQuery = ref('')
const selectedDepartment = ref('')
const selectedStatus = ref('')
const selectedPosition = ref('')
const formTab = ref('personal')
const detailsDialog = ref(false)
const detailsTab = ref('info')
const selectedEmployeeDetails = ref(null)
const loadingFingerprint = ref(false)
const route = useRoute()

// Dialog states for details
const leaveDialog = ref(false)
const attendanceDialog = ref(false)
const evaluationDialog = ref(false)
const skillDialog = ref(false)
const certificateDialog = ref(false)
const salaryDialog = ref(false)

// Forms for details
const leaveForm = ref({
  startDate: '',
  endDate: '',
  type: '',
  reason: '',
  status: 'قيد المراجعة'
})

const attendanceForm = ref({
  date: '',
  checkIn: '',
  checkOut: '',
  status: 'حاضر'
})

const evaluationForm = ref({
  date: '',
  rating: 5,
  evaluator: '',
  comments: ''
})

const skillForm = ref({
  name: ''
})

const certificateForm = ref({
  name: '',
  issuer: '',
  date: '',
  expiryDate: ''
})

const salaryForm = ref({
  month: '',
  baseSalary: 0,
  bonuses: 0,
  deductions: 0,
  status: 'قيد الدفع'
})

// Options
const leaveTypeOptions = ref([
  { title: 'سنوية', value: 'سنوية' },
  { title: 'مرضية', value: 'مرضية' },
  { title: 'طارئة', value: 'طارئة' },
  { title: 'أمومة', value: 'أمومة' },
  { title: 'أخرى', value: 'أخرى' }
])

const leaveStatusOptions = ref([
  { title: 'قيد المراجعة', value: 'قيد المراجعة' },
  { title: 'موافق', value: 'موافق' },
  { title: 'مرفوض', value: 'مرفوض' }
])

const attendanceStatusOptions = ref([
  { title: 'حاضر', value: 'حاضر' },
  { title: 'غائب', value: 'غائب' },
  { title: 'متأخر', value: 'متأخر' }
])

const salaryStatusOptions = ref([
  { title: 'قيد الدفع', value: 'قيد الدفع' },
  { title: 'مدفوع', value: 'مدفوع' },
  { title: 'معلق', value: 'معلق' }
])

// ========================================
// بيانات النموذج
// ========================================
const employeeForm = ref({
  name: '',
  idNumber: '',
  birthDate: '',
  gender: '',
  maritalStatus: '',
  nationality: '',
  address: '',
  phone: '',
  phone2: '',
  email: '',
  emergencyContact: '',
  emergencyPhone: '',
  department: '',
  position: '',
  hireDate: '',
  contractType: '',
  salary: 0,
  status: 'نشط',
  education: '',
  experience: '',
  notes: '',
  fingerprint: null,
  fingerprintDate: null
})

// ========================================
// بيانات الموظفين
// ========================================
const employees = ref([
  {
    id: 1,
    name: 'أحمد محمد علي',
    idNumber: '1234567890',
    phone: '07701234567',
    email: 'ahmed@example.com',
    department: 'المبيعات',
    position: 'مدير مبيعات',
    salary: 1500000,
    hireDate: '2023-01-15',
    status: 'نشط',
    birthDate: '1990-05-20',
    gender: 'ذكر',
    maritalStatus: 'متزوج',
    nationality: 'عراقي',
    address: 'بغداد - الكرادة',
    phone2: '07901234567',
    emergencyContact: 'زوجته - سارة',
    emergencyPhone: '07701234568',
    contractType: 'دائم',
    education: 'بكالوريوس إدارة أعمال',
    experience: '10',
    notes: 'موظف متميز',
    fingerprint: 'FP_20240115_abc123xyz',
    fingerprintDate: '2024-01-15',
    leaves: [
      { id: 1, startDate: '2024-01-10', endDate: '2024-01-15', days: 5, type: 'سنوية', status: 'موافق' },
      { id: 2, startDate: '2024-03-20', endDate: '2024-03-22', days: 2, type: 'مرضية', status: 'موافق' }
    ],
    attendance: [
      { id: 1, date: '2024-01-15', checkIn: '08:00', checkOut: '17:00', hours: 9, status: 'حاضر' },
      { id: 2, date: '2024-01-16', checkIn: '08:15', checkOut: '17:30', hours: 9.25, status: 'حاضر' }
    ],
    evaluations: [
      { id: 1, date: '2024-01-01', rating: 5, evaluator: 'المدير العام', comments: 'أداء ممتاز' },
      { id: 2, date: '2024-06-01', rating: 4, evaluator: 'المدير العام', comments: 'أداء جيد جداً' }
    ],
    skills: ['إدارة المبيعات', 'التواصل', 'القيادة', 'التخطيط الاستراتيجي'],
    certificates: [
      { name: 'شهادة إدارة المبيعات', issuer: 'معهد التدريب', date: '2023-05-10' },
      { name: 'شهادة القيادة', issuer: 'جامعة بغداد', date: '2022-12-15' }
    ],
    salaryHistory: [
      { id: 1, month: '2024-01', baseSalary: 1500000, bonuses: 200000, deductions: 0, netSalary: 1700000, status: 'مدفوع' },
      { id: 2, month: '2024-02', baseSalary: 1500000, bonuses: 150000, deductions: 50000, netSalary: 1600000, status: 'مدفوع' }
    ]
  },
  {
    id: 2,
    name: 'فاطمة حسن',
    idNumber: '0987654321',
    phone: '07701234568',
    email: 'fatima@example.com',
    department: 'المحاسبة',
    position: 'محاسبة',
    salary: 1200000,
    hireDate: '2023-03-20',
    status: 'نشط',
    birthDate: '1992-08-15',
    gender: 'أنثى',
    maritalStatus: 'عزباء',
    nationality: 'عراقية',
    address: 'بغداد - المنصور',
    phone2: '',
    emergencyContact: 'والدها - حسن',
    emergencyPhone: '07701234569',
    contractType: 'دائم',
    education: 'بكالوريوس محاسبة',
    experience: '5',
    notes: '',
    fingerprint: null,
    fingerprintDate: null,
    leaves: [],
    attendance: [],
    evaluations: [],
    skills: ['المحاسبة', 'البرامج المحاسبية'],
    certificates: [],
    salaryHistory: []
  },
  {
    id: 3,
    name: 'خالد إبراهيم',
    idNumber: '1122334455',
    phone: '07701234569',
    email: 'khalid@example.com',
    department: 'الموارد البشرية',
    position: 'أخصائي موارد بشرية',
    salary: 1300000,
    hireDate: '2023-05-10',
    status: 'في إجازة',
    birthDate: '1988-12-05',
    gender: 'ذكر',
    maritalStatus: 'متزوج',
    nationality: 'عراقي',
    address: 'بغداد - الجادرية',
    phone2: '',
    emergencyContact: '',
    emergencyPhone: '',
    contractType: 'دائم',
    education: 'بكالوريوس إدارة',
    experience: '8',
    notes: '',
    leaves: [],
    attendance: [],
    evaluations: [],
    skills: [],
    certificates: [],
    salaryHistory: []
  }
])

// ========================================
// الخيارات
// ========================================
const statusOptions = ref([
  { title: 'نشط', value: 'نشط' },
  { title: 'في إجازة', value: 'في إجازة' },
  { title: 'معطل', value: 'معطل' }
])

const genderOptions = ref([
  { title: 'ذكر', value: 'ذكر' },
  { title: 'أنثى', value: 'أنثى' }
])

const maritalStatusOptions = ref([
  { title: 'عزباء', value: 'عزباء' },
  { title: 'متزوج', value: 'متزوج' },
  { title: 'مطلق', value: 'مطلق' },
  { title: 'أرمل', value: 'أرمل' }
])

const departments = ref([
  { title: 'المبيعات', value: 'المبيعات' },
  { title: 'المحاسبة', value: 'المحاسبة' },
  { title: 'الموارد البشرية', value: 'الموارد البشرية' },
  { title: 'التسويق', value: 'التسويق' },
  { title: 'التطوير', value: 'التطوير' },
  { title: 'الدعم الفني', value: 'الدعم الفني' }
])

const departmentOptions = computed(() => {
  const depts = [...new Set(employees.value.map(emp => emp.department))]
  return depts.map(dept => ({ title: dept, value: dept }))
})

const positionOptions = computed(() => {
  const positions = [...new Set(employees.value.map(emp => emp.position))]
  return positions.map(pos => ({ title: pos, value: pos }))
})

// ========================================
// رؤوس الجداول
// ========================================
const headers = ref([
  { title: '#', key: 'serial', align: 'center', sortable: false },
  { title: 'اسم الموظف', key: 'name', align: 'right' },
  { title: 'رقم الهوية', key: 'idNumber', align: 'center' },
  { title: 'القسم', key: 'department', align: 'center' },
  { title: 'المنصب', key: 'position', align: 'right' },
  { title: 'رقم الهاتف', key: 'phone', align: 'center' },
  { title: 'البريد الإلكتروني', key: 'email', align: 'right' },
  { title: 'الراتب', key: 'salary', align: 'center' },
  { title: 'تاريخ التوظيف', key: 'hireDate', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' },
  { title: 'الإجراءات', key: 'actions', align: 'center', sortable: false }
])

const leaveHeaders = ref([
  { title: 'تاريخ البداية', key: 'startDate', align: 'center' },
  { title: 'تاريخ النهاية', key: 'endDate', align: 'center' },
  { title: 'عدد الأيام', key: 'days', align: 'center' },
  { title: 'نوع الإجازة', key: 'type', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' }
])

const attendanceHeaders = ref([
  { title: 'التاريخ', key: 'date', align: 'center' },
  { title: 'وقت الدخول', key: 'checkIn', align: 'center' },
  { title: 'وقت الخروج', key: 'checkOut', align: 'center' },
  { title: 'عدد الساعات', key: 'hours', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' }
])

const evaluationHeaders = ref([
  { title: 'التاريخ', key: 'date', align: 'center' },
  { title: 'التقييم', key: 'rating', align: 'center' },
  { title: 'المقيّم', key: 'evaluator', align: 'right' },
  { title: 'ملاحظات', key: 'comments', align: 'right' }
])

const salaryHeaders = ref([
  { title: 'الشهر', key: 'month', align: 'center' },
  { title: 'الراتب الأساسي', key: 'baseSalary', align: 'center' },
  { title: 'المكافآت', key: 'bonuses', align: 'center' },
  { title: 'الخصومات', key: 'deductions', align: 'center' },
  { title: 'صافي الراتب', key: 'netSalary', align: 'center' },
  { title: 'الحالة', key: 'status', align: 'center' }
])

// ========================================
// Computed Properties
// ========================================
const totalEmployees = computed(() => employees.value.length)
const activeEmployees = computed(() => {
  return employees.value.filter(emp => emp.status === 'نشط').length
})
const totalDepartments = computed(() => {
  return new Set(employees.value.map(emp => emp.department)).size
})
const onLeave = computed(() => {
  return employees.value.filter(emp => emp.status === 'في إجازة').length
})
const totalSalaries = computed(() => {
  return employees.value.reduce((sum, emp) => sum + (emp.salary || 0), 0)
})
const averageSalary = computed(() => {
  if (employees.value.length === 0) return 0
  return totalSalaries.value / employees.value.length
})

const filteredEmployees = computed(() => {
  let filtered = [...employees.value]
  
  if (selectedDepartment.value) {
    filtered = filtered.filter(emp => emp.department === selectedDepartment.value)
  }
  
  if (selectedStatus.value) {
    filtered = filtered.filter(emp => emp.status === selectedStatus.value)
  }
  
  if (selectedPosition.value) {
    filtered = filtered.filter(emp => emp.position === selectedPosition.value)
  }
  
  return filtered
})

// ========================================
// Methods
// ========================================
const formatCurrency = (amount) => {
  if (!amount) return '0 IQD'
  const formatted = new Intl.NumberFormat('en-US', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
  return formatted + ' IQD'
}

const formatNumber = (number) => {
  if (!number) return '0'
  return new Intl.NumberFormat('en-US', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(number)
}

const getStatusText = (status) => {
  return status
}

const searchEmployees = () => {
  console.log('البحث عن:', searchQuery.value)
}

const openAddEmployeeDialog = () => {
  employeeDialog.value = true
  isEditing.value = false
  selectedEmployee.value = null
  formTab.value = 'personal'
  employeeForm.value = {
    name: '',
    idNumber: '',
    birthDate: '',
    gender: '',
    maritalStatus: '',
    nationality: '',
    address: '',
    phone: '',
    phone2: '',
    email: '',
    emergencyContact: '',
    emergencyPhone: '',
    department: '',
    position: '',
    hireDate: '',
    contractType: '',
    salary: 0,
    status: 'نشط',
    education: '',
    experience: '',
    notes: '',
    fingerprint: null,
    fingerprintDate: null
  }
}

const closeEmployeeDialog = () => {
  employeeDialog.value = false
  isEditing.value = false
  selectedEmployee.value = null
  formTab.value = 'personal'
  employeeForm.value = {
    name: '',
    idNumber: '',
    birthDate: '',
    gender: '',
    maritalStatus: '',
    nationality: '',
    address: '',
    phone: '',
    phone2: '',
    email: '',
    emergencyContact: '',
    emergencyPhone: '',
    department: '',
    position: '',
    hireDate: '',
    contractType: '',
    salary: 0,
    status: 'نشط',
    education: '',
    experience: '',
    notes: '',
    fingerprint: null,
    fingerprintDate: null
  }
}

const editEmployee = (item) => {
  selectedEmployee.value = item
  isEditing.value = true
  employeeForm.value = { ...item }
  employeeDialog.value = true
  formTab.value = 'personal'
}

const viewEmployeeDetails = (item) => {
  selectedEmployeeDetails.value = item
  detailsDialog.value = true
  detailsTab.value = 'info'
}

const closeDetailsDialog = () => {
  detailsDialog.value = false
  selectedEmployeeDetails.value = null
  detailsTab.value = 'info'
}

const deleteEmployee = (item) => {
  if (confirm(`هل أنت متأكد من حذف الموظف "${item.name}"؟`)) {
    const index = employees.value.findIndex(e => e.id === item.id)
    if (index > -1) {
      employees.value.splice(index, 1)
    }
  }
}

const saveEmployee = () => {
  if (!employeeFormValid.value) return

  if (isEditing.value && selectedEmployee.value) {
    // تحديث الموظف الموجود
    const index = employees.value.findIndex(e => e.id === selectedEmployee.value.id)
    if (index > -1) {
      employees.value[index] = {
        ...employeeForm.value,
        id: selectedEmployee.value.id,
        salary: parseFloat(employeeForm.value.salary),
        leaves: employees.value[index].leaves || [],
        attendance: employees.value[index].attendance || [],
        evaluations: employees.value[index].evaluations || [],
        skills: employees.value[index].skills || [],
        certificates: employees.value[index].certificates || [],
        salaryHistory: employees.value[index].salaryHistory || [],
        fingerprint: employees.value[index].fingerprint || null,
        fingerprintDate: employees.value[index].fingerprintDate || null
      }
    }
  } else {
    // إضافة موظف جديد
    const newEmployee = {
      ...employeeForm.value,
      id: Date.now(),
      salary: parseFloat(employeeForm.value.salary),
      leaves: [],
      attendance: [],
      evaluations: [],
      skills: [],
      certificates: [],
      salaryHistory: [],
      fingerprint: null,
      fingerprintDate: null
    }
    employees.value.push(newEmployee)
  }

  closeEmployeeDialog()
}

const registerFingerprint = async () => {
  loadingFingerprint.value = true
  
  try {
    // محاكاة عملية تسجيل البصمة
    // في التطبيق الحقيقي، سيتم الاتصال بقارئ البصمة هنا
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // توليد معرف بصمة وهمي (في التطبيق الحقيقي سيأتي من قارئ البصمة)
    const fingerprintId = 'FP_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9)
    const fingerprintDate = new Date().toISOString().split('T')[0]
    
    employeeForm.value.fingerprint = fingerprintId
    employeeForm.value.fingerprintDate = fingerprintDate
    
    alert('تم تسجيل البصمة بنجاح!')
  } catch (error) {
    alert('حدث خطأ أثناء تسجيل البصمة. يرجى المحاولة مرة أخرى.')
    console.error('Fingerprint registration error:', error)
  } finally {
    loadingFingerprint.value = false
  }
}

const registerFingerprintFromDetails = async () => {
  if (!selectedEmployeeDetails.value) return
  
  loadingFingerprint.value = true
  
  try {
    // محاكاة عملية تسجيل البصمة
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // توليد معرف بصمة وهمي
    const fingerprintId = 'FP_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9)
    const fingerprintDate = new Date().toISOString().split('T')[0]
    
    selectedEmployeeDetails.value.fingerprint = fingerprintId
    selectedEmployeeDetails.value.fingerprintDate = fingerprintDate
    
    // تحديث في مصفوفة الموظفين الرئيسية
    const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
    if (index > -1) {
      employees.value[index].fingerprint = fingerprintId
      employees.value[index].fingerprintDate = fingerprintDate
    }
    
    alert('تم تسجيل البصمة بنجاح!')
  } catch (error) {
    alert('حدث خطأ أثناء تسجيل البصمة. يرجى المحاولة مرة أخرى.')
    console.error('Fingerprint registration error:', error)
  } finally {
    loadingFingerprint.value = false
  }
}

const deleteFingerprint = () => {
  if (!selectedEmployeeDetails.value) return
  
  if (!confirm('هل أنت متأكد من حذف بصمة الموظف؟')) {
    return
  }
  
  selectedEmployeeDetails.value.fingerprint = null
  selectedEmployeeDetails.value.fingerprintDate = null
  
  // تحديث في مصفوفة الموظفين الرئيسية
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].fingerprint = null
    employees.value[index].fingerprintDate = null
  }
  
  alert('تم حذف البصمة بنجاح')
}

const openAddLeaveDialog = () => {
  if (!selectedEmployeeDetails.value) return
  leaveForm.value = {
    startDate: '',
    endDate: '',
    type: '',
    reason: '',
    status: 'قيد المراجعة'
  }
  leaveDialog.value = true
}

const closeLeaveDialog = () => {
  leaveDialog.value = false
  leaveForm.value = {
    startDate: '',
    endDate: '',
    type: '',
    reason: '',
    status: 'قيد المراجعة'
  }
}

const saveLeave = () => {
  if (!selectedEmployeeDetails.value) return
  if (!leaveForm.value.startDate || !leaveForm.value.endDate || !leaveForm.value.type) {
    alert('يرجى ملء جميع الحقول المطلوبة')
    return
  }

  const start = new Date(leaveForm.value.startDate)
  const end = new Date(leaveForm.value.endDate)
  const days = Math.ceil((end - start) / (1000 * 60 * 60 * 24)) + 1

  const newLeave = {
    id: Date.now(),
    startDate: leaveForm.value.startDate,
    endDate: leaveForm.value.endDate,
    days: days,
    type: leaveForm.value.type,
    reason: leaveForm.value.reason,
    status: leaveForm.value.status
  }

  if (!selectedEmployeeDetails.value.leaves) {
    selectedEmployeeDetails.value.leaves = []
  }
  selectedEmployeeDetails.value.leaves.push(newLeave)

  // Update in main employees array
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].leaves = selectedEmployeeDetails.value.leaves
  }

  closeLeaveDialog()
}

const openAddAttendanceDialog = () => {
  if (!selectedEmployeeDetails.value) return
  const today = new Date().toISOString().split('T')[0]
  attendanceForm.value = {
    date: today,
    checkIn: '08:00',
    checkOut: '17:00',
    status: 'حاضر'
  }
  attendanceDialog.value = true
}

const closeAttendanceDialog = () => {
  attendanceDialog.value = false
  attendanceForm.value = {
    date: '',
    checkIn: '',
    checkOut: '',
    status: 'حاضر'
  }
}

const saveAttendance = () => {
  if (!selectedEmployeeDetails.value) return
  if (!attendanceForm.value.date || !attendanceForm.value.checkIn) {
    alert('يرجى ملء جميع الحقول المطلوبة')
    return
  }

  const checkIn = new Date(`${attendanceForm.value.date}T${attendanceForm.value.checkIn}`)
  const checkOut = attendanceForm.value.checkOut 
    ? new Date(`${attendanceForm.value.date}T${attendanceForm.value.checkOut}`)
    : null
  
  const hours = checkOut 
    ? ((checkOut - checkIn) / (1000 * 60 * 60)).toFixed(2)
    : 0

  const newAttendance = {
    id: Date.now(),
    date: attendanceForm.value.date,
    checkIn: attendanceForm.value.checkIn,
    checkOut: attendanceForm.value.checkOut || null,
    hours: parseFloat(hours),
    status: attendanceForm.value.status
  }

  if (!selectedEmployeeDetails.value.attendance) {
    selectedEmployeeDetails.value.attendance = []
  }
  selectedEmployeeDetails.value.attendance.push(newAttendance)

  // Update in main employees array
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].attendance = selectedEmployeeDetails.value.attendance
  }

  closeAttendanceDialog()
}

const openAddEvaluationDialog = () => {
  if (!selectedEmployeeDetails.value) return
  const today = new Date().toISOString().split('T')[0]
  evaluationForm.value = {
    date: today,
    rating: 5,
    evaluator: '',
    comments: ''
  }
  evaluationDialog.value = true
}

const closeEvaluationDialog = () => {
  evaluationDialog.value = false
  evaluationForm.value = {
    date: '',
    rating: 5,
    evaluator: '',
    comments: ''
  }
}

const saveEvaluation = () => {
  if (!selectedEmployeeDetails.value) return
  if (!evaluationForm.value.date || !evaluationForm.value.evaluator) {
    alert('يرجى ملء جميع الحقول المطلوبة')
    return
  }

  const newEvaluation = {
    id: Date.now(),
    date: evaluationForm.value.date,
    rating: evaluationForm.value.rating,
    evaluator: evaluationForm.value.evaluator,
    comments: evaluationForm.value.comments
  }

  if (!selectedEmployeeDetails.value.evaluations) {
    selectedEmployeeDetails.value.evaluations = []
  }
  selectedEmployeeDetails.value.evaluations.push(newEvaluation)

  // Update in main employees array
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].evaluations = selectedEmployeeDetails.value.evaluations
  }

  closeEvaluationDialog()
}

const openAddSkillDialog = () => {
  if (!selectedEmployeeDetails.value) return
  skillForm.value = {
    name: ''
  }
  skillDialog.value = true
}

const closeSkillDialog = () => {
  skillDialog.value = false
  skillForm.value = {
    name: ''
  }
}

const saveSkill = () => {
  if (!selectedEmployeeDetails.value) return
  if (!skillForm.value.name) {
    alert('يرجى إدخال اسم المهارة')
    return
  }

  if (!selectedEmployeeDetails.value.skills) {
    selectedEmployeeDetails.value.skills = []
  }
  
  if (!selectedEmployeeDetails.value.skills.includes(skillForm.value.name)) {
    selectedEmployeeDetails.value.skills.push(skillForm.value.name)
  }

  // Update in main employees array
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].skills = selectedEmployeeDetails.value.skills
  }

  closeSkillDialog()
}

const deleteSkill = (skillName) => {
  if (!selectedEmployeeDetails.value || !selectedEmployeeDetails.value.skills) return
  const index = selectedEmployeeDetails.value.skills.indexOf(skillName)
  if (index > -1) {
    selectedEmployeeDetails.value.skills.splice(index, 1)
    
    // Update in main employees array
    const empIndex = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
    if (empIndex > -1) {
      employees.value[empIndex].skills = selectedEmployeeDetails.value.skills
    }
  }
}

const openAddCertificateDialog = () => {
  if (!selectedEmployeeDetails.value) return
  certificateForm.value = {
    name: '',
    issuer: '',
    date: '',
    expiryDate: ''
  }
  certificateDialog.value = true
}

const closeCertificateDialog = () => {
  certificateDialog.value = false
  certificateForm.value = {
    name: '',
    issuer: '',
    date: '',
    expiryDate: ''
  }
}

const saveCertificate = () => {
  if (!selectedEmployeeDetails.value) return
  if (!certificateForm.value.name || !certificateForm.value.issuer || !certificateForm.value.date) {
    alert('يرجى ملء جميع الحقول المطلوبة')
    return
  }

  const newCertificate = {
    id: Date.now(),
    name: certificateForm.value.name,
    issuer: certificateForm.value.issuer,
    date: certificateForm.value.date,
    expiryDate: certificateForm.value.expiryDate || null
  }

  if (!selectedEmployeeDetails.value.certificates) {
    selectedEmployeeDetails.value.certificates = []
  }
  selectedEmployeeDetails.value.certificates.push(newCertificate)

  // Update in main employees array
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].certificates = selectedEmployeeDetails.value.certificates
  }

  closeCertificateDialog()
}

const deleteCertificate = (certId) => {
  if (!selectedEmployeeDetails.value || !selectedEmployeeDetails.value.certificates) return
  const index = selectedEmployeeDetails.value.certificates.findIndex(c => c.id === certId)
  if (index > -1) {
    selectedEmployeeDetails.value.certificates.splice(index, 1)
    
    // Update in main employees array
    const empIndex = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
    if (empIndex > -1) {
      employees.value[empIndex].certificates = selectedEmployeeDetails.value.certificates
    }
  }
}

const openAddSalaryDialog = () => {
  if (!selectedEmployeeDetails.value) return
  const today = new Date()
  const month = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, '0')}`
  salaryForm.value = {
    month: month,
    baseSalary: selectedEmployeeDetails.value.salary || 0,
    bonuses: 0,
    deductions: 0,
    status: 'قيد الدفع'
  }
  salaryDialog.value = true
}

const closeSalaryDialog = () => {
  salaryDialog.value = false
  salaryForm.value = {
    month: '',
    baseSalary: 0,
    bonuses: 0,
    deductions: 0,
    status: 'قيد الدفع'
  }
}

const saveSalary = () => {
  if (!selectedEmployeeDetails.value) return
  if (!salaryForm.value.month || !salaryForm.value.baseSalary) {
    alert('يرجى ملء جميع الحقول المطلوبة')
    return
  }

  const netSalary = parseFloat(salaryForm.value.baseSalary) + 
                   parseFloat(salaryForm.value.bonuses || 0) - 
                   parseFloat(salaryForm.value.deductions || 0)

  const newSalary = {
    id: Date.now(),
    month: salaryForm.value.month,
    baseSalary: parseFloat(salaryForm.value.baseSalary),
    bonuses: parseFloat(salaryForm.value.bonuses || 0),
    deductions: parseFloat(salaryForm.value.deductions || 0),
    netSalary: netSalary,
    status: salaryForm.value.status
  }

  if (!selectedEmployeeDetails.value.salaryHistory) {
    selectedEmployeeDetails.value.salaryHistory = []
  }
  selectedEmployeeDetails.value.salaryHistory.push(newSalary)

  // Update in main employees array
  const index = employees.value.findIndex(e => e.id === selectedEmployeeDetails.value.id)
  if (index > -1) {
    employees.value[index].salaryHistory = selectedEmployeeDetails.value.salaryHistory
  }

  closeSalaryDialog()
}

// مراقبة تغيير hash في الرابط للانتقال إلى قسم البصمة
watch(() => route.hash, (newHash) => {
  if (newHash === '#fingerprint' || newHash === '#بصمة') {
    nextTick(() => {
      // إذا كان هناك موظف محدد، فتح تبويب البصمة
      if (selectedEmployeeDetails.value) {
        detailsTab.value = 'fingerprint'
      } else if (employees.value.length > 0) {
        // فتح أول موظف وتبويب البصمة
        viewEmployeeDetails(employees.value[0])
        nextTick(() => {
          detailsTab.value = 'fingerprint'
        })
      }
    })
  }
}, { immediate: true })

onMounted(() => {
  console.log('صفحة الموارد البشرية المتكاملة جاهزة')
  
  // التحقق من وجود hash في الرابط للانتقال إلى تبويب البصمة
  const hash = route.hash || window.location.hash
  if (hash === '#fingerprint' || hash === '#بصمة') {
    // الانتقال إلى صفحة التفاصيل إذا كان هناك موظف محدد
    // أو فتح تبويب البصمة في أول موظف
    if (employees.value.length > 0) {
      setTimeout(() => {
        viewEmployeeDetails(employees.value[0])
        nextTick(() => {
          detailsTab.value = 'fingerprint'
        })
      }, 500)
    }
  }
})
</script>

<style scoped>
/* ========================================
   تنسيقات الصفحة الأساسية
   ======================================== */

.data-page {
  background: #ffffff !important;
  min-height: 100vh;
  padding: 0;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
}

.human-resources-page {
  background: #ffffff !important;
  min-height: 100vh;
  padding: 0;
  overflow-x: hidden;
  width: 100%;
  box-sizing: border-box;
}

.centered-content {
  max-width: 1400px !important;
  margin: 0 auto !important;
  width: 100% !important;
  box-sizing: border-box !important;
}

.fullscreen-content {
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
  box-sizing: border-box;
  overflow-x: hidden;
  padding: 0 20px;
}

/* Header Styles - نفس تنسيق صفحة المهندسين */
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

.header-content {
  padding: 12px 16px !important;
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-height: auto !important;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.1) 0%, rgba(21, 101, 192, 0.05) 100%);
  backdrop-filter: blur(10px);
  position: relative;
  z-index: 3;
  animation: fadeInUp 1.2s ease-out 0.3s both;
  max-width: calc(100vw - 320px);
  margin: 0 auto;
  gap: 1rem;
}

.header-left-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
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

.header-text {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
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
}

.stats-row {
  margin-bottom: 30px;
  width: 100%;
  overflow-x: hidden;
  display: flex;
  flex-wrap: wrap;
  gap: 0;
  justify-content: center;
}

.stats-row .v-col {
  padding: 8px;
  display: flex;
  align-items: stretch;
  flex: 1 1 auto;
  min-width: 0;
}

.stats-section {
  margin-bottom: 30px;
  padding: 0 20px;
}

/* Modern Statistics Cards - نفس تصميم صفحة إدارة المهام */
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
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 250, 252, 0.95) 100%);
  border-radius: 12px;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid rgba(226, 232, 240, 0.8);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 12px 10px;
  position: relative;
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, transparent 0%, rgba(99, 102, 241, 0.5) 50%, transparent 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-6px) scale(1.02);
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.15), 0 4px 16px rgba(99, 102, 241, 0.2);
  border-color: rgba(99, 102, 241, 0.4);
}

.stat-card:hover::before {
  opacity: 1;
}

.filters-card {
  border-radius: 12px !important;
  overflow: hidden !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1) !important;
  margin-bottom: 24px;
}

/* رأس شريط الفلترة */
.filters-header-new {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  border-radius: 12px 12px 0 0 !important;
  padding: 0.5rem 0.75rem !important;
  border-bottom: 1px solid #1e40af !important;
  position: relative !important;
  overflow: hidden !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.3) !important;
}

.filters-header-new,
.filters-header-new *,
.filters-header-new div,
.filters-header-new span {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.filters-header-new::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, #93c5fd, #60a5fa, #3b82f6);
  border-radius: 12px 12px 0 0;
}

.filters-header-title {
  color: #ef4444 !important;
  font-weight: 700 !important;
  font-size: 0.95rem !important;
  letter-spacing: 0.3px !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3) !important;
}

.filters-header-new :deep(.v-icon) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  fill: #ffffff !important;
}

/* محتوى الفلترة */
.filters-content {
  padding: 0.5rem 0.75rem !important;
  background: #ffffff !important;
  background-color: #ffffff !important;
  border-radius: 0 0 12px 12px !important;
}

.filters-content,
.filters-content *,
.filters-content .v-row,
.filters-content .v-col {
  background: #ffffff !important;
  background-color: #ffffff !important;
}

/* حقول البحث والفلترة الجديدة */
.search-field-new :deep(.v-field),
.filter-field-new :deep(.v-field),
.search-field-new :deep(.v-field__input),
.filter-field-new :deep(.v-field__input),
.search-field-new :deep(.v-input),
.filter-field-new :deep(.v-input) {
  background: #ffffff !important;
  background-color: #ffffff !important;
}

.search-field-new :deep(.v-field),
.filter-field-new :deep(.v-field) {
  border-radius: 8px !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid #fca5a5 !important;
  transition: all 0.3s ease !important;
}

.search-field-new :deep(.v-field:hover),
.filter-field-new :deep(.v-field:hover) {
  border-color: #f87171 !important;
  box-shadow: 0 2px 6px rgba(239, 68, 68, 0.1) !important;
}

.search-field-new :deep(.v-field--focused),
.filter-field-new :deep(.v-field--focused) {
  border-color: #f87171 !important;
  box-shadow: 0 0 0 2px rgba(239, 68, 68, 0.1) !important;
}

.search-field-new :deep(.v-field__input),
.filter-field-new :deep(.v-field__input) {
  color: #1e293b !important;
  font-size: 0.875rem !important;
  padding: 8px 12px !important;
}

.search-field-new :deep(.v-label),
.filter-field-new :deep(.v-label) {
  color: #1e40af !important;
  font-weight: 500 !important;
  font-size: 0.875rem !important;
}

.search-field-new :deep(.v-field--focused .v-label),
.filter-field-new :deep(.v-field--focused .v-label) {
  color: #1e40af !important;
}

.search-field-new :deep(.v-field__prepend-inner .v-icon) {
  color: #1e40af !important;
}

.search-field-new :deep(.v-field__input::placeholder) {
  color: #94a3b8 !important;
  opacity: 1 !important;
}

.filter-field-new :deep(.v-field__append-inner .v-icon) {
  color: #1e40af !important;
}

/* ========================================
   تأثيرات بصرية
   ======================================== */

.glass-effect {
  background: rgba(255, 255, 255, 0.1) !important;
  backdrop-filter: blur(20px) !important;
  -webkit-backdrop-filter: blur(20px) !important;
}

.gradient-animation {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  animation: gradient-shift 3s ease infinite;
  background-size: 200% 200%;
}

@keyframes gradient-shift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.text-glow {
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3), 0 0 30px rgba(255, 255, 255, 0.2);
}

.fade-in {
  animation: fadeIn 0.6s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.star-twinkle {
  animation: star-twinkle 2s ease-in-out infinite;
}

@keyframes star-twinkle {
  0%, 100% { 
    transform: scale(1) rotate(0deg); 
    opacity: 1;
  }
  50% { 
    transform: scale(1.1) rotate(180deg); 
    opacity: 0.8;
  }
}

.hover-lift {
  transition: all 0.3s ease;
}

.hover-lift:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.card-glow {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.card-glow:hover {
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
}

.smooth-transition {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.icon-glow {
  filter: drop-shadow(0 2px 6px rgba(0, 0, 0, 0.15));
  transition: all 0.3s ease;
}

.stat-card:hover .icon-glow {
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.25));
}

.full-width {
  width: 100% !important;
}

/* ========================================
   تنسيقات الجدول
   ======================================== */

.data-table-card {
  background: rgba(255, 249, 249, 0.95);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin-top: 24px;
}

.table-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  padding: 8px 12px !important;
  border-radius: 8px 8px 0 0 !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.2) !important;
  margin: 0 !important;
  margin-bottom: 16px !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-align: right !important;
  display: flex !important;
  min-height: 36px !important;
  align-items: center !important;
  justify-content: space-between !important;
}

.indigo-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  font-weight: 600!important;
}

.title-text {
  color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.95rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.3px !important;
}

.project-table {
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  width: 100% !important;
  display: block !important;
  visibility: visible !important;
  opacity: 1 !important;
  margin-top: 16px !important;
}

.project-table .v-data-table__th {
  background: #ffffff !important;
  color: #1a1a1a !important;
  font-weight: 500 !important;
  font-size: 0.55rem !important;
  text-align: center !important;
  padding: 0.1rem 0.2rem !important;
  border-bottom: 1px solid rgba(148, 163, 184, 0.2) !important;
  text-shadow: none !important;
  letter-spacing: 0.1px !important;
  min-height: 16px !important;
}

/* تصغير عنوان "اسم الموظف" تحديداً */
.project-table .v-data-table__th:nth-child(2) {
  font-size: 0.6rem !important;
  font-weight: 400 !important;
}

.project-table .v-data-table__td {
  padding: 0.15rem 0.3rem !important;
  border-bottom: 1px solid #e5e7eb !important;
  text-align: center !important;
  background: #ffffff !important;
  color: #000000 !important;
  font-weight: 400 !important;
  font-size: 0.65rem !important;
  vertical-align: middle !important;
  min-height: 24px !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(even) {
  background: linear-gradient(135deg, #faf5ff 0%, #f3e8ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:nth-child(odd) {
  background: linear-gradient(135deg, #ffffff 0%, #faf5ff 100%) !important;
}

.project-table .v-data-table__wrapper tbody tr:hover {
  background: linear-gradient(135deg, #f3e8ff 0%, #e9d5ff 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.15) !important;
  transition: all 0.3s ease !important;
}

.project-table .serial-number {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.55rem !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 1px 4px !important;
  border-radius: 3px !important;
  display: inline-block !important;
  min-width: 20px !important;
  text-align: center !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1) !important;
}

.project-table .project-name {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.6rem !important;
  text-align: right !important;
  padding: 1px 4px !important;
  border-radius: 3px !important;
  display: inline-block !important;
  max-width: 150px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.project-table .date-text {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 0.6rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
}

.project-table .cost-text {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 0.6rem !important;
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: ltr !important;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%) !important;
  padding: 2px 6px !important;
  border-radius: 4px !important;
  display: inline-block !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08) !important;
  border: 1px solid #d1d5db !important;
  white-space: nowrap !important;
  min-width: 70px !important;
  text-align: center !important;
}

.project-table .v-chip {
  font-weight: 500 !important;
  font-size: 0.5rem !important;
  padding: 0 4px !important;
  height: 16px !important;
  border-radius: 4px !important;
  color: #000000 !important;
  background: #f3f4f6 !important;
  border: 1px solid #e5e7eb !important;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08) !important;
}

.action-buttons {
  display: flex;
  gap: 4px;
  align-items: center;
  justify-content: center;
}

.action-btn {
  min-width: 28px !important;
  height: 28px !important;
  border-radius: 6px !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.action-btn:hover {
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15) !important;
}

/* ========================================
   تنسيق الحقول والنماذج
   ======================================== */

.image-style-dialog {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%);
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(5, 150, 105, 0.3);
}

.dialog-header {
  background: linear-gradient(135deg, #059669, #10b981);
  padding: 20px;
  color: white;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-title {
  font-size: 1.5rem;
  font-weight: 700;
}

.close-btn {
  color: white !important;
}

.dialog-body {
  padding: 30px;
  background: white !important;
}

.form-tabs {
  margin-bottom: 20px;
}

.form-fields {
  margin-bottom: 20px;
}

.form-field {
  margin-bottom: 16px;
}

.form-field .v-field__input {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  background: white !important;
  padding: 12px 16px !important;
}

.form-field .v-field__input input,
.form-field .v-field__input textarea {
  color: #000000 !important;
  font-weight: 600 !important;
  background: white !important;
}

.form-field .v-label,
.form-field .v-field__label,
.form-field .v-label--active,
.form-field .v-field__label--active,
.form-field .v-label--floating,
.form-field .v-field__label--floating {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
  background: white !important;
  padding: 0 8px !important;
}

.form-field .v-field--focused .v-label,
.form-field .v-field--focused .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}

.form-field .v-field__outline {
  border-color: #d1d5db !important;
  border-width: 2px !important;
}

.form-field .v-field--focused .v-field__outline {
  border-color: #6b7280 !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(107, 114, 128, 0.1) !important;
}

.form-field .v-field:hover .v-field__outline {
  border-color: #9ca3af !important;
}

/* تنسيق خاص للقوائم المنسدلة */
.form-field .v-select .v-field__input {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
}

.form-field .v-select .v-field__outline {
  border-color: #d1d5db !important;
  border-width: 2px !important;
}

.form-field .v-select .v-label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  background: white !important;
}

.form-field .v-select .v-field--focused .v-field__outline {
  border-color: #6b7280 !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(107, 114, 128, 0.1) !important;
}

/* تنسيق القوائم المنسدلة المفتوحة - ألوان سوداء */
.form-field .black-list .v-menu__content,
.form-field .black-list .v-overlay__content {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.form-field .black-list .v-menu__content .v-list-item,
.form-field .black-list .v-overlay__content .v-list-item {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  background: white !important;
}

.form-field .black-list .v-menu__content .v-list-item:hover,
.form-field .black-list .v-overlay__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.form-field .black-list .v-menu__content .v-list-item--active,
.form-field .black-list .v-overlay__content .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

/* تنسيق شامل لجميع القوائم المنسدلة في الحوار */
.image-style-dialog .v-menu__content .v-list-item,
.image-style-dialog .v-overlay__content .v-list-item {
  color: #000000 !important;
  background: white !important;
}

.image-style-dialog .v-menu__content .v-list-item:hover,
.image-style-dialog .v-overlay__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.image-style-dialog .v-menu__content .v-list-item--active,
.image-style-dialog .v-overlay__content .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
}

/* تنسيقات إضافية قوية للتسميات - ضمان الوضوح */
.image-style-dialog .form-field .v-text-field .v-label,
.image-style-dialog .form-field .v-select .v-label,
.image-style-dialog .form-field .v-textarea .v-label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}

.image-style-dialog .form-field .v-text-field .v-field--focused .v-label,
.image-style-dialog .form-field .v-select .v-field--focused .v-label,
.image-style-dialog .form-field .v-textarea .v-field--focused .v-label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}

.image-style-dialog .form-field .v-text-field .v-field__input input,
.image-style-dialog .form-field .v-textarea .v-field__input textarea {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.dialog-actions {
  padding: 20px 30px;
  background: #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>

<style scoped>
/* تنسيقات قوية للتسميات باستخدام :deep() */
.image-style-dialog :deep(.v-label),
.image-style-dialog :deep(.v-field__label) {
  color: #000000 !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
  opacity: 1 !important;
  background: white !important;
}

.image-style-dialog :deep(.v-label--active),
.image-style-dialog :deep(.v-field__label--active),
.image-style-dialog :deep(.v-label--floating),
.image-style-dialog :deep(.v-field__label--floating) {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  background: white !important;
}

.image-style-dialog :deep(.v-field--focused .v-label),
.image-style-dialog :deep(.v-field--focused .v-field__label) {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
  background: white !important;
}

.image-style-dialog :deep(.v-field__input),
.image-style-dialog :deep(.v-field__input input),
.image-style-dialog :deep(.v-field__input textarea) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  background: white !important;
}
</style>

<style>
/* تنسيقات شاملة للتسميات بدون scope */
.image-style-dialog .v-label,
.image-style-dialog .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}

.image-style-dialog .v-field--focused .v-label,
.image-style-dialog .v-field--focused .v-field__label {
  color: #000000 !important;
  font-weight: 800 !important;
  opacity: 1 !important;
}

.image-style-dialog .v-field__input,
.image-style-dialog .v-field__input input,
.image-style-dialog .v-field__input textarea {
  color: #000000 !important;
  font-weight: 600 !important;
  background: white !important;
}

/* ========================================
   تنسيق نافذة التفاصيل
   ======================================== */

.details-dialog {
  border-radius: 20px;
  overflow: hidden;
}

.details-header {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%);
}

.details-body {
  padding: 30px;
  background: white !important;
}

.details-tabs {
  margin-bottom: 20px;
}

.info-content {
  padding: 10px 0;
}

.info-card {
  margin-bottom: 20px;
  border-radius: 12px;
}

.info-card-title {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%);
  color: white !important;
  padding: 15px 20px !important;
  font-weight: 700 !important;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 12px 0;
  border-bottom: 1px solid #e5e7eb;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  font-weight: 700;
  color: #6b7280;
}

.info-value {
  font-weight: 600;
  color: #1a1a1a;
}

/* تحسين وضوح النصوص في قسم البصمة */
.fingerprint-status-section .info-label,
.fingerprint-info .info-label {
  color: #374151 !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
}

.fingerprint-status-section .info-value,
.fingerprint-info .info-value {
  color: #111827 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.section-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #1a1a1a;
  margin-bottom: 20px;
}

.details-table {
  margin-top: 20px;
}

.rating-display {
  display: flex;
  align-items: center;
  justify-content: center;
}

.certificate-item {
  padding: 12px;
  background: #f9fafb;
  border-radius: 8px;
  border-left: 4px solid #059669;
}

.gap-2 {
  gap: 8px;
}

/* ========================================
   تنسيق القياسات والأرقام
   ======================================== */

.stat-icon {
  margin-bottom: 8px !important;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  border-radius: 12px;
  background: linear-gradient(135deg, rgba(99, 102, 241, 0.1) 0%, rgba(139, 92, 246, 0.1) 100%);
  transition: all 0.3s ease;
}

.stat-card:hover .stat-icon {
  transform: scale(1.1) rotate(5deg);
  background: linear-gradient(135deg, rgba(99, 102, 241, 0.2) 0%, rgba(139, 92, 246, 0.2) 100%);
}

.stat-card .stat-icon .v-icon {
  transition: all 0.3s ease;
}

.stat-card:hover .stat-icon .v-icon {
  transform: scale(1.15);
}

.stat-number {
  font-size: 1.1rem !important;
  line-height: 1.2 !important;
  font-weight: 800 !important;
  text-align: center !important;
  margin: 4px 0 !important;
  letter-spacing: -0.5px;
}

.stat-currency {
  font-size: 1.3rem !important;
  line-height: 1.3 !important;
  font-weight: 800 !important;
  text-align: center !important;
  margin: 4px 0 !important;
  letter-spacing: -0.5px;
}

.stat-card h3 {
  min-height: 45px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  flex-wrap: wrap !important;
  text-align: center !important;
  padding: 2px 4px !important;
  margin: 0 !important;
}

.stat-card p {
  margin: 0 !important;
  font-size: 0.85rem !important;
  font-weight: 600 !important;
  letter-spacing: 0.3px;
  opacity: 0.9;
}

.stat-card {
  min-height: 150px !important;
  display: flex !important;
  flex-direction: column !important;
  justify-content: flex-start !important;
  align-items: center !important;
  text-align: center !important;
  gap: 4px;
}

/* تحسين وضوح النص في مربع متوسط الراتب */
.average-salary-value {
  color: #7c3aed !important;
  font-weight: 900 !important;
  text-shadow: 0 2px 4px rgba(124, 58, 237, 0.2) !important;
  letter-spacing: 0.5px !important;
}

.average-salary-label {
  color: #6d28d9 !important;
  font-weight: 700 !important;
  text-shadow: 0 1px 2px rgba(109, 40, 217, 0.15) !important;
  opacity: 1 !important;
}

.average-salary-icon .v-icon {
  color: #7c3aed !important;
  filter: drop-shadow(0 2px 4px rgba(124, 58, 237, 0.3)) !important;
}

/* Responsive Design */
@media (max-width: 1200px) {
  .centered-content,
  .fullscreen-content {
    max-width: 100% !important;
    padding: 0 15px !important;
  }
}

.fingerprint-card {
  border-radius: 12px;
  border: 2px solid rgba(199, 210, 254, 0.5);
  background: linear-gradient(135deg, rgba(238, 242, 255, 0.5) 0%, rgba(224, 231, 255, 0.3) 100%);
}

.fingerprint-title {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%);
  color: white !important;
  font-weight: 700 !important;
}

.fingerprint-status {
  text-align: center;
  padding: 20px 0;
}

.fingerprint-hint {
  text-align: center;
  font-style: italic;
  color: #6366f1 !important;
}

/* ألوان البصمة المسجلة - بنفسجي/نيلي */
.fingerprint-registered {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.3) !important;
}

.fingerprint-registered-small {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: white !important;
}

.fingerprint-status-chip {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: white !important;
}

/* ألوان البصمة غير المسجلة - بنفسجي فاتح/أزرق */
.fingerprint-pending {
  background: linear-gradient(135deg, #6366f1 0%, #818cf8 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3) !important;
}

.fingerprint-pending-small {
  background: linear-gradient(135deg, #6366f1 0%, #818cf8 100%) !important;
  color: white !important;
}

/* أزرار البصمة */
.fingerprint-btn {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(124, 58, 237, 0.3) !important;
  transition: all 0.3s ease !important;
}

.fingerprint-btn:hover {
  background: linear-gradient(135deg, #6d28d9 0%, #7c3aed 100%) !important;
  box-shadow: 0 6px 16px rgba(124, 58, 237, 0.4) !important;
  transform: translateY(-2px) !important;
}

.fingerprint-btn:disabled {
  background: linear-gradient(135deg, #a5b4fc 0%, #c7d2fe 100%) !important;
  opacity: 0.6 !important;
}

.fingerprint-delete-btn {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3) !important;
  transition: all 0.3s ease !important;
}

.fingerprint-delete-btn:hover {
  background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%) !important;
  box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4) !important;
  transform: translateY(-2px) !important;
}

.fingerprint-chip {
  transition: all 0.3s ease !important;
}

.fingerprint-chip:hover {
  transform: scale(1.05) !important;
  box-shadow: 0 6px 16px rgba(124, 58, 237, 0.4) !important;
}

.fingerprint-management-card {
  border-radius: 12px;
  border: 2px solid rgba(199, 210, 254, 0.5);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(238, 242, 255, 0.8) 100%);
}

.fingerprint-card-title {
  background: linear-gradient(135deg, #7c3aed 0%, #8b5cf6 100%);
  color: white !important;
  font-weight: 700 !important;
  padding: 20px !important;
}

.fingerprint-content {
  padding: 10px 0;
}

.fingerprint-status-section {
  text-align: center;
  padding: 20px;
  background: linear-gradient(135deg, rgba(238, 242, 255, 0.5) 0%, rgba(224, 231, 255, 0.3) 100%);
  border-radius: 12px;
  border: 1px solid rgba(199, 210, 254, 0.3);
}

.fingerprint-status-section .section-title {
  color: #1a1a1a !important;
  font-weight: 800 !important;
  font-size: 1.6rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.05) !important;
}

.fingerprint-info {
  text-align: right;
  padding: 15px;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 8px;
  border-left: 4px solid #7c3aed;
}

.fingerprint-info .info-label {
  color: #374151 !important;
  font-weight: 800 !important;
  font-size: 1rem !important;
}

.fingerprint-info .info-value {
  color: #111827 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
}

.fingerprint-actions {
  padding: 20px 0;
}

@media (max-width: 768px) {
  .human-resources-page {
    padding: 10px !important;
  }

  .stats-row .v-col {
    padding: 6px !important;
  }
  
  .stat-card {
    min-height: 140px !important;
    padding: 10px 8px !important;
  }
  
  .stat-icon {
    width: 48px !important;
    height: 48px !important;
    margin-bottom: 6px !important;
  }
  
  .stat-number {
    font-size: 1rem !important;
  }
  
  .stat-currency {
    font-size: 1.1rem !important;
  }
  
  .stat-card p {
    font-size: 0.75rem !important;
  }
  
  .stat-card h3 {
    min-height: 35px !important;
  }

  .page-title {
    font-size: 2rem;
  }
  
  .page-subtitle {
    font-size: 1rem;
  }
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

@keyframes iconBounce {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-8px);
  }
}

</style>
