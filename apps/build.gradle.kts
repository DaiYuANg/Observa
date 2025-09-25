import io.quarkus.gradle.QuarkusPlugin

plugins {
  alias(libs.plugins.quarkus) apply false
}

subprojects {
  apply<QuarkusPlugin>()
}