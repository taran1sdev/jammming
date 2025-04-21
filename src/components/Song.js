import React from 'react';


export function Song({track, handleClick, search}) {   
    if(track){
        return (
        <div className='Song' onClick={handleClick} id={track.id}>
            <img src={track.img} id={track.id} alt='Album art'/>
            <div className='Info' id={track.id}>
                <h4 id={track.id}>{
                track.name.length > 25 ? track.name.slice(0, 25) : track.name
                }</h4>
                <p id={track.id}>{track.artist}</p>
            </div>
            {
              /*  !search ? <div className='Play'>
                    {
                        trackPlaying === track.id ? <img src={require('../resources/pause.png')} name='pause' id={track.id} />
                        : <img src={require('../resources/play.png')} name='play' id={track.id} />
                    }  
                </div> 
                : <></> */
            }
        </div>
    )}
};

export default Song;