
FROM golang:1.21-alpine as go

COPY main.go go.mod go.sum ./
COPY internal internal
COPY filter filter
COPY ent ent
COPY templates templates

RUN CGO_ENABLED go build -o wg

FROM linuxserver/wireguard:1.0.20210914 as wg

COPY ./entrypoint-backend.sh /custom-services.d/
COPY --from=go /go/wg /wg
RUN chmod +x /wg
