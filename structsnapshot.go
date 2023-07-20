package structsnapshot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

// Snapshot represents a snapshot of a struct
type Snapshot struct {
	Name   string          `json:"name"`
	Fields []FieldSnapshot `json:"fields"`
}

// ToJSON converts the snapshot to a JSON representation
func (s *Snapshot) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Snapshot) SaveToFile() error {
	filePath := fmt.Sprintf("__structsnapshot__/%s.json", s.Name)

	jsonData, err := s.ToJSON()
	if err != nil {
		return fmt.Errorf("error converting snapshot to JSON: %v", err)
	}

	err = os.MkdirAll("__structsnapshot__", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating structsnapshot directory: %v", err)
	}

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error saving JSON snapshot to file: %v", err)
	}

	return nil
}

// FieldSnapshot represents a field of a struct
type FieldSnapshot struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tag  string `json:"tag"`
}

func getFields(fieldType reflect.Type) []FieldSnapshot {
	var fields []FieldSnapshot

	for i := 0; i < fieldType.NumField(); i++ {
		fieldType := fieldType.Field(i)

		if fieldType.Anonymous {
			fields = append(fields, getFields(fieldType.Type)...)
			continue
		}

		field := FieldSnapshot{
			Name: fieldType.Name,
			Type: fieldType.Type.String(),
			Tag:  string(fieldType.Tag),
		}

		fields = append(fields, field)
	}

	return fields
}

// TakeSnapshot creates a snapshot of a struct
func TakeSnapshot(data interface{}) (*Snapshot, error) {
	snapshot := &Snapshot{
		Name: reflect.TypeOf(data).Name(),
	}

	snapshot.Fields = getFields(reflect.ValueOf(data).Type())

	return snapshot, nil
}

// LoadSnapshot loads a JSON snapshot from a file
func LoadSnapshot(data interface{}) (*Snapshot, error) {
	filePath := fmt.Sprintf("__structsnapshot__/%s.json", reflect.TypeOf(data).Name())

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}

	var snapshot Snapshot
	err = json.Unmarshal(jsonData, &snapshot)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON snapshot: %v", err)
	}

	return &snapshot, nil
}
