import * as React from 'react';

import { Link } from '@tanstack/react-location'

import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import ArrowCircleUpOutlinedIcon from '@mui/icons-material/ArrowCircleUpOutlined';
import ArrowCircleDownOutlinedIcon from '@mui/icons-material/ArrowCircleDownOutlined';
import AddIcon from '@mui/icons-material/Add';
import Paper from '@mui/material/Paper';
import Fab from '@mui/material/Fab';


function Footer() {
  const [value, setValue] = React.useState(0);
  const ref = React.useRef<HTMLDivElement>(null);

  return (
    <Box sx={{ pb: 7 }} ref={ref}>
      <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
        <BottomNavigation
          showLabels
          value={value}
          onChange={(event, newValue) => {
            setValue(newValue);
          }}
        >

          <BottomNavigationAction
            label="借りたい"
            icon={<ArrowCircleDownOutlinedIcon />}
            component={Link}
            to="/rent"
          />
          
          <Fab sx={{ mt: 0.5 }} size="medium" color="primary" aria-label="back" component={Link} to="/add">
            <AddIcon />
          </Fab>
          <BottomNavigationAction
            label="貸せる"
            icon={<ArrowCircleUpOutlinedIcon />}
            component={Link}
            to="/lend"
          />
        </BottomNavigation>
      </Paper>
    </Box>
  );
}

export default Footer
