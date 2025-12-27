# 🚀 دليل النشر السريع - Quick Deploy Guide

## النشر على Vercel (الأسهل والأسرع)

### خطوات سريعة:

1. **ادفع المشروع إلى GitHub:**
   ```bash
   git init
   git add .
   git commit -m "Ready for deployment"
   git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO.git
   git push -u origin main
   ```

2. **اذهب إلى [vercel.com](https://vercel.com) و:**
   - سجل دخول بحساب GitHub
   - اضغط "New Project"
   - اختر المستودع
   - في "Environment Variables" أضف:
     ```
     VITE_API_URL = https://your-backend-api.com/api
     ```
   - اضغط "Deploy"

3. **جاهز!** 🎉
   - ستحصل على رابط مثل: `https://your-project.vercel.app`

---

## النشر على Netlify

1. **ادفع المشروع إلى GitHub** (كما أعلاه)

2. **اذهب إلى [netlify.com](https://netlify.com) و:**
   - سجل دخول بحساب GitHub
   - اضغط "New site from Git"
   - اختر المستودع
   - في "Environment variables" أضف:
     ```
     VITE_API_URL = https://your-backend-api.com/api
     ```
   - اضغط "Deploy site"

---

## ⚠️ مهم جداً قبل النشر:

1. **حدّث رابط API في:**
   - `vercel.json` (سطر 18)
   - `netlify.toml` (سطر 5)

2. **تأكد من:**
   - ✅ API الخاص بك متاح على الإنترنت
   - ✅ API يدعم CORS
   - ✅ API يستخدم HTTPS

---

## 🧪 اختبار البناء محلياً:

```bash
npm run build
npm run preview
```

---

## 📝 ملاحظات:

- لا ترفع ملفات `.env` إلى GitHub
- استخدم متغيرات البيئة في منصة النشر
- جميع المسارات تعيد التوجيه تلقائياً إلى `index.html`

---

**للمزيد من التفاصيل، راجع `DEPLOYMENT_GUIDE.md`**

