import React from "react";
import { 
  Typography, 
  Container, 
  Button, 
  Box, 
  Paper, 
  Grid,
  Card,
  CardContent,
} from "@mui/material";
import { 
  NoteAlt as NoteIcon, 
  FolderSpecial as FolderIcon, 
  Difference as DiffIcon, 
  PictureAsPdf as PdfIcon,
  Login as LoginIcon,
  PersonAdd as SignupIcon
} from "@mui/icons-material";
import { Link } from "react-router-dom";

const Top: React.FC = () => {
  return (
    <div className="min-h-screen bg-white">
      {/* ヒーローセクション */}
      <Box className="bg-gradient-to-b from-blue-50 to-white py-12 md:py-20">
        <Container maxWidth="lg">
          <Box className="text-center mb-8 md:mb-12 px-4">
            <Typography 
              variant="h3" 
              component="h1" 
              className="font-bold mb-4"
              sx={{ 
                fontSize: { xs: '1.75rem', sm: '2.25rem', md: '3rem' },
                lineHeight: { xs: 1.2, md: 1.3 },
                color: '#1a237e',
                marginBottom: { xs: 3, md: 4 }
              }}
            >
              エンジニアのための<br className="sm:hidden" />スマートメモアプリ
            </Typography>

            <Typography 
              variant="h6" 
              className="mx-auto text-gray-600 text-center"
              sx={{ 
                fontWeight: 400,
                maxWidth: '650px',
                fontSize: { xs: '1rem', md: '1.2rem' },
                lineHeight: 1.6,
                marginBottom: { xs: 5, md: 8 },
                margin: '0 auto',
                textAlign: 'center'
              }}
            >
              コードのように構造化された思考を、直感的に記録・整理・共有できます
            </Typography>

            <Box 
              className="flex justify-center flex-wrap"
              sx={{
                gap: { xs: 2, md: 3 },
                flexDirection: { xs: 'column', sm: 'row' },
                alignItems: 'center',
                '& > *': { 
                  width: { xs: '100%', sm: 'auto' },
                  maxWidth: { xs: '280px', sm: 'none' }
                }
              }}
            >
              <Button
                variant="contained"
                color="primary"
                component={Link}
                to="/signup"
                size="large"
                startIcon={<SignupIcon />}
                sx={{ 
                  borderRadius: '28px',
                  px: { xs: 3, md: 4 },
                  py: 1.5,
                  boxShadow: '0 4px 14px 0 rgba(0,118,255,0.39)',
                  backgroundColor: '#3f51b5',
                  '&:hover': {
                    backgroundColor: '#303f9f'
                  },
                  fontWeight: 600
                }}
              >
                無料で始める
              </Button>

              <Button
                variant="outlined"
                color="primary"
                component={Link}
                to="/login"
                size="large"
                startIcon={<LoginIcon />}
                sx={{ 
                  borderRadius: '28px',
                  px: { xs: 3, md: 4 },
                  py: 1.5,
                  borderColor: '#3f51b5',
                  borderWidth: 2,
                  '&:hover': {
                    borderColor: '#303f9f',
                    borderWidth: 2,
                    backgroundColor: 'rgba(63, 81, 181, 0.04)'
                  },
                  fontWeight: 600
                }}
              >
                ログイン
              </Button>
            </Box>
          </Box>
        </Container>
      </Box>

      {/* 機能紹介セクション */}
      <Container maxWidth="lg" className="py-12">
        <Typography 
          variant="h4" 
          component="h2" 
          className="text-center font-bold mb-12"
          sx={{ color: '#1a237e' }}
        >
          主な機能
        </Typography>

        <Grid container spacing={4}>
          <Grid item xs={12} md={6} lg={3}>
            <Card 
              elevation={0} 
              className="h-full transition-all duration-300 hover:shadow-md"
              sx={{ 
                border: '1px solid #f0f0f0',
                borderRadius: '12px'
              }}
            >
              <CardContent className="p-6 flex flex-col items-center text-center">
                <FolderIcon 
                  sx={{ 
                    fontSize: 56, 
                    color: '#3f51b5', 
                    mb: 2 
                  }} 
                />
                <Typography variant="h6" className="font-bold mb-2">
                  直感的なディレクトリ
                </Typography>
                <Typography variant="body2" color="textSecondary">
                  メモを作成するときにディレクトリの作成を直感的なUIから行うことができます
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={6} lg={3}>
            <Card 
              elevation={0} 
              className="h-full transition-all duration-300 hover:shadow-md"
              sx={{ 
                border: '1px solid #f0f0f0',
                borderRadius: '12px'
              }}
            >
              <CardContent className="p-6 flex flex-col items-center text-center">
                <NoteIcon 
                  sx={{ 
                    fontSize: 56, 
                    color: '#3f51b5', 
                    mb: 2 
                  }} 
                />
                <Typography variant="h6" className="font-bold mb-2">
                  マークダウン対応
                </Typography>
                <Typography variant="body2" color="textSecondary">
                  マークダウンでメモを作成することができ、コードブロックやテーブルも美しく表示します
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={6} lg={3}>
            <Card 
              elevation={0} 
              className="h-full transition-all duration-300 hover:shadow-md"
              sx={{ 
                border: '1px solid #f0f0f0',
                borderRadius: '12px'
              }}
            >
              <CardContent className="p-6 flex flex-col items-center text-center">
                <DiffIcon 
                  sx={{ 
                    fontSize: 56, 
                    color: '#3f51b5', 
                    mb: 2 
                  }} 
                />
                <Typography variant="h6" className="font-bold mb-2">
                  差分表示
                </Typography>
                <Typography variant="body2" color="textSecondary">
                  メモの差分を確認してから保存でき、変更履歴も簡単に閲覧できます
                </Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={6} lg={3}>
            <Card 
              elevation={0} 
              className="h-full transition-all duration-300 hover:shadow-md"
              sx={{ 
                border: '1px solid #f0f0f0',
                borderRadius: '12px'
              }}
            >
              <CardContent className="p-6 flex flex-col items-center text-center">
                <PdfIcon 
                  sx={{ 
                    fontSize: 56, 
                    color: '#3f51b5', 
                    mb: 2 
                  }} 
                />
                <Typography variant="h6" className="font-bold mb-2">
                  エクスポート機能
                </Typography>
                <Typography variant="body2" color="textSecondary">
                  メモをPDFで出力したり、画像などを保存したりすることができます
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Container>

      {/* CTAセクション */}
      <Box className="bg-blue-50 py-16">
        <Container maxWidth="md">
          <Paper 
            elevation={0} 
            className="p-8 text-center"
            sx={{ 
              borderRadius: '16px',
              border: '1px solid #e0e0ff'
            }}
          >
            <Typography variant="h5" className="font-bold mb-4" sx={{ color: '#1a237e' }}>
              今すぐ始めましょう
            </Typography>
            <Typography variant="body1" className="mb-6 text-gray-600">
              無料アカウントを作成して、より効率的なメモ体験を始めましょう
            </Typography>
            <Button
              variant="contained"
              color="primary"
              component={Link}
              to="/signup"
              size="large"
              sx={{ 
                borderRadius: '28px',
                px: 6,
                py: 1.5,
                boxShadow: '0 4px 14px 0 rgba(0,118,255,0.39)',
                backgroundColor: '#3f51b5',
                '&:hover': {
                  backgroundColor: '#303f9f'
                }
              }}
            >
              無料で登録する
            </Button>
          </Paper>
        </Container>
      </Box>

      {/* フッター */}
      <Box className="py-8 border-t border-gray-100">
        <Container maxWidth="lg">
          <Typography variant="body2" className="text-center text-gray-500">
            © 2024 エンジニアメモアプリ
          </Typography>
        </Container>
      </Box>
    </div>
  );
};

export default Top;
