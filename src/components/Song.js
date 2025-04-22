import React from 'react';


export function Song({track, handleClick, search, trackPlaying}) {   
    if(track){
        return (
        <div className='Song' onClick={handleClick} id={track.uri}>
            <img src={track.img} id={track.uri} alt='Album art'/>
            <div className='Info' id={track.uri}>
                <h4 id={track.uri}>{
                track.name.length > 25 ? track.name.slice(0, 25) : track.name
                }</h4>
                <p id={track.uri}>{track.artist}</p>
            </div>
            {
               !search ? <div className='Play'>
                    {
                        trackPlaying === track.uri ? <img src={require('../resources/pause.png')} name='pause' 
                        id={track.uri} alt='pause button' />
                        : <img src={require('../resources/play.png')} name='play' id={track.uri} alt='play button' />
                    }  
                </div> 
                : <></> 
            }
        </div>
    )}
};

export default Song;