import Router from '@koa/router'
import  bannerRoutes  from './banner/banner' 

let router = new Router()
// api
router.prefix('/api')
router.get('/getAPI', async ctx => {
    let apis = [{ name: 'banners' }, { name: 'users' }]
    ctx.body = apis
})

router.use(bannerRoutes)

export default router.routes()
