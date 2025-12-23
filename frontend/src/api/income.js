import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

// ============ Income ============

export async function listIncome({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/income?${query}`, { method: 'GET' })
  // Return full pagination response
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

export async function getIncomeById(id) {
  const result = await apiFetch(`/income/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function getIncomeStats() {
  const result = await apiFetch('/income/stats', { method: 'GET' })
  return result?.data || result
}

export async function createIncome(payload) {
  const result = await apiFetch('/income', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateIncome(id, payload) {
  const result = await apiFetch(`/income/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteIncome(id) {
  const result = await apiFetch(`/income/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
