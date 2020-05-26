FROM alpine:3.11 as gobase

RUN apk add --no-cache build-base go git

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o rss-merger -ldflags '-w -extldflags "-static"' cmd/rss-merger/main.go

FROM alpine:3.11 as webbase

RUN apk add --no-cache nodejs npm openjdk11

COPY . /app
WORKDIR /app/web

RUN npm install -g shadow-cljs
RUN shadow-cljs release app
RUN cp -r assets/* /app/dist/

FROM scratch

# ENV vars...

WORKDIR /app
COPY --from=gobase /app/rss-merger /app/rss-merger
COPY --from=webbase /app/dist /app/dist

ENTRYPOINT ["/app/rss-merger"]