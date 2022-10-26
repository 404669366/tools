package tools

import (
	engname "github.com/yelinaung/eng-name"
	"math/rand"
	"time"
)

func RandEnName() string {
	seed := time.Now().UnixNano()
	randName := engname.New(seed)
	rand.Seed(seed)
	if rand.Intn(1) > 0 {
		return randName.GetMenName()
	}
	return randName.GetWomenName()
}

func RandEnMenName() string {
	seed := time.Now().UnixNano()
	randName := engname.New(seed)
	return randName.GetMenName()
}

func RandEnWomenName() string {
	seed := time.Now().UnixNano()
	randName := engname.New(seed)
	return randName.GetWomenName()
}
