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
      },
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
    docsDir: 'docs',
    nav: [
      {
        text: 'Regen Ledger',
        link: '/ledger/',
      },
      {
        text: 'Modules',
        link: '/modules/',
      },
      {
        text: 'Validators',
        link: '/validators/',
      },
      {
        text: 'Commands',
        link: '/commands/',
      },
      {
        text: 'Tutorials',
        link: '/tutorials/',
      },
    ],
    sidebar: {
      '/ledger/': [
        {
          title: 'Introduction',
          collapsable: false,
          children: [
            '/ledger/',
          ],
        },
        {
          title: 'Get Started',
          collapsable: false,
          children: [
            '/ledger/get-started/',
            '/ledger/get-started/local-testnet',
            '/ledger/get-started/live-networks',
          ],
        },
        {
          title: 'Infrastructure',
          collapsable: false,
          children: [
            '/ledger/infrastructure/',
            '/ledger/infrastructure/interfaces',
          ],
        },
        {
          title: 'Migration Guides',
          collapsable: false,
          children: [
            '/ledger/migrations/',
            '/ledger/migrations/v4.0-migration',
          ],
        },
      ],
      '/modules/': [
        {
          title: 'Modules',
          collapsable: false,
          sidebarDepth: 0,
          children: [
            {
              title: 'List of Modules',
              path: '/modules/',
            },
          ],
        },
        {
          title: 'Ecocredit Module',
          collapsable: false,
          sidebarDepth: 0,
          children: [
            {
              title: 'Overview',
              path: '/modules/ecocredit/',
            },
            '/modules/ecocredit/01_concepts',
            '/modules/ecocredit/02_state',
            '/modules/ecocredit/03_messages',
            '/modules/ecocredit/04_queries',
            '/modules/ecocredit/05_events',
            '/modules/ecocredit/06_types',
            '/modules/ecocredit/07_client',
            '/modules/ecocredit/features/',
          ],
        },
        {
          title: 'Data Module',
          collapsable: false,
          sidebarDepth: 0,
          children: [
            {
              title: 'Overview',
              path: '/modules/data/',
            },
            '/modules/data/01_concepts',
            '/modules/data/02_state',
            '/modules/data/03_messages',
            '/modules/data/04_queries',
            '/modules/data/05_events',
            '/modules/data/06_types',
            '/modules/data/07_client',
            '/modules/data/features/',
          ],
        },
      ],
      '/validators/': [
        {
          title: 'Validators',
          collapsable: false,
          children: [
            {
              title: 'Overview',
              path: '/validators/',
            },
          ],
        },
        {
          title: 'Get Started',
          collapsable: false,
          children: [
            '/validators/get-started/',
            '/validators/get-started/install-regen',
            '/validators/get-started/initialize-node',
            '/validators/get-started/create-a-validator',
            '/validators/get-started/using-quickstart',
            '/validators/get-started/using-state-sync',
            '/validators/get-started/using-cosmovisor',
          ]
        },
        {
          title: 'Upgrade Guides',
          collapsable: false,
          children: [
            '/validators/upgrades/',
            '/validators/upgrades/v2.0-upgrade',
            '/validators/upgrades/v3.0-upgrade',
            '/validators/upgrades/v4.0-upgrade',
          ],
        },
      ],
      '/commands/': [
        {
          title: 'Commands',
          collapsable: false,
          children: [
            {
              title: 'List of Commands',
              path: '/commands/',
            },
          ],
        },
        {
          title: 'Regen App',
          collapsable: false,
          sidebarDepth: 0,
          children: [
              '/commands/regen',
              '/commands/regen_add-genesis-account',
              '/commands/regen_collect-gentxs',
              '/commands/regen_config',
              '/commands/regen_debug',
              '/commands/regen_export',
              '/commands/regen_gentx',
              '/commands/regen_init',
              '/commands/regen_keys',
              '/commands/regen_migrate',
              '/commands/regen_query',
              '/commands/regen_rosetta',
              '/commands/regen_start',
              '/commands/regen_status',
              '/commands/regen_tendermint',
              '/commands/regen_testnet',
              '/commands/regen_tx',
              '/commands/regen_unsafe-reset-all',
              '/commands/regen_validate-genesis',
              '/commands/regen_version',
          ]
        },
      ],
      '/tutorials/': [
        {
          title: 'Tutorials',
          collapsable: false,
          sidebarDepth: 0,
          children: [
            {
              title: 'List of Tutorials',
              path: '/tutorials/',
            },
          ],
        },
        {
          title: 'User Tutorials',
          collapsable: false,
          children: [
            '/tutorials/ibc-transfers',
          ],
        },
        {
          title: 'Developer Tutorials',
          collapsable: false,
          children: [],
        },
      ],
    },
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