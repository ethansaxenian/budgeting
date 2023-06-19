FROM node:lts

WORKDIR "/api"

COPY package.json .

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]
