package savers

import (
	"Lab4/services/models/rerq_types"
	"encoding/json"
	"encoding/xml"
	"github.com/gocarina/gocsv"
	"os"
)

type Marshaller struct {
}

func NewMarshaller() *Marshaller {
	return &Marshaller{}
}

func (m *Marshaller) MarshalToJSON(delivery []rerq_types.DeliveryResponse, name string) (error, string) {
	jsonData, err := json.MarshalIndent(delivery, "", " ")
	if err != nil {
		return err, ""
	}

	file, err := os.Create(name + ".json")
	if err != nil {
		return err, ""
	}
	_, err = file.Write(jsonData)
	if err != nil {
		return err, ""
	}

	return nil, file.Name()
}

func (m *Marshaller) MarshalToXML(delivery []rerq_types.DeliveryResponse, name string) (error, string) {
	xmlData, err := xml.MarshalIndent(delivery, "", " ")
	if err != nil {
		return err, ""
	}

	file, err := os.Create(name + ".xml")
	if err != nil {
		return err, ""
	}
	_, err = file.Write(xmlData)
	if err != nil {
		return err, file.Name()
	}

	return nil, ""
}

func (m *Marshaller) MarshalToCSV(delivery []rerq_types.DeliveryResponse, name string) (error, string) {
	file, err := os.Create(name + ".csv")
	if err != nil {
		return err, ""
	}
	defer file.Close()

	err = gocsv.MarshalFile(delivery, file)
	if err != nil {
		return err, ""
	}

	return nil, file.Name()
}

func (m *Marshaller) UnmarshalFromJSON(pathToFile string) (rerq_types.DeliveryRequest, error) {
	dreq := &rerq_types.DeliveryRequest{}
	file, err := os.Open(pathToFile)
	if err != nil {
		return rerq_types.DeliveryRequest{}, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(dreq)
	if err != nil {
		return rerq_types.DeliveryRequest{}, err
	}
	return *dreq, nil
}

func (m *Marshaller) UnmarshalFromXML(pathToFile string) (rerq_types.DeliveryRequest, error) {
	dreq := &rerq_types.DeliveryRequest{}
	file, err := os.Open(pathToFile)
	if err != nil {
		return rerq_types.DeliveryRequest{}, err
	}
	defer file.Close()

	err = xml.NewDecoder(file).Decode(dreq)
	if err != nil {
		return rerq_types.DeliveryRequest{}, err
	}
	return *dreq, nil

}
