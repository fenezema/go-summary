package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func ret(a, b int) (int, error) {
	return a - b, errors.New("lala")
}

func main() {
	g, _ := errgroup.WithContext(context.Background())

	_, err := add(1, 3)
	res1 := 0
	res2 := 0
	res3 := 0
	g.Go(func() error {
		res1, err = add(1, 2)
		return err
	})
	g.Go(func() error {
		res2, err = ret(3, 2)
		return err
	})
	g.Go(func() error {
		res3, err = add(1, 2)
		return err
	})

	// Wait group
	if err := g.Wait(); err != nil {
		fmt.Println("lala")
	} else {
		if res1 > 0 && res2 > 0 && res3 > 0 {
			fmt.Println("hehe")
		} else {
			fmt.Println("not hehe")
		}
	}
}
