import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

export async function listProjects({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT, ...otherParams } = {}) {
  const query = new URLSearchParams({ page, limit, ...otherParams }).toString()
  const result = await apiFetch(`/projects?${query}`, { method: 'GET' })
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

export async function getProject(id) {
  const result = await apiFetch(`/projects/${id}`, { method: 'GET' })
  return result?.data
}

export async function getProjectStats(params = {}) {
  const query = new URLSearchParams(params).toString()
  const suffix = query ? `?${query}` : ''
  const result = await apiFetch(`/projects/stats${suffix}`, { method: 'GET' })
  return result?.data || {}
}

export async function getProjectWorkdays(id) {
  const result = await apiFetch(`/projects/${id}/workdays`, { method: 'GET' })
  return result?.data || []
}

export async function createProject(payload) {
  const result = await apiFetch('/projects', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateProject(id, payload) {
  const result = await apiFetch(`/projects/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

