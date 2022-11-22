package main

import (
	"log"
	"math/rand"
)

  var DEATH_PENALTY = -100
  var KILL_BONUS = 1

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


func killOrAvoid(myHead Coord, otherHead Coord, isMoveBeneficial map[string]int, enemyLength int, myLength int, width int, height int) {
  var colisionScore = KILL_BONUS
  if(enemyLength >= myLength) {
    colisionScore = DEATH_PENALTY
  }
  if GAME_MODE != Wrapped {
    if(myHead.X == otherHead.X){
      if(myHead.Y + 2 == otherHead.Y) {
        isMoveBeneficial["up"] += colisionScore
      } else if(myHead.Y - 2 == otherHead.Y) {
        isMoveBeneficial["down"] += colisionScore
      }
    } else if(myHead.Y == otherHead.Y){
      if(myHead.X + 2 == otherHead.X) {
        isMoveBeneficial["right"] += colisionScore
      } else if(myHead.X - 2 == otherHead.X) {
        isMoveBeneficial["left"] += colisionScore
      }
    }
    if(myHead.X == otherHead.X + 1){
      if(myHead.Y + 1 == otherHead.Y) {
        isMoveBeneficial["up"] += colisionScore
        isMoveBeneficial["left"] += colisionScore
      } else if(myHead.Y - 1 == otherHead.Y) {
        isMoveBeneficial["down"] += colisionScore
        isMoveBeneficial["left"] += colisionScore
      }
    } else if(myHead.X == otherHead.X - 1){
      if(myHead.Y + 1 == otherHead.Y) {
        isMoveBeneficial["up"] += colisionScore
        isMoveBeneficial["right"] += colisionScore
      } else if(myHead.Y - 1 == otherHead.Y) {
        isMoveBeneficial["down"] += colisionScore
        isMoveBeneficial["right"] += colisionScore
      }
    }
  } else {
    if(myHead.X == otherHead.X){
      if((myHead.Y + 2)%height == otherHead.Y) {
        isMoveBeneficial["up"] += colisionScore
      } else if((myHead.Y - 2)%height == otherHead.Y) {
        isMoveBeneficial["down"] += colisionScore
      }
    } else if(myHead.Y == otherHead.Y){
      if((myHead.X + 2)%width == otherHead.X) {
        isMoveBeneficial["right"] += colisionScore
      } else if(myHead.X == (otherHead.X + 2)%width) {
        isMoveBeneficial["left"] += colisionScore
      }
    }
    if(myHead.X == (otherHead.X + 1)%width){
      if((myHead.Y + 1)%height == otherHead.Y) {
        isMoveBeneficial["up"] += colisionScore
        isMoveBeneficial["left"] += colisionScore
      } else if(myHead.Y == (otherHead.Y + 1)%height) {
        isMoveBeneficial["down"] += colisionScore
        isMoveBeneficial["left"] += colisionScore
      }
    } else if((myHead.X +1)%width == otherHead.X){
      if((myHead.Y + 1)%height  == otherHead.Y) {
        isMoveBeneficial["up"] += colisionScore
        isMoveBeneficial["right"] += colisionScore
      } else if(myHead.Y == (otherHead.Y + 1)%height) {
        isMoveBeneficial["down"] += colisionScore
        isMoveBeneficial["right"] += colisionScore
      }
    }
  }
}


func avoidCollision(myHead Coord, bodyPart Coord, isMoveSafe map[string]bool, width int, height int) {
  if GAME_MODE != Wrapped {
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
  } else {
    if(myHead.X == bodyPart.X){
      if((myHead.Y + 1)%height == bodyPart.Y) {
        isMoveSafe["up"] = false
      } else if(myHead.Y == (bodyPart.Y +1)%height) {
        isMoveSafe["down"] = false
      }
    } else if(myHead.Y == bodyPart.Y){
      if((myHead.X + 1)%width == bodyPart.X) {
        isMoveSafe["right"] = false
      } else if(myHead.X == (bodyPart.X +1)%width) {
        isMoveSafe["left"] = false
      }
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
  isMoveBeneficial := map[string]int{
		"up":    0,
		"down":  0,
		"left":  0,
		"right": 0,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head

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

  
  // Prevent Killer Whale from colliding with itself or other Battlesnakes
  for _, snake := range state.Board.Snakes {
    if(snake.ID != state.You.ID) {
      // Next turn the head will be replaced by neck
      avoidCollision(myHead, snake.Body[0], isMoveSafe, state.Board.Width, state.Board.Height)
      killOrAvoid(myHead, snake.Body[0], isMoveBeneficial, snake.Length, state.You.Length, state.Board.Width, state.Board.Height)
    }
    for j := 1; j < len(snake.Body) - 1; j++ { // We ignore the queue, it will move
      avoidCollision(myHead, snake.Body[j], isMoveSafe, state.Board.Width, state.Board.Height)
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

  // Remove safeMoves with a lower benefit
  var maxMoveBenefits int = isMoveBeneficial[safeMoves[0]];
  var currentBenefits int
  for j := 1; j < len(safeMoves); j++ {
    currentBenefits = isMoveBeneficial[safeMoves[j]];
    if(currentBenefits > maxMoveBenefits) {
      maxMoveBenefits = currentBenefits;
    }
  }

  bestMoves := []string{}
  for j := 0; j < len(safeMoves); j++ {
    currentBenefits = isMoveBeneficial[safeMoves[j]];
    if(currentBenefits == maxMoveBenefits) {
      bestMoves = append(bestMoves, safeMoves[j])
    }
  }
  
	// Choose a random move from the safe ones
	nextMove := bestMoves[rand.Intn(len(bestMoves))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food
  printObj(bestMoves)
  printObj(isMoveBeneficial)
  printObj(isMoveSafe)

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
