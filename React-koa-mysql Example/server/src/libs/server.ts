// require没有类型信息 node中的js是不认ts的
// const Koa = require('koa')
// import有类型信息，支持ts
import Koa from 'koa';
import {port} from '~/config/server'

const app = new Koa();
app.listen(port, () => {
    console.log(`server start at localhost:${port}`);
})

export default app