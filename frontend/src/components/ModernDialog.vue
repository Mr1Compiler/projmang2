<template>
  <TransitionRoot as="template" :show="open">
    <Dialog as="div" class="relative z-50" @close="closeDialog">
      <TransitionChild
        as="template"
        enter="ease-out duration-300"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="ease-in duration-200"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
      </TransitionChild>

      <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
        <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
          <TransitionChild
            as="template"
            enter="ease-out duration-300"
            enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            enter-to="opacity-100 translate-y-0 sm:scale-100"
            leave="ease-in duration-200"
            leave-from="opacity-100 translate-y-0 sm:scale-100"
            leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
          >
            <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
              <!-- Header -->
              <div class="bg-blue-600 px-3 py-2 sm:px-4 dialog-header-small">
                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <component :is="icon" class="h-5 w-5 text-white mr-2" />
                    <DialogTitle as="h3" class="text-base font-semibold leading-5 text-white">
                      {{ title }}
                    </DialogTitle>
                  </div>
                  <button
                    type="button"
                    class="rounded-md bg-blue-600 text-white hover:bg-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 close-header-btn"
                    @click="closeDialog"
                  >
                    <span class="sr-only">إغلاق</span>
                    <XMarkIcon class="h-5 w-5" aria-hidden="true" />
                  </button>
                </div>
              </div>

              <!-- Content -->
              <div class="bg-white px-4 py-5 sm:p-6">
                <slot />
              </div>

              <!-- Footer -->
              <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6 dialog-footer">
                <button
                  type="button"
                  class="save-btn-dialog"
                  @click="saveAction"
                >
                  <CheckIcon class="h-4 w-4 mr-2" />
                  {{ saveText }}
                </button>
                <button
                  type="button"
                  class="cancel-btn-dialog"
                  @click="closeDialog"
                >
                  إلغاء
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<script setup>
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
import { XMarkIcon, CheckIcon } from '@heroicons/vue/24/outline'

const props = defineProps({
  open: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: 'إضافة عنصر'
  },
  saveText: {
    type: String,
    default: 'حفظ'
  },
  icon: {
    type: [String, Object],
    default: 'PlusIcon'
  }
})

const emit = defineEmits(['close', 'save'])

const closeDialog = () => {
  emit('close')
}

const saveAction = () => {
  emit('save')
}
</script>

<style scoped>
/* تحسينات إضافية للتصميم */
.bg-blue-600 {
  background-color: #1976d2;
}

.hover\:bg-blue-500:hover {
  background-color: #1565c0;
}

.focus\:ring-blue-500:focus {
  --tw-ring-color: #1976d2;
}

/* تنسيق شريط العنوان المصغر */
.dialog-header-small {
  padding: 0.75rem 1rem !important;
}

.dialog-header-small h3 {
  font-size: var(--font-size-base) !important;
  font-weight: 700 !important;
  line-height: 1.4 !important;
}

.close-header-btn {
  padding: 0.25rem !important;
  min-width: auto !important;
  width: 32px !important;
  height: 32px !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

/* تنسيق أزرار القائمة */
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 16px;
  padding: 1.25rem 1.5rem !important;
  background: #f8f9fa !important;
  border-top: 1px solid rgba(0, 0, 0, 0.05) !important;
}

.save-btn-dialog {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 14px !important;
  padding: 14px 32px !important;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 50%, #2563eb 100%) !important;
  color: #ffffff !important;
  font-size: var(--font-size-base) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: 140px !important;
  box-shadow: 
    0 6px 20px rgba(59, 130, 246, 0.4),
    0 3px 10px rgba(37, 99, 235, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.25) !important;
  border: 2px solid rgba(255, 255, 255, 0.4) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
  position: relative;
  overflow: hidden;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  gap: 10px;
}

.save-btn-dialog svg {
  width: 20px !important;
  height: 20px !important;
  stroke-width: 2.5 !important;
  color: #ffffff !important;
  flex-shrink: 0;
}

.save-btn-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  transition: left 0.6s ease;
  z-index: 1;
}

.save-btn-dialog:hover::before {
  left: 100%;
}

.save-btn-dialog:hover {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 50%, #1d4ed8 100%) !important;
  box-shadow: 
    0 8px 24px rgba(59, 130, 246, 0.5),
    0 4px 12px rgba(37, 99, 235, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.4) !important;
  transform: translateY(-2px) scale(1.02) !important;
  border-color: rgba(255, 255, 255, 0.5) !important;
}

.cancel-btn-dialog {
  font-weight: 700 !important;
  text-transform: none !important;
  border-radius: 14px !important;
  padding: 14px 32px !important;
  border: 2px solid rgba(156, 163, 175, 0.5) !important;
  color: #374151 !important;
  background: rgba(255, 255, 255, 1) !important;
  backdrop-filter: blur(10px) !important;
  font-size: var(--font-size-base) !important;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif !important;
  min-width: 140px !important;
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.1),
    0 2px 8px rgba(156, 163, 175, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  letter-spacing: 0.5px !important;
  line-height: 1.5 !important;
  position: relative;
  overflow: hidden;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.cancel-btn-dialog::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.cancel-btn-dialog:hover::before {
  left: 100%;
}

.cancel-btn-dialog:hover {
  background: rgba(255, 255, 255, 1) !important;
  border-color: rgba(156, 163, 175, 0.5) !important;
  color: #374151 !important;
  transform: translateY(-2px) scale(1.02) !important;
  box-shadow: 
    0 6px 20px rgba(0, 0, 0, 0.12),
    0 3px 10px rgba(156, 163, 175, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 1) !important;
}
</style>
