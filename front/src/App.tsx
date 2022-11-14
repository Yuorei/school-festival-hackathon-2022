import * as React from 'react';
import Box from '@mui/material/Box';

import { Router, Outlet, Link } from '@tanstack/react-location'
import { routes, location } from './routes'

import Header from './components/header'
import Footer from './components/footer'
import BackToUp from './components/backToUp'

function App() {

  return (
    <div className="App">
      <Router routes={routes} location={location}>
        <Header />
        <Box sx={{p:2}}></Box>
        <Outlet />
        <BackToUp />
        <Footer />
      </Router>
    </div>
  )
}

export default App
