package main

import (
    "time"
    "math/rand"
    "log"
    "os"
    "encoding/json"
    "flag"
)

var (
    Filename = flag.String("f", "../configs/enigma.json", "output filename")
    SizeRotors = flag.Int("sr", 3, "number of rotors")
    SizeAlphabet = flag.Int("as", 256, "size of alphabet")
)

type Enigma struct {
    SizeAlphabet int `json:"size_alphabet"`
    SizeRotors int `json:"size_rotors"`
    Rotors [][]int `json:"rotors"`
    Reflector []int `json:"reflector"`
    Panel []int `json:"panel"`
}

func shuffle_pairs(n int) []int {
	s2 := make([]int, n/2)

	for i := range s2 {
		s2[i] = n/2 + i
	}

	shuffled := make([]int, n)
	k := 0
	for i := len(s2); i > 0; i-- {
		index := rand.Intn(i)

		shuffled[k] = s2[index]

		s2[index] = s2[i-1]

		k++
	}

	for i := 0; i < n/2; i++ {
		shuffled[shuffled[i]] = i
	}

	return shuffled
}

func main() {
    flag.Parse()

    rotors := make([][]int, *SizeRotors)

    rand.Seed(time.Now().Unix())
    for i := range rotors {
        rotors[i] = rand.Perm(*SizeAlphabet)
    }
    reflector := shuffle_pairs(*SizeAlphabet)
    panel := shuffle_pairs(*SizeAlphabet)

    file, err := os.OpenFile(*Filename, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        log.Fatal("who hui (open)")
    }

    defer file.Close()

    e := Enigma{
        SizeAlphabet: *SizeAlphabet,
        SizeRotors: *SizeRotors,
        Rotors: rotors,
        Reflector: reflector,
        Panel: panel,
    }

    b, err := json.MarshalIndent(&e, "", "  ")
    if err != nil {
        log.Fatal("cant fork (marshal)")
    }

    _, err = file.Write(b)
    if err != nil {
            log.Fatal("in english (write)")
    }
}