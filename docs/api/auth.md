# Authentication API

## Login

```
POST /api/v1/auth/login
```

**Auth Required:** No

### Request Body

```json
{
  "username": "string (required)",
  "password": "string (required)"
}
```

### Response

```json
{
  "success": true,
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": 1,
      "username": "john_doe",
      "fullName": "John Doe",
      "email": "john@example.com",
      "phone": "+1234567890",
      "avatar": null,
      "jobTitle": "Project Manager",
      "status": "active",
      "lastLogin": "2025-12-23T10:30:00Z",
      "createdAt": "2025-01-01T00:00:00Z"
    }
  }
}
```

### Error Responses

**401 - Invalid Credentials:**
```json
{
  "success": false,
  "error": "invalid username or password"
}
```

**401 - Inactive Account:**
```json
{
  "success": false,
  "error": "user account is inactive"
}
```

---

## Refresh Token

```
POST /api/v1/auth/refresh
```

**Auth Required:** No

### Request Body

```json
{
  "refreshToken": "string (required)"
}
```

### Response

```json
{
  "success": true,
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": 1,
      "username": "john_doe",
      "fullName": "John Doe",
      "email": "john@example.com"
    }
  }
}
```

### Error Response

**401 - Invalid Token:**
```json
{
  "success": false,
  "error": "invalid or expired refresh token"
}
```

---

## Logout

```
POST /api/v1/auth/logout
```

**Auth Required:** No

### Request Body

```json
{
  "refreshToken": "string (required)"
}
```

### Response

```json
{
  "success": true,
  "data": null
}
```

---

## Get My Pages

Get accessible pages for the authenticated user.

```
GET /api/v1/auth/pages
```

**Auth Required:** Yes

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Dashboard",
      "route": "/dashboard",
      "icon": "mdi-view-dashboard",
      "permissions": ["read"]
    },
    {
      "id": 2,
      "name": "Projects",
      "route": "/projects",
      "icon": "mdi-folder-multiple",
      "permissions": ["read", "write", "delete"]
    },
    {
      "id": 3,
      "name": "Users",
      "route": "/users",
      "icon": "mdi-account-multiple",
      "permissions": ["read", "write"]
    }
  ]
}
```

### Permissions Array

The `permissions` field is an array of strings indicating allowed actions:
- `"read"` - Can view/GET the resource
- `"write"` - Can create (POST) and update (PUT/PATCH) the resource
- `"delete"` - Can delete (DELETE) the resource

**Frontend Usage:**
```javascript
const hasRead = page.permissions.includes('read');
const hasWrite = page.permissions.includes('write');
const hasDelete = page.permissions.includes('delete');
```
