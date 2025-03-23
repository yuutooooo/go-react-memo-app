import React, { useState, useEffect } from "react";
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
  CircularProgress,
  Drawer,
  useMediaQuery,
  useTheme,
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
  Menu as MenuIcon,
} from "@mui/icons-material";

// インターフェース定義
interface NoteResponse {
  id: string;
  title: string;
  content: string;
  folder_id: string;
  created_at: string;
  updated_at: string;
}

interface FolderResponse {
  id: string;
  name: string;
  path: string;
  parentFolderID: string;
  createdAt: string;
  updatedAt: string;
}

interface FolderNoteTree {
  folder: FolderResponse;
  notes: NoteResponse[];
  children: FolderNoteTree[];
}

interface SideBarProps {
  folderTree: FolderNoteTree[];
  loading: boolean;
  onNoteSelect: (note: NoteResponse) => void;
  selectedNote: NoteResponse | null;
}

const SideBar: React.FC<SideBarProps> = ({ 
  folderTree, 
  loading, 
  onNoteSelect,
  selectedNote
}) => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  
  // 状態管理
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({});
  const [selectedNodeId, setSelectedNodeId] = useState<string | null>(null);
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [createType, setCreateType] = useState<"file" | "folder">("file");
  const [newItemName, setNewItemName] = useState("");
  const [selectedFolder, setSelectedFolder] = useState<FolderResponse | null>(null);
  const [mobileOpen, setMobileOpen] = useState(false);
  
  // 初期状態でいくつかのフォルダを開いておく
  useEffect(() => {
    if (folderTree.length > 0) {
      const initialExpanded: Record<string, boolean> = {};
      // ルートフォルダを展開
      folderTree.forEach(item => {
        initialExpanded[item.folder.id] = true;
        
        // ノートを持つフォルダを展開
        if (item.notes.length > 0) {
          initialExpanded[item.folder.id] = true;
        }
      });
      setExpandedNodes(initialExpanded);
    }
  }, [folderTree]);

  // ノードを展開/折りたたむ関数
  const toggleNodeExpansion = (folderId: string, e: React.MouseEvent) => {
    e.stopPropagation();
    setExpandedNodes(prev => ({
      ...prev,
      [folderId]: !prev[folderId]
    }));
  };

  // メニューを開く関数
  const handleMenuOpen = (folder: FolderResponse, e: React.MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    setSelectedFolder(folder);
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

  // フォルダをクリックしたときの関数
  const handleFolderClick = (folder: FolderResponse, e: React.MouseEvent) => {
    toggleNodeExpansion(folder.id, e);
    setSelectedNodeId(folder.id);
  };

  // ノートをクリックしたときの関数
  const handleNoteClick = (note: NoteResponse) => {
    setSelectedNodeId(note.id);
    onNoteSelect(note);
    
    // モバイル表示の場合はサイドバーを閉じる
    if (isMobile) {
      setMobileOpen(false);
    }
  };

  // モバイル表示のトグル
  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  // フォルダとノートのツリーを再帰的にレンダリングする関数
  const renderFolderTree = (items: FolderNoteTree[], level: number = 0) => {
    return items.map(item => (
      <React.Fragment key={item.folder.id}>
        {/* フォルダ */}
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            p: 1,
            pl: 1 + level * 2,
            cursor: "pointer",
            borderRadius: "4px",
            transition: "background-color 0.2s",
            "&:hover": {
              backgroundColor: "rgba(0, 0, 0, 0.04)",
            },
            ...(selectedNodeId === item.folder.id && {
              backgroundColor: "rgba(63, 81, 181, 0.08)",
            }),
          }}
          onClick={(e) => handleFolderClick(item.folder, e)}
        >
          <IconButton
            size="small"
            onClick={(e) => toggleNodeExpansion(item.folder.id, e)}
            sx={{ mr: 0.5 }}
          >
            {expandedNodes[item.folder.id] ? (
              <CollapseIcon fontSize="small" />
            ) : (
              <ExpandIcon fontSize="small" />
            )}
          </IconButton>
          
          <FolderIcon sx={{ mr: 1, color: "#3f51b5" }} />
          
          <Typography
            variant="body2"
            sx={{
              flexGrow: 1,
              overflow: "hidden",
              textOverflow: "ellipsis",
              whiteSpace: "nowrap",
              fontWeight: item.notes.length > 0 ? 600 : 400,
            }}
          >
            {item.folder.name}
            {item.notes.length > 0 && (
              <Typography component="span" sx={{ ml: 1, color: "text.secondary", fontSize: '0.8rem' }}>
                ({item.notes.length})
              </Typography>
            )}
          </Typography>
          
          <Tooltip title="メニュー">
            <IconButton
              size="small"
              onClick={(e) => handleMenuOpen(item.folder, e)}
            >
              <MoreIcon fontSize="small" />
            </IconButton>
          </Tooltip>
        </Box>

        {/* 展開されていれば子要素を表示 */}
        {expandedNodes[item.folder.id] && (
          <>
            {/* ノート */}
            {item.notes.map(note => (
              <Box
                key={note.id}
                sx={{
                  display: "flex",
                  alignItems: "center",
                  p: 1,
                  pl: 4 + level * 2,
                  cursor: "pointer",
                  borderRadius: "4px",
                  transition: "background-color 0.2s",
                  "&:hover": {
                    backgroundColor: "rgba(0, 0, 0, 0.04)",
                  },
                  ...(selectedNodeId === note.id && {
                    backgroundColor: "rgba(63, 81, 181, 0.08)",
                  }),
                }}
                onClick={() => handleNoteClick(note)}
              >
                <FileIcon sx={{ mr: 1, color: "#607d8b" }} />
                <Typography
                  variant="body2"
                  sx={{
                    flexGrow: 1,
                    overflow: "hidden",
                    textOverflow: "ellipsis",
                    whiteSpace: "nowrap",
                  }}
                >
                  {note.title}
                </Typography>
              </Box>
            ))}
            
            {/* 子フォルダ */}
            {item.children.length > 0 && renderFolderTree(item.children, level + 1)}
          </>
        )}
      </React.Fragment>
    ));
  };

  const sidebarContent = (
    <>
      <Box sx={{ p: 2 }}>
        <Typography variant="h6" component="h2" sx={{ fontWeight: 600, mb: 1 }}>
          フォルダとノート
        </Typography>
        
        <Divider sx={{ mb: 2 }} />
        
        {/* アクションボタン */}
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-around",
            mb: 2,
          }}
        >
          <Button
            size="small"
            startIcon={<NewFolderIcon />}
            variant="outlined"
            onClick={() => {
              if (folderTree.length > 0) {
                setSelectedFolder(folderTree[0].folder);
                handleCreateDialogOpen("folder");
              }
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
              if (folderTree.length > 0) {
                setSelectedFolder(folderTree[0].folder);
                handleCreateDialogOpen("file");
              }
            }}
            sx={{ borderRadius: '20px' }}
          >
            新規ノート
          </Button>
        </Box>
        
        <Divider sx={{ mb: 2 }} />
      </Box>

      {/* フォルダ構造 */}
      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
          <CircularProgress />
        </Box>
      ) : folderTree.length === 0 ? (
        <Box sx={{ p: 2, textAlign: 'center' }}>
          <Typography variant="body2" color="text.secondary">
            フォルダが見つかりません
          </Typography>
        </Box>
      ) : (
        <Box sx={{ overflow: 'auto', maxHeight: 'calc(100vh - 220px)' }}>
          {renderFolderTree(folderTree)}
        </Box>
      )}

      {/* 新規作成ダイアログ */}
      <Dialog
        open={createDialogOpen}
        onClose={() => setCreateDialogOpen(false)}
        fullWidth
        maxWidth="xs"
      >
        <DialogTitle>
          {createType === "folder" ? "新規フォルダの作成" : "新規ノートの作成"}
        </DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label={createType === "folder" ? "フォルダ名" : "ノートタイトル"}
            fullWidth
            variant="outlined"
            value={newItemName}
            onChange={(e) => setNewItemName(e.target.value)}
            placeholder={createType === "folder" ? "新しいフォルダ" : "新しいノート"}
            sx={{ mt: 1 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setCreateDialogOpen(false)}>キャンセル</Button>
          <Button
            onClick={() => {
              // TODO: APIを呼び出して新規作成処理を行う
              console.log("Create", createType, newItemName, selectedFolder);
              setCreateDialogOpen(false);
            }}
            variant="contained"
            color="primary"
            disabled={!newItemName.trim()}
          >
            作成
          </Button>
        </DialogActions>
      </Dialog>

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
          ノート作成
        </MenuItem>
        <Divider />
        <MenuItem onClick={handleMenuClose} sx={{ color: "error.main" }}>
          <DeleteIcon fontSize="small" sx={{ mr: 1 }} />
          削除
        </MenuItem>
      </Menu>
    </>
  );

  // レスポンシブ対応
  if (isMobile) {
    return (
      <>
        <IconButton
          color="inherit"
          aria-label="メニューを開く"
          edge="start"
          onClick={handleDrawerToggle}
          sx={{ 
            position: 'fixed', 
            left: 16, 
            bottom: 16, 
            zIndex: 1200, 
            bgcolor: 'primary.main',
            color: 'white',
            '&:hover': {
              bgcolor: 'primary.dark',
            },
            boxShadow: 3
          }}
        >
          <MenuIcon />
        </IconButton>
        
        <Drawer
          variant="temporary"
          open={mobileOpen}
          onClose={handleDrawerToggle}
          sx={{
            '& .MuiDrawer-paper': { 
              width: 280,
              boxSizing: 'border-box',
            },
          }}
        >
          {sidebarContent}
        </Drawer>
      </>
    );
  }

  return (
    <Paper
      elevation={0}
      sx={{
        width: 280,
        height: "100%",
        borderRight: "1px solid #e0e0e0",
        overflow: "hidden",
        display: "flex",
        flexDirection: "column",
      }}
    >
      {sidebarContent}
    </Paper>
  );
};

export default SideBar;
