import React from 'react'

type Props = React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement> & {
  state: any
  setState: any
}

const Input: React.FC<Props> = ({ state, setState, ...props }: Props) => {
  const error = state[`${props.name !== undefined ? props.name : ''}Error`]
  return (
    <>
      <label title={error}>{props.placeholder}</label>
      <input
        {...props}
        title={error}
        placeholder={props.placeholder}
        data-testid={props.name}
        readOnly
        className="mb-3 p-3 form-control block w-full bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out" data-status={error ? 'invalid' : 'valid'}
        onFocus={e => { e.target.readOnly = false }}
        onChange={e => { setState({ ...state, [e.target.name]: e.target.value }) }}
      />
    </>
  )
}

export default Input
