import (
	"Golang_project/quicksort"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func client() []int {
	// Connect to the server [Done]
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer conn.Close()

	// Data to send
	data := []int{14, 8, 2, 6, 18, 13, 9, 7, 11, 1, 16, 20, 10, 5, 19, 15, 17, 4, 3, 12}
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
	
	// 2nd step: Sending segments of the data to the server
	for k := 0; k < sizeData ; k++ {
		
		// Conversion into bytes for buffer
		bufferSend := make([]byte, 4)
		binary.BigEndian.PutUint32(bufferSend, uint32(data[k]))
		
		// Send data to the server
		_, err = conn.Write(bufferSend)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
	}

	// Read and process data from the server and convert the binary data back to integers
	var dataTreated []int
	bufferReceive := make([]byte, 4)
	for i := 0; i < sizeData; i++ {
		_, err := conn.Read(bufferReceive)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return nil
		}
				
		val := int(binary.BigEndian.Uint32(bufferReceive))
		dataTreated = append(dataTreated, val)
		}

	return dataTreated
}
