FROM ubuntu
LABEL authors="aaditya"

RUN mkdir /app
WORKDIR /app
COPY ../backend/build/build ./build
RUN chmod +x ./build

ENTRYPOINT ["./build"]