package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	firstHalf := fmt.Sprintf("Happy birthday %s! ", name)
    secondHalf := fmt.Sprintf("You are now %v years old!", age)
    return firstHalf + secondHalf
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcome := Welcome(name) 
    assignment := fmt.Sprintf("You have been assigned to table %03d.", table)
    directionString := fmt.Sprintf(" Your table is %s,", direction)
    distanceString := fmt.Sprintf(" exactly %.1f meters from here.\n", distance)
    neighborString := fmt.Sprintf("You will be sitting next to %s.", neighbor)

    return welcome + "\n" + assignment + directionString + distanceString + neighborString
}
