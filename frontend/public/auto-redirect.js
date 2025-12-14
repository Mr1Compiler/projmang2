// إعادة التوجيه التلقائي من المنفذ 3000 إلى 3004
(function() {
    'use strict';
    
    // التحقق من المنفذ الحالي
    const currentPort = window.location.port;
    const currentHost = window.location.hostname;
    
    // إذا كان المنفذ 3000، قم بإعادة التوجيه إلى 3004
    if (currentPort === '3000' || currentHost === 'localhost:3000') {
        const newUrl = window.location.href.replace(':3000', ':3004');
        
        // إظهار رسالة للمستخدم
        if (confirm('تم تحديث عنوان الموقع إلى المنفذ الجديد (3004). هل تريد الانتقال الآن؟')) {
            window.location.href = newUrl;
        } else {
            // إعادة التوجيه التلقائي بعد 3 ثوان
            setTimeout(() => {
                window.location.href = newUrl;
            }, 3000);
        }
    }
    
    // إضافة رسالة في وحدة التحكم
    console.log('✅ نظام إعادة التوجيه التلقائي نشط');
    console.log('المنفذ الحالي:', currentPort);
    console.log('المنفذ المطلوب: 3004');
})();
