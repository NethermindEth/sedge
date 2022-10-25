FROM --platform=$BUILDPLATFORM golang:1.18.3 AS build
WORKDIR /src
ARG TARGETOS
ARG TARGETARCH
ARG LDFLAGS
ARG OUTPUT_NAME
ARG PACKAGE
COPY . ./
RUN go mod download
RUN if [ "$TARGETARCH" = "arm64" ] ;  \
    then apt update && apt install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu -y &&  \
      CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "${LDFLAGS}" -o /src/sedge $PACKAGE ;  \
    else  \
      CGO_ENABLED=1 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "${LDFLAGS}" -o /src/sedge $PACKAGE;  \
    fi

FROM golang:1.18.3
COPY --from=build /src/sedge /sedge