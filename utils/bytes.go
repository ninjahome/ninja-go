package utils

import "fmt"

const (
	K = 1 << 10
	M = 1 << 20
	G = 1 << 30
	T = 1 << 40
)

type ByteSize int64

func (bs ByteSize) String() string {
	if bs > T {
		return fmt.Sprintf("%.2fT", float64(bs)/float64(T))
	}
	if bs > G {
		return fmt.Sprintf("%.2fG", float64(bs)/float64(G))
	}
	if bs > M {
		return fmt.Sprintf("%.2fM", float64(bs)/float64(M))
	}
	if bs > K {
		return fmt.Sprintf("%.2fK", float64(bs)/float64(K))
	}
	return fmt.Sprintf("%d", bs)

}
