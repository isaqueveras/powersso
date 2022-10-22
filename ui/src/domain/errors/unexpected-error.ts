export class UnexpectedError extends Error {
  constructor (message: string | undefined) {
    super(message !== undefined ? message : 'Something went wrong. Please try again soon.')
    this.name = 'UnexpectedError'
  }
}
