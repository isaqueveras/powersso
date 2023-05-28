import React from 'react'
import { useParams } from 'react-router-dom'

import { ActivationPage } from '../../../presentation/pages'
import { makeRemoteActivation } from '../usecases'

export const makeActivationPage: React.FC = () => {
  const { token } = useParams<{ token: string }>()
  return <ActivationPage usecase={makeRemoteActivation(token)}/>
}
