import { makeApiUrl, makeAxiosHttpClient } from '../../../main/factories/http'
import { RemoteRegisterUser } from '../../../data/usecases'
import { RegisterUser } from '../../../domain/usecases'

export const makeRemoteCreateAccount = (): RegisterUser =>
  new RemoteRegisterUser(makeApiUrl('auth/register'), makeAxiosHttpClient())
