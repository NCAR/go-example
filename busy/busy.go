/*Package busy contains a simple example utilizing go routines*/
package busy

import (
	"fmt"
	"github.com/NCAR/go-example/say"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Second()))
}

func randDuration() (time.Duration, int) {
	r := rand.Intn(1000)
	return time.Duration(r) * time.Microsecond, r
}

type data struct {
	When time.Time
	s    say.Sayer //remember, say.Sayer is an interface
}

/*our data struct is now a Sayer.  Conceptually, the
following is now legal:
  s := data{Time: time.Now()}
  s.s = &s;
*/
func (d *data) Say() string {
	return "Doh!"
}

/*Basic shows a very basic usage of channels*/
func Basic(writers, readers int) {
	// the make keywork creates the channel on the heap of type "chan data"
	chanData := make(chan data, 0)

	writer := func(id int) {
		//wait's type are determined by the return value of randDuration(),
		// and the second returned argument is being discarded via the _
		// keyword
		wait, _ := randDuration()

		// time.After() returns a (single fire) channel , which
		// this will block on until it can be read (and discarded)
		// via the <- operator
		<-time.After(wait)

		//closure here in accessing chanData.  This blocks
		// until someone does a synchronous read from chanData
		chanData <- data{When: time.Now(), s: say.NewI(int64(id))}
	}

	now := time.Now()

	for i := 0; i < writers; i++ {
		//the 'go' keyword invokes calling the passed function on a separate
		// goroutine.  In this case, we are attempting to write data objects
		// to chanData 'writers' times
		go writer(i)
	}

	for i := 0; i < readers; i++ {
		// d's type is determined from the definition of chanData
		// this time, we are using the <- operator to synchronously read
		// from one of the above goroutines.  Which is not
		d := <-chanData
		fmt.Printf("%s: %v\n", time.Since(d.When), d.s.Say())
	}
	fmt.Println("Operation took ", time.Since(now))
}

/*MapRace purposefully forces race conditions, to demonstrate the
functionality of the 'go test -race' race detector*/
func MapRace(N int) {
	//wait group, so we don't exit MapRace while data is being modified
	wg := sync.WaitGroup{}
	wg.Add(N)

	//data is an empty has map of int->timestamps
	//maps are not thread safe by definition
	data := map[int]time.Time{}

	f := func() {
		d, _ := randDuration()
		<-time.After(d) //random delay
		//Pick a random int from 0:9 and attach the current time stamp
		data[rand.Intn(10)] = time.Now()
		wg.Done() //tells the wait group this is done
	}

	for i := 0; i < N; i++ {
		go f()
	}

	wg.Wait() //waits for all go routines to return
}
