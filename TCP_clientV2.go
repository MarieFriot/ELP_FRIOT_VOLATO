import (
	"Golang_project/quicksort"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func bigListeGen(size int) []int {
	liste := []int{}
	for i := 0; i < size; i++ {
		liste = append(liste, rand.Intn(1000))
	}
	return liste
}

func test_client() []int {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer conn.Close()

	// Data to send
	data := bigListeGen(250)
	print(data)
	sizeData := len(data)

	// 1st step: Send data size
	// Conversion into bytes for buffer
	buffLen := make([]byte, 4)
	binary.BigEndian.PutUint32(buffLen, uint32(sizeData))

	// Send data to the server
	_, err = conn.Write(buffLen)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Determining amount of buffers of size x to send (here x = 100)
	sizeBuff := 100
	var iteration int
	iteration = (sizeData / sizeBuff) + 1

	// 2nd step: Sending segments of the data to the server
	for i := 0; i < iteration; i++ {

		buffSend := make([]byte, 0)

		if i == (iteration - 1) {
			buffSend = make([]byte, 4*(sizeData-i*sizeBuff))
			for j, v := range data[i*sizeBuff : sizeData] {
				binary.BigEndian.PutUint32(buffSend[j*4:], uint32(v))
			}
		} else {
			// Conversion into bytes for buffer
			buffSend = make([]byte, 4*sizeBuff)
			for j, v := range data[i*sizeBuff : (i+1)*sizeBuff] {
				binary.BigEndian.PutUint32(buffSend[j*4:], uint32(v))
			}
		}

		// Send data to the server
		_, err = conn.Write(buffSend)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	}

	// Read and process data from the server and convert the binary data back to integersm [TBD]
	var dataTreated []int
	for i := 0; i < iteration; i++ {

		buffRecep := make([]byte, 0)

		if i == (iteration - 1) {
			buffRecep = make([]byte, 4*(sizeData-i*sizeBuff))
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return nil
			}

			for j := 0; j < 4*(sizeData-i*sizeBuff); j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))
				dataTreated = append(dataTreated, val)
			}

		} else {
			buffRecep = make([]byte, 4*sizeBuff)
			_, err := conn.Read(buffRecep)
			if err != nil {
				fmt.Println("Error reading data:", err)
				return nil
			}

			for j := 0; j < 4*sizeBuff; j += 4 {
				val := int(binary.BigEndian.Uint32(buffRecep[j : j+4]))
				dataTreated = append(dataTreated, val)
			}
		}
	}

	return dataTreated
}
