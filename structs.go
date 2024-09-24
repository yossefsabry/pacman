package main

type Sprite struct {
	row int
	col int
}

type Config struct {
    Player   string `json:"player"`
    Ghost    string `json:"ghost"`
    Wall     string `json:"wall"`
    UnderWall     string `json:"underWall"`
    Dot      string `json:"dot"`
    Pill     string `json:"pill"`
    Death    string `json:"death"`
    Space    string `json:"space"`
    UseEmoji bool   `json:"use_emoji"`
}


