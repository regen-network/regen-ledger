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
              children: [
                {
                  title: 'Overview',
                  path: '/modules/data/'
                },
                // '/modules/data/01_concepts',
                // '/modules/data/02_state',
                // '/modules/data/03_messages',
                // '/modules/data/04_events',
                // '/modules/data/05_client',
                {
                  title: 'Protobuf',
                  path: '/modules/data/protobuf'
                }
              ]
            },
            {
              title: 'Ecocredit Module',
              children: [
                {
                  title: 'Overview',
                  path: '/modules/ecocredit/'
                },
                // '/modules/ecocredit/01_concepts',
                // '/modules/ecocredit/02_state',
                // '/modules/ecocredit/03_messages',
                // '/modules/ecocredit/04_events',
                // '/modules/ecocredit/05_client',
                {
                  title: 'Protobuf',
                  path: '/modules/ecocredit/protobuf'
                }
              ]
            },
            {
              title: 'Group Module',
              children: [
                {
                  title: 'Overview',
                  path: '/modules/group/'
                },
                '/modules/group/01_concepts',
                '/modules/group/02_state',
                '/modules/group/03_messages',
                '/modules/group/04_events',
                // '/modules/group/05_client',
                {
                  title: 'Protobuf',
                  path: '/modules/group/protobuf'
                }
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
