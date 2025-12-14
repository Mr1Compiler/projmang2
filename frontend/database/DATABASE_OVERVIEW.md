# نظرة عامة على قاعدة البيانات
## Database Overview

## 📊 مخطط العلاقات (ERD)

### النظام الهرمي

```
نظام إدارة المشاريع
│
├── نظام المستخدمين والصلاحيات
│   ├── Users (المستخدمون)
│   ├── Roles (الأدوار)
│   ├── Permissions (الصلاحيات)
│   ├── User_Roles (ربط المستخدمين بالأدوار)
│   └── Role_Permissions (ربط الأدوار بالصلاحيات)
│
├── التنظيم الإداري
│   ├── Departments (الأقسام)
│   ├── Team_Members (أعضاء الفريق)
│   ├── Expense_Categories (فئات المصاريف)
│   ├── Income_Categories (فئات الإيرادات)
│   └── Expense_Types (أنواع المصاريف)
│
├── المشاريع والمهام
│   ├── Projects (المشاريع)
│   ├── Project_Team (فريق المشروع)
│   ├── Tasks (المهام)
│   ├── Task_Comments (تعليقات المهام)
│   └── Task_Attachments (مرفقات المهام)
│
├── الإدارة المالية
│   ├── Expenses (المصاريف)
│   ├── Project_Expenses (مصاريف المشاريع)
│   ├── Income (الإيرادات)
│   ├── Debtors (المدينون)
│   ├── Debtor_Payments (دفعات المدينين)
│   └── Payment_Installments (أقساط الدفع)
│
└── أيام العمل
    ├── Work_Days (أيام العمل)
    ├── Work_Day_Labor (العمالة)
    ├── Work_Day_Equipment (المعدات)
    └── Work_Day_Materials (المواد)
```

## 📋 الجداول الرئيسية

### 1. نظام المستخدمين والصلاحيات (User & Permission System)

#### users - المستخدمون
```
id (PK)
username
email (UNIQUE)
password
full_name
phone
avatar
status [active, inactive, suspended, pending]
last_login
email_verified
created_at
updated_at
```

**العلاقات:**
- `users` → `team_members` (1:1)
- `users` → `user_roles` (1:N)
- `users` → `projects` (manager_id) (1:N)
- `users` → `tasks` (created_by, assigned_by) (1:N)

#### roles - الأدوار
```
id (PK)
name (UNIQUE)
display_name
description
is_system_role
created_at
updated_at
```

**الأدوار الافتراضية:**
- super_admin (مدير النظام)
- admin (مدير)
- project_manager (مدير مشروع)
- engineer (مهندس)
- accountant (محاسب)
- viewer (مشاهد)

### 2. الأقسام والفئات (Departments & Categories)

#### departments - الأقسام
```
id (PK)
name
code (UNIQUE)
description
manager_id (FK → users)
parent_department_id (FK → departments)
budget
status [active, inactive]
created_at
updated_at
```

**العلاقات:**
- `departments` → `team_members` (1:N)
- `departments` → `projects` (1:N)
- `departments` → `expenses` (1:N)

### 3. أعضاء الفريق (Team Members)

#### team_members - أعضاء الفريق
```
id (PK)
user_id (FK → users)
name
email
phone
department_id (FK → departments)
position
role
specialization
hire_date
salary
hourly_rate
status [active, inactive, on_leave, resigned]
created_at
updated_at
```

**العلاقات:**
- `team_members` → `tasks` (assigned_to) (1:N)
- `team_members` → `project_team` (N:M مع projects)

### 4. المشاريع (Projects)

#### projects - المشاريع
```
id (PK)
project_number (UNIQUE, auto-generated)
name
description
client_name
client_phone
client_email
client_address
work_location
start_date
end_date
planned_end_date
contract_value
actual_cost
estimated_cost
status [pending, in_progress, on_hold, completed, cancelled]
priority [low, medium, high, urgent]
progress_percentage
manager_id (FK → users)
department_id (FK → departments)
created_by (FK → users)
created_at
updated_at
```

**العلاقات:**
- `projects` → `tasks` (1:N)
- `projects` → `project_team` (N:M مع team_members)
- `projects` → `project_expenses` (N:M مع expenses)
- `projects` → `income` (1:N)
- `projects` → `debtors` (1:N)
- `projects` → `work_days` (1:N)

### 5. المهام (Tasks)

#### tasks - المهام
```
id (PK)
task_number (UNIQUE, auto-generated)
title
description
project_id (FK → projects)
parent_task_id (FK → tasks)
assigned_to (FK → team_members)
assigned_by (FK → users)
status [pending, in_progress, review, completed, cancelled, on_hold]
priority [low, medium, high, urgent]
start_date
due_date
completed_date
estimated_hours
actual_hours
progress_percentage
created_by (FK → users)
created_at
updated_at
```

**العلاقات:**
- `tasks` → `task_comments` (1:N)
- `tasks` → `task_attachments` (1:N)
- `tasks` → `tasks` (parent_task_id) (1:N) - مهام فرعية

### 6. المصاريف (Expenses)

#### expenses - المصاريف
```
id (PK)
expense_number (UNIQUE, auto-generated)
description
amount
expense_date
expense_type_id (FK → expense_types)
category_id (FK → expense_categories)
project_id (FK → projects via project_expenses)
department_id (FK → departments)
paid_to
payment_method [cash, bank_transfer, check, credit_card, other]
receipt_number
status [pending, approved, rejected, paid]
approved_by (FK → users)
approved_at
created_by (FK → users)
created_at
updated_at
```

**العلاقات:**
- `expenses` → `project_expenses` (N:M مع projects)

### 7. الإيرادات (Income)

#### income - الإيرادات
```
id (PK)
income_number (UNIQUE, auto-generated)
description
amount
income_date
category_id (FK → income_categories)
project_id (FK → projects)
client_name
payment_method [cash, bank_transfer, check, credit_card, other]
receipt_number
invoice_number
status [pending, confirmed, cancelled]
created_by (FK → users)
created_at
updated_at
```

### 8. المدينون (Debtors)

#### debtors - المدينون
```
id (PK)
debtor_number (UNIQUE, auto-generated)
name
email
phone
address
company_name
total_debt
paid_amount
remaining_amount
project_id (FK → projects)
due_date
status [active, partially_paid, fully_paid, overdue, cancelled]
created_by (FK → users)
created_at
updated_at
```

**العلاقات:**
- `debtors` → `debtor_payments` (1:N)
- `debtors` → `payment_installments` (1:N)

### 9. أيام العمل (Work Days)

#### work_days - أيام العمل
```
id (PK)
work_day_number (UNIQUE, auto-generated)
project_id (FK → projects)
work_date
work_location
form_number
work_period_from
work_period_to
day_name
work_type
description
status [pending, in_progress, completed, cancelled]
total_labor_cost (calculated)
total_equipment_cost (calculated)
total_materials_cost (calculated)
total_cost (calculated)
created_by (FK → users)
created_at
updated_at
```

**العلاقات:**
- `work_days` → `work_day_labor` (1:N)
- `work_days` → `work_day_equipment` (1:N)
- `work_days` → `work_day_materials` (1:N)

## 🔗 العلاقات الرئيسية

### 1. علاقات المستخدمين
```
users ←→ user_roles ←→ roles ←→ role_permissions ←→ permissions
users → team_members (1:1)
users → projects (manager_id)
users → tasks (assigned_by, created_by)
```

### 2. علاقات المشاريع
```
projects → tasks (1:N)
projects ←→ project_team ←→ team_members (N:M)
projects ←→ project_expenses ←→ expenses (N:M)
projects → income (1:N)
projects → debtors (1:N)
projects → work_days (1:N)
```

### 3. علاقات المهام
```
tasks → project (N:1)
tasks → team_member (assigned_to) (N:1)
tasks → tasks (parent_task_id) (1:N)
tasks → task_comments (1:N)
tasks → task_attachments (1:N)
```

### 4. علاقات أيام العمل
```
work_days → project (N:1)
work_days → work_day_labor (1:N)
work_days → work_day_equipment (1:N)
work_days → work_day_materials (1:N)
```

## 📈 المشاهد (Views)

### view_project_summary
ملخص شامل للمشاريع مع الإحصائيات

**الأعمدة:**
- معلومات المشروع الأساسية
- اسم المدير
- اسم القسم
- عدد المهام (الإجمالي والمكتمل)
- حجم الفريق
- إجمالي المصاريف
- إجمالي الإيرادات

### view_task_summary
ملخص المهام مع المعلومات المرتبطة

### view_expense_summary
ملخص المصاريف مع التصنيفات

### view_income_summary
ملخص الإيرادات مع المشاريع

### view_debtor_summary
ملخص المدينين مع الأقساط والدفعات

### view_work_day_summary
ملخص أيام العمل مع التكاليف

## 🔄 المشغلات (Triggers)

### 1. مشغلات التحديث التلقائي
- `after_work_day_labor_insert/update/delete` → تحديث تكاليف العمالة
- `after_work_day_equipment_insert/update/delete` → تحديث تكاليف المعدات
- `after_work_day_materials_insert/update/delete` → تحديث تكاليف المواد
- `after_debtor_payment_insert` → تحديث حالة المدين
- `after_task_update` → تحديث تقدم المشروع

### 2. مشغلات الأرقام التلقائية
- `before_project_insert` → رقم المشروع (PRJ-YYYY-XXXX)
- `before_expense_insert` → رقم المصروف (EXP-YYYY-XXXX)
- `before_income_insert` → رقم الإيراد (INC-YYYY-XXXX)
- `before_task_insert` → رقم المهمة (TSK-YYYY-XXXX)
- `before_work_day_insert` → رقم يوم العمل (WD-YYYY-XXXX)
- `before_debtor_insert` → رقم المدين (DBT-YYYY-XXXX)

### 3. مشغلات الحسابات
- `before_work_day_labor_insert/update` → حساب التكلفة الإجمالية
- `before_work_day_equipment_insert/update` → حساب التكلفة الإجمالية
- `before_work_day_materials_insert/update` → حساب التكلفة الإجمالية

## ⚙️ الإجراءات المخزنة (Stored Procedures)

| الإجراء | الوصف |
|---------|-------|
| `calculate_project_costs` | حساب تكاليف المشروع |
| `calculate_project_progress` | حساب تقدم المشروع |
| `update_debtor_status` | تحديث حالة المدين |
| `calculate_work_day_costs` | حساب تكاليف يوم العمل |
| `get_monthly_financial_report` | التقرير المالي الشهري |
| `assign_task` | تعيين مهمة |
| `add_debtor_payment` | إضافة دفعة مدين |
| `get_dashboard_stats` | إحصائيات لوحة التحكم |
| `get_overdue_projects` | المشاريع المتأخرة |
| `get_team_performance` | أداء الفريق |

## 📊 الإحصائيات

### حجم قاعدة البيانات

```sql
SELECT 
    COUNT(*) as total_tables
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'project_management_db';
-- النتيجة: 35+ جدول

SELECT 
    COUNT(*) as total_views
FROM information_schema.VIEWS
WHERE TABLE_SCHEMA = 'project_management_db';
-- النتيجة: 6 views

SELECT 
    COUNT(*) as total_procedures
FROM information_schema.ROUTINES
WHERE ROUTINE_SCHEMA = 'project_management_db'
    AND ROUTINE_TYPE = 'PROCEDURE';
-- النتيجة: 10 إجراءات مخزنة

SELECT 
    COUNT(*) as total_triggers
FROM information_schema.TRIGGERS
WHERE TRIGGER_SCHEMA = 'project_management_db';
-- النتيجة: 25+ مشغل
```

## 🔒 الأمان والصلاحيات

### مستويات الوصول

1. **مدير النظام (Super Admin)**
   - صلاحيات كاملة على جميع الجداول
   - إدارة المستخدمين والأدوار
   - الوصول إلى الإعدادات

2. **المدير (Admin)**
   - معظم الصلاحيات
   - لا يمكنه تعديل الإعدادات الحرجة

3. **مدير المشروع (Project Manager)**
   - إدارة المشاريع والمهام
   - إدارة الفريق
   - عرض التقارير المالية

4. **المهندس (Engineer)**
   - عرض وتحديث المهام المعينة له
   - عرض المشاريع
   - عرض أيام العمل

5. **المحاسب (Accountant)**
   - إدارة المصاريف والإيرادات
   - إدارة المدينين
   - عرض التقارير المالية

6. **المشاهد (Viewer)**
   - صلاحيات القراءة فقط

## 🎯 أفضل الممارسات

### 1. النسخ الاحتياطي
```bash
# نسخة احتياطية يومية
./database/maintenance/backup.sh full

# نسخة احتياطية للبيانات فقط
./database/maintenance/backup.sh data
```

### 2. الصيانة الدورية
```sql
-- تنفيذ أسبوعياً
SOURCE /path/to/database/maintenance/optimize.sql;
```

### 3. مراقبة الأداء
```sql
-- مراقبة الاستعلامات البطيئة
SET GLOBAL slow_query_log = 'ON';
SET GLOBAL long_query_time = 2;
```

### 4. التأمين
- تغيير كلمات المرور الافتراضية فوراً
- استخدام SSL للاتصالات
- تقييد الوصول للمنافذ
- تفعيل سجلات التدقيق

## 📞 للمزيد من المعلومات

راجع الملفات التالية:
- `README.md` - الدليل الشامل
- `INSTALLATION.md` - دليل التثبيت
- `schema.sql` - مخطط قاعدة البيانات
- `seed_data.sql` - البيانات الأولية
- `stored_procedures.sql` - الإجراءات المخزنة
- `triggers.sql` - المشغلات
- `queries/common_queries.sql` - الاستعلامات الشائعة

---

© 2024 ارتكاز للحلول البرمجية - جميع الحقوق محفوظة

