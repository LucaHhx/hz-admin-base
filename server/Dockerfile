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

WORKDIR /srv/hab
COPY . /srv/hab

EXPOSE 9688 9689

CMD ["./hab", "-type=all"]
