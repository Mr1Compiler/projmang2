<template>
  <div class="work-days-page">
    <!-- Page Title Header -->
    <div class="page-title-header">
      <div class="header-left-section">
        <v-btn icon="mdi-arrow-left" @click="goBack" class="nav-btn">
          <v-icon>mdi-arrow-left</v-icon>
        </v-btn>
        <v-btn icon="mdi-plus" @click="addWorkDay" class="nav-btn add-btn">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
        <v-icon class="title-icon">mdi-calendar-clock</v-icon>
        <h1 class="page-title-text" style="color: #ffffff !important;">ايام العمل</h1>
      </div>
      <div class="header-right-section">
        <div class="search-section">
          <div class="search-box">
            <v-icon class="search-icon">mdi-magnify</v-icon>
            <v-text-field
              v-model="searchQuery"
              placeholder="البحث..."
              variant="plain"
              hide-details
              class="search-input"
            />
            <v-btn color="primary" variant="elevated" class="search-btn">
              بحث
            </v-btn>
          </div>
        </div>
      </div>
    </div>

    <!-- Add Work Day Modal -->
    <v-dialog v-model="showAddForm" max-width="600" persistent>
      <v-card class="modal-card">
        <v-card-title class="modal-header">
          <span class="modal-title">اضافة يوم عمل</span>
          <v-btn icon="mdi-close" @click="closeAddForm" variant="text" class="close-btn">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-title>

        <v-card-text class="modal-content">
          <v-form ref="form" v-model="formValid">
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="workDayForm.workLocation"
                  label="مكان العمل"
                  variant="outlined"
                  :rules="[v => !!v || 'مكان العمل مطلوب']"
                  required
                />
              </v-col>
              
              <v-col cols="12">
                <v-text-field
                  v-model="workDayForm.formNumber"
                  label="رقم الاستمارة"
                  variant="outlined"
                  :rules="[v => !!v || 'رقم الاستمارة مطلوب']"
                  required
                />
              </v-col>
              
              <v-col cols="12">
                <div class="work-period-section">
                  <label class="section-label">فتره العمل</label>
                  <div class="period-inputs">
                    <v-text-field
                      v-model="workDayForm.workPeriodFrom"
                      label="من"
                      variant="outlined"
                      type="time"
                      class="period-input"
                    />
                    <span class="period-separator">-</span>
                    <v-text-field
                      v-model="workDayForm.workPeriodTo"
                      label="الى"
                      variant="outlined"
                      type="time"
                      class="period-input"
                    />
                  </div>
                </div>
              </v-col>
              
              <v-col cols="12">
                <v-select
                  v-model="workDayForm.day"
                  label="اليوم"
                  :items="dayOptions"
                  variant="outlined"
                  :rules="[v => !!v || 'اليوم مطلوب']"
                  required
                  class="day-select"
                />
              </v-col>
              
              <v-col cols="12">
                <v-select
                  v-model="workDayForm.workType"
                  label="نوع العمل"
                  :items="workTypeOptions"
                  variant="outlined"
                  :rules="[v => !!v || 'نوع العمل مطلوب']"
                  required
                  class="work-type-select"
                />
              </v-col>
              
              <v-col cols="12">
                <v-text-field
                  v-model="workDayForm.about"
                  label="عن"
                  variant="outlined"
                  placeholder="وصف العمل"
                  :rules="[v => !!v || 'الوصف مطلوب']"
                  required
                />
              </v-col>
              
              <v-col cols="12">
                <v-text-field
                  v-model="workDayForm.date"
                  label="التاريخ"
                  variant="outlined"
                  type="date"
                  :rules="[v => !!v || 'التاريخ مطلوب']"
                  required
                  class="date-field"
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="modal-footer">
          <v-spacer />
          <v-btn
            variant="outlined"
            @click="closeAddForm"
            class="cancel-btn"
            size="large"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="primary"
            variant="elevated"
            @click="saveWorkDay"
            :disabled="!formValid"
            :loading="saving"
            class="save-btn"
            size="large"
          >
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Data Table -->
    <div class="table-container">
      <v-data-table
        :headers="tableHeaders"
        :items="workDaysData"
        :search="searchQuery"
        class="work-days-table"
        no-data-text="لا توجد بيانات"
        loading-text="جاري التحميل..."
        items-per-page="10"
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
              icon="mdi-file-excel"
              size="small"
              color="success"
              variant="text"
              @click="exportToExcel"
              class="action-btn export-action-btn"
              title="تحميل Excel"
            />
            <v-btn
              icon="mdi-delete"
              size="small"
              color="red"
              variant="text"
              @click="deleteWorkDay(item)"
              class="action-btn"
            />
            <v-btn
              icon="mdi-pencil"
              size="small"
              color="black"
              variant="text"
              @click="editWorkDay(item)"
              class="action-btn"
            />
            <v-btn
              icon="mdi-eye"
              size="small"
              color="blue"
              variant="text"
              @click="viewWorkDay(item)"
              class="action-btn"
              title="عرض التفاصيل"
            />
            <v-btn
              icon="mdi-circle"
              size="small"
              color="black"
              variant="text"
              @click="toggleStatus(item)"
              class="action-btn"
            />
          </div>
        </template>
      </v-data-table>
    </div>


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
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// State
const showAddForm = ref(false)
const formValid = ref(false)
const isEditing = ref(false)
const saving = ref(false)
const searchQuery = ref('')
const showSuccessMessage = ref(false)
const successMessage = ref('')

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

// Work days data (matching the image)
const workDaysData = ref([
  {
    id: 1,
    serial: 1,
    date: '10/03/2022',
    day: 'الثلاثاء',
    workPeriod: '8-2',
    workType: 'الاعمال الترابية',
    about: 'اعمال المسوحات الهندسية'
  },
  {
    id: 2,
    serial: 2,
    date: '30/04/2022',
    day: 'الخميس',
    workPeriod: '2-10',
    workType: 'اعمال الباور كيربر',
    about: 'اعمال صب الباور كيربر'
  },
  {
    id: 3,
    serial: 3,
    date: '15/05/2022',
    day: 'السبت',
    workPeriod: '6-12',
    workType: 'اعمال الحفر',
    about: 'حفر اساسات المبنى الرئيسي'
  },
  {
    id: 4,
    serial: 4,
    date: '22/05/2022',
    day: 'الأحد',
    workPeriod: '8-4',
    workType: 'اعمال الخرسانة',
    about: 'صب خرسانة الاساسات'
  }
])

// Form data
const workDayForm = ref({
  workLocation: '',
  formNumber: '',
  workPeriodFrom: '',
  workPeriodTo: '',
  day: '',
  workType: '',
  about: '',
  date: ''
})

// Options
const dayOptions = [
  'السبت', 'الأحد', 'الاثنين', 'الثلاثاء', 'الأربعاء', 'الخميس', 'الجمعة'
]

const workTypeOptions = [
  'الاعمال الترابية',
  'اعمال الباور كيربر',
  'اعمال الحفر',
  'اعمال الخرسانة',
  'اعمال البناء',
  'اعمال الكهرباء',
  'اعمال السباكة',
  'اعمال التشطيبات'
]

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
  allSheets += `تاريخ التصدير: ${new Date().toLocaleDateString('ar-SA')}\n`
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
    ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
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
    ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
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
    ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
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
    ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
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
    ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
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
    ['تاريخ التصدير', new Date().toLocaleDateString('ar-SA')],
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
  workDayForm.value = { ...item }
  showAddForm.value = true
}

const viewWorkDay = (item) => {
  // Navigate to work day details page using Vue Router
  router.push('/work-day-details')
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

const deleteWorkDay = (item) => {
  const index = workDaysData.value.findIndex(day => day.id === item.id)
  if (index > -1) {
    workDaysData.value.splice(index, 1)
    // Update serial numbers
    workDaysData.value.forEach((day, idx) => {
      day.serial = idx + 1
    })
    successMessage.value = 'تم حذف يوم العمل بنجاح'
    showSuccessMessage.value = true
  }
}

const toggleStatus = (item) => {
  console.log('Toggle status:', item)
}

const saveWorkDay = async () => {
  saving.value = true
  
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 1000))
  
  if (isEditing.value) {
    // Update existing work day
    const index = workDaysData.value.findIndex(day => day.id === workDayForm.value.id)
    if (index > -1) {
      workDaysData.value[index] = { ...workDayForm.value }
    }
    successMessage.value = 'تم تحديث يوم العمل بنجاح'
  } else {
    // Add new work day
    const newWorkDay = {
      ...workDayForm.value,
      id: Date.now(),
      serial: workDaysData.value.length + 1,
      workPeriod: `${workDayForm.value.workPeriodFrom}-${workDayForm.value.workPeriodTo}`
    }
    workDaysData.value.push(newWorkDay)
    successMessage.value = 'تم إضافة يوم العمل بنجاح'
  }
  
  saving.value = false
  showAddForm.value = false
  showSuccessMessage.value = true
  resetForm()
}


const resetForm = () => {
  workDayForm.value = {
    workLocation: '',
    formNumber: '',
    workPeriodFrom: '',
    workPeriodTo: '',
    day: '',
    workType: '',
    about: '',
    date: ''
  }
  formValid.value = false
}

// Load data on mount
onMounted(() => {
  const savedData = localStorage.getItem('workDaysData')
  if (savedData) {
    workDaysData.value = JSON.parse(savedData)
  }
})

// Save data when changed
const saveToLocalStorage = () => {
  localStorage.setItem('workDaysData', JSON.stringify(workDaysData.value))
}

// Watch for changes and save
import { watch } from 'vue'
watch(workDaysData, saveToLocalStorage, { deep: true })
</script>


<style scoped>
.work-days-page {
  background: #f5f5f5;
  min-height: 100vh;
  direction: rtl;
}

/* Top Navigation Bar */
/* Page Title Header */
.page-title-header {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  padding: 0.6rem 1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: white !important;
  box-shadow: 
    0 4px 16px rgba(59, 130, 246, 0.3),
    0 2px 8px rgba(37, 99, 235, 0.2),
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
  gap: 0.6rem;
}

.header-right-section {
  display: flex;
  align-items: center;
}

.nav-btn {
  background: rgba(255, 255, 255, 0.95) !important;
  color: #2563eb !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15) !important;
  border: 2px solid rgba(255, 255, 255, 0.3) !important;
  transition: all 0.3s ease !important;
  width: 32px !important;
  height: 32px !important;
  min-width: 32px !important;
}

.nav-btn :deep(.v-icon) {
  font-size: 18px !important;
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
  font-size: 1.4rem !important;
  color: #ffffff !important;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.3)) !important;
  animation: iconPulse 2s ease-in-out infinite;
}

@keyframes iconPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

.page-title-text {
  font-size: 1.2rem;
  font-weight: 700;
  color: #ffffff !important;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4), 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  letter-spacing: 0.5px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.3;
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


/* Modal Styles */
.modal-card {
  border-radius: 12px !important;
  overflow: hidden;
}

.modal-header {
  background: #007bff !important;
  color: white !important;
  padding: 1rem 1.5rem !important;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-title {
  font-size: 1.4rem;
  font-weight: 800;
  color: white;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  letter-spacing: 0.5px;
}

.close-btn {
  color: white !important;
  min-width: 32px !important;
  height: 32px !important;
}

.modal-content {
  padding: 2rem !important;
  background: white;
}

.modal-content :deep(.v-field) {
  border-radius: 8px !important;
  background: #fafafa !important;
  transition: all 0.3s ease !important;
}

.modal-content :deep(.v-field:hover) {
  background: #f5f5f5 !important;
  border-color: #666 !important;
}

.modal-content :deep(.v-field__input) {
  color: #1e3a8a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  text-shadow: none !important;
}

.modal-content :deep(.v-label) {
  color: #000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  text-shadow: none !important;
}

.modal-content :deep(.v-field__outline) {
  border-color: #333 !important;
  border-width: 2px !important;
  border-style: solid !important;
  box-shadow: 0 0 0 1px rgba(51, 51, 51, 0.2) !important;
}

.modal-content :deep(.v-field--focused .v-field__outline) {
  border-color: #007bff !important;
  border-width: 3px !important;
  border-style: solid !important;
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.3) !important;
  outline: none !important;
}

.modal-content :deep(.v-field__input::placeholder) {
  color: #1e3a8a !important;
  opacity: 0.7 !important;
  font-weight: 500 !important;
}

/* Dropdown Menu Styling - نفس تنسيقات إضافة موظف */
.modal-content :deep(.v-list),
.modal-content :deep(.v-menu__content),
.modal-content :deep(.v-overlay__content) {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
  max-height: 300px !important;
  overflow-y: auto !important;
}

.modal-content :deep(.v-list .v-list-item__content),
.modal-content :deep(.v-menu__content .v-list-item__content),
.modal-content :deep(.v-overlay__content .v-list-item__content) {
  color: #000000 !important;
}

.modal-content :deep(.v-list .v-list-item-title),
.modal-content :deep(.v-menu__content .v-list-item-title),
.modal-content :deep(.v-overlay__content .v-list-item-title) {
  color: #000000 !important;
  font-weight: 600 !important;
}

.modal-content :deep(.v-list-item),
.modal-content :deep(.v-menu__content .v-list-item),
.modal-content :deep(.v-overlay__content .v-list-item) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  background: white !important;
  border-bottom: 1px solid #e9ecef !important;
  transition: all 0.2s ease !important;
}

.modal-content :deep(.v-list-item:hover),
.modal-content :deep(.v-menu__content .v-list-item:hover),
.modal-content :deep(.v-overlay__content .v-list-item:hover) {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-list-item--active),
.modal-content :deep(.v-menu__content .v-list-item--active),
.modal-content :deep(.v-overlay__content .v-list-item--active),
.modal-content :deep(.v-list-item[aria-selected="true"]),
.modal-content :deep(.v-menu__content .v-list-item[aria-selected="true"]),
.modal-content :deep(.v-overlay__content .v-list-item[aria-selected="true"]) {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

.modal-content :deep(.v-list-item--active:hover),
.modal-content :deep(.v-menu__content .v-list-item--active:hover),
.modal-content :deep(.v-overlay__content .v-list-item--active:hover) {
  background: #000000 !important;
  color: white !important;
}

/* Vuetify Select Dropdown Styling - نفس تنسيقات إضافة موظف */
.modal-content :deep(.v-select__menu) {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
  max-height: 300px !important;
  overflow-y: auto !important;
}

.modal-content :deep(.v-select__menu .v-list-item__content) {
  color: #000000 !important;
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

/* Table Container */
.table-container {
  background: white;
  margin: 1rem;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border: 1px solid #e9ecef;
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
}

.work-days-table :deep(.v-data-table-header th),
.work-days-table :deep(.v-data-table__th),
.work-days-table :deep(.v-data-table-header) {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 50%, #6d28d9 100%) !important;
  color: #ffffff !important;
  -webkit-text-fill-color: #ffffff !important;
  font-weight: 700 !important;
  font-size: 0.85rem !important;
  border-bottom: 3px solid #5b21b6 !important;
  text-align: center !important;
  padding: 12px 10px !important;
  letter-spacing: 0.5px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4), 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  position: relative !important;
  border-radius: 0 !important;
  box-shadow: 0 2px 8px rgba(139, 92, 246, 0.3) !important;
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
  font-weight: 700 !important;
  font-size: 0.85rem !important;
  text-align: center !important;
  padding: 12px 10px !important;
  border-bottom: 3px solid #5b21b6 !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.4), 0 1px 2px rgba(0, 0, 0, 0.2) !important;
  letter-spacing: 0.5px !important;
  position: relative !important;
  box-shadow: 0 2px 8px rgba(139, 92, 246, 0.3) !important;
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
  padding: 0.75rem 0.5rem !important;
  border-bottom: 2px solid #e9ecef !important;
  background: white !important;
  font-size: 0.75rem !important;
  color: #212529 !important;
  font-weight: 400 !important;
}

.work-days-table :deep(.v-data-table__tbody tr:hover) {
  background: #e3f2fd !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 2px 8px rgba(0, 123, 255, 0.15) !important;
  transition: all 0.2s ease !important;
}

.work-days-table :deep(.v-data-table__tbody tr:hover td) {
  color: #1976d2 !important;
  font-weight: 600 !important;
}

/* Table Cell Styles */
.serial-number,
.date-text,
.day-text,
.period-text,
.work-type-text,
.about-text {
  font-size: 0.75rem !important;
  color: #212529 !important;
  font-weight: 400 !important;
  text-shadow: 0 1px 1px rgba(0, 0, 0, 0.05) !important;
}

.serial-number {
  background: #f8f9fa !important;
  padding: 0.4rem 0.5rem !important;
  border-radius: 6px !important;
  font-weight: 500 !important;
  font-size: 0.75rem !important;
  color: #007bff !important;
  border: 1px solid #e9ecef !important;
}

.date-text {
  color: #28a745 !important;
  font-weight: 500 !important;
  font-size: 0.75rem !important;
}

.day-text {
  color: #dc3545 !important;
  font-weight: 500 !important;
  font-size: 0.75rem !important;
}

.period-text {
  color: #ffc107 !important;
  font-weight: 500 !important;
  font-size: 0.75rem !important;
  background: #fff3cd !important;
  padding: 0.2rem 0.4rem !important;
  border-radius: 4px !important;
}

.work-type-text {
  color: #6f42c1 !important;
  font-weight: 500 !important;
  font-size: 0.75rem !important;
}

.about-text {
  color: #17a2b8 !important;
  font-weight: 400 !important;
  font-size: 0.75rem !important;
  max-width: 200px !important;
  overflow: hidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.actions-buttons {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
  align-items: center;
  padding: 0.5rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.action-btn {
  min-width: 36px !important;
  height: 36px !important;
  border-radius: 6px !important;
  transition: all 0.2s ease !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1) !important;
}

.action-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15) !important;
}

.actions-buttons .v-btn[color="red"] {
  background: #1e3a8a !important;
  color: white !important;
}

.actions-buttons .v-btn[color="black"] {
  background: #6c757d !important;
  color: white !important;
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
</style>