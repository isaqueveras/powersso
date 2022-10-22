import React from 'react'
import { useRecoilValue } from 'recoil'

import { loginState } from './atoms'
import { FormStatusBase } from '../../../../presentation/components'

const FormStatus: React.FC = () => {
  return <FormStatusBase state={useRecoilValue(loginState)} />
}

export default FormStatus
