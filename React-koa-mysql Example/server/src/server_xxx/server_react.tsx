// let num:Number = 3
// console.log(num)

import Koa from 'koa'
import Router from '@koa/router'
// import main from '../componernts/main.js'

import React from 'react'
import ReactDomServer from 'react-dom/server'

import App from '@/App'

const app = new Koa()
const router = new Router()

// 服务端进行渲染
router.get('/', async ctx => {
    ctx.body = ReactDomServer.renderToString(<App />)
})

app.use(router.routes())
app.listen(7001, () => {
    console.log("server start at localhost:7001")
})
