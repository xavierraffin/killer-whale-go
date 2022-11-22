package main

import (
	"log"
	"math/rand"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Xavier Raffin",
		Color:      "#000000",
		Head:       "orca",
		Tail:       "shiny",
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
	// printObj(state)

	GAME_MODE = ReadMode(state.Game.Ruleset.Name)
	log.Printf("Game mode is %s", GAME_MODE)
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

/*
func killOrAvoid(myHead Coord, bodyPart Coord, isMoveSafe map[string]bool, enemyLength int, myLength int) {
  if()
}
*/

func avoidCollision(myHead Coord, bodyPart Coord, isMoveSafe map[string]bool) {
  if(myHead.X == bodyPart.X){
    if(myHead.Y + 1 == bodyPart.Y) {
      isMoveSafe["up"] = false
    } else if(myHead.Y - 1 == bodyPart.Y) {
      isMoveSafe["down"] = false
    }
  } else if(myHead.Y == bodyPart.Y){
    if(myHead.X + 1 == bodyPart.X) {
      isMoveSafe["right"] = false
    } else if(myHead.X - 1 == bodyPart.X) {
      isMoveSafe["left"] = false
    }
  }
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {

	GAME_MODE = ReadMode(state.Game.Ruleset.Name)
	log.Printf("MOVE mode is %s", GAME_MODE)

	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		isMoveSafe["left"] = false

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		isMoveSafe["right"] = false

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		isMoveSafe["down"] = false

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		isMoveSafe["up"] = false
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	if GAME_MODE != Wrapped {
    if myHead.X == 0 { 
		  isMoveSafe["left"] = false
  
  	} else if myHead.X == state.Board.Width -1 { 
  		isMoveSafe["right"] = false
  	}
    if myHead.Y == 0 { 
  		isMoveSafe["down"] = false
  
  	} else if myHead.Y == state.Board.Height -1 {
  		isMoveSafe["up"] = false
  
  	}
	}

  if GAME_MODE != Wrapped {
	  // Prevent Killer Whale from colliding with itself or other Battlesnakes
    for _, snake := range state.Board.Snakes {
      if(snake.ID != state.You.ID) {
        // Next turn the head will be replaced by neck
        avoidCollision(myHead, snake.Body[0], isMoveSafe)
  	    // killOrAvoid(myHead, snake.Body[0], isMoveSafe, snake.Length, state.You.Length)
      }
      for j := 1; j < len(snake.Body) - 1; j++ { // We ignore the queue, it will move
  	    avoidCollision(myHead, snake.Body[j], isMoveSafe)
      }
    }
  }

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := safeMoves[rand.Intn(len(safeMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
