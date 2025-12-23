# Projects API

All endpoints require JWT authentication.

---

## List Projects

```
GET /api/v1/projects
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
      "name": "Project Alpha",
      "type": "construction",
      "clientPhone": "+1234567890",
      "location": "New York",
      "startDate": "2024-01-15T00:00:00Z",
      "status": "in_progress",
      "progressPercentage": 75.5,
      "warningCost": 80000.00,
      "totalCost": 100000.00,
      "currentSpending": 60000.00
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Project

```
GET /api/v1/projects/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Project Alpha",
    "type": "construction",
    "description": "Main building project",
    "clientPhone": "+1234567890",
    "location": "New York",
    "startDate": "2024-01-15T00:00:00Z",
    "duration": 90,
    "warningCost": 80000.00,
    "totalCost": 100000.00,
    "status": "in_progress",
    "progressPercentage": 75.5,
    "notes": "On track",
    "createdBy": 1,
    "createdAt": "2024-01-10T08:00:00Z",
      "currentSpending": 60000.00
  }
}
```

---

## Get Project Stats

```
GET /api/v1/projects/stats
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
    "pending": 10,
    "inProgress": 25,
    "completed": 15,
    "totalBudget": 5000000.00,
    "averageProgress": 62.5
  }
}
```

---

## Create Project

```
POST /api/v1/projects
```

### Request Body

```json
{
  "name": "string (required)",
  "type": "string",
  "description": "string",
  "clientPhone": "string",
  "location": "string",
  "startDate": "2024-01-15T00:00:00Z",
  "duration": 90,
  "warningCost": 80000.00,
  "totalCost": 100000.00,
  "notes": "string",
  "createdBy": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Project name |
| `type` | string | No | Project type |
| `description` | string | No | Project description |
| `clientPhone` | string | No | Client phone number |
| `location` | string | No | Project location |
| `startDate` | datetime | No | Project start date |
| `duration` | int | No | Duration in days |
| `warningCost` | float | Yes | Warning budget threshold |
| `totalCost` | float | Yes | Total budget |
| `notes` | string | No | Additional notes |

### Response

Returns the created project object.

---

## Update Project

```
PUT /api/v1/projects/:id
```

### Request Body

All fields are optional. Only include fields you want to update.

```json
{
  "name": "string",
  "type": "string",
  "description": "string",
  "clientPhone": "string",
  "location": "string",
  "startDate": "2024-01-15T00:00:00Z",
  "duration": 90,
  "warningCost": 80000.00,
  "totalCost": 100000.00,
  "status": "in_progress",
  "progressPercentage": 75.5,
  "notes": "string"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `status` | string | `pending`, `in_progress`, `completed` |
| `progressPercentage` | float | 0-100 |

### Response

Returns the updated project object.

---

## Delete Project

```
DELETE /api/v1/projects/:id
```

### Response

```json
{
  "success": true,
  "data": null
}
```

---

## Get Project Workdays

```
GET /api/v1/projects/:id/workdays
```

Returns all workdays for a specific project.

---

## Get Project Expenses

```
GET /api/v1/projects/:id/expenses
```

Returns all expenses for a specific project.

---

## Get Project Team Members

```
GET /api/v1/projects/:id/team-members
```

Returns all team members assigned to a specific project.
