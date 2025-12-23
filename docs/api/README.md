# API Documentation

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

Most endpoints require JWT authentication. Include the access token in the `Authorization` header:

```
Authorization: Bearer <access_token>
```

## Pagination

All list endpoints support pagination with query parameters:

| Parameter | Type | Default | Max | Description |
|-----------|------|---------|-----|-------------|
| `page` | int | 1 | - | Page number |
| `limit` | int | 20 | 100 | Items per page |

### Paginated Response Format

```json
{
  "data": [],
  "total": 100,
  "page": 1,
  "limit": 20,
  "totalPages": 5
}
```

## Common Response Format

### Success Response

```json
{
  "success": true,
  "data": { ... }
}
```

### Created Response (201)

```json
{
  "success": true,
  "data": { ... }
}
```

### Error Response

```json
{
  "success": false,
  "error": "error message"
}
```

### Validation Error Response

```json
{
  "success": false,
  "errors": {
    "fieldName": "field is required",
    "email": "email must be a valid email"
  }
}
```

## API Sections

| Resource | Description | File |
|----------|-------------|------|
| [Authentication](./auth.md) | Login, logout, token refresh | auth.md |
| [Dashboard](./dashboard.md) | Stats and analytics | dashboard.md |
| [Projects](./projects.md) | Project management | projects.md |
| [Workdays](./workdays.md) | Workdays, materials, labor, equipment | workdays.md |
| [Finance](./finance.md) | Expenses and income | finance.md |
| [Users](./users.md) | User management | users.md |
| [Roles](./roles.md) | Roles, pages, permissions | roles.md |
| [Work Categories](./work-categories.md) | Work categories and subcategories | work-categories.md |
| [Debtors](./debtors.md) | Debtor management | debtors.md |
| [Team Members](./team-members.md) | Team member assignments | team-members.md |

## Date Format

All dates use ISO 8601 format: `2024-01-15T10:30:00Z`

## Health Check

```
GET /health
```

No authentication required. Returns server health status.
