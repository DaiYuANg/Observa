package org.observa.service;

import io.smallrye.mutiny.Uni;
import jakarta.enterprise.context.ApplicationScoped;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.eclipse.microprofile.reactive.messaging.Channel;
import org.eclipse.microprofile.reactive.messaging.Emitter;
import org.observa.model.EventDTO;

@ApplicationScoped
@Slf4j
@RequiredArgsConstructor
public class EventService {

  @Channel("quote")
  Emitter<String> quoteRequestEmitter;

  public Uni<Void> sendEvent(EventDTO event) {
    log.info("Sending event: {}", event);
    return Uni.createFrom().completionStage(quoteRequestEmitter.send(event.eventId()));
  }
}
