import { apiFetch } from './client'

// ============ WorkDay Materials ============

export async function listMaterialsByWorkDay(workDayId) {
  const result = await apiFetch(`/workdays/${workDayId}/materials`, { method: 'GET' })
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function createMaterial(payload) {
  const result = await apiFetch(`/workdayMaterials`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateMaterial(id, payload) {
  const result = await apiFetch(`/workdayMaterials/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteMaterial(id) {
  const result = await apiFetch(`/workdayMaterials/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

// ============ WorkDay Labor (Daily Expenses) ============

export async function listLaborByWorkDay(workDayId) {
  const result = await apiFetch(`/workdays/${workDayId}/labor`, { method: 'GET' })
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function createLabor(payload) {
  const result = await apiFetch(`/workdayLabor`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateLabor(id, payload) {
  const result = await apiFetch(`/workdayLabor/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteLabor(id) {
  const result = await apiFetch(`/workdayLabor/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

// ============ WorkDay Equipment ============

export async function listEquipmentByWorkDay(workDayId) {
  const result = await apiFetch(`/workdays/${workDayId}/equipment`, { method: 'GET' })
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function createEquipment(payload) {
  const result = await apiFetch(`/workdayEquipment`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateEquipment(id, payload) {
  const result = await apiFetch(`/workdayEquipment/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteEquipment(id) {
  const result = await apiFetch(`/workdayEquipment/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
