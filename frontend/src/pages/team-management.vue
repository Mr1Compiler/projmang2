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
  console.log('Edit member:', member)
}

const resetPassword = (member) => {
  // TODO: Implement reset password functionality
  console.log('Reset password for:', member)
}

const deleteMember = (member) => {
  // TODO: Implement delete functionality
  console.log('Delete member:', member)
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

<style scoped>
/* CSS قوي للتأكد من تطبيق التغييرات على عناوين الجدول */
.v-data-table.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}

/* Main Page Styles */
.team-management-page {
  background: linear-gradient(135deg, #e0e7ff 0%, #c7d2fe 50%, #a5b4fc 100%);
  min-height: 100vh;
  padding: 2rem 1rem;
}

/* Header Styles */
.page-header {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 50%, #e0e7ff 100%);
  border-radius: 20px;
  padding: 2rem;
  margin-bottom: 2rem;
  box-shadow: 0 8px 32px rgba(67, 56, 202, 0.1);
  border: 1px solid rgba(67, 56, 202, 0.2);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.main-title {
  font-size: 2.5rem;
  font-weight: 900;
  color: #1e293b;
  margin: 0;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.main-subtitle {
  font-size: 1.2rem;
  color: #64748b;
  margin: 0.5rem 0 0 0;
  font-weight: 500;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

/* Statistics Cards */
.stats-section {
  margin-bottom: 2rem;
}

.stat-card {
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.stat-info {
  flex: 1;
}

.stat-number {
  font-size: 2rem;
  font-weight: 900;
  color: #1e293b;
  margin: 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.stat-label {
  font-size: 1rem;
  color: #64748b;
  margin: 0;
  font-weight: 600;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

/* Search and Filter */
.search-filter-card {
  border-radius: 16px;
  margin-bottom: 2rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
}

.search-filter-content {
  padding: 1rem 0;
}

.search-field .v-field__input,
.filter-field .v-field__input {
  color: #1a1a1a !important;
  background: white !important;
  font-weight: 600;
  font-size: 1rem;
  padding: 12px 16px;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.search-field .v-field__input:focus,
.filter-field .v-field__input:focus {
  border-color: #4338ca;
  box-shadow: 0 0 0 4px rgba(67, 56, 202, 0.1);
}

.search-field .v-label,
.filter-field .v-label {
  color: #1a1a1a !important;
  font-weight: 800;
  font-size: 1rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.reset-btn {
  height: 56px;
  border-radius: 12px;
  font-weight: 600;
  text-transform: none;
}

/* Team Table */
.team-table-card {
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
}

.table-title {
  background: linear-gradient(135deg, rgba(5, 150, 105, 0.1) 0%, rgba(16, 185, 129, 0.1) 100%);
  border-bottom: 2px solid rgba(5, 150, 105, 0.2);
  padding: 1.5rem;
}

.team-data-table {
  background: transparent;
}

.team-data-table th,
.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-align: center !important;
  vertical-align: middle !important;
  padding: 1.2rem 0.8rem !important;
  border: none !important;
  box-shadow: 0 3px 12px rgba(4, 120, 87, 0.4) !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS إضافي لضمان تطبيق التغييرات على Vuetify */
.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%) !important;
  color: white !important;
}

.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: white !important;
}

/* التأكد من أن جميع عناوين الجدول تستخدم اللون الجديد */
.team-data-table table thead tr th {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%) !important;
  color: white !important;
  font-weight: 700 !important;
}

.team-data-table td {
  background: rgba(248, 250, 252, 0.9);
  color: #1a1a1a !important;
  font-weight: 500;
  text-align: center;
  vertical-align: middle;
  padding: 1rem 0.5rem;
  border-bottom: 1px solid rgba(226, 232, 240, 0.5);
}

.team-data-table tbody tr:hover td {
  background: rgba(5, 150, 105, 0.05);
  transform: scale(1.01);
  transition: all 0.3s ease;
}

/* Member Info */
.member-info {
  text-align: right;
}

.member-name {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 0.25rem 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-email {
  font-size: 0.9rem;
  color: #64748b;
  margin: 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-avatar {
  border: 2px solid rgba(67, 56, 202, 0.3);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* Tasks Info */
.tasks-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
}

.task-count {
  font-size: 1.2rem;
  font-weight: 700;
  color: #1e293b;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.task-label {
  font-size: 0.8rem;
  color: #64748b;
  font-weight: 500;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
  flex-wrap: wrap;
}

.action-btn {
  border-radius: 8px;
  font-weight: 600;
  text-transform: none;
  min-width: 36px;
  height: 36px;
}

.view-btn {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%) !important;
  color: white !important;
  box-shadow: 0 2px 8px rgba(59, 130, 246, 0.3);
}

.edit-btn {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%) !important;
  color: white !important;
  box-shadow: 0 2px 8px rgba(245, 158, 11, 0.3);
}

.reset-btn {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%) !important;
  color: white !important;
  box-shadow: 0 2px 8px rgba(139, 92, 246, 0.3);
}

.delete-btn {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%) !important;
  color: white !important;
  box-shadow: 0 2px 8px rgba(239, 68, 68, 0.3);
}

.action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}

/* Dialog Styles - تنسيق مشابه لصفحات أخرى */
.add-member-dialog.image-style-dialog {
  background: linear-gradient(135deg, #059669 0%, #10b981 100%);
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(5, 150, 105, 0.3);
}

.add-member-dialog .dialog-header {
  background: linear-gradient(135deg, #059669, #10b981);
  padding: 20px;
  color: white;
}

.add-member-dialog .header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.add-member-dialog .header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.add-member-dialog .dialog-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: white;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.add-member-dialog .dialog-body {
  padding: 30px;
  background: white !important;
}

.add-member-dialog .form-fields {
  margin-bottom: 20px;
}

.add-member-dialog .form-field {
  margin-bottom: 16px;
}

/* تحسين ألوان الحقول */
.add-member-dialog .form-field .v-field__input {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  background: white !important;
  border-radius: 8px !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
}

.add-member-dialog .form-field .v-field__input input,
.add-member-dialog .form-field .v-field__input textarea {
  color: #1a1a1a !important;
  background: white !important;
}

.add-member-dialog .form-field .v-field__outline {
  border-color: #d1d5db !important;
  border-width: 2px !important;
}

.add-member-dialog .form-field .v-label {
  color: #000000 !important;
  font-weight: 700 !important;
  font-size: 1rem !important;
  background: white !important;
  padding: 0 8px !important;
  opacity: 1 !important;
}

.add-member-dialog .form-field .v-field--focused .v-field__outline {
  border-color: #6b7280 !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(107, 114, 128, 0.1) !important;
}

.add-member-dialog .form-field .v-field:hover .v-field__outline {
  border-color: #9ca3af !important;
}

.add-member-dialog .form-field .v-field--focused .v-label {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

/* تحسين القوائم المنسدلة */
.add-member-dialog .form-field .v-select .v-field__input {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
}

.add-member-dialog .form-field .v-select .v-field__outline {
  border-color: #d1d5db !important;
  border-width: 2px !important;
}

.add-member-dialog .form-field .v-select .v-field--focused .v-field__outline {
  border-color: #6b7280 !important;
  border-width: 3px !important;
  box-shadow: 0 0 0 3px rgba(107, 114, 128, 0.1) !important;
}

.add-member-dialog .form-field .v-select .v-field:hover .v-field__outline {
  border-color: #9ca3af !important;
}

.add-member-dialog .form-field .v-select .v-label {
  color: #000000 !important;
  font-weight: 700 !important;
  background: white !important;
  opacity: 1 !important;
}

/* تحسين ألوان رسائل الخطأ */
.add-member-dialog .form-field .v-messages__message {
  color: #dc2626 !important;
  font-weight: 500 !important;
  font-size: 0.8rem !important;
}

.add-member-dialog .dialog-actions {
  padding: 20px 30px;
  background: #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* تنسيق القوائم المنسدلة المفتوحة - ألوان سوداء */
.add-member-dialog .form-field .v-list {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.add-member-dialog .form-field .v-list-item {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  background: white !important;
}

.add-member-dialog .form-field .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.add-member-dialog .form-field .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

/* إصلاح شامل لجميع عناصر القوائم المنسدلة - ألوان سوداء */
.add-member-dialog .v-menu__content .v-list,
.add-member-dialog .v-menu__content .v-list-item,
.add-member-dialog .v-menu__content .v-list-item *,
.add-member-dialog .v-overlay__content .v-list,
.add-member-dialog .v-overlay__content .v-list-item,
.add-member-dialog .v-overlay__content .v-list-item * {
  color: #1a1a1a !important;
  background: white !important;
}

.add-member-dialog .v-menu__content .v-list-item:hover,
.add-member-dialog .v-overlay__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.add-member-dialog .v-menu__content .v-list-item--active,
.add-member-dialog .v-overlay__content .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
}

/* إصلاح نص القوائم المنسدلة */
.add-member-dialog .v-list-item-title,
.add-member-dialog .v-list-item-content,
.add-member-dialog .v-list-item-content * {
  color: inherit !important;
  background: inherit !important;
}

/* View Member Dialog - يحتفظ بتصميمه الأصلي */
.view-member-dialog {
  border-radius: 20px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.view-member-dialog .dialog-header {
  background: linear-gradient(135deg, #4338ca 0%, #6366f1 100%);
  color: white;
  padding: 1.5rem;
  border-radius: 20px 20px 0 0;
}

.view-member-dialog .dialog-title {
  font-size: 1.5rem;
  font-weight: 700;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.view-member-dialog .dialog-actions {
  padding: 1.5rem;
  background: rgba(248, 250, 252, 0.5);
  border-radius: 0 0 20px 20px;
}

/* Member Details */
.member-details {
  text-align: center;
  padding: 1rem 0;
}

.member-avatar-large {
  margin-bottom: 1.5rem;
}

.member-avatar-large .v-avatar {
  border: 4px solid rgba(67, 56, 202, 0.3);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

.member-name-large {
  font-size: 1.8rem;
  font-weight: 800;
  color: #1e293b;
  margin: 0 0 0.5rem 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-email-large {
  font-size: 1.1rem;
  color: #64748b;
  margin: 0 0 1.5rem 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem;
  background: rgba(248, 250, 252, 0.8);
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.5);
}

.stat-item span {
  font-weight: 600;
  color: #1e293b;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-notes {
  text-align: right;
  padding: 1rem;
  background: rgba(248, 250, 252, 0.8);
  border-radius: 12px;
  border: 1px solid rgba(226, 232, 240, 0.5);
}

.member-notes h4 {
  color: #1e293b;
  font-weight: 700;
  margin: 0 0 0.5rem 0;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

.member-notes p {
  color: #64748b;
  margin: 0;
  line-height: 1.6;
  font-family: 'Cairo', 'Tajawal', 'Arial', sans-serif;
}

/* Responsive Design */
@media (max-width: 768px) {
  .team-management-page {
    padding: 1rem 0.5rem;
  }
  
  .page-header {
    padding: 1.5rem;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
  }
  
  .main-title {
    font-size: 2rem;
  }
  
  .main-subtitle {
    font-size: 1rem;
  }
  
  .stats-section .v-col {
    margin-bottom: 1rem;
  }
  
  .stat-content {
    flex-direction: column;
    text-align: center;
  }
  
  .action-buttons {
    flex-direction: column;
    gap: 0.25rem;
  }
  
  .member-stats {
    grid-template-columns: 1fr;
  }
}

/* CSS إضافي للتأكد من تطبيق التغييرات */
.v-data-table.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}
</style>

<style>
/* CSS غير محدود النطاق للتأكد من تطبيق التغييرات */
.v-data-table.team-data-table .v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table.team-data-table .v-data-table__wrapper table thead tr th .v-data-table-header__content .v-data-table-header__sort-icon {
  color: #ffffff !important;
  filter: drop-shadow(0 1px 1px rgba(0, 0, 0, 0.3)) !important;
}

/* CSS إضافي لجميع عناوين الجداول */
.v-data-table table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

/* CSS لضمان التطبيق على جميع عناصر Vuetify */
.v-data-table__wrapper table thead tr th {
  background: linear-gradient(135deg, #047857 0%, #059669 100%) !important;
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}

.v-data-table__wrapper table thead tr th .v-data-table-header__content {
  color: #ffffff !important;
  font-weight: 800 !important;
  font-size: 1.1rem !important;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3) !important;
  letter-spacing: 0.5px !important;
}
</style>

<style>
/* تنسيقات شاملة للقوائم المنسدلة في حوار إضافة عضو - ألوان سوداء */
.add-member-dialog .v-menu__content,
.add-member-dialog .v-overlay__content {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.add-member-dialog .v-menu__content .v-list,
.add-member-dialog .v-overlay__content .v-list {
  background: white !important;
  padding: 4px 0 !important;
}

.add-member-dialog .v-menu__content .v-list-item,
.add-member-dialog .v-overlay__content .v-list-item {
  color: #1a1a1a !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  min-height: 48px !important;
  background: white !important;
  transition: all 0.2s ease !important;
}

.add-member-dialog .v-menu__content .v-list-item:hover,
.add-member-dialog .v-overlay__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.add-member-dialog .v-menu__content .v-list-item--active,
.add-member-dialog .v-overlay__content .v-list-item--active,
.add-member-dialog .v-menu__content .v-list-item[aria-selected="true"],
.add-member-dialog .v-overlay__content .v-list-item[aria-selected="true"] {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

.add-member-dialog .v-list-item-title,
.add-member-dialog .v-list-item-content,
.add-member-dialog .v-list-item__content {
  color: inherit !important;
  background: inherit !important;
}

.add-member-dialog .v-list-item-title *,
.add-member-dialog .v-list-item-content *,
.add-member-dialog .v-list-item__content * {
  color: inherit !important;
  background: inherit !important;
}

/* إصلاح أيقونات القوائم المنسدلة */
.add-member-dialog .v-list-item .v-icon {
  color: inherit !important;
}

.add-member-dialog .v-list-item--active .v-icon,
.add-member-dialog .v-list-item[aria-selected="true"] .v-icon {
  color: white !important;
}

/* تنسيقات شاملة لكل القوائم المنسدلة في الحوار - ألوان سوداء */
.v-dialog:has(.add-member-dialog) .v-overlay__content .v-list,
.v-dialog:has(.add-member-dialog) .v-menu__content .v-list {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.v-dialog:has(.add-member-dialog) .v-overlay__content .v-list-item,
.v-dialog:has(.add-member-dialog) .v-menu__content .v-list-item {
  color: #000000 !important;
  font-weight: 600 !important;
  background: white !important;
}

.v-dialog:has(.add-member-dialog) .v-overlay__content .v-list-item:hover,
.v-dialog:has(.add-member-dialog) .v-menu__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.v-dialog:has(.add-member-dialog) .v-overlay__content .v-list-item--active,
.v-dialog:has(.add-member-dialog) .v-menu__content .v-list-item--active,
.v-dialog:has(.add-member-dialog) .v-overlay__content .v-list-item[aria-selected="true"],
.v-dialog:has(.add-member-dialog) .v-menu__content .v-list-item[aria-selected="true"] {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

/* تنسيق شامل لجميع القوائم المنسدلة */
body .v-overlay__content .v-list-item,
body .v-menu__content .v-list-item {
  color: #000000 !important;
  background: white !important;
}

body .v-overlay__content .v-list-item:hover,
body .v-menu__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

body .v-overlay__content .v-list-item--active,
body .v-menu__content .v-list-item--active,
body .v-overlay__content .v-list-item[aria-selected="true"],
body .v-menu__content .v-list-item[aria-selected="true"] {
  background: #1a1a1a !important;
  color: white !important;
}

/* تنسيق خاص للقوائم داخل الحوار */
.v-dialog__content .v-overlay__content .v-list,
.v-dialog__content .v-menu__content .v-list {
  background: white !important;
  border: 2px solid #1a1a1a !important;
}

.v-dialog__content .v-overlay__content .v-list-item,
.v-dialog__content .v-menu__content .v-list-item {
  color: #000000 !important;
  background: white !important;
}

.v-dialog__content .v-overlay__content .v-list-item:hover,
.v-dialog__content .v-menu__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.v-dialog__content .v-overlay__content .v-list-item--active,
.v-dialog__content .v-menu__content .v-list-item--active {
  background: #1a1a1a !important;
  color: white !important;
}
</style>

<style scoped>
/* تنسيقات قوية جداً للقوائم المنسدلة - ألوان سوداء */
.add-member-dialog :deep(.v-menu__content),
.add-member-dialog :deep(.v-overlay__content) {
  background: white !important;
  border: 2px solid #1a1a1a !important;
  border-radius: 8px !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2) !important;
}

.add-member-dialog :deep(.v-menu__content .v-list),
.add-member-dialog :deep(.v-overlay__content .v-list) {
  background: white !important;
  padding: 4px 0 !important;
}

.add-member-dialog :deep(.v-menu__content .v-list-item),
.add-member-dialog :deep(.v-overlay__content .v-list-item) {
  color: #000000 !important;
  font-weight: 600 !important;
  font-size: 1rem !important;
  padding: 12px 16px !important;
  min-height: 48px !important;
  background: white !important;
}

.add-member-dialog :deep(.v-menu__content .v-list-item:hover),
.add-member-dialog :deep(.v-overlay__content .v-list-item:hover) {
  background: #f5f5f5 !important;
  color: #000000 !important;
  font-weight: 700 !important;
}

.add-member-dialog :deep(.v-menu__content .v-list-item--active),
.add-member-dialog :deep(.v-overlay__content .v-list-item--active),
.add-member-dialog :deep(.v-menu__content .v-list-item[aria-selected="true"]),
.add-member-dialog :deep(.v-overlay__content .v-list-item[aria-selected="true"]) {
  background: #1a1a1a !important;
  color: white !important;
  font-weight: 700 !important;
}

.add-member-dialog :deep(.v-list-item-title),
.add-member-dialog :deep(.v-list-item__title),
.add-member-dialog :deep(.v-list-item-content),
.add-member-dialog :deep(.v-list-item__content) {
  color: inherit !important;
  background: inherit !important;
}

.add-member-dialog :deep(.v-list-item-title *),
.add-member-dialog :deep(.v-list-item__title *),
.add-member-dialog :deep(.v-list-item-content *),
.add-member-dialog :deep(.v-list-item__content *) {
  color: inherit !important;
  background: inherit !important;
}
</style>

<style>
/* تنسيقات شاملة لكل القوائم المنسدلة في الصفحة - ألوان سوداء */
.v-dialog .v-overlay__content .v-list-item,
.v-dialog .v-menu__content .v-list-item {
  color: #000000 !important;
  background: white !important;
}

.v-dialog .v-overlay__content .v-list-item:hover,
.v-dialog .v-menu__content .v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.v-dialog .v-overlay__content .v-list-item--active,
.v-dialog .v-menu__content .v-list-item--active,
.v-dialog .v-overlay__content .v-list-item[aria-selected="true"],
.v-dialog .v-menu__content .v-list-item[aria-selected="true"] {
  background: #1a1a1a !important;
  color: white !important;
}

/* تنسيق شامل لجميع عناصر القائمة */
.v-list-item,
.v-list-item__title,
.v-list-item__content,
.v-list-item-title,
.v-list-item-content {
  color: #000000 !important;
}

.v-list-item:hover {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.v-list-item--active,
.v-list-item[aria-selected="true"] {
  background: #1a1a1a !important;
  color: white !important;
}

/* تنسيق خاص لحقول v-select في الحوار */
.black-dropdown-select :deep(.v-menu__content .v-list-item),
.black-dropdown-select :deep(.v-overlay__content .v-list-item) {
  color: #000000 !important;
  background: white !important;
}

.black-dropdown-select :deep(.v-menu__content .v-list-item:hover),
.black-dropdown-select :deep(.v-overlay__content .v-list-item:hover) {
  background: #f5f5f5 !important;
  color: #000000 !important;
}

.black-dropdown-select :deep(.v-menu__content .v-list-item--active),
.black-dropdown-select :deep(.v-overlay__content .v-list-item--active),
.black-dropdown-select :deep(.v-menu__content .v-list-item[aria-selected="true"]),
.black-dropdown-select :deep(.v-overlay__content .v-list-item[aria-selected="true"]) {
  background: #1a1a1a !important;
  color: white !important;
}

/* تنسيق شامل لجميع التسميات - ألوان سوداء */
.add-member-dialog .v-label,
.add-member-dialog .v-field__label,
.add-member-dialog .v-field__label--active,
.add-member-dialog .v-label--active,
.add-member-dialog .v-label--floating,
.add-member-dialog .v-field__label--floating {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.add-member-dialog .v-field--focused .v-label,
.add-member-dialog .v-field--focused .v-field__label,
.add-member-dialog .v-field--focused .v-label--active,
.add-member-dialog .v-field--focused .v-field__label--active {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.add-member-dialog .v-text-field .v-label,
.add-member-dialog .v-select .v-label,
.add-member-dialog .v-textarea .v-label {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.add-member-dialog .v-text-field .v-field--focused .v-label,
.add-member-dialog .v-select .v-field--focused .v-label,
.add-member-dialog .v-textarea .v-field--focused .v-label {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}
</style>

<style scoped>
/* تنسيقات إضافية للتسميات */
.add-member-dialog :deep(.v-label),
.add-member-dialog :deep(.v-field__label) {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.add-member-dialog :deep(.v-label--active),
.add-member-dialog :deep(.v-field__label--active),
.add-member-dialog :deep(.v-label--floating),
.add-member-dialog :deep(.v-field__label--floating) {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

.add-member-dialog :deep(.v-field--focused .v-label),
.add-member-dialog :deep(.v-field--focused .v-field__label) {
  color: #000000 !important;
  font-weight: 700 !important;
  opacity: 1 !important;
}

/* تنسيق شامل لجميع حدود الحقول - رمادي فاتح */
.add-member-dialog :deep(.v-field__outline) {
  border-color: #d1d5db !important;
}

.add-member-dialog :deep(.v-field__outline__start),
.add-member-dialog :deep(.v-field__outline__notch),
.add-member-dialog :deep(.v-field__outline__end) {
  border-color: #d1d5db !important;
}

.add-member-dialog :deep(.v-field--focused .v-field__outline) {
  border-color: #6b7280 !important;
  border-width: 3px !important;
}

.add-member-dialog :deep(.v-field--focused .v-field__outline__start),
.add-member-dialog :deep(.v-field--focused .v-field__outline__notch),
.add-member-dialog :deep(.v-field--focused .v-field__outline__end) {
  border-color: #6b7280 !important;
}

.add-member-dialog :deep(.v-field:hover .v-field__outline) {
  border-color: #9ca3af !important;
}

.add-member-dialog :deep(.v-text-field .v-field__outline),
.add-member-dialog :deep(.v-select .v-field__outline),
.add-member-dialog :deep(.v-textarea .v-field__outline) {
  border-color: #d1d5db !important;
}

.add-member-dialog :deep(.v-text-field .v-field--focused .v-field__outline),
.add-member-dialog :deep(.v-select .v-field--focused .v-field__outline),
.add-member-dialog :deep(.v-textarea .v-field--focused .v-field__outline) {
  border-color: #6b7280 !important;
  border-width: 3px !important;
}
</style>
