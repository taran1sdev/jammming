export async function transferPlayback(device_id) {
    const postData = {"device_ids": [device_id]};

    try {
        

        const response = await fetch("/transferPlayback", {
            headers: {
                "Content-Type": "application/json"
            },
            method: "POST",
            body: JSON.stringify(postData)
        });

        if (!response.ok) {
            throw new Error("API returned with status code:", response.status);
        }
        
        const data = await response.json();
        if (data.error) {
            console.error(data.error);
            return false;
        }

        return true;
    } catch (error) {
        console.error(error);
        return false;
    }
};

export default transferPlayback;