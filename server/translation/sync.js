const fs = require('fs')
const path = require('path')

function loadJsonFiles(dir) {
  const files = {}

  function readDir(currentPath) {
    const items = fs.readdirSync(currentPath)

    items.forEach(item => {
      const fullPath = path.join(currentPath, item)
      const stat = fs.statSync(fullPath)

      if (stat.isDirectory()) {
        readDir(fullPath)
      } else if (item.endsWith('.json')) {
        const content = JSON.parse(fs.readFileSync(fullPath, 'utf8'))
        const relativePath = path.relative(dir, fullPath)
        files[relativePath] = content
      }
    })
  }

  readDir(dir)
  return files
}

function getValueByPath(obj, path) {
  return path.split('.').reduce((current, key) => current && current[key], obj)
}

function findMissingKeys(zhObj, enObj) {
  const missingKeys = []

  function compareObjects(zh, en, path = '') {
    for (const key in zh) {
      const currentPath = path ? `${path}.${key}` : key
      const zhValue = zh[key]
      const enValue = en[key]

      if (typeof zhValue === 'object' && zhValue !== null) {
        if (!enValue || typeof enValue !== 'object') {
          missingKeys.push(currentPath)
        } else {
          compareObjects(zhValue, enValue, currentPath)
        }
      } else if (!(key in en)) {
        missingKeys.push(currentPath)
      }
    }
  }

  compareObjects(zhObj, enObj)
  return missingKeys
}

const zhcn = loadJsonFiles(path.join(__dirname, 'zh-CN'))
const enus = loadJsonFiles(path.join(__dirname, 'en-US'))

// 对比每个文件中的键值
for (const filePath in zhcn) {
  const missingKeys = findMissingKeys(zhcn[filePath], enus[filePath] || {})

  if (missingKeys.length > 0) {
    missingKeys.forEach(key => {
      const zhValue = getValueByPath(zhcn[filePath], key)
      console.log(`${filePath} | ${key} | ${zhValue}`)
    })
  }
}
