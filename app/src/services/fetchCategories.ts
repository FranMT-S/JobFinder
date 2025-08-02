import { URL_API } from "../constants/constants"
import type { Category } from "../types/Categories"
import { Result, type SafeResult } from "../types/Result"
import { safeFetch } from "./safeFetch"

export const fetchCategories = async (signal?:AbortSignal):Promise<SafeResult<Category[]>> => {
    const url = `${URL_API}/categories`
    
    return safeFetch<Category[]>(url,{signal})
}