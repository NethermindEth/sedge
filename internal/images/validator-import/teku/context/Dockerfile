FROM busybox:latest

RUN mkdir -p /scripts

COPY validator-init.sh /scripts

WORKDIR /scripts

RUN dos2unix validator-init.sh

RUN chmod +x validator-init.sh

CMD ["/bin/sh", "validator-init.sh"]
