
export class ValidationError extends Error {
  errors: string[]
  name: string = 'Validation'
  
  constructor(errs: string[]) {
    super()
    Object.setPrototypeOf(this, new.target.prototype)
    this.errors = errs
  }
}

export class RequestError extends Error {
  name: string = 'Request'
  message: string

  constructor(message: string) {
    super()
    this.message = message
    Object.setPrototypeOf(this, new.target.prototype)
  }
}
