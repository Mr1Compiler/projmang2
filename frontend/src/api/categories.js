import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

// ============ Work Categories ============

export async function listCategories({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/workCategories?${query}`, { method: 'GET' })
  const paginatedData = result?.data || {}
  if (Array.isArray(paginatedData)) {
    return { data: paginatedData, total: paginatedData.length, page, limit, totalPages: 1 }
  }
  return {
    data: paginatedData.data || [],
    total: paginatedData.total || 0,
    page: paginatedData.page || page,
    limit: paginatedData.limit || limit,
    totalPages: paginatedData.totalPages || 0
  }
}

export async function getCategory(id) {
  const result = await apiFetch(`/workCategories/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createCategory(payload) {
  const result = await apiFetch(`/workCategories`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateCategory(id, payload) {
  const result = await apiFetch(`/workCategories/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteCategory(id) {
  const result = await apiFetch(`/workCategories/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

export async function getCategoryStats({ period = '' } = {}) {
  const query = period ? `?period=${period}` : ''
  const result = await apiFetch(`/workCategories/stats${query}`, { method: 'GET' })
  return result?.data || result
}

// ============ Work SubCategories ============

export async function listSubCategories({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/workSubcategories?${query}`, { method: 'GET' })
  const paginatedData = result?.data || {}
  if (Array.isArray(paginatedData)) {
    return { data: paginatedData, total: paginatedData.length, page, limit, totalPages: 1 }
  }
  return {
    data: paginatedData.data || [],
    total: paginatedData.total || 0,
    page: paginatedData.page || page,
    limit: paginatedData.limit || limit,
    totalPages: paginatedData.totalPages || 0
  }
}

export async function getSubCategory(id) {
  const result = await apiFetch(`/workSubcategories/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createSubCategory(payload) {
  const result = await apiFetch(`/workSubcategories`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateSubCategory(id, payload) {
  const result = await apiFetch(`/workSubcategories/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteSubCategory(id) {
  const result = await apiFetch(`/workSubcategories/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
