-- Reverse migration: Drop all tables

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
