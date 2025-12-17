/**
 * Formatting utilities for the application
 * Centralizes all formatting functions to ensure consistency
 */

/**
 * Format a number as Iraqi Dinar currency
 * @param {number} amount - The amount to format
 * @param {Object} options - Formatting options
 * @param {string} [options.locale='en-US'] - Locale for formatting (changed to EN for english digits)
 * @param {boolean} [options.showSymbol=true] - Whether to show currency symbol
 * @returns {string} Formatted currency string
 */
export function formatCurrency(amount, options = {}) {
  const { locale = 'en-US', showSymbol = true } = options

  if (amount === null || amount === undefined) {
    return showSymbol ? '0 /.9' : '0'
  }

  const formatted = new Intl.NumberFormat(locale, {
    style: 'decimal',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(amount)

  return showSymbol ? `${formatted} د.ع` : formatted
}

/**
 * Format a number as currency with full Intl support
 * @param {number} amount - The amount to format
 * @param {string} [currency='IQD'] - Currency code
 * @param {string} [locale='en-US'] - Locale for formatting
 * @returns {string} Formatted currency string
 */
export function formatCurrencyFull(amount, currency = 'IQD', locale = 'en-US') {
  if (amount === null || amount === undefined) {
    return new Intl.NumberFormat(locale, {
      style: 'currency',
      currency,
      minimumFractionDigits: 0,
    }).format(0)
  }

  return new Intl.NumberFormat(locale, {
    style: 'currency',
    currency,
    minimumFractionDigits: 0,
  }).format(amount)
}

/**
 * Format amount with simple locale string (shorthand)
 * @param {number} amount - The amount to format
 * @returns {string} Formatted amount with /.9 suffix
 */
export function formatAmount(amount) {
  if (amount === null || amount === undefined) {
    return '0 /.9'
  }
  return `${amount.toLocaleString('en-US')} د.ع`
}

/**
 * Format a date for display
 * @param {string|Date} date - The date to format
 * @param {Object} options - Formatting options
 * @param {string} [options.locale='en-US'] - Locale for formatting
 * @param {string} [options.format='short'] - Format type: 'short', 'long', 'full'
 * @returns {string} Formatted date string
 */
export function formatDate(date, options = {}) {
  const { locale = 'en-US', format = 'short' } = options

  if (!date) {
    return 'تاريخ غير محدد'
  }

  const dateObj = typeof date === 'string' ? new Date(date) : date

  if (isNaN(dateObj.getTime())) {
    return 'تاريخ غير صالح'
  }

  const formatOptions = {
    short: { year: 'numeric', month: '2-digit', day: '2-digit' },
    long: { year: 'numeric', month: 'long', day: 'numeric' },
    full: { year: 'numeric', month: 'long', day: 'numeric', weekday: 'long' },
  }

  return dateObj.toLocaleDateString(locale, formatOptions[format] || formatOptions.short)
}

/**
 * Format a date for HTML input fields (YYYY-MM-DD)
 * @param {string|Date} date - The date to format
 * @returns {string} Date string in YYYY-MM-DD format
 */
export function formatDateForInput(date) {
  if (!date) return ''

  const dateObj = typeof date === 'string' ? new Date(date) : date

  if (isNaN(dateObj.getTime())) {
    return ''
  }

  return dateObj.toISOString().split('T')[0]
}

/**
 * Format a date and time for display
 * @param {string|Date} date - The date to format
 * @param {string} [locale='en-US'] - Locale for formatting
 * @returns {string} Formatted date and time string
 */
export function formatDateTime(date, locale = 'en-US') {
  if (!date) return 'تاريخ غير محدد'

  const dateObj = typeof date === 'string' ? new Date(date) : date

  if (isNaN(dateObj.getTime())) {
    return 'تاريخ غير صالح'
  }

  return dateObj.toLocaleString(locale, {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  })
}

/**
 * Format a time string
 * @param {string} time - Time string (HH:MM or HH:MM:SS)
 * @param {boolean} [showSeconds=false] - Whether to show seconds
 * @returns {string} Formatted time string
 */
export function formatTime(time, showSeconds = false) {
  if (!time) return ''

  const parts = time.split(':')
  if (parts.length < 2) return time

  return showSeconds && parts.length >= 3
    ? `${parts[0]}:${parts[1]}:${parts[2]}`
    : `${parts[0]}:${parts[1]}`
}

/**
 * Format a number with thousands separator
 * @param {number} number - The number to format
 * @param {string} [locale='en-US'] - Locale for formatting
 * @returns {string} Formatted number string
 */
export function formatNumber(number, locale = 'en-US') {
  if (number === null || number === undefined) {
    return '0'
  }

  return new Intl.NumberFormat(locale).format(number)
}

/**
 * Format a percentage
 * @param {number} value - The value to format (0-100 or 0-1)
 * @param {Object} options - Formatting options
 * @param {boolean} [options.isDecimal=false] - Whether value is decimal (0-1) or percentage (0-100)
 * @param {number} [options.decimals=0] - Number of decimal places
 * @returns {string} Formatted percentage string
 */
export function formatPercentage(value, options = {}) {
  const { isDecimal = false, decimals = 0 } = options

  if (value === null || value === undefined) {
    return '0%'
  }

  const percentage = isDecimal ? value * 100 : value
  return `${percentage.toFixed(decimals)}%`
}

/**
 * Format a phone number (Iraqi format)
 * @param {string} phone - Phone number to format
 * @returns {string} Formatted phone number
 */
export function formatPhone(phone) {
  if (!phone) return ''

  // Remove non-digit characters
  const digits = phone.replace(/\D/g, '')

  // Iraqi phone format: 07XX XXX XXXX
  if (digits.length === 11 && digits.startsWith('07')) {
    return `${digits.slice(0, 4)} ${digits.slice(4, 7)} ${digits.slice(7)}`
  }

  return phone
}

/**
 * Format file size
 * @param {number} bytes - File size in bytes
 * @returns {string} Formatted file size
 */
export function formatFileSize(bytes) {
  if (bytes === 0) return '0 Bytes'

  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`
}

/**
 * Format relative time (e.g., "2 hours ago")
 * @param {string|Date} date - The date to format
 * @param {string} [locale='ar'] - Locale for formatting
 * @returns {string} Relative time string
 */
export function formatRelativeTime(date, locale = 'ar') {
  if (!date) return ''

  const dateObj = typeof date === 'string' ? new Date(date) : date
  const now = new Date()
  const diffInSeconds = Math.floor((now - dateObj) / 1000)

  const rtf = new Intl.RelativeTimeFormat(locale, { numeric: 'auto' })

  if (diffInSeconds < 60) {
    return rtf.format(-diffInSeconds, 'second')
  } else if (diffInSeconds < 3600) {
    return rtf.format(-Math.floor(diffInSeconds / 60), 'minute')
  } else if (diffInSeconds < 86400) {
    return rtf.format(-Math.floor(diffInSeconds / 3600), 'hour')
  } else if (diffInSeconds < 2592000) {
    return rtf.format(-Math.floor(diffInSeconds / 86400), 'day')
  } else if (diffInSeconds < 31536000) {
    return rtf.format(-Math.floor(diffInSeconds / 2592000), 'month')
  } else {
    return rtf.format(-Math.floor(diffInSeconds / 31536000), 'year')
  }
}

// Default export with all formatters
export default {
  formatCurrency,
  formatCurrencyFull,
  formatAmount,
  formatDate,
  formatDateForInput,
  formatDateTime,
  formatTime,
  formatNumber,
  formatPercentage,
  formatPhone,
  formatFileSize,
  formatRelativeTime,
}
