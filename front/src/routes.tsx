import { useContext } from 'react';
import { Route, ReactLocation, Navigate } from '@tanstack/react-location'
import Lend from './pages/lend'
import Rent from './pages/rent'
import Add from './pages/add'
import Login from './pages/login'

import { AuthWrapper } from './authWrapper';
 
export const location = new ReactLocation()

export const routes: Route[] = [
  { path: '/', element: (<AuthWrapper><Rent /></AuthWrapper>) },
  { path: '/lend', element: (<AuthWrapper><Lend /></AuthWrapper>) },
  { path: '/rent', element:(<AuthWrapper><Rent /></AuthWrapper>) },
  { path: '/add', element: (<AuthWrapper><Add /></AuthWrapper>)},
  { path: '/login', element: <Login /> },
]