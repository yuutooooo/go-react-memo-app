import React from 'react';
import { 
  Box, 
  Container, 
  Typography, 
  Button, 
  Grid,
  Card,
  CardContent,
} from '@mui/material';
import { Link } from 'react-router-dom';
import { 
  NoteAlt as NoteIcon,
  Security as SecurityIcon,
  DevicesOutlined as DevicesIcon,
  Cake as CakeIcon 
} from '@mui/icons-material';
import { useAuth } from '../contexts/AuthContext';
import LogoutButton from '../components/LogoutButton';

const Landing: React.FC = () => {
  const { isAuthenticated } = useAuth();

  return (
    <Box sx={{ minHeight: '100vh', bgcolor: '#f9fafb' }}>
      {/* ヘッダー */}
      <Box sx={{ 
        bgcolor: '#fff', 
        boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
        py: 2,
        borderBottom: '1px solid #eaecef'
      }}>
        <Container maxWidth="lg">
          <Box sx={{ 
            display: 'flex', 
            justifyContent: 'space-between',
            alignItems: 'center'
          }}>
            <Typography 
              variant="h5" 
              component="h1" 
              sx={{ 
                fontWeight: 'bold',
                color: '#3f51b5'
              }}
            >
              メモアプリ
            </Typography>
            
            <Box sx={{ display: 'flex', gap: 2 }}>
              {isAuthenticated ? (
                <>
                  <Button 
                    component={Link} 
                    to="/home" 
                    variant="outlined" 
                    color="primary"
                  >
                    マイメモ
                  </Button>
                  <LogoutButton variant="contained" color="primary" />
                </>
              ) : (
                <>
                  <Button 
                    component={Link} 
                    to="/login" 
                    variant="outlined" 
                    color="primary"
                  >
                    ログイン
                  </Button>
                  <Button 
                    component={Link} 
                    to="/signup" 
                    variant="contained" 
                    color="primary"
                  >
                    アカウント作成
                  </Button>
                </>
              )}
            </Box>
          </Box>
        </Container>
      </Box>

      {/* ヒーローセクション */}
      <Box sx={{ 
        bgcolor: '#3f51b5',
        color: 'white',
        py: { xs: 6, md: 10 },
        textAlign: 'center'
      }}>
        <Container maxWidth="md">
          <Typography 
            variant="h2" 
            component="h2" 
            sx={{ 
              fontWeight: 'bold',
              mb: 2,
              fontSize: { xs: '2rem', sm: '3rem', md: '3.5rem' }
            }}
          >
            マークダウン対応のメモアプリ
          </Typography>
          
          <Typography 
            variant="h5" 
            sx={{ 
              mb: 4,
              opacity: 0.9,
              maxWidth: '700px',
              mx: 'auto',
              lineHeight: 1.5
            }}
          >
            シンプルで使いやすい、マークダウン記法に対応したメモアプリです。
            アイデアや考えをすばやく整理しましょう。
          </Typography>
          
          {isAuthenticated ? (
            <Button 
              component={Link} 
              to="/home"
              variant="contained" 
              color="secondary" 
              size="large"
              sx={{ 
                py: 1.5, 
                px: 4, 
                fontSize: '1.1rem',
                fontWeight: 'bold',
                borderRadius: '8px',
                boxShadow: '0 4px 14px 0 rgba(0, 0, 0, 0.2)',
              }}
            >
              メモを作成する
            </Button>
          ) : (
            <Button 
              component={Link} 
              to="/signup"
              variant="contained" 
              color="secondary" 
              size="large"
              sx={{ 
                py: 1.5, 
                px: 4, 
                fontSize: '1.1rem',
                fontWeight: 'bold',
                borderRadius: '8px',
                boxShadow: '0 4px 14px 0 rgba(0, 0, 0, 0.2)',
              }}
            >
              無料で始める
            </Button>
          )}
        </Container>
      </Box>

      {/* 特徴セクション */}
      <Box sx={{ py: 8 }}>
        <Container>
          <Typography 
            variant="h4" 
            component="h3" 
            align="center"
            sx={{ 
              mb: 6,
              fontWeight: 'bold'
            }}
          >
            メモアプリの特徴
          </Typography>
          
          <Grid container spacing={4}>
            <Grid item xs={12} sm={6} md={3}>
              <Card sx={{ height: '100%', boxShadow: '0 2px 12px rgba(0,0,0,0.08)' }}>
                <CardContent sx={{ textAlign: 'center', py: 4 }}>
                  <NoteIcon sx={{ fontSize: 60, color: '#3f51b5', mb: 2 }} />
                  <Typography gutterBottom variant="h5" component="div" fontWeight="bold">
                    マークダウン対応
                  </Typography>
                  <Typography variant="body1" color="text.secondary">
                    フォーマット豊かなドキュメントを簡単に作成できます。
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
            
            <Grid item xs={12} sm={6} md={3}>
              <Card sx={{ height: '100%', boxShadow: '0 2px 12px rgba(0,0,0,0.08)' }}>
                <CardContent sx={{ textAlign: 'center', py: 4 }}>
                  <SecurityIcon sx={{ fontSize: 60, color: '#3f51b5', mb: 2 }} />
                  <Typography gutterBottom variant="h5" component="div" fontWeight="bold">
                    セキュア
                  </Typography>
                  <Typography variant="body1" color="text.secondary">
                    HTTP-Onlyクッキーで安全に認証情報を管理します。
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
            
            <Grid item xs={12} sm={6} md={3}>
              <Card sx={{ height: '100%', boxShadow: '0 2px 12px rgba(0,0,0,0.08)' }}>
                <CardContent sx={{ textAlign: 'center', py: 4 }}>
                  <DevicesIcon sx={{ fontSize: 60, color: '#3f51b5', mb: 2 }} />
                  <Typography gutterBottom variant="h5" component="div" fontWeight="bold">
                    レスポンシブ
                  </Typography>
                  <Typography variant="body1" color="text.secondary">
                    どのデバイスからでも快適に利用できます。
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
            
            <Grid item xs={12} sm={6} md={3}>
              <Card sx={{ height: '100%', boxShadow: '0 2px 12px rgba(0,0,0,0.08)' }}>
                <CardContent sx={{ textAlign: 'center', py: 4 }}>
                  <CakeIcon sx={{ fontSize: 60, color: '#3f51b5', mb: 2 }} />
                  <Typography gutterBottom variant="h5" component="div" fontWeight="bold">
                    使いやすさ
                  </Typography>
                  <Typography variant="body1" color="text.secondary">
                    直感的なUIで、すぐに使い始めることができます。
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Container>
      </Box>

      {/* フッター */}
      <Box sx={{ 
        bgcolor: '#f1f5f9',
        py: 4,
        borderTop: '1px solid #e2e8f0'
      }}>
        <Container>
          <Typography 
            variant="body2" 
            align="center"
            color="text.secondary"
          >
            © {new Date().getFullYear()} メモアプリ. All rights reserved.
          </Typography>
        </Container>
      </Box>
    </Box>
  );
};

export default Landing; 