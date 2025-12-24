import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

/**
 * List all team members with pagination
 * @param {Object} params - Query parameters
 * @param {number} params.page - Page number (default: DEFAULT_PAGE)
 * @param {number} params.limit - Items per page (default: DEFAULT_LIMIT)
 * @returns {Promise<{data: Array, total: number, page: number, limit: number, totalPages: number}>}
 */
export async function listTeamMembers({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/team-members?${query}`)
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
 * Get a single team member by ID
 * @param {number|string} id - Team member ID
 * @returns {Promise<Object>} Team member data
 */
export async function getTeamMember(id) {
  const result = await apiFetch(`/team-members/${id}`)
  return result?.data || result
}

/**
 * Get team member statistics
 * @param {Object} params - Query parameters
 * @param {string} params.period - Period filter ('month', 'year', or empty for all time)
 * @returns {Promise<Object>} Stats object with total, uniqueUsers, uniqueProjects, avgPerProject
 */
export async function getTeamMemberStats(params = {}) {
  const query = new URLSearchParams(params).toString()
  const suffix = query ? `?${query}` : ''
  const result = await apiFetch(`/team-members/stats${suffix}`)
  return result?.data || {}
}

/**
 * Create a new team member (add user to project)
 * @param {Object} teamMemberData - Team member data
 * @param {number} teamMemberData.projectId - Project ID (required)
 * @param {number} teamMemberData.userId - User ID (required)
 * @returns {Promise<Object>} Created team member data
 */
export async function createTeamMember(teamMemberData) {
  const result = await apiFetch('/team-members', {
    method: 'POST',
    body: teamMemberData
  })
  return result?.data || result
}

/**
 * Delete a team member (remove user from project)
 * @param {number|string} id - Team member ID
 * @returns {Promise<Object>} Response
 */
export async function deleteTeamMember(id) {
  const result = await apiFetch(`/team-members/${id}`, {
    method: 'DELETE'
  })
  return result?.data || result
}
