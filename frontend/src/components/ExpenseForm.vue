<template>
  <div class="space-y-6">
    <!-- نوع المصروف -->
    <div>
      <label for="expense-type" class="block text-sm font-medium text-gray-700 mb-2">
        نوع المصروف
      </label>
      <input
        id="expense-type"
        v-model="formData.type"
        type="text"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
        placeholder="أدخل نوع المصروف"
        required
      />
    </div>

    <!-- الصف الثاني: العدد والأجرة اليومية -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="quantity" class="block text-sm font-medium text-gray-700 mb-2">
          العدد
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
        <label for="daily-wage" class="block text-sm font-medium text-gray-700 mb-2">
          الأجرة اليومية
        </label>
        <input
          id="daily-wage"
          v-model="formData.price"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
    </div>

    <!-- الصف الثالث: المجموع والتفاصيل -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="total" class="block text-sm font-medium text-gray-700 mb-2">
          المجموع
        </label>
        <input
          id="total"
          v-model="formData.totalPrice"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
      <div>
        <label for="details" class="block text-sm font-medium text-gray-700 mb-2">
          التفاصيل
        </label>
        <input
          id="details"
          v-model="formData.details"
          type="text"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
          placeholder="أدخل التفاصيل"
          required
        />
      </div>
    </div>

    <!-- ملاحظات إضافية -->
    <div>
      <label for="additional-notes" class="block text-sm font-medium text-gray-700 mb-2">
        ملاحظات إضافية
      </label>
      <textarea
        id="additional-notes"
        v-model="formData.additionalNotes"
        rows="3"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm px-3 py-2 border"
        placeholder="أدخل ملاحظات إضافية"
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
  details: '',
  additionalNotes: '',
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
    details: '',
    additionalNotes: ''
  }
}

// دالة للتحقق من صحة البيانات
const isValid = () => {
  return formData.value.type && 
         formData.value.quantity && 
         formData.value.price && 
         formData.value.totalPrice && 
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
