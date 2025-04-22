import React, {useEffect, useState} from 'react';

import LoginButton from '../components/LoginButton';
import SearchBar from '../components/SearchBar';
import {Song} from '../components/Song'; 

import searchSpotify from '../apiCalls/searchSpotify';
import playTrack from '../apiCalls/playTrack';
import pausePlayback from '../apiCalls/pausePlayback';

function SearchContainer({selectedSongs, setSelectedSongs, authenticated}) {
    const [searchTerm, setSearchTerm] = useState('');
    const [searchResults, setSearchResults] = useState([]);

    useEffect(() => {
        if(searchTerm){
            searchSpotify(searchTerm).then((results) => {
                const tracks = results.map((track) => {
                    return {
                        uri: track.URI,
                        name: track.Name,
                        artist: track.Artist,
                        img: track.Image
                    }
                })
                
                setSearchResults(tracks);
            });
        }
    },[searchTerm])

    const handleTextInput = (e) => setSearchTerm(e.target.value);

    const [trackPlaying, setTrackPlaying] = useState('');

    const handleClick = (e) => {
        
        if(e.target.name === "play"){
            playTrack(e.target.id).then((result) => {
                if (result) {
                    setTrackPlaying(e.target.id);
                    console.log("Playback Successful");
                } else {
                    console.log("An error occured during playback");
                }
            });
        } else if (e.target.name === "pause") {
            pausePlayback().then((result) => {
                if (result) {
                    setTrackPlaying('');
                    console.log("Playback paused");
                } else {
                    console.log("An error occured pausing playback");
                }
            })
        } else {
            const trackToAdd = searchResults.find((track) => {
                return track.uri === e.target.id; 
            })
    
            if(!selectedSongs.includes(trackToAdd)){
                setSelectedSongs([...selectedSongs, trackToAdd]);
            }
        }
    }

    return (
        <div className="Search Container">
            {
                authenticated ? 
                    <SearchBar handleTextInput={handleTextInput} 
                        searchTerm={searchTerm}
                        placeholder={"Search for Songs:"} /> 
                    : <LoginButton />
            }
            {
                searchResults && searchTerm ? 
                    searchResults.map((track) => <Song track={track}
                                                    handleClick={handleClick}
                                                    trackPlaying={trackPlaying}
                                                    /> )  
                : <></>
            }             
        </div>
    );
} 

export default SearchContainer;