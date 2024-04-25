FROM golang:1.22
WORKDIR /app
COPY . .
RUN go build -o /uugnet
EXPOSE 23
CMD /uugnet serve