# Interface Hierarchy and Usage Flow:

## 1) Config:

- **Purpose:** The top-level interface representing the entire configuration for your application or microservice.
- **Methods:**
    - `GetLoggingConfig()`: Retrieves the configuration specific to logging.
    - `GetServiceConfig()`: Retrieves configuration specific to the service itself (e.g., port, database connection, etc.).

## LoggingConfig:

- **Purpose:** Holds the configuration details required to set up the logger.
- **Methods:**
    - `GetLevel()`: Retrieves the desired logging level (e.g., `constants.LogLevelDebug`, `constants.LogLevelInfo`).
    - `GetLoggerKind()`: Specifies the type or implementation of the logger (e.g., `constants.LoggerKindZap`, `constants.LoggerKindLogrus`).

## LoggerFactory:

- **Purpose:** Responsible for creating and configuring a Logger instance based on the `LoggingConfig`.
- **Methods:**
    - `GetLogger()`: Returns a configured Logger instance.
    - `GetLoggerConfig()`: This might be optional, but it allows you to retrieve the original `LoggingConfig` that was used to set up the logger.

## Logger:

- **Purpose:** The actual interface for logging messages at different levels.
- **Methods:**
    - `Debug`, `Info`, `Warn`, `Error`, `Fatal`: Log messages at their respective levels.
    - `Debugf`, `Infof`, `Warnf`, `Errorf`, `Fatalf`: Log formatted messages.
    - `SetLevel`: Dynamically change the logging level.
    - `GetLevel`, `GetLoggerKind`: Potentially useful for introspection or logging diagnostics.
