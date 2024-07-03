package types

import (
	"encoding/json"
	"fmt"
)

type DataLocationId uint64

const (
	Fiamma DataLocationId = iota
	NubitDA
	AvailDA
)

func (t *DataLocationId) String() string {
	return [...]string{"Fiamma", "NubitDA", "AvailDA", "AvailDA_BitVM", "SP1"}[*t]
}

func DataLocationIdFromString(dataLocation string) (DataLocationId, error) {
	switch dataLocation {
	case "Fiamma":
		return Fiamma, nil
	case "NubitDA":
		return NubitDA, nil
	case "AvailDA":
		return AvailDA, nil
	}
	return 0, fmt.Errorf("unknown proof data location: %s", dataLocation)
}

func DataLocationIdToString(DataLocationId DataLocationId) (string, error) {
	switch DataLocationId {
	case Fiamma:
		return "Fiamma", nil
	case NubitDA:
		return "NubitDA", nil
	case AvailDA:
		return "AvailDA", nil
	}
	return "", fmt.Errorf("unknown proof data location id : %d", DataLocationId)
}

func (t *DataLocationId) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t, err = DataLocationIdFromString(s)
	return err
}

func (t *DataLocationId) Unmarshal(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t, err = DataLocationIdFromString(s)
	return err
}

func (t DataLocationId) MarshalJSON() ([]byte, error) {
	// Check if the enum value has a corresponding string representation
	if str, ret := DataLocationIdToString(t); ret == nil {
		// If yes, marshal the string representation
		return json.Marshal(str)
	}
	// If not, return an error
	return nil, fmt.Errorf("invalid DataLocationId value: %d", t)
}
