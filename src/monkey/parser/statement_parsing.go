package parser

import (
	"monkey/ast"
	"monkey/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IDENT:
		switch p.peekToken.Type {
		case token.ASSIGN:
			return p.parseAssignStatement()
		case token.ADD_ASSIGN:
			return p.parseShortcutAssignStatement(token.ADD_ASSIGN)
		case token.SUB_ASSIGN:
			return p.parseShortcutAssignStatement(token.SUB_ASSIGN)
		case token.MULT_ASSIGN:
			return p.parseShortcutAssignStatement(token.MULT_ASSIGN)
		case token.DIV_ASSIGN:
			return p.parseShortcutAssignStatement(token.DIV_ASSIGN)
		case token.MOD_ASSIGN:
			return p.parseShortcutAssignStatement(token.MOD_ASSIGN)
		case token.POW_ASSIGN:
			return p.parseShortcutAssignStatement(token.POW_ASSIGN)
		default:
			return p.parseExpressionStatement()
		}
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken, Statements: []ast.Statement{}}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.curToken}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek("=") {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseShortcutAssignStatement(tokenType token.TokenType) *ast.ShortcutAssignStatement {
	operator := string(tokenType)

	stmt := &ast.ShortcutAssignStatement{Token: p.curToken}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	stmt.Operator = operator[0 : len(operator)-1]

	if !p.expectPeek(tokenType) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
