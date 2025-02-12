# Matrix - Computer Systems Management

## Overview

The **Matrix - Computer Systems Management** is a project that demonstrates the implementation of gRPC for managing CPU-related resources. The system simulates the management and monitoring of CPU hardware components such as keyboard, laptop, memory, processor, screen, and storage. It leverages **Protocol Buffers** for defining data structures and **gRPC** for communication between services.

This project models real-world interactions with a CPU system using multiple components, which are defined as Protocol Buffers in the `proto/` directory.

### Key Features:

- **gRPC-based communication**: All inter-service communication is done using gRPC, enabling high-performance and scalable communication.
- **Protocol Buffers (protobuf)**: All data exchanged between services is serialized and deserialized using Protocol Buffers, ensuring efficient and compact data transmission.
- **Real-life use case**: This system models a real-world scenario of computer management with various services interacting with each other, making it suitable for testing and exploring the potential of gRPC in a production-like environment.
