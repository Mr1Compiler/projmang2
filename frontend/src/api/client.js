import { toast } from 'vue3-toastify'
import { config } from '@/config/env'
import { mockRequest, isMockEnabled } from '@/mocks/mockApi'

const TOKEN_KEY = 'accessToken'
const REFRESH_TOKEN_KEY = 'refreshToken'

class ApiClient {
  constructor() {
    this.baseURL = config.apiBaseUrl
    this.timeout = config.apiTimeout
  }

  /**
   * Check if mock mode is enabled
   */
  useMock() {
    return isMockEnabled()
  }

  getToken() {
    return localStorage.getItem(TOKEN_KEY)
  }

  setTokens(accessToken, refreshToken) {
    localStorage.setItem(TOKEN_KEY, accessToken)
    if (refreshToken) {
      localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken)
    }
  }

  clearTokens() {
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem(REFRESH_TOKEN_KEY)
  }

  getRefreshToken() {
    return localStorage.getItem(REFRESH_TOKEN_KEY)
  }

  async request(endpoint, options = {}) {
    // Use mock data if enabled
    if (this.useMock()) {
      const method = options.method || 'GET'
      const body = options.body ? JSON.parse(options.body) : null
      return mockRequest(method, endpoint, body)
    }

    const url = `${this.baseURL}${endpoint}`
    const token = this.getToken()

    const requestConfig = {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...(token && { Authorization: `Bearer ${token}` }),
        ...options.headers,
      },
      signal: AbortSignal.timeout(this.timeout),
    }

    try {
      const response = await fetch(url, requestConfig)

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}))
        const errorMessage = errorData.message || `HTTP error! status: ${response.status}`

        // Handle 401 Unauthorized - clear tokens and redirect to login
        if (response.status === 401) {
          this.clearTokens()
          // Only redirect if not already on login page
          if (!window.location.pathname.includes('login')) {
            window.location.href = '/login'
          }
        }

        throw new Error(errorMessage)
      }

      // Handle 204 No Content
      if (response.status === 204) {
        return null
      }

      return await response.json()
    } catch (error) {
      // Don't show toast for aborted requests
      if (error.name !== 'AbortError') {
        toast.error(error.message || 'Request failed')
      }
      throw error
    }
  }

  get(endpoint, options = {}) {
    return this.request(endpoint, { ...options, method: 'GET' })
  }

  post(endpoint, data, options = {}) {
    return this.request(endpoint, {
      ...options,
      method: 'POST',
      body: JSON.stringify(data),
    })
  }

  put(endpoint, data, options = {}) {
    return this.request(endpoint, {
      ...options,
      method: 'PUT',
      body: JSON.stringify(data),
    })
  }

  patch(endpoint, data, options = {}) {
    return this.request(endpoint, {
      ...options,
      method: 'PATCH',
      body: JSON.stringify(data),
    })
  }

  delete(endpoint, options = {}) {
    return this.request(endpoint, { ...options, method: 'DELETE' })
  }
}

const apiClient = new ApiClient()

export default apiClient
export { apiClient, TOKEN_KEY, REFRESH_TOKEN_KEY }
