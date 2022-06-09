//go:build go1.18

package parquet

import (
	"io"
	"reflect"
)

// GenericWriter is similar to a Writer but uses a type parameter to define the
// Go type representing the schema of rows being written.
//
// Using this type over Writer has multiple advantages:
//
// - By leveraging type information, the Go compiler can provide greater
//   guarantees that the code is correct. For example, the parquet.Writer.Write
//   method accepts an argument of type interface{}, which delays type checking
//   until runtime. The parquet.GenericWriter[T].Write method ensures at
//   compile time that the values it receives will be of type T, reducing the
//   risk of introducing errors.
//
// - Since type information is known at compile time, the implementation of
//   parquet.GenericWriter[T] can make safe assumptions, removing the need for
//   runtime validation of how the parameters are passed to its methods.
//   Optimizations relying on type information are more effective, some of the
//   writer's state can be precomputed at initialization, which was not possible
//   with parquet.Writer.
//
// - The parquet.GenericWriter[T].Write method uses a data-oriented design,
//   accepting an slice of T instead of a single value, creating more
//   opportunities to amortize the runtime cost of abstractions.
//   This optimization is not available for parquet.Writer because its Write
//   method's argument would be of type []interface{}, which would require
//   conversions back and forth from concrete types to empty interfaces (since
//   a []T cannot be interpreted as []interface{} in Go), would make the API
//   more difficult to use and waste compute resources in the type conversions,
//   defeating the purpose of the optimization in the first place.
//
// Note that this type is only available when compiling with Go 1.18 or later.
type GenericWriter[T any] struct {
	// At this time GenericWriter is expressed in terms of Writer to reuse the
	// underlying logic. In the future, and if we accepted to break backward
	// compatibility on the Write method, we could modify Writer to be an alias
	// to GenericWriter with:
	//
	//	type Writer = GenericWriter[any]
	//
	base Writer
	// This function writes rows of type T to the writer, it gets generated by
	// the NewGenericWriter function based on the type T and the underlying
	// schema of the parquet file.
	write writeFunc[T]
	// This field is used to leverage the optimized writeRowsFunc algorithms.
	buffers columnBufferWriter
}

// NewGenericWriter is like NewWriter but returns a GenericWriter[T] suited to
// write rows of Go type T.
//
// The type parameter T should be a map, struct, or any. Any other types will
// cause a panic at runtime. Type checking is a lot more effective when the
// generic parameter is a struct type, using map and interface types is somewhat
// similar to using a Writer.
//
// If the option list may explicitly declare a schema, it must be compatible
// with the schema generated from T.
func NewGenericWriter[T any](output io.Writer, options ...WriterOption) *GenericWriter[T] {
	config, err := NewWriterConfig(options...)
	if err != nil {
		panic(err)
	}

	t := typeOf[T]()
	schema := schemaOf(dereference(t))
	if config.Schema == nil {
		config.Schema = schema
	}

	return &GenericWriter[T]{
		base: Writer{
			output: output,
			config: config,
			schema: schema,
			writer: newWriter(output, config),
		},
		write: writeFuncOf[T](t, config.Schema),
	}
}

type writeFunc[T any] func(*GenericWriter[T], []T) (int, error)

func writeFuncOf[T any](t reflect.Type, schema *Schema) writeFunc[T] {
	switch t.Kind() {
	case reflect.Interface, reflect.Map:
		return (*GenericWriter[T]).writeRows

	case reflect.Struct:
		return makeWriteFunc[T](t, schema)

	case reflect.Pointer:
		if e := t.Elem(); e.Kind() == reflect.Struct {
			return makeWriteFunc[T](t, schema)
		}
	}
	panic("cannot create writer for values of type " + t.String())
}

func makeWriteFunc[T any](t reflect.Type, schema *Schema) writeFunc[T] {
	size := t.Size()
	writeRows := writeRowsFuncOf(t, schema, nil)
	return func(w *GenericWriter[T], rows []T) (n int, err error) {
		defer w.buffers.clear()
		err = writeRows(&w.buffers, makeArray(rows), size, 0, columnLevels{})
		if err == nil {
			n = len(rows)
		}
		return n, err
	}
}

func (w *GenericWriter[T]) Close() error {
	return w.base.Close()
}

func (w *GenericWriter[T]) Flush() error {
	return w.base.Flush()
}

func (w *GenericWriter[T]) Reset(output io.Writer) {
	w.base.Reset(output)
}

func (w *GenericWriter[T]) Write(rows []T) (int, error) {
	if w.buffers.columns == nil {
		w.buffers = columnBufferWriter{
			columns: make([]ColumnBuffer, len(w.base.writer.columns)),
			values:  make([]Value, 0, defaultValueBufferSize),
		}

		for i, c := range w.base.writer.columns {
			// These fields are usually lazily initialized when writing rows,
			// we need them to exist now tho.
			c.columnBuffer = c.newColumnBuffer()
			c.maxValues = int32(c.columnBuffer.Cap())
			w.buffers.columns[i] = c.columnBuffer
		}
	}

	n, err := w.write(w, rows)
	if err != nil {
		return n, err
	}

	for _, c := range w.base.writer.columns {
		c.numValues = int32(c.columnBuffer.NumValues())

		if c.numValues > 0 && c.numValues >= c.maxValues {
			if err := c.flush(); err != nil {
				return 0, err
			}
		}
	}

	return n, nil
}

func (w *GenericWriter[T]) WriteRows(rows []Row) (int, error) {
	return w.base.WriteRows(rows)
}

func (w *GenericWriter[T]) WriteRowGroup(rowGroup RowGroup) (int64, error) {
	return w.base.WriteRowGroup(rowGroup)
}

func (w *GenericWriter[T]) ReadRowsFrom(rows RowReader) (int64, error) {
	return w.base.ReadRowsFrom(rows)
}

func (w *GenericWriter[T]) Schema() *Schema {
	return w.base.Schema()
}

func (w *GenericWriter[T]) writeRows(rows []T) (int, error) {
	if cap(w.base.rowbuf) < len(rows) {
		w.base.rowbuf = make([]Row, len(rows))
	} else {
		w.base.rowbuf = w.base.rowbuf[:len(rows)]
	}
	defer clearRows(w.base.rowbuf)

	schema := w.base.Schema()
	for i := range rows {
		w.base.rowbuf[i] = schema.Deconstruct(w.base.rowbuf[i], &rows[i])
	}

	return w.base.WriteRows(w.base.rowbuf)
}

var (
	_ RowWriterWithSchema = (*GenericWriter[any])(nil)
	_ RowReaderFrom       = (*GenericWriter[any])(nil)
	_ RowGroupWriter      = (*GenericWriter[any])(nil)

	_ RowWriterWithSchema = (*GenericWriter[struct{}])(nil)
	_ RowReaderFrom       = (*GenericWriter[struct{}])(nil)
	_ RowGroupWriter      = (*GenericWriter[struct{}])(nil)

	_ RowWriterWithSchema = (*GenericWriter[map[struct{}]struct{}])(nil)
	_ RowReaderFrom       = (*GenericWriter[map[struct{}]struct{}])(nil)
	_ RowGroupWriter      = (*GenericWriter[map[struct{}]struct{}])(nil)
)
