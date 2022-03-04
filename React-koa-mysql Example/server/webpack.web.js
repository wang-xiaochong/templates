const path = require('path');

module.exports = {
    entry: '../web/src/App.tsx',
    target: 'node',
    output: {
        // libraryTarget:'ump', //可当作模块被引入，不是普通js
        path: path.resolve(__dirname, 'componernts'),
        filename: 'main.js'
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                // use: 'ts-loader', //隐藏掉报错
                use: 'awesome-typescript-loader', // 明显提示错误是什么
                exclude: /node_modules/,
            },
            {
                test: /\.(jpg|png|svg|gif|css)/,
                use:'file-loader'
            }
        ],
    },
    
};