<template>
  <div class="login-page">
    <div class="login-container">
      <!-- Login Card with Header -->
      <div class="login-box">
        <!-- Header -->
        <div class="login-header">
          <div class="header-gradient-line"></div>
          <div class="header-content">
            <v-icon size="40" color="white">mdi-shield-lock</v-icon>
            <div class="header-text">
              <h1 class="main-title">نظام إدارة المشاريع</h1>
              <p class="subtitle">تسجيل الدخول</p>
            </div>
          </div>
        </div>

        <!-- Login Form -->
        <div class="login-form-section">
          <form @submit.prevent="handleLogin">
            <!-- Error Message -->
            <v-alert
              v-if="errorMessage"
              type="error"
              variant="tonal"
              class="mb-4"
              density="compact"
            >
              {{ errorMessage }}
            </v-alert>

            <!-- Username Field -->
            <v-text-field
              v-model="formData.username"
              label="اسم المستخدم"
              prepend-inner-icon="mdi-account"
              variant="outlined"
              density="comfortable"
              class="mb-4"
              required
              :disabled="loading"
              hide-details="auto"
            ></v-text-field>

            <!-- Password Field -->
            <v-text-field
              v-model="formData.password"
              :type="showPassword ? 'text' : 'password'"
              label="كلمة المرور"
              prepend-inner-icon="mdi-lock"
              :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
              @click:append-inner="showPassword = !showPassword"
              variant="outlined"
              density="comfortable"
              class="mb-6"
              required
              :disabled="loading"
              hide-details="auto"
            ></v-text-field>

            <!-- Login Button -->
            <v-btn
              type="submit"
              size="large"
              block
              :loading="loading"
              :disabled="loading"
              class="login-btn"
            >
              <v-icon start>mdi-login</v-icon>
              دخول
            </v-btn>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { handleLogin as loginUser } from '@/services/authService'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const formData = ref({
  username: '',
  password: ''
})

const loading = ref(false)
const errorMessage = ref('')
const showPassword = ref(false)

const handleLogin = async () => {
  // التحقق من الحقول
  if (!formData.value.username || !formData.value.password) {
    errorMessage.value = 'يرجى إدخال اسم المستخدم وكلمة المرور'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    const session = await loginUser({
      username: formData.value.username.trim(),
      password: formData.value.password,
    })

    // Set user in store
    authStore.setUser({ id: session.userId, username: session.username })

    // Fetch user pages and permissions
    await authStore.fetchPages()

    await router.push('/')
  } catch (error) {
    const status = error?.status
    const message = error?.message || ''
    if (status === 0) {
      errorMessage.value = 'فشل الاتصال بالخادم. يرجى التأكد من تشغيل الخادم والمحاولة مجدداً.'
    } else if (status === 401) {
      if (message.includes('inactive')) {
        errorMessage.value = 'حسابك غير مفعل. يرجى التواصل مع المسؤول.'
      } else {
        errorMessage.value = 'اسم المستخدم أو كلمة المرور غير صحيحة.'
      }
    } else {
      errorMessage.value = error?.message || 'حدث خطأ غير متوقع. يرجى المحاولة مرة أخرى.'
    }
    console.error('Login error:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #eef2ff 0%, #e0e7ff 50%, #c7d2fe 100%);
  padding: 20px;
}

.login-container {
  width: 100%;
  max-width: 420px;
}

.login-box {
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 10px 25px -5px rgba(67, 56, 202, 0.2), 0 8px 10px -6px rgba(67, 56, 202, 0.1);
}

/* Header - Indigo gradient like sidebar */
.login-header {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 100%);
  position: relative;
  overflow: hidden;
}

.header-gradient-line {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #a5b4fc 0%, #c7d2fe 50%, #e0e7ff 100%);
}

.header-content {
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.main-title {
  color: white;
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0;
  letter-spacing: 0.3px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.subtitle {
  color: rgba(255, 255, 255, 0.9);
  font-size: 0.875rem;
  margin: 0;
}

/* Form Section */
.login-form-section {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  padding: 24px;
}

/* Text field styling */
.login-form-section :deep(.v-field__input) {
  color: #1e293b !important;
}

.login-form-section :deep(.v-label) {
  color: #64748b !important;
}

.login-form-section :deep(.v-field--focused .v-label) {
  color: #4338ca !important;
}

.login-form-section :deep(.v-field__outline) {
  color: #c7d2fe !important;
}

.login-form-section :deep(.v-field--focused .v-field__outline) {
  color: #4338ca !important;
}

.login-form-section :deep(.v-icon) {
  color: #64748b !important;
}

.login-form-section :deep(.v-field--focused .v-icon) {
  color: #4338ca !important;
}

/* Login Button - Indigo gradient */
.login-btn {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 100%) !important;
  color: white !important;
  font-weight: 600 !important;
  letter-spacing: 0.5px;
  box-shadow: 0 4px 6px -1px rgba(67, 56, 202, 0.3) !important;
}

.login-btn:hover {
  background: linear-gradient(135deg, #3730a3 0%, #4338ca 100%) !important;
  box-shadow: 0 6px 12px rgba(67, 56, 202, 0.4) !important;
}

.login-btn :deep(.v-icon) {
  color: white !important;
}

/* Responsive */
@media (max-width: 480px) {
  .login-container {
    max-width: 100%;
  }

  .header-content {
    padding: 20px;
  }

  .main-title {
    font-size: 1.1rem;
  }

  .login-form-section {
    padding: 20px;
  }
}
</style>
