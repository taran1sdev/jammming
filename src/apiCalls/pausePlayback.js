export async function pausePlayback() {
    try {
        const response = await fetch("/pause");

        if (!response.ok) {
            throw new Error("API returned status code: " + response.status)
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

export default pausePlayback;