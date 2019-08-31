package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	wWidth  = 1023
	wHeight = wWidth
	tilebuf = 18
)

type DLLST struct {
	State *State
	Next  *DLLST
	Prev  *DLLST
}

func fillLST(state *State) (*DLLST, *DLLST) {
	lst := new(DLLST)
	lst.State = state
	var head, tail *DLLST
	if state.Parent != nil {
		head, tail = fillLST(state.Parent)
		lst.Next = head
		lst.Next.Prev = lst
	} else {
		head, tail = lst, lst
	}
	return lst, tail
}

func DisplayGui(info *Info) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		wWidth, wHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	defer surface.Free()
	if err != nil {
		panic(err)
	}

	// state := info.End
	_, base_state := fillLST(info.End)

	state := base_state

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_r:
					if t.Type == sdl.KEYDOWN {
						state = base_state
					}
				case sdl.K_RIGHT:
					if t.Type == sdl.KEYDOWN {
						if state.Prev != nil {
							state = state.Prev
						}
					}
				case sdl.K_LEFT:
					if t.Type == sdl.KEYDOWN {
						if state.Next != nil {
							state = state.Next
						}
					}
				default:
					fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
						t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				}
			}
		}
		surface.FillRect(nil, 0)
		drawState(surface, state.State)
		window.UpdateSurface()
	}
}

func drawState(surface *sdl.Surface, state *State) {
	var tilesize, x, y int32
	fmt.Println(state)
	tilesize = int32(wWidth/state.Size - (tilebuf - tilebuf/state.Size))
	for ii, n := range state.Board {
		x = int32(GetX(ii, state.Size))
		y = int32(GetY(ii, state.Size))
		rect := sdl.Rect{x*tilesize + x*tilebuf, y*tilesize + y*tilebuf, tilesize, tilesize}
		if n != 0 {
			surface.FillRect(&rect, 0xffff0000)
		} else {
			surface.FillRect(&rect, 0xff00ff00)
		}
	}
}
