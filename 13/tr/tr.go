

package main
import("bufio";"os";"fmt";"flag")
func die(s string, v... interface{}) {
  fmt.Fprintf(os.Stderr, "tu: ");
  fmt.Fprintf(os.Stderr, s, v...);
  fmt.Fprintf(os.Stderr, "\n");
  os.Exit(1)

}

func main() {
  flag.Parse()
  if 1 > flag.NArg() { die("missing operand"); }
  if 2 > flag.NArg() { die("missing operand after `%s'", flag.Arg(0)); }
  if 2 < flag.NArg() { die("extra operand after `%s'", flag.Arg(1)); }
  tab := make(map[int]int)
  set1 := []int(flag.Arg(0))
  set2 := []int(flag.Arg(1))
  j := 0
  for i := 0; i < len(set1); i++ {
    tab[set1[i]] = set2[j]
    if j < len(set2) - 1 { j++ }
  }
  in := bufio.NewReader(os.Stdin)
  out := bufio.NewWriter(os.Stdout)
  flush := func() {
    if er := out.Flush(); er != nil { die("flush: %s", er.String()) }
  }
  writeRune := func(r int) {
    if _, er := out.WriteRune(r); er != nil { die("write: %s", er.String()) }
  }
  for done := false; !done; {
    switch r,_,er := in.ReadRune(); er {
    case os.EOF: done = true
    case nil:
      if s,found := tab[r]; found {
        writeRune(s)
      } else {
        writeRune(r)
      }
      if '\n' == r { flush() }
    default: die("%s: %s", os.Stdin.Name(), er.String())
    }
  }
  flush()
}

