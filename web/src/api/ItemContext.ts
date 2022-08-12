import { createContext } from 'react';

type ItemContextType = {
  itemValue: string;
  saveItem: (value: string) => void;
}

const FormContext = createContext(null)

export default FormContext;
