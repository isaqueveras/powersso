export class Oops extends Error {
  constructor (message: string) {
    super(message)
    this.name = 'OopsError'
  }
}
