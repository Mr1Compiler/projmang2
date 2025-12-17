/**
 * Mock Data for UI Testing
 *
 * Toggle MOCK_ENABLED to enable/disable all mock data.
 * When backend is ready, delete this entire /mocks folder.
 */

export const MOCK_ENABLED = true

// ============================================
// PROJECTS
// ============================================
export const mockProjects = [
  {
    id: 1,
    name: 'مشروع برج السلام',
    type: 'بناء',
    description: 'بناء برج سكني مكون من 15 طابق',
    location: 'الرياض - حي النخيل',
    startDate: '2025-01-01',
    duration: 24,
    initialCost: 5000000,
    criticalCost: 4500000,
    totalCost: 4800000,
    status: 'active',
    progressPercentage: 35,
    clientPhone: '0501234567',
    notes: 'مشروع ذو أولوية عالية',
    priority: 'high',
    category: 'بناء'
  },
  {
    id: 2,
    name: 'صيانة مجمع الواحة',
    type: 'صيانة',
    description: 'أعمال صيانة شاملة للمجمع التجاري',
    location: 'جدة - حي الروضة',
    startDate: '2025-02-15',
    duration: 6,
    initialCost: 800000,
    criticalCost: 750000,
    totalCost: 720000,
    status: 'active',
    progressPercentage: 60,
    clientPhone: '0559876543',
    notes: '',
    priority: 'medium',
    category: 'صيانة'
  },
  {
    id: 3,
    name: 'تطوير مركز الأعمال',
    type: 'تطوير',
    description: 'تطوير وتحديث مركز الأعمال القائم',
    location: 'الدمام - حي الفيصلية',
    startDate: '2024-11-01',
    duration: 12,
    initialCost: 2500000,
    criticalCost: 2300000,
    totalCost: 2400000,
    status: 'completed',
    progressPercentage: 100,
    clientPhone: '0561112233',
    notes: 'تم التسليم بنجاح',
    priority: 'low',
    category: 'تطوير'
  },
  {
    id: 4,
    name: 'مشروع فيلا الأمير',
    type: 'بناء',
    description: 'بناء فيلا فاخرة مع حديقة ومسبح',
    location: 'الرياض - حي الملقا',
    startDate: '2025-03-01',
    duration: 18,
    initialCost: 3500000,
    criticalCost: 3200000,
    totalCost: 0,
    status: 'pending',
    progressPercentage: 0,
    clientPhone: '0504445566',
    notes: 'في انتظار الموافقات',
    priority: 'urgent',
    category: 'بناء'
  },
  {
    id: 5,
    name: 'استشارات مشروع المطار',
    type: 'استشارات',
    description: 'تقديم استشارات هندسية لتوسعة المطار',
    location: 'جدة - مطار الملك عبدالعزيز',
    startDate: '2024-12-01',
    duration: 8,
    initialCost: 1200000,
    criticalCost: 1100000,
    totalCost: 900000,
    status: 'cancelled',
    progressPercentage: 25,
    clientPhone: '0507778899',
    notes: 'تم إلغاء المشروع بناءً على طلب العميل',
    priority: 'high',
    category: 'استشارات'
  }
]

// ============================================
// USERS
// ============================================
export const mockUsers = [
  {
    id: 1,
    username: 'admin',
    fullName: 'أحمد محمد الأحمد',
    email: 'admin@company.com',
    phone: '0501234567',
    jobTitle: 'مدير النظام',
    status: 'active',
    avatar: null,
    roles: ['admin']
  },
  {
    id: 2,
    username: 'manager1',
    fullName: 'خالد عبدالله السعيد',
    email: 'khaled@company.com',
    phone: '0559998877',
    jobTitle: 'مدير مشاريع',
    status: 'active',
    avatar: null,
    roles: ['manager']
  },
  {
    id: 3,
    username: 'engineer1',
    fullName: 'محمد سعد العتيبي',
    email: 'mohammed@company.com',
    phone: '0561112233',
    jobTitle: 'مهندس مدني',
    status: 'active',
    avatar: null,
    roles: ['engineer']
  },
  {
    id: 4,
    username: 'accountant1',
    fullName: 'سارة أحمد القحطاني',
    email: 'sara@company.com',
    phone: '0504445566',
    jobTitle: 'محاسب',
    status: 'inactive',
    avatar: null,
    roles: ['accountant']
  }
]

// ============================================
// ENGINEERS
// ============================================
export const mockEngineers = [
  {
    id: 1,
    name: 'محمد سعد العتيبي',
    email: 'mohammed@company.com',
    phone: '0561112233',
    specialization: 'هندسة مدنية',
    rating: 4.5,
    skills: ['إدارة المشاريع', 'التصميم الإنشائي', 'AutoCAD'],
    experience: 8,
    avatar: null,
    status: 'active',
    projects: [1, 2]
  },
  {
    id: 2,
    name: 'فهد ناصر الدوسري',
    email: 'fahad@company.com',
    phone: '0507778899',
    specialization: 'هندسة معمارية',
    rating: 4.8,
    skills: ['التصميم المعماري', 'Revit', '3D Modeling'],
    experience: 12,
    avatar: null,
    status: 'active',
    projects: [1, 4]
  },
  {
    id: 3,
    name: 'عبدالرحمن علي الشهري',
    email: 'abdulrahman@company.com',
    phone: '0551234567',
    specialization: 'هندسة كهربائية',
    rating: 4.2,
    skills: ['التمديدات الكهربائية', 'أنظمة التحكم'],
    experience: 5,
    avatar: null,
    status: 'active',
    projects: [2]
  },
  {
    id: 4,
    name: 'يوسف سالم المطيري',
    email: 'yousef@company.com',
    phone: '0569876543',
    specialization: 'هندسة ميكانيكية',
    rating: 3.9,
    skills: ['التكييف والتبريد', 'السباكة'],
    experience: 6,
    avatar: null,
    status: 'inactive',
    projects: []
  }
]

// ============================================
// DEBTORS
// ============================================
export const mockDebtors = [
  {
    id: 1,
    name: 'شركة البناء المتقدم',
    phone: '0112345678',
    email: 'info@advanced-build.com',
    amount: 250000,
    status: 'pending',
    notes: 'دفعة متأخرة من مشروع البرج'
  },
  {
    id: 2,
    name: 'مؤسسة الخليج للتجارة',
    phone: '0119876543',
    email: 'gulf@trade.com',
    amount: 85000,
    status: 'pending',
    notes: 'فاتورة مواد البناء'
  },
  {
    id: 3,
    name: 'محمد عبدالله الغامدي',
    phone: '0551234567',
    email: 'mghamdi@gmail.com',
    amount: 45000,
    status: 'paid',
    notes: 'تم السداد بتاريخ 2025-01-10'
  },
  {
    id: 4,
    name: 'شركة المقاولات الذهبية',
    phone: '0114567890',
    email: 'golden@contractors.com',
    amount: 320000,
    status: 'overdue',
    notes: 'متأخر منذ 3 أشهر - يحتاج متابعة'
  },
  {
    id: 5,
    name: 'مصنع الحديد الوطني',
    phone: '0117778899',
    email: 'national@steel.com',
    amount: 180000,
    status: 'pending',
    notes: 'دفعة ثانية من توريد الحديد'
  }
]

// ============================================
// EXPENSES
// ============================================
export const mockExpenses = [
  {
    id: 1,
    description: 'شراء مواد بناء - اسمنت',
    amount: 75000,
    projectId: 1,
    projectName: 'مشروع برج السلام',
    date: '2025-01-15',
    notes: '500 كيس اسمنت',
    workLocation: 'الموقع الرئيسي',
    expenseType: 'بناء'
  },
  {
    id: 2,
    description: 'أجور عمال - يناير',
    amount: 120000,
    projectId: 1,
    projectName: 'مشروع برج السلام',
    date: '2025-01-30',
    notes: 'رواتب 40 عامل',
    workLocation: 'الموقع الرئيسي',
    expenseType: 'أخرى'
  },
  {
    id: 3,
    description: 'استئجار معدات ثقيلة',
    amount: 45000,
    projectId: 2,
    projectName: 'صيانة مجمع الواحة',
    date: '2025-02-20',
    notes: 'رافعة + حفار لمدة أسبوع',
    workLocation: 'جدة',
    expenseType: 'صيانة'
  },
  {
    id: 4,
    description: 'مواد كهربائية',
    amount: 32000,
    projectId: 2,
    projectName: 'صيانة مجمع الواحة',
    date: '2025-02-25',
    notes: 'أسلاك ومفاتيح ولوحات',
    workLocation: 'جدة',
    expenseType: 'صيانة'
  },
  {
    id: 5,
    description: 'استشارات هندسية',
    amount: 50000,
    projectId: 3,
    projectName: 'تطوير مركز الأعمال',
    date: '2024-12-01',
    notes: 'استشارة تصميم',
    workLocation: 'الدمام',
    expenseType: 'استشارات'
  }
]

// ============================================
// INCOME
// ============================================
export const mockIncome = [
  {
    id: 1,
    description: 'دفعة أولى - مشروع البرج',
    amount: 1500000,
    projectId: 1,
    projectName: 'مشروع برج السلام',
    date: '2025-01-05',
    notes: '30% من قيمة العقد'
  },
  {
    id: 2,
    description: 'دفعة ثانية - مشروع البرج',
    amount: 1000000,
    projectId: 1,
    projectName: 'مشروع برج السلام',
    date: '2025-02-01',
    notes: 'بعد إنجاز 35%'
  },
  {
    id: 3,
    description: 'دفعة كاملة - صيانة الواحة',
    amount: 800000,
    projectId: 2,
    projectName: 'صيانة مجمع الواحة',
    date: '2025-02-15',
    notes: 'مبلغ العقد كاملاً'
  },
  {
    id: 4,
    description: 'تسوية نهائية - مركز الأعمال',
    amount: 500000,
    projectId: 3,
    projectName: 'تطوير مركز الأعمال',
    date: '2024-12-20',
    notes: 'الدفعة الأخيرة'
  }
]

// ============================================
// WORKDAYS
// ============================================
export const mockWorkdays = [
  {
    id: 1,
    projectId: 1,
    date: '2025-01-15',
    description: 'صب الأساسات - المرحلة الأولى',
    notes: 'تم بنجاح',
    status: 'completed',
    materials: [
      { id: 1, name: 'اسمنت', quantity: 100, unitPrice: 25, unit: 'كيس', notes: '' },
      { id: 2, name: 'حديد تسليح', quantity: 5, unitPrice: 3000, unit: 'طن', notes: '' }
    ],
    labor: [
      { id: 1, workerName: 'فريق الصب', hours: 10, hourlyRate: 50, role: 'عمال صب', notes: '15 عامل' }
    ],
    equipment: [
      { id: 1, name: 'خلاطة اسمنت', quantity: 2, dailyRate: 500, notes: '' }
    ]
  },
  {
    id: 2,
    projectId: 1,
    date: '2025-01-16',
    description: 'تركيب الهيكل الحديدي',
    notes: 'جاري العمل',
    status: 'in-progress',
    materials: [
      { id: 3, name: 'حديد هيكلي', quantity: 8, unitPrice: 4000, unit: 'طن', notes: '' }
    ],
    labor: [
      { id: 2, workerName: 'فريق اللحام', hours: 8, hourlyRate: 80, role: 'لحام', notes: '5 عمال' }
    ],
    equipment: [
      { id: 2, name: 'رافعة', quantity: 1, dailyRate: 2000, notes: '' }
    ]
  },
  {
    id: 3,
    projectId: 2,
    date: '2025-02-20',
    description: 'صيانة نظام التكييف',
    notes: '',
    status: 'completed',
    materials: [
      { id: 4, name: 'فلاتر تكييف', quantity: 20, unitPrice: 150, unit: 'قطعة', notes: '' },
      { id: 5, name: 'غاز فريون', quantity: 10, unitPrice: 200, unit: 'اسطوانة', notes: '' }
    ],
    labor: [
      { id: 3, workerName: 'فني تكييف', hours: 6, hourlyRate: 100, role: 'فني', notes: '' }
    ],
    equipment: []
  }
]

// ============================================
// TASKS (Project Tasks)
// ============================================
export const mockTasks = [
  {
    id: 1,
    title: 'مراجعة المخططات الهندسية',
    description: 'مراجعة واعتماد المخططات النهائية',
    status: 'completed',
    dueDate: '2025-01-10',
    projectId: 1
  },
  {
    id: 2,
    title: 'طلب تصاريح البناء',
    description: 'تقديم طلب التصاريح للبلدية',
    status: 'completed',
    dueDate: '2025-01-15',
    projectId: 1
  },
  {
    id: 3,
    title: 'توريد مواد البناء',
    description: 'التنسيق مع الموردين لتوريد الاسمنت والحديد',
    status: 'in-progress',
    dueDate: '2025-02-01',
    projectId: 1
  },
  {
    id: 4,
    title: 'فحص نظام الكهرباء',
    description: 'فحص شامل للتمديدات الكهربائية',
    status: 'pending',
    dueDate: '2025-02-28',
    projectId: 2
  },
  {
    id: 5,
    title: 'تسليم التقرير النهائي',
    description: 'إعداد وتسليم التقرير النهائي للعميل',
    status: 'cancelled',
    dueDate: '2025-01-20',
    projectId: 5
  }
]

// ============================================
// EMPLOYEES (HR)
// ============================================
export const mockEmployees = [
  {
    id: 1,
    name: 'أحمد محمد الأحمد',
    department: 'الإدارة',
    position: 'مدير عام',
    status: 'نشط',
    salary: 25000,
    email: 'ahmed@company.com',
    phone: '0501234567',
    leaves: [
      { id: 1, startDate: '2025-01-01', endDate: '2025-01-03', days: 3, reason: 'إجازة سنوية', status: 'approved' }
    ],
    attendance: [
      { id: 1, date: '2025-01-15', checkIn: '08:00', checkOut: '17:00', hours: 9, notes: '' }
    ],
    evaluations: [
      { id: 1, date: '2024-12-15', score: 95, feedback: 'أداء ممتاز', evaluator: 'مجلس الإدارة' }
    ],
    skills: ['إدارة المشاريع', 'القيادة', 'التخطيط الاستراتيجي'],
    certificates: [
      { id: 1, name: 'PMP', issueDate: '2020-05-01', expiryDate: '2026-05-01', number: 'PMP-123456' }
    ],
    salaryHistory: [
      { id: 1, date: '2024-01-01', amount: 22000, notes: 'زيادة سنوية' },
      { id: 2, date: '2025-01-01', amount: 25000, notes: 'ترقية' }
    ],
    fingerprint: true,
    fingerprintDate: '2024-01-15'
  },
  {
    id: 2,
    name: 'خالد عبدالله السعيد',
    department: 'المشاريع',
    position: 'مدير مشاريع',
    status: 'نشط',
    salary: 18000,
    email: 'khaled@company.com',
    phone: '0559998877',
    leaves: [],
    attendance: [
      { id: 2, date: '2025-01-15', checkIn: '07:30', checkOut: '18:00', hours: 10.5, notes: 'عمل إضافي' }
    ],
    evaluations: [
      { id: 2, date: '2024-12-15', score: 88, feedback: 'أداء جيد جداً', evaluator: 'المدير العام' }
    ],
    skills: ['إدارة المشاريع', 'AutoCAD', 'MS Project'],
    certificates: [],
    salaryHistory: [],
    fingerprint: true,
    fingerprintDate: '2024-02-01'
  },
  {
    id: 3,
    name: 'سارة أحمد القحطاني',
    department: 'المالية',
    position: 'محاسب',
    status: 'في إجازة',
    salary: 12000,
    email: 'sara@company.com',
    phone: '0504445566',
    leaves: [
      { id: 2, startDate: '2025-01-10', endDate: '2025-01-20', days: 10, reason: 'إجازة مرضية', status: 'approved' }
    ],
    attendance: [],
    evaluations: [
      { id: 3, date: '2024-12-15', score: 82, feedback: 'أداء جيد', evaluator: 'المدير المالي' }
    ],
    skills: ['المحاسبة', 'Excel', 'التقارير المالية'],
    certificates: [
      { id: 2, name: 'CPA', issueDate: '2022-03-01', expiryDate: null, number: 'CPA-789012' }
    ],
    salaryHistory: [],
    fingerprint: true,
    fingerprintDate: '2024-03-01'
  },
  {
    id: 4,
    name: 'محمد سعد العتيبي',
    department: 'الهندسة',
    position: 'مهندس مدني',
    status: 'نشط',
    salary: 15000,
    email: 'mohammed@company.com',
    phone: '0561112233',
    leaves: [],
    attendance: [
      { id: 3, date: '2025-01-15', checkIn: '08:00', checkOut: '16:00', hours: 8, notes: '' }
    ],
    evaluations: [],
    skills: ['التصميم الإنشائي', 'AutoCAD', 'Revit'],
    certificates: [],
    salaryHistory: [],
    fingerprint: false,
    fingerprintDate: null
  }
]

// ============================================
// TEAM MEMBERS (Project Assignments)
// ============================================
export const mockTeamMembers = [
  { id: 1, userId: 2, projectId: 1, role: 'مدير المشروع' },
  { id: 2, userId: 3, projectId: 1, role: 'مهندس موقع' },
  { id: 3, userId: 2, projectId: 2, role: 'مدير المشروع' },
  { id: 4, userId: 3, projectId: 2, role: 'مهندس موقع' }
]

// ============================================
// WORK CATEGORIES
// ============================================
export const mockWorkCategories = [
  {
    id: 1,
    name: 'أعمال الهيكل الإنشائي',
    description: 'جميع أعمال البناء الإنشائية',
    subcategories: [
      { id: 1, name: 'صب الخرسانة', description: '' },
      { id: 2, name: 'تركيب الحديد', description: '' },
      { id: 3, name: 'أعمال الطوب', description: '' }
    ]
  },
  {
    id: 2,
    name: 'أعمال التشطيبات',
    description: 'أعمال التشطيب الداخلي والخارجي',
    subcategories: [
      { id: 4, name: 'الدهانات', description: '' },
      { id: 5, name: 'السيراميك والبلاط', description: '' },
      { id: 6, name: 'الجبس والديكور', description: '' }
    ]
  },
  {
    id: 3,
    name: 'الأعمال الكهربائية',
    description: 'التمديدات والتركيبات الكهربائية',
    subcategories: [
      { id: 7, name: 'التمديدات الرئيسية', description: '' },
      { id: 8, name: 'الإنارة', description: '' },
      { id: 9, name: 'أنظمة التحكم', description: '' }
    ]
  },
  {
    id: 4,
    name: 'الأعمال الميكانيكية',
    description: 'أعمال التكييف والسباكة',
    subcategories: [
      { id: 10, name: 'التكييف والتبريد', description: '' },
      { id: 11, name: 'السباكة', description: '' },
      { id: 12, name: 'مكافحة الحريق', description: '' }
    ]
  }
]

// ============================================
// ROLES & PERMISSIONS
// ============================================
export const mockRoles = [
  {
    id: 1,
    name: 'admin',
    description: 'مدير النظام - صلاحيات كاملة'
  },
  {
    id: 2,
    name: 'manager',
    description: 'مدير مشاريع - صلاحيات إدارية'
  },
  {
    id: 3,
    name: 'engineer',
    description: 'مهندس - صلاحيات محدودة'
  },
  {
    id: 4,
    name: 'accountant',
    description: 'محاسب - صلاحيات مالية'
  }
]

export const mockPages = [
  { id: 1, name: 'لوحة التحكم', route: '/dashboard', icon: 'mdi-view-dashboard', order: 1 },
  { id: 2, name: 'المشاريع', route: '/projects', icon: 'mdi-folder-multiple', order: 2 },
  { id: 3, name: 'المصروفات', route: '/expenses', icon: 'mdi-cash-minus', order: 3 },
  { id: 4, name: 'الإيرادات', route: '/income', icon: 'mdi-cash-plus', order: 4 },
  { id: 5, name: 'المدينون', route: '/debtors', icon: 'mdi-account-cash', order: 5 },
  { id: 6, name: 'المهندسون', route: '/engineers', icon: 'mdi-hard-hat', order: 6 },
  { id: 7, name: 'المستخدمون', route: '/users', icon: 'mdi-account-group', order: 7 },
  { id: 8, name: 'الموارد البشرية', route: '/human-resources', icon: 'mdi-badge-account', order: 8 },
  { id: 9, name: 'التصنيفات', route: '/categories', icon: 'mdi-tag-multiple', order: 9 }
]

// Admin has all permissions
export const mockRolePages = [
  // Admin - full access to all pages
  ...mockPages.map((page, index) => ({
    id: index + 1,
    roleId: 1,
    pageId: page.id,
    canRead: true,
    canWrite: true,
    canDelete: true
  })),
  // Manager - read/write most, no delete users
  { id: 10, roleId: 2, pageId: 1, canRead: true, canWrite: true, canDelete: false },
  { id: 11, roleId: 2, pageId: 2, canRead: true, canWrite: true, canDelete: true },
  { id: 12, roleId: 2, pageId: 3, canRead: true, canWrite: true, canDelete: false },
  { id: 13, roleId: 2, pageId: 5, canRead: true, canWrite: true, canDelete: false },
  { id: 14, roleId: 2, pageId: 6, canRead: true, canWrite: true, canDelete: false },
  // Engineer - limited access
  { id: 15, roleId: 3, pageId: 1, canRead: true, canWrite: false, canDelete: false },
  { id: 16, roleId: 3, pageId: 2, canRead: true, canWrite: true, canDelete: false },
  // Accountant - financial pages
  { id: 17, roleId: 4, pageId: 1, canRead: true, canWrite: false, canDelete: false },
  { id: 18, roleId: 4, pageId: 3, canRead: true, canWrite: true, canDelete: false },
  { id: 19, roleId: 4, pageId: 4, canRead: true, canWrite: true, canDelete: false },
  { id: 20, roleId: 4, pageId: 5, canRead: true, canWrite: true, canDelete: false }
]

// ============================================
// HELPER: Get all mock data as a single object
// ============================================
export const getAllMockData = () => ({
  projects: mockProjects,
  users: mockUsers,
  engineers: mockEngineers,
  debtors: mockDebtors,
  expenses: mockExpenses,
  income: mockIncome,
  workdays: mockWorkdays,
  tasks: mockTasks,
  employees: mockEmployees,
  teamMembers: mockTeamMembers,
  workCategories: mockWorkCategories,
  roles: mockRoles,
  pages: mockPages,
  rolePages: mockRolePages
})
