import { ValidationComposite } from '../../../main/composites'
import { ValidationBuilder } from '../../../main/builders'

export const makeCreateAccountValidation = (): ValidationComposite => ValidationComposite.build([
  ...ValidationBuilder.field('firstName').required().build(),
  ...ValidationBuilder.field('lastName').required().build(),
  ...ValidationBuilder.field('email').required().email().build(),
  ...ValidationBuilder.field('password').required().min(5).build(),
  ...ValidationBuilder.field('confirmPassword').required().min(5).sameAs('password').build()
])
