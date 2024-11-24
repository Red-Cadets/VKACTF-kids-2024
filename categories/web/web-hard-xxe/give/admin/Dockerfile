FROM node:lts-alpine
WORKDIR /admin
COPY . .
RUN npm install

EXPOSE 3017
CMD [ "node", "index" ]