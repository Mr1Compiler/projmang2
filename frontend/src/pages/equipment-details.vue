<template>
  <div class="equipment-details-page">
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
          <h1 class="page-title text-glow fade-in">تفاصيل الآليات</h1>
          <p class="page-subtitle fade-in">إدارة المعدات والآلات المستخدمة في المشروع</p>
        </div>
      </div>
    </div>

    <!-- Equipment Section -->
    <v-card class="section-card mb-4" elevation="2">
      <v-card-title class="section-title">
        <v-icon class="me-2" color="indigo">mdi-truck</v-icon>
        الآليات والمعدات
      </v-card-title>
      
      <!-- Search Bar and Add Button -->
      <v-card-text class="pb-0">
        <div class="search-container">
          <div class="search-box">
            <v-text-field
              v-model="equipmentSearch"
              label="البحث في الآليات..."
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              clearable
              hide-details
              class="search-field"
              @keyup.enter="searchEquipment"
            />
            <v-btn 
              color="primary" 
              variant="elevated" 
              class="search-btn"
              @click="searchEquipment"
              size="large"
            >
              <v-icon class="me-2">mdi-magnify</v-icon>
              بحث
            </v-btn>
            <v-btn 
              color="indigo" 
              variant="elevated" 
              class="add-btn"
              @click="openAddDialog"
              size="large"
            >
              <v-icon class="me-2">mdi-plus</v-icon>
              إضافة آلة
            </v-btn>
          </div>
        </div>
      </v-card-text>

      <!-- Equipment Table -->
      <v-data-table
        :headers="equipmentHeaders"
        :items="equipmentData"
        :search="equipmentSearch"
        :items-per-page="10"
        class="data-table equipment-table"
        no-data-text="لا توجد بيانات متاحة"
        loading-text="جاري التحميل..."
        density="comfortable"
      >
        <template v-slot:item.actions="{ item }">
          <v-btn
            icon="mdi-delete"
            size="small"
            color="error"
            variant="text"
            @click="deleteEquipment(item)"
            title="حذف الآلة"
          />
        </template>
      </v-data-table>
    </v-card>

    <!-- Add New Equipment Dialog -->
    <v-dialog v-model="showAddDialog" max-width="600" persistent>
      <v-card class="dialog-card">
        <v-card-title class="dialog-title">
          <v-icon class="me-2">mdi-truck-plus</v-icon>
          إضافة آلة جديدة
        </v-card-title>
        <v-card-text class="dialog-content">
          <v-form ref="form" v-model="formValid">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.name"
                  label="اسم الآلة"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم الآلة مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.type"
                  label="نوع الآلة"
                  variant="outlined"
                  :rules="[v => !!v || 'نوع الآلة مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.model"
                  label="الموديل"
                  variant="outlined"
                  :rules="[v => !!v || 'الموديل مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.rentalCost"
                  label="تكلفة الإيجار اليومية"
                  variant="outlined"
                  type="number"
                  :rules="[v => !!v || 'تكلفة الإيجار مطلوبة']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.operator"
                  label="اسم السائق/المشغل"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم المشغل مطلوب']"
                  required
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="newEquipment.status"
                  label="الحالة"
                  variant="outlined"
                  :rules="[v => !!v || 'الحالة مطلوبة']"
                  required
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <v-btn @click="closeAddDialog" color="grey">
            إلغاء
          </v-btn>
          <v-btn @click="saveEquipment" color="primary" :disabled="!formValid">
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Simple Dialog for Add Equipment -->
    <SimpleDialog
      :open="showAddDialog"
      title="إضافة آلة"
      save-text="حفظ"
      icon="mdi-truck-plus"
      @close="closeAddDialog"
      @save="saveEquipmentModern"
    >
      <EquipmentForm
        v-model="newEquipment"
        ref="equipmentFormRef"
      />
    </SimpleDialog>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import SimpleDialog from '@/components/SimpleDialog.vue'
import EquipmentForm from '@/components/EquipmentForm.vue'

const router = useRouter()

// Reactive data
const equipmentSearch = ref('')
const showAddDialog = ref(false)
const formValid = ref(false)

// Modern dialog data
const newEquipment = ref({
  type: '',
  dailyCost: '',
  daysUsed: '',
  totalCost: '',
  driverName: '',
  driverPhone: '',
  notes: ''
})

const equipmentFormRef = ref(null)

// Form data
const newEquipmentItem = ref({
  name: '',
  type: '',
  model: '',
  rentalCost: '',
  operator: '',
  status: ''
})

// Table headers
const equipmentHeaders = [
  { title: 'التسلسل', key: 'id', sortable: false },
  { title: 'اسم الآلة', key: 'name', sortable: true },
  { title: 'نوع الآلة', key: 'type', sortable: true },
  { title: 'الموديل', key: 'model', sortable: true },
  { title: 'تكلفة الإيجار اليومية', key: 'rentalCost', sortable: true },
  { title: 'المشغل', key: 'operator', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false }
]

// Sample data
const equipmentData = ref([
  {
    id: 1,
    name: 'حفار هيدروليكي',
    type: 'حفار',
    model: 'CAT 320D',
    rentalCost: 150000,
    operator: 'أحمد محمد',
    status: 'نشط'
  },
  {
    id: 2,
    name: 'رافعة برجية',
    type: 'رافعة',
    model: 'Potain MDT 178',
    rentalCost: 200000,
    operator: 'محمد علي',
    status: 'نشط'
  },
  {
    id: 3,
    name: 'خلاطة خرسانة',
    type: 'خلاطة',
    model: 'Schwing SP 1000',
    rentalCost: 80000,
    operator: 'علي حسن',
    status: 'نشط'
  },
  {
    id: 4,
    name: 'شاحنة نقل',
    type: 'شاحنة',
    model: 'Mercedes Actros',
    rentalCost: 60000,
    operator: 'حسن أحمد',
    status: 'نشط'
  },
  {
    id: 5,
    name: 'جرافة',
    type: 'جرافة',
    model: 'Caterpillar D6T',
    rentalCost: 120000,
    operator: 'أحمد علي',
    status: 'صيانة'
  }
])

// Methods
const goBack = () => {
  router.push('/work-day-details')
}

const searchEquipment = () => {
}

const openAddDialog = () => {
  newEquipmentItem.value = {
    name: '',
    type: '',
    model: '',
    rentalCost: '',
    operator: '',
    status: ''
  }
  // Reset modern form data
  newEquipment.value = {
    type: '',
    dailyCost: '',
    daysUsed: '',
    totalCost: '',
    driverName: '',
    driverPhone: '',
    notes: ''
  }
  showAddDialog.value = true
}

const closeAddDialog = () => {
  showAddDialog.value = false
  newEquipment.value = {
    name: '',
    type: '',
    model: '',
    rentalCost: '',
    operator: '',
    status: ''
  }
}

const saveEquipment = () => {
  const newMachine = {
    id: equipmentData.value.length + 1,
    name: newEquipmentItem.value.name,
    type: newEquipmentItem.value.type,
    model: newEquipmentItem.value.model,
    rentalCost: parseInt(newEquipmentItem.value.rentalCost),
    operator: newEquipmentItem.value.operator,
    status: newEquipmentItem.value.status
  }
  equipmentData.value.push(newMachine)
  closeAddDialog()
}

// Modern dialog function for equipment
const saveEquipmentModern = () => {
  if (equipmentFormRef.value && equipmentFormRef.value.isValid()) {
    const newMachine = {
      id: equipmentData.value.length + 1,
      name: newEquipment.value.type,
      type: newEquipment.value.type,
      model: '-',
      rentalCost: parseInt(newEquipment.value.dailyCost),
      operator: newEquipment.value.driverName,
      status: 'نشط',
      daysUsed: parseInt(newEquipment.value.daysUsed),
      totalCost: parseInt(newEquipment.value.totalCost),
      driverPhone: newEquipment.value.driverPhone,
      notes: newEquipment.value.notes
    }
    
    equipmentData.value.push(newMachine)
    
    // Reset form
    equipmentFormRef.value.resetForm()
    newEquipment.value = {
      type: '',
      dailyCost: '',
      daysUsed: '',
      totalCost: '',
      driverName: '',
      driverPhone: '',
      notes: ''
    }
    
    showAddDialog.value = false
  }
}

const deleteEquipment = (item) => {
  if (confirm('هل أنت متأكد من حذف هذه الآلة؟')) {
    const index = equipmentData.value.findIndex(e => e.id === item.id)
    if (index > -1) {
      equipmentData.value.splice(index, 1)
    }
  }
}

// Lifecycle
onMounted(() => {
})
</script>


<style>
/* Import page styles */
@import './styles/equipment-details.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
