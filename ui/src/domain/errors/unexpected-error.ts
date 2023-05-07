export class UnexpectedError extends Error {
  constructor () {
    super('Something went wrong. Please try again soon.')
    this.name = 'UnexpectedError'
  }
}
