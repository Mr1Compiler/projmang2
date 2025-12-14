# دليل التثبيت والإعداد
## Installation and Setup Guide

## 🚀 التثبيت السريع

### المتطلبات الأساسية

قبل البدء، تأكد من توفر:

- **MySQL** 8.0 أو أحدث (أو MariaDB 10.5+)
- **PHP** 7.4 أو أحدث
- **Composer** (اختياري)
- صلاحيات إنشاء قاعدة بيانات

### الخطوة 1: إنشاء قاعدة البيانات

```bash
# الاتصال بـ MySQL
mysql -u root -p

# أو باستخدام مستخدم آخر
mysql -u your_username -p
```

```sql
-- إنشاء قاعدة البيانات
CREATE DATABASE project_management_db
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;

-- إنشاء مستخدم خاص (اختياري ومُوصى به)
CREATE USER 'pm_user'@'localhost' IDENTIFIED BY 'strong_password_here';
GRANT ALL PRIVILEGES ON project_management_db.* TO 'pm_user'@'localhost';
FLUSH PRIVILEGES;

-- استخدام قاعدة البيانات
USE project_management_db;
```

### الخطوة 2: تنفيذ ملف المخطط

```bash
# من الطرفية (Terminal)
cd /Users/msi/Projects/project\ managemen./database/

# تنفيذ ملف المخطط
mysql -u pm_user -p project_management_db < schema.sql
```

أو من داخل MySQL:

```sql
USE project_management_db;
SOURCE /Users/msi/Projects/project\ managemen./database/schema.sql;
```

### الخطوة 3: إدراج البيانات الأولية

```bash
# تنفيذ ملف البيانات الأولية
mysql -u pm_user -p project_management_db < seed_data.sql
```

أو من داخل MySQL:

```sql
SOURCE /Users/msi/Projects/project\ managemen./database/seed_data.sql;
```

### الخطوة 4: تنفيذ الإجراءات المخزنة والمشغلات

```bash
# تنفيذ ملف الإجراءات المخزنة
mysql -u pm_user -p project_management_db < stored_procedures.sql

# تنفيذ ملف المشغلات
mysql -u pm_user -p project_management_db < triggers.sql
```

### الخطوة 5: إعداد ملف البيئة

```bash
# نسخ ملف .env.example
cp database/.env.example .env

# تحرير ملف .env
nano .env
```

قم بتحديث الإعدادات:

```env
DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=project_management_db
DB_USERNAME=pm_user
DB_PASSWORD=your_password_here
```

### الخطوة 6: التحقق من التثبيت

```sql
-- التحقق من الجداول
SHOW TABLES;

-- التحقق من البيانات
SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM projects;
SELECT COUNT(*) FROM tasks;

-- التحقق من الإجراءات المخزنة
SHOW PROCEDURE STATUS WHERE Db = 'project_management_db';

-- التحقق من المشغلات
SHOW TRIGGERS;

-- التحقق من الـ Views
SHOW FULL TABLES WHERE TABLE_TYPE LIKE 'VIEW';
```

## 📋 التثبيت المفصل

### التثبيت على Windows

#### 1. تثبيت XAMPP أو WAMP

قم بتنزيل وتثبيت [XAMPP](https://www.apachefriends.org/) أو [WAMP](https://www.wampserver.com/)

#### 2. بدء خدمات MySQL

افتح لوحة تحكم XAMPP/WAMP وقم بتشغيل MySQL

#### 3. فتح phpMyAdmin

افتح المتصفح واذهب إلى: `http://localhost/phpmyadmin`

#### 4. إنشاء قاعدة البيانات

- انقر على "New" (جديد)
- اسم قاعدة البيانات: `project_management_db`
- الترميز: `utf8mb4_unicode_ci`
- انقر "Create" (إنشاء)

#### 5. استيراد الملفات

- اختر قاعدة البيانات
- انقر على تبويب "Import" (استيراد)
- اختر ملف `schema.sql`
- انقر "Go" (تنفيذ)
- كرر العملية لملفات: `seed_data.sql`, `stored_procedures.sql`, `triggers.sql`

### التثبيت على macOS

#### استخدام Homebrew

```bash
# تثبيت MySQL
brew install mysql

# بدء خدمة MySQL
brew services start mysql

# تأمين التثبيت
mysql_secure_installation

# الاتصال بـ MySQL
mysql -u root -p
```

ثم اتبع الخطوات العامة أعلاه.

### التثبيت على Linux (Ubuntu/Debian)

```bash
# تحديث النظام
sudo apt update

# تثبيت MySQL
sudo apt install mysql-server

# تأمين التثبيت
sudo mysql_secure_installation

# الاتصال بـ MySQL
sudo mysql -u root -p
```

ثم اتبع الخطوات العامة أعلاه.

## 🔧 الإعداد المتقدم

### تحسين الأداء

قم بتحرير ملف `my.cnf` أو `my.ini`:

```ini
[mysqld]
# حجم ذاكرة التخزين المؤقت
innodb_buffer_pool_size = 1G

# حجم ملف السجل
innodb_log_file_size = 256M

# عدد الاتصالات المتزامنة
max_connections = 200

# حجم جداول الذاكرة المؤقتة
tmp_table_size = 64M
max_heap_table_size = 64M

# التخزين المؤقت للاستعلامات
query_cache_type = 1
query_cache_size = 64M
query_cache_limit = 2M
```

أعد تشغيل MySQL:

```bash
# Linux
sudo systemctl restart mysql

# macOS
brew services restart mysql

# Windows (كمسؤول)
net stop MySQL
net start MySQL
```

### إعداد النسخ الاحتياطي التلقائي

#### على Linux/macOS

```bash
# إعطاء صلاحيات التنفيذ
chmod +x database/maintenance/backup.sh

# إضافة مهمة cron للنسخ الاحتياطي اليومي
crontab -e

# إضافة السطر التالي (نسخة احتياطية يومية عند الساعة 2 صباحاً)
0 2 * * * /path/to/database/maintenance/backup.sh full
```

#### على Windows

استخدم Task Scheduler:

1. افتح Task Scheduler
2. Create Basic Task
3. Trigger: Daily at 2:00 AM
4. Action: Start a Program
5. Program: `mysql/bin/mysqldump.exe`
6. Arguments: `-u username -ppassword project_management_db > backup.sql`

### إعداد الصيانة الدورية

```bash
# إنشاء نص صيانة دوري
crontab -e

# تنفيذ الصيانة أسبوعياً (كل أحد عند الساعة 3 صباحاً)
0 3 * * 0 mysql -u username -p < /path/to/database/maintenance/optimize.sql
```

## 🔐 الأمان

### تأمين قاعدة البيانات

1. **تغيير كلمات المرور الافتراضية**

```sql
-- تغيير كلمات مرور المستخدمين
UPDATE users SET password = '$2y$10$NEW_HASHED_PASSWORD' WHERE username = 'admin';
```

2. **تقييد الوصول**

```sql
-- إنشاء مستخدم للقراءة فقط
CREATE USER 'reader'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT ON project_management_db.* TO 'reader'@'localhost';
```

3. **تفعيل SSL**

```ini
[mysqld]
require_secure_transport = ON
ssl-ca=/path/to/ca.pem
ssl-cert=/path/to/server-cert.pem
ssl-key=/path/to/server-key.pem
```

### جدار الحماية

```bash
# Linux - السماح فقط للاتصالات المحلية
sudo ufw allow from 127.0.0.1 to any port 3306

# أو السماح من شبكة محددة
sudo ufw allow from 192.168.1.0/24 to any port 3306
```

## 🧪 الاختبار

### اختبار الاتصال

```php
<?php
// test_connection.php
require_once 'database/Database.php';

try {
    $db = Database::getInstance();
    echo "✅ الاتصال بقاعدة البيانات ناجح!\n";
    
    // اختبار استعلام بسيط
    $result = $db->selectOne("SELECT COUNT(*) as count FROM users");
    echo "عدد المستخدمين: " . $result['count'] . "\n";
    
} catch (Exception $e) {
    echo "❌ خطأ في الاتصال: " . $e->getMessage() . "\n";
}
?>
```

```bash
php test_connection.php
```

### تشغيل الأمثلة

```bash
# تشغيل ملف الأمثلة
php database/examples/usage_examples.php
```

## 🔄 الترقية

### من إصدار سابق

```sql
-- 1. نسخة احتياطية أولاً!
mysqldump -u username -p project_management_db > backup_before_upgrade.sql

-- 2. تنفيذ نصوص الترقية (إن وجدت)
SOURCE /path/to/upgrade_v1_to_v2.sql;

-- 3. تحديث الإجراءات المخزنة والمشغلات
SOURCE /path/to/stored_procedures.sql;
SOURCE /path/to/triggers.sql;

-- 4. تحليل وتحسين الجداول
SOURCE /path/to/maintenance/optimize.sql;
```

## ❓ استكشاف الأخطاء

### المشكلة: لا يمكن الاتصال بقاعدة البيانات

```bash
# التحقق من أن MySQL يعمل
sudo systemctl status mysql

# التحقق من المنفذ
netstat -tulpn | grep 3306

# التحقق من ملف السجل
sudo tail -f /var/log/mysql/error.log
```

### المشكلة: خطأ في الترميز

```sql
-- التحقق من الترميز
SHOW VARIABLES LIKE 'character_set%';

-- تغيير الترميز
ALTER DATABASE project_management_db 
    CHARACTER SET utf8mb4 
    COLLATE utf8mb4_unicode_ci;
```

### المشكلة: بطء الاستعلامات

```sql
-- تفعيل سجل الاستعلامات البطيئة
SET GLOBAL slow_query_log = 'ON';
SET GLOBAL long_query_time = 2;

-- عرض الاستعلامات البطيئة
SELECT * FROM mysql.slow_log ORDER BY query_time DESC LIMIT 10;
```

## 📞 الدعم

للحصول على المساعدة:

- **البريد الإلكتروني**: support@erticaz.com
- **الموقع**: www.erticaz.com
- **الوثائق**: راجع ملف `README.md`

## 📄 الترخيص

© 2024 ارتكاز للحلول البرمجية - جميع الحقوق محفوظة

