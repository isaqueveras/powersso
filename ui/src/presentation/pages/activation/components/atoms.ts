import { atom } from 'recoil'

export const activationState = atom({
  key: 'activationState',
  default: {
    isLoading: false,
    errorMessage: ''
  }
})
