const path = require('path')
const webpack = require('webpack')
const packageInfo = require(path.join(__dirname, 'package.json'))

module.exports = {
  publicPath: '/',
  outputDir: '../dist/frontend',

  configureWebpack: {
    module: {
      rules: [
        {
          test: /\.scss$/,
          use: [
            'vue-style-loader',
            'css-loader',
            'sass-loader',
          ],
        },
      ],
    },
    plugins: [
      new webpack.DefinePlugin({
        'process.env': {
          NAME: `"${packageInfo.name}"`,
          VERSION: `"${packageInfo.version}"`,
          HOMEPAGE: `"${packageInfo.homepage}"`,
        },
      }),
    ],
  },

  // Load a global SASS file.
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "@/global.sass";`,
      },
    },
  },

  runtimeCompiler: true,

  // Use `frontend` instead of the default `src` directory.
  // See https://vuejsdevelopers.com/2019/03/18/vue-cli-3-rename-src-folder/
  chainWebpack: config => {
    config.entry('app').clear().add(path.join(__dirname, 'main.js')).end()
    config.resolve.alias.set('@', __dirname)
  },
}
