package org.observa;

import jakarta.persistence.*;
import org.observa.Environment;

import java.util.List;

@Entity
@Table(name = "project")
public class Project {
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Long id;

  @Column(nullable = false, unique = true)
  private String name;

  private String description;

  @OneToMany(mappedBy = "project")
  private List<Environment> environments;
}
