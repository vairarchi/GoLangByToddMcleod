package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("-------------------------")
	fmt.Println("context:", ctx)
	fmt.Println("Context error:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
	fmt.Println("-------------------------")
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("context:", ctx)
	fmt.Println("Context error:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
	fmt.Println("cancel:", cancel)
	fmt.Printf("Cancel type: %T\n", cancel)
	fmt.Println("-------------------------")

	cancel()
	fmt.Println("context:", ctx)
	fmt.Println("Context error:", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
	fmt.Println("cancel:", cancel)
	fmt.Printf("Cancel type: %T\n", cancel)
	fmt.Println("-------------------------")
}
