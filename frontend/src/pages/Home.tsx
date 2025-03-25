import React, { useEffect, useState } from "react";
import { Container, Typography, Paper, Box } from "@mui/material";
import SideBar from "../components/SideBar";
import MarkdownEditor from "../components/MarkdownEditor";
import { authenticatedApi } from "../api";
import Header from "../components/Header";

// データ型の定義
interface User {
  id: string;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

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

interface UserIndexResponse {
  user: User;
  folderAndNoteTree: FolderNoteTree[];
}

const Home: React.FC = () => {
  const [markdownContent, setMarkdownContent] = useState<string>("");
  const [userData, setUserData] = useState<User | null>(null);
  const [folderTree, setFolderTree] = useState<FolderNoteTree[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [selectedNote, setSelectedNote] = useState<NoteResponse | null>(null);
  const [selectedNoteId, setSelectedNoteId] = useState<string | null>(null);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        setLoading(true);
        const response = await authenticatedApi("GET", "user/index");
        console.log(response);

        if (response && response.user && response.folderAndNoteTree) {
          setUserData(response.user);
          setFolderTree(response.folderAndNoteTree);

          // 初期表示するノートがあれば選択
          if (response.folderAndNoteTree.length > 0) {
            const firstFolderWithNotes = findFirstFolderWithNotes(response.folderAndNoteTree);
            if (firstFolderWithNotes && firstFolderWithNotes.notes.length > 0) {
              const firstNote = firstFolderWithNotes.notes[0];
              setSelectedNote(firstNote);
              setSelectedNoteId(firstNote.id);
              setMarkdownContent(firstNote.content);
            }
          }
        }
        console.log(selectedNote)
      } catch (error) {
        console.error("データの取得に失敗しました:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchUserData();
  }, []);

  // 最初にノートを持つフォルダを見つける補助関数
  const findFirstFolderWithNotes = (folders: FolderNoteTree[]): FolderNoteTree | null => {
    for (const folder of folders) {
      if (folder.notes.length > 0) {
        return folder;
      }
      if (folder.children.length > 0) {
        const childResult = findFirstFolderWithNotes(folder.children);
        if (childResult) {
          return childResult;
        }
      }
    }
    return null;
  };

  const handleSave = (content: string) => {
    setMarkdownContent(content);
    // ここで保存処理を実装（APIを呼び出すなど）
    console.log("保存処理:", content);
  };

  const handleCancel = () => {
    // キャンセル処理
    console.log("キャンセル");
  };

  const handleNoteSelect = (note: NoteResponse) => {
    setSelectedNote(note);
    setSelectedNoteId(note.id);
    setMarkdownContent(note.content);
    console.log(note.content)
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
      <Header />
      <Box sx={{ 
        display: "flex", 
        height: "calc(100vh - 64px)",
        background: "linear-gradient(135deg, #f5f7fa 0%, #e4e8eb 100%)"
      }}>
        {/* サイドバー */}
        <SideBar
          folderTree={folderTree}
          loading={loading}
          onNoteSelect={handleNoteSelect}
          selectedNote={selectedNote}
        />
        
        {/* メインコンテンツ */}
        <Box 
          sx={{ 
            flexGrow: 1, 
            overflow: "auto",
            p: { xs: 2, sm: 3 },
            background: "linear-gradient(135deg, #f5f7fa 0%, #e4e8eb 100%)"
          }}
        >
          <Container maxWidth="lg" sx={{ py: 3 }}>
            <Paper 
              elevation={0} 
              sx={{ 
                p: 4, 
                borderRadius: "16px",
                background: "rgba(255, 255, 255, 0.9)",
                backdropFilter: "blur(10px)",
                boxShadow: "0 8px 32px rgba(0, 0, 0, 0.1)",
                border: "1px solid rgba(255, 255, 255, 0.2)",
                transition: "all 0.3s ease",
                "&:hover": {
                  boxShadow: "0 12px 48px rgba(0, 0, 0, 0.15)",
                  transform: "translateY(-2px)"
                }
              }}
            >
              <Typography 
                variant="h4" 
                component="h1" 
                sx={{ 
                  mb: 4, 
                  fontWeight: 700,
                  textAlign: "center",
                  background: "linear-gradient(45deg, #3f51b5, #2196f3)",
                  backgroundClip: "text",
                  WebkitBackgroundClip: "text",
                  color: "transparent",
                  textShadow: "0 2px 4px rgba(0,0,0,0.1)"
                }}
              >
                {selectedNote ? selectedNote.title : "メモエディタ"}
              </Typography>
              
              <MarkdownEditor
                noteContent={markdownContent}
                onSave={handleSave}
                onCancel={handleCancel}
              />
            </Paper>
          </Container>
        </Box>
      </Box>
    </div>
  );
};

export default Home;
