// Tailwind 扩展配置 — 登录模式配置
// 前端可将此配置合并到项目的 tailwind.config.js 中

/** @type {import('tailwindcss').Config} */
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#EFF6FF',
          100: '#DBEAFE',
          500: '#2563EB',
          600: '#1D4ED8',
          700: '#1E40AF',
        },
        'brand-bg': '#0F6BF8',
      },
      fontFamily: {
        sans: ['system-ui', '-apple-system', 'Segoe UI', 'PingFang SC', 'Noto Sans SC', 'sans-serif'],
      },
      boxShadow: {
        'card': '0 4px 24px rgba(0, 0, 0, 0.08)',
        'btn': '0 6px 16px rgba(37, 99, 235, 0.3)',
        'btn-hover': '0 10px 24px rgba(37, 99, 235, 0.35)',
      },
      spacing: {
        '18': '4.5rem',  // 72px
      },
    },
  },
}
