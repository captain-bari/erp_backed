FROM debian:bookworm-slim
RUN mkdir -p /app/ /app/log
COPY --chmod=755 ./erp /app
COPY ./start.sh /app
COPY ./key_open.pem /app
COPY ./cert.pem /app
RUN chmod +x /app/start.sh /app/erp
ENTRYPOINT ["/app/start.sh"]