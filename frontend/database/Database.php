<?php
/**
 * فئة الاتصال بقاعدة البيانات
 * Database Connection Class
 * 
 * @package ProjectManagementSystem
 * @author Erticaz Solutions
 * @version 1.0.0
 */

class Database {
    private static $instance = null;
    private $connection;
    private $config;
    private $queryLog = [];

    /**
     * Constructor - منع الإنشاء المباشر
     */
    private function __construct() {
        $this->loadConfig();
        $this->connect();
    }

    /**
     * منع الاستنساخ
     */
    private function __clone() {}

    /**
     * الحصول على نسخة واحدة من الفئة (Singleton)
     * 
     * @return Database
     */
    public static function getInstance() {
        if (self::$instance === null) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    /**
     * تحميل إعدادات قاعدة البيانات
     */
    private function loadConfig() {
        $configFile = __DIR__ . '/config.php';
        if (!file_exists($configFile)) {
            throw new Exception('Database configuration file not found');
        }
        
        define('APP_ACCESS', true);
        $this->config = require $configFile;
    }

    /**
     * الاتصال بقاعدة البيانات
     */
    private function connect() {
        try {
            $connection = $this->config['connections'][$this->config['default']];
            
            $dsn = sprintf(
                "%s:host=%s;port=%s;dbname=%s;charset=%s",
                $connection['driver'],
                $connection['host'],
                $connection['port'],
                $connection['database'],
                $connection['charset']
            );

            $this->connection = new PDO(
                $dsn,
                $connection['username'],
                $connection['password'],
                $this->config['options']
            );

            $this->log('Database connected successfully');
        } catch (PDOException $e) {
            $this->handleError($e);
        }
    }

    /**
     * الحصول على الاتصال
     * 
     * @return PDO
     */
    public function getConnection() {
        return $this->connection;
    }

    /**
     * تنفيذ استعلام SELECT
     * 
     * @param string $query
     * @param array $params
     * @return array
     */
    public function select($query, $params = []) {
        return $this->execute($query, $params)->fetchAll();
    }

    /**
     * تنفيذ استعلام SELECT وإرجاع صف واحد
     * 
     * @param string $query
     * @param array $params
     * @return array|null
     */
    public function selectOne($query, $params = []) {
        $result = $this->execute($query, $params)->fetch();
        return $result ?: null;
    }

    /**
     * تنفيذ استعلام INSERT
     * 
     * @param string $query
     * @param array $params
     * @return int آخر معرف مدرج
     */
    public function insert($query, $params = []) {
        $this->execute($query, $params);
        return $this->connection->lastInsertId();
    }

    /**
     * تنفيذ استعلام UPDATE
     * 
     * @param string $query
     * @param array $params
     * @return int عدد الصفوف المتأثرة
     */
    public function update($query, $params = []) {
        return $this->execute($query, $params)->rowCount();
    }

    /**
     * تنفيذ استعلام DELETE
     * 
     * @param string $query
     * @param array $params
     * @return int عدد الصفوف المحذوفة
     */
    public function delete($query, $params = []) {
        return $this->execute($query, $params)->rowCount();
    }

    /**
     * تنفيذ استعلام عام
     * 
     * @param string $query
     * @param array $params
     * @return PDOStatement
     */
    public function execute($query, $params = []) {
        try {
            $startTime = microtime(true);
            
            $stmt = $this->connection->prepare($query);
            $stmt->execute($params);
            
            $executionTime = microtime(true) - $startTime;
            
            // تسجيل الاستعلام
            $this->logQuery($query, $params, $executionTime);
            
            // تسجيل الاستعلامات البطيئة
            if ($executionTime > $this->config['logging']['slow_query_time']) {
                $this->log("Slow query detected: {$query} - Time: {$executionTime}s", 'warning');
            }
            
            return $stmt;
        } catch (PDOException $e) {
            $this->handleError($e, $query, $params);
        }
    }

    /**
     * بدء معاملة
     */
    public function beginTransaction() {
        return $this->connection->beginTransaction();
    }

    /**
     * تنفيذ المعاملة
     */
    public function commit() {
        return $this->connection->commit();
    }

    /**
     * التراجع عن المعاملة
     */
    public function rollback() {
        return $this->connection->rollBack();
    }

    /**
     * التحقق من وجود جدول
     * 
     * @param string $tableName
     * @return bool
     */
    public function tableExists($tableName) {
        $query = "SHOW TABLES LIKE :table";
        $result = $this->selectOne($query, ['table' => $tableName]);
        return $result !== null;
    }

    /**
     * الحصول على جميع الجداول
     * 
     * @return array
     */
    public function getAllTables() {
        $query = "SHOW TABLES";
        return $this->select($query);
    }

    /**
     * تسجيل الاستعلام
     * 
     * @param string $query
     * @param array $params
     * @param float $executionTime
     */
    private function logQuery($query, $params, $executionTime) {
        if ($this->config['logging']['enabled']) {
            $this->queryLog[] = [
                'query' => $query,
                'params' => $params,
                'execution_time' => $executionTime,
                'timestamp' => date('Y-m-d H:i:s'),
            ];
        }
    }

    /**
     * الحصول على سجل الاستعلامات
     * 
     * @return array
     */
    public function getQueryLog() {
        return $this->queryLog;
    }

    /**
     * تسجيل رسالة
     * 
     * @param string $message
     * @param string $level
     */
    private function log($message, $level = 'info') {
        if (!$this->config['logging']['enabled']) {
            return;
        }

        $logFile = $this->config['logging']['log_file'];
        $logDir = dirname($logFile);
        
        if (!is_dir($logDir)) {
            mkdir($logDir, 0755, true);
        }

        $timestamp = date('Y-m-d H:i:s');
        $logMessage = "[{$timestamp}] [{$level}] {$message}\n";
        
        file_put_contents($logFile, $logMessage, FILE_APPEND);
    }

    /**
     * معالجة الأخطاء
     * 
     * @param PDOException $e
     * @param string $query
     * @param array $params
     */
    private function handleError($e, $query = '', $params = []) {
        $errorMessage = sprintf(
            "Database Error: %s\nQuery: %s\nParams: %s",
            $e->getMessage(),
            $query,
            json_encode($params)
        );
        
        $this->log($errorMessage, 'error');
        
        // في بيئة الإنتاج، لا تعرض تفاصيل الخطأ
        if (getenv('APP_ENV') === 'production') {
            throw new Exception('Database error occurred. Please contact support.');
        } else {
            throw new Exception($errorMessage);
        }
    }

    /**
     * إغلاق الاتصال
     */
    public function disconnect() {
        $this->connection = null;
        $this->log('Database disconnected');
    }

    /**
     * Destructor
     */
    public function __destruct() {
        $this->disconnect();
    }
}

