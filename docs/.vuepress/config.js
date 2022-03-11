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
          title: 'Introduction',
          collapsable: false,
          children: [
            '/intro/',
          ]
        },
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
          title: 'Migrations',
          collapsable: false,
          children: [
            '/migrations/upgrade',
            '/migrations/v2.0-upgrade',
            '/migrations/v3.0-upgrade',
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
              title: 'Ecocredit Module',
              collapsable: false,
              children: [
                {
                  title: 'Overview',
                  path: '/modules/ecocredit/'
                },
                '/modules/ecocredit/01_concepts',
                '/modules/ecocredit/02_state',
                '/modules/ecocredit/03_messages',
                '/modules/ecocredit/04_events',
                '/modules/ecocredit/05_client',
                {
                  title: 'Protobuf - Core',
                  path: 'https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.v1'
                },
                {
                  title: 'Protobuf - Basket',
                  path: 'https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.basket.v1'
                },
                {
                  title: 'Protobuf - Marketplace',
                  path: 'https://buf.build/regen/regen-ledger/docs/main/regen.ecocredit.marketplace.v1'
                }
              ]
            },
            {
              title: 'Data Module',
              collapsable: false,
              children: [
                {
                  title: 'Overview',
                  path: '/modules/data/'
                },
                '/modules/data/01_concepts',
                '/modules/data/02_state',
                '/modules/data/03_messages',
                '/modules/data/04_events',
                '/modules/data/05_client',
                {
                  title: 'Protobuf',
                  path: 'https://buf.build/regen/regen-ledger/docs/main/regen.data.v1'
                }
              ]
            },
            {
              title: 'Group Module',
              collapsable: false,
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
                  path: 'https://buf.build/regen/regen-ledger/docs/main/regen.group.v1alpha1'
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
            '/tutorials/ibc-transfers'
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
  ],
  markdown: {
    extendMarkdown: md => {
      md.use(require('./markdown-it-gh'))
    }
  }
}
