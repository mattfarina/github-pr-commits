FROM debian:9.5-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY bin/github-pr-commits /usr/local/bin/github-pr-commits

CMD /usr/local/bin/github-pr-commits