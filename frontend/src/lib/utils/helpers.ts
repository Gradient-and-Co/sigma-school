export const isValidEmail = (email) => {
	const EMAIL_REGEX =
		/[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/;
	return EMAIL_REGEX.test(email.trim());
};

export const isValidPasswordStrong = (password) => {
	const strongRegex = new RegExp('^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})');

	return strongRegex.test(password.trim());
};

export const isValidPasswordMedium = (password) => {
	const mediumRegex = new RegExp(
		'^(((?=.*[a-z])(?=.*[A-Z]))|((?=.*[a-z])(?=.*[0-9]))|((?=.*[A-Z])(?=.*[0-9])))(?=.{6,})'
	);

	return mediumRegex.test(password.trim());
};

export const setAuthToken = ({cookies, token}) => {
	cookies.set('AuthorizationToken', `Bearer ${token}`, {
		maxAge: 60 * 60 * 24,
		path: '/'
	});
	console.log('Set Cookie: ', cookies.get('AuthorizationToken'));
};