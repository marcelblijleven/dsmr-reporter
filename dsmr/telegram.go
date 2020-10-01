package dsmr

import (
	"strings"
	"time"
)

type Telegram struct {
	Header      string                `json:"header"`
	DateTime    time.Time             `json:"datetime"`
	Version     string                `json:"version"`
	EquipmentID string                `json:"equipment_id"`
	Msg         string                `json:"msg"`
	DataObjects map[string]DataObject `json:"data_objects"`
}

func (t Telegram) String() string {
	powerDelivered := t.DataObjects[ObisActualElectricityPowerDelivered.Code]
	return powerDelivered.Value
}

type DataObject struct {
	Code        string `json:"-"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Unit        string `json:"unit,omitempty"`
}

func ParseTelegram(data []byte, tz string) (Telegram, error) {
	msg := string(data)
	t := Telegram{}
	t.DataObjects = make(map[string]DataObject)

	for _, line := range strings.Split(msg, "\n") {
		line = strings.TrimSpace(line)

		if isHeader(line) {
			t.Header = strings.TrimLeft(line, "/")
			continue
		}

		if isEmptyOrEnd(line) {
			continue
		}

		do, err := parseLine(line)

		if err != nil {
			return t, err
		}

		switch do.Code {
		case ObisVersion.Code:
			t.Version = do.Value
		case ObisEquipmentIdentifier.Code:
			t.EquipmentID = do.Value
		case ObisTimestamp.Code:
			datetime, err := parseTimestamp(do.Value, tz)

			if err != nil {
				return t, err
			}

			t.DateTime = datetime
		default:
			t.DataObjects[do.Code] = do
		}
	}

	t.Msg = msg
	return t, nil
}

func parseLine(line string) (DataObject, error) {
	do := DataObject{}
	code, err := getObisCode(line)

	if err != nil {
		return do, err
	}

	obis, err := GetObis(code)

	if err != nil {
		return do, err
	}

	do.Code = code
	do.Description = obis.Name

	// Remove the Obis identifier from the line and get the value
	value := strings.Split(line, code)[1]

	// Check if value is a multipart value, e.g. for power outages
	if m, isMulti := isMultipartValue(value); isMulti {
		// TODO: figure out how to represent this data in a useful way
		// Sequence of (timestamp)(outage in seconds)
		do.Value = strings.Join(m, " ")
		return do, nil
	}

	do.Value, do.Unit = getValueUnit(value)
	return do, nil
}
