package main

import (
	"./binding"
	"io/ioutil"
	"os"
)

var stemMap map[string]*binding.Gohunspell
var suggestMap map[string]*binding.Gohunspell

func init() {
	stemMap = make(map[string]*binding.Gohunspell)
	suggestMap = make(map[string]*binding.Gohunspell)

	initHunspellLang(stemMap, "en", "en_US.aff", "en_US.dic")
	initHunspellLang(stemMap, "fr", "fr_FR.aff", "fr_FR.dic")
	initHunspellLang(stemMap, "de", "de_DE.aff", "de_DE.dic")
	initHunspellLang(stemMap, "es", "es_ES.aff", "es_ES.dic")

	initHunspellLang(suggestMap, "en", "en_US.aff", "en_US.dic")
	initHunspellLang(suggestMap, "fr", "fr-moderne.aff", "fr-moderne.dic")
	initHunspellLang(suggestMap, "de", "de_DE_frami.aff", "de_DE_frami.dic")
	initHunspellLang(suggestMap, "es", "es_ANY.aff", "es_ANY.dic")
}

func initHunspellLang(cacheMap map[string]*binding.Gohunspell, lang string, affFile string, dicFile string)  {
	file, err := os.Open("./binding/include/dictionaries/" + affFile)
	if err != nil {
		return
	}
	aff, e := ioutil.ReadAll(file)
	if e != nil {
		return
	}
	e = file.Close()
	if e != nil {
		return
	}
	file, e = os.Open("./binding/include/dictionaries/" + dicFile)
	if e != nil {
		return
	}
	dic, e := ioutil.ReadAll(file)
	e = file.Close()
	if e != nil {
		return
	}
	cacheMap[lang] = binding.NewGohunspell(aff, dic)
}

func StemWord(word string, lang string)[]string {
	speller := stemMap[lang]
	_, list := speller.Stem(word)
	return list
}

func Suggest(word string, lang string)[]string {
	speller := suggestMap[lang]
	_, _, list := speller.CheckSuggestions(word)
	return list
}