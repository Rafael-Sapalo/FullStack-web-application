FROM golang:latest
WORKDIR /home/rafael/Biggest_Project/FullStack-web-application/src/server
COPY . .
RUN go build -o server .
ENV PORT=8080
EXPOSE 8080
CMD ["./server"]
