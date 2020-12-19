## ILogger
ILogger is an opinionated, minimal, lazy-loading logger. Uses a mix of global variables and dependency injection for simplicity since Logger should be present just about everywhere.

## Use 
Use the logger by calling `Logger().{LogLevel}()` anywhere.

##### For Example: 
- `Logger().Error("yikes")`
- `Logger().Info("a thing happened.")`