# spiri API

Sample project. Contains one endpoint for Spiri app.

# Project structure
- domain - domain entities and business logic
- external - external sources e.g. database, grpc interface, web, http clients, etc.
- services - use case business logic that connects external sources to domain
- spiri - GRPC contract and generated code

# Architecture
Inspired by http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/

Main points:
- Packages (layers) that import one another from top to bottom:
  * External
  * Services
  * Domain
- Interfaces are defined on the receiving side (lower layer) to avoid circular dependencies
- Higher layers implement interfaces defined in lower levels
- Main application creates all dependencies and passes them downwards
