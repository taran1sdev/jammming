export async function playTrack(song_uri) {
    const postData = {"uris": [song_uri]};

    try {
        const response = await fetch("/playTrack", {
            headers: {
                "Content-Type": "application/json"
            },
            method: "POST",
            body: JSON.stringify(postData) 
        });

        if (!response.ok) {
            throw new Error("API returned with status code: " + response.status);
        }

        const data = await response.json();

        if (data.error) {
            console.error(data.error);
            return false;
        }

        return true;
    } catch(error) {
        console.error(error);
        return false;
    }
};

export default playTrack;