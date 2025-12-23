import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

/**
 * List all users with pagination
 * @param {Object} params - Query parameters
 * @param {number} params.page - Page number (default: DEFAULT_PAGE)
 * @param {number} params.limit - Items per page (max 100, default: DEFAULT_LIMIT)
 * @returns {Promise<{data: Array, total: number, page: number, limit: number, totalPages: number}>}
 */
export async function listUsers({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/users?${query}`)
  const paginatedData = result?.data || {}
  return {
    data: paginatedData.data || [],
    total: paginatedData.total || 0,
    page: paginatedData.page || page,
    limit: paginatedData.limit || limit,
    totalPages: paginatedData.totalPages || 0
  }
}

/**
 * Get a single user by ID
 * @param {number|string} id - User ID
 * @returns {Promise<Object>} User data
 */
export async function getUser(id) {
  const result = await apiFetch(`/users/${id}`)
  return result?.data || result
}

/**
 * Create a new user
 * @param {Object} userData - User data
 * @param {string} userData.username - Username (required)
 * @param {string} userData.email - Email (required, valid email format)
 * @param {string} userData.password - Password (required, min 8 characters)
 * @param {string} userData.fullName - Full name (required)
 * @param {string} userData.phone - Phone number (required)
 * @param {string} userData.jobTitle - Job title (required)
 * @returns {Promise<Object>} Created user data
 */
export async function createUser(userData) {
  const result = await apiFetch('/users', {
    method: 'POST',
    body: userData
  })
  return result?.data || result
}

/**
 * Update an existing user
 * @param {number|string} id - User ID
 * @param {Object} userData - User data to update (all fields optional)
 * @param {string} [userData.username] - Username
 * @param {string} [userData.email] - Email
 * @param {string} [userData.fullName] - Full name
 * @param {string} [userData.phone] - Phone number
 * @param {string} [userData.avatar] - Avatar URL
 * @param {string} [userData.jobTitle] - Job title
 * @param {string} [userData.status] - Status
 * @returns {Promise<Object>} Updated user data
 */
export async function updateUser(id, userData) {
  const result = await apiFetch(`/users/${id}`, {
    method: 'PUT',
    body: userData
  })
  return result?.data || result
}

/**
 * Update user password
 * @param {number|string} id - User ID
 * @param {string} password - New password (min 8 characters)
 * @returns {Promise<Object>} Response
 */
export async function updateUserPassword(id, password) {
  const result = await apiFetch(`/users/${id}/password`, {
    method: 'PUT',
    body: { password }
  })
  return result?.data || result
}

/**
 * Update user status
 * @param {number|string} id - User ID
 * @param {string} status - New status ('active' or 'inactive')
 * @returns {Promise<Object>} Response
 */
export async function updateUserStatus(id, status) {
  const result = await apiFetch(`/users/${id}/status`, {
    method: 'PATCH',
    body: { status }
  })
  return result?.data || result
}

/**
 * Delete a user
 * @param {number|string} id - User ID
 * @returns {Promise<Object>} Response
 */
export async function deleteUser(id) {
  const result = await apiFetch(`/users/${id}`, {
    method: 'DELETE'
  })
  return result?.data || result
}

/**
 * Get user roles
 * @param {number|string} id - User ID
 * @returns {Promise<Array>} User roles
 */
export async function getUserRoles(id) {
  const result = await apiFetch(`/users/${id}/roles`)
  return result?.data || result
}

/**
 * Get user team members
 * @param {number|string} id - User ID
 * @returns {Promise<Array>} Team members
 */
export async function getUserTeamMembers(id) {
  const result = await apiFetch(`/users/${id}/team-members`)
  return result?.data || result
}
