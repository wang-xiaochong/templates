
import app from '~/libs/server'

// 使用
import * as redis from '~/libs/redis'
(async () => {
    redis.set('name', '456')
})();

(async () => {
    let a = await redis.get('name')
    console.log(a)
})();


app.use(async ctx => {
    // let [rows] = await db.query('SELECT * FROM banner_table');
    ctx.body = "rows";
})


