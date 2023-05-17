import { HttpClient, HttpStatusCode } from '../../data/protocols/http'
import { CreateAccount } from '../../domain/usecases'
import { Oops } from '../../domain/errors'

export class RemoteCreateAccount implements CreateAccount {
  constructor (
    private readonly url: string,
    private readonly httpClient: HttpClient<CreateAccount.Model>
  ) {}

  async register (params: CreateAccount.Params): Promise<CreateAccount.Model> {
    const httpResponse = await this.httpClient.request({ url: this.url, method: 'post', body: params })
    switch (httpResponse.statusCode) {
      case HttpStatusCode.ok: return httpResponse.body
      default: throw new Oops(httpResponse.body.message)
    }
  }
}
