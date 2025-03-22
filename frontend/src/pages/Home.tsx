import React, { useEffect, useState } from "react";
import { Container, Typography, Paper, Box } from "@mui/material";
import SideBar from "../components/SideBar";
import MarkdownEditor from "../components/MarkdownEditor";
import { authenticatedApi } from "../api";

const Home: React.FC = () => {
  const [markdownContent, setMarkdownContent] = useState<string>("");
  // const [user, setUser] = useState<User | null>(null);
  // const [folders, setFolders] = useState<Folder[]>([]);

  useEffect(() => {
    const fetchUser = async () => {
      const response = await authenticatedApi("GET", "user/index");
      console.log(response);
    };
    fetchUser();
  }, []);

  const handleSave = (content: string) => {
    setMarkdownContent(content);
    // ここで保存処理を実装（APIを呼び出すなど）
    console.log("保存処理:", content);
  };

  const handleCancel = () => {
    // キャンセル処理
    console.log("キャンセル");
  };

  return (
    <div className="min-h-screen bg-white">
      <Box sx={{ display: "flex", height: "100vh" }}>
        {/* サイドバー */}
        <SideBar />
        
        {/* メインコンテンツ */}
        <Box 
          sx={{ 
            flexGrow: 1, 
            overflow: "auto",
            p: { xs: 2, sm: 3 },
            backgroundColor: "#f9fafc"
          }}
        >
          <Container maxWidth="lg" sx={{ py: 3 }}>
            <Paper 
              elevation={0} 
              sx={{ 
                p: 4, 
                borderRadius: "12px",
                boxShadow: "0 2px 12px rgba(0,0,0,0.05)"
              }}
            >
              <Typography 
                variant="h4" 
                component="h1" 
                sx={{ 
                  mb: 3, 
                  fontWeight: 600,
                  textAlign: "center",
                  color: "#3f51b5"
                }}
              >
                メモエディタ
              </Typography>
              
              <MarkdownEditor
                initialContent={markdownContent}
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
