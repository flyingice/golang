package service

// Service type
type Service struct {
}

// Args type
type Args struct {
	Left, Right int
}

// Add method
func (*Service) Add(args *Args, reply *int) error {
	*reply = args.Left + args.Right
	return nil
}
