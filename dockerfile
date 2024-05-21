FROM node:22-alpine3.18

WORKDIR /app
COPY package.json ./
COPY . .
RUN npm install
RUN npm run build

EXPOSE 8080
CMD ["npm", "run", "start"]