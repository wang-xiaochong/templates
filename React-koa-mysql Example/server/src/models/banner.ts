import  db  from '~/libs/database';
import {BannerData} from 'models/banner'

export async function getAllBanners():Promise<BannerData[]> {
    let [rows] = await db.query('SELECT * FROM banner_table');
    let banners = rows as BannerData[]
    return banners
}

