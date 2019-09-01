package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	wWidth  = 750
	wHeight = wWidth
	tilebuf = 5
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
	defer renderer.Destroy()

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
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.FillRect(nil)
		drawState(renderer, font, state.State)
		renderer.Present()
		// window.UpdateSurface()
	}
}

func drawState(renderer *sdl.Renderer, font *ttf.Font, state *State) {
	var tilesize, x, y int32
	fmt.Println(state)
	tilesize = int32(wWidth/state.Size - (tilebuf - tilebuf/state.Size))
	for ii, n := range state.Board {
		x = int32(GetY(ii, state.Size))
		y = int32(GetX(ii, state.Size))
		// rect := sdl.Rect{x*tilesize + x*tilebuf, y*tilesize + y*tilebuf, tilesize, tilesize}
		// rect := new(sdl.Rect)
		rect := &sdl.Rect{x*tilesize + x*tilebuf, y*tilesize + y*tilebuf, tilesize, tilesize}
		if n != 0 {
			renderer.SetDrawColor(0, 255, 255, 255)
			renderer.FillRect(rect)

			fontSurf, err := font.RenderUTF8Solid(strconv.Itoa(n), sdl.Color{255, 100, 200, 255})
			if err != nil {
				panic(err) // TODO: address error
			}
			texture, err := renderer.CreateTextureFromSurface(fontSurf)
			if err != nil {
				panic(err) // TODO: address error
			}

			renderer.Copy(texture, &sdl.Rect{0, 0, 100, 100}, rect)

			fontSurf.Free()
			texture.Destroy()
		} else {
			renderer.SetDrawColor(255, 0, 0, 255)
			renderer.FillRect(rect)
		}
	}
}
