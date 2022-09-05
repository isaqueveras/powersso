export interface FieldValidation {
  field: string
  validate: (input: any) => Error | null
}
