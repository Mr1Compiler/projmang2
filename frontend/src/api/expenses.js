import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

// ============ Expenses ============

export async function listExpenses({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/expenses?${query}`, { method: 'GET' })
  // Return full pagination response
  const paginatedData = result?.data || {}
  return {
    data: paginatedData.data || [],
    total: paginatedData.total || 0,
    page: paginatedData.page || page,
    limit: paginatedData.limit || limit,
    totalPages: paginatedData.totalPages || 0
  }
}

export async function getExpenseById(id) {
  const result = await apiFetch(`/expenses/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function getExpenseStats() {
  const result = await apiFetch('/expenses/stats', { method: 'GET' })
  return result?.data || result
}

export async function createExpense(payload) {
  const result = await apiFetch('/expenses', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateExpense(id, payload) {
  const result = await apiFetch(`/expenses/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteExpense(id) {
  const result = await apiFetch(`/expenses/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

// ============ Project Expenses ============

export async function listExpensesByProject(projectId, { page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/projects/${projectId}/expenses?${query}`, { method: 'GET' })
  // Return full pagination response
  const paginatedData = result?.data || {}
  if (Array.isArray(paginatedData)) {
    return { data: paginatedData, total: paginatedData.length, page, limit, totalPages: 1 }
  }
  return {
    data: paginatedData.data || paginatedData || [],
    total: paginatedData.total || 0,
    page: paginatedData.page || page,
    limit: paginatedData.limit || limit,
    totalPages: paginatedData.totalPages || 0
  }
}
