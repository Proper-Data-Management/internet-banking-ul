FROM golang:latest as build_env
LABEL maintainer = "Alexandr (Alex M.A.K.) Mikhailenko <alex-m.a.k@yandex.kz>"

RUN apt update && apt list --upgradable && apt install -y git tar unzip libaio-dev wget pkg-config gcc openssl && apt clean

ENV GO111MODULE=on
ENV GOPROXY="https://proxy.golang.org,direct"
ENV GOSUMDB="off"
ENV GOPATH /root/go
ENV LD_LIBRARY_PATH /opt/oracle/instantclient_21_7
ENV PKG_CONFIG_PATH /opt/oracle/instantclient_21_7
ENV AL_HILAL_CORE_PROJECT ${GOPATH}/src/github.com/internet-banking-ul

ADD go.mod go.sum /root/
WORKDIR /root

RUN go mod download

COPY . ${AL_HILAL_CORE_PROJECT}
WORKDIR ${AL_HILAL_CORE_PROJECT}

RUN make docker-deps
RUN make build

FROM golang:latest
RUN apt update && apt list --upgradable && apt install -y libaio-dev && apt clean
RUN mkdir /root/al_hilal_core

ENV GOPATH /root/go
ENV LD_LIBRARY_PATH /opt/oracle/instantclient_21_7
ENV PKG_CONFIG_PATH /opt/oracle/instantclient_21_7
ENV AL_HILAL_CORE_PROJECT ${GOPATH}/src/github.com/internet-banking-ul

COPY --from=build_env ${AL_HILAL_CORE_PROJECT}/out/bin/al_hilal_core /root/al_hilal_core/al_hilal_core
COPY --from=build_env /opt/oracle /opt/oracle

WORKDIR /root/al_hilal_core
ENTRYPOINT ["/root/al_hilal_core/al_hilal_core"]

EXPOSE 8000
EXPOSE 8001