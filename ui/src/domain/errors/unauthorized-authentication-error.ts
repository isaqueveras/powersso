export class UnauthorizedAuthenticationError extends Error {
  constructor () {
    super('Unauthorized authentication')
    this.name = 'UnauthorizedAuthentication'
  }
}
