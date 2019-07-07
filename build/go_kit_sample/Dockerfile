FROM busybox

COPY cmd/go_kit_sample/go_kit_sample /opt/

WORKDIR "/opt/"

EXPOSE 8000

ENTRYPOINT ["./service"]