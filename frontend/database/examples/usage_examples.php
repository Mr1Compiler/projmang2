<?php
/**
 * أمثلة استخدام قاعدة البيانات
 * Database Usage Examples
 * 
 * @package ProjectManagementSystem
 * @author Erticaz Solutions
 * @version 1.0.0
 */

// تضمين ملف الفئة
require_once __DIR__ . '/../Database.php';

// الحصول على نسخة من قاعدة البيانات
$db = Database::getInstance();

/**
 * مثال 1: استعلامات SELECT البسيطة
 * Example 1: Simple SELECT queries
 */
echo "=== مثال 1: استعلامات SELECT ===\n";

// جلب جميع المشاريع
$projects = $db->select("SELECT * FROM projects WHERE status = :status", ['status' => 'in_progress']);
echo "عدد المشاريع النشطة: " . count($projects) . "\n\n";

// جلب مشروع واحد
$project = $db->selectOne("SELECT * FROM projects WHERE id = :id", ['id' => 1]);
echo "اسم المشروع: " . ($project['name'] ?? 'غير موجود') . "\n\n";

/**
 * مثال 2: استعلامات INSERT
 * Example 2: INSERT queries
 */
echo "=== مثال 2: إضافة بيانات ===\n";

try {
    $newProjectId = $db->insert(
        "INSERT INTO projects (project_number, name, description, status, created_by) 
         VALUES (:project_number, :name, :description, :status, :created_by)",
        [
            'project_number' => 'PRJ-' . date('Y') . '-' . rand(1000, 9999),
            'name' => 'مشروع تجريبي',
            'description' => 'وصف المشروع التجريبي',
            'status' => 'pending',
            'created_by' => 1
        ]
    );
    echo "تم إضافة المشروع بنجاح. المعرف: {$newProjectId}\n\n";
} catch (Exception $e) {
    echo "خطأ في إضافة المشروع: " . $e->getMessage() . "\n\n";
}

/**
 * مثال 3: استعلامات UPDATE
 * Example 3: UPDATE queries
 */
echo "=== مثال 3: تحديث بيانات ===\n";

try {
    $affectedRows = $db->update(
        "UPDATE projects SET progress_percentage = :progress WHERE id = :id",
        [
            'progress' => 50.00,
            'id' => 1
        ]
    );
    echo "تم تحديث {$affectedRows} سجل\n\n";
} catch (Exception $e) {
    echo "خطأ في التحديث: " . $e->getMessage() . "\n\n";
}

/**
 * مثال 4: استعلامات DELETE
 * Example 4: DELETE queries
 */
echo "=== مثال 4: حذف بيانات ===\n";

try {
    $deletedRows = $db->delete(
        "DELETE FROM projects WHERE status = :status AND created_at < :date",
        [
            'status' => 'cancelled',
            'date' => date('Y-m-d', strtotime('-1 year'))
        ]
    );
    echo "تم حذف {$deletedRows} سجل\n\n";
} catch (Exception $e) {
    echo "خطأ في الحذف: " . $e->getMessage() . "\n\n";
}

/**
 * مثال 5: المعاملات (Transactions)
 * Example 5: Transactions
 */
echo "=== مثال 5: المعاملات ===\n";

try {
    $db->beginTransaction();
    
    // إضافة مصروف
    $expenseId = $db->insert(
        "INSERT INTO expenses (expense_number, description, amount, expense_date, status, created_by) 
         VALUES (:expense_number, :description, :amount, :expense_date, :status, :created_by)",
        [
            'expense_number' => 'EXP-' . date('Y') . '-' . rand(1000, 9999),
            'description' => 'مصروف تجريبي',
            'amount' => 100000,
            'expense_date' => date('Y-m-d'),
            'status' => 'pending',
            'created_by' => 1
        ]
    );
    
    // ربط المصروف بالمشروع
    $db->insert(
        "INSERT INTO project_expenses (project_id, expense_id, allocated_amount) 
         VALUES (:project_id, :expense_id, :allocated_amount)",
        [
            'project_id' => 1,
            'expense_id' => $expenseId,
            'allocated_amount' => 100000
        ]
    );
    
    $db->commit();
    echo "تمت المعاملة بنجاح\n\n";
} catch (Exception $e) {
    $db->rollback();
    echo "فشلت المعاملة: " . $e->getMessage() . "\n\n";
}

/**
 * مثال 6: استخدام Views
 * Example 6: Using Views
 */
echo "=== مثال 6: استخدام Views ===\n";

$projectSummary = $db->select("SELECT * FROM view_project_summary LIMIT 5");
echo "ملخص المشاريع:\n";
foreach ($projectSummary as $project) {
    echo "- {$project['name']}: {$project['progress_percentage']}% مكتمل\n";
}
echo "\n";

/**
 * مثال 7: استعلامات معقدة مع JOIN
 * Example 7: Complex queries with JOIN
 */
echo "=== مثال 7: استعلامات معقدة ===\n";

$query = "
    SELECT 
        p.name as project_name,
        COUNT(t.id) as task_count,
        AVG(t.progress_percentage) as avg_progress
    FROM projects p
    LEFT JOIN tasks t ON p.id = t.project_id
    WHERE p.status = :status
    GROUP BY p.id, p.name
    ORDER BY task_count DESC
";

$projectStats = $db->select($query, ['status' => 'in_progress']);
echo "إحصائيات المشاريع النشطة:\n";
foreach ($projectStats as $stat) {
    echo "- {$stat['project_name']}: {$stat['task_count']} مهمة، متوسط التقدم: {$stat['avg_progress']}%\n";
}
echo "\n";

/**
 * مثال 8: التحقق من وجود جدول
 * Example 8: Check table existence
 */
echo "=== مثال 8: التحقق من الجداول ===\n";

if ($db->tableExists('projects')) {
    echo "جدول projects موجود\n";
} else {
    echo "جدول projects غير موجود\n";
}
echo "\n";

/**
 * مثال 9: الحصول على جميع الجداول
 * Example 9: Get all tables
 */
echo "=== مثال 9: جميع الجداول ===\n";

$tables = $db->getAllTables();
echo "عدد الجداول في قاعدة البيانات: " . count($tables) . "\n";
foreach (array_slice($tables, 0, 5) as $table) {
    echo "- " . current($table) . "\n";
}
echo "\n";

/**
 * مثال 10: الحصول على سجل الاستعلامات
 * Example 10: Get query log
 */
echo "=== مثال 10: سجل الاستعلامات ===\n";

$queryLog = $db->getQueryLog();
echo "عدد الاستعلامات المنفذة: " . count($queryLog) . "\n";
echo "إجمالي وقت التنفيذ: " . array_sum(array_column($queryLog, 'execution_time')) . " ثانية\n";
echo "\n";

/**
 * مثال 11: استعلامات التقارير المالية
 * Example 11: Financial reports queries
 */
echo "=== مثال 11: التقارير المالية ===\n";

$financialSummary = $db->selectOne("
    SELECT 
        (SELECT SUM(amount) FROM expenses WHERE status = 'paid') as total_expenses,
        (SELECT SUM(amount) FROM income WHERE status = 'confirmed') as total_income,
        (SELECT SUM(amount) FROM income WHERE status = 'confirmed') - 
        (SELECT SUM(amount) FROM expenses WHERE status = 'paid') as net_profit
");

if ($financialSummary) {
    echo "إجمالي المصاريف: " . number_format($financialSummary['total_expenses']) . " دينار\n";
    echo "إجمالي الإيرادات: " . number_format($financialSummary['total_income']) . " دينار\n";
    echo "صافي الربح: " . number_format($financialSummary['net_profit']) . " دينار\n";
}
echo "\n";

/**
 * مثال 12: استعلامات تقارير الفريق
 * Example 12: Team reports queries
 */
echo "=== مثال 12: تقارير الفريق ===\n";

$teamPerformance = $db->select("
    SELECT 
        tm.name,
        COUNT(t.id) as total_tasks,
        SUM(CASE WHEN t.status = 'completed' THEN 1 ELSE 0 END) as completed_tasks,
        AVG(t.progress_percentage) as avg_progress
    FROM team_members tm
    LEFT JOIN tasks t ON tm.id = t.assigned_to
    WHERE tm.status = 'active'
    GROUP BY tm.id, tm.name
    ORDER BY completed_tasks DESC
    LIMIT 5
");

echo "أداء الفريق:\n";
foreach ($teamPerformance as $member) {
    $completionRate = $member['total_tasks'] > 0 
        ? round(($member['completed_tasks'] / $member['total_tasks']) * 100, 2) 
        : 0;
    echo "- {$member['name']}: {$member['completed_tasks']}/{$member['total_tasks']} مهمة مكتملة ({$completionRate}%)\n";
}
echo "\n";

/**
 * مثال 13: استعلامات أيام العمل
 * Example 13: Work days queries
 */
echo "=== مثال 13: تقارير أيام العمل ===\n";

$workDaySummary = $db->select("
    SELECT 
        p.name as project_name,
        COUNT(wd.id) as work_days_count,
        SUM(wd.total_cost) as total_cost
    FROM work_days wd
    JOIN projects p ON wd.project_id = p.id
    WHERE wd.status = 'completed'
    GROUP BY p.id, p.name
    ORDER BY total_cost DESC
");

echo "ملخص أيام العمل حسب المشروع:\n";
foreach ($workDaySummary as $summary) {
    echo "- {$summary['project_name']}: {$summary['work_days_count']} يوم عمل، التكلفة: " 
         . number_format($summary['total_cost']) . " دينار\n";
}
echo "\n";

echo "=== انتهت الأمثلة ===\n";

