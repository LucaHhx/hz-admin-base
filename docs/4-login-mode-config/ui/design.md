# UI 设计系统

> 需求: login-mode-config | 角色: ui

## 调色板

### 主色（Primary）
| 名称 | 色值 | 用途 |
|------|------|------|
| primary-500 | `#2563EB` | 主按钮、链接、激活态 |
| primary-600 | `#1D4ED8` | 主按钮悬停 |
| primary-700 | `#1E40AF` | 主按钮按下 |
| primary-50 | `#EFF6FF` | 主色背景（极淡） |
| primary-100 | `#DBEAFE` | 主色卡片背景 |

### 中性色（Neutral）
| 名称 | 色值 | 用途 |
|------|------|------|
| gray-50 | `#F9FAFB` | 页面背景 |
| gray-100 | `#F3F4F6` | 输入框背景 |
| gray-200 | `#E5E7EB` | 分割线、边框 |
| gray-300 | `#D1D5DB` | 禁用态边框 |
| gray-400 | `#9CA3AF` | 占位符文字 |
| gray-500 | `#6B7280` | 辅助文字 |
| gray-700 | `#374151` | 正文文字 |
| gray-900 | `#111827` | 标题文字 |

### 语义色
| 名称 | 色值 | 用途 |
|------|------|------|
| success | `#10B981` | 成功提示 |
| warning | `#F59E0B` | 警告提示 |
| error | `#EF4444` | 错误提示、必填标记 |
| info | `#3B82F6` | 信息提示 |

### 装饰色
| 名称 | 色值 | 用途 |
|------|------|------|
| brand-bg | `#0F6BF8` | 桌面端右侧装饰背景（沿用现有设计） |

## 字体方案

| 层级 | 字号 | 行高 | 字重 | 用途 |
|------|------|------|------|------|
| h1 | 36px (text-4xl) | 1.2 | bold (700) | 应用名称 |
| h2 | 24px (text-2xl) | 1.3 | semibold (600) | 页面标题 |
| body | 16px (text-base) | 1.5 | normal (400) | 正文 |
| small | 14px (text-sm) | 1.5 | normal (400) | 辅助说明文字 |
| caption | 12px (text-xs) | 1.4 | normal (400) | 模式标签、微注释 |

**字体族**: `system-ui, -apple-system, 'Segoe UI', 'PingFang SC', 'Noto Sans SC', sans-serif`

## 间距系统

基础单位: 4px

| Token | 值 | 用途 |
|-------|------|------|
| space-1 | 4px | 图标与文字间距 |
| space-2 | 8px | 紧凑元素间距 |
| space-3 | 12px | 表单元素内部间距 |
| space-4 | 16px | 表单项之间间距 |
| space-6 | 24px | 区块间距 |
| space-8 | 32px | 大区块间距 |
| space-9 | 36px | 标题与表单间距 |
| space-12 | 48px | 顶部/底部内边距 |

## 圆角与阴影

| Token | 值 | 用途 |
|-------|------|------|
| rounded-sm | 4px | 输入框、小按钮 |
| rounded-md | 8px | 卡片、按钮 |
| rounded-lg | 12px | 登录卡片 |
| rounded-full | 9999px | 圆形按钮 |
| shadow-card | `0 4px 24px rgba(0,0,0,0.08)` | 登录卡片阴影 |
| shadow-btn | `0 6px 16px rgba(37,99,235,0.3)` | 主按钮阴影 |

## 组件规范

### 按钮

| 类型 | 高度 | 圆角 | 样式 |
|------|------|------|------|
| 主按钮 | 44px (h-11) | rounded-md | bg-primary-500, 白色文字, shadow-btn |
| 次按钮 | 44px (h-11) | rounded-md | 白色背景, gray-300 边框, gray-700 文字 |
| 圆形按钮 | 44x44px | rounded-full | 渐变蓝背景, 白色图标 |

### 输入框

| 属性 | 值 |
|------|------|
| 高度 | 44px (size="large") |
| 边框 | 1px solid gray-200 |
| 圆角 | rounded-md (8px) |
| 聚焦态 | border-primary-500, ring-2 ring-primary-100 |
| 占位符色 | gray-400 |

### 验证码图片

| 属性 | 值 |
|------|------|
| 宽度 | 128px (w-32) |
| 高度 | 44px（与输入框等高） |
| 边框 | 1px solid gray-300 |
| 圆角 | rounded-md |
| 鼠标 | cursor-pointer |

### 登录卡片

| 属性 | 值 |
|------|------|
| 最大宽度 | 384px (md:w-96) |
| 内边距 | pt-12 pb-10 |
| 背景 | 白色 |
| 阴影 | shadow-card（桌面端） |

## 设计决策

1. **沿用现有视觉风格**: 保持与当前 strict 模式登录页一致的蓝白配色和整体布局，降低用户认知成本
2. **模式差异最小化**: 三种模式共用同一套布局框架，仅表单内容不同，保持视觉一致性
3. **simple 模式去除验证码和分步**: 最简洁的表单，只有用户名+密码+登录按钮
4. **captcha 模式在 simple 基础上加验证码**: 用户名+密码+验证码，一步完成
5. **strict 模式保持现有分步流程**: 完全沿用现有 UI 和交互，确保生产环境零影响
6. **Passkey 按钮**: simple 和 captcha 模式不显示 Passkey 按钮（Passkey 是 strict 模式特有功能）
7. **响应式策略**: 移动端全宽卡片，桌面端左右分栏（左侧表单 + 右侧装饰图）
