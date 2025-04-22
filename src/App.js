import './App.css';

import React, {useState, useEffect} from 'react';

import getAuthenticated from './apiCalls/getAuthenticated';

import SearchContainer from './containers/SearchContainer'
import PlaylistContainer from './containers/PlaylistContainer'

import WebPlayback from './components/WebPlayback';

function App() {
    const [selectedSongs, setSelectedSongs] = useState([]);
    const [authenticated, setAuthenticated] = useState(false);

	useEffect(() => {
		getAuthenticated().then((result) => {
			setAuthenticated(result);
		})
	}) 
	return (
    <div className="App">
        <h1>Jamming: Spotify Playlist Builder</h1>
    	<SearchContainer
		selectedSongs={selectedSongs}
		setSelectedSongs={setSelectedSongs} 
		authenticated={authenticated} />

	<PlaylistContainer
		selectedSongs={selectedSongs}
		setSelectedSongs={setSelectedSongs} />
    
	{
		authenticated ? <WebPlayback/> 
		: <></>
	}
	</div>
  );
}

export default App;
