# syntax=docker/dockerfile:1

# Build the application from source
FROM --platform=$BUILDPLATFORM golang:1.23 AS build-stage

ARG TARGETOS
ARG TARGETARCH

RUN echo "I am running on $TARGETOS, building for $TARGETARCH" > /log

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY datafetcher/ datafetcher/
COPY static/ static/

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /up-bank-go


# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /up-bank-go /up-bank-go
COPY ./static ./static/

EXPOSE 63312

USER nonroot:nonroot

ENTRYPOINT ["/up-bank-go"]
