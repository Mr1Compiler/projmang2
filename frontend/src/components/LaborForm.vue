<template>
  <div class="space-y-6">
    <!-- اسم العامل -->
    <div>
      <label for="worker-name" class="block text-sm font-medium text-gray-700 mb-2">
        اسم العامل
      </label>
      <input
        id="worker-name"
        v-model="formData.name"
        type="text"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
        placeholder="أدخل اسم العامل"
        required
      />
    </div>

    <!-- الصف الثاني: المهنة والراتب -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="profession" class="block text-sm font-medium text-gray-700 mb-2">
          المهنة
        </label>
        <input
          id="profession"
          v-model="formData.profession"
          type="text"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
          placeholder="أدخل المهنة"
          required
        />
      </div>
      <div>
        <label for="salary" class="block text-sm font-medium text-gray-700 mb-2">
          الراتب اليومي (IQD)
        </label>
        <input
          id="salary"
          v-model="formData.salary"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
    </div>

    <!-- الصف الثالث: عدد الأيام والمجموع -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="days-worked" class="block text-sm font-medium text-gray-700 mb-2">
          عدد الأيام
        </label>
        <input
          id="days-worked"
          v-model="formData.daysWorked"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
      <div>
        <label for="total-salary" class="block text-sm font-medium text-gray-700 mb-2">
          المجموع (IQD)
        </label>
        <input
          id="total-salary"
          v-model="formData.totalSalary"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
    </div>

    <!-- الصف الرابع: رقم الهاتف والعنوان -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="phone" class="block text-sm font-medium text-gray-700 mb-2">
          رقم الهاتف
        </label>
        <input
          id="phone"
          v-model="formData.phone"
          type="tel"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
          placeholder="07XX XXX XXXX"
          required
        />
      </div>
      <div>
        <label for="address" class="block text-sm font-medium text-gray-700 mb-2">
          العنوان
        </label>
        <input
          id="address"
          v-model="formData.address"
          type="text"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
          placeholder="أدخل العنوان"
          required
        />
      </div>
    </div>

    <!-- ملاحظات -->
    <div>
      <label for="notes" class="block text-sm font-medium text-gray-700 mb-2">
        ملاحظات
      </label>
      <textarea
        id="notes"
        v-model="formData.notes"
        rows="3"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-green-500 focus:ring-green-500 sm:text-sm px-3 py-2 border"
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
  name: '',
  profession: '',
  salary: '',
  daysWorked: '',
  totalSalary: '',
  phone: '',
  address: '',
  notes: '',
  ...props.initialData
})

// مراقبة التغييرات وإرسالها للوالد
watch(formData, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// دالة لإعادة تعيين النموذج
const resetForm = () => {
  formData.value = {
    name: '',
    profession: '',
    salary: '',
    daysWorked: '',
    totalSalary: '',
    phone: '',
    address: '',
    notes: ''
  }
}

// دالة للتحقق من صحة البيانات
const isValid = () => {
  return formData.value.name && 
         formData.value.profession && 
         formData.value.salary && 
         formData.value.daysWorked && 
         formData.value.totalSalary &&
         formData.value.phone &&
         formData.value.address
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
  border-color: #4caf50;
  box-shadow: 0 0 0 3px rgba(76, 175, 80, 0.1);
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
