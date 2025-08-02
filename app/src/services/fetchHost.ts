import { URL_API } from "../constants/constants"
import type { HostScrapper } from "../types/HostScrapper"
import type { SafeResult } from "../types/Result"
import { safeFetch } from "./safeFetch"

export const fetchHost = async (signal?:AbortSignal):Promise<SafeResult<HostScrapper[]>> => {
    const url = `${URL_API}/scraper/host`
    
    return safeFetch<HostScrapper[]>(url,{signal})
}