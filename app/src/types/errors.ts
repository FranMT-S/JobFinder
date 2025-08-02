export class ErrorResponse extends Error {
    apiMessage: string
    constructor(message: string, apiMessage: string) {
        super(message)
        this.apiMessage = apiMessage
    }
}