<template>
  <div class="work-days-page">
    <!-- Page Title Header -->
    <div class="page-title-header">
      <div class="header-left-section">
        <v-icon class="title-icon">mdi-calendar-clock</v-icon>
        <h1 class="page-title-text" style="color: #ffffff !important;">ايام العمل</h1>
      </div>
      <div class="header-right-section">
        <v-btn icon="mdi-arrow-right" @click="goBack" class="nav-btn">
          <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
      </div>
    </div>

    <!-- Add Work Day Modal - Clean Form Style -->
    <v-dialog v-model="showAddForm" max-width="900" scrollable persistent>
      <v-card class="clean-dialog-card clean-form-card">
        <!-- Header Section -->
        <v-card-title class="clean-dialog-header clean-form-header">
          <h2 class="clean-form-title">
            {{ isEditing ? 'تعديل يوم العمل' : 'معلومات يوم العمل' }}
          </h2>
        </v-card-title>

        <!-- Form Content -->
        <v-card-text class="clean-form-content">
          <p class="clean-form-instruction">
            لإتمام {{ isEditing ? 'تعديل' : 'إضافة' }} يوم العمل، يرجى توفير المعلومات التالية. يرجى ملاحظة أن جميع الحقول المميزة بعلامة النجمة (*) مطلوبة.
          </p>

          <v-form ref="form" v-model="formValid">
            <!-- الصف الأول: التاريخ -->
            <v-row class="clean-form-row">
              <v-col cols="12" md="6">
                <v-menu
                  v-model="dateMenu"
                  :close-on-content-click="false"
                  location="bottom"
                >
                  <template v-slot:activator="{ props }">
                    <v-text-field
                      v-model="formattedDate"
                      label="التاريخ *"
                      readonly
                      v-bind="props"
                      variant="outlined"
                      density="comfortable"
                      :rules="[v => !!v || 'التاريخ مطلوب']"
                      required
                      hide-details="auto"
                      prepend-inner-icon="mdi-calendar"
                    />
                  </template>
                  <v-date-picker
                    v-model="selectedDate"
                    @update:model-value="onDateSelected"
                    color="primary"
                    theme="light"
                  />
                </v-menu>
              </v-col>

              <v-col cols="12" md="6">
                <v-select
                  v-model="workDayForm.status"
                  label="الحالة"
                  :items="statusOptions"
                  variant="outlined"
                  density="comfortable"
                  hide-details="auto"
                />
              </v-col>
            </v-row>

            <!-- الصف الثاني: تصنيف العمل -->
            <v-row class="clean-form-row">
              <v-col cols="12">
                <v-select
                  v-model="workDayForm.workSubCategoryId"
                  label="تصنيف العمل"
                  :items="workSubCategories"
                  :loading="loadingSubCategories"
                  variant="outlined"
                  density="comfortable"
                  hide-details="auto"
                  clearable
                />
              </v-col>
            </v-row>

            <!-- الصف الثالث: الوصف -->
            <v-row class="clean-form-row">
              <v-col cols="12">
                <v-textarea
                  v-model="workDayForm.description"
                  label="الوصف"
                  variant="outlined"
                  rows="3"
                  density="comfortable"
                  hide-details="auto"
                />
              </v-col>
            </v-row>

            <!-- الصف الثالث: الملاحظات -->
            <v-row class="clean-form-row">
              <v-col cols="12">
                <v-textarea
                  v-model="workDayForm.notes"
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

        <v-card-actions class="clean-form-actions">
          <v-spacer />
          <v-btn
            class="clean-form-cancel-btn"
            variant="outlined"
            @click="closeAddForm"
          >
            إلغاء
          </v-btn>
          <v-btn
            class="clean-form-continue-btn"
            variant="elevated"
            :disabled="!formValid || (isEditing && !hasFormChanges)"
            :loading="saving"
            @click="saveWorkDay"
          >
            {{ isEditing ? 'تحديث يوم العمل' : 'حفظ يوم العمل' }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Data Table -->
    <v-card class="data-table-card card-glow smooth-transition centered-table" elevation="2">
      <v-card-title class="d-flex align-center justify-space-between" style="padding: 10px 16px !important;">
        <div class="d-flex align-center gap-2" style="flex: 1;">
          <v-icon class="me-2" style="color: #ffffff;" size="20">mdi-calendar-clock</v-icon>
          <span class="text-h6 font-weight-black" style="color: #ffffff; font-family: 'Arial', 'Helvetica', sans-serif; text-shadow: 0 3px 6px rgba(0, 0, 0, 0.3), 0 1px 3px rgba(0, 0, 0, 0.2); letter-spacing: 0.5px; font-size: 1.25rem !important;">أيام العمل</span>
          <v-chip class="ms-3 work-days-count-chip" size="x-small" variant="elevated">{{ workDaysData.length || 0 }}</v-chip>
          <div class="table-search-box">
            <v-text-field
              v-model="searchQuery"
              placeholder="البحث..."
              variant="outlined"
              density="compact"
              hide-details
              prepend-inner-icon="mdi-magnify"
              class="table-search-input"
              style="max-width: 250px;"
            />
          </div>
        </div>
        <v-btn
          class="add-button btn-glow light-sweep smooth-transition"
          @click="addWorkDay"
          elevation="2"
          color="primary"
          size="small"
          style="background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important; height: 36px !important;"
        >
          <v-icon class="me-2 icon-glow" size="18">mdi-plus</v-icon>
          إضافة يوم عمل جديد
        </v-btn>
      </v-card-title>

      <v-card-text class="pa-0">
        <div class="table-container">
          <v-data-table
            :headers="tableHeaders"
            :items="workDaysData"
            :search="searchQuery"
            class="work-days-table"
            no-data-text="لا توجد بيانات"
            loading-text="جاري التحميل..."
            :items-per-page="-1"
            hide-default-footer
          >
        <!-- Serial Number Column -->
        <template v-slot:item.serial="{ item }">
          <span class="serial-number">{{ item.serial }}</span>
        </template>

        <!-- Date Column -->
        <template v-slot:item.date="{ item }">
          <span class="date-text">{{ item.date }}</span>
        </template>

        <!-- Day Column -->
        <template v-slot:item.day="{ item }">
          <span class="day-text">{{ item.day }}</span>
        </template>

        <!-- Work Period Column -->
        <template v-slot:item.workPeriod="{ item }">
          <span class="period-text">{{ item.workPeriod }}</span>
        </template>

        <!-- Work Type Column -->
        <template v-slot:item.workType="{ item }">
          <span class="work-type-text">{{ item.workType }}</span>
        </template>

        <!-- About Column -->
        <template v-slot:item.about="{ item }">
          <span class="about-text">{{ item.about }}</span>
        </template>

        <!-- Actions Column -->
        <template v-slot:item.actions="{ item }">
          <div class="actions-buttons">
            <v-btn
              size="small"
              variant="tonal"
              color="error"
              @click="deleteWorkDay(item)"
              class="action-btn-styled"
              title="حذف"
            >
              <v-icon size="18">mdi-delete-outline</v-icon>
            </v-btn>
            <v-btn
              size="small"
              variant="tonal"
              color="grey-darken-1"
              @click="editWorkDay(item)"
              class="action-btn-styled"
              title="تعديل"
            >
              <v-icon size="18">mdi-pencil-outline</v-icon>
            </v-btn>
            <v-btn
              size="small"
              variant="tonal"
              color="primary"
              @click="viewWorkDay(item)"
              class="action-btn-styled"
              title="عرض التفاصيل"
            >
              <v-icon size="18">mdi-eye-outline</v-icon>
            </v-btn>
            <v-btn
              v-if="item.status !== 'completed'"
              size="small"
              variant="tonal"
              color="success"
              @click="markAsComplete(item)"
              class="action-btn-styled complete-btn"
              title="تحديد كمكتمل"
            >
              <v-icon size="18">mdi-check</v-icon>
            </v-btn>
            <v-btn
              v-else
              size="small"
              variant="tonal"
              color="warning"
              @click="markAsUncomplete(item)"
              class="action-btn-styled uncomplete-btn"
              title="إلغاء الإكتمال"
            >
              <v-icon size="18">mdi-close</v-icon>
            </v-btn>
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
        </div>
      </v-card-text>
    </v-card>


    <!-- Confirm Complete Dialog -->
    <v-dialog v-model="showConfirmComplete" max-width="450" persistent>
      <v-card class="confirm-dialog-card-light">
        <v-card-title class="confirm-dialog-header-light">
          <v-icon class="me-2" color="primary">mdi-check-circle-outline</v-icon>
          تأكيد إتمام يوم العمل
        </v-card-title>
        <v-card-text class="confirm-dialog-content-light">
          <p>هل أنت متأكد من تحديد يوم العمل كمكتمل؟</p>
          <p class="text-medium-emphasis mt-2">
            سيتم تحديث نسبة إنجاز المشروع تلقائياً.
          </p>
        </v-card-text>
        <v-card-actions class="confirm-dialog-actions-light">
          <v-spacer />
          <v-btn
            variant="outlined"
            class="cancel-btn-light"
            @click="cancelComplete"
          >
            إلغاء
          </v-btn>
          <v-btn
            variant="elevated"
            class="confirm-btn-light"
            @click="confirmComplete"
          >
            <v-icon class="me-1">mdi-check</v-icon>
            تأكيد
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Confirm Uncomplete Dialog -->
    <v-dialog v-model="showConfirmUncomplete" max-width="450" persistent>
      <v-card class="confirm-dialog-card-light">
        <v-card-title class="confirm-dialog-header-light">
          <v-icon class="me-2" color="primary">mdi-undo-variant</v-icon>
          تأكيد إلغاء الإكتمال
        </v-card-title>
        <v-card-text class="confirm-dialog-content-light">
          <p>هل أنت متأكد من إلغاء إكتمال يوم العمل؟</p>
          <p class="text-medium-emphasis mt-2">
            سيتم تقليل نسبة إنجاز المشروع تلقائياً.
          </p>
        </v-card-text>
        <v-card-actions class="confirm-dialog-actions-light">
          <v-spacer />
          <v-btn
            variant="outlined"
            class="cancel-btn-light"
            @click="cancelUncomplete"
          >
            إلغاء
          </v-btn>
          <v-btn
            variant="elevated"
            class="confirm-btn-light"
            @click="confirmUncomplete"
          >
            <v-icon class="me-1">mdi-undo</v-icon>
            تأكيد
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Success Snackbar -->
    <v-snackbar
      v-model="showSuccessMessage"
      color="success"
      timeout="3000"
      location="top"
    >
      <v-icon class="me-2">mdi-check-circle</v-icon>
      {{ successMessage }}
    </v-snackbar>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { listWorkDays, createWorkDay, updateWorkDay, deleteWorkDay as apiDeleteWorkDay, completeWorkDay, uncompleteWorkDay, listWorkSubCategories } from '@/api/workdays'
import { DEFAULT_LIMIT } from '@/constants/pagination'

const router = useRouter()
const route = useRoute()

// State
const showAddForm = ref(false)
const formValid = ref(false)
const isEditing = ref(false)
const saving = ref(false)
const searchQuery = ref('')
const showSuccessMessage = ref(false)
const successMessage = ref('')
const loading = ref(false)
const error = ref('')
const projectId = computed(() => route.query.projectId || '')

// Table headers
const tableHeaders = [
  { title: 'التسلسل', key: 'serial', sortable: true, width: '80px' },
  { title: 'التاريخ', key: 'date', sortable: true, width: '120px' },
  { title: 'اليوم', key: 'day', sortable: true, width: '100px' },
  { title: 'فترة العمل', key: 'workPeriod', sortable: true, width: '120px' },
  { title: 'نوع العمل', key: 'workType', sortable: true, width: '150px' },
  { title: 'عن', key: 'about', sortable: true, width: '200px' },
  { title: 'الاجراءات', key: 'actions', sortable: false, width: '120px' }
]

// Work days data
const workDaysData = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(DEFAULT_LIMIT)
const totalItems = ref(0)
const totalPages = ref(0)

// Form data
const workDayForm = ref({
  date: '',
  description: '',
  notes: '',
  status: 'draft',
  workSubCategoryId: null
})

// Original form data for change detection
const originalFormData = ref(null)

// Check if form has changes (for edit mode)
const hasFormChanges = computed(() => {
  if (!isEditing.value || !originalFormData.value) return true
  return (
    workDayForm.value.date !== originalFormData.value.date ||
    workDayForm.value.description !== originalFormData.value.description ||
    workDayForm.value.notes !== originalFormData.value.notes ||
    workDayForm.value.status !== originalFormData.value.status ||
    workDayForm.value.workSubCategoryId !== originalFormData.value.workSubCategoryId
  )
})

// Work subcategories
const workSubCategories = ref([])
const loadingSubCategories = ref(false)

// Options (removed 'completed' - use markAsComplete button instead)
const statusOptions = [
  { title: 'مسودة', value: 'draft' },
  { title: 'قيد التنفيذ', value: 'in_progress' }
]

// Date picker state
const dateMenu = ref(false)
const selectedDate = ref(null)

// Computed property for formatted date display
const formattedDate = computed(() => {
  if (!workDayForm.value.date) return ''
  // Parse date string (YYYY-MM-DD) without timezone conversion
  const [year, month, day] = workDayForm.value.date.split('-').map(Number)
  const date = new Date(year, month - 1, day, 12, 0, 0)
  return date.toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })
})

// Handle date selection from picker
const onDateSelected = (date) => {
  if (date) {
    // Convert to YYYY-MM-DD format for the form
    const d = new Date(date)
    workDayForm.value.date = d.toISOString().split('T')[0]
  }
  dateMenu.value = false
}

// دالة تصدير البيانات إلى Excel
const exportToExcel = () => {
  try {
    // إنشاء ملف Excel متعدد الشيتات
    const workbook = createExcelWorkbook()
    
    // إنشاء رابط التحميل
    const link = document.createElement('a')
    const url = URL.createObjectURL(workbook)
    link.setAttribute('href', url)
    link.setAttribute('download', `تقرير_شامل_${new Date().toISOString().split('T')[0]}.csv`)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)

    // إظهار رسالة نجاح
    successMessage.value = 'تم تصدير التقرير الشامل بنجاح'
    showSuccessMessage.value = true

  } catch (error) {
    console.error('خطأ في تصدير البيانات:', error)
    successMessage.value = 'حدث خطأ في تصدير البيانات'
    showSuccessMessage.value = true
  }
}

// دالة إنشاء ملف Excel متعدد الشيتات
const createExcelWorkbook = () => {
  // إنشاء محتوى CSV شامل يحتوي على جميع الشيتات
  const allSheetsCSV = createAllSheetsCSV()
  
  // إضافة BOM للدعم العربي
  const BOM = '\uFEFF'
  const blob = new Blob([BOM + allSheetsCSV], { 
    type: 'text/csv;charset=utf-8;' 
  })
  
  return blob
}

// دالة إنشاء ملف CSV شامل يحتوي على جميع الشيتات
const createAllSheetsCSV = () => {
  const workDaysCSV = createWorkDaysCSV()
  const machineryCSV = createMachineryCSV()
  const dailyExpensesCSV = createDailyExpensesCSV()
  const materialsCSV = createMaterialsCSV()
  const laborCSV = createLaborCSV()
  const summaryCSV = createSummaryCSV()
  
  // دمج جميع الشيتات في ملف واحد مع فواصل واضحة
  let allSheets = ''
  
  // إضافة عنوان رئيسي
  allSheets += 'تقرير شامل - أيام العمل والمصاريف\n'
  allSheets += `تاريخ التصدير: ${new Date().toLocaleDateString('en-US')}\n`
  allSheets += '='.repeat(80) + '\n\n'
  
  // إضافة كل شيت مع عنوان واضح
  allSheets += '📋 شيت 1: أيام العمل\n'
  allSheets += '='.repeat(40) + '\n'
  allSheets += workDaysCSV + '\n\n'
  
  allSheets += '🔧 شيت 2: الآليات والمعدات\n'
  allSheets += '='.repeat(40) + '\n'
  allSheets += machineryCSV + '\n\n'
  
  allSheets += '💰 شيت 3: المصاريف اليومية\n'
  allSheets += '='.repeat(40) + '\n'
  allSheets += dailyExpensesCSV + '\n\n'
  
  allSheets += '🏗️ شيت 4: المواد والمواد الخام\n'
  allSheets += '='.repeat(40) + '\n'
  allSheets += materialsCSV + '\n\n'
  
  allSheets += '👷 شيت 5: الأيدي العاملة\n'
  allSheets += '='.repeat(40) + '\n'
  allSheets += laborCSV + '\n\n'
  
  allSheets += '📈 شيت 6: ملخص التكاليف\n'
  allSheets += '='.repeat(40) + '\n'
  allSheets += summaryCSV + '\n\n'
  
  // إضافة ملاحظات
  allSheets += 'ملاحظات:\n'
  allSheets += '- يمكن فتح هذا الملف في Excel أو أي برنامج جداول بيانات\n'
  allSheets += '- كل شيت منفصل بوضوح لسهولة القراءة\n'
  allSheets += '- جميع التكاليف محسوبة تلقائياً\n'
  allSheets += '- البيانات محدثة حتى تاريخ التصدير\n'
  
  return allSheets
}

// شيت أيام العمل - CSV
const createWorkDaysCSV = () => {
  const workDaysData = getWorkDaysData()
  const csvData = [
    ['تقرير أيام العمل'],
    ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
    [''],
    ['التسلسل', 'التاريخ', 'اليوم', 'فترة العمل', 'نوع العمل', 'الوصف', 'التكلفة']
  ]
  
  workDaysData.forEach(item => {
    csvData.push([
      item.serial,
      item.date,
      item.day,
      item.workPeriod,
      item.workType,
      item.about,
      item.cost
    ])
  })
  
  const totalCost = workDaysData.reduce((sum, item) => sum + item.cost, 0)
  csvData.push(['', '', '', '', '', '', ''])
  csvData.push(['إجمالي أيام العمل', '', '', '', '', '', workDaysData.length])
  csvData.push(['إجمالي التكلفة', '', '', '', '', '', totalCost])
  
  return csvData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

// شيت الآليات - CSV
const createMachineryCSV = () => {
  const machineryData = getMachineryData()
  const csvData = [
    ['تقرير الآليات والمعدات'],
    ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
    [''],
    ['اسم الآلة', 'نوع الآلة', 'ساعات التشغيل', 'التكلفة/ساعة', 'إجمالي التكلفة', 'الحالة', 'ملاحظات']
  ]
  
  machineryData.forEach(item => {
    csvData.push([
      item.name,
      item.type,
      item.hours,
      item.costPerHour,
      item.totalCost,
      item.status,
      item.notes
    ])
  })
  
  const totalCost = machineryData.reduce((sum, item) => sum + item.totalCost, 0)
  csvData.push(['', '', '', '', '', '', ''])
  csvData.push(['إجمالي الآليات', '', '', '', '', '', machineryData.length])
  csvData.push(['إجمالي التكلفة', '', '', '', '', '', totalCost])
  
  return csvData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

// شيت المصاريف اليومية - CSV
const createDailyExpensesCSV = () => {
  const dailyExpenses = getDailyExpenses()
  const csvData = [
    ['تقرير المصاريف اليومية'],
    ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
    [''],
    ['التاريخ', 'نوع المصروف', 'المبلغ', 'الوصف', 'المشروع', 'المسؤول', 'الحالة']
  ]
  
  dailyExpenses.forEach(item => {
    csvData.push([
      item.date,
      item.type,
      item.amount,
      item.description,
      item.project,
      item.responsible,
      item.status
    ])
  })
  
  const totalAmount = dailyExpenses.reduce((sum, item) => sum + item.amount, 0)
  csvData.push(['', '', '', '', '', '', ''])
  csvData.push(['إجمالي المصاريف', '', '', '', '', '', dailyExpenses.length])
  csvData.push(['إجمالي المبلغ', '', '', '', '', '', totalAmount])
  
  return csvData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

// شيت المواد - CSV
const createMaterialsCSV = () => {
  const materialsData = getMaterialsData()
  const csvData = [
    ['تقرير المواد والمواد الخام'],
    ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
    [''],
    ['اسم المادة', 'الكمية', 'الوحدة', 'سعر الوحدة', 'إجمالي التكلفة', 'المورد', 'تاريخ الشراء']
  ]
  
  materialsData.forEach(item => {
    csvData.push([
      item.name,
      item.quantity,
      item.unit,
      item.unitPrice,
      item.totalCost,
      item.supplier,
      item.purchaseDate
    ])
  })
  
  const totalCost = materialsData.reduce((sum, item) => sum + item.totalCost, 0)
  csvData.push(['', '', '', '', '', '', ''])
  csvData.push(['إجمالي المواد', '', '', '', '', '', materialsData.length])
  csvData.push(['إجمالي التكلفة', '', '', '', '', '', totalCost])
  
  return csvData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

// شيت الأيدي العاملة - CSV
const createLaborCSV = () => {
  const laborData = getLaborData()
  const csvData = [
    ['تقرير الأيدي العاملة'],
    ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
    [''],
    ['اسم العامل', 'المهنة', 'ساعات العمل', 'أجر الساعة', 'إجمالي الأجر', 'المشروع', 'التاريخ']
  ]
  
  laborData.forEach(item => {
    csvData.push([
      item.name,
      item.profession,
      item.hours,
      item.hourlyWage,
      item.totalWage,
      item.project,
      item.date
    ])
  })
  
  const totalWage = laborData.reduce((sum, item) => sum + item.totalWage, 0)
  csvData.push(['', '', '', '', '', '', ''])
  csvData.push(['إجمالي العمال', '', '', '', '', '', laborData.length])
  csvData.push(['إجمالي الأجور', '', '', '', '', '', totalWage])
  
  return csvData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

// شيت الملخص - CSV
const createSummaryCSV = () => {
  const totalWorkCost = workDaysData.value.reduce((sum, item) => 
    sum + calculateWorkDayCost(item.workType, item.workPeriod), 0)
  const machineryData = getMachineryData()
  const dailyExpenses = getDailyExpenses()
  const materialsData = getMaterialsData()
  const laborData = getLaborData()
  
  const csvData = [
    ['ملخص التكاليف الشامل'],
    ['تاريخ التصدير', new Date().toLocaleDateString('en-US')],
    [''],
    ['نوع التكلفة', 'المبلغ', 'النسبة المئوية'],
    ['تكلفة العمل', totalWorkCost, ''],
    ['تكلفة الآليات', machineryData.reduce((sum, item) => sum + item.totalCost, 0), ''],
    ['المصاريف اليومية', dailyExpenses.reduce((sum, item) => sum + item.amount, 0), ''],
    ['تكلفة المواد', materialsData.reduce((sum, item) => sum + item.totalCost, 0), ''],
    ['تكلفة الأيدي العاملة', laborData.reduce((sum, item) => sum + item.totalWage, 0), ''],
    ['', '', ''],
    ['الإجمالي العام', totalWorkCost + 
      machineryData.reduce((sum, item) => sum + item.totalCost, 0) +
      dailyExpenses.reduce((sum, item) => sum + item.amount, 0) +
      materialsData.reduce((sum, item) => sum + item.totalCost, 0) +
      laborData.reduce((sum, item) => sum + item.totalWage, 0), '']
  ]
  
  return csvData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
}

// دالة الحصول على بيانات أيام العمل
const getWorkDaysData = () => {
  return workDaysData.value.map(item => ({
    serial: item.serial,
    date: item.date,
    day: item.day,
    workPeriod: item.workPeriod,
    workType: item.workType,
    about: item.about,
    cost: calculateWorkDayCost(item.workType, item.workPeriod)
  }))
}

// دالة حساب تكلفة يوم العمل
const calculateWorkDayCost = (workType, workPeriod) => {
  const hourlyRates = {
    'بناء': 50,
    'كهرباء': 60,
    'سباكة': 55,
    'دهان': 45,
    'نجارة': 65,
    'حدادة': 70,
    'أخرى': 40
  }
  
  const hours = workPeriod.includes('8') ? 8 : workPeriod.includes('6') ? 6 : 4
  const rate = hourlyRates[workType] || 40
  return hours * rate
}

// بيانات الآليات والمعدات
const getMachineryData = () => [
  { name: 'حفار صغير', type: 'حفارة', hours: 24, costPerHour: 200, totalCost: 4800, status: 'نشط', notes: 'حالة جيدة' },
  { name: 'خلاطة خرسانة', type: 'معدات بناء', hours: 16, costPerHour: 150, totalCost: 2400, status: 'نشط', notes: 'تحتاج صيانة' },
  { name: 'رافعة صغيرة', type: 'رافعة', hours: 12, costPerHour: 300, totalCost: 3600, status: 'متوقف', notes: 'عطل في المحرك' },
  { name: 'مولد كهرباء', type: 'مولد', hours: 32, costPerHour: 100, totalCost: 3200, status: 'نشط', notes: 'يعمل بكفاءة' }
]

// بيانات المصاريف اليومية
const getDailyExpenses = () => [
  { date: '2024-01-15', type: 'وقود', amount: 500, description: 'وقود للآليات', project: 'مشروع البناء', responsible: 'أحمد محمد', status: 'مدفوع' },
  { date: '2024-01-16', type: 'طعام', amount: 300, description: 'غداء العمال', project: 'مشروع البناء', responsible: 'سارة أحمد', status: 'مدفوع' },
  { date: '2024-01-17', type: 'نقل', amount: 200, description: 'نقل المواد', project: 'مشروع البناء', responsible: 'محمد علي', status: 'مدفوع' },
  { date: '2024-01-18', type: 'أدوات', amount: 150, description: 'أدوات يدوية', project: 'مشروع البناء', responsible: 'فاطمة حسن', status: 'مدفوع' }
]

// بيانات المواد والمواد الخام
const getMaterialsData = () => [
  { name: 'أسمنت', quantity: 50, unit: 'كيس', unitPrice: 25, totalCost: 1250, supplier: 'شركة البناء', purchaseDate: '2024-01-10' },
  { name: 'حديد تسليح', quantity: 2, unit: 'طن', unitPrice: 3000, totalCost: 6000, supplier: 'مصنع الحديد', purchaseDate: '2024-01-12' },
  { name: 'رمل', quantity: 10, unit: 'متر مكعب', unitPrice: 80, totalCost: 800, supplier: 'محجر الرمل', purchaseDate: '2024-01-14' },
  { name: 'طوب', quantity: 1000, unit: 'قطعة', unitPrice: 2, totalCost: 2000, supplier: 'مصنع الطوب', purchaseDate: '2024-01-16' },
  { name: 'دهان', quantity: 20, unit: 'لتر', unitPrice: 45, totalCost: 900, supplier: 'شركة الدهانات', purchaseDate: '2024-01-18' }
]

// بيانات الأيدي العاملة
const getLaborData = () => [
  { name: 'أحمد محمد', profession: 'بناء', hours: 40, hourlyWage: 50, totalWage: 2000, project: 'مشروع البناء', date: '2024-01-15' },
  { name: 'سارة أحمد', profession: 'كهربائي', hours: 32, hourlyWage: 60, totalWage: 1920, project: 'مشروع البناء', date: '2024-01-15' },
  { name: 'محمد علي', profession: 'سباك', hours: 24, hourlyWage: 55, totalWage: 1320, project: 'مشروع البناء', date: '2024-01-16' },
  { name: 'فاطمة حسن', profession: 'دهان', hours: 16, hourlyWage: 45, totalWage: 720, project: 'مشروع البناء', date: '2024-01-17' },
  { name: 'علي محمود', profession: 'نجار', hours: 28, hourlyWage: 65, totalWage: 1820, project: 'مشروع البناء', date: '2024-01-18' }
]

// Methods
const goBack = () => {
  router.push('/project-management')
}

const addWorkDay = () => {
  if (!projectId.value) {
    successMessage.value = 'يجب اختيار مشروع أولاً'
    showSuccessMessage.value = true
    return
  }
  isEditing.value = false
  resetForm()
  showAddForm.value = true
}

const closeAddForm = () => {
  showAddForm.value = false
  resetForm()
}

const editWorkDay = (item) => {
  isEditing.value = true

  // Parse date properly to avoid timezone issues
  let dateValue = ''
  let dateForPicker = null
  if (item.workDate) {
    const dateStr = item.workDate.split('T')[0]
    dateValue = dateStr
    // Create date at noon to avoid timezone shifting
    const [year, month, day] = dateStr.split('-').map(Number)
    dateForPicker = new Date(year, month - 1, day, 12, 0, 0)
  }

  // Map item fields to form fields properly
  const formData = {
    id: item.id,
    date: dateValue,
    description: item.description || '',
    notes: item.notes || '',
    status: item.status || 'draft',
    workSubCategoryId: item.workSubCategoryId ? Number(item.workSubCategoryId) : null,
  }
  workDayForm.value = { ...formData }
  // Store original data for change detection
  originalFormData.value = { ...formData }
  // Initialize selected date for date picker
  selectedDate.value = dateForPicker
  showAddForm.value = true
}

const viewWorkDay = (item) => {
  router.push({ path: '/work-day-details', query: { id: item.id, projectId: projectId.value } })
}

const testDetailsPage = () => {
  // Navigate to test details page
  router.push('/work-day-details/test-123')
}

const testSimplePage = () => {
  // Navigate to simple test page
  router.push('/work-day-details-simple')
}

const testFixedPage = () => {
  // Navigate to fixed test page
  router.push('/work-day-details-fixed/test-123')
}

const testHtmlPage = () => {
  // Navigate to HTML page in same tab
  window.location.href = '/test-work-details.html'
}

const deleteWorkDay = async (item) => {
  if (!item?.id) return
  await apiDeleteWorkDay(item.id)
  await loadWorkDays()
  successMessage.value = 'تم حذف يوم العمل بنجاح'
  showSuccessMessage.value = true
}

const toggleStatus = (item) => {
  console.log('Toggle status:', item)
}

// Confirm complete dialog
const showConfirmComplete = ref(false)
const itemToComplete = ref(null)

// Show confirmation dialog before completing
const markAsComplete = (item) => {
  if (!item?.id) return
  itemToComplete.value = item
  showConfirmComplete.value = true
}

// Cancel completion
const cancelComplete = () => {
  showConfirmComplete.value = false
  itemToComplete.value = null
}

// Confirm and complete work day
const confirmComplete = async () => {
  if (!itemToComplete.value?.id) return
  try {
    await completeWorkDay(itemToComplete.value.id)
    await loadWorkDays()
    successMessage.value = 'تم تحديد يوم العمل كمكتمل وتحديث نسبة إنجاز المشروع'
    showSuccessMessage.value = true
  } catch (error) {
    console.error('Failed to complete work day:', error)
    successMessage.value = error?.message || 'فشل تحديد يوم العمل كمكتمل'
    showSuccessMessage.value = true
  } finally {
    showConfirmComplete.value = false
    itemToComplete.value = null
  }
}

// Confirm uncomplete dialog
const showConfirmUncomplete = ref(false)
const itemToUncomplete = ref(null)

// Show confirmation dialog before uncompleting
const markAsUncomplete = (item) => {
  if (!item?.id) return
  itemToUncomplete.value = item
  showConfirmUncomplete.value = true
}

// Cancel uncomplete
const cancelUncomplete = () => {
  showConfirmUncomplete.value = false
  itemToUncomplete.value = null
}

// Confirm and uncomplete work day
const confirmUncomplete = async () => {
  if (!itemToUncomplete.value?.id) return
  try {
    await uncompleteWorkDay(itemToUncomplete.value.id)
    await loadWorkDays()
    successMessage.value = 'تم إلغاء إكتمال يوم العمل وتحديث نسبة إنجاز المشروع'
    showSuccessMessage.value = true
  } catch (error) {
    console.error('Failed to uncomplete work day:', error)
    successMessage.value = error?.message || 'فشل إلغاء إكتمال يوم العمل'
    showSuccessMessage.value = true
  } finally {
    showConfirmUncomplete.value = false
    itemToUncomplete.value = null
  }
}

const saveWorkDay = async () => {
  if (!projectId.value) {
    successMessage.value = 'يجب اختيار مشروع أولاً'
    showSuccessMessage.value = true
    return
  }

  saving.value = true
  try {
    const payload = {
      projectId: Number(projectId.value),
      workSubCategoryId: workDayForm.value.workSubCategoryId ? Number(workDayForm.value.workSubCategoryId) : null,
      workDate: workDayForm.value.date ? new Date(workDayForm.value.date).toISOString() : new Date().toISOString(),
      description: workDayForm.value.description || null,
      notes: workDayForm.value.notes || null,
    }

    if (isEditing.value && workDayForm.value.id) {
      const updatePayload = {
        workSubCategoryId: payload.workSubCategoryId,
        workDate: payload.workDate,
        description: payload.description,
        notes: payload.notes,
        status: workDayForm.value.status || null,
      }
      await updateWorkDay(workDayForm.value.id, updatePayload)
      successMessage.value = 'تم تحديث يوم العمل بنجاح'
    } else {
      await createWorkDay(payload)
      successMessage.value = 'تم إضافة يوم العمل بنجاح'
    }

    await loadWorkDays()
    showAddForm.value = false
    showSuccessMessage.value = true
    resetForm()
  } catch (error) {
    console.error('Failed to save work day', error)
    successMessage.value = error?.message || 'فشل حفظ يوم العمل'
    showSuccessMessage.value = true
  } finally {
    saving.value = false
  }
}


const resetForm = () => {
  // تعيين التاريخ الحالي كافتراضي
  const today = new Date().toISOString().split('T')[0]
  workDayForm.value = {
    date: today,
    description: '',
    notes: '',
    status: 'draft',
    workSubCategoryId: null
  }
  originalFormData.value = null
  formValid.value = false
}

const loadWorkDays = async (page = currentPage.value) => {
  loading.value = true
  error.value = ''
  try {
    const response = await listWorkDays({ projectId: projectId.value, page, limit: itemsPerPage.value })
    const normalized = Array.isArray(response.data) ? response.data : (Array.isArray(response) ? response : [])

    // Update pagination state
    totalItems.value = response.total || normalized.length
    totalPages.value = response.totalPages || 1
    currentPage.value = response.page || page

    workDaysData.value = normalized.map((w, idx) => ({
      id: w.id,
      serial: (currentPage.value - 1) * itemsPerPage.value + idx + 1,
      date: formatDate(w.workDate),
      day: formatDayName(w.workDate),
      workPeriod: '-', // backend does not provide; keep placeholder
      workType: w.status || '',
      about: w.description || w.notes || '',
      status: w.status || 'draft', // Keep status for complete button visibility
      // Preserve original API fields for editing
      workDate: w.workDate,
      description: w.description || '',
      notes: w.notes || '',
      workSubCategoryId: w.workSubCategoryId || null,
    }))
  } catch (err) {
    console.error('فشل جلب أيام العمل', err)
    error.value = 'فشل جلب أيام العمل من الخادم'
  } finally {
    loading.value = false
  }
}

// Handle page change
const onPageChange = (page) => {
  currentPage.value = page
  loadWorkDays(page)
}

const formatDate = (value) => {
  if (!value) return ''
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return value
  return d.toLocaleDateString('en-US', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

const formatDayName = (value) => {
  if (!value) return ''
  const d = new Date(value)
  if (Number.isNaN(d.getTime())) return ''
  const days = ['الأحد', 'الاثنين', 'الثلاثاء', 'الأربعاء', 'الخميس', 'الجمعة', 'السبت']
  return days[d.getDay()] || ''
}

// Load work subcategories
const loadWorkSubCategories = async () => {
  loadingSubCategories.value = true
  try {
    const data = await listWorkSubCategories()
    workSubCategories.value = data.map(item => ({
      title: item.name,
      value: item.id
    }))
  } catch (err) {
    console.error('Failed to load work subcategories:', err)
  } finally {
    loadingSubCategories.value = false
  }
}

onMounted(() => {
  loadWorkSubCategories()
  if (projectId.value) {
    loadWorkDays()
  }
})

watch(projectId, (val, prev) => {
  if (val && val !== prev) loadWorkDays()
})
</script>


<style scoped>
.work-days-page {
  background: #f5f5f5;
  min-height: 100vh;
  direction: rtl;
}

/* Form Content Styling */
.clean-form-content {
  padding: 16px !important;
  background: #f8fafc !important;
  border: 2px solid #d1d5db !important;
  border-radius: 8px !important;
  margin: 12px !important;
}

.clean-form-content :deep(.v-text-field),
.clean-form-content :deep(.v-textarea),
.clean-form-content :deep(.v-select) {
  background: #ffffff !important;
  border-radius: 8px !important;
}

.clean-form-content :deep(.v-field) {
  background: #ffffff !important;
}

.clean-form-content :deep(.v-field__outline) {
  color: #cbd5e1 !important;
}

.clean-form-content :deep(.v-field--focused .v-field__outline) {
  color: #3b82f6 !important;
}

.clean-form-content :deep(.v-label) {
  color: #475569 !important;
  font-weight: 500 !important;
}

.clean-form-content :deep(.v-field__input) {
  color: #1e293b !important;
}

.clean-form-row {
  margin-bottom: 12px !important;
}

/* Top Navigation Bar */
/* Page Title Header */
.page-title-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  padding: 0.4rem 0.75rem !important;
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: white !important;
  box-shadow: 
    0 2px 8px rgba(59, 130, 246, 0.25),
    0 1px 4px rgba(37, 99, 235, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border-bottom: 2px solid rgba(255, 255, 255, 0.3) !important;
  position: relative;
  overflow: hidden;
  min-height: auto;
}

/* Ensure all text in header is white, except buttons */
.page-title-header h1,
.page-title-header h2,
.page-title-header h3,
.page-title-header span:not(.v-btn span),
.page-title-header p,
.page-title-header label {
  color: white !important;
}

/* Force white color with deep selector */
.page-title-header :deep(h1),
.page-title-header :deep(h2),
.page-title-header :deep(h3),
.page-title-header :deep(span):not(.v-btn span),
.page-title-header :deep(p),
.page-title-header :deep(label) {
  color: white !important;
  -webkit-text-fill-color: white !important;
}

/* Exclude buttons from white color */
.page-title-header .v-btn,
.page-title-header .nav-btn {
  color: inherit !important;
}

.page-title-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  animation: shimmer 3s infinite;
  pointer-events: none;
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

.header-left-section {
  display: flex;
  align-items: center;
  gap: 0.4rem !important;
}

.header-right-section {
  display: flex;
  align-items: center;
}

.nav-btn {
  background: rgba(255, 255, 255, 0.95) !important;
  color: #2563eb !important;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.12) !important;
  border: 1px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s ease !important;
  width: 28px !important;
  height: 28px !important;
  min-width: 28px !important;
}

.nav-btn :deep(.v-icon) {
  font-size: 16px !important;
}

.nav-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  transform: translateY(-2px) scale(1.05) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.nav-btn :deep(.v-icon) {
  color: #2563eb !important;
}

.add-btn:hover :deep(.v-icon) {
  transform: rotate(90deg) !important;
}

.title-icon {
  font-size: 1rem !important;
  color: #ffffff !important;
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.3)) !important;
  animation: iconPulse 2s ease-in-out infinite;
}

@keyframes iconPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

.page-title-text {
  font-size: 1rem !important;
  font-weight: 600 !important;
  color: #ffffff !important;
  margin: 0 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3), 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  letter-spacing: 0.3px !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.3 !important;
}

/* Force white color with higher specificity */
.page-title-header .page-title-text,
.page-title-header h1.page-title-text,
.header-left-section .page-title-text,
.header-left-section h1.page-title-text,
.page-title-header :deep(.page-title-text),
.page-title-header :deep(h1.page-title-text) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* Override any inline styles or other rules */
.page-title-header .header-left-section .page-title-text {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

/* Force white color for all text elements in header */
/* Force white color for text in header sections */
.page-title-header .header-left-section {
  color: white !important;
}

.page-title-header .header-right-section {
  color: white !important;
}

/* Specific text elements - force white */
.page-title-header .header-left-section h1,
.page-title-header .header-left-section h2,
.page-title-header .header-left-section h3,
.page-title-header .header-left-section p,
.page-title-header .header-left-section label,
.page-title-header .header-left-section div:not(.search-box):not(.search-section) {
  color: white !important;
}

.search-section {
  display: flex;
  align-items: center;
}

.search-box {
  display: flex;
  align-items: center;
  background: rgba(255, 255, 255, 1) !important;
  border-radius: 12px !important;
  padding: 0.5rem 1rem !important;
  gap: 0.75rem !important;
  box-shadow: 
    0 4px 12px rgba(0, 0, 0, 0.15),
    0 2px 6px rgba(0, 0, 0, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.9) !important;
  border: 2px solid rgba(255, 255, 255, 0.8) !important;
  min-width: 300px;
}

.search-icon {
  color: #2563eb !important;
  font-size: 1.2rem !important;
  margin-left: 0.5rem;
}

.search-input {
  min-width: 200px;
  flex: 1;
}

.search-input :deep(.v-field) {
  background: transparent !important;
  border: none !important;
  box-shadow: none !important;
}

.search-input :deep(.v-field__input) {
  padding: 0.5rem 0 !important;
  font-size: 0.95rem !important;
  color: #000000 !important;
  font-weight: 600 !important;
}

.search-input :deep(.v-field__input input) {
  color: #000000 !important;
  font-weight: 600 !important;
}

.search-input :deep(.v-field__input::placeholder) {
  color: #9ca3af !important;
  font-weight: 500 !important;
}

.search-btn {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 50%, #3b82f6 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  padding: 0.5rem 1.25rem !important;
  border-radius: 8px !important;
  font-size: 0.85rem !important;
  text-transform: none !important;
  box-shadow: 
    0 2px 8px rgba(30, 64, 175, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s ease !important;
  min-width: 80px;
}

.search-btn:hover {
  background: linear-gradient(135deg, #1e3a8a 0%, #1e40af 50%, #2563eb 100%) !important;
  box-shadow: 
    0 4px 12px rgba(30, 64, 175, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
  transform: translateY(-1px) scale(1.02) !important;
}

.search-btn :deep(.v-btn__content) {
  color: white !important;
  font-weight: 600 !important;
  font-size: 0.85rem !important;
}


/* Modal Styles (old design ما زال مستخدماً في أماكن أخرى) */
.modal-card {
  border-radius: 12px !important;
  overflow: hidden;
}

.close-btn {
  color: white !important;
  min-width: 32px !important;
  height: 32px !important;
}

/* Dropdown Menu Styling - نفس تنسيقات نموذج المهمة */
.modal-content :deep(.v-select__selection) {
  text-align: right !important;
  direction: rtl !important;
  align-items: center !important;
  min-height: 40px !important;
  font-weight: 500 !important;
}

.modal-content :deep(.v-field__append-inner .v-icon) {
  color: #1e40af !important;
  opacity: 1 !important;
}

.modal-content :deep(.v-overlay__content) {
  border-radius: 10px !important;
  box-shadow:
    0 10px 30px rgba(15, 23, 42, 0.18),
    0 4px 10px rgba(148, 163, 184, 0.25) !important;
  border: 1px solid #e5e7eb !important;
  background: #ffffff !important;
  overflow: hidden !important;
}

.modal-content :deep(.v-overlay__content .v-list) {
  padding: 4px 0 !important;
}

.modal-content :deep(.v-overlay__content .v-list-item) {
  min-height: 34px !important;
  padding-inline: 12px !important;
}

.modal-content :deep(.v-overlay__content .v-list-item-title) {
  font-size: 0.85rem !important;
  color: #0f172a !important;
  text-align: right !important;
}

.modal-content :deep(.v-overlay__content .v-list-item--active),
.modal-content :deep(.v-overlay__content .v-list-item:hover) {
  background: #eff6ff !important;
  color: #1d4ed8 !important;
}

.modal-content :deep(.v-select__menu .v-list-item-title) {
  color: #000000 !important;
  font-weight: 600 !important;
}

.modal-content :deep(.v-select__menu .v-list-item) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  background: white !important;
  border-bottom: 1px solid #e9ecef !important;
  transition: all 0.2s ease !important;
}

.modal-content :deep(.v-select__menu .v-list-item:hover) {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-select__menu .v-list-item--active),
.modal-content :deep(.v-select__menu .v-list-item[aria-selected="true"]) {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-select__menu .v-list-item--active:hover) {
  background: #000000 !important;
  color: white !important;
}

/* Select Field Arrow */
.modal-content :deep(.v-field__append-inner) {
  color: #333 !important;
}

.modal-content :deep(.v-field--focused .v-field__append-inner) {
  color: #007bff !important;
}

/* Global styling for all dropdown items in modal - نفس تنسيقات إضافة موظف */
.modal-content :deep(.v-list-item),
.modal-content :deep(.v-list-item__content),
.modal-content :deep(.v-list-item-title),
.modal-content :deep(.v-list-item__title) {
  color: #000000 !important;
  font-weight: 600 !important;
}

.modal-content :deep(.v-select .v-field__input) {
  color: #000000 !important;
}

.modal-content :deep(.v-select .v-field__input::placeholder) {
  color: #000000 !important;
  opacity: 0.7 !important;
}

/* Force color for all text in dropdowns */
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-title),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-subtitle),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-body) {
  color: #1e3a8a !important;
  font-weight: 600 !important;
}

/* Ultra-specific targeting for dropdown text */
.modal-content :deep(.v-list-item) * {
  color: #1e3a8a !important;
}

.modal-content :deep(.v-list-item__content) * {
  color: #1e3a8a !important;
}

.modal-content :deep(.v-select__menu .v-list-item) * {
  color: #1e3a8a !important;
}

.modal-content :deep(.v-menu__content .v-list-item) * {
  color: #1e3a8a !important;
}

.modal-content :deep(.v-overlay__content .v-list-item) * {
  color: #1e3a8a !important;
}

/* Target all possible text elements */
.modal-content :deep(span),
.modal-content :deep(div),
.modal-content :deep(p),
.modal-content :deep(.v-list-item__content span),
.modal-content :deep(.v-list-item__content div) {
  color: #1e3a8a !important;
}

/* Comprehensive dropdown styling */
.modal-content :deep(.v-list-item__overlay),
.modal-content :deep(.v-list-item__underlay) {
  color: #1e3a8a !important;
}

.modal-content :deep(.v-list-item__prepend),
.modal-content :deep(.v-list-item__append) {
  color: #1e3a8a !important;
}

/* Override any white text specifically */
.modal-content :deep(.v-list-item),
.modal-content :deep(.v-list-item__content),
.modal-content :deep(.v-list-item-title),
.modal-content :deep(.v-list-item__title) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
}

/* Force override for any conflicting styles */
.modal-content :deep(.v-list-item[style*="color"]),
.modal-content :deep(.v-list-item__content[style*="color"]),
.modal-content :deep(.v-list-item-title[style*="color"]) {
  color: #1e3a8a !important;
}

/* Ultimate text color override */
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-title),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-subtitle),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-body),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-media),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-action) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
  fill: #1e3a8a !important;
}

/* Target all text nodes */
.modal-content :deep(.v-list-item .v-list-item__content > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > * > *) {
  color: #1e3a8a !important;
}

/* Specific targeting for select menus */
.modal-content :deep(.v-select__menu .v-list-item .v-list-item__content .v-list-item-title),
.modal-content :deep(.v-select__menu .v-list-item .v-list-item__content .v-list-item-subtitle),
.modal-content :deep(.v-select__menu .v-list-item .v-list-item__content .v-list-item-body) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
}

/* Nuclear approach - override everything */
.modal-content :deep(.v-list-item),
.modal-content :deep(.v-list-item__content),
.modal-content :deep(.v-list-item-title),
.modal-content :deep(.v-list-item__title),
.modal-content :deep(.v-list-item-subtitle),
.modal-content :deep(.v-list-item-body),
.modal-content :deep(.v-list-item-media),
.modal-content :deep(.v-list-item-action) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
  fill: #1e3a8a !important;
  stroke: #1e3a8a !important;
}

/* Force all text elements */
.modal-content :deep(.v-list-item .v-list-item__content),
.modal-content :deep(.v-list-item .v-list-item__content > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > * > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > * > * > *) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
}

/* Override any inline styles */
.modal-content :deep(.v-list-item[style]),
.modal-content :deep(.v-list-item__content[style]),
.modal-content :deep(.v-list-item-title[style]),
.modal-content :deep(.v-list-item__title[style]) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
}

/* Ultimate dropdown override */
.modal-content :deep(.v-select__menu),
.modal-content :deep(.v-menu__content),
.modal-content :deep(.v-overlay__content),
.modal-content :deep(.v-list) {
  --v-list-item-color: #1e3a8a !important;
  --v-list-item-title-color: #1e3a8a !important;
  --v-list-item-subtitle-color: #1e3a8a !important;
}

.modal-content :deep(.v-select__menu .v-list-item),
.modal-content :deep(.v-menu__content .v-list-item),
.modal-content :deep(.v-overlay__content .v-list-item),
.modal-content :deep(.v-list .v-list-item) {
  --v-list-item-color: #1e3a8a !important;
  --v-list-item-title-color: #1e3a8a !important;
  --v-list-item-subtitle-color: #1e3a8a !important;
}

/* CSS Variables override */
.modal-content :deep(.v-list-item) {
  color: var(--v-list-item-color, #1e3a8a) !important;
}

.modal-content :deep(.v-list-item-title) {
  color: var(--v-list-item-title-color, #1e3a8a) !important;
}

.modal-content :deep(.v-list-item-subtitle) {
  color: var(--v-list-item-subtitle-color, #1e3a8a) !important;
}

/* Final attempt - target all possible elements */
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-title),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-subtitle),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-body),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-media),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-action),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-prepend),
.modal-content :deep(.v-list-item .v-list-item__content .v-list-item-append) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
  fill: #1e3a8a !important;
  stroke: #1e3a8a !important;
}

/* Override all possible text elements */
.modal-content :deep(.v-list-item .v-list-item__content > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > * > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > * > * > *),
.modal-content :deep(.v-list-item .v-list-item__content > * > * > * > * > *) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
}

/* Global dropdown text color - outside modal */
:deep(.v-list-item) {
  color: #1e3a8a !important;
}

:deep(.v-list-item__content) {
  color: #1e3a8a !important;
}

:deep(.v-list-item-title) {
  color: #1e3a8a !important;
}

:deep(.v-list-item__title) {
  color: #1e3a8a !important;
}

:deep(.v-list-item-subtitle) {
  color: #1e3a8a !important;
}

:deep(.v-list-item-body) {
  color: #1e3a8a !important;
}

/* Global select menu styling */
:deep(.v-select__menu .v-list-item) {
  color: #1e3a8a !important;
}

:deep(.v-select__menu .v-list-item__content) {
  color: #1e3a8a !important;
}

:deep(.v-select__menu .v-list-item-title) {
  color: #1e3a8a !important;
}

/* Global menu content styling */
:deep(.v-menu__content .v-list-item) {
  color: #1e3a8a !important;
}

:deep(.v-menu__content .v-list-item__content) {
  color: #1e3a8a !important;
}

:deep(.v-menu__content .v-list-item-title) {
  color: #1e3a8a !important;
}

/* Ultimate global override for all dropdown text */
:deep(.v-list-item *),
:deep(.v-list-item__content *),
:deep(.v-select__menu *),
:deep(.v-menu__content *),
:deep(.v-overlay__content *) {
  color: #1e3a8a !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #1e3a8a !important;
}

/* Force all text elements globally */
:deep(span),
:deep(div),
:deep(p),
:deep(.v-list-item span),
:deep(.v-list-item div),
:deep(.v-list-item p) {
  color: #1e3a8a !important;
}

/* Override any CSS variables globally */
:deep(.v-list-item) {
  --v-list-item-color: #1e3a8a !important;
  --v-list-item-title-color: #1e3a8a !important;
  --v-list-item-subtitle-color: #1e3a8a !important;
}

:deep(.v-list-item-title) {
  color: var(--v-list-item-title-color, #1e3a8a) !important;
}

:deep(.v-list-item-subtitle) {
  color: var(--v-list-item-subtitle-color, #1e3a8a) !important;
}

/* تنسيق حقول اليوم ونوع العمل */
.day-select :deep(.v-field),
.work-type-select :deep(.v-field) {
  background: rgba(255, 255, 255, 1) !important;
  border: 2px solid rgba(156, 163, 175, 0.3) !important;
  border-radius: 10px !important;
  transition: all 0.3s ease !important;
}

.day-select :deep(.v-field__input),
.work-type-select :deep(.v-field__input) {
  color: #000000 !important;
  font-weight: 600 !important;
  padding: 12px 16px !important;
}

.day-select :deep(.v-field__input input),
.work-type-select :deep(.v-field__input input) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.day-select :deep(.v-label),
.work-type-select :deep(.v-label) {
  color: #6b7280 !important;
  font-weight: 600 !important;
  background: rgba(255, 255, 255, 1) !important;
  padding: 0 8px !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.day-select :deep(.v-field--focused .v-label),
.work-type-select :deep(.v-field--focused .v-label) {
  color: #2563eb !important;
  background: rgba(255, 255, 255, 1) !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.day-select :deep(.v-field:hover),
.work-type-select :deep(.v-field:hover) {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(59, 130, 246, 0.5) !important;
}

.day-select :deep(.v-field--focused),
.work-type-select :deep(.v-field--focused) {
  border-color: #2563eb !important;
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15) !important;
  background: rgba(255, 255, 255, 1) !important;
}

.day-select :deep(.v-icon),
.work-type-select :deep(.v-icon) {
  color: #2563eb !important;
}

/* تنسيق حقل التاريخ */
.date-field :deep(.v-field) {
  background: rgba(255, 255, 255, 1) !important;
  border: 2px solid rgba(156, 163, 175, 0.3) !important;
  border-radius: 10px !important;
  transition: all 0.3s ease !important;
}

.date-field :deep(.v-field__input) {
  color: #000000 !important;
  font-weight: 600 !important;
  padding: 12px 16px !important;
  text-decoration: none !important;
  border-bottom: none !important;
}

.date-field :deep(.v-field__input input) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.date-field :deep(.v-label) {
  color: #6b7280 !important;
  font-weight: 600 !important;
  background: rgba(255, 255, 255, 1) !important;
  padding: 0 8px !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.date-field :deep(.v-field--focused .v-label) {
  color: #2563eb !important;
  background: rgba(255, 255, 255, 1) !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.date-field :deep(.v-field:hover) {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(59, 130, 246, 0.5) !important;
}

.date-field :deep(.v-field--focused) {
  border-color: #2563eb !important;
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15) !important;
  background: rgba(255, 255, 255, 1) !important;
}

.date-field :deep(.v-icon) {
  color: #2563eb !important;
}

/* Specific styling for day and work type selects - نفس تنسيقات إضافة موظف */
.day-select :deep(.v-menu__content),
.day-select :deep(.v-overlay__content),
.work-type-select :deep(.v-menu__content),
.work-type-select :deep(.v-overlay__content) {
  background: white !important;
  border: 2px solid #2563eb !important;
  border-radius: 10px !important;
  box-shadow: 0 4px 12px rgba(37, 99, 235, 0.2) !important;
}

.day-select :deep(.v-list-item),
.work-type-select :deep(.v-list-item) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  background: white !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #000000 !important;
}

.day-select :deep(.v-list-item__content),
.work-type-select :deep(.v-list-item__content) {
  color: #000000 !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #000000 !important;
}

.day-select :deep(.v-list-item-title),
.work-type-select :deep(.v-list-item-title) {
  color: #000000 !important;
  font-weight: 600 !important;
  text-shadow: none !important;
  -webkit-text-fill-color: #000000 !important;
}

.day-select :deep(.v-list-item:hover),
.work-type-select :deep(.v-list-item:hover) {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.day-select :deep(.v-list-item--active),
.day-select :deep(.v-list-item[aria-selected="true"]),
.work-type-select :deep(.v-list-item--active),
.work-type-select :deep(.v-list-item[aria-selected="true"]) {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

.day-select :deep(.v-list-item--active:hover),
.work-type-select :deep(.v-list-item--active:hover) {
  background: #000000 !important;
  color: white !important;
}

.modal-footer {
  padding: 1rem 1.5rem !important;
  background: #f8f9fa !important;
  border-top: 1px solid #dee2e6 !important;
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
  color: #ffffff !important;
  font-size: 0.75rem !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: 100px !important;
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
  color: #ffffff !important;
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
  color: #ffffff !important;
  box-shadow: 0 2px 8px rgba(156, 163, 175, 0.2) !important;
  transform: none !important;
  opacity: 0.6 !important;
  cursor: not-allowed !important;
}

.save-btn:disabled::before {
  display: none !important;
}

/* Work Period Section */
.work-period-section {
  margin-bottom: 1rem;
  padding: 1.5rem;
  background: rgba(255, 255, 255, 0.95) !important;
  border: 2px solid rgba(59, 130, 246, 0.2) !important;
  border-radius: 12px !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08) !important;
}

.section-label {
  display: block;
  font-weight: 800;
  color: #1e40af !important;
  margin-bottom: 1rem;
  font-size: 1.1rem;
  text-shadow: none !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.period-inputs {
  display: flex;
  align-items: center;
  gap: 1rem;
  width: 100%;
}

.test-btn {
  font-size: 0.8rem !important;
  height: 32px !important;
}

.period-input {
  flex: 1;
  min-width: 0;
}

.period-input :deep(.v-field) {
  background: rgba(255, 255, 255, 1) !important;
  border: 2px solid rgba(156, 163, 175, 0.3) !important;
  border-radius: 10px !important;
  transition: all 0.3s ease !important;
}

.period-input :deep(.v-field__input) {
  color: #000000 !important;
  font-weight: 600 !important;
  padding: 12px 16px !important;
  text-decoration: none !important;
  border-bottom: none !important;
}

.period-input :deep(.v-field__input input) {
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.period-input :deep(.v-field__input input) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
}

.period-input :deep(.v-label) {
  color: #6b7280 !important;
  font-weight: 600 !important;
  background: rgba(255, 255, 255, 1) !important;
  padding: 0 8px !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.period-input :deep(.v-field--focused .v-label) {
  color: #2563eb !important;
  background: rgba(255, 255, 255, 1) !important;
  text-decoration: none !important;
  border-bottom: none !important;
  box-shadow: none !important;
}

.period-input :deep(.v-field__input) {
  text-decoration: none !important;
  border-bottom: none !important;
}

.period-input :deep(.v-field__input::after),
.period-input :deep(.v-field__input::before) {
  display: none !important;
}

.period-input :deep(.v-field:hover) {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(59, 130, 246, 0.5) !important;
}

.period-input :deep(.v-field--focused) {
  border-color: #2563eb !important;
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.15) !important;
  background: rgba(255, 255, 255, 1) !important;
}

.period-input :deep(.v-icon) {
  color: #2563eb !important;
}

.period-separator {
  font-size: 1.5rem;
  font-weight: 800;
  color: #2563eb !important;
  margin: 0 0.5rem;
  text-shadow: none !important;
  flex-shrink: 0;
}

/* Data Table Card */
.data-table-card {
  border-radius: 20px !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1) !important;
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
  margin: 1rem;
  margin-top: 2rem !important;
}

.data-table-card:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15) !important;
}

.data-table-card .v-card-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  padding: 10px 16px !important;
  min-height: auto !important;
  border-radius: 20px 20px 0 0 !important;
}

.data-table-card .v-card-title .v-icon {
  font-size: 20px !important;
  width: 20px !important;
  height: 20px !important;
}

.data-table-card .v-card-title .text-h6 {
  font-size: 1.25rem !important;
  line-height: 1.4 !important;
}

.data-table-card .v-card-title .v-chip {
  font-size: 0.75rem !important;
  height: 24px !important;
  padding: 0 8px !important;
}

.data-table-card .v-card-title .work-days-count-chip {
  background: rgba(255, 255, 255, 0.2) !important;
  color: #ffffff !important;
  border: 1px solid rgba(255, 255, 255, 0.3) !important;
}

.data-table-card .v-card-title .work-days-count-chip :deep(.v-chip__content) {
  color: #ffffff !important;
}

.data-table-card .v-card-text {
  padding: 0 !important;
}

/* Table Search Box */
.table-search-box {
  display: flex;
  align-items: center;
}

.table-search-input :deep(.v-field) {
  background: rgba(255, 255, 255, 0.95) !important;
  border-radius: 8px !important;
  min-height: 36px !important;
  max-height: 36px !important;
}

.table-search-input :deep(.v-field__input) {
  padding: 8px 12px !important;
  font-size: 0.875rem !important;
  min-height: 36px !important;
}

.table-search-input :deep(.v-field__outline) {
  border-color: rgba(255, 255, 255, 0.3) !important;
  border-width: 1px !important;
}

.table-search-input :deep(.v-field--focused .v-field__outline) {
  border-color: rgba(255, 255, 255, 0.6) !important;
  border-width: 2px !important;
}

.table-search-input :deep(.v-field__prepend-inner) {
  padding-top: 0 !important;
  padding-bottom: 0 !important;
  align-items: center !important;
}

.table-search-input :deep(.v-field__prepend-inner .v-icon) {
  font-size: 18px !important;
  color: rgba(255, 255, 255, 0.7) !important;
}

.table-search-input :deep(input) {
  color: #1e293b !important;
  font-size: 0.875rem !important;
}

.table-search-input :deep(input::placeholder) {
  color: #94a3b8 !important;
  font-size: 0.875rem !important;
}

.gap-2 {
  gap: 8px !important;
}

/* Add Button Styles */
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
    0 0 0 1px rgba(255, 255, 255, 0.1) inset !important;
  position: relative !important;
  overflow: hidden !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
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
}

.add-button:active {
  transform: translateY(-1px) scale(1.02) !important;
  box-shadow: 
    0 4px 12px rgba(30, 64, 175, 0.4),
    0 2px 6px rgba(37, 99, 235, 0.3),
    0 0 25px rgba(59, 130, 246, 0.5),
    inset 0 1px 0 rgba(255, 255, 255, 0.3) !important;
}

.add-button:hover :deep(.v-icon) {
  transform: rotate(90deg) scale(1.1) !important;
}

.icon-glow {
  position: relative !important;
}

.icon-glow::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 100%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.3) 0%, transparent 70%);
  border-radius: 50%;
  animation: glow 2s ease-in-out infinite alternate;
  pointer-events: none;
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

@keyframes glow {
  0% { opacity: 0.5; transform: translate(-50%, -50%) scale(0.8); }
  100% { opacity: 1; transform: translate(-50%, -50%) scale(1.2); }
}

.smooth-transition {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

/* Table Container */
.table-container {
  background: white;
  margin: 0;
  border-radius: 0;
  overflow: hidden;
  box-shadow: none;
  border: none;
}

.table-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #007bff, #0056b3);
}

.work-days-table {
  direction: rtl;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 250, 252, 0.95) 100%) !important;
  border-radius: 20px !important;
  overflow: hidden !important;
  box-shadow: 0 12px 40px rgba(67, 56, 202, 0.15) !important;
  backdrop-filter: blur(15px) !important;
  border: 2px solid #cbd5e1 !important;
  margin-top: 20px !important;
}

.work-days-table :deep(table) {
  border-collapse: separate !important;
  border-spacing: 0 !important;
  border: 2px solid #e2e8f0 !important;
  border-radius: 8px !important;
  overflow: hidden !important;
}

.work-days-table :deep(.v-data-table__wrapper) {
  border: 1px solid #e2e8f0 !important;
}

.work-days-table :deep(.v-data-table-header th),
.work-days-table :deep(.v-data-table__th),
.work-days-table :deep(.v-data-table-header) {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 50%, #8b5cf6 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 500 !important;
  font-size: 0.55rem !important;
  border-bottom: 1px solid #ffffff !important;
  text-align: center !important;
  padding: 3px 4px !important;
  letter-spacing: 0px !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  position: relative !important;
  border-radius: 0 !important;
  min-height: 24px !important;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
}

/* Force white color for all text elements in table header */
.work-days-table :deep(.v-data-table-header th *),
.work-days-table :deep(.v-data-table-header th span),
.work-days-table :deep(.v-data-table-header th div),
.work-days-table :deep(.v-data-table-header th .v-data-table-header__content),
.work-days-table :deep(.v-data-table-header th .v-data-table-header__content *) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.work-days-table :deep(.v-data-table-header th::after) {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #ffffff 0%, rgba(255, 255, 255, 0.3) 50%, #ffffff 100%);
  opacity: 0.3;
}

/* إظهار أسهم الفرز دائماً في جميع الأعمدة القابلة للفرز */
.work-days-table :deep(.v-data-table-header th) {
  position: relative;
}

/* إظهار الأسهم دائماً - جميع الحالات */
.work-days-table :deep(.v-data-table-header th .v-data-table-header__content) {
  position: relative;
}

.work-days-table :deep(.v-data-table-header th .v-data-table-header__sort-badge),
.work-days-table :deep(.v-data-table-header th .v-data-table-header__sort-icon),
.work-days-table :deep(.v-data-table-header th .v-icon),
.work-days-table :deep(.v-data-table-header th[aria-sort] .v-data-table-header__sort-badge),
.work-days-table :deep(.v-data-table-header th[aria-sort] .v-data-table-header__sort-icon),
.work-days-table :deep(.v-data-table-header th:not([aria-sort]) .v-data-table-header__sort-badge),
.work-days-table :deep(.v-data-table-header th:not([aria-sort]) .v-data-table-header__sort-icon),
.work-days-table :deep(.v-data-table-header th:hover .v-data-table-header__sort-badge),
.work-days-table :deep(.v-data-table-header th:hover .v-data-table-header__sort-icon) {
  opacity: 1 !important;
  visibility: visible !important;
  display: inline-flex !important;
}

/* إظهار الأسهم حتى في الحالة الافتراضية (غير مفروز) */
.work-days-table :deep(.v-data-table-header th.sortable .v-data-table-header__sort-badge),
.work-days-table :deep(.v-data-table-header th[data-sortable="true"] .v-data-table-header__sort-badge),
.work-days-table :deep(.v-data-table-header th.sortable .v-data-table-header__sort-icon),
.work-days-table :deep(.v-data-table-header th[data-sortable="true"] .v-data-table-header__sort-icon) {
  opacity: 1 !important;
  visibility: visible !important;
  display: inline-flex !important;
}

/* إظهار أيقونة السهم في الأعمدة القابلة للفرز - دائماً */
.work-days-table :deep(.v-data-table-header th.sortable::after),
.work-days-table :deep(.v-data-table-header th[data-sortable="true"]::after) {
  content: '⇅' !important;
  display: inline-block !important;
  margin-right: 4px !important;
  opacity: 1 !important;
  font-size: 0.9rem !important;
  color: rgba(255, 255, 255, 1) !important;
  visibility: visible !important;
}

/* إظهار الأسهم في جميع حالات الفرز */
.work-days-table :deep(.v-data-table-header th[aria-sort="ascending"]::after) {
  content: '↑' !important;
  opacity: 1 !important;
  visibility: visible !important;
}

.work-days-table :deep(.v-data-table-header th[aria-sort="descending"]::after) {
  content: '↓' !important;
  opacity: 1 !important;
  visibility: visible !important;
}

.work-days-table :deep(.v-data-table__wrapper table thead tr th) {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 50%, #6d28d9 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 600 !important;
  font-size: 0.65rem !important;
  text-align: center !important;
  padding: 4px 6px !important;
  border-bottom: 1px solid #5b21b6 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4), 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  letter-spacing: 0.2px !important;
  position: relative !important;
  box-shadow: 0 1px 4px rgba(139, 92, 246, 0.3) !important;
}

/* Force white color for all nested text elements */
.work-days-table :deep(.v-data-table__wrapper table thead tr th *),
.work-days-table :deep(.v-data-table__wrapper table thead tr th span),
.work-days-table :deep(.v-data-table__wrapper table thead tr th div),
.work-days-table :deep(.v-data-table__wrapper table thead tr th .v-data-table-header__content),
.work-days-table :deep(.v-data-table__wrapper table thead tr th .v-data-table-header__content *) {
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
}

.work-days-table :deep(.v-data-table__tbody td) {
  text-align: center !important;
  padding: 6px 6px !important;
  min-height: 40px !important;
  border-bottom: 1px solid #e2e8f0 !important;
  border-right: 1px solid #e2e8f0 !important;
  background: rgba(255, 255, 255, 0.9) !important;
  font-size: 0.7rem !important;
  color: #1e293b !important;
  font-weight: 400 !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.4 !important;
  word-wrap: break-word !important;
  overflow-wrap: break-word !important;
}

.work-days-table :deep(.v-data-table__tbody tr:nth-child(even)) {
  background: linear-gradient(135deg, rgba(238, 242, 255, 0.8) 0%, rgba(224, 231, 255, 0.6) 100%) !important;
}

.work-days-table :deep(.v-data-table__tbody tr:nth-child(odd)) {
  background: rgba(255, 255, 255, 0.95) !important;
}

.work-days-table :deep(.v-data-table__tbody tr) {
  border-bottom: 1px solid #e2e8f0 !important;
}

.work-days-table :deep(.v-data-table__tbody tr:last-child) {
  border-bottom: none !important;
}

.work-days-table :deep(.v-data-table__tbody tr:hover) {
  background: linear-gradient(135deg, rgba(67, 56, 202, 0.08) 0%, rgba(99, 102, 241, 0.05) 100%) !important;
  transform: translateY(-2px) !important;
  box-shadow: 0 8px 20px rgba(67, 56, 202, 0.2) !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.work-days-table :deep(.v-data-table__tbody tr:hover td) {
  color: #4338ca !important;
  font-weight: 600 !important;
}

/* Table Cell Styles */
.serial-number,
.date-text,
.day-text,
.period-text,
.work-type-text,
.about-text {
  font-size: 0.7rem !important;
  color: #212529 !important;
  font-weight: 400 !important;
  text-shadow: 0 1px 1px rgba(0, 0, 0, 0.05) !important;
}

.serial-number {
  background: #f8f9fa !important;
  padding: 0.25rem 0.4rem !important;
  border-radius: 4px !important;
  font-weight: 500 !important;
  font-size: 0.7rem !important;
  color: #007bff !important;
  border: 1px solid #e9ecef !important;
}

.date-text {
  color: #28a745 !important;
  font-weight: 500 !important;
  font-size: 0.7rem !important;
}

.day-text {
  color: #dc3545 !important;
  font-weight: 500 !important;
  font-size: 0.7rem !important;
}

.period-text {
  color: #ffc107 !important;
  font-weight: 500 !important;
  font-size: 0.7rem !important;
  background: #fff3cd !important;
  padding: 0.15rem 0.3rem !important;
  border-radius: 4px !important;
}

.work-type-text {
  color: #6f42c1 !important;
  font-weight: 500 !important;
  font-size: 0.7rem !important;
}

.about-text {
  color: #17a2b8 !important;
  font-weight: 400 !important;
  font-size: 0.7rem !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.actions-buttons {
  display: flex;
  gap: 4px;
  justify-content: center;
  align-items: center;
  padding: 4px 6px;
  background: #f8fafc;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
}

.action-btn-styled {
  min-width: 32px !important;
  width: 32px !important;
  height: 32px !important;
  border-radius: 8px !important;
  padding: 0 !important;
  transition: all 0.2s ease !important;
}

.action-btn-styled:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.12) !important;
}

.action-btn-styled .v-icon {
  font-size: 18px !important;
}

/* Complete button specific styling */
.complete-btn {
  background: rgba(34, 197, 94, 0.15) !important;
}

.complete-btn:hover {
  background: rgba(34, 197, 94, 0.25) !important;
}

.complete-btn .v-icon,
.uncomplete-btn .v-icon {
  transform: scaleX(1) !important;
  direction: ltr !important;
}

/* Uncomplete button specific styling */
.uncomplete-btn {
  background: rgba(245, 158, 11, 0.15) !important;
}

.uncomplete-btn:hover {
  background: rgba(245, 158, 11, 0.25) !important;
}

/* Dialog Styles */
.dialog-header {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  padding: 1.5rem !important;
}

.dialog-content {
  padding: 2rem !important;
}

.dialog-actions {
  padding: 1rem 1.5rem !important;
  background: #f8f9fa !important;
}

/* Pagination */
.work-days-table :deep(.v-data-table-footer) {
  background: #f8f9fa !important;
  border-top: 1px solid #dee2e6 !important;
  padding: 1rem !important;
}

/* Responsive */
@media (max-width: 768px) {
  .main-header {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }
  
  .search-box {
    width: 100%;
  }
  
  .search-input {
    min-width: auto;
    flex: 1;
  }
  
  .header-title {
    align-self: center;
  }
  
  .table-container {
    margin: 0.5rem;
  }
  
  .work-days-table :deep(.v-data-table-header th),
  .work-days-table :deep(.v-data-table__tbody td) {
    padding: 0.5rem 0.25rem !important;
    font-size: 0.8rem !important;
  }
  
  .actions-buttons {
    flex-direction: column;
    gap: 0.1rem;
  }
  
  .action-btn {
    min-width: 24px !important;
    height: 24px !important;
  }
}

/* زر تصدير Excel في عمود الإجراءات */
.export-action-btn {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  border-radius: 6px !important;
  box-shadow: 0 2px 6px rgba(16, 185, 129, 0.3) !important;
  transition: all 0.3s ease !important;
  margin-bottom: 4px !important;
}

.export-action-btn:hover {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  transform: scale(1.1) !important;
  box-shadow: 0 4px 10px rgba(16, 185, 129, 0.4) !important;
}

.export-action-btn .v-icon {
  color: white !important;
  font-size: 1rem !important;
}

/* Confirm Dialog Styles - Light Theme */
.confirm-dialog-card-light {
  border-radius: 16px !important;
  overflow: hidden;
  background: #ffffff !important;
  border: 2px solid #e2e8f0 !important;
  box-shadow: 0 10px 40px rgba(59, 130, 246, 0.15) !important;
}

.confirm-dialog-header-light {
  background: linear-gradient(135deg, #f0f7ff 0%, #e0edff 100%) !important;
  color: #1e40af !important;
  font-weight: 700 !important;
  padding: 1.25rem 1.5rem !important;
  border-bottom: 2px solid #bfdbfe !important;
  font-size: 1.1rem !important;
}

.confirm-dialog-content-light {
  padding: 1.5rem !important;
  font-size: 1rem !important;
  color: #374151 !important;
  background: #ffffff !important;
}

.confirm-dialog-content-light p {
  color: #1e293b !important;
  font-weight: 500 !important;
}

.confirm-dialog-actions-light {
  padding: 1rem 1.5rem !important;
  background: #f8fafc !important;
  border-top: 1px solid #e2e8f0 !important;
}

.cancel-btn-light {
  color: #64748b !important;
  border-color: #cbd5e1 !important;
  font-weight: 600 !important;
  border-radius: 10px !important;
  padding: 8px 20px !important;
}

.cancel-btn-light:hover {
  background: #f1f5f9 !important;
  border-color: #94a3b8 !important;
}

.confirm-btn-light {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  border-radius: 10px !important;
  padding: 8px 20px !important;
  box-shadow: 0 4px 12px rgba(37, 99, 235, 0.3) !important;
}

.confirm-btn-light:hover {
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%) !important;
  box-shadow: 0 6px 16px rgba(37, 99, 235, 0.4) !important;
  transform: translateY(-1px) !important;
}
</style>