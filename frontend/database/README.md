# قاعدة بيانات نظام إدارة المشاريع
## Project Management System Database

## 📋 نظرة عامة

قاعدة بيانات شاملة ومتكاملة لنظام إدارة المشاريع الهندسية، مصممة لإدارة جميع جوانب المشاريع من المالية، الموارد البشرية، المهام، وأيام العمل.

## 🗄️ الجداول الرئيسية

### 1. نظام المستخدمين والصلاحيات
- **users** - جدول المستخدمين
- **roles** - جدول الأدوار
- **permissions** - جدول الصلاحيات
- **user_roles** - ربط المستخدمين بالأدوار
- **role_permissions** - ربط الأدوار بالصلاحيات

### 2. التنظيم الإداري
- **departments** - الأقسام
- **team_members** - أعضاء الفريق/المهندسين
- **expense_categories** - فئات المصاريف
- **income_categories** - فئات الإيرادات
- **expense_types** - أنواع المصاريف

### 3. المشاريع والمهام
- **projects** - المشاريع
- **project_team** - فريق المشروع
- **tasks** - المهام
- **task_comments** - تعليقات المهام
- **task_attachments** - مرفقات المهام

### 4. الإدارة المالية
- **expenses** - المصاريف
- **project_expenses** - مصاريف المشاريع
- **income** - الإيرادات
- **debtors** - المدينين
- **debtor_payments** - دفعات المدينين
- **payment_installments** - أقساط الدفع

### 5. أيام العمل
- **work_days** - أيام العمل
- **work_day_labor** - العمالة
- **work_day_equipment** - المعدات
- **work_day_materials** - المواد

## 🚀 التثبيت والإعداد

### المتطلبات الأساسية
- MySQL 8.0 أو أحدث
- أو MariaDB 10.5 أو أحدث

### خطوات التثبيت

#### 1. إنشاء قاعدة البيانات

```sql
CREATE DATABASE project_management_db
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;

USE project_management_db;
```

#### 2. تنفيذ ملف المخطط (Schema)

```bash
mysql -u username -p project_management_db < schema.sql
```

أو من داخل MySQL:

```sql
USE project_management_db;
SOURCE /path/to/schema.sql;
```

#### 3. إدراج البيانات الأولية

```bash
mysql -u username -p project_management_db < seed_data.sql
```

أو من داخل MySQL:

```sql
USE project_management_db;
SOURCE /path/to/seed_data.sql;
```

## 🔐 المستخدمون الافتراضيون

| اسم المستخدم | البريد الإلكتروني | كلمة المرور | الدور |
|--------------|-------------------|-------------|-------|
| admin | admin@erticaz.com | password123 | مدير النظام |
| manager1 | manager1@erticaz.com | password123 | مدير مشروع |
| engineer1 | engineer1@erticaz.com | password123 | مهندس |
| accountant1 | accountant@erticaz.com | password123 | محاسب |

> **تنبيه:** يجب تغيير كلمات المرور الافتراضية فوراً في بيئة الإنتاج!

## 📊 Views (المشاهد)

تحتوي قاعدة البيانات على عدة views جاهزة للتقارير:

- **view_project_summary** - ملخص المشاريع
- **view_task_summary** - ملخص المهام
- **view_expense_summary** - ملخص المصاريف
- **view_income_summary** - ملخص الإيرادات
- **view_debtor_summary** - ملخص المدينين
- **view_work_day_summary** - ملخص أيام العمل

### أمثلة الاستخدام

```sql
-- عرض ملخص جميع المشاريع
SELECT * FROM view_project_summary;

-- عرض المشاريع النشطة فقط
SELECT * FROM view_project_summary WHERE status = 'in_progress';

-- عرض المهام المتأخرة
SELECT * FROM view_task_summary WHERE days_remaining < 0;

-- إجمالي المصاريف حسب المشروع
SELECT project_name, SUM(amount) as total
FROM view_expense_summary
GROUP BY project_name;
```

## 🔍 الفهارس (Indexes)

تم إنشاء فهارس على الأعمدة الأكثر استخداماً لتحسين الأداء:

- فهارس على المفاتيح الأجنبية
- فهارس على الحالات (status)
- فهارس على التواريخ
- فهارس على الأرقام المرجعية

## 🔗 العلاقات

### علاقات المشاريع
- مشروع ← مدير (users)
- مشروع ← قسم (departments)
- مشروع ← فريق (team_members)
- مشروع ← مهام (tasks)
- مشروع ← مصاريف (expenses)
- مشروع ← إيرادات (income)

### علاقات المهام
- مهمة ← مشروع (projects)
- مهمة ← مسؤول (team_members)
- مهمة ← منشئ (users)
- مهمة ← تعليقات (task_comments)
- مهمة ← مرفقات (task_attachments)

### علاقات أيام العمل
- يوم عمل ← مشروع (projects)
- يوم عمل ← عمالة (work_day_labor)
- يوم عمل ← معدات (work_day_equipment)
- يوم عمل ← مواد (work_day_materials)

## 📈 استعلامات مفيدة

### إحصائيات المشاريع

```sql
-- عدد المشاريع حسب الحالة
SELECT status, COUNT(*) as count
FROM projects
GROUP BY status;

-- إجمالي قيمة المشاريع
SELECT 
    COUNT(*) as total_projects,
    SUM(contract_value) as total_contract_value,
    SUM(actual_cost) as total_actual_cost,
    AVG(progress_percentage) as avg_progress
FROM projects;
```

### التقارير المالية

```sql
-- ملخص المصاريف والإيرادات
SELECT 
    (SELECT SUM(amount) FROM expenses WHERE status = 'paid') as total_expenses,
    (SELECT SUM(amount) FROM income WHERE status = 'confirmed') as total_income,
    (SELECT SUM(amount) FROM income WHERE status = 'confirmed') - 
    (SELECT SUM(amount) FROM expenses WHERE status = 'paid') as net_profit;

-- المصاريف حسب الشهر
SELECT 
    DATE_FORMAT(expense_date, '%Y-%m') as month,
    COUNT(*) as expense_count,
    SUM(amount) as total_amount
FROM expenses
GROUP BY DATE_FORMAT(expense_date, '%Y-%m')
ORDER BY month DESC;
```

### تقارير الفريق

```sql
-- عدد المهام لكل عضو
SELECT 
    tm.name,
    COUNT(t.id) as total_tasks,
    SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) as completed_tasks,
    SUM(CASE WHEN t.status = 'in_progress' THEN 1 ELSE 0 END) as in_progress_tasks
FROM team_members tm
LEFT JOIN tasks t ON tm.id = t.assigned_to
GROUP BY tm.id, tm.name;
```

### تقارير أيام العمل

```sql
-- إجمالي التكاليف حسب المشروع
SELECT 
    p.name as project_name,
    COUNT(wd.id) as work_days_count,
    SUM(wd.total_labor_cost) as total_labor,
    SUM(wd.total_equipment_cost) as total_equipment,
    SUM(wd.total_materials_cost) as total_materials,
    SUM(wd.total_cost) as grand_total
FROM projects p
LEFT JOIN work_days wd ON p.id = wd.project_id
GROUP BY p.id, p.name;
```

## 🔒 الأمان

### نصائح أمنية

1. **كلمات المرور**
   - استخدم bcrypt أو argon2 لتشفير كلمات المرور
   - لا تحفظ كلمات المرور بنص واضح أبداً
   - فرض سياسة كلمات مرور قوية

2. **الصلاحيات**
   - استخدم مبدأ الحد الأدنى من الصلاحيات
   - راجع الصلاحيات بانتظام
   - سجل جميع العمليات الحساسة

3. **النسخ الاحتياطي**
   - احتفظ بنسخ احتياطية يومية
   - اختبر عملية الاستعادة بانتظام
   - احفظ النسخ في مواقع متعددة

4. **مراقبة الأداء**
   - راقب الاستعلامات البطيئة
   - حدث الإحصائيات بانتظام
   - راجع خطط التنفيذ للاستعلامات

## 🔄 الصيانة

### النسخ الاحتياطي

```bash
# نسخة احتياطية كاملة
mysqldump -u username -p project_management_db > backup_$(date +%Y%m%d).sql

# نسخة احتياطية للبيانات فقط
mysqldump -u username -p --no-create-info project_management_db > data_backup_$(date +%Y%m%d).sql

# نسخة احتياطية للمخطط فقط
mysqldump -u username -p --no-data project_management_db > schema_backup_$(date +%Y%m%d).sql
```

### الاستعادة

```bash
mysql -u username -p project_management_db < backup_20240101.sql
```

### تحسين الأداء

```sql
-- تحديث إحصائيات الجداول
ANALYZE TABLE projects, tasks, expenses, income;

-- تحسين الجداول
OPTIMIZE TABLE projects, tasks, expenses, income;

-- إعادة بناء الفهارس
ALTER TABLE projects ENGINE=InnoDB;
```

## 📝 ملاحظات مهمة

1. **الترميز (Encoding)**
   - جميع الجداول تستخدم `utf8mb4` لدعم الأحرف العربية بشكل كامل
   - `utf8mb4_unicode_ci` للمقارنة غير الحساسة لحالة الأحرف

2. **المحرك (Engine)**
   - جميع الجداول تستخدم `InnoDB` لدعم المعاملات (Transactions)
   - دعم المفاتيح الأجنبية (Foreign Keys)

3. **التواريخ**
   - استخدم `TIMESTAMP` للتواريخ التي تحتاج توقيت (created_at, updated_at)
   - استخدم `DATE` للتواريخ فقط (birth_date, due_date)
   - استخدم `TIME` للوقت فقط (work_period_from, work_period_to)

4. **الأرقام المالية**
   - استخدم `DECIMAL(15, 2)` للمبالغ المالية
   - تجنب `FLOAT` أو `DOUBLE` للقيم المالية

## 🆘 استكشاف الأخطاء

### مشاكل شائعة

#### 1. خطأ في الترميز
```sql
-- التحقق من ترميز قاعدة البيانات
SHOW CREATE DATABASE project_management_db;

-- تغيير الترميز
ALTER DATABASE project_management_db 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;
```

#### 2. خطأ في المفاتيح الأجنبية
```sql
-- تعطيل فحص المفاتيح الأجنبية مؤقتاً
SET FOREIGN_KEY_CHECKS = 0;
-- تنفيذ العمليات
SET FOREIGN_KEY_CHECKS = 1;
```

#### 3. بطء الاستعلامات
```sql
-- تحليل الاستعلام
EXPLAIN SELECT * FROM view_project_summary WHERE status = 'active';

-- عرض الاستعلامات البطيئة
SHOW PROCESSLIST;
```

## 📞 الدعم

للمساعدة والدعم:
- البريد الإلكتروني: support@erticaz.com
- الموقع: www.erticaz.com

## 📄 الترخيص

© 2024 ارتكاز للحلول البرمجية - جميع الحقوق محفوظة

---

**تم إنشاء قاعدة البيانات بواسطة:** فريق ارتكاز للحلول البرمجية  
**التاريخ:** 2024  
**الإصدار:** 1.0.0

