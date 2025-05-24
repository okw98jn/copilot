# GitHub Copilotカスタムインストラクションの使用ガイド

このガイドでは、GitHub Copilotのカスタムインストラクション機能を使用して、提供されたテンプレートを実際に適用する方法を説明します。

## 目次

1. [VS Codeでの設定方法](#vs-codeでの設定方法)
2. [JetBrains IDEでの設定方法](#jetbrains-ideでの設定方法)
3. [テンプレートの適用例](#テンプレートの適用例)
4. [カスタムインストラクションの効果の確認方法](#カスタムインストラクションの効果の確認方法)

## VS Codeでの設定方法

1. **VS Codeを開く**

2. **GitHub Copilotの設定にアクセスする**
   - VS Codeの左下の歯車アイコンをクリックして「設定」を開く
   - もしくはキーボードショートカット `Ctrl+,`（Windows/Linux）または `Cmd+,`（Mac）を使用

3. **GitHub Copilotの設定を検索**
   - 設定の検索バーに「GitHub Copilot」と入力

4. **カスタムインストラクションの設定を見つける**
   - 「GitHub Copilot: Enable Chat Custom Instructions」の項目を見つけてチェックを入れる

5. **カスタムインストラクションを入力**
   - 「GitHub Copilot: Chat Custom Instructions」の項目をクリックして編集
   - このリポジトリのテンプレートから選んだテキストを貼り付ける

6. **設定を保存**
   - 自動的に保存されますが、設定ページを閉じるとよいでしょう

## JetBrains IDEでの設定方法

1. **JetBrains IDE（IntelliJ、PyCharm、WebStormなど）を開く**

2. **設定にアクセス**
   - Windows/Linuxの場合: `Ctrl+Alt+S`、または「File」→「Settings」
   - Macの場合: `Cmd+,`、または「IntelliJ IDEA」→「Preferences」

3. **プラグインの設定に移動**
   - 左側のナビゲーションから「Plugins」を選択
   - インストールされたプラグインから「GitHub Copilot」を見つける

4. **カスタムインストラクションの設定**
   - 「GitHub Copilot」の設定ページを開く
   - 「Custom Instructions」セクションを見つける
   - このリポジトリのテンプレートからコピーしたテキストを貼り付ける

5. **設定を保存**
   - 「Apply」または「OK」ボタンをクリックして設定を保存

## テンプレートの適用例

以下は、このリポジトリの汎用テンプレートを適用する具体的な手順です。

### 例: 汎用テンプレートの適用

1. [templates/general.md](/templates/general.md) ファイルを開く

2. テンプレート部分（コードブロックで囲まれた部分）をコピー

```
あなたは私のプログラミングアシスタントです。以下のガイドラインに従ってください：

## 応答形式
- すべての回答は日本語で提供してください
- 簡潔かつ明確な説明を心がけてください
- コードには適切なコメントを含めてください
- 可能な限り具体例を示してください

## コーディングスタイル
- 変数・関数名は説明的で、目的が明確になるようにしてください
- 一貫した命名規則を使用してください（例：camelCaseまたはsnake_case）
- 適切なインデントとフォーマットを守ってください
- 関数やメソッドには適切なドキュメンテーションコメントを追加してください

## コード品質
- シンプルで読みやすいコードを提案してください
- DRY原則（Don't Repeat Yourself）を守ってください
- エラーハンドリングを適切に行ってください
- セキュリティのベストプラクティスに従ってください

## その他の指示
- [ここにプロジェクト固有の要件や追加指示を記入してください]

どうぞよろしくお願いします。
```

3. プロジェクトの要件に合わせて内容を編集（例: その他の指示の部分を具体的に記載）

4. VS Codeの設定で「GitHub Copilot: Chat Custom Instructions」に編集済みのテキストを貼り付ける

5. 設定を保存する

## カスタムインストラクションの効果の確認方法

カスタムインストラクションが正しく適用されているか確認するには：

1. **VS CodeでGitHub Copilot Chatを開く**
   - キーボードショートカット `Ctrl+Shift+I`（Windows/Linux）または `Cmd+Shift+I`（Mac）
   - もしくは、VS Codeの左側のアクティビティバーでCopilotアイコンをクリック

2. **初期メッセージを確認**
   - チャットが開くと、Copilotがあなたの指示に基づいた初期メッセージを表示します
   - このメッセージで指定した言語（日本語など）や形式が使われているか確認

3. **テスト質問をする**
   - 例: 「簡単なカウンター関数を作成してください」など、プロジェクトに関連する質問をします
   - 返答が指定したスタイルやガイドラインに従っているか確認します

4. **フィードバックと調整**
   - 必要に応じてカスタムインストラクションを調整して再度試してみましょう

---

# GitHub Copilot Custom Instructions Usage Guide

This guide explains how to apply the custom instruction templates provided in this repository to GitHub Copilot.

## Table of Contents

1. [Setting up in VS Code](#setting-up-in-vs-code)
2. [Setting up in JetBrains IDEs](#setting-up-in-jetbrains-ides)
3. [Template Application Example](#template-application-example)
4. [Verifying Custom Instructions Effect](#verifying-custom-instructions-effect)

## Setting up in VS Code

1. **Open VS Code**

2. **Access GitHub Copilot settings**
   - Click the gear icon in the bottom left to open Settings
   - Or use the keyboard shortcut `Ctrl+,` (Windows/Linux) or `Cmd+,` (Mac)

3. **Search for GitHub Copilot settings**
   - Type "GitHub Copilot" in the settings search bar

4. **Find custom instructions settings**
   - Find and check the "GitHub Copilot: Enable Chat Custom Instructions" option

5. **Enter custom instructions**
   - Click on "GitHub Copilot: Chat Custom Instructions" to edit
   - Paste the text from your chosen template in this repository

6. **Save settings**
   - Settings will be saved automatically, but you can close the settings page

## Setting up in JetBrains IDEs

1. **Open your JetBrains IDE** (IntelliJ, PyCharm, WebStorm, etc.)

2. **Access settings**
   - Windows/Linux: `Ctrl+Alt+S` or "File" → "Settings"
   - Mac: `Cmd+,` or "IntelliJ IDEA" → "Preferences"

3. **Navigate to plugins settings**
   - Select "Plugins" from the left navigation
   - Find "GitHub Copilot" among installed plugins

4. **Configure custom instructions**
   - Open the "GitHub Copilot" settings page
   - Find the "Custom Instructions" section
   - Paste the text copied from one of the templates in this repository

5. **Save settings**
   - Click "Apply" or "OK" to save your settings

## Template Application Example

Here's a specific example of applying a general template from this repository:

### Example: Applying the General Template

1. Open the [templates/general.md](/templates/general.md) file

2. Copy the template section (inside the code block)

```
You are my programming assistant. Please follow these guidelines:

## Response Format
- Provide all responses in Japanese
- Keep explanations concise and clear
- Include appropriate comments in code
- Provide examples wherever possible

## Coding Style
- Use descriptive variable and function names that clearly indicate their purpose
- Use consistent naming conventions (e.g., camelCase or snake_case)
- Maintain proper indentation and formatting
- Add appropriate documentation comments for functions and methods

## Code Quality
- Suggest simple and readable code
- Follow the DRY (Don't Repeat Yourself) principle
- Implement proper error handling
- Follow security best practices

## Additional Instructions
- [Insert project-specific requirements or additional instructions here]

Thank you for your assistance.
```

3. Edit the content to match your project needs (e.g., fill in the Additional Instructions section)

4. Paste the edited text into "GitHub Copilot: Chat Custom Instructions" in VS Code settings

5. Save your settings

## Verifying Custom Instructions Effect

To verify your custom instructions are properly applied:

1. **Open GitHub Copilot Chat in VS Code**
   - Keyboard shortcut `Ctrl+Shift+I` (Windows/Linux) or `Cmd+Shift+I` (Mac)
   - Or click the Copilot icon in the activity bar on the left side of VS Code

2. **Check the initial message**
   - When the chat opens, Copilot will show an initial message based on your instructions
   - Verify that it uses your specified language (e.g., Japanese) and format

3. **Ask a test question**
   - For example: "Create a simple counter function" or something relevant to your project
   - Check if the response follows your specified style and guidelines

4. **Feedback and adjustments**
   - Adjust your custom instructions as needed and try again