import * as React from 'react';

import SignInScreen from '../components/signIn'


function Login() {

  const [value, setValue] = React.useState('Controlled');

  return (
    <SignInScreen />
    
  );
}

export default Login
