# build stage
FROM node:14-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install -s
COPY . .
EXPOSE 8080
CMD ["npm", "run", "--silent", "serve"]
