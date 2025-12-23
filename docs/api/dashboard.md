# Dashboard API

All endpoints require JWT authentication.

---

## Get Dashboard Stats

```
GET /api/v1/dashboard/stats
```

### Response

```json
{
  "success": true,
  "data": {
    "totalProjects": 50,
    "activeProjects": 20,
    "completedProjects": 25,
    "totalEngineers": 15
  }
}
```

---

## Get Financial Stats

```
GET /api/v1/dashboard/financial
```

### Query Parameters

| Parameter | Type | Required | Default | Values | Description |
|-----------|------|----------|---------|--------|-------------|
| `period` | string | No | `all` | `all`, `month` | Filter period |
| `month` | string | No* | - | `2024-01` | Specific month (YYYY-MM format). *Required when period=month |

### Response

```json
{
  "success": true,
  "data": {
    "totalIncome": 150000.00,
    "totalExpenses": 85000.00
  }
}
```

---

## Get Project Progress

```
GET /api/v1/dashboard/project-progress
```

### Query Parameters

| Parameter | Type | Required | Default | Values | Description |
|-----------|------|----------|---------|--------|-------------|
| `status` | string | No | `in_progress` | `pending`, `in_progress`, `completed`, `all` | Filter by status |
| `limit` | int | No | 5 | 1-20 | Max projects to return |

### Response

```json
{
  "success": true,
  "data": {
    "projects": [
      {
        "id": 1,
        "name": "Project Alpha",
        "progress": 75.5
      },
      {
        "id": 2,
        "name": "Project Beta",
        "progress": 45.0
      }
    ]
  }
}
```

---

## Get Task Summary

```
GET /api/v1/dashboard/task-summary
```

### Query Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `month` | string | Yes | Month in YYYY-MM format |

### Response

```json
{
  "success": true,
  "data": {
    "labels": ["Week 1", "Week 2", "Week 3", "Week 4"],
    "completed": [10, 15, 8, 12],
    "pending": [5, 3, 7, 4]
  }
}
```

---

## Get Activities

Returns recent user activities (audit logs) for the dashboard activity feed.

```
GET /api/v1/dashboard/activities
```

### Query Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `page` | int | No | 1 | Page number |
| `limit` | int | No | 20 | Items per page (max 100) |

### Response

```json
{
  "success": true,
  "data": {
    "data": [
      {
        "id": 45,
        "actorName": "Ahmed",
        "action": "create",
        "targetType": "project",
        "targetName": "Office Building Phase 2",
        "createdAt": "2025-12-18T10:30:00Z"
      },
      {
        "id": 44,
        "actorName": "Sara",
        "action": "update",
        "targetType": "expense",
        "targetName": "Cement Purchase",
        "createdAt": "2025-12-18T09:15:00Z"
      }
    ],
    "total": 150,
    "page": 1,
    "limit": 20,
    "totalPages": 8
  }
}
```

### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | int64 | Activity ID |
| `actorName` | string | Username of the user who performed the action |
| `action` | string | Action type (see Action Values below) |
| `targetType` | string | Type of resource affected (see Target Types below) |
| `targetName` | string | Name of the affected resource |
| `createdAt` | string | ISO 8601 timestamp |

### Action Values

The `action` field contains one of these values. **Frontend should translate these for Arabic UI:**

| Value | English | Arabic |
|-------|---------|--------|
| `create` | Created | أضاف |
| `update` | Updated | حدّث |
| `delete` | Deleted | حذف |
| `login` | Logged in | سجّل دخول |
| `logout` | Logged out | سجّل خروج |
| `status_change` | Changed status | غيّر الحالة |
| `refresh` | Refreshed token | جدّد الجلسة |

### Target Types

The `targetType` field contains one of these values. **Frontend should translate these for Arabic UI:**

| Value | English | Arabic |
|-------|---------|--------|
| `project` | Project | مشروع |
| `expense` | Expense | مصروف |
| `income` | Income | دخل |
| `debtor` | Debtor | مدين |
| `user` | User | مستخدم |
| `work_category` | Work Category | فئة عمل |
| `work_subcategory` | Work Subcategory | فئة فرعية |
| `team_member` | Team Member | عضو فريق |
| `auth` | Authentication | المصادقة |
| `role` | Role | دور |
| `role_page` | Role Permission | صلاحية الدور |
| `page` | Page | صفحة |

### Frontend Translation Example

```javascript
// Translation maps for Arabic
const actionMapAr = {
  create: "أضاف",
  update: "حدّث",
  delete: "حذف",
  login: "سجّل دخول",
  logout: "سجّل خروج",
  status_change: "غيّر الحالة",
  refresh: "جدّد الجلسة"
};

const targetTypeMapAr = {
  project: "مشروع",
  expense: "مصروف",
  income: "دخل",
  debtor: "مدين",
  user: "مستخدم",
  work_category: "فئة عمل",
  work_subcategory: "فئة فرعية",
  team_member: "عضو فريق",
  auth: "المصادقة",
  role: "دور",
  role_page: "صلاحية الدور",
  page: "صفحة"
};

// Build activity text
function formatActivity(activity, locale = 'ar') {
  if (locale === 'ar') {
    return `${activity.actorName} ${actionMapAr[activity.action]} ${targetTypeMapAr[activity.targetType]} '${activity.targetName}'`;
    // Example: "Ahmed أضاف مشروع 'Office Building Phase 2'"
  }
  return `${activity.actorName} ${activity.action}d ${activity.targetType} '${activity.targetName}'`;
}
```

### Notes

- Activities are sorted by `createdAt` descending (newest first)
- Only successful operations are returned (failed operations are excluded)
- `actorName` shows "System" for operations without a logged-in user
- `targetName` shows "Unknown" if the target resource was deleted or not found
