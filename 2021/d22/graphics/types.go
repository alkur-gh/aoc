package graphics

type Borders struct {
    xmin, xmax, ymin, ymax, zmin, zmax int
}

type Rectangle struct {
    x0, x1, y0, y1 int
    r rune
}

type Canvas struct {
    length, height int
    v [][]rune
}
