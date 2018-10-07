package merkletree

import (
	"os"
	"fmt"
	"testing"
	"io/ioutil"
	"crypto/rand"
	//
	"golang.org/x/crypto/sha3"
)

func TestPackage(t *testing.T) {

	os.Mkdir("tmp", 0777)

	for _, n := range []int{0, 1, 2, 3, 10, 100} {

		t.Run(
			fmt.Sprintf("TESTING WITH LEAVES %d", n),
			func (t *testing.T) {

				tree := New(sha3.New256)
				for x := 0; x < n; x++ {

					k := fmt.Sprintf("%d", x)

					b := make([]byte, 1024)
					rand.Read(b)

					tree.Add(k, b)

					err := ioutil.WriteFile("tmp/"+k, b, 0777)
					if err != nil {
						panic(err)
					}
				}
				root1 := fmt.Sprintf("%x", tree.Root())

				tree = New(sha3.New256)
				for x := 0; x < n; x++ {

					k := fmt.Sprintf("%d", x)

					b, err := ioutil.ReadFile("tmp/"+k)
					if err != nil {
						panic(err)
					}

					tree.Add(k, b)

				}
 				root2 := fmt.Sprintf("%x", tree.Root())

				if root1 != root2 {

					t.Errorf("Failed to verify!")
					return

				}

			},
		)


	}

}
