provider "kafka" {
  bootstrap_servers = ["localhost:9092"]
}

resource "kafka_topic" "my_test_topic" {
  name               = "my_test_topic"
  replication_factor = 1
  partitions         = 1

  config = {
    "segment.ms"   = "4000"
    "retention.ms" = "86400000"
  }
}