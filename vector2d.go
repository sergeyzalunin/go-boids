package main

import "math"

// Vector2D is main structure to control boids
type Vector2D struct {
	x float64
	y float64
}

// Add multiplicates two vectors
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v1.x + v2.x, v1.y + v2.y}
}

// Subtract between two vectors
func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{v1.x - v2.x, v1.y - v2.y}
}

// Multiply two vectors
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{v1.x * v2.x, v1.y * v2.y}
}

// AddV adds number to vector position
func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{v1.x + d, v1.y + d}
}

// MultiplyV multiplies number to vector position
func (v1 Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{v1.x * d, v1.y * d}
}

// DivisionV divides number from vecto position
func (v1 Vector2D) DivisionV(d float64) Vector2D {
	return Vector2D{v1.x / d, v1.y / d}
}

// Limit creates new vector with specified limit
func (v1 Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{
		math.Min(math.Max(v1.x, lower), upper),
		math.Min(math.Max(v1.y, lower), upper),
	}
}

// Distance calculates distance between two vectors
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
