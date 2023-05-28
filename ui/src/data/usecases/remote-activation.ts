import { HttpClient, HttpStatusCode } from '../protocols/http'
import { Activation } from '../../domain/usecases'
import { Err } from '../../domain/errors'

export class RemoteActivation implements Activation {
  constructor (
    private readonly url: string,
    private readonly httpClient: HttpClient<Activation.Model>
  ) {}

  async activate (params: Activation.Params): Promise<Activation.Model> {
    const httpResponse = await this.httpClient.request({ url: this.url, method: 'post', body: params })
    switch (httpResponse.statusCode) {
      case HttpStatusCode.ok: return httpResponse.body
      default: throw new Err(httpResponse.body.message)
    }
  }
}
