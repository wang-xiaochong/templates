
import app from '~/libs/server'
import db from '~/libs/database'

app.use(async ctx => {
    let [rows] = await db.query('SELECT * FROM banner_table');
    ctx.body = rows;
})


