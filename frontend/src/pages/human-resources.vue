<template>
  <v-container class="fill-height data-page human-resources-page" fluid>
    <div class="fullscreen-content centered-content">
      <!-- Header Section -->
      <div class="page-header glass-effect gradient-animation">
        <span class="page-icon star-twinkle">👥</span>
        <h1 class="page-title text-glow fade-in">الموارد البشرية</h1>
        <p class="page-subtitle fade-in">إدارة وتتبع جميع الموظفين والموارد البشرية</p>
      </div>

      <!-- Summary Cards - Expanded -->
      <v-row class="mb-6 stats-row full-width">
        <v-col cols="12" sm="6" md="4" lg="4" xl="4">
          <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-2 icon-glow">
              <v-icon size="48" color="primary">mdi-account-group</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-primary mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(totalEmployees) || '0' }}</h3>
            <p class="text-subtitle-1 text-primary mb-0">إجمالي الموظفين</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="4" xl="4">
          <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-2 icon-glow">
              <v-icon size="48" color="success">mdi-account-check</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-success mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(activeEmployees) || '0' }}</h3>
            <p class="text-subtitle-1 text-success mb-0">موظفين نشطين</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="4" xl="4">
          <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-2 icon-glow">
              <v-icon size="48" color="info">mdi-office-building</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-info mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(totalDepartments) || '0' }}</h3>
            <p class="text-subtitle-1 text-info mb-0">إجمالي الأقسام</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="4" xl="4">
          <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-2 icon-glow">
              <v-icon size="48" color="warning">mdi-calendar-clock</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-warning mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr;">{{ formatNumber(onLeave) || '0' }}</h3>
            <p class="text-subtitle-1 text-warning mb-0">في إجازة</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="4" xl="4">
          <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-2 icon-glow">
              <v-icon size="48" color="error">mdi-currency-usd</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold text-error mb-2 stat-number" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; font-size: 1.1rem !important;">{{ formatCurrency(totalSalaries) || '0 د.ع' }}</h3>
            <p class="text-subtitle-1 text-error mb-0">إجمالي الرواتب</p>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="4" lg="4" xl="4">
          <v-card class="stat-card pa-3 pb-4 text-center hover-lift card-glow smooth-transition" elevation="2">
            <div class="stat-icon mb-2 icon-glow average-salary-icon">
              <v-icon size="48" color="#7c3aed">mdi-chart-line</v-icon>
            </div>
            <h3 class="text-h3 font-weight-bold mb-2 stat-number average-salary-value" style="font-family: 'Arial', 'Helvetica', sans-serif; direction: ltr; font-size: 1.1rem !important;">{{ formatCurrency(averageSalary) || '0 د.ع' }}</h3>
            <p class="text-subtitle-1 mb-0 average-salary-label">متوسط الراتب</p>
          </v-card>
        </v-col>
      </v-row>

      <!-- Search Bar -->
      <v-card class="search-card mb-4" elevation="2">
        <v-card-text class="pa-4">
          <v-row class="align-center">
            <v-col cols="12" md="3">
              <v-text-field
                v-model="searchQuery"
                label="البحث في الموظفين..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="search-field"
                style="background: #f5f5f5;"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-select
                v-model="selectedDepartment"
                :items="departmentOptions"
                label="القسم"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
                class="black-list"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-select
                v-model="selectedStatus"
                :items="statusOptions"
                label="الحالة"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-select
                v-model="selectedPosition"
                :items="positionOptions"
                label="المنصب"
                variant="outlined"
                density="comfortable"
                clearable
                hide-details
              />
            </v-col>
            <v-col cols="12" md="1">
              <v-btn
                color="primary"
                variant="elevated"
                size="large"
                class="search-btn"
                @click="searchEmployees"
                block
              >
                بحث
              </v-btn>
            </v-col>
            <v-col cols="12" md="2">
              <v-btn
                color="success"
                variant="elevated"
                size="large"
                class="add-expense-btn"
                @click="openAddEmployeeDialog"
                block
              >
                <v-icon class="me-2">mdi-plus</v-icon>
                إضافة موظف
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>

      <!-- Employees Table -->
      <v-card class="data-table-card" elevation="2">
        <v-card-title class="table-title indigo-title">
          <span class="title-text">قائمة الموظفين</span>
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
                <v-icon size="16">mdi-eye</v-icon>
              </v-btn>
              <v-btn
                size="small"
                color="success"
                variant="text"
                @click="editEmployee(item)"
                icon
                class="action-btn"
                title="تعديل"
              >
                <v-icon size="16">mdi-pencil</v-icon>
              </v-btn>
              <v-btn
                size="small"
                color="error"
                variant="text"
                @click="deleteEmployee(item)"
                icon
                class="action-btn"
                title="حذف"
              >
                <v-icon size="16">mdi-delete</v-icon>
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
import { toast } from 'vue3-toastify'

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
  if (!amount) return '0 د.ع'
  const formatted = new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
  return formatted + ' د.ع'
}

const formatNumber = (number) => {
  if (!number) return '0'
  return new Intl.NumberFormat('ar-IQ', {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(number)
}

const getStatusText = (status) => {
  return status
}

const searchEmployees = () => {
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
    
    toast.success('تم تسجيل البصمة بنجاح!')
  } catch (error) {
    toast.error('حدث خطأ أثناء تسجيل البصمة. يرجى المحاولة مرة أخرى.')
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
    
    toast.success('تم تسجيل البصمة بنجاح!')
  } catch (error) {
    toast.error('حدث خطأ أثناء تسجيل البصمة. يرجى المحاولة مرة أخرى.')
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
  
  toast.success('تم حذف البصمة بنجاح')
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
    toast.warning('يرجى ملء جميع الحقول المطلوبة')
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
    toast.warning('يرجى ملء جميع الحقول المطلوبة')
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
    toast.warning('يرجى ملء جميع الحقول المطلوبة')
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
    toast.warning('يرجى إدخال اسم المهارة')
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
    toast.warning('يرجى ملء جميع الحقول المطلوبة')
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
    toast.warning('يرجى ملء جميع الحقول المطلوبة')
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


<style>
/* Import page styles */
@import './styles/human-resources.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
