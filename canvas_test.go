package tracer

import (
	"fmt"
	"os"
	"testing"
)

func TestCanvas(t *testing.T) {
	c := NewCanvas(5, 3)

	c.WritePixel(0, 0, Color(1.5, 0, 0))
	c.WritePixel(2, 1, Color(0, 0.5, 0))
	c.WritePixel(4, 2, Color(-0.5, 0, 1))

	fmt.Println(c.ToPPM())

	c = NewCanvas(10, 2)

	for y := range c {
		for x := range c[y] {
			c.WritePixel(x, y, Color(1, 0.8, 0.6))
		}
	}

	fmt.Println(c.ToPPM())
}

func TestPicture(t *testing.T) {
	type env struct {
		gravity Tuple
		wind    Tuple
	}

	type projectile struct {
		position Tuple
		velocity Tuple
	}

	tick := func(e env, p projectile) projectile {
		return projectile{
			position: p.position.Add(p.velocity),
			velocity: p.velocity.Add(e.gravity.Add(e.wind)),
		}
	}

	p := projectile{Point(0, 1, 0), Vector(1, 1.8, 0).Normalize().Multiply(11.25)}
	e := env{Vector(0, -0.1, 0), Vector(-0.01, 0, 0)}

	width := 900
	height := 550
	size := 3
	c := NewCanvas(width, height)

	for p.position.x() >= 0 && p.position.x()+float64(size) < float64(width) &&
		p.position.y() >= 0 && p.position.y()+float64(size) < float64(height) {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				c.WritePixel(int(p.position.x())+i, height-(int(p.position.y())+j), Color(1, 0.5, 0.25))
			}
		}

		p = tick(e, p)
	}

	// write to file
	f, err := os.Create("test.ppm")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(c.ToPPM())
	if err != nil {
		t.Error(err)
		return
	}
}
