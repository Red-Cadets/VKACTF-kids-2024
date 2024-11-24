FROM node:lts-alpine
RUN npm install -g http-server

WORKDIR /client


COPY . .


EXPOSE 31337
CMD [ "http-server", "dist", "-p 31337" ]