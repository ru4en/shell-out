# React Builder
FROM node:lts AS react-builder
WORKDIR /app
COPY frontend/ ./
RUN npm install
RUN npm run build

# Nginx Web Server (frontend)
FROM nginx:alpine AS nginx-server
WORKDIR /usr/share/nginx/html
COPY --from=react-builder /app/dist .
COPY frontend/nginx/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
