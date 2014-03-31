//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:22
type yySymType struct {
	yys       int
	stmt_func ast.Stmt
	stmt_if   ast.Stmt
	stmt_else ast.Stmt
	stmt_for  ast.Stmt
	stmts     []ast.Stmt
	stmt      ast.Stmt
	teim      ast.Expr
	expr      ast.Expr
	tok       Token
	idents    []string
	exprs     []ast.Expr
	pair      *ast.PairExpr
	pairs     []*ast.PairExpr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VAR = 57350
const VARARG = 57351
const FUNC = 57352
const RETURN = 57353
const IF = 57354
const ELSE = 57355
const FOR = 57356
const IN = 57357
const EQ = 57358
const NE = 57359
const GE = 57360
const LE = 57361
const OR = 57362
const AND = 57363
const UNARY = 57364

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VAR",
	"VARARG",
	"FUNC",
	"RETURN",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQ",
	"NE",
	"GE",
	"LE",
	"OR",
	"AND",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:284

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 52
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 370

var yyAct = []int{

	1, 7, 55, 20, 21, 22, 23, 24, 83, 77,
	84, 43, 87, 86, 103, 85, 52, 51, 93, 53,
	98, 57, 94, 129, 83, 126, 93, 124, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 73, 122, 121, 120, 118, 127, 79,
	117, 53, 82, 113, 110, 32, 33, 35, 37, 39,
	41, 27, 28, 29, 30, 31, 81, 108, 106, 111,
	45, 109, 107, 50, 49, 26, 90, 34, 36, 38,
	40, 76, 97, 46, 47, 99, 91, 74, 100, 101,
	80, 125, 104, 95, 56, 105, 46, 32, 33, 35,
	37, 39, 41, 92, 78, 48, 42, 112, 54, 114,
	115, 116, 6, 5, 119, 4, 3, 26, 123, 34,
	36, 38, 40, 2, 0, 0, 0, 0, 128, 32,
	33, 35, 37, 39, 41, 27, 28, 29, 30, 31,
	0, 102, 0, 0, 0, 0, 0, 0, 0, 26,
	0, 34, 36, 38, 40, 32, 33, 35, 37, 39,
	41, 27, 28, 29, 30, 31, 0, 0, 0, 0,
	96, 0, 0, 0, 0, 26, 0, 34, 36, 38,
	40, 32, 33, 35, 37, 39, 41, 27, 28, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 26, 89, 34, 36, 38, 40, 32, 33, 35,
	37, 39, 41, 27, 28, 29, 30, 31, 0, 0,
	0, 0, 88, 0, 0, 0, 0, 26, 0, 34,
	36, 38, 40, 32, 33, 35, 37, 39, 41, 27,
	28, 29, 30, 31, 0, 75, 0, 0, 0, 0,
	0, 0, 0, 26, 0, 34, 36, 38, 40, 32,
	33, 35, 37, 39, 41, 27, 28, 29, 30, 31,
	0, 25, 0, 0, 0, 0, 0, 0, 0, 26,
	0, 34, 36, 38, 40, 32, 33, 35, 37, 39,
	41, 27, 28, 29, 30, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 26, 0, 34, 36, 38,
	40, 32, 33, 35, 37, 39, 41, 0, 0, 29,
	30, 31, 14, 13, 16, 0, 8, 0, 10, 9,
	11, 26, 12, 34, 36, 38, 40, 14, 13, 16,
	0, 15, 0, 44, 0, 0, 0, 0, 19, 0,
	18, 0, 0, 0, 17, 0, 15, 0, 0, 0,
	0, 0, 0, 19, 0, 18, 0, 0, 0, 17,
}
var yyPact = []int{

	318, -1000, 318, 318, 318, 318, 318, 243, 102, 333,
	66, 54, 101, -1000, 44, 333, -1000, 333, 88, 333,
	-1000, -1000, -1000, -1000, -1000, -1000, 333, 333, 333, 333,
	333, 333, 333, 333, 333, 333, 333, 333, 333, 333,
	333, 333, 58, 217, 53, 51, 100, 333, 75, 333,
	333, 81, -27, 269, -20, -1000, -22, 191, 165, 295,
	295, 81, 81, 81, 269, 269, 269, 269, 269, 269,
	269, 269, 269, 269, 333, -1000, 99, -9, 84, 139,
	333, -11, 269, 333, -1000, 88, -1000, 333, -1000, -1000,
	113, -17, 83, 91, 36, 41, 35, 39, -1000, 269,
	-1000, 269, -1000, 22, 38, -1000, 318, 21, 318, 318,
	318, 18, 14, 318, 13, 12, 11, 318, -1000, -6,
	78, -1000, -1000, -8, -1000, 16, -1000, 318, -10, -1000,
}
var yyPgo = []int{

	0, 0, 123, 116, 115, 113, 112, 1, 16, 2,
	108, 9,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	3, 3, 5, 4, 6, 9, 10, 10, 10, 11,
	11, 11, 8, 8, 8, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 2, 5, 3,
	8, 9, 11, 7, 7, 3, 0, 1, 3, 0,
	1, 3, 0, 1, 3, 1, 1, 2, 1, 4,
	3, 7, 8, 3, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 8, 11,
	10, 12, 14, 5, 4, 23, 6, 36, 32, 30,
	-1, -1, -1, -1, -1, 28, 36, 22, 23, 24,
	25, 26, 16, 17, 38, 18, 39, 19, 40, 20,
	41, 21, 4, -7, 10, 4, 30, 30, 4, 30,
	29, -7, -8, -7, -10, -9, 6, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, 29, 28, 30, -11, 4, -7,
	15, -8, -7, 35, 37, 35, 33, 34, 31, 37,
	-7, -11, 4, 35, 31, 9, 31, -7, 31, -7,
	-9, -7, 28, 31, 9, 4, 32, 31, 32, 32,
	32, 31, -1, 32, -1, -1, -1, 32, 33, -1,
	33, 33, 33, -1, 33, 13, 33, 32, -1, 33,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 1, 0, 0, 0,
	0, 0, 0, 25, 26, 0, 28, 22, 16, 0,
	2, 3, 4, 5, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 19, 0, 0, 22,
	0, 27, 0, 23, 0, 17, 0, 0, 0, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	48, 49, 50, 51, 0, 9, 19, 0, 20, 0,
	0, 0, 47, 0, 30, 0, 33, 0, 35, 34,
	0, 0, 20, 0, 0, 0, 0, 0, 29, 24,
	18, 15, 8, 0, 0, 21, 1, 0, 1, 1,
	1, 0, 0, 1, 0, 0, 0, 1, 31, 0,
	13, 14, 10, 0, 32, 0, 11, 1, 0, 12,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 26, 41, 3,
	30, 31, 24, 22, 35, 23, 3, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 34, 28,
	39, 29, 38, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 36, 3, 37, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 32, 40, 33,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	27,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:47
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:54
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-1].stmt.Position(l.pos)
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:62
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:69
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:76
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:83
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:91
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-1].expr}
		}
	case 8:
		//line parser.go.y:95
		{
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expr}
		}
	case 9:
		//line parser.go.y:99
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-1].expr}
		}
	case 10:
		//line parser.go.y:104
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 11:
		//line parser.go.y:108
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 12:
		//line parser.go.y:113
		{
			yyVAL.stmt_else = &ast.IfStmt{Expr: yyS[yypt-8].expr, ThenStmts: yyS[yypt-5].stmts, ElseStmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:118
		{
			yyVAL.stmt_if = &ast.IfStmt{Expr: yyS[yypt-4].expr, ThenStmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:123
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:128
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 16:
		//line parser.go.y:133
		{
			yyVAL.pairs = []*ast.PairExpr{}
		}
	case 17:
		//line parser.go.y:137
		{
			yyVAL.pairs = []*ast.PairExpr{yyS[yypt-0].pair}
		}
	case 18:
		//line parser.go.y:141
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 19:
		//line parser.go.y:146
		{
			yyVAL.idents = []string{}
		}
	case 20:
		//line parser.go.y:150
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 21:
		//line parser.go.y:154
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 22:
		//line parser.go.y:159
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 23:
		//line parser.go.y:163
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 24:
		//line parser.go.y:167
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 25:
		//line parser.go.y:172
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 26:
		//line parser.go.y:176
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 27:
		//line parser.go.y:180
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 28:
		//line parser.go.y:184
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 29:
		//line parser.go.y:188
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 30:
		//line parser.go.y:192
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 31:
		//line parser.go.y:196
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 32:
		//line parser.go.y:200
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 33:
		//line parser.go.y:204
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.Key] = v.Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 34:
		//line parser.go.y:212
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 35:
		//line parser.go.y:216
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 36:
		//line parser.go.y:220
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 37:
		//line parser.go.y:224
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 38:
		//line parser.go.y:228
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:232
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:236
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 41:
		//line parser.go.y:240
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 42:
		//line parser.go.y:244
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 43:
		//line parser.go.y:248
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 44:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 45:
		//line parser.go.y:256
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:260
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 47:
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
		}
	case 48:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 49:
		//line parser.go.y:272
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 50:
		//line parser.go.y:276
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 51:
		//line parser.go.y:280
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
