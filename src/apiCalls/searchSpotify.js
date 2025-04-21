export async function searchSpotify(searchTerm) {
    try {
        const response = await fetch(`/search?searchTerm=${encodeURIComponent(searchTerm)}`);
    
        if(!response.ok) {
            throw new Error(`Api returned Status: ${response.status}`);
        }

        const data = await response.json();
        return data.Tracks;
    }catch(error) {
        console.error(error);
        return null;
    }
}

export default searchSpotify