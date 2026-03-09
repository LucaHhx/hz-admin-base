/**
 * 网站配置文件
 */
const greenText = (text) => `\x1b[32m${text}\x1b[0m`

const config = {
  appName: 'ADMIN',
  appLogo: 'logo.svg',
  showViteLogo: true,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(greenText('Vite Logo'))
  }
}

export default config

