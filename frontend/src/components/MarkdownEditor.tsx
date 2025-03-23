import React, { useState } from "react";
import { Box, Button, TextField, Tabs, Tab, Paper, Typography } from "@mui/material";
// @ts-ignore
import ReactMarkdown from 'react-markdown';
// @ts-ignore
import remarkGfm from 'remark-gfm';
// @ts-ignore
import rehypeRaw from 'rehype-raw';
// @ts-ignore
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
// @ts-ignore
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';
// CSSスタイル
import '../pages/markdown.css';

interface MarkdownEditorProps {
  initialContent?: string;
  onSave?: (content: string) => void;
  onCancel?: () => void;
}

const MarkdownEditor: React.FC<MarkdownEditorProps> = ({
  initialContent,
  onSave,
  onCancel
}) => {
  const [markdownContent, setMarkdownContent] = useState<string>(
    initialContent || 
`# マークダウンエディタ

ここに**マークダウン**を入力できます。

## 見出し

### 小見出し

## リスト

- 項目1
- 項目2
  - ネストした項目
  - もう一つのネスト項目
- 項目3

## 番号付きリスト

1. 最初の項目
2. 2番目の項目
3. 3番目の項目

## チェックボックス

- [x] 完了したタスク
- [ ] 未完了のタスク

## コード

インラインコード: \`const message = "Hello World";\`

\`\`\`javascript
// コードブロック
function hello() {
  console.log("Hello, world!");
}
\`\`\`

## テーブル

| 名前 | 年齢 | 職業 |
|------|------|------|
| 田中 | 28   | エンジニア |
| 佐藤 | 32   | デザイナー |

## 引用

> これは引用です。
> 複数行にまたがります。

## 水平線

---

## リンク

[Google](https://www.google.com)

## 画像

画像も表示できます（URLを入力）

## HTML埋め込み

<div style="padding: 10px; background-color: #f0f0f0; border-radius: 5px;">
  HTMLタグも使えます
</div>
`);
  const [tabValue, setTabValue] = useState<number>(0);

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  const handleSave = () => {
    if (onSave) {
      onSave(markdownContent);
    }
  };

  const handleCancel = () => {
    if (onCancel) {
      onCancel();
    }
  };

  // コードブロックのカスタムレンダラー
  const components = {
    h1: ({ node, children, ...props }: any) => (
      <Typography variant="h3" component="h1" sx={{ my: 2, fontWeight: 600 }} {...props}>
        {children}
      </Typography>
    ),
    h2: ({ node, children, ...props }: any) => (
      <Typography variant="h4" component="h2" sx={{ my: 2, fontWeight: 600 }} {...props}>
        {children}
      </Typography>
    ),
    h3: ({ node, children, ...props }: any) => (
      <Typography variant="h5" component="h3" sx={{ my: 1.5, fontWeight: 600 }} {...props}>
        {children}
      </Typography>
    ),
    h4: ({ node, children, ...props }: any) => (
      <Typography variant="h6" component="h4" sx={{ my: 1.5, fontWeight: 600 }} {...props}>
        {children}
      </Typography>
    ),
    p: ({ node, children, ...props }: any) => (
      <Typography variant="body1" sx={{ my: 1.5 }} {...props}>
        {children}
      </Typography>
    ),
    a: ({ node, children, ...props }: any) => (
      <a {...props} style={{ color: '#3f51b5', textDecoration: 'none' }}>
        {children}
      </a>
    ),
    blockquote: ({ node, children, ...props }: any) => (
      <Box 
        component="blockquote" 
        sx={{ 
          borderLeft: '4px solid #3f51b5', 
          pl: 2, 
          py: 0.5, 
          my: 2,
          color: 'text.secondary',
          bgcolor: 'rgba(63, 81, 181, 0.05)',
          borderRadius: '0 4px 4px 0'
        }}
        {...props}
      >
        {children}
      </Box>
    ),
    table: ({ node, children, ...props }: any) => (
      <Box sx={{ overflowX: 'auto', my: 2 }}>
        <table style={{ borderCollapse: 'collapse', width: '100%' }} {...props}>
          {children}
        </table>
      </Box>
    ),
    thead: ({ node, children, ...props }: any) => (
      <thead style={{ backgroundColor: '#f5f5f5' }} {...props}>
        {children}
      </thead>
    ),
    th: ({ node, children, ...props }: any) => (
      <th style={{ padding: '12px 16px', textAlign: 'left', borderBottom: '2px solid #e0e0e0' }} {...props}>
        {children}
      </th>
    ),
    td: ({ node, children, ...props }: any) => (
      <td style={{ padding: '8px 16px', borderBottom: '1px solid #e0e0e0' }} {...props}>
        {children}
      </td>
    ),
    hr: ({ node, ...props }: any) => (
      <hr style={{ border: 'none', height: '1px', backgroundColor: '#e0e0e0', margin: '24px 0' }} {...props} />
    ),
    ol: ({ node, children, ...props }: any) => (
      <ol className="ordered-list" style={{ 
        paddingLeft: '24px', 
        margin: '16px 0',
        listStyleType: 'decimal',
        display: 'block'
      }} {...props}>
        {children}
      </ol>
    ),
    ul: ({ node, children, ...props }: any) => (
      <ul className="list-disc pl-6 my-4" style={{ 
        color: '#000',
        listStyleType: 'disc',
        display: 'block'
      }} {...props}>
        {children}
      </ul>
    ),
    li: ({ node, children, ordered, ...props }: any) => {
      // 親要素がolかulかを確認
      const isOrderedList = node.parent && node.parent.tagName === 'ol';
      return (
        <li 
          className={isOrderedList ? "ordered-list-item" : "unordered-list-item"} 
          style={{ 
            display: 'list-item',
            margin: '8px 0',
            listStyleType: isOrderedList ? 'decimal' : 'disc'
          }} 
          {...props}
        >
          {children}
        </li>
      );
    },
    code: ({ node, inline, className, children, ...props }: any) => {
      const match = /language-(\w+)/.exec(className || '');
      return !inline && match ? (
        <Box sx={{ my: 2, borderRadius: '8px', overflow: 'hidden' }}>
          <SyntaxHighlighter
            style={vscDarkPlus}
            language={match[1]}
            PreTag="div"
            {...props}
          >
            {String(children).replace(/\n$/, '')}
          </SyntaxHighlighter>
        </Box>
      ) : (
        <code 
          className={className} 
          style={{ 
            backgroundColor: '#f5f5f5', 
            padding: '2px 6px', 
            borderRadius: '4px',
            fontFamily: 'monospace',
            fontSize: '0.9em',
            color: '#e53935'
          }} 
          {...props}
        >
          {children}
        </code>
      );
    }
  };

  // マークダウンレンダリング共通コンポーネント
  const MarkdownRenderer = ({ content }: { content: string }) => (
    <Box className="markdown-body">
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[rehypeRaw]}
        components={components}
      >
        {content}
      </ReactMarkdown>
    </Box>
  );

  return (
    <>
      {/* タブ切り替え */}
      <Tabs 
        value={tabValue} 
        onChange={handleTabChange} 
        sx={{ mb: 2 }}
        variant="fullWidth"
      >
        <Tab label="編集" />
        <Tab label="プレビュー" />
        <Tab label="分割表示" />
      </Tabs>

      <Box sx={{ mb: 3 }}>
        {tabValue === 0 && (
          <TextField
            fullWidth
            multiline
            minRows={20}
            maxRows={40}
            value={markdownContent}
            onChange={(e) => setMarkdownContent(e.target.value)}
            variant="outlined"
            placeholder="マークダウンを入力してください..."
            sx={{
              fontFamily: 'monospace',
              '& .MuiOutlinedInput-root': {
                borderRadius: '8px',
              }
            }}
          />
        )}

        {tabValue === 1 && (
          <Paper
            elevation={0}
            sx={{
              p: 3,
              minHeight: '600px',
              maxHeight: '800px',
              border: '1px solid #e0e0e0',
              borderRadius: '8px',
              overflowY: 'auto'
            }}
          >
            <MarkdownRenderer content={markdownContent} />
          </Paper>
        )}

        {tabValue === 2 && (
          <Box sx={{ display: 'flex', gap: 2 }}>
            <TextField
              fullWidth
              multiline
              minRows={20}
              maxRows={40}
              value={markdownContent}
              onChange={(e) => setMarkdownContent(e.target.value)}
              variant="outlined"
              placeholder="マークダウンを入力してください..."
              sx={{ 
                width: '50%',
                fontFamily: 'monospace' 
              }}
            />
            <Paper
              elevation={0}
              sx={{
                p: 3,
                width: '50%',
                minHeight: '600px',
                maxHeight: '800px',
                border: '1px solid #e0e0e0',
                borderRadius: '8px',
                overflowY: 'auto'
              }}
            >
              <MarkdownRenderer content={markdownContent} />
            </Paper>
          </Box>
        )}
      </Box>

      <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
        <Button 
          variant="outlined" 
          color="primary"
          onClick={handleCancel}
          sx={{
            borderRadius: "28px",
            px: 3,
            py: 1,
            borderWidth: 2,
            '&:hover': {
              borderWidth: 2
            }
          }}
        >
          キャンセル
        </Button>

        <Button 
          variant="contained" 
          color="primary"
          onClick={handleSave}
          sx={{
            borderRadius: "28px",
            px: 3,
            py: 1,
            boxShadow: '0 4px 10px rgba(63, 81, 181, 0.2)',
            '&:hover': {
              boxShadow: '0 6px 12px rgba(63, 81, 181, 0.3)'
            }
          }}
        >
          保存
        </Button>
      </Box>
    </>
  );
};

export default MarkdownEditor; 