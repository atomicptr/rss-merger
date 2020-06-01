FROM alpine:3.12 as gobase

RUN apk add --no-cache build-base go git

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o rss-merger -ldflags '-w -extldflags "-static"' cmd/rss-merger/main.go

FROM alpine:3.12 as webbase

RUN apk add --no-cache nodejs yarn openjdk11

COPY . /app
WORKDIR /app/web

RUN yarn && yarn build

FROM alpine:3.12

ENV RSS_MERGER_STORAGEDIR "/data"
ENV RSS_MERGER_PORT 8081

WORKDIR /app
COPY --from=gobase /app/rss-merger /app/rss-merger
COPY --from=webbase /app/public /app/public

ENTRYPOINT ["/app/rss-merger"]
