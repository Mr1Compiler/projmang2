<template>
  <div>
    <!-- Team Management Section -->
    <v-expand-transition>
      <div v-if="show" class="team-management-section">
        <v-card class="team-management-card">
          <v-card-title class="team-section-header">
            <v-icon size="32" color="success" class="mr-2">mdi-account-group</v-icon>
            <span class="text-h5 font-weight-bold">إدارة الفريق</span>
            <v-spacer />
            <v-btn
              color="success"
              variant="elevated"
              size="small"
              prepend-icon="mdi-plus"
              @click="showAddMemberDialog = true"
            >
              إضافة عضو
            </v-btn>
          </v-card-title>

          <v-card-text>
            <!-- Team Statistics -->
            <div class="team-stats-row mb-4">
              <v-row>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="team-stat-card">
                    <v-card-text class="text-center">
                      <v-icon size="32" color="primary" class="mb-2">mdi-account-multiple</v-icon>
                      <h3 class="text-h6 font-weight-bold">{{ teamMembers.length }}</h3>
                      <p class="text-caption">إجمالي الأعضاء</p>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="team-stat-card">
                    <v-card-text class="text-center">
                      <v-icon size="32" color="success" class="mb-2">mdi-account-check</v-icon>
                      <h3 class="text-h6 font-weight-bold">{{ activeTeamMembers }}</h3>
                      <p class="text-caption">أعضاء نشطين</p>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="team-stat-card">
                    <v-card-text class="text-center">
                      <v-icon size="32" color="info" class="mb-2">mdi-office-building</v-icon>
                      <h3 class="text-h6 font-weight-bold">{{ teamDepartments }}</h3>
                      <p class="text-caption">الأقسام</p>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="team-stat-card">
                    <v-card-text class="text-center">
                      <v-icon size="32" color="warning" class="mb-2">mdi-clipboard-list</v-icon>
                      <h3 class="text-h6 font-weight-bold">{{ totalTeamTasks }}</h3>
                      <p class="text-caption">المهام</p>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
            </div>

            <!-- Team Members Table -->
            <v-data-table
              :headers="teamHeaders"
              :items="teamMembers"
              :items-per-page="5"
              class="team-members-table"
              :loading="loading"
            >
              <!-- Avatar Column -->
              <template #item.avatar="{ item }">
                <v-avatar size="32" class="member-avatar">
                  <v-img v-if="item.avatar" :src="item.avatar" />
                  <v-icon v-else size="20" color="primary">mdi-account</v-icon>
                </v-avatar>
              </template>

              <!-- Name Column -->
              <template #item.name="{ item }">
                <div class="member-info">
                  <h4 class="member-name">{{ item.name }}</h4>
                  <p class="member-email">{{ item.email }}</p>
                </div>
              </template>

              <!-- Department Column -->
              <template #item.department="{ item }">
                <v-chip
                  :color="getDepartmentColor(item.department)"
                  variant="elevated"
                  size="small"
                >
                  {{ item.department }}
                </v-chip>
              </template>

              <!-- Role Column -->
              <template #item.role="{ item }">
                <v-chip
                  :color="getRoleColor(item.role)"
                  variant="elevated"
                  size="small"
                >
                  {{ item.role }}
                </v-chip>
              </template>

              <!-- Status Column -->
              <template #item.status="{ item }">
                <v-chip
                  :color="item.status === 'نشط' ? 'success' : 'error'"
                  variant="elevated"
                  size="small"
                >
                  <v-icon start>{{ item.status === 'نشط' ? 'mdi-check-circle' : 'mdi-close-circle' }}</v-icon>
                  {{ item.status }}
                </v-chip>
              </template>

              <!-- Tasks Column -->
              <template #item.tasks="{ item }">
                <div class="tasks-info">
                  <span class="task-count">{{ item.tasksCount }}</span>
                  <span class="task-label">مهمة</span>
                </div>
              </template>

              <!-- Actions Column -->
              <template #item.actions="{ item }">
                <div class="action-buttons">
                  <v-btn
                    icon="mdi-eye"
                    size="x-small"
                    color="info"
                    variant="elevated"
                    @click="$emit('view-member', item)"
                  />
                  <v-btn
                    icon="mdi-pencil"
                    size="x-small"
                    color="warning"
                    variant="elevated"
                    @click="$emit('edit-member', item)"
                  />
                  <v-btn
                    icon="mdi-delete"
                    size="x-small"
                    color="error"
                    variant="elevated"
                    @click="$emit('delete-member', item)"
                  />
                </div>
              </template>
            </v-data-table>
          </v-card-text>
        </v-card>
      </div>
    </v-expand-transition>

    <!-- Add Team Member Dialog -->
    <v-dialog v-model="showAddMemberDialog" max-width="600">
      <v-card class="team-member-dialog-card">
        <v-card-title class="team-member-dialog-header">
          <v-icon size="28" color="white" class="mr-2">mdi-account-plus</v-icon>
          <span class="text-h5 font-weight-black">إضافة عضو جديد</span>
        </v-card-title>

        <v-card-text>
          <v-form ref="teamMemberFormRef" v-model="teamMemberFormValid">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="teamMemberForm.name"
                  label="اسم العضو"
                  variant="outlined"
                  :rules="[v => !!v || 'اسم العضو مطلوب']"
                  required
                  class="team-member-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="teamMemberForm.email"
                  label="البريد الإلكتروني"
                  variant="outlined"
                  type="email"
                  :rules="[v => !!v || 'البريد الإلكتروني مطلوب', v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح']"
                  required
                  class="team-member-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="teamMemberForm.phone"
                  label="رقم الهاتف"
                  variant="outlined"
                  :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                  required
                  class="team-member-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="teamMemberForm.department"
                  :items="departmentOptions"
                  label="القسم"
                  variant="outlined"
                  required
                  class="team-member-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="teamMemberForm.role"
                  :items="roleOptions"
                  label="المنصب"
                  variant="outlined"
                  required
                  class="team-member-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="teamMemberForm.status"
                  :items="statusOptions"
                  label="الحالة"
                  variant="outlined"
                  required
                  class="team-member-form-field"
                  style="--v-theme-primary: #000000; --v-field-label-color: #000000; text-align: center; direction: rtl;"
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>

        <v-card-actions class="pa-4">
          <v-spacer></v-spacer>
          <v-btn
            color="grey"
            variant="outlined"
            @click="closeAddMemberDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="success"
            variant="elevated"
            @click="saveTeamMember"
            :disabled="!teamMemberFormValid"
          >
            إضافة العضو
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  teamMembers: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['view-member', 'edit-member', 'delete-member', 'add-member'])

// Dialog state
const showAddMemberDialog = ref(false)
const teamMemberFormValid = ref(false)
const teamMemberFormRef = ref(null)
const teamMemberForm = ref({
  name: '',
  email: '',
  phone: '',
  department: '',
  role: '',
  status: 'نشط'
})

// Table headers
const teamHeaders = [
  { title: 'الصورة', key: 'avatar', sortable: false, width: '60px' },
  { title: 'الاسم', key: 'name', sortable: true },
  { title: 'القسم', key: 'department', sortable: true },
  { title: 'الدور', key: 'role', sortable: true },
  { title: 'الحالة', key: 'status', sortable: true },
  { title: 'المهام', key: 'tasks', sortable: true },
  { title: 'الإجراءات', key: 'actions', sortable: false, width: '120px' }
]

// Dropdown options
const departmentOptions = [
  'الهندسة المدنية',
  'الهندسة المعمارية',
  'الهندسة الكهربائية',
  'الهندسة الميكانيكية',
  'إدارة المشاريع',
  'الموارد البشرية',
  'المحاسبة',
  'التسويق'
]

const roleOptions = [
  'مدير المشروع',
  'مهندس أول',
  'مهندس',
  'مساعد مهندس',
  'فني',
  'مساعد إداري',
  'محاسب',
  'مستشار'
]

const statusOptions = [
  'نشط',
  'غير نشط',
  'في إجازة',
  'مستقيل'
]

// Computed statistics
const activeTeamMembers = computed(() => props.teamMembers.filter(member => member.status === 'نشط').length)
const teamDepartments = computed(() => new Set(props.teamMembers.map(member => member.department)).size)
const totalTeamTasks = computed(() => props.teamMembers.reduce((sum, member) => sum + (member.tasksCount || 0), 0))

// Helper functions
const getDepartmentColor = (department) => {
  const colors = {
    'الهندسة المدنية': 'blue',
    'الهندسة المعمارية': 'purple',
    'الهندسة الكهربائية': 'orange',
    'الهندسة الميكانيكية': 'green',
    'إدارة المشاريع': 'red',
    'الموارد البشرية': 'pink',
    'المحاسبة': 'teal',
    'التسويق': 'indigo'
  }
  return colors[department] || 'grey'
}

const getRoleColor = (role) => {
  if (role.includes('مدير')) return 'red'
  if (role.includes('أول')) return 'blue'
  if (role.includes('مهندس')) return 'green'
  return 'grey'
}

// Dialog functions
const closeAddMemberDialog = () => {
  showAddMemberDialog.value = false
  teamMemberForm.value = {
    name: '',
    email: '',
    phone: '',
    department: '',
    role: '',
    status: 'نشط'
  }
  teamMemberFormValid.value = false
}

const saveTeamMember = () => {
  if (teamMemberFormValid.value) {
    const newMember = {
      ...teamMemberForm.value,
      id: Date.now(),
      tasksCount: 0,
      avatar: null
    }
    emit('add-member', newMember)
    closeAddMemberDialog()
  }
}
</script>

<style>
/* Import shared project styles */
@import './styles/projects.css';
</style>

<style scoped>
/* Component-specific styles for TeamManagement */

/* Team members table */
.team-members-table {
  border-radius: 12px !important;
}

/* Member info layout */
.member-info {
  display: flex;
  flex-direction: column;
}

.member-name {
  font-weight: 600;
  font-size: 0.95rem;
  color: #1f2937;
  margin: 0;
}

.member-email {
  font-size: 0.8rem;
  color: #6b7280;
  margin: 0;
}

.member-avatar {
  background: linear-gradient(135deg, #e0e7ff 0%, #c7d2fe 100%) !important;
}

/* Tasks display */
.tasks-info {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.task-count {
  font-weight: 700;
  font-size: 1rem;
  color: #1976d2;
}

.task-label {
  font-size: 0.8rem;
  color: #6b7280;
}

/* Action buttons */
.action-buttons {
  display: flex;
  gap: 0.25rem;
}

/* Team member dialog */
.team-member-dialog-card {
  border-radius: 16px !important;
  overflow: hidden;
}

.team-member-dialog-header {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%) !important;
  color: white !important;
  padding: 1.5rem !important;
}

.team-member-form-field {
  margin-bottom: 0.5rem;
}
</style>
