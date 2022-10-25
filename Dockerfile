FROM --platform=$BUILDPLATFORM golang:1.18.3 AS build
WORKDIR /src
ARG TARGETOS
ARG TARGETARCH
ARG LDFLAGS
ARG OUTPUT_NAME
ARG PACKAGE
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "${LDFLAGS}" -o /src/sedge $package

FROM golang:1.18.3
COPY --from=build /src/sedge /sedge