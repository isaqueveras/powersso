import React, { useEffect } from 'react'
import { useRecoilState, useRecoilValue, useResetRecoilState } from 'recoil'
import { useHistory } from 'react-router-dom'

import { Validation } from '../../protocols'
import { Authentication } from '../../../domain/usecases'
import { loginState } from './components'
import { currentAccountState } from '../../components'

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

  const navigate = useHistory()
  const validate = (field: string): void => {
    const { email, password } = state
    const formData = { email, password }
    setState(old => ({ ...old, [`${field}Error`]: validation.validate(field, formData) }))
    setState(old => ({ ...old, isFormInvalid: !!old.emailError || !!old.passwordError }))
  }

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault()
    try {
      if (state.isLoading || state.isFormInvalid) {
        return
      }

      setState(old => ({ ...old, isLoading: true }))
      const account = await authentication.auth({
        email: state.email,
        password: state.password
      })

      if (account !== undefined) {
        setCurrentAccount(account)
      }
      navigate.replace('/')
    } catch (error: any) {
      setState(old => ({
        ...old,
        isLoading: false,
        mainError: error.message
      }))
    }
  }

  return (
    <section className="h-screen">
      <div className="px-6 h-full text-gray-800">
        <div className="flex xl:justify-center lg:justify-between justify-center items-center flex-wrap h-full g-6">
          <div className="xl:ml-20 xl:w-5/12 lg:w-5/12 md:w-8/12 mb-12 md:mb-0">
            <form data-testid="form" onSubmit={handleSubmit}>
              <input type="text" className="mb-6 form-control block w-full px-4 py-2 text-md font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-black focus:outline-none" placeholder="Email address"/>
              <input type="password" className="mb-6 form-control block w-full px-4 py-2 text-md font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-black focus:outline-none" placeholder="Password"/>
              <div className="flex justify-between items-center mb-6">
                <div className="form-group form-check">
                  <input type="checkbox" className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-black checked:border-black focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer" id="rememberMe" />
                  <label className="form-check-label inline-block text-gray-800" htmlFor="rememberMe">Remember me</label>
                </div>
                <a href="#!" className="text-gray-800">Forgot password?</a>
              </div>
              <div className="text-center lg:text-left">
                <button type="submit" className="w-full inline-block px-7 py-3 bg-black text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:bg-black hover:shadow-lg focus:black focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out">Login</button>
                <p className="text-sm font-semibold mt-2 pt-1 mb-0">
                  {"Don't have an account?"}
                  <a href="#!" className="text-red-600 hover:text-red-700 focus:text-red-700 transition duration-200 ease-in-out ml-1">Register</a>
                </p>
              </div>
            </form>
          </div>
        </div>
      </div>
    </section>
  )
}

export default Login
