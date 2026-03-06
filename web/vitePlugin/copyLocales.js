import fs from 'fs'
import path from 'path'

export const copyLocalesPlugin = () => {
  return {
    name: 'copy-locales',
    closeBundle() {
      const srcDir = path.resolve(__dirname, '../src/locales')
      const destDir = path.resolve(__dirname, '../dist/locales')

      // 确保目标目录存在
      if (!fs.existsSync(destDir)) {
        fs.mkdirSync(destDir, { recursive: true })
      }

      // 复制整个 locales 目录
      function copyDir(src, dest) {
        if (!fs.existsSync(dest)) {
          fs.mkdirSync(dest, { recursive: true })
        }

        const entries = fs.readdirSync(src, { withFileTypes: true })

        for (const entry of entries) {
          const srcPath = path.join(src, entry.name)
          const destPath = path.join(dest, entry.name)

          if (entry.isDirectory()) {
            copyDir(srcPath, destPath)
          } else {
            fs.copyFileSync(srcPath, destPath)
          }
        }
      }

      copyDir(srcDir, destDir)
    }
  }
}
