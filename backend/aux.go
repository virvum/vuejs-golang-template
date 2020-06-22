package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// RandBytes returns a byte slice of n random bytes.
func RandBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("rand.Read: %v", err)
	}

	return b, nil
}

// RandToken returns a random token, which is a SHA256 sum of 128 random bytes.
func RandToken() (string, error) {
	var n uint32 = 128

	b, err := RandBytes(n)
	if err != nil {
		return "", fmt.Errorf("RandBytes(%v): %v", n, err)
	}

	h := sha256.New()
	h.Write(b)

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// RandString returns a randon alphanumeric string of length n.
func RandString(n uint32) (string, error) {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b, err := RandBytes(n)
	if err != nil {
		return "", fmt.Errorf("RandBytes(%v): %v", n, err)
	}

	for i := range b {
		b[i] = charSet[int(b[i])%len(charSet)]
	}

	return string(b), nil
}

// FormatDuration is like Go's Duration.String(), but unit values are rounded. The parameter `min` specifies at which unit to start.
func FormatDuration(d time.Duration, min string) string {
	units := []struct {
		Name   string
		Length uint64
	}{
		{"ns", 1000},
		{"µs", 1000},
		{"ms", 1000},
		{"s", 60},
		{"m", 60},
		{"h", 24},
		{"d", 365},
		{"y", 0},
	}

	u := uint64(d.Nanoseconds())
	var s []string
	var start = false

	switch min {
	case "ns", "µs", "ms", "s", "m", "h", "g", "y":
	default:
		panic("invalid parameter min")
	}

	if u == 0 {
		return "0ns"
	}

	for _, unit := range units {
		if unit.Name == min {
			start = true
		}

		if unit.Length > 0 {
			if start || u < unit.Length {
				if m := u % unit.Length; m > 0 {
					s = append([]string{fmt.Sprintf("%d%s", m, unit.Name)}, s...)
				}
			}

			u = uint64(u / unit.Length)

			if u == 0 {
				break
			}
		} else {
			s = append([]string{fmt.Sprintf("%d%s", u, unit.Name)}, s...)
		}
	}

	return strings.Join(s, "")
}

const (
	kibibyte float64 = 1024
	mebibyte         = kibibyte * 1024
	gibibyte         = mebibyte * 1024
	tebibyte         = gibibyte * 1024
	pebibyte         = tebibyte * 1024
	exbibyte         = pebibyte * 1024
)

// BytesToSize converts b to a human-readable size.
func BytesToSize(b uint64) string {
	var (
		u string
		v float64 = float64(b)
	)

	switch {
	case v >= exbibyte:
		u = "EiB"
		v /= exbibyte
	case v >= tebibyte:
		u = "TiB"
		v /= tebibyte
	case v >= gibibyte:
		u = "GiB"
		v /= gibibyte
	case v >= mebibyte:
		u = "MiB"
		v /= mebibyte
	case v >= kibibyte:
		u = "KiB"
		v /= kibibyte
	default:
		return fmt.Sprintf("%d B", int(v))
	}

	return fmt.Sprintf("%.1f %s", v, u)
}

// OxfordJoin joins a slice of strings the Oxford-way, which means that all
// elements are joined by ", ", except for the last two, which are joined by
// ", w", where w is parameter w.
func OxfordJoin(values []string, format string, w string) string {
	var s strings.Builder

	for i := 0; i < len(values); i++ {
		switch {
		case i > 0 && i < len(values)-1:
			s.WriteString(", ")
		case i > 0:
			s.WriteString(", ")
			s.WriteString(w)
			s.WriteString(" ")
		}

		s.WriteString(fmt.Sprintf(format, values[i]))
	}

	return s.String()
}
