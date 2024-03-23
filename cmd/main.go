package main

import (
	"context"
	"flag"
	"fmt"
	"task/internal/floodcontrol"
)

func main() {

	n := flag.Int("n", 10, "Количество секунд для проверки")
	k := flag.Int("k", 5, "Максимальное количество вызовов")
	flag.Parse()

	fc := floodcontrol.NewFloodControl(*n, *k)
	ctx := context.Background()

	ok, err := fc.Check(ctx, 2)
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("Флуд-контроль не пройден")
	} else {
		fmt.Println("Флуд-контроль пройден")
	}
}
