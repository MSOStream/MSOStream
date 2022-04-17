use log::{warn, info};
use rdkafka::{consumer::{DefaultConsumerContext, StreamConsumer, Consumer, CommitMode}, ClientConfig, config::RDKafkaLogLevel, Message};

#[tokio::main]
async fn main() {
    let context = DefaultConsumerContext::default();

    let consumer: StreamConsumer = ClientConfig::new()
        .set("group.id", "video_api")
        .set("bootstrap.servers", "localhost:9092")
        .set("enable.partition.eof", "false")
        .set("session.timeout.ms", "6000")
        .set("enable.auto.commit", "true")
        .set_log_level(RDKafkaLogLevel::Debug)
        .create_with_context(context)
        .expect("Consumer creation failed");

    consumer.subscribe(&vec!["upload"])
        .expect("Can't subscribe to upload");
    
        loop {
            match consumer.recv().await {
                Err(e) => warn!("Kafka error: {:?}", e),
                Ok(m) => {
                    let payload = match m.payload_view::<str>() {
                        None => "",
                        Some(Ok(s)) => s,
                        Some(Err(e)) => { warn!("Error while deserializing payload: {:?}", e); "" },
                    };
                    info!("key: '{:?}', payload: '{}', topic: {}, partition: {}, offset: {}, timestamp: {:?}",
                      m.key(), payload, m.topic(), m.partition(), m.offset(), m.timestamp());
                    consumer.commit_message(&m, CommitMode::Async).unwrap();
                }
            }
        }
}