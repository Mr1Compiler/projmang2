-- =============================================
-- ROLLBACK STATUS ENUMS MIGRATION
-- =============================================

-- Drop trigger first
DROP TRIGGER IF EXISTS trigger_project_in_progress ON workDays;
DROP FUNCTION IF EXISTS set_project_in_progress();

-- =============================================
-- REVERT TABLES BACK TO VARCHAR
-- =============================================

-- Projects
ALTER TABLE projects
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text,
    ALTER COLUMN status SET DEFAULT 'draft';

-- Work Days
ALTER TABLE workDays
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Expenses
ALTER TABLE expenses
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Income
ALTER TABLE income
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Debtors
ALTER TABLE debtors
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text,
    ALTER COLUMN status SET DEFAULT 'active';

-- Pages
ALTER TABLE pages
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text,
    ALTER COLUMN status SET DEFAULT 'active';

-- Users
ALTER TABLE users
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text;

-- Work Categories
ALTER TABLE workCategories
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text;

-- Work Sub Categories
ALTER TABLE workSubCategories
    ALTER COLUMN status TYPE VARCHAR(50) USING status::text;

-- =============================================
-- DROP ENUM TYPES
-- =============================================

DROP TYPE IF EXISTS project_status;
DROP TYPE IF EXISTS work_day_status;
DROP TYPE IF EXISTS financial_status;
DROP TYPE IF EXISTS debtor_status;
DROP TYPE IF EXISTS general_status;
