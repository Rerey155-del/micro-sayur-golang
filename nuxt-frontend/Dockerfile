# --- Build Stage ---
FROM node:18-alpine AS builder

# Set npm registry to use HTTP or skip SSL if necessary due to local network certificate issues
RUN npm config set strict-ssl false

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Build for production using Nitro engine
RUN npm run build

# --- Runtime Stage ---
FROM node:18-alpine

WORKDIR /app

# Copy the output from the build stage
COPY --from=builder /app/.output ./

EXPOSE 3000

CMD ["node", "server/index.mjs"]
