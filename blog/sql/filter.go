package sql

import "github.com/untanky/homepage/blog"

type Filter interface {
	Condition() string
	Args() []any
}

type basicFilter struct {
	condition string
	args      any
}

func (filter basicFilter) Condition() string {
	return filter.condition
}

func (filter basicFilter) Args() []any {
	return []any{filter.args}
}

func MatchPostID(id blog.PostID) Filter {
	return basicFilter{
		condition: "id = $2",
		args:      id,
	}
}

func MatchSlug(slug string) Filter {
	return basicFilter{
		condition: "slug = $2",
		args:      slug,
	}
}
