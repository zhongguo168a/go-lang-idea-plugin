package main
func main() {
    switch a++; i := x[i].(type) {
        case a:
             return 1
    }
}
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  FunctionDeclaration(main)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('main')
    PsiElement(()('(')
    PsiElement())(')')
    PsiWhiteSpace(' ')
    BlockStmtImpl
      PsiElement({)('{')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      SwitchTypeStmtImpl
        PsiElement(KEYWORD_SWITCH)('switch')
        PsiWhiteSpace(' ')
        IncDecStmt
          LiteralExpressionImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('a')
          PsiElement(++)('++')
        PsiElement(;)(';')
        PsiWhiteSpace(' ')
        SwitchTypeGuardImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('i')
          PsiWhiteSpace(' ')
          PsiElement(:=)(':=')
          PsiWhiteSpace(' ')
          IndexExpressionImpl
            LiteralExpressionImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('x')
            PsiElement([)('[')
            LiteralExpressionImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('i')
            PsiElement(])(']')
          PsiElement(.)('.')
          PsiElement(()('(')
          PsiElement(KEYWORD_TYPE)('type')
          PsiElement())(')')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('        ')
        SwitchTypeCaseImpl
          PsiElement(KEYWORD_CASE)('case')
          PsiWhiteSpace(' ')
          TypeNameImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('a')
          PsiElement(:)(':')
          PsiWhiteSpace('\n')
          PsiWhiteSpace('             ')
          ReturnStmtImpl
            PsiElement(KEYWORD_RETURN)('return')
            PsiWhiteSpace(' ')
            LiteralExpressionImpl
              LiteralIntegerImpl
                PsiElement(LITERAL_INT)('1')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('    ')
        PsiElement(})('}')
      PsiWhiteSpace('\n')
      PsiElement(})('}')
