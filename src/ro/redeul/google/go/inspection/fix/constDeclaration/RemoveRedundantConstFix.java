package ro.redeul.google.go.inspection.fix.constDeclaration;

import com.intellij.codeInspection.LocalQuickFix;
import com.intellij.codeInspection.ProblemDescriptor;
import com.intellij.openapi.editor.Editor;
import com.intellij.openapi.fileEditor.FileEditorManager;
import com.intellij.openapi.project.Project;
import com.intellij.psi.PsiElement;
import org.jetbrains.annotations.NotNull;
import ro.redeul.google.go.inspection.fix.DeleteStmtFix;
import ro.redeul.google.go.lang.psi.declarations.GoConstDeclaration;
import ro.redeul.google.go.lang.psi.expressions.GoExpr;
import ro.redeul.google.go.lang.psi.expressions.literals.GoLiteralIdentifier;

public class RemoveRedundantConstFix implements LocalQuickFix {
    @NotNull
    @Override
    public String getName() {
        return "Remove redundant constant";
    }

    @NotNull
    @Override
    public String getFamilyName() {
        return "Constant";
    }

    @Override
    public void applyFix(@NotNull Project project, @NotNull ProblemDescriptor descriptor) {
        PsiElement element = descriptor.getStartElement();
        Editor editor = FileEditorManager.getInstance(project).getSelectedTextEditor();
        if (editor == null) {
            return;
        }

        if (!(element instanceof GoConstDeclaration)) {
            return;
        }

        GoConstDeclaration cd = (GoConstDeclaration) element;
        GoLiteralIdentifier[] ids = cd.getIdentifiers();
        GoExpr[] expressions = cd.getExpressions();

        if (expressions.length == 0) {
            DeleteStmtFix.deleteStatement(cd);
            return;
        }

        if (ids.length <= expressions.length) {
            return;
        }

        cd.deleteChildRange(ids[expressions.length - 1].getNextSibling(), ids[ids.length - 1]);
    }
}
