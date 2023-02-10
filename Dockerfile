FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 6011

ENV TZ=Asia/Jakarta

CMD air