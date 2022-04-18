FROM bufbuild/buf:1.3.1 as buf
FROM rust:1.60.0-buster as build

ENV PKG_CONFIG_ALLOW_CROSS=1

WORKDIR /usr/src/api-video
COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf
COPY . .

RUN apt update && apt upgrade -y && apt install -y cmake
RUN cargo install --path .


FROM gcr.io/distroless/cc-debian10

COPY --from=build /usr/src/api-video /usr/local/bin/api-video

CMD [ "api-video" ]