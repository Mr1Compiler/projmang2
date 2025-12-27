<template>
  <v-app dir="rtl" class="arabic-app">
    <!-- الشريط الجانبي الأيمن -->
    <v-navigation-drawer
      v-if="!isLoginPage"
      v-model="drawer"
      location="right"
      :temporary="$vuetify.display.mobile"
      width="320"
      class="modern-sidebar rtl-sidebar"
    >
      <!-- الشعار والعنوان -->
      <div class="sidebar-header pa-6">
        <div class="text-center mb-6 position-relative">
          <!-- زر إغلاق الـ sidebar (يظهر على الشاشات الصغيرة أو عند الحاجة) -->
          <v-btn
            v-if="$vuetify.display.mobile"
            icon
            variant="text"
            size="small"
            @click="drawer = false"
            class="sidebar-close-btn"
          >
            <v-icon color="white">mdi-close</v-icon>
          </v-btn>
          <v-img
            class="mb-4 mx-auto logo-enhanced"
            height="80"
            width="80"
            src="@/assets/logo.png"
          />
          <h1 class="text-h5 font-weight-bold text-white mb-2">نظام إدارة المشاريع</h1>
          <p class="text-body-2 text-white text-opacity-80">نظام شامل لإدارة المشاريع والمهندسين والمصاريف</p>
        </div>
      </div>

      <!-- قائمة التنقل الرئيسية -->
      <v-list class="pa-2" nav>
        <v-list-item
          v-for="item in mainMenuItems"
          :key="item.title"
          :to="item.to"
          class="menu-item"
          rounded="xl"
          :class="{ 'active-menu-item': item.active }"
        >
          <template v-slot:prepend>
            <v-icon :color="item.active ? 'primary' : 'white'" class="me-3">
              {{ item.icon }}
            </v-icon>
          </template>
          <v-list-item-title class="text-body-1 font-weight-medium text-white">
            {{ item.title }}
          </v-list-item-title>
          <template v-slot:append v-if="item.badge">
            <v-chip
              :color="item.badgeColor"
              size="small"
              class="text-white"
            >
              {{ item.badge }}
            </v-chip>
          </template>
        </v-list-item>
      </v-list>

      <!-- زر تسجيل الخروج -->
      <v-list class="pa-2" nav>
        <v-list-item
          @click="handleLogout"
          class="menu-item logout-menu-item"
          rounded="xl"
        >
          <template v-slot:prepend>
            <v-icon color="white" class="me-3">mdi-logout</v-icon>
          </template>
          <v-list-item-title class="text-body-1 font-weight-medium text-white">
            تسجيل الخروج
          </v-list-item-title>
        </v-list-item>
      </v-list>

      <!-- Footer -->
      <template v-slot:append>
        <div class="sidebar-footer">
          <!-- معلومات المستخدم -->
          <div v-if="currentUsername" class="user-info-card">
            <div class="user-avatar">
              <v-icon color="white" size="24">mdi-account-circle</v-icon>
            </div>
            <div class="user-details">
              <p class="user-greeting">مرحباً</p>
              <p class="user-name">{{ currentUsername }}</p>
            </div>
          </div>
        </div>
      </template>
    </v-navigation-drawer>

    <!-- المحتوى الرئيسي -->
    <v-main>
      <!-- زر إظهار/إخفاء الـ Sidebar المحسن -->
      <v-btn
        v-if="!isLoginPage"
        icon
        @click="drawer = !drawer"
        :class="['sidebar-toggle-btn-fixed', { 'sidebar-open': drawer }]"
        color="primary"
        size="large"
        elevation="4"
      >
        <v-icon class="toggle-icon">{{ drawer ? 'mdi-menu-open' : 'mdi-menu' }}</v-icon>
        <div class="toggle-ripple"></div>
      </v-btn>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { handleLogout as logoutUser, getCurrentUser } from '@/services/authService'
import { useAuthStore } from '@/stores/auth'

// Note: authStore is still used for logout and permissions loading

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const drawer = ref(!authStore.isMobile) // Close by default on mobile

// Load permissions from storage on mount
onMounted(() => {
  authStore.loadFromStorage()
})

// اسم المستخدم الحالي
const currentUsername = computed(() => {
  const user = getCurrentUser()
  return user?.username || localStorage.getItem('username') || ''
})

// التحقق من صفحة تسجيل الدخول
const isLoginPage = computed(() => {
  return route.path === '/login'
})

// دالة تسجيل الخروج
const handleLogout = async () => {
  authStore.clearAuth()
  await logoutUser()
  router.push('/login')
}

// قائمة التنقل المبنية ديناميكياً من صفحات المستخدم
// الصفحات التي لها مسارات Vue فقط ستظهر في السايدبار
const mainMenuItems = computed(() => {
  return authStore.pages
    .filter(page => {
      // فقط الصفحات التي لها مسارات Vue (تبدأ بـ / وليست API routes)
      const apiOnlyRoutes = [
        '/workdays',
        '/workdayLabor',
        '/workdayEquipment',
        '/workdayMaterials',
        '/rolePages',
        '/userRoles',
        '/teamMembers',
        '/workSubcategories'
      ]
      return !apiOnlyRoutes.includes(page.route)
    })
    .map(page => ({
      title: page.name,
      icon: page.icon || 'mdi-file-document',
      to: page.route,
      active: route.path === page.route
    }))
})

// Active state is now computed in mainMenuItems
</script>

<style>
/* دعم RTL للنظام العربي */
.arabic-app {
  direction: rtl;
  text-align: right;
  font-family: 'Cairo', 'Tajawal', 'Noto Sans Arabic', 'Arial', sans-serif;
}

/* ========================================
   السايد بار العصري
   ======================================== */
.modern-sidebar {
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.7) 0%, rgba(21, 101, 192, 0.7) 100%) !important;
  border-left: none !important;
  box-shadow: -4px 0 20px rgba(25, 118, 210, 0.2) !important;
  z-index: 5 !important;
  backdrop-filter: blur(10px);
}

.rtl-sidebar {
  direction: rtl;
  text-align: right;
}

.sidebar-header {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 0 0 20px 20px;
  margin: -16px -16px 16px -16px;
}

.logo-enhanced {
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  padding: 8px;
}

/* عناصر القائمة */
.menu-item {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border-radius: 8px !important;
  margin-bottom: 0 !important;
  margin-top: 0 !important;
}

.menu-item .v-list-item-title {
  font-size: 0.65rem !important;
}

.modern-sidebar .v-list-item-title {
  font-size: 0.65rem !important;
}

.modern-sidebar p,
.modern-sidebar span {
  font-size: 0.6rem !important;
}

.modern-sidebar .text-h5 {
  font-size: 0.75rem !important;
}

.modern-sidebar .text-body-2 {
  font-size: 0.55rem !important;
}

/* تكبير نص رأس السايد بار */
.sidebar-header .text-h5 {
  font-size: 0.8rem !important;
  text-align: center !important;
}

.sidebar-header .text-body-2 {
  font-size: 0.6rem !important;
  text-align: center !important;
}

.sidebar-header .text-center {
  text-align: center !important;
  display: flex !important;
  flex-direction: column !important;
  align-items: center !important;
  justify-content: center !important;
}

.sidebar-header h1,
.sidebar-header p {
  text-align: center !important;
  width: 100% !important;
}

.modern-sidebar .text-caption {
  font-size: 0.5rem !important;
}

.modern-sidebar h1,
.modern-sidebar h2,
.modern-sidebar h3 {
  font-size: 0.65rem !important;
}

/* تكبير عنوان "المميزات الرئيسية" */
.modern-sidebar h3.text-h6 {
  font-size: 0.7rem !important;
  font-weight: 700 !important;
}

/* توحيد المسافات بين العناصر */
.modern-sidebar .v-list {
  padding: 4px !important;
  gap: 0 !important;
}

.modern-sidebar .v-list-item {
  padding: 2px 6px !important;
  min-height: 28px !important;
  margin-bottom: 0 !important;
  margin-top: 0 !important;
}

.modern-sidebar .v-list-item.menu-item {
  margin-bottom: 0 !important;
  margin-top: 0 !important;
}

.modern-sidebar .v-list-item.menu-item:last-child {
  margin-bottom: 0 !important;
}

.sidebar-header {
  padding: 8px !important;
  margin-bottom: 4px !important;
}

.sidebar-header .mb-6 {
  margin-bottom: 4px !important;
}

.feature-grid {
  gap: 6px !important;
}

.feature-item {
  padding: 6px 4px !important;
}

.menu-item:hover {
  background: rgba(255, 255, 255, 0.2) !important;
  transform: translateX(-4px);
}

.active-menu-item {
  background: linear-gradient(135deg, #1565c0 0%, #1976d2 100%) !important;
  box-shadow: 0 4px 15px rgba(255, 255, 255, 0.4);
  transform: translateX(-4px);
}

.active-menu-item .v-list-item-title {
  font-weight: 600 !important;
}

/* شبكة المميزات */
.feature-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.feature-item {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  padding: 12px 8px;
  text-align: center;
  transition: all 0.3s ease;
  backdrop-filter: blur(5px);
}

.feature-item:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-2px);
}

/* رابط إدارة الفريق */
.feature-link {
  text-decoration: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: rgba(255, 255, 255, 0.2) !important;
  border: 1px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
}

.feature-link:hover {
  background: rgba(255, 255, 255, 0.4) !important;
  border-color: rgba(255, 255, 255, 0.5);
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 8px 25px rgba(255, 255, 255, 0.3);
}

.feature-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.feature-link:hover::before {
  left: 100%;
}

/* إصلاح اتجاه العناصر */
.v-list-item {
  text-align: right;
}

.v-card-title {
  text-align: right;
}

.v-card-text {
  text-align: right;
}

/* إصلاح اتجاه الأيقونات */
.v-icon {
  transform: scaleX(-1);
}

/* إصلاح اتجاه الأزرار */
.v-btn {
  direction: rtl;
}

/* إصلاح اتجاه النصوص */
h1, h2, h3, h4, h5, h6, p, span, div {
  text-align: right;
}

/* إصلاح اتجاه الجداول */
.v-data-table {
  direction: rtl;
}

.v-data-table th,
.v-data-table td {
  text-align: right;
}

/* إصلاح اتجاه النماذج */
.v-text-field input {
  text-align: right;
}

.v-select .v-field__input {
  text-align: right;
}

/* إصلاح اتجاه القوائم */
.v-menu .v-list {
  direction: rtl;
}

/* إصلاح اتجاه الشريط الجانبي */
.v-navigation-drawer {
  direction: rtl;
}

/* Mini Sidebar */
.mini-sidebar {
  background: #ffffff !important;
  border-left: 1px solid rgba(0, 0, 0, 0.1) !important;
  box-shadow: -2px 0 10px rgba(0, 0, 0, 0.05) !important;
  z-index: 6 !important;
}



/* إصلاح اتجاه المحتوى */
.v-main {
  direction: rtl;
}


/* إصلاح اتجاه الحاويات */
.v-container {
  direction: rtl;
}

/* إصلاح اتجاه الصفوف والأعمدة */
.v-row {
  direction: rtl;
}

.v-col {
  direction: rtl;
}

/* ========================================
   تخصيص عناوين الجداول - Table Headers
   ======================================== */
.v-data-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-align: center !important;
  vertical-align: middle !important;
  padding: 1.2rem 0.8rem !important;
  border: none !important;
  box-shadow: 0 3px 12px rgba(4, 120, 87, 0.4) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}

/* CSS إضافي لضمان التطبيق */
.v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لجميع عناوين الجداول */
table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS محدد لصفحة إدارة الفريق */
.team-data-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.team-data-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لضمان التطبيق على جميع عناصر Vuetify */
.v-data-table.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* ========================================
   تنسيق Footer السايد بار - تسجيل الخروج
   ======================================== */
.sidebar-footer {
  padding: 8px;
  background: transparent;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  margin-top: auto;
}

/* بطاقة معلومات المستخدم */
.user-info-card {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 10px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
  margin-bottom: 6px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  position: relative;
  overflow: hidden;
}

.user-info-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(255, 255, 255, 0.1),
    transparent
  );
  animation: pulse-sweep 3s infinite;
}

.user-info-card::after {
  content: '';
  position: absolute;
  inset: 0;
  border-radius: 8px;
  padding: 1px;
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.2),
    rgba(255, 255, 255, 0.05),
    rgba(255, 255, 255, 0.2)
  );
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  animation: pulse-border 2s ease-in-out infinite;
  pointer-events: none;
}

@keyframes pulse-sweep {
  0% {
    left: -100%;
  }
  50% {
    left: 100%;
  }
  100% {
    left: 100%;
  }
}

@keyframes pulse-border {
  0%, 100% {
    opacity: 0.5;
  }
  50% {
    opacity: 1;
  }
}

.user-info-card:hover {
  background: rgba(255, 255, 255, 0.08);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border-color: rgba(255, 255, 255, 0.15);
}

.user-info-card:hover::before {
  animation-duration: 2s;
}

.user-info-card:hover::after {
  animation-duration: 1.5s;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0.1) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid rgba(255, 255, 255, 0.3);
  flex-shrink: 0;
  position: relative;
  z-index: 1;
}

.user-details {
  flex: 1;
  min-width: 0;
  position: relative;
  z-index: 1;
}

.user-greeting {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.55rem;
  margin: 0;
  line-height: 1.2;
  font-weight: 400;
}

.user-name {
  color: #ffffff;
  font-size: 0.65rem;
  margin: 1px 0 0 0;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* زر تسجيل الخروج - نفس تنسيق عناصر القائمة */
.logout-menu-item {
  cursor: pointer;
  margin-top: 4px !important;
  border-top: 1px solid rgba(255, 255, 255, 0.1) !important;
  padding-top: 4px !important;
}

.logout-menu-item:hover {
  background: rgba(239, 68, 68, 0.2) !important;
}

.logout-menu-item:hover .v-icon {
  color: #fca5a5 !important;
}

.logout-menu-item:hover .v-list-item-title {
  color: #fca5a5 !important;
}

/* ========================================
   Responsive Design - التصميم المتجاوب
   دعم جميع أجهزة iPad والهواتف
   ======================================== */

/* الهواتف الصغيرة (320px - 480px) */
@media (max-width: 480px) {
  .modern-sidebar {
    width: 280px !important;
  }

  .sidebar-header {
    padding: 12px !important;
  }

  .sidebar-header .text-h5 {
    font-size: 0.7rem !important;
  }

  .sidebar-header .text-body-2 {
    font-size: 0.5rem !important;
  }

  .logo-enhanced {
    height: 60px !important;
    width: 60px !important;
  }

  .modern-sidebar .v-list-item-title {
    font-size: 0.6rem !important;
  }

  .menu-item {
    min-height: 40px !important;
    padding: 4px 8px !important;
  }

  .user-info-card {
    padding: 8px !important;
    gap: 8px !important;
  }

  .user-avatar {
    width: 32px !important;
    height: 32px !important;
  }

  .user-name {
    font-size: 0.6rem !important;
  }

  .user-greeting {
    font-size: 0.5rem !important;
  }

  .sidebar-toggle-btn-fixed {
    top: 12px !important;
    right: 12px !important;
    width: 40px !important;
    height: 40px !important;
  }

  /* v-main padding removed to let Vuetify handle layout */
}

/* الهواتف الكبيرة (481px - 767px) */
@media (min-width: 481px) and (max-width: 767px) {
  .modern-sidebar {
    width: 300px !important;
  }

  .sidebar-header {
    padding: 16px !important;
  }

  .sidebar-header .text-h5 {
    font-size: 0.75rem !important;
  }

  .sidebar-header .text-body-2 {
    font-size: 0.55rem !important;
  }

  .logo-enhanced {
    height: 70px !important;
    width: 70px !important;
  }

  .modern-sidebar .v-list-item-title {
    font-size: 0.65rem !important;
  }

  .menu-item {
    min-height: 44px !important;
    padding: 6px 10px !important;
  }

  .user-info-card {
    padding: 10px !important;
    gap: 10px !important;
  }

  .user-avatar {
    width: 36px !important;
    height: 36px !important;
  }

  .user-name {
    font-size: 0.7rem !important;
  }

  .user-greeting {
    font-size: 0.55rem !important;
  }

  .sidebar-toggle-btn-fixed {
    top: 14px !important;
    right: 14px !important;
  }

  /* v-main padding removed */
}

/* iPad عمودي (768px - 1024px) */
@media (min-width: 768px) and (max-width: 1024px) {
  .modern-sidebar {
    width: 280px !important;
  }

  .sidebar-header {
    padding: 20px !important;
  }

  .sidebar-header .text-h5 {
    font-size: 0.85rem !important;
  }

  .sidebar-header .text-body-2 {
    font-size: 0.6rem !important;
  }

  .logo-enhanced {
    height: 75px !important;
    width: 75px !important;
  }

  .modern-sidebar .v-list-item-title {
    font-size: 0.7rem !important;
  }

  .menu-item {
    min-height: 48px !important;
    padding: 8px 12px !important;
  }

  .user-info-card {
    padding: 12px !important;
    gap: 12px !important;
  }

  .user-avatar {
    width: 40px !important;
    height: 40px !important;
  }

  .user-name {
    font-size: 0.75rem !important;
  }

  .user-greeting {
    font-size: 0.6rem !important;
  }

  /* v-main padding removed */

  /* تحسين الجداول على iPad */
  .v-data-table {
    font-size: 0.85rem !important;
  }

  .v-data-table table thead tr th {
    font-size: 0.9rem !important;
    padding: 12px 8px !important;
  }

  .v-data-table table tbody tr td {
    font-size: 0.8rem !important;
    padding: 10px 8px !important;
  }
}

/* iPad أفقي والشاشات المتوسطة (1025px - 1366px) */
@media (min-width: 1025px) and (max-width: 1366px) {
  .modern-sidebar {
    width: 300px !important;
  }

  .sidebar-header {
    padding: 24px !important;
  }

  .sidebar-header .text-h5 {
    font-size: 0.9rem !important;
  }

  .sidebar-header .text-body-2 {
    font-size: 0.65rem !important;
  }

  .logo-enhanced {
    height: 80px !important;
    width: 80px !important;
  }

  .modern-sidebar .v-list-item-title {
    font-size: 0.75rem !important;
  }

  .menu-item {
    min-height: 50px !important;
    padding: 10px 14px !important;
  }

  /* v-main padding removed */
}

/* الشاشات الكبيرة (1367px+) */
@media (min-width: 1367px) {
  .modern-sidebar {
    width: 320px !important;
  }

  /* v-main padding removed */
}

/* تحسينات عامة للشاشات الصغيرة والمتوسطة */
@media (max-width: 1024px) {
  .user-info-card {
    padding: 10px;
    gap: 10px;
  }

  .logout-btn {
    height: 44px !important;
  }

  .logout-text {
    font-size: 0.9rem;
  }

  /* تحسين الأزرار */
  .v-btn {
    min-width: auto !important;
    padding: 8px 16px !important;
  }

  /* تحسين الحقول */
  .v-text-field {
    font-size: 0.9rem !important;
  }

  /* تحسين البطاقات */
  .v-card {
    margin-bottom: 16px !important;
  }

  .v-card-title {
    font-size: 1rem !important;
    padding: 12px 16px !important;
  }

  .v-card-text {
    padding: 12px 16px !important;
    font-size: 0.9rem !important;
  }
}

/* تحسينات خاصة للهواتف */
@media (max-width: 767px) {
  /* إخفاء بعض العناصر غير الضرورية على الهواتف */
  .sidebar-header .text-body-2 {
    display: none;
  }

  /* تحسين المسافات */
  .v-list {
    padding: 4px !important;
  }

  .v-list-item {
    padding: 4px 8px !important;
    min-height: 40px !important;
  }

  /* تحسين الجداول - تفعيل التمرير الأفقي */
  .v-data-table__wrapper {
    overflow-x: auto !important;
    -webkit-overflow-scrolling: touch !important;
  }

  .v-data-table table {
    min-width: 600px !important;
  }
}

/* دعم اللمس على الأجهزة اللوحية */
@media (hover: none) and (pointer: coarse) {
  .menu-item {
    min-height: 48px !important;
    padding: 12px 16px !important;
  }

  .v-btn {
    min-height: 44px !important;
    padding: 12px 20px !important;
  }

  .sidebar-toggle-btn-fixed {
    width: 48px !important;
    height: 48px !important;
  }
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لجميع عناوين الجداول في التطبيق */
.v-application .v-data-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-application .v-data-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* ========================================
   زر Toggle Sidebar المحسن
   ======================================== */
.sidebar-toggle-btn-fixed {
  position: fixed !important;
  top: 20px !important;
  right: 20px !important;
  z-index: 1000 !important;
  width: 56px !important;
  height: 56px !important;
  min-width: 56px !important;
  border-radius: 50% !important;
  box-shadow: 0 8px 24px rgba(25, 118, 210, 0.3), 0 4px 12px rgba(0, 0, 0, 0.15) !important;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 50%, #0d47a1 100%) !important;
  overflow: visible !important;
}

.sidebar-toggle-btn-fixed::before {
  content: '';
  position: absolute;
  inset: -4px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(25, 118, 210, 0.3), rgba(21, 101, 192, 0.3));
  z-index: -1;
  opacity: 0;
  transition: opacity 0.3s ease;
  animation: pulse-ring 2s ease-in-out infinite;
}

.sidebar-toggle-btn-fixed:hover::before {
  opacity: 1;
}

.sidebar-toggle-btn-fixed:hover {
  transform: scale(1.15) rotate(5deg) !important;
  box-shadow: 0 12px 32px rgba(25, 118, 210, 0.5), 0 8px 16px rgba(0, 0, 0, 0.2) !important;
  background: linear-gradient(135deg, #1e88e5 0%, #1976d2 50%, #1565c0 100%) !important;
}

.sidebar-toggle-btn-fixed:active {
  transform: scale(1.05) !important;
}

.sidebar-toggle-btn-fixed.sidebar-open {
  background: linear-gradient(135deg, #43a047 0%, #388e3c 50%, #2e7d32 100%) !important;
  box-shadow: 0 8px 24px rgba(67, 160, 71, 0.3), 0 4px 12px rgba(0, 0, 0, 0.15) !important;
}

.sidebar-toggle-btn-fixed.sidebar-open:hover {
  background: linear-gradient(135deg, #4caf50 0%, #43a047 50%, #388e3c 100%) !important;
  box-shadow: 0 12px 32px rgba(67, 160, 71, 0.5), 0 8px 16px rgba(0, 0, 0, 0.2) !important;
}

.sidebar-toggle-btn-fixed .toggle-icon,
.sidebar-toggle-btn-fixed .v-icon {
  color: white !important;
  font-size: 28px !important;
  width: 28px !important;
  height: 28px !important;
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1) !important;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2)) !important;
}

.sidebar-toggle-btn-fixed:hover .toggle-icon,
.sidebar-toggle-btn-fixed:hover .v-icon {
  transform: scale(1.1) !important;
}

.sidebar-toggle-btn-fixed.sidebar-open .toggle-icon,
.sidebar-toggle-btn-fixed.sidebar-open .v-icon {
  transform: rotate(180deg) !important;
}

.toggle-ripple {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 0;
  height: 0;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.3);
  pointer-events: none;
  transition: width 0.6s ease, height 0.6s ease, opacity 0.6s ease;
}

.sidebar-toggle-btn-fixed:active .toggle-ripple {
  width: 100px;
  height: 100px;
  opacity: 0;
}

@keyframes pulse-ring {
  0% {
    transform: scale(0.95);
    opacity: 0.7;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.3;
  }
  100% {
    transform: scale(0.95);
    opacity: 0.7;
  }
}

/* زر إغلاق الـ sidebar */
.sidebar-close-btn {
  position: absolute !important;
  top: 8px !important;
  left: 8px !important;
  z-index: 10 !important;
  background: rgba(255, 255, 255, 0.15) !important;
  backdrop-filter: blur(8px) !important;
  border-radius: 50% !important;
  width: 36px !important;
  height: 36px !important;
  transition: all 0.3s ease !important;
}

.sidebar-close-btn:hover {
  background: rgba(255, 255, 255, 0.3) !important;
  transform: scale(1.15) rotate(90deg) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

/* تحسينات للشاشات الصغيرة */
@media (max-width: 960px) {
  .sidebar-toggle-btn-fixed {
    top: 16px !important;
    right: 16px !important;
    width: 52px !important;
    height: 52px !important;
    min-width: 52px !important;
  }

  .sidebar-toggle-btn-fixed .toggle-icon,
  .sidebar-toggle-btn-fixed .v-icon {
    font-size: 24px !important;
    width: 24px !important;
    height: 24px !important;
  }
}

@media (max-width: 480px) {
  .sidebar-toggle-btn-fixed {
    top: 12px !important;
    right: 12px !important;
    width: 48px !important;
    height: 48px !important;
    min-width: 48px !important;
  }

  .sidebar-toggle-btn-fixed .toggle-icon,
  .sidebar-toggle-btn-fixed .v-icon {
    font-size: 22px !important;
    width: 22px !important;
    height: 22px !important;
  }
}
</style>
