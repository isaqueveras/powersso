import { AccountModel } from '@/domain/models'

import { atom } from 'recoil'

export const currentAccountState = atom({
  key: 'currentAccountState',
  default: {
    getCurrentAccount: null as unknown as () => AccountModel,
    setCurrentAccount: null as unknown as (account: AccountModel) => void
  }
})
