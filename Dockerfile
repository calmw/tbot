# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM  golang AS builder
ADD . /src
WORKDIR /src
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
RUN go mod tidy
RUN go build -o /tbot .

# # final stage
FROM ubuntu
RUN apt-get -y update && apt-get -y upgrade && apt-get install ca-certificates wget -y
RUN wget -P /usr/local/bin/ https://chainbridge.ams3.digitaloceanspaces.com/subkey-rc6 \
  && mv /usr/local/bin/subkey-rc6 /usr/local/bin/subkey \
  && chmod +x /usr/local/bin/subkey
#RUN subkey --version

COPY --from=builder /tbot ./
RUN chmod +x ./tbot

ENTRYPOINT ["./tbot"]
