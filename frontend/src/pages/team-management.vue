<template>
  <v-container fluid class="team-management-page">
    <!-- Header Section -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <v-icon size="48" color="primary" class="title-icon">mdi-account-group</v-icon>
          <h1 class="main-title">إدارة الفريق</h1>
          <p class="main-subtitle">إدارة شاملة لأعضاء الفريق وتوزيع المهام</p>
        </div>
        <div class="header-actions">
          <v-btn
            color="primary"
            variant="elevated"
            size="large"
            prepend-icon="mdi-plus"
            @click="showAddMemberDialog = true"
          >
            إضافة عضو جديد
          </v-btn>
        </div>
      </div>
    </div>

    <!-- Statistics Cards -->
    <div class="stats-section">
      <v-row>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card total-members">
            <v-card-text>
              <div class="stat-content">
                <v-icon size="40" color="primary">mdi-account-multiple</v-icon>
                <div class="stat-info">
                  <h3 class="stat-number">{{ totalMembers }}</h3>
                  <p class="stat-label">إجمالي الأعضاء</p>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card active-members">
            <v-card-text>
              <div class="stat-content">
                <v-icon size="40" color="success">mdi-account-check</v-icon>
                <div class="stat-info">
                  <h3 class="stat-number">{{ activeMembers }}</h3>
                  <p class="stat-label">أعضاء نشطين</p>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card departments">
            <v-card-text>
              <div class="stat-content">
                <v-icon size="40" color="info">mdi-office-building</v-icon>
                <div class="stat-info">
                  <h3 class="stat-number">{{ totalDepartments }}</h3>
                  <p class="stat-label">الأقسام</p>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card tasks">
            <v-card-text>
              <div class="stat-content">
                <v-icon size="40" color="warning">mdi-clipboard-list</v-icon>
                <div class="stat-info">
                  <h3 class="stat-number">{{ totalTasks }}</h3>
                  <p class="stat-label">المهام المكلفة</p>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <!-- Search and Filter Section -->
    <v-card class="search-filter-card">
      <v-card-text>
        <div class="search-filter-content">
          <v-row>
            <v-col cols="12" md="4">
              <v-text-field
                v-model="searchQuery"
                label="البحث في الأعضاء..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                clearable
                class="search-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="selectedDepartment"
                :items="departmentOptions"
                label="القسم"
                variant="outlined"
                clearable
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="3">
              <v-select
                v-model="selectedRole"
                :items="roleOptions"
                label="الدور"
                variant="outlined"
                clearable
                class="filter-field"
              />
            </v-col>
            <v-col cols="12" md="2">
              <v-btn
                color="primary"
                variant="outlined"
                @click="resetFilters"
                class="reset-btn"
              >
                إعادة تعيين
              </v-btn>
            </v-col>
          </v-row>
        </div>
      </v-card-text>
    </v-card>

    <!-- Team Members Table -->
    <v-card class="team-table-card">
      <v-card-title class="table-title">
        <v-icon size="28" color="primary" class="mr-2">mdi-account-group</v-icon>
        <span class="text-h5 font-weight-bold">قائمة أعضاء الفريق</span>
        <v-spacer />
        <v-chip color="primary" variant="elevated" size="large">
          {{ filteredMembers.length }} عضو
        </v-chip>
      </v-card-title>
      
      <v-data-table
        :headers="headers"
        :items="filteredMembers"
        :items-per-page="10"
        class="team-data-table"
        :loading="loading"
        :header-props="{
          style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
        }"
      >
        <!-- Avatar Column -->
        <template #item.avatar="{ item }">
          <v-avatar size="40" class="member-avatar">
            <v-img v-if="item.avatar" :src="item.avatar" />
            <v-icon v-else size="24" color="primary">mdi-account</v-icon>
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
              size="small"
              color="info"
              variant="elevated"
              @click="viewMember(item)"
              class="action-btn view-btn"
            />
            <v-btn
              icon="mdi-pencil"
              size="small"
              color="warning"
              variant="elevated"
              @click="editMember(item)"
              class="action-btn edit-btn"
            />
            <v-btn
              icon="mdi-key"
              size="small"
              color="primary"
              variant="elevated"
              @click="resetPassword(item)"
              class="action-btn reset-btn"
            />
            <v-btn
              icon="mdi-delete"
              size="small"
              color="error"
              variant="elevated"
              @click="deleteMember(item)"
              class="action-btn delete-btn"
            />
          </div>
        </template>
      </v-data-table>
    </v-card>

    <!-- Add Member Dialog -->
    <v-dialog v-model="showAddMemberDialog" max-width="600px" persistent>
      <v-card class="add-member-dialog image-style-dialog">
        <v-card-title class="dialog-header">
          <div class="header-content">
            <div class="header-left">
              <v-icon size="32" color="white" class="mr-2">mdi-account-plus</v-icon>
              <span class="dialog-title">إضافة عضو جديد</span>
            </div>
          </div>
        </v-card-title>
        
        <v-card-text class="dialog-body">
          <v-form ref="addMemberForm" v-model="addFormValid">
            <v-row class="form-fields">
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-text-field
                    v-model="newMember.name"
                    label="الاسم الكامل"
                    variant="outlined"
                    :rules="[v => !!v || 'الاسم مطلوب']"
                    required
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-text-field
                    v-model="newMember.email"
                    label="البريد الإلكتروني"
                    type="email"
                    variant="outlined"
                    :rules="emailRules"
                    required
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-text-field
                    v-model="newMember.phone"
                    label="رقم الهاتف"
                    variant="outlined"
                    :rules="[v => !!v || 'رقم الهاتف مطلوب']"
                    required
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-select
                    v-model="newMember.department"
                    :items="departmentOptions"
                    label="القسم"
                    variant="outlined"
                    :rules="[v => !!v || 'القسم مطلوب']"
                    required
                    class="black-dropdown-select"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-select
                    v-model="newMember.role"
                    :items="roleOptions"
                    label="الدور"
                    variant="outlined"
                    :rules="[v => !!v || 'الدور مطلوب']"
                    required
                    class="black-dropdown-select"
                  />
                </div>
              </v-col>
              <v-col cols="12" md="6">
                <div class="form-field">
                  <v-select
                    v-model="newMember.status"
                    :items="statusOptions"
                    label="الحالة"
                    variant="outlined"
                    :rules="[v => !!v || 'الحالة مطلوبة']"
                    required
                    class="black-dropdown-select"
                  />
                </div>
              </v-col>
              <v-col cols="12">
                <div class="form-field">
                  <v-textarea
                    v-model="newMember.notes"
                    label="ملاحظات إضافية"
                    variant="outlined"
                    rows="3"
                  />
                </div>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        
        <v-card-actions class="dialog-actions">
          <v-spacer />
          <v-btn
            color="grey"
            variant="text"
            @click="closeAddMemberDialog"
          >
            إلغاء
          </v-btn>
          <v-btn
            color="primary"
            variant="elevated"
            :disabled="!addFormValid"
            :loading="saving"
            @click="saveNewMember"
          >
            حفظ
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- View Member Dialog -->
    <v-dialog v-model="showViewMemberDialog" max-width="500px">
      <v-card class="view-member-dialog">
        <v-card-title class="dialog-header">
          <v-icon size="32" color="info" class="mr-2">mdi-account-details</v-icon>
          <span class="dialog-title">تفاصيل العضو</span>
        </v-card-title>
        
        <v-card-text v-if="selectedMember">
          <div class="member-details">
            <div class="member-avatar-large">
              <v-avatar size="80">
                <v-img v-if="selectedMember.avatar" :src="selectedMember.avatar" />
                <v-icon v-else size="40" color="primary">mdi-account</v-icon>
              </v-avatar>
            </div>
            
            <div class="member-info-details">
              <h3 class="member-name-large">{{ selectedMember.name }}</h3>
              <p class="member-email-large">{{ selectedMember.email }}</p>
              
              <div class="member-stats">
                <div class="stat-item">
                  <v-icon color="primary">mdi-phone</v-icon>
                  <span>{{ selectedMember.phone }}</span>
                </div>
                <div class="stat-item">
                  <v-icon color="success">mdi-office-building</v-icon>
                  <span>{{ selectedMember.department }}</span>
                </div>
                <div class="stat-item">
                  <v-icon color="info">mdi-account-tie</v-icon>
                  <span>{{ selectedMember.role }}</span>
                </div>
                <div class="stat-item">
                  <v-icon :color="selectedMember.status === 'نشط' ? 'success' : 'error'">
                    {{ selectedMember.status === 'نشط' ? 'mdi-check-circle' : 'mdi-close-circle' }}
                  </v-icon>
                  <span>{{ selectedMember.status }}</span>
                </div>
              </div>
              
              <div v-if="selectedMember.notes" class="member-notes">
                <h4>ملاحظات:</h4>
                <p>{{ selectedMember.notes }}</p>
              </div>
            </div>
          </div>
        </v-card-text>
        
        <v-card-actions class="dialog-actions">
          <v-spacer />
          <v-btn
            color="primary"
            variant="elevated"
            @click="closeViewMemberDialog"
          >
            إغلاق
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

// Reactive data
const loading = ref(false)
const saving = ref(false)
const searchQuery = ref('')
const selectedDepartment = ref(null)
const selectedRole = ref(null)
const showAddMemberDialog = ref(false)
const showViewMemberDialog = ref(false)
const selectedMember = ref(null)
const addFormValid = ref(false)

// Form data
const newMember = ref({
  name: '',
  email: '',
  phone: '',
  department: '',
  role: '',
  status: 'نشط',
  notes: ''
})

// Sample data
const teamMembers = ref([
  {
    id: 1,
    name: 'أحمد محمد علي',
    email: 'ahmed.mohamed@company.com',
    phone: '+964 770 123 4567',
    department: 'الهندسة المدنية',
    role: 'مهندس أول',
    status: 'نشط',
    tasksCount: 5,
    avatar: null,
    notes: 'خبرة 5 سنوات في المشاريع الكبيرة'
  },
  {
    id: 2,
    name: 'فاطمة حسن',
    email: 'fatima.hassan@company.com',
    phone: '+964 770 234 5678',
    department: 'الهندسة المعمارية',
    role: 'مهندسة معمارية',
    status: 'نشط',
    tasksCount: 3,
    avatar: null,
    notes: 'متخصصة في التصميم المعماري الحديث'
  },
  {
    id: 3,
    name: 'محمد عبدالله',
    email: 'mohamed.abdullah@company.com',
    phone: '+964 770 345 6789',
    department: 'إدارة المشاريع',
    role: 'مدير مشروع',
    status: 'نشط',
    tasksCount: 8,
    avatar: null,
    notes: 'مدير مشاريع ذو خبرة واسعة'
  },
  {
    id: 4,
    name: 'سارة أحمد',
    email: 'sara.ahmed@company.com',
    phone: '+964 770 456 7890',
    department: 'الهندسة الكهربائية',
    role: 'مهندسة كهربائية',
    status: 'غير نشط',
    tasksCount: 2,
    avatar: null,
    notes: 'متخصصة في الأنظمة الكهربائية'
  },
  {
    id: 5,
    name: 'علي محمود',
    email: 'ali.mahmoud@company.com',
    phone: '+964 770 567 8901',
    department: 'الهندسة الميكانيكية',
    role: 'مهندس ميكانيكي',
    status: 'نشط',
    tasksCount: 4,
    avatar: null,
    notes: 'خبرة في أنظمة التكييف والتهوية'
  }
])

// Options
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
  'مدير مشروع',
  'مهندس أول',
  'مهندس',
  'مهندس مساعد',
  'مهندسة معمارية',
  'مهندسة كهربائية',
  'مهندسة ميكانيكية',
  'مدير قسم',
  'مشرف',
  'مساعد إداري'
]

const statusOptions = [
  'نشط',
  'غير نشط',
  'إجازة',
  'مستقيل'
]

// Table headers
const headers = [
  { 
    title: 'الصورة', 
    key: 'avatar', 
    sortable: false, 
    width: '80px',
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  },
  { 
    title: 'الاسم', 
    key: 'name', 
    sortable: true,
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  },
  { 
    title: 'القسم', 
    key: 'department', 
    sortable: true,
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  },
  { 
    title: 'الدور', 
    key: 'role', 
    sortable: true,
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  },
  { 
    title: 'الحالة', 
    key: 'status', 
    sortable: true,
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  },
  { 
    title: 'المهام', 
    key: 'tasks', 
    sortable: true,
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  },
  { 
    title: 'الإجراءات', 
    key: 'actions', 
    sortable: false, 
    width: '200px',
    headerProps: {
      style: 'background: linear-gradient(135deg, #059669 0%, #10b981 100%); color: white; font-weight: 700;'
    }
  }
]

// Validation rules
const emailRules = [
  v => !!v || 'البريد الإلكتروني مطلوب',
  v => /.+@.+\..+/.test(v) || 'البريد الإلكتروني غير صحيح'
]

// Computed properties
const totalMembers = computed(() => teamMembers.value.length)
const activeMembers = computed(() => teamMembers.value.filter(member => member.status === 'نشط').length)
const totalDepartments = computed(() => new Set(teamMembers.value.map(member => member.department)).size)
const totalTasks = computed(() => teamMembers.value.reduce((sum, member) => sum + member.tasksCount, 0))

const filteredMembers = computed(() => {
  let filtered = teamMembers.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(member =>
      member.name.toLowerCase().includes(query) ||
      member.email.toLowerCase().includes(query) ||
      member.phone.includes(query)
    )
  }

  if (selectedDepartment.value) {
    filtered = filtered.filter(member => member.department === selectedDepartment.value)
  }

  if (selectedRole.value) {
    filtered = filtered.filter(member => member.role === selectedRole.value)
  }

  return filtered
})

// Methods
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

const resetFilters = () => {
  searchQuery.value = ''
  selectedDepartment.value = null
  selectedRole.value = null
}

const viewMember = (member) => {
  selectedMember.value = member
  showViewMemberDialog.value = true
}

const editMember = (member) => {
  // TODO: Implement edit functionality
}

const resetPassword = (member) => {
  // TODO: Implement reset password functionality
}

const deleteMember = (member) => {
  // TODO: Implement delete functionality
}

const closeAddMemberDialog = () => {
  showAddMemberDialog.value = false
  newMember.value = {
    name: '',
    email: '',
    phone: '',
    department: '',
    role: '',
    status: 'نشط',
    notes: ''
  }
}

const saveNewMember = async () => {
  if (!addFormValid.value) return

  saving.value = true
  
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 1000))
  
  const newId = Math.max(...teamMembers.value.map(m => m.id)) + 1
  const member = {
    id: newId,
    ...newMember.value,
    tasksCount: 0,
    avatar: null
  }
  
  teamMembers.value.push(member)
  closeAddMemberDialog()
  saving.value = false
}

const closeViewMemberDialog = () => {
  showViewMemberDialog.value = false
  selectedMember.value = null
}

onMounted(() => {
  // Initialize data if needed
})
</script>


<style>
/* Import page styles */
@import './styles/team-management.css';
</style>

<style scoped>
/* Component-specific overrides (if any) */
</style>
