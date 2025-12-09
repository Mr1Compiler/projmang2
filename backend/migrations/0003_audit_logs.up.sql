CREATE TABLE auditLogs (
    id SERIAL PRIMARY KEY,
    actorId INT REFERENCES users(id) ON DELETE SET NULL,
    action VARCHAR(50) NOT NULL,
    targetType VARCHAR(50) NOT NULL,
    targetId INT,
    method VARCHAR(10) NOT NULL,
    status VARCHAR(20) NOT NULL,
    ipAddress VARCHAR(45) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_auditLogs_actorId ON auditLogs(actorId);
CREATE INDEX idx_auditLogs_target ON auditLogs(targetType, targetId);
CREATE INDEX idx_auditLogs_createdAt ON auditLogs(createdAt);
