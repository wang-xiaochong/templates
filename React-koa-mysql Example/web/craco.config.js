const CracoAlias = require('craco-alias');
const path = require('path')
module.exports = {
    plugins: [
        {
            plugin: CracoAlias,
            options: {
                aliases: {
                        'models': path.resolve(__dirname, './src/models'),
                        '@': path.resolve(__dirname, './src')
                    
                }
            }
        }
    ]
}