FROM golang AS build
WORKDIR /app
COPY copilot-gpt4-service/copilot-gpt4-service .
RUN CGO_ENABLED=1 GOOS=linux  go build -o copilot-gpt4-service .

FROM node:18-alpine AS base

FROM base AS deps

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn config set registry 'https://registry.npmmirror.com/'
RUN yarn install

FROM base AS builder

RUN apk update && apk add --no-cache git

ENV OPENAI_API_KEY=""
ENV GOOGLE_API_KEY=""
ENV CODE=""

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN yarn build

FROM base AS runner
WORKDIR /app

RUN apk add proxychains-ng

ENV PROXY_URL=""
ENV OPENAI_API_KEY="http://localhost:8080"
ENV GOOGLE_API_KEY=""
ENV CODE="helloworld"
FROM ubuntu
RUN apt update
RUN apt install -y curl
RUN curl -fsSL https://deb.nodesource.com/setup_16.x |  bash
RUN apt-get install -y nodejs
WORKDIR /app
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/.next/server ./.next/server
COPY --from=build /app/copilot-gpt4-service /app/copilot-gpt4-service
EXPOSE 3000
CMD ["bash","-c","nohup ./copilot-gpt4-service > /dev/null  & node server.js"]
# CMD nohup /app/copilot-gpt4-service && if [ -n "$PROXY_URL" ]; then \
#     export HOSTNAME="127.0.0.1"; \
#     protocol=$(echo $PROXY_URL | cut -d: -f1); \
#     host=$(echo $PROXY_URL | cut -d/ -f3 | cut -d: -f1); \
#     port=$(echo $PROXY_URL | cut -d: -f3); \
#     conf=/etc/proxychains.conf; \
#     echo "strict_chain" > $conf; \
#     echo "proxy_dns" >> $conf; \
#     echo "remote_dns_subnet 224" >> $conf; \
#     echo "tcp_read_time_out 15000" >> $conf; \
#     echo "tcp_connect_time_out 8000" >> $conf; \
#     echo "localnet 127.0.0.0/255.0.0.0" >> $conf; \
#     echo "localnet ::1/128" >> $conf; \
#     echo "[ProxyList]" >> $conf; \
#     echo "$protocol $host $port" >> $conf; \
#     cat /etc/proxychains.conf; \
#     proxychains -f $conf node server.js; \
#     else \
#     node server.js; \
#     fi
