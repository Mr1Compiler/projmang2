-- Reverse migration: Drop all tables

-- Drop triggers first (in reverse order of creation)
DROP TRIGGER IF EXISTS update_income_updated_at ON income;
DROP TRIGGER IF EXISTS update_expenses_updated_at ON expenses;
DROP TRIGGER IF EXISTS update_debtors_updated_at ON debtors;
DROP TRIGGER IF EXISTS update_workDayEquipment_updated_at ON workDayEquipment;
DROP TRIGGER IF EXISTS update_workDayLabor_updated_at ON workDayLabor;
DROP TRIGGER IF EXISTS update_workDayMaterials_updated_at ON workDayMaterials;
DROP TRIGGER IF EXISTS update_workDays_updated_at ON workDays;
DROP TRIGGER IF EXISTS update_projects_updated_at ON projects;
DROP TRIGGER IF EXISTS update_workSubCategories_updated_at ON workSubCategories;
DROP TRIGGER IF EXISTS update_workCategories_updated_at ON workCategories;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP TRIGGER IF EXISTS update_roles_updated_at ON roles;
DROP TRIGGER IF EXISTS update_pages_updated_at ON pages;

-- Drop indexes
DROP INDEX IF EXISTS idx_workDays_workDate;
DROP INDEX IF EXISTS idx_projects_status;
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_expenses_debtorId;
DROP INDEX IF EXISTS idx_expenses_projectId;
DROP INDEX IF EXISTS idx_workDayEquipment_workDayId;
DROP INDEX IF EXISTS idx_workDayLabor_workDayId;
DROP INDEX IF EXISTS idx_workDayMaterials_workDayId;
DROP INDEX IF EXISTS idx_workDays_workSubCategoryId;
DROP INDEX IF EXISTS idx_workDays_projectId;
DROP INDEX IF EXISTS idx_workSubCategories_categoryId;
DROP INDEX IF EXISTS idx_userRoles_roleId;
DROP INDEX IF EXISTS idx_userRoles_userId;
DROP INDEX IF EXISTS idx_rolePages_pageId;
DROP INDEX IF EXISTS idx_rolePages_roleId;

-- Drop tables (in reverse order of dependencies)
DROP TABLE IF EXISTS income;
DROP TABLE IF EXISTS expenses;
DROP TABLE IF EXISTS debtors;
DROP TABLE IF EXISTS workDayEquipment;
DROP TABLE IF EXISTS workDayLabor;
DROP TABLE IF EXISTS workDayMaterials;
DROP TABLE IF EXISTS workDays;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS workSubCategories;
DROP TABLE IF EXISTS workCategories;
DROP TABLE IF EXISTS userRoles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS rolePages;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS pages;

-- Drop function (must be last, after all triggers are dropped)
DROP FUNCTION IF EXISTS update_updated_at();
