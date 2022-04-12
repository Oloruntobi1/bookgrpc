package main

import (
	"fmt"
	"net"
	"time"

	"github.com/Oloruntobi1/bookgrpc/serverstream/gateway"
	documents "github.com/Oloruntobi1/bookgrpc/serverstream/v1/dok"

	"log"

	"google.golang.org/grpc"
)

type server struct {
	documents.UnimplementedDocumentsServer
}

func main() {
	fmt.Println("Starting server...")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	documents.RegisterDocumentsServer(s, &server{})

	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }

	log.Println("Serving gRPC on 0.0.0.0:3000")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gateway.StartGateway()
}

// container struct
type container struct {
	documents []*documents.Document
}

// GetDocuments function
func (*server) GetDocuments(req *documents.EmptyReq, stream documents.Documents_GetDocumentsServer) error {
	fmt.Println("GetDocuemnts function")

	// Initialize the container struct and call the initDocuments function
	// to get dummy data to send on the stream response message.
	docs := container{}.initDocuments()

	// Iterate over the documents
	for _, v := range docs {
		// Run some validation on each object
		if v.Size > 250 {
			// Create the response object
			res := &documents.GetDocumentsRes{
				Document: v,
			}

			// Use the stream object to send the response stream message
			stream.Send(res)

			// Sleep for a little bit..
			time.Sleep(1000 * time.Millisecond)
		}
	}
	return nil
}

// initDocuments function
func (c container) initDocuments() []*documents.Document {
	c.documents = append(c.documents, c.getDocument("Doc One", "nat", 345))
	c.documents = append(c.documents, c.getDocument("Doc Tow", "zip", 245))
	c.documents = append(c.documents, c.getDocument("Doc Three", "nat", 445))
	c.documents = append(c.documents, c.getDocument("Doc Four", "pid", 545))
	c.documents = append(c.documents, c.getDocument("Doc Five", "nat", 145))
	return c.documents
}

// getDocument function
func (c container) getDocument(name, documentType string, size int64) *documents.Document {
	return &documents.Document{
		Name:         name,
		DocumentType: documentType,
		Size:         size,
	}
}