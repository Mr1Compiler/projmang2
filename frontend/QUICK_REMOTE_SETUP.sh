#!/bin/bash

# سكريبت سريع لإضافة remote جديد
# استخدم: bash QUICK_REMOTE_SETUP.sh

echo "🔗 إضافة Remote جديد - publicProj"
echo ""

# التحقق من وجود remote
if git remote | grep -q "publicProj"; then
    echo "⚠️  Remote 'publicProj' موجود بالفعل!"
    echo "لحذفه: git remote remove publicProj"
    exit 1
fi

# طلب رابط المستودع
echo "أدخل رابط المستودع (GitHub, GitLab, etc):"
read REPO_URL

if [ -z "$REPO_URL" ]; then
    echo "❌ لم يتم إدخال رابط!"
    exit 1
fi

# إضافة remote
echo ""
echo "➕ إضافة remote..."
git remote add publicProj "$REPO_URL"

if [ $? -eq 0 ]; then
    echo "✅ تم إضافة remote بنجاح!"
    echo ""
    echo "📋 Remotes الحالية:"
    git remote -v
    echo ""
    echo "🚀 لرفع التغييرات:"
    echo "   git push publicProj main"
else
    echo "❌ فشل إضافة remote!"
    exit 1
fi

