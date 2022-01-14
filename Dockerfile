FROM debian:stable
ADD release/updatemgr.linux.amd64 /app/updatemgr
ENTRYPOINT ["/app/updatemgr"]
