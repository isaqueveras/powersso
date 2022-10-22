import { ValidationComposite } from '../../../main/composites'
import { ValidationBuilder } from '../../../main/builders'

export const makeLoginValidation = (): ValidationComposite => ValidationComposite.build([
  ...ValidationBuilder.field('email').required().email().build(),
  ...ValidationBuilder.field('password').required().min(5).build()
])
