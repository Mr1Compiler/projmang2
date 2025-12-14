<template>
  <div class="space-y-6">
    <!-- نوع المادة -->
    <div>
      <label for="material-type" class="block text-sm font-medium text-gray-700 mb-2">
        نوع المادة
      </label>
      <input
        id="material-type"
        v-model="formData.type"
        type="text"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
        placeholder="أدخل نوع المادة"
        required
      />
    </div>

    <!-- الصف الثاني: الكمية والسعر -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="quantity" class="block text-sm font-medium text-gray-700 mb-2">
          الكمية
        </label>
        <input
          id="quantity"
          v-model="formData.quantity"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
      <div>
        <label for="price" class="block text-sm font-medium text-gray-700 mb-2">
          IQD السعر
        </label>
        <input
          id="price"
          v-model="formData.price"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
    </div>

    <!-- الصف الثالث: السعر الكلي واسم السائق -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="total-price" class="block text-sm font-medium text-gray-700 mb-2">
          IQD السعر الكلي
        </label>
        <input
          id="total-price"
          v-model="formData.totalPrice"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
      <div>
        <label for="driver-name" class="block text-sm font-medium text-gray-700 mb-2">
          اسم السائق
        </label>
        <input
          id="driver-name"
          v-model="formData.driver"
          type="text"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="أدخل اسم السائق"
          required
        />
      </div>
    </div>

    <!-- التفاصيل -->
    <div>
      <label for="details" class="block text-sm font-medium text-gray-700 mb-2">
        التفاصيل
      </label>
      <textarea
        id="details"
        v-model="formData.details"
        rows="3"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
        placeholder="أدخل التفاصيل الإضافية"
        required
      />
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  initialData: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const formData = ref({
  type: '',
  quantity: '',
  price: '',
  totalPrice: '',
  driver: '',
  details: '',
  ...props.initialData
})

// مراقبة التغييرات وإرسالها للوالد
watch(formData, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// دالة لإعادة تعيين النموذج
const resetForm = () => {
  formData.value = {
    type: '',
    quantity: '',
    price: '',
    totalPrice: '',
    driver: '',
    details: ''
  }
}

// دالة للتحقق من صحة البيانات
const isValid = () => {
  return formData.value.type && 
         formData.value.quantity && 
         formData.value.price && 
         formData.value.totalPrice && 
         formData.value.driver && 
         formData.value.details
}

defineExpose({
  resetForm,
  isValid,
  formData
})
</script>

<style scoped>
/* تحسينات إضافية للتصميم */
input:focus,
textarea:focus {
  outline: none;
  border-color: #1976d2;
  box-shadow: 0 0 0 3px rgba(25, 118, 210, 0.1);
}

label {
  color: #374151;
  font-weight: 500;
}

input,
textarea {
  transition: all 0.2s ease-in-out;
}

input:hover,
textarea:hover {
  border-color: #9ca3af;
}
</style>
