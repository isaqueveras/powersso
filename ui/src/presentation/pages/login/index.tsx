import React, { useEffect } from 'react'
import { useRecoilState, useRecoilValue, useResetRecoilState } from 'recoil'
import { useHistory } from 'react-router-dom'

import { Validation } from '../../protocols'
import { Authentication } from '../../../domain/usecases'
import { currentAccountState } from '../../components'
import { loginState, Input, FormStatus, SubmitButton } from './components'

type Props = {
  validation: Validation
  usecase: Authentication
}

const LoginPage: React.FC<Props> = ({ validation, usecase }: Props) => {
  const resetLoginState = useResetRecoilState(loginState)
  const [state, setState] = useRecoilState(loginState)
  const { setCurrentAccount } = useRecoilValue(currentAccountState)

  useEffect(() => resetLoginState(), [])
  useEffect(() => validate('email'), [state.email])
  useEffect(() => validate('password'), [state.password])

  const history = useHistory()
  const validate = (field: string): void => {
    const { email, password } = state
    const formData = { email, password }
    setState(old => ({ ...old, [`${field}Error`]: validation.validate(field, formData) }))
    setState(old => ({ ...old, isFormInvalid: !!old.emailError || !!old.passwordError }))
  }

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault()
    try {
      if (state.isLoading || state.isFormInvalid) return
      setState(old => ({ ...old, isLoading: true }))
      const account = await usecase.auth({ email: state.email, password: state.password })
      setCurrentAccount(account)
      history.replace('/')
    } catch (error: any) {
      setState(old => ({ ...old, isLoading: false, messageError: error.message }))
    }
  }

  return (
    <section className='h-screen flex justify-center items-center bg-pink-50'>
      <form data-testid='form' onSubmit={handleSubmit}>
        <div className="space-y-6 bg-white p-12 rounded-md shadow-md">
          <div className="border-b border-gray-900/10 pb-6">
            <h2 className="text-2xl font-bold leading-7 text-gray-900">Welcome to PowerSSO</h2>
            <p className="mt-1 text-sm leading-6 text-gray-600 mb-6">Sign in to your account.</p>
            <div className='border-t border-gray-900/10 pt-3 grid grid-cols-1 gap-x-6 sm:grid-cols-6'>
              <div className="sm:col-span-6">
                <FormStatus />
              </div>
              <div className="sm:col-span-6">
                <Input type='text' name='email' placeholder='Email address'/>
              </div>
              <div className="sm:col-span-6">
                <Input type='password' name='password' placeholder='Password'/>
              </div>
            </div>
          </div>
          <SubmitButton text='Login' />
        </div>
      </form>
    </section>
  )
}

export default LoginPage
