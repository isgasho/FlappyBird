package main

import (
    //_ "image/png"

    "github.com/hajimehoshi/ebiten" //"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
    pipes [PIPE_NUM]*Pipe
    bird *Bird
    first_pipe_idx int
    last_pipe_idx int
}

func new_game() *Game {
    game := Game{}
    for a := 0; a < PIPE_NUM; a++ {
        game.pipes[a] = new_pipe()
    }
    game.bird = new_bird()
    game.first_pipe_idx = 0
    game.last_pipe_idx = PIPE_NUM - 1
    game.release_new_pipe()
    return &game
}

func (game *Game) Update(screen *ebiten.Image) error {
    game.bird.flap()
    for _, pipe := range game.pipes {
        pipe.move()
    }
    if game.pipes[game.last_pipe_idx].longitude <= SCREEN_WIDTH - PIPE_WIDTH - DISTANCE {
        game.release_new_pipe()
    }
    if game.first_pipe().longitude < -PIPE_WIDTH {
        game.reset_first_pipe()
    }
    if game.bird.touch_pipe(game.first_pipe()) {
        game.bird.die()
    }
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

	op_bg := &ebiten.DrawImageOptions{}
	op_base := &ebiten.DrawImageOptions{}
    op_base.GeoM.Translate(0, BG_HEIGHT)
    screen.DrawImage(background, op_bg)
    screen.DrawImage(base, op_base)
    for i := 1; i < BG_NUM; i++ {
        op_bg.GeoM.Translate(BG_WIDTH, 0)
        op_base.GeoM.Translate(BG_WIDTH, 0)
        screen.DrawImage(background, op_bg)
        screen.DrawImage(base, op_base)
    }
    game.bird.draw(screen)
    for _, pipe := range game.pipes {
        pipe.draw(screen)
    }
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}


func (game *Game) release_new_pipe() {
    game.last_pipe_idx += 1
    game.last_pipe_idx %= PIPE_NUM
    game.last_pipe().visible = true
}

func (game *Game) reset_first_pipe() {
    game.first_pipe().reset()
    game.first_pipe_idx += 1
    game.first_pipe_idx %= PIPE_NUM
}

func (game *Game) first_pipe() *Pipe {
    return game.pipes[game.first_pipe_idx]
}

func (game *Game) last_pipe() *Pipe {
    return game.pipes[game.last_pipe_idx]
}
