UPDATE schema_migrations SET dirty = false;
-- Optionally ensure version is correct if you know all tables exist
-- UPDATE schema_migrations SET version = 6, dirty = false;
