<template>
  <v-container class="pa-6">
    <v-row>
      <v-col cols="12">
        <h1 class="text-h3 mb-6 text-center">اختبار الأيقونات النهائي</h1>
        
        <!-- Test 1: Direct icon tags -->
        <v-card class="mb-6" elevation="2">
          <v-card-title>اختبار 1: أيقونات مباشرة</v-card-title>
          <v-card-text>
            <div class="d-flex flex-wrap ga-4">
              <div class="text-center">
                <i class="mdi mdi-home" style="font-size: 48px; color: #4caf50;"></i>
                <p class="mt-2">mdi-home</p>
              </div>
              <div class="text-center">
                <i class="mdi mdi-trending-up" style="font-size: 48px; color: #f44336;"></i>
                <p class="mt-2">mdi-trending-up</p>
              </div>
              <div class="text-center">
                <i class="mdi mdi-currency-usd" style="font-size: 48px; color: #2196f3;"></i>
                <p class="mt-2">mdi-currency-usd</p>
              </div>
              <div class="text-center">
                <i class="mdi mdi-chart-line" style="font-size: 48px; color: #ff9800;"></i>
                <p class="mt-2">mdi-chart-line</p>
              </div>
            </div>
          </v-card-text>
        </v-card>

        <!-- Test 2: In buttons -->
        <v-card class="mb-6" elevation="2">
          <v-card-title>اختبار 2: أيقونات في الأزرار</v-card-title>
          <v-card-text>
            <div class="d-flex flex-wrap ga-4">
              <v-btn color="primary">
                <i class="mdi mdi-plus me-2"></i>
                إضافة
              </v-btn>
              <v-btn color="success">
                <i class="mdi mdi-check me-2"></i>
                موافق
              </v-btn>
              <v-btn color="error">
                <i class="mdi mdi-delete me-2"></i>
                حذف
              </v-btn>
              <v-btn color="warning">
                <i class="mdi mdi-pencil me-2"></i>
                تعديل
              </v-btn>
            </div>
          </v-card-text>
        </v-card>

        <!-- Test 3: Different sizes -->
        <v-card elevation="2">
          <v-card-title>اختبار 3: أحجام مختلفة</v-card-title>
          <v-card-text>
            <div class="d-flex align-center ga-4">
              <i class="mdi mdi-heart" style="font-size: 16px; color: #e91e63;"></i>
              <i class="mdi mdi-heart" style="font-size: 24px; color: #e91e63;"></i>
              <i class="mdi mdi-heart" style="font-size: 32px; color: #e91e63;"></i>
              <i class="mdi mdi-heart" style="font-size: 48px; color: #e91e63;"></i>
              <i class="mdi mdi-heart" style="font-size: 64px; color: #e91e63;"></i>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
// No script needed for this test
</script>

<style scoped>
/* Component-specific styles if needed */
</style>
