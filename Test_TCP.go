package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read and process data from the client
	buff := make([]byte, 4*10) // Assuming a maximum of 10 integers
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Convert the binary data back to integers
	var data []int
	for i := 0; i < len(buff); i += 4 {
		val := int(binary.BigEndian.Uint32(buff[i : i+4]))
		data = append(data, val)
	}

	// Perform treatment on data
	//inserer quicksort

	// Convert the treated data back to binary
	responseBuff := make([]byte, 4*len(data))
	for i, v := range data {
		binary.BigEndian.PutUint32(responseBuff[i*4:], uint32(v))
	}

	// Send the treated data back to the client
	_, err = conn.Write(responseBuff)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
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

	fmt.Println("Server is listening on port 8080")

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

func test_client() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Data to send
	data := []int{1, 4, 6, 8, 3, 5, 7, 9, 2}

	// Conversion into bytes for buffer
	buff := make([]byte, 4*len(data))
	for i, v := range data {
		binary.BigEndian.PutUint32(buff[i*4:], uint32(v))
	}

	// Send data to the server
	_, err = conn.Write(buff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and process data from the server
	buffAnswer := make([]byte, 4*10) // Assuming a maximum of 10 integers
	_, err = conn.Read(buffAnswer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Convert the binary data back to integers
	var dataTreated []int
	for i := 0; i < len(buff); i += 4 {
		val := int(binary.BigEndian.Uint32(buffAnswer[i : i+4]))
		dataTreated = append(dataTreated, val)
	}
}
