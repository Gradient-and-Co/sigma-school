import {AUTH_JWT_SECRET} from "$env/static/private";
import jwt from 'jsonwebtoken'

export const handle = async ({event, resolve}) => {
    const authCookie = event.cookies.get('AuthorizationToken');
    console.log('HOOK: ', authCookie);
    if (authCookie) {
        const token = authCookie.split(' ')[1];
        console.log('TOK: ', token);
        try {
            const session = jwt.verify(token, AUTH_JWT_SECRET);
            if (session) {
                event.locals.session = session;
            }
        } catch (error) {
            console.log(error);
        }
    }
    return await resolve(event);
};