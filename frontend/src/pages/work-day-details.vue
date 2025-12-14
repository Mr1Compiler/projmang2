<template>
  <div class="work-day-details-page">
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
          <h1 class="page-title text-glow fade-in">تفاصيل يوم العمل</h1>
          <p class="page-subtitle fade-in">معلومات مفصلة عن يوم العمل المحدد</p>
        </div>
      </div>
    </div>


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
      <v-card class="dialog-card">
        <v-card-title class="dialog-title">
          <v-icon class="me-2" color="primary">{{ selectedCategory.icon }}</v-icon>
          {{ selectedCategory.title }}
        </v-card-title>
        <v-card-text class="dialog-content">
          <p class="dialog-description">{{ selectedCategory.description }}</p>
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
        </v-card-text>
        <v-card-actions class="dialog-actions">
          <v-spacer></v-spacer>
          <v-btn @click="showCategoryDialog = false" color="primary" variant="elevated">
            <v-icon class="me-2">mdi-close</v-icon>
            إغلاق
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'

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
  console.log('تم الضغط على:', category) // للتأكد من أن الدالة تعمل
  
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
    console.log('النافذة المنبثقة:', showCategoryDialog.value) // للتأكد من التحديث
  }, 100)
}

const editWorkDay = () => {
  alert('تم الضغط على: تعديل يوم العمل')
}

const deleteWorkDay = () => {
  if (confirm('هل أنت متأكد من حذف يوم العمل؟')) {
    alert('تم الضغط على: حذف يوم العمل')
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
  
  console.log('✅ صفحة تفاصيل يوم العمل تم تحميلها بنجاح!')
  console.log('جميع العناصر ظاهرة ومتاحة للاستخدام')
})
</script>

<style scoped>
.work-day-details-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #e3f2fd 0%, #bbdefb 25%, #90caf9 50%, #64b5f6 75%, #42a5f5 100%);
  padding: 2rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  position: relative;
  overflow: hidden;
}

.work-day-details-page::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 20% 80%, rgba(120, 144, 156, 0.1) 0%, transparent 50%),
              radial-gradient(circle at 80% 20%, rgba(100, 181, 246, 0.1) 0%, transparent 50%),
              radial-gradient(circle at 40% 40%, rgba(66, 165, 245, 0.1) 0%, transparent 50%);
  pointer-events: none;
}

.page-header {
  background: linear-gradient(135deg, #1976d2 0%, #1e88e5 25%, #2196f3 50%, #42a5f5 75%, #64b5f6 100%);
  background-size: 400% 400%;
  animation: gradientShift 8s ease infinite;
  color: white;
  padding: 2rem;
  border-radius: 20px;
  margin-bottom: 4rem;
  position: relative;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(25, 118, 210, 0.3), 0 0 0 1px rgba(255, 255, 255, 0.3);
  border: 2px solid rgba(255, 255, 255, 0.4);
  backdrop-filter: blur(15px);
  z-index: 1;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 1rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.back-btn {
  background: rgba(255, 255, 255, 0.25) !important;
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 600;
  font-size: 1rem;
  letter-spacing: 0.5px;
  line-height: 1.2;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.35) !important;
  transform: translateX(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
}

.page-title {
  font-size: 3rem;
  font-weight: 800;
  margin: 0;
  text-shadow: 0 4px 12px rgba(0, 0, 0, 0.5), 0 0 0 2px rgba(0, 0, 0, 0.2);
  color: #ffffff !important;
  letter-spacing: 2px;
  text-stroke: 1px rgba(0, 0, 0, 0.3);
  -webkit-text-stroke: 1px rgba(0, 0, 0, 0.3);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.2;
  background: linear-gradient(135deg, #1a237e 0%, #283593 50%, #3949ab 100%);
  padding: 0.5rem 1rem;
  border-radius: 15px;
  display: inline-block;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 25px rgba(26, 35, 126, 0.4);
}

.page-subtitle {
  font-size: 1.6rem;
  opacity: 1;
  margin: 0.5rem 0 0 0;
  color: #ffffff !important;
  text-shadow: 0 4px 15px rgba(0, 0, 0, 0.8), 0 0 0 2px rgba(255, 255, 255, 0.1);
  font-weight: 600;
  letter-spacing: 1px;
  text-stroke: 0.5px rgba(255, 255, 255, 0.2);
  -webkit-text-stroke: 0.5px rgba(255, 255, 255, 0.2);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.3;
}

.info-card, .categories-card, .actions-card {
  border-radius: 16px;
  overflow: hidden;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: rgba(255, 255, 255, 0.8);
  box-shadow: 0 8px 25px rgba(25, 118, 210, 0.15), 0 0 0 1px rgba(255, 255, 255, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.4);
  backdrop-filter: blur(20px);
  position: relative;
  z-index: 1;
  margin-top: 3rem;
}

.info-card-title, .categories-title {
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  background-size: 200% 200%;
  animation: gradientShift 3s ease infinite;
  color: #ffffff !important;
  font-weight: 800;
  font-size: 1.6rem;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.3), 0 0 20px rgba(255, 255, 255, 0.2);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 1px;
  line-height: 1.3;
  margin-bottom: 2rem;
  padding: 1.25rem 1.5rem !important;
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(59, 130, 246, 0.4);
}

.categories-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: shimmer 3s ease-in-out infinite;
  pointer-events: none;
}

.categories-title::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.15) 0%, transparent 70%);
  animation: pulse 4s ease-in-out infinite;
  pointer-events: none;
}

@keyframes shimmer {
  0% { left: -100%; }
  100% { left: 100%; }
}

@keyframes pulse {
  0%, 100% { 
    opacity: 0.3; 
    transform: scale(0.8); 
  }
  50% { 
    opacity: 0.6; 
    transform: scale(1.2); 
  }
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: rgba(255, 255, 255, 0.7);
  border-radius: 8px;
  margin-bottom: 0.5rem;
  border-right: 4px solid #3c41d1;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  box-shadow: 0 2px 8px rgba(25, 118, 210, 0.1);
  backdrop-filter: blur(10px);
}

.info-item label {
  font-weight: 600;
  color: #2c3e50;
  font-size: 1.2rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px;
  line-height: 1.3;
}

.info-item span {
  color: #34495e;
  font-size: 1.2rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 500;
  letter-spacing: 0.3px;
  line-height: 1.3;
}

.category-card {
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  height: 100%;
  min-height: 300px;
  border-radius: 20px;
  overflow: hidden;
  position: relative;
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 255, 255, 0.2);
  animation: float 4s ease-in-out infinite;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  margin-bottom: 2rem;
  margin-top: 1rem;
}

.category-card:hover {
  transform: translateY(-15px) scale(1.03);
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(255, 255, 255, 0.4) !important;
  animation: pulse 1.5s ease-in-out infinite;
}

.materials-expenses {
  background: linear-gradient(135deg, #1976d2 0%, #1e88e5 25%, #2196f3 50%, #42a5f5 75%, #64b5f6 100%) !important;
  background-size: 400% 400%;
  animation: gradientShift 5s ease infinite;
  color: white;
  position: relative;
  overflow: hidden;
}

.materials-expenses::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  animation: shimmer 3s infinite;
  border-radius: 20px;
  z-index: 1;
}

.materials-expenses .category-icon {
  animation: bounce 2s ease-in-out infinite, pulse 3s ease-in-out infinite;
}

.materials-expenses .category-icon:hover {
  transform: scale(1.3) rotate(180deg);
  animation-play-state: paused;
}

.materials-expenses::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmerDiagonal 4s infinite;
  border-radius: 20px;
  z-index: 2;
}

.labor {
  background: linear-gradient(135deg, #1976d2 0%, #1e88e5 25%, #2196f3 50%, #42a5f5 75%, #64b5f6 100%) !important;
  background-size: 400% 400%;
  animation: gradientShift 6s ease infinite;
  color: white;
  position: relative;
  overflow: hidden;
}

.labor::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  animation: shimmer 3.5s infinite;
  border-radius: 20px;
  z-index: 1;
}

.labor .category-icon {
  animation: float 3s ease-in-out infinite, wiggle 2s ease-in-out infinite;
}

.labor .category-icon:hover {
  transform: scale(1.2) rotate(360deg);
  animation-play-state: paused;
}

.labor::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmerDiagonal 4.5s infinite;
  border-radius: 20px;
  z-index: 2;
}

.equipment {
  background: linear-gradient(135deg, #1976d2 0%, #1e88e5 25%, #2196f3 50%, #42a5f5 75%, #64b5f6 100%) !important;
  background-size: 400% 400%;
  animation: gradientShift 7s ease infinite;
  color: white;
  position: relative;
  overflow: hidden;
}

.equipment::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
  animation: shimmer 4s infinite;
  border-radius: 20px;
  z-index: 1;
}

.equipment .category-icon {
  animation: rotate 4s linear infinite, bounce 2.5s ease-in-out infinite;
}

.equipment .category-icon:hover {
  transform: scale(1.25) rotate(720deg);
  animation-play-state: paused;
}

.equipment::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmerDiagonal 5s infinite;
  border-radius: 20px;
  z-index: 2;
}

.category-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  animation: bounce 2s ease-in-out infinite, rotate 4s linear infinite;
  transition: all 0.3s ease;
  display: inline-block;
}

.category-icon:hover {
  transform: scale(1.2) rotate(360deg);
  animation-play-state: paused;
}

.category-title {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
  text-shadow: 0 5px 12px rgba(0, 0, 0, 0.7), 0 0 0 2px rgba(255, 255, 255, 0.2);
  color: #ffffff !important;
  text-stroke: 0.5px rgba(255, 255, 255, 0.3);
  -webkit-text-stroke: 0.5px rgba(255, 255, 255, 0.3);
  letter-spacing: 1px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.2;
}

.category-description {
  opacity: 1;
  margin-bottom: 1rem;
  text-shadow: 0 3px 8px rgba(0, 0, 0, 0.6);
  color: #ffffff !important;
  font-weight: 500;
  font-size: 1.2rem;
  text-stroke: 0.3px rgba(255, 255, 255, 0.2);
  -webkit-text-stroke: 0.3px rgba(255, 255, 255, 0.2);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  line-height: 1.4;
  letter-spacing: 0.5px;
}

.count-chip {
  font-weight: 700;
  font-size: 1.3rem;
  color: #1a1a1a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  background: rgba(255, 255, 255, 0.95) !important;
  border: 2px solid rgba(0, 0, 0, 0.1) !important;
  backdrop-filter: blur(10px);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px;
  line-height: 1.2;
}

.details-btn {
  font-weight: 600;
  font-size: 1.2rem;
  color: #1a1a1a !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  background: rgba(255, 255, 255, 0.95) !important;
  border: 2px solid rgba(0, 0, 0, 0.1) !important;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px;
  line-height: 1.2;
}

.details-btn:hover {
  background: rgba(255, 255, 255, 1) !important;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.actions-card {
  background: white;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

/* Animations */
.glass-effect {
  backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.1);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.gradient-animation {
  background-size: 400% 400%;
  animation: gradientShift 15s ease infinite;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.text-glow {
  text-shadow: 0 0 20px rgba(255, 255, 255, 0.5);
}

.fade-in {
  animation: fadeIn 1s ease-in;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
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

@keyframes rotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

@keyframes wiggle {
  0%, 7% { transform: rotate(0deg); }
  15% { transform: rotate(-5deg); }
  20% { transform: rotate(3deg); }
  25% { transform: rotate(-3deg); }
  30% { transform: rotate(2deg); }
  35% { transform: rotate(-1deg); }
  40%, 100% { transform: rotate(0deg); }
}

@keyframes shimmerDiagonal {
  0% { transform: translateX(-100%) translateY(-100%); }
  100% { transform: translateX(100%) translateY(100%); }
}

@keyframes gradientShift {
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

@keyframes shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
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

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

/* Dialog Styles */
.dialog-card {
  border-radius: 16px;
  overflow: hidden;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  background: rgba(255, 255, 255, 0.85);
  box-shadow: 0 12px 30px rgba(25, 118, 210, 0.2), 0 0 0 1px rgba(255, 255, 255, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(25px);
  position: relative;
  z-index: 1;
}

.dialog-title {
  background: linear-gradient(135deg, #3f51b5 0%, #5c6bc0 100%);
  color: white;
  font-weight: 600;
  font-size: 1.4rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  letter-spacing: 0.5px;
  line-height: 1.3;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.dialog-content {
  padding: 2rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
}

.dialog-description {
  font-size: 1.2rem;
  color: #34495e;
  margin-bottom: 1rem;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  font-weight: 500;
  letter-spacing: 0.3px;
  line-height: 1.4;
}

.dialog-actions {
  padding: 1rem 2rem;
  background: rgba(248, 249, 250, 0.7);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  backdrop-filter: blur(10px);
}

/* Responsive */
@media (max-width: 768px) {
  .work-day-details-page {
    padding: 1rem;
    font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  }
  
  .page-title {
    font-size: 2rem;
    font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
    font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  }
  
  .info-item {
    flex-direction: column;
    text-align: center;
    gap: 0.5rem;
    font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  }
  
  .dialog-content {
    padding: 1rem;
    font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  }
}

/* Toggle Switch Styles */
.toggle-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 1rem;
  gap: 0.5rem;
}

.toggle-switch {
  transform: scale(1.2);
}

.toggle-label {
  font-size: 1rem;
  font-weight: 800;
  color: #1b5e20 !important;
  text-shadow: 
    0 2px 4px rgba(0, 0, 0, 0.3),
    0 1px 2px rgba(0, 0, 0, 0.2),
    0 0 8px rgba(27, 94, 32, 0.4) !important;
  letter-spacing: 0.5px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  display: inline-block;
  padding: 2px 8px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 4px;
  backdrop-filter: blur(5px);
}

/* Disabled Card Styles */
.disabled-card {
  opacity: 0.6;
  background: linear-gradient(135deg, #f5f5f5 0%, #e0e0e0 100%) !important;
  cursor: not-allowed !important;
}

.disabled-card:hover {
  transform: none !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1) !important;
}

.disabled-icon {
  opacity: 0.5;
  filter: grayscale(100%);
}

.disabled-text {
  color: #9e9e9e !important;
  opacity: 0.7;
}

/* Enhanced Category Card Styles */
.category-card {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.category-card:hover:not(.disabled-card) {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 12px 30px rgba(0, 0, 0, 0.15);
}

.category-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s ease;
}

.category-card:hover:not(.disabled-card)::before {
  left: 100%;
}

/* Materials & Expenses Card */
.materials-expenses {
  background: linear-gradient(135deg, #ff9800 0%, #f57c00 100%);
  color: white;
}

.materials-expenses .category-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

/* Labor Card */
.labor {
  background: linear-gradient(135deg, #2196f3 0%, #1976d2 100%);
  color: white;
}

.labor .category-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

/* Equipment Card */
.equipment {
  background: linear-gradient(135deg, #4caf50 0%, #388e3c 100%);
  color: white;
}

.equipment .category-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

/* Enhanced Button Styles */
.details-btn {
  font-weight: 600;
  text-transform: none;
  border-radius: 12px;
  padding: 8px 20px;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.details-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
}

.details-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Count Chip Enhancement */
.count-chip {
  font-weight: 700;
  font-size: 1.1rem;
  padding: 8px 16px;
  border-radius: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

/* Category Title Enhancement */
.category-title {
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  line-height: 1.3;
}

.category-description {
  font-size: 1rem;
  opacity: 0.9;
  margin-bottom: 1rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

/* Status Chip Styles */
.status-chip {
  font-weight: 700;
  font-size: 1rem;
  padding: 10px 20px;
  border-radius: 20px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.25);
  transition: all 0.3s ease;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.status-chip.v-chip--color-success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: #ffffff !important;
  box-shadow: 0 4px 16px rgba(16, 185, 129, 0.4) !important;
}

.status-chip.v-chip--color-success:hover {
  background: linear-gradient(135deg, #059669 0%, #047857 100%) !important;
  transform: translateY(-2px);
  box-shadow: 0 6px 24px rgba(16, 185, 129, 0.5) !important;
}

.status-chip:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
}

/* Enhanced Categories Title - تم التنسيق في الأعلى */
</style>