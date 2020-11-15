FROM golang:latest as build

ARG VERSION
ARG COMMIT
ARG BRANCH
# ARG SENTRY_RELEASE_NAME

COPY . /go/src
WORKDIR /go/src

# RUN mv pkg /go/pkg; exit 0 # copy built packages from build step

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app -ldflags="-w -s \
-X main.VERSION=${VERSION} \
-X main.COMMIT=${COMMIT} \
-X main.BRANCH=${BRANCH}"
# -X main.SENTRY_RELEASE=${COMMIT} \
# -X ${SENTRY_BASE}.sentryRelease=${SENTRY_RELEASE_NAME}"

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=build /go/bin/app /
# COPY --from=build /go/src/db/migrations /db/migrations
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/app"]
