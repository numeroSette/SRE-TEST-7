FROM golang:1.13

WORKDIR /src
COPY . .

ENV GO111MODULE=on

RUN cd cmd && go build -o /bin/cmd

ENTRYPOINT ["/bin/cmd"]