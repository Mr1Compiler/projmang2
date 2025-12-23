# Finance API

All endpoints require JWT authentication.

---

# Expenses

## List Expenses

```
GET /api/v1/expenses
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
      "name": "Office Supplies",
      "amount": 500.00,
      "type": "operational",
      "expenseDate": "2024-01-15T00:00:00Z",
      "projectId": 1,
      "isDebtor": false,
      "debtorId": null,
      "status": "paid"
    }
  ],
  "total": 100,
  "page": 1,
  "limit": 20,
  "totalPages": 5
}
```

---

## Get Expense Stats

```
GET /api/v1/expenses/stats
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
    "total": 250,
    "totalAmount": 150000.00,
    "pending": 50,
    "approved": 180,
    "rejected": 20,
    "debtorCount": 15,
    "averageAmount": 600.00
  }
}
```

---

## Get Expense

```
GET /api/v1/expenses/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Office Supplies",
    "amount": 500.00,
    "type": "operational",
    "expenseDate": "2024-01-15T00:00:00Z",
    "projectId": 1,
    "isDebtor": false,
    "debtorId": null,
    "status": "paid",
    "notes": "Monthly supplies",
    "createdBy": 1,
    "createdAt": "2024-01-15T08:00:00Z"
  }
}
```

---

## Create Expense

```
POST /api/v1/expenses
```

### Request Body

```json
{
  "name": "Office Supplies",
  "amount": 500.00,
  "type": "operational",
  "expenseDate": "2024-01-15T00:00:00Z",
  "projectId": 1,
  "isDebtor": false,
  "debtorId": null,
  "status": "paid",
  "notes": "string",
  "createdBy": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Expense name |
| `amount` | float | Yes | Expense amount |
| `type` | string | No | Expense type |
| `expenseDate` | datetime | Yes | Date of expense |
| `projectId` | int | No | Related project ID |
| `isDebtor` | bool | No | Is this a debtor expense |
| `debtorId` | int | No | Related debtor ID (if isDebtor is true) |
| `status` | string | No | Payment status |
| `notes` | string | No | Additional notes |

---

## Update Expense

```
PUT /api/v1/expenses/:id
```

### Request Body

```json
{
  "name": "string",
  "amount": 500.00,
  "type": "string",
  "expenseDate": "2024-01-15T00:00:00Z",
  "projectId": 1,
  "isDebtor": false,
  "debtorId": null,
  "status": "string",
  "notes": "string"
}
```

---

## Delete Expense

```
DELETE /api/v1/expenses/:id
```

---

# Income

## List Income

```
GET /api/v1/income
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
      "name": "Project Payment",
      "amount": 10000.00,
      "type": "project",
      "incomeDate": "2024-01-15T00:00:00Z",
      "status": "received"
    }
  ],
  "total": 50,
  "page": 1,
  "limit": 20,
  "totalPages": 3
}
```

---

## Get Income

```
GET /api/v1/income/:id
```

### Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Project Payment",
    "amount": 10000.00,
    "type": "project",
    "incomeDate": "2024-01-15T00:00:00Z",
    "status": "received",
    "notes": "First milestone payment",
    "createdBy": 1,
    "createdAt": "2024-01-15T08:00:00Z"
  }
}
```

---

## Create Income

```
POST /api/v1/income
```

### Request Body

```json
{
  "name": "Project Payment",
  "amount": 10000.00,
  "type": "project",
  "incomeDate": "2024-01-15T00:00:00Z",
  "status": "received",
  "notes": "string",
  "createdBy": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Income name |
| `amount` | float | Yes | Income amount |
| `type` | string | No | Income type |
| `incomeDate` | datetime | Yes | Date of income |
| `status` | string | No | Receipt status |
| `notes` | string | No | Additional notes |

---

## Update Income

```
PUT /api/v1/income/:id
```

### Request Body

```json
{
  "name": "string",
  "amount": 10000.00,
  "type": "string",
  "incomeDate": "2024-01-15T00:00:00Z",
  "status": "string",
  "notes": "string"
}
```

---

## Delete Income

```
DELETE /api/v1/income/:id
```

---

## Get Income Stats

```
GET /api/v1/income/stats
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
    "totalAmount": 500000.00,
    "pending": 10,
    "approved": 35,
    "rejected": 5,
    "averageAmount": 10000.00
  }
}
```
