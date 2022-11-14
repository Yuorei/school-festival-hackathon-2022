import { createContext } from 'react';

const initialState = {
  isLoggedIn: false,
}
export const AuthContext = createContext(initialState);