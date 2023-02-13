import React, { useEffect } from 'react'
import { useRecoilState, useRecoilValue, useResetRecoilState } from 'recoil'
import { useHistory } from 'react-router-dom'

import { Validation } from '../../protocols'
import { Authentication } from '../../../domain/usecases'
import { currentAccountState } from '../../components'
import { loginState, Input, FormStatus } from './components'

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
      const account = await authentication.auth({
        email: state.email,
        password: state.password
      })
      setCurrentAccount(account)
      history.replace('/')
    } catch (error: any) {
      setState(old => ({
        ...old,
        isLoading: false,
        mainError: error.message
      }))
    }
  }

  return (
    <section className="h-screen flex justify-center items-center">
      <div className="flex xl:justify-center lg:justify-between items-center flex-wrap">
        <form data-testid="form" className='w-96 max-w-4xl' onSubmit={handleSubmit}>
          <Input type="text" name="email" placeholder="Email address"/>
          <Input type="password" name="password" placeholder="Password"/>
          {/* <div className="flex justify-between items-center my-2">
            <div className="form-group form-check">
              <input type="checkbox" className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-black checked:border-black focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer" id="rememberMe" />
              <label className="form-check-label inline-block text-gray-800" htmlFor="rememberMe">Remember me</label>
            </div>
            <a href="#!" className="text-gray-800">Forgot password?</a>
          </div> */}
          <div className="text-center lg:text-left mb-3">
            <button type="submit" className="w-full inline-block px-7 py-3 bg-black text-white font-medium text-sm leading-snug uppercase rounded hover:bg-black hover:shadow-lg focus:outline-none focus:ring-0 transition duration-150 ease-in-out">Login</button>
            <p className="text-sm pt-1">
              {"Don't have an account?"}
              <a href="#!" className="text-blue-600 hover:text-blue-700 focus:text-red-700 transition duration-200 ease-in-out ml-1">Register</a>
            </p>
          </div>
          <FormStatus />
        </form>
      </div>
    </section>
  )
}

export default Login
