package models

import (
	"Lab4/controllers"
	"Lab4/services/models/rerq_types"
	"Lab4/services/models/savers"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Server struct {
	dcontrol   *controllers.DeliverController
	marchaller *savers.Marshaller
	encrypter  *savers.Encrypter
	compressor *savers.Compressor
}

func NewServer() *Server {
	return &Server{
		dcontrol:   controllers.NewDeliveryController(),
		marchaller: savers.NewMarshaller(),
		encrypter:  savers.NewEncrypter(),
		compressor: savers.NewCompressor(),
	}
}

func (s *Server) Print(msg string) {
	fmt.Println(msg)
}
func (s *Server) PrintWithWait(msg string) string {
	var input string
	fmt.Println(msg)
	_, _ = fmt.Scanln(&input)
	return input
}

func (s *Server) ProceedFile(pathToFile string) (rerq_types.DeliveryRequest, error) {
	fileExt := filepath.Ext(pathToFile)

	switch fileExt {
	case ".json":
		return s.marchaller.UnmarshalFromJSON(pathToFile)
	case ".xml":
		return s.marchaller.UnmarshalFromXML(pathToFile)
	}

	return rerq_types.DeliveryRequest{}, errors.New("invalid file extension")
}

func (s *Server) ProceedCalculation(data rerq_types.DeliveryRequest) ([]rerq_types.DeliveryResponse, error) {
	deliverResp := make([]rerq_types.DeliveryResponse, len(data.Batches))

	for i, batch := range data.Batches {
		s.dcontrol.AddCargo(controllers.CargoInfo{
			Amount: batch.CargoNumber,
			Type:   batch.CargoType,
		})

		s.dcontrol.SetDeliveryInfo(controllers.DeliveryInfo{
			batch.TransportInstance,
			batch.TransportType, batch.DeliveryDistance,
		})

		total, time, err := s.dcontrol.GetDeliveyResult()
		if err != nil {
			return nil, err
		}
		deliverResp[i].TotoalCost = total
		deliverResp[i].Time = time
	}

	return deliverResp, nil
}

func (s *Server) ProceedOne() {
	cargoAmount := s.PrintWithWait("Enter amount of batches ")
	i, err := strconv.Atoi(cargoAmount)
	if err != nil {
		fmt.Println(err)
	}

	for range i {
		s.Print(fmt.Sprintf("Cargo %d", i))
		cargoNumber := s.PrintWithWait("Enter number of cargos")
		cargoType := s.PrintWithWait("Enter type of cargo")

		amount, _ := strconv.Atoi(cargoNumber)
		s.dcontrol.AddCargo(controllers.CargoInfo{
			amount,
			cargoType,
		})

	}

	distance := s.PrintWithWait("Enter distance to the point of the destination")
	transportType := s.PrintWithWait("Enter type of transport")
	inner := s.GetTransportInfo(transportType)

	distanceFloat, _ := strconv.ParseInt(distance, 10, 32)

	s.dcontrol.SetDeliveryInfo(controllers.DeliveryInfo{
		transportType,
		inner, int(distanceFloat),
	})

	s.Print("The cost of delivery: ")

	total, time, err := s.dcontrol.GetDeliveyResult()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Total delivery cost: ", total)
	fmt.Println("Time delivery cost: ", time)
}

func (s *Server) Show() {
	fmt.Println("Choose type of input:")
	fmt.Println("1. By user")
	fmt.Println("2. By file")

	var input string
	_, _ = fmt.Scanln(&input)
	switch input {
	case "1":
		s.ProceedOne()
	case "2":
		fmt.Println("Enter filename")
		var fileName string
		_, _ = fmt.Scanln(&fileName)
		data, err := s.ProceedFile(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		resultDeliver, err := s.ProceedCalculation(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		s.FileOperations(resultDeliver)
	}

}

func (s *Server) GetTransportInfo(trtype string) string {
	fmt.Println("Print one of the variant: ")
	switch trtype {
	case "whater":
		fmt.Println("1. tanker")
	case "air":
		fmt.Println("1. plane")
		fmt.Println("2. helicopter")
	case "ground":
		fmt.Println("1. train")
		fmt.Println("2. truck")
	default:
		fmt.Println("Non type in")
		return ""
	}

	var input string
	_, _ = fmt.Scanln(&input)
	return input
}

func (s *Server) SaveOutput(data []rerq_types.DeliveryResponse) error {
	fmt.Println("Choose type of outpu to save:")
	fmt.Println("1. xml")
	fmt.Println("2. json")
	fmt.Println("3. csv")

	var input string
	_, _ = fmt.Scanln(&input)

	var fileName string
	var err error
	switch input {
	case "1":
		err, fileName = s.marchaller.MarshalToXML(data, "delivery")
		if err != nil {
			return err
		}
	case "2":
		err, fileName = s.marchaller.MarshalToJSON(data, "delivery")
		if err != nil {
			return err
		}
	case "3":
		err, fileName = s.marchaller.MarshalToCSV(data, "delivery")
		if err != nil {
			return err
		}
	}
	s.ZipOperation(fileName)
	fmt.Println("Output saved")

	return nil
}

func (s *Server) PrintOutput(data []rerq_types.DeliveryResponse) {
	for i, d := range data {
		fmt.Println(fmt.Sprintf("Batch number %d", i+1))
		fmt.Println(d.TotoalCost)
		fmt.Println(d.Time)
		fmt.Println("--------------------------------")
	}
}

func (s *Server) Encrypt(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = s.encrypter.GenerateKey()
	if err != nil {
		return err
	}
	fmt.Println("Your secret key, please save to decrypt : " + s.encrypter.GetKey())

	encryptes, err := s.encrypter.Encrypt(file)

	encFile, err := os.Create(fileName + ".txt")
	if err != nil {
		return err
	}
	_, err = encFile.WriteString(encryptes)
	if err != nil {
		return err
	}
	encFile.Close()
	fmt.Println("Successfully encrypted")

	err = s.ZipOperation(fileName + ".txt")
	if err != nil {
		return err
	}

	return nil

}

func (s *Server) ToZip(fileName string) error {
	err := s.compressor.Compress(fileName)
	if err != nil {
		return err
	}
	fmt.Println("Saved to ZIP")
	return nil
}

func (s *Server) FileOperations(data []rerq_types.DeliveryResponse) {
	fmt.Println("Select if you want to procced")
	fmt.Println("1. SaveFile")
	fmt.Println("2. Exit")

	var input string
	_, _ = fmt.Scanln(&input)
	switch input {
	case "1":
		err := s.SaveOutput(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":
		os.Exit(0)
	}

}

func (s *Server) ZipOperation(filename string) error {
	fmt.Println(filename)
	fmt.Println("Select if you want to procced")
	fmt.Println("1. ZipFile")
	fmt.Println("2. Encrypt")
	fmt.Println("3. Exit")

	var input string
	_, _ = fmt.Scanln(&input)
	switch input {
	case "1":
		err := s.ToZip("/home/udainoko/Documents/BSUIR/Sem4/OOP/Lab4/" + filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	case "2":
		err := s.Encrypt("/home/udainoko/Documents/BSUIR/Sem4/OOP/Lab4/" + filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	case "3":
		os.Exit(0)
	}
	return nil
}

func (s *Server) EcryptOpration(filename string) error {
	fmt.Println("Select if you want to procced")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Exit")

	var input string
	_, _ = fmt.Scanln(&input)
	switch input {
	case "1":
		err := s.Encrypt(filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	case "2":
		os.Exit(0)
	}
	return nil
}
