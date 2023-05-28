import React from 'react'

import { CreateAccountPage } from '../../../presentation/pages'
import { makeRemoteCreateAccount } from '../usecases'
import { makeCreateAccountValidation } from '../validation'

export const makeCreateAccountPage: React.FC = () => {
  return (
    <CreateAccountPage
      usecase={makeRemoteCreateAccount()}
      validation={makeCreateAccountValidation()}
    />
  )
}
