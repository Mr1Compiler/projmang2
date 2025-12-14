-- ====================================
-- الاستعلامات الشائعة
-- Common Queries
-- ====================================

-- ====================================
-- 1. استعلامات المشاريع
-- Project Queries
-- ====================================

-- جميع المشاريع النشطة
SELECT * FROM projects WHERE status = 'in_progress' ORDER BY start_date DESC;

-- المشاريع المتأخرة
SELECT 
    p.*,
    DATEDIFF(CURRENT_DATE, p.end_date) as days_overdue
FROM projects p
WHERE p.end_date < CURRENT_DATE 
    AND p.status NOT IN ('completed', 'cancelled')
ORDER BY days_overdue DESC;

-- المشاريع حسب التقدم
SELECT 
    id,
    project_number,
    name,
    progress_percentage,
    status
FROM projects
WHERE status = 'in_progress'
ORDER BY progress_percentage ASC;

-- ملخص مشروع محدد
SELECT * FROM view_project_summary WHERE id = 1;

-- أعلى 5 مشاريع من حيث التكلفة
SELECT 
    name,
    contract_value,
    actual_cost,
    (contract_value - actual_cost) as profit
FROM projects
WHERE status = 'completed'
ORDER BY actual_cost DESC
LIMIT 5;

-- ====================================
-- 2. استعلامات المهام
-- Task Queries
-- ====================================

-- المهام المعينة لعضو معين
SELECT 
    t.*,
    p.name as project_name
FROM tasks t
LEFT JOIN projects p ON t.project_id = p.id
WHERE t.assigned_to = 1
    AND t.status NOT IN ('completed', 'cancelled')
ORDER BY t.due_date ASC;

-- المهام المتأخرة
SELECT 
    t.id,
    t.title,
    t.due_date,
    DATEDIFF(CURRENT_DATE, t.due_date) as days_overdue,
    tm.name as assigned_to,
    p.name as project_name
FROM tasks t
LEFT JOIN team_members tm ON t.assigned_to = tm.id
LEFT JOIN projects p ON t.project_id = p.id
WHERE t.due_date < CURRENT_DATE
    AND t.status NOT IN ('completed', 'cancelled')
ORDER BY days_overdue DESC;

-- المهام حسب الحالة
SELECT 
    status,
    COUNT(*) as count,
    AVG(progress_percentage) as avg_progress
FROM tasks
GROUP BY status;

-- المهام العالية الأولوية
SELECT 
    t.id,
    t.title,
    t.priority,
    t.due_date,
    tm.name as assigned_to,
    p.name as project_name
FROM tasks t
LEFT JOIN team_members tm ON t.assigned_to = tm.id
LEFT JOIN projects p ON t.project_id = p.id
WHERE t.priority IN ('high', 'urgent')
    AND t.status NOT IN ('completed', 'cancelled')
ORDER BY 
    FIELD(t.priority, 'urgent', 'high'),
    t.due_date ASC;

-- ====================================
-- 3. استعلامات الفريق
-- Team Queries
-- ====================================

-- أعضاء الفريق النشطين
SELECT * FROM team_members WHERE status = 'active' ORDER BY name;

-- أعضاء الفريق حسب القسم
SELECT 
    d.name as department,
    COUNT(tm.id) as member_count,
    AVG(tm.salary) as avg_salary
FROM departments d
LEFT JOIN team_members tm ON d.id = tm.department_id
WHERE tm.status = 'active'
GROUP BY d.id, d.name
ORDER BY member_count DESC;

-- عبء العمل لكل عضو
SELECT 
    tm.id,
    tm.name,
    COUNT(t.id) as total_tasks,
    SUM(CASE WHEN t.status = 'in_progress' THEN 1 ELSE 0 END) as active_tasks,
    SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) as completed_tasks,
    AVG(t.progress_percentage) as avg_progress
FROM team_members tm
LEFT JOIN tasks t ON tm.id = t.assigned_to
WHERE tm.status = 'active'
GROUP BY tm.id, tm.name
ORDER BY active_tasks DESC;

-- أفضل أداء (معدل الإنجاز)
SELECT 
    tm.name,
    COUNT(t.id) as total_tasks,
    SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) as completed_tasks,
    ROUND((SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) / COUNT(t.id)) * 100, 2) as completion_rate
FROM team_members tm
JOIN tasks t ON tm.id = t.assigned_to
WHERE tm.status = 'active'
    AND t.created_at >= DATE_SUB(CURRENT_DATE, INTERVAL 3 MONTH)
GROUP BY tm.id, tm.name
HAVING COUNT(t.id) >= 5
ORDER BY completion_rate DESC
LIMIT 10;

-- ====================================
-- 4. استعلامات المصاريف
-- Expense Queries
-- ====================================

-- المصاريف الشهرية
SELECT 
    DATE_FORMAT(expense_date, '%Y-%m') as month,
    COUNT(*) as expense_count,
    SUM(amount) as total_amount
FROM expenses
WHERE status = 'paid'
GROUP BY DATE_FORMAT(expense_date, '%Y-%m')
ORDER BY month DESC
LIMIT 12;

-- المصاريف حسب النوع
SELECT 
    et.name as expense_type,
    COUNT(e.id) as count,
    SUM(e.amount) as total
FROM expenses e
LEFT JOIN expense_types et ON e.expense_type_id = et.id
WHERE e.status = 'paid'
    AND e.expense_date >= DATE_SUB(CURRENT_DATE, INTERVAL 1 YEAR)
GROUP BY et.id, et.name
ORDER BY total DESC;

-- المصاريف حسب المشروع
SELECT 
    p.name as project,
    COUNT(e.id) as expense_count,
    SUM(e.amount) as total_expenses
FROM projects p
JOIN project_expenses pe ON p.id = pe.project_id
JOIN expenses e ON pe.expense_id = e.id
WHERE e.status = 'paid'
GROUP BY p.id, p.name
ORDER BY total_expenses DESC;

-- أعلى 10 مصاريف
SELECT 
    e.expense_number,
    e.description,
    e.amount,
    e.expense_date,
    et.name as type,
    p.name as project
FROM expenses e
LEFT JOIN expense_types et ON e.expense_type_id = et.id
LEFT JOIN project_expenses pe ON e.id = pe.expense_id
LEFT JOIN projects p ON pe.project_id = p.id
WHERE e.status = 'paid'
ORDER BY e.amount DESC
LIMIT 10;

-- ====================================
-- 5. استعلامات الإيرادات
-- Income Queries
-- ====================================

-- الإيرادات الشهرية
SELECT 
    DATE_FORMAT(income_date, '%Y-%m') as month,
    COUNT(*) as income_count,
    SUM(amount) as total_amount
FROM income
WHERE status = 'confirmed'
GROUP BY DATE_FORMAT(income_date, '%Y-%m')
ORDER BY month DESC
LIMIT 12;

-- الإيرادات حسب المشروع
SELECT 
    p.name as project,
    COUNT(i.id) as income_count,
    SUM(i.amount) as total_income
FROM projects p
LEFT JOIN income i ON p.id = i.project_id
WHERE i.status = 'confirmed'
GROUP BY p.id, p.name
ORDER BY total_income DESC;

-- مقارنة الإيرادات والمصاريف الشهرية
SELECT 
    months.month,
    COALESCE(income_data.total_income, 0) as income,
    COALESCE(expense_data.total_expenses, 0) as expenses,
    COALESCE(income_data.total_income, 0) - COALESCE(expense_data.total_expenses, 0) as profit
FROM (
    SELECT DISTINCT DATE_FORMAT(income_date, '%Y-%m') as month FROM income
    UNION
    SELECT DISTINCT DATE_FORMAT(expense_date, '%Y-%m') FROM expenses
) months
LEFT JOIN (
    SELECT 
        DATE_FORMAT(income_date, '%Y-%m') as month,
        SUM(amount) as total_income
    FROM income
    WHERE status = 'confirmed'
    GROUP BY DATE_FORMAT(income_date, '%Y-%m')
) income_data ON months.month = income_data.month
LEFT JOIN (
    SELECT 
        DATE_FORMAT(expense_date, '%Y-%m') as month,
        SUM(amount) as total_expenses
    FROM expenses
    WHERE status = 'paid'
    GROUP BY DATE_FORMAT(expense_date, '%Y-%m')
) expense_data ON months.month = expense_data.month
ORDER BY months.month DESC
LIMIT 12;

-- ====================================
-- 6. استعلامات المدينين
-- Debtor Queries
-- ====================================

-- المدينين النشطين
SELECT * FROM debtors WHERE status IN ('active', 'partially_paid') ORDER BY due_date;

-- المدينين المتأخرين
SELECT 
    d.*,
    DATEDIFF(CURRENT_DATE, d.due_date) as days_overdue
FROM debtors d
WHERE d.due_date < CURRENT_DATE
    AND d.status NOT IN ('fully_paid', 'cancelled')
ORDER BY days_overdue DESC;

-- أعلى 10 ديون
SELECT 
    name,
    company_name,
    total_debt,
    paid_amount,
    remaining_amount,
    status
FROM debtors
ORDER BY remaining_amount DESC
LIMIT 10;

-- ملخص الديون
SELECT 
    COUNT(*) as total_debtors,
    SUM(total_debt) as total_debt,
    SUM(paid_amount) as total_paid,
    SUM(remaining_amount) as total_remaining
FROM debtors
WHERE status NOT IN ('cancelled');

-- ====================================
-- 7. استعلامات أيام العمل
-- Work Day Queries
-- ====================================

-- أيام العمل حسب المشروع
SELECT 
    p.name as project,
    COUNT(wd.id) as work_days,
    SUM(wd.total_cost) as total_cost
FROM work_days wd
JOIN projects p ON wd.project_id = p.id
WHERE wd.status = 'completed'
GROUP BY p.id, p.name
ORDER BY total_cost DESC;

-- تفاصيل يوم عمل محدد
SELECT 
    wd.*,
    p.name as project_name,
    (SELECT COUNT(*) FROM work_day_labor WHERE work_day_id = wd.id) as labor_count,
    (SELECT COUNT(*) FROM work_day_equipment WHERE work_day_id = wd.id) as equipment_count,
    (SELECT COUNT(*) FROM work_day_materials WHERE work_day_id = wd.id) as materials_count
FROM work_days wd
LEFT JOIN projects p ON wd.project_id = p.id
WHERE wd.id = 1;

-- ملخص تكاليف أيام العمل
SELECT 
    DATE_FORMAT(work_date, '%Y-%m') as month,
    COUNT(*) as work_days_count,
    SUM(total_labor_cost) as labor_cost,
    SUM(total_equipment_cost) as equipment_cost,
    SUM(total_materials_cost) as materials_cost,
    SUM(total_cost) as total_cost
FROM work_days
WHERE status = 'completed'
GROUP BY DATE_FORMAT(work_date, '%Y-%m')
ORDER BY month DESC
LIMIT 12;

-- ====================================
-- 8. التقارير المالية
-- Financial Reports
-- ====================================

-- التقرير المالي الشامل
SELECT 
    'الإيرادات المؤكدة' as category,
    COALESCE(SUM(amount), 0) as amount
FROM income
WHERE status = 'confirmed'

UNION ALL

SELECT 
    'المصاريف المدفوعة' as category,
    COALESCE(SUM(amount), 0) as amount
FROM expenses
WHERE status = 'paid'

UNION ALL

SELECT 
    'الديون المستحقة' as category,
    COALESCE(SUM(remaining_amount), 0) as amount
FROM debtors
WHERE status IN ('active', 'partially_paid', 'overdue')

UNION ALL

SELECT 
    'صافي الربح' as category,
    (SELECT COALESCE(SUM(amount), 0) FROM income WHERE status = 'confirmed') -
    (SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE status = 'paid') as amount;

-- الميزانية العمومية
SELECT 
    d.name as department,
    d.budget,
    COALESCE(SUM(e.amount), 0) as spent,
    d.budget - COALESCE(SUM(e.amount), 0) as remaining,
    ROUND((COALESCE(SUM(e.amount), 0) / d.budget) * 100, 2) as spent_percentage
FROM departments d
LEFT JOIN expenses e ON d.id = e.department_id AND e.status = 'paid'
WHERE d.status = 'active'
GROUP BY d.id, d.name, d.budget
ORDER BY spent_percentage DESC;

-- ====================================
-- 9. لوحة التحكم
-- Dashboard
-- ====================================

-- إحصائيات سريعة
SELECT 
    (SELECT COUNT(*) FROM projects WHERE status = 'in_progress') as active_projects,
    (SELECT COUNT(*) FROM tasks WHERE status IN ('pending', 'in_progress')) as active_tasks,
    (SELECT COUNT(*) FROM team_members WHERE status = 'active') as active_team_members,
    (SELECT SUM(amount) FROM income WHERE status = 'confirmed' AND MONTH(income_date) = MONTH(CURRENT_DATE)) as monthly_income,
    (SELECT SUM(amount) FROM expenses WHERE status = 'paid' AND MONTH(expense_date) = MONTH(CURRENT_DATE)) as monthly_expenses;

-- المشاريع القادمة
SELECT 
    id,
    project_number,
    name,
    start_date,
    manager_id
FROM projects
WHERE status = 'pending'
    AND start_date > CURRENT_DATE
ORDER BY start_date ASC
LIMIT 5;

-- الأنشطة الأخيرة
(SELECT 
    'مشروع' as type,
    name as title,
    created_at
FROM projects
ORDER BY created_at DESC
LIMIT 5)
UNION ALL
(SELECT 
    'مهمة' as type,
    title,
    created_at
FROM tasks
ORDER BY created_at DESC
LIMIT 5)
UNION ALL
(SELECT 
    'مصروف' as type,
    description,
    created_at
FROM expenses
ORDER BY created_at DESC
LIMIT 5)
ORDER BY created_at DESC
LIMIT 10;

-- ====================================
-- 10. البحث والفلترة
-- Search and Filter
-- ====================================

-- البحث في المشاريع
SELECT * FROM projects 
WHERE name LIKE '%keyword%' 
    OR description LIKE '%keyword%'
    OR project_number LIKE '%keyword%';

-- البحث في المهام
SELECT * FROM tasks 
WHERE title LIKE '%keyword%' 
    OR description LIKE '%keyword%'
    OR task_number LIKE '%keyword%';

-- البحث في أعضاء الفريق
SELECT * FROM team_members 
WHERE name LIKE '%keyword%' 
    OR email LIKE '%keyword%'
    OR phone LIKE '%keyword%';

-- فلترة متقدمة للمشاريع
SELECT * FROM projects 
WHERE status = 'in_progress'
    AND progress_percentage < 50
    AND end_date BETWEEN '2024-01-01' AND '2024-12-31'
    AND department_id = 1
ORDER BY end_date ASC;

-- ====================================
-- نهاية الاستعلامات الشائعة
-- End of Common Queries
-- ====================================

