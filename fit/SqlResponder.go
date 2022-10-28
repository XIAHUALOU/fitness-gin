package FitGin

import "log"

type Query interface {
	Sql() string
	Args() []interface{}
	Mapping() map[string]string
	First() bool
	Key() string
	Get() interface{}
}
type SimpleQueryWithArgs struct {
	sql        string
	args       []interface{}
	mapping    map[string]string
	fetchFirst bool
	datakey    string
}

func NewSimpleQueryWithArgs(sql string, args []interface{}) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, args: args}
}
func NewSimpleQueryWithMapping(sql string, mapping map[string]string) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, mapping: mapping}
}
func NewSimpleQueryWithFetchFirst(sql string) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, fetchFirst: true}
}
func NewSimpleQueryWithKey(sql string, key string) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, datakey: key}
}
func (self *SimpleQueryWithArgs) Sql() string {
	return self.sql
}
func (self *SimpleQueryWithArgs) Mapping() map[string]string {
	return self.mapping
}
func (self *SimpleQueryWithArgs) Args() []interface{} {
	return self.args
}
func (self *SimpleQueryWithArgs) First() bool {
	return self.fetchFirst
}
func (self *SimpleQueryWithArgs) Key() string {
	return self.datakey
}
func (self *SimpleQueryWithArgs) WithMapping(mapping map[string]string) *SimpleQueryWithArgs {
	self.mapping = mapping
	return self
}
func (self *SimpleQueryWithArgs) WithFirst() *SimpleQueryWithArgs {
	self.fetchFirst = true
	return self
}
func (self *SimpleQueryWithArgs) WithKey(key string) *SimpleQueryWithArgs {
	self.datakey = key
	return self
}
func (self *SimpleQueryWithArgs) Get() interface{} {
	ret, err := queryForMapsByInterface(self)
	if err != nil {
		log.Println("query get error:", err)
		return nil
	}
	return ret
}

type SimpleQuery string

func (self SimpleQuery) WithArgs(args ...interface{}) *SimpleQueryWithArgs {
	return NewSimpleQueryWithArgs(string(self), args)
}
func (self SimpleQuery) WithMapping(mapping map[string]string) *SimpleQueryWithArgs {
	return NewSimpleQueryWithMapping(string(self), mapping)
}
func (self SimpleQuery) WithFirst() *SimpleQueryWithArgs {
	return NewSimpleQueryWithFetchFirst(string(self))
}
func (self SimpleQuery) WithKey(key string) *SimpleQueryWithArgs {
	return NewSimpleQueryWithKey(string(self), key)
}
func (self SimpleQuery) First() bool {
	return false
}
func (self SimpleQuery) Sql() string {
	return string(self)
}
func (self SimpleQuery) Key() string {
	return ""
}
func (self SimpleQuery) Args() []interface{} {
	return []interface{}{}
}
func (self SimpleQuery) Mapping() map[string]string {
	return map[string]string{}
}
func (self SimpleQuery) Get() interface{} {
	return NewSimpleQueryWithArgs(string(self), nil).Get()
}
