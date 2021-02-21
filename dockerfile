FROM frolvlad/alpine-glibc
WORKDIR /app
COPY core  /app/
EXPOSE 8000
ENTRYPOINT ["/app/core"]