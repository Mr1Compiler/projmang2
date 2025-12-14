<?php
/**
 * ملف تكوين قاعدة البيانات
 * Database Configuration File
 * 
 * @package ProjectManagementSystem
 * @author Erticaz Solutions
 * @version 1.0.0
 */

// منع الوصول المباشر
if (!defined('APP_ACCESS')) {
    die('Direct access not permitted');
}

/**
 * إعدادات قاعدة البيانات
 * Database Settings
 */
return [
    // إعدادات الاتصال
    'connections' => [
        'mysql' => [
            'driver' => 'mysql',
            'host' => getenv('DB_HOST') ?: 'localhost',
            'port' => getenv('DB_PORT') ?: '3306',
            'database' => getenv('DB_DATABASE') ?: 'project_management_db',
            'username' => getenv('DB_USERNAME') ?: 'root',
            'password' => getenv('DB_PASSWORD') ?: '',
            'charset' => 'utf8mb4',
            'collation' => 'utf8mb4_unicode_ci',
            'prefix' => '',
            'strict' => true,
            'engine' => 'InnoDB',
        ],
    ],

    // الاتصال الافتراضي
    'default' => getenv('DB_CONNECTION') ?: 'mysql',

    // إعدادات إضافية
    'options' => [
        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
        PDO::ATTR_EMULATE_PREPARES => false,
        PDO::MYSQL_ATTR_INIT_COMMAND => "SET NAMES 'utf8mb4' COLLATE 'utf8mb4_unicode_ci'",
    ],

    // إعدادات التخزين المؤقت
    'cache' => [
        'enabled' => true,
        'ttl' => 3600, // ساعة واحدة
    ],

    // إعدادات السجلات
    'logging' => [
        'enabled' => true,
        'slow_query_time' => 2, // ثانيتين
        'log_file' => __DIR__ . '/../logs/database.log',
    ],
];

