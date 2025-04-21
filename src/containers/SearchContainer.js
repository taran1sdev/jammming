import React, {useEffect, useState} from 'react';

import LoginButton from '../components/LoginButton';
import SearchBar from '../components/SearchBar';
import {Song} from '../components/Song'; 

import searchSpotify from '../apiCalls/searchSpotify';


function SearchContainer({selectedSongs, setSelectedSongs, authenticated}) {
    const [searchTerm, setSearchTerm] = useState('');
    const [searchResults, setSearchResults] = useState([]);

    useEffect(() => {
        if(searchTerm){
            searchSpotify(searchTerm).then((results) => {
                const tracks = results.map((track) => {
                    return {
                        id: track.ID,
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

    //const [trackPlaying, setTrackPlaying] = useState(null);
    //const [audio, setAudio] = useState(null);

    const handleClick = (e) => {
        const trackToAdd = searchResults.find((track) => {
            return track.id === e.target.id; 
        })

        if(!selectedSongs.includes(trackToAdd)){
            setSelectedSongs([...selectedSongs, trackToAdd]);
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
                                                    /> )  
                : <></>
            }             
        </div>
    );
} 

export default SearchContainer;