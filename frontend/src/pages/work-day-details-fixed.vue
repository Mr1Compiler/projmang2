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
  
  console.log('Page loaded successfully!')
  console.log('Work Day ID:', workDayId.value)
})

const goBack = () => {
  router.push('/work-days')
}
</script>

<style scoped>
.work-details-page {
  min-height: 100vh;
  background: #f5f5f5;
  padding: 2rem;
  direction: rtl;
  font-family: Arial, sans-serif;
}

.page-header {
  background: #1e3a8a;
  color: white;
  padding: 2rem;
  border-radius: 12px;
  margin-bottom: 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-header h1 {
  margin: 0;
  font-size: 2rem;
}

.back-btn {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

.debug-info {
  background: white;
  padding: 1.5rem;
  border-radius: 12px;
  margin-bottom: 2rem;
  border-left: 4px solid #10b981;
}

.debug-info h2 {
  margin: 0 0 1rem 0;
}

.info-section {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  margin-bottom: 2rem;
}

.info-section h3 {
  margin: 0 0 1.5rem 0;
  color: #1e3a8a;
  font-size: 1.5rem;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
}

.info-item label {
  font-weight: bold;
  color: #333;
}

.info-item span {
  color: #666;
}

.categories-section {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  margin-bottom: 2rem;
}

.categories-section h3 {
  margin: 0 0 1.5rem 0;
  color: #1e3a8a;
  font-size: 1.5rem;
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}

.category-card {
  padding: 2rem;
  border-radius: 16px;
  text-align: center;
  color: white;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.category-card:hover {
  transform: translateY(-8px);
}

.category-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.category-card h4 {
  margin: 0 0 0.5rem 0;
  font-size: 1.4rem;
}

.category-card p {
  margin: 0 0 1rem 0;
  opacity: 0.9;
}

.count {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-weight: bold;
}

.materials {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
}

.expenses {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
}

.labor {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.equipment {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
}

.actions-section {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  text-align: center;
}

.action-btn {
  padding: 0.75rem 2rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: bold;
  margin: 0 0.5rem;
}

.edit-btn {
  background: #3b82f6;
  color: white;
}

.delete-btn {
  background: #ef4444;
  color: white;
}

.action-btn:hover {
  opacity: 0.8;
}

@media (max-width: 768px) {
  .work-details-page {
    padding: 1rem;
  }
  
  .page-header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
  
  .categories-grid {
    grid-template-columns: 1fr;
  }
  
  .info-grid {
    grid-template-columns: 1fr;
  }
}
</style>
