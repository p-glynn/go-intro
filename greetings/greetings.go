package greetings

import (
	"errors"
	"fmt"
  "math/rand"
)

// Hello returns a personalized greeting message for the named person
func Hello(name string) (string, error) {
  // if no name was given, return an error with a message
  if name == "" {
    return "", errors.New("did you forget to add a name?")
  }
  // create a message using a random format if a name was received
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

// Hellos returns a map that associates each of
// the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
  // a map to associate names with messages
  messages := make(map[string]string)
  // loop through received slice of names
  // call Hello function to get a message for each name
  for _, name := range names {
    message, err  := Hello(name)
    if err != nil {
      return nil, err
    }
    // in the map, associated the retrieved message with the name
    messages[name] = message
  }
  return messages, nil
}

// randomFormat returns one of a set of greeting messages
// the returned message is selected at random
func randomFormat() string {
  // a slice of message formats
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
    "Hiya, %v! It's a pleasure to meet you",
    "Good morrow, %v!",
    "Good day, %v!",
  }

  // return randomly selected message format by specifying
  // a random index for slice of formats
  return formats[rand.Intn(len(formats))]
}