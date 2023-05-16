import { HttpClient, HttpStatusCode } from '../protocols/http'
import { RegisterUser } from '../../domain/usecases'
import { Oops } from '../../domain/errors'
import { Err } from '@/domain/models'

export class RemoteRegisterUser implements RegisterUser {
  constructor (
    private readonly url: string,
    private readonly httpClient: HttpClient<Err>
  ) {}

  async register (params: RegisterUser.Params): Promise<RegisterUser.Model> {
    const httpResponse = await this.httpClient.request({ url: this.url, method: 'post', body: params })
    switch (httpResponse.statusCode) {
      case HttpStatusCode.ok: return httpResponse.body
      default: throw new Oops(httpResponse.body.message)
    }
  }
}
