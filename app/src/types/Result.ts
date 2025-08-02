
export type SafeResult<T> = [error:Error,data:null] | [error:null,data:T]

export class Result {
    static Success<T>(data:T):SafeResult<T> {
       return [null,data];
    }
 
    static Fail<T>(error:Error):SafeResult<T> {
       return [error,null];
    }
 }