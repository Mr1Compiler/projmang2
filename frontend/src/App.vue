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
          :ripple="false"
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
/* ========================================
   السايد بار العصري - Protected from page CSS leaks
   ======================================== */
.modern-sidebar,
.v-navigation-drawer.modern-sidebar,
.v-application .modern-sidebar,
.v-application .v-navigation-drawer.modern-sidebar,
body .modern-sidebar,
body .v-navigation-drawer.modern-sidebar {
  background: var(--sidebar-bg-gradient) !important;
  border-left: none !important;
  box-shadow: -4px 0 20px rgba(37, 99, 235, 0.3) !important;
  color: var(--text-white) !important;
}

/* Make inner list transparent to show gradient */
.modern-sidebar .v-list,
.v-navigation-drawer.modern-sidebar .v-list {
  background: transparent !important;
}

/* Protect sidebar from page-level color overrides */
.modern-sidebar *,
.v-navigation-drawer.modern-sidebar *,
body .modern-sidebar * {
  color: var(--text-white) !important;
  -webkit-text-fill-color: var(--text-white) !important;
}

.modern-sidebar .v-icon,
.v-navigation-drawer.modern-sidebar .v-icon {
  color: var(--text-white) !important;
}

.rtl-sidebar {
  direction: rtl;
  text-align: right;
}

.sidebar-header {
  background: rgba(255, 255, 255, 0.1) !important;
  backdrop-filter: blur(10px);
  border-radius: 0 0 20px 20px;
  margin: -16px -16px 8px -16px; /* Reduced bottom margin */
}

.logo-enhanced {
  filter: drop-shadow(0 4px 8px var(--shadow-medium));
  border-radius: var(--radius-full);
  background: rgba(255, 255, 255, 0.1);
  padding: var(--space-1); /* Reduced padding */
}

/* عناصر القائمة */
.menu-item {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border-radius: var(--radius-xl) !important;
  margin-bottom: 4px; /* Reduced spacing */
}

.menu-item .v-list-item-title {
  font-size: var(--font-size-base-minus) !important;
}

.modern-sidebar .v-list-item-title {
  font-size: var(--font-size-base-minus) !important;
  color: var(--text-white) !important;
  -webkit-text-fill-color: var(--text-white) !important;
}

/* Override page-level CSS that leaks into sidebar */
/* Limit global impact by checking parent classes */
/* Limit global impact by checking parent classes */
.modern-sidebar .v-list-item, 
.v-navigation-drawer.modern-sidebar .v-list-item {
  background-color: transparent !important;
  color: var(--text-white) !important;
}

/* Force specific internal elements to be transparent */
.modern-sidebar .v-list-item .v-list-item__overlay,
.modern-sidebar .v-list-item .v-list-item__content,
.modern-sidebar .v-list-item .v-list-item__prepend,
.modern-sidebar .v-list-item .v-list-item__append,
.modern-sidebar .v-list-item .v-list-item-title,
.v-navigation-drawer.modern-sidebar .v-list-item .v-list-item__overlay,
.v-navigation-drawer.modern-sidebar .v-list-item .v-list-item__content,
.v-navigation-drawer.modern-sidebar .v-list-item .v-list-item__prepend,
.v-navigation-drawer.modern-sidebar .v-list-item .v-list-item__append {
  background-color: transparent !important;
  background: transparent !important;
}

.modern-sidebar .v-list-item__content,
.modern-sidebar .v-list-item__title {
  color: var(--text-white) !important;
}

.modern-sidebar .v-list-item:hover,
.v-navigation-drawer.modern-sidebar .v-list-item:hover {
  background-color: rgba(255, 255, 255, 0.1) !important;
  animation: none !important;
}

/* Icon specific fix */
.modern-sidebar .v-icon,
.v-navigation-drawer.modern-sidebar .v-icon {
  background-color: transparent !important;
  background: transparent !important;
  box-shadow: none !important;
  display: inline-flex !important;
  align-items: center;
}

/* Alignment Fix */
.modern-sidebar .v-list-item,
.v-navigation-drawer.modern-sidebar .v-list-item {
  display: flex !important;
  align-items: center !important;
  min-height: 48px; /* Ensure consistent height */
}

/* Ensure prepend container alignment and transparency */
.modern-sidebar .v-list-item__prepend,
.v-navigation-drawer.modern-sidebar .v-list-item__prepend {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  margin-inline-end: 12px;
  background-color: transparent !important;
  background: transparent !important;
}

/* Kill all Vuetify overlay effects (ripples, focus, etc) */
.modern-sidebar .v-list-item::before,
.modern-sidebar .v-list-item::after,
.modern-sidebar .v-list-item__overlay,
.v-navigation-drawer.modern-sidebar .v-list-item::before,
.v-navigation-drawer.modern-sidebar .v-list-item::after,
.v-navigation-drawer.modern-sidebar .v-list-item__overlay {
  display: none !important;
  opacity: 0 !important;
  background-color: transparent !important;
}

/* Ensure no background change on focus/active unless it's our custom class */
.modern-sidebar .v-list-item:focus,
.modern-sidebar .v-list-item:active,
.modern-sidebar .v-list-item--active,
.v-navigation-drawer.modern-sidebar .v-list-item:focus,
.v-navigation-drawer.modern-sidebar .v-list-item:active,
.v-navigation-drawer.modern-sidebar .v-list-item--active {
  background-color: transparent !important;
}

/* Strictly remove background from icons in sidebar */
.modern-sidebar .v-list-item__prepend > .v-icon,
.modern-sidebar .v-icon,
.v-navigation-drawer.modern-sidebar .v-icon {
  background-color: transparent !important;
  background: transparent !important;
  box-shadow: none !important;
  border: none !important;
  display: inline-flex !important;
}

/* Ensure proper coloring for active state icons */
.modern-sidebar .active-menu-item .v-icon,
.v-navigation-drawer.modern-sidebar .active-menu-item .v-icon {
  color: var(--text-white) !important;
  opacity: 1 !important;
}

.modern-sidebar p,
.modern-sidebar span {
  font-size: var(--font-size-sm-plus) !important;
  color: var(--text-white) !important;
  -webkit-text-fill-color: var(--text-white) !important;
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
  background: rgba(255, 255, 255, 0.1) !important;
  transform: translateX(-4px);
  animation: none !important;
}

.modern-sidebar .active-menu-item,
.modern-sidebar .v-list-item.active-menu-item,
body .modern-sidebar .active-menu-item {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0.1) 100%) !important;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  transform: translateX(-4px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.modern-sidebar .active-menu-item .v-list-item-title,
body .modern-sidebar .active-menu-item .v-list-item-title {
  font-weight: 700 !important;
  color: white !important;
}

/* Protect chips/badges in sidebar */
.modern-sidebar .v-chip,
.v-navigation-drawer.modern-sidebar .v-chip {
  color: var(--text-white) !important;
  font-weight: bold !important;
  border-radius: var(--radius-lg);
  /* Allow background color from prop to work, but reset if needed */
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
}

/* Ensure chip content doesn't inherit transparent background */
.modern-sidebar .v-chip .v-chip__content,
.v-navigation-drawer.modern-sidebar .v-chip .v-chip__content {
  color: var(--text-white) !important;
  background-color: transparent !important;
}

/* Fix specific colors for badges if they look washed out */
.modern-sidebar .v-chip.bg-error,
.modern-sidebar .v-chip.text-error {
  background-color: var(--color-error) !important;
}

.modern-sidebar .v-chip.bg-success,
.modern-sidebar .v-chip.text-success {
  background-color: var(--color-success) !important;
}

.modern-sidebar .v-chip.bg-primary,
.modern-sidebar .v-chip.text-primary {
  background-color: var(--color-primary) !important;
}

/* شبكة المميزات */
.feature-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--space-3);
}

.feature-item {
  background: rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-lg);
  padding: var(--space-3) 8px;
  text-align: center;
  transition: all 0.3s ease;
  backdrop-filter: blur(5px);
}

.feature-item:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateY(-2px);
}

/* رابط إدارة الفريق */
.feature-link {
  text-decoration: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: rgba(255, 255, 255, 0.15) !important;
  border: 1px solid rgba(255, 255, 255, 0.2);
  position: relative;
  overflow: hidden;
}

.feature-link:hover {
  background: rgba(255, 255, 255, 0.25) !important;
  border-color: rgba(255, 255, 255, 0.4);
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
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
