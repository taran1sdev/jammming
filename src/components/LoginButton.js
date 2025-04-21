import React from 'react';

import spotifyLogo from '../resources/spotify-logo.png';

function LoginButton() {
    return (
        <>
            <a href='#'
            onClick={() => {window.location.href='http://localhost:5000/auth/login'}}>
                <div className='Login'>
                    <img className='Logo' src={spotifyLogo} alt='spotify logo'/>
                    <h4>Click to Login!</h4>
                </div>
            </a>
            
        </>
    );
}

export default LoginButton;