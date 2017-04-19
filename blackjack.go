package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"log"
	"net/http"
	"flag"
)

var (
	clubs = []string {"cking","cqueen","cjack","cace","c1","c2","c3","c4","c5","c6","c7","c8","c9"}
	hearts = []string{"hking","hqueen","hjack","hace","h1","h2","h3","h4","h5","h6","h7","h8","h9"}
	spades = []string{"sking","squeen","sjack","sace","s1","s2","s3","s4","s5","s6","s7","s8","s9"}
	diamonds = []string{"dking","dqueen","djack","dace","d1","d2","d3","d4","d5","d6","d7","d8","d9"}	
)

/* struct Dealer Player */

type Player struct {
	dealer bool 
	playerType string
	playerCards []string
}


var deckSize int = len(clubs) + len(hearts) + len(spades) + len(diamonds)
var handSize int = 10
var addr = flag.String("addr", ":8500", "http service address")

func main() {
	
	    http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
	
}

func QR(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("<HTML><B>Hello, World!</HTML>")
	fmt.Fprintf(w, "<HTML><B>Hello, World! %s</HTML>", req.URL.Path[1:])
	cards := make([]string, deckSize)
	hand := make([]string, handSize)
	dealerPull := make([]int, deckSize)
	
	cards, dealerPull = shuffle(cards)
	dealer := &Player{ true, "Dealer", cards }
	player := &Player{ false, "Player 1", hand }

	fmt.Fprintf(w,"<P>Would you like to play a game?<P>")		
	fmt.Fprintf(w,"<P>Dealer's Cards %s", dealer)
	fmt.Fprintf(w,"<P>Player's Cards %s", player)
	fmt.Fprintf(w,"<P>Dealer Pull Order %s", dealerPull)
}


func shuffle(deck []string) ([]string, []int) {
	for _, clubCard := range clubs {
		deck = append(deck,clubCard)
	}
	for _, heartCard := range hearts {
		deck = append(deck,heartCard)
	}
	for _, spadeCard := range spades {
		deck = append(deck,spadeCard)
	}
	for _, diamondCard := range diamonds {
		deck = append(deck,diamondCard)
	}
	//Fisher-Yates shuffle
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	//lets try Perm
	newlist := make([]int, deckSize) 
	newlist = rand.Perm(52)
	return deck, newlist
}

func cardValue (card string) int {
	var value int
		
	switch (card) {
	case "ace", "king", "queen", "jack":
		value = 10
	default:
		value, _ = strconv.Atoi(card)
	}
 
	return value
}
