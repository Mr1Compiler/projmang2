-- ====================================
-- المشغلات (Triggers)
-- Database Triggers
-- ====================================

DELIMITER //

-- ====================================
-- 1. مشغل لتحديث تكاليف يوم العمل عند إضافة عمالة
-- Trigger to update work day costs when adding labor
-- ====================================

DROP TRIGGER IF EXISTS after_work_day_labor_insert //

CREATE TRIGGER after_work_day_labor_insert
AFTER INSERT ON work_day_labor
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_labor_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_labor
            WHERE work_day_id = NEW.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = NEW.work_day_id;
END //

DROP TRIGGER IF EXISTS after_work_day_labor_update //

CREATE TRIGGER after_work_day_labor_update
AFTER UPDATE ON work_day_labor
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_labor_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_labor
            WHERE work_day_id = NEW.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = NEW.work_day_id;
END //

DROP TRIGGER IF EXISTS after_work_day_labor_delete //

CREATE TRIGGER after_work_day_labor_delete
AFTER DELETE ON work_day_labor
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_labor_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_labor
            WHERE work_day_id = OLD.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = OLD.work_day_id;
END //

-- ====================================
-- 2. مشغل لتحديث تكاليف يوم العمل عند إضافة معدات
-- Trigger to update work day costs when adding equipment
-- ====================================

DROP TRIGGER IF EXISTS after_work_day_equipment_insert //

CREATE TRIGGER after_work_day_equipment_insert
AFTER INSERT ON work_day_equipment
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_equipment_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_equipment
            WHERE work_day_id = NEW.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = NEW.work_day_id;
END //

DROP TRIGGER IF EXISTS after_work_day_equipment_update //

CREATE TRIGGER after_work_day_equipment_update
AFTER UPDATE ON work_day_equipment
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_equipment_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_equipment
            WHERE work_day_id = NEW.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = NEW.work_day_id;
END //

DROP TRIGGER IF EXISTS after_work_day_equipment_delete //

CREATE TRIGGER after_work_day_equipment_delete
AFTER DELETE ON work_day_equipment
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_equipment_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_equipment
            WHERE work_day_id = OLD.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = OLD.work_day_id;
END //

-- ====================================
-- 3. مشغل لتحديث تكاليف يوم العمل عند إضافة مواد
-- Trigger to update work day costs when adding materials
-- ====================================

DROP TRIGGER IF EXISTS after_work_day_materials_insert //

CREATE TRIGGER after_work_day_materials_insert
AFTER INSERT ON work_day_materials
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_materials_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_materials
            WHERE work_day_id = NEW.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = NEW.work_day_id;
END //

DROP TRIGGER IF EXISTS after_work_day_materials_update //

CREATE TRIGGER after_work_day_materials_update
AFTER UPDATE ON work_day_materials
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_materials_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_materials
            WHERE work_day_id = NEW.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = NEW.work_day_id;
END //

DROP TRIGGER IF EXISTS after_work_day_materials_delete //

CREATE TRIGGER after_work_day_materials_delete
AFTER DELETE ON work_day_materials
FOR EACH ROW
BEGIN
    UPDATE work_days
    SET 
        total_materials_cost = (
            SELECT COALESCE(SUM(total_cost), 0)
            FROM work_day_materials
            WHERE work_day_id = OLD.work_day_id
        ),
        total_cost = total_labor_cost + total_equipment_cost + total_materials_cost
    WHERE id = OLD.work_day_id;
END //

-- ====================================
-- 4. مشغل لتحديث حالة المدين عند إضافة دفعة
-- Trigger to update debtor status when adding payment
-- ====================================

DROP TRIGGER IF EXISTS after_debtor_payment_insert //

CREATE TRIGGER after_debtor_payment_insert
AFTER INSERT ON debtor_payments
FOR EACH ROW
BEGIN
    DECLARE total_paid DECIMAL(15, 2);
    DECLARE total_debt DECIMAL(15, 2);
    DECLARE remaining DECIMAL(15, 2);
    DECLARE new_status VARCHAR(20);
    
    -- حساب المبلغ المدفوع
    SELECT COALESCE(SUM(amount), 0) INTO total_paid
    FROM debtor_payments
    WHERE debtor_id = NEW.debtor_id;
    
    -- الحصول على إجمالي الدين
    SELECT total_debt INTO total_debt
    FROM debtors
    WHERE id = NEW.debtor_id;
    
    SET remaining = total_debt - total_paid;
    
    -- تحديد الحالة
    IF remaining <= 0 THEN
        SET new_status = 'fully_paid';
    ELSE
        SET new_status = 'partially_paid';
    END IF;
    
    -- تحديث المدين
    UPDATE debtors
    SET 
        paid_amount = total_paid,
        remaining_amount = remaining,
        status = new_status
    WHERE id = NEW.debtor_id;
END //

-- ====================================
-- 5. مشغل لتحديث تقدم المشروع عند تحديث مهمة
-- Trigger to update project progress when task is updated
-- ====================================

DROP TRIGGER IF EXISTS after_task_update //

CREATE TRIGGER after_task_update
AFTER UPDATE ON tasks
FOR EACH ROW
BEGIN
    DECLARE total_tasks INT;
    DECLARE completed_tasks INT;
    DECLARE progress DECIMAL(5, 2);
    
    IF NEW.project_id IS NOT NULL THEN
        -- عدد المهام الإجمالي
        SELECT COUNT(*) INTO total_tasks
        FROM tasks
        WHERE project_id = NEW.project_id;
        
        -- عدد المهام المكتملة
        SELECT COUNT(*) INTO completed_tasks
        FROM tasks
        WHERE project_id = NEW.project_id AND status = 'completed';
        
        -- حساب النسبة المئوية
        IF total_tasks > 0 THEN
            SET progress = (completed_tasks / total_tasks) * 100;
        ELSE
            SET progress = 0;
        END IF;
        
        -- تحديث تقدم المشروع
        UPDATE projects
        SET progress_percentage = progress
        WHERE id = NEW.project_id;
    END IF;
END //

-- ====================================
-- 6. مشغل لتحديد رقم المشروع تلقائياً
-- Trigger to auto-generate project number
-- ====================================

DROP TRIGGER IF EXISTS before_project_insert //

CREATE TRIGGER before_project_insert
BEFORE INSERT ON projects
FOR EACH ROW
BEGIN
    IF NEW.project_number IS NULL OR NEW.project_number = '' THEN
        SET NEW.project_number = CONCAT(
            'PRJ-',
            YEAR(CURRENT_DATE),
            '-',
            LPAD((SELECT COALESCE(MAX(CAST(SUBSTRING(project_number, -4) AS UNSIGNED)), 0) + 1
                  FROM projects
                  WHERE project_number LIKE CONCAT('PRJ-', YEAR(CURRENT_DATE), '-%')), 4, '0')
        );
    END IF;
END //

-- ====================================
-- 7. مشغل لتحديد رقم المصروف تلقائياً
-- Trigger to auto-generate expense number
-- ====================================

DROP TRIGGER IF EXISTS before_expense_insert //

CREATE TRIGGER before_expense_insert
BEFORE INSERT ON expenses
FOR EACH ROW
BEGIN
    IF NEW.expense_number IS NULL OR NEW.expense_number = '' THEN
        SET NEW.expense_number = CONCAT(
            'EXP-',
            YEAR(CURRENT_DATE),
            '-',
            LPAD((SELECT COALESCE(MAX(CAST(SUBSTRING(expense_number, -4) AS UNSIGNED)), 0) + 1
                  FROM expenses
                  WHERE expense_number LIKE CONCAT('EXP-', YEAR(CURRENT_DATE), '-%')), 4, '0')
        );
    END IF;
END //

-- ====================================
-- 8. مشغل لتحديد رقم الإيراد تلقائياً
-- Trigger to auto-generate income number
-- ====================================

DROP TRIGGER IF EXISTS before_income_insert //

CREATE TRIGGER before_income_insert
BEFORE INSERT ON income
FOR EACH ROW
BEGIN
    IF NEW.income_number IS NULL OR NEW.income_number = '' THEN
        SET NEW.income_number = CONCAT(
            'INC-',
            YEAR(CURRENT_DATE),
            '-',
            LPAD((SELECT COALESCE(MAX(CAST(SUBSTRING(income_number, -4) AS UNSIGNED)), 0) + 1
                  FROM income
                  WHERE income_number LIKE CONCAT('INC-', YEAR(CURRENT_DATE), '-%')), 4, '0')
        );
    END IF;
END //

-- ====================================
-- 9. مشغل لتحديد رقم المهمة تلقائياً
-- Trigger to auto-generate task number
-- ====================================

DROP TRIGGER IF EXISTS before_task_insert //

CREATE TRIGGER before_task_insert
BEFORE INSERT ON tasks
FOR EACH ROW
BEGIN
    IF NEW.task_number IS NULL OR NEW.task_number = '' THEN
        SET NEW.task_number = CONCAT(
            'TSK-',
            YEAR(CURRENT_DATE),
            '-',
            LPAD((SELECT COALESCE(MAX(CAST(SUBSTRING(task_number, -4) AS UNSIGNED)), 0) + 1
                  FROM tasks
                  WHERE task_number LIKE CONCAT('TSK-', YEAR(CURRENT_DATE), '-%')), 4, '0')
        );
    END IF;
END //

-- ====================================
-- 10. مشغل لتحديد رقم يوم العمل تلقائياً
-- Trigger to auto-generate work day number
-- ====================================

DROP TRIGGER IF EXISTS before_work_day_insert //

CREATE TRIGGER before_work_day_insert
BEFORE INSERT ON work_days
FOR EACH ROW
BEGIN
    IF NEW.work_day_number IS NULL OR NEW.work_day_number = '' THEN
        SET NEW.work_day_number = CONCAT(
            'WD-',
            YEAR(CURRENT_DATE),
            '-',
            LPAD((SELECT COALESCE(MAX(CAST(SUBSTRING(work_day_number, -4) AS UNSIGNED)), 0) + 1
                  FROM work_days
                  WHERE work_day_number LIKE CONCAT('WD-', YEAR(CURRENT_DATE), '-%')), 4, '0')
        );
    END IF;
END //

-- ====================================
-- 11. مشغل لتحديد رقم المدين تلقائياً
-- Trigger to auto-generate debtor number
-- ====================================

DROP TRIGGER IF EXISTS before_debtor_insert //

CREATE TRIGGER before_debtor_insert
BEFORE INSERT ON debtors
FOR EACH ROW
BEGIN
    IF NEW.debtor_number IS NULL OR NEW.debtor_number = '' THEN
        SET NEW.debtor_number = CONCAT(
            'DBT-',
            YEAR(CURRENT_DATE),
            '-',
            LPAD((SELECT COALESCE(MAX(CAST(SUBSTRING(debtor_number, -4) AS UNSIGNED)), 0) + 1
                  FROM debtors
                  WHERE debtor_number LIKE CONCAT('DBT-', YEAR(CURRENT_DATE), '-%')), 4, '0')
        );
    END IF;
END //

-- ====================================
-- 12. مشغل لحساب التكلفة الإجمالية للعمالة
-- Trigger to calculate total labor cost
-- ====================================

DROP TRIGGER IF EXISTS before_work_day_labor_insert //

CREATE TRIGGER before_work_day_labor_insert
BEFORE INSERT ON work_day_labor
FOR EACH ROW
BEGIN
    IF NEW.total_cost IS NULL OR NEW.total_cost = 0 THEN
        IF NEW.hourly_rate IS NOT NULL AND NEW.hours_worked IS NOT NULL THEN
            SET NEW.total_cost = NEW.hourly_rate * NEW.hours_worked;
        ELSEIF NEW.daily_rate IS NOT NULL THEN
            SET NEW.total_cost = NEW.daily_rate;
        END IF;
    END IF;
END //

DROP TRIGGER IF EXISTS before_work_day_labor_update //

CREATE TRIGGER before_work_day_labor_update
BEFORE UPDATE ON work_day_labor
FOR EACH ROW
BEGIN
    IF NEW.hourly_rate IS NOT NULL AND NEW.hours_worked IS NOT NULL THEN
        SET NEW.total_cost = NEW.hourly_rate * NEW.hours_worked;
    ELSEIF NEW.daily_rate IS NOT NULL THEN
        SET NEW.total_cost = NEW.daily_rate;
    END IF;
END //

-- ====================================
-- 13. مشغل لحساب التكلفة الإجمالية للمعدات
-- Trigger to calculate total equipment cost
-- ====================================

DROP TRIGGER IF EXISTS before_work_day_equipment_insert //

CREATE TRIGGER before_work_day_equipment_insert
BEFORE INSERT ON work_day_equipment
FOR EACH ROW
BEGIN
    IF NEW.total_cost IS NULL OR NEW.total_cost = 0 THEN
        IF NEW.unit_price IS NOT NULL AND NEW.quantity IS NOT NULL THEN
            SET NEW.total_cost = NEW.unit_price * NEW.quantity;
        END IF;
    END IF;
END //

DROP TRIGGER IF EXISTS before_work_day_equipment_update //

CREATE TRIGGER before_work_day_equipment_update
BEFORE UPDATE ON work_day_equipment
FOR EACH ROW
BEGIN
    IF NEW.unit_price IS NOT NULL AND NEW.quantity IS NOT NULL THEN
        SET NEW.total_cost = NEW.unit_price * NEW.quantity;
    END IF;
END //

-- ====================================
-- 14. مشغل لحساب التكلفة الإجمالية للمواد
-- Trigger to calculate total materials cost
-- ====================================

DROP TRIGGER IF EXISTS before_work_day_materials_insert //

CREATE TRIGGER before_work_day_materials_insert
BEFORE INSERT ON work_day_materials
FOR EACH ROW
BEGIN
    IF NEW.total_cost IS NULL OR NEW.total_cost = 0 THEN
        IF NEW.unit_price IS NOT NULL AND NEW.quantity IS NOT NULL THEN
            SET NEW.total_cost = NEW.unit_price * NEW.quantity;
        END IF;
    END IF;
END //

DROP TRIGGER IF EXISTS before_work_day_materials_update //

CREATE TRIGGER before_work_day_materials_update
BEFORE UPDATE ON work_day_materials
FOR EACH ROW
BEGIN
    IF NEW.unit_price IS NOT NULL AND NEW.quantity IS NOT NULL THEN
        SET NEW.total_cost = NEW.unit_price * NEW.quantity;
    END IF;
END //

DELIMITER ;

-- ====================================
-- ملاحظات
-- Notes
-- ====================================

/*
المشغلات المنشأة:
1. تحديث تكاليف يوم العمل تلقائياً عند إضافة/تعديل/حذف العمالة
2. تحديث تكاليف يوم العمل تلقائياً عند إضافة/تعديل/حذف المعدات
3. تحديث تكاليف يوم العمل تلقائياً عند إضافة/تعديل/حذف المواد
4. تحديث حالة المدين تلقائياً عند إضافة دفعة
5. تحديث تقدم المشروع تلقائياً عند تحديث المهام
6. إنشاء أرقام تلقائية للمشاريع، المصاريف، الإيرادات، المهام، أيام العمل، والمدينين
7. حساب التكلفة الإجمالية تلقائياً للعمالة، المعدات، والمواد
*/

