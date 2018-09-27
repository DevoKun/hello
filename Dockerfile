FROM   scratch
EXPOSE 8765
ENV    HTTP_PORT=8765
ADD    target/hello-linux-amd64 /
CMD    ["/hello-linux-amd64"]