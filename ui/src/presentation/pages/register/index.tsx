import { RegisterUser } from '@/domain/usecases'
import { Validation } from '@/presentation/protocols'
import { useRecoilState, useResetRecoilState } from 'recoil'
import { registerUserState } from './components/atoms'
import React, { useEffect } from 'react'
import { useHistory } from 'react-router-dom'

type Props = {
  validation: Validation
  usecase: RegisterUser
}

const Register: React.FC<Props> = ({
  validation,
  usecase: registerUser
}: Props) => {
  const resetRegisterUserState = useResetRecoilState(registerUserState)
  const [state, setState] = useRecoilState(registerUserState)

  useEffect(() => resetRegisterUserState(), [])
  useEffect(() => validate('email'), [state.email])
  useEffect(() => validate('password'), [state.password])

  const validate = (field: string): void => {
    const { email, password } = state
    const formData = { email, password }
    setState((old) => ({
      ...old,
      [`${field}Error`]: validation.validate(field, formData)
    }))
    setState((old) => ({
      ...old,
      isFormInvalid: !!old.emailError || !!old.passwordError
    }))
  }

  const history = useHistory()

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault()
    try {
      if (state.isLoading || state.isFormInvalid) return
      setState((old) => ({ ...old, isLoading: true }))
      await registerUser.register({
        email: state.email,
        password: state.password,
        fisrt_name: state.fisrt_name,
        last_name: state.last_name
      })
      history.replace('/')
    } catch (error) {
      setState((old) => ({ ...old, isLoading: false }))
    }
  }

  return (
    <section className='h-screen flex justify-center items-center bg-gray-100'>
      <form data-testid='form-register-user' onSubmit={handleSubmit}>
        <div className="space-y-12 bg-white p-12 rounded-md shadow-md">
          <div className="border-b border-gray-900/10 pb-6">
            <h2 className="text-2xl font-bold leading-7 text-gray-900">Create account</h2>
            <p className="mt-1 text-sm leading-6 text-gray-600 mb-6">Use a permanent address where you can receive mail.</p>
            <div className="border-t border-gray-900/10 pt-6 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
              <div className="sm:col-span-3">
                <label className="block text-sm font-medium leading-6 text-gray-900">First name</label>
                <div className="mt-2">
                  <input type="text" name="first-name" id="first-name" autoComplete="given-name" className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
                </div>
              </div>

              <div className="sm:col-span-3">
                <label className="block text-sm font-medium leading-6 text-gray-900">Last name</label>
                <div className="mt-2">
                  <input type="text" name="last-name" id="last-name" autoComplete="family-name" className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
                </div>
              </div>

              <div className="sm:col-span-6">
                <label className="block text-sm font-medium leading-6 text-gray-900">Email address</label>
                <div className="mt-2">
                  <input id="email" name="email" type="email" autoComplete="email" className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
                </div>
              </div>

              <div className="sm:col-span-6">
                <label className="block text-sm font-medium leading-6 text-gray-900">Password</label>
                <div className="mt-2">
                  <input name="password" type="password" className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
                </div>
              </div>

              <div className="sm:col-span-6">
                <label className="block text-sm font-medium leading-6 text-gray-900">Confirm password</label>
                <div className="mt-2">
                  <input name="confirm-password" type="password" className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
                </div>
              </div>
            </div>
          </div>
          <button type="submit" className="block rounded-md bg-pink-600 w-full px-3 py-3 text-sm font-semibold text-white shadow-sm hover:bg-pink-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-pink-600">Create account</button>
        </div>
      </form>
    </section>
  )
}

export default Register
