package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"

	b "spacegame/base"
)


type Rect struct {
	Sprite *b.Sprite
}

func NewRect(sprite *b.Sprite) *Rect {
	return &Rect{
		Sprite: sprite,
	}
}


// BoundingBoxCollision checks if two rectangles intersect using the Separating Axis Theorem.
// It doesn't check for pixel perfect collision.
func (r *Rect) BoundingBoxCollision(other *Rect) bool {
    if other == nil || r == nil {
        return false
    }
	return r.Sprite.Position.X <= other.MaxX() &&
		   other.Sprite.Position.X <= r.MaxX() &&
		   r.Sprite.Position.Y <= other.MaxY() &&
		   other.Sprite.Position.Y <= r.MaxY()
}

// MaxX returns the maximum X coordinate of the rectangle, which is the sum of the rectangle's
// position X and its width.
func (r *Rect) MaxX() float64 {
	return r.Sprite.Position.X + float64(r.Sprite.Size.Width)
}

// MaxY returns the maximum Y coordinate of the rectangle, which is the sum of the rectangle's
// position Y and its height.
func (r *Rect) MaxY() float64 {
	return r.Sprite.Position.Y + float64(r.Sprite.Size.Height)
}

// PixelPerfectCollision checks for a pixel-level collision between two rectangles.
// It first verifies if their bounding boxes overlap using BoundingBoxCollision. 
// If they do, it iterates over the overlapping pixels and compares the alpha values 
// of the pixels from both rectangles. If any overlapping pixel pair has non-zero 
// alpha values, indicating visible pixels, it returns true, indicating a collision. 
// Otherwise, it returns false.
func (r *Rect) PixelPerfectCollision(other *Rect) bool {
	if !r.BoundingBoxCollision(other) {
		return false
	}

	left := max(r.Sprite.Position.X, other.Sprite.Position.X)
    right := min(r.Sprite.Position.X+float64(r.Sprite.Size.Width), other.Sprite.Position.X+float64(other.Sprite.Size.Width))
    top := max(r.Sprite.Position.Y, other.Sprite.Position.Y)
    bottom := min(r.Sprite.Position.Y+float64(r.Sprite.Size.Height), other.Sprite.Position.Y+float64(other.Sprite.Size.Height))

    for y := top; y < bottom; y++ {
        for x := left; x < right; x++ {
            ax := int(x - r.Sprite.Position.X)
            ay := int(y - r.Sprite.Position.Y)
            bx := int(x - other.Sprite.Position.X)
            by := int(y - other.Sprite.Position.Y)

            if ax < 0 || ay < 0 || bx < 0 || by < 0 {
                continue
            }

            aAlpha := alphaAt(r.Sprite.Img, ax, ay)
            bAlpha := alphaAt(other.Sprite.Img, bx, by)

            if aAlpha > 0 && bAlpha > 0 {
                return true
            }
        }
    }

    return false
}

// alphaAt returns the alpha value of the pixel at (x, y) in img, or 0 if the pixel is out of bounds.
func alphaAt(img *ebiten.Image, x, y int) uint32 {
    c := img.At(x, y).(color.RGBA)
    return uint32(c.A)
}

// IsBetween returns true if x is between start and end.
func IsBetween(x, start, end float64) bool {
	return x <= end && x >= start
}