package qasircore

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type Tracing struct {
}

func (t *Tracing) createMetadata(ctx context.Context) context.Context {
	var result context.Context
	u4 := uuid.Must(uuid.NewRandom())
	md := metadata.New(map[string]string{"uuid": fmt.Sprintf("%d", u4)})
	log.Println(md)
	result = metadata.NewOutgoingContext(context.Background(), md)

	return result
}
