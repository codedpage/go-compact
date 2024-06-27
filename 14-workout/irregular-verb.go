package main

import (
	"fmt"
	"math/rand"
	"time"
)

var vb = map[string][]string{}

func main() {
	c := make(chan string)
	go verb(c)

	for val := range c {

		for _, kv := range vb[val] {
			fmt.Print(kv, "\t\t")
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
	}
}

func verb(c chan string) {

	keys := make([]string, 0, len(vb))
	for key := range vb {
		keys = append(keys, key)
	}

	indices := rand.Perm(len(keys))
	for _, idx := range indices {
		c <- keys[idx]
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func init() {
	vb = map[string][]string{

		"awake":      {"awake", "awoke", "awoken"},
		"be":         {"be", "was, were", "been"},
		"beat":       {"beat", "beat", "beaten"},
		"become":     {"become", "became", "become"},
		"begin":      {"begin", "began", "begun"},
		"bend":       {"bend", "bent", "bent"},
		"bet":        {"bet", "bet", "bet"},
		"bid":        {"bid", "bid", "bid"},
		"bite":       {"bite", "bit", "bitten"},
		"blow":       {"blow", "blew", "blown"},
		"break":      {"break", "broke", "broken"},
		"bring":      {"bring", "brought", "brought"},
		"broadcast":  {"broadcast", "broadcast", "broadcast"},
		"build":      {"build", "built", "built"},
		"burn":       {"burn", "burned / burnt", "burned / burnt"},
		"buy":        {"buy", "bought", "bought"},
		"catch":      {"catch", "caught", "caught"},
		"choose":     {"choose", "chose", "chosen"},
		"come":       {"come", "came", "come"},
		"cost":       {"cost", "cost", "cost"},
		"cut":        {"cut", "cut", "cut"},
		"dig":        {"dig", "dug", "dug"},
		"do":         {"do", "did", "done"},
		"draw":       {"draw", "drew", "drawn"},
		"dream":      {"dream", "dreamed / dreamt", "dreamed / dreamt"},
		"drive":      {"drive", "drove", "driven"},
		"drink":      {"drink", "drank", "drunk"},
		"eat":        {"eat", "ate", "eaten"},
		"fall":       {"fall", "fell", "fallen"},
		"feel":       {"feel", "felt", "felt"},
		"fight":      {"fight", "fought", "fought"},
		"find":       {"find", "found", "found"},
		"fly":        {"fly", "flew", "flown"},
		"forget":     {"forget", "forgot", "forgotten"},
		"forgive":    {"forgive", "forgave", "forgiven"},
		"freeze":     {"freeze", "froze", "frozen"},
		"get":        {"get", "got", "got"},
		"give":       {"give", "gave", "given"},
		"go":         {"go", "went", "gone"},
		"grow":       {"grow", "grew", "grown"},
		"hang":       {"hang", "hung", "hung"},
		"have":       {"have", "had", "had"},
		"hear":       {"hear", "heard", "heard"},
		"hide":       {"hide", "hid", "hidden"},
		"hit":        {"hit", "hit", "hit"},
		"hold":       {"hold", "held", "held"},
		"hurt":       {"hurt", "hurt", "hurt"},
		"keep":       {"keep", "kept", "kept"},
		"know":       {"know", "knew", "known"},
		"lay":        {"lay", "laid", "laid"},
		"lead":       {"lead", "led", "led"},
		"learn":      {"learn", "learned / learnt", "learned / learnt"},
		"leave":      {"leave", "left", "left"},
		"lend":       {"lend", "lent", "lent"},
		"let":        {"let", "let", "let"},
		"lie":        {"lie", "lay", "lain"},
		"lose":       {"lose", "lost", "lost"},
		"make":       {"make", "made", "made"},
		"mean":       {"mean", "meant", "meant"},
		"meet":       {"meet", "met", "met"},
		"pay":        {"pay", "paid", "paid"},
		"put":        {"put", "put", "put"},
		"read":       {"read", "read", "read"},
		"ride":       {"ride", "rode", "ridden"},
		"ring":       {"ring", "rang", "rung"},
		"rise":       {"rise", "rose", "risen"},
		"run":        {"run", "ran", "run"},
		"say":        {"say", "said", "said"},
		"see":        {"see", "saw", "seen"},
		"sell":       {"sell", "sold", "sold"},
		"send":       {"send", "sent", "sent"},
		"show":       {"show", "showed", "showed / shown"},
		"shut":       {"shut", "shut", "shut"},
		"sing":       {"sing", "sang", "sung"},
		"sink":       {"sink", "sank", "sunk"},
		"sit":        {"sit", "sat", "sat"},
		"sleep":      {"sleep", "slept", "slept"},
		"speak":      {"speak", "spoke", "spoken"},
		"spend":      {"spend", "spent", "spent"},
		"stand":      {"stand", "stood", "stood"},
		"stink":      {"stink", "stank", "stunk"},
		"swim":       {"swim", "swam", "swum"},
		"take":       {"take", "took", "taken"},
		"teach":      {"teach", "taught", "taught"},
		"tear":       {"tear", "tore", "torn"},
		"tell":       {"tell", "told", "told"},
		"think":      {"think", "thought", "thought"},
		"throw":      {"throw", "threw", "thrown"},
		"understand": {"understand", "understood", "understood"},
		"wake":       {"wake", "woke", "woken"},
		"wear":       {"wear", "wore", "worn"},
		"win":        {"win", "won", "won"},
		"write":      {"write", "wrote", "written"},
	}
}
