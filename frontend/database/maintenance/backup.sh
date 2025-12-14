#!/bin/bash

# ====================================
# نص النسخ الاحتياطي لقاعدة البيانات
# Database Backup Script
# ====================================

# الألوان للرسائل
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# المتغيرات
DB_HOST="localhost"
DB_PORT="3306"
DB_NAME="project_management_db"
DB_USER="root"
DB_PASS=""

# مسارات النسخ الاحتياطي
BACKUP_DIR="/Users/msi/Projects/project managemen./database/backups"
DATE=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="$BACKUP_DIR/backup_$DATE.sql"
BACKUP_FILE_GZ="$BACKUP_DIR/backup_$DATE.sql.gz"

# عدد الأيام للاحتفاظ بالنسخ الاحتياطية
RETENTION_DAYS=30

# ====================================
# الدوال
# ====================================

# دالة لطباعة الرسائل
print_message() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# دالة للتحقق من وجود mysqldump
check_mysqldump() {
    if ! command -v mysqldump &> /dev/null; then
        print_message $RED "❌ خطأ: mysqldump غير مثبت"
        exit 1
    fi
}

# دالة لإنشاء مجلد النسخ الاحتياطي
create_backup_dir() {
    if [ ! -d "$BACKUP_DIR" ]; then
        mkdir -p "$BACKUP_DIR"
        print_message $GREEN "✅ تم إنشاء مجلد النسخ الاحتياطي: $BACKUP_DIR"
    fi
}

# دالة لعمل نسخة احتياطية كاملة
full_backup() {
    print_message $YELLOW "📦 بدء النسخ الاحتياطي الكامل..."
    
    mysqldump \
        --host=$DB_HOST \
        --port=$DB_PORT \
        --user=$DB_USER \
        --password=$DB_PASS \
        --databases $DB_NAME \
        --routines \
        --triggers \
        --events \
        --single-transaction \
        --quick \
        --lock-tables=false \
        > "$BACKUP_FILE"
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✅ تم إنشاء النسخة الاحتياطية: $BACKUP_FILE"
        
        # ضغط النسخة الاحتياطية
        gzip "$BACKUP_FILE"
        print_message $GREEN "✅ تم ضغط النسخة الاحتياطية: $BACKUP_FILE_GZ"
        
        # حساب حجم الملف
        FILE_SIZE=$(du -h "$BACKUP_FILE_GZ" | cut -f1)
        print_message $GREEN "📊 حجم النسخة الاحتياطية: $FILE_SIZE"
    else
        print_message $RED "❌ فشل إنشاء النسخة الاحتياطية"
        exit 1
    fi
}

# دالة لعمل نسخة احتياطية للبيانات فقط
data_only_backup() {
    local DATA_BACKUP_FILE="$BACKUP_DIR/data_backup_$DATE.sql"
    
    print_message $YELLOW "📦 بدء نسخ احتياطي للبيانات فقط..."
    
    mysqldump \
        --host=$DB_HOST \
        --port=$DB_PORT \
        --user=$DB_USER \
        --password=$DB_PASS \
        --no-create-info \
        --skip-triggers \
        $DB_NAME \
        > "$DATA_BACKUP_FILE"
    
    if [ $? -eq 0 ]; then
        gzip "$DATA_BACKUP_FILE"
        print_message $GREEN "✅ تم إنشاء نسخة احتياطية للبيانات: ${DATA_BACKUP_FILE}.gz"
    else
        print_message $RED "❌ فشل إنشاء نسخة احتياطية للبيانات"
    fi
}

# دالة لعمل نسخة احتياطية للمخطط فقط
schema_only_backup() {
    local SCHEMA_BACKUP_FILE="$BACKUP_DIR/schema_backup_$DATE.sql"
    
    print_message $YELLOW "📦 بدء نسخ احتياطي للمخطط فقط..."
    
    mysqldump \
        --host=$DB_HOST \
        --port=$DB_PORT \
        --user=$DB_USER \
        --password=$DB_PASS \
        --no-data \
        --routines \
        --triggers \
        --events \
        $DB_NAME \
        > "$SCHEMA_BACKUP_FILE"
    
    if [ $? -eq 0 ]; then
        gzip "$SCHEMA_BACKUP_FILE"
        print_message $GREEN "✅ تم إنشاء نسخة احتياطية للمخطط: ${SCHEMA_BACKUP_FILE}.gz"
    else
        print_message $RED "❌ فشل إنشاء نسخة احتياطية للمخطط"
    fi
}

# دالة لحذف النسخ الاحتياطية القديمة
cleanup_old_backups() {
    print_message $YELLOW "🧹 تنظيف النسخ الاحتياطية القديمة (أقدم من $RETENTION_DAYS يوم)..."
    
    find "$BACKUP_DIR" -name "backup_*.sql.gz" -type f -mtime +$RETENTION_DAYS -delete
    find "$BACKUP_DIR" -name "data_backup_*.sql.gz" -type f -mtime +$RETENTION_DAYS -delete
    find "$BACKUP_DIR" -name "schema_backup_*.sql.gz" -type f -mtime +$RETENTION_DAYS -delete
    
    print_message $GREEN "✅ تم تنظيف النسخ الاحتياطية القديمة"
}

# دالة لعرض قائمة النسخ الاحتياطية
list_backups() {
    print_message $YELLOW "📋 قائمة النسخ الاحتياطية:"
    echo ""
    ls -lh "$BACKUP_DIR"/*.sql.gz 2>/dev/null | awk '{print $9, "(" $5 ")"}'
    echo ""
}

# دالة لاستعادة نسخة احتياطية
restore_backup() {
    local RESTORE_FILE=$1
    
    if [ ! -f "$RESTORE_FILE" ]; then
        print_message $RED "❌ الملف غير موجود: $RESTORE_FILE"
        exit 1
    fi
    
    print_message $YELLOW "⚠️  تحذير: سيتم استبدال قاعدة البيانات الحالية!"
    read -p "هل أنت متأكد؟ (yes/no): " confirm
    
    if [ "$confirm" != "yes" ]; then
        print_message $YELLOW "⚠️  تم إلغاء عملية الاستعادة"
        exit 0
    fi
    
    print_message $YELLOW "📥 بدء استعادة النسخة الاحتياطية..."
    
    # فك ضغط الملف إذا كان مضغوطاً
    if [[ $RESTORE_FILE == *.gz ]]; then
        gunzip -c "$RESTORE_FILE" | mysql \
            --host=$DB_HOST \
            --port=$DB_PORT \
            --user=$DB_USER \
            --password=$DB_PASS \
            $DB_NAME
    else
        mysql \
            --host=$DB_HOST \
            --port=$DB_PORT \
            --user=$DB_USER \
            --password=$DB_PASS \
            $DB_NAME < "$RESTORE_FILE"
    fi
    
    if [ $? -eq 0 ]; then
        print_message $GREEN "✅ تمت الاستعادة بنجاح"
    else
        print_message $RED "❌ فشلت عملية الاستعادة"
        exit 1
    fi
}

# ====================================
# البرنامج الرئيسي
# ====================================

print_message $GREEN "===================================="
print_message $GREEN "  نظام النسخ الاحتياطي لقاعدة البيانات"
print_message $GREEN "===================================="
echo ""

# التحقق من mysqldump
check_mysqldump

# إنشاء مجلد النسخ الاحتياطي
create_backup_dir

# معالجة الأوامر
case "$1" in
    full|"")
        full_backup
        cleanup_old_backups
        ;;
    data)
        data_only_backup
        ;;
    schema)
        schema_only_backup
        ;;
    list)
        list_backups
        ;;
    restore)
        if [ -z "$2" ]; then
            print_message $RED "❌ يجب تحديد ملف النسخة الاحتياطية"
            echo "الاستخدام: $0 restore <backup_file>"
            exit 1
        fi
        restore_backup "$2"
        ;;
    cleanup)
        cleanup_old_backups
        ;;
    *)
        print_message $YELLOW "الاستخدام: $0 [full|data|schema|list|restore|cleanup]"
        echo ""
        echo "الأوامر:"
        echo "  full     - نسخة احتياطية كاملة (افتراضي)"
        echo "  data     - نسخة احتياطية للبيانات فقط"
        echo "  schema   - نسخة احتياطية للمخطط فقط"
        echo "  list     - عرض قائمة النسخ الاحتياطية"
        echo "  restore  - استعادة نسخة احتياطية"
        echo "  cleanup  - حذف النسخ الاحتياطية القديمة"
        exit 1
        ;;
esac

print_message $GREEN ""
print_message $GREEN "✅ تمت العملية بنجاح"
print_message $GREEN "===================================="

