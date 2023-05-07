import React, { useEffect } from 'react'
import { useRecoilState, useRecoilValue, useResetRecoilState } from 'recoil'
import { useHistory } from 'react-router-dom'

import { Validation } from '../../protocols'
import { Authentication } from '../../../domain/usecases'
import { currentAccountState } from '../../components'
import { loginState, Input, FormStatus, SubmitButton } from './components'

type Props = {
  validation: Validation
  authentication: Authentication
}

const Login: React.FC<Props> = ({ validation, authentication }: Props) => {
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
      const account = await authentication.auth({ email: state.email, password: state.password })
      setCurrentAccount(account)
      history.replace('/')
    } catch (error: any) {
      setState(old => ({ ...old, isLoading: false, mainError: error.message }))
    }
  }

  return (
    <section className='h-screen flex justify-center items-center bg-gray-200'>
      <section>
        <section className='mb-4'>
          <h3 className='font-bold text-3xl'>Welcome to PowerSSO</h3>
          <p className='text-gray-600 pt-2'>Sign in to your account.</p>
        </section>
        <section className='flex xl:justify-center lg:justify-between items-center flex-wrap bg-white p-8 rounded py-16'>
          <form data-testid='form' className='w-96 max-w-4xl' onSubmit={handleSubmit}>
            <FormStatus />
            <Input type='text' name='email' placeholder='Email address'/>
            <Input type='password' name='password' placeholder='Password'/>
            <SubmitButton text='Login' />
          </form>
        </section>
      </section>
    </section>
  )
}

export default Login
