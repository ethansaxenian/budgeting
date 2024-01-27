FROM node:lts

WORKDIR "/web"

COPY package.json .

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]
