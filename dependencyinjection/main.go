package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

// Datastore for this app
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// Factory function for SimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Jonathan",
		},
	}
}

// # Business Logic:
// Next we want to write some business logic

// # Dependencies:
// Datastore: The business logic needsdata to work with, so it requires a datastore 
// Logger: and it shall log when it get invoked, so it needs a logger
// Thus here we define the dependencies (so that we can _inject_ it later!)

// First it needs some data to work with, so it *depends* on a datastore
type DataStore interface {
	// that is able to quere for a user
	UserNameForID(userID string) (string, bool)
}

// We want the business logic to log when it is invoked, so it *depends* on a logger
type Logger interface {
	Log(message string)
}

// To make your LogOutput function meet this interface, we define a function type
// with a method, that meets the interface
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// At this point LoggerAdapter and SimpleDataStore both meet the interfaces of our
// business logic (Datastore and Logger interface) without either type having
// an idea that it does - it's implicit

// Now that the _dependencies_ are defined, we can implement the business logic
// Note how nothing in SimpleLogic mentions the concrete type of the Logger and
// DataStore so we have _no_ dependency on them - we're decoupled:
// We won't later have a problem swapping in new implementations from an entirely
// different provider, because the provider has nothing to do with this interface
//
// This is very different from languages like Java:
// Even though Java uses an interface to decouple the implementation from the
// interface, the explicit interfaces binds the client and the provider together
// This makes replacing a dependency in Java (and other langs with explicit
// interfaces) far mor difficult than it is in Go
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// Note how the fields in the SimpleLogic struct are unexported!
// Thus they can only be accessed by code in this package.
// Though you can't enforce immutability in Go, you can limit which code can access
// these fields making accidental modification less likely
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func (sl SimpleLogic) SayHello(userid string) (string, error) {
	sl.l.Log("in SayHello with userid " + userid)
	name, ok := sl.ds.UserNameForID(userid)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userid string) (string, error) {
	sl.l.Log("in SayGiidbye with userid " + userid)
	name, ok := sl.ds.UserNameForID(userid)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Bye, " + name, nil
}

// Client Code

// Our API shall have a simple endpoint /hello, which says hello to the person
// whose user ID was supplied via query parameters (this is not best practice)
// your controller needs business logic that says hello, so we define an interface
// for it
type Logic interface {
	SayHello(userid string) (string, error)
}

// Again the concretetype SimpleLogic, implements this interface without being
// aware of it. But not with the "SayGoodbye()" method

// Because the "Logic" interface is owned by clientcode so its method set is
// customized to the needs of the client (no SayGoodby implemented):
type Controller struct {
	l     Logger
	logic Logic
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// With this signature, this method matches the http.Handler interface!
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHello from Controller called")
	userID := r.URL.Query().Get("user_id") // get user from query parameter
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// The main function is the only part of the code that knows what all the
// concrete types actually are. You could swap in different implementatios here
// Externalizing the dependencies via dependency injection means that you
// can limit the changes that are eeded to evolve your code over time
func main() {
	l := LoggerAdapter(LogOutput) // LoggerAdapter implements the Logger interface
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds) // and can be passed in here as a Logger!
	c := NewController(l, logic)
	// can be used as a Handler, because the method matches the http.Handler 
	// interface! 
	// The http.HandleFunc() will convert the c.SayHello to a http.HandleFunc
	// _function type_. After this, http.HandleFunc implement the http.Handler
	// interface, it can be used as a http.Handler!
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
