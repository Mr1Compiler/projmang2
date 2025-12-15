<template>
  <div class="labor-details-page">
    <!-- Header Section -->
    <div class="page-header glass-effect gradient-animation">
      <div class="header-content">
        <v-btn 
          icon="mdi-arrow-left" 
          @click="goBack" 
          class="back-btn"
          size="large"
          color="white"
        >
          <v-icon>mdi-arrow-left</v-icon>
        </v-btn>
        <div class="header-text">
          <h1 class="page-title text-glow fade-in">تفاصيل الأيدي العاملة</h1>
          <p class="page-subtitle fade-in">إدارة العمال والموظفين المشاركين في المشروع</p>
        </div>
      </div>
    </div>

    <!-- Labor Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="success">mdi-account-group</v-icon>
        الأيدي العاملة
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="laborSearch"
              label="البحث في العمال..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchLabor"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchLabor"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="success" 
              variant="elevated" 
              class="add-btn"
              @click="openAddDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة عامل
            </v-btn>
          </div>
        </div>
      </v-card-text>
    </v-card>

    <!-- Labor Table Card -->
    <v-card class="data-table-card" elevation="2">
      <v-card-title class="table-title indigo-title">
        <span class="title-text">قائمة العمال</span>
      </v-card-title>

      <v-data-table
        :headers="laborHeaders"
        :items="laborData"
        :search="laborSearch"
        :items-per-page="10"
        class="project-table"
        no-data-text="لا توجد بيانات متاحة"
        loading-text="جاري التحميل..."
        hover
      >
        <!-- Serial Number Column -->
        <template #item.id="{ index }">
          <span class="serial-number">{{ index + 1 }}</span>
        </template>

        <!-- Name Column -->
        <template #item.name="{ item }">
          <span class="project-name">{{ item.name }}</span>
        </template>

        <!-- Position Column -->
        <template #item.position="{ item }">
          <span class="project-name">{{ item.position }}</span>
        </template>

        <!-- Experience Column -->
        <template #item.experience="{ item }">
          <span class="date-text">{{ item.experience }}</span>
        </template>

        <!-- Salary Column -->
        <template #item.salary="{ item }">
          <span class="cost-text">{{ formatCurrency(item.salary) }}</span>
        </template>

        <!-- Phone Column -->
        <template #item.phone="{ item }">
          <span class="date-text">{{ item.phone }}</span>
        </template>

        <!-- Specialty Column -->
        <template #item.specialty="{ item }">
          <v-chip class="category-chip" size="small">
            {{ item.specialty }}
          </v-chip>
        </template>

        <!-- Actions Column -->
        <template #item.actions="{ item }">
          <div class="action-buttons">
            <v-btn
              icon="mdi-delete"
              size="small"
              color="error"
              variant="text"
              @click="deleteLabor(item)"
              title="حذف العامل"
              class="action-btn"
            />
          </div>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add New Labor Dialog -->
    <v-dialog v-model="showAddDialog" max-width="600" persistent>
      <v-card class="modal-card">
        <v-card-title class="modal-header">
          <span class="modal-title">إضافة عامل جديد</span>
          <v-btn icon="mdi-close" @click="closeAddDialog" variant="text" class="close-btn">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-title>
        <v-card-text class="modal-content">
          <v-form ref="form" v-model="formValid">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newLabor.name"
                  label="اسم العامل"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم العامل مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newLabor.position"
                  label="المنصب/الوظيفة"
                  variant="outlined"
                  :rules="[v => !!v || 'المنصب مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newLabor.experience"
                  label="سنوات الخبرة"
                  variant="outlined"
                  type="number"
                  :rules="[v => !!v || 'سنوات الخبرة مطلوبة']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newLabor.salary"
                  label="الراتب اليومي"
                  variant="outlined"
                  type="number"
                  :rules="[v => !!v || 'الراتب مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newLabor.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                  :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newLabor.specialty"
                  label="التخصص"
                  variant="outlined"
                  :rules="[v => !!v || 'التخصص مطلوب']"
                  required
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions class="modal-footer">
          <v-spacer />
          <v-btn @click="closeAddDialog" color="grey" variant="text">
            إلغاء
          </v-btn>
          <v-btn @click="saveLabor" color="primary" variant="elevated" :disabled="!formValid" class="save-btn">
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Simple Dialog for Add Worker -->
    <SimpleDialog
      :open="showAddDialog"
      title="إضافة عامل"
      save-text="حفظ"
      icon="mdi-account-plus"
      @close="closeAddDialog"
      @save="saveWorker"
    >
      <LaborForm
        v-model="newWorker"
        ref="laborFormRef"
      />
    </SimpleDialog>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import SimpleDialog from '@/components/SimpleDialog.vue'
import LaborForm from '@/components/LaborForm.vue'

const router = useRouter()

// Reactive data
const laborSearch = ref('')
const showAddDialog = ref(false)
const formValid = ref(false)

// Modern dialog data
const newWorker = ref({
  name: '',
  profession: '',
  salary: '',
  daysWorked: '',
  totalSalary: '',
  phone: '',
  address: '',
  notes: ''
})

const laborFormRef = ref(null)

// Form data
const newLabor = ref({
  name: '',
  position: '',
  experience: '',
  salary: '',
  phone: '',
  specialty: ''
})

// Table headers
const laborHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'اسم العامل', key: 'name', sortable: true },
  { title: 'المنصب', key: 'position', sortable: true },
  { title: 'سنوات الخبرة', key: 'experience', sortable: true },
  { title: 'الراتب اليومي', key: 'salary', sortable: true },
  { title: 'رقم الهاتف', key: 'phone', sortable: true },
  { title: 'التخصص', key: 'specialty', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Sample data
const laborData = ref([
  {
    id: 1,
    name: 'أحمد محمد علي',
    position: 'مهندس مدني',
    experience: 5,
    salary: 50000,
    phone: '07701234567',
    specialty: 'البناء والتشييد'
  },
  {
    id: 2,
    name: 'محمد حسن أحمد',
    position: 'عامل بناء',
    experience: 3,
    salary: 25000,
    phone: '07801234567',
    specialty: 'البناء العام'
  },
  {
    id: 3,
    name: 'علي أحمد محمد',
    position: 'عامل بناء',
    experience: 2,
    salary: 20000,
    phone: '07901234567',
    specialty: 'البناء العام'
  },
  {
    id: 4,
    name: 'حسن علي محمد',
    position: 'عامل بناء',
    experience: 4,
    salary: 30000,
    phone: '07501234567',
    specialty: 'البناء العام'
  },
  {
    id: 5,
    name: 'أحمد علي حسن',
    position: 'عامل بناء',
    salary: 22000,
    experience: 1,
    phone: '07301234567',
    specialty: 'البناء العام'
  }
])

// Methods
const goBack = () => {
  router.push('/work-day-details')
}

const searchLabor = () => {
}

const openAddDialog = () => {
  newLabor.value = {
    name: '',
    position: '',
    experience: '',
    salary: '',
    phone: '',
    specialty: ''
  }
  // Reset modern form data
  newWorker.value = {
    name: '',
    profession: '',
    salary: '',
    daysWorked: '',
    totalSalary: '',
    phone: '',
    address: '',
    notes: ''
  }
  showAddDialog.value = true
}

const closeAddDialog = () => {
  showAddDialog.value = false
  newLabor.value = {
    name: '',
    position: '',
    experience: '',
    salary: '',
    phone: '',
    specialty: ''
  }
}

const saveLabor = () => {
  const newWorker = {
    id: laborData.value.length + 1,
    name: newLabor.value.name,
    position: newLabor.value.position,
    experience: parseInt(newLabor.value.experience),
    salary: parseInt(newLabor.value.salary),
    phone: newLabor.value.phone,
    specialty: newLabor.value.specialty
  }
  laborData.value.push(newWorker)
  closeAddDialog()
}

// Modern dialog function for workers
const saveWorker = () => {
  if (laborFormRef.value && laborFormRef.value.isValid()) {
    const newWorkerItem = {
      id: laborData.value.length + 1,
      name: newWorker.value.name,
      position: newWorker.value.profession,
      experience: `${newWorker.value.daysWorked} يوم`,
      salary: parseInt(newWorker.value.salary),
      phone: newWorker.value.phone,
      specialty: newWorker.value.address,
      status: 'نشط',
      totalSalary: parseInt(newWorker.value.totalSalary),
      notes: newWorker.value.notes
    }
    
    laborData.value.push(newWorkerItem)
    
    // Reset form
    laborFormRef.value.resetForm()
    newWorker.value = {
      name: '',
      profession: '',
      salary: '',
      daysWorked: '',
      totalSalary: '',
      phone: '',
      address: '',
      notes: ''
    }
    
    showAddDialog.value = false
  }
}

const deleteLabor = (item) => {
  if (confirm('هل أنت متأكد من حذف هذا العامل؟')) {
    const index = laborData.value.findIndex(l => l.id === item.id)
    if (index > -1) {
      laborData.value.splice(index, 1)
    }
  }
}

// Format currency
const formatCurrency = (value) => {
  if (!value) return '0 د.ع'
  return new Intl.NumberFormat('ar-IQ', {
    style: 'currency',
    currency: 'IQD',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(value).replace('IQD', 'د.ع')
}

// Lifecycle
onMounted(() => {
})
</script>


<style>
/* Import page styles */
@import './styles/labor-details.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
