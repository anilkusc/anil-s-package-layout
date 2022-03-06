# Anil's Package Layout

Project purpose

# Table of contents

1. [Package Structure](#infrastructure-of-demo-projects)
    * [App](#tech-stack)
    * [Api](#tech-stack)
    * [Domain](#tech-stack)
    * [Pkg](#tech-stack)
    * [Repository](#tech-stack)
2. [Migrating To Microservice](#getting-started)
3. [Tests](#tests)
    * [Mocking](#elastic-container-service)
        - [Database](#elastic-container-service)
        - [API](#elastic-container-service)
    * [Constructors&Destructors](#elastic-container-service)
    * [Comparing Errors](#elastic-container-service)
    * [Coverage](#elastic-container-service)
        - [Coverage Percantage](#elastic-container-service)
        - [Get Missing Tests](#elastic-container-service)
4. [Further Information](#further-information)

## Tests
Purpose is testing every unit independently
### 
go test . -coverprofile test.out && go tool cover -html=test.out -o coverage.html