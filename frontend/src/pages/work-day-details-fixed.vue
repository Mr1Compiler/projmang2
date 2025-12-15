<template>
  <div class="work-details-page">
    <!-- Header -->
    <div class="page-header">
      <h1>تفاصيل يوم العمل</h1>
      <button @click="goBack" class="back-btn">العودة</button>
    </div>

    <!-- Debug Info -->
    <div class="debug-info">
      <h2 style="color: green;">✅ الصفحة تعمل بشكل صحيح</h2>
      <p>ID: {{ workDayId }}</p>
      <p>الوقت الحالي: {{ currentTime }}</p>
    </div>

    <!-- Work Day Info -->
    <div class="info-section">
      <h3>معلومات يوم العمل</h3>
      <div class="info-grid">
        <div class="info-item">
          <label>مكان العمل:</label>
          <span>موقع البناء الرئيسي</span>
        </div>
        <div class="info-item">
          <label>رقم الاستمارة:</label>
          <span>FORM-001</span>
        </div>
        <div class="info-item">
          <label>فترة العمل:</label>
          <span>08:00 - 16:00</span>
        </div>
        <div class="info-item">
          <label>اليوم:</label>
          <span>الاثنين</span>
        </div>
        <div class="info-item">
          <label>نوع العمل:</label>
          <span>بناء</span>
        </div>
        <div class="info-item">
          <label>التاريخ:</label>
          <span>2024-01-15</span>
        </div>
      </div>
    </div>

    <!-- Categories -->
    <div class="categories-section">
      <h3>الأقسام</h3>
      <div class="categories-grid">
        <!-- Materials -->
        <div class="category-card materials">
          <div class="category-icon">🔨</div>
          <h4>المواد</h4>
          <p>مواد البناء والإنشاء</p>
          <div class="count">5 مادة</div>
        </div>

        <!-- Expenses -->
        <div class="category-card expenses">
          <div class="category-icon">💰</div>
          <h4>المصاريف اليومية</h4>
          <p>النفقات والمصاريف</p>
          <div class="count">3 مصروف</div>
        </div>

        <!-- Labor -->
        <div class="category-card labor">
          <div class="category-icon">👷</div>
          <h4>الأيدي العاملة</h4>
          <p>العمال والموظفين</p>
          <div class="count">8 عامل</div>
        </div>

        <!-- Equipment -->
        <div class="category-card equipment">
          <div class="category-icon">🚜</div>
          <h4>الآلات</h4>
          <p>المعدات والآلات</p>
          <div class="count">4 آلة</div>
        </div>
      </div>
    </div>

    <!-- Actions -->
    <div class="actions-section">
      <button class="action-btn edit-btn">تعديل</button>
      <button class="action-btn delete-btn">حذف</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const workDayId = ref('')
const currentTime = ref('')

onMounted(() => {
  workDayId.value = route.params.id || 'test-123'
  currentTime.value = new Date().toLocaleString('ar-SA')
})

const goBack = () => {
  router.push('/work-days')
}
</script>


<style>
/* Import page styles */
@import './styles/work-day-details-fixed.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
