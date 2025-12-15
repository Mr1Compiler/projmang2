# AI Agent Rules - Frontend

**IMPORTANT**: Read this file before making ANY changes to the frontend.

---

## Project Structure (DO NOT CHANGE)

```
src/
├── api/           # API calls only - no business logic
├── stores/        # Pinia stores only
├── composables/   # Reusable logic (useXxx.js)
├── components/    # Reusable Vue components
├── pages/         # Route pages only
├── utils/         # Pure utility functions
├── config/        # Environment config
└── styles/        # Global CSS only
```

---

## Critical Rules

### 1. API Layer
- Base URL: `/api/v1/` (NOT `/api/`)
- All API calls go through `api/client.js`
- NO direct `fetch()` calls in components
- NO hardcoded URLs

### 2. Authentication
- Tokens stored in `localStorage`
- Use `stores/auth.js` for auth state
- Sidebar is DYNAMIC (from `/api/v1/auth/pages`)
- NO hardcoded menu items in sidebar

### 3. State Management
- Use Pinia stores for shared state
- NO duplicate data across components
- Stores location: `src/stores/`

### 4. Components
- Max 300-500 lines per file
- Use `<script setup>` syntax
- Props must be defined with types
- Emits must be declared

### 5. Styling
- NO `alert()` - use `vue3-toastify`
- NO `console.log` in production code
- Avoid `!important` in CSS
- Use scoped styles

### 6. File Naming
- Components: `PascalCase.vue`
- Composables: `useXxx.js`
- Stores: `lowercase.js`
- API modules: `lowercase.js`

### 7. Pinia Stores (CRITICAL)
- Use **Setup Stores** (Composition API syntax), NOT Options API
- ALWAYS use `storeToRefs()` when destructuring state/getters
- Actions can be destructured directly (no storeToRefs needed)
- NO circular dependencies between stores

```js
// ✅ CORRECT
import { storeToRefs } from 'pinia'
const store = useProjectsStore()
const { projects, loading } = storeToRefs(store)
const { fetchProjects } = store

// ❌ WRONG - breaks reactivity
const { projects, loading } = store
```

### 8. Composables
- Return plain objects containing refs (allows destructuring)
- Can nest other composables
- Naming: `useXxx.js`

```js
// ✅ CORRECT - return plain object with refs
export function useProjects() {
  const store = useProjectsStore()
  const { projects } = storeToRefs(store)
  return { projects, fetchProjects: store.fetchProjects }
}
```

### 9. Error Handling
- Use `toast.error(message)` / `toast.success(message)` from vue3-toastify
- Use try/catch in all API calls
- NO `console.log` - use `console.error` only in development

```js
// ✅ CORRECT
try {
  await api.createProject(data)
  toast.success('Project created')
} catch (error) {
  toast.error(error.message || 'Failed to create project')
}
```

---

## DO NOT

- Create new folders outside the structure above
- Add TypeScript (project uses JavaScript)
- Change API base URL structure
- Hardcode sidebar menu items
- Put business logic in components
- Create files > 500 lines

---

## Before Making Changes

1. Check `refactoring-plan.md` for current phase
2. Follow the established patterns
3. Keep files small and focused
4. Test after each change

---

**Last Updated**: 2025-12-15 (v2 - Added Pinia, composables, error handling rules)
