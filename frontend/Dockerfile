# # Base Image: Node.js 8.X
# FROM node:carbon

# # Create app directory
# WORKDIR /usr/src/app

# # Install app dependencies
# # A wildcard is used to ensure both package.json AND package-lock.json are copied
# # where available (npm@5+)
# COPY package*.json ./

# RUN npm install

# COPY . .

# EXPOSE 3000

# CMD ["npm", "start"]

# base image
FROM node:12.2.0-alpine

# set working directory
WORKDIR /usr/src/app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /app/package.json
RUN npm install --silent
RUN npm install react-scripts@3.0.1 -g --silent
# start app
CMD ["npm", "start"]