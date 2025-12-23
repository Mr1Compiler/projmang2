-- =============================================
-- STATUS ENUMS MIGRATION
-- =============================================

-- Create enum types
CREATE TYPE project_status AS ENUM ('pending', 'in_progress', 'completed');
CREATE TYPE work_day_status AS ENUM ('pending', 'in_progress', 'completed');
CREATE TYPE financial_status AS ENUM ('pending', 'approved', 'rejected');
CREATE TYPE debtor_status AS ENUM ('active', 'paid');
CREATE TYPE general_status AS ENUM ('active', 'inactive');

-- =============================================
-- ALTER TABLES TO USE ENUMS
-- =============================================

-- Projects
ALTER TABLE projects
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE project_status USING status::project_status,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Work Days
ALTER TABLE workDays
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE work_day_status USING status::work_day_status,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Expenses
ALTER TABLE expenses
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE financial_status USING status::financial_status,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Income
ALTER TABLE income
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE financial_status USING status::financial_status,
    ALTER COLUMN status SET DEFAULT 'pending';

-- Debtors
ALTER TABLE debtors
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE debtor_status USING status::debtor_status,
    ALTER COLUMN status SET DEFAULT 'active';

-- Pages
ALTER TABLE pages
    ALTER COLUMN status DROP DEFAULT,
    ALTER COLUMN status TYPE general_status USING status::general_status,
    ALTER COLUMN status SET DEFAULT 'active';

-- Users
ALTER TABLE users
    ALTER COLUMN status TYPE general_status USING status::general_status;

-- Work Categories
ALTER TABLE workCategories
    ALTER COLUMN status TYPE general_status USING status::general_status;

-- Work Sub Categories
ALTER TABLE workSubCategories
    ALTER COLUMN status TYPE general_status USING status::general_status;

-- =============================================
-- TRIGGER: Auto set project to 'in_progress' on first workDay
-- =============================================

CREATE OR REPLACE FUNCTION set_project_in_progress()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE projects
    SET status = 'in_progress'
    WHERE id = NEW.projectId
      AND status = 'pending';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_project_in_progress
    AFTER INSERT ON workDays
    FOR EACH ROW EXECUTE FUNCTION set_project_in_progress();
