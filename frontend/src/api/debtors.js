import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

// ============ Debtors ============

export async function listDebtors({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/debtors?${query}`, { method: 'GET' })
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

export async function getDebtorById(id) {
  const result = await apiFetch(`/debtors/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function getDebtorStats() {
  const result = await apiFetch('/debtors/stats', { method: 'GET' })
  return result?.data || result
}

export async function createDebtor(payload) {
  const result = await apiFetch('/debtors', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateDebtor(id, payload) {
  const result = await apiFetch(`/debtors/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteDebtor(id) {
  const result = await apiFetch(`/debtors/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
