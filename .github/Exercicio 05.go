package main

import "fmt"

type Playlist struct {
	Nome    string
	Musicas []Musica
}

type Musica struct {
	Titulo  string
	Artista string
	Duracao int
}

func main() {
	p := []Playlist{
		{
			Nome: "Playlista Vrau",
			Musicas: []Musica{
				{
					Titulo:  "Vrau 1",
					Artista: "Vrauson",
					Duracao: 100,
				},
				{
					Titulo:  "Vrau 2",
					Artista: "Vrauson",
					Duracao: 150,
				},
				{
					Titulo:  "Vrau 3",
					Artista: "Vrausonelson",
					Duracao: 200,
				},
			},
		},
		{
			Nome: "Playlista Vrau 2",
			Musicas: []Musica{
				{
					Titulo:  "Vrau 1",
					Artista: "Vrauson",
					Duracao: 100,
				},
				{
					Titulo:  "Vrau 2",
					Artista: "Vrauson",
					Duracao: 150,
				},
			},
		},
	}

	fmt.Print(FindPlaylists(p, "Vrau 4"))
}

func FindPlaylists(p []Playlist, nomeMusica string) []Playlist {
	var found []Playlist

	for _, playlist := range p {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == nomeMusica {
				found = append(found, playlist)
			}
		}
	}

	return found
}
