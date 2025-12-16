# CSS Centralization Roadmap

> **Goal:** Replace all hardcoded CSS values with CSS variables so one change updates the entire app.

---

## Current State (Audit Results)

| Property | Hardcoded | Using Variables | Files Affected |
|----------|-----------|-----------------|----------------|
| Colors (`#4338ca` etc.) | 71 | 16 | 9 files |
| Fonts (`font-family`) | 294 | ~89 | 25 files |
| Border-radius (`12px` etc.) | 677 | ~20 | 33 files |
| Box-shadow | 762 | ~10 | 35 files |
| Transitions | 426 | ~22 | 35 files |
| Z-index | 121 | 0 | 17 files |
| RGBA colors | 2,406 | ~8 | 38 files |
| **Total** | **~4,700+** | **~165** | - |

---

## Phase 1: Colors (Priority: High)

### What to Replace

```css
/* BEFORE - Hardcoded */
color: #4338ca;
background: #6366f1;
border-color: #3730a3;

/* AFTER - Centralized */
color: var(--color-primary);
background: var(--color-primary-light);
border-color: var(--color-primary-dark);
```

### Files to Update

- [ ] `src/pages/styles/project-management.css` (13 occurrences)
- [ ] `src/pages/styles/task-management.css` (28 occurrences)
- [ ] `src/pages/styles/human-resources.css` (3 occurrences)
- [ ] `src/pages/styles/team-management.css` (2 occurrences)
- [ ] `src/pages/styles/users.css` (1 occurrence)
- [ ] `src/pages/styles/project-expenses.css` (1 occurrence)
- [ ] `src/components/projects/styles/projects.css` (2 occurrences)
- [ ] `src/styles/indigo-theme.css` (15 occurrences)

### Available Variables (in `src/styles/colors.css`)

```css
--color-primary: #4338ca;
--color-primary-light: #6366f1;
--color-primary-dark: #3730a3;
--color-success: #4caf50;
--color-warning: #ff9800;
--color-error: #f44336;
--color-info: #2196f3;
```

---

## Phase 2: Border Radius (Priority: High)

### What to Replace

```css
/* BEFORE - Hardcoded */
border-radius: 4px;
border-radius: 8px;
border-radius: 12px;
border-radius: 16px;
border-radius: 24px;
border-radius: 50%;

/* AFTER - Centralized */
border-radius: var(--radius-base);
border-radius: var(--radius-lg);
border-radius: var(--radius-xl);
border-radius: var(--radius-2xl);
border-radius: var(--radius-3xl);
border-radius: var(--radius-full);
```

### Mapping Guide

| Hardcoded | Variable |
|-----------|----------|
| `2px` | `var(--radius-sm)` |
| `4px` | `var(--radius-base)` |
| `6px` | `var(--radius-md)` |
| `8px` | `var(--radius-lg)` |
| `12px` | `var(--radius-xl)` |
| `16px` | `var(--radius-2xl)` |
| `24px` | `var(--radius-3xl)` |
| `50%` / `9999px` | `var(--radius-full)` |

### Files to Update (Top Priority)

- [ ] `src/pages/styles/project-management.css` (138 occurrences)
- [ ] `src/pages/styles/categories.css` (55 occurrences)
- [ ] `src/pages/styles/engineers.css` (33 occurrences)
- [ ] `src/pages/styles/expenses.css` (33 occurrences)
- [ ] `src/pages/styles/users.css` (29 occurrences)
- [ ] `src/pages/styles/expense-types.css` (28 occurrences)
- [ ] `src/pages/styles/project-expenses.css` (27 occurrences)

---

## Phase 3: Box Shadows (Priority: High)

### What to Replace

```css
/* BEFORE - Hardcoded */
box-shadow: 0 1px 3px rgba(0,0,0,0.1);
box-shadow: 0 4px 6px rgba(0,0,0,0.1);
box-shadow: 0 10px 15px rgba(0,0,0,0.1);
box-shadow: 0 20px 25px rgba(0,0,0,0.1);

/* AFTER - Centralized */
box-shadow: var(--shadow-sm);
box-shadow: var(--shadow-base);
box-shadow: var(--shadow-md);
box-shadow: var(--shadow-lg);
```

### Available Variables (in `src/styles/spacing.css`)

```css
--shadow-xs: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
--shadow-sm: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
--shadow-base: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
--shadow-md: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
--shadow-lg: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
--shadow-xl: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
```

### Files to Update (Top Priority)

- [ ] `src/pages/styles/project-management.css` (189 occurrences)
- [ ] `src/pages/styles/categories.css` (62 occurrences)
- [ ] `src/pages/styles/users.css` (54 occurrences)
- [ ] `src/pages/styles/engineers.css` (37 occurrences)
- [ ] `src/pages/styles/expenses.css` (30 occurrences)

---

## Phase 4: Fonts (Priority: Medium)

### What to Replace

```css
/* BEFORE - Hardcoded */
font-family: 'Cairo', 'Tajawal', sans-serif;

/* AFTER - Centralized */
font-family: var(--font-primary);
```

### Available Variables (in `src/styles/typography.css`)

```css
--font-primary: 'Cairo', 'Tajawal', 'Noto Sans Arabic', sans-serif;
--font-display: 'Cairo', 'Tajawal', 'Noto Sans Arabic', sans-serif;
```

---

## Phase 5: Spacing (Priority: Low)

### What to Replace

```css
/* BEFORE - Hardcoded */
padding: 16px;
margin: 24px;
gap: 8px;

/* AFTER - Centralized */
padding: var(--space-4);
margin: var(--space-6);
gap: var(--space-2);
```

### Mapping Guide

| Hardcoded | Variable |
|-----------|----------|
| `4px` | `var(--space-1)` |
| `8px` | `var(--space-2)` |
| `12px` | `var(--space-3)` |
| `16px` | `var(--space-4)` |
| `24px` | `var(--space-6)` |
| `32px` | `var(--space-8)` |
| `48px` | `var(--space-12)` |

---

## Execution Strategy

### Option A: File-by-File (Recommended)

1. Start with one file (e.g., `project-management.css`)
2. Replace ALL hardcoded values in that file
3. Test the page
4. Move to next file
5. Repeat

### Option B: Property-by-Property

1. Replace all colors across all files
2. Test
3. Replace all border-radius
4. Test
5. Continue...

---

## Testing Checklist

After each file update:

- [ ] Page loads without errors
- [ ] Colors appear correct
- [ ] Border-radius looks consistent
- [ ] Shadows render properly
- [ ] RTL layout still works
- [ ] Responsive design intact

---

## Benefits After Completion

| Change | Before | After |
|--------|--------|-------|
| Update primary color | Edit 9+ files | Edit 1 line in `colors.css` |
| Update border radius | Edit 33+ files | Edit 1 line in `spacing.css` |
| Update shadows | Edit 35+ files | Edit 1 line in `spacing.css` |
| Theme consistency | Manual checking | Guaranteed |

---

## Notes

- Always use `var(--variable-name)` syntax
- Keep scoped styles for layout-specific tweaks only
- Test RTL after each change
- Run `npm run dev` to verify no build errors

---

---

## Phase 6: Transitions (Priority: Medium)

### What to Replace

```css
/* BEFORE - Hardcoded */
transition: all 0.3s ease;
transition: all 0.2s ease-in-out;
transition: transform 0.3s;

/* AFTER - Centralized */
transition: all var(--transition-normal);
transition: all var(--transition-fast);
transition: transform var(--transition-normal);
```

### Add to `src/styles/spacing.css`

```css
/* Transitions */
--transition-fast: 0.15s ease;
--transition-normal: 0.3s ease;
--transition-slow: 0.5s ease;
--transition-bounce: 0.3s cubic-bezier(0.68, -0.55, 0.265, 1.55);
```

---

## Phase 7: Z-Index (Priority: Medium)

### What to Replace

```css
/* BEFORE - Hardcoded */
z-index: 1;
z-index: 10;
z-index: 100;
z-index: 1000;
z-index: 9999;

/* AFTER - Centralized */
z-index: var(--z-base);
z-index: var(--z-dropdown);
z-index: var(--z-sticky);
z-index: var(--z-modal);
z-index: var(--z-tooltip);
```

### Already in `src/styles/spacing.css`

```css
--z-0: 0;
--z-10: 10;
--z-20: 20;
--z-30: 30;
--z-40: 40;
--z-50: 50;
```

### Recommended: Add Semantic Z-Index

```css
--z-base: 1;
--z-dropdown: 100;
--z-sticky: 200;
--z-fixed: 300;
--z-modal-backdrop: 400;
--z-modal: 500;
--z-popover: 600;
--z-tooltip: 700;
```

---

## Phase 8: Vuetify Theme Integration (Priority: Low)

### Vuetify 3 CSS Variables

Vuetify 3 exposes theme colors as RGB values. Use them like this:

```css
/* Vuetify 3 theme variables */
color: rgb(var(--v-theme-primary));
background: rgba(var(--v-theme-primary), 0.1);
border-color: rgb(var(--v-theme-error));
```

### Sync Custom Variables with Vuetify

In `src/plugins/vuetify.js`:

```javascript
export default createVuetify({
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors: {
          primary: '#4338ca',
          secondary: '#424242',
          success: '#4caf50',
          warning: '#ff9800',
          error: '#f44336',
          info: '#2196f3',
        }
      }
    }
  }
})
```

This ensures Vuetify components use the same colors as your CSS variables.

---

## Phase 9: Vuetify Global Defaults (Priority: Medium) ⚠️ NEW

### Current State

Your `src/plugins/vuetify.js` is minimal and doesn't set global component defaults. This means each Vuetify component uses its own default props, leading to inconsistency.

### What to Add

```javascript
// src/plugins/vuetify.js
export default createVuetify({
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors: {
          primary: '#4338ca',
          secondary: '#424242',
          success: '#4caf50',
          warning: '#ff9800',
          error: '#f44336',
          info: '#2196f3',
        }
      }
    }
  },
  defaults: {
    // Global defaults for ALL Vuetify components
    global: {
      ripple: true,
    },
    // Buttons
    VBtn: {
      elevation: 0,
      rounded: 'lg',
      variant: 'flat',
    },
    // Cards
    VCard: {
      elevation: 2,
      rounded: 'xl',
    },
    // Text Fields
    VTextField: {
      variant: 'outlined',
      density: 'comfortable',
      hideDetails: 'auto',
      rounded: 'lg',
    },
    // Select
    VSelect: {
      variant: 'outlined',
      density: 'comfortable',
      hideDetails: 'auto',
      rounded: 'lg',
    },
    // Data Tables
    VDataTable: {
      hover: true,
      density: 'comfortable',
    },
    // Dialogs
    VDialog: {
      maxWidth: 600,
    },
    // Chips
    VChip: {
      rounded: 'lg',
    },
    // Avatars
    VAvatar: {
      rounded: 'lg',
    },
  }
})
```

### Benefits

| Change | Before | After |
|--------|--------|-------|
| Update button style | Edit every `<v-btn>` in codebase | Edit 1 line in `vuetify.js` |
| Update card elevation | Find all `<v-card>` instances | Edit 1 line in `vuetify.js` |
| Update input variant | Edit every form field | Edit 1 line in `vuetify.js` |

---

## Phase 10: Add Missing CSS Variables (Priority: High) ⚠️ NEW

### Variables to Add to `src/styles/spacing.css`

The roadmap mentions transition variables but they DON'T EXIST yet. Add these:

```css
/* Add to :root in spacing.css */

/* Transitions - MISSING */
--transition-fast: 0.15s ease;
--transition-normal: 0.3s ease;
--transition-slow: 0.5s ease;
--transition-bounce: 0.3s cubic-bezier(0.68, -0.55, 0.265, 1.55);

/* Semantic Z-Index - MISSING */
--z-base: 1;
--z-dropdown: 100;
--z-sticky: 200;
--z-fixed: 300;
--z-modal-backdrop: 400;
--z-modal: 500;
--z-popover: 600;
--z-tooltip: 700;
```

### Why This is High Priority

You CANNOT replace hardcoded transitions/z-index until these variables exist!

---

## Phase 11: RTL-Safe CSS Properties (Priority: HIGH) ⚠️ CRITICAL

### Current Problem

You have **340 occurrences** of RTL-unsafe CSS across 21 files. These properties break when switching text direction.

### What to Replace

```css
/* BEFORE - RTL UNSAFE */
left: 10px;
right: 10px;
margin-left: 16px;
margin-right: 16px;
padding-left: 24px;
padding-right: 24px;
text-align: left;
border-left: 1px solid;

/* AFTER - RTL SAFE (CSS Logical Properties) */
inset-inline-start: 10px;
inset-inline-end: 10px;
margin-inline-start: 16px;
margin-inline-end: 16px;
padding-inline-start: 24px;
padding-inline-end: 24px;
text-align: start;
border-inline-start: 1px solid;
```

### Mapping Guide

| Unsafe Property | Safe Logical Property |
|-----------------|----------------------|
| `left:` | `inset-inline-start:` |
| `right:` | `inset-inline-end:` |
| `margin-left:` | `margin-inline-start:` |
| `margin-right:` | `margin-inline-end:` |
| `padding-left:` | `padding-inline-start:` |
| `padding-right:` | `padding-inline-end:` |
| `border-left:` | `border-inline-start:` |
| `border-right:` | `border-inline-end:` |
| `text-align: left` | `text-align: start` |
| `text-align: right` | `text-align: end` |

### Files to Update (Top Priority)

| File | Occurrences |
|------|-------------|
| `project-management.css` | 83 |
| `categories.css` | 43 |
| `engineers.css` | 33 |
| `expense-types.css` | 22 |
| `work-day-details.css` | 18 |
| `index.css` | 18 |
| `task-management.css` | 16 |
| `materials-expenses.css` | 16 |

### Why This is Critical

Your app is Arabic (RTL). Using `left`/`right` instead of logical properties means:
- Manual overrides needed for every RTL fix
- Inconsistent behavior when switching languages
- More CSS to maintain

---

## Phase 12: Vuetify RTL & Locale Configuration (Priority: HIGH) ⚠️ NEW

### Current State

Your `vuetify.js` is missing RTL and locale configuration:

```javascript
// CURRENT - Missing RTL config
export default createVuetify({
  theme: {
    defaultTheme: 'system',
  },
})
```

### What to Add

```javascript
// src/plugins/vuetify.js
import { createVuetify } from 'vuetify'
import { ar } from 'vuetify/locale'  // Arabic locale

export default createVuetify({
  locale: {
    locale: 'ar',
    fallback: 'en',
    messages: { ar },
    rtl: {
      ar: true,  // Enable RTL for Arabic
    },
  },
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors: {
          primary: '#4338ca',
          secondary: '#424242',
          success: '#4caf50',
          warning: '#ff9800',
          error: '#f44336',
          info: '#2196f3',
        }
      }
    }
  },
  defaults: {
    // ... global defaults from Phase 9
  }
})
```

### Benefits

- Vuetify components automatically flip for RTL
- Built-in Arabic translations for Vuetify components
- `v-locale-provider` component available for mixed-direction content
- `useLocale()` composable for dynamic locale switching

---

## Summary: Full Centralization Checklist

| Phase | Property | Count | Priority | Status |
|-------|----------|-------|----------|--------|
| 1 | Colors | 71 | High | ✅ **DONE** |
| 2 | Border-radius | 677 | High | ✅ **DONE** |
| 3 | Box-shadow | 762 | High | ✅ **DONE** |
| 4 | Fonts | 294 | Medium | ✅ **DONE** |
| 5 | Spacing | 456+ | Low | ✅ **DONE** (709 values centralized) |
| 6 | Transitions | 426 | Medium | ✅ **DONE** (411 centralized, 10 multi-property) |
| 7 | Z-index | 121 | Medium | ✅ **DONE** (117 centralized) |
| 8 | Vuetify Theme | - | Low | ✅ **DONE** (colors.js single source) |
| 9 | Vuetify Global Defaults | - | Medium | ✅ **DONE** (already in vuetify.js) |
| 10 | Add Missing CSS Variables | - | High | ✅ **DONE** |
| 11 | RTL-Safe CSS Properties | 340 | **HIGH** | ✅ **DONE** (265 converted to logical properties) |
| 12 | Vuetify RTL & Locale Config | - | **HIGH** | ✅ **DONE** (already in vuetify.js) |

---

## Recommended Execution Order

### Step 1: Foundation (Do First)
1. **Phase 10** - Add missing CSS variables (required before phases 6-7)
2. **Phase 12** - Vuetify RTL & Locale config (enables proper RTL support)

### Step 2: High Priority CSS
3. **Phase 11** - RTL-Safe CSS Properties (340 replacements - critical for Arabic)
4. **Phase 1** - Colors (most visible impact)
5. **Phase 2** - Border-radius
6. **Phase 3** - Box-shadow

### Step 3: Vuetify Configuration
7. **Phase 8** - Vuetify Theme (sync colors)
8. **Phase 9** - Vuetify Global Defaults (component consistency)

### Step 4: Remaining Properties
9. **Phase 4** - Fonts
10. **Phase 5** - Spacing
11. **Phase 6** - Transitions
12. **Phase 7** - Z-index

---

## Total Counts

| Category | Count |
|----------|-------|
| Hardcoded CSS values | ~4,700+ |
| RTL-unsafe properties | 340 |
| Vuetify config items | 3 (theme, defaults, locale) |
| **Total replacements** | **~5,000+** |

---

**Created:** 2025-12-16
**Updated:** 2025-12-16 (Added Phases 9-12 from Context7 audit)
**Status:** Pending
**Estimated Effort:** ~12-16 hours for full centralization
