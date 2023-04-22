package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf("Hi, %#v. Welcome!", name)
    //message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error, int) {
    messages := make(map[string]string)
    for seq, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err, seq
        }
        messages[name] = message
    }
    return messages, nil, -1
}
