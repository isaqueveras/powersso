import React from 'react'

import { useRecoilValue } from 'recoil'
import { createAccountState } from './atoms'
import { SubmitButtonBase } from '../../../../presentation/components'

type Props = {
  text: string
}

const SubmitButton: React.FC<Props> = ({ text }: Props) => {
  return <SubmitButtonBase text={text} state={useRecoilValue(createAccountState)} />
}

export default SubmitButton
