package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GamesType struct {
	Games map[string]GameState
}

type GameState struct {
	Field         [9]string `json:"field"`
	CurrentPlayer string    `json:"currentPlayer"`
	Message       string    `json:"message"`
}

//var gameState = GameState{[9]string{"", "", "", "", "", "", "", "", ""}, "X", "Start the Game!"}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users/{game}", func(w http.ResponseWriter, r *http.Request) {
		/*

		   GET http://localhost:8080/users/Tictactoe HTTP/1.1

		*/
		vars := mux.Vars(r)
		if gsname, ok := vars["game"]; ok {

			Games := GetGames("./games.json")

			if Games.Games == nil {
				Games = &GamesType{map[string]GameState{}}
			}
			gs, ok := Games.Games[gsname]
			if !ok {
				//if ok {
				gs = GameState{[9]string{}, "X", "Start the Game!"}

				Games.Games[gsname] = gs
				WriteGames("./games.json", Games)

			}
			if currentAsJSON, err := json.Marshal(gs); err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusOK)
				w.Write(currentAsJSON)
				//fmt.Fprintf(w, "%v", string(currentAsJSON))

			} else {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error while creating JSON : %v", err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Please provide Game in URL")
		}

	}).Methods("GET")

	r.HandleFunc("/users/{game}/move/{field:[1-9]}", func(w http.ResponseWriter, r *http.Request) {
		/*

		   POST http://localhost:8080/users/Tictactoe/move/3 HTTP/1.1
		   POST http://localhost:8080/users/Tictactoe/move/4 HTTP/1.1
		   POST http://localhost:8080/users/Tictactoe/move/2 HTTP/1.1

		*/

		vars := mux.Vars(r)
		if gsname, ok := vars["games"]; ok {
			Games := GetGames("./games.json")

			if Games.Games == nil {
				Games = &GamesType{map[string]GameState{}}
			}
			gs, ok := Games.Games[gsname]
			if !ok {
				gs = GameState{[9]string{}, "X", "Start the Game!"}
				Games.Games[gsname] = gs
				WriteGames("./games.json", Games)

			}
			if gs.gameOngoing() {

				if fieldstring, ok := vars["field"]; ok {
					if field, err := strconv.Atoi(fieldstring); err == nil {
						if gs.CurrentPlayer == "X" {
							gs.CurrentPlayer = "O"
						} else {
							gs.CurrentPlayer = "X"
						}
						gs.playerPut(field - 1)

					}
				}

			}

			gs.gameOngoing()

			if currentAsJSON, err := json.Marshal(gs); err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusOK)
				w.Write(currentAsJSON)
				//fmt.Fprintf(w, "%v", string(currentAsJSON))

			} else {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error while creating JSON : %v", err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Please provide Game in URL")
		}

	}).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func (gs *GameState) gameOngoing() bool {

	if !gs.playerWon() {
		return !gs.draw()
	}
	return false
}

func (gs *GameState) playerPut(field int) {

	gs.Field[field] = gs.CurrentPlayer

}

func compare(tocompare ...string) bool {
	res := true
	current := ""
	if len(tocompare) > 0 {
		current = tocompare[0]
	} else {
		return false
	}
	for _, v := range tocompare {
		if len(v) == 0 {
			res = false
			break
		}

		if v != current {
			res = false
			break
		}

	}
	return res

}

func (gs *GameState) playerWon() bool {

	res := false

	for i := 0; i < 3; i++ {

		res = compare(gs.Field[i*3], gs.Field[i*3+1], gs.Field[i*3+2])

		res = compare(gs.Field[i], gs.Field[i+3], gs.Field[i+6])

	}

	res = compare(gs.Field[0], gs.Field[4], gs.Field[8])

	res = compare(gs.Field[2], gs.Field[4], gs.Field[6])

	if res {
		gs.Message = fmt.Sprintf("Player %s has won !", gs.CurrentPlayer)
	}

	return res

}

func (gs *GameState) draw() bool {
	i := 0
	for _, v := range gs.Field {
		if v == "X" || v == "Y" {
			i++
		}
	}
	if i == 8 {

		gs.Message = fmt.Sprint("Draw !")

		return true
	}
	return false

}

func GetGames(path string) *GamesType {

	games := &GamesType{}
	file, e := ioutil.ReadFile(path)
	if e != nil {
		panic(fmt.Sprint("no games.json :", e))
	}
	json.Unmarshal(file, games)

	return games
}

func WriteGames(path string, Games *GamesType) {

	d1, _ := json.Marshal(Games)
	ioutil.WriteFile(path, d1, 0644)

}
