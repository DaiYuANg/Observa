package org.observa;

import jakarta.persistence.*;

import java.time.Instant;

@Entity
@Table(name = "event",
  indexes = {@Index(name = "idx_event_project_time",
    columnList = "timestamp"
  )})
public class Event {
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Long id;

  @ManyToOne
  @JoinColumn(name = "project_id", nullable = false)
  private Project project;

  @ManyToOne
  @JoinColumn(name = "environment_id")
  private Environment environment;

  @Column(nullable = false)
  private String level; // ERROR, WARN, INFO

  @Column(nullable = false)
  private String message;

  @Column(name = "timestamp", nullable = false)
  private Instant timestamp;

  @Column(name = "stacktrace", columnDefinition = "TEXT")
  private String stacktrace;

  @Column(name = "sdk_version")
  private String sdkVersion;

  @Column(name = "platform")
  private String platform; // Java, Go, JS 等
}
