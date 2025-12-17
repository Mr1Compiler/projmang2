/**
 * Mock API Handler
 *
 * Intercepts API calls and returns mock data when MOCK_ENABLED is true.
 * Delete this file and /mocks folder when backend is ready.
 */

import {
  MOCK_ENABLED,
  mockProjects,
  mockUsers,
  mockEngineers,
  mockDebtors,
  mockExpenses,
  mockIncome,
  mockWorkdays,
  mockEmployees,
  mockTeamMembers,
  mockWorkCategories,
  mockRoles,
  mockPages,
  mockRolePages,
} from './mockData'

// Simulated delay to mimic network request (ms)
const MOCK_DELAY = 300

/**
 * Helper to simulate async API delay
 */
const delay = (ms = MOCK_DELAY) => new Promise(resolve => setTimeout(resolve, ms))

/**
 * Helper to paginate data
 */
const paginate = (data, page = 1, limit = 10) => {
  const start = (page - 1) * limit
  const end = start + limit
  const paginatedData = data.slice(start, end)

  return {
    data: paginatedData,
    page,
    limit,
    total: data.length,
    totalPages: Math.ceil(data.length / limit),
  }
}

/**
 * Helper to generate new ID
 */
const generateId = (items) => {
  const maxId = items.reduce((max, item) => Math.max(max, item.id || 0), 0)
  return maxId + 1
}

// In-memory storage for mutations (so changes persist during session)
let projectsData = [...mockProjects]
let usersData = [...mockUsers]
let engineersData = [...mockEngineers]
let debtorsData = [...mockDebtors]
let expensesData = [...mockExpenses]
let incomeData = [...mockIncome]
let workdaysData = [...mockWorkdays]
let employeesData = [...mockEmployees]
let teamMembersData = [...mockTeamMembers]
let workCategoriesData = [...mockWorkCategories]

/**
 * Mock API handlers by endpoint pattern
 */
const mockHandlers = {
  // ============================================
  // PROJECTS
  // ============================================
  'GET /api/v1/projects': async (params) => {
    await delay()
    let filtered = [...projectsData]

    if (params?.search) {
      const search = params.search.toLowerCase()
      filtered = filtered.filter(p =>
        p.name.toLowerCase().includes(search) ||
        p.description?.toLowerCase().includes(search)
      )
    }
    if (params?.status) {
      filtered = filtered.filter(p => p.status === params.status)
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'GET /api/v1/projects/:id': async (id) => {
    await delay()
    const project = projectsData.find(p => p.id === parseInt(id))
    if (!project) throw new Error('Project not found')
    return { data: project }
  },

  'POST /api/v1/projects': async (data) => {
    await delay()
    const newProject = {
      id: generateId(projectsData),
      ...data,
      progressPercentage: 0,
      status: data.status || 'pending',
      createdAt: new Date().toISOString(),
    }
    projectsData.unshift(newProject)
    return { data: newProject }
  },

  'PUT /api/v1/projects/:id': async (id, data) => {
    await delay()
    const index = projectsData.findIndex(p => p.id === parseInt(id))
    if (index === -1) throw new Error('Project not found')
    projectsData[index] = { ...projectsData[index], ...data }
    return { data: projectsData[index] }
  },

  'DELETE /api/v1/projects/:id': async (id) => {
    await delay()
    projectsData = projectsData.filter(p => p.id !== parseInt(id))
    return null
  },

  // ============================================
  // USERS
  // ============================================
  'GET /api/v1/users': async (params) => {
    await delay()
    let filtered = [...usersData]

    if (params?.search) {
      const search = params.search.toLowerCase()
      filtered = filtered.filter(u =>
        u.fullName?.toLowerCase().includes(search) ||
        u.username?.toLowerCase().includes(search) ||
        u.email?.toLowerCase().includes(search)
      )
    }
    if (params?.status) {
      filtered = filtered.filter(u => u.status === params.status)
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'GET /api/v1/users/:id': async (id) => {
    await delay()
    const user = usersData.find(u => u.id === parseInt(id))
    if (!user) throw new Error('User not found')
    return { data: user }
  },

  'POST /api/v1/users': async (data) => {
    await delay()
    const newUser = {
      id: generateId(usersData),
      ...data,
      status: 'active',
      createdAt: new Date().toISOString(),
    }
    usersData.unshift(newUser)
    return { data: newUser }
  },

  'PUT /api/v1/users/:id': async (id, data) => {
    await delay()
    const index = usersData.findIndex(u => u.id === parseInt(id))
    if (index === -1) throw new Error('User not found')
    usersData[index] = { ...usersData[index], ...data }
    return { data: usersData[index] }
  },

  'DELETE /api/v1/users/:id': async (id) => {
    await delay()
    usersData = usersData.filter(u => u.id !== parseInt(id))
    return null
  },

  // ============================================
  // DEBTORS
  // ============================================
  'GET /api/v1/debtors': async (params) => {
    await delay()
    let filtered = [...debtorsData]

    if (params?.search) {
      const search = params.search.toLowerCase()
      filtered = filtered.filter(d =>
        d.name.toLowerCase().includes(search) ||
        d.phone?.includes(search)
      )
    }
    if (params?.status) {
      filtered = filtered.filter(d => d.status === params.status)
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'GET /api/v1/debtors/:id': async (id) => {
    await delay()
    const debtor = debtorsData.find(d => d.id === parseInt(id))
    if (!debtor) throw new Error('Debtor not found')
    return { data: debtor }
  },

  'POST /api/v1/debtors': async (data) => {
    await delay()
    const newDebtor = {
      id: generateId(debtorsData),
      ...data,
      status: data.status || 'pending',
      createdAt: new Date().toISOString(),
    }
    debtorsData.unshift(newDebtor)
    return { data: newDebtor }
  },

  'PUT /api/v1/debtors/:id': async (id, data) => {
    await delay()
    const index = debtorsData.findIndex(d => d.id === parseInt(id))
    if (index === -1) throw new Error('Debtor not found')
    debtorsData[index] = { ...debtorsData[index], ...data }
    return { data: debtorsData[index] }
  },

  'DELETE /api/v1/debtors/:id': async (id) => {
    await delay()
    debtorsData = debtorsData.filter(d => d.id !== parseInt(id))
    return null
  },

  // ============================================
  // EXPENSES
  // ============================================
  'GET /api/v1/expenses': async (params) => {
    await delay()
    let filtered = [...expensesData]

    if (params?.search) {
      const search = params.search.toLowerCase()
      filtered = filtered.filter(e =>
        e.description.toLowerCase().includes(search)
      )
    }
    if (params?.projectId) {
      filtered = filtered.filter(e => e.projectId === parseInt(params.projectId))
    }
    if (params?.startDate) {
      filtered = filtered.filter(e => e.date >= params.startDate)
    }
    if (params?.endDate) {
      filtered = filtered.filter(e => e.date <= params.endDate)
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'GET /api/v1/expenses/:id': async (id) => {
    await delay()
    const expense = expensesData.find(e => e.id === parseInt(id))
    if (!expense) throw new Error('Expense not found')
    return { data: expense }
  },

  'POST /api/v1/expenses': async (data) => {
    await delay()
    const project = projectsData.find(p => p.id === data.projectId)
    const newExpense = {
      id: generateId(expensesData),
      ...data,
      projectName: project?.name || '',
      createdAt: new Date().toISOString(),
    }
    expensesData.unshift(newExpense)
    return { data: newExpense }
  },

  'PUT /api/v1/expenses/:id': async (id, data) => {
    await delay()
    const index = expensesData.findIndex(e => e.id === parseInt(id))
    if (index === -1) throw new Error('Expense not found')
    expensesData[index] = { ...expensesData[index], ...data }
    return { data: expensesData[index] }
  },

  'DELETE /api/v1/expenses/:id': async (id) => {
    await delay()
    expensesData = expensesData.filter(e => e.id !== parseInt(id))
    return null
  },

  // ============================================
  // INCOME
  // ============================================
  'GET /api/v1/income': async (params) => {
    await delay()
    let filtered = [...incomeData]

    if (params?.projectId) {
      filtered = filtered.filter(i => i.projectId === parseInt(params.projectId))
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'POST /api/v1/income': async (data) => {
    await delay()
    const project = projectsData.find(p => p.id === data.projectId)
    const newIncome = {
      id: generateId(incomeData),
      ...data,
      projectName: project?.name || '',
      createdAt: new Date().toISOString(),
    }
    incomeData.unshift(newIncome)
    return { data: newIncome }
  },

  // ============================================
  // WORKDAYS
  // ============================================
  'GET /api/v1/workdays': async (params) => {
    await delay()
    let filtered = [...workdaysData]

    if (params?.projectId) {
      filtered = filtered.filter(w => w.projectId === parseInt(params.projectId))
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'POST /api/v1/workdays': async (data) => {
    await delay()
    const newWorkday = {
      id: generateId(workdaysData),
      ...data,
      materials: [],
      labor: [],
      equipment: [],
      status: 'pending',
      createdAt: new Date().toISOString(),
    }
    workdaysData.unshift(newWorkday)
    return { data: newWorkday }
  },

  // ============================================
  // TEAM MEMBERS
  // ============================================
  'GET /api/v1/team-members': async (params) => {
    await delay()
    let filtered = [...teamMembersData]

    if (params?.projectId) {
      filtered = filtered.filter(tm => tm.projectId === parseInt(params.projectId))
    }
    if (params?.userId) {
      filtered = filtered.filter(tm => tm.userId === parseInt(params.userId))
    }

    return paginate(filtered, params?.page, params?.limit)
  },

  'POST /api/v1/team-members': async (data) => {
    await delay()
    const newTeamMember = {
      id: generateId(teamMembersData),
      ...data,
      createdAt: new Date().toISOString(),
    }
    teamMembersData.push(newTeamMember)
    return { data: newTeamMember }
  },

  'DELETE /api/v1/team-members/:id': async (id) => {
    await delay()
    teamMembersData = teamMembersData.filter(tm => tm.id !== parseInt(id))
    return null
  },

  // ============================================
  // WORK CATEGORIES
  // ============================================
  'GET /api/v1/work-categories': async (params) => {
    await delay()
    return paginate(workCategoriesData, params?.page, params?.limit)
  },

  // ============================================
  // AUTH (Mock login)
  // ============================================
  'POST /api/v1/auth/login': async (data) => {
    await delay()
    const user = usersData.find(u =>
      u.username === data.username || u.email === data.username
    )
    if (!user) throw new Error('Invalid credentials')

    return {
      accessToken: 'mock-access-token-' + Date.now(),
      refreshToken: 'mock-refresh-token-' + Date.now(),
      user: {
        id: user.id,
        username: user.username,
        fullName: user.fullName,
        email: user.email,
        roles: user.roles || ['user'],
      },
    }
  },

  'GET /api/v1/auth/pages': async () => {
    await delay()
    // Return all pages with full permissions for mock mode
    return {
      data: mockPages.map(page => ({
        ...page,
        canRead: true,
        canWrite: true,
        canDelete: true,
      })),
    }
  },

  'POST /api/v1/auth/logout': async () => {
    await delay()
    return { message: 'Logged out successfully' }
  },

  'POST /api/v1/auth/refresh': async () => {
    await delay()
    return {
      accessToken: 'mock-access-token-' + Date.now(),
      refreshToken: 'mock-refresh-token-' + Date.now(),
    }
  },

  // ============================================
  // ROLES
  // ============================================
  'GET /api/v1/roles': async (params) => {
    await delay()
    return paginate(mockRoles, params?.page, params?.limit)
  },

  // ============================================
  // PAGES
  // ============================================
  'GET /api/v1/pages': async (params) => {
    await delay()
    return paginate(mockPages, params?.page, params?.limit)
  },
}

/**
 * Match endpoint to handler
 */
const matchEndpoint = (method, endpoint) => {
  const key = `${method} ${endpoint}`

  // Exact match
  if (mockHandlers[key]) {
    return { handler: mockHandlers[key], params: {} }
  }

  // Pattern match (e.g., /api/v1/projects/:id)
  for (const pattern of Object.keys(mockHandlers)) {
    const [patternMethod, patternPath] = pattern.split(' ')
    if (patternMethod !== method) continue

    const patternParts = patternPath.split('/')
    const endpointParts = endpoint.split('?')[0].split('/')

    if (patternParts.length !== endpointParts.length) continue

    let match = true
    const params = {}

    for (let i = 0; i < patternParts.length; i++) {
      if (patternParts[i].startsWith(':')) {
        params[patternParts[i].slice(1)] = endpointParts[i]
      } else if (patternParts[i] !== endpointParts[i]) {
        match = false
        break
      }
    }

    if (match) {
      return { handler: mockHandlers[pattern], params }
    }
  }

  return null
}

/**
 * Parse query params from URL
 */
const parseQueryParams = (endpoint) => {
  const [, queryString] = endpoint.split('?')
  if (!queryString) return {}

  const params = {}
  new URLSearchParams(queryString).forEach((value, key) => {
    params[key] = value
  })
  return params
}

/**
 * Main mock request handler
 * Call this from the API client when MOCK_ENABLED is true
 */
export const mockRequest = async (method, endpoint, data = null) => {
  if (!MOCK_ENABLED) {
    throw new Error('Mock mode is disabled')
  }

  const match = matchEndpoint(method.toUpperCase(), endpoint)

  if (!match) {
    console.warn(`[Mock API] No handler for: ${method} ${endpoint}`)
    await delay()
    return { data: [], page: 1, limit: 10, total: 0, totalPages: 0 }
  }

  const queryParams = parseQueryParams(endpoint)
  const { handler, params } = match

  try {
    // Call handler with appropriate args
    if (params.id) {
      return await handler(params.id, data || queryParams)
    }
    return await handler(data || queryParams)
  } catch (error) {
    console.error(`[Mock API] Error: ${error.message}`)
    throw error
  }
}

/**
 * Check if mock mode is enabled
 */
export const isMockEnabled = () => MOCK_ENABLED

/**
 * Reset all mock data to initial state
 */
export const resetMockData = () => {
  projectsData = [...mockProjects]
  usersData = [...mockUsers]
  engineersData = [...mockEngineers]
  debtorsData = [...mockDebtors]
  expensesData = [...mockExpenses]
  incomeData = [...mockIncome]
  workdaysData = [...mockWorkdays]
  employeesData = [...mockEmployees]
  teamMembersData = [...mockTeamMembers]
  workCategoriesData = [...mockWorkCategories]
}

export default { mockRequest, isMockEnabled, resetMockData }
