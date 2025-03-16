package game

import (
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
)

func DrawRectangle(screen *ebiten.Image, x, y, width, height float64, clr color.Color) {
    rect := ebiten.NewImage(int(width), int(height))
    rect.Fill(clr)
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(x, y)
    screen.DrawImage(rect, op)
}
