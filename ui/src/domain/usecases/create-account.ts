import { Err } from "../models"

export interface RegisterUser {
  register: (params: RegisterUser.Params) => Promise<RegisterUser.Model>
}

export namespace RegisterUser {
  export type Params = {
    first_name: string
    last_name: string
    email: string
    password: string
  }

  export type Model = Err
}