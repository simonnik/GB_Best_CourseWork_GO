package parser

// Query represents a parsed query
type Query struct {
	Type          Type
	TableName     string
	Conditions    []Condition
	LastCondWhere CondWhere
	Fields        []string // Used for SELECT (i.e. SELECTed field names)
}

// Type is the type of SQL query, e.g. SELECT
type Type int

const (
	// UnknownType is the zero value for a Type
	UnknownType Type = iota
	// Select represents a SELECT query
	Select
)

// Operator is between operands in a condition
type Operator int

const (
	// UnknownOperator is the zero value for an Operator
	UnknownOperator Operator = iota
	// Eq -> "="
	Eq
	// Ne -> "!="
	Ne
	// Gt -> ">"
	Gt
	// Lt -> "<"
	Lt
	// Gte -> ">="
	Gte
	// Lte -> "<="
	Lte
)

// CondWhere are AndCondition or OrCondition
type CondWhere int

const (
	// UnknownCondWhere is the zero value for an CondWhere
	UnknownCondWhere CondWhere = iota
	// And -> "AND"
	And
	// Or -> "Or"
	Or
)

// Condition is a single boolean condition in a WHERE clause
type Condition struct {
	// Operand1 is the left-hand side operand
	Operand1 string
	// Operand1IsField determines if Operand1 is a literal or a field name
	Operand1IsField bool
	// Operator is e.g. "=", ">"
	Operator Operator
	// Operand1 is the right-hand side operand
	Operand2 string
	// Operand2IsField determines if Operand2 is a literal or a field name
	Operand2IsField bool
	// Condition determines whether Condition is an And condition or an Or condition
	Condition CondWhere
}
