FROM golang:1.18
RUN mkdir -p /app/ /app/log
COPY ./erp /app
COPY ./start.sh /app
COPY ./key_open.pem /app
COPY ./cert.pem /app
RUN chmod +x /app/start.sh
ENTRYPOINT ["/app/start.sh"]