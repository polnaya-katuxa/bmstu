FROM amd64/node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
COPY . .
RUN npm install
RUN npm run build

FROM amd64/nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]