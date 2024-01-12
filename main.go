package kafkamessagehelpers

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func messageToJson(messageData *kafka.Message) (string, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	// Convert the message value to a byte array
	messageValueBytes := []byte(messageValue)
	// Convert the byte array to a JSON object
	var dat map[string]interface{}
	if err := json.Unmarshal(messageValueBytes, &dat); err != nil {
		return "", err
	}
	// Convert the JSON object to a string
	jsonString, err := json.Marshal(dat)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

//Convert kafka message to a string
func messageToString(messageData *kafka.Message) (string, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	return messageValue, nil
}

//Convert kafka message to a byte array
func messageToBytes(messageData *kafka.Message) ([]byte, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	// Convert the message value to a byte array
	messageValueBytes := []byte(messageValue)
	return messageValueBytes, nil
}

//convert the kafka message to a map
func messageToMap(messageData *kafka.Message) (map[string]interface{}, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	// Convert the message value to a byte array
	messageValueBytes := []byte(messageValue)
	// Convert the byte array to a JSON object
	var dat map[string]interface{}
	if err := json.Unmarshal(messageValueBytes, &dat); err != nil {
		return nil, err
	}
	return dat, nil
}

//return a map of specific fields from the kafka message
func messageToMapFields(messageData *kafka.Message, fields []string) (map[string]interface{}, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	// Convert the message value to a byte array
	messageValueBytes := []byte(messageValue)
	// Convert the byte array to a JSON object
	var dat map[string]interface{}
	if err := json.Unmarshal(messageValueBytes, &dat); err != nil {
		return nil, err
	}
	// Create a new map to hold the specific fields
	mapFields := make(map[string]interface{})
	// Iterate over the fields and add them to the map
	for _, field := range fields {
		mapFields[field] = dat[field]
	}
	return mapFields, nil
}

//verify that a kafka message contains specific fields
func messageContainsFields(messageData *kafka.Message, fields []string) (bool, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	// Convert the message value to a byte array
	messageValueBytes := []byte(messageValue)
	// Convert the byte array to a JSON object
	var dat map[string]interface{}
	if err := json.Unmarshal(messageValueBytes, &dat); err != nil {
		return false, err
	}
	// Iterate over the fields and check if they exist in the map
	for _, field := range fields {
		if _, ok := dat[field]; !ok {
			return false, nil
		}
	}
	return true, nil
}
