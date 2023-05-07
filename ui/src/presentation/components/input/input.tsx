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
      <label title={error} className="text-gray-600 py-2">{props.placeholder}</label>
      <input
        {...props}
        ref={inputRef}
        title={error}
        placeholder={props.placeholder}
        data-testid={props.name}
        readOnly
        className="mb-3 p-3 form-control block w-full border border-solid rounded transition ease-in-out" data-status={error ? 'invalid' : 'valid'}
        onFocus={e => { e.target.readOnly = false }}
        onChange={e => { setState({ ...state, [e.target.name]: e.target.value }) }}
      />
    </div>
  )
}

export default Input
