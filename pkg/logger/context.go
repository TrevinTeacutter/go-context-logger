package logger

import (
	"strings"
)

type Context struct {
	prefix        string
	fields        map[Field]interface{}
	children      []*Context
	configuration *Configuration
}

func CreateContext(prefix string) *Context {
	return &Context{
		prefix:   prefix,
		fields:   make(map[Field]interface{}),
		children: make([]*Context, 0),
	}
}

func CreateRootContext(configuration Configuration) *Context {
	root := &Context{
		prefix:        configuration.Prefix,
		fields:        make(map[Field]interface{}),
		children:      make([]*Context, 0),
		configuration: &configuration,
	}

	if !configuration.OmitMetadata {
		root.children = append(root.children, CreateBuildContext())
	}

	return root
}

func (ctx *Context) Copy() *Context {
	newCtx := &Context{
		prefix:        ctx.prefix,
		configuration: ctx.configuration,
		fields:        make(map[Field]interface{}),
		children:      make([]*Context, len(ctx.children)),
	}

	for k, v := range ctx.fields {
		newCtx.fields[k] = v
	}

	for i, child := range ctx.children {
		newCtx.children[i] = child.Copy()
	}

	return newCtx
}

func (ctx *Context) WithFields(fields map[Field]interface{}) *Context {
	newCtx := ctx.Copy()

	for k, v := range fields {
		newCtx.fields[k] = v
	}

	return newCtx
}

func (ctx *Context) WithChildren(children ...*Context) *Context {
	newCtx := ctx.Copy()

	newCtx.children = append(newCtx.children, children...)

	return newCtx
}

func (ctx *Context) GetFields() map[string]interface{} {
	if ctx.configuration.Flatten {
		return ctx.getFlattenedFields([]string{}, ctx.configuration.Separator, ctx.configuration.Verbose)
	}

	return ctx.getStructuredFields(ctx.configuration.Verbose)
}

func (ctx *Context) getStructuredFields(verbose bool) map[string]interface{} {
	fields := make(map[string]interface{})

	for k, v := range ctx.fields {
		if k.shouldBeLogged(verbose) {
			continue
		}

		fields[k.name] = v
	}

	for _, child := range ctx.children {
		fields[child.prefix] = child.getStructuredFields(verbose)
	}

	return fields
}

func (ctx *Context) getFlattenedFields(paths []string, separator string, verbose bool) map[string]interface{} {
	fields := make(map[string]interface{})

	if ctx.prefix != "" {
		paths = append(paths, ctx.prefix)
	}

	for k, v := range ctx.fields {
		if k.shouldBeLogged(verbose) {
			continue
		}

		fullPath := append(paths, k.name)

		fields[strings.Join(fullPath, separator)] = v
	}

	for _, child := range ctx.children {
		results := child.getFlattenedFields(paths, separator, verbose)

		for k, v := range results {
			fields[k] = v
		}
	}

	return fields
}
