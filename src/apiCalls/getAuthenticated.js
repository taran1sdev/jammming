export async function getAuthenticated() {
    try {
        const response = await fetch("/auth");

        if(!response.ok) {
            throw new Error('No response from api');
        }

        const data = await response.json();

        return data.auth;
    } catch(error) {
        console.error(error);
        return null;
    }
}

export default getAuthenticated