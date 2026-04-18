package airportrobot

import "fmt"
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface{
    LanguageName() string
    Greet(person string) string
}

type German struct {}

type Italian struct {}

type Portuguese struct {}

func (g German) LanguageName() string {
    return "German";
}

func (i Italian) LanguageName() string {
    return "Italian";
}

func (p Portuguese) LanguageName() string {
    return "Portuguese";
}

func (g German) Greet(person string) string{
    return fmt.Sprintf("Hallo %s!", person)
}

func (i Italian) Greet(person string) string{
    return fmt.Sprintf("Ciao %s!", person)
}

func (p Portuguese) Greet(person string) string{
    return fmt.Sprintf("Olá %s!", person)
}

func SayHello(name string, g Greeter) string{
    return fmt.Sprintf("I can speak %s: %s",g.LanguageName(), g.Greet(name));
}