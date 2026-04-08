# --- Build Stage ---
FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm config set strict-ssl false && npm install

COPY . .

# Build for production using Nitro engine
RUN npm run build

# --- Runtime Stage ---
FROM node:20-alpine

WORKDIR /app

# Copy the output from the build stage
COPY --from=builder /app/.output ./

EXPOSE 3000

CMD ["node", "server/index.mjs"]
