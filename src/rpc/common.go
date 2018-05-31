package main

// START OMIT
type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Add(args *Args, reply *int) error { // HL
	*reply = args.A + args.B
	return nil
}

// END OMIT

const serverAddress = "localhost"
