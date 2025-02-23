# Step 1: Build the React application
FROM node:16 as build

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

# Step 2: Serve the React application with npm start (development mode)
FROM node:16

WORKDIR /app

# Copy package.json and package-lock.json separately to optimize cache
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . .

EXPOSE 3000

CMD ["npm", "start"]
