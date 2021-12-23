package graphics

func MakeCuboid(x0, x1, y0, y1, z0, z1 int, r rune) *Cuboid {
    return &Cuboid{x0, x1, y0, y1, z0, z1, r}
}

func (c *Cuboid) XLine() *Line {
    return &Line{c.x0, c.x1, ' '}
}

func (c *Cuboid) YLine() *Line {
    return &Line{c.y0, c.y1, ' '}
}

func (c *Cuboid) ZLine() *Line {
    return &Line{c.z0, c.z1, ' '}
}

func (c *Cuboid) IsValid() bool {
    return (c.x0 <= c.x1) && (c.y0 <= c.y1) && (c.z0 <= c.z1)
}

func (c *Cuboid) Points() (int, int, int, int, int, int) {
    return c.x0, c.x1, c.y0, c.y1, c.z0, c.z1
}

func (c *Cuboid) Volume() int {
    return (c.x1 - c.x0 + 1) * (c.y1 - c.y0 + 1) * (c.z1 - c.z0 + 1)
}

func (c1 *Cuboid) IsIntersect(c2 *Cuboid) bool {
    return c1.XLine().IsIntersect(c2.XLine()) &&
           c1.YLine().IsIntersect(c2.YLine()) &&
           c1.ZLine().IsIntersect(c2.ZLine())
}

func (base *Cuboid) Difference(diff *Cuboid) []*Cuboid {
    if !base.IsIntersect(diff) {
        return []*Cuboid{base}
    }

    cs := []*Cuboid{}

    xi := base.XLine().Intersect(diff.XLine())
    yi := base.YLine().Intersect(diff.YLine())
    zi := base.ZLine().Intersect(diff.ZLine())

    cs = append(cs, &Cuboid{
        base.x0, xi.x0 - 1,
        base.y0, yi.x0 - 1,
        base.z0, zi.x0 - 1,
        '1',
    }, &Cuboid{
        xi.x0, xi.x1,
        base.y0, yi.x0 - 1,
        base.z0, zi.x0 - 1,
        '2',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        base.y0, yi.x0 - 1,
        base.z0, zi.x0 - 1,
        '3',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        yi.x0, yi.x1,
        base.z0, zi.x0 - 1,
        '4',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        yi.x1 + 1, base.y1,
        base.z0, zi.x0 - 1,
        '5',
    }, &Cuboid{
        xi.x0, xi.x1,
        yi.x1 + 1, base.y1,
        base.z0, zi.x0 - 1,
        '6',
    }, &Cuboid{
        base.x0, xi.x0 - 1,
        yi.x1 + 1, base.y1,
        base.z0, zi.x0 - 1,
        '7',
    }, &Cuboid{
        base.x0, xi.x0 - 1,
        yi.x0, yi.x1,
        base.z0, zi.x0 - 1,
        '8',
    }, &Cuboid{
        xi.x0, xi.x1,
        yi.x0, yi.x1,
        base.z0, zi.x0 - 1,
        '9',
    })

    cs = append(cs, &Cuboid{
        base.x0, xi.x0 - 1,
        base.y0, yi.x0 - 1,
        zi.x1 + 1, base.z1,
        '1',
    }, &Cuboid{
        xi.x0, xi.x1,
        base.y0, yi.x0 - 1,
        zi.x1 + 1, base.z1,
        '2',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        base.y0, yi.x0 - 1,
        zi.x1 + 1, base.z1,
        '3',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        yi.x0, yi.x1,
        zi.x1 + 1, base.z1,
        '4',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        yi.x1 + 1, base.y1,
        zi.x1 + 1, base.z1,
        '5',
    }, &Cuboid{
        xi.x0, xi.x1,
        yi.x1 + 1, base.y1,
        zi.x1 + 1, base.z1,
        '6',
    }, &Cuboid{
        base.x0, xi.x0 - 1,
        yi.x1 + 1, base.y1,
        zi.x1 + 1, base.z1,
        '7',
    }, &Cuboid{
        base.x0, xi.x0 - 1,
        yi.x0, yi.x1,
        zi.x1 + 1, base.z1,
        '8',
    }, &Cuboid{
        xi.x0, xi.x1,
        yi.x0, yi.x1,
        zi.x1 + 1, base.z1,
        '9',
    })

    cs = append(cs, &Cuboid{
        base.x0, xi.x0 - 1,
        base.y0, yi.x0 - 1,
        zi.x0, zi.x1,
        '1',
    }, &Cuboid{
        xi.x0, xi.x1,
        base.y0, yi.x0 - 1,
        zi.x0, zi.x1,
        '2',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        base.y0, yi.x0 - 1,
        zi.x0, zi.x1,
        '3',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        yi.x0, yi.x1,
        zi.x0, zi.x1,
        '4',
    }, &Cuboid{
        xi.x1 + 1, base.x1,
        yi.x1 + 1, base.y1,
        zi.x0, zi.x1,
        '5',
    }, &Cuboid{
        xi.x0, xi.x1,
        yi.x1 + 1, base.y1,
        zi.x0, zi.x1,
        '6',
    }, &Cuboid{
        base.x0, xi.x0 - 1,
        yi.x1 + 1, base.y1,
        zi.x0, zi.x1,
        '7',
    }, &Cuboid{
        base.x0, xi.x0 - 1,
        yi.x0, yi.x1,
        zi.x0, zi.x1,
        '8',
    })

    valid := []*Cuboid{}
    for _, c := range cs {
        if c.IsValid() {
            valid = append(valid, c)
        }
    }
    return valid
}
