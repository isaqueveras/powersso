import React from 'react'
import { useRecoilValue } from 'recoil'

import { createAccountState } from './atoms'
import { FormStatusBase } from '../../../../presentation/components'

const FormStatus: React.FC = () => {
  const state = useRecoilValue(createAccountState)
  return <FormStatusBase messageError={state.messageError} />
}

export default FormStatus
