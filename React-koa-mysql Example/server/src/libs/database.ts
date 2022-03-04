
// mysql2中有promise版本,但引入promise版本后所有的mysql操作都是异步
import mysql from 'mysql2/promise';


// createConnection     单连接
// createPool           连接池
// createPoolCluster    集群

import { db } from '~/config/database'

export default mysql.createPool(db);


// 普通版本mysql2 返回信息是回调方式 koa是异步的，async await ，所以需要更换版本
// conn.query('SELECT * FROM banner_table', (err, rows) => {
//     if(!err) console.log(rows)
// })

