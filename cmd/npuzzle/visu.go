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

func DisplayVisu(info *Info) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

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

	// other setup
	base_head, base_tail := fillLST(info.End)

	state := base_tail

	// ttf setup
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("./assets/Bebas-Regular.ttf", 64)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	renderer.Clear()

	textures := make([]*sdl.Texture, state.State.Size*state.State.Size)
	for ii := 0; ii < state.State.Size*state.State.Size; ii += 1 {
		fontSurf, err := font.RenderUTF8Blended(strconv.Itoa(ii), sdl.Color{203, 247, 237, 180})
		if err != nil {
			panic(err) // TODO: address error
		}
		texture, err := renderer.CreateTextureFromSurface(fontSurf)
		if err != nil {
			panic(err) // TODO: address error
		}
		textures[ii] = texture
		fontSurf.Free()
	}

	defer func() {
		for ii := range textures {
			textures[ii].Destroy()
		}
	}()

	// render loop
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_q:
						fallthrough
					case sdl.K_ESCAPE:
						running = false
					case sdl.K_e:
						state = base_head
					case sdl.K_s:
						state = base_tail
					case sdl.K_RIGHT:
						if state.Prev != nil {
							state = state.Prev
						}
					case sdl.K_LEFT:
						if state.Next != nil {
							state = state.Next
						}
					}
				}
			}
		}
		renderer.Clear()
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.FillRect(nil)
		drawState(renderer, font, textures, state.State)
		renderer.Present()
	}
}

func drawState(renderer *sdl.Renderer, font *ttf.Font, textures []*sdl.Texture, state *State) {
	var tilesize, x, y int32
	tilesize = int32(wWidth/state.Size - (tilebuf - tilebuf/state.Size))
	for ii, n := range state.Board {
		x = int32(GetY(ii, state.Size))
		y = int32(GetX(ii, state.Size))
		rect := &sdl.Rect{x*tilesize + x*tilebuf, y*tilesize + y*tilebuf, tilesize, tilesize}
		if n != 0 {
			renderer.SetDrawColor(64, 110, 142, 255)
			renderer.FillRect(rect)
			// fix weird sizing issues with 0-9
			if n < 10 {
				rect.W /= 2
				rect.X += rect.W / 2
			}
			renderer.Copy(textures[n], &sdl.Rect{0, 0, 100, 100}, rect)
			// fontSurf.Free()
			// texture.Destroy()
		} else {
			renderer.SetDrawColor(22, 25, 37, 255)
			renderer.FillRect(rect)
		}
	}
}
