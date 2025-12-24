import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore, PERMISSIONS } from '@/stores/auth'

export function usePermissions(pagePath = null) {
  const route = useRoute()
  const authStore = useAuthStore()

  // Get current page path (from param or current route)
  const currentPath = computed(() => {
    return pagePath || route.path
  })

  // Get permissions for current page
  const permissions = computed(() => {
    return authStore.getPagePermissions(currentPath.value)
  })

  // Permission check functions
  const hasPermission = (permission) => {
    return permissions.value.includes(permission)
  }

  const canRead = computed(() => hasPermission(PERMISSIONS.READ))
  const canCreate = computed(() => hasPermission(PERMISSIONS.CREATE))
  const canUpdate = computed(() => hasPermission(PERMISSIONS.UPDATE))
  const canWrite = computed(() => hasPermission(PERMISSIONS.CREATE) || hasPermission(PERMISSIONS.UPDATE))
  const canDelete = computed(() => hasPermission(PERMISSIONS.DELETE))
  const canUpdatePassword = computed(() => hasPermission(PERMISSIONS.UPDATE_PASSWORD))
  const canUpdateStatus = computed(() => hasPermission(PERMISSIONS.UPDATE_STATUS))

  // Check permission for a specific page (not current)
  const hasPermissionFor = (path, permission) => {
    return authStore.hasPermission(path, permission)
  }

  // Check if user can access a page
  const canAccessPage = (path) => {
    return authStore.canAccessPage(path)
  }

  return {
    permissions,
    hasPermission,
    canRead,
    canCreate,
    canUpdate,
    canWrite,
    canDelete,
    canUpdatePassword,
    canUpdateStatus,
    hasPermissionFor,
    canAccessPage,
    PERMISSIONS
  }
}
