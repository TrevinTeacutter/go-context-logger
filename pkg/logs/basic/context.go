package basic

import (
	"strings"

	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

var _ logs.Context = (*Context)(nil)

type Context struct {
	prefix        string
	fields        []logs.Field
	children      []logs.Context
}

func CreateContext(prefix string) *Context {
	return &Context{
		prefix:   prefix,
		fields:   make([]logs.Field, 0),
		children: make([]logs.Context, 0),
	}
}

func CreateRootContext(configuration logs.Configuration) *Context {
	root := &Context{
		prefix:        configuration.Prefix,
		fields:        make([]logs.Field, 0),
		children:      make([]logs.Context, 0),
	}

	if !configuration.OmitMetadata {
		root.children = append(root.children, CreateBuildContext())
	}

	return root
}

func (ctx *Context) Prefix() string {
	return ctx.prefix
}

func (ctx *Context) Copy() logs.Context {
	return ctx.copy()
}

func (ctx *Context) WithFields(fields ...logs.Field) logs.Context {
	newCtx := ctx.copy()

	newCtx.fields = append(newCtx.fields, fields...)

	return newCtx
}

func (ctx *Context) WithChildren(children ...logs.Context) logs.Context {
	newCtx := ctx.copy()

	newCtx.children = append(newCtx.children, children...)

	return newCtx
}

func (ctx *Context) Fields(verbose, flatten bool, separator string) map[string]interface{} {
	if flatten {
		return ctx.getFlattenedFields(verbose, flatten, separator)
	}

	return ctx.getStructuredFields(verbose, flatten, separator)
}

func (ctx *Context) copy() *Context {
	newCtx := &Context{
		prefix:        ctx.prefix,
		fields:        make([]logs.Field, len(ctx.fields)),
		children:      make([]logs.Context, len(ctx.children)),
	}

	for i, field := range ctx.fields {
		newCtx.fields[i] = field
	}

	for i, child := range ctx.children {
		newCtx.children[i] = child.Copy()
	}

	return newCtx
}

func (ctx *Context) getStructuredFields(verbose, flatten bool, separator string) map[string]interface{} {
	fields := make(map[string]interface{})

	for _, field := range ctx.fields {
		if !verbose && field.Verbose() {
			continue
		}

		fields[field.Key()] = field.Value()
	}

	for _, child := range ctx.children {
		fields[child.Prefix()] = child.Fields(verbose, flatten, separator)
	}

	return fields
}

func (ctx *Context) getFlattenedFields(verbose, flatten bool, separator string) map[string]interface{} {
	fields := make(map[string]interface{})

	for _, field := range ctx.fields {
		if !verbose && field.Verbose() {
			continue
		}

		fields[strings.Join([]string{ctx.prefix, field.Key()}, separator)] = field.Value()
	}

	for _, child := range ctx.children {
		results := child.Fields(verbose, flatten, separator)

		for k, v := range results {
			fields[strings.Join([]string{ctx.prefix, k}, separator)] = v
		}
	}

	return fields
}
