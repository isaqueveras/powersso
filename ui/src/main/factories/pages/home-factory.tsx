import { currentAccountState } from '../../../presentation/components'
import React from 'react'
import { useRecoilValue } from 'recoil'

export const makeHomePage: React.FC = () => {
  const { getCurrentAccount } = useRecoilValue(currentAccountState)
  const name = getCurrentAccount().first_name
  return <h1>Hello, {name}!</h1>
}
