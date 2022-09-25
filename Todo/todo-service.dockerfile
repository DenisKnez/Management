FROM alpine:latest

RUN mkdir /app
COPY todoApp /app
CMD ["/app/todoApp"]