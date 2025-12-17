<template>
  <div class="work-day-details-page">
    <!-- Header Section -->
    <div class="d-flex align-center mb-4">
      <v-btn
        icon="mdi-arrow-left"
        @click="goBack"
        class="back-btn me-3"
        size="large"
        variant="text"
        color="primary"
      >
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>
    </div>
    <PageHeader
      title="تفاصيل يوم العمل"
      subtitle="معلومات مفصلة عن يوم العمل المحدد"
      mdi-icon="mdi-calendar-clock"
    />


    <!-- Work Day Info Card -->
    <v-card class="info-card mb-4" elevation="2">
      <v-card-title class="info-card-title">
        <v-icon class="me-2" color="primary">mdi-information</v-icon>
        معلومات يوم العمل
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12" md="6">
            <div class="info-item">
              <label>مكان العمل:</label>
              <span>{{ workDayInfo.location || 'موقع البناء الرئيسي' }}</span>
            </div>
          </v-col>
          <v-col cols="12" md="6">
            <div class="info-item">
              <label>رقم الاستمارة:</label>
              <span>{{ workDayInfo.formNumber || 'FORM-001' }}</span>
            </div>
          </v-col>
          <v-col cols="12" md="6">
            <div class="info-item">
              <label>فترة العمل:</label>
              <span>{{ workDayInfo.workPeriod || '08:00 - 16:00' }}</span>
            </div>
          </v-col>
          <v-col cols="12" md="6">
            <div class="info-item">
              <label>اليوم:</label>
              <span>{{ workDayInfo.day || 'الاثنين' }}</span>
            </div>
          </v-col>
          <v-col cols="12" md="6">
            <div class="info-item">
              <label>نوع العمل:</label>
              <span>{{ workDayInfo.workType || 'بناء' }}</span>
            </div>
          </v-col>
          <v-col cols="12" md="6">
            <div class="info-item">
              <label>التاريخ:</label>
              <span>{{ workDayInfo.date || '2024-01-15' }}</span>
            </div>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <!-- Categories Section -->
    <v-card class="categories-card mb-4" elevation="2">
      <v-card-title class="categories-title">
        <v-icon class="me-2" color="success">mdi-view-grid</v-icon>
        أقسام إدارة يوم العمل
        <v-spacer />
        <v-chip 
          :color="getOverallStatus().color" 
          variant="elevated"
          class="status-chip"
        >
          <v-icon :icon="getOverallStatus().icon" class="me-1" />
          {{ getOverallStatus().text }}
        </v-chip>
      </v-card-title>
      <v-card-text>
        <v-row>
          <!-- Materials & Expenses Combined -->
          <v-col cols="12" md="4">
            <v-card 
              class="category-card materials-expenses" 
              :class="{ 'disabled-card': !categoriesEnabled.materialsExpenses }"
              elevation="3"
              @click="categoriesEnabled.materialsExpenses ? showCategoryDetails('materials-expenses') : null"
            >
              <v-card-text class="text-center">
                <!-- Toggle Switch -->
                <div class="toggle-container">
                  <v-switch
                    v-model="categoriesEnabled.materialsExpenses"
                    color="success"
                    hide-details
                    @click.stop
                    class="toggle-switch"
                  />
                  <span class="toggle-label">{{ categoriesEnabled.materialsExpenses ? 'مفعل' : 'معطل' }}</span>
                </div>
                
                <div class="category-icon" :class="{ 'disabled-icon': !categoriesEnabled.materialsExpenses }">🔨💰</div>
                <h4 class="category-title" :class="{ 'disabled-text': !categoriesEnabled.materialsExpenses }">المواد والمصاريف اليومية</h4>
                <p class="category-description" :class="{ 'disabled-text': !categoriesEnabled.materialsExpenses }">مواد البناء والنفقات اليومية</p>
                <v-chip 
                  class="count-chip" 
                  :color="categoriesEnabled.materialsExpenses ? 'white' : 'grey'" 
                  size="large"
                >
                  5 مواد + 3 مصاريف
                </v-chip>
                <v-btn 
                  class="details-btn mt-2" 
                  :color="categoriesEnabled.materialsExpenses ? 'white' : 'grey'" 
                  variant="elevated"
                  :disabled="!categoriesEnabled.materialsExpenses"
                  @click.stop="showCategoryDetails('materials-expenses')"
                >
                  التفاصيل
                </v-btn>
              </v-card-text>
            </v-card>
          </v-col>

          <!-- Labor -->
          <v-col cols="12" md="4">
            <v-card 
              class="category-card labor" 
              :class="{ 'disabled-card': !categoriesEnabled.labor }"
              elevation="3"
              @click="categoriesEnabled.labor ? showCategoryDetails('labor') : null"
            >
              <v-card-text class="text-center">
                <!-- Toggle Switch -->
                <div class="toggle-container">
                  <v-switch
                    v-model="categoriesEnabled.labor"
                    color="success"
                    hide-details
                    @click.stop
                    class="toggle-switch"
                  />
                  <span class="toggle-label">{{ categoriesEnabled.labor ? 'مفعل' : 'معطل' }}</span>
                </div>
                
                <div class="category-icon" :class="{ 'disabled-icon': !categoriesEnabled.labor }">👷</div>
                <h4 class="category-title" :class="{ 'disabled-text': !categoriesEnabled.labor }">الأيدي العاملة</h4>
                <p class="category-description" :class="{ 'disabled-text': !categoriesEnabled.labor }">العمال والموظفين</p>
                <v-chip 
                  class="count-chip" 
                  :color="categoriesEnabled.labor ? 'white' : 'grey'" 
                  size="large"
                >
                  8 عامل
                </v-chip>
                <v-btn 
                  class="details-btn mt-2" 
                  :color="categoriesEnabled.labor ? 'white' : 'grey'" 
                  variant="elevated"
                  :disabled="!categoriesEnabled.labor"
                  @click.stop="showCategoryDetails('labor')"
                >
                  التفاصيل
                </v-btn>
              </v-card-text>
            </v-card>
          </v-col>

          <!-- Equipment -->
          <v-col cols="12" md="4">
            <v-card 
              class="category-card equipment" 
              :class="{ 'disabled-card': !categoriesEnabled.equipment }"
              elevation="3"
              @click="categoriesEnabled.equipment ? showCategoryDetails('equipment') : null"
            >
              <v-card-text class="text-center">
                <!-- Toggle Switch -->
                <div class="toggle-container">
                  <v-switch
                    v-model="categoriesEnabled.equipment"
                    color="success"
                    hide-details
                    @click.stop
                    class="toggle-switch"
                  />
                  <span class="toggle-label">{{ categoriesEnabled.equipment ? 'مفعل' : 'معطل' }}</span>
                </div>
                
                <div class="category-icon" :class="{ 'disabled-icon': !categoriesEnabled.equipment }">🚜</div>
                <h4 class="category-title" :class="{ 'disabled-text': !categoriesEnabled.equipment }">الآليات</h4>
                <p class="category-description" :class="{ 'disabled-text': !categoriesEnabled.equipment }">المعدات والآلات</p>
                <v-chip 
                  class="count-chip" 
                  :color="categoriesEnabled.equipment ? 'white' : 'grey'" 
                  size="large"
                >
                  4 آلة
                </v-chip>
                <v-btn 
                  class="details-btn mt-2" 
                  :color="categoriesEnabled.equipment ? 'white' : 'grey'" 
                  variant="elevated"
                  :disabled="!categoriesEnabled.equipment"
                  @click.stop="showCategoryDetails('equipment')"
                >
                  التفاصيل
                </v-btn>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <!-- Actions Section -->
    <v-card class="actions-card" elevation="2">
      <v-card-text class="text-center">
        <v-btn 
          color="primary" 
          size="large" 
          class="me-2"
          @click="editWorkDay"
        >
          <v-icon class="me-2">mdi-pencil</v-icon>
          تعديل يوم العمل
        </v-btn>
        <v-btn 
          color="error" 
          size="large" 
          class="ms-2"
          @click="deleteWorkDay"
        >
          <v-icon class="me-2">mdi-delete</v-icon>
          حذف يوم العمل
        </v-btn>
      </v-card-text>
    </v-card>

    <!-- Category Details Dialog -->
    <v-dialog v-model="showCategoryDialog" max-width="800" persistent scrollable>
      <v-card class="image-style-dialog">
        <!-- Header Section -->
        <div class="dialog-header">
          <div class="header-content">
            <div class="header-left">
              <v-icon size="24" color="white" class="header-icon">{{ selectedCategory.icon }}</v-icon>
              <span class="header-title">{{ selectedCategory.title }}</span>
            </div>
            <v-btn
              icon="mdi-close"
              variant="text"
              size="small"
              color="white"
              @click="showCategoryDialog = false"
              class="close-btn"
            />
          </div>
        </div>

        <!-- Dialog Body -->
        <div class="dialog-body">
          <p class="mb-4">{{ selectedCategory.description }}</p>
          <v-alert type="info" variant="tonal" class="mt-4">
            <v-icon class="me-2">mdi-information</v-icon>
            هنا ستظهر التفاصيل الكاملة لهذا القسم
          </v-alert>

          <!-- تفاصيل إضافية حسب النوع -->
          <div v-if="selectedCategory.title === 'المواد والمصاريف اليومية'" class="mt-4">
            <h4>قائمة المواد:</h4>
            <v-list>
              <v-list-item>أسمنت - 50 كيس</v-list-item>
              <v-list-item>رمل - 10 متر مكعب</v-list-item>
              <v-list-item>حصى - 8 متر مكعب</v-list-item>
              <v-list-item>حديد - 2 طن</v-list-item>
              <v-list-item>طوب - 1000 قطعة</v-list-item>
            </v-list>
          </div>

          <div v-else-if="selectedCategory.title === 'الأيدي العاملة'" class="mt-4">
            <h4>قائمة العمال:</h4>
            <v-list>
              <v-list-item>أحمد محمد - مهندس</v-list-item>
              <v-list-item>محمد علي - عامل بناء</v-list-item>
              <v-list-item>علي حسن - عامل بناء</v-list-item>
              <v-list-item>حسن أحمد - عامل بناء</v-list-item>
              <v-list-item>أحمد علي - عامل بناء</v-list-item>
            </v-list>
          </div>

          <div v-else-if="selectedCategory.title === 'الآليات'" class="mt-4">
            <h4>قائمة الآليات:</h4>
            <v-list>
              <v-list-item>حفار - 1 آلة</v-list-item>
              <v-list-item>رافعة - 1 آلة</v-list-item>
              <v-list-item>خلاطة - 1 آلة</v-list-item>
              <v-list-item>شاحنة - 1 آلة</v-list-item>
            </v-list>
          </div>
        </div>

        <!-- Dialog Actions -->
        <div class="dialog-actions">
          <v-btn
            color="primary"
            variant="elevated"
            @click="showCategoryDialog = false"
            class="save-btn"
          >
            <v-icon class="me-2">mdi-check</v-icon>
            إغلاق
          </v-btn>
        </div>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { toast } from 'vue3-toastify'
import { PageHeader } from '@/components/shared'

const router = useRouter()

// Reactive data for categories toggle
const categoriesEnabled = ref({
  materialsExpenses: true,
  labor: true,
  equipment: true
})

// حفظ حالة الكروت في localStorage
const saveCategoriesState = () => {
  localStorage.setItem('categoriesEnabled', JSON.stringify(categoriesEnabled.value))
}

// تحميل حالة الكروت من localStorage
const loadCategoriesState = () => {
  const saved = localStorage.getItem('categoriesEnabled')
  if (saved) {
    try {
      categoriesEnabled.value = JSON.parse(saved)
    } catch (error) {
      console.error('خطأ في تحميل حالة الكروت:', error)
    }
  }
}

// Reactive data
const showCategoryDialog = ref(false)
const selectedCategory = ref({})

const workDayInfo = ref({
  location: 'موقع البناء الرئيسي',
  formNumber: 'FORM-001',
  workPeriod: '08:00 - 16:00',
  day: 'الاثنين',
  workType: 'بناء',
  date: '2024-01-15'
})

// Methods
const goBack = () => {
  router.push('/work-days')
}

// دالة للحصول على الحالة العامة للكروت
const getOverallStatus = () => {
  const enabledCount = Object.values(categoriesEnabled.value).filter(Boolean).length
  const totalCount = Object.keys(categoriesEnabled.value).length
  
  if (enabledCount === totalCount) {
    return {
      text: 'جميع الأقسام مفعلة',
      color: 'success',
      icon: 'mdi-check-circle'
    }
  } else if (enabledCount === 0) {
    return {
      text: 'جميع الأقسام معطلة',
      color: 'error',
      icon: 'mdi-close-circle'
    }
  } else {
    return {
      text: `${enabledCount} من ${totalCount} مفعلة`,
      color: 'warning',
      icon: 'mdi-alert-circle'
    }
  }
}

const showCategoryDetails = (category) => {
  if (category === 'materials-expenses') {
    // الانتقال إلى صفحة المواد والمصاريف
    router.push('/materials-expenses-details')
    return
  }
  
  if (category === 'labor') {
    // الانتقال إلى صفحة الأيدي العاملة
    router.push('/labor-details')
    return
  }
  
  if (category === 'equipment') {
    // الانتقال إلى صفحة الآليات
    router.push('/equipment-details')
    return
  }
  
  const categories = {}
  
  selectedCategory.value = categories[category] || {}
  
  // إضافة تأخير صغير للتأكد من التحديث
  setTimeout(() => {
    showCategoryDialog.value = true
  }, 100)
}

const editWorkDay = () => {
  toast.info('تم الضغط على: تعديل يوم العمل')
}

const deleteWorkDay = () => {
  if (confirm('هل أنت متأكد من حذف يوم العمل؟')) {
    toast.success('تم حذف يوم العمل')
  }
}

// Lifecycle
// مراقبة تغييرات حالة الكروت وحفظها تلقائياً
watch(categoriesEnabled, () => {
  saveCategoriesState()
}, { deep: true })

onMounted(() => {
  // تحميل حالة الكروت المحفوظة
  loadCategoriesState()
})
</script>


<style scoped>
/* Import page styles - scoped to this component only */
@import './styles/work-day-details.css';
</style>
