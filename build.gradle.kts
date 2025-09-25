import name.remal.gradle_plugins.lombok.LombokPlugin

plugins {
  `java-library`
  alias(libs.plugins.version.check)
  alias(libs.plugins.git)
  alias(libs.plugins.lombok)
  alias(libs.plugins.dotenv)
  alias(libs.plugins.spotless)
}

group = "org.observa"
version = "1.0-SNAPSHOT"

allprojects {
  repositories {
    mavenLocal()
    mavenCentral()
    google()
  }
}

subprojects {
  apply<JavaLibraryPlugin>()
  apply<LombokPlugin>()
  tasks.test {
    useJUnitPlatform()
  }

  dependencies {
    implementation(enforcedPlatform(rootProject.libs.quarkus.bom))
    annotationProcessor(enforcedPlatform(rootProject.libs.quarkus.bom))
    implementation(rootProject.libs.mapstruct)
    annotationProcessor(rootProject.libs.mapstruct.annotation.processor)
    implementation(rootProject.libs.record.builder.core)
    annotationProcessor(rootProject.libs.record.builder.processor)
    implementation(rootProject.libs.jetbrains.annotation)
  }
}


