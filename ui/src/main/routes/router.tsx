import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { RecoilRoot } from 'recoil'

import { PrivateRoute } from '../proxies'
import { setCurrentAccountAdapter, getCurrentAccountAdapter } from '../../main/adapters'
import { currentAccountState } from '../../presentation/components'
import { makeLoginPage, makeHomePage, makeCreateAccountPage, makeActivationPage } from '../../main/factories/pages'

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
            <Route path="/auth/login" exact component={makeLoginPage} />
            <Route path="/auth/register" exact component={makeCreateAccountPage} />
            <Route path="/auth/activation/:token" exact component={makeActivationPage} />
            <PrivateRoute path="/" exact component={makeHomePage} />
          </Switch>
        </BrowserRouter>
      </div>
    </RecoilRoot>
  )
}

export default Router
