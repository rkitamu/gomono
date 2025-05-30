# gomono - アプリケーション仕様書

## 1. 概要

`gomono`は、指定したGoファイルで使用されているローカルパッケージのソースコードを解析し、依存関係を解決して1つのファイルにマージするCLIツールです。LLMへのコード提供やデバッグ時の全体像把握に便利です。

## 2. 主要機能

### 2.1 パッケージ解析

- 指定されたGoファイルのimport文を解析
- ローカルパッケージ（相対パス・モジュール内パッケージ）を特定
- 依存関係のグラフを構築し、循環参照を検出

### 2.2 ソースコードマージ

- 依存関係順にパッケージをマージ
- パッケージ宣言の統一（`package main`に変更）
- 重複するimport文の除去・統合
- 型定義・関数・変数の競合回避

### 2.3 出力制御

- 標準出力への直接出力
- ファイル出力（`-o`オプション）
- マージ過程の詳細ログ（`-v`オプション）

## 3. コマンドライン仕様

### 3.1 基本構文

```bash
gomono [OPTIONS] <target-file>
```

### 3.2 オプション

| オプション | 短縮形 | 説明 | デフォルト |
|-----------|--------|------|-----------|
| `--output` | `-o` | 出力ファイルパス | 標準出力 |
| `--verbose` | `-v` | 詳細ログ出力 | false |
| `--exclude` | `-e` | 除外パッケージパターン | なし |
| `--dry-run` | `-n` | 実際のマージを行わず解析結果のみ表示 | false |
| `--help` | `-h` | ヘルプ表示 | - |
| `--version` | | バージョン情報表示 | - |

### 3.3 使用例

```bash
# 基本的な使用方法
gomono main.go

# ファイルに出力
gomono -o merged.go main.go

# 詳細ログ付きで実行
gomono -v main.go

# 特定パッケージを除外
gomono -e "internal/test*" main.go

# ドライラン（解析のみ）
gomono -n main.go
```

## 4. 実装詳細

### 4.1 処理フロー

1. **引数解析**: コマンドライン引数の検証・パース
2. **ファイル解析**: 対象Goファイルの構文解析
3. **依存関係解決**: ローカルパッケージの依存関係を再帰的に解析
4. **ソートと検証**: トポロジカルソートで順序決定、循環参照チェック
5. **マージ処理**: パッケージを順次マージし、競合を解決
6. **出力**: マージ結果を指定された形式で出力

### 4.2 マージルール

- **パッケージ宣言**: すべて`package main`に統一
- **import文**: 重複除去し、標準・外部・ローカルの順でグループ化
- **型定義**: 同名型の競合時は`パッケージ名_型名`形式でリネーム
- **関数・変数**: 非公開（小文字開始）はパッケージプレフィックス付与
- **コメント**: パッケージ境界を示すセパレータコメントを挿入

### 4.3 エラーハンドリング

- **ファイル不存在**: 適切なエラーメッセージで終了
- **構文エラー**: 該当行番号と詳細を表示
- **循環参照**: 依存関係パスを表示して警告
- **型競合**: 競合内容と解決方法を通知

## 5. 対応環境

### 5.1 実行環境
- **対応OS**: Windows 10+, macOS 10.15+, Linux (Ubuntu 18.04+)
- **Go Version**: Go 1.19以上
- **実行権限**: 一般ユーザー権限
- **必要ディスク容量**: 50MB以下

### 5.2 セットアップ

```bash
# Go 1.19以上がインストール済みであることを確認
go version

# ツールのインストール
go install github.com/username/gomono@latest

# 実行確認
gomono --version
```

### 5.3 プロジェクト構造対応

- **Go Modules**: ✅ 完全対応（推奨）
- **GOPATH**: ❌ 現在未対応（v2.0で対応予定）
- **Go 1.18+ workspaces**: ❌ 現在未対応（v2.0で対応予定）

## 6. 制限事項

### 6.1 現在の制限

- 標準ライブラリ・外部パッケージは含めない（importのみ保持）
- CGOを使用するパッケージは未対応
- アセンブリファイル（.s）は無視
- ビルドタグ（`// +build`）は考慮しない

### 6.2 将来の拡張予定

- 外部パッケージの選択的マージ機能
- 設定ファイル（`.merger.yaml`）による詳細制御
- マージ結果の最適化（未使用コード除去）
- IDE統合（VS Code拡張など）

## 7. 出力例

### 7.1 標準出力例

```go
// Merged by gomono
// Generated at: 2025-01-30 15:30:45
// Source files: main.go, utils/helper.go, models/user.go

package main

import (
    "fmt"
    "strings"
    
    "github.com/external/package"
)

// ========== main ==========

func main() {
    // ... main function code
}

// ========== utils/helper ==========

func utils_FormatString(s string) string {
    // ... helper function code
}

// ========== models/user ==========

type models_User struct {
    // ... user struct definition
}
```

### 7.2 詳細ログ例（-v使用時）

```
[INFO ] Analyzing target file: main.go
[INFO ] Found local imports: utils/helper, models/user
[INFO ] Building dependency graph...
[INFO ] Dependency order: models/user → utils/helper → main
[DEBUG] Processing package: models/user
[DEBUG] Renaming private function: formatName → models_formatName  
[INFO ] Successfully merged 3 packages (247 lines)
[INFO ] Output written to: stdout
```

## 8. 開発・テスト

### 8.1 開発環境

WIP

### 8.2 テスト方針

WIP
