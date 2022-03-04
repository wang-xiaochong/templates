import axios from "@/libs/axios";

export interface BannerData {
    ID: number,
    img: string,
    href: string,
    sort: number
}

export async function getAllBanners(): Promise<BannerData[]> {
    let { data } = await axios('api/banner/getAllBanners', {
        headers: { a: 12 }
    });
    return data
}
