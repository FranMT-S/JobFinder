
export interface ResponseData<T = any>{
    ok:boolean
    data:T
    message:string
}

export class ResponseError extends Error {
    ok:boolean
	status_code:number    
	message:string 
	error_uuid:string 
    
    constructor(message: string, status_code: number, error_uuid: string) {
        super(message)
        this.ok = false
        this.status_code = status_code
        this.message = message
        this.error_uuid = error_uuid
    }
}

