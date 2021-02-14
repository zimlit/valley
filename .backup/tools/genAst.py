import sys

def defineAst(outputdir, types):
    path = outputDir + "/" + "expr.go"
    outFile = open(path, "w")
    outFile.write("package expr\n\n")
    outFile.write("import \"valley/token\"\n\n")
    outFile.write("type Expr interface {\n  Accept(visitor Visitor) interface {}\n}\n")
    outFile.write("\ntype Visitor interface {\n")
    for t in types:
         typeName = t.split(":")[0].strip()
         outFile.write(f'   visit{typeName}Expr(expr {typeName}) interface {"{}"}\n')
    outFile.write("}\n")
    for t in types:
        typeName = t.split(":")[0].strip()
        fields = t.split(":")[1].strip()
        outFile.write(f'\ntype {typeName} struct {"{"}\n')
        for field in fields.split(", "):
            outFile.write(f'    {field}\n')

        outFile.write("}\n\n")
        outFile.write(f'func ({typeName.lower()} {typeName}) Accept(visitor Visitor) interface {"{}"} {"{"}\n')
        outFile.write(f'    return visitor.visit{typeName}Expr({typeName.lower()})\n')
        outFile.write("}\n")

args = sys.argv
if len(args) != 2:
    print("Usage: genAst <outputdir>")
    exit(64)

outputDir = args[1]
defineAst(outputDir, [
    "Binary   : Left Expr, Operator token.Token, Right Expr",
    "Grouping : Expression Expr",
    "Literal  : Value interface {}",
    "Unary    : Operator token.Token, Right Expr"
])
