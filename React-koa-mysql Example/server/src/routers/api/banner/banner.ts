import Router from '@koa/router'
import { getAllBanners } from '~/models/banner'

let router = new Router()
// api
router.prefix('/banner')
router.get('/getAllBanners', async ctx => {
    let banners = await getAllBanners()
    ctx.body = banners
})

export default router.routes()
