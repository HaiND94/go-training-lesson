# **Why Twirp?**
1. Language-agnostic: Twirp relies on protobuf, whichj supports various languages, make it easier to build cross-language services
2. Simple and Clear: Unlike the complexities of gPRC, Twirp  provides a much simpler interface and clearer error handling
3. HTTP-based: Twirp use HTTP/1.1 and HTTP/2, allowing for easier debugging and compatibility
# **Alternatives**
1. Twirp vs gRPC 
    - Protobufs: Both Twirp and gRPC use protobufs as their Interface Definition Language (IDL), ensuring strong typing and efficitent serialization/deserialization
    - Simplicity: Twrip emphasize in its design. Unlike gRPC, which introduces concepts like stream and channels, Twirp's design is more straightforword with unary RPCs.
    - Transport layer: gRPC use HTTP/2 by default and is optimized for it, offering feature like multiplexing and streaming. Twirp, on the other hand, operates over standard HTTP, allowing for broader compatibility with existing infrastructure and tools. However, it can miss out on some of the efficiencies of HTTP/2.
    - Error Handling: Twirp has built-in error codes similar to HTTP status codes, making errors more comprehensible and directly mappable to HTTP responses. gRPC uses its set of status codes.
    - Language Support: gRPC has a wide range of supported languages due to its larger ecosystem. Twirp's main implementations are in Go and Python, though community support for other languages exists.
2. Twirp vs. JSON-RPC
    - Data Format: Twirp uses protobufs, offering both a JSON and protobuf-based wire format. JSON-RPC, as the name suggests, only uses JSON.

    - IDL: While Twirp has a clear IDL (protobufs), JSON-RPC doesn't prescribe a specific IDL, which can lead to looser contracts between services.

    - Error Handling: Twirp offers more structured error handling with built-in error codes, while JSON-RPC has a more generic error object structure.

    - HTTP Methods: Twirp uses standard HTTP methods (POST for service calls). In contrast, JSON-RPC typically uses POST for all requests, regardless of the operation.
3. Twirp vs. REST
    - Style: REST is an architectural style, whereas Twirp is an RPC framework. While REST focuses on representing resources and their state, Twirp, like other RPC systems, emphasizes calling functions or methods.

    - Data Format: Twirp uses protobufs, allowing for efficient binary serialization. REST, while not limited to, commonly uses JSON.

    - Standardization: Twirp provides a more structured and strict contract between client and server due to protobufs. REST doesn't enforce a specific contract, leading to potential inconsistencies across implementations.

    - Verbosity: Twirp routes are more verbose (often reflecting the full method name), whereas REST emphasizes concise, resource-based URLs.

    - Error Handling: REST relies on standard HTTP status codes for errors. Twirp has its own set of error codes but maps them transparently to HTTP status codes.
