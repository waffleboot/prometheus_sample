FROM busybox

COPY cmd/service/service /opt/

WORKDIR "/opt/"

EXPOSE 8000

ENTRYPOINT ["./service"]