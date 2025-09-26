package org.observa.service;

import io.quarkus.hibernate.reactive.panache.common.WithSession;
import io.smallrye.mutiny.Uni;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import lombok.val;
import org.eclipse.microprofile.reactive.messaging.Incoming;
import org.hibernate.SessionFactory;
import org.hibernate.reactive.mutiny.Mutiny;
import org.observa.TestEntity;

@ApplicationScoped
@Slf4j
@RequiredArgsConstructor
public class EventConsumer {

  private final Mutiny.SessionFactory sessionFactory;

  @Incoming("quote")
  @WithSession
  public Uni<Void> consume(String eventId) {
    log.info("Received event from 'quote' channel: {}", eventId);
    // TODO: 在这里处理消息，比如存数据库、调用其他服务等
    val testEntity = new TestEntity();
    testEntity.setEv(eventId);
    return sessionFactory.withTransaction(session -> session.persist(testEntity));
  }
}
