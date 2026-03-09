# 资源交付清单

> 需求: login-mode-config | 创建: 2026-03-09

## AI 生成资源（必须交付）

以下资源由 UI 设计师直接生成并交付，**不可留空**:

| # | 资源 | 路径 | 状态 |
|---|------|------|------|
| 1 | SVG 图标（9个） | `icons/*.svg` | 已交付 |
| 2 | CSS 变量 (Design Tokens) | `tokens.css` | 已交付 |
| 3 | Tailwind 扩展配置 | `tailwind.config.js` | 已交付 |

### 图标清单

| 文件 | 描述 | 尺寸 |
|------|------|------|
| `icons/user.svg` | 用户图标 | 20x20 |
| `icons/lock.svg` | 锁图标 | 20x20 |
| `icons/shield.svg` | 盾牌图标 | 20x20 |
| `icons/image-captcha.svg` | 图片验证码图标 | 20x20 |
| `icons/arrow-right.svg` | 右箭头图标 | 20x20 |
| `icons/arrow-left.svg` | 左箭头图标 | 20x20 |
| `icons/key.svg` | 钥匙图标 | 20x20 |
| `icons/globe.svg` | 地球图标 | 20x20 |
| `icons/refresh.svg` | 刷新图标 | 20x20 |

## 需人工提供资源（记录 + 占位）

以下资源需要人工提供，UI 设计师需记录需求并提供占位方案:

| # | 资源描述 | 用途 | 占位方案 | 状态 |
|---|----------|------|----------|------|
| 1 | 应用 Logo | 登录页顶部品牌标识 | 蓝色圆角方块+首字母"H"（项目已有 `@/assets/logo.svg` 可直接使用） | 无需额外提供 |
| 2 | 右侧装饰图 | 桌面端右侧装饰区域 | 蓝色渐变背景+文字（项目已有 `@/assets/login_bg.svg` 可直接使用） | 无需额外提供 |

## 自检清单

- [x] `icons/` 目录包含所有设计稿中使用的 SVG 图标（9个）
- [x] `tokens.css` 包含完整的 CSS 变量（颜色、间距、字号）
- [x] `tailwind.config.js` 包含设计系统的 Tailwind 扩展配置
- [x] merge.html 中无外部 URL 引用本地应有的资源
- [x] 所有需人工提供的资源已记录并有占位方案
- [x] 占位方案在视觉上可接受（不影响布局预览）
