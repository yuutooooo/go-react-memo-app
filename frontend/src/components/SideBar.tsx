import React, { useState } from "react";
import {
  Box,
  Typography,
  IconButton,
  Menu,
  MenuItem,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  Divider,
  Paper,
  Tooltip,
} from "@mui/material";
import {
  FolderOutlined as FolderIcon,
  DescriptionOutlined as FileIcon,
  MoreVert as MoreIcon,
  CreateNewFolder as NewFolderIcon,
  NoteAdd as NewFileIcon,
  ChevronRight as ExpandIcon,
  ExpandMore as CollapseIcon,
  Delete as DeleteIcon,
} from "@mui/icons-material";

// ディレクトリ構造のインターフェース
interface FileNode {
  id: string;
  name: string;
  type: "file" | "folder";
  children?: FileNode[];
}

// モックデータ - 実際のアプリケーションではAPIから取得します
const initialFileStructure: FileNode[] = [
  {
    id: "1",
    name: "プロジェクト",
    type: "folder",
    children: [
      {
        id: "2",
        name: "ドキュメント",
        type: "folder",
        children: [
          { id: "3", name: "仕様書.md", type: "file" },
          { id: "4", name: "議事録.md", type: "file" },
        ],
      },
      {
        id: "5",
        name: "ソースコード",
        type: "folder",
        children: [
          { id: "6", name: "app.js", type: "file" },
          { id: "7", name: "utils.js", type: "file" },
          {
            id: "8",
            name: "components",
            type: "folder",
            children: [
              { id: "9", name: "Button.jsx", type: "file" },
              { id: "10", name: "Modal.jsx", type: "file" },
            ],
          },
        ],
      },
      { id: "11", name: "README.md", type: "file" },
    ],
  },
  {
    id: "12",
    name: "個人メモ",
    type: "folder",
    children: [
      { id: "13", name: "アイデア.md", type: "file" },
      { id: "14", name: "タスク.md", type: "file" },
    ],
  },
];

const SideBar: React.FC = () => {
  // 状態管理
  const [fileStructure, setFileStructure] = useState<FileNode[]>(initialFileStructure);
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({});
  const [selectedNode, setSelectedNode] = useState<FileNode | null>(null);
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [createType, setCreateType] = useState<"file" | "folder">("file");
  const [newItemName, setNewItemName] = useState("");
  const [currentParentId, setCurrentParentId] = useState<string | null>(null);

  // ノードを展開/折りたたむ関数
  const toggleNodeExpansion = (nodeId: string, e: React.MouseEvent) => {
    e.stopPropagation();
    setExpandedNodes((prev) => ({
      ...prev,
      [nodeId]: !prev[nodeId],
    }));
  };

  // メニューを開く関数
  const handleMenuOpen = (
    node: FileNode,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.stopPropagation();
    setSelectedNode(node);
    setAnchorEl(e.currentTarget);
  };

  // メニューを閉じる関数
  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  // ダイアログを開く関数
  const handleCreateDialogOpen = (type: "file" | "folder") => {
    setCreateType(type);
    setNewItemName("");
    setCreateDialogOpen(true);
    handleMenuClose();
  };

  // ノードをクリックしたときの関数
  const handleNodeClick = (node: FileNode) => {
    if (node.type === "folder") {
      toggleNodeExpansion(node.id, { stopPropagation: () => {} } as React.MouseEvent);
    } else {
      // ファイルを選択した場合の処理
      console.log("ファイルを選択:", node.name);
    }
  };

  // 新しいファイル/フォルダを追加する関数
  const handleCreateItem = () => {
    if (!newItemName.trim() || !selectedNode) {
      return;
    }

    const newId = Date.now().toString();
    const newItem: FileNode = {
      id: newId,
      name: newItemName,
      type: createType,
      ...(createType === "folder" && { children: [] }),
    };

    // フォルダ構造を更新する関数
    const updateFolderStructure = (nodes: FileNode[]): FileNode[] => {
      return nodes.map((node) => {
        if (node.id === selectedNode.id && node.type === "folder") {
          return {
            ...node,
            children: [...(node.children || []), newItem],
          };
        } else if (node.children) {
          return {
            ...node,
            children: updateFolderStructure(node.children),
          };
        }
        return node;
      });
    };

    setFileStructure(updateFolderStructure(fileStructure));
    setCreateDialogOpen(false);
    
    // 親フォルダを展開
    setExpandedNodes((prev) => ({
      ...prev,
      [selectedNode.id]: true,
    }));
  };

  // フォルダ構造を再帰的にレンダリングする関数
  const renderFileStructure = (nodes: FileNode[], level: number = 0) => {
    return nodes.map((node) => (
      <Box key={node.id}>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            p: 1,
            pl: 1 + level * 2,
            cursor: "pointer",
            borderRadius: "4px",
            "&:hover": {
              backgroundColor: "rgba(0, 0, 0, 0.04)",
            },
            ...(selectedNode?.id === node.id && {
              backgroundColor: "rgba(63, 81, 181, 0.08)",
            }),
          }}
          onClick={() => handleNodeClick(node)}
        >
          {node.type === "folder" && (
            <IconButton
              size="small"
              onClick={(e) => toggleNodeExpansion(node.id, e)}
              sx={{ mr: 0.5 }}
            >
              {expandedNodes[node.id] ? (
                <CollapseIcon fontSize="small" />
              ) : (
                <ExpandIcon fontSize="small" />
              )}
            </IconButton>
          )}
          {node.type === "folder" ? (
            <FolderIcon sx={{ mr: 1, color: "#3f51b5" }} />
          ) : (
            <FileIcon sx={{ mr: 1, color: "#607d8b" }} />
          )}
          <Typography
            variant="body2"
            sx={{
              flexGrow: 1,
              overflow: "hidden",
              textOverflow: "ellipsis",
              whiteSpace: "nowrap",
            }}
          >
            {node.name}
          </Typography>
          {node.type === "folder" && (
            <Tooltip title="メニュー">
              <IconButton
                size="small"
                onClick={(e) => handleMenuOpen(node, e)}
              >
                <MoreIcon fontSize="small" />
              </IconButton>
            </Tooltip>
          )}
        </Box>
        {node.children &&
          expandedNodes[node.id] &&
          renderFileStructure(node.children, level + 1)}
      </Box>
    ));
  };

  return (
    <Paper
      elevation={0}
      sx={{
        width: 280,
        height: "100%",
        borderRight: "1px solid #e0e0e0",
        overflow: "auto",
        p: 1,
      }}
    >
      <Box sx={{ p: 1, mb: 1 }}>
        <Typography variant="h6" component="h2" sx={{ fontWeight: 600 }}>
          フォルダ
        </Typography>
      </Box>

      <Divider sx={{ mb: 2 }} />

      {/* ルートレベルのアクション */}
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-around",
          mb: 2,
          px: 1,
        }}
      >
        <Button
          size="small"
          startIcon={<NewFolderIcon />}
          variant="outlined"
          onClick={() => {
            setSelectedNode(fileStructure[0]);
            handleCreateDialogOpen("folder");
          }}
          sx={{ borderRadius: '20px' }}
        >
          新規フォルダ
        </Button>
        <Button
          size="small"
          startIcon={<NewFileIcon />}
          variant="outlined"
          onClick={() => {
            setSelectedNode(fileStructure[0]);
            handleCreateDialogOpen("file");
          }}
          sx={{ borderRadius: '20px' }}
        >
          新規ファイル
        </Button>
      </Box>

      <Divider sx={{ mb: 2 }} />

      {/* ファイル構造の表示 */}
      <Box sx={{ mb: 2 }}>{renderFileStructure(fileStructure)}</Box>

      {/* フォルダ操作メニュー */}
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleMenuClose}
      >
        <MenuItem
          onClick={() => handleCreateDialogOpen("folder")}
          sx={{ minWidth: 150 }}
        >
          <NewFolderIcon fontSize="small" sx={{ mr: 1 }} />
          フォルダ作成
        </MenuItem>
        <MenuItem onClick={() => handleCreateDialogOpen("file")}>
          <NewFileIcon fontSize="small" sx={{ mr: 1 }} />
          ファイル作成
        </MenuItem>
        <Divider />
        <MenuItem onClick={handleMenuClose} sx={{ color: "error.main" }}>
          <DeleteIcon fontSize="small" sx={{ mr: 1 }} />
          削除
        </MenuItem>
      </Menu>

      {/* 新規作成ダイアログ */}
      <Dialog
        open={createDialogOpen}
        onClose={() => setCreateDialogOpen(false)}
        fullWidth
        maxWidth="xs"
      >
        <DialogTitle>
          {createType === "folder" ? "新規フォルダの作成" : "新規ファイルの作成"}
        </DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label={createType === "folder" ? "フォルダ名" : "ファイル名"}
            fullWidth
            variant="outlined"
            value={newItemName}
            onChange={(e) => setNewItemName(e.target.value)}
            placeholder={createType === "folder" ? "新しいフォルダ" : "新しいファイル.md"}
            sx={{ mt: 1 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setCreateDialogOpen(false)}>キャンセル</Button>
          <Button
            onClick={handleCreateItem}
            variant="contained"
            color="primary"
            disabled={!newItemName.trim()}
          >
            作成
          </Button>
        </DialogActions>
      </Dialog>
    </Paper>
  );
};

export default SideBar;
