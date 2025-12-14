-- ====================================
-- نظام إدارة المشاريع - قاعدة البيانات الكاملة
-- Project Management System - Complete Database Schema
-- ====================================

-- حذف الجداول إذا كانت موجودة (للتطوير فقط)
DROP TABLE IF EXISTS payment_installments;
DROP TABLE IF EXISTS debtor_payments;
DROP TABLE IF EXISTS debtors;
DROP TABLE IF EXISTS work_day_materials;
DROP TABLE IF EXISTS work_day_equipment;
DROP TABLE IF EXISTS work_day_labor;
DROP TABLE IF EXISTS work_days;
DROP TABLE IF EXISTS task_comments;
DROP TABLE IF EXISTS task_attachments;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS project_team;
DROP TABLE IF EXISTS project_expenses;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS expenses;
DROP TABLE IF EXISTS income;
DROP TABLE IF EXISTS team_members;
DROP TABLE IF EXISTS departments;
DROP TABLE IF EXISTS expense_types;
DROP TABLE IF EXISTS income_categories;
DROP TABLE IF EXISTS expense_categories;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS role_permissions;

-- ====================================
-- 1. نظام المستخدمين والصلاحيات
-- Users and Permissions System
-- ====================================

-- جدول الصلاحيات
CREATE TABLE permissions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    module VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول الأدوار
CREATE TABLE roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_system_role BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول ربط الأدوار بالصلاحيات
CREATE TABLE role_permissions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role_id INT NOT NULL,
    permission_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,
    UNIQUE KEY unique_role_permission (role_id, permission_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول المستخدمين
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    avatar VARCHAR(500),
    status ENUM('active', 'inactive', 'suspended', 'pending') DEFAULT 'active',
    last_login TIMESTAMP NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    email_verified_at TIMESTAMP NULL,
    remember_token VARCHAR(100),
    reset_token VARCHAR(100),
    reset_token_expires TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_username (username),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول ربط المستخدمين بالأدوار
CREATE TABLE user_roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_role (user_id, role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 2. الأقسام والفئات
-- Departments and Categories
-- ====================================

-- جدول الأقسام
CREATE TABLE departments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE,
    description TEXT,
    manager_id INT,
    parent_department_id INT,
    budget DECIMAL(15, 2) DEFAULT 0,
    status ENUM('active', 'inactive') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (manager_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (parent_department_id) REFERENCES departments(id) ON DELETE SET NULL,
    INDEX idx_status (status),
    INDEX idx_code (code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول فئات المصاريف
CREATE TABLE expense_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE,
    description TEXT,
    parent_category_id INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_category_id) REFERENCES expense_categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول فئات الإيرادات
CREATE TABLE income_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE,
    description TEXT,
    parent_category_id INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_category_id) REFERENCES income_categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول أنواع المصاريف
CREATE TABLE expense_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE,
    description TEXT,
    category_id INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES expense_categories(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 3. أعضاء الفريق والمهندسين
-- Team Members and Engineers
-- ====================================

-- جدول أعضاء الفريق
CREATE TABLE team_members (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(20),
    department_id INT,
    position VARCHAR(255),
    role VARCHAR(100),
    specialization VARCHAR(255),
    hire_date DATE,
    salary DECIMAL(12, 2),
    hourly_rate DECIMAL(10, 2),
    status ENUM('active', 'inactive', 'on_leave', 'resigned') DEFAULT 'active',
    address TEXT,
    national_id VARCHAR(50),
    emergency_contact VARCHAR(255),
    emergency_phone VARCHAR(20),
    notes TEXT,
    avatar VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE SET NULL,
    INDEX idx_email (email),
    INDEX idx_status (status),
    INDEX idx_department (department_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 4. المشاريع
-- Projects
-- ====================================

-- جدول المشاريع
CREATE TABLE projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    project_number VARCHAR(100) UNIQUE,
    name VARCHAR(500) NOT NULL,
    description TEXT,
    client_name VARCHAR(255),
    client_phone VARCHAR(20),
    client_email VARCHAR(255),
    client_address TEXT,
    work_location TEXT,
    start_date DATE,
    end_date DATE,
    planned_end_date DATE,
    contract_value DECIMAL(15, 2) DEFAULT 0,
    actual_cost DECIMAL(15, 2) DEFAULT 0,
    estimated_cost DECIMAL(15, 2) DEFAULT 0,
    status ENUM('pending', 'in_progress', 'on_hold', 'completed', 'cancelled') DEFAULT 'pending',
    priority ENUM('low', 'medium', 'high', 'urgent') DEFAULT 'medium',
    progress_percentage DECIMAL(5, 2) DEFAULT 0,
    manager_id INT,
    department_id INT,
    notes TEXT,
    attachments JSON,
    metadata JSON,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (manager_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_project_number (project_number),
    INDEX idx_status (status),
    INDEX idx_start_date (start_date),
    INDEX idx_manager (manager_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول فريق المشروع
CREATE TABLE project_team (
    id INT AUTO_INCREMENT PRIMARY KEY,
    project_id INT NOT NULL,
    team_member_id INT NOT NULL,
    role VARCHAR(100),
    assigned_date DATE,
    removed_date DATE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (team_member_id) REFERENCES team_members(id) ON DELETE CASCADE,
    INDEX idx_project (project_id),
    INDEX idx_team_member (team_member_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 5. المهام
-- Tasks
-- ====================================

-- جدول المهام
CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task_number VARCHAR(100) UNIQUE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    project_id INT,
    parent_task_id INT,
    assigned_to INT,
    assigned_by INT,
    status ENUM('pending', 'in_progress', 'review', 'completed', 'cancelled', 'on_hold') DEFAULT 'pending',
    priority ENUM('low', 'medium', 'high', 'urgent') DEFAULT 'medium',
    start_date DATE,
    due_date DATE,
    completed_date DATE,
    estimated_hours DECIMAL(10, 2),
    actual_hours DECIMAL(10, 2),
    progress_percentage DECIMAL(5, 2) DEFAULT 0,
    tags JSON,
    attachments JSON,
    metadata JSON,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_task_id) REFERENCES tasks(id) ON DELETE CASCADE,
    FOREIGN KEY (assigned_to) REFERENCES team_members(id) ON DELETE SET NULL,
    FOREIGN KEY (assigned_by) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_task_number (task_number),
    INDEX idx_status (status),
    INDEX idx_priority (priority),
    INDEX idx_assigned_to (assigned_to),
    INDEX idx_project (project_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول تعليقات المهام
CREATE TABLE task_comments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task_id INT NOT NULL,
    user_id INT NOT NULL,
    comment TEXT NOT NULL,
    parent_comment_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_comment_id) REFERENCES task_comments(id) ON DELETE CASCADE,
    INDEX idx_task (task_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول مرفقات المهام
CREATE TABLE task_attachments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task_id INT NOT NULL,
    file_name VARCHAR(500) NOT NULL,
    file_path VARCHAR(1000) NOT NULL,
    file_type VARCHAR(100),
    file_size BIGINT,
    uploaded_by INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
    FOREIGN KEY (uploaded_by) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_task (task_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 6. المصاريف والإيرادات
-- Expenses and Income
-- ====================================

-- جدول المصاريف
CREATE TABLE expenses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    expense_number VARCHAR(100) UNIQUE,
    description TEXT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    expense_date DATE NOT NULL,
    expense_type_id INT,
    category_id INT,
    project_id INT,
    department_id INT,
    paid_to VARCHAR(255),
    payment_method ENUM('cash', 'bank_transfer', 'check', 'credit_card', 'other') DEFAULT 'cash',
    receipt_number VARCHAR(100),
    status ENUM('pending', 'approved', 'rejected', 'paid') DEFAULT 'pending',
    approved_by INT,
    approved_at TIMESTAMP NULL,
    notes TEXT,
    attachments JSON,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (expense_type_id) REFERENCES expense_types(id) ON DELETE SET NULL,
    FOREIGN KEY (category_id) REFERENCES expense_categories(id) ON DELETE SET NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE SET NULL,
    FOREIGN KEY (approved_by) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_expense_number (expense_number),
    INDEX idx_expense_date (expense_date),
    INDEX idx_status (status),
    INDEX idx_project (project_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول مصاريف المشاريع
CREATE TABLE project_expenses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    project_id INT NOT NULL,
    expense_id INT NOT NULL,
    allocation_percentage DECIMAL(5, 2) DEFAULT 100,
    allocated_amount DECIMAL(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (expense_id) REFERENCES expenses(id) ON DELETE CASCADE,
    INDEX idx_project (project_id),
    INDEX idx_expense (expense_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول الإيرادات
CREATE TABLE income (
    id INT AUTO_INCREMENT PRIMARY KEY,
    income_number VARCHAR(100) UNIQUE,
    description TEXT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    income_date DATE NOT NULL,
    category_id INT,
    project_id INT,
    client_name VARCHAR(255),
    payment_method ENUM('cash', 'bank_transfer', 'check', 'credit_card', 'other') DEFAULT 'cash',
    receipt_number VARCHAR(100),
    invoice_number VARCHAR(100),
    status ENUM('pending', 'confirmed', 'cancelled') DEFAULT 'confirmed',
    notes TEXT,
    attachments JSON,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES income_categories(id) ON DELETE SET NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_income_number (income_number),
    INDEX idx_income_date (income_date),
    INDEX idx_status (status),
    INDEX idx_project (project_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 7. المدينين
-- Debtors
-- ====================================

-- جدول المدينين
CREATE TABLE debtors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    debtor_number VARCHAR(100) UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(20),
    address TEXT,
    company_name VARCHAR(255),
    total_debt DECIMAL(15, 2) NOT NULL DEFAULT 0,
    paid_amount DECIMAL(15, 2) DEFAULT 0,
    remaining_amount DECIMAL(15, 2) DEFAULT 0,
    project_id INT,
    due_date DATE,
    status ENUM('active', 'partially_paid', 'fully_paid', 'overdue', 'cancelled') DEFAULT 'active',
    notes TEXT,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_debtor_number (debtor_number),
    INDEX idx_status (status),
    INDEX idx_due_date (due_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول دفعات المدينين
CREATE TABLE debtor_payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    debtor_id INT NOT NULL,
    payment_number VARCHAR(100) UNIQUE,
    amount DECIMAL(15, 2) NOT NULL,
    payment_date DATE NOT NULL,
    payment_method ENUM('cash', 'bank_transfer', 'check', 'credit_card', 'other') DEFAULT 'cash',
    receipt_number VARCHAR(100),
    notes TEXT,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (debtor_id) REFERENCES debtors(id) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_debtor (debtor_id),
    INDEX idx_payment_date (payment_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول أقساط المدينين
CREATE TABLE payment_installments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    debtor_id INT NOT NULL,
    installment_number INT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    due_date DATE NOT NULL,
    paid_date DATE,
    status ENUM('pending', 'paid', 'overdue', 'cancelled') DEFAULT 'pending',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (debtor_id) REFERENCES debtors(id) ON DELETE CASCADE,
    INDEX idx_debtor (debtor_id),
    INDEX idx_due_date (due_date),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- 8. أيام العمل وتفاصيلها
-- Work Days and Details
-- ====================================

-- جدول أيام العمل
CREATE TABLE work_days (
    id INT AUTO_INCREMENT PRIMARY KEY,
    work_day_number VARCHAR(100) UNIQUE,
    project_id INT NOT NULL,
    work_date DATE NOT NULL,
    work_location VARCHAR(500),
    form_number VARCHAR(100),
    work_period_from TIME,
    work_period_to TIME,
    day_name VARCHAR(50),
    work_type VARCHAR(255),
    description TEXT,
    status ENUM('pending', 'in_progress', 'completed', 'cancelled') DEFAULT 'pending',
    total_labor_cost DECIMAL(15, 2) DEFAULT 0,
    total_equipment_cost DECIMAL(15, 2) DEFAULT 0,
    total_materials_cost DECIMAL(15, 2) DEFAULT 0,
    total_cost DECIMAL(15, 2) DEFAULT 0,
    notes TEXT,
    created_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_project (project_id),
    INDEX idx_work_date (work_date),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول العمالة في أيام العمل
CREATE TABLE work_day_labor (
    id INT AUTO_INCREMENT PRIMARY KEY,
    work_day_id INT NOT NULL,
    worker_name VARCHAR(255) NOT NULL,
    job_title VARCHAR(255),
    hours_worked DECIMAL(5, 2),
    hourly_rate DECIMAL(10, 2),
    daily_rate DECIMAL(10, 2),
    total_cost DECIMAL(12, 2),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (work_day_id) REFERENCES work_days(id) ON DELETE CASCADE,
    INDEX idx_work_day (work_day_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول المعدات في أيام العمل
CREATE TABLE work_day_equipment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    work_day_id INT NOT NULL,
    equipment_name VARCHAR(255) NOT NULL,
    equipment_type VARCHAR(255),
    quantity DECIMAL(10, 2),
    unit_of_measure VARCHAR(50),
    unit_price DECIMAL(12, 2),
    total_cost DECIMAL(15, 2),
    rental_hours DECIMAL(5, 2),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (work_day_id) REFERENCES work_days(id) ON DELETE CASCADE,
    INDEX idx_work_day (work_day_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- جدول المواد في أيام العمل
CREATE TABLE work_day_materials (
    id INT AUTO_INCREMENT PRIMARY KEY,
    work_day_id INT NOT NULL,
    material_name VARCHAR(255) NOT NULL,
    material_type VARCHAR(255),
    quantity DECIMAL(10, 2),
    unit_of_measure VARCHAR(50),
    unit_price DECIMAL(12, 2),
    total_cost DECIMAL(15, 2),
    supplier_name VARCHAR(255),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (work_day_id) REFERENCES work_days(id) ON DELETE CASCADE,
    INDEX idx_work_day (work_day_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ====================================
-- إنشاء الفهارس الإضافية للأداء
-- Additional Indexes for Performance
-- ====================================

CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_projects_created_at ON projects(created_at);
CREATE INDEX idx_tasks_created_at ON tasks(created_at);
CREATE INDEX idx_expenses_created_at ON expenses(created_at);
CREATE INDEX idx_income_created_at ON income(created_at);
CREATE INDEX idx_work_days_created_at ON work_days(created_at);

-- ====================================
-- إنشاء Views للتقارير
-- Create Views for Reports
-- ====================================

-- عرض ملخص المشاريع
CREATE OR REPLACE VIEW view_project_summary AS
SELECT 
    p.id,
    p.project_number,
    p.name,
    p.status,
    p.start_date,
    p.end_date,
    p.contract_value,
    p.actual_cost,
    p.estimated_cost,
    p.progress_percentage,
    u.full_name as manager_name,
    d.name as department_name,
    COUNT(DISTINCT t.id) as total_tasks,
    SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) as completed_tasks,
    COUNT(DISTINCT pt.team_member_id) as team_size,
    COALESCE(SUM(e.amount), 0) as total_expenses,
    COALESCE(SUM(i.amount), 0) as total_income
FROM projects p
LEFT JOIN users u ON p.manager_id = u.id
LEFT JOIN departments d ON p.department_id = d.id
LEFT JOIN tasks t ON p.id = t.project_id
LEFT JOIN project_team pt ON p.id = pt.project_id AND pt.is_active = TRUE
LEFT JOIN project_expenses pe ON p.id = pe.project_id
LEFT JOIN expenses e ON pe.expense_id = e.id
LEFT JOIN income i ON p.id = i.project_id
GROUP BY p.id;

-- عرض ملخص المهام
CREATE OR REPLACE VIEW view_task_summary AS
SELECT 
    t.id,
    t.task_number,
    t.title,
    t.status,
    t.priority,
    t.start_date,
    t.due_date,
    t.progress_percentage,
    p.name as project_name,
    tm.name as assigned_to_name,
    u.full_name as created_by_name,
    DATEDIFF(t.due_date, CURRENT_DATE) as days_remaining
FROM tasks t
LEFT JOIN projects p ON t.project_id = p.id
LEFT JOIN team_members tm ON t.assigned_to = tm.id
LEFT JOIN users u ON t.created_by = u.id;

-- عرض ملخص المصاريف
CREATE OR REPLACE VIEW view_expense_summary AS
SELECT 
    e.id,
    e.expense_number,
    e.description,
    e.amount,
    e.expense_date,
    e.status,
    et.name as expense_type,
    ec.name as category_name,
    p.name as project_name,
    d.name as department_name,
    u.full_name as created_by_name
FROM expenses e
LEFT JOIN expense_types et ON e.expense_type_id = et.id
LEFT JOIN expense_categories ec ON e.category_id = ec.id
LEFT JOIN projects p ON e.project_id = p.id
LEFT JOIN departments d ON e.department_id = d.id
LEFT JOIN users u ON e.created_by = u.id;

-- عرض ملخص الإيرادات
CREATE OR REPLACE VIEW view_income_summary AS
SELECT 
    i.id,
    i.income_number,
    i.description,
    i.amount,
    i.income_date,
    i.status,
    ic.name as category_name,
    p.name as project_name,
    i.client_name,
    u.full_name as created_by_name
FROM income i
LEFT JOIN income_categories ic ON i.category_id = ic.id
LEFT JOIN projects p ON i.project_id = p.id
LEFT JOIN users u ON i.created_by = u.id;

-- عرض ملخص المدينين
CREATE OR REPLACE VIEW view_debtor_summary AS
SELECT 
    d.id,
    d.debtor_number,
    d.name,
    d.company_name,
    d.total_debt,
    d.paid_amount,
    d.remaining_amount,
    d.due_date,
    d.status,
    p.name as project_name,
    DATEDIFF(d.due_date, CURRENT_DATE) as days_until_due,
    COUNT(dp.id) as payment_count,
    COALESCE(SUM(dp.amount), 0) as total_payments
FROM debtors d
LEFT JOIN projects p ON d.project_id = p.id
LEFT JOIN debtor_payments dp ON d.id = dp.debtor_id
GROUP BY d.id;

-- عرض ملخص أيام العمل
CREATE OR REPLACE VIEW view_work_day_summary AS
SELECT 
    wd.id,
    wd.work_day_number,
    wd.work_date,
    wd.work_location,
    wd.work_type,
    wd.status,
    p.name as project_name,
    wd.total_labor_cost,
    wd.total_equipment_cost,
    wd.total_materials_cost,
    wd.total_cost,
    COUNT(DISTINCT wdl.id) as labor_count,
    COUNT(DISTINCT wde.id) as equipment_count,
    COUNT(DISTINCT wdm.id) as materials_count
FROM work_days wd
LEFT JOIN projects p ON wd.project_id = p.id
LEFT JOIN work_day_labor wdl ON wd.id = wdl.work_day_id
LEFT JOIN work_day_equipment wde ON wd.id = wde.work_day_id
LEFT JOIN work_day_materials wdm ON wd.id = wdm.work_day_id
GROUP BY wd.id;

-- ====================================
-- نهاية ملف المخطط
-- End of Schema File
-- ====================================

