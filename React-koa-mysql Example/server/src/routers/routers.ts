import Router from '@koa/router'
import APIRouters from '~/routers/api/api'
import RenderRouters from '~/routers/render/render'

const router = new Router();

router.use(APIRouters)
router.use(RenderRouters)

export default router.routes();