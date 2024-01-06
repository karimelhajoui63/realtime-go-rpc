# Use an official Node.js runtime as a parent image
FROM node

# # Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install --quiet

# Copy the rest of the application code
COPY . .

# ADd script to ensure that RabbitMQ server is running before trying to connect to it
RUN wget -O ./wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x ./wait-for-it.sh

# Build the Vite project
RUN npm run build

# Expose the port
EXPOSE 45871

# Start the application
CMD ["npm", "run", "preview"]