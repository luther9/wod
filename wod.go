/*
wod - A World of Darkness dice roller.

Written in 2016 by Luther Thompson, lutheroto@gmail.com

To the extent possible under law, the author(s) have dedicated all copyright and
related and neighboring rights to this software to the public domain worldwide.
This software is distributed without any warranty.

You should have received a copy of the CC0 Public Domain Dedication along with
this software. If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func badArgument() {
	fmt.Println("Need an integer argument")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		badArgument()
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		badArgument()
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := []int{}
	for i := 0; i < n; i++ {
		for {
			r := rng.Intn(10)
			roll = append(roll, r)
			if r != 9 {
				break
			}
		}
	}
	fmt.Println(roll)
	successes := 0
	for _, r := range roll {
		if r >= 7 {
			successes++
		}
	}
	fmt.Println(successes)
}
