package main

import (
	"flag"
	"fmt"
	"time"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error{
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid %s", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f,name,usage)
	return &f.Celsius
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	//var period = flag.Duration("period", 1*time.Second,"sleep period")
	//flag.Parse()
	//fmt.Printf("Sleeping for %v...", *period)
	//time.Sleep(*period)
	//fmt.Println()

	//var id int
	//var name string
	//fmt.Println("parsed? = ", flag.Parsed())
	//flag.IntVar(&id, "id", 1, "help msg for id")
	//flag.StringVar(&name, "name", "h", "help msg for name")
	//flag.Parse()
	//fmt.Println("parsed? = ", flag.Parsed())
	//
	//for i, v := range flag.Args() {
	//	fmt.Printf("arg[%d] = (%s).\n", i, v)
	//}
	//
	//flag.Visit(func(f *flag.Flag) {
	//	fmt.Println("1", f.Name, f.Value, f.Usage, f.DefValue)
	//})
	//
	//flag.VisitAll(func(f *flag.Flag) {
	//	fmt.Println("2", f.Name, f.Value, f.Usage, f.DefValue)
	//})
	//fmt.Println(flag.NArg(), flag.NFlag())

	var temp = CelsiusFlag("temp", 20.0, "help")
	flag.Parse()
	fmt.Println(*temp)


}
