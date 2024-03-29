FROM bufbuild/buf:1.4.0 as buf
FROM rust:1.60.0-buster as build

# create a new empty shell project
RUN USER=root cargo new --bin auth
WORKDIR /auth

# copy over your manifests
COPY ./Cargo.lock ./Cargo.lock
COPY ./Cargo.toml ./Cargo.toml

# this build step will cache your dependencies
RUN apt-get update && apt-get upgrade -y && apt-get install -y cmake
RUN mv src/main.rs src/service.rs
RUN cargo build --release
RUN rm src/*.rs

# copy your source tree, build file and buf
COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf
COPY buf.work.yaml buf.work.yaml
COPY buf.gen.yaml buf.gen.yaml
COPY diesel.toml diesel.toml
COPY ./build.rs ./build.rs
COPY ./migrations ./migrations
COPY ./proto ./proto
COPY ./src ./src

# build for release
RUN rm ./target/release/deps/auth_service*
RUN cargo build --release

# our final base
FROM debian:bullseye-slim

RUN apt-get update
RUN apt-get install -y libpq-dev
COPY --from=build /auth/target/release/auth_service .

EXPOSE 8080
# set the startup command to run your binary
CMD ["./auth_service"]

