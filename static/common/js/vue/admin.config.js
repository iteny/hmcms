// nodejs 中的path模块
var path = require('path');
var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    // 入口文件，path.resolve()方法，可以结合我们给定的两个参数最后生成绝对路径，最终指向的就是我们的index.js文件
    entry: path.resolve(__dirname, '../vue/admin.js'),
        
    // 输出配置
    output: {
        // 输出路径是 myProject/output/static
        path: path.resolve(__dirname, '../dist'),
        publicPath: '/',
        filename: 'xh.js',
        // chunkFilename: '[id].[chunkhash].js'
    },    
    resolve: {
        extensions: ['', '.js', '.vue','.css']
    },
    vue: {
        loaders: {
            css: ExtractTextPlugin.extract(
                "style-loader",
                "css-loader",
                {
                    publicPath: path.resolve(__dirname, '../dist/'),
                }
            ),
        }
    },
    module: {        
        loaders: [
            // 使用vue-loader 加载 .vue 结尾的文件
            { test: /\.vue$/,loader: 'vue'},
            { test: /\.js$/,loader: 'babel',exclude: /node_modules/},
            { test: /\.css$/, loader: 'style!css!autoprefixer'},
            { test: /\.(gif|jpg|png|woff|svg|eot|ttf)\??.*$/, loader: 'url-loader?limit=8192'},
            { test: /\.(html|tpl)$/, loader: 'html-loader' }
        ]
    },
     // 转es5
    babel: {
        presets: ['es2015'],
        plugins: ['transform-runtime']
    },
    plugins: [
        // new webpack.optimize.UglifyJsPlugin({
        //     compress: {
        //         warnings: false
        //     }
        // }),
        new ExtractTextPlugin("[name].css",{ allChunks : true,resolve : ['modules'] }),     
    ],
}