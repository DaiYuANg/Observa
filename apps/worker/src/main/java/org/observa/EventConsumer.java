package org.observa.service;

import io.smallrye.mutiny.Uni;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import lombok.extern.slf4j.Slf4j;
import org.eclipse.microprofile.reactive.messaging.Incoming;

@ApplicationScoped
@Slf4j
public class EventConsumer {

  @Incoming("quote")
  public void consume(String eventId) {
    log.info("Received event from 'quote' channel: {}", eventId);
    // TODO: 在这里处理消息，比如存数据库、调用其他服务等
  }
}
