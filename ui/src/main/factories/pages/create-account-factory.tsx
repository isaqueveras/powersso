import React from 'react'

import { CreateAccount } from '../../../presentation/pages'
import { makeRemoteCreateAccount } from '../usecases'
import { makeCreateAccountValidation } from '../validation'

export const makeCreateAccount: React.FC = () => {
  return (
    <CreateAccount
      usecase={makeRemoteCreateAccount()}
      validation={makeCreateAccountValidation()}
    />
  )
}
