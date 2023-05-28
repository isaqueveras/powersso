import React from 'react'
import { useRecoilValue } from 'recoil'

import { activationState } from './atoms'
import { FormStatusBase } from '../../../../presentation/components'

const FormStatus: React.FC = () => {
  const state = useRecoilValue(activationState)
  return <FormStatusBase errorMessage={state.errorMessage} />
}

export default FormStatus
