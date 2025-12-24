import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getMyPages } from '@/api/auth'

// Central permission constants
export const PERMISSIONS = {
  READ: 'read',
  CREATE: 'create',
  UPDATE: 'update',
  DELETE: 'delete',
  UPDATE_PASSWORD: 'updatePassword',
  UPDATE_STATUS: 'updateStatus'
}

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref(null)
  const pages = ref([])
  const isLoading = ref(false)

  // Computed
  const isAuthenticated = computed(() => {
    return localStorage.getItem('isAuthenticated') === 'true' &&
           Boolean(localStorage.getItem('accessToken'))
  })

  const allowedRoutes = computed(() => {
    return pages.value.map(p => p.route)
  })

  // Get permissions map: { '/projects': ['read', 'write', 'delete'], ... }
  const permissionsMap = computed(() => {
    const map = {}
    pages.value.forEach(page => {
      map[page.route] = page.permissions || []
    })
    return map
  })

  // Actions
  async function fetchPages() {
    if (isLoading.value) return
    isLoading.value = true
    try {
      const data = await getMyPages()
      pages.value = data
      // Persist to localStorage for page refresh
      localStorage.setItem('userPages', JSON.stringify(data))
    } catch (error) {
      console.error('Failed to fetch pages:', error)
      pages.value = []
    } finally {
      isLoading.value = false
    }
  }

  function loadFromStorage() {
    // Load user
    const userRaw = localStorage.getItem('user')
    if (userRaw) {
      try {
        user.value = JSON.parse(userRaw)
      } catch (e) {
        user.value = null
      }
    }
    // Load pages
    const pagesRaw = localStorage.getItem('userPages')
    if (pagesRaw) {
      try {
        pages.value = JSON.parse(pagesRaw)
      } catch (e) {
        pages.value = []
      }
    }
  }

  function setUser(userData) {
    user.value = userData
  }

  function clearAuth() {
    user.value = null
    pages.value = []
    localStorage.removeItem('userPages')
  }

  // Permission check helpers
  function hasPermission(route, permission) {
    const perms = permissionsMap.value[route]
    return perms ? perms.includes(permission) : false
  }

  function canAccessPage(route) {
    return allowedRoutes.value.includes(route)
  }

  function getPagePermissions(route) {
    return permissionsMap.value[route] || []
  }

  return {
    // State
    user,
    pages,
    isLoading,
    // Computed
    isAuthenticated,
    allowedRoutes,
    permissionsMap,
    // Actions
    fetchPages,
    loadFromStorage,
    setUser,
    clearAuth,
    hasPermission,
    canAccessPage,
    getPagePermissions
  }
})
