-- ====================================
-- الإجراءات المخزنة
-- Stored Procedures
-- ====================================

DELIMITER //

-- ====================================
-- 1. إجراء لحساب تكاليف المشروع
-- Procedure to calculate project costs
-- ====================================

DROP PROCEDURE IF EXISTS calculate_project_costs //

CREATE PROCEDURE calculate_project_costs(IN p_project_id INT)
BEGIN
    DECLARE total_expenses DECIMAL(15, 2);
    DECLARE total_work_days_cost DECIMAL(15, 2);
    
    -- حساب المصاريف المباشرة
    SELECT COALESCE(SUM(e.amount), 0) INTO total_expenses
    FROM project_expenses pe
    JOIN expenses e ON pe.expense_id = e.id
    WHERE pe.project_id = p_project_id;
    
    -- حساب تكاليف أيام العمل
    SELECT COALESCE(SUM(total_cost), 0) INTO total_work_days_cost
    FROM work_days
    WHERE project_id = p_project_id;
    
    -- تحديث التكلفة الفعلية للمشروع
    UPDATE projects
    SET actual_cost = total_expenses + total_work_days_cost
    WHERE id = p_project_id;
    
    -- إرجاع النتائج
    SELECT 
        total_expenses,
        total_work_days_cost,
        total_expenses + total_work_days_cost as total_cost;
END //

-- ====================================
-- 2. إجراء لحساب تقدم المشروع
-- Procedure to calculate project progress
-- ====================================

DROP PROCEDURE IF EXISTS calculate_project_progress //

CREATE PROCEDURE calculate_project_progress(IN p_project_id INT)
BEGIN
    DECLARE total_tasks INT;
    DECLARE completed_tasks INT;
    DECLARE progress DECIMAL(5, 2);
    
    -- عدد المهام الإجمالي
    SELECT COUNT(*) INTO total_tasks
    FROM tasks
    WHERE project_id = p_project_id;
    
    -- عدد المهام المكتملة
    SELECT COUNT(*) INTO completed_tasks
    FROM tasks
    WHERE project_id = p_project_id AND status = 'completed';
    
    -- حساب النسبة المئوية
    IF total_tasks > 0 THEN
        SET progress = (completed_tasks / total_tasks) * 100;
    ELSE
        SET progress = 0;
    END IF;
    
    -- تحديث تقدم المشروع
    UPDATE projects
    SET progress_percentage = progress
    WHERE id = p_project_id;
    
    -- إرجاع النتائج
    SELECT total_tasks, completed_tasks, progress;
END //

-- ====================================
-- 3. إجراء لتحديث حالة المدين
-- Procedure to update debtor status
-- ====================================

DROP PROCEDURE IF EXISTS update_debtor_status //

CREATE PROCEDURE update_debtor_status(IN p_debtor_id INT)
BEGIN
    DECLARE total DECIMAL(15, 2);
    DECLARE paid DECIMAL(15, 2);
    DECLARE remaining DECIMAL(15, 2);
    DECLARE new_status VARCHAR(20);
    
    -- حساب المبالغ
    SELECT 
        d.total_debt,
        COALESCE(SUM(dp.amount), 0)
    INTO total, paid
    FROM debtors d
    LEFT JOIN debtor_payments dp ON d.id = dp.debtor_id
    WHERE d.id = p_debtor_id
    GROUP BY d.id, d.total_debt;
    
    SET remaining = total - paid;
    
    -- تحديد الحالة الجديدة
    IF remaining <= 0 THEN
        SET new_status = 'fully_paid';
    ELSEIF paid > 0 THEN
        SET new_status = 'partially_paid';
    ELSEIF CURRENT_DATE > (SELECT due_date FROM debtors WHERE id = p_debtor_id) THEN
        SET new_status = 'overdue';
    ELSE
        SET new_status = 'active';
    END IF;
    
    -- تحديث المدين
    UPDATE debtors
    SET 
        paid_amount = paid,
        remaining_amount = remaining,
        status = new_status
    WHERE id = p_debtor_id;
    
    -- إرجاع النتائج
    SELECT total, paid, remaining, new_status;
END //

-- ====================================
-- 4. إجراء لحساب تكاليف يوم العمل
-- Procedure to calculate work day costs
-- ====================================

DROP PROCEDURE IF EXISTS calculate_work_day_costs //

CREATE PROCEDURE calculate_work_day_costs(IN p_work_day_id INT)
BEGIN
    DECLARE labor_cost DECIMAL(15, 2);
    DECLARE equipment_cost DECIMAL(15, 2);
    DECLARE materials_cost DECIMAL(15, 2);
    
    -- حساب تكلفة العمالة
    SELECT COALESCE(SUM(total_cost), 0) INTO labor_cost
    FROM work_day_labor
    WHERE work_day_id = p_work_day_id;
    
    -- حساب تكلفة المعدات
    SELECT COALESCE(SUM(total_cost), 0) INTO equipment_cost
    FROM work_day_equipment
    WHERE work_day_id = p_work_day_id;
    
    -- حساب تكلفة المواد
    SELECT COALESCE(SUM(total_cost), 0) INTO materials_cost
    FROM work_day_materials
    WHERE work_day_id = p_work_day_id;
    
    -- تحديث يوم العمل
    UPDATE work_days
    SET 
        total_labor_cost = labor_cost,
        total_equipment_cost = equipment_cost,
        total_materials_cost = materials_cost,
        total_cost = labor_cost + equipment_cost + materials_cost
    WHERE id = p_work_day_id;
    
    -- إرجاع النتائج
    SELECT labor_cost, equipment_cost, materials_cost, 
           (labor_cost + equipment_cost + materials_cost) as total_cost;
END //

-- ====================================
-- 5. إجراء للتقارير المالية الشهرية
-- Procedure for monthly financial reports
-- ====================================

DROP PROCEDURE IF EXISTS get_monthly_financial_report //

CREATE PROCEDURE get_monthly_financial_report(
    IN p_year INT,
    IN p_month INT
)
BEGIN
    SELECT 
        'الإيرادات' as category,
        COALESCE(SUM(amount), 0) as total,
        COUNT(*) as count
    FROM income
    WHERE YEAR(income_date) = p_year 
        AND MONTH(income_date) = p_month
        AND status = 'confirmed'
    
    UNION ALL
    
    SELECT 
        'المصاريف' as category,
        COALESCE(SUM(amount), 0) as total,
        COUNT(*) as count
    FROM expenses
    WHERE YEAR(expense_date) = p_year 
        AND MONTH(expense_date) = p_month
        AND status = 'paid'
    
    UNION ALL
    
    SELECT 
        'صافي الربح' as category,
        (
            SELECT COALESCE(SUM(amount), 0)
            FROM income
            WHERE YEAR(income_date) = p_year 
                AND MONTH(income_date) = p_month
                AND status = 'confirmed'
        ) - (
            SELECT COALESCE(SUM(amount), 0)
            FROM expenses
            WHERE YEAR(expense_date) = p_year 
                AND MONTH(expense_date) = p_month
                AND status = 'paid'
        ) as total,
        NULL as count;
END //

-- ====================================
-- 6. إجراء لتعيين مهمة لعضو الفريق
-- Procedure to assign task to team member
-- ====================================

DROP PROCEDURE IF EXISTS assign_task //

CREATE PROCEDURE assign_task(
    IN p_task_id INT,
    IN p_team_member_id INT,
    IN p_assigned_by INT
)
BEGIN
    -- التحقق من أن المهمة موجودة
    IF NOT EXISTS (SELECT 1 FROM tasks WHERE id = p_task_id) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'المهمة غير موجودة';
    END IF;
    
    -- التحقق من أن عضو الفريق موجود ونشط
    IF NOT EXISTS (
        SELECT 1 FROM team_members 
        WHERE id = p_team_member_id AND status = 'active'
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'عضو الفريق غير موجود أو غير نشط';
    END IF;
    
    -- تعيين المهمة
    UPDATE tasks
    SET 
        assigned_to = p_team_member_id,
        assigned_by = p_assigned_by,
        status = 'pending'
    WHERE id = p_task_id;
    
    SELECT 'تم تعيين المهمة بنجاح' as message;
END //

-- ====================================
-- 7. إجراء لإضافة دفعة مدين
-- Procedure to add debtor payment
-- ====================================

DROP PROCEDURE IF EXISTS add_debtor_payment //

CREATE PROCEDURE add_debtor_payment(
    IN p_debtor_id INT,
    IN p_amount DECIMAL(15, 2),
    IN p_payment_date DATE,
    IN p_payment_method VARCHAR(20),
    IN p_receipt_number VARCHAR(100),
    IN p_notes TEXT,
    IN p_created_by INT
)
BEGIN
    DECLARE payment_number VARCHAR(100);
    DECLARE new_payment_id INT;
    
    -- إنشاء رقم الدفعة
    SET payment_number = CONCAT('PAY-', YEAR(p_payment_date), '-', 
                                LPAD(FLOOR(RAND() * 10000), 4, '0'));
    
    -- إضافة الدفعة
    INSERT INTO debtor_payments (
        debtor_id, payment_number, amount, payment_date, 
        payment_method, receipt_number, notes, created_by
    ) VALUES (
        p_debtor_id, payment_number, p_amount, p_payment_date,
        p_payment_method, p_receipt_number, p_notes, p_created_by
    );
    
    SET new_payment_id = LAST_INSERT_ID();
    
    -- تحديث حالة المدين
    CALL update_debtor_status(p_debtor_id);
    
    -- إرجاع معلومات الدفعة
    SELECT * FROM debtor_payments WHERE id = new_payment_id;
END //

-- ====================================
-- 8. إجراء لإحصائيات لوحة التحكم
-- Procedure for dashboard statistics
-- ====================================

DROP PROCEDURE IF EXISTS get_dashboard_stats //

CREATE PROCEDURE get_dashboard_stats()
BEGIN
    -- إحصائيات المشاريع
    SELECT 
        'المشاريع' as category,
        COUNT(*) as total,
        SUM(CASE WHEN status = 'in_progress' THEN 1 ELSE 0 END) as in_progress,
        SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as completed,
        SUM(CASE WHEN status = 'pending' THEN 1 ELSE 0 END) as pending,
        AVG(progress_percentage) as avg_progress
    FROM projects
    
    UNION ALL
    
    -- إحصائيات المهام
    SELECT 
        'المهام' as category,
        COUNT(*) as total,
        SUM(CASE WHEN status = 'in_progress' THEN 1 ELSE 0 END) as in_progress,
        SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as completed,
        SUM(CASE WHEN status = 'pending' THEN 1 ELSE 0 END) as pending,
        AVG(progress_percentage) as avg_progress
    FROM tasks
    
    UNION ALL
    
    -- إحصائيات الفريق
    SELECT 
        'أعضاء الفريق' as category,
        COUNT(*) as total,
        SUM(CASE WHEN status = 'active' THEN 1 ELSE 0 END) as in_progress,
        NULL as completed,
        NULL as pending,
        NULL as avg_progress
    FROM team_members;
END //

-- ====================================
-- 9. إجراء لتقرير المشاريع المتأخرة
-- Procedure for overdue projects report
-- ====================================

DROP PROCEDURE IF EXISTS get_overdue_projects //

CREATE PROCEDURE get_overdue_projects()
BEGIN
    SELECT 
        p.id,
        p.project_number,
        p.name,
        p.end_date,
        DATEDIFF(CURRENT_DATE, p.end_date) as days_overdue,
        p.progress_percentage,
        u.full_name as manager_name,
        d.name as department_name
    FROM projects p
    LEFT JOIN users u ON p.manager_id = u.id
    LEFT JOIN departments d ON p.department_id = d.id
    WHERE p.end_date < CURRENT_DATE
        AND p.status NOT IN ('completed', 'cancelled')
    ORDER BY days_overdue DESC;
END //

-- ====================================
-- 10. إجراء لتقرير أداء الفريق
-- Procedure for team performance report
-- ====================================

DROP PROCEDURE IF EXISTS get_team_performance //

CREATE PROCEDURE get_team_performance(IN p_period_days INT)
BEGIN
    SELECT 
        tm.id,
        tm.name,
        tm.position,
        d.name as department_name,
        COUNT(t.id) as total_tasks,
        SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) as completed_tasks,
        SUM(CASE WHEN t.status = 'in_progress' THEN 1 ELSE 0 END) as in_progress_tasks,
        SUM(CASE WHEN t.status = 'pending' THEN 1 ELSE 0 END) as pending_tasks,
        AVG(t.progress_percentage) as avg_progress,
        SUM(t.actual_hours) as total_hours_worked
    FROM team_members tm
    LEFT JOIN departments d ON tm.department_id = d.id
    LEFT JOIN tasks t ON tm.id = t.assigned_to 
        AND t.created_at >= DATE_SUB(CURRENT_DATE, INTERVAL p_period_days DAY)
    WHERE tm.status = 'active'
    GROUP BY tm.id, tm.name, tm.position, d.name
    ORDER BY completed_tasks DESC;
END //

DELIMITER ;

-- ====================================
-- أمثلة الاستخدام
-- Usage Examples
-- ====================================

/*
-- حساب تكاليف مشروع
CALL calculate_project_costs(1);

-- حساب تقدم مشروع
CALL calculate_project_progress(1);

-- تحديث حالة مدين
CALL update_debtor_status(1);

-- حساب تكاليف يوم عمل
CALL calculate_work_day_costs(1);

-- التقرير المالي الشهري
CALL get_monthly_financial_report(2024, 10);

-- تعيين مهمة
CALL assign_task(1, 2, 1);

-- إضافة دفعة مدين
CALL add_debtor_payment(1, 500000, '2024-10-21', 'cash', 'RCP-001', 'دفعة نقدية', 1);

-- إحصائيات لوحة التحكم
CALL get_dashboard_stats();

-- المشاريع المتأخرة
CALL get_overdue_projects();

-- أداء الفريق (آخر 30 يوم)
CALL get_team_performance(30);
*/

