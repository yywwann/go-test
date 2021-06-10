```shell script
# 进去 kafka docker 容器
$ docker exec -it kafka /bin/bash

# 创建 topic
$ kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic test

# 查询所有 topic
$ kafka-topics.sh --list --zookeeper zookeeper:2181

# 查看 topic 信息
$ kafka-topics.sh --describe --zookeeper zookeeper:2181 --topic test

# 启动一个生产者
$ kafka-console-producer.sh --broker-list localhost:9092 --topic test

# 启动一个消费者
$ kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning

$ kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 2 --partitions 2 --topic my-replicated-topic
$ kafka-topics.sh --describe --zookeeper zookeeper:2181 --topic my-replicated-topic
# Topic: my-replicated-topic	PartitionCount: 2	ReplicationFactor: 2	Configs:
#	  Topic: my-replicated-topic	Partition: 0	Leader: 2	Replicas: 2,1Isr: 2
#	  Topic: my-replicated-topic	Partition: 1	Leader: 1	Replicas: 1,2Isr: 1,2
$ kafka-topics.sh --alter --zookeeper zookeeper:2181 --partitions 16 --topic goim-zb_otc-topic
$ kafka-topics.sh --delete --zookeeper zookeeper:2181 --topic goim-zb_otc-topic

# new
$ kafka-topics.sh --bootstrap-server localhost:9092 --list
$ kafka-topics.sh --bootstrap-server localhost:9092 --topic test16 --create --replication-factor 1 --partitions 16
$ kafka-topics.sh --bootstrap-server localhost:9092 --topic pubv2  --describe 
$ kafka-console-producer.sh --bootstrap-server localhost:9092 --topic quickstart-events 
$ kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic quickstart-events --from-beginning 
$ kafka-topics.sh --bootstrap-server localhost:9092 --topic pubv2 --alter --partitions 16
# 查看所有消费者组
$ kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
# 查看某个消费者组的具体消费情况
$ kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group <group_name> --describe
```