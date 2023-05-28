export class Err extends Error {
  constructor (message: string) {
    super(message)
    this.name = 'ErrError'
  }
}
