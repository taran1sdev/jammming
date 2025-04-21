export async function getAccessToken() {
	try{
		const response = await fetch("/access");

		if(!response.ok) {
			throw new Error('No response from api');
		}

		const data = await response.json();

		return data.access_token;
	} catch(error) {
		console.error(error);
		return null;
	}

}

export default getAccessToken
