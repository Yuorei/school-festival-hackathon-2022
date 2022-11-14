import * as React from 'react';
import { Route, ReactLocation } from '@tanstack/react-location'
import Lend from './pages/lend'
import Rent from './pages/rent'
import Add from './pages/add'
import Login from './pages/login'
import Protected from './components/protected'
 
export const location = new ReactLocation()
let isProtedted = false

export const routes: Route[] = [
  { path: '/', element: <Rent /> },
  { path: '/lend', element: <Lend /> },
  { path: '/rent', element:<Rent /> },
  { path: '/add', element: <Add />},
  { path: '/login', element: <Login /> },
]