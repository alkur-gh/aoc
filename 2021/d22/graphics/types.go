package graphics

type Borders struct {
    xmin, xmax, ymin, ymax, zmin, zmax int
}

type Line struct {
    x0, x1 int
    r rune
}

type Rectangle struct {
    x0, x1, y0, y1 int
    r rune
}

type Cuboid struct {
    x0, x1, y0, y1, z0, z1 int
    r rune
}

type Canvas struct {
    length, height int
    v [][]rune
}
