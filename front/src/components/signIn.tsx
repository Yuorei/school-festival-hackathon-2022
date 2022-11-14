import React, {useContext, useState, useEffect} from 'react';
import firebase from '../firebase';
import 'firebase/compat/auth';
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth';

const uiConfig = {
    signInFlow: 'popup',
    signInSuccessUrl: "/",
    signInOptions: [
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        //firebase.auth.FacebookAuthProvider.PROVIDER_ID,
        //firebase.auth.TwitterAuthProvider.PROVIDER_ID,
        //firebase.auth.GithubAuthProvider.PROVIDER_ID,
        //firebase.auth.EmailAuthProvider.PROVIDER_ID,
        //firebase.auth.PhoneAuthProvider.PROVIDER_ID,
        //firebaseui.auth.AnonymousAuthProvider.PROVIDER_ID
    ],
}

import { AuthContext } from '../context/authContext'


const SignInScreen = (props) => {
    const { isLoggedIn, setIsLoggedIn } = useContext(AuthContext)
    useEffect(() => {
        firebase.auth().onAuthStateChanged(user => {
            setIsLoggedIn(true);
        });
    }, []);

    return (
        <div>
            <p>Please sign-in:</p>
            <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebase.auth()} />
        </div>
    );
}
export default SignInScreen;
