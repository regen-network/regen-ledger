import dotenv from 'dotenv'
dotenv.config()

import { defaultTheme, defineUserConfig } from 'vuepress-vite'

const { description } = require('../package')

import { backToTopPlugin } from '@vuepress/plugin-back-to-top'
import { mediumZoomPlugin } from '@vuepress/plugin-medium-zoom'
import { googleAnalyticsPlugin } from '@vuepress/plugin-google-analytics'

/**
 * Ref：https://v2.vuepress.vuejs.org/guide/configuration.html#config-file
 */
export default defineUserConfig({
  title: 'Regen Ledger Documentation',
  description: description,
  head: [
    ['meta', { name: 'theme-color', content: '#3eaf7c' }],
    ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }],
  ],
  theme: defaultTheme({
    repo: 'regen-network/regen-ledger',
    docsDir: 'docs',
    contributors: false,
    editLink: false,
    lastUpdated: false,
    navbar: [
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
      // TODO: add to navigation when specs are up-to-date
      // {
      //   text: 'Specifications',
      //   link: '/specs/',
      // },
    ],
    sidebar: {
      '/ledger/': [
        {
          text: 'Introduction',
          children: [
            '/ledger/',
            '/ledger/architecture',
            '/ledger/interfaces',
          ],
        },
        {
          text: 'Get Started',
          children: [
            '/ledger/get-started/',
            '/ledger/get-started/manage-keys',
            '/ledger/get-started/local-testnet',
            '/ledger/get-started/regen-mainnet',
            '/ledger/get-started/redwood-testnet',
          ],
        },
        {
          text: 'Migration Guides',
          children: [
            '/ledger/migrations/',
            '/ledger/migrations/v4.0-migration',
            '/ledger/migrations/v5.0-migration',
            '/ledger/migrations/v5.1-migration',
          ],
        },
      ],
      '/modules/': [
        {
          text: 'Modules',
          children: [
            {
              text: 'List of Modules',
              link: '/modules/',
            },
          ],
        },
        {
          text: 'Data Module',
          children: [
            {
              text: 'Overview',
              link: '/modules/data/',
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
        {
          text: 'Ecocredit Module',
          children: [
            {
              text: 'Overview',
              link: '/modules/ecocredit/',
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
          text: 'Intertx Module',
          children: [
            {
              text: 'Overview',
              link: '/modules/intertx/',
            },
            '/modules/intertx/01_messages',
            '/modules/intertx/02_queries',
          ],
        },
      ],
      '/validators/': [
        {
          text: 'Validators',
          children: [
            {
              text: 'Overview',
              link: '/validators/',
            },
          ],
        },
        {
          text: 'Get Started',
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
          text: 'Upgrade Guides',
          children: [
            '/validators/upgrades/',
            '/validators/upgrades/v2.0-upgrade',
            '/validators/upgrades/v3.0-upgrade',
            '/validators/upgrades/v4.0-upgrade',
            '/validators/upgrades/v4.1-upgrade',
            '/validators/upgrades/v5.0-upgrade',
            '/validators/upgrades/v5.1-upgrade',
          ],
        },
      ],
      '/commands/': [
        {
          text: 'Commands',
          children: [
            {
              text: 'List of Commands',
              link: '/commands/',
            },
          ],
        },
        {
          text: 'Regen App',
          children: [
            { text: 'regen', link: '/commands/regen' },
            { text: 'regen add-genesis-account', link: '/commands/regen_add-genesis-account' },
            { text: 'regen collect-gentxs', link: '/commands/regen_collect-gentxs' },
            { text: 'regen config', link: '/commands/regen_config' },
            { text: 'regen debug', link: '/commands/regen_debug' },
            { text: 'regen export', link: '/commands/regen_export' },
            { text: 'regen gentx', link: '/commands/regen_gentx' },
            { text: 'regen init', link: '/commands/regen_init' },
            { text: 'regen keys', link: '/commands/regen_keys' },
            { text: 'regen migrate', link: '/commands/regen_migrate' },
            { text: 'regen query', link: '/commands/regen_query' },
            { text: 'regen rollback', link: '/commands/regen_rollback' },
            { text: 'regen rosetta', link: '/commands/regen_rosetta' },
            { text: 'regen start', link: '/commands/regen_start' },
            { text: 'regen status', link: '/commands/regen_status' },
            { text: 'regen tendermint', link: '/commands/regen_tendermint' },
            { text: 'regen testnet', link: '/commands/regen_testnet' },
            { text: 'regen tx', link: '/commands/regen_tx' },
            { text: 'regen validate-genesis', link: '/commands/regen_validate-genesis' },
            { text: 'regen version', link: '/commands/regen_version' },
          ]
        },
      ],
      '/tutorials/': [
        {
          text: 'Tutorials',
          children: [
            {
              text: 'List of Tutorials',
              link: '/tutorials/',
            },
          ],
        },
        {
          text: 'User Tutorials',
          children: [
            '/tutorials/user/ibc-transfers',
            '/tutorials/user/currency-allowlist-proposal',
            '/tutorials/user/credit-class-project-batch-management',
          ],
        },
        {
          text: 'Developer Tutorials',
          children: [
            '/tutorials/developer/tendermint-postgres-indexer',
          ],
        },
      ],
      '/specs/': [
        {
          text: 'Specifications',
          children: [
            {
              text: 'Overview',
              link: '/specs/',
            },
            '/specs/regen-ledger',
          ],
        },
        {
          text: 'RFCs',
          children: [
            {
              text: 'RFC Overview',
              link: '/specs/rfcs/',
            },
            '/specs/rfcs/001-ecocredit-module/',
            '/specs/rfcs/002-basket-functionality/',
          ],
        },
        {
          text: 'ADRs',
          children: [
            {
              text: 'ADR Overview',
              link: '/specs/adrs/',
            },
          ],
        },
      ],
    },
  }),
  plugins: [
    backToTopPlugin(),
    mediumZoomPlugin(),
    googleAnalyticsPlugin({
      id: process.env.GOOGLE_ANALYTICS_ID
    }),
  ]
})
