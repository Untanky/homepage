package posts

type Group struct {
	Create CreateCommand `cmd:"" help:"Create a new blog post"`
}
