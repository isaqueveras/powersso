import { makeApiUrl, makeAxiosHttpClient } from '../http'
import { RemoteActivation } from '../../../data/usecases'
import { Activation } from '../../../domain/usecases'

export const makeRemoteActivation = (token: string): Activation =>
  new RemoteActivation(makeApiUrl(`auth/activation/${token}`), makeAxiosHttpClient())
