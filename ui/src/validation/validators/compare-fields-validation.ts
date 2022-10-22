import { FieldValidation } from '../../validation/protocols'
import { InvalidFieldError } from '../../validation/errors'

export class CompareFieldsValidation implements FieldValidation {
  constructor (
    readonly field: string,
    private readonly fieldToCompare: string
  ) {}

  validate (input: any): Error | null {
    return input[this.field] !== input[this.fieldToCompare] ? new InvalidFieldError() : null
  }
}
