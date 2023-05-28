import { Err } from "@/domain/models"

export interface Activation {
  activate: (params: Activation.Params) => Promise<Activation.Model>
}

export namespace Activation {
  export type Params = {
    id: string
  }

  export type Model = Err
}