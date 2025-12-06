# Authentication & Authorization System

## Overview

The system uses **JWT-based authentication** with **Role-Based Access Control (RBAC)** for authorization.

## Authentication (Who are you?)

### Login Flow
1. User sends `POST /api/v1/auth/login` with username & password
2. Server validates credentials (bcrypt)
3. Server returns JWT token containing `user_id` and `username`
4. Client stores token and sends it in `Authorization: Bearer <token>` header

### JWT Structure
```json
{
  "user_id": 1,
  "username": "john",
  "exp": 1234567890,
  "iat": 1234567890
}
```

## Authorization (What can you do?)

### Database Structure

```
users
  └── userRoles (many-to-many)
        └── roles
              └── rolePages (many-to-many)
                    └── pages
                          └── permissions (JSON array)
```

### Two-Level Access Control

**Level 1 - Page Access:**
- User's role determines which pages they can access
- No role-page mapping = no access to that page

**Level 2 - Action Permissions:**
- `permissions` field in `rolePages` defines allowed actions per page
- Stored as JSON array: `["read", "write", "delete"]`

### Permission Mapping

| HTTP Method | Permission |
|-------------|------------|
| GET         | read       |
| POST        | write      |
| PUT/PATCH   | write      |
| DELETE      | delete     |

### Example

User "John" has role "Manager":

| Page     | Permissions              |
|----------|--------------------------|
| users    | `["read"]`               |
| projects | `["read", "write"]`      |
| reports  | `["read", "write", "delete"]` |

John can:
- View users (read only)
- View and create/edit projects
- Full access to reports

## API Endpoints

### Public (No Auth)
- `POST /api/v1/auth/login` - Login

### Protected (JWT Required)
- `GET /api/v1/me/pages` - Get current user's allowed pages with permissions
- All other `/api/v1/*` routes

## Frontend Integration

1. **Login:** Store JWT token
2. **Fetch permissions:** Call `GET /api/v1/me/pages` after login
3. **Sidebar:** Show only pages user has access to
4. **UI Controls:** Show/hide buttons based on permissions array
   - Has `"read"` → can view
   - Has `"write"` → show edit/create buttons
   - Has `"delete"` → show delete button

## Response Format

### Login Response
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "userId": 1,
    "username": "john"
  }
}
```

### User Pages Response
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Users",
      "route": "/users",
      "icon": "users-icon",
      "permissions": ["read"]
    },
    {
      "id": 2,
      "name": "Projects",
      "route": "/projects",
      "icon": "projects-icon",
      "permissions": ["read", "write", "delete"]
    }
  ]
}
```

### Unauthorized Response (401)
```json
{
  "success": false,
  "error": "authorization header required"
}
```

### Forbidden Response (403)
```json
{
  "success": false,
  "error": "access denied"
}
```
