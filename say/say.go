/*Package say is an example go package with no real functionality

It covers the basics of:
 - interfaces
 - type syntax
*/
package say

import (
	"fmt"
	sc "strconv" //import strconv package as "sc"
)

/*I typedef for int64 (64 bit int).  Typedef'd parameters can
add any sort of functionality*/
type I int64

var q I      //q is of type() Isay defaults to 0
var r *I     //r defaults to a null pointer, which is  nil
var s = I(4) // s is a pointer to an Isay struct of value 5

/*init is called on module init at runtime. runtime order is as follows:
- Package wide variables initialized (such as q, r, and s above
- whatever actions init() performs
*/
func init() {
	var val = I(54) // syntax is allowed
	val2 := I(55)   // but this form is better
	q = val2 - val  //setting q to something non-zero

	r = &val //this is both syntactically ok and safe, unlike C/C++
	if *r != val {
		panic("r is the same as I")
	}
}

/*Sayer is an exported interface declaration. Title case
declarations are exported.*/
type Sayer interface {
	Say() string
}

/*usay is a un-exported by way of being a lower case
character: non-package member cannot (easily) access.*/
type usay uint

/*Say is a function receiver attached to the definition of
usay. u, in this context, is unalterable but accessible for
comparison and use. usay is a Sayer via the same signature*/
func (u usay) Say() string {
	switch u {
	case 2, 4, 6:
		fallthrough //case blocks don't 'fallthrough' automatically
	case 8, 10:
		return "Even!"
	case 1, 3, 5, 7, 9:
		return "Odd!"
	case 0: //ordering in case is irrelevant
		return "Zero!!!!"
	default:
		return "Too Big!!"
	}
}

/*NewI returns a instantiated I from the passed value i*/
func NewI(i int64) *I {
	r := I(i)
	return &r //safe & legal
}

/*asUsay returns a usay type from I.  CamelCase is preferred in
Go, and tools such as golint will nag if you use snake_case.
asUsay can only be directly called by things inside the package,
the definition is not exported to outside callers*/
func (i I) asUsay() usay {
	return usay(uint(i))
}

/*Say conforms to the Sayer interface. Say is a pointer receiver on I.*/
func (i *I) Say() (r string) {
	/*r is equal to "" - all variables have a initialization
	state. For strings, this is an empty string*/
	if *i < 0 { //no need for () around most conditionals
		*i = -*i
	}

	r = sc.FormatInt((int64(*i)), 10) + ": " + i.asUsay().Say()
	//type safety mandates using int64() in an explicit conversion

	//one annoyance is that every function needs an explicit return
	return
}

/*CountDown prints some interesting factoids*/
func (i I) CountDown() {
	for j := I(0); j < i; j++ {
		fmt.Println(j.Say())
	}
}

/*Parse is an exposed pointer receiver.  It takes any single
argument and attempt to set i to the value.  interface{} can
be seen as 'anything' variable, but is a bit more nuanced.*/
func (i *I) CanConvert(iface interface{}) bool {
	//type switch:  uses introspection internally
	switch v := iface.(type) {
	//various forms of ints. A int is always 32 bits.
	case uint8, int8, int16, uint16, int32, uint32, int, uint, int64, uint64:
		return true
	case float32, float64:
		return true
	case []byte, string:
		return false
	case usay, I:
		return true
	case Sayer: //can check for a sayer interface
		return true
	case func() int:
		//functions are first class citizens, and can be treated as POD
		return true
	default:
		fmt.Println(`This wont compile unless v is used.
backtick quotes span multiple lines`, v)
	}
	return false
}
