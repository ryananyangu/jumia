############################
# STEP 1 build executable binary
############################
# golang alpine 1.14
FROM golang:alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/jumia
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/jumia .

############################
# STEP 2 build a small image
############################
FROM scratch


# Copy our static executable
COPY --from=builder /go/bin/jumia /go/bin/jumia

# Use an unprivileged user.
USER appuser:appuser


# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group


ARG ENV=ENV
ARG DB_HOST=DB_HOST
ARG DB_PORT=DB_PORT
ARG DB_USER=DB_USER
ARG DB_PASSWORD=DB_PASSWORD
ARG DB_NAME=DB_NAME

ENV ENV=${ENV}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}

# ENV PROJECT_ID=${PROJECT_ID}

# Run the binary.
EXPOSE 8080

ENTRYPOINT ["/go/bin/jumia"]