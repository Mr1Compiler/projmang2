# Users API

All endpoints require JWT authentication.

---

## List Users

```
GET /api/v1/users
```

### Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | int | 1 | Page number |
| `limit` | int | 20 | Items per page (max 100) |

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "fullName": "John Doe",
      "userName": "john_doe",
      "email": "john@example.com",
      "phone": "+1234567890",
      "jobTitle": "Engineer",
      "status": "active",
      "lastLogin": "2024-01-15T10:30:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get User

```
GET /api/v1/users/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "john_doe",
    "fullName": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "avatar": "https://example.com/avatar.jpg",
    "jobTitle": "Engineer",
    "status": "active",
    "lastLogin": "2024-01-15T10:30:00Z",
    "createdAt": "2024-01-01T08:00:00Z"
  }
}
```

---

## Create User

```
POST /api/v1/users
```

### Request Body

```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword123",
  "fullName": "John Doe",
  "phone": "+1234567890",
  "jobTitle": "Engineer"
}
```

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `username` | string | Yes | - |
| `email` | string | Yes | Valid email format |
| `password` | string | Yes | Minimum 8 characters |
| `fullName` | string | Yes | - |
| `phone` | string | Yes | - |
| `jobTitle` | string | Yes | - |

---

## Update User

```
PUT /api/v1/users/:id
```

### Request Body

```json
{
  "username": "string",
  "email": "string",
  "fullName": "string",
  "phone": "string",
  "avatar": "string",
  "jobTitle": "string",
  "status": "string"
}
```

All fields are optional.

---

## Update Password

```
PUT /api/v1/users/:id/password
```

### Request Body

```json
{
  "password": "newpassword123"
}
```

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `password` | string | Yes | Minimum 8 characters |

---

## Update Status

```
PATCH /api/v1/users/:id/status
```

### Request Body

```json
{
  "status": "inactive"
}
```

| Field | Type | Required | Values |
|-------|------|----------|--------|
| `status` | string | Yes | `active`, `inactive` |

---

## Delete User

```
DELETE /api/v1/users/:id
```

---

## Get User Roles

```
GET /api/v1/users/:id/roles
```

Returns all roles assigned to a specific user.

---

## Get User Team Members

```
GET /api/v1/users/:id/team-members
```

Returns all team member assignments for a specific user.
