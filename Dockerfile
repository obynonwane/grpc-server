FROM --platform=linux/amd64 alpine:latest

RUN mkdir /app

COPY loggerApp /app

CMD [ "/app/loggerApp" ]



