package tests

import (
	"testing"

	"github.com/stscoundrel/old-norwegian-dictionary-go/dictionary"
)

func TestDictionaryContainsCorrectAmountOfEntries(t *testing.T) {
	result := dictionary.GetDictionary()
	expected := 42021

	if len(result) != expected {
		t.Error("Did not return expected amount of entries. Received", len(result), "expected ", expected)
	}
}

func TestDictionaryContainsExpectedContent(t *testing.T) {
	result := dictionary.GetDictionary()

	expected1 := dictionary.DictionaryEntry{
		Headword:     "-æri",
		PartOfSpeech: "uten ordklasse",
		Definition:   "-æri (af ár dvs. Aar) i hallæri.",
	}

	expected2 := dictionary.DictionaryEntry{
		Headword:     "fri",
		PartOfSpeech: "m",
		Definition:   "fri, m. = friðill. Hým. 9.",
	}

	expected3 := dictionary.DictionaryEntry{
		Headword:     "náðuliga",
		PartOfSpeech: "adv",
		Definition:   "náðuliga, adv.  1)  i Stilhed, ubemærket; hann bauð at hafa Hánef þar á launþar til, er skip kemr n. at, svá athonum mætti útan koma Vem. 591; B.biskup biðr nú því öruggari til guðsaf öllu hjarta, sem hann má þat náð-uligar ok leyniligar gera fyrir mönn-um Mar. 116911 fg; biðjandi því meðmeira athuga, sem hann mátti leyni-ligar ok auðveldligar (&vl náðuligar)Mar. 83710. 34.  2)  naadigen; biðjom vér,at þér takir þessum várum erendumbetr ok náðuligar, en vér erum verðirDN. VII, 19013.",
	}

	expected4 := dictionary.DictionaryEntry{
		Headword:     "þyrnir",
		PartOfSpeech: "m",
		Definition:   "þyrnir, m. Tjørn, Tornebusk. Stj. 39611;Hom. 10218; Post. 75034; Klm. 54615;Mar. 3378. 10351.",
	}

	if result[0] != expected1 {
		t.Error("Did not return expected entry. Received", result[0], "expected ", expected1)
	}

	if result[10000] != expected2 {
		t.Error("Did not return expected entry. Received", result[10000], "expected ", expected2)
	}

	if result[25000] != expected3 {
		t.Error("Did not return expected entry. Received", result[25000], "expected ", expected3)
	}

	if result[42000] != expected4 {
		t.Error("Did not return expected entry. Received", result[42000], "expected ", expected4)
	}
}
