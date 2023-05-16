import { atom } from 'recoil'

export const createAccountState = atom({
  key: 'createAccountState',
  default: {
    isLoading: false,
    isFormInvalid: true,
    email: '',
    password: '',
    confirmPassword: '',
    firstName: '',
    lastName: '',
    messageError: '',
    passwordError: '',
    firstNameError: '',
    lastNameError: '',
    confirmPasswordError: '',
    mainError: ''
  }
})
