# Work Categories API

All endpoints require JWT authentication.

---

# Work Categories

## List Categories

```
GET /api/v1/work-categories
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Foundation",
      "description": "Foundation work",
      "status": "active"
    }
  ],
  "total": 20,
  "page": 1,
  "limit": 20,
  "totalPages": 1
}
```

---

## Get Category

```
GET /api/v1/work-categories/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Foundation",
    "description": "Foundation work",
    "status": "active",
    "createdBy": 1,
    "createdAt": "2024-01-01T08:00:00Z"
  }
}
```

---

## Create Category

```
POST /api/v1/work-categories
```

### Request Body

```json
{
  "name": "Electrical",
  "description": "Electrical work",
  "status": "active",
  "createdBy": 1
}
```

| Field | Type | Required |
|-------|------|----------|
| `name` | string | Yes |
| `description` | string | No |
| `status` | string | No |

---

## Update Category

```
PUT /api/v1/work-categories/:id
```

### Request Body

```json
{
  "name": "string",
  "description": "string",
  "status": "string"
}
```

---

## Delete Category

```
DELETE /api/v1/work-categories/:id
```

---

## Get Work Category Stats

```
GET /api/v1/work-categories/stats
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
    "total": 20,
    "active": 15,
    "inactive": 5,
    "totalSubcategory": 50
  }
}
```

---

# Work Subcategories

## List Subcategories

```
GET /api/v1/work-subcategories
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "categoryId": 1,
      "name": "Concrete Pouring",
      "description": "Pouring foundation concrete",
      "percentage": 25.5,
      "status": "active"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Subcategory

```
GET /api/v1/work-subcategories/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "categoryId": 1,
    "name": "Concrete Pouring",
    "description": "Pouring foundation concrete",
    "percentage": 25.5,
    "status": "active",
    "createdBy": 1,
    "createdAt": "2024-01-01T08:00:00Z"
  }
}
```

---

## Create Subcategory

```
POST /api/v1/work-subcategories
```

### Request Body

```json
{
  "categoryId": 1,
  "name": "Concrete Pouring",
  "description": "Pouring foundation concrete",
  "percentage": 25.5,
  "status": "active",
  "createdBy": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `categoryId` | int | Yes | Parent category ID |
| `name` | string | Yes | Subcategory name |
| `description` | string | No | Description |
| `percentage` | float | No | Progress percentage weight |
| `status` | string | No | Status |

---

## Update Subcategory

```
PUT /api/v1/work-subcategories/:id
```

### Request Body

```json
{
  "name": "string",
  "description": "string",
  "percentage": 25.5,
  "status": "string"
}
```

---

## Delete Subcategory

```
DELETE /api/v1/work-subcategories/:id
```
