FROM golang:1.13 as builder

ENV GOPROXY https://goproxy.cn
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o appMain

FROM scratch as prod

COPY --from=builder /app/appMain /app
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/configs /app/configs

CMD [ "/app/appMain" ]
