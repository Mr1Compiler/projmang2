# نظام التنسيق الشامل - نظام إدارة المشاريع
## Comprehensive Styling System - Project Management System

### 📋 نظرة عامة
نظام CSS شامل ومنظم تم تصميمه خصيصاً لدعم اللغة العربية والاتجاه من اليمين إلى اليسار (RTL) في نظام إدارة المشاريع.

### 🗂️ هيكل الملفات

```
src/styles/
├── main.css          # الملف الرئيسي - يستورد جميع الملفات
├── colors.css        # نظام الألوان الشامل
├── typography.css    # نظام الخطوط والتنسيق النصي
├── spacing.css       # نظام المسافات والطبقات
├── settings.scss     # إعدادات Vuetify
└── README.md         # هذا الملف
```

### 🎨 نظام الألوان (colors.css)

#### الألوان الأساسية
- **Primary**: `#1976d2` - اللون الأساسي للنظام
- **Secondary**: `#424242` - اللون الثانوي
- **Success**: `#4caf50` - لون النجاح
- **Warning**: `#ff9800` - لون التحذير
- **Error**: `#f44336` - لون الخطأ
- **Info**: `#2196f3` - لون المعلومات

#### ألوان النص
- **Primary**: `#212121` - النص الأساسي
- **Secondary**: `#757575` - النص الثانوي
- **Tertiary**: `#9e9e9e` - النص المساعد
- **White**: `#ffffff` - النص الأبيض

#### التدرجات
```css
--gradient-primary: linear-gradient(135deg, #667eea 0%, #764ba2 100%)
--gradient-success: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)
```

### ✍️ نظام الخطوط (typography.css)

#### الخطوط المستخدمة
- **Primary**: Cairo - الخط الأساسي للنصوص
- **Display**: Cairo - خط العناوين
- **Secondary**: Tajawal - الخط الثانوي

#### أحجام الخطوط
```css
--font-size-xs: 0.75rem    /* 12px */
--font-size-sm: 0.875rem   /* 14px */
--font-size-base: 1rem     /* 16px */
--font-size-lg: 1.125rem   /* 18px */
--font-size-xl: 1.25rem    /* 20px */
--font-size-2xl: 1.5rem    /* 24px */
--font-size-3xl: 1.875rem  /* 30px */
--font-size-4xl: 2.25rem   /* 36px */
```

#### أوزان الخطوط
```css
--font-weight-light: 300
--font-weight-normal: 400
--font-weight-medium: 500
--font-weight-semibold: 600
--font-weight-bold: 700
--font-weight-extrabold: 800
--font-weight-black: 900
```

### 📏 نظام المسافات (spacing.css)

#### المسافات الأساسية
```css
--space-0: 0
--space-1: 0.25rem   /* 4px */
--space-2: 0.5rem    /* 8px */
--space-3: 0.75rem   /* 12px */
--space-4: 1rem      /* 16px */
--space-5: 1.25rem   /* 20px */
--space-6: 1.5rem    /* 24px */
--space-8: 2rem      /* 32px */
--space-12: 3rem     /* 48px */
--space-16: 4rem     /* 64px */
```

#### الحدود الدائرية
```css
--radius-sm: 0.125rem   /* 2px */
--radius-base: 0.25rem  /* 4px */
--radius-md: 0.375rem   /* 6px */
--radius-lg: 0.5rem     /* 8px */
--radius-xl: 0.75rem    /* 12px */
--radius-2xl: 1rem      /* 16px */
--radius-full: 9999px
```

#### الظلال
```css
--shadow-sm: 0 1px 3px 0 rgba(0, 0, 0, 0.1)
--shadow-base: 0 4px 6px -1px rgba(0, 0, 0, 0.1)
--shadow-md: 0 10px 15px -3px rgba(0, 0, 0, 0.1)
--shadow-lg: 0 20px 25px -5px rgba(0, 0, 0, 0.1)
--shadow-xl: 0 25px 50px -12px rgba(0, 0, 0, 0.25)
```

### 🎯 الفئات المساعدة

#### فئات النص
```css
.text-primary      /* النص الأساسي */
.text-secondary    /* النص الثانوي */
.text-white        /* النص الأبيض */
.text-success      /* النص الأخضر */
.text-warning      /* النص البرتقالي */
.text-error        /* النص الأحمر */
```

#### فئات الأوزان
```css
.font-weight-light
.font-weight-normal
.font-weight-medium
.font-weight-semibold
.font-weight-bold
.font-weight-extrabold
```

#### فئات الأحجام
```css
.text-xs, .text-sm, .text-base, .text-lg, .text-xl
.text-2xl, .text-3xl, .text-4xl, .text-5xl, .text-6xl
```

#### فئات المسافات
```css
.p-0, .p-1, .p-2, .p-3, .p-4, .p-5, .p-6, .p-8
.px-0, .px-1, .px-2, .px-3, .px-4, .px-5, .px-6
.py-0, .py-1, .py-2, .py-3, .py-4, .py-5, .py-6
.m-0, .m-1, .m-2, .m-3, .m-4, .m-5, .m-6, .m-8
.mx-0, .mx-1, .mx-2, .mx-3, .mx-4, .mx-5, .mx-6
.my-0, .my-1, .my-2, .my-3, .my-4, .my-5, .my-6
```

### 🎨 الفئات المخصصة

#### البطاقات المحسنة
```css
.enhanced-card     /* بطاقة محسنة مع تأثيرات */
.stat-card         /* بطاقة إحصائية مع تدرج علوي */
```

#### الأزرار المحسنة
```css
.enhanced-btn      /* زر محسن مع تأثيرات */
.btn-primary       /* زر أساسي */
.btn-secondary     /* زر ثانوي */
.btn-success       /* زر نجاح */
.btn-warning       /* زر تحذير */
.btn-error         /* زر خطأ */
.btn-outlined      /* زر محاط */
```

#### النصوص المحسنة
```css
.arabic-text       /* نص عربي محسن */
.enhanced-title    /* عنوان محسن مع تدرج */
.enhanced-logo     /* شعار محسن */
```

### 📱 الاستجابة

#### نقاط الكسر
```css
/* شاشات صغيرة */
@media (max-width: 600px) { ... }

/* شاشات متوسطة */
@media (min-width: 601px) and (max-width: 960px) { ... }

/* شاشات كبيرة */
@media (min-width: 961px) { ... }
```

#### فئات الاستجابة
```css
.sm\:block         /* عرض كتلة في الشاشات الصغيرة */
.md\:flex          /* عرض مرن في الشاشات المتوسطة */
.lg\:hidden        /* إخفاء في الشاشات الكبيرة */
```

### 🔧 التخصيص

#### تخصيص الألوان
```css
:root {
  --color-primary: #your-color;
  --color-secondary: #your-color;
}
```

#### تخصيص الخطوط
```css
:root {
  --font-primary: 'Your-Font', sans-serif;
  --font-display: 'Your-Display-Font', sans-serif;
}
```

#### تخصيص المسافات
```css
:root {
  --space-4: 1.5rem;  /* تخصيص المسافة 4 */
  --radius-lg: 1rem;  /* تخصيص الحدود الدائرية */
}
```

### 🚀 الاستخدام

#### في المكونات Vue
```vue
<template>
  <div class="enhanced-card p-6">
    <h2 class="enhanced-title text-2xl mb-4">عنوان المحتوى</h2>
    <p class="arabic-text text-base text-secondary leading-relaxed">
      نص المحتوى باللغة العربية
    </p>
    <v-btn class="enhanced-btn btn-primary mt-4">
      زر العمل
    </v-btn>
  </div>
</template>
```

#### في CSS المخصص
```css
.my-custom-component {
  background: var(--gradient-primary);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-lg);
  padding: var(--space-6);
}
```

### 📚 المراجع

- [Vuetify CSS Variables](https://vuetifyjs.com/features/sass-variables/)
- [CSS Custom Properties](https://developer.mozilla.org/en-US/docs/Web/CSS/--*)
- [Arabic Typography Guidelines](https://www.arabictypography.com/)
- [RTL CSS Best Practices](https://rtlstyling.com/)

### 🤝 المساهمة

عند إضافة فئات أو متغيرات جديدة، يرجى:
1. تحديث هذا الملف
2. اتباع نظام التسمية الموجود
3. إضافة تعليقات توضيحية
4. اختبار الاستجابة

### 📝 التحديثات

- **v1.0.0** - إطلاق النظام الأساسي
- **v1.1.0** - إضافة نظام الألوان الشامل
- **v1.2.0** - إضافة نظام الخطوط المتقدم
- **v1.3.0** - إضافة نظام المسافات والطبقات
- **v1.4.0** - تحسينات الاستجابة والأداء