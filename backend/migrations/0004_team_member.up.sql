CREATE TABLE teamMembers (
    id SERIAL PRIMARY KEY,
    projectId INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    userId INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(projectId, userId)
);

CREATE INDEX idx_teamMembers_projectId ON teamMembers(projectId);
CREATE INDEX idx_teamMembers_userId ON teamMembers(userId);

CREATE TRIGGER update_teamMembers_updated_at
    BEFORE UPDATE ON teamMembers
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();