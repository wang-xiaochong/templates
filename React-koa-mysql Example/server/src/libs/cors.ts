
import Koa from 'Koa'

export function init(app: Koa) {
    app.use(async (ctx, next) => {
        // CORS=>Cross-orgin Resurce Share 跨域资源共享
        ctx.set('Access-Control-Allow-Origin', '*');    // 域
        ctx.set('Access-Control-Allow-Headers', '*');   // 头
        if (ctx.method == 'OPTIONS') {
            ctx.body = 'OK'
        } else {
            await next();
        }
    })
}

