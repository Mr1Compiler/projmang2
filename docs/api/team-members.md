# Team Members API

All endpoints require JWT authentication.

---

## List Team Members

```
GET /api/v1/team-members
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
      "projectId": 1,
      "userId": 5,
      "createdAt": "2024-01-15T08:00:00Z",
      "updatedAt": "2024-01-15T08:00:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Team Member

```
GET /api/v1/team-members/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "projectId": 1,
    "userId": 5,
    "createdAt": "2024-01-15T08:00:00Z",
    "updatedAt": "2024-01-15T08:00:00Z"
  }
}
```

---

## Add Team Member

```
POST /api/v1/team-members
```

### Request Body

```json
{
  "projectId": 1,
  "userId": 5
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `projectId` | int | Yes | Project ID |
| `userId` | int | Yes | User ID to assign |

---

## Remove Team Member

```
DELETE /api/v1/team-members/:id
```

---

## Get Team Members by Project

You can also get team members for a specific project:

```
GET /api/v1/projects/:id/team-members
```

---

## Get Team Members by User

You can also get team member assignments for a specific user:

```
GET /api/v1/users/:id/team-members
```

---

## Get Team Member Stats

```
GET /api/v1/team-members/stats
```

### Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `period` | string | `month` | Time period filter: `all`, `month`, `year` |

**Period Options:**
- `month` - Current month (from day 1 of the month)
- `year` - Current year (from January 1)
- `all` - All time (no filter)

### Response

```json
{
  "success": true,
  "data": {
    "total": 50,
    "uniqueUsers": 25,
    "uniqueProjects": 10,
    "avgPerProject": 5.0
  }
}
```
