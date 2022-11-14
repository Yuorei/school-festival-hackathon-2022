import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';

export type Props = {
  text: string;
  children: React.ReactNode;
};

export const Item : React.FC<Props> = ({ itemcontents } : object) => {
  return (
    <Card sx={{ m: 1}}>
      <CardActionArea>
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            { itemcontents.name }
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Lizards are a widespread group of squamate reptiles, with over 6,000
            species, ranging across all continents except Antarctica
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}
