-- ====================================
-- نظام إدارة المشاريع - البيانات الأولية
-- Project Management System - Seed Data
-- ====================================

-- ====================================
-- 1. الصلاحيات الأساسية
-- Basic Permissions
-- ====================================

INSERT INTO permissions (name, description, module) VALUES
-- صلاحيات المستخدمين
('users.view', 'عرض المستخدمين', 'users'),
('users.create', 'إضافة مستخدمين', 'users'),
('users.edit', 'تعديل المستخدمين', 'users'),
('users.delete', 'حذف المستخدمين', 'users'),
('users.manage_roles', 'إدارة أدوار المستخدمين', 'users'),

-- صلاحيات المشاريع
('projects.view', 'عرض المشاريع', 'projects'),
('projects.create', 'إضافة مشاريع', 'projects'),
('projects.edit', 'تعديل المشاريع', 'projects'),
('projects.delete', 'حذف المشاريع', 'projects'),
('projects.manage_team', 'إدارة فريق المشروع', 'projects'),

-- صلاحيات المهام
('tasks.view', 'عرض المهام', 'tasks'),
('tasks.create', 'إضافة مهام', 'tasks'),
('tasks.edit', 'تعديل المهام', 'tasks'),
('tasks.delete', 'حذف المهام', 'tasks'),
('tasks.assign', 'تعيين المهام', 'tasks'),

-- صلاحيات المصاريف
('expenses.view', 'عرض المصاريف', 'expenses'),
('expenses.create', 'إضافة مصاريف', 'expenses'),
('expenses.edit', 'تعديل المصاريف', 'expenses'),
('expenses.delete', 'حذف المصاريف', 'expenses'),
('expenses.approve', 'اعتماد المصاريف', 'expenses'),

-- صلاحيات الإيرادات
('income.view', 'عرض الإيرادات', 'income'),
('income.create', 'إضافة إيرادات', 'income'),
('income.edit', 'تعديل الإيرادات', 'income'),
('income.delete', 'حذف الإيرادات', 'income'),

-- صلاحيات المدينين
('debtors.view', 'عرض المدينين', 'debtors'),
('debtors.create', 'إضافة مدينين', 'debtors'),
('debtors.edit', 'تعديل المدينين', 'debtors'),
('debtors.delete', 'حذف المدينين', 'debtors'),
('debtors.manage_payments', 'إدارة دفعات المدينين', 'debtors'),

-- صلاحيات أيام العمل
('workdays.view', 'عرض أيام العمل', 'workdays'),
('workdays.create', 'إضافة أيام عمل', 'workdays'),
('workdays.edit', 'تعديل أيام العمل', 'workdays'),
('workdays.delete', 'حذف أيام العمل', 'workdays'),

-- صلاحيات الفريق
('team.view', 'عرض أعضاء الفريق', 'team'),
('team.create', 'إضافة أعضاء فريق', 'team'),
('team.edit', 'تعديل أعضاء الفريق', 'team'),
('team.delete', 'حذف أعضاء الفريق', 'team'),

-- صلاحيات الأقسام
('departments.view', 'عرض الأقسام', 'departments'),
('departments.create', 'إضافة أقسام', 'departments'),
('departments.edit', 'تعديل الأقسام', 'departments'),
('departments.delete', 'حذف الأقسام', 'departments'),

-- صلاحيات التقارير
('reports.financial', 'عرض التقارير المالية', 'reports'),
('reports.projects', 'عرض تقارير المشاريع', 'reports'),
('reports.team', 'عرض تقارير الفريق', 'reports'),
('reports.export', 'تصدير التقارير', 'reports'),

-- صلاحيات الإعدادات
('settings.view', 'عرض الإعدادات', 'settings'),
('settings.edit', 'تعديل الإعدادات', 'settings');

-- ====================================
-- 2. الأدوار الأساسية
-- Basic Roles
-- ====================================

INSERT INTO roles (name, display_name, description, is_system_role) VALUES
('super_admin', 'مدير النظام', 'صلاحيات كاملة على جميع أجزاء النظام', TRUE),
('admin', 'مدير', 'صلاحيات إدارية واسعة', TRUE),
('project_manager', 'مدير مشروع', 'إدارة المشاريع والفرق', FALSE),
('engineer', 'مهندس', 'تنفيذ المهام والمشاريع', FALSE),
('accountant', 'محاسب', 'إدارة المالية والمصاريف', FALSE),
('viewer', 'مشاهد', 'عرض البيانات فقط', FALSE);

-- ====================================
-- 3. ربط الصلاحيات بالأدوار
-- Role Permissions Mapping
-- ====================================

-- مدير النظام - جميع الصلاحيات
INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id FROM permissions;

-- مدير - معظم الصلاحيات
INSERT INTO role_permissions (role_id, permission_id)
SELECT 2, id FROM permissions WHERE name NOT LIKE 'settings.edit';

-- مدير المشروع
INSERT INTO role_permissions (role_id, permission_id)
SELECT 3, id FROM permissions WHERE 
    module IN ('projects', 'tasks', 'team', 'workdays') OR
    name IN ('expenses.view', 'income.view', 'reports.projects', 'reports.team');

-- المهندس
INSERT INTO role_permissions (role_id, permission_id)
SELECT 4, id FROM permissions WHERE 
    name IN ('projects.view', 'tasks.view', 'tasks.edit', 'workdays.view', 'team.view');

-- المحاسب
INSERT INTO role_permissions (role_id, permission_id)
SELECT 5, id FROM permissions WHERE 
    module IN ('expenses', 'income', 'debtors') OR
    name IN ('reports.financial', 'projects.view');

-- المشاهد
INSERT INTO role_permissions (role_id, permission_id)
SELECT 6, id FROM permissions WHERE name LIKE '%.view';

-- ====================================
-- 4. المستخدمين الافتراضيين
-- Default Users
-- ====================================

-- كلمة المرور للجميع: password123 (يجب تشفيرها في التطبيق الفعلي)
INSERT INTO users (username, email, password, full_name, phone, status, email_verified) VALUES
('admin', 'admin@erticaz.com', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'مدير النظام', '07701234567', 'active', TRUE),
('manager1', 'manager1@erticaz.com', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'أحمد محمد علي', '07701234568', 'active', TRUE),
('engineer1', 'engineer1@erticaz.com', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'محمد عبدالله', '07701234569', 'active', TRUE),
('accountant1', 'accountant@erticaz.com', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'فاطمة حسن', '07701234570', 'active', TRUE);

-- ربط المستخدمين بالأدوار
INSERT INTO user_roles (user_id, role_id) VALUES
(1, 1), -- admin -> super_admin
(2, 3), -- manager1 -> project_manager
(3, 4), -- engineer1 -> engineer
(4, 5); -- accountant1 -> accountant

-- ====================================
-- 5. الأقسام
-- Departments
-- ====================================

INSERT INTO departments (name, code, description, manager_id, budget, status) VALUES
('الهندسة المدنية', 'CIVIL', 'قسم الهندسة المدنية', 2, 500000.00, 'active'),
('الهندسة المعمارية', 'ARCH', 'قسم الهندسة المعمارية', 2, 400000.00, 'active'),
('الهندسة الكهربائية', 'ELEC', 'قسم الهندسة الكهربائية', NULL, 300000.00, 'active'),
('الهندسة الميكانيكية', 'MECH', 'قسم الهندسة الميكانيكية', NULL, 300000.00, 'active'),
('إدارة المشاريع', 'PM', 'قسم إدارة المشاريع', 2, 200000.00, 'active'),
('المحاسبة', 'ACC', 'قسم المحاسبة', 4, 150000.00, 'active'),
('الموارد البشرية', 'HR', 'قسم الموارد البشرية', NULL, 100000.00, 'active');

-- ====================================
-- 6. فئات المصاريف
-- Expense Categories
-- ====================================

INSERT INTO expense_categories (name, code, description, is_active) VALUES
('مصاريف مباشرة', 'DIRECT', 'مصاريف مرتبطة مباشرة بالمشاريع', TRUE),
('مصاريف غير مباشرة', 'INDIRECT', 'مصاريف إدارية وعامة', TRUE),
('رواتب وأجور', 'SALARY', 'رواتب الموظفين والعمالة', TRUE),
('مواد ومستلزمات', 'MATERIALS', 'مواد البناء والمستلزمات', TRUE),
('معدات', 'EQUIPMENT', 'معدات وأجهزة', TRUE),
('نقل ومواصلات', 'TRANSPORT', 'مصاريف النقل والمواصلات', TRUE),
('صيانة', 'MAINTENANCE', 'مصاريف الصيانة', TRUE),
('مرافق', 'UTILITIES', 'كهرباء وماء وغاز', TRUE);

-- ====================================
-- 7. أنواع المصاريف
-- Expense Types
-- ====================================

INSERT INTO expense_types (name, code, description, category_id, is_active) VALUES
('رواتب شهرية', 'SAL-MON', 'رواتب الموظفين الشهرية', 3, TRUE),
('أجور يومية', 'WAG-DAY', 'أجور العمالة اليومية', 3, TRUE),
('أسمنت', 'MAT-CEM', 'مادة الأسمنت', 4, TRUE),
('حديد', 'MAT-STL', 'حديد التسليح', 4, TRUE),
('طابوق', 'MAT-BRK', 'الطابوق', 4, TRUE),
('رمل', 'MAT-SND', 'مادة الرمل', 4, TRUE),
('حصى', 'MAT-GRV', 'مادة الحصى', 4, TRUE),
('كرين', 'EQP-CRN', 'تأجير كرين', 5, TRUE),
('خلاطة', 'EQP-MIX', 'تأجير خلاطة', 5, TRUE),
('شاحنات', 'TRN-TRK', 'أجور شاحنات النقل', 6, TRUE),
('وقود', 'TRN-FUL', 'وقود المركبات', 6, TRUE),
('كهرباء', 'UTL-ELE', 'فواتير الكهرباء', 8, TRUE),
('ماء', 'UTL-WAT', 'فواتير الماء', 8, TRUE);

-- ====================================
-- 8. فئات الإيرادات
-- Income Categories
-- ====================================

INSERT INTO income_categories (name, code, description, is_active) VALUES
('إيرادات مشاريع', 'PROJ', 'إيرادات من المشاريع', TRUE),
('دفعات عملاء', 'CLIENT', 'دفعات من العملاء', TRUE),
('دفعات مقدمة', 'ADVANCE', 'دفعات مقدمة للمشاريع', TRUE),
('إيرادات أخرى', 'OTHER', 'إيرادات متنوعة', TRUE);

-- ====================================
-- 9. أعضاء الفريق
-- Team Members
-- ====================================

INSERT INTO team_members (user_id, name, email, phone, department_id, position, role, specialization, hire_date, salary, hourly_rate, status) VALUES
(2, 'أحمد محمد علي', 'ahmed.mohamed@erticaz.com', '07701234567', 1, 'مدير مشروع', 'مهندس أول', 'الهندسة المدنية', '2020-01-15', 2500000.00, 15000.00, 'active'),
(3, 'محمد عبدالله', 'mohamed.abdullah@erticaz.com', '07701234568', 1, 'مهندس', 'مهندس', 'الهندسة المدنية', '2021-03-20', 1800000.00, 12000.00, 'active'),
(4, 'فاطمة حسن', 'fatima.hassan@erticaz.com', '07701234569', 6, 'محاسبة', 'محاسب أول', 'محاسبة', '2020-06-10', 1500000.00, 10000.00, 'active'),
(NULL, 'سارة أحمد', 'sara.ahmed@erticaz.com', '07701234570', 2, 'مهندسة معمارية', 'مهندسة', 'الهندسة المعمارية', '2021-08-15', 1700000.00, 11000.00, 'active'),
(NULL, 'علي محمود', 'ali.mahmoud@erticaz.com', '07701234571', 3, 'مهندس كهربائي', 'مهندس', 'الهندسة الكهربائية', '2022-01-10', 1600000.00, 11000.00, 'active'),
(NULL, 'حسن عبدالرحمن', 'hassan.abdelrahman@erticaz.com', '07701234572', 4, 'مهندس ميكانيكي', 'مهندس', 'الهندسة الميكانيكية', '2022-03-15', 1600000.00, 11000.00, 'active'),
(NULL, 'زينب علي', 'zainab.ali@erticaz.com', '07701234573', 5, 'منسقة مشاريع', 'مساعد', 'إدارة مشاريع', '2022-05-20', 1200000.00, 8000.00, 'active'),
(NULL, 'يوسف خالد', 'yousif.khaled@erticaz.com', '07701234574', 1, 'مشرف موقع', 'مشرف', 'الهندسة المدنية', '2021-11-01', 1400000.00, 9000.00, 'active');

-- ====================================
-- 10. المشاريع
-- Projects
-- ====================================

INSERT INTO projects (project_number, name, description, client_name, client_phone, work_location, start_date, end_date, contract_value, estimated_cost, status, priority, progress_percentage, manager_id, department_id, created_by) VALUES
('PRJ-2024-001', 'مشروع بناء مجمع سكني', 'بناء مجمع سكني متكامل يحتوي على 50 وحدة سكنية في منطقة الكرادة', 'شركة البناء المتقدم', '07701234580', 'الكرادة - بغداد', '2024-01-15', '2024-12-31', 5000000000.00, 4500000000.00, 'in_progress', 'high', 35.50, 1, 1, 1),
('PRJ-2024-002', 'مشروع مدرسة ابتدائية', 'إنشاء مدرسة ابتدائية حديثة في منطقة الزعفرانية', 'وزارة التربية', '07701234581', 'الزعفرانية - بغداد', '2024-02-01', '2024-11-30', 3000000000.00, 2800000000.00, 'in_progress', 'high', 45.00, 1, 1, 1),
('PRJ-2024-003', 'مشروع مركز تجاري', 'بناء مركز تجاري متعدد الطوابق في المنصور', 'مجموعة المنصور التجارية', '07701234582', 'المنصور - بغداد', '2024-03-10', '2025-03-10', 8000000000.00, 7200000000.00, 'pending', 'medium', 0.00, 1, 1, 1),
('PRJ-2023-015', 'مشروع جسر الدورة', 'إنشاء جسر حديث في منطقة الدورة', 'أمانة بغداد', '07701234583', 'الدورة - بغداد', '2023-06-01', '2024-06-01', 4500000000.00, 4200000000.00, 'completed', 'high', 100.00, 1, 1, 1);

-- ====================================
-- 11. فريق المشروع
-- Project Team
-- ====================================

INSERT INTO project_team (project_id, team_member_id, role, assigned_date, is_active) VALUES
(1, 1, 'مدير المشروع', '2024-01-15', TRUE),
(1, 2, 'مهندس تنفيذ', '2024-01-15', TRUE),
(1, 4, 'مهندس معماري', '2024-01-15', TRUE),
(1, 8, 'مشرف موقع', '2024-01-20', TRUE),
(2, 1, 'مدير المشروع', '2024-02-01', TRUE),
(2, 2, 'مهندس تنفيذ', '2024-02-01', TRUE),
(2, 5, 'مهندس كهرباء', '2024-02-05', TRUE),
(3, 1, 'مدير المشروع', '2024-03-10', TRUE),
(4, 1, 'مدير المشروع', '2023-06-01', FALSE),
(4, 2, 'مهندس تنفيذ', '2023-06-01', FALSE);

-- ====================================
-- 12. المهام
-- Tasks
-- ====================================

INSERT INTO tasks (task_number, title, description, project_id, assigned_to, assigned_by, status, priority, start_date, due_date, estimated_hours, progress_percentage, created_by) VALUES
('TSK-2024-001', 'أعمال الحفر والتسوية', 'حفر وتسوية الأرض للمجمع السكني', 1, 2, 1, 'completed', 'high', '2024-01-15', '2024-02-15', 320, 100.00, 1),
('TSK-2024-002', 'صب القواعد', 'صب قواعد المبنى', 1, 2, 1, 'completed', 'high', '2024-02-16', '2024-03-15', 280, 100.00, 1),
('TSK-2024-003', 'بناء الجدران الأرضية', 'بناء جدران الطابق الأرضي', 1, 2, 1, 'in_progress', 'high', '2024-03-16', '2024-05-15', 400, 65.00, 1),
('TSK-2024-004', 'تركيب الأبواب والنوافذ', 'تركيب جميع الأبواب والنوافذ', 1, 4, 1, 'pending', 'medium', '2024-05-16', '2024-06-30', 200, 0.00, 1),
('TSK-2024-005', 'التشطيبات الخارجية', 'أعمال التشطيبات الخارجية', 1, 4, 1, 'pending', 'medium', '2024-07-01', '2024-08-31', 300, 0.00, 1),
('TSK-2024-010', 'أعمال الحفر', 'حفر وتسوية أرض المدرسة', 2, 2, 1, 'completed', 'high', '2024-02-01', '2024-02-20', 200, 100.00, 1),
('TSK-2024-011', 'صب الأساسات', 'صب أساسات المبنى', 2, 2, 1, 'in_progress', 'high', '2024-02-21', '2024-03-20', 250, 80.00, 1),
('TSK-2024-012', 'بناء الجدران', 'بناء جدران المدرسة', 2, 2, 1, 'pending', 'high', '2024-03-21', '2024-05-30', 350, 0.00, 1);

-- ====================================
-- 13. المصاريف
-- Expenses
-- ====================================

INSERT INTO expenses (expense_number, description, amount, expense_date, expense_type_id, category_id, project_id, department_id, paid_to, payment_method, status, created_by) VALUES
('EXP-2024-001', 'رواتب شهر يناير للمشروع', 15000000.00, '2024-01-31', 1, 3, 1, 1, 'الموظفين', 'bank_transfer', 'paid', 1),
('EXP-2024-002', 'شراء أسمنت للمشروع', 25000000.00, '2024-02-05', 3, 4, 1, 1, 'شركة الأسمنت العراقية', 'bank_transfer', 'paid', 1),
('EXP-2024-003', 'شراء حديد تسليح', 40000000.00, '2024-02-10', 4, 4, 1, 1, 'شركة الحديد والصلب', 'bank_transfer', 'paid', 1),
('EXP-2024-004', 'أجور عمالة يومية', 5000000.00, '2024-02-15', 2, 3, 1, 1, 'العمال', 'cash', 'paid', 1),
('EXP-2024-005', 'تأجير كرين', 8000000.00, '2024-02-20', 8, 5, 1, 1, 'شركة المعدات الثقيلة', 'check', 'paid', 1),
('EXP-2024-010', 'رواتب شهر فبراير للمشروع الثاني', 12000000.00, '2024-02-28', 1, 3, 2, 1, 'الموظفين', 'bank_transfer', 'paid', 1),
('EXP-2024-011', 'شراء مواد بناء', 18000000.00, '2024-03-05', 5, 4, 2, 1, 'شركة مواد البناء', 'bank_transfer', 'approved', 1),
('EXP-2024-012', 'أجور نقل مواد', 3000000.00, '2024-03-10', 10, 6, 2, 1, 'شركة النقل السريع', 'cash', 'pending', 1);

-- ====================================
-- 14. الإيرادات
-- Income
-- ====================================

INSERT INTO income (income_number, description, amount, income_date, category_id, project_id, client_name, payment_method, status, created_by) VALUES
('INC-2024-001', 'دفعة مقدمة - مشروع المجمع السكني', 1000000000.00, '2024-01-10', 3, 1, 'شركة البناء المتقدم', 'bank_transfer', 'confirmed', 1),
('INC-2024-002', 'دفعة أولى - مشروع المجمع السكني', 1500000000.00, '2024-03-15', 1, 1, 'شركة البناء المتقدم', 'bank_transfer', 'confirmed', 1),
('INC-2024-003', 'دفعة مقدمة - مشروع المدرسة', 600000000.00, '2024-01-25', 3, 2, 'وزارة التربية', 'bank_transfer', 'confirmed', 1),
('INC-2024-004', 'دفعة أولى - مشروع المدرسة', 900000000.00, '2024-03-20', 1, 2, 'وزارة التربية', 'bank_transfer', 'confirmed', 1),
('INC-2023-050', 'دفعة نهائية - مشروع جسر الدورة', 500000000.00, '2024-06-05', 1, 4, 'أمانة بغداد', 'bank_transfer', 'confirmed', 1);

-- ====================================
-- 15. المدينين
-- Debtors
-- ====================================

INSERT INTO debtors (debtor_number, name, email, phone, company_name, total_debt, paid_amount, remaining_amount, project_id, due_date, status, created_by) VALUES
('DBT-2024-001', 'شركة البناء المتقدم', 'info@advancedbuild.iq', '07701234590', 'شركة البناء المتقدم', 2500000000.00, 2500000000.00, 0.00, 1, '2024-12-31', 'partially_paid', 1),
('DBT-2024-002', 'وزارة التربية', 'finance@moe.gov.iq', '07701234591', 'وزارة التربية', 1500000000.00, 1500000000.00, 0.00, 2, '2024-11-30', 'partially_paid', 1),
('DBT-2024-003', 'مجموعة المنصور التجارية', 'finance@mansour.iq', '07701234592', 'مجموعة المنصور التجارية', 8000000000.00, 0.00, 8000000000.00, 3, '2025-03-10', 'active', 1);

-- ====================================
-- 16. دفعات المدينين
-- Debtor Payments
-- ====================================

INSERT INTO debtor_payments (debtor_id, payment_number, amount, payment_date, payment_method, receipt_number, created_by) VALUES
(1, 'PAY-2024-001', 1000000000.00, '2024-01-10', 'bank_transfer', 'RCP-001', 1),
(1, 'PAY-2024-002', 1500000000.00, '2024-03-15', 'bank_transfer', 'RCP-002', 1),
(2, 'PAY-2024-003', 600000000.00, '2024-01-25', 'bank_transfer', 'RCP-003', 1),
(2, 'PAY-2024-004', 900000000.00, '2024-03-20', 'bank_transfer', 'RCP-004', 1);

-- ====================================
-- 17. أقساط المدينين
-- Payment Installments
-- ====================================

INSERT INTO payment_installments (debtor_id, installment_number, amount, due_date, status) VALUES
(1, 1, 1000000000.00, '2024-01-10', 'paid'),
(1, 2, 1500000000.00, '2024-03-15', 'paid'),
(1, 3, 1000000000.00, '2024-06-15', 'pending'),
(2, 1, 600000000.00, '2024-01-25', 'paid'),
(2, 2, 900000000.00, '2024-03-20', 'paid'),
(2, 3, 1500000000.00, '2024-07-20', 'pending'),
(3, 1, 2000000000.00, '2024-06-10', 'pending'),
(3, 2, 3000000000.00, '2024-09-10', 'pending'),
(3, 3, 3000000000.00, '2024-12-10', 'pending');

-- ====================================
-- 18. أيام العمل
-- Work Days
-- ====================================

INSERT INTO work_days (work_day_number, project_id, work_date, work_location, form_number, work_period_from, work_period_to, day_name, work_type, description, status, created_by) VALUES
('WD-2024-001', 1, '2024-02-15', 'موقع المجمع السكني - الكرادة', 'F-001', '07:00:00', '17:00:00', 'الخميس', 'أعمال حفر وتسوية', 'أعمال الحفر والتسوية للقواعد', 'completed', 1),
('WD-2024-002', 1, '2024-02-16', 'موقع المجمع السكني - الكرادة', 'F-002', '07:00:00', '17:00:00', 'الجمعة', 'أعمال خرسانة', 'صب خرسانة القواعد', 'completed', 1),
('WD-2024-003', 1, '2024-03-20', 'موقع المجمع السكني - الكرادة', 'F-003', '07:00:00', '17:00:00', 'الأربعاء', 'أعمال بناء', 'بناء الجدران الأرضية', 'completed', 1),
('WD-2024-010', 2, '2024-02-10', 'موقع المدرسة - الزعفرانية', 'F-010', '07:00:00', '16:00:00', 'السبت', 'أعمال حفر', 'حفر أساسات المدرسة', 'completed', 1),
('WD-2024-011', 2, '2024-02-25', 'موقع المدرسة - الزعفرانية', 'F-011', '07:00:00', '16:00:00', 'الأحد', 'أعمال خرسانة', 'صب الأساسات', 'in_progress', 1);

-- ====================================
-- 19. عمالة أيام العمل
-- Work Day Labor
-- ====================================

INSERT INTO work_day_labor (work_day_id, worker_name, job_title, hours_worked, hourly_rate, daily_rate, total_cost) VALUES
(1, 'عامل الحفر 1', 'عامل حفر', 10.00, 5000.00, NULL, 50000.00),
(1, 'عامل الحفر 2', 'عامل حفر', 10.00, 5000.00, NULL, 50000.00),
(1, 'مشغل حفار', 'مشغل آليات', 10.00, 8000.00, NULL, 80000.00),
(2, 'عامل خرسانة 1', 'عامل خرسانة', 10.00, 6000.00, NULL, 60000.00),
(2, 'عامل خرسانة 2', 'عامل خرسانة', 10.00, 6000.00, NULL, 60000.00),
(2, 'حداد', 'حداد', 10.00, 7000.00, NULL, 70000.00),
(3, 'بناء 1', 'بناء', 9.00, 7000.00, NULL, 63000.00),
(3, 'بناء 2', 'بناء', 9.00, 7000.00, NULL, 63000.00),
(3, 'مساعد بناء', 'مساعد', 9.00, 5000.00, NULL, 45000.00);

-- ====================================
-- 20. معدات أيام العمل
-- Work Day Equipment
-- ====================================

INSERT INTO work_day_equipment (work_day_id, equipment_name, equipment_type, quantity, unit_of_measure, unit_price, total_cost, rental_hours) VALUES
(1, 'حفار كاتربلر', 'معدات حفر', 1.00, 'يوم', 500000.00, 500000.00, 10.00),
(1, 'شاحنة قلاب', 'معدات نقل', 2.00, 'يوم', 150000.00, 300000.00, 10.00),
(2, 'خلاطة خرسانة', 'معدات خرسانة', 1.00, 'يوم', 200000.00, 200000.00, 8.00),
(2, 'مضخة خرسانة', 'معدات خرسانة', 1.00, 'يوم', 400000.00, 400000.00, 6.00),
(3, 'سقالات', 'معدات بناء', 50.00, 'متر', 5000.00, 250000.00, NULL),
(3, 'خلاطة مونة', 'معدات بناء', 1.00, 'يوم', 100000.00, 100000.00, 9.00);

-- ====================================
-- 21. مواد أيام العمل
-- Work Day Materials
-- ====================================

INSERT INTO work_day_materials (work_day_id, material_name, material_type, quantity, unit_of_measure, unit_price, total_cost, supplier_name) VALUES
(2, 'أسمنت', 'مواد بناء', 50.00, 'كيس', 12000.00, 600000.00, 'شركة الأسمنت العراقية'),
(2, 'حصى', 'مواد بناء', 10.00, 'متر مكعب', 40000.00, 400000.00, 'محجر الرشيد'),
(2, 'رمل', 'مواد بناء', 8.00, 'متر مكعب', 35000.00, 280000.00, 'محجر الرشيد'),
(2, 'حديد تسليح', 'مواد إنشائية', 2.00, 'طن', 1500000.00, 3000000.00, 'شركة الحديد والصلب'),
(3, 'طابوق', 'مواد بناء', 5000.00, 'حبة', 500.00, 2500000.00, 'معمل الطابوق الحديث'),
(3, 'أسمنت', 'مواد بناء', 30.00, 'كيس', 12000.00, 360000.00, 'شركة الأسمنت العراقية'),
(3, 'رمل', 'مواد بناء', 5.00, 'متر مكعب', 35000.00, 175000.00, 'محجر الرشيد');

-- ====================================
-- تحديث التكاليف الإجمالية لأيام العمل
-- Update Total Costs for Work Days
-- ====================================

UPDATE work_days wd SET
    total_labor_cost = (SELECT COALESCE(SUM(total_cost), 0) FROM work_day_labor WHERE work_day_id = wd.id),
    total_equipment_cost = (SELECT COALESCE(SUM(total_cost), 0) FROM work_day_equipment WHERE work_day_id = wd.id),
    total_materials_cost = (SELECT COALESCE(SUM(total_cost), 0) FROM work_day_materials WHERE work_day_id = wd.id);

UPDATE work_days SET total_cost = total_labor_cost + total_equipment_cost + total_materials_cost;

-- ====================================
-- نهاية ملف البيانات الأولية
-- End of Seed Data File
-- ====================================

