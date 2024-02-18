package sse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// formatServerSentEvent takes name of an event and any kind of data and transforms
// it into a server sent event payload structure.
// Data is sent as a json object, { "data": <your_data> }.
//
// Example:
//
//	Input:
//		event="price-update"
//		data=10
//	Output:
//		event: price-update\n
//		data: "{\"data\":10}"\n\n
func FormatServerSentEvent(event string, data any) (string, error) {
	m := map[string]any{
		"data": data,
	}

	buff := bytes.NewBuffer([]byte{})

	encoder := json.NewEncoder(buff)

	err := encoder.Encode(m)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", buff.String()))

	return sb.String(), nil
}
