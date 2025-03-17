import React from "react";
import { Container, Typography, Paper, Box, Button } from "@mui/material";
import { Link } from "react-router-dom";
import SideBar from "../components/SideBar";

const Home: React.FC = () => {
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
          <Container maxWidth="md" sx={{ py: 3 }}>
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
                ホームページ
              </Typography>
              
              <Typography 
                variant="body1" 
                sx={{ 
                  mb: 4,
                  lineHeight: 1.6
                }}
              >
                ログインに成功しました。左側のサイドバーからフォルダとファイルを管理できます。
                フォルダを選択して新しいメモを作成したり、既存のメモを閲覧したりすることができます。
              </Typography>
              
              <Box sx={{ textAlign: "center" }}>
                <Button 
                  component={Link} 
                  to="/"
                  variant="outlined" 
                  color="primary"
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
                  トップページに戻る
                </Button>
              </Box>
            </Paper>
          </Container>
        </Box>
      </Box>
    </div>
  );
};

export default Home;
