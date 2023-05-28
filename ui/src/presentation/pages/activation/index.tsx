import React, { useEffect } from 'react'
import { Activation } from '@/domain/usecases/activation'
import { useRecoilState, useResetRecoilState } from 'recoil'
import { activationState } from './components/atoms'
import FormStatus from './components/form-status'
import { Link } from 'react-router-dom'

type Props = {
  usecase: Activation
}

const ActivationPage: React.FC<Props> = ({ usecase }: Props) => {
  const resetActivationState = useResetRecoilState(activationState)
  const [state, setState] = useRecoilState(activationState)

  useEffect(() => {
    resetActivationState()
    handleSubmit()
  }, [])

  const handleSubmit = async (): Promise<void> => {
    try {
      if (state.isLoading) return
      setState((old) => ({ ...old, isLoading: true }))
      await usecase.activate({ id: 'asdasdasd' })
    } catch (error: any) {
      setState((old) => ({ ...old, isLoading: false, errorMessage: error.message }))
    }
  }

  return (
    <section className='h-screen flex justify-center items-center bg-pink-50'>
      {state.errorMessage === '' && (
        <div className='space-y-6 bg-white p-12 rounded-sm shadow-sm'>
          <p className='mb-4 text-base'>Activation successful! now you can login.</p>
          <Link to='/auth/login'><p className='text-blue-500'>Login</p></Link>
        </div>
      )}
      <FormStatus />
    </section>
  )
}

export default ActivationPage
