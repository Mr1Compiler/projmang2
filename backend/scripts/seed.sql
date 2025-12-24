-- =============================================
-- Seed Script: Fresh Database Setup
-- Run this after migrations on a clean database
-- =============================================

-- =============================================
-- 1. CLEAR ALL DATA (for fresh start)
-- =============================================
TRUNCATE TABLE rolepages, userroles, users, pages, roles RESTART IDENTITY CASCADE;

-- =============================================
-- 2. INSERT PAGES
-- =============================================

INSERT INTO pages (name, route, icon, status) VALUES
-- صفحات السايدبار (تظهر في القائمة الجانبية)
('لوحة التحكم', '/', 'mdi-view-dashboard', 'active'),
('إدارة المشاريع', '/projects', 'mdi-folder-multiple', 'active'),
('المصروفات', '/expenses', 'mdi-chart-line', 'active'),
('الإيرادات', '/income', 'mdi-trending-up', 'active'),
('المديونون', '/debtors', 'mdi-credit-card', 'active'),
('المستخدمين', '/users', 'mdi-account-multiple', 'active'),
('التصنيفات', '/categories', 'mdi-tag-multiple', 'active'),
('الأدوار', '/roles', 'mdi-shield-account', 'active'),
('إدارة الصفحات', '/pages', 'mdi-file-document-multiple', 'active'),
-- صفحات API فقط (لا تظهر في السايدبار - للصلاحيات فقط)
('أيام العمل', '/workdays', 'mdi-calendar-check', 'active'),
('العمالة', '/workday-labor', 'mdi-account-hard-hat', 'active'),
('المعدات', '/workday-equipment', 'mdi-tools', 'active'),
('المواد', '/workday-materials', 'mdi-package-variant', 'active'),
('صلاحيات الأدوار', '/role-pages', 'mdi-key', 'active'),
('أدوار المستخدمين', '/user-roles', 'mdi-account-key', 'active'),
('أعضاء الفريق', '/team-members', 'mdi-account-group', 'active'),
('تصنيفات العمل', '/work-categories', 'mdi-shape', 'active'),
('تصنيفات العمل الفرعية', '/work-subcategories', 'mdi-shape-outline', 'active');

-- =============================================
-- 3. INSERT ROLES
-- =============================================

INSERT INTO roles (name, description) VALUES
('مدير النظام', 'مدير النظام - صلاحيات إدارة المستخدمين والأدوار والصفحات فقط');

-- =============================================
-- 4. ASSIGN SUPERADMIN PERMISSIONS
-- Only management pages: users, roles, pages + their APIs
-- =============================================

DO $$
DECLARE
    admin_role_id INTEGER;
    page_id INTEGER;
BEGIN
    SELECT id INTO admin_role_id FROM roles WHERE name = 'مدير النظام';

    -- المستخدمين
    SELECT id INTO page_id FROM pages WHERE route = '/users';
    INSERT INTO rolepages (roleid, pageid, permissions)
    VALUES (admin_role_id, page_id, '["read","create","update","delete","updatePassword","updateStatus"]');

    -- الأدوار
    SELECT id INTO page_id FROM pages WHERE route = '/roles';
    INSERT INTO rolepages (roleid, pageid, permissions)
    VALUES (admin_role_id, page_id, '["read","create","update","delete"]');

    -- إدارة الصفحات
    SELECT id INTO page_id FROM pages WHERE route = '/pages';
    INSERT INTO rolepages (roleid, pageid, permissions)
    VALUES (admin_role_id, page_id, '["read","create","update","delete"]');

    -- صلاحيات الأدوار (API)
    SELECT id INTO page_id FROM pages WHERE route = '/role-pages';
    INSERT INTO rolepages (roleid, pageid, permissions)
    VALUES (admin_role_id, page_id, '["read","create","update","delete"]');

    -- أدوار المستخدمين (API)
    SELECT id INTO page_id FROM pages WHERE route = '/user-roles';
    INSERT INTO rolepages (roleid, pageid, permissions)
    VALUES (admin_role_id, page_id, '["read","create","update","delete"]');
END $$;

-- =============================================
-- 5. CREATE SUPERADMIN USER
-- =============================================

-- Password: Admin@123 (bcrypt hashed)
-- CHANGE THIS PASSWORD AFTER FIRST LOGIN!
INSERT INTO users (username, email, password, fullname, phone, jobtitle, status) VALUES
('superadmin', 'admin@system.com', '$2b$10$TZSD0FXAuC6R2SwRWsUhsuDJcVBgoNIcfj76wtdUI9lggU0euFVBC', 'مدير النظام', '0000000000', 'مدير النظام', 'active');

-- =============================================
-- 6. ASSIGN ROLE TO SUPERADMIN USER
-- =============================================

DO $$
DECLARE
    superadmin_user_id INTEGER;
    admin_role_id INTEGER;
BEGIN
    SELECT id INTO superadmin_user_id FROM users WHERE username = 'superadmin';
    SELECT id INTO admin_role_id FROM roles WHERE name = 'مدير النظام';

    INSERT INTO userroles (userid, roleid)
    VALUES (superadmin_user_id, admin_role_id);
END $$;

-- =============================================
-- DONE! Login with:
-- Username: superadmin
-- Password: Admin@123
--
-- Then create مدير role with all pages for client admin
-- =============================================
