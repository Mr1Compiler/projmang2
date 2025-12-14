#!/bin/bash

# ====================================
# نص التثبيت السريع لقاعدة البيانات
# Quick Database Installation Script
# ====================================

# الألوان
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# طباعة رسالة ملونة
print_message() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# طباعة العنوان
print_header() {
    echo ""
    print_message $BLUE "===================================="
    print_message $BLUE "  نظام إدارة المشاريع"
    print_message $BLUE "  Project Management System"
    print_message $BLUE "  نص التثبيت السريع"
    print_message $BLUE "===================================="
    echo ""
}

# التحقق من MySQL
check_mysql() {
    if ! command -v mysql &> /dev/null; then
        print_message $RED "❌ خطأ: MySQL غير مثبت على النظام"
        print_message $YELLOW "يرجى تثبيت MySQL أولاً"
        exit 1
    fi
    print_message $GREEN "✅ MySQL مثبت"
}

# الحصول على معلومات الاتصال
get_credentials() {
    echo ""
    print_message $YELLOW "معلومات الاتصال بقاعدة البيانات:"
    echo ""
    
    read -p "اسم المستخدم [root]: " DB_USER
    DB_USER=${DB_USER:-root}
    
    read -sp "كلمة المرور: " DB_PASS
    echo ""
    
    read -p "اسم قاعدة البيانات [project_management_db]: " DB_NAME
    DB_NAME=${DB_NAME:-project_management_db}
    
    read -p "المضيف [localhost]: " DB_HOST
    DB_HOST=${DB_HOST:-localhost}
    
    echo ""
}

# اختبار الاتصال
test_connection() {
    print_message $YELLOW "🔌 اختبار الاتصال..."
    
    if mysql -h"$DB_HOST" -u"$DB_USER" -p"$DB_PASS" -e "SELECT 1;" &> /dev/null; then
        print_message $GREEN "✅ الاتصال ناجح"
        return 0
    else
        print_message $RED "❌ فشل الاتصال"
        print_message $YELLOW "يرجى التحقق من اسم المستخدم وكلمة المرور"
        exit 1
    fi
}

# إنشاء قاعدة البيانات
create_database() {
    print_message $YELLOW "📦 إنشاء قاعدة البيانات..."
    
    mysql -h"$DB_HOST" -u"$DB_USER" -p"$DB_PASS" << EOF
CREATE DATABASE IF NOT EXISTS $DB_NAME
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;
EOF
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✅ تم إنشاء قاعدة البيانات: $DB_NAME"
    else
        print_message $RED "❌ فشل إنشاء قاعدة البيانات"
        exit 1
    fi
}

# تنفيذ ملف SQL
execute_sql_file() {
    local file=$1
    local description=$2
    
    if [ ! -f "$file" ]; then
        print_message $RED "❌ الملف غير موجود: $file"
        return 1
    fi
    
    print_message $YELLOW "⏳ $description..."
    
    mysql -h"$DB_HOST" -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$file"
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✅ تم بنجاح: $description"
        return 0
    else
        print_message $RED "❌ فشل: $description"
        return 1
    fi
}

# إعداد ملف البيئة
setup_env_file() {
    print_message $YELLOW "⚙️  إعداد ملف البيئة..."
    
    local env_file=".env"
    local env_example="database/.env.example"
    
    if [ -f "$env_file" ]; then
        print_message $YELLOW "⚠️  ملف .env موجود بالفعل"
        read -p "هل تريد استبداله؟ (y/n): " replace
        if [ "$replace" != "y" ]; then
            print_message $YELLOW "تم تخطي إعداد ملف .env"
            return 0
        fi
    fi
    
    if [ -f "$env_example" ]; then
        cp "$env_example" "$env_file"
        
        # تحديث الإعدادات
        sed -i.bak "s/DB_HOST=.*/DB_HOST=$DB_HOST/" "$env_file"
        sed -i.bak "s/DB_DATABASE=.*/DB_DATABASE=$DB_NAME/" "$env_file"
        sed -i.bak "s/DB_USERNAME=.*/DB_USERNAME=$DB_USER/" "$env_file"
        sed -i.bak "s/DB_PASSWORD=.*/DB_PASSWORD=$DB_PASS/" "$env_file"
        
        rm -f "${env_file}.bak"
        
        print_message $GREEN "✅ تم إنشاء ملف .env"
    else
        print_message $YELLOW "⚠️  ملف .env.example غير موجود"
    fi
}

# عرض معلومات ما بعد التثبيت
show_post_install_info() {
    echo ""
    print_message $GREEN "===================================="
    print_message $GREEN "  اكتمل التثبيت بنجاح! 🎉"
    print_message $GREEN "===================================="
    echo ""
    
    print_message $BLUE "معلومات الاتصال:"
    echo "  المضيف: $DB_HOST"
    echo "  قاعدة البيانات: $DB_NAME"
    echo "  المستخدم: $DB_USER"
    echo ""
    
    print_message $BLUE "المستخدمون الافتراضيون:"
    echo "  admin@erticaz.com - password123 (مدير النظام)"
    echo "  manager1@erticaz.com - password123 (مدير مشروع)"
    echo "  engineer1@erticaz.com - password123 (مهندس)"
    echo "  accountant@erticaz.com - password123 (محاسب)"
    echo ""
    
    print_message $RED "⚠️  تحذير: يجب تغيير كلمات المرور الافتراضية فوراً!"
    echo ""
    
    print_message $BLUE "الخطوات التالية:"
    echo "  1. تحديث ملف .env بالإعدادات الصحيحة"
    echo "  2. تغيير كلمات المرور الافتراضية"
    echo "  3. اختبار الاتصال بقاعدة البيانات"
    echo "  4. راجع ملف README.md للمزيد من المعلومات"
    echo ""
    
    print_message $BLUE "الوثائق:"
    echo "  - README.md - الدليل الشامل"
    echo "  - INSTALLATION.md - دليل التثبيت"
    echo "  - DATABASE_OVERVIEW.md - نظرة عامة"
    echo "  - SUMMARY.md - الملخص"
    echo ""
    
    print_message $GREEN "===================================="
    print_message $GREEN "شكراً لاستخدامك نظام إدارة المشاريع"
    print_message $GREEN "© 2024 ارتكاز للحلول البرمجية"
    print_message $GREEN "===================================="
    echo ""
}

# البرنامج الرئيسي
main() {
    print_header
    
    # التحقق من MySQL
    check_mysql
    
    # الحصول على معلومات الاتصال
    get_credentials
    
    # اختبار الاتصال
    test_connection
    
    # إنشاء قاعدة البيانات
    create_database
    
    # تنفيذ ملفات SQL
    echo ""
    print_message $BLUE "تنفيذ ملفات قاعدة البيانات..."
    echo ""
    
    execute_sql_file "database/schema.sql" "تنفيذ مخطط قاعدة البيانات"
    execute_sql_file "database/seed_data.sql" "إدراج البيانات الأولية"
    execute_sql_file "database/stored_procedures.sql" "إنشاء الإجراءات المخزنة"
    execute_sql_file "database/triggers.sql" "إنشاء المشغلات"
    
    # إعداد ملف البيئة
    echo ""
    setup_env_file
    
    # عرض معلومات ما بعد التثبيت
    show_post_install_info
}

# تشغيل البرنامج
main

