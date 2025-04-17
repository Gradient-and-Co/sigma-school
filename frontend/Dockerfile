FROM node:22-alpine

WORKDIR /app

COPY . .
RUN npm install

EXPOSE 5173
#ENV HOST=0.0.0.0
#ENV AUTH_JWT_SECRET=jwtSecret
#ENV AUTH_FINGERPRINT=fingerprint

CMD [ "npm", "run", "dev" ]