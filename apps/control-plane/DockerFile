# syntax=docker/dockerfile:1
FROM golang:1.25-alpine AS build
RUN apk add --no-cache ca-certificates git
WORKDIR /src

# Cache deps
COPY apps/control-plane/go.mod apps/control-plane/go.sum ./apps/control-plane/
RUN cd apps/control-plane && go mod download

# Copy repo
COPY . .

# Auto-detect main package (best-effort)
ARG MAIN_PKG=""
RUN set -eux; \
  cd apps/control-plane; \
  if [ -z "$MAIN_PKG" ]; then \
    MAIN_PKG="$(go list -f '{{if eq .Name \"main\"}}{{.ImportPath}}{{end}}' ./... | head -n 1)"; \
  fi; \
  echo "Building MAIN_PKG=$MAIN_PKG"; \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/control-plane "$MAIN_PKG"

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/control-plane /control-plane
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/control-plane"]
