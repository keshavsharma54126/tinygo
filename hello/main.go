package main

import (
	"flag"
	"fmt"
)

func main(){
	var lang string 
	flag.StringVar(&lang,"lang","en","The required language, eg. en,ur,hi...")
	flag.Parse()
	greeting:=greet(language(lang))
	phrasebookgreeting:=phrasebookgreet(language(lang))
	fmt.Println(greeting)
	fmt.Println(phrasebookgreeting)
}

type language string 

var phrasebook= map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
    "en": "Hello world",       // English
    "fr": "Bonjour le monde",  // French
    "he": "שלום עולם",         // Hebrew
    "ur": "ہیلو دنیا",         // Urdu
    "vi": "Xin chào Thế Giới", // Vietnamese
	"hi":  "नमस्ते दुनिया",   //hindi my mother toungue
}

func phrasebookgreet(l language)string{

	greeting,ok:=phrasebook[l]
	if !ok{
		return fmt.Sprintf("unsupported language:%q",l)
	}

	return greeting
}

func greet(l language) string{

	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
} 