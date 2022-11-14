import * as React from 'react';

import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import { spacing } from '@mui/system';

import { Item } from '../components/item'
import useSWR from 'swr'
const fetcher = (url: string): Promise<any> => fetch(process.env.BASE_URL + url).then(res => res.json());


function getLendList() {
  const { data, error } = useSWR('/api/user/123', fetcher)

  if (error) return <div>failed to load</div>
  if (!data) return <div>loading...</div>
}

function Lend() {
  

  const items : Array<{id: number, user_id: number, name: string, comment:string, image_url:string, deadline:string}> = [
    {
      id: 1,
      user_id: 1,
      name:  "LEND",
      comment:  "test",
      image_url: "http://example.com",
      deadline: "2021-10-10"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    },
    {
      id: 2,
      user_id: 2,
      name:  "test2",
      comment:  "test2",
      image_url: "http://example.com",
      deadline: "2021-10-11"
    }
  ]

  return (
    <Box sx={{mx: 'auto'}}>
      <Grid container spacing={0} alignItems="center" justifyContent="center">
        {items.map((item: object) => (
          <Grid item xs={4} md={3} lg={2}>
            <Item sx={{mx: 'auto'}} itemcontents={item} />
          </Grid>  
        ))}
      </Grid>
    </Box>
  );
}

export default Lend
