dependencies {
  implementation(libs.bundles.quarkus.application)
  implementation(libs.bundles.quarkus.rest)
  implementation(libs.bundles.hibernate.reactive)
  implementation(libs.bundles.hibernate.orm)
  implementation(projects.libs.dbms)
}
