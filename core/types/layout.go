package types

const (
	PADDING        = 100
	DEFAULT_HEIGHT = 1024
	DEFAULT_WIDTH  = 768
)

type Item struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Row struct {
	Height int
	Width  int

	Items []Item
}

type ScreenSpace struct {
	Remaining_Height int
	Remaining_Width  int
	Rows             []*Row
}

func Create_Flexbox() *ScreenSpace {
	ss := ScreenSpace{
		Remaining_Height: DEFAULT_HEIGHT,
		Remaining_Width:  DEFAULT_WIDTH,
		Rows:             make([]*Row, 0),
	}
	ss.Rows = append(ss.Rows, Create_Row())
	return &ss
}

func (ss *ScreenSpace) Add_Item(x int, y int, height int, width int) (float32, float32, float32, float32) {
	ss.Remaining_Space()
	rec := &Item{float32(x), float32(y), float32(height), float32(width)}
	// If out of width create a new row
	if ss.Remaining_Width-width <= 0 {
		ss.Rows = append(ss.Rows, Create_Row())
		ss.Remaining_Width = DEFAULT_WIDTH
	}
	selected_row := ss.Rows[len(ss.Rows)-1]
	rec.X = float32(DEFAULT_WIDTH - ss.Remaining_Width + PADDING)
	rec.Y = float32(DEFAULT_HEIGHT - ss.Remaining_Height + PADDING)
	selected_row.Add_Item(rec)
	return rec.X, rec.Y, rec.Width, rec.Height
}

func (ss *ScreenSpace) Remaining_Space() {
	var total_height int
	for r := range ss.Rows {
		total_height += ss.Rows[r].Height
	}
	ss.Remaining_Height = DEFAULT_HEIGHT - total_height
	ss.Remaining_Width -= ss.Rows[len(ss.Rows)-1].Width
}

func (row *Row) Add_Item(rec *Item) {
	row.Items = append(row.Items, *rec)
	row.Width += int(rec.Width)
	if rec.Height > float32(row.Height) {
		row.Height = int(rec.Height)
	}
}

func Create_Row() (row *Row) {
	return &Row{
		Height: 0,
		Width:  0,
		Items:  []Item{},
	}
}
