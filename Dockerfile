FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
      netbase \
      && rm -rf /var/lib/apt/lists/ \
      && apt-get autoremove -y && apt-get autoclean -y

COPY . /src/

WORKDIR /src

EXPOSE 8765
#8888 docker build -t rchd-uploadservice:v1 .
CMD ["/src/hzminioupload" ,"-ctype" ,"nacos" ,"-chost" ,"192.168.130.32:8848"]
