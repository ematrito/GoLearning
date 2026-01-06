package airportrobot

import (
    "fmt"
)
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(visitor string) string
}

type German struct {
}
     
func (g German) LanguageName() string {
     return "German"    
 }   

func (g German) Greet(visitor string) string {
     return fmt.Sprintf(": Hallo %s!", visitor)
}

type Italian struct {
}
     
func (i Italian) LanguageName() string {
     return "Italian"    
 }   

func (i Italian) Greet(visitor string) string {
     return fmt.Sprintf(": Ciao %s!", visitor)
}

type Portuguese struct {
}
     
func (p Portuguese) LanguageName() string {
     return "Portuguese"    
 }   

func (p Portuguese) Greet(visitor string) string {
     return fmt.Sprintf(": Olá %s!", visitor)
}


func SayHello(visitorName string, greeter Greeter) string {
    langName := greeter.LanguageName()
    greeting := greeter.Greet(visitorName)
    return "I can speak " + langName + greeting
}