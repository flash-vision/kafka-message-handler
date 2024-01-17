# Kafka Message Wranglers Module

## Overview
The Kafka Message Wranglers module provides a versatile and easy-to-use interface for processing Kafka messages in Go. It supports converting Kafka messages to various formats like JSON, string, byte array, and maps. With an elegant design and straightforward usage, this module will be a fantastic addition to your Kafka-related projects.

## Features
- **Convert Kafka messages to JSON, string, byte array, and map formats.**
- **Extract specific fields from Kafka messages.**
- **Check if Kafka messages contain specific fields.**
- **Easily extendable interface for additional message processing methods.**

## Installation
To use the Kafka Message Wranglers in your project, you first need to have Go installed. Then, you can import this module into your Go project:

```go
import wranglers "github.com/flash-vision/kafkamessagewranglers"
```

# Example Usage

To use a wrangler, you invoke WrangleMessage against a declared wrangler type:
First declare the wrangler type:

> mywranglertype := wranglers.<WranglerType>{Params}

Then use the wrangler:

> wrangled, err := mywranglertype.WrangleMessage(messageData)

The example below illustrates the use of multiple wranglers:



```
//Check for the existence of particular fields in the message
exists := wranglers.ContainsFieldsWrangler{Fields: []string{"field1", "field2", "address"}}
existing, err := exists.WrangleMessage(messageData)

if err != nil {
    log.Printf("[ERROR][MESSAGE_HANDLER] Error wrangling message: %v\n", err)
}
fieldsExist, ok := existing.(bool)
if !ok {
    log.Printf("[ERROR][MESSAGE_HANDLER] Error casting wrangled message to bool: %v\n", err)
}

if fieldsExist {
		//extract fields from the message
		extractor := wranglers.MapFieldsWrangler{Fields: []string{"address"}}
		extracted, err := extractor.WrangleMessage(messageData)
		if err != nil {
			log.Printf("[ERROR][MESSAGE_HANDLER] Error wrangling message: %v\n", err)
		}
		//cast the extracted fields to a map[string]interface{}
		extractedMap, ok := extracted.(map[string]interface{})
		if !ok {
			log.Printf("[ERROR][MESSAGE_HANDLER] Error casting wrangled message to map[string]interface{}: %v\n", err)
		} else {
			//print the extracted fields to stdout
			fmt.Printf("[INFO][MESSAGE_HANDLER] Extracted fields: %v\n", extractedMap)
		}

		wrangler := wranglers.JsonWrangler{}
		dat, err := wrangler.WrangleMessage(messageData)
		if err != nil {
			log.Printf("[ERROR][MESSAGE_HANDLER] Error wrangling message: %v\n", err)
		}
		//print the message to stdout
		fmt.Println(dat)
	} else {
		log.Println("[INFO][MESSAGE_HANDLER][MISSING FIELDS] Message does not contain the required fields")
	}

}
```

