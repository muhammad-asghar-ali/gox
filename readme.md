# Project Overview

This repository contains several independent projects, each with its own functionality. Below is an overview of each project and its structure.

## Projects

### [ai-fun](./ai-fun)

This project is designed to explore and implement various artificial intelligence models and algorithms.

### [bookstore](./bookstore)

A simple bookstore application with separate sections for:

- **cmd**: Main GO server.
- **internal**:
  - **config**: Configuration settings.
  - **handlers**: HTTP request handlers.
  - **models**: Database models.
  - **routes**: API route definitions.
  - **utils**: Utility functions and helpers.

### [cache](./cache)

A caching service that helps store and retrieve data quickly.

- **internal**: Core logic for the caching mechanism.

### [chronos](./chronos)

Chronos is a bot application that integrates with Slack and provides various functionalities like age calculation, system monitoring, and more. The bot is built using Go and the Slack API, allowing it to interact with Slack users through custom commands.

- **cmd**: Main GO server.
- **config**: Configuration settings.
- **slack**: Slack-related integrations.

### [climatrax](./climatrax)

A climate-related project that likely fetches weather data.

- **cmd**: Main GO server.
- **config**: Configuration settings.
- **handlers**: HTTP request handlers.
- **openweather**: Integration with the OpenWeather API.

### [crud](./crud)

A simple CRUD application for managing movies.

- **cmd**: Main GO server.
- **internal**:
  - **movies**: Logic for managing movie-related operations.

### [ect](./ect)

A simple email checker tool.

### [fileck](./fileck)

A simple Slack bot that uploads files to a Slack channel using the Slack API. The bot uses an OAuth token and file paths to upload files to specified channels in Slack.

- **cmd**: Main GO server.
- **config**: Configuration settings.
- **slack**: Slack-related integrations.

### [go-server](./go-server)

A Go-based server application.

- **static**: Static files served by the server.

### [matrix](./matrix)

A computer systems management application, designed to simulate and manage various components of computer systems, including laptops, processors, memory, screens, keyboards, and storage devices. The system uses **gRPC** for communication and **Protocol Buffers (protobuf)** for defining data structures, enabling efficient and scalable management of hardware resources.

- **cmd**: Main GO server.
- **internal**:
  - **fns**: Functions related to various management operations, such as generating random data for system components or other utility functions.
  - **pb**: Protocol Buffers-related files defining the data structures and services for different system components like keyboards, memory, processors, and laptops.
    - **pbconnect**: Logic for establishing connections and handling communication between different protocol buffer services.
  - **serializer**: Logic for serializing and deserializing data, including support for JSON and file serialization.
  - **services**: Service implementations for handling the business logic of managing system components like laptops, processors, and storage.
  - **tmp**: Temporary or experimental code, used for testing new features or approaches.
- **proto**: Protocol Buffers definitions for services such as laptop, memory, processor, and storage, which define how data is communicated between services.

### [mchkx](./mchkx)

A simple Go program to check DNS records for a domain's mail configuration. It verifies if the domain has MX, SPF, and DMARC records and outputs the results in a CSV format.

### [pulse](./pulse)

Pulse is a real-time notification system built using Go and Kafka. This system includes a producer to send notifications and a consumer to process them.

- **cmd**:
  - **consumer**: Main GO server for consuming messages.
  - **producer**: Main GO server for producing messages.
- **internal**:
  - **errors**: Error handling logic.
  - **handlers**: HTTP request handlers.
  - **kafka**: Kafka-related integrations.
  - **middlewares**: Middlewares for various operations.
  - **models**: Database models.
  - **store**: Data storage logic.
  - **utils**: Utility functions and helpers.

### [stocks](./stocks)

A stock-related project CRUD with postgres database.

- **cmd**:
  - **migrate**: Database migrations.
    - **migrations**: SQL or other migration files.
- **internal**:
  - **config**: Configuration settings.
  - **handlers**: HTTP request handlers.
  - **models**: Database models.
  - **routes**: API route definitions.

### [ums](./ums)

A user management system with mongo database.

- **cmd**: Main GO server.
- **internal**:
  - **config**: Configuration settings.
  - **handlers**: HTTP request handlers.
  - **models**: Database models.
  - **routes**: API route definitions.
