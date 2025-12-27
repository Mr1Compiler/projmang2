# 🚂 دليل النشر على Railway

## النشر على Railway

Railway هو منصة سحابية سهلة للنشر. هذا الدليل يوضح كيفية رفع المشروع على Railway.

## 📋 المتطلبات

- حساب على [Railway](https://railway.app)
- المشروع موجود على GitHub (موصى به)

## 🚀 خطوات النشر

### الطريقة 1: من خلال GitHub (موصى بها)

1. **ادفع المشروع إلى GitHub:**
   ```bash
   git init
   git add .
   git commit -m "Ready for Railway deployment"
   git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO.git
   git push -u origin main
   ```

2. **سجل دخول إلى Railway:**
   - اذهب إلى [railway.app](https://railway.app)
   - سجل دخول بحساب GitHub

3. **إنشاء مشروع جديد:**
   - اضغط على "New Project"
   - اختر "Deploy from GitHub repo"
   - اختر المستودع الخاص بك

4. **إعداد متغيرات البيئة:**
   - في صفحة المشروع، اذهب إلى "Variables"
   - أضف المتغيرات التالية:
     ```
     VITE_API_URL=https://your-backend-api.com/api
     NODE_ENV=production
     ```

5. **إعدادات البناء:**
   - Railway سيكتشف تلقائياً أنه مشروع Vite
   - إذا لم يكتشف، استخدم:
     - **Build Command:** `npm run build`
     - **Start Command:** `npm run preview` (للاختبار) أو استخدم خدمة static hosting

6. **إعداد Static Files:**
   - بعد البناء، اضغط على "Settings"
   - في "Root Directory" اكتب: `dist`
   - في "Build Command" اكتب: `npm run build`
   - في "Start Command" اتركه فارغاً أو استخدم: `npx serve dist -p $PORT`

### الطريقة 2: استخدام Railway CLI

```bash
# تثبيت Railway CLI
npm i -g @railway/cli

# تسجيل الدخول
railway login

# تهيئة المشروع
railway init

# ربط المشروع بمشروع Railway موجود
railway link

# إضافة متغيرات البيئة
railway variables set VITE_API_URL=https://your-backend-api.com/api

# النشر
railway up
```

## ⚙️ إعدادات متقدمة

### استخدام Nginx للـ Static Files (موصى به)

1. **إنشاء ملف `nginx.conf`:**
   ```nginx
   server {
     listen $PORT;
     server_name _;
     root /app/dist;
     index index.html;

     location / {
       try_files $uri $uri/ /index.html;
     }

     location /api {
       proxy_pass https://your-backend-api.com;
       proxy_set_header Host $host;
       proxy_set_header X-Real-IP $remote_addr;
     }
   }
   ```

2. **تحديث `package.json` لإضافة serve:**
   ```json
   {
     "scripts": {
       "start": "npx serve dist -p $PORT -s"
     }
   }
   ```

### استخدام Vercel/Netlify للـ Frontend + Railway للـ Backend

إذا كان لديك Backend على Railway و Frontend منفصل:

1. **نشر Frontend على Vercel/Netlify:**
   - اتبع دليل النشر على Vercel/Netlify
   - استخدم رابط Railway API في متغيرات البيئة

2. **إعداد CORS في Backend:**
   - تأكد من أن Backend يدعم CORS من نطاق Frontend

## 🔧 تحديث المشروع الموجود

إذا كان لديك مشروع موجود على Railway:

1. **ربط المشروع المحلي:**
   ```bash
   railway link
   ```

2. **دفع التغييرات:**
   ```bash
   git add .
   git commit -m "Update for Railway"
   git push
   ```

3. **Railway سيبني تلقائياً بعد push**

## 📝 متغيرات البيئة المطلوبة

في Railway Dashboard > Variables:

```
VITE_API_URL=https://your-backend-api.com/api
NODE_ENV=production
PORT=3000
```

## 🌐 إعدادات النطاق (Domain)

1. في Railway Dashboard، اذهب إلى "Settings"
2. في قسم "Domains"، يمكنك:
   - استخدام النطاق الافتراضي: `projectmang.up.railway.app`
   - إضافة نطاق مخصص

## ✅ التحقق من النشر

بعد النشر:

1. ✅ تحقق من أن الموقع يعمل: `https://projectmang.up.railway.app`
2. ✅ تحقق من أن API يعمل
3. ✅ تحقق من جميع الصفحات
4. ✅ تحقق من التصميم المتجاوب

## 🐛 حل المشاكل

### المشكلة: البناء فشل
- تحقق من أن `package.json` يحتوي على `build` script
- تحقق من logs في Railway Dashboard

### المشكلة: الموقع لا يعمل
- تحقق من `Root Directory` في Settings (يجب أن يكون `dist`)
- تحقق من `Start Command`

### المشكلة: API لا يعمل
- تحقق من `VITE_API_URL` في Variables
- تحقق من CORS في Backend

## 📞 الدعم

- [Railway Documentation](https://docs.railway.app)
- [Railway Discord](https://discord.gg/railway)

---

**ملاحظة:** تأكد من تحديث `VITE_API_URL` في متغيرات البيئة قبل النشر!

