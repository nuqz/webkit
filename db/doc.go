// Package db contains a set of abstractions to represent the data model for
// almost any kind of database.
//
// The first basic idea is that every record (row, document, etc.) in the
// database has a so-called primary key, which is unique in the table
// (collection, etc.). Each data model must implement the PK method to provide
// information about the primary key.
//
// The second basic idea is that information about certain model may or may not
// be present in the database. In order to know exactly, the model must
// implement IsNew method. To mark model as "stored in the database" model
// should implement Store method. This package contains implementation of the
// idea which was described above - M.
//
// The third basic idea is that every model can be converted to
// map[string]interface{} via model's AsMap method.
package db
