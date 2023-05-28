import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { RecoilRoot } from 'recoil'

import { setCurrentAccountAdapter, getCurrentAccountAdapter } from '../../main/adapters'
import { currentAccountState } from '../../presentation/components'
import { makeLogin, makeHome, makeCreateAccount } from '../../main/factories/pages'
import { PrivateRoute } from '../proxies'

const Router: React.FC = () => {
  const state = {
    setCurrentAccount: setCurrentAccountAdapter,
    getCurrentAccount: getCurrentAccountAdapter
  }
  return (
    <RecoilRoot initializeState={({ set }) => set(currentAccountState, state)}>
      <div className='h-screen'>
        <BrowserRouter>
          <Switch>
            <Route path="/auth/login" exact component={makeLogin} />
            <Route path="/auth/register" exact component={makeCreateAccount} />
            <PrivateRoute path="/" exact component={makeHome} />
          </Switch>
        </BrowserRouter>
      </div>
    </RecoilRoot>
  )
}

export default Router
