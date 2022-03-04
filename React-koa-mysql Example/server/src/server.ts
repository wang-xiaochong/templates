import app from '~/libs/server'
import { init as corsInit } from '~/libs/cors'
import routers from '~/routers/routers'
import koaStatic from 'koa-static'
import { staticRoot } from '~/config/app'
import { enableCors } from '~/config/app'

if (enableCors) corsInit(app);

app.use(routers)
app.use(koaStatic(staticRoot))