package dsmr

// codes retrieved from https://www.netbeheernederland.nl/_upload/Files/Slimme_meter_15_32ffe3cc38.pdf

type Obis struct {
	Code string
	Name string
}

// GetObis returns the obis from the map if the code exists
func GetObis(code string) (Obis, error) {
	if obis, ok := obisMap[code]; ok {
		return obis, nil
	}

	return Obis{}, RunUnknownObis
}

var obisMap = map[string]Obis{
	ObisVersion.Code:             ObisVersion,
	ObisTimestamp.Code:           ObisTimestamp,
	ObisEquipmentIdentifier.Code: ObisEquipmentIdentifier,
	ObisMeterReadingElectricityDeliveredToClientTariff1.Code: ObisMeterReadingElectricityDeliveredToClientTariff1,
	ObisMeterReadingElectricityDeliveredToClientTariff2.Code: ObisMeterReadingElectricityDeliveredToClientTariff2,
	ObisMeterReadingElectricityDeliveredByClientTariff1.Code: ObisMeterReadingElectricityDeliveredByClientTariff1,
	ObisMeterReadingElectricityDeliveredByClientTariff2.Code: ObisMeterReadingElectricityDeliveredByClientTariff2,
	ObisTariffIndicatorElectricity.Code:                      ObisTariffIndicatorElectricity,
	ObisActualElectricityPowerDelivered.Code:                 ObisActualElectricityPowerDelivered,
	ObisActualElectricityPowerReceived.Code:                  ObisActualElectricityPowerReceived,
	ObisNumberOfPowerFailuresInAnyPhase.Code:                 ObisNumberOfPowerFailuresInAnyPhase,
	ObisNumberOfLongPowerFailuresInAnyPhase.Code:             ObisNumberOfLongPowerFailuresInAnyPhase,
	ObisPowerFailureEventLog.Code:                            ObisPowerFailureEventLog,
	ObisNumberOfVoltageSagsInPhaseL1.Code:                    ObisNumberOfVoltageSagsInPhaseL1,
	ObisNumberOfVoltageSagsInPhaseL2.Code:                    ObisNumberOfVoltageSagsInPhaseL2,
	ObisNumberOfVoltageSagsInPhaseL3.Code:                    ObisNumberOfVoltageSagsInPhaseL3,
	ObisNumberOfVoltageSwellsInPhaseL1.Code:                  ObisNumberOfVoltageSwellsInPhaseL1,
	ObisNumberOfVoltageSwellsInPhaseL2.Code:                  ObisNumberOfVoltageSwellsInPhaseL2,
	ObisNumberOfVoltageSwellsInPhaseL3.Code:                  ObisNumberOfVoltageSwellsInPhaseL3,
	ObisTextMessageCodes.Code:                                ObisTextMessageCodes,
	ObisTextMessageMax1024Character.Code:                     ObisTextMessageMax1024Character,
	ObisInstantaneousVoltageL1.Code:                          ObisInstantaneousVoltageL1,
	ObisInstantaneousVoltageL2.Code:                          ObisInstantaneousVoltageL2,
	ObisInstantaneousVoltageL3.Code:                          ObisInstantaneousVoltageL3,
	ObisInstantaneousCurrentL1.Code:                          ObisInstantaneousCurrentL1,
	ObisInstantaneousCurrentL2.Code:                          ObisInstantaneousCurrentL2,
	ObisInstantaneousCurrentL3.Code:                          ObisInstantaneousCurrentL3,
	ObisInstantaneousActivePowerL1PosP.Code:                  ObisInstantaneousActivePowerL1PosP,
	ObisInstantaneousActivePowerL2PosP.Code:                  ObisInstantaneousActivePowerL2PosP,
	ObisInstantaneousActivePowerL3PosP.Code:                  ObisInstantaneousActivePowerL3PosP,
	ObisInstantaneousActivePowerL1MinP.Code:                  ObisInstantaneousActivePowerL1MinP,
	ObisInstantaneousActivePowerL2MinP.Code:                  ObisInstantaneousActivePowerL2MinP,
	ObisInstantaneousActivePowerL3MinP.Code:                  ObisInstantaneousActivePowerL3MinP,
}

var (
	ObisVersion = Obis{
		Code: "1-3:0.2.8",
		Name: "Version",
	}
	ObisTimestamp = Obis{
		Code: "0-0:1.0.0",
		Name: "Timestamp",
	}
	ObisEquipmentIdentifier = Obis{
		Code: "0-0:96.1.1",
		Name: "Equipment identifier",
	}
	ObisMeterReadingElectricityDeliveredToClientTariff1 = Obis{
		Code: "1-0:1.8.1",
		Name: "Electricity delivered to client tariff 1",
	}
	ObisMeterReadingElectricityDeliveredToClientTariff2 = Obis{
		Code: "1-0:1.8.2",
		Name: "Electricity delivered to client tariff 2",
	}
	ObisMeterReadingElectricityDeliveredByClientTariff1 = Obis{
		Code: "1-0:2.8.1",
		Name: "Electricity delivered by client tariff 1",
	}
	ObisMeterReadingElectricityDeliveredByClientTariff2 = Obis{
		Code: "1-0:2.8.2",
		Name: "Electricity delivered by client tariff 2",
	}
	ObisTariffIndicatorElectricity = Obis{
		Code: "0-0:96.14.0",
		Name: "Tariff indicator",
	}
	ObisActualElectricityPowerDelivered = Obis{
		Code: "1-0:1.7.0",
		Name: "Actual power delivered",
	}
	ObisActualElectricityPowerReceived = Obis{
		Code: "1-0:2.7.0",
		Name: "Actual power received",
	}
	ObisNumberOfPowerFailuresInAnyPhase = Obis{
		Code: "0-0:96.7.21",
		Name: "Number of power failures in any phase",
	}
	ObisNumberOfLongPowerFailuresInAnyPhase = Obis{
		Code: "0-0:96.7.9",
		Name: "Number of long power failures in any phase",
	}
	ObisPowerFailureEventLog = Obis{
		Code: "1-0:99.97.0",
		Name: "Power failure event log"}
	ObisNumberOfVoltageSagsInPhaseL1 = Obis{
		Code: "1-0:32.32.0",
		Name: "Number of voltages sags in phase L1",
	}
	ObisNumberOfVoltageSagsInPhaseL2 = Obis{
		Code: "1-0:52.32.0",
		Name: "Number of voltages sags in phase L2",
	}
	ObisNumberOfVoltageSagsInPhaseL3 = Obis{
		Code: "1-0:72.32.0",
		Name: "Number of voltages sags in phase L3",
	}
	ObisNumberOfVoltageSwellsInPhaseL1 = Obis{
		Code: "1-0:32.36.0",
		Name: "Number of voltages swells in phase L1",
	}
	ObisNumberOfVoltageSwellsInPhaseL2 = Obis{
		Code: "1-0:52.36.0",
		Name: "Number of voltages swells in phase L2",
	}
	ObisNumberOfVoltageSwellsInPhaseL3 = Obis{
		Code: "1-0:72.36.0",
		Name: "Number of voltages swells in phase L3",
	}
	ObisTextMessageCodes = Obis{
		Code: "0-0:96.13.1",
		Name: "Text message codes",
	}
	ObisTextMessageMax1024Character = Obis{
		Code: "0-0:96.13.0",
		Name: "Text message codes max 1024 characters",
	}
	ObisInstantaneousVoltageL1 = Obis{
		Code: "1-0:32.7.0",
		Name: "Instantaneous voltage L1",
	}
	ObisInstantaneousVoltageL2 = Obis{
		Code: "1-0:52.7.0",
		Name: "Instantaneous voltage L2",
	}
	ObisInstantaneousVoltageL3 = Obis{
		Code: "1-0:72.7.0",
		Name: "Instantaneous voltage L3",
	}
	ObisInstantaneousCurrentL1 = Obis{
		Code: "1-0:31.7.0",
		Name: "Instantaneous current L1",
	}
	ObisInstantaneousCurrentL2 = Obis{
		Code: "1-0:51.7.0",
		Name: "Instantaneous current L2",
	}
	ObisInstantaneousCurrentL3 = Obis{
		Code: "1-0:71.7.0",
		Name: "Instantaneous current L3",
	}
	ObisInstantaneousActivePowerL1PosP = Obis{
		Code: "1-0:21.7.0",
		Name: "Instantaneous active power L1 Pos P",
	}
	ObisInstantaneousActivePowerL2PosP = Obis{
		Code: "1-0:41.7.0",
		Name: "Instantaneous active power L2 Pos P",
	}
	ObisInstantaneousActivePowerL3PosP = Obis{
		Code: "1-0:61.7.0",
		Name: "Instantaneous active power L3 Pos P",
	}
	ObisInstantaneousActivePowerL1MinP = Obis{
		Code: "1-0:22.7.0",
		Name: "Instantaneous active power L1 Min P",
	}
	ObisInstantaneousActivePowerL2MinP = Obis{
		Code: "1-0:42.7.0",
		Name: "Instantaneous active power L2 Min P",
	}
	ObisInstantaneousActivePowerL3MinP = Obis{
		Code: "1-0:62.7.0",
		Name: "Instantaneous active power L3 Min P",
	}
)
