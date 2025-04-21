import React, { useState} from "react";

import createPlaylist from "../apiCalls/createPlaylist";
import addTracks from "../apiCalls/addTracks";

import PlaylistName from "../components/PlaylistName";
import Song from "../components/Song";

function PlaylistContainer({selectedSongs, setSelectedSongs}) {
    
    const [playlistName, setPlaylistName] = useState(null);

    const handleTextInput = (e) => setPlaylistName(e.target.value);
    
    const handleClick = (e) => {
        const trackToRemove = selectedSongs.find((track) => {
            return track.id === e.target.id;
        })
        setSelectedSongs(selectedSongs.filter((song) => song.id !== trackToRemove.id));
    };

    const handleUpload = (e) => {
        createPlaylist(playlistName).then((id) => {
            if (id !== null) {
                const urisToAdd = selectedSongs.map((track) => {
                    return track.uri;
                });
                console.log(urisToAdd);
                addTracks(urisToAdd).then((result) => {
                    if (result) {
                        console.log("Playlist Created and Updated Successfully");
                        setSelectedSongs([]);
                        setPlaylistName(null);
                    } 
                })
            }
        })
    }

    return (
        <div className="Playlist Container">
            <PlaylistName handleTextInput={handleTextInput} 
                handleClick={handleUpload}
                playlistName={playlistName} />
            
            {
                selectedSongs ?                  
                    selectedSongs.map((track) => <Song track={track} handleClick={handleClick} search={true} />) 
                    : <></>
            }
        </div>
    );
};

export default PlaylistContainer;
