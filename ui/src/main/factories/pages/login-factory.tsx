import React from 'react'

import { makeLoginValidation } from '../../../main/factories/validation'
import { makeRemoteAuthentication } from '../../../main/factories/usecases'
import { Login } from '../../../presentation/pages'

export const makeLogin: React.FC = () => {
  return (
    <Login
      usecase={makeRemoteAuthentication()}
      validation={makeLoginValidation()}
    />
  )
}
