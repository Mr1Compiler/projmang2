# Workdays API

All endpoints require JWT authentication.

---

# Workdays

## List Workdays

```
GET /api/v1/workdays
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
      "workSubCategoryId": 5,
      "workDate": "2024-01-15T00:00:00Z",
      "status": "completed",
      "totalCost": 5000.00
    }
  ],
  "total": 100,
  "page": 1,
  "limit": 20,
  "totalPages": 5
}
```

---

## Get Workday

```
GET /api/v1/workdays/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "projectId": 1,
    "workSubCategoryId": 5,
    "workDate": "2024-01-15T00:00:00Z",
    "description": "Foundation work",
    "status": "completed",
    "totalCost": 5000.00,
    "notes": "Completed on time",
    "createdBy": 1,
    "createdAt": "2024-01-15T08:00:00Z"
  }
}
```

---

## Create Workday

```
POST /api/v1/workdays
```

### Request Body

```json
{
  "projectId": 1,
  "workSubCategoryId": 5,
  "workDate": "2024-01-15T00:00:00Z",
  "description": "string",
  "notes": "string",
  "createdBy": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `projectId` | int | Yes | Project ID |
| `workSubCategoryId` | int | No | Work subcategory ID |
| `workDate` | datetime | Yes | Work date |
| `description` | string | No | Description |
| `notes` | string | No | Notes |

---

## Update Workday

```
PUT /api/v1/workdays/:id
```

### Request Body

```json
{
  "workSubCategoryId": 5,
  "workDate": "2024-01-15T00:00:00Z",
  "description": "string",
  "status": "completed",
  "totalCost": 5000.00,
  "notes": "string"
}
```

---

## Complete Workday

```
PATCH /api/v1/workdays/:id/complete
```

Marks a workday as complete.

### Response

Returns the updated workday object.

---

## Delete Workday

```
DELETE /api/v1/workdays/:id
```

---

# Workday Materials

## List Materials

```
GET /api/v1/workday-materials
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "workDayId": 1,
      "materialName": "Cement",
      "quantity": 100.0,
      "cost": 500.00,
      "notes": "Type A cement",
      "createdAt": "2024-01-15T08:00:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Material

```
GET /api/v1/workday-materials/:id
```

---

## Create Material

```
POST /api/v1/workday-materials
```

### Request Body

```json
{
  "workDayId": 1,
  "materialName": "Cement",
  "quantity": 100.0,
  "cost": 500.00,
  "notes": "string"
}
```

| Field | Type | Required |
|-------|------|----------|
| `workDayId` | int | Yes |
| `materialName` | string | Yes |
| `quantity` | float | Yes |
| `cost` | float | Yes |
| `notes` | string | No |

---

## Update Material

```
PUT /api/v1/workday-materials/:id
```

### Request Body

```json
{
  "materialName": "string",
  "quantity": 100.0,
  "cost": 500.00,
  "notes": "string"
}
```

---

## Delete Material

```
DELETE /api/v1/workday-materials/:id
```

---

# Workday Labor

## List Labor

```
GET /api/v1/workday-labor
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "workDayId": 1,
      "workerName": "John Doe",
      "jobTitle": "Mason",
      "phone": "+1234567890",
      "address": "123 Main St",
      "quantity": 8.0,
      "cost": 200.00,
      "notes": "Full day",
      "createdAt": "2024-01-15T08:00:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Labor

```
GET /api/v1/workday-labor/:id
```

---

## Create Labor

```
POST /api/v1/workday-labor
```

### Request Body

```json
{
  "workDayId": 1,
  "workerName": "John Doe",
  "jobTitle": "Mason",
  "phone": "+1234567890",
  "address": "123 Main St",
  "quantity": 8.0,
  "cost": 200.00,
  "notes": "string"
}
```

| Field | Type | Required |
|-------|------|----------|
| `workDayId` | int | Yes |
| `workerName` | string | Yes |
| `jobTitle` | string | No |
| `phone` | string | No |
| `address` | string | No |
| `quantity` | float | Yes |
| `cost` | float | Yes |
| `notes` | string | No |

---

## Update Labor

```
PUT /api/v1/workday-labor/:id
```

### Request Body

```json
{
  "workerName": "string",
  "jobTitle": "string",
  "phone": "string",
  "address": "string",
  "quantity": 8.0,
  "cost": 200.00,
  "notes": "string"
}
```

---

## Delete Labor

```
DELETE /api/v1/workday-labor/:id
```

---

# Workday Equipment

## List Equipment

```
GET /api/v1/workday-equipment
```

### Response

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "workDayId": 1,
      "equipmentName": "Excavator",
      "quantity": 1.0,
      "cost": 1000.00,
      "notes": "Rented",
      "createdAt": "2024-01-15T08:00:00Z"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Equipment

```
GET /api/v1/workday-equipment/:id
```

---

## Create Equipment

```
POST /api/v1/workday-equipment
```

### Request Body

```json
{
  "workDayId": 1,
  "equipmentName": "Excavator",
  "quantity": 1.0,
  "cost": 1000.00,
  "notes": "string"
}
```

| Field | Type | Required |
|-------|------|----------|
| `workDayId` | int | Yes |
| `equipmentName` | string | Yes |
| `quantity` | float | Yes |
| `cost` | float | Yes |
| `notes` | string | No |

---

## Update Equipment

```
PUT /api/v1/workday-equipment/:id
```

### Request Body

```json
{
  "equipmentName": "string",
  "quantity": 1.0,
  "cost": 1000.00,
  "notes": "string"
}
```

---

## Delete Equipment

```
DELETE /api/v1/workday-equipment/:id
```
