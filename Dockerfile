FROM debian:stable
RUN echo "deb http://deb.debian.org/debian sid main" > /etc/apt/sources.list
ADD release/updatemgr.linux.amd64 /app/updatemgr
ENTRYPOINT ["/app/updatemgr"]
