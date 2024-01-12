# kafka-message-helpers
 module of helper functions for converting, validating kafka topic messages



```
messageData := &kafka.Message{ /* ... message data ... */ }

// For JSON handling
jsonHandler := JsonHandler{}
result, err := MessageHandler(messageData, jsonHandler)

// For String handling
stringHandler := StringHandler{}
result, err := MessageHandler(messageData, stringHandler)

// And so on for other handlers

```
