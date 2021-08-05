package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PlayingCard struct {
	Suit          string
	LongSuit      string
	UnicodeSuit   string
	Rank          string
	LongRank      string
	UnicodeSymbol string
	RankIndex     uint
}

type CribbageCard struct {
	PlayingCard
	Value uint
}

type PlayingCardCollection []PlayingCard

func (c PlayingCardCollection) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c PlayingCardCollection) Less(i int, j int) bool {
	return c[i].RankIndex < c[j].RankIndex
}

func IsValidPlayingCardString(cardString string) bool {
	matched, err := regexp.MatchString(`^(10|[0-9JQKAjqka])[HDSChdsc]$`, cardString)
	if err != nil {
		fmt.Println("Playing card validity check failed.")
		fmt.Println(err)
		return false
	}
	return matched
}

func GetPlayingCardFromString(cardString string) (*PlayingCard, error) {
	if !IsValidPlayingCardString(cardString) {
		return nil, fmt.Errorf("invalid card string: %s", cardString)
	}

	suitRegexp := regexp.MustCompile(`[HDSChdsc]`)
	rankRegexp := regexp.MustCompile(`(10|[0-9JQKAjqka])`)

	suit := strings.ToUpper(suitRegexp.FindString(cardString))
	rank := strings.ToUpper(rankRegexp.FindString(cardString))

	var longSuit, unicodeSuit, longRank string
	var suitMask, rankMask byte

	switch suit {
	case "S":
		longSuit = "Spades"
		unicodeSuit = "\u2660"
		suitMask = 0xa0
	case "H":
		longSuit = "Hearts"
		unicodeSuit = "\u2665"
		suitMask = 0xb0
	case "D":
		longSuit = "Diamonds"
		unicodeSuit = "\u2666"
		suitMask = 0xc0
	case "C":
		longSuit = "Clubs"
		unicodeSuit = "\u2663"
		suitMask = 0xd0
	}

	rankIndex, _ := strconv.ParseUint(rank, 10, 32)

	switch rank {
	case "J":
		longRank = "Jack"
		rankMask = 0x0b
	case "Q":
		longRank = "Queen"
		rankMask = 0x0c
	case "K":
		longRank = "King"
		rankMask = 0x0d
	case "A":
		longRank = "Ace"
		rankMask = 0x0e
	default:
		longRank = rank
		rankMask = byte(rankIndex)
	}

	return &PlayingCard{
		Suit:          suit,
		LongSuit:      longSuit,
		Rank:          rank,
		LongRank:      longRank,
		UnicodeSuit:   unicodeSuit,
		UnicodeSymbol: string([]byte{0x01, 0xf0, suitMask | rankMask}),
	}, nil
}

func GetCribbageCardFromString(cardString string) (*CribbageCard, error) {
	playingCard, err := GetPlayingCardFromString(cardString)
	if err != nil {
		return nil, err
	}

	cardVal, err := strconv.ParseUint(playingCard.Rank, 10, 32)

	if err != nil {
		switch playingCard.Rank {
		case "A":
			cardVal = 1
		default:
			cardVal = 10
		}
	}

	return &CribbageCard{
		PlayingCard: *playingCard,
		Value:       uint(cardVal),
	}, nil
}

func (c *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", c.LongRank, c.LongSuit)
}

func (c *PlayingCard) GetUnicodeString() string {
	return fmt.Sprintf("%s%s", c.Rank, c.UnicodeSuit)
}

func (c *PlayingCard) GetUnicodeSymbol() string {
	return c.UnicodeSymbol
}
