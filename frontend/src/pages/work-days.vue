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



<style>
/* Import page styles */
@import './styles/work-days.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
