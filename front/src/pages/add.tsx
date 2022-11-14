import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Grid from '@mui/material/Grid';
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormControl from '@mui/material/FormControl';
import FormLabel from '@mui/material/FormLabel';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';



function Add() {

  const [value, setValue] = React.useState('');
  const [date, setDate] = React.useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value);
  };

  return (
    <Card sx={{ mx: 3, p: 1 ,}} component="form">
      <CardContent>
        <Grid container spacing={4} alignItems="center" justifyContent="center">
        <Grid item xs={10}>
        <RadioGroup
          row
          aria-labelledby="demo-row-radio-buttons-group-label"
          name="row-radio-buttons-group"
        >
          <FormControlLabel value="rent" control={<Radio />} label="借りたい" />
          <FormControlLabel value="lend" control={<Radio />} label="貸せる" />
        </RadioGroup>
  </Grid>
  <Grid item xs={10}>
  <TextField id="standard-basic" fullWidth label="商品名" variant="standard" />
  </Grid>
  <Grid item xs={10}>
  <TextField
          id="outlined-multiline-flexible"
          label="説明"
          multiline
          fullWidth
              maxRows={4}
              value={value}
          />
          </Grid>
          <Grid item xs={10}>
          <Button variant="contained" component="label">
  画像をアップロード
  <input hidden accept="image/*" multiple type="file" />
            </Button>
          </Grid>
          <Grid item xs={10}>
            <LocalizationProvider dateAdapter={AdapterDateFns}>
      <DatePicker
        label="掲載期限"
        value={date}
        onChange={(newDate) => {
          setDate(newDate);
        }}
        renderInput={(params) => <TextField {...params} />}
      />
            </LocalizationProvider>
          </Grid>
          <Grid item xs={10}>
            <Button variant="contained">登録</Button>
            </Grid>
          </Grid>
      </CardContent>
    </Card>
  );
}

export default Add
