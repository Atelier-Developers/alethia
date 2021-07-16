# build stage
FROM node:14-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install -s
COPY . .

# development stage
FROM build-stage as development-stage
EXPOSE 8080
CMD ["npm", "run", "--silent", "serve"]
# production stage
FROM build-stage as production-stage
RUN npm run --silent build
FROM nginx:1.20-alpine
COPY nginx_config/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
