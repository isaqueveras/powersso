import { RouteProps, Route, Redirect } from 'react-router-dom'
import { useRecoilValue } from 'recoil'
import React from 'react'

import { currentAccountState } from '../../presentation/components'

const PrivateRoute: React.FC<RouteProps> = (props: RouteProps) => {
  const { getCurrentAccount } = useRecoilValue(currentAccountState)
  return getCurrentAccount()?.jwt_token
    ? <Route {...props} />
    : <Route {...props} component={() => <Redirect to="/auth/login" />} />
}

export default PrivateRoute
