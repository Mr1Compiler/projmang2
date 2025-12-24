import { apiFetch } from './client'

/**
 * List all pages with pagination
 * @param {Object} params - Query parameters
 * @param {number} params.page - Page number
 * @param {number} params.limit - Items per page
 * @returns {Promise<Object>} Paginated pages data
 */
export async function listPages(params = {}) {
  const queryParams = new URLSearchParams()
  if (params.page) queryParams.append('page', params.page)
  if (params.limit) queryParams.append('limit', params.limit)

  const queryString = queryParams.toString()
  const path = queryString ? `/pages?${queryString}` : '/pages?limit=100'

  const result = await apiFetch(path)
  return result?.data || { data: [], total: 0, page: 1, totalPages: 0 }
}

/**
 * Get a single page by ID
 * @param {number|string} id - Page ID
 * @returns {Promise<Object>} Page data
 */
export async function getPage(id) {
  const result = await apiFetch(`/pages/${id}`)
  return result?.data || result
}

/**
 * Create a new page
 * @param {Object} pageData - Page data
 * @param {string} pageData.name - Page name (required)
 * @param {string} pageData.route - Page route (required)
 * @param {string} pageData.icon - Page icon
 * @param {string} pageData.status - Page status (active/inactive)
 * @returns {Promise<Object>} Created page data
 */
export async function createPage(pageData) {
  const result = await apiFetch('/pages', {
    method: 'POST',
    body: pageData
  })
  return result?.data || result
}

/**
 * Update an existing page
 * @param {number|string} id - Page ID
 * @param {Object} pageData - Page data to update
 * @returns {Promise<Object>} Updated page data
 */
export async function updatePage(id, pageData) {
  const result = await apiFetch(`/pages/${id}`, {
    method: 'PUT',
    body: pageData
  })
  return result?.data || result
}

/**
 * Delete a page
 * @param {number|string} id - Page ID
 * @returns {Promise<Object>} Response
 */
export async function deletePage(id) {
  const result = await apiFetch(`/pages/${id}`, {
    method: 'DELETE'
  })
  return result?.data || result
}
