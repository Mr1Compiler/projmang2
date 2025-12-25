import { apiFetch } from './client'

/**
 * List all role-page assignments with pagination
 * @param {Object} params - Query parameters
 * @param {number} params.page - Page number
 * @param {number} params.limit - Items per page
 * @returns {Promise<Object>} Paginated role-pages data
 */
export async function listRolePages(params = {}) {
  const queryParams = new URLSearchParams()
  if (params.page) queryParams.append('page', params.page)
  if (params.limit) queryParams.append('limit', params.limit)

  const queryString = queryParams.toString()
  const path = queryString ? `/rolePages?${queryString}` : '/rolePages?limit=100'

  const result = await apiFetch(path)
  return result?.data || { data: [], total: 0, page: 1, totalPages: 0 }
}

/**
 * Get a single role-page assignment by ID
 * @param {number|string} id - RolePage ID
 * @returns {Promise<Object>} RolePage data
 */
export async function getRolePage(id) {
  const result = await apiFetch(`/rolePages/${id}`)
  return result?.data || result
}

/**
 * Get all pages assigned to a specific role
 * @param {number|string} roleId - Role ID
 * @returns {Promise<Array>} Array of role-page assignments
 */
export async function getRolePagesByRoleId(roleId) {
  const result = await apiFetch(`/rolePages?roleId=${roleId}&limit=100`)
  // Handle response - when using roleId filter, backend returns array directly
  const data = result?.data || []
  const rolePages = Array.isArray(data) ? data : []

  // Parse permissions from JSON string to array
  return rolePages.map(rp => ({
    ...rp,
    permissions: rp.permissions ? JSON.parse(rp.permissions) : []
  }))
}

/**
 * Create a new role-page assignment (assign page to role with permissions)
 * @param {Object} rolePageData - RolePage data
 * @param {number} rolePageData.roleId - Role ID (required)
 * @param {number} rolePageData.pageId - Page ID (required)
 * @param {Array<string>} rolePageData.permissions - Permissions array (e.g., ['read', 'write', 'delete'])
 * @returns {Promise<Object>} Created role-page data
 */
export async function createRolePage(rolePageData) {
  // Backend expects permissions as JSON string, not array
  const payload = {
    roleId: rolePageData.roleId,
    pageId: rolePageData.pageId,
    permissions: JSON.stringify(rolePageData.permissions)
  }
  const result = await apiFetch('/rolePages', {
    method: 'POST',
    body: payload
  })
  return result?.data || result
}

/**
 * Update an existing role-page assignment (update permissions)
 * @param {number|string} id - RolePage ID
 * @param {Object} rolePageData - RolePage data to update
 * @param {Array<string>} rolePageData.permissions - New permissions array
 * @returns {Promise<Object>} Updated role-page data
 */
export async function updateRolePage(id, rolePageData) {
  // Backend expects permissions as JSON string, not array
  const payload = {
    permissions: JSON.stringify(rolePageData.permissions)
  }
  const result = await apiFetch(`/rolePages/${id}`, {
    method: 'PUT',
    body: payload
  })
  return result?.data || result
}

/**
 * Delete a role-page assignment (remove page from role)
 * @param {number|string} id - RolePage ID
 * @returns {Promise<Object>} Response
 */
export async function deleteRolePage(id) {
  const result = await apiFetch(`/rolePages/${id}`, {
    method: 'DELETE'
  })
  return result?.data || result
}

/**
 * Bulk update role pages - atomically replaces all pages for a role in a single transaction
 * @param {number|string} roleId - Role ID
 * @param {Array<Object>} pagePermissions - Array of { pageId, permissions }
 * @returns {Promise<Array>} Updated role-pages
 */
export async function bulkUpdateRolePages(roleId, pagePermissions) {
  // Filter out pages with no permissions and format for backend
  const pages = pagePermissions
    .filter(pp => pp.permissions && pp.permissions.length > 0)
    .map(pp => ({
      pageId: Number(pp.pageId),
      permissions: JSON.stringify(pp.permissions)
    }))

  const result = await apiFetch(`/rolePages/role/${roleId}`, {
    method: 'PUT',
    body: { pages }
  })

  return result?.data || result
}
