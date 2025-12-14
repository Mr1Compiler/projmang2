-- ====================================
-- نصوص الصيانة والتحسين
-- Maintenance and Optimization Scripts
-- ====================================

-- ====================================
-- 1. تحليل الجداول
-- Analyze Tables
-- ====================================

ANALYZE TABLE users;
ANALYZE TABLE roles;
ANALYZE TABLE permissions;
ANALYZE TABLE departments;
ANALYZE TABLE team_members;
ANALYZE TABLE projects;
ANALYZE TABLE tasks;
ANALYZE TABLE expenses;
ANALYZE TABLE income;
ANALYZE TABLE debtors;
ANALYZE TABLE work_days;
ANALYZE TABLE work_day_labor;
ANALYZE TABLE work_day_equipment;
ANALYZE TABLE work_day_materials;

-- ====================================
-- 2. تحسين الجداول
-- Optimize Tables
-- ====================================

OPTIMIZE TABLE users;
OPTIMIZE TABLE roles;
OPTIMIZE TABLE permissions;
OPTIMIZE TABLE departments;
OPTIMIZE TABLE team_members;
OPTIMIZE TABLE projects;
OPTIMIZE TABLE tasks;
OPTIMIZE TABLE expenses;
OPTIMIZE TABLE income;
OPTIMIZE TABLE debtors;
OPTIMIZE TABLE work_days;
OPTIMIZE TABLE work_day_labor;
OPTIMIZE TABLE work_day_equipment;
OPTIMIZE TABLE work_day_materials;

-- ====================================
-- 3. إصلاح الجداول
-- Repair Tables
-- ====================================

REPAIR TABLE users;
REPAIR TABLE projects;
REPAIR TABLE tasks;
REPAIR TABLE expenses;
REPAIR TABLE income;

-- ====================================
-- 4. فحص الجداول
-- Check Tables
-- ====================================

CHECK TABLE users;
CHECK TABLE projects;
CHECK TABLE tasks;
CHECK TABLE expenses;
CHECK TABLE income;

-- ====================================
-- 5. عرض حالة الجداول
-- Show Table Status
-- ====================================

SELECT 
    TABLE_NAME,
    ENGINE,
    TABLE_ROWS,
    ROUND(DATA_LENGTH / 1024 / 1024, 2) AS 'Data Size (MB)',
    ROUND(INDEX_LENGTH / 1024 / 1024, 2) AS 'Index Size (MB)',
    ROUND((DATA_LENGTH + INDEX_LENGTH) / 1024 / 1024, 2) AS 'Total Size (MB)',
    TABLE_COLLATION
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'project_management_db'
ORDER BY (DATA_LENGTH + INDEX_LENGTH) DESC;

-- ====================================
-- 6. عرض الفهارس المستخدمة
-- Show Index Usage
-- ====================================

SELECT 
    TABLE_NAME,
    INDEX_NAME,
    SEQ_IN_INDEX,
    COLUMN_NAME,
    CARDINALITY,
    INDEX_TYPE
FROM information_schema.STATISTICS
WHERE TABLE_SCHEMA = 'project_management_db'
ORDER BY TABLE_NAME, INDEX_NAME, SEQ_IN_INDEX;

-- ====================================
-- 7. البحث عن الجداول الكبيرة
-- Find Large Tables
-- ====================================

SELECT 
    TABLE_NAME,
    ROUND((DATA_LENGTH + INDEX_LENGTH) / 1024 / 1024, 2) AS 'Size (MB)',
    TABLE_ROWS
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'project_management_db'
    AND (DATA_LENGTH + INDEX_LENGTH) > 10485760  -- أكبر من 10 MB
ORDER BY (DATA_LENGTH + INDEX_LENGTH) DESC;

-- ====================================
-- 8. عرض الاستعلامات البطيئة
-- Show Slow Queries (if enabled)
-- ====================================

-- تفعيل سجل الاستعلامات البطيئة
SET GLOBAL slow_query_log = 'ON';
SET GLOBAL long_query_time = 2;
SET GLOBAL log_queries_not_using_indexes = 'ON';

-- عرض إعدادات السجل
SHOW VARIABLES LIKE 'slow_query%';
SHOW VARIABLES LIKE 'long_query_time';

-- ====================================
-- 9. تنظيف البيانات القديمة
-- Cleanup Old Data
-- ====================================

-- حذف المهام الملغاة القديمة (أقدم من سنة)
DELETE FROM tasks 
WHERE status = 'cancelled' 
    AND created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);

-- حذف المصاريف المرفوضة القديمة
DELETE FROM expenses 
WHERE status = 'rejected' 
    AND created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);

-- حذف تعليقات المهام المحذوفة
DELETE FROM task_comments 
WHERE task_id NOT IN (SELECT id FROM tasks);

-- ====================================
-- 10. إعادة بناء الفهارس
-- Rebuild Indexes
-- ====================================

ALTER TABLE projects ENGINE=InnoDB;
ALTER TABLE tasks ENGINE=InnoDB;
ALTER TABLE expenses ENGINE=InnoDB;
ALTER TABLE income ENGINE=InnoDB;
ALTER TABLE work_days ENGINE=InnoDB;

-- ====================================
-- 11. إحصائيات قاعدة البيانات
-- Database Statistics
-- ====================================

SELECT 
    'إجمالي المستخدمين' as metric,
    COUNT(*) as value
FROM users

UNION ALL

SELECT 
    'إجمالي المشاريع' as metric,
    COUNT(*) as value
FROM projects

UNION ALL

SELECT 
    'إجمالي المهام' as metric,
    COUNT(*) as value
FROM tasks

UNION ALL

SELECT 
    'إجمالي المصاريف' as metric,
    COUNT(*) as value
FROM expenses

UNION ALL

SELECT 
    'إجمالي الإيرادات' as metric,
    COUNT(*) as value
FROM income

UNION ALL

SELECT 
    'إجمالي أعضاء الفريق' as metric,
    COUNT(*) as value
FROM team_members

UNION ALL

SELECT 
    'إجمالي أيام العمل' as metric,
    COUNT(*) as value
FROM work_days;

-- ====================================
-- 12. تحديث الإحصائيات
-- Update Statistics
-- ====================================

ANALYZE TABLE 
    users, roles, permissions, user_roles, role_permissions,
    departments, team_members,
    projects, project_team, project_expenses,
    tasks, task_comments, task_attachments,
    expenses, expense_types, expense_categories,
    income, income_categories,
    debtors, debtor_payments, payment_installments,
    work_days, work_day_labor, work_day_equipment, work_day_materials;

-- ====================================
-- 13. عرض متغيرات النظام المهمة
-- Show Important System Variables
-- ====================================

SHOW VARIABLES LIKE 'innodb_buffer_pool_size';
SHOW VARIABLES LIKE 'innodb_log_file_size';
SHOW VARIABLES LIKE 'max_connections';
SHOW VARIABLES LIKE 'query_cache_size';
SHOW VARIABLES LIKE 'tmp_table_size';
SHOW VARIABLES LIKE 'max_heap_table_size';

-- ====================================
-- 14. عرض حالة النظام
-- Show System Status
-- ====================================

SHOW STATUS LIKE 'Threads_connected';
SHOW STATUS LIKE 'Threads_running';
SHOW STATUS LIKE 'Questions';
SHOW STATUS LIKE 'Slow_queries';
SHOW STATUS LIKE 'Uptime';

-- ====================================
-- 15. تحسين الذاكرة
-- Memory Optimization
-- ====================================

-- مسح ذاكرة التخزين المؤقت
RESET QUERY CACHE;
FLUSH QUERY CACHE;

-- مسح السجلات
FLUSH LOGS;

-- إعادة تحميل الصلاحيات
FLUSH PRIVILEGES;

-- ====================================
-- نهاية نصوص الصيانة
-- End of Maintenance Scripts
-- ====================================

