FROM docker.4pd.io/golang:1.16-alpine as builder 

WORKDIR /build

ARG TARGETARCH=amd64

COPY go.mod go.mod
COPY go.sum go.sum

COPY cmd/ cmd/
COPY pkg/ pkg/

ENV ENABLE_PROXY=true

RUN if [ "$ENABLE_PROXY" = "true" ] ; then go env -w GOPROXY=https://goproxy.cn,direct ; fi \
	&& go mod download


# build
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} \
	go build -a -ldflags "-s -w" -o openaios-iam-${TARGETARCH} cmd/main.go

WORKDIR /dist
RUN cp /build/openaios-iam-${TARGETARCH} ./openaios-iam-${TARGETARCH}


FROM docker.4pd.io/alpine:3.15.0

ARG TARGETARCH=amd64

WORKDIR /

COPY entrypoint.sh /usr/local/bin/ 

# This is required by daemon connnecting with cri
RUN apk add --no-cache ca-certificates bash tzdata \
	&& dos2unix /usr/local/bin/entrypoint.sh \
	&& chmod +x /usr/local/bin/entrypoint.sh \
	&& cp /usr/share/zoneinfo/Hongkong /etc/localtime

COPY --from=builder /dist/openaios-iam-${TARGETARCH} /usr/local/bin/openaios-iam

ENTRYPOINT ["entrypoint.sh"]

CMD [ "openaios-iam" ]