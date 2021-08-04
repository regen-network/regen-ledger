require('dotenv').config()

const { description } = require('../package')
const webpack = require('webpack')

module.exports = {
  configureWebpack: (config) => {
    return { plugins: [
      new webpack.EnvironmentPlugin({ ...process.env })
    ]}
  },
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
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }],
    /**
     * Google Analytics 4 is not supported in vuepress v1 but will be in v2.
     * The following is a workaround until we update to vuepress v2.
     *
     * ref：https://github.com/vuejs/vuepress/issues/2713
     */
    [
      'script',
      {
        async: true,
        src: 'https://www.googletagmanager.com/gtag/js?id=' + process.env.GOOGLE_ANALYTICS_ID,
      }
    ],
    [
      'script',
      {},
      [
        "window.dataLayer = window.dataLayer || [];\nfunction gtag(){dataLayer.push(arguments);}\ngtag('js', new Date());\ngtag('config', '" + process.env.GOOGLE_ANALYTICS_ID + "');",
      ],
    ],
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
          collapsable: false,
          children: [
            '/getting-started/',
            '/getting-started/live-networks',
            '/getting-started/running-a-full-node',
            '/getting-started/running-a-validator',
            '/getting-started/prerequisites'
          ]
        },
        {
          title: 'Regen Ledger',
          collapsable: false,
          children: [
            '/regen-ledger/',
            '/regen-ledger/interfaces',
          ]
        },
        {
          title: 'Modules',
          collapsable: false,
          children: [
            {
              title: 'Data Module',
              collapsable: false,
              children: [
                '/modules/data/',
                {
                  title: 'Protobuf Documentation',
                  path: '/modules/data/protobuf'
                }
              ]
            },
            {
              title: 'Ecocredit Module',
              collapsable: false,
              children: [
                '/modules/ecocredit/',
                {
                  title: 'Protobuf Documentation',
                  path: '/modules/ecocredit/protobuf'
                }
              ]
            },
          ]
        },
        {
          title: 'Tutorials',
          collapsable: false,
          children: [
            '/tutorials/',
            // '/tutorials/data-cli',
            // '/tutorials/data-grpc',
          ]
        },
      ],
    }
  },
  /**
   * Apply plugins，ref：https://v1.vuepress.vuejs.org/plugin/
   */
  plugins: [
    '@vuepress/plugin-back-to-top',
    '@vuepress/plugin-medium-zoom',
  ]
}
