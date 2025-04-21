import React from 'react';

function SearchBar({searchTerm, handleTextInput, placeholder}) {
    return (
        <>
            <input type='text' name='songSearchBar' value={searchTerm || ""} 
            onChange={handleTextInput} placeholder={placeholder}/>
        </>    
    );
}

export default SearchBar;