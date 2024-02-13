package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/mircodezorzi/go-grpc-bench/pb"
)

func Bytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, sizes)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}

	return fmt.Sprintf(f, val, suffix)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randLabels(n, m int) (l []*pb.Label) {
	for i := 0; i < n; i++ {
		l = append(l, &pb.Label{
			Name:  randString(m),
			Value: randString(m),
		})
	}
	return
}

func randSamples(n int) (s []*pb.Sample) {
	for i := 0; i < n; i++ {
		s = append(s, &pb.Sample{
			Value:     rand.Float64(),
			Timestamp: rand.Int63(),
		})
	}
	return
}

func randomFolders(n, m, e int) (f map[string]*pb.Folder) {
	f = make(map[string]*pb.Folder)
	for i := 0; i < n; i++ {
		f[randString(10)] = &pb.Folder{
			Description: randString(100),
			People:      randomPeople(m, e),
		}
	}
	return
}

func randomPeople(n, m int) (p []*pb.Person) {
	for i := 0; i < n; i++ {
		person := &pb.Person{
			Name:   randString(20),
			Age:    rand.Int31(),
			Emails: randomStrings(m, 10),
		}

		if rand.Int()%2 == 0 {
			person.ContactInfo = &pb.Person_Address{
				Address: &pb.Address{
					Street:  randString(10),
					City:    randString(10),
					State:   randString(10),
					ZipCode: randString(10),
				},
			}
		} else {
			person.ContactInfo = &pb.Person_PhoneNumber{
				PhoneNumber: &pb.PhoneNumber{
					Number: randString(10),
					Type:   pb.PhoneType(rand.Int() % 4),
				},
			}
		}

		p = append(p, person)
	}
	return
}

func randomStrings(n, m int) (s []string) {
	for i := 0; i < n; i++ {
		s = append(s, randString(m))
	}
	return
}
