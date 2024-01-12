import (
	"Golang_project/quicksort"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	
	// 1st step : Read the length of the data (to optimize buffer size if possible)
	buffLen := make([]byte, 4)
	_, err := conn.Read(buffLen)
	if err != nil {
		fmt.Println("Error reading length:", err)
		return
	}

	sizeData := int(binary.BigEndian.Uint32(buffLen))

	// Doing the partition step and read data from client.
	var l1 []int
	var l2 []int
	var pivot int
	bufferReceive := make([]byte, 4)
	for i := 0; i < sizeData; i++ {
		_, err := conn.Read(bufferReceive)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		val := int(binary.BigEndian.Uint32(bufferReceive))
		if i == 0 {
			pivot = val
		}

		if val <= pivot {
			l1 = append(l1, val)
		} else {
			l2 = append(l2, val)
		}
	}

	// Perform treatment on data
	result1 := quicksort.QuicksortParallel(l1)
	result2 := quicksort.QuicksortParallel(l2)

	var data []int
	data = append(data, result1...)
	data = append(data, result2...)

	// Sending back segments of the data to the client
	for k := 0; k < sizeData ; k++ {
		
		// Conversion into bytes for buffer
		bufferSend := make([]byte, 4)
		binary.BigEndian.PutUint32(bufferSend, uint32(data[k]))

		// Send data to the server
		_, err = conn.Write(bufferSend)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Data treated and response sent successfully")
}

func test_server() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080.")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}
