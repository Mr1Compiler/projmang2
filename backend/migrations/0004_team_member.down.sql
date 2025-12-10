DROP TRIGGER IF EXISTS update_teamMembers_updated_at ON teamMembers;
DROP INDEX IF EXISTS idx_teamMembers_projectId;
DROP INDEX IF EXISTS idx_teamMembers_userId;
DROP TABLE IF EXISTS teamMembers;