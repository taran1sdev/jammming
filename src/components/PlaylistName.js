import React from "react";
import UploadButton from "./UploadButton";
import SearchBar from "./SearchBar";

function PlaylistName({handleTextInput, handleClick, playlistName }){
    return (
        <div className="PlaylistName">
            <SearchBar name='playlistNameBar' searchTerm={playlistName}
            handleTextInput={handleTextInput} placeholder="Playlist Name:" />
            
            {
                playlistName ? 
                    <UploadButton handleClick={handleClick}/>
                    : <></>
            }
            
            
        </div>
    )
};

export default PlaylistName;