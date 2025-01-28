FROM node:18-alpine AS builder

WORKDIR /app
COPY . /app

RUN yarn install && yarn build

EXPOSE 5173

CMD ["yarn", "dev"]

# FROM node:18-alpine

# RUN mkdir /app
# COPY --from=builder /app/build /app/build
# COPY --from=builder /app/package*.json /app/yarn.lock /app/

# RUN cd /app && \
#   yarn install --production && \
#   yarn cache clean

# WORKDIR /app
# EXPOSE 3000

# CMD ["node", "build/index.js"]
