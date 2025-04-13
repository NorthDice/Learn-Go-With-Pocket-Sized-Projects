/*
Package pocketlog provides a logging API for your application.

First, create a logger using pocketlog.New with a specified log level threshold.
Messages with a lower priority level will not be logged.

The caller is responsible for sharing the logger instance across different parts of the application.

The logger supports three log levels:
- Debug: primarily used for debugging and tracing step-by-step execution
- Info: useful messages highlighting key milestones in a process
- Error: error messages to help diagnose issues and failures
*/
package pocketLog
