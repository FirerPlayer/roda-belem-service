package entity

type ReactionEnum string

const (
	Like   ReactionEnum = "like"
	Love   ReactionEnum = "love"
	Thanks ReactionEnum = "thanks"
)

type Reaction struct {
	React  ReactionEnum
	UserID string
}

func NewReaction(react ReactionEnum, userID string) *Reaction {
	return &Reaction{
		React:  react,
		UserID: userID,
	}
}
