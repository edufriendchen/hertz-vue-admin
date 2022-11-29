/**
 * 网站配置文件
 */

const config = {
  appName: 'Hertz-Vue-Admin',
  appLogo: '@/assets/nav_logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Hertz-Vue-Admin，开源地址：https://github.com/EduFriendChen/hertz-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本: 先行版`
      )
    )
    console.log('\n')
  }
}

export default config
