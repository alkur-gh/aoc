package graphics

func MakeBorders(xmin, xmax, ymin, ymax, zmin, zmax int) *Borders {
    return &Borders{xmin, xmax, ymin, ymax, zmin, zmax}
}

func MakeRectangle(x0, x1, y0, y1 int, r rune) *Rectangle {
    return &Rectangle{x0, x1, y0, y1, r}
}

func MakeCanvas(b *Borders) *Canvas {
    length := b.xmax - b.xmin + 1
    height := b.ymax - b.ymin + 1
    v := [][]rune{}
    for i := 0; i < height; i++ {
        v = append(v, make([]rune, height))
    }
    return &Canvas{length, height, v}
}
