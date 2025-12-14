<template>
  <div>
    <v-container fluid class="pa-0">
        <!-- Hero Section -->
        <div class="page-header glass-effect gradient-animation">
          <div class="header-top-content">
            <h1 class="page-title">مرحباً بك في نظام إدارة المشاريع</h1>
            <span class="page-icon">🏠</span>
          </div>
          <p class="page-subtitle">نظام متكامل لإدارة مشاريعك وفريقك ومواردك المالية</p>
        </div>

        <!-- Quick Actions Grid -->
        <div class="quick-actions-container">
            <div class="quick-actions-grid">
              <v-card
                v-for="action in quickActions"
                :key="action.title"
                class="quick-action-card"
                @click="$router.push(action.route)"
              >
                <v-card-text class="text-center pa-6">
                  <v-icon :color="action.color" size="72" class="mb-4">{{ action.icon }}</v-icon>
                  <h3 class="text-h6 font-weight-bold mb-2 text-primary">{{ action.title }}</h3>
                  <p class="text-body-2 text-secondary">{{ action.description }}</p>
                </v-card-text>
              </v-card>
            </div>
        </div>

        <!-- Statistics Section -->
        <div class="stats-section pa-6">
          <h2 class="text-h4 font-weight-bold text-center mb-6">إحصائيات المشاريع</h2>
          <v-row>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="error">mdi-currency-usd</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-error mb-2">{{ projectStats.totalBudget }}</h3>
                <p class="text-subtitle-1 text-error mb-0">إجمالي الميزانية</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="info">mdi-folder-multiple</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-info mb-2">{{ projectStats.totalProjects }}</h3>
                <p class="text-subtitle-1 text-info mb-0">إجمالي المشاريع</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="warning">mdi-check-circle</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-warning mb-2">{{ projectStats.completedProjects }}</h3>
                <p class="text-subtitle-1 text-warning mb-0">مشاريع مكتملة</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="success">mdi-clock-alert</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-success mb-2">{{ projectStats.activeProjects }}</h3>
                <p class="text-subtitle-1 text-success mb-0">مشاريع نشطة</p>
              </v-card>
            </v-col>
            <v-col cols="12" md="3">
              <v-card class="stat-card pa-6 text-center" elevation="2">
                <div class="stat-icon mb-3">
                  <v-icon size="84" color="primary">mdi-account-hard-hat</v-icon>
                </div>
                <h3 class="text-h3 font-weight-bold text-primary mb-2">{{ projectStats.totalEngineers }}</h3>
                <p class="text-subtitle-1 text-primary mb-0">المهندسين</p>
              </v-card>
            </v-col>
          </v-row>
        </div>
      </v-container>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

// Sample data - in real app, this would come from API
const incomeData = ref(175000) // Total income
const expenseData = ref(103000) // Total expenses

const projectStats = ref({
  totalProjects: 24,
  completedProjects: 18,
  activeProjects: 6,
  totalEngineers: 12
})

// الإجراءات السريعة
const quickActions = ref([
  {
    title: 'إدارة المشاريع',
    description: 'إدارة وتتبع جميع المشاريع والمهام',
    icon: 'mdi-folder-multiple',
    color: 'primary',
    route: '/project-management'
  },
  {
    title: 'المهندسين',
    description: 'إدارة فريق المهندسين والمطورين',
    icon: 'mdi-account-hard-hat',
    color: 'success',
    route: '/engineers'
  },
  {
    title: 'المصاريف الإدارية',
    description: 'إدارة وتتبع المصاريف الإدارية',
    icon: 'mdi-currency-usd',
    color: 'error',
    route: '/expenses'
  },
  {
    title: 'التصنيفات',
    description: 'إدارة وتنظيم تصنيفات المشاريع والمهام',
    icon: 'mdi-tag-multiple',
    color: 'info',
    route: '/categories'
  },
  {
    title: 'المديونون',
    description: 'إدارة حسابات المديونون والمستحقات',
    icon: 'mdi-credit-card',
    color: 'error',
    route: '/debtors'
  },
  {
    title: 'المخزون',
    description: 'إدارة وتتبع جميع عناصر المخزون والمواد',
    icon: 'mdi-package-variant',
    color: 'purple',
    route: '/inventory'
  },
  {
    title: 'المبيعات',
    description: 'إدارة وتتبع جميع عمليات البيع والمبيعات',
    icon: 'mdi-cash-multiple',
    color: 'success',
    route: '/sales'
  },
  {
    title: 'المشتريات',
    description: 'إدارة وتتبع جميع عمليات الشراء والمشتريات',
    icon: 'mdi-cart',
    color: 'error',
    route: '/purchases'
  },
  {
    title: 'الموارد البشرية',
    description: 'إدارة وتتبع جميع الموظفين والموارد البشرية',
    icon: 'mdi-account-group',
    color: 'info',
    route: '/human-resources'
  }
])


// Computed properties
const totalIncome = computed(() => incomeData.value)
const totalExpenses = computed(() => expenseData.value)
const netProfit = computed(() => totalIncome.value - totalExpenses.value)

// Methods
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('ar-SA', {
    style: 'currency',
    currency: 'IQD'
  }).format(amount)
}
</script>

<style scoped>
/* ========================================
   CSS محسن ومنظم - تم نقله إلى rtl.css
   ======================================== */


/* رأس الصفحة المحسن - متناسق مع الصفحات الأخرى */
.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #667eea 100%) !important;
  backdrop-filter: blur(20px);
  border-radius: 16px;
  padding: 20px 30px;
  margin-bottom: 20px;
  text-align: center !important;
  border: 2px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
  animation: gradient-animation 8s ease infinite;
  background-size: 200% 200%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: white;
}

.header-top-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 10px;
  width: 100%;
  text-align: center;
}

.page-icon {
  font-size: 2rem;
  display: inline-block;
  animation: iconFloat 3s ease-in-out infinite;
}

@keyframes iconFloat {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes gradient-animation {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.gradient-animation {
  background-size: 200% 200%;
  animation: gradient-animation 3s ease infinite;
}

.page-title {
  font-size: 2rem;
  font-weight: 800;
  color: rgb(246, 246, 246);
  margin: 0;
  padding: 0;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
  text-align: center !important;
  width: auto;
  line-height: 1.2;
  letter-spacing: 0.5px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.page-subtitle {
  font-size: 1rem;
  color: rgba(243, 240, 240, 0.9);
  margin: 0;
  padding: 0;
  text-align: center !important;
  width: 100%;
  line-height: 1.4;
  letter-spacing: 0.3px;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
  font-weight: 500;
}

.quick-actions-container {
  padding: 0 2rem;
  margin-bottom: 2rem;
}

/* تأثيرات متحركة للخلفية */
.hero-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    radial-gradient(circle at 20% 80%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
    radial-gradient(circle at 80% 20%, rgba(255, 255, 255, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 40% 40%, rgba(59, 130, 246, 0.2) 0%, transparent 50%);
  animation: backgroundShift 20s ease-in-out infinite;
}

/* تأثيرات الجسيمات المتحركة المحسنة */
.hero-section::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    radial-gradient(3px 3px at 20px 30px, rgba(255, 255, 255, 0.4), transparent),
    radial-gradient(2px 2px at 40px 70px, rgba(255, 255, 255, 0.3), transparent),
    radial-gradient(1px 1px at 90px 40px, rgba(255, 255, 255, 0.5), transparent),
    radial-gradient(1px 1px at 130px 80px, rgba(255, 255, 255, 0.4), transparent),
    radial-gradient(2px 2px at 160px 30px, rgba(255, 255, 255, 0.3), transparent),
    radial-gradient(1px 1px at 200px 50px, rgba(59, 130, 246, 0.3), transparent),
    radial-gradient(2px 2px at 250px 20px, rgba(147, 51, 234, 0.2), transparent);
  background-repeat: repeat;
  background-size: 300px 150px;
  animation: particleFloat 20s linear infinite;
}

/* تأثيرات إضافية للخلفية */
.hero-section {
  position: relative;
}

.hero-section .floating-shapes {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
  pointer-events: none;
}

.hero-section .floating-shapes::before,
.hero-section .floating-shapes::after {
  content: '';
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  animation: floatUp 15s infinite linear;
}

.hero-section .floating-shapes::before {
  width: 80px;
  height: 80px;
  top: 20%;
  left: 10%;
  animation-delay: 0s;
}

.hero-section .floating-shapes::after {
  width: 120px;
  height: 120px;
  top: 60%;
  right: 15%;
  animation-delay: 7s;
}

@keyframes floatUp {
  0% {
    transform: translateY(100vh) rotate(0deg);
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    transform: translateY(-100vh) rotate(360deg);
    opacity: 0;
  }
}

/* حركة الخلفية */
@keyframes backgroundShift {
  0%, 100% { transform: translateX(0) translateY(0); }
  25% { transform: translateX(-10px) translateY(-5px); }
  50% { transform: translateX(5px) translateY(-10px); }
  75% { transform: translateX(-5px) translateY(5px); }
}

/* حركة الجسيمات */
@keyframes particleFloat {
  0% { transform: translateY(0) translateX(0); }
  100% { transform: translateY(-100px) translateX(50px); }
}

.hero-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grain" width="100" height="100" patternUnits="userSpaceOnUse"><circle cx="25" cy="25" r="1" fill="rgba(255,255,255,0.1)"/><circle cx="75" cy="75" r="1" fill="rgba(255,255,255,0.1)"/><circle cx="50" cy="10" r="0.5" fill="rgba(255,255,255,0.05)"/><circle cx="10" cy="60" r="0.5" fill="rgba(255,255,255,0.05)"/><circle cx="90" cy="40" r="0.5" fill="rgba(255,255,255,0.05)"/></pattern></defs><rect width="100" height="100" fill="url(%23grain)"/></svg>');
  opacity: 0.3;
}

.hero-content {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
  text-align: center;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 700;
  color: #ffffff;
  text-shadow: 0 4px 12px rgba(0,0,0,0.4);
  margin-bottom: 1rem;
  line-height: 1.2;
  letter-spacing: 0.5px;
  position: relative;
  animation: titleGlow 3s ease-in-out infinite;
}


/* تأثير التوهج المحسن للعنوان */
@keyframes titleGlow {
  0%, 100% { 
    text-shadow: 
      0 4px 12px rgba(0,0,0,0.4),
      0 0 20px rgba(255, 255, 255, 0.1),
      0 0 40px rgba(59, 130, 246, 0.1);
  }
  50% { 
    text-shadow: 
      0 4px 12px rgba(0,0,0,0.4),
      0 0 40px rgba(255, 255, 255, 0.3),
      0 0 60px rgba(59, 130, 246, 0.2),
      0 0 80px rgba(147, 51, 234, 0.15);
  }
}

/* تأثير النبض للعنوان الفرعي */
.hero-subtitle {
  animation: subtitlePulse 4s ease-in-out infinite;
}

@keyframes subtitlePulse {
  0%, 100% { 
    opacity: 1;
    transform: translateY(0);
    text-shadow: 0 2px 12px rgba(0,0,0,0.6), 0 1px 4px rgba(0,0,0,0.8);
  }
  50% { 
    opacity: 1;
    transform: translateY(-2px);
    text-shadow: 0 2px 12px rgba(0,0,0,0.6), 0 1px 4px rgba(0,0,0,0.8), 0 0 20px rgba(255,255,255,0.3);
  }
}

.hero-subtitle {
  font-size: 1.25rem;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 3rem;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
  text-shadow: 0 2px 8px rgba(0,0,0,0.3);
}

/* شبكة الإجراءات السريعة - صف واحد */
.quick-actions-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.quick-action-card {
  border-radius: 24px;
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(25px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 
    0 12px 40px rgba(0, 0, 0, 0.08),
    0 0 0 1px rgba(255, 255, 255, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
  position: relative;
  overflow: hidden;
  min-height: 220px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center !important;
}

/* تأثير الإضاءة المتحركة */
.quick-action-card::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(
    from 0deg,
    transparent 0deg,
    rgba(59, 130, 246, 0.1) 60deg,
    rgba(147, 51, 234, 0.1) 120deg,
    rgba(59, 130, 246, 0.1) 180deg,
    transparent 240deg,
    transparent 360deg
  );
  animation: rotateGlow 4s linear infinite;
  opacity: 0;
  transition: opacity 0.3s ease;
}

@keyframes rotateGlow {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* تأثير الإضاءة الداخلية */
.quick-action-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.1) 0%, transparent 50%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

/* تأثير الحدود المتحركة */
.quick-action-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 20px;
  padding: 2px;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.3), rgba(147, 51, 234, 0.3), rgba(59, 130, 246, 0.3));
  mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  mask-composite: exclude;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.quick-action-card:hover {
  transform: translateY(-16px) scale(1.05);
  box-shadow: 
    0 32px 64px rgba(0, 0, 0, 0.12),
    0 0 0 1px rgba(255, 255, 255, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.3),
    0 0 40px rgba(59, 130, 246, 0.1);
  background: rgba(255, 255, 255, 1);
}

.quick-action-card:hover::before {
  opacity: 1;
  animation: rotateGlow 2s linear infinite;
}

.quick-action-card:hover::after {
  opacity: 1;
  animation: borderGlow 1.5s ease-in-out infinite;
}

/* تأثير النبض للبطاقات */
.quick-action-card:active {
  transform: translateY(-8px) scale(0.98);
  transition: all 0.1s ease;
}

/* تأثير الحدود المتوهجة */
@keyframes borderGlow {
  0%, 100% { 
    background: linear-gradient(135deg, rgba(59, 130, 246, 0.3), rgba(147, 51, 234, 0.3), rgba(59, 130, 246, 0.3));
  }
  50% { 
    background: linear-gradient(135deg, rgba(59, 130, 246, 0.6), rgba(147, 51, 234, 0.6), rgba(59, 130, 246, 0.6));
  }
}

/* قسم الإحصائيات */
.stats-section {
  background: #f8f9fa;
}

.stat-card {
  border-radius: 16px;
  transition: all 0.3s ease;
  height: 100%;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(0,0,0,0.1);
}

/* تأثيرات عامة */
.v-card {
  border-radius: 16px;
  transition: all 0.3s ease;
}

.v-card:hover {
  box-shadow: 0 8px 25px rgba(0,0,0,0.15);
}

.text-primary {
  color: #1976d2 !important;
}

/* تصميم متجاوب */
@media (max-width: 768px) {
  .hero-title {
    font-size: 2.5rem;
  }
  
  .hero-subtitle {
    font-size: 1.1rem;
  }
  
  .quick-actions-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }
}

/* للشاشات الصغيرة جداً */
@media (max-width: 600px) {
  .quick-actions-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }
  
  .feature-grid {
    grid-template-columns: 1fr;
  }
}

/* تأثيرات الحركة */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.hero-content > * {
  animation: fadeInUp 0.6s ease-out;
}

.hero-content > *:nth-child(2) {
  animation-delay: 0.2s;
}

.hero-content > *:nth-child(3) {
  animation-delay: 0.4s;
}

/* دعم RTL شامل */
.hero-section {
  direction: rtl;
  text-align: right;
}

.hero-title {
  text-align: right;
  direction: rtl;
}

.hero-subtitle {
  text-align: right;
  direction: rtl;
}

.quick-actions-grid {
  direction: rtl;
}

.quick-action-card {
  direction: rtl;
  text-align: right;
}

.quick-action-card .v-card-text {
  text-align: center;
  position: relative;
  z-index: 2;
}

/* تأثيرات الأيقونات المحسنة */
.quick-action-card .v-icon {
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  filter: drop-shadow(0 8px 20px rgba(0, 0, 0, 0.15));
  position: relative;
  font-size: 72px !important;
  width: 72px !important;
  height: 72px !important;
}

.quick-action-card:hover .v-icon {
  transform: scale(1.08) rotate(5deg);
  filter: drop-shadow(0 20px 40px rgba(0, 0, 0, 0.25));
  animation: iconPulse 2s ease-in-out infinite;
}

/* تأثير النبض للأيقونات */
@keyframes iconPulse {
  0%, 100% { 
    transform: scale(1.08) rotate(5deg);
    filter: drop-shadow(0 20px 40px rgba(0, 0, 0, 0.25));
  }
  50% { 
    transform: scale(1.12) rotate(5deg);
    filter: drop-shadow(0 24px 48px rgba(0, 0, 0, 0.35));
  }
}

/* تأثير الإضاءة للأيقونات */
.quick-action-card:hover .v-icon::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 130%;
  height: 130%;
  background: radial-gradient(circle, rgba(59, 130, 246, 0.3) 0%, transparent 70%);
  border-radius: 50%;
  transform: translate(-50%, -50%);
  animation: iconGlow 2s ease-in-out infinite;
}

@keyframes iconGlow {
  0%, 100% { 
    opacity: 0.5;
    transform: translate(-50%, -50%) scale(1);
  }
  50% { 
    opacity: 0.8;
    transform: translate(-50%, -50%) scale(1.4);
  }
}

/* تأثيرات النصوص */
.quick-action-card h3 {
  transition: all 0.3s ease;
  font-weight: 700;
  letter-spacing: 0.5px;
  font-size: 1.1rem;
  margin-bottom: 8px;
  text-align: center !important;
  width: 100%;
  color: #1976d2 !important;
}

.quick-action-card:hover h3 {
  color: #1565c0 !important;
  transform: translateY(-2px);
  font-size: 1.15rem;
}

.quick-action-card p {
  transition: all 0.3s ease;
  opacity: 0.8;
  font-size: 0.9rem;
  line-height: 1.4;
  text-align: center !important;
  width: 100%;
  color: #757575 !important;
}

.quick-action-card:hover p {
  opacity: 1;
  transform: translateY(-1px);
  font-size: 0.95rem;
  color: #616161 !important;
}

.stats-section {
  direction: rtl;
  text-align: right;
}

/* تنسيق خاص للأرقام الإنجليزية في الإحصائيات */
.stat-card h3 {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  font-variant-numeric: tabular-nums;
  direction: ltr !important;
  text-align: center !important;
}

.stat-card p {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: rtl !important;
  text-align: center !important;
}

/* تنسيق خاص لبطاقة السعر بالدينار العراقي */
.currency-amount {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  font-variant-numeric: tabular-nums !important;
  direction: ltr !important;
  text-align: center !important;
  font-size: 3.5rem !important;
  font-weight: 900 !important;
  letter-spacing: 2px !important;
  color: #d32f2f !important;
  line-height: 1 !important;
  margin-bottom: 0.5rem !important;
  text-shadow: 0 4px 8px rgba(211, 47, 47, 0.5) !important;
  border: 3px solid #d32f2f !important;
  padding: 10px !important;
  border-radius: 10px !important;
  background: linear-gradient(135deg, #fff3e0 0%, #ffebee 100%) !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 8px 16px rgba(211, 47, 47, 0.3) !important;
}

.currency-amount:hover {
  transform: scale(1.1) !important;
  text-shadow: 0 6px 12px rgba(211, 47, 47, 0.7) !important;
  box-shadow: 0 12px 24px rgba(211, 47, 47, 0.5) !important;
}

.currency-label {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  direction: rtl !important;
  text-align: center !important;
  font-size: 1.5rem !important;
  font-weight: 700 !important;
  color: #d32f2f !important;
  margin-bottom: 0 !important;
  line-height: 1.4 !important;
  letter-spacing: 1px !important;
  text-shadow: 0 2px 4px rgba(211, 47, 47, 0.4) !important;
  border: 2px solid #d32f2f !important;
  padding: 8px !important;
  border-radius: 8px !important;
  background: linear-gradient(135deg, #fff3e0 0%, #ffebee 100%) !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 4px 8px rgba(211, 47, 47, 0.2) !important;
}

.currency-label:hover {
  transform: scale(1.05) !important;
  text-shadow: 0 3px 6px rgba(211, 47, 47, 0.6) !important;
  box-shadow: 0 6px 12px rgba(211, 47, 47, 0.4) !important;
}

/* تنسيق أقوى للتأكد من التطبيق */
.stat-card .currency-amount {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  font-size: 3.5rem !important;
  font-weight: 900 !important;
  color: #d32f2f !important;
  direction: ltr !important;
  text-align: center !important;
  letter-spacing: 2px !important;
  text-shadow: 0 4px 8px rgba(211, 47, 47, 0.5) !important;
  border: 3px solid #d32f2f !important;
  padding: 10px !important;
  border-radius: 10px !important;
  background: linear-gradient(135deg, #fff3e0 0%, #ffebee 100%) !important;
  box-shadow: 0 8px 16px rgba(211, 47, 47, 0.3) !important;
}

.stat-card .currency-label {
  font-family: 'Arial', 'Helvetica', sans-serif !important;
  font-size: 1.5rem !important;
  font-weight: 700 !important;
  color: #d32f2f !important;
  direction: rtl !important;
  text-align: center !important;
  letter-spacing: 1px !important;
  text-shadow: 0 2px 4px rgba(211, 47, 47, 0.4) !important;
  border: 2px solid #d32f2f !important;
  padding: 8px !important;
  border-radius: 8px !important;
  background: linear-gradient(135deg, #fff3e0 0%, #ffebee 100%) !important;
  box-shadow: 0 4px 8px rgba(211, 47, 47, 0.2) !important;
}

.stat-card {
  direction: rtl;
  text-align: center;
}

/* إصلاح اتجاه النصوص */
h1, h2, h3, h4, h5, h6, p, span, div {
  text-align: right;
  direction: rtl;
}

/* إصلاح اتجاه الأيقونات */
.v-icon {
  transform: scaleX(-1);
}

/* إصلاح اتجاه الأزرار */
.v-btn {
  direction: rtl;
}

/* إصلاح اتجاه القوائم */
.v-list {
  direction: rtl;
  text-align: right;
}

/* إصلاح اتجاه الكروت */
.v-card {
  direction: rtl;
  text-align: right;
}

.v-card-title {
  text-align: right;
  direction: rtl;
}

.v-card-text {
  text-align: right;
  direction: rtl;
}

/* بطاقات الإحصائيات المحسنة */
.stat-card {
  background: var(--bg-white) !important;
  border-radius: 16px !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  min-height: 200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center !important;
}

.stat-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15) !important;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: var(--gradient-primary);
  border-radius: 16px 16px 0 0;
}

.stat-icon {
  position: relative;
  z-index: 2;
}

.stat-card .v-icon {
  font-size: 84px !important;
  width: 84px !important;
  height: 84px !important;
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.1));
  transition: all 0.3s ease;
}

.stat-card:hover .v-icon {
  transform: scale(1.1);
  filter: drop-shadow(0 8px 16px rgba(0, 0, 0, 0.2));
}

.stat-card h3 {
  font-size: 2.5rem !important;
  font-weight: 700 !important;
  margin-bottom: 8px !important;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center !important;
  width: 100%;
}

.stat-card p {
  font-size: 1rem !important;
  font-weight: 500 !important;
  opacity: 0.8;
  text-align: center !important;
  width: 100%;
}

/* تأثيرات خاصة لكل لون */
.stat-card:nth-child(1)::before {
  background: var(--gradient-primary);
}

.stat-card:nth-child(2)::before {
  background: var(--gradient-success);
}

.stat-card:nth-child(3)::before {
  background: var(--gradient-warning);
}

.stat-card:nth-child(4)::before {
  background: var(--gradient-info);
}

/* التصميم المتجاوب للإحصائيات */
@media (max-width: 960px) {
  .stat-card {
    min-height: 200px;
  }
  
  .stat-card h3 {
    font-size: 2.2rem !important;
  }
  
  .stat-card .v-icon {
    font-size: 76px !important;
    width: 76px !important;
    height: 76px !important;
  }
}

@media (max-width: 600px) {
  .stat-card {
    min-height: 180px;
    padding: 20px !important;
  }
  
  .stat-card h3 {
    font-size: 2rem !important;
  }
  
  .stat-card .v-icon {
    font-size: 68px !important;
    width: 68px !important;
    height: 68px !important;
  }
}
</style>
