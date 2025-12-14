<template>
  <div class="test-dialog-page">
    <v-container>
      <h1>اختبار النافذة المنبثقة</h1>
      
      <v-btn @click="showDialog = true" color="primary" size="large">
        عرض النافذة المنبثقة
      </v-btn>
      
      <v-dialog v-model="showDialog" max-width="600">
        <v-card>
          <v-card-title>
            <v-icon class="me-2">mdi-test-tube</v-icon>
            اختبار النافذة المنبثقة
          </v-card-title>
          <v-card-text>
            <p>هذه نافذة اختبار للتأكد من عمل النوافذ المنبثقة</p>
            <v-alert type="success" variant="tonal">
              النافذة تعمل بشكل مثالي! ✅
            </v-alert>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn @click="showDialog = false" color="primary">
              إغلاق
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-container>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const showDialog = ref(false)
</script>

<style scoped>
.test-dialog-page {
  padding: 2rem;
  min-height: 100vh;
  background: #f5f5f5;
}
</style>
