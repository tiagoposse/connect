
FROM cosmtrek/air as air

FROM golang:1.21-alpine as go

FROM linuxserver/wireguard as wg

COPY --from=air /go/bin/air /usr/bin/air
COPY --from=go /usr/local/go/ /usr/local/go/

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH
