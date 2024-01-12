package kafkamessagehelpers

import (
	"encoding/json"
	"log"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)
/**
 * Package kafkamessagehelpers provides helper functions for working with Kafka messages.
 */

/**
 * Interface MessageWranglerInterface defines the contract for handling Kafka messages.
 */
type MessageWranglerInterface interface {
    WrangleMessage(*kafka.Message) (interface{}, error)
}

type JsonWrangler struct{}
func (j JsonWrangler) WrangleMessage(m *kafka.Message) (interface{}, error) {
    return MessageToJson(m)
}

type StringWrangler struct{}
func (s StringWrangler) WrangleMessage(m *kafka.Message) (interface{}, error) {
    return MessageToString(m)
}

type BytesWrangler struct{}
func (b BytesWrangler) WrangleMessage(m *kafka.Message) (interface{}, error) {
	return MessageToBytes(m)
}

type MapWrangler struct{}
func (mh MapWrangler) WrangleMessage(m *kafka.Message) (interface{}, error) {
	return MessageToMap(m)
}

type MapFieldsWrangler struct{
	Fields []string
}
func (mfh MapFieldsWrangler) WrangleMessage(m *kafka.Message) (interface{}, error) {
	return MessageToMapFields(m, mfh.Fields)
}

type ContainsFieldsWrangler struct{
	Fields []string
}
func (cfh ContainsFieldsWrangler) WrangleMessage(m *kafka.Message) (interface{}, error) {
	return MessageContainsFields(m, cfh.Fields)
}




func MessageToJson(messageData *kafka.Message) (string, error) {
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

// Convert kafka message to a string
func MessageToString(messageData *kafka.Message) (string, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	return messageValue, nil
}

// Convert kafka message to a byte array
func MessageToBytes(messageData *kafka.Message) ([]byte, error) {
	// Convert the message value to a string
	messageValue := string(messageData.Value)
	// Convert the message value to a byte array
	messageValueBytes := []byte(messageValue)
	return messageValueBytes, nil
}

// convert the kafka message to a map
func MessageToMap(messageData *kafka.Message) (map[string]interface{}, error) {
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

// return a map of specific fields from the kafka message
func MessageToMapFields(messageData *kafka.Message, fields []string) (map[string]interface{}, error) {
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

// verify that a kafka message contains specific fields
func MessageContainsFields(messageData *kafka.Message, fields []string) (bool, error) {
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

func MessageWrangler(messageData *kafka.Message, Wrangler MessageWranglerInterface) any {
    dat, err := Wrangler.WrangleMessage(messageData)
    if err != nil {
        log.Println("Error handling message:", err)
        return nil
    }
    return dat
}



