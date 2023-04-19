FROM ubuntu:jammy
RUN apt update && apt install -y ca-certificates curl && rm -rf /var/lib/apt/lists/*
COPY numary /usr/bin/numary
EXPOSE 3068
ENTRYPOINT ["numary"]
ENV OTEL_SERVICE_NAME ledger
CMD ["server", "start"]
