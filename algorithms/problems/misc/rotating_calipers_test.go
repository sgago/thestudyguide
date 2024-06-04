package misc

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
)

// Point represents a point in 2D space
type Point struct {
	x, y float64
}

// byX implements sort.Interface for []Point based on the x field.
type byX []Point

func (a byX) Len() int           { return len(a) }
func (a byX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byX) Less(i, j int) bool { return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].y < a[j].y) }

// cross product of vectors OA and OB
// a positive cross product indicates a counter-clockwise turn, negative indicates a clockwise turn, and zero indicates collinear points
func cross(o, a, b Point) float64 {
	return (a.x-o.x)*(b.y-o.y) - (a.y-o.y)*(b.x-o.x)
}

// ConvexHull computes the convex hull of a set of 2D points using Andrew's monotone chain algorithm
func ConvexHull(points []Point) []Point {
	sort.Sort(byX(points))

	lower := []Point{}
	for _, p := range points {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	upper := []Point{}
	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}

// distance calculates the Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

// maxDistance finds the maximum distance between any two points in a slice using Rotating Calipers
func maxDistanceRotatingCalipers(points []Point) float64 {
	hull := ConvexHull(points)
	n := len(hull)
	if n == 0 {
		return 0
	}

	maxDist := 0.0
	k := 1
	for i := 0; i < n; i++ {
		for {
			currentDist := distance(hull[i], hull[(i+1)%n])
			nextDist := distance(hull[i], hull[(k+1)%n])
			if nextDist > currentDist {
				k = (k + 1) % n
			} else {
				break
			}
		}
		maxDist = math.Max(maxDist, distance(hull[i], hull[k]))
	}
	return maxDist
}

func TestRotatingCalipers(t *testing.T) {
	points := []Point{
		{x: 0, y: 0},
		{x: 1, y: 2},
		{x: 4, y: 6},
		{x: 7, y: 8},
		{x: 2, y: 1},
		{x: 3, y: 5},
		{x: 8, y: 8},
		{x: 9, y: 1},
		{x: 100, y: 0},
	}

	fmt.Printf("The maximum distance between any two points is: %f\n", maxDistanceRotatingCalipers(points))
}

// GenerateRandomPoints generates a slice of random points
func GenerateRandomPoints(numPoints int) []Point {
	points := make([]Point, numPoints)
	for i := 0; i < numPoints; i++ {
		points[i] = Point{
			x: rand.Float64() * 1000,
			y: rand.Float64() * 1000,
		}
	}
	return points
}

// maxDistanceBruteForce finds the maximum distance between any two points using brute force
func maxDistanceBruteForce(points []Point) float64 {
	maxDist := 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := distance(points[i], points[j])
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return maxDist
}

// Benchmark for Brute Force Method
func BenchmarkMaxDistanceBruteForce(b *testing.B) {
	points := GenerateRandomPoints(1000)
	for i := 0; i < b.N; i++ {
		maxDistanceBruteForce(points)
	}
}

// Benchmark for Convex Hull + Rotating Calipers Method
func BenchmarkMaxDistanceRotatingCalipers(b *testing.B) {
	points := GenerateRandomPoints(1000)
	for i := 0; i < b.N; i++ {
		maxDistanceRotatingCalipers(points)
	}
}
