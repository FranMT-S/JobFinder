import { URL_API } from "../constants/constants"
import { Category } from "../enums/category"
import type { JobRequest, JobScrapeated } from "../types/Job"
import { type SafeResult } from "../types/Result"
import { safeFetch } from "./safeFetch"

export const fetchScrapper = async (page:number,limit:number,jobRequest:JobRequest,signal?:AbortSignal):Promise<SafeResult<JobScrapeated[]>> => {
    const url = `${URL_API}/scraper`
    const params = new URLSearchParams({
        page: page.toString(),
        max: limit.toString()
    })

    const apiUrl = `${url}?${params.toString()}`
    const body = {
        "location":  "",
        "level": jobRequest?.level || "",
        "skills": jobRequest?.skills || [],
        "modalities": jobRequest?.modalities || "",
        "minimumSalaryExpectation": jobRequest?.minimumSalaryExpectation || -1,
        "maximumSalaryExpectation": jobRequest?.maximumSalaryExpectation || -1,
        "category": jobRequest?.category || Category.NotCategory,
        "host": jobRequest?.hostSelected || []
      }   
      
    return safeFetch<JobScrapeated[]>(apiUrl,{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body),
        signal
    })
}