// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

import fmt
// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct{ h, m int } // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
  c := Clock{hour, minute}
  return c
}

func (Clock *c) String() string {
 return fmt.Sprintf("%d:%d", c.h, c.m)
}

func (Clock *c) Add(minutes int) Clock {
  return
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
