const path = require('path')
const webpackNodeExternals = require('webpack-node-externals')

module.exports = {
    entry: '~/server.ts',
    output: {
        path: path.resolve(__dirname, 'build'),
        filename:'server.js'
    },
    externals: [
      webpackNodeExternals()  
    ],
    target:'node',

    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader', //隐藏掉报错
                // use: 'awesome-typescript-loader', // 明显提示错误是什么
                exclude: /node_modules/,
            },
            {
                test: /\.css$/,
                use:['css-loader']
            },
            {
                test: /\.(jpg|png|svg|gif)/,
                use:'file-loader'
            }
        ]
    },

    resolve: {
        extensions: ['.ts', '.tsx', 'js', 'jsx'],
        alias: {
            '@': path.resolve(__dirname, '../web/src'),
            '~': path.resolve(__dirname, './src'),
            'models':path.resolve(__dirname,'../web/src/models')
        }
    }
}