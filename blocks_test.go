package main

import (
	"fmt"
	"testing"
)

func TestAddFigureOutsideFrame(t *testing.T) {
	b := newBlocks(80, 25)
	f := newFigure(0, 0, 0, 79)

	for i := 0; i < 100; i++ {

		b.addFigure(f)

		if b.areHere(-1, 78) {
			t.Fatal("b[79,0] is busy")
		}

		if b.areHere(-1, 79) {
			t.Fatal("b[79,-1] is busy")
		}

		if b.areHere(0, 78) {
			t.Fatal("b[78,0] is busy")
		}

		if b.areHere(0, 79) {
			t.Fatal("b[79,0] is busy")
		}
	}
}

func TestAddFigureInsideFrame(t *testing.T) {
	b := newBlocks(25, 80)

	x := 4
	y := 60

	f := newFigure(0, 0, x, y)

	b.addFigure(f)

	for bn, fb := range f.blocks() {

		line_number := fb[1]
		column_number := fb[0]

		if !b.areHere(column_number, line_number) {
			t.Fatalf("(%d) b[%d, %d] not busy\nFigure.blocks(): %v", bn, column_number, line_number, f.blocks())
		}
	}

}

func TestAddFigureInsideFrame100Times(t *testing.T) {
	b := newBlocks(25, 80)
	f := newFigure(0, 0, 1, 79)

	for i := 0; i < 100; i++ {

		b.addFigure(f)

		for bn, fb := range f.blocks() {

			line_number := fb[1]
			column_number := fb[0]

			t.Run(fmt.Sprintf("Block[%d,%d] should be busy", column_number, line_number), func(t *testing.T) {
				if !b.areHere(column_number, line_number) {
					t.Fatalf("(%d) b[%d, %d] not busy\nFigure.blocks(): %v", bn, column_number, line_number, f.blocks())
				}
			})
		}

		for line_number := 0; line_number < b.height; line_number++ {
			for column_number := 0; column_number < b.width; column_number++ {
				if !f.isHere(column_number, line_number) {
					t.Run(fmt.Sprintf("Block[%d,%d] should NOT be busy", column_number, line_number), func(t *testing.T) {
						if b.areHere(column_number, line_number) {
							t.Fatalf("b[%d, %d] IS busy\nFigure.blocks(): %v", column_number, line_number, f.blocks())
						}
					})
				}
			}
		}
	}
}

func TestFillLine(t *testing.T) {
	b := newBlocks(20, 80)

	for figure_number := 0; figure_number < 10; figure_number++ {

		f := newFigure(0, 0, figure_number*2, 79)

		b.addFigure(f)

	}

	t.Run("Blocks should form deletable line", func(t *testing.T) {
		if !b.deletable_lines {
			t.Fatalf("b.deletable_lines is: %v", b.deletable_lines)
		}
	})

}

func TestFillLine100Times(t *testing.T) {
	b := newBlocks(20, 20)

	for round_number := 0; round_number < 100; round_number++ {

		for figure_number := 0; figure_number < 10; figure_number++ {

			f := newFigure(0, 0, figure_number*2, 19)

			b.addFigure(f)

		}

		t.Run(fmt.Sprintf("(%d) Blocks should form deletable line", round_number), func(t *testing.T) {
			if !b.deletable_lines {
				t.Fatalf("(%d) b.deletable_lines is: %v", round_number, b.deletable_lines)
			}
		})

		lines_deleted := b.delete_deletable_lines()

		t.Run(fmt.Sprintf("(%d) 2 deletable lines should be deleted", round_number), func(t *testing.T) {
			if lines_deleted != 2 {
				t.Fatalf("(%d) b.delete_deletable_lines returned: %v", round_number, lines_deleted)
			}
		})

		for line_number := 0; line_number < b.height; line_number++ {
			for column_number := 0; column_number < b.width; column_number++ {

				t.Run(fmt.Sprintf("(%d) Block[%d,%d] should NOT be busy", round_number, column_number, line_number), func(t *testing.T) {
					if b.areHere(column_number, line_number) {
						t.Fatalf("(%d) b[%d, %d] IS busy", round_number, column_number, line_number)
					}
				})

			}
		}

	} // round

}
