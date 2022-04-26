FROM node:17.7.2

ENV HOME="/app"

RUN mkdir -p ${HOME}
WORKDIR ${HOME}
COPY package*.json ./
RUN yarn install
COPY . ${HOME}

ENTRYPOINT [ "npm", "start" ]

