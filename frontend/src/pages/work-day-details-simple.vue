<template>
  <div class="simple-test-page">
    <h1>اختبار صفحة تفاصيل يوم العمل</h1>
    <p>هذه صفحة اختبار بسيطة</p>
    
    <v-card class="test-card">
      <v-card-title>معلومات يوم العمل</v-card-title>
      <v-card-text>
        <p><strong>مكان العمل:</strong> موقع البناء الرئيسي</p>
        <p><strong>رقم الاستمارة:</strong> FORM-001</p>
        <p><strong>فترة العمل:</strong> 08:00 - 16:00</p>
        <p><strong>اليوم:</strong> الاثنين</p>
        <p><strong>نوع العمل:</strong> بناء</p>
        <p><strong>التاريخ:</strong> 2024-01-15</p>
      </v-card-text>
    </v-card>

    <!-- Categories -->
    <div class="categories-grid">
      <v-card class="category-card materials-card">
        <v-card-text>
          <h3>المواد</h3>
          <p>مواد البناء والإنشاء</p>
          <div class="count">5 مادة</div>
        </v-card-text>
      </v-card>

      <v-card class="category-card expenses-card">
        <v-card-text>
          <h3>المصاريف اليومية</h3>
          <p>النفقات والمصاريف</p>
          <div class="count">3 مصروف</div>
        </v-card-text>
      </v-card>

      <v-card class="category-card labor-card">
        <v-card-text>
          <h3>الأيدي العاملة</h3>
          <p>العمال والموظفين</p>
          <div class="count">8 عامل</div>
        </v-card-text>
      </v-card>

      <v-card class="category-card equipment-card">
        <v-card-text>
          <h3>الآلات</h3>
          <p>المعدات والآلات</p>
          <div class="count">4 آلة</div>
        </v-card-text>
      </v-card>
    </div>

    <v-btn @click="goBack" color="primary">العودة</v-btn>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  router.push('/work-days')
}
</script>

<style scoped>
.simple-test-page {
  padding: 2rem;
  background: #f5f5f5;
  min-height: 100vh;
  direction: rtl;
}

.test-card {
  margin-bottom: 2rem;
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.category-card {
  border-radius: 12px;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.category-card:hover {
  transform: translateY(-5px);
}

.materials-card {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
}

.expenses-card {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  color: white;
}

.labor-card {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
}

.equipment-card {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
  color: white;
}

.category-card h3 {
  color: white;
  margin-bottom: 0.5rem;
}

.category-card p {
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 1rem;
}

.count {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-weight: 600;
  text-align: center;
}
</style>
