<template>
  <v-footer
    app
    height="48"
    class="app-footer"
  >
    <div class="footer-content">
      <!-- Company Info -->
      <div class="company-info">
        <span class="company-name">ارتكاز للحلول البرمجية</span>
        <span class="separator">|</span>
        <span class="copyright">
          &copy; {{ currentYear }} - جميع الحقوق محفوظة
        </span>
      </div>

      <!-- Social Links -->
      <div class="social-links">
        <a
          v-for="item in socialItems"
          :key="item.title"
          class="social-link"
          :href="item.href"
          rel="noopener noreferrer"
          target="_blank"
          :title="item.title"
        >
          <v-icon :icon="item.icon" size="18" />
        </a>
      </div>
    </div>
  </v-footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())

const socialItems = [
  {
    title: 'تواصل معنا',
    icon: 'mdi-email',
    href: 'mailto:info@irtikaz.com'
  },
  {
    title: 'واتساب',
    icon: 'mdi-whatsapp',
    href: 'https://wa.me/966500000000'
  },
  {
    title: 'تويتر',
    icon: 'mdi-twitter',
    href: 'https://twitter.com/irtikaz'
  }
]
</script>

<style scoped>
@import './styles/shared.css';

.app-footer {
  background: var(--text-white) !important;
  border-top: 1px solid var(--border-light);
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.05);
}

.footer-content {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 var(--space-6);
}

.company-info {
  display: flex;
  align-items: center;
  gap: var(--space-2);
  color: var(--color-gray-600);
  font-size: var(--font-size-sm);
}

.company-name {
  font-weight: 600;
  color: var(--color-primary);
}

.separator {
  color: var(--color-gray-300);
}

.copyright {
  color: var(--color-gray-500);
}

.social-links {
  display: flex;
  align-items: center;
  gap: var(--space-3);
}

.social-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: var(--radius-full);
  background: var(--color-gray-100);
  color: var(--color-gray-500);
  text-decoration: none;
  transition: all var(--transition-normal);
}

.social-link:hover {
  background: var(--color-primary);
  color: var(--text-white);
  transform: translateY(-2px);
}

/* RTL Support */
.footer-content {
  direction: rtl;
}

/* Responsive */
@media (max-width: 600px) {
  .footer-content {
    flex-direction: column;
    gap: var(--space-2);
    padding: var(--space-2);
  }

  .company-info {
    flex-wrap: wrap;
    justify-content: center;
    text-align: center;
  }

  .separator {
    display: none;
  }
}
</style>
