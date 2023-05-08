import * as vscode from 'vscode'

import {
    ActiveTextEditor,
    ActiveTextEditorSelection,
    ActiveTextEditorVisibleContent,
    Editor,
} from '@sourcegraph/cody-shared/src/editor'
import { SURROUNDING_LINES } from '@sourcegraph/cody-shared/src/prompt/constants'

import { FileChatProvider } from '../command/FileChatProvider'

export class VSCodeEditor implements Editor {
    constructor(public fileChatProvider: FileChatProvider) {}

    public getWorkspaceRootPath(): string | null {
        const uri = vscode.window.activeTextEditor?.document?.uri
        if (uri) {
            const wsFolder = vscode.workspace.getWorkspaceFolder(uri)
            if (wsFolder) {
                return wsFolder.uri.fsPath
            }
        }
        return vscode.workspace.workspaceFolders?.[0]?.uri?.fsPath ?? null
    }

    public getActiveTextEditor(): ActiveTextEditor | null {
        const activeEditor = this.getActiveTextEditorInstance()
        if (!activeEditor) {
            return null
        }
        const documentUri = activeEditor.document.uri
        const documentText = activeEditor.document.getText()
        return { content: documentText, filePath: documentUri.fsPath }
    }

    private getActiveTextEditorInstance(): vscode.TextEditor | null {
        const activeEditor = vscode.window.activeTextEditor
        return activeEditor && activeEditor.document.uri.scheme === 'file' ? activeEditor : null
    }

    public getActiveTextEditorSelection(): ActiveTextEditorSelection | null {
        const activeEditor = this.getActiveTextEditorInstance()
        if (!activeEditor && this.fileChatProvider.selection) {
            return this.fileChatProvider.selection
        }
        if (!activeEditor) {
            return null
        }
        const selection = activeEditor.selection
        if (!selection || selection?.start.isEqual(selection.end)) {
            void vscode.window.showErrorMessage('No code selected. Please select some code and try again.')
            return null
        }
        return this.createActiveTextEditorSelection(activeEditor, selection)
    }

    public getActiveTextEditorSelectionOrEntireFile(): ActiveTextEditorSelection | null {
        const activeEditor = this.getActiveTextEditorInstance()
        if (!activeEditor) {
            return null
        }
        let selection = activeEditor.selection
        if (!selection || selection.isEmpty) {
            selection = new vscode.Selection(0, 0, activeEditor.document.lineCount, 0)
        }
        return this.createActiveTextEditorSelection(activeEditor, selection)
    }

    private createActiveTextEditorSelection(
        activeEditor: vscode.TextEditor,
        selection: vscode.Selection
    ): ActiveTextEditorSelection {
        const precedingText = activeEditor.document.getText(
            new vscode.Range(
                new vscode.Position(Math.max(0, selection.start.line - SURROUNDING_LINES), 0),
                selection.start
            )
        )
        const followingText = activeEditor.document.getText(
            new vscode.Range(selection.end, new vscode.Position(selection.end.line + SURROUNDING_LINES, 0))
        )

        return {
            fileName: vscode.workspace.asRelativePath(activeEditor.document.uri.fsPath),
            selectedText: activeEditor.document.getText(selection),
            precedingText,
            followingText,
        }
    }

    public getActiveTextEditorVisibleContent(): ActiveTextEditorVisibleContent | null {
        const activeEditor = this.getActiveTextEditorInstance()
        if (!activeEditor) {
            return null
        }
        const visibleRanges = activeEditor.visibleRanges
        if (visibleRanges.length === 0) {
            return null
        }
        const visibleRange = visibleRanges[0]
        const content = activeEditor.document.getText(
            new vscode.Range(
                new vscode.Position(visibleRange.start.line, 0),
                new vscode.Position(visibleRange.end.line + 1, 0)
            )
        )
        return {
            fileName: vscode.workspace.asRelativePath(activeEditor.document.uri.fsPath),
            content,
        }
    }

    public async replaceSelection(fileName: string, selectedText: string, replacement: string): Promise<void> {
        const startTime = performance.now()
        const activeEditor = this.getActiveTextEditorInstance() || (await this.fileChatProvider.getEditor())
        if (!activeEditor || vscode.workspace.asRelativePath(activeEditor.document.uri.fsPath) !== fileName) {
            // TODO: should return something indicating success or failure
            console.error('Missing editor')
            return
        }
        let selection = activeEditor.selection
        const taskID = this.fileChatProvider.diffFiles.shift() || ''

        // Stop tracking for file changes to perfotm replacement
        this.fileChatProvider.isInProgress = false
        const chatSelection = this.fileChatProvider.getSelectionRange()

        if (this.fileChatProvider.selectionRange) {
            selection = new vscode.Selection(chatSelection.start, new vscode.Position(chatSelection.end.line + 1, 0))
        }
        if (!selection) {
            console.error('Missing selection')
            return
        }
        if (activeEditor.document.getText(selection) !== selectedText && !this.fileChatProvider.selectionRange) {
            // TODO: Be robust to this.
            await vscode.window.showErrorMessage(
                'The selection changed while Cody was working. The text will not be edited.'
            )
            return
        }
        const trimmedReplacement = replacement
        const startLine = selection.start.line

        // Perform edits
        await activeEditor.edit(edit => {
            edit.delete(selection)
            edit.insert(new vscode.Position(startLine, 0), trimmedReplacement + '\n')
        })

        const lineCount = {
            origin: selectedText.split('\n').length,
            replacement: trimmedReplacement.split('\n').length,
        }

        if (activeEditor.document) {
            // Highlight from the start line to the length of the replacement content
            const newRange = new vscode.Range(startLine, 0, startLine + lineCount.replacement - 1, 0)
            this.fileChatProvider.setReplacementRange(newRange, taskID)
        }

        // check performance time
        console.info('Replacement duration:', performance.now() - startTime)
        return
    }

    public async showQuickPick(labels: string[]): Promise<string | undefined> {
        const label = await vscode.window.showQuickPick(labels)
        return label
    }

    public async showWarningMessage(message: string): Promise<void> {
        await vscode.window.showWarningMessage(message)
    }
}
