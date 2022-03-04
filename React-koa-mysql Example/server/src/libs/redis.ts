import redis from 'redis'
import redisConf from '~/config/redis'
import { promisify } from 'util'


let client = redis.createClient(redisConf)

export const get = promisify(client.get).bind(client);
export const set =  promisify(client.set).bind(client);


// 回调取值 服务器为异步 所以需要更改
// client.get('name', (err, data) => {
//     console.log(err, data)
// })

// 回调函数转成await方式
// const get = promisify(client.get).bind(client);
// (async () => {
//     let ret = await get('name')
//     console.log(ret)
// })();

