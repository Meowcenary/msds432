# Build
# Set image and working dir
FROM node:16-alpine AS build
WORKDIR /app

# Copy package.json and package-lock.json (if available)
# and install dependencies
COPY ./dashboard/package*.json ./
RUN npm install

# Copy the rest of the application files and build
COPY ./dashboard/ .
RUN npm run build

# Run
# Serve the app using nginx
FROM nginx:1.21-alpine

# Copy the built app from the previous stage
COPY --from=build /app/build /usr/share/nginx/html

# Expose ports
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
