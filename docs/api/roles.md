# Roles API

All endpoints require JWT authentication.

---

# Roles

## List Roles

```
GET /api/v1/roles
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Admin",
      "description": "Full system access",
      "createdAt": "2024-01-01T08:00:00Z"
    }
  ],
  "total": 10,
  "page": 1,
  "limit": 20,
  "totalPages": 1
}
```

---

## Get Role

```
GET /api/v1/roles/:id
```

---

## Create Role

```
POST /api/v1/roles
```

### Request Body

```json
{
  "name": "Manager",
  "description": "Project management access"
}
```

| Field | Type | Required |
|-------|------|----------|
| `name` | string | Yes |
| `description` | string | No |

---

## Update Role

```
PUT /api/v1/roles/:id
```

### Request Body

```json
{
  "name": "string",
  "description": "string"
}
```

---

## Delete Role

```
DELETE /api/v1/roles/:id
```

---

# Pages

## List Pages

```
GET /api/v1/pages
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Dashboard",
      "icon": "dashboard",
      "route": "/dashboard",
      "status": "active",
      "createdAt": "2024-01-01T08:00:00Z"
    }
  ],
  "total": 20,
  "page": 1,
  "limit": 20,
  "totalPages": 1
}
```

---

## Get Page

```
GET /api/v1/pages/:id
```

---

## Create Page

```
POST /api/v1/pages
```

### Request Body

```json
{
  "name": "Reports",
  "icon": "chart",
  "route": "/reports",
  "status": "active"
}
```

| Field | Type | Required |
|-------|------|----------|
| `name` | string | Yes |
| `icon` | string | No |
| `route` | string | Yes |
| `status` | string | No |

---

## Update Page

```
PUT /api/v1/pages/:id
```

### Request Body

```json
{
  "name": "string",
  "icon": "string",
  "route": "string",
  "status": "string"
}
```

---

## Delete Page

```
DELETE /api/v1/pages/:id
```

---

# Role Pages (Permissions)

## List Role Pages

```
GET /api/v1/role-pages
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "roleId": 1,
      "pageId": 1,
      "permissions": "read,write,delete",
      "createdAt": "2024-01-01T08:00:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Role Page

```
GET /api/v1/role-pages/:id
```

---

## Create Role Page

```
POST /api/v1/role-pages
```

### Request Body

```json
{
  "roleId": 1,
  "pageId": 1,
  "permissions": "read,write"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `roleId` | int | Yes | Role ID |
| `pageId` | int | Yes | Page ID |
| `permissions` | string | No | Comma-separated: `read`, `write`, `delete` |

---

## Update Role Page

```
PUT /api/v1/role-pages/:id
```

### Request Body

```json
{
  "permissions": "read,write,delete"
}
```

---

## Delete Role Page

```
DELETE /api/v1/role-pages/:id
```

---

# User Roles

## List User Roles

```
GET /api/v1/user-roles
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "userId": 1,
      "roleId": 1,
      "createdAt": "2024-01-01T08:00:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get User Role

```
GET /api/v1/user-roles/:id
```

---

## Assign Role to User

```
POST /api/v1/user-roles
```

### Request Body

```json
{
  "userId": 1,
  "roleId": 1
}
```

| Field | Type | Required |
|-------|------|----------|
| `userId` | int | Yes |
| `roleId` | int | Yes |

---

## Remove Role from User

```
DELETE /api/v1/user-roles/:id
```
