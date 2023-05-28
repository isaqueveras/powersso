import React from 'react'

import { makeLoginValidation } from '../../../main/factories/validation'
import { makeRemoteAuthentication } from '../../../main/factories/usecases'
import { LoginPage } from '../../../presentation/pages'

export const makeLoginPage: React.FC = () => {
  return (
    <LoginPage
      usecase={makeRemoteAuthentication()}
      validation={makeLoginValidation()}
    />
  )
}
