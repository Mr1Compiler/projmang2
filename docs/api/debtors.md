# Debtors API

All endpoints require JWT authentication.

---

## List Debtors

```
GET /api/v1/debtors
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
      "name": "ABC Company",
      "email": "abc@example.com",
      "phone": "+1234567890",
      "totalDebt": 5000.00,
      "currency": "USD",
      "dueDate": "2024-02-15T00:00:00Z",
      "status": "pending"
    }
  ],
  "total": 30,
  "page": 1,
  "limit": 20,
  "totalPages": 2
}
```

---

## Get Debtor

```
GET /api/v1/debtors/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "ABC Company",
    "email": "abc@example.com",
    "phone": "+1234567890",
    "totalDebt": 5000.00,
    "currency": "USD",
    "dueDate": "2024-02-15T00:00:00Z",
    "status": "pending",
    "notes": "Payment expected next month",
    "createdBy": 1,
    "createdAt": "2024-01-01T08:00:00Z"
  }
}
```

---

## Create Debtor

```
POST /api/v1/debtors
```

### Request Body

```json
{
  "name": "ABC Company",
  "email": "abc@example.com",
  "phone": "+1234567890",
  "totalDebt": 5000.00,
  "currency": "USD",
  "dueDate": "2024-02-15T00:00:00Z",
  "status": "pending",
  "notes": "string",
  "createdBy": 1
}
```

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `name` | string | Yes | - |
| `email` | string | No | Valid email format |
| `phone` | string | No | - |
| `totalDebt` | float | Yes | - |
| `currency` | string | Yes | e.g., USD, EUR |
| `dueDate` | datetime | No | - |
| `status` | string | No | - |
| `notes` | string | No | - |

---

## Update Debtor

```
PUT /api/v1/debtors/:id
```

### Request Body

```json
{
  "name": "string",
  "email": "string",
  "phone": "string",
  "totalDebt": 5000.00,
  "currency": "string",
  "dueDate": "2024-02-15T00:00:00Z",
  "status": "string",
  "notes": "string"
}
```

All fields are optional.

---

## Delete Debtor

```
DELETE /api/v1/debtors/:id
```

---

## Get Debtor Stats

```
GET /api/v1/debtors/stats
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
    "total": 30,
    "active": 20,
    "paid": 10,
    "totalDebt": 150000.00,
    "activeDebt": 100000.00,
    "averageDebt": 5000.00
  }
}
```
