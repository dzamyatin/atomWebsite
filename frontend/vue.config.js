const webpack = require('webpack');

module.exports = {
  // chainWebpack: config => {
  //   config.resolve.alias.set('vue', '@vue/compat')
  //
  //   config.module
  //       .rule('vue')
  //       .use('vue-loader')
  //       .tap(options => {
  //         return {
  //           ...options,
  //           compilerOptions: {
  //             compatConfig: {
  //               MODE: 2,
  //             },
  //           },
  //         }
  //       })
  // },
  configureWebpack: {
    // Set up all the aliases we use in our app.
    plugins: [
      new webpack.optimize.LimitChunkCountPlugin({
        maxChunks: 6
      })
    ]
  },
  pwa: {
    name: 'Vue Argon Design',
    themeColor: '#172b4d',
    msTileColor: '#172b4d',
    appleMobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: '#172b4d'
  },
  css: {
    // Enable CSS source maps.
    sourceMap: process.env.NODE_ENV !== 'production'
  }
};
