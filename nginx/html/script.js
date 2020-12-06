const jwtKey = "session";
const baseUrl = "http://localhost:8080";

const loginSuccessId = "loginSuccess";
const loginErrorId = "loginError";

const registerSuccessId = "registerSuccess";
const registerErrorId = "registerError";

/**
 * Handles setting the text content of a form response message element.
 * @param {string} id - Id of element to set message text to.
 * @param {any} data - Data containing message.
 * @returns {string}
 */
const handleMessageEl = (id, data) => {
	const errEl = document.getElementById(id);
	errEl.textContent = data.message ? data.message : data;
	errEl.style.display = "block";

	return id;
};

/**
 * Handles resetting the text content of a form message element.
 * @param {string} id - Id of element to set message text to.
 * @returns {*}
 */
const resetMessageEl = (id) => {
	const errEl = document.getElementById(id);
	errEl.textContent = "";
	errEl.style.display = "none";

	return id;
};

/**
 * Handles user login with server.
 * @param {Event} e - Form submit event.
 */
const login = (e) => {
	e.preventDefault();

	const email = document.getElementById("email").value;
	const password = document.getElementById("password").value;
	const payload = { email, password };

	fetch(`${baseUrl}/login`, {
		method: "POST",
		body: JSON.stringify(payload),
		headers: {
			"Content-Type": "application/json; charset=UTF-8"
		}
	}).then(async res => {
		const data = await res.json();
		if (res.status !== 200) {
			handleMessageEl(loginErrorId, data);
			return resetMessageEl(loginSuccessId);
		}

		const { access_token } = data;
		if (!access_token) {
			console.error("No access_token in server response::");
			const err = `Sorry, something went wrong. 
			If the problem persists please contact admin: danny@52inc.com`;
			handleMessageEl(loginErrorId, err);
			return resetMessageEl(loginSuccessId);
		}

		sessionStorage.setItem(jwtKey, access_token);
		handleMessageEl(loginSuccessId, "Login successful!");
		resetMessageEl(loginErrorId);

		window.location.pathname = "/members.html";
	}).catch(err => {
		handleMessageEl(loginErrorId, err);
		return resetMessageEl(loginSuccessId);
	});
};

/**
 * Handles new user registration.
 * @param {Event} e - Form submit event.
 */
const register = (e) => {
	e.preventDefault();

	const email = document.getElementById("email").value;
	const password = document.getElementById("password").value;
	const payload = { email, password };

	fetch(`${baseUrl}/register`, {
		method: "POST",
		body: JSON.stringify(payload),
		headers: {
			"Content-Type": "application/json; charset=UTF-8"
		}
	}).then(async res => {
		const data = await res.json();
		if (res.status !== 200) {
			handleMessageEl(registerErrorId, data);
			return resetMessageEl(registerSuccessId);
		}

		handleMessageEl(registerSuccessId, data);
		return resetMessageEl(registerErrorId);
	}).catch(err => {
		handleMessageEl(registerErrorId, err);
		return resetMessageEl(registerSuccessId);
	});
};

/**
 * Get customers and their credit scores from db.
 * @returns {Promise<Response>}
 */
const getCustomers = () => {
	return fetch(`${baseUrl}/customers`, {
		method: "GET",
		headers: {
			"Authorization": "Bearer " + sessionStorage.getItem(jwtKey),
			"Content-Type": "application/json; charset=UTF-8"
		}
	}).then(res => {
		if (res.status === 401) {
			throw Error("Session expired, redirecting to login...");
		}

		return res.json();
	});
};

/**
 * Handles logging out with server.
 * @returns {Promise<Response>}
 */
const logout = () => {
	return fetch(`${baseUrl}/logout`, {
		method: "POST",
		headers: {
			"Authorization": "Bearer " + sessionStorage.getItem(jwtKey),
			"Content-Type": "application/json; charset=UTF-8"
		}
	});
};