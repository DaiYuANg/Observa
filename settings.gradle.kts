pluginManagement {
  repositories {
    mavenLocal()
    mavenCentral()
    gradlePluginPortal()
    google()
  }
}

plugins {
  id("org.gradle.toolchains.foojay-resolver-convention") version "0.10.0"
  id("org.danilopianini.gradle-pre-commit-git-hooks") version "2.0.23"
}
enableFeaturePreview("TYPESAFE_PROJECT_ACCESSORS")
rootProject.name = "Observa"
include("apps:gateway")
include("apps:worker")
include("apps:admin")
include("libs:dbms")