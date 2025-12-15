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


<style>
/* Import page styles */
@import './styles/work-day-details-simple.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
