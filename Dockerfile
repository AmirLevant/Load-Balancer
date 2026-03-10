FROM golang:1.25.6

WORKDIR /app

COPY . .

RUN make build . 


CMD [ "/app/bin/lb"]
