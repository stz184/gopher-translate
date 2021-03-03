package translator

import "testing"

func TestTranslateWord(t *testing.T) {
	words := []struct {
		English string
		Gopher  string
	}{
		{
			"apple",
			"gapple",
		},
		{
			"Earth",
			"Gearth",
		},
		{
			"IRON",
			"GIRON",
		},
		{
			"Oath",
			"Goath",
		},
		{
			"universe",
			"guniverse",
		},
		{
			"xray",
			"gexray",
		},
		{
			"Xray",
			"Gexray",
		},
		{
			"XRAY",
			"GEXRAY",
		},
		{
			"Squirrel",
			"Irrelsquogo",
		},
		{
			"square",
			"aresquogo",
		},
		{
			"SQUASH",
			"ASHSQUOGO",
		},
		{
			"chair",
			"airchogo",
		},
		{
			"Tropical",
			"Opicaltrogo",
		},
		{
			"BANANA",
			"ANANABOGO",
		},
		{
			"smth",
			"smthogo",
		},
	}

	for _, word := range words {
		gopherWord := TranslateWord(word.English)

		if gopherWord != word.Gopher {
			t.Log("Expected", word.English, "to be translated as", word.Gopher, "but got", gopherWord)
			t.Fail()
		}
	}
}

func TestTranslateSentence(t *testing.T) {
	sentences := []struct {
		English string
		Gopher  string
	}{
		{
			"May I have an apple for breakfast?",
			"Aymogo GI avehogo gan gapple orfogo eakfastbrogo?",
		},
		{
			"I want to leave the Earth.",
			"GI antwogo otogo eavelogo ethogo Gearth.",
		},
		{
			"Gophers speak better English than me!",
			"Ophersgogo eakspogo etterbogo Genglish anthogo emogo!",
		},
		{
			"That's skipped",
			"That's ippedskogo",
		},
		{
			"This sentence has punctuation, exclamation mark and multiple   spaces!",
			"Isthogo entencesogo ashogo unctuationpogo, gexclamation arkmogo gand ultiplemogo acesspogo!",
		},
	}

	for _, sentence := range sentences {
		gopherSentence := TranslateSentence(sentence.English)
		if gopherSentence != sentence.Gopher {
			t.Log("Expected \"", sentence.English, "\" to be translated as \"", sentence.Gopher, "\" but got", gopherSentence)
			t.Fail()
		}
	}
}
