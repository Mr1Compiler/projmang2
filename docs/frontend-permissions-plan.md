# Frontend Permissions Implementation Plan

## Overview

Implement role-based access control (RBAC) in the frontend by fetching user permissions from backend and conditionally rendering UI elements.

---

## Tasks

### 1. Create Pinia Auth Store
**File:** `frontend/src/stores/auth.js` (NEW)

**Purpose:** Centralized state for user, tokens, and permissions

**State:**
- `user` - current user info
- `accessToken` / `refreshToken`
- `allowedPages` - array of pages user can access
- `permissions` - map of page → permissions array

**Actions:**
- `login()` - authenticate and fetch permissions
- `logout()` - clear state and tokens
- `fetchPermissions()` - call `/auth/pages` API

---

### 2. Create Permissions Composable
**File:** `frontend/src/composables/usePermissions.js` (NEW)

**Purpose:** Helper functions for permission checks

**Functions:**
- `hasPermission(page, action)` - check if user has specific permission
- `canRead(page)` - shorthand for read check
- `canWrite(page)` - shorthand for write check
- `canDelete(page)` - shorthand for delete check
- `canUpdatePassword(page)` - shorthand for password update check
- `canUpdateStatus(page)` - shorthand for status update check

---

### 3. Modify Login Flow
**File:** `frontend/src/pages/login.vue` (MODIFY)

**Changes:**
- After successful login → call `GET /api/v1/auth/pages`
- Store returned pages and permissions in Pinia store
- Then redirect to dashboard

---

### 4. Filter Sidebar Dynamically
**File:** `frontend/src/App.vue` (MODIFY)

**Changes:**
- Import auth store
- Filter `mainMenuItems` against `allowedPages`
- Only render menu items user has access to

---

### 5. Add Route Guards
**File:** `frontend/src/router/index.js` (MODIFY)

**Changes:**
- Add `meta: { page: 'pageName' }` to each route
- In `beforeEach` guard:
  - Check if page is in user's `allowedPages`
  - If not allowed → redirect to dashboard or 403 page

---

### 6. Add Permission Checks to Page Components
**Files:** All page components (MODIFY)

**Changes per page:**
- Import `usePermissions` composable
- Add `v-if` checks on buttons:
  - **Add/Create button:** `v-if="canWrite('pageName')"`
  - **Edit button/icon:** `v-if="canWrite('pageName')"`
  - **Delete button/icon:** `v-if="canDelete('pageName')"`

---

## Permission Mapping

| Permission | HTTP Method | UI Element |
|------------|-------------|------------|
| `read` | GET | View page, View icon |
| `write` | POST, PUT | Add button, Edit icon |
| `delete` | DELETE | Delete icon |
| `updatePassword` | PUT | Change password button (Users page only) |
| `updateStatus` | PUT | Activate/Deactivate toggle (Users page only) |

---

## UI Behavior by Role

| Role | Sidebar | Table Actions |
|------|---------|---------------|
| Viewer | Limited pages | View only |
| Editor | More pages | View + Edit + Add |
| Admin | All pages | View + Edit + Add + Delete |

---

## API Endpoint

```
GET /api/v1/auth/pages

Response:
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Projects",
      "route": "/projects",
      "icon": "mdi-folder",
      "permissions": ["read", "write", "delete"]
    }
  ]
}
```

---

## Files Summary

| File | Action | Description |
|------|--------|-------------|
| `stores/auth.js` | CREATE | Pinia store for auth state |
| `composables/usePermissions.js` | CREATE | Permission helper functions |
| `pages/login.vue` | MODIFY | Fetch permissions after login |
| `App.vue` | MODIFY | Filter sidebar by allowed pages |
| `router/index.js` | MODIFY | Add route guards |
| `pages/*.vue` | MODIFY | Add v-if on action buttons |

---

## Implementation Order

1. Create Pinia store
2. Create composable
3. Modify login flow
4. Filter sidebar
5. Add route guards
6. Update page components (one by one)