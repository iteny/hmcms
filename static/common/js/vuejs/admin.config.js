const path = require('path'),
    webpack = require('webpack'),
    NODE_ENV = process.env.NODE_ENV || "DEV",//环境类型
    NODE_RUN = process.env.NODE_RUN || "0",//是否运行
    ROOT_PATH = path.resolve(__dirname,'../vuejs/admin.js'),
    OUT_PATH = path.resolve(__dirname, '../dist'),
    SERVER_PATH =  process.env.SERVER || "./build/",//服务路径
    ExtractTextPlugin = require("extract-text-webpack-plugin"),
    HtmlWebpackPlugin = require("html-webpack-plugin");
module.exports = {
    entry:{//定义入口文件
        page: [ROOT_PATH],
        //打包第三方库作为公共包
        common:['vue','vue-router','element-ui']
    },
    output:{//输出配置
        path: OUT_PATH,
        publicPath: '/',
        filename: '[name].js?t=[hash:5]',
    },
    externals:[require('webpack-require-http')],
    resolve: {
        extensions: ['.js', '.vue', '.jsx', '.less', '.scss', '.css']
    },    
    module:{
        rules: [{
            test: /\.html$/,
            use: [{
                loader: 'html-loader',
                options: {
                    //root: resolve(__dirname, 'src'),
                    attrs: ['img:src', 'link:href']
                }
           }]
        },{
            test: /\.js(x)*$/,
            exclude: /^node_modules$/,
            //loader: 'babel-loader'
            use: ['babel-loader'],
        },{
            test: /\.vue$/,
            use: ['vue-loader']
        },{
            test: /\.css$/,
            exclude: /^node_modules$/,
            loader: ExtractTextPlugin.extract({
                fallbackLoader: "style-loader",
                loader: "css-loader",
                publicPath: path.resolve(__dirname, '../dist/')
            })
        },{
            test: /\.less/,
            exclude: /^node_modules$/,
            loader: ExtractTextPlugin.extract({
                fallbackLoader: 'style-loader',
                loader: "css-loader!less-loader",
                publicPath: path.resolve(__dirname, '../dist/')
            })
        },{
            test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
            use: [{
                loader: "url-loader",
                query: {
                    limit: 10000,
                    name: 'imgs/[name].[hash:7].[ext]'
                }
            }]
        },{
            test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
            use: [{
                loader: "url-loader",
                query: {
                    limit: 5000,
                    name: '/fonts/[name].[ext]'
                }
            }]
        }]
    },

    plugins: [
        new webpack.optimize.UglifyJsPlugin({
            test: /(\.js|\.vue)$/,
            compress: {
                warnings: false
            },
            comments: false
        }),
        new webpack.LoaderOptionsPlugin({
            test:/\.vue$/,//把.vue组件里的css分离出来
            options: {
                vue: {
                    loaders: {
                      css: ExtractTextPlugin.extract({fallback:'vue-style-loader', use:'css-loader'}),
                    },
                },
            }
        }),
        new webpack.optimize.CommonsChunkPlugin({name: 'common', filename: 'common.js' }),
        new ExtractTextPlugin('vue.css?t=[hash:5]')//css分离
    ],
}