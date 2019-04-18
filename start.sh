docker rm zookeeper kafka1 kafka2 kafka3

docker run --name zookeeper \
  -p 2181:2181 \
  -itd \
  -e ALLOW_ANONYMOUS_LOGIN=yes \
  bitnami/zookeeper:latest

docker run --name kafka1 \
  -itd \
  --link zookeeper:zookeeper \
  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e ALLOW_PLAINTEXT_LISTENER=yes \
  -p 0.0.0.0:9092:9092 \
  bitnami/kafka:latest

#docker run --name kafka2 \
#  -itd \
#  --network app-tier \
#  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
#  -e ALLOW_PLAINTEXT_LISTENER=yes \
#  -p 19092:9092 \
#  bitnami/kafka:latest
#
#docker run --name kafka3 \
#  -itd \
#  --network app-tier \
#  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
#  -e ALLOW_PLAINTEXT_LISTENER=yes \
#  -p 19093:9092 \
#  bitnami/kafka:latest
