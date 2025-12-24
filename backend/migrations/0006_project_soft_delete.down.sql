-- Remove soft delete column
DROP INDEX IF EXISTS idx_projects_is_active;
ALTER TABLE projects DROP COLUMN IF EXISTS isActive;
