import React, { useRef } from 'react'

type Props = React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement> & {
  state: any
  setState: any
}

const Input: React.FC<Props> = ({ state, setState, ...props }: Props) => {
  const inputRef = useRef<HTMLInputElement>(null)
  const error = state[`${props.name}Error`]
  return (
    <div>
      <label title={error} className="block text-sm font-medium leading-6 text-gray-900 py-2">{props.placeholder}</label>
      <input
        {...props}
        ref={inputRef}
        title={error}
        placeholder={props.placeholder}
        data-testid={props.name}
        readOnly
        className="mb-3 block w-full form-control rounded-md border-0 p-3 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-pink-600 sm:text-sm sm:leading-6"
        data-status={error ? 'invalid' : 'valid'}
        onFocus={e => { e.target.readOnly = false }}
        onChange={e => { setState({ ...state, [e.target.name]: e.target.value }) }}
      />
    </div>
  )
}

export default Input
