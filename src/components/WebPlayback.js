import React, {useEffect, useState} from "react";

import getAccessToken from "../apiCalls/getAccessToken"; 
import transferPlayback from "../apiCalls/transferPlayback";

function WebPlayback() {
    const [player, setPlayer] = useState(undefined);

    useEffect(() => {
        const script = document.createElement('script');
        script.src = 'https://sdk.scdn.co/spotify-player.js';
        script.async = true;
    
        document.body.appendChild(script);
    
        window.onSpotifyWebPlaybackSDKReady = () => {
    
            const player = new window.Spotify.Player({
                name: "Web Playback SDK",
                getOAuthToken: cb => { cb(getAccessToken()); },
                volume: 0.5
            });
    
            setPlayer(player);
    
            player.addListener('ready', ({ device_id }) => {
                console.log('SDK ready with Device ID: ', device_id);
                transferPlayback(device_id).then((success) => {
                    if(success) {                
                        console.log("Device playback transferred successfully");
                    } else {
                        console.log("An error occured while transferring playback");
                    }
                })
            });
    
            player.addListener('not_ready', ({ device_id }) => {
                console.log('Device ID has gone offline: ', device_id);
            });
    
            player.connect();
        };
    }, []);

    return (
        <div className="WebPlayback">
        </div>
    )
};

export default WebPlayback;