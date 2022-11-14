import React, { createContext, useState } from 'react';
import Box from '@mui/material/Box';

import { Router, Outlet, Link } from '@tanstack/react-location'
import { routes, location } from './routes'

import Header from './components/header'
import Footer from './components/footer'
import BackToUp from './components/backToUp'

import { AuthContext } from './context/authContext'

function App() {
  const isLoggedIn = useState(false);

  return (
    <div className="App">
      <AuthContext.Provider value={isLoggedIn}>
        <Router routes={routes} location={location}>
          <Header />
          <Box sx={{p:2}}></Box>
          <Outlet />
          <BackToUp />
          <Footer />  
        </Router>
      </AuthContext.Provider>
    </div>
  )
}

export default App
