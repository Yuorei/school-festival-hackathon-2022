import * as React from 'react';
import { useContext } from 'react';
import { Navigate } from '@tanstack/react-location';
import { AuthContext } from './context/authContext';




export const AuthWrapper = ({ children }) => {
  const { isLoggedIn, setIsLoggedIn } = useContext(AuthContext);
  console.log(`AuthWrapper: isLoggedIn: ${isLoggedIn}`);
  if (!isLoggedIn) {
    return <Navigate to="/" />;
  } else {
    return children
  }
}