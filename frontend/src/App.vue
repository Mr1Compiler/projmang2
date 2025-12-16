<template>
  <v-app dir="rtl" class="arabic-app">
    <!-- الشريط الجانبي الأيمن -->
    <v-navigation-drawer
      v-model="drawer"
      location="right"
      permanent
      width="320"
      class="modern-sidebar rtl-sidebar"
    >
      <!-- الشعار والعنوان -->
      <div class="sidebar-header pa-6">
        <div class="text-center mb-6">
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
          class="menu-item mb-2"
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

      <!-- قسم المميزات -->
      <div class="px-4 mt-4">
        <h3 class="text-h6 font-weight-bold text-white mb-3">المميزات الرئيسية</h3>
        <div class="feature-grid">
          <div
            v-for="feature in features"
            :key="feature.title"
            class="feature-item"
          >
            <v-icon :color="feature.color" size="20" class="mb-1">{{ feature.icon }}</v-icon>
            <div class="text-caption text-white text-opacity-80">{{ feature.title }}</div>
          </div>
          <!-- رابط إدارة الفريق -->
          <router-link to="/team-management" class="feature-item feature-link">
            <v-icon color="success" size="20" class="mb-1">mdi-account-group</v-icon>
            <div class="text-caption text-white text-opacity-80">إدارة الفريق</div>
            <v-chip size="x-small" color="success" class="mt-1">جديد</v-chip>
          </router-link>
        </div>
      </div>

      <!-- Footer -->
      <template v-slot:append>
        <div class="pa-4 text-center">
                 <p class="text-caption text-white text-opacity-60">
                
                 </p>
                 <p class="text-caption text-white text-opacity-60">
                 </p>
        </div>
      </template>
    </v-navigation-drawer>

    <!-- المحتوى الرئيسي -->
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const drawer = ref(true)

// قائمة التنقل الرئيسية
const mainMenuItems = ref([
  { title: 'الرئيسية', icon: 'mdi-view-dashboard', to: '/', active: false },
  { title: 'المشاريع', icon: 'mdi-folder-multiple', to: '/project-management', active: false },
  { title: 'إدارة المهام', icon: 'mdi-clipboard-list', to: '/task-management', active: false, badge: 'جديد', badgeColor: 'success' },
  { title: 'المهندسين', icon: 'mdi-account-hard-hat', to: '/engineers', active: false },
  { title: 'التصنيفات', icon: 'mdi-tag-multiple', to: '/categories', active: false },
  { title: 'المصروفات الإدارية', icon: 'mdi-chart-line', to: '/expenses', active: false },
  { title: 'أنواع المصروفات', icon: 'mdi-format-list-bulleted-type', to: '/expense-types', active: false },
  { title: 'الإيرادات', icon: 'mdi-trending-up', to: '/income', active: false },
  { title: 'المديونون', icon: 'mdi-credit-card', to: '/debtors', active: false, badge: '3', badgeColor: 'error' },
  { title: 'المستخدمين', icon: 'mdi-account-multiple', to: '/users', active: false },
  { title: 'الموارد البشرية', icon: 'mdi-account-group', to: '/human-resources', active: false },
  { title: 'بصمة الموظفين', icon: 'mdi-fingerprint', to: '/human-resources#fingerprint', active: false, badge: 'جديد', badgeColor: 'primary' }
])

const features = ref([
  {
    icon: 'mdi-chart-line',
    title: 'تقارير المشاريع',
    color: 'primary'
  },
  {
    icon: 'mdi-calculator',
    title: 'إدارة المالية',
    color: 'warning'
  },
  {
    icon: 'mdi-shield-check',
    title: 'أمان عالي',
    color: 'error'
  }
])

// تحديث العنصر النشط حسب الصفحة الحالية
const updateActiveMenuItem = () => {
  mainMenuItems.value.forEach(item => {
    item.active = item.to === route.path
  })
}

// مراقبة تغيير الصفحة
watch(() => route.path, updateActiveMenuItem, { immediate: true })
</script>

<style>
/* دعم RTL للنظام العربي */
.arabic-app {
  direction: rtl;
  text-align: right;
  font-family: 'Cairo', 'Tajawal', 'Noto Sans Arabic', 'Arial', sans-serif;
}

/* ========================================
   السايد بار العصري - Protected from page CSS leaks
   ======================================== */
.modern-sidebar,
.v-navigation-drawer.modern-sidebar,
.v-application .modern-sidebar,
.v-application .v-navigation-drawer.modern-sidebar,
body .modern-sidebar,
body .v-navigation-drawer.modern-sidebar {
  background: linear-gradient(135deg, #2563eb 0%, #7c3aed 100%) !important;
  border-left: none !important;
  box-shadow: -4px 0 20px rgba(37, 99, 235, 0.3) !important;
  color: white !important;
}

/* Protect sidebar from page-level color overrides */
.modern-sidebar *,
.v-navigation-drawer.modern-sidebar *,
body .modern-sidebar * {
  color: white !important;
  -webkit-text-fill-color: white !important;
}

.modern-sidebar .v-icon,
.v-navigation-drawer.modern-sidebar .v-icon {
  color: white !important;
}

.rtl-sidebar {
  direction: rtl;
  text-align: right;
}

.sidebar-header {
  background: rgba(255, 255, 255, 0.1) !important;
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
  border-radius: 12px !important;
  margin-bottom: 8px;
}

.menu-item .v-list-item-title {
  font-size: var(--font-size-base-minus) !important;
}

.modern-sidebar .v-list-item-title {
  font-size: var(--font-size-base-minus) !important;
  color: white !important;
  -webkit-text-fill-color: white !important;
}

/* Override page-level CSS that leaks into sidebar */
.modern-sidebar .v-list .v-list-item,
.modern-sidebar .v-list .v-list-item *,
.modern-sidebar .v-list .v-list-item .v-list-item__content,
.modern-sidebar .v-list .v-list-item .v-list-item__title,
.modern-sidebar .v-list-item,
.modern-sidebar .v-list-item *,
.modern-sidebar .v-list-item .v-list-item-title {
  color: white !important;
  -webkit-text-fill-color: white !important;
  text-shadow: none !important;
  font-weight: 500 !important;
  font-size: var(--font-size-base-minus) !important;
  animation: none !important;
}

.modern-sidebar .v-list-item:hover {
  animation: none !important;
}

.modern-sidebar p,
.modern-sidebar span {
  font-size: var(--font-size-sm-plus) !important;
  color: white !important;
  -webkit-text-fill-color: white !important;
}

.modern-sidebar .text-h5 {
  font-size: var(--font-size-base) !important;
}

.modern-sidebar .text-body-2 {
  font-size: var(--font-size-sm-minus) !important;
}

.modern-sidebar .text-caption {
  font-size: var(--font-size-xs) !important;
}

.modern-sidebar .menu-item,
.modern-sidebar .v-list-item.menu-item,
body .modern-sidebar .menu-item {
  background: transparent !important;
}

.modern-sidebar .menu-item:hover,
.modern-sidebar .v-list-item.menu-item:hover,
body .modern-sidebar .menu-item:hover {
  background: rgba(37, 99, 235, 0.2) !important;
  transform: translateX(-4px);
  animation: none !important;
}

.modern-sidebar .active-menu-item,
.modern-sidebar .v-list-item.active-menu-item,
body .modern-sidebar .active-menu-item {
  background: linear-gradient(135deg, #1d4ed8 0%, #6d28d9 100%) !important;
  box-shadow: 0 4px 15px rgba(37, 99, 235, 0.4);
  transform: translateX(-4px);
}

.modern-sidebar .active-menu-item .v-list-item-title,
body .modern-sidebar .active-menu-item .v-list-item-title {
  font-weight: 600 !important;
  color: white !important;
}

/* Protect chips/badges in sidebar */
.modern-sidebar .v-chip,
body .modern-sidebar .v-chip {
  color: white !important;
}

.modern-sidebar .v-chip .v-chip__content,
body .modern-sidebar .v-chip .v-chip__content {
  color: white !important;
}

/* شبكة المميزات */
.feature-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.feature-item {
  background: rgba(5, 150, 105, 0.2);
  border-radius: 8px;
  padding: 12px 8px;
  text-align: center;
  transition: all 0.3s ease;
  backdrop-filter: blur(5px);
}

.feature-item:hover {
  background: rgba(5, 150, 105, 0.3);
  transform: translateY(-2px);
}

/* رابط إدارة الفريق */
.feature-link {
  text-decoration: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: rgba(34, 197, 94, 0.2) !important;
  border: 1px solid rgba(34, 197, 94, 0.3);
  position: relative;
  overflow: hidden;
}

.feature-link:hover {
  background: rgba(34, 197, 94, 0.4) !important;
  border-color: rgba(34, 197, 94, 0.5);
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 8px 25px rgba(34, 197, 94, 0.3);
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

/* إصلاح اتجاه الأيقونات الاتجاهية فقط (السهام والشيفرونات) */
/* Note: Don't flip all icons - only directional ones need flipping for RTL */
.v-icon.mdi-chevron-left,
.v-icon.mdi-chevron-right,
.v-icon.mdi-arrow-left,
.v-icon.mdi-arrow-right,
.v-icon.mdi-menu-left,
.v-icon.mdi-menu-right {
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
  font-size: var(--font-size-base-plus) !important;
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
  font-size: var(--font-size-base-plus) !important;
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
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لجميع عناوين الجداول */
table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS محدد لصفحة إدارة الفريق */
.team-data-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.team-data-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لضمان التطبيق على جميع عناصر Vuetify */
.v-data-table.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لجميع عناوين الجداول في التطبيق */
.v-application .v-data-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-application .v-data-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: var(--font-size-base-plus) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* ========================================
   Main Content Area - Ensure proper colors
   ======================================== */

/* Ensure page content has proper colors */
.v-main .data-page,
.v-main .project-card,
.v-main .stat-card {
  color: inherit;
}

/* Allow page-specific text colors to work */
.v-main .data-page .page-content-section p,
.v-main .data-page .page-content-section span,
.v-main .data-page .page-content-section h1,
.v-main .data-page .page-content-section h2,
.v-main .data-page .page-content-section h3,
.v-main .data-page .page-content-section h4,
.v-main .data-page .page-content-section h5,
.v-main .data-page .page-content-section h6 {
  -webkit-text-fill-color: unset;
}

/* ========================================
   Toast Notifications RTL Support
   ======================================== */
.Toastify__toast-container {
  direction: rtl !important;
  z-index: 9999 !important;
}

.Toastify__toast {
  direction: rtl !important;
  font-family: 'Cairo', 'Tajawal', 'Noto Sans Arabic', 'Arial', sans-serif !important;
}

.Toastify__toast-body {
  direction: rtl !important;
  text-align: right !important;
}

/* Position toast on left for RTL */
.Toastify__toast-container--top-right {
  right: auto !important;
  left: 1em !important;
}

.Toastify__toast-container--bottom-right {
  right: auto !important;
  left: 1em !important;
}
</style>
