# FIXME: (@isaqueveras) fix version
FROM node:latest

WORKDIR /app

COPY package*.json ./

RUN npm install --force

COPY . .

EXPOSE 3000
