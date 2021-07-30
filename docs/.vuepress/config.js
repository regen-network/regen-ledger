const { description } = require('../package')

module.exports = {
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#title
   */
  title: 'Regen Ledger Documentation',
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#description
   */
  description: description,

  /**
   * Extra tags to be injected to the page HTML `<head>`
   *
   * ref：https://v1.vuepress.vuejs.org/config/#head
   */
  head: [
    ['meta', { name: 'theme-color', content: '#3eaf7c' }],
    ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }]
  ],

  /**
   * Theme configuration, here is the default theme configuration for VuePress.
   *
   * ref：https://v1.vuepress.vuejs.org/theme/default-theme-config.html
   */
  themeConfig: {
    repo: 'regen-network/regen-ledger',
    editLinks: false,
    docsDir: 'docs',
    lastUpdated: false,
    nav: [
      {
        text: 'Getting Started',
        link: '/getting-started/',
      }
    ],
    sidebar: {
      '/': [
        {
          title: 'Getting Started',
          children: [
            '/getting-started/',
            '/getting-started/live-networks',
            '/getting-started/running-a-full-node',
            '/getting-started/running-a-validator',
            '/getting-started/prerequisites'
          ]
        },
        'core-functionality',
        'api',
        {
          title: 'Modules',
          children: [
            {
              title: 'Data Module',
              collapsable: false,
              children: [
                '/modules/data/',
                {title: 'Protobuf Documentation', path: '/modules/data/protobuf', collapsable: true}
              ]
            },
            {
              title: 'Ecocredit Module',
              collapsable: false,
              children: [
                '/modules/ecocredit/',
                {title: 'Protobuf Documentation', path: '/modules/ecocredit/protobuf', collapsable: true}
              ]
            },
          ]
        }
      ],
    }
  },

  /**
   * Apply plugins，ref：https://v1.vuepress.vuejs.org/zh/plugin/
   */
  plugins: [
    '@vuepress/plugin-back-to-top',
    '@vuepress/plugin-medium-zoom',
  ]
}
