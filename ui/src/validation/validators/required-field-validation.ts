import { FieldValidation } from '../../validation/protocols'
import { RequiredFieldError } from '../../validation/errors'

export class RequiredFieldValidation implements FieldValidation {
  constructor (readonly field: string) {}

  validate (input: any): Error | null {
    return input[this.field] ? null : new RequiredFieldError()
  }
}
