const AntdDayjsWebpackPlugin = require('antd-dayjs-webpack-plugin');
const withSass = require('@zeit/next-sass')

module.exports = {
  webpack: (config, { buildId, dev, isServer, defaultLoaders }) => {
    // Perform customizations to webpack config
    // Important: return the modified config
    config.plugins.push(new AntdDayjsWebpackPlugin())
    return config
  },
};