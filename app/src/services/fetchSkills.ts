import { URL_API } from "../constants/constants"
import { Result, type SafeResult } from "../types/Result"
import { safeFetch } from "./safeFetch"

export const fetchSkills = async (signal?:AbortSignal):Promise<SafeResult<string[]>> => {
    const url = `${URL_API}/skills` 
    
    return  safeFetch<string[]>(url,{signal})
}