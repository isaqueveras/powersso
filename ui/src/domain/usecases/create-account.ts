import { Err } from '@/domain/models'

export interface CreateAccount {
  register: (params: CreateAccount.Params) => Promise<CreateAccount.Model>
}

export namespace CreateAccount {
  export type Params = {
    first_name: string
    last_name: string
    email: string
    password: string
  }

  export type Model = Err
}