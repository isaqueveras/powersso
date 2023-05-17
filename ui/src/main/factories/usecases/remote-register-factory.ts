import { makeApiUrl, makeAxiosHttpClient } from '../../../main/factories/http'
import { RemoteCreateAccount } from '../../../data/usecases'
import { CreateAccount } from '../../../domain/usecases'

export const makeRemoteCreateAccount = (): CreateAccount =>
  new RemoteCreateAccount(makeApiUrl('auth/register'), makeAxiosHttpClient())
