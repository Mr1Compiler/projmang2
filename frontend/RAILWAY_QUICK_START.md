# 🚂 النشر السريع على Railway

## خطوات سريعة للنشر على Railway

### 1. ادفع المشروع إلى GitHub:
```bash
git init
git add .
git commit -m "Ready for Railway"
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO.git
git push -u origin main
```

### 2. في Railway Dashboard:

1. **اذهب إلى [railway.app](https://railway.app)**
2. **سجل دخول بحساب GitHub**
3. **اضغط "New Project"**
4. **اختر "Deploy from GitHub repo"**
5. **اختر المستودع الخاص بك**

### 3. إعدادات المشروع:

في صفحة المشروع:

1. **Settings > Root Directory:**
   - اكتب: `dist`

2. **Settings > Build Command:**
   - اكتب: `npm run build`

3. **Settings > Start Command:**
   - اكتب: `npm start` أو اتركه فارغاً

4. **Variables (متغيرات البيئة):**
   - اضغط على "Variables"
   - أضف:
     ```
     VITE_API_URL = https://your-backend-api.com/api
     PORT = 3000
     ```

### 4. النشر:

- Railway سيبني تلقائياً بعد ربط GitHub
- أو اضغط "Deploy" يدوياً

### 5. النطاق:

- بعد النشر، ستحصل على نطاق مثل: `projectmang.up.railway.app`
- يمكنك إضافة نطاق مخصص من Settings > Domains

---

## ⚠️ مهم جداً:

1. **حدّث `VITE_API_URL`** في Variables برابط API الخاص بك
2. **تأكد من أن Backend متاح** ويدعم CORS
3. **تحقق من البناء** في قسم "Deployments"

---

## 🔍 التحقق:

بعد النشر، تحقق من:
- ✅ الموقع يعمل: `https://projectmang.up.railway.app`
- ✅ جميع الصفحات تعمل
- ✅ API يعمل بشكل صحيح

---

**للمزيد من التفاصيل، راجع `DEPLOY_RAILWAY.md`**

