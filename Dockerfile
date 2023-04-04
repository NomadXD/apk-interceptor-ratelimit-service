# Set the base image to Node.js v14
FROM node:18

# Create a new directory to store the app
WORKDIR /app

# Copy package.json and package-lock.json to the app directory
COPY package*.json ./

# Install app dependencies
RUN npm install

# Copy the rest of the app files to the app directory
COPY . .

# Expose port 3000
EXPOSE 8080

# Start the app
CMD ["npm", "start"]