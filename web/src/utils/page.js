// import { fmtTitle } from '@/utils/fmtRouterTitle'
import config from '@/core/config'

export default function getPageTitle(pageTitle, route) {
  var env = 'DEV'
  switch (window.location.hostname) {
    case 'bankinf-b.onepay.vip':
      env = 'PROD'
      break
    case 'bankinf-b.onepayment.vip':
      env = 'PRE'
      break
  }
  return `${config.appName} - ${env}`
}
