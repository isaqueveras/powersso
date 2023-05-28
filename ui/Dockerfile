FROM node:16.3.0-alpine

WORKDIR /app/frontend

COPY package*.json ./

RUN npm install --force

COPY . .

RUN npm ci 
RUN npm run build

# ENV NODE_ENV production

EXPOSE 3000

CMD [ "npx", "serve", "build" ]