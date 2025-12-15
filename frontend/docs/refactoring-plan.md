# Frontend Refactoring Plan

## Current Situation

### Code Quality Issues
- ❌ **16 `alert()` calls** - Should use toast notifications
- ❌ **75 `console.log` statements** - Debug code in production
- ❌ **No Pinia stores** - Empty `stores/` folder, no state management
- ❌ **No error handling strategy** - Inconsistent error feedback
- ❌ **No tests** - No unit or integration tests

### File Size Analysis
- `project-management.vue`: **9,158 lines** ⚠️
- `engineers.vue`: **4,104 lines** ⚠️
- `expenses.vue`: **3,575 lines** ⚠️
- `debtors.vue`: **2,241 lines** ⚠️
- `human-resources.vue`: **3,568 lines** ⚠️
- `users.vue`: **2,983 lines** ⚠️

### Problems
- ❌ Files too large (best practice: 300-500 lines max)
- ❌ Hard to maintain and navigate
- ❌ Slow IDE performance
- ❌ Difficult to test
- ❌ Merge conflicts
- ❌ Poor code reusability

---

## Proposed Structure

### 1. API Layer (`src/api/`)
```
src/api/
├── client.js          # Base HTTP client (base URL: /api/v1/)
├── auth.js            # Login, logout, refresh, getPages
├── projects.js        # Project endpoints
├── debtors.js         # Debtor endpoints
├── expenses.js        # Expense endpoints
├── income.js          # Income endpoints
├── users.js           # User endpoints
├── teams.js           # Team member endpoints
├── workdays.js        # Workday + labor + materials + equipment
├── roles.js           # Roles & role-pages (RBAC)
└── categories.js      # Work categories & subcategories
```

**Notes**:
- Base URL: `/api/v1/` (not `/api/`)
- Engineers: No separate endpoint - use `users.js` + `workdays.js`

**Purpose**: Centralize all API calls, use env variables for base URL

**API Client with Error Handling (`client.js`):**
```js
import { toast } from 'vue3-toastify'

const BASE_URL = import.meta.env.VITE_API_URL || '/api/v1'

async function request(endpoint, options = {}) {
  const token = localStorage.getItem('accessToken')

  const config = {
    headers: {
      'Content-Type': 'application/json',
      ...(token && { Authorization: `Bearer ${token}` }),
      ...options.headers,
    },
    ...options,
  }

  try {
    const response = await fetch(`${BASE_URL}${endpoint}`, config)

    if (!response.ok) {
      const error = await response.json().catch(() => ({}))
      throw new Error(error.message || `HTTP ${response.status}`)
    }

    return response.json()
  } catch (error) {
    toast.error(error.message || 'Request failed')
    throw error
  }
}

export const api = {
  get: (endpoint) => request(endpoint),
  post: (endpoint, data) => request(endpoint, { method: 'POST', body: JSON.stringify(data) }),
  put: (endpoint, data) => request(endpoint, { method: 'PUT', body: JSON.stringify(data) }),
  delete: (endpoint) => request(endpoint, { method: 'DELETE' }),
}
```

**API Module Pattern (returns loading/error states):**
```js
// api/projects.js
import { api } from './client'

export const projectsApi = {
  getAll: () => api.get('/projects'),
  getById: (id) => api.get(`/projects/${id}`),
  create: (data) => api.post('/projects', data),
  update: (id, data) => api.put(`/projects/${id}`, data),
  delete: (id) => api.delete(`/projects/${id}`),
}
```

### 2. Stores (`src/stores/`) - Pinia Setup Stores
```
src/stores/
├── auth.js            # Token, user, pages, permissions
├── projects.js        # Projects state
├── debtors.js         # Debtors state
├── expenses.js        # Expenses state
├── users.js           # Users state
└── ui.js              # Sidebar state, loading states
```

**Purpose**: Centralized state management with Pinia

**IMPORTANT - Use Setup Stores (Composition API syntax):**
```js
// ✅ Setup Store (USE THIS)
import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useProjectsStore = defineStore('projects', () => {
  const projects = ref([])
  const loading = ref(false)
  const error = ref(null)

  const projectCount = computed(() => projects.value.length)

  async function fetchProjects() {
    loading.value = true
    try {
      projects.value = await projectsApi.getAll()
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  return { projects, loading, error, projectCount, fetchProjects }
})
```

**Using stores in components - ALWAYS use `storeToRefs()`:**
```js
import { storeToRefs } from 'pinia'
import { useProjectsStore } from '@/stores/projects'

const store = useProjectsStore()
const { projects, loading } = storeToRefs(store) // ✅ Keeps reactivity
const { fetchProjects } = store // Actions don't need storeToRefs
```

### 3. Composables (`src/composables/`)
```
src/composables/
├── useAuth.js         # Login/logout logic
├── usePermissions.js  # Check read/write/delete permissions
├── useProjects.js     # Project state & logic
├── useDebtors.js      # Debtor state & logic
├── useExpenses.js     # Expense state & logic
├── useTeamManagement.js  # Team logic
├── useTasks.js        # Task logic
└── useUsers.js        # User logic
```

**Purpose**: Reusable reactive logic, separate from UI

### 4. Components (`src/components/`)
```
src/components/
├── projects/
│   ├── ProjectHeader.vue      # Header section
│   ├── ProjectStats.vue        # Statistics cards
│   ├── ProjectTable.vue        # Data table
│   ├── ProjectForm.vue         # Add/Edit dialog
│   ├── ProjectFilters.vue      # Search & filters
│   ├── TeamManagement.vue      # Team section
│   ├── ExpenseManagement.vue   # Expense section
│   └── TaskList.vue            # Task management
├── debtors/
│   ├── DebtorTable.vue
│   ├── DebtorForm.vue
│   └── DebtorStats.vue
└── engineers/
    ├── EngineerTable.vue
    ├── EngineerForm.vue
    └── EngineerStats.vue
```

**Purpose**: Break down large pages into smaller, reusable components

### 5. Utils (`src/utils/`)
```
src/utils/
├── formatters.js      # Currency, date, number formatting
├── validators.js      # Form validation helpers
└── constants.js       # Constants & enums
```

**Purpose**: Shared utility functions

### 6. Authentication & Authorization (RBAC)

**Login Flow:**
1. User logs in → `POST /api/v1/auth/login`
2. Store `accessToken` + `refreshToken` in localStorage
3. Fetch user pages → `GET /api/v1/auth/pages`
4. Store pages in `stores/auth.js`
5. Sidebar shows only allowed pages

**Dynamic Sidebar:**
- Sidebar reads from `auth store` (not hardcoded)
- Each menu item comes from `/api/v1/auth/pages` response
- Shows only pages user has access to

**Permission Checks:**
- `usePermissions.js` composable checks if user can:
  - `canRead(page)` → show page
  - `canWrite(page)` → show add/edit buttons
  - `canDelete(page)` → show delete button

**Route Guards:**
- Vue Router guards check if user can access route
- Redirect to login if no token
- Redirect to 403 if no permission

---

## Refactoring Strategies

### Option A: Big Bang Approach
**Description**: Refactor everything at once

**Pros**:
- ✅ Consistent structure throughout
- ✅ No mixed patterns
- ✅ Complete transformation

**Cons**:
- ❌ High risk
- ❌ Long development time
- ❌ All features broken during refactor
- ❌ Difficult to test incrementally

**Best for**: Small projects or when you can afford downtime

---

### Option B: Incremental Approach ⭐ **RECOMMENDED**
**Description**: Refactor one section at a time, test after each

**Pros**:
- ✅ Lower risk
- ✅ Continuous progress
- ✅ Can test after each step
- ✅ Less disruptive
- ✅ Can pause/resume anytime

**Cons**:
- ❌ Temporary mixed patterns
- ❌ Takes longer overall
- ❌ Need discipline to finish

**Best for**: Production apps, large codebases

**Steps**:
1. Extract API layer first
2. Extract one component at a time
3. Move logic to composables
4. Test after each extraction
5. Repeat for next section

---

### Option C: Feature-Based Approach
**Description**: Refactor complete features (projects → debtors → engineers)

**Pros**:
- ✅ Complete features at once
- ✅ Can deploy feature by feature
- ✅ Clear milestones

**Cons**:
- ❌ Each feature takes longer
- ❌ Mixed patterns between features
- ❌ Harder to see overall progress

**Best for**: When features are independent

---

## Recommended Approach: Incremental (Option B)

### Phase 0: Quick Cleanup (Before Refactoring) ✅ COMPLETED
**Goal**: Clean up code quality issues first
**Status**: ✅ Completed on 2025-12-15

#### 1. ✅ Replace `alert()` with Toast Notifications
- Replaced all 16 `alert()` calls with `toast.success()` / `toast.error()` / `toast.warning()`
- Files fixed: `human-resources.vue`, `project-expenses.vue`, `work-day-details.vue`, `expense-types.vue`

#### 2. ✅ Remove `console.log` Statements
- Removed 74 `console.log` statements across 22 files
- Kept 1 in `router/index.js` (legitimate error handling)

#### 3. ✅ Setup Pinia Stores
- Created `stores/auth.js` with login/logout, token management, and permissions
- Uses Setup Store syntax (Composition API) as per AI_RULES.md

#### 4. ✅ Configure API Layer with Error Handling
- Updated `api/client.js` with auth token injection, toast errors, 401 handling
- Created `api/projects.js` with full CRUD + nested resources

#### 5. ✅ Setup Utils
- Created `utils/formatters.js` with currency, date, number formatting functions

#### ⏳ Pending (Optional for Phase 0)
- Setup Testing (Vitest) - deferred to later phase

**Deliverable**: ✅ Clean codebase ready for refactoring

---

### Phase 1: Foundation ✅ COMPLETED
**Goal**: Set up infrastructure
**Status**: ✅ Completed on 2025-12-15

#### 1. ✅ Create `api/` folder structure
- `api/client.js` - Base HTTP client with auth, error handling, 401 redirect
- `api/auth.js` - Login, logout, refresh, getPages
- `api/projects.js` - Projects CRUD + workdays, expenses, team members
- `api/debtors.js` - Debtors CRUD
- `api/expenses.js` - Expenses CRUD with filters
- `api/income.js` - Income CRUD with filters
- `api/users.js` - Users CRUD + password change, status update, roles
- `api/teams.js` - Team members management
- `api/workdays.js` - Workdays + materials + labor + equipment
- `api/roles.js` - Roles + pages + role-pages (RBAC)
- `api/categories.js` - Work categories & subcategories
- `api/index.js` - Barrel export

#### 2. ✅ Create `stores/` folder (Pinia Setup Stores)
- `stores/auth.js` - Auth state, tokens, permissions (from Phase 0)
- `stores/projects.js` - Projects state & CRUD
- `stores/debtors.js` - Debtors state & CRUD
- `stores/expenses.js` - Expenses state & CRUD with filters
- `stores/users.js` - Users state & CRUD
- `stores/ui.js` - Sidebar, loading, modals, dark mode, breadcrumbs
- `stores/index.js` - Barrel export

#### 3. ✅ Create `composables/` folder
- `composables/usePermissions.js` - Permission checking (canRead, canWrite, canDelete)
- `composables/useProjects.js` - Projects logic with permission checks
- `composables/index.js` - Barrel export

#### 4. ✅ Complete `utils/` folder
- `utils/formatters.js` - Currency, date, number formatting (from Phase 0)
- `utils/validators.js` - Form validation helpers (required, email, phone, etc.)
- `utils/constants.js` - Constants & enums (statuses, routes, modals, etc.)
- `utils/index.js` - Barrel export

**Deliverable**: ✅ Infrastructure ready, no page changes yet

---

### Phase 2: Extract Components ✅ COMPLETED
**Goal**: Break down `project-management.vue`
**Status**: ✅ Completed on 2025-12-15

**Created Components** (`src/components/projects/`):
1. ✅ **ProjectStats.vue** - Statistics cards (5 cards: total, active, pending, budget, progress)
2. ✅ **ProjectCard.vue** - Individual project card with status, details, progress bar
3. ✅ **ProjectFilters.vue** - Search, status filter, sort options
4. ✅ **ProjectForm.vue** - Add/Edit project dialog (v-model compatible)
5. ✅ **TeamManagement.vue** - Expandable team section with stats, table, add member dialog
6. ✅ **DeleteConfirmDialog.vue** - Reusable delete confirmation dialog
7. ✅ **index.js** - Barrel export for all components

**Component Features**:
- All components use `<script setup>` syntax (Composition API)
- Props and emits properly defined
- Scoped styles with responsive breakpoints
- RTL (Arabic) support maintained
- Original UI/UX styling preserved

**Usage Example**:
```vue
import {
  ProjectStats,
  ProjectCard,
  ProjectFilters,
  ProjectForm,
  TeamManagement,
  DeleteConfirmDialog
} from '@/components/projects'
```

**Deliverable**: ✅ Components ready for integration

---

### Phase 3: Move Logic to Composables
**Goal**: Extract business logic

**Extract**:
- State management (refs, reactive)
- Computed properties
- Methods/functions
- API calls

**Create**:
- `useProjects.js` - Project logic
- `useTeamManagement.js` - Team logic
- `useExpenses.js` - Expense logic
- `useTasks.js` - Task logic

**Deliverable**: `project-management.vue` reduced to ~200-300 lines

---

### Phase 4: Repeat for Other Pages
**Priority**:
1. `engineers.vue` (4,104 lines)
2. `expenses.vue` (3,575 lines)
3. `debtors.vue` (2,241 lines)
4. `human-resources.vue` (3,568 lines)
5. `users.vue` (2,983 lines)

**Process**: Same as Phase 2-3

---

## Expected Results

### Before
```
project-management.vue: 9,158 lines
engineers.vue: 4,104 lines
expenses.vue: 3,575 lines
```

### After
```
project-management.vue: ~200-300 lines
  ├── Uses: ProjectHeader, ProjectStats, ProjectTable, etc.
  ├── Uses: useProjects, useTeamManagement composables
  └── Uses: projectsApi from api/

components/projects/: ~8 components (200-400 lines each)
composables/: ~6 composables (100-200 lines each)
api/: ~6 modules (50-100 lines each)
```

---

## Benefits

### Code Quality
- ✅ Smaller, focused files
- ✅ Single Responsibility Principle
- ✅ Easier to understand
- ✅ Better code organization

### Maintainability
- ✅ Easy to find code
- ✅ Easy to modify
- ✅ Less merge conflicts
- ✅ Better git history

### Reusability
- ✅ Components can be reused
- ✅ Composables shared across pages
- ✅ Utils available everywhere

### Testing
- ✅ Components testable in isolation
- ✅ Composables testable separately
- ✅ API layer mockable

### Performance
- ✅ Better code splitting
- ✅ Lazy loading components
- ✅ Faster IDE performance

---

## Migration Checklist

### For Each Page Refactor:

- [ ] Create API module in `api/`
- [ ] Create composable in `composables/`
- [ ] Extract components to `components/`
- [ ] Update main page to use new structure
- [ ] Test all functionality
- [ ] Remove old code
- [ ] Update imports
- [ ] Verify no regressions
- [ ] Update documentation

---

## Questions to Consider

1. **Which approach?** Incremental (recommended) vs Big Bang vs Feature-based
2. **Start with?** `project-management.vue` (biggest) or smaller file first?
3. **Keep functionality?** Maintain all features during refactor?
4. **Add tests?** Write tests during refactor or after?
5. **Approach?** Quick wins or thorough refactor?
6. **Team size?** Solo or team (affects merge strategy)

---

## Next Steps

1. **Decide on approach** (recommend: Incremental)
2. **Start with Phase 1** (Foundation)
3. **Pick first component** to extract (recommend: ProjectStats)
4. **Test after each step**
5. **Iterate and improve**

---

## Notes

- Keep existing functionality working during refactor
- Test after each extraction
- Use git branches for each phase
- Document component props/emits
- Follow Vue 3 Composition API best practices

---

**Last Updated**: 2025-12-15 (v5 - Phase 2 completed)
**Status**: Phase 0 ✅ Complete | Phase 1 ✅ Complete | Phase 2 ✅ Complete | Phase 3 Ready to Start
