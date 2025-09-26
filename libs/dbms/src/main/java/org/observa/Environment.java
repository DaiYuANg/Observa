package org.observa;

import jakarta.persistence.*;

@Entity
@Table(name = "environment")
public class Environment {
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Long id;

  @ManyToOne
  @JoinColumn(name = "project_id", nullable = false)
  private Project project;

  @Column(nullable = false)
  private String name;
}
