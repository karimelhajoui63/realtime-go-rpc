# Use an official Node.js runtime as a parent image
FROM node:alpine

# # Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install --quiet

# Copy the rest of the application code
COPY . .

# Build the Vite project
RUN npm run build

# Expose the port
EXPOSE 45871

# Start the application
CMD ["npm", "run", "preview"]