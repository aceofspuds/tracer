package tracer

import (
	"math"
	"os"
	"testing"
)

func TestCanvas(t *testing.T) {
	c := NewCanvas(5, 3)

	c.WritePixel(0, 0, Color(1.5, 0, 0))
	c.WritePixel(2, 1, Color(0, 0.5, 0))
	c.WritePixel(4, 2, Color(-0.5, 0, 1))
	s := c.ToPPM()
	if s != canvasPPM1 {
		t.Errorf("expected %s, returned %s", canvasPPM1, s)
	}

	c = NewCanvas(10, 2)

	for y := range c {
		for x := range c[y] {
			c.WritePixel(x, y, Color(1, 0.8, 0.6))
		}
	}

	s = c.ToPPM()
	if s != canvasPPM2 {
		t.Errorf("expected %s, returned %s", canvasPPM2, s)
	}

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

func TestClock(t *testing.T) {
	width := 900
	height := 900
	size := 3
	c := NewCanvas(width, height)

	p := Point(0, 1, 0)
	scale := ScalingMatrix(0, float64(height)/3., 0)
	translate := TranslationMatrix(float64(height)/2., float64(width)/2., 0)

	for i := 0; i < 12; i++ {
		transform := translate.Multiply(RotationZMatrix((float64(i) * 2. * math.Pi) / 12.).Multiply(scale))
		p1 := transform.MultiplyT(p)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				c.WritePixel(int(p1.x())+i, height-(int(p1.y())+j), Color(1, 0.5, 0.25))
			}
		}
	}

	// write to file
	f, err := os.Create("clock.ppm")
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

const canvasPPM1 string = `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255`

const canvasPPM2 string = `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`
