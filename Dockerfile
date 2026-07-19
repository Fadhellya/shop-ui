# Base images di-mirror ke internal registry agar build tidak tergantung internet
FROM image-registry.openshift-image-registry.svc:5000/virtus-shop-ci/golang:1.26-alpine AS builder

ARG VERSION=1.0.0
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -mod=vendor -ldflags "-s -w" -o /out/shop-ui .

FROM image-registry.openshift-image-registry.svc:5000/virtus-shop-ci/runtime-base:latest

ARG VERSION=1.0.0
ENV VERSION=${VERSION}
WORKDIR /home/app
COPY --from=builder /out/shop-ui .
USER app
EXPOSE 8080
CMD ["./shop-ui"]
