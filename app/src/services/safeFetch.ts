import { type ResponseData, type ResponseError } from "../types/Response"
import { Result, type SafeResult } from "../types/Result"


// SafeFetch wrapper the fetch to return a safeResult that is a tuple [error,data]
export const safeFetch = async <T>(input: RequestInfo | URL, init?: RequestInit):Promise<SafeResult<T>> => {
    try {
        const response = await fetch(input, init)


        let res:any  = await response.json();
        
        if(!res.ok){
            const error = res as ResponseError;
            return Result.Fail(new Error(error.message || "Error processing the data"))
        }

        const resData = res as ResponseData<T>
        return Result.Success(resData.data)
    } catch (error) {
        return Result.Fail(error as Error)
    }
}
    
