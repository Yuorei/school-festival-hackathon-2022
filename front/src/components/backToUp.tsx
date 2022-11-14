import * as React from 'react';

import Fab from '@mui/material/Fab';
import ExpandLessIcon from '@mui/icons-material/ExpandLess';


export default function backToUp() {
  const returnTop = () => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    })
  }

  return (
    <Fab sx={{position: 'fixed', bottom: 80, right: 30,}} color="primary" aria-label="back" onClick={returnTop}>
      <ExpandLessIcon />
    </Fab>
  );
}