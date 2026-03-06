// utils/copyToClipboard.js
export async function copyToClipboard(text) {
  if (text == null) text = ''
  text = String(text)

  // 1) 优先：现代 Clipboard API（需要安全上下文 https / localhost）
  if (navigator.clipboard && window.isSecureContext) {
    await navigator.clipboard.writeText(text)
    return true
  }

  // 2) 降级：execCommand
  const ta = document.createElement('textarea')
  ta.value = text
  ta.setAttribute('readonly', '')
  ta.style.position = 'fixed'
  ta.style.left = '-9999px'
  ta.style.top = '0'
  ta.style.opacity = '0'

  document.body.appendChild(ta)
  ta.select()
  ta.setSelectionRange(0, ta.value.length)

  const ok = document.execCommand('copy')
  document.body.removeChild(ta)

  if (!ok) throw new Error('Copy failed')
  return true
}
