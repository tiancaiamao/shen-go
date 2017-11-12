package main

import (
	"fmt"
	"os"

	"github.com/tiancaiamao/shen-go/kl"
)

func main() {
	r := kl.NewSexpReader(os.Stdin)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		res := kl.Eval(sexp)
		fmt.Println(kl.ObjString(res))
	}
}
