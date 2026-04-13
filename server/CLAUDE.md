# server

**Parent**: [default](../AGENTS.md) | **Repository**: This project is a Redis server implementation written in Go, aiming to replicate core Redis functionality.

> Implements the primary request handling mechanism for the application, orchestrating the full lifecycle of RESP-formatted command processing from deserialization to execution and response serialization.

## Key Components

*No components detected*

## Folder Overview

## See Also

- [Parent overview →](../AGENTS.md) - Repository-level concepts and architecture
- [client/ →](../client/AGENTS.md) - Implements a client application designed to connect with a Redis-like server. It manages user input, serializes commands using the RESP protocol, transmits data over TCP, and displays server responses.
- [resp/ →](../resp/AGENTS.md) - Provides comprehensive functionality for encoding and decoding data using the Redis Serialization Protocol (RESP). It includes dedicated components for both serialization and deserialization processes, ensuring proper protocol adherence.
