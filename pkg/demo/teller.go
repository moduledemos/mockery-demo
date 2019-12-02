package demo

type Teller struct {
}

func (teller Teller) say(words []byte) {
	return
}

func (teller Teller) Tell(story Story) error {
	for i := 0; story.HasChapter(i); i++ {
		words, err := story.ReadChapter(i)
		if err != nil {
			return err
		}
		teller.say(words)
	}
	return nil
}
