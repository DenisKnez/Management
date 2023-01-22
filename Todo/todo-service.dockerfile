FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/config 
COPY todoApp /app
CMD ["/app/todoApp"]