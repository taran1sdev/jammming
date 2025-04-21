export async function addTracks(urisToAdd) {
    try {
        if(urisToAdd === undefined || urisToAdd.length < 1) {
            console.log("No tracks to add.")
            return true;
        }
        const response = await fetch("/addTracks", {
            headers: {
                "Content-Type": "application/json"
            },
            method: "POST",
            body: JSON.stringify(urisToAdd)
        });

        if (!response.ok) {
            throw new Error("Api returned status code: " + response.status);
        }

        const data = await response.json();

        if(data.error) {
            console.error(data.error);
            return false;
        }
        console.log(data.success);
        return true;
    } catch (error) {
        console.error(error);
    }
}

export default addTracks;