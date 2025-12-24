-- Add isActive column for soft deletion
ALTER TABLE projects ADD COLUMN isActive BOOLEAN NOT NULL DEFAULT TRUE;

-- Create index for better query performance
CREATE INDEX idx_projects_is_active ON projects(isActive);
