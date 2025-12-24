/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import { routes } from 'vue-router/auto-routes'
import { isAuthenticated } from '@/services/authService'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

// خريطة تحويل مسارات Vue إلى مسارات API للصلاحيات
// فقط للصفحات التي لها أسماء مختلفة في Vue عن API
const routeToPermissionMap = {
  '/project-management': '/projects',
  '/task-management': '/projects',
  '/engineers': '/users',
  '/expense-types': '/expenses',
  '/human-resources': '/users',
  '/team-management': '/projects',
  '/working-day-details': '/work-day-details', // مسار بديل لنفس الصفحة
  '/work-days': '/workdays', // أيام العمل
  '/work-day-details': '/workdays', // تفاصيل يوم العمل
  '/labor-details': '/workday-labor', // تفاصيل العمالة
  '/equipment-details': '/workday-equipment', // تفاصيل المعدات
  '/materials-expenses-details': '/workday-materials' // تفاصيل المواد
}

// دالة لاستخراج المسار الأساسي من المسار الكامل
function getBaseRoute(path) {
  // إزالة المعرفات الرقمية من المسار (مثل /projects/1 -> /projects)
  const segments = path.split('/').filter(Boolean)
  if (segments.length === 0) return '/'

  // أخذ الجزء الأول فقط
  return '/' + segments[0]
}

// دالة للحصول على مسار الصلاحيات
function getPermissionRoute(path, authStore) {
  const baseRoute = getBaseRoute(path)

  // أولاً: تحقق إذا كان المسار الأساسي موجود مباشرة في صلاحيات المستخدم
  if (authStore.canAccessPage(baseRoute)) {
    return baseRoute
  }

  // ثانياً: تحقق من خريطة التحويل للمسار الكامل
  if (routeToPermissionMap[path]) {
    return routeToPermissionMap[path]
  }

  // ثالثاً: تحقق من خريطة التحويل للمسار الأساسي
  if (routeToPermissionMap[baseRoute]) {
    return routeToPermissionMap[baseRoute]
  }

  // رابعاً: أعد المسار الأساسي كما هو
  return baseRoute
}

// Navigation Guard - حماية الصفحات
router.beforeEach((to, from, next) => {
  // الصفحات المسموح الوصول إليها بدون تسجيل دخول
  const publicPages = ['/login']
  const isPublicPage = publicPages.includes(to.path)

  // إذا كان المستخدم غير مسجل دخول ومحاولة الوصول لصفحة محمية
  if (!isAuthenticated() && !isPublicPage) {
    // إعادة التوجيه إلى صفحة تسجيل الدخول
    next('/login')
    return
  }

  // إذا كان المستخدم مسجل دخول ومحاولة الوصول لصفحة تسجيل الدخول
  if (isAuthenticated() && isPublicPage) {
    // إعادة التوجيه إلى الصفحة الرئيسية
    next('/')
    return
  }

  // التحقق من صلاحيات الوصول للصفحة
  if (isAuthenticated() && !isPublicPage) {
    const authStore = useAuthStore()

    // Load permissions from storage if not loaded
    if (authStore.pages.length === 0) {
      authStore.loadFromStorage()
    }

    // Handle dashboard - redirect to first available page if no access
    if (to.path === '/') {
      if (!authStore.canAccessPage('/')) {
        // Redirect to first available page
        const firstPage = authStore.pages.find(p => p.route !== '/')
        if (firstPage) {
          next(firstPage.route)
          return
        }
      }
      next()
      return
    }

    // تحويل مسار Vue إلى مسار API للصلاحيات
    const permissionRoute = getPermissionRoute(to.path, authStore)

    // Check if user has access to this page
    if (authStore.pages.length > 0 && !authStore.canAccessPage(permissionRoute)) {
      // Redirect to dashboard if not allowed
      next('/')
      return
    }
  }

  // السماح بالوصول
  next()
})

// Workaround for https://github.com/vitejs/vite/issues/11804
router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (localStorage.getItem('vuetify:dynamic-reload')) {
      console.error('Dynamic import error, reloading page did not fix it', err)
    } else {
      console.log('Reloading page to fix dynamic import error')
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router
