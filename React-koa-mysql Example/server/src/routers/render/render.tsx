import Router from '@koa/router'
import React from 'react'
import ReactDomServer from 'react-dom/server'
import App from '@/App'
import path from 'path'
import fs from 'fs'
import { staticRoot } from '~/config/app'

let router = new Router()

if (process.env.NODE_ENV === 'production') {
    // 服务端进行渲染
    router.get('/', async ctx => {
        let index = fs.readFileSync(path.resolve(staticRoot, 'index.html')).toString()
        let str = ReactDomServer.renderToString(<App />)
        let ret = index.replace('<div id="root"></div>', `<div id="root">${str}</div>`)
        ctx.body = ret
    })
} else {
    router.get('/', async ctx => {
        ctx.body = "render disabled because env is development"
    })
}

export default router.routes()
