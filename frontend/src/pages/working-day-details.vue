<template>
  <div class="working-day-details-page">
    <v-container>
      <v-row justify="center">
        <v-col cols="12" md="8">
          <v-card class="pa-6">
            <v-card-title class="text-h4 text-center mb-6">
              صفحة فارغة
            </v-card-title>
            
            <v-card-text class="text-center">
              <p>هذه صفحة فارغة كما طلبت</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const workingDayId = ref(null)
const projectId = ref(null)

onMounted(() => {
  // Get data from URL parameters
  workingDayId.value = route.query.id
  projectId.value = route.query.projectId
  
  console.log('=== WORKING DAY DETAILS PAGE LOADED ===')
  console.log('Working day ID:', workingDayId.value)
  console.log('Project ID:', projectId.value)
  console.log('All query params:', route.query)
})
</script>

<style scoped>
.working-day-details-page {
  min-height: 100vh;
  background: #f5f5f5;
  padding: 20px 0;
}
</style>