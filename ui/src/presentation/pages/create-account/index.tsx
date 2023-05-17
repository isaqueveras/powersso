import React, { useEffect } from 'react'
import { useHistory } from 'react-router-dom'
import { useRecoilState, useResetRecoilState } from 'recoil'

import { CreateAccount } from '@/domain/usecases'
import { Validation } from '@/presentation/protocols'
import { createAccountState, FormStatus, Input, SubmitButton } from './components'

type Props = {
  validation: Validation
  usecase: CreateAccount
}

const CreateAccountPage: React.FC<Props> = ({ validation, usecase }: Props) => {
  const resetCreateAccountState = useResetRecoilState(createAccountState)
  const [state, setState] = useRecoilState(createAccountState)

  useEffect(() => resetCreateAccountState(), [])
  useEffect(() => validate('fisrtName'), [state.firstName])
  useEffect(() => validate('lastName'), [state.lastName])
  useEffect(() => validate('email'), [state.email])
  useEffect(() => validate('password'), [state.password])
  useEffect(() => validate('confirmPassword'), [state.confirmPassword])

  const validate = (field: string): void => {
    const { firstName: fisrtName, lastName, email, password, confirmPassword } = state
    const formData = { fisrtName, lastName, email, password, confirmPassword }
    setState((old) => ({
      ...old,
      [`${field}Error`]: validation.validate(field, formData)
    }))
    setState((old) => ({
      ...old,
      isFormInvalid:
        !!old.firstNameError ||
        !!old.lastNameError ||
        !!old.messageError ||
        !!old.passwordError ||
        !!old.confirmPasswordError
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
      await usecase.register({
        email: state.email,
        password: state.password,
        first_name: state.firstName,
        last_name: state.lastName
      })
      history.replace('/')
    } catch (error: any) {
      setState((old) => ({ ...old, isLoading: false, messageError: error.message }))
    }
  }

  return (
    <section className='h-screen flex justify-center items-center bg-pink-50'>
      <form data-testid='form-create-account' onSubmit={handleSubmit}>
        <div className='space-y-6 bg-white p-12 rounded-md shadow-md'>
          <div className='border-b border-gray-900/10 pb-6'>
            <h2 className='text-2xl font-semibold leading-7 text-gray-900'>
              Create account
            </h2>
            <p className='mt-1 text-sm leading-6 text-gray-600 mb-6'>
              Use a permanent address where you can receive mail.
            </p>
            <div className='border-t border-gray-900/10 pt-3 grid grid-cols-1 gap-x-6 sm:grid-cols-6'>
              <div className="sm:col-span-6">
                <FormStatus />
              </div>
              <div className='sm:col-span-3'>
                <Input name='firstName' type='text' placeholder='First name' />
              </div>
              <div className='sm:col-span-3'>
                <Input name='lastName' type='text' placeholder='Last name' />
              </div>
              <div className='sm:col-span-6'>
                <Input name='email' type='email' placeholder='E-mail' />
              </div>
              <div className='sm:col-span-6'>
                <Input name='password' type='password' placeholder='Password' />
              </div>
              <div className='sm:col-span-6'>
                <Input
                  name='confirmPassword'
                  type='password'
                  placeholder='Confirm password'
                />
              </div>
            </div>
          </div>
          <SubmitButton text='Create Account' />
        </div>
      </form>
    </section>
  )
}

export default CreateAccountPage
