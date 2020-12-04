FROM golang:1.15 as builder

WORKDIR /app
COPY ./ /app

RUN pwd
RUN go get -d -v

# Statically compile our app for use in a distroless container
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o api .

# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static

# Copy over binary
COPY --from=builder /app/api /api
COPY --from=builder /app/db/seed.sql /seed.sql

ENTRYPOINT ["/api"]
