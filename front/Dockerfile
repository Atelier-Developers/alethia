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

# pre production stage
FROM build-stage as pre-production-stage
RUN npm run --silent build

FROM nginx:1.20-alpine as production-stage
COPY nginx_config/default.conf /etc/nginx/conf.d/default.conf
COPY --from=pre-production-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
