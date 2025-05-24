# GitHub Copilotのカスタムインストラクション機能

## 概要

GitHub Copilotのカスタムインストラクション機能を使用すると、AIが生成するコードや提案をユーザーの好みやプロジェクトの要件に合わせてカスタマイズすることができます。この機能により、より一貫性があり、プロジェクト固有のコーディング規約に準拠したコードの提案を受けることが可能になります。

## カスタムインストラクションの設定方法

1. GitHub Copilotがインストールされているエディタ（VS Code等）を開きます
2. GitHub Copilotの設定にアクセスします
   - VS Codeの場合: 設定から「GitHub Copilot」を検索
   - JetBrains IDEの場合: 設定から「GitHub Copilot」を選択
3. カスタムインストラクションセクションを見つけ、指示を入力します
4. 変更を保存します

## カスタムインストラクションの効果的な書き方

効果的なカスタムインストラクションを書くためのヒント：

1. **具体的に記述する**: 「良いコードを書いて」ではなく「このプロジェクトではキャメルケースを使用し、関数には説明的な名前を付けてください」など具体的に指示します
2. **優先順位を明確に**: 最も重要な指示を最初に記載します
3. **プロジェクトの文脈を提供**: 開発中のプロジェクトの種類や目的を説明します
4. **コーディング規約を明示**: プロジェクトで使用しているコーディング規約を明確に伝えます
5. **言語や技術スタックを指定**: 使用しているプログラミング言語やフレームワークを明記します
6. **応答の言語を指定**: 例えば「すべての回答は日本語で提供してください」と指定できます

## カスタムインストラクションの例

```
あなたは私のペアプログラマーです。このプロジェクトはTypeScriptとReactを使用したウェブアプリケーションです。
以下のガイドラインに従ってください：

1. コードはTypeScriptの厳格なタイプチェックに準拠すること
2. React関数コンポーネントとフックを使用すること
3. コンポーネント名はPascalCase、変数と関数名はcamelCaseを使用すること
4. コードには適切なコメントを日本語で追加すること
5. すべての提案と説明は日本語で行うこと

プロジェクトは以下のディレクトリ構造に従っています：
- src/components/: すべてのReactコンポーネント
- src/hooks/: カスタムフック
- src/utils/: ユーティリティ関数
- src/types/: 型定義

どうぞよろしくお願いします。
```

## カスタマイズ可能な主な要素

1. **コーディングスタイル**: インデント、命名規則、コメントスタイルなど
2. **アーキテクチャとパターン**: 使用する設計パターンや構造的アプローチ
3. **言語と技術スタック**: 使用するプログラミング言語やフレームワーク
4. **プロジェクト構造**: ファイルやディレクトリの構成
5. **応答言語**: AIからの応答に使用する自然言語（日本語、英語など）
6. **コード品質基準**: パフォーマンス、セキュリティ、テスト可能性などの優先事項

## 注意点

- カスタムインストラクションはプロジェクトごとに異なる場合があります
- 定期的に更新して、プロジェクトの進化に合わせることをお勧めします
- チーム全体で同じカスタムインストラクションを使用することで、一貫性のあるコード生成が可能になります

---

# GitHub Copilot Custom Instructions Feature

## Overview

GitHub Copilot's custom instructions feature allows you to tailor the AI-generated code and suggestions to your preferences and project requirements. This feature enables more consistent code suggestions that adhere to project-specific coding conventions.

## How to Set Up Custom Instructions

1. Open your editor with GitHub Copilot installed (e.g., VS Code)
2. Access GitHub Copilot settings
   - In VS Code: Search for "GitHub Copilot" in settings
   - In JetBrains IDEs: Select "GitHub Copilot" in settings
3. Find the custom instructions section and enter your instructions
4. Save your changes

## Tips for Writing Effective Custom Instructions

1. **Be specific**: Instead of "write good code," say "use camelCase and descriptive function names in this project"
2. **Prioritize**: Put the most important instructions first
3. **Provide context**: Explain the type and purpose of the project you're working on
4. **Specify coding conventions**: Clearly communicate the coding conventions used in your project
5. **Specify language and tech stack**: Mention the programming languages and frameworks you're using
6. **Specify response language**: For example, "provide all responses in Japanese"

## Example of Custom Instructions

```
You are my pair programmer. This project is a web application using TypeScript and React.
Please follow these guidelines:

1. Code should comply with strict TypeScript type checking
2. Use React functional components and hooks
3. Use PascalCase for component names and camelCase for variables and function names
4. Add appropriate comments to the code in Japanese
5. Provide all suggestions and explanations in Japanese

The project follows this directory structure:
- src/components/: All React components
- src/hooks/: Custom hooks
- src/utils/: Utility functions
- src/types/: Type definitions

Thank you for your assistance.
```

## Main Elements That Can Be Customized

1. **Coding style**: Indentation, naming conventions, comment style, etc.
2. **Architecture and patterns**: Design patterns and structural approaches to use
3. **Language and tech stack**: Programming languages and frameworks being used
4. **Project structure**: File and directory organization
5. **Response language**: Natural language for AI responses (Japanese, English, etc.)
6. **Code quality criteria**: Priorities for performance, security, testability, etc.

## Important Notes

- Custom instructions may differ from project to project
- Regular updates are recommended to keep up with project evolution
- Using the same custom instructions across the team ensures consistent code generation