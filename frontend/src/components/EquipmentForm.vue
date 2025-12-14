<template>
  <div class="space-y-6">
    <!-- نوع الآلة -->
    <div>
      <label for="equipment-type" class="block text-sm font-medium text-gray-700 mb-2">
        نوع الآلة
      </label>
      <input
        id="equipment-type"
        v-model="formData.type"
        type="text"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
        placeholder="أدخل نوع الآلة"
        required
      />
    </div>

    <!-- الصف الثاني: التكلفة اليومية وعدد الأيام -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="daily-cost" class="block text-sm font-medium text-gray-700 mb-2">
          التكلفة اليومية (IQD)
        </label>
        <input
          id="daily-cost"
          v-model="formData.dailyCost"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
      <div>
        <label for="days-used" class="block text-sm font-medium text-gray-700 mb-2">
          عدد الأيام المستخدمة
        </label>
        <input
          id="days-used"
          v-model="formData.daysUsed"
          type="number"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
          placeholder="0"
          required
        />
      </div>
    </div>

    <!-- الصف الثالث: المجموع -->
    <div>
      <label for="total-cost" class="block text-sm font-medium text-gray-700 mb-2">
        المجموع (IQD)
      </label>
      <input
        id="total-cost"
        v-model="formData.totalCost"
        type="number"
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
        placeholder="0"
        required
      />
    </div>

    <!-- الصف الرابع: اسم السائق ورقم الهاتف -->
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <div>
        <label for="driver-name" class="block text-sm font-medium text-gray-700 mb-2">
          اسم السائق
        </label>
        <input
          id="driver-name"
          v-model="formData.driverName"
          type="text"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
          placeholder="أدخل اسم السائق"
          required
        />
      </div>
      <div>
        <label for="driver-phone" class="block text-sm font-medium text-gray-700 mb-2">
          رقم هاتف السائق
        </label>
        <input
          id="driver-phone"
          v-model="formData.driverPhone"
          type="tel"
          class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
          placeholder="07XX XXX XXXX"
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
        class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm px-3 py-2 border"
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
  dailyCost: '',
  daysUsed: '',
  totalCost: '',
  driverName: '',
  driverPhone: '',
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
    type: '',
    dailyCost: '',
    daysUsed: '',
    totalCost: '',
    driverName: '',
    driverPhone: '',
    notes: ''
  }
}

// دالة للتحقق من صحة البيانات
const isValid = () => {
  return formData.value.type && 
         formData.value.dailyCost && 
         formData.value.daysUsed && 
         formData.value.totalCost &&
         formData.value.driverName &&
         formData.value.driverPhone
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
textarea:focus,
select:focus {
  outline: none;
  border-color: #6366f1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

label {
  color: #374151;
  font-weight: 500;
}

input,
textarea,
select {
  transition: all 0.2s ease-in-out;
}

input:hover,
textarea:hover,
select:hover {
  border-color: #9ca3af;
}

select {
  background-color: white;
  cursor: pointer;
}
</style>
