import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { RecoilRoot } from 'recoil'

import { setCurrentAccountAdapter, getCurrentAccountAdapter } from '../../main/adapters'
import { currentAccountState } from '../../presentation/components'
import { makeLogin } from '../../main/factories/pages'

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
            <Route path="auth/login" exact component={makeLogin} />
            {/* <Route path="/" element={<Dashboard />} />
            <Route path="projects" element={<Projects />} />
            <Route path="users" element={<Users />} />
            <Route path="users/new" element={<NewUser />} /> */}
          </Switch>
        </BrowserRouter>
      </div>
    </RecoilRoot>
  )
}

export default Router
