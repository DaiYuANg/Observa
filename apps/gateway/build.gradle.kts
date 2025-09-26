dependencies {
  implementation(libs.bundles.quarkus.application)
  implementation(libs.bundles.quarkus.kafka)
  implementation(libs.bundles.quarkus.rest)
  implementation(libs.quarkus.bucket4j)
}

tasks.quarkusDev{
  extensionJvmOptions {
    jvmArguments.addAll(listOf("-XX:+UseCompactObjectHeaders"))
  }
}