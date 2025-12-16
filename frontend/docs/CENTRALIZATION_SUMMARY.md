# CSS Centralization - Summary

## What We're Doing

Replacing scattered, hardcoded styles with centralized variables so one change updates the entire app.

---

## Current Problem

Right now, styles are scattered across 35+ files. To change the primary color, you must edit 9 files manually. To change border radius, you must edit 33 files. This is slow, error-prone, and inconsistent.

---

## After Centralization

All styles come from one place. Change a color once, and every page updates automatically. New pages inherit all styles without extra work.

---

## What Gets Centralized

- Colors (primary, secondary, status)
- Fonts (family and sizes)
- Border radius (rounded corners)
- Shadows (card/button depth)
- Transitions (animation timing)
- Z-index (layer ordering)
- Vuetify components (buttons, cards, tables, inputs)

---

## Benefits

- Consistency across all pages
- Theme changes in seconds
- New pages inherit everything
- Dark mode ready for later
- RTL support works automatically
- Less duplicate code

---

## Scope

Around 5,000 hardcoded values need to be replaced with centralized variables.

---

## Result

Before: Change color means editing many files and hoping nothing breaks.

After: Change one variable and everything updates.

---

See CENTRALIZATION_ROADMAP.md for the detailed execution plan.
