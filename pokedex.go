package Pokedex

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"github.com/gorilla/mux"
	"strings"
	"strconv"
	"sort"
)

type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective types, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak types that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
}

// Move is an attack information. The
type Move struct {
	// The ID of the move
	ID int `json:"id"`
	// Name of the attack
	Name string `json:"name"`
	// Type of attack
	Type string `json:"type"`
	// The damage that enemy will take
	Damage int `json:"damage"`
	// Energy requirement of the attack
	Energy int `json:"energy"`
	// Dps is Damage Per Second
	Dps float64 `json:"dps"`
	// The duration
	Duration int `json:"duration"`
}

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}

func (base *BaseData) listHandler(w http.ResponseWriter, r *http.Request) {
	/*
	* List http request handler method
	*/
	params := mux.Vars(r) //Request's parameters!
	q1 := params["q1"]
	q2 := params["q2"]
	if len(q1) == 0 { //If there is no parameter, so print all pokemons!
		fmt.Fprint(w, "All the pokemons are:", "\n")
		for i := 0; i < len(base.Pokemons); i++ {
			fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
		}
	} else if (len(q2) == 0 && strings.Contains(q1, "=")) { // If there is only one parameters!
		s := strings.Split(q1, "=")
		key, val := s[0], s[1]
		//fmt.Println(len(s))
		if (strings.EqualFold(key, "type")) { //If this parameter is type, print pokemons of that type.
			fmt.Fprint(w,val, " types of pokemons are:", "\n")
			for i := 0; i < len(base.Pokemons); i++ {
				if (len(base.Pokemons[i].TypeII) != 0 && len(base.Pokemons[i].TypeI) != 0 && (strings.EqualFold(base.Pokemons[i].TypeI[0], val) || strings.EqualFold(base.Pokemons[i].TypeII[0], val))) {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}

			}
		} else if (strings.EqualFold(key, "sortby")) { // If this parameter is sortby, first sort accordingly, then print all pokemons.
			if (strings.EqualFold(val, "BaseAttack")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseAttack < base.Pokemons[j].BaseAttack
				})
				fmt.Fprint(w, "All the pokemons are(sorted by base attacks):", "\n")
				for i := 0; i < len(base.Pokemons); i++ {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			} else if (strings.EqualFold(val, "BaseDefense")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseDefense < base.Pokemons[j].BaseDefense
				})
				fmt.Fprint(w, "All the pokemons are(sorted by BaseDefense):", "\n")
				for i := 0; i < len(base.Pokemons); i++ {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			} else if (strings.EqualFold(val, "BaseStamina")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseStamina < base.Pokemons[j].BaseStamina
				})
				fmt.Fprint(w, "All the pokemons are(sorted by BaseStamina):", "\n")
				for i := 0; i < len(base.Pokemons); i++ {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			} else if (strings.EqualFold(val, "Weight")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].Weight < base.Pokemons[j].Weight
				})
				fmt.Fprint(w, "All the pokemons are(sorted by Weight):", "\n")
				for i := 0; i < len(base.Pokemons); i++ {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			} else if (strings.EqualFold(val, "Height")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].Height < base.Pokemons[j].Height
				})
				fmt.Fprint(w, "All the pokemons are(sorted by Height):", "\n")
				for i := 0; i < len(base.Pokemons); i++ {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			}

		}

	} else if (len(q2) == 0 && strings.EqualFold(q1, "type")) { // If parameter is type but no "=" sign, all types printed!
		var i int;
		fmt.Fprint(w, "All the types are: \n")
		for i = 0; i < len(base.Types); i++ {
			fmt.Fprint(w, base.Types[i].toString())
		}
	} else if (len(q2) == 0 && strings.EqualFold(q1, "move")) { // If parameter is move but no "=" sign, all moves printed!
		var i int;
		fmt.Fprint(w, "All the moves are: \n")
		for i = 0; i < len(base.Moves); i++ {
			fmt.Fprint(w, base.Moves[i].toString())
		}
	} else if (len(q2) == 0 && strings.EqualFold(q1, "pokemon")) { // If parameter is pokemon but no "=" sign, all pokemons printed!
		fmt.Fprint(w, "All the pokemons are:", "\n")
		for i := 0; i < len(base.Pokemons); i++ {
			fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
		}
	} else if(strings.Contains(q1,"=") && strings.Contains(q2,"=")){ //If there are 2 parameters.

		s1 := strings.Split(q1, "=")
		key1, val1 := s1[0], s1[1]
		s2 := strings.Split(q2, "=")
		key2, val2 := s2[0], s2[1]
		if strings.EqualFold(key1, "type") && strings.EqualFold(key2, "sortby") { // Sort and print given type of pokemons
			if (strings.EqualFold(val2, "BaseAttack")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseAttack < base.Pokemons[j].BaseAttack
				})
				fmt.Fprint(w, val1, " type of pokemons are(sorted by base attacks):", "\n")
			} else if (strings.EqualFold(val2, "BaseDefense")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseDefense < base.Pokemons[j].BaseDefense
				})
				fmt.Fprint(w, val1, "type of pokemons are(sorted by BaseDefense):", "\n")
			} else if (strings.EqualFold(val2, "BaseStamina")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseStamina < base.Pokemons[j].BaseStamina
				})
				fmt.Fprint(w, val1, "type of pokemons are(sorted by BaseStamina):", "\n")
			} else if (strings.EqualFold(val2, "Weight")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].Weight < base.Pokemons[j].Weight
				})
				fmt.Fprint(w, val1, "type of pokemons are(sorted by Weight):", "\n")
			} else if (strings.EqualFold(val2, "Height")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].Height < base.Pokemons[j].Height
				})
				fmt.Fprint(w, val1, "type of pokemons are(sorted by Height):", "\n")
			}
			for i := 0; i < len(base.Pokemons); i++ {
				if (len(base.Pokemons[i].TypeII) != 0 && len(base.Pokemons[i].TypeI) != 0 && (strings.EqualFold(base.Pokemons[i].TypeI[0], val1) || strings.EqualFold(base.Pokemons[i].TypeII[0], val1))) {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			}
		} else if strings.EqualFold(key1, "sortby") && strings.EqualFold(key2, "type") {
			if (strings.EqualFold(val1, "BaseAttack")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseAttack < base.Pokemons[j].BaseAttack
				})
				fmt.Fprint(w, val2, " type of pokemons are(sorted by base attacks):", "\n")
			} else if (strings.EqualFold(val1, "BaseDefense")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseDefense < base.Pokemons[j].BaseDefense
				})
				fmt.Fprint(w, val2, "type of pokemons are(sorted by BaseDefense):", "\n")
			} else if (strings.EqualFold(val1, "BaseStamina")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].BaseStamina < base.Pokemons[j].BaseStamina
				})
				fmt.Fprint(w, val2, "type of pokemons are(sorted by BaseStamina):", "\n")
			} else if (strings.EqualFold(val1, "Weight")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].Weight < base.Pokemons[j].Weight
				})
				fmt.Fprint(w, val2, "type of pokemons are(sorted by Weight):", "\n")
			} else if (strings.EqualFold(val1, "Height")) {
				sort.Slice(base.Pokemons, func(i, j int) bool {
					return base.Pokemons[i].Height < base.Pokemons[j].Height
				})
				fmt.Fprint(w, val2, "type of pokemons are(sorted by Height):", "\n")
			}
			for i := 0; i < len(base.Pokemons); i++ {
				if (len(base.Pokemons[i].TypeII) != 0 && len(base.Pokemons[i].TypeI) != 0 && (strings.EqualFold(base.Pokemons[i].TypeI[0], val2) || strings.EqualFold(base.Pokemons[i].TypeII[0], val2))) {
					fmt.Fprint(w, base.Pokemons[i].toString(), "\n")
				}
			}
		}
	}else{ //OTHER EVERYTHING
		otherwise(w,r)
	}

}

func (t Type) toString() string {
	/*
	* Type struct type object informations to string method.
	*/
	var s string = "\n" + t.Name + "\nIt is effective against: "
	var str string = strings.Join(t.EffectiveAgainst, ", ")
	s += str
	s += "\nIt is weak against: "
	var str2 string = strings.Join(t.WeakAgainst, ", ")
	s += str2 + "\n"
	return s
}
func (pokemon *Pokemon) toString() string {
	/*
	* Pokemon struct type object informations to string method.
 	*/
	s := []string{pokemon.Name, "\tWeight: " + pokemon.Weight, "\tHeight: " + pokemon.Height, "\tBaseAttack: " + strconv.Itoa(pokemon.BaseAttack), "\tBaseDefense: " + strconv.Itoa(pokemon.BaseDefense), "\tBaseStamina: " + strconv.Itoa(pokemon.BaseStamina), "\tNext evolutions: \n"}
	var evs [] string
	evs = make([]string, len(pokemon.NextEvolutions))
	var i int
	for i = 0; i < len(pokemon.NextEvolutions); i++ {
		evs[i] = "\t\t-" + pokemon.NextEvolutions[i].Name
	}
	var str string = strings.Join(s, "\n")
	var str2 string = strings.Join(evs, "\n")
	return str + str2
}

func (move Move) toString() string {
	/*
	* Move struct type object informations to string method.
 	*/
	var s string = "\n\nName: " + move.Name + "\nDamage:" + strconv.Itoa(move.Damage) + "\nDps:" + strconv.FormatFloat(move.Dps, 'f', -1, 64) + "\nDuration:" + strconv.Itoa(move.Duration) + "\nEnergy:" + strconv.Itoa(move.Energy) + "\nId:" + strconv.Itoa(move.ID) + "\nType:" + move.Type
	return s
}

func (Base *BaseData) getHandler(w http.ResponseWriter, r *http.Request) {
	/*
	* Get http request Handler
 	*/
	params := mux.Vars(r)
	q1 := params["q1"]
	i := Base.getPokemon(q1)
	if (i >= 0) {
		fmt.Fprint(w, Base.Pokemons[i].toString())
		return
	}
	j := Base.getMove(q1)
	if (j >= 0) {
		fmt.Fprint(w, Base.Moves[j].toString())
		return
	}
	k := Base.getType(q1)
	if (k >= 0) {
		fmt.Fprint(w, Base.Types[k].toString())
		return
	}
	otherwise(w,r)

}
func (base *BaseData) getPokemon(string2 string) int {
	/*
	* Getting a pokemon from its name.
 	*/
	var i int
	for i = 0; i < len(base.Pokemons); i++ {
		if strings.EqualFold(string2, base.Pokemons[i].Name) {
			return i
		}
	}
	return -1
}
func (base *BaseData) getMove(string2 string) int {
	/*
	* Getting a move from its name.
 	*/
	var i int
	for i = 0; i < len(base.Moves); i++ {
		if strings.EqualFold(string2, base.Moves[i].Name) {
			return i
		}
	}
	return -1
}
func (base *BaseData) getType(string2 string) int {
	/*
	* Getting a type from its name.
	 */
	var i int
	for i = 0; i < len(base.Types); i++ {
		if strings.EqualFold(string2, base.Types[i].Name) {
			return i
		}
	}
	return -1
}
func otherwise(w http.ResponseWriter, r *http.Request) {
	/*
	* Introduction and NotFound page design
	 */
	fmt.Fprint(w, "Either This is your first time, or you tried to do something I couldn't!\n")
	fmt.Fprint(w, "There are 2 ways of getting information from this API!:\n")
	fmt.Fprint(w, "*Get:\n\tYou can use \"get\" to get detailed information about anything... Such as:\n\t- http://localhost:8080/get/Pikachu")
	fmt.Fprint(w, "\n\t- http://localhost:8080/get/Bug\n\t- http://localhost:8080/get/Wrap\n\n")
	fmt.Fprint(w, "*List:\n\tYou can use \"list\" to get listed information about anything... Such as:\n\t- http://localhost:8080/list/Move")
	fmt.Fprint(w, "\n\t- http://localhost:8080/list/Type\n\t- http://localhost:8080/list/pokemon\n")
	fmt.Fprint(w, "{!}Also you can use \"list\" with some filters(type and sortby) Such as:\n\t- http://localhost:8080/list/type=Bug")
	fmt.Fprint(w, "\n\t- http://localhost:8080/list/sortby=BaseAttack\n\t- http://localhost:8080/list/sortby=BaseAttack/type=Bug\n\n")
	fmt.Fprint(w, "Thanks a lot! Pika, Pika!!")
}

func main() {
	//read data.json to a BaseData
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("opening json file", err.Error())
	}
	var b BaseData
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&b); err != nil {
		fmt.Println("parsing json file", err.Error())
	}

	//Handlers declerations
	r := mux.NewRouter()
	r.HandleFunc("/list{sl0:[/]*}{q1:[a-zA-Z0-9=]*}{sl1:[/]*}{q2:[a-zA-Z0-9=]*}{sl2:[/]*}", b.listHandler)
	r.HandleFunc("/get{sl0:[/]*}{q1:[a-zA-Z0-9=]*}{sl1:[/]*}", b.getHandler)
	r.HandleFunc("/", otherwise)
	r.NotFoundHandler = http.HandlerFunc(otherwise)

	http.Handle("/", r)
	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
