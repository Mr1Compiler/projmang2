# UI/UX Development Guide

> **For AI Agents & UI/UX Developers**
> This guide ensures UI changes are isolated from business logic.

---

## CRITICAL RULES

### NEVER MODIFY These Files/Folders

```
src/api/           # Backend communication - DO NOT TOUCH
src/stores/        # State management - DO NOT TOUCH
src/composables/   # Business logic - DO NOT TOUCH
src/router/        # Routing logic - DO NOT TOUCH
src/utils/         # Utility functions - DO NOT TOUCH
src/config/        # Configuration - DO NOT TOUCH
```

### NEVER Change

- `<script setup>` logic in components (imports, functions, computed, refs)
- Props definitions (`defineProps`)
- Emits definitions (`defineEmits`)
- API calls or store interactions
- Event handlers logic (what they do, not how they look)
- Route definitions
- Any `.js` files outside styles

---

## SAFE TO MODIFY

### 1. Global Styles (Theme System)

Location: `src/styles/`

| File | Purpose | Safe Changes |
|------|---------|--------------|
| `colors.css` | Color palette & CSS variables | All color values, gradients |
| `typography.css` | Fonts, sizes, weights | Font families, sizes, weights |
| `spacing.css` | Margins, padding, gaps | Spacing values, border-radius |
| `visual-effects.css` | Animations, transitions | All effects, animations |
| `main.css` | Global styles, Vuetify overrides | Vuetify component styling |
| `indigo-theme.css` | Theme-specific colors | Theme colors |
| `form-dialog.css` | Dialog/modal styling | Form appearance |
| `table-headers.css` | Table styling | Table appearance |
| `rtl.css` | RTL-specific styles | RTL adjustments |

### 2. Page-Specific Styles

Location: `src/pages/styles/`

| File | Affects Page |
|------|--------------|
| `project-management.css` | Project management page |
| `engineers.css` | Engineers page |
| `expenses.css` | Expenses page |
| `debtors.css` | Debtors page |
| `human-resources.css` | HR page |
| `users.css` | Users page |
| `work-days.css` | Work days page |
| `income.css` | Income page |
| `categories.css` | Categories page |
| `inventory.css` | Inventory page |
| `sales.css` | Sales page |
| `purchases.css` | Purchases page |
| `project-details.css` | Project details page |
| `project-expenses.css` | Project expenses page |
| `work-day-details.css` | Work day details page |
| `labor-details.css` | Labor details page |
| `equipment-details.css` | Equipment details page |
| `materials-expenses.css` | Materials expenses page |
| `task-management.css` | Task management page |
| `team-management.css` | Team management page |
| `expense-types.css` | Expense types page |

### 3. Component-Specific Styles

Location: `src/components/[feature]/styles/`

| File | Affects Components |
|------|-------------------|
| `components/projects/styles/projects.css` | All project components |
| `components/engineers/styles/engineers.css` | All engineer components |
| `components/expenses/styles/expenses.css` | All expense components |
| `components/debtors/styles/debtors.css` | All debtor components |
| `components/human-resources/styles/human-resources.css` | All HR components |
| `components/users/styles/users.css` | All user components |

### 4. Template HTML/Classes (in `.vue` files)

**SAFE to modify in `<template>` section:**
- CSS class names
- Vuetify component props (color, size, variant, elevation)
- Icon names (`mdi-*`)
- Style bindings (`:style`, `:class`)
- Layout structure (v-row, v-col arrangements)

**Example - SAFE changes:**
```vue
<!-- BEFORE -->
<v-btn color="primary" size="default">Save</v-btn>

<!-- AFTER - Safe to change -->
<v-btn color="indigo" size="large" rounded="lg" elevation="4">Save</v-btn>
```

### 5. Scoped Styles in Components

**SAFE to modify `<style scoped>` sections in:**

```
src/components/projects/*.vue
src/components/engineers/*.vue
src/components/expenses/*.vue
src/components/debtors/*.vue
src/components/human-resources/*.vue
src/components/users/*.vue
src/components/AppFooter.vue
src/components/ModernDialog.vue
src/components/SimpleDialog.vue
```

---

## CSS Variables Reference

### Colors (defined in `src/styles/colors.css`)

```css
/* Primary Colors - Indigo Theme */
--color-primary: #4338ca;
--color-primary-light: #6366f1;
--color-primary-dark: #3730a3;

/* Status Colors */
--color-success: #4caf50;
--color-warning: #ff9800;
--color-error: #f44336;
--color-info: #2196f3;

/* Text Colors */
--text-primary: #212121;
--text-secondary: #757575;
--text-tertiary: #9e9e9e;
--text-white: #ffffff;

/* Background Colors */
--background-primary: #ffffff;
--background-secondary: #f8fafc;
--background-tertiary: #f1f5f9;
--background-indigo-light: #eef2ff;

/* Gradients */
--gradient-primary: linear-gradient(135deg, #4338ca 0%, #6366f1 100%);
--gradient-success: linear-gradient(135deg, #10b981 0%, #059669 100%);
--gradient-warning: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
--gradient-error: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
```

### Typography (defined in `src/styles/typography.css`)

```css
/* Font Families */
--font-primary: 'Cairo', 'Tajawal', 'Noto Sans Arabic', sans-serif;
--font-display: 'Cairo', 'Tajawal', 'Noto Sans Arabic', sans-serif;

/* Font Sizes */
--font-size-xs: 0.75rem;    /* 12px */
--font-size-sm: 0.875rem;   /* 14px */
--font-size-base: 1rem;     /* 16px */
--font-size-lg: 1.125rem;   /* 18px */
--font-size-xl: 1.25rem;    /* 20px */
--font-size-2xl: 1.5rem;    /* 24px */
--font-size-3xl: 1.875rem;  /* 30px */
--font-size-4xl: 2.25rem;   /* 36px */

/* Font Weights */
--font-weight-normal: 400;
--font-weight-medium: 500;
--font-weight-semibold: 600;
--font-weight-bold: 700;
```

### Spacing (defined in `src/styles/spacing.css`)

```css
/* Spacing Scale */
--space-1: 0.25rem;   /* 4px */
--space-2: 0.5rem;    /* 8px */
--space-3: 0.75rem;   /* 12px */
--space-4: 1rem;      /* 16px */
--space-6: 1.5rem;    /* 24px */
--space-8: 2rem;      /* 32px */
--space-12: 3rem;     /* 48px */

/* Border Radius */
--radius-sm: 0.125rem;   /* 2px */
--radius-md: 0.375rem;   /* 6px */
--radius-lg: 0.5rem;     /* 8px */
--radius-xl: 0.75rem;    /* 12px */
--radius-2xl: 1rem;      /* 16px */
--radius-full: 9999px;

/* Shadows */
--shadow-sm: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
--shadow-md: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
--shadow-lg: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
--shadow-xl: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
```

### Visual Effects (defined in `src/styles/visual-effects.css`)

```css
/* Available Effect Classes */
.smooth-transition      /* Standard transition */
.hover-lift            /* Lift on hover */
.hover-glow            /* Glow on hover */
.light-sweep           /* Light sweep animation */
.card-glow             /* Card glow effect */
.glass-effect          /* Glassmorphism */
.gradient-animation    /* Animated gradient */
.fade-in               /* Fade in animation */
.icon-spin             /* Spin icon on hover */
.icon-bounce           /* Bouncing icon */
.icon-glow             /* Glowing icon */
.btn-glow              /* Button glow effect */
.text-glow             /* Text glow effect */
```

---

## How to Change Themes

### Change Primary Color Theme

Edit `src/styles/colors.css`:

```css
:root {
  /* Change these values for new theme */
  --color-primary: #YOUR_COLOR;
  --color-primary-light: #YOUR_LIGHTER;
  --color-primary-dark: #YOUR_DARKER;

  /* Update gradient */
  --gradient-primary: linear-gradient(135deg, #YOUR_COLOR 0%, #YOUR_LIGHTER 100%);
}
```

### Change Fonts

Edit `src/styles/typography.css`:

```css
/* Add your font import */
@import url('https://fonts.googleapis.com/css2?family=YOUR_FONT&display=swap');

:root {
  --font-primary: 'YOUR_FONT', 'Cairo', sans-serif;
  --font-display: 'YOUR_FONT', 'Cairo', sans-serif;
}
```

### Change Component Appearance

Edit the specific CSS file or scoped styles:

```css
/* Example: Change all cards */
.v-card {
  border-radius: var(--radius-2xl) !important;
  box-shadow: var(--shadow-lg) !important;
}

/* Example: Change all buttons */
.v-btn {
  border-radius: var(--radius-lg) !important;
  font-weight: var(--font-weight-semibold) !important;
}
```

---

## Vuetify Props Reference (Safe to Change)

### Colors
```vue
color="primary|secondary|success|warning|error|info|indigo|..."
```

### Sizes
```vue
size="x-small|small|default|large|x-large"
```

### Variants
```vue
variant="flat|text|elevated|tonal|outlined|plain"
```

### Density
```vue
density="default|comfortable|compact"
```

### Rounded
```vue
rounded="0|sm|md|lg|xl|pill|circle"
```

### Elevation
```vue
elevation="0|1|2|3|4|...|24"
```

---

## Component Structure Overview

```
src/components/
├── projects/
│   ├── ProjectStats.vue      # Statistics cards
│   ├── ProjectCard.vue       # Project card display
│   ├── ProjectFilters.vue    # Search/filter UI
│   ├── ProjectForm.vue       # Add/Edit dialog
│   ├── TeamManagement.vue    # Team section
│   ├── DeleteConfirmDialog.vue  # Delete confirmation
│   └── styles/projects.css   # Shared styles
│
├── engineers/
│   ├── EngineerStats.vue
│   ├── EngineerTable.vue
│   ├── EngineerForm.vue
│   ├── EngineerDetails.vue
│   ├── EngineerProjects.vue
│   └── styles/engineers.css
│
├── expenses/
│   ├── ExpenseStats.vue
│   ├── ExpenseFilters.vue
│   ├── ExpenseTable.vue
│   ├── ExpenseForm.vue
│   └── styles/expenses.css
│
├── debtors/
│   ├── DebtorStats.vue
│   ├── DebtorFilters.vue
│   ├── DebtorTable.vue
│   ├── DebtorForm.vue
│   ├── DebtorDetails.vue
│   ├── DebtorPayments.vue
│   └── styles/debtors.css
│
├── human-resources/
│   ├── HRStats.vue
│   ├── HRFilters.vue
│   ├── HRTable.vue
│   ├── EmployeeForm.vue
│   ├── EmployeeDetails.vue
│   ├── EmployeeSubDialogs.vue
│   └── styles/human-resources.css
│
└── users/
    ├── UserStats.vue
    ├── UserFilters.vue
    ├── UserTable.vue
    ├── UserForm.vue
    ├── UserDetails.vue
    ├── UserDialogs.vue
    └── styles/users.css
```

---

## Example: Safe UI Modification

### Task: Change button style across the app

**DO THIS:**

Edit `src/styles/main.css`:
```css
.v-btn {
  border-radius: 12px !important;
  font-weight: 600 !important;
  text-transform: none !important;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1) !important;
}

.v-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15) !important;
}
```

**DON'T DO THIS:**

```javascript
// NEVER modify button click handlers
const handleClick = () => {
  // This is LOGIC - don't touch!
  store.saveData()
}
```

---

## Example: Change a Component's Look

### Task: Restyle the stats cards

**SAFE - Edit template classes:**

```vue
<!-- In ProjectStats.vue template -->
<v-card class="stat-card pa-6 text-center hover-lift card-glow"
        elevation="4"
        rounded="xl">
```

**SAFE - Edit scoped styles:**

```vue
<style scoped>
.stat-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8fafc 100%);
  border: 2px solid var(--border-indigo-light);
}
</style>
```

**NOT SAFE - Don't modify script:**

```vue
<script setup>
// DON'T TOUCH ANY OF THIS
const props = defineProps({...})
const formattedBudget = computed(() => {...})
</script>
```

---

## RTL Support Notes

This project uses Arabic (RTL) by default. When modifying styles:

1. Use logical properties when possible:
   - `margin-inline-start` instead of `margin-left`
   - `padding-inline-end` instead of `padding-right`

2. Test all changes in RTL mode

3. Direction-specific styles are in `src/styles/rtl.css`

---

## Vue Scoped Styles Best Practices

Based on official Vue documentation:

### How Scoped Styles Work

When using `<style scoped>`, Vue adds a unique attribute (like `data-v-7ba5bd90`) to elements:

```css
/* Your code */
.stat-card { color: red; }

/* Compiled output */
.stat-card[data-v-7ba5bd90] { color: red; }
```

### Deep Selectors (Style Child Components)

To style elements inside child components, use `:deep()`:

```css
/* Style child component elements */
:deep(.v-card-title) {
  font-size: 1.5rem;
}

/* Style v-html content */
:deep(.html-content p) {
  margin-bottom: 1rem;
}
```

### Global Styles Within Scoped Block

Use `:global()` for single global rules without creating new style block:

```css
<style scoped>
:global(.some-global-class) {
  color: red;
}
</style>
```

### Performance Tip

Use classes instead of element selectors in scoped styles:

```css
/* SLOW - avoid */
p { color: red; }

/* FAST - preferred */
.paragraph { color: red; }
```

---

## Vuetify SASS Variables (Advanced)

For deeper Vuetify customization, edit `src/styles/settings.scss`:

```scss
// Override Vuetify defaults
@use 'vuetify/settings' with (
  $button-height: 44px,
  $card-border-radius: 16px,
  $button-border-radius: 8px
);
```

**Important:** Only use `@use 'vuetify/settings'` - never `@use 'vuetify/styles'` (it's precompiled CSS).

---

## Pinia Store Rules (DO NOT MODIFY)

The project uses Pinia stores with `storeToRefs()` for reactivity. This is LOGIC - never modify:

```javascript
// This pattern is in pages - DO NOT CHANGE
const store = useProjectsStore()
const { projects, loading } = storeToRefs(store)  // Keeps reactivity
const { fetchProjects } = store                    // Actions don't need storeToRefs
```

**Why this matters:** Destructuring without `storeToRefs()` breaks reactivity. The logic developer set this up correctly - don't touch it.

---

## Quick Checklist Before Making Changes

- [ ] Am I only modifying CSS/styles?
- [ ] Am I NOT touching any `.js` files in api/, stores/, composables/, utils/?
- [ ] Am I NOT changing any function logic in `<script setup>`?
- [ ] Am I NOT modifying props or emits definitions?
- [ ] Am I using CSS variables from the design system?
- [ ] Will this change work with RTL layout?

---

## File Modification Summary

| Action | Safe Files | Forbidden Files |
|--------|-----------|-----------------|
| Change colors | `styles/colors.css` | Any `.js` file |
| Change fonts | `styles/typography.css` | `api/*.js` |
| Change spacing | `styles/spacing.css` | `stores/*.js` |
| Change animations | `styles/visual-effects.css` | `composables/*.js` |
| Change component look | `components/**/styles/*.css` | `utils/*.js` |
| Change page look | `pages/styles/*.css` | `router/*.js` |
| Change Vuetify theme | `plugins/vuetify.js` | `config/*.js` |
| Change template classes | `*.vue` `<template>` section | `*.vue` `<script>` section |

---

## Contact for Logic Changes

If a UI change requires logic modifications (new props, new events, etc.), escalate to the backend/logic developer. Do NOT attempt to modify:

- Store actions/getters
- API endpoints
- Composable functions
- Route guards
- Form validation logic
- Data transformations

---

**Last Updated:** 2025-12-16
**Project:** Project Management System
**Framework:** Vue 3 + Vuetify 3 + Pinia
**Verified Against:** Vue 3 Official Docs, Vuetify 3 Docs, Pinia Docs (via Context7)
