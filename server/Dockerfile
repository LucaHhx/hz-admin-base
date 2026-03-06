FROM alpine:latest

RUN apk --no-cache add \
    gcompat \
    bash \
    tzdata \
    font-noto \
    font-noto-cjk \
    curl \
    fontconfig

ENV TZ=Asia/Shanghai

WORKDIR /srv/hz-admin-base
COPY . /srv/hz-admin-base

EXPOSE 9688 9689

CMD ["./hz-admin-base", "-type=all"]
