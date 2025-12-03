-- Project Management Database Schema
-- Generated from ERD diagram

-- =============================================
-- ROLES AND PERMISSIONS
-- =============================================

CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    route VARCHAR(255),
    status VARCHAR(50),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rolePages (
    id SERIAL PRIMARY KEY,
    roleId INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    pageId INTEGER NOT NULL REFERENCES pages(id) ON DELETE CASCADE,
    permissions VARCHAR(255),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- USERS
-- =============================================

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    fullName VARCHAR(255),
    phone VARCHAR(50),
    avatar VARCHAR(255),
    jobTitle VARCHAR(255),
    status VARCHAR(50),
    lastLogin TIMESTAMP,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE userRoles (
    id SERIAL PRIMARY KEY,
    userId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    roleId INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- WORK CATEGORIES
-- =============================================

CREATE TABLE workCategories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workSubCategories (
    id SERIAL PRIMARY KEY,
    categoryId INTEGER NOT NULL REFERENCES workCategories(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    percentage DECIMAL(5, 2),
    status VARCHAR(50),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- PROJECTS
-- =============================================

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255),
    description TEXT,
    clientPhone VARCHAR(50),
    location VARCHAR(255),
    startDate DATE,
    duration INTEGER,
    warningCost DECIMAL(15, 2),
    totalCost DECIMAL(15, 2),
    status VARCHAR(50),
    progressPercentage DECIMAL(5, 2),
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- WORK DAYS
-- =============================================

CREATE TABLE workDays (
    id SERIAL PRIMARY KEY,
    projectId INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    workSubCategoryId INTEGER REFERENCES workSubCategories(id) ON DELETE SET NULL,
    workDate DATE NOT NULL,
    description TEXT,
    status VARCHAR(50),
    totalCost DECIMAL(15, 2),
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workDayMaterials (
    id SERIAL PRIMARY KEY,
    workDayId INTEGER NOT NULL REFERENCES workDays(id) ON DELETE CASCADE,
    materialName VARCHAR(255) NOT NULL,
    quantity DECIMAL(15, 2) NOT NULL DEFAULT 1,
    cost DECIMAL(15, 2) NOT NULL DEFAULT 0,
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workDayLabor (
    id SERIAL PRIMARY KEY,
    workDayId INTEGER NOT NULL REFERENCES workDays(id) ON DELETE CASCADE,
    workerName VARCHAR(255) NOT NULL,
    jobTitle VARCHAR(255),
    phone VARCHAR(50),
    address VARCHAR(255),
    quantity DECIMAL(15, 2) NOT NULL DEFAULT 1,
    cost DECIMAL(15, 2) NOT NULL DEFAULT 0,
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workDayEquipment (
    id SERIAL PRIMARY KEY,
    workDayId INTEGER NOT NULL REFERENCES workDays(id) ON DELETE CASCADE,
    equipmentName VARCHAR(255) NOT NULL,
    equipmentType VARCHAR(255),
    quantity DECIMAL(15, 2) NOT NULL DEFAULT 1,
    cost DECIMAL(15, 2) NOT NULL DEFAULT 0,
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- DEBTORS
-- =============================================

CREATE TABLE debtors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50),
    totalDebt DECIMAL(15, 2),
    currency VARCHAR(10),
    dueDate DATE,
    status VARCHAR(50),
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- FINANCES
-- =============================================

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    type VARCHAR(255),
    expenseDate DATE,
    projectId INTEGER REFERENCES projects(id) ON DELETE SET NULL,
    isDebtor BOOLEAN DEFAULT FALSE,
    debtorId INTEGER DEFAULT NULL REFERENCES debtors(id) ON DELETE SET NULL,
    status VARCHAR(50),
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE income (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    type VARCHAR(255),
    incomeDate DATE,
    status VARCHAR(50),
    notes TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =============================================
-- INDEXES
-- =============================================

CREATE INDEX idx_rolePages_roleId ON rolePages(roleId);
CREATE INDEX idx_rolePages_pageId ON rolePages(pageId);
CREATE INDEX idx_userRoles_userId ON userRoles(userId);
CREATE INDEX idx_userRoles_roleId ON userRoles(roleId);
CREATE INDEX idx_workSubCategories_categoryId ON workSubCategories(categoryId);
CREATE INDEX idx_workDays_projectId ON workDays(projectId);
CREATE INDEX idx_workDays_workSubCategoryId ON workDays(workSubCategoryId);
CREATE INDEX idx_workDayMaterials_workDayId ON workDayMaterials(workDayId);
CREATE INDEX idx_workDayLabor_workDayId ON workDayLabor(workDayId);
CREATE INDEX idx_workDayEquipment_workDayId ON workDayEquipment(workDayId);
CREATE INDEX idx_expenses_projectId ON expenses(projectId);
CREATE INDEX idx_expenses_debtorId ON expenses(debtorId);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_projects_status ON projects(status);
CREATE INDEX idx_workDays_workDate ON workDays(workDate);
