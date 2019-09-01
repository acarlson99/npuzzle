package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	wWidth  = 750
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
	// graphics setup
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(wWidth, wHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// surface, err := window.GetSurface()
	// defer surface.Free()
	// if err != nil {
	// 	panic(err)
	// }

	// ttf setup
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	font, err := ttf.OpenFont("./assets/ComicSans.ttf", 24)
	if err != nil {
		panic(err)
	}

	fontSurf, err := font.RenderUTF8Solid("HELLO THERE", sdl.Color{255, 100, 200, 255})
	if err != nil {
		panic(err) // TODO: address error
	}
	texture, err := renderer.CreateTextureFromSurface(fontSurf)
	if err != nil {
		panic(err) // TODO: address error
	}

	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.Clear()

	// render loop
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
				}
			}
		}
		renderer.Clear()
		// surface.FillRect(nil, 0)
		renderer.SetDrawColor(255, 0, 0, 255)
		drawState(renderer, state.State)
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Copy(texture, &sdl.Rect{0, 0, 100, 100}, &sdl.Rect{100, 100, 100, 100})
		renderer.Present()
		// window.UpdateSurface()
	}
}

func drawState(renderer *sdl.Renderer, state *State) {
	var tilesize, x, y int32
	fmt.Println(state)
	tilesize = int32(wWidth/state.Size - (tilebuf - tilebuf/state.Size))
	for ii, n := range state.Board {
		x = int32(GetX(ii, state.Size))
		y = int32(GetY(ii, state.Size))
		rect := sdl.Rect{x*tilesize + x*tilebuf, y*tilesize + y*tilebuf, tilesize, tilesize}
		if n != 0 {
			renderer.FillRect(&rect)
		} else {
			renderer.FillRect(&rect)
		}
	}
}
