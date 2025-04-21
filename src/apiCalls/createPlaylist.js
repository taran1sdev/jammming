export async function createPlaylist(name) {
    try {
        const response = await fetch('/createPlaylist', {
            headers: {
                'Content-Type': 'application/json'
            },
            method: "POST",
            body: JSON.stringify({"name": name, "description": "Made with Jammming!", "public": false})
        });

        if (!response.ok) {
            throw new Error("api returned status code: " + response.status);
        }

        const data = await response.json();

        if (data.error) {
            console.error(data.error)
            return null
        }
        
        return data.success
    } catch (error) {
        console.error(error);
        return null;
    }
}

export default createPlaylist;