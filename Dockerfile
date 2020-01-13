FROM golang:1.13-alpine

ARG HOST_UID
ARG GO111MODULE
ARG CGO_ENABLED
ARG GOARCH
ARG GOOS

WORKDIR /go/src/revolut-stocks-list
COPY . .

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk --update --no-cache add -u \
        shadow \
        curl \
        make \
        bash \
        nano \
        g++ \
        leptonica-dev \
        tesseract-ocr-dev \
        git && \
    # apk del --purge \
        # git && \
    rm -rf /var/cache/apk/* \
        /tmp/*

RUN adduser -D -u ${HOST_UID} -s /bin/bash docker && \
    chown -R docker:docker .
USER docker

RUN make -f Makefile.native get

ENTRYPOINT ["./entrypoint.sh"]
CMD ["run"]
