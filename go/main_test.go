package main

import "testing"

func TestGcd(t *testing.T) {
	patterns := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 8, 2},
		{4, 44, 4},
		{11, 33, 11},
	}

	for i, pattern := range patterns {
		ans := gcd(pattern.x, pattern.y)
		if pattern.expected != ans {
			t.Errorf("pattern %d: expected %d: actual; %d", i, pattern.expected, ans)
		}
	}
}

func TestLcm(t *testing.T) {
	patterns := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 4, 4},
		{3, 5, 15},
		{4, 7, 28},
	}

	for i, pattern := range patterns {
		ans := lcm(pattern.x, pattern.y)
		if pattern.expected != ans {
			t.Errorf("pattern %d: expected %d: actual; %d", i, pattern.expected, ans)
		}
	}
}
