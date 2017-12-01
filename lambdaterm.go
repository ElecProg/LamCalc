package LamCalc

import "reflect"

// LamTerm is a general type to represent both LamExprns and LamFuncs
type LamTerm interface {
	// Allow us to manipulate it as a list
	Len() int
	Index(int) interface{}
	Append(...interface{}) LamTerm

	// Equality checking (is the same as alpha equivalence because of the De Bruijn indexes)
	Equals(other LamTerm) bool

	// Lambda print
	String() string
	deDeBruijn(boundLetters []string, nextletter int) string

	// Lambda expand
	Substitute(int, LamTerm) LamTerm
	ExpandOnce() LamTerm
	Expand() LamFunc
}

// LamExpr is a list of lamfuncs, lamexprns and De Bruijn indexes (all lowered by one) which isn't a function itself.
type LamExpr []interface{}

// Len returns the lenght of the LamExpr
func (lx LamExpr) Len() int {
	return len(lx)
}

// Index returns the ith element of the LamExpr
func (lx LamExpr) Index(i int) interface{} {
	return lx[i]
}

// Append allows us to append to the LamExpr
func (lx LamExpr) Append(terms ...interface{}) LamTerm {
	return append(lx, terms...)
}

// Equals checks wether the LamExp is identical to a LamTerm
func (lx LamExpr) Equals(other LamTerm) bool {

	if reflect.TypeOf(other).String() != "LamCalc.LamExpr" {
		return false

	} else if len(lx) != other.Len() {
		return false
	}

	for i := range lx {
		switch elem := lx[i].(type) {
		case int:
			if reflect.TypeOf(other.Index(i)).Kind() != reflect.Int || elem != other.Index(i).(int) {
				return false
			}

		case LamExpr:
			if reflect.TypeOf(other.Index(i)).String() != "LamCalc.LamExpr" || !elem.Equals(other.Index(i).(LamExpr)) {
				return false
			}

		case LamFunc:
			if reflect.TypeOf(other.Index(i)).String() != "LamCalc.LamFunc" || !LamExpr(elem).Equals(LamExpr(other.Index(i).(LamFunc))) {
				return false
			}

		default:
			return false
		}
	}

	return true
}

// LamFunc is a list of lamfuncs, lamexprns and De Bruijn indexes (all lowered by one)
type LamFunc []interface{}

// Len returns the lenght of the LamFunc
func (lf LamFunc) Len() int {
	return len(lf)
}

// Index returns the ith element of the LamFunc
func (lf LamFunc) Index(i int) interface{} {
	return lf[i]
}

// Append allows us to append to the LamFunc
func (lf LamFunc) Append(terms ...interface{}) LamTerm {
	return append(lf, terms...)
}

// Equals checks wether two LamFuncs are identical
func (lf LamFunc) Equals(other LamTerm) bool {

	if reflect.TypeOf(other).String() != "LamCalc.LamFunc" {
		return false
	}

	return LamExpr(lf).Equals(LamExpr(other.(LamFunc)))
}
